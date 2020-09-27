// Code generated by go-bindata.
// sources:
// templates/views/cache.html
// templates/views/subscriptions.html
// locales/ru/LC_MESSAGES/cache.mo
// locales/ru/LC_MESSAGES/mqtt.mo
// locales/ru/LC_MESSAGES/subscriptions.mo
// DO NOT EDIT!

package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/elazarl/go-bindata-assetfs"
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

var _templatesViewsCacheHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x4d\x8f\xdb\x36\x13\xbe\xef\xaf\x18\x30\x1b\x58\x06\xa2\x8f\x4d\xde\xbc\x2d\x36\xb2\x8b\xa2\x45\xd0\x43\x0a\x14\xed\x16\x3d\x14\xc5\x82\x12\xc7\x16\xb7\x14\xa9\x92\xa3\x75\x5c\x75\xff\x7b\x41\xca\xb2\xe5\xad\xe5\x6c\x8a\xa4\x3a\x51\xc3\x99\x67\x3e\x9f\x91\xba\x0e\x04\xae\xa4\x46\x60\xa5\xd1\x84\x9a\x18\x3c\x3c\x5c\xe4\x42\xde\x43\xa9\xb8\x73\x0b\xd6\xf0\x35\xc6\x24\x49\x21\x5b\x5e\x00\x00\x8c\x2f\x83\xfc\x56\xe1\x8a\x76\x97\x41\xa1\x7a\xb5\xec\x3a\x90\x57\x5f\x6a\x60\xdf\xf0\xb2\x42\x06\x09\x3c\x3c\xe4\x69\xf5\x6a\x07\x91\x0a\x79\xbf\xbc\x98\x80\xb3\x72\x5d\x1d\xe1\xad\x8c\xad\x07\x15\x7f\x8e\x2b\x63\xe5\x9f\x46\x13\x57\x10\xde\x15\x2f\x50\xc5\x21\x0c\xb0\x46\x61\xaf\xc6\xa0\x46\xaa\x8c\x58\xb0\xc6\x38\x62\x20\xc5\x82\x39\xd4\x62\x04\xfd\x38\x80\x80\xb6\xb6\xa6\x6d\xa0\x69\x95\x8a\x1f\x87\x72\xca\x46\xea\xa6\xa5\xde\xe8\x84\x66\xd0\x0e\xf1\x0d\xfa\x05\xe9\xb3\xda\xc1\xa2\x68\x89\x8c\x06\xda\x36\xb8\x60\xae\x2d\x6a\x39\x4e\x60\x04\x05\x1e\x4e\x96\xa6\x3f\x08\xae\xd7\x68\xc3\xd1\xd5\x6c\x12\x7e\x78\x04\x27\x1e\x93\x59\xaf\x7d\xc9\x6a\x23\xb8\x62\x3b\x19\xb7\x6b\xa4\x05\x7b\xd6\x0b\x9f\x06\x14\x74\xfb\x51\x59\xb0\xc3\x00\x18\xbd\x92\xb6\x86\x95\x6a\x5d\x05\xe5\x61\x1c\x3e\x0a\xb5\x30\x62\x3b\x02\xfd\x45\x2a\x05\x05\x8e\x41\x13\xf8\xda\x22\x6c\x4d\x0b\xae\xb5\xf8\xd5\xbf\xf0\x51\x72\xa5\x0a\x5e\xfe\xbe\x60\xbc\x24\x69\x74\x34\x0b\xf0\xb3\xf9\x9b\x33\x9d\xf2\x4f\x2e\xf7\xf3\xc3\x1d\xac\x78\x8c\x96\x3b\xb4\x0c\x1e\xd7\xe2\xad\xc7\xdb\x45\xb6\xcc\x53\x79\x66\x00\xd2\x7e\x02\x26\x06\x2a\x0d\x13\x75\x62\x2e\x7b\x62\x9d\x11\xe5\xa9\x9f\xf1\x23\x1a\x0e\x6c\x1c\x0f\x75\xa9\x90\xdb\x95\x7c\xef\xc3\xfc\xe7\xad\x35\x9b\x7e\x1a\x05\xde\xcb\x12\xdd\x89\xcd\xf0\xfe\xb6\xe1\x1a\xd5\x98\xc6\x47\xb7\xc3\xba\x99\xe6\x22\xf1\x42\x61\x6c\xd1\x35\x46\x3b\x79\x8f\xa7\x68\x18\x74\x8e\x0c\xa0\x37\x73\x64\x65\x83\x22\xb4\xb7\x97\x0b\x1a\x61\x81\x36\x1b\xcb\x9b\x3e\x09\x49\x58\x3b\x06\x8e\xb6\xbe\x59\x1b\x29\xa8\xba\xbe\xca\xb2\xe7\x53\x6c\xa6\x0a\xb9\x98\xba\xb3\x67\x7a\x4a\xd5\x61\x2f\xde\x98\x46\x96\xc3\x5e\xa4\xea\x89\x56\xdf\x72\x42\x92\x35\x7e\xb4\xe1\x0f\x7c\xab\x0c\x17\x1f\xb2\xcb\xd3\xa9\x04\xbc\xcd\x99\xb4\x3d\x3f\x4f\xdf\x75\x1d\x58\xbf\x95\xe0\x52\xbe\x80\x4b\x5f\x6b\xb8\x5e\x40\x12\x38\x7b\x1b\x4a\xef\xbf\x38\xa7\x51\xcf\x16\x53\xf8\xec\x02\x60\x12\x8a\x99\xfc\x44\x56\xea\x75\x9f\xdf\x44\xa0\x83\x65\xee\x4a\x2b\x1b\xda\xed\x57\xde\x34\x4a\x96\xdc\x73\x3e\xbd\xe3\xf7\xbc\xbf\x64\x4b\x61\xca\xb6\x46\x4d\xc9\xc6\x4a\xc2\x48\x70\xc2\x1b\xd3\x7b\x89\x66\x7b\xe7\x43\x4f\x92\xb7\xc6\xd6\x9c\x80\xbd\xcc\xb2\xff\xc7\xd9\x55\x9c\xbd\xbc\xb9\x7a\x7d\x9d\xfd\xef\x3a\x7b\x1d\x67\x5f\x5c\x67\x99\xff\xb6\xce\xe6\xf3\x3c\xed\x1d\x2c\x9f\x10\x67\x63\x71\x99\x97\x46\xa0\x4f\xb6\xb1\x52\xd3\x0a\xd8\x73\xc7\x76\xbe\x77\x6d\xfd\xcb\xf2\x4d\xc8\x3b\x68\xe6\x69\xb0\x9a\x04\x9f\x6e\x72\xd7\x01\x6a\x31\xd9\x8f\x74\xa2\xcd\x79\x1a\x28\xf6\x81\xbd\xb3\x7f\x3d\x5e\x3b\x07\x9f\x17\xa3\xbf\x11\x3f\x6b\x6c\x08\xa4\xeb\xc0\x11\x27\x59\x7e\x77\xf3\xfd\x3b\x88\xfa\xf3\xcf\x3f\xbe\x03\x96\x0a\xee\xaa\xc2\x70\x2b\x52\xee\x1c\x92\x4b\xef\x51\x0b\x63\x5d\xba\x27\xbe\x4b\x34\x52\x5c\xb8\xb4\x74\xbd\xf4\xa6\x97\x16\xc6\x90\x23\xcb\x9b\xa4\x96\x3a\x29\x9d\x63\xb0\xe2\xca\xe1\xfc\x13\x7a\x3d\x2c\x9c\x21\x80\x83\xe4\x7c\x00\xa7\xab\x72\xe7\x3e\x61\x4d\xd2\x3b\x97\xde\xfd\xd1\xa2\xdd\x26\xa3\xb2\xf8\x58\xee\x3e\x47\x2d\x0a\xe7\x1d\x4e\x36\xe0\xb3\xf8\x3c\x54\xfb\x91\xef\x51\x1b\xfe\x03\xe7\xbb\xdc\x27\x7b\x7f\xec\xbe\xe7\xc8\x93\x16\xd4\x9e\x5d\x97\xd1\xb0\xab\xe6\x89\x45\x2e\xb6\xd1\xaa\xd5\xe1\x27\x06\xa2\x39\x74\x47\xc4\xbc\x8c\x66\xcf\xc2\xde\x9d\xcd\xfd\xee\xea\x4b\x12\xcd\x13\x63\x05\xda\xe8\xd7\xab\x17\x30\x13\xe8\xca\xd9\x6f\xf3\x44\x58\xbe\x89\xe6\x6f\xf6\xe6\x0f\xbb\xf3\x7e\x83\x8d\xe6\xf4\xef\x00\x00\x00\xff\xff\x56\xf7\x58\xde\x4b\x0c\x00\x00"

func templatesViewsCacheHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsCacheHtml,
		"templates/views/cache.html",
	)
}

func templatesViewsCacheHtml() (*asset, error) {
	bytes, err := templatesViewsCacheHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/cache.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesViewsSubscriptionsHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\xc1\x6e\xda\x40\x10\xbd\xf3\x15\xa3\x6d\x2a\xb5\x52\xf1\x86\x9c\xaa\xc8\x70\xe9\xa5\x07\xaa\x28\x85\x9e\xab\xf5\xee\x50\x2f\x75\x76\xdd\x9d\x0d\x04\x59\xfc\x7b\x65\x1b\xc2\x42\xbc\x8e\x44\x49\x4f\x98\x99\x79\xf3\xde\xcc\xb3\xa7\xaa\x40\xe1\x42\x1b\x04\x26\xad\xf1\x68\x3c\x83\xed\x76\x30\x48\x95\x5e\x81\x2c\x04\xd1\x98\x39\xbb\x66\xa0\xd5\x98\x29\x5c\x69\x89\xc4\x26\x03\x00\x80\xb0\xe4\xe9\x67\x29\x0c\x16\xbb\xcc\xcb\xac\xd7\xbe\xc0\x20\xdb\x54\xe4\x37\x93\xaa\x02\x3d\xfa\x6c\x80\xcd\x1e\x33\x92\x4e\x97\x5e\x5b\x43\x0c\x12\xd8\x6e\x53\x9e\xdf\x9c\x20\x82\x9e\xb2\x40\xe1\x16\xfa\x89\x4d\x52\xae\xf4\x2a\x20\x3e\xf9\x7b\xa4\x63\x3f\x63\xbc\xaf\x17\x59\x81\x43\x87\x54\x5a\x43\x7a\x75\x2a\xba\x29\x6f\x6a\x8e\x00\xd0\xc2\xc8\x3b\x5d\xa2\x02\x25\xbc\x68\xe3\xca\x07\xbd\xc0\xd8\xb5\x13\x25\x03\xf2\x9b\x02\xc7\x6c\xad\x95\xcf\x6f\x47\xd7\xd7\xef\x3b\x58\x5a\xa6\x1c\x85\x8a\xe5\x5c\x77\x62\x07\xdc\xeb\x7b\x50\x43\x69\x8b\xe1\x88\x4d\xde\xa5\xdc\xe7\xbd\x98\x83\x1f\x73\x5b\x6a\xb9\xf7\xe1\x15\xd4\x4b\xa6\xe7\x36\xf7\x77\xb3\x7f\x6f\xf2\x45\x14\x45\x26\xe4\x6f\xba\x4c\xab\x57\xdb\xa4\x3c\xb6\xda\x1a\xd3\x63\x48\x66\xd5\xa6\x3b\x57\x55\xe0\x84\xf9\x85\x70\xa5\x3f\xc1\x15\x05\x2f\x3b\xdc\x8e\x21\x09\x03\x54\x7f\x7f\xdd\xfd\x7b\x0d\x57\xf5\x94\x42\x29\xb8\xd2\x30\x6a\xe7\x8b\x08\x0d\xea\x8f\xa4\x24\x8d\xe9\xe7\x41\xef\xef\x66\xe7\x01\xa7\x68\xce\x03\x36\x56\xf6\x42\xe3\x3e\x56\x15\xa0\x51\xd1\x45\xf3\x88\x93\x29\x6f\x3e\xea\x93\xe3\x11\x3d\x3f\xbb\xc7\xdd\xcf\x81\x73\x10\x1c\xdd\xfa\x75\x62\x7b\x21\x55\x05\xe4\x85\xd7\xf2\xeb\xfc\xdb\x14\x3e\xb4\xcf\x3f\xbe\x4f\x81\x71\x25\x28\xcf\xac\x70\x8a\x0b\x22\xf4\xc4\x57\x68\x94\x75\xc4\x9f\x4f\x0d\x25\x06\xfd\x30\x23\x2e\xa9\x8d\xce\xdb\x68\x66\xad\x27\xef\x44\x99\x3c\x68\x93\x48\x22\x06\x0b\x51\x10\x7e\xbc\x20\xeb\xe1\xc4\xed\x05\x1c\x22\xfd\x02\xba\xb7\xb2\xa4\x0b\xee\x84\x2f\x89\x2f\xff\x3c\xa2\xdb\x24\xc1\x5a\x6a\x2d\xcb\xb7\xd8\x45\x46\x35\x61\xd4\x80\x37\xe1\x3c\x6c\xfb\x84\x3b\xb0\xe1\x3f\x90\xef\x66\x8f\x7a\xbf\xec\xb6\xfe\x6f\x00\x00\x00\xff\xff\x93\xe6\x0e\xbf\x84\x08\x00\x00"

func templatesViewsSubscriptionsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsSubscriptionsHtml,
		"templates/views/subscriptions.html",
	)
}

