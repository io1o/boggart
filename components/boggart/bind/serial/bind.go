package serial

import (
	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/protocols/serial_tcp"
)

type Bind struct {
	boggart.BindBase

	server *serial_tcp.Server
}

func (b *Bind) Run() error {
	go func() {
		b.UpdateStatus(boggart.BindStatusOnline)

		err := b.server.ListenAndServe()
		if err != nil {
			b.Logger().Error("Failed serve with error " + err.Error())

			b.UpdateStatus(boggart.BindStatusOffline)
		}
	}()

	return nil
}

func (b *Bind) Close() (err error) {
	return b.server.Close()
}
