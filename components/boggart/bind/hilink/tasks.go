package hilink

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/kihamo/boggart/components/boggart/tasks"
	"github.com/kihamo/boggart/providers/hilink"
	"github.com/kihamo/boggart/providers/hilink/client/device"
	"github.com/kihamo/boggart/providers/hilink/client/global"
	"github.com/kihamo/boggart/providers/hilink/client/monitoring"
	"github.com/kihamo/boggart/providers/hilink/client/net"
	"github.com/kihamo/boggart/providers/hilink/client/sms"
	"github.com/kihamo/boggart/providers/hilink/static/models"
	"go.uber.org/multierr"
)

func (b *Bind) Tasks() []tasks.Task {
	return []tasks.Task{
		tasks.NewTask().
			WithName("serial-number").
			WithHandler(
				b.Workers().WrapTaskHandlerIsOnline(
					tasks.HandlerFuncFromShortToLong(b.taskSerialNumberHandler),
				),
			).
			WithSchedule(
				tasks.ScheduleWithSuccessLimit(
					tasks.ScheduleWithDuration(tasks.ScheduleNow(), time.Second*30),
					1,
				),
			),
	}
}

func (b *Bind) taskSerialNumberHandler(ctx context.Context) error {
	deviceInfo, err := b.client.Device.GetDeviceInformation(device.NewGetDeviceInformationParamsWithContext(ctx))
	if err != nil {
		return fmt.Errorf("get device information failed: %w", err)
	}

	if deviceInfo.Payload.SerialNumber == "" {
		return errors.New("device returns empty serial number")
	}

	if deviceInfo.Payload.MacAddress1 == "" && deviceInfo.Payload.MacAddress2 == "" {
		return errors.New("device returns empty MAC address")
	}

	if deviceInfo.Payload.MacAddress1 != "" {
		err = b.Meta().SetMACAsString(deviceInfo.Payload.MacAddress1)
	} else if deviceInfo.Payload.MacAddress2 != "" {
		err = b.Meta().SetMACAsString(deviceInfo.Payload.MacAddress2)
	}

	if err != nil {
		return err
	}

	b.Meta().SetSerialNumber(deviceInfo.Payload.SerialNumber)

	// set settings
	settings, err := b.client.Global.GetGlobalModuleSwitch(global.NewGetGlobalModuleSwitchParamsWithContext(ctx))
	if err != nil {
		return fmt.Errorf("get global module switch failed: %w", err)
	}

	cfg := b.config()
	ussdEnabled := settings.Payload.USSDEnabled > 0
	smsEnabled := settings.Payload.SMSEnabled > 0

	if ussdEnabled {
		b.Workers().RegisterTask(tasks.NewTask().
			WithName("balance-updater").
			WithHandler(
				b.Workers().WrapTaskHandlerIsOnline(
					tasks.HandlerWithTimeout(
						tasks.HandlerFuncFromShortToLong(b.taskBalanceUpdaterHandler),
						cfg.BalanceUpdaterTimeout,
					),
				),
			).
			WithSchedule(tasks.ScheduleWithDuration(tasks.ScheduleNow(), cfg.BalanceUpdaterInterval)),
		)

		b.Workers().RegisterTask(tasks.NewTask().
			WithName("limit-traffic-updater").
			WithHandler(
				b.Workers().WrapTaskHandlerIsOnline(
					tasks.HandlerWithTimeout(
						tasks.HandlerFuncFromShortToLong(b.taskLimitTrafficUpdaterHandler),
						cfg.LimitTrafficUpdaterTimeout,
					),
				),
			).
			WithSchedule(tasks.ScheduleWithDuration(tasks.ScheduleNow(), cfg.LimitTrafficUpdaterInterval)),
		)
	}

	if smsEnabled {
		b.Workers().RegisterTask(tasks.NewTask().
			WithName("sms-checker").
			WithHandler(
				b.Workers().WrapTaskHandlerIsOnline(
					tasks.HandlerWithTimeout(
						tasks.HandlerFuncFromShortToLong(b.taskSMSCheckerHandler),
						cfg.SMSCheckerTimeout,
					),
				),
			).
			WithSchedule(tasks.ScheduleWithDuration(tasks.ScheduleNow(), cfg.SMSCheckerInterval)),
		)

		b.Workers().RegisterTask(tasks.NewTask().
			WithName("cleaner").
			WithHandler(
				b.Workers().WrapTaskHandlerIsOnline(
					tasks.HandlerFuncFromShortToLong(b.taskCleanerHandler),
				),
			).
			WithSchedule(tasks.ScheduleWithDuration(tasks.ScheduleNow(), cfg.CleanerInterval)),
		)
	}

	b.Workers().RegisterTask(tasks.NewTask().
		WithName("system-updater").
		WithHandler(
			b.Workers().WrapTaskHandlerIsOnline(
				tasks.HandlerWithTimeout(
					tasks.HandlerFuncFromShortToLong(b.taskSystemUpdaterHandler),
					cfg.SystemUpdaterTimeout,
				),
			),
		).
		WithSchedule(tasks.ScheduleWithDuration(tasks.ScheduleNow(), cfg.SystemUpdaterInterval)),
	)

	return nil
}

