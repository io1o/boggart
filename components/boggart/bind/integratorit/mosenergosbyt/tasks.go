package mosenergosbyt

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/kihamo/boggart/components/boggart/tasks"
	"go.uber.org/multierr"
)

func (b *Bind) Tasks() []tasks.Task {
	return []tasks.Task{
		tasks.NewTask().
			WithName("updater").
			WithHandler(
				b.Workers().WrapTaskHandlerIsOnline(
					tasks.HandlerFuncFromShortToLong(b.taskUpdaterHandler),
				),
			).
			WithSchedule(tasks.ScheduleWithDuration(tasks.ScheduleNow(), b.config().UpdaterInterval)),
	}
}

func (b *Bind) taskUpdaterHandler(ctx context.Context) error {
	account, err := b.Account(ctx)
	if err != nil {
		return err
	}

	cfg := b.config()

	if balance, e := b.client.CurrentBalance(ctx, account.Provider.IDAbonent); e != nil {
		err = multierr.Append(e, err)
	} else {
		metricBalance.With("account", account.NNAccount).Set(balance.Balance)

		if e := b.MQTT().PublishAsync(ctx, cfg.TopicBalance.Format(account.NNAccount), balance.Balance); e != nil {
			err = multierr.Append(e, err)
		}

		var serviceID string

		for i, service := range balance.Services {
			if id, ok := services[strings.ToLower(service.Service)]; ok {
				serviceID = id
			} else {
				serviceID = strconv.FormatInt(int64(i), 10)
			}

			metricServiceBalance.With("account", account.NNAccount, "service", serviceID).Set(service.Balance)

			if e := b.MQTT().PublishAsync(ctx, cfg.TopicServiceBalance.Format(account.NNAccount, serviceID), service.Balance); e != nil {
				err = multierr.Append(e, err)
			}
		}
	}

	// last bill
	if details, e := b.client.ChargeDetail(ctx, account.Provider.IDAbonent, time.Now().Add(-cfg.BalanceDetailsInterval), time.Now()); e != nil {
		err = multierr.Append(e, err)
	} else {
		lastBillTime := time.Time{}
		lastBillUUID := ""

		for _, detail := range details {
			if lastBillTime.IsZero() || detail.Period.After(lastBillTime) {
				lastBillTime = detail.Period.Time
				lastBillUUID = detail.Services[0].ReportUUID
			}
		}

		if !lastBillTime.IsZero() {
			billLink, e := b.Widget().URL(map[string]string{
				"action": "bill",
				"period": lastBillTime.Format(layoutPeriod),
				"uuid":   lastBillUUID,
			})

			if e == nil {
				if e := b.MQTT().PublishAsync(ctx, cfg.TopicLastBill.Format(account.NNAccount), billLink); e != nil {
					err = multierr.Append(err, e)
				}
			}
		}
	}

	return err
}