func templatesViewsSubscriptionsHtml() (*asset, error) {
	bytes, err := templatesViewsSubscriptionsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/subscriptions.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesCacheMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x4f\x6b\x14\x41\x10\xc5\xdf\xfa\x17\xe6\xe8\xd9\x43\xe5\x60\xd0\x43\xc7\x99\x35\x07\xe9\xdd\xde\x15\x43\x02\xe2\x2e\x84\x30\x7a\xb6\xd8\x69\x67\x07\x67\xbb\x97\xee\x1e\x31\x90\x83\xe0\xc5\x8b\x78\xf2\xe2\xc5\x8f\x60\x40\x34\xa0\xe8\x67\xe8\xf9\x02\x7e\x16\x99\xd9\x68\xf0\x5d\xea\xfd\x78\xaf\x8a\xfa\x7d\xe3\xca\x07\x00\xb8\x0c\xe0\x26\x80\x5d\x00\x57\x01\xcc\xb0\xd1\xb3\x73\x66\x00\xd7\x01\x94\x00\xa6\x03\xc0\x02\xb8\x06\xe0\xf3\x00\xd8\x02\xf0\x6d\x00\x0c\x70\xa1\x4b\xe7\x37\xb1\xc7\x8b\xa5\xc6\x21\x1f\xd7\x96\x0b\x1c\xe9\xb5\x75\x41\xcc\x7d\x59\x15\xe2\x61\x53\x7a\x91\x5b\x49\x85\x7e\xf9\xe0\x45\xb5\xe4\x95\xdd\x71\x4d\x32\x63\x1f\x44\xee\xd8\xf8\x9a\x83\x75\x92\x1e\xf7\x11\xcd\x1b\xc7\x2b\x5b\x58\x1a\xff\xd7\x9f\x24\x33\x36\x65\xc3\xa5\x16\xb9\xe6\x95\xa4\x7f\x2c\xe9\xa8\xf1\xbe\x62\x93\xcc\x1f\xcd\xf7\xc5\x53\xed\x7c\x65\x8d\xa4\x6c\x27\x4d\xf6\xac\x09\xda\x04\x91\x1f\xaf\xb5\xa4\xa0\x5f\x85\xbb\xeb\x9a\x2b\x33\xa2\xc5\x92\x9d\xd7\x41\x3d\xc9\x0f\xc4\xfd\x8b\x5e\xf7\xcf\x73\xed\xc4\xbe\x59\xd8\xa2\x32\xa5\xa4\xe4\xb0\x6e\x1c\xd7\xe2\xc0\xba\x95\x97\x64\xd6\x3d\x7a\x75\x6f\x44\x1b\xab\xcc\xad\x2c\x55\x2a\xa3\xed\x6d\xea\x6c\xba\xa5\xb2\x8c\xa6\x94\x92\xec\x79\xa2\x86\x7f\xa3\xb1\xda\xed\xec\xed\xbe\x36\xce\x52\x3a\x39\xd9\xac\x4c\xd4\x30\xbd\x43\x53\xca\x48\xd2\x70\x84\xf8\xb1\x7d\xd7\xbe\x45\xfc\x14\x7f\xc5\x1f\xf1\x6b\x3c\x8b\x3f\xe3\x69\xfb\x9e\xba\x11\xbf\xb4\xaf\xdb\x37\xf1\x2c\x7e\x8f\xa7\xf8\x13\x00\x00\xff\xff\xdb\xec\x6e\x96\xd8\x01\x00\x00"

func localesRuLc_messagesCacheMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesCacheMo,
		"locales/ru/LC_MESSAGES/cache.mo",
	)
}