func (b *Bind) taskBalanceUpdaterHandler(ctx context.Context) error {
	if b.simStatus.Load() != 1 {
		return nil
	}

	balance, err := b.Balance(ctx)

	if err == nil {
		sn := b.Meta().SerialNumber()

		metricBalance.With("serial_number", sn).Set(balance)

		if e := b.MQTT().PublishAsync(ctx, b.config().TopicBalance.Format(sn), balance); e != nil {
			err = multierr.Append(err, e)
		}
	}

	return err
}

func (b *Bind) taskLimitTrafficUpdaterHandler(ctx context.Context) error {
	if b.simStatus.Load() != 1 || b.operator.IsEmpty() {
		return nil
	}

	op, err := b.operatorSettings()
	if err != nil {
		return err
	}

	if op.LimitTrafficUSSD == "" {
		return nil
	}

	_, err = b.USSD(ctx, op.LimitTrafficUSSD)

	return err
}

func (b *Bind) taskSMSCheckerHandler(ctx context.Context) error {
	if b.simStatus.Load() != 1 {
		return nil
	}

	// sms counters
	paramsCount := sms.NewGetSMSCountParamsWithContext(ctx)
	responseCount, err := b.client.Sms.GetSMSCount(paramsCount)

	if err != nil {
		return fmt.Errorf("get sms count failed: %w", err)
	}

	sn := b.Meta().SerialNumber()
	cfg := b.config()

	metricSMSUnread.With("serial_number", sn).Set(float64(responseCount.Payload.LocalUnread))
	metricSMSInbox.With("serial_number", sn).Set(float64(responseCount.Payload.LocalInbox))

	if e := b.MQTT().PublishAsync(ctx, cfg.TopicSMSUnread.Format(sn), responseCount.Payload.LocalUnread); e != nil {
		err = multierr.Append(err, e)
	}

	if e := b.MQTT().PublishAsync(ctx, cfg.TopicSMSInbox.Format(sn), responseCount.Payload.LocalInbox); e != nil {
		err = multierr.Append(err, e)
	}

	// ----- read new sms -----
	paramsList := sms.NewGetSMSListParamsWithContext(ctx).
		WithRequest(&models.SMSListRequest{
			PageIndex: 1,
			ReadCount: 20,
			BoxType:   1,
		})

	responseList, e := b.client.Sms.GetSMSList(paramsList)
	if e != nil {
		return multierr.Append(err, fmt.Errorf("get sms list failed: %w", err))
	}

	for _, s := range responseList.Payload.Messages {
		if s.Status != 1 {
			payload, e := json.Marshal(s)

			if err != nil {
				err = multierr.Append(err, e)
				continue
			}

			isSpecial := b.checkSpecialSMS(ctx, s)
			if !isSpecial {
				if e = b.MQTT().PublishAsync(ctx, cfg.TopicSMS.Format(sn), payload); e != nil {
					err = multierr.Append(err, e)
					continue
				}
			}

			if isSpecial && cfg.CleanerSpecial {
				params := sms.NewRemoveSMSParamsWithContext(ctx)
				params.Request.Index = s.Index

				_, e = b.client.Sms.RemoveSMS(params)
			} else {
				params := sms.NewReadSMSParamsWithContext(ctx)
				params.Request.Index = s.Index

				_, e = b.client.Sms.ReadSMS(params)
			}

			if e != nil {
				err = multierr.Append(err, e)
			}
		}
	}

	// ----- delete old sms -----

	return err
}

