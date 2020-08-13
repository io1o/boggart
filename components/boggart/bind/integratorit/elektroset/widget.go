package elektroset

import (
	"encoding/json"
	"io"
	"sort"
	"time"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/boggart/protocols/http"
	"github.com/kihamo/boggart/providers/integratorit/elektroset"
	"github.com/kihamo/shadow/components/dashboard"
)

func (b *Bind) WidgetHandler(w *dashboard.Response, r *dashboard.Request) {
	vars := map[string]interface{}{}
	ctx := r.Context()
	query := r.URL().Query()
	widget := b.Widget()

	action := query.Get("action")
	vars["action"] = action

	switch action {
	case "bill":
		account := query.Get("account")

		period, err := time.Parse(layoutPeriod, query.Get("period"))
		if err != nil {
			widget.NotFound(w, r)
			return
		}

		var provider elektroset.Provider
		err = json.Unmarshal([]byte(query.Get("provider")), &provider)
		if err != nil {
			widget.NotFound(w, r)
			return
		}

		billLink, err := b.client.BillFile(ctx, account, provider, period)
		if err != nil {
			widget.InternalError(w, r, err)
			return
		}

		response, err := http.NewClient().Get(ctx, billLink)
		if err != nil {
			widget.InternalError(w, r, err)
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename=\"elektroset_bill_"+period.Format("20060102")+".pdf\"")
		w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
		w.Header().Set("Content-Length", response.Header.Get("Content-Length"))

		_, err = io.Copy(w, response.Body)
		if err != nil {
			widget.InternalError(w, r, err)
			return
		}

		response.Body.Close()
		return

	default:
		type row struct {
			Date   time.Time
			Values map[string][3]float64
		}

		rows := make([]*row, 0)
		rowsByTime := make(map[time.Time]*row)
		tariffs := make(map[int64]bool, 3)

		accounts, err := b.client.Accounts(ctx)
		if err == nil {
			dateStart := time.Now().Add(-time.Hour * 24 * 365)
			dateEnd := time.Now()

			for _, account := range accounts {
				for _, service := range account.Services {
					if balances, e := b.client.BalanceDetails(ctx, account.Number, service.Provider, dateStart, dateEnd); e == nil {
						for _, balance := range balances {
							switch balance.KDTariffPlanEntity {
							// показания
							case elektroset.TariffPlanEntityValue:
								r, ok := rowsByTime[balance.DatetimeEntity.Time]
								if !ok && (balance.ValueT1 != nil || balance.ValueT2 != nil || balance.ValueT3 != nil) {
									r = &row{
										Date:   balance.DatetimeEntity.Time,
										Values: make(map[string][3]float64, 3),
									}
									rowsByTime[balance.DatetimeEntity.Time] = r
									rows = append(rows, r)
								}

								if balance.ValueT1 != nil {
									r.Values["tariff1"] = [3]float64{*balance.ValueT1}
									tariffs[1] = true
								}

								if balance.ValueT2 != nil {
									r.Values["tariff2"] = [3]float64{*balance.ValueT2}
									tariffs[2] = true
								}

								if balance.ValueT3 != nil {
									r.Values["tariff3"] = [3]float64{*balance.ValueT3}
									tariffs[3] = true
								}
							}
						}
					} else {
						widget.FlashError(r, "Get balance details failed with error %v", "", err)
					}
				}

				// deltas 1
				for i, row := range rows {
					if i > 0 {
						for tariff, current := range row.Values {
							if prev, ok := rows[i-1].Values[tariff]; ok {
								current[1] = current[0] - prev[0]
								row.Values[tariff] = current
							}
						}
					}
				}

				// trends 2
				for i, row := range rows {
					if i > 0 {
						for tariff, current := range row.Values {
							if prev, ok := rows[i-1].Values[tariff]; ok {
								current[2] = current[1] - prev[1]
								row.Values[tariff] = current
							}
						}
					}
				}
			}
		} else {
			widget.FlashError(r, "Get accounts failed with error %v", "")
		}

		sort.SliceStable(rows, func(i, j int) bool {
			if rows[i].Date.Year() == rows[j].Date.Year() {
				return int(rows[i].Date.Month()) > int(rows[j].Date.Month())
			}

			return rows[i].Date.Year() > rows[j].Date.Year()
		})

		vars["rows"] = rows
		vars["tariffs"] = tariffs
	}

	widget.Render(ctx, "widget", vars)
}

func (b *Bind) WidgetAssetFS() *assetfs.AssetFS {
	return assetFS()
}