func localesRuLc_messagesCacheMo() (*asset, error) {
	bytes, err := localesRuLc_messagesCacheMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/cache.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesMqttMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\x31\x8f\xd3\x30\x1c\xc5\x5f\x20\x20\x91\x05\x89\x99\xe1\xcf\xc0\x09\x06\x1f\x49\x61\x40\x6e\xdc\x22\x4e\x57\x09\xd1\x48\x50\x05\x26\x16\x93\xb8\x69\x44\x62\xa7\xb6\x83\x40\xea\xc2\xc6\x00\x23\x0c\x7c\x2a\xba\xb0\x20\x31\xf0\x49\x50\x93\x02\x3a\x4f\xbf\x9f\xde\x7b\x96\xe5\x5f\x37\xc2\x2f\x00\x10\x02\xb8\x09\x20\x05\x70\x05\xc0\x2b\x8c\xa7\x03\x70\x1d\xc0\x16\x40\x04\xe0\x03\x80\x6b\x00\x3e\x01\x98\x07\xc0\xd7\xe3\xf6\x47\x00\x5c\x05\xf0\x33\x18\xfd\x77\x00\x04\xc7\xec\x32\x80\x4b\xc7\xfb\x50\x18\xbd\xae\x2b\xe6\xe5\xeb\xb0\xdd\x7a\x8f\x56\xe9\x3e\x3c\x93\xc5\x46\x8d\x98\x3d\xcf\x73\xac\x54\x67\xac\x67\x99\xab\xea\x92\x3d\xee\x2b\xc7\x72\xc3\xa9\x54\x6f\x1f\xbd\xa9\x37\xb2\x35\xa7\xb6\x8f\x96\xd2\x79\x96\x5b\xa9\x5d\x23\xbd\xb1\x9c\x9e\x0e\x11\x65\xbd\x95\xad\x29\x0d\xa5\x17\xfa\xb3\x68\x29\x75\xd5\xcb\x4a\xb1\x5c\xc9\x96\xd3\x3f\xe7\xb4\xea\x9d\xab\xa5\x8e\xb2\x27\xd9\x39\x7b\xa9\xac\xab\x8d\xe6\x94\x9c\xc6\xd1\x99\xd1\x5e\x69\xcf\xf2\xf7\x9d\xe2\xe4\xd5\x3b\x7f\xaf\x6b\x64\xad\xa7\x54\x6c\xa4\x75\xca\x8b\x17\xf9\x82\x3d\xfc\xdf\x3b\xbc\x67\xad\x2c\x3b\xd7\x85\x29\x6b\x5d\x71\x8a\x9e\x35\xbd\x95\x0d\x5b\x18\xdb\x3a\x4e\xba\x1b\xd4\x89\xfb\x53\x1a\x51\xe8\xdb\x49\x2c\x44\x42\x27\x27\x74\xc0\xf8\x96\x48\x12\x9a\x53\x4c\x7c\xf0\x99\x98\xfc\x8d\x52\xf1\xe0\x80\x77\x86\x5a\x9a\xc4\xb4\xdb\x8d\x93\x99\x98\xc4\x77\x69\x4e\x09\x71\x9a\x4c\x31\x7c\xe2\xf7\x6f\xfb\xcf\xfb\x8f\x23\xff\x09\x00\x00\xff\xff\xf9\x0d\x64\x37\xe7\x01\x00\x00"

func localesRuLc_messagesMqttMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesMqttMo,
		"locales/ru/LC_MESSAGES/mqtt.mo",
	)
}

