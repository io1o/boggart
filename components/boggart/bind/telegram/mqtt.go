package telegram

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/mqtt"
	"github.com/kihamo/boggart/components/storage"
)

const (
	MQTTPrefix mqtt.Topic = boggart.ComponentName + "/telegram/+/"

	MQTTSubscribeTopicSendMessage  = MQTTPrefix + "send/+/message"
	MQTTSubscribeTopicSendFile     = MQTTPrefix + "send/+/file"
	MQTTSubscribeTopicSendFileURL  = MQTTPrefix + "send/+/file/url"
	MQTTPublishTopicReceiveMessage = MQTTPrefix + "receive/+/message"
	MQTTPublishTopicReceiveAudio   = MQTTPrefix + "receive/+/audio"
	MQTTPublishTopicReceiveVoice   = MQTTPrefix + "receive/+/voice"
)

func (b *Bind) MQTTPublishes() []mqtt.Topic {
	sn := mqtt.NameReplace(b.SerialNumber())

	return []mqtt.Topic{
		mqtt.Topic(MQTTPublishTopicReceiveMessage.Format(sn)),
		mqtt.Topic(MQTTPublishTopicReceiveAudio.Format(sn)),
		mqtt.Topic(MQTTPublishTopicReceiveVoice.Format(sn)),
	}
}

func (b *Bind) MQTTSubscribers() []mqtt.Subscriber {
	sn := mqtt.NameReplace(b.SerialNumber())

	return []mqtt.Subscriber{
		mqtt.NewSubscriber(MQTTSubscribeTopicSendMessage.Format(sn), 0, func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 4) {
				return nil
			}

			routes := mqtt.RouteSplit(message.Topic())
			if len(routes) < 1 {
				return errors.New("bad topic name")
			}

			return b.SendMessage(routes[len(routes)-2], message.String())
		}),
		mqtt.NewSubscriber(MQTTSubscribeTopicSendFile.Format(sn), 0, func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 4) {
				return nil
			}

			routes := mqtt.RouteSplit(message.Topic())
			if len(routes) < 1 {
				return errors.New("bad topic name")
			}

			var (
				mime storage.MIMEType
				name string
				url  string

				payload FilePayload
			)

			if err := message.UnmarshalJSON(&payload); err == nil {
				mime = storage.MIMEType(payload.MIME)
				name = payload.Name
				url = payload.URL
			} else {
				url = message.String()
			}

			if url == "" {
				return errors.New("url fields is empty")
			}

			request, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				return err
			}

			response, err := http.DefaultClient.Do(request)
			if err != nil {
				return err
			}
			defer response.Body.Close()

			if mime == storage.MIMETypeUnknown {
				mime, err = storage.MimeTypeFromHTTPHeader(response.Header)
				if err != nil {
					copyBody := &bytes.Buffer{}
					if _, err := io.CopyN(copyBody, response.Body, 128); err != nil {
						return err
					}

					mime, err = storage.MimeTypeFromData(bytes.NewBuffer(copyBody.Bytes()))
					if err != nil {
						return err
					}

					// довычитываем все остальное, так как body уже порвался на две части до 128 и послке
					if _, err := io.Copy(copyBody, response.Body); err != nil {
						return err
					}

					// присваиваем обратно, чтобы с этим можно было проджолжать работать
					response.Body = ioutil.NopCloser(copyBody)
					defer copyBody.Reset()
				}
			}

			to := routes[len(routes)-2]

			if name == "" {
				name = "File at " + time.Now().Format(time.RFC1123Z)
			}

			switch mime {
			case storage.MIMETypeJPEG:
				err = b.SendPhoto(to, name, response.Body, response.ContentLength)

			case storage.MIMETypeMPEG, storage.MIMETypeWAVE, storage.MIMETypeOGG:
				err = b.SendAudio(to, name, response.Body, response.ContentLength)

			default:
				err = b.SendDocument(to, name, response.Body, response.ContentLength)
			}

			return err
		}),
		mqtt.NewSubscriber(MQTTSubscribeTopicSendFileURL.Format(sn), 0, func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 5) {
				return nil
			}

			routes := mqtt.RouteSplit(message.Topic())
			if len(routes) < 1 {
				return errors.New("bad topic name")
			}

			return b.SendFileAsURL(routes[len(routes)-3], "File at "+time.Now().Format(time.RFC1123Z), message.String())
		}),
	}
}