func (b *Bind) taskSystemUpdaterHandler(ctx context.Context) (err error) {
	sn := b.Meta().SerialNumber()
	cfg := b.config()

	// status
	if response, e := b.client.Monitoring.GetMonitoringStatus(monitoring.NewGetMonitoringStatusParamsWithContext(ctx)); e == nil {
		b.simStatus.Set(uint32(response.Payload.SimStatus))

		// только с активной SIM картой
		if response.Payload.SimStatus == 1 {
			if e := b.MQTT().PublishAsync(ctx, cfg.TopicSignalLevel.Format(sn), response.Payload.SignalIcon); e != nil {
				err = multierr.Append(err, e)
			}

			metricSignalLevel.With("serial_number", sn).Set(float64(response.Payload.SignalIcon))

			if b.operator.IsEmpty() {
				plmn, e := b.client.Net.GetCurrentPLMN(net.NewGetCurrentPLMNParamsWithContext(ctx))
				if e == nil && plmn.Payload.FullName != "" {
					b.operator.Set(plmn.Payload.FullName)

					if e := b.MQTT().PublishAsync(ctx, cfg.TopicOperator.Format(sn), plmn.Payload.FullName); e != nil {
						err = multierr.Append(err, e)
					}
				} else {
					err = multierr.Append(err, e)
				}
			}
		}
	} else {
		err = multierr.Append(err, fmt.Errorf("get monitoring status failed: %w", e))
	}

	// все проверки ниже только с активной SIM картой
	if b.simStatus.Load() != 1 {
		return err
	}

	// signal
	if response, e := b.client.Device.GetDeviceSignal(device.NewGetDeviceSignalParamsWithContext(ctx)); e == nil {
		const dBmPostfix = "dBm"

		if raw := strings.TrimRight(response.Payload.RSSI, dBmPostfix); raw != "" {
			if val, e := parseInt(raw); e == nil {
				metricSignalRSSI.With("serial_number", sn).Set(float64(val))

				if e = b.MQTT().PublishAsync(ctx, cfg.TopicSignalRSSI.Format(sn), val); e != nil {
					err = multierr.Append(err, e)
				}
			} else {
				err = multierr.Append(err, fmt.Errorf("parse RSSI value '%s' failed: %w", raw, e))
			}
		}

		if raw := strings.TrimRight(response.Payload.RSRP, dBmPostfix); raw != "" {
			if val, e := parseInt(raw); e == nil {
				metricSignalRSRP.With("serial_number", sn).Set(float64(val))

				if e = b.MQTT().PublishAsync(ctx, cfg.TopicSignalRSRP.Format(sn), val); e != nil {
					err = multierr.Append(err, e)
				}
			} else {
				err = multierr.Append(err, fmt.Errorf("parse RSRP value '%s' failed: %w", raw, e))
			}
		}

		if raw := strings.TrimRight(response.Payload.RSRQ, dBmPostfix); raw != "" {
			if val, e := parseInt(raw); e == nil {
				metricSignalRSRQ.With("serial_number", sn).Set(float64(val))

				if e = b.MQTT().PublishAsync(ctx, cfg.TopicSignalRSRQ.Format(sn), val); e != nil {
					err = multierr.Append(err, e)
				}
			} else {
				err = multierr.Append(err, fmt.Errorf("parse RSRQ value '%s' failed: %w", raw, e))
			}
		}

		if raw := strings.TrimRight(response.Payload.SINR, dBmPostfix); raw != "" {
			if val, e := parseInt(raw); e == nil {
				metricSignalSINR.With("serial_number", sn).Set(float64(val))

				if e = b.MQTT().PublishAsync(ctx, cfg.TopicSignalSINR.Format(sn), val); e != nil {
					err = multierr.Append(err, e)
				}
			} else {
				err = multierr.Append(err, fmt.Errorf("parse SINR value '%s' failed: %w", raw, e))
			}
		}
	} else {
		err = multierr.Append(err, fmt.Errorf("get device signal failed: %w", e))
	}

	// traffic
	if response, e := b.client.Monitoring.GetMonitoringTrafficStatistics(monitoring.NewGetMonitoringTrafficStatisticsParamsWithContext(ctx)); e == nil {
		metricTotalConnectTime.With("serial_number", sn).Set(float64(response.Payload.TotalConnectTime))
		metricTotalDownload.With("serial_number", sn).Set(float64(response.Payload.TotalDownload))
		metricTotalUpload.With("serial_number", sn).Set(float64(response.Payload.TotalUpload))

		if e := b.MQTT().PublishAsync(ctx, cfg.TopicConnectionTime.Format(sn), response.Payload.TotalConnectTime); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, cfg.TopicConnectionDownload.Format(sn), response.Payload.TotalDownload); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTT().PublishAsync(ctx, cfg.TopicConnectionUpload.Format(sn), response.Payload.TotalUpload); e != nil {
			err = multierr.Append(err, e)
		}
	} else {
		err = multierr.Append(err, fmt.Errorf("get monitoring traffic statistics failed: %w", e))
	}

	return err
}

func (b *Bind) taskCleanerHandler(ctx context.Context) error {
	if b.simStatus.Load() != 1 {
		return nil
	}

	var page int64 = 1
	cfg := b.config()

	for {
		paramsList := sms.NewGetSMSListParamsWithContext(ctx).
			WithRequest(&models.SMSListRequest{
				PageIndex: page,
				ReadCount: 20,
				BoxType:   1,
			})

		response, err := b.client.Sms.GetSMSList(paramsList)
		if err != nil {
			return fmt.Errorf("get sms list failed: %w", err)
		}

		if len(response.Payload.Messages) == 0 {
			return nil
		}

		for _, s := range response.Payload.Messages {
			remove := cfg.CleanerSpecial && b.checkSpecialSMS(ctx, s)
			if !remove {
				d, e := time.Parse(hilink.TimeFormat, s.Date)

				if e != nil {
					continue
				}

				remove = time.Since(d) > cfg.CleanerDuration
			}

			if remove {
				params := sms.NewRemoveSMSParamsWithContext(ctx)
				params.Request.Index = s.Index

				if _, err := b.client.Sms.RemoveSMS(params); err != nil {
					return fmt.Errorf("call remove sms failed: %w", err)
				}

				b.Logger().Warn("Clean sms",
					"date", s.Date,
					"content", s.Content,
					"phone", s.Phone,
				)

				time.Sleep(time.Second)
			}
		}

		page++
	}
}