func localesRuLc_messagesMqttMo() (*asset, error) {
	bytes, err := localesRuLc_messagesMqttMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/mqtt.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesSubscriptionsMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\x4f\x6b\x13\x41\x18\xc6\x9f\xae\xad\xe0\x8a\x82\xc5\xa3\x87\x57\xd0\xa2\x87\xad\xbb\xa9\xa0\x6c\xb2\x89\x58\x5a\x10\x13\xd4\x76\xf5\x3e\x4d\xc6\xed\xd2\xcd\xcc\x32\xb3\x2b\x0a\xb9\x78\xf2\xe0\x4d\xf0\xe2\x45\xf0\xa4\x47\xff\x61\x4c\x31\x7e\x00\x2f\x6f\xbe\x80\x9f\x45\x76\x93\x5a\xfa\x9c\x7e\xcf\xfb\x3e\x33\xf3\x30\x7f\x57\x97\xdf\x02\xc0\x69\x00\x97\x00\x74\x01\x9c\x05\x30\xc2\x5c\x1f\x01\x9c\x01\xf0\x09\xc0\x0a\x80\xaf\x00\x4e\x01\xf8\x09\xe0\x1c\x80\x5f\x8b\xf9\x1f\x00\x9d\x25\x60\x06\x60\x15\xc0\x05\x07\x38\x0f\xe0\x8a\x33\xcf\x6f\x38\xc0\x45\x00\xb7\x1c\xc0\xad\xde\x71\x80\xa5\xc5\x5d\x95\x9c\x45\x87\x4a\xcb\x38\xd6\xca\x11\x6c\x8a\x2c\xdb\x13\xfd\x03\x5b\x93\xc5\xa3\x07\xbb\xd8\x2d\xf7\x6c\xdf\xa4\x79\x91\x6a\x65\x11\xeb\x3c\xed\x63\x47\xe6\xda\x14\x5e\xcf\x26\xe9\xc0\xbb\x5b\x26\xd6\x8b\x75\x48\x03\xf9\xec\xce\x41\xba\x2f\x86\x7a\xdd\x94\x6e\x57\xd8\xc2\x8b\x8d\x50\x36\x13\x85\x36\x21\xdd\xaf\x57\xd4\x2b\x8d\x18\xea\x81\xa6\xd6\x89\x7c\xdb\xed\x0a\x95\x94\x22\x91\x5e\x2c\xc5\x30\xa4\xff\x3e\xa4\x9d\xd2\xda\x54\x28\xb7\x77\xaf\xb7\xe5\x3d\x91\xc6\xa6\x5a\x85\x14\xac\xfb\xee\xa6\x56\x85\x54\x85\x17\xbf\xc8\x65\x48\x85\x7c\x5e\xdc\xc8\x33\x91\xaa\x26\xf5\xf7\x85\xb1\xb2\x88\x1e\xc7\xdb\xde\xed\xe3\x5c\xd5\xe7\xa9\x34\xde\x96\xea\xeb\x41\xaa\x92\x90\xdc\x87\x59\x69\x44\xe6\x6d\x6b\x33\xb4\x21\xa9\xbc\xb6\x36\xda\x68\xd2\x1c\x23\x75\x35\xf0\xa3\x28\xa0\xb5\x35\xaa\xd0\xbf\x1c\x05\x01\x75\xc8\xa7\xb0\xf6\xed\xa8\x71\xb4\x6a\x45\x37\x2b\xbc\x56\xc7\x5a\x81\x4f\xa3\xd1\xfc\x48\x3b\x6a\xf8\xd7\xa9\x43\x01\x85\xd4\x68\x82\xdf\xf1\x94\x0f\xf9\x90\x3f\xf3\x77\x9e\xf0\x94\xbf\x80\xdf\xcc\x5e\xf3\x8f\x0a\x6b\x5b\xfd\x3c\xbf\xe7\x29\x7f\xe3\xdf\x3c\x9e\xbd\x9c\xbd\xe2\x31\x4f\x78\x0c\xfe\xc0\xd3\x6a\xc4\x13\xfc\x0b\x00\x00\xff\xff\xe1\x49\x78\xbe\x57\x02\x00\x00"

func localesRuLc_messagesSubscriptionsMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesSubscriptionsMo,
		"locales/ru/LC_MESSAGES/subscriptions.mo",
	)
}

func localesRuLc_messagesSubscriptionsMo() (*asset, error) {
	bytes, err := localesRuLc_messagesSubscriptionsMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/subscriptions.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/views/cache.html":              templatesViewsCacheHtml,
	"templates/views/subscriptions.html":      templatesViewsSubscriptionsHtml,
	"locales/ru/LC_MESSAGES/cache.mo":         localesRuLc_messagesCacheMo,
	"locales/ru/LC_MESSAGES/mqtt.mo":          localesRuLc_messagesMqttMo,
	"locales/ru/LC_MESSAGES/subscriptions.mo": localesRuLc_messagesSubscriptionsMo,
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
				"cache.mo":         &bintree{localesRuLc_messagesCacheMo, map[string]*bintree{}},
				"mqtt.mo":          &bintree{localesRuLc_messagesMqttMo, map[string]*bintree{}},
				"subscriptions.mo": &bintree{localesRuLc_messagesSubscriptionsMo, map[string]*bintree{}},
			}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"cache.html":         &bintree{templatesViewsCacheHtml, map[string]*bintree{}},
			"subscriptions.html": &bintree{templatesViewsSubscriptionsHtml, map[string]*bintree{}},
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
