package webos

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow/components/dashboard"
)

func (b *Bind) WidgetHandler(w *dashboard.Response, r *dashboard.Request) {
	data := make(map[string]interface{})
	widget := b.Widget()

	if !b.Meta().IsStatusOnline() {
		r.Session().FlashBag().Error(widget.Translate(r.Context(), "Device is offline", ""))
	}

	if r.IsPost() {
		toast := r.Original().FormValue("toast")
		data["toast"] = toast

		if toast != "" {
			if err := b.Toast(toast); err != nil {
				r.Session().FlashBag().Error(err.Error())
			} else {
				r.Session().FlashBag().Success(widget.Translate(r.Context(), "Message send success", ""))
				widget.Redirect(r.URL().Path, http.StatusFound, w, r)
				return
			}
		} else {
			data["error"] = widget.Translate(r.Context(), "Message is empty", "")
		}
	}

	widget.Render(r.Context(), "widget", data)
}

func (b *Bind) WidgetAssetFS() *assetfs.AssetFS {
	return assetFS()
}
