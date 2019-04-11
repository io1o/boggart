// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package pulsar

import (
	"github.com/elazarl/go-bindata-assetfs"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x9a\x6d\x6f\xe2\x46\x10\xc7\xdf\xe7\x53\x8c\xac\x9c\x2e\xa8\xc5\x36\x24\xd7\x9e\x38\x43\x55\xe5\xae\xa7\x4a\x6d\x15\x5d\xd3\xbe\xe8\x9b\x68\xf1\x2e\x61\xab\x65\xed\xdb\x1d\x43\x22\xcb\xdf\xa9\x9f\xa1\x9f\xac\x5a\xdb\x18\x87\xf0\x60\x83\xe9\x1d\x04\x29\xc1\xec\xc3\x7f\x66\x76\x7e\xcc\x98\x84\x38\x06\xca\x46\x5c\x32\xb0\xfc\x40\x22\x93\x68\x41\x92\x9c\x79\xda\x57\x3c\x44\xc0\xc7\x90\xf5\x2d\x12\x86\x82\xfb\x04\x79\x20\x9d\xbf\xc9\x94\x64\x93\xd6\xe0\x0c\x00\x60\x14\x49\xdf\xcc\xc0\x94\x88\x88\xfd\x14\xa8\x09\xc1\x0b\xd9\x82\x38\x9d\x35\x0f\xc5\x30\x52\x12\x24\x9b\xc1\xcf\x12\x85\xfd\x5b\x34\x19\x32\x95\xaf\x6c\xd9\xa3\xf9\x96\x77\xe9\x8e\xe4\xcc\x73\x32\x03\x83\x33\x8f\xf2\x29\xf8\x82\x68\xdd\xb7\x54\x30\xcb\x2d\x96\x47\xfd\x40\xb4\x27\xb4\xdd\xe9\x82\xb9\xd2\x93\xf9\xd5\x83\x6e\x77\xba\xf9\xfa\xe5\x3d\x0f\x77\x21\x91\x4c\x94\x66\x9f\xaf\x98\x1f\xc6\xd3\x35\xc5\x3a\x15\x08\xd6\xb7\x90\x0c\x57\x29\x15\x2b\x23\x31\x17\x94\x64\x0a\x92\x4c\xdb\x48\x86\x1a\x86\x44\xdd\x99\x0b\x6b\x21\x23\xb8\x5e\x65\xab\x50\x12\x3c\x5f\x1b\x2a\xa6\x99\xc4\x34\x17\x56\x1c\x03\x1f\x01\xfb\x0c\x36\xc9\x52\x60\x99\xe4\xcd\x6d\x9a\xb1\x29\x33\x8b\x98\xa4\x90\x24\x03\x8f\xc0\x58\xb1\x51\x7f\xcd\xbe\x45\x66\x7b\xd3\x80\xd3\x0b\xb7\xf5\xce\xec\x15\x9a\x41\x92\xc4\x31\xd8\x9f\xd8\xe7\x88\x69\xb4\xff\xf8\xf4\x8b\x7d\x43\x70\x9c\x0d\x67\xe2\xd6\xc0\x88\x76\xde\x4a\xb0\xae\x23\xa5\x98\x44\x98\x30\xa2\x23\xc5\x26\x4c\xa2\xb6\xc0\x86\x24\xf1\x1c\x32\xf0\x1c\xc1\x1b\x08\x94\x28\x7f\x6c\xa2\xdb\x31\xde\xd2\xf6\x5d\xc2\xfe\x21\xd3\xe9\xe7\x32\xab\x4e\xe1\xc7\xb9\x85\xed\x81\x7b\x4e\x24\xd6\xcc\x94\x98\x44\x32\x6c\xaf\xa7\x72\x13\x9d\x65\x05\x33\x02\x23\x42\x19\x64\xc7\x05\x5c\x6e\x50\x33\x8f\xf5\x87\x97\x24\x1b\x37\x7a\x86\x6b\x56\x32\x2e\x18\xa4\xbf\xdb\x1a\x15\x0f\x19\x05\x4a\x90\x64\xe3\x14\xdb\x8a\xe9\x30\x90\xda\xf8\x24\x83\x99\x22\xa1\x05\x1a\x1f\x4d\x20\x33\x4e\x71\xdc\xeb\xb8\xee\xab\x2d\xae\x66\x56\xc7\x8c\xd0\x2a\xeb\xd4\xf6\x45\xb9\xe0\x22\xab\xef\x09\x16\x29\xc5\xf1\x0e\x02\x1f\x24\x53\xf7\x8f\xc0\x25\x7c\xf4\x89\xa8\x23\xe5\x39\x55\x3c\x36\x5a\x15\xe3\x1f\x06\xf4\x71\xfb\xba\x38\x06\x45\xe4\x3d\x83\x73\xfe\x2d\x9c\x6b\x24\x08\xbd\x3e\xd8\xe6\x42\xc3\x16\x02\xa0\xde\x39\xd3\x41\xb5\xb6\x43\x03\x3f\x32\x55\xc5\x9e\x29\x8e\xec\x82\x12\x64\xb7\xc1\xef\xa8\xb8\xbc\xbf\x78\x1d\xc7\x99\x97\xb6\x49\x95\x9d\x35\x19\xb0\xba\xae\xfb\x5d\xdb\xed\xb4\xdd\xee\x6d\xe7\x4d\xcf\xbd\xea\xb9\x6f\xfe\x72\xbf\xef\xb9\xae\xa9\x00\xaf\x5b\xad\xa2\xe5\x78\x0e\x56\x38\xbc\xbd\xfc\x2d\xf7\xc9\xc2\xdb\x9c\x8b\x24\xa9\xed\x4b\x35\x2e\x8a\x12\x55\x45\x6f\x3b\x19\x9e\x93\xbe\x6f\xb7\x56\x8e\xbc\x8c\x1e\x71\x9d\x80\x3a\x0c\xc3\xf2\xdb\xfd\xd7\x45\x0f\xac\x5b\x36\x9e\x69\xfd\x69\xb8\xa9\x5d\x7c\xbe\x4c\xd5\x80\xfa\xa7\x46\x9f\x56\x59\x20\x92\x02\xf2\xc9\x22\xe2\x8a\xc9\x82\xa2\x67\xd9\xa6\x32\x18\x09\xfb\x83\x52\x81\xaa\x02\xff\x92\x4b\x05\x94\xec\x01\xdb\xd4\x94\x41\x95\xb6\xf8\x25\xe5\x42\xbf\xb6\x97\x55\xde\x1f\x2b\xfc\xaa\x6e\xa3\xd8\xd4\x48\x65\x5d\xc4\x9d\xc2\x58\xbf\xbc\xd6\xf6\x3b\xcb\xa4\x64\x25\xd3\xef\x99\x40\x02\x6e\xdd\x63\xcb\x4e\x01\x55\x20\xef\x57\xa6\xb5\xb6\xd8\xc2\xbd\x7b\x6c\xc6\xbd\xb2\x6a\xe7\xad\xbc\x11\x91\x22\x02\xac\x6f\x5e\x51\xd0\xcc\x0f\x24\xb5\xca\x2f\xcc\x4d\xf5\x92\xd5\x73\x90\x5c\x14\x3f\xcb\xb3\x3b\x7a\xb4\x23\xa6\x1b\x02\x2a\xc7\xf3\x45\xc2\xa9\xd6\x09\x97\x1f\x9e\x93\xe1\xb3\x13\xc5\x3b\xd8\xac\x5f\x4d\xaa\xdb\xa8\xd6\x17\x60\xaf\x32\x7e\xcb\x26\x21\x53\x04\x23\x95\x7e\xcc\xd8\xbd\x8e\xe3\x42\xe8\x8e\xcb\xe6\xab\xf9\x2a\xfd\xff\xbf\xa6\x37\x71\x33\xb9\x1c\x4a\x5a\xa6\x9f\xde\x55\x42\x91\x9f\x7f\xff\xb9\xb6\x76\x0a\xf0\xeb\xc5\x2c\x88\xb0\x21\xce\x82\x08\x0f\x0b\x5a\x61\xe0\x04\x48\x33\xb1\xbc\x34\xd4\xa8\x69\x42\x0d\xc1\x96\x6a\x1d\x16\xb7\x92\x89\x13\x00\x2e\x8b\xe6\x25\x20\x77\x13\xcc\x98\xda\x03\xb3\xd0\xec\x6f\x1e\xad\x92\xec\x91\xe2\x94\x45\xb0\x11\xa1\x8f\x3e\x11\xce\xf8\x14\x28\xca\xfe\xbc\xb4\x07\x46\x2c\x15\x68\x9e\xa3\xb2\xee\x91\x82\x94\x87\xb0\x95\xa4\x53\xe0\xe8\x9a\x84\xc4\xe7\xb8\x0f\x49\x7e\x2e\xd1\x3c\x4b\x4f\x95\x8f\x94\xa6\x22\x88\x35\x3c\x29\x32\x83\x8b\x2c\x19\x13\x4f\x47\xe1\xe0\xd2\x73\xcc\x93\xd5\x3a\x05\xbc\x02\xa9\xa3\x49\x98\xfe\xcf\x6f\x0f\xc2\x16\x2a\x07\x80\x6c\x59\xfc\x58\x39\x2b\xc5\x51\x13\x35\x67\x7c\x12\xb0\xdd\x44\xe6\xfc\xb9\x0c\x23\x84\xce\x3e\x77\x58\x91\x16\xe6\xc3\x76\x18\xe1\x5d\xe7\x00\x77\x5a\xcf\xe5\x8f\x14\xb9\xa7\x91\x6c\xec\x97\xa1\xc9\x8d\x3e\x85\x8e\x59\xa6\xac\xdb\x10\x65\xdd\xc3\x52\xd6\x3d\x19\xca\xba\x2f\x91\xb2\xcb\x86\x28\xbb\x3c\x2c\x65\x97\x27\x43\xd9\xe5\x4b\xa4\xec\xaa\x21\xca\xae\x0e\x4b\xd9\xd5\xc9\x50\x76\x75\x3c\x94\x35\xfb\x15\x92\xcd\xde\x79\x0e\xe5\xd3\x75\xdf\xab\x5b\x39\xb5\x62\x78\x69\xa8\xf4\x32\xbf\xcc\x9f\x0a\x77\xfe\x0b\x00\x00\xff\xff\xbf\xce\xb7\xd9\x37\x2b\x00\x00"

func templatesViewsWidgetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsWidgetHtml,
		"templates/views/widget.html",
	)
}

