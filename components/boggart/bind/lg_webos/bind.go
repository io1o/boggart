package lg_webos

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kihamo/boggart/components/boggart"
	"github.com/snabb/webostv"
)

var defaultDialerLGWebOS = webostv.Dialer{
	DisableTLS: true,
	WebsocketDialer: &websocket.Dialer{
		Proxy: http.ProxyFromEnvironment,
		NetDial: (&net.Dialer{
			Timeout:   time.Second * 5,
			KeepAlive: time.Second * 30, // ensure we notice if the TV goes away
		}).Dial,
	},
}

type Bind struct {
	boggart.DeviceBindBase
	boggart.DeviceBindMQTT

	mutex    sync.RWMutex
	initOnce sync.Once

	client *webostv.Tv
	config *Config
}

func (b *Bind) Client() (*webostv.Tv, error) {
	b.mutex.RLock()
	c := b.client
	b.mutex.RUnlock()

	if c != nil {
		return c, nil
	}

	client, err := defaultDialerLGWebOS.Dial(b.config.Host)
	if err != nil {
		return nil, err
	}

	go client.MessageHandler()

	b.mutex.Lock()
	b.client = client
	b.mutex.Unlock()

	return client, nil
}

func (b *Bind) UpdateStatus(status boggart.DeviceStatus) {
	if status == boggart.DeviceStatusOffline && status != b.Status() {
		b.mutex.Lock()
		b.client = nil
		b.mutex.Unlock()
	}

	b.DeviceBindBase.UpdateStatus(status)
}
