package mobile

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/kihamo/boggart/protocols/http"
	"github.com/kihamo/shadow/components/tracing"
	tracingHttp "github.com/kihamo/shadow/components/tracing/http"
	"github.com/opentracing/opentracing-go/log"
)

const (
	MegafonLkComponentName = "megafon"

	MegafonLkURL           = "https://lk.megafon.ru"
	MegafonLkLoginFormURL  = MegafonLkURL + "/login/"
	MegafonLkBalanceURL    = MegafonLkURL + "/api/lk/balance/get"
	MegafonLkRemaindersURL = MegafonLkURL + "/api/options/remaindersMini"
)

var reCSRF = regexp.MustCompile(`"CSRF":"([^"]+)"`)

type Megafon struct {
	connection *http.Client
	phone      string
	password   string
}

type MegafonRemainders struct {
	Voice                float64
	Sms                  float64
	Internet             float64
	InternetProlongation float64
}

func NewMegafon(phone, password string) *Megafon {
	return &Megafon{
		connection: http.NewClient(),
		phone:      phone,
		password:   password,
	}
}

func (m *Megafon) Number() string {
	return m.phone
}

func (m *Megafon) auth(ctx context.Context) (string, error) {
	m.connection.Reset()

	ctx = tracingHttp.ComponentNameToContext(ctx, MegafonLkComponentName)
	ctx = tracingHttp.OperationNameToContext(ctx, MegafonLkComponentName+".auth.get")

	response, err := m.connection.Get(ctx, MegafonLkLoginFormURL)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return "", err
	}

	selectorFormAction := doc.Find("form[class~=form-login-autofill]")
	if selectorFormAction.Length() == 0 {
		return "", errors.New("selector of login form returns empty result")
	}

	action, ok := selectorFormAction.First().Attr("action")
	if !ok || action == "" {
		return "", errors.New("action for login form not found")
	}

	selectorCSRF := selectorFormAction.Find("input[name=CSRF]")
	if selectorCSRF.Length() == 0 {
		return "", errors.New("selector of CSRF token returns empty result")
	}

	token, ok := selectorCSRF.First().Attr("value")
	if !ok || token == "" {
		return "", errors.New("CSRF token not found")
	}

	ctx = tracingHttp.OperationNameToContext(ctx, MegafonLkComponentName+".auth.post")

	response, err = m.connection.Post(ctx, MegafonLkURL+action, map[string]string{
		"j_username": m.phone,
		"j_password": m.password,
		"CSRF":       token,
	})
	if err != nil {
		return "", err
	}

	submatch := reCSRF.FindStringSubmatch(http.BodyFromResponse(response))
	if len(submatch) != 2 {
		return "", errors.New("CSRF token not found in page")
	}

	return submatch[1], nil
}

func (m *Megafon) Balance(ctx context.Context) (float64, error) {
	ctx = tracingHttp.ComponentNameToContext(ctx, MegafonLkComponentName)

	span, ctx := tracing.StartSpanFromContext(ctx, MegafonLkComponentName, "balance")
	defer span.Finish()

	csrf, err := m.auth(ctx)
	if err != nil {
		tracing.SpanError(span, err)
		return -1, err
	}

	response, err := m.connection.GetAjax(ctx, fmt.Sprintf("%s?CSRF=%s&_=%d", MegafonLkBalanceURL, csrf, time.Now().Unix()))
	if err != nil {
		tracing.SpanError(span, err)
		return -1, err
	}

	decoder := json.NewDecoder(response.Body)
	reply := struct {
		Balance float64 `json:"balance"`
	}{}

	if err := decoder.Decode(&reply); err != nil {
		tracing.SpanError(span, err)
		return -1, err
	}

	span.LogFields(log.Float64("balance", reply.Balance))

	return reply.Balance, nil
}

func (m *Megafon) Remainders(ctx context.Context) (*MegafonRemainders, error) {
	ctx = tracingHttp.ComponentNameToContext(ctx, MegafonLkComponentName)

	span, ctx := tracing.StartSpanFromContext(ctx, MegafonLkComponentName, "remainders")
	defer span.Finish()

	csrf, err := m.auth(ctx)
	if err != nil {
		tracing.SpanError(span, err)
		return nil, err
	}

	response, err := m.connection.GetAjax(ctx, fmt.Sprintf("%s?CSRF=%s&_=%d", MegafonLkRemaindersURL, csrf, time.Now().Unix()))
	if err != nil {
		tracing.SpanError(span, err)
		return nil, err
	}

	decoder := json.NewDecoder(response.Body)
	reply := struct {
		Models []struct {
			Remainders []struct {
				Name      string `json:"name"`
				Available struct {
					Value float64 `json:"value"`
				} `json:"availableValue"`
				Total struct {
					Value float64 `json:"value"`
				} `json:"totalValue"`
			} `json:"remainders"`
		} `json:"models"`
	}{}

	if err := decoder.Decode(&reply); err != nil {
		tracing.SpanError(span, err)
		return nil, err
	}

	ret := &MegafonRemainders{}

	for _, model := range reply.Models {
		for _, remainder := range model.Remainders {
			switch remainder.Name {
			case "Минуты по тарифному плану":
				ret.Voice = remainder.Total.Value - remainder.Available.Value

			case "SMS по тарифному плану":
				ret.Sms = remainder.Total.Value - remainder.Available.Value

			case "Интернет по тарифному плану":
				ret.Internet = remainder.Total.Value - remainder.Available.Value

			case "Автопродление":
				ret.InternetProlongation = remainder.Total.Value - remainder.Available.Value
			}
		}
	}

	span.LogFields(
		log.Float64("voice", ret.Voice),
		log.Float64("sms", ret.Sms),
		log.Float64("internet", ret.Internet),
		log.Float64("internet_prolongation", ret.InternetProlongation),
	)

	return ret, nil
}