func templatesViewsWidgetHtml() (*asset, error) {
	bytes, err := templatesViewsWidgetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/widget.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x53\x4d\x68\x5c\x55\x14\xfe\x3a\x93\x99\xd4\xb1\x5a\x8d\x7f\xad\x08\xde\x22\x69\xad\xe1\x25\xf3\x93\x6a\x7d\xc9\xa4\xd5\x98\x8a\xda\xd0\x58\xc6\x2e\x95\xcb\xcc\xed\xcc\xd3\x99\xf7\x86\xf7\x53\x0d\x54\xc8\x0f\x21\x34\x11\x15\x54\x14\x11\x15\xc1\x8d\xa0\x4e\x63\x06\xc7\x36\x33\x01\xc1\x85\xe0\xe2\x5c\x70\xe1\x46\x57\x2e\x05\x05\x41\x5c\xca\x7d\xf7\x99\xf4\x25\x94\xdc\xc5\x9c\xf3\x7d\xe7\xdc\xf3\xf3\xdd\x37\xbf\x0d\xf4\xbd\x07\x00\xf7\x03\x78\x00\xc0\x1f\x00\x06\x01\x6c\xec\x43\x78\x66\x12\xc0\xdd\x00\x9e\x4f\x00\xf7\x02\xb8\x98\x00\xfa\x01\xbc\x9e\x00\xf6\x03\x58\x4a\x00\xb7\x02\x78\x33\xca\xfb\x30\x01\xf4\x01\xf8\x3a\x01\xdc\x06\x60\x3d\x01\xa4\x01\x7c\x9f\x00\x6e\x07\xf0\x63\x14\xff\x25\xe2\x7f\x8d\xee\xff\x9e\x00\x52\x00\xfe\x8a\xee\xfd\x13\xd9\xbe\xa4\xb6\x03\x91\x65\x49\xe0\x4e\x00\xc3\x49\x5d\xef\x74\x12\x38\x08\xe0\x5c\x52\xdf\x7f\x29\xca\x13\x11\xef\x25\x75\x9f\xc5\x24\x50\xdb\x07\xac\x24\x81\x02\x80\xbf\xfb\x80\x47\x01\x1c\x4b\x01\x19\x00\x2f\x46\xd6\x4a\x01\x07\x00\xf8\x29\xe0\x21\x00\xf3\x29\xbd\xe7\xe7\x29\xbd\xff\x57\x29\xdd\xf7\x87\x94\xd6\xec\xe7\x28\xfe\x67\x4a\xef\xf1\x6f\x4a\xeb\xd0\x97\x06\xee\x00\x70\x38\x0d\x3c\x08\xe0\x91\xc8\x3e\x1b\xd9\x97\x23\xbb\x92\x06\x8e\x00\xf8\x22\xad\xfb\xfd\x94\x06\x8e\x29\x1d\xa2\xfb\x87\xfa\x75\xbf\xe3\xfd\x7a\xef\x93\xfd\xc0\x80\xaa\xd3\x0f\xa8\x27\x4a\xe9\x67\xc2\x61\x6c\x9f\xfd\x88\x9f\x7b\xa0\x67\x3f\x04\x5d\x23\x09\x5d\x43\xcd\x7b\xd7\x0d\x79\xea\x5d\x0f\x46\xbe\x7a\xa3\xfb\x22\x5f\xed\x79\x20\xf2\x13\xd0\x7a\x2a\xad\xd4\x9e\x6a\xc6\x5b\xa0\x35\x57\x73\x62\xb0\xc2\x3c\x51\x76\xec\xca\xb6\xe7\x61\x68\x9b\x1d\xba\x81\x7e\xc2\x2d\xd7\xac\x4b\x02\x93\xbc\xc9\xcb\x96\x3f\x8b\x49\xc7\xf6\x82\x46\xd3\xb7\x1c\x1b\x93\x81\xeb\x0a\xdb\x67\x0d\xc1\xbd\xc0\x15\x0d\x61\xfb\x1e\x9e\xe2\xbe\x08\x7f\x18\xb7\x2b\xcc\xb7\x1a\x02\x53\xb6\x70\xab\xb3\x91\x61\x96\xcd\x9e\x2e\xf3\x3a\xb6\x7e\x46\x6a\x98\xde\xae\x80\x19\xe7\x55\xe1\x62\x26\xa8\x7b\x82\x59\x76\x33\xf0\x59\x2e\x86\xf2\x31\x54\x88\xa1\x51\x94\x44\xa3\x29\x5c\xee\x07\xae\x60\x15\x51\xf7\x79\x8c\xb1\xec\x18\x74\x02\x1f\x17\x78\x3d\x10\x68\x8c\x7b\x41\x73\xa2\x30\x3e\xa2\x4c\x1c\x8d\xd4\xd0\x54\x2d\x3c\x9c\x17\x4d\xc7\xf5\x8d\x69\xaf\x6a\x55\x8c\x27\x83\xaa\x67\x94\x1c\x93\x55\xc4\xa5\xd3\xaf\x58\x35\xde\x70\x86\xdd\x20\x33\x73\xae\x64\x4c\xba\x82\x2b\x81\x0c\x25\x83\xc9\xf2\xd9\xdc\xe3\x46\xb6\x60\xe4\x1f\x63\xf9\x82\x79\xe2\xc4\x50\xb6\x90\xcd\x66\xce\x72\xcf\x37\x4a\x2e\xb7\xbd\x3a\xf7\x1d\xd7\x64\xcf\x85\x35\xd8\x74\xe0\xf2\x86\x53\x71\xd8\x78\xac\xf0\x44\xe6\x2c\xb7\xab\x01\xaf\x0a\xa3\x24\x78\xc3\x64\x5b\xd8\x64\xe7\x03\xcf\xb3\xb8\x9d\x99\x7e\x66\x7a\xca\xb8\x20\x5c\xcf\x72\x6c\x93\xe5\x86\xb3\x99\x49\xc7\xf6\x85\xed\x1b\xa5\xd9\xa6\x30\x99\x2f\x5e\xf3\x47\x9a\x75\x6e\xd9\x63\xac\x5c\xe3\xae\x27\xfc\xe2\x0b\xa5\x33\xc6\xc9\xed\x3c\x35\xcf\x45\xe1\x1a\x53\x76\xd9\xa9\x58\x76\xd5\x64\x99\x99\x7a\xe0\xf2\xba\x71\xc6\x71\x1b\x9e\xc9\xec\x66\x08\xbd\x62\x61\x8c\x69\xb7\x68\x0f\xe6\xb2\xc5\x62\x8e\x1d\x3d\xca\x94\x9b\x3d\x52\xcc\xe5\xd8\x29\x96\x65\x66\x88\x27\x8a\xf9\xff\x43\xe3\xc5\x51\xe5\x3e\x1c\xa6\x8d\xe7\xb2\xec\xf2\x65\x7d\x65\xa2\x98\xcf\x1e\x67\xa7\x58\x8e\x99\x2c\x3f\xa6\xbe\x4b\x39\x4f\x6d\xba\x26\x17\xa9\x4b\xeb\xd4\xda\xc9\xc8\xd5\x9d\x4c\xf8\xd5\xee\xb8\xb4\x93\x92\xab\xbb\x28\xd0\xdb\x72\x4e\x2e\x51\x87\xd6\x40\x1f\xd3\x55\xb9\x42\x6d\xda\x00\x7d\x4a\x2d\x39\x2f\x97\xa8\xa7\x52\x3e\xd3\xf9\xf2\x0a\x75\xa8\xcd\x68\x93\x7a\x74\x8d\x5a\xf4\x2d\xb5\xe4\x02\xb5\xe9\x3a\x75\x40\xef\x86\xa0\xb5\xe5\x30\xea\x30\x5a\x93\x73\xaa\x9c\x7c\x0b\xf4\x25\x75\xa9\x2d\xe7\xe8\x1b\xea\xec\x82\x8c\xd6\x18\xbd\x13\xd6\xbc\x8e\x5d\xce\x88\x5c\x06\x7d\xb2\xb3\xa7\x7c\x03\xf4\x11\xf5\xe4\x15\xea\x52\x4f\xce\xcb\x05\x45\x7c\x40\x1b\xb4\x29\x17\x55\x54\xce\x53\x57\xae\xd2\x77\x6a\x88\x70\x0d\x96\xdb\x23\x9e\xdf\x23\x5e\xd8\x23\x3e\xaa\xa6\x6c\x87\x1b\x6f\x52\x8b\xd6\x59\x38\xe8\x06\x6d\x86\x64\x4b\x2e\xc8\x45\x2d\x65\x8c\x52\x56\x4b\xba\x4e\x2d\xb9\xac\xa4\xbc\x49\x4e\x8f\xae\x6a\x82\xba\xd4\x52\x1a\xbe\x1f\x3a\xcb\xd4\xa6\xae\x7a\x18\xd0\x46\xec\x4f\x1c\x87\xa1\x8a\x9d\xd8\xfc\x3d\x5a\xc3\x7f\x01\x00\x00\xff\xff\xab\x3e\xa5\xca\x5d\x07\x00\x00"

func localesRuLc_messagesWidgetMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesWidgetMo,
		"locales/ru/LC_MESSAGES/widget.mo",
	)
}

func localesRuLc_messagesWidgetMo() (*asset, error) {
	bytes, err := localesRuLc_messagesWidgetMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/widget.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/views/widget.html": templatesViewsWidgetHtml,
	"locales/ru/LC_MESSAGES/widget.mo": localesRuLc_messagesWidgetMo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"locales": &bintree{nil, map[string]*bintree{
		"ru": &bintree{nil, map[string]*bintree{
			"LC_MESSAGES": &bintree{nil, map[string]*bintree{
				"widget.mo": &bintree{localesRuLc_messagesWidgetMo, map[string]*bintree{}},
			}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"widget.html": &bintree{templatesViewsWidgetHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
