// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package mercury

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x7b\x93\xdb\xb6\x11\xff\xdf\x9f\x62\x07\x73\xae\x4f\x49\xa8\x77\xdc\xf4\x22\x5d\xfe\xa8\xdb\x69\x66\xe2\x4e\xc6\xb9\xb6\xd3\x26\xe9\x0d\x24\x82\x27\xc4\x24\x40\x03\xa0\xce\xd7\x1b\x7d\xf7\x0e\x08\x3e\x40\x3d\x49\x1e\x69\x5b\x31\x35\xe3\xb3\x88\xc7\x2e\xf0\xdb\x1f\x76\x21\x72\xc1\xc7\x47\x70\x89\x47\x19\x01\xb4\xe4\x4c\x11\xa6\x10\x6c\x36\xcf\x66\x72\x29\x68\xa8\x40\x3d\x84\x64\x8e\x70\x18\xfa\x74\x89\x15\xe5\x6c\xf0\x1b\x5e\x63\x53\x89\xae\x9f\x01\x00\x78\x11\x5b\xea\x1a\x58\x63\x3f\x22\x7f\xe5\x22\xc0\xea\x92\xf5\xe0\x31\xae\xd5\x1f\x41\x54\x24\x18\x30\x72\x0f\xdf\x33\xe5\xf7\xff\x1e\x05\x0b\x22\x92\x96\xbd\xbe\x97\x76\xf9\x36\xee\xb1\x79\x36\x1b\x18\x05\xd7\xcf\x66\x2e\x5d\xc3\xd2\xc7\x52\xce\x91\xe0\xf7\x89\x46\xbb\x74\xc9\x7d\x27\x70\x9d\xd1\x18\xf4\x37\x19\xa4\xdf\xde\x4b\x67\x34\x4e\xda\x6f\xf7\x79\x7f\x1b\x62\x46\x7c\xab\x76\xb7\x45\x0a\x46\xb1\x4d\xd6\x4e\x70\x9f\xcc\x91\xc2\x8b\x7d\x92\xb2\x96\x91\x9f\x0a\x64\x78\x0d\x0c\xaf\x1d\x85\x17\x12\x16\x58\xdc\xea\x2f\x28\x17\xe3\x53\xb9\x4f\x57\x26\xc9\xa7\x49\xdb\x50\x10\x49\x98\x8a\x6d\x81\x1e\x1f\x81\x7a\x40\xde\x41\x1f\x1b\x13\x20\x6d\xbc\x54\xa7\x2e\x5b\x13\xdd\x88\x30\x17\x36\x9b\xeb\x19\x86\x95\x20\xde\xfc\x40\xbf\xdc\xb2\x57\x6b\x4e\xdd\xcb\x61\xef\x5b\xdd\xd7\x97\x04\x36\x9b\xc7\x47\xe8\xbf\x21\xef\x22\x22\x55\xff\x1f\x6f\x7e\xe8\xff\x88\xd5\xca\x14\x1b\xe1\xe8\x5a\x0b\x1d\x7d\xc3\x00\xfd\x39\x12\x82\x30\x05\x01\xc1\x32\x12\x24\x20\x4c\x49\x04\x7d\xd8\x6c\x66\x03\x7c\x3d\x1b\xf8\xb4\x81\x89\x06\x9c\xa9\x95\xff\x50\x77\xbe\x56\xf7\x3a\xd3\xfe\xce\xc8\x99\x27\x62\xfe\x80\x7d\x7f\x3e\xda\x87\xc5\x6b\xd3\x00\xa4\x9e\x88\x54\x74\xd9\x3c\x12\x2e\x95\xa1\x8f\x6b\x23\x61\x75\x7f\x0a\x12\x89\x98\x7d\x18\xbc\x4a\x35\x34\x3c\xf1\x15\xf7\xa9\x8b\x1f\x64\xdd\x99\xdb\xfd\x9f\x32\xf5\x54\xce\xbe\xb9\xff\x2d\xd3\x71\x7a\xf2\xb3\x41\xe4\x1f\xa8\xb1\x5c\x93\xc2\x0b\xe7\xb0\x73\x2a\xf4\xd8\x72\x52\xb6\x04\x5d\x02\x1e\x76\x09\x18\xc4\x80\xb2\x23\xd2\xf4\xe7\xe8\x1a\x3a\xda\x33\x1e\x4f\x98\x69\x27\xef\x95\x13\x44\x8a\xb8\xe0\x71\xa6\x9c\xd1\x04\x02\x67\xe1\x4c\x86\x27\xf4\x5b\xe3\x10\xf8\x1e\x2e\x8b\xbe\xc6\xc5\x8a\x80\xa2\x01\x01\xce\xc0\x25\x6b\xba\x24\xf0\x3c\x46\x9d\x51\x3f\xfb\x77\x19\x0a\xca\x14\xa0\x42\x74\xfb\xe5\x40\x78\xfb\x05\x5d\xbb\x7c\x19\x69\xf7\xd5\xbf\x17\x54\x91\x4b\xad\xe4\x86\xff\xa4\x04\x65\x77\x97\x2f\x10\x5c\xf6\x75\x49\xdf\x04\x32\x40\xe3\xe1\xf0\xa5\x33\x1c\x39\xc3\xf1\xcd\xe8\xeb\xab\xe1\xf4\x6a\xf8\xf5\x7f\x86\x7f\xbc\x1a\x0e\x51\x0f\xd0\x8b\x5e\x2f\x0b\x6b\xa8\xd7\x2b\x85\xd9\x20\xbc\x7e\x76\xba\x95\x8e\x1e\xc4\xb2\xad\x4f\x20\xfe\xeb\x48\x25\x68\x48\x5c\x8d\x0d\x36\xe5\xae\x72\x04\x91\x21\x67\x52\x9b\x9c\xf1\x7b\x81\x43\x04\x52\x3d\x68\x9e\xdc\x53\x57\xad\xae\x46\xc3\xe1\xf3\x92\x96\x98\xa9\x15\xc1\x6e\xd9\xb6\xa2\x5c\xc3\x44\xf0\x96\x07\x4d\x97\x8f\x5a\xd5\x94\x72\x33\x82\x01\xfc\xab\x01\x39\xe3\x86\xe4\x4c\x1a\x92\x33\xad\x25\x67\x36\x28\x6b\x0f\x2d\xb3\x82\x95\x17\xdc\x7d\xa8\xb0\x8e\xd9\x1d\x81\x0b\xfa\x15\x5c\xe8\xf8\x28\xe1\x6a\x0e\x7d\xf3\xad\xc4\xf2\x30\x0a\x85\xf1\x4b\x98\xb9\x70\x49\xde\x25\x82\xfa\x31\x69\xe0\xc2\xac\xcf\xf8\xa2\x67\x57\xff\x9b\x60\x91\xd6\xea\xef\x3d\x2b\x82\xc8\x68\xb9\x24\x52\x5a\x21\xa4\x82\x65\xdc\xcc\x32\xf6\x40\xfa\xc6\x65\xc0\x85\x56\xf3\xf8\x58\x18\x45\x6c\xb5\x92\xf0\xa6\x3a\x4a\x37\x8e\x3b\x94\xda\xc9\x6f\x79\x3a\x7b\x2b\x9f\x0f\xf8\x66\x04\x9b\x8d\xe5\xc7\x2a\x8d\xc3\xd8\xe9\x4e\xe5\xc2\x6e\x84\x06\x78\x58\xd6\xd6\xc5\x49\x85\x98\xa5\x16\xf3\xf1\x82\xf8\x10\xff\x75\x5c\x4d\x29\x81\xae\x1b\x9d\xf5\x2b\xe2\x2b\x5c\x9c\x3a\x0c\x9a\x05\xd6\x60\x51\x50\x31\x1b\xe8\x49\x56\x06\x39\xde\xb3\x50\x0f\xfc\x96\x91\x4e\xd7\x49\x07\x35\xe4\x8e\xa5\x1d\xa8\x29\xf3\x78\x87\x73\xd1\x79\x98\x09\x34\xec\x3c\x88\x87\x23\x5f\xb5\x0f\x75\x5d\x1c\x62\x3c\xcb\x87\x8a\xf3\x09\x2c\xe3\x26\x03\xcb\xf8\x5c\x02\xcb\xb8\xfd\x55\x38\x6e\x35\xb0\xb4\x82\x74\x2b\x81\xe5\x3c\xa1\xb6\x02\x4b\x2b\x50\x37\x1f\x58\xce\x13\x67\xdb\x79\x9c\x4d\x60\xd9\x07\x75\x17\x58\xb6\x40\x9a\x34\x19\x58\x26\xe7\x12\x58\x26\xed\xaf\xc2\x49\xab\x81\xa5\x15\xa4\x5b\x09\x2c\xe7\x09\xb5\x15\x58\x5a\x81\xba\xf9\xc0\x72\x9e\x38\xdb\xce\xe3\x6c\x02\xcb\x3e\xa8\xbb\xc0\xb2\x05\xd2\xb4\xc9\xc0\x32\x3d\x97\xc0\x32\x6d\x7f\x15\x4e\x5b\x0d\x2c\xad\x20\xdd\x4a\x60\x39\x4f\xa8\xad\xc0\xd2\x0a\xd4\xcd\x07\x96\xf3\xc4\xd9\x76\x1e\x67\x13\x58\xf6\x41\xfd\xa9\x05\x96\xf2\x4f\xc6\xaa\x8d\x62\x36\x28\xf9\x6c\x6c\x36\x88\x9f\xd5\x9e\x7c\x18\x6f\xad\xb8\x7d\xb9\x1c\xa7\xf5\x78\x5c\x04\x40\xdd\x79\xde\x2b\xe1\x81\xae\x70\x56\x5c\xd0\xff\x71\xa6\xb0\x0f\xf1\xb5\xa1\x85\x4f\x3c\x85\xe2\x07\xca\x8e\xe2\x77\x77\x3e\x99\xa3\x35\xf6\xa9\x8b\x15\x17\x08\x02\xa2\x56\xdc\x9d\xa3\x90\x1f\x4d\x66\x2a\x8c\xc2\xca\x6c\x88\xf5\xdc\x09\x1e\x85\x25\x3b\xc7\x02\x0c\x6b\xb3\x6c\x30\xa6\x04\xf7\x9d\xa4\xd0\xe4\x86\x4d\xd3\xd4\xb0\xa9\x95\x19\xa6\x67\x35\x47\x01\x77\xc9\xad\x1a\x59\x09\x1b\x3f\xad\xf8\x3d\xdc\x8c\xcc\x03\xd5\xe2\xfa\x10\xe4\x5d\x44\x05\x71\xd1\xf5\x17\x55\x59\x3b\x1b\xc4\x23\xaa\xd0\x61\x37\xc5\xed\x65\x3a\x8d\x97\x7b\x13\xdc\x4a\x49\xa5\x2c\x8c\xd2\xb5\xbc\x5c\x91\xe5\xdb\x05\x7f\x9f\xd9\xfd\x37\xe9\xc8\x7b\xaa\x96\x2b\x04\x0c\x07\x24\x47\x27\x66\x49\x76\x91\x25\x82\x24\xd9\x14\xfd\xa4\xa6\x07\x48\x89\x88\x98\x6c\x1c\x2d\x9b\xb8\xd9\x12\x81\x41\x15\xac\x5c\xba\x2e\xbb\x58\xcb\x37\xfd\x14\x88\x36\xde\x21\xda\xb8\x23\x1a\xc9\xd1\xb1\x88\x36\x3e\x48\xb4\x71\x47\xb4\x93\x44\x9b\xec\x10\x6d\xd2\x11\x8d\xe4\xe8\x58\x44\x9b\x1c\x24\xda\xa4\x23\xda\x49\xa2\x4d\x77\x88\x36\xed\x88\x46\x72\x74\x2c\xa2\x4d\x0f\x12\x6d\xda\x11\xed\x14\xd1\x70\xc0\x23\xa6\xb6\xc9\x96\x94\x76\x84\x2b\xa2\x94\x93\x2e\x2d\xd8\x4f\x3c\x53\xdb\x91\xef\x14\xf9\x42\x7e\x4f\xc4\x36\xf7\x4c\x61\x47\xbd\x02\x46\x39\xf3\x92\xeb\xfd\xc4\x8b\x2b\x3b\xde\x9d\x8c\xae\x34\x20\xdb\xb4\x8b\xcb\x3a\xd6\xd9\x08\x59\x31\x36\xbe\x3c\x10\x65\x69\x40\x3a\xca\x9d\xa2\x9c\x8b\xd5\x0e\xe5\xe2\xb2\x8e\x72\x36\x42\x39\xe5\xcc\xe5\x0e\xe5\x74\x71\xc7\xb6\x23\x6c\xd3\xcb\xb1\x78\xe7\xed\x7b\xa6\x88\x58\x9b\xfb\x8d\xc0\xb8\x82\x65\x72\xb6\x47\x61\x41\x3d\xef\x33\x60\x20\x8b\xcf\xc5\xa6\x6c\x4b\x11\x8a\xb9\x96\x5d\xc4\xb7\xd6\xe3\xd3\x6c\xfd\xa4\x0c\x36\x1b\x04\x21\x56\x8a\x08\x36\x47\xff\xfd\xd9\xf9\xf2\xd7\xef\x7e\x1e\x3a\x7f\xfa\xf5\x8b\x0b\xd4\xb1\x2c\xd8\xba\xed\x56\x60\x59\xc7\xb0\xfc\xb6\x5b\x76\xb1\xc3\xb0\x71\xc7\xb0\x53\x0c\x9b\x1c\x62\x58\xbc\xd3\xfd\xca\x9c\x4c\xc4\xcc\xfd\x5c\x62\xe9\x3e\x9e\x4d\x6c\x9e\x4d\xf6\xf0\x6c\xd2\xf1\xec\x14\xcf\xa6\x87\x78\x86\x3d\x45\x04\x84\x82\x48\x09\x0b\xbc\x7c\x0b\x8b\x48\x29\xce\x3e\x53\xaa\x4d\x6d\xaa\x4d\xf7\x50\x6d\xfa\xd1\xa8\x56\x99\x6b\x3e\xbb\x95\xdc\xa7\x2e\xba\xfe\xa0\x64\x2d\x67\xb8\x94\xb8\xdc\xf3\x24\x51\xce\xa4\xaa\x21\x0d\x4b\x13\x4b\x0a\x22\x89\xca\x76\xdf\x0b\xc5\x60\xa1\x98\xc9\xf8\x48\x2d\x68\x9a\xe4\x6b\xe0\x8d\xe9\x62\x0e\xbf\x1a\x61\x4f\x19\x81\x8c\x16\x01\xdd\x1d\x42\x9a\xdf\x93\x8e\x42\xe2\x75\xe1\x47\x92\xbe\xac\x39\x86\xe6\x7d\xd5\x6c\xa0\x0d\x5e\x2b\x7d\xc0\x7e\x21\xc2\x69\x3d\xbf\xf3\x73\xe7\xaf\xf2\x37\x56\xd4\x3e\x9d\x5d\xeb\xec\xfa\x27\x70\x36\x1b\x76\xce\x67\xbb\xf8\x21\x3e\x9d\x9d\x52\xa4\x52\xf2\x4d\x15\x1b\x00\x64\xe7\xa8\xb5\xd2\xfc\xe5\x0a\xa8\xf2\x51\x69\x4b\x94\x39\x92\xad\xe5\xed\x1c\xc8\xae\x98\x73\x5a\xda\x38\xf0\x89\xa4\x08\x7d\xb6\x4b\x39\x9d\x4f\xe0\x3a\x3a\x4a\xd9\x3f\x02\x5f\xe7\x6f\x26\xaa\xb9\xc6\x8f\x08\xff\xa7\x0e\x12\x4f\x76\x1d\xaf\x88\x49\x90\xa3\xe9\x26\xee\xec\x1c\x48\xdd\x65\x9f\xba\x5f\x65\x7e\x31\xe5\x37\xbc\xab\xaf\x7e\x73\x5b\x2e\xbe\x19\xa7\xc5\xf4\x89\x10\x5c\xd4\xca\x93\x54\x6e\xe1\xfd\x35\x69\x6e\xb5\xde\x51\x16\xa5\xf7\xff\x92\xe8\xa8\x35\xda\xb2\x2b\x76\xcf\xf8\xaa\xe9\xca\x3a\xd6\x49\xef\x2c\xbe\x07\xa7\x80\x41\xbc\x43\x2a\xf1\x46\x1c\xd8\x6c\x5e\xd4\x4d\x9f\x4f\x3f\xc6\xba\x8c\x58\xea\xdd\xfa\xa9\xb0\x39\x22\x4a\x70\x76\xb7\xd7\xdc\xb5\x85\x66\x19\xff\x8d\x0e\x35\x95\x3c\xfa\x86\xfd\xe8\x47\x02\xfb\x80\xbe\x7c\xee\x82\x24\x4b\xce\x5c\x64\x5f\x48\xb4\xa3\xf9\xa2\xf0\x66\xa4\xed\xda\x27\x8c\xea\x09\x34\x3e\x32\x31\x7b\x5e\x1f\x6d\x5a\xd5\x92\x8e\xb7\x3f\xb3\x81\xa1\x56\x6d\xb6\xd7\xd4\x5f\xcf\x13\x55\xd7\x55\xf0\xdf\x85\x17\x73\xa5\x7e\x1c\x38\x83\x80\xa8\xf4\xb9\x79\x9b\x5b\xaf\xa7\x45\x9f\xd7\xf8\x2d\xb1\xee\xd5\xd5\x8e\x3c\x01\x7e\x6b\x1e\x11\xb5\x13\x7a\xb6\xc4\x7f\xbc\xd8\x53\x1c\xcc\xa1\x20\x50\x6f\x0b\xdf\x00\x19\x9f\x64\xce\x0f\xc9\xbb\x1f\xb0\x54\xe6\x9e\x31\xf0\xf4\x81\x44\x6d\xf2\xf9\x58\x2a\x93\x6a\x71\xcb\x3d\xef\xb6\xdd\x5d\xd0\x51\x65\x1f\x8f\x98\x0d\xed\x6f\x0e\xcd\xae\xe6\x76\xe7\x43\x2e\x82\xf3\x20\x3b\x6b\x90\xeb\xec\x03\x52\x7d\x5b\xd7\xef\x8a\xe9\xac\x23\x7a\x2b\xb7\x65\x8e\xcf\xef\xc8\x7d\xde\x03\x55\x7b\x8a\xb7\x8a\xac\xcb\xe4\x6b\xf2\x5f\x36\x9c\xff\x07\x00\x00\xff\xff\xcb\xc5\x76\x31\x08\x5c\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x6f\x6c\x5b\x57\x15\xff\x15\xe2\x98\x38\x69\x1c\x36\xfe\x8e\xb2\xde\xc1\x5a\x3a\x5a\xb7\xb6\x93\xb2\xe1\x36\x2d\xa5\x7f\xd6\x76\xcb\x56\x5a\x33\xf8\x84\x78\xb5\xaf\x63\x53\xfb\x3d\xf3\xde\x73\xbb\x48\x43\x4a\x5c\xda\x6e\x6a\x58\x10\x30\x75\x12\x83\x09\x18\xe2\x0b\x12\x6e\x52\xaf\x6e\x12\x3b\x5f\x10\x0c\x69\x1f\xce\x45\xa0\x31\x24\x10\x1a\x08\xc1\x3e\x20\x04\x12\xda\x24\x24\xd0\x79\xf7\x36\x76\x12\x07\x85\x0f\xad\xd4\x9c\x73\xcf\x39\xf7\x9c\xdf\x39\xe7\x77\x9f\xfc\xfb\xbb\x7a\x9e\x03\x80\x07\x01\x6c\x01\xf0\xab\x4d\xc0\x09\x00\xd3\xef\x44\xf0\xaf\xd6\x03\xbc\x07\xc0\xf5\x1e\xe0\x7d\x00\x7e\xd6\x03\x84\x00\xfc\xae\x07\xe8\x05\xf0\xc7\x1e\x60\x07\x80\x37\x7b\x80\x0f\x01\xb8\x3b\xa4\xcf\x7b\x42\xfa\xfc\xf9\x10\x70\x2f\x00\xdf\xc8\x6f\x86\x74\xbe\xd9\x10\xb0\x19\xc0\xcf\x43\x00\x97\xfa\x75\x08\x78\x17\x80\xd7\x43\x40\x3f\x80\x37\x42\x40\x18\xc0\xdf\x8c\xfd\x5f\x21\x60\x1b\x80\xff\x18\xb9\xad\x17\xb8\x0f\xc0\x23\xbd\xda\x5f\xea\x05\xee\x07\x70\xbe\x57\xd7\xfd\x76\x2f\xb0\x15\xc0\x4d\x13\xf7\x5a\xaf\xce\xf7\x76\x2f\xd0\x03\x60\x53\x58\xcb\xde\x30\x30\x08\xa0\x3f\xac\xf1\xdc\x13\x06\xfa\x00\xec\x08\xeb\x3e\x93\x61\x8d\xef\xa1\xb0\xc6\xb5\xdf\xd8\x8f\x87\x81\xbb\x00\x3c\x16\xd6\xf5\xad\xb0\xce\xff\x25\xe3\x77\x4d\xfe\x27\x4d\xbe\xaf\x18\xff\x15\x23\x67\x8c\xbc\x66\xe4\x8b\x26\xff\x8f\x4c\xfc\x5c\x18\x88\x00\x68\x9a\xf3\x2b\x26\xee\x97\x46\xfe\xd6\xc8\x37\x8c\x7c\xd3\xd4\xfd\x47\x18\xc8\x6f\x02\xde\x0a\x03\xc3\x8c\xab\x0f\xf8\x04\x80\xe7\xfa\x80\x01\x00\xaf\x1a\xf9\x7a\x1f\xf0\x45\x00\x7f\xee\x03\x0e\x02\x78\x38\x02\xe4\x18\x5f\x04\x38\xca\x73\x8f\x00\xfb\x00\x8c\xf4\x03\x23\x00\xca\xfd\xc0\x47\x01\xfc\xb0\x5f\xf3\xe0\x17\xfd\xba\xef\x3f\xf4\xeb\xf9\xfd\xb5\x1f\x18\x02\xf0\x56\xbf\x9e\x4b\xdf\x80\xb6\x6f\x19\xd0\x75\x76\x0e\xe8\xfc\xd3\x03\xc0\x17\x00\xfc\x73\x00\xf8\x30\x80\xd3\x9b\x81\x93\x9c\x7f\x33\x70\x0a\xc0\xc2\x66\x20\x0d\x60\x60\x50\x9f\xad\x41\x8d\x77\x76\x50\xd7\x5b\x32\xf2\x95\x41\xe0\x63\x00\xd4\xa0\xe6\xc3\xdb\x83\x7a\xcf\x5b\xa3\xda\x7f\x28\xaa\xf9\x79\x32\xaa\xf9\xf6\x99\xa8\x9e\x67\x2e\xaa\xe3\xbf\x1c\xd5\x79\xaf\x45\x35\xce\x97\xa2\x1a\xff\x5c\x54\xe3\xff\xa9\x39\xbf\x16\x05\xde\x0f\xe0\x2f\x46\xb2\x91\xe5\x07\x8c\x8c\x0f\x69\x9e\x1d\x1f\xd2\xfd\xe4\x87\x74\xfe\xcb\xc6\xfe\xe3\x21\xbd\xd7\x57\x8d\xfc\x8d\x91\x7f\x32\xf2\xef\x43\xba\xce\xbf\x87\x80\x4d\xd0\x18\xf7\x42\x63\x63\xde\xf1\x0e\x77\x41\xef\x96\xfb\xe2\x5d\x24\x01\x7c\x04\xc0\x7b\xb9\x3e\x80\x18\xcf\x17\x9a\xf7\x51\x00\xbb\x19\x1f\x34\x8e\xce\x7f\xef\x58\x75\xe6\xbe\xb7\x1b\xfd\xae\x0e\x3b\xcf\x6b\x0f\x80\x04\x00\x01\xcd\xe3\x0f\x42\xe3\x7d\x00\x7a\xce\x43\x26\x76\x4b\xc7\xbd\xb0\x91\xcc\x0f\x9e\xfd\xbb\xcd\xf9\x1e\x23\x99\x87\xf7\x76\xc4\xf3\xfe\x78\x86\x77\x43\x7f\x33\xee\x37\x76\x9e\xe3\xc7\xa1\xdf\xe3\x56\xf3\x6d\x62\xee\xf1\x3b\xc0\xb6\xac\xf0\x64\xc6\xb1\xb3\x6d\xcd\xc3\xce\xb6\x75\x67\x87\xf9\x50\xd9\x2d\x14\x71\xa8\x32\x5e\xf1\x7c\x1c\xce\x5b\xf6\xb8\x14\xd9\x82\x57\x2e\x5a\x13\xa2\xe4\x64\xa5\xc8\x59\x85\xa2\xcc\x8a\x0b\x05\x3f\x2f\xa4\xeb\x3a\xae\xd8\xe6\x75\x0d\xf4\x2a\x99\x8c\xf4\xd6\xf8\xfc\x42\x69\x63\x49\x82\xc0\xe5\x24\x15\xd7\x95\xb6\x2f\xb2\x96\x2f\x85\x65\x67\xb5\xd7\xb1\x45\x49\xfa\xd2\x5d\xe9\xbe\xed\xca\xca\xf3\x85\x8c\x0c\x32\x1b\x77\x49\x5a\x5e\xc5\x95\x25\x69\xfb\x1e\x8e\x74\xa6\xc2\x11\x6b\x02\x47\x64\x46\x96\xce\x4a\x17\x47\xa4\x97\x71\x0b\x65\xbf\xe0\xd8\x38\xa2\xe1\xe0\x98\x3c\xeb\x56\x2c\x77\x02\x0f\x4b\x7f\x03\x13\xe9\x8c\x5a\xbf\x65\x8e\xca\x3b\xc5\x42\xd6\x9a\xf0\xba\x47\x1c\x37\x5e\x9c\xb0\x7d\xe9\x9e\xb7\x8a\x22\xe7\xb8\xc2\xca\xf9\xd2\x15\x65\x57\x7a\x9e\x38\x6b\x65\xce\x89\xb3\x15\xdf\x77\xec\x95\x41\x19\xd3\xb5\x6f\xb9\x85\x5c\x6e\xa5\xcf\x76\xfc\xff\xe9\x2f\x3b\x17\xa4\xbb\x4b\x23\xe7\x19\xf1\x60\x71\xd2\xb2\x83\x09\x9c\xac\x14\xf9\x8f\x2d\xf1\xa8\xe5\xf9\x3a\x56\x38\xb9\xdc\x8a\xa3\x8d\x31\xeb\x9c\xd4\x17\xc7\x2c\x37\x93\xc7\x98\x35\x81\xb1\xf6\x06\x30\xe6\xd8\x7e\x5e\xff\x2d\x4e\x08\xcf\xb7\xfc\x82\xe7\x17\x32\x1e\x1e\x73\xce\xeb\x3d\x3c\x9e\xf1\x1d\x96\xa7\xa5\x27\x7d\x9c\xb1\xce\x4b\x9c\x91\x65\x5f\x3b\xcf\xe4\x9d\x0b\x22\x9d\x30\x32\x69\xe4\xb0\x91\x23\x5a\x5a\x25\xa7\x62\xfb\x5a\x0f\xa0\x04\x5a\x00\x51\xab\xc1\xf2\xd3\x09\xb1\x47\x7c\x2e\x8f\x74\xd2\xc8\x61\x23\x47\xb4\x7c\xc2\x2a\x56\x24\x4e\xcb\xb2\xe3\xfa\xb1\x31\x6f\xbc\x90\x8d\x7d\xba\x32\xee\xc5\xd2\x4e\x8a\x59\xf6\xa9\x73\x85\xbc\x55\x72\x76\xbb\x95\xc8\xa9\xc7\xd3\xb1\xc3\xae\xb4\x98\x3a\x31\x66\x58\x4a\x24\xe3\x89\x4f\xc6\xe2\xc3\xb1\xe4\x83\x22\x39\x9c\xda\xbb\x77\x67\x7c\x38\x1e\x8f\xf0\xa8\x62\x69\xd7\xb2\xbd\xa2\xe5\x3b\x6e\x4a\x3c\x12\xe4\x10\x63\x15\xd7\x2a\x39\x59\x47\xec\x5f\x91\xf8\x40\xe4\x51\xcb\x1e\xaf\x58\xe3\x32\x96\x96\x56\x29\x25\x96\xcf\x29\x71\xba\xe2\x79\x05\xcb\x8e\x8c\x9d\x18\x3b\x1a\x7b\x42\xba\x5e\xc1\xb1\x53\x22\xb1\x3b\x1e\x39\xec\xd8\xbe\xb4\xfd\x58\x7a\xa2\x2c\x53\xc2\x97\x4f\xfa\x7b\xca\x45\xab\x60\xef\x13\x99\xbc\xe5\x7a\xd2\x1f\xfd\x6c\xfa\x58\xec\xa1\x76\x1c\xe3\xc9\x49\x37\x76\xd4\xce\x38\xd9\x82\x3d\x9e\x12\x91\x53\xc5\x8a\x6b\x15\x63\xc7\x1c\xb7\xe4\xa5\x84\x5d\x0e\x8e\xde\xe8\xf0\x3e\xa1\xd5\x51\x7b\x5b\x22\x3e\x3a\x9a\x10\xdb\xb7\x0b\x56\xe3\xf7\x8d\x26\x12\xe2\xa0\x88\x8b\x54\x70\x3e\x30\x9a\xbc\xed\xda\x3f\x3a\xc2\xea\x8e\x20\x6c\x7f\x22\x2e\x9e\x7a\x4a\x5f\x39\x30\x9a\x8c\x3f\x20\x0e\x8a\x84\x48\x89\xe4\x3e\xfe\x4a\xa9\x29\xaa\xd3\xbc\xba\x48\x4d\xba\x41\xb5\xd5\x16\x75\x75\xb5\x25\xf8\x86\xad\xba\xb4\xda\xa4\xae\xae\x31\x81\xbe\x4e\x4b\x6a\x92\xea\xb4\xa0\xa6\xf9\x30\x4b\x73\xea\xa2\x9a\x52\x55\xd0\xf3\x74\x93\x16\xa9\x4e\xcd\xe0\x7f\x83\xea\x22\x08\x7c\x99\x1a\xb4\x48\x2d\x9a\x15\xd4\x08\x0a\x35\x68\x9e\x6a\xea\x32\x35\xa8\x21\xe8\x26\xd5\x68\x96\xea\x6a\x52\x3d\x4d\x0d\x5a\xa0\x96\x9a\x52\xd3\x42\x4d\x09\x6a\x05\x96\xeb\x34\x4f\x2d\xba\xc5\x2f\x9b\x7e\xa0\xa6\x68\x89\xea\xea\x69\x6a\x52\x8b\xb3\x75\xd4\x53\x57\x3b\xaa\xa9\xab\x5d\x6a\x75\x03\xc8\x51\xaa\xca\xe5\x69\x96\x6a\x5c\xfe\xce\xc3\x5c\x5d\x73\x1d\xb0\xdf\xd7\x43\x57\xcf\x68\xa0\x37\xa8\xa6\xaa\x54\x13\x0c\x66\x36\xe8\x74\x51\xcd\x08\x6a\x52\x4d\xa8\x29\x75\x85\xea\xaa\xaa\xae\x04\x19\xea\x9d\x77\x6b\x1c\xb4\x7c\x37\x88\x0e\xb6\xa5\x26\x19\x2e\x6b\xdc\x56\x00\x7b\x55\xc1\x25\x6a\x31\x1c\xee\x9c\xd1\xd2\x02\x83\xfa\x56\x37\x14\x6c\xe6\xd6\xa6\xb5\xc2\x97\xae\xab\x49\x3e\x7e\x97\x96\xa8\xa1\xa6\xa8\xa6\xc7\xcd\x1b\x58\xd1\x27\xdf\x7d\x89\xea\x41\xaa\x9a\x66\xd4\x8b\xd4\xa2\x05\x75\x91\x3b\xba\x33\x24\x5a\x5b\xe0\x8e\x90\xa0\x4b\x99\xa5\xa0\xcb\x9b\x74\x83\x0d\xea\x0a\x93\x41\x5d\x12\xc1\xb1\x4e\xb7\xfe\xcf\xf4\x6b\x73\x69\x96\x34\x35\xd1\x57\x34\xd4\xb5\x99\x25\xce\x4f\x0b\x7c\xab\x49\x35\x7a\x39\xd8\x6c\x83\xe9\x32\xcf\xbc\xa5\x25\x9a\xdf\x60\x2a\xf5\x35\xee\x40\x4d\xd2\x9c\x3e\x56\x97\x89\x54\xa7\x39\x6a\xb1\xa1\xa6\x26\xa9\xa1\xbe\x4a\xb5\x8d\x61\x6b\xf2\xda\xdb\x69\x1a\xea\x52\x67\x12\xbd\x1e\xa6\xfd\x45\x5a\xe4\xc7\xbe\xb1\xa4\x8b\xd4\x52\xcf\x70\x6b\x01\xeb\x1b\xbb\x96\x19\xac\x37\x14\x70\x3a\x78\x29\x9c\xf0\x27\xd4\xe4\x54\x9a\xc6\xcf\xab\x67\x35\x39\x59\x69\x1a\x96\xea\xe1\x99\xed\xd5\x79\x47\x55\x9a\xa7\x05\xf5\x6c\x7b\xe5\x5d\xe3\x66\xd7\x46\x2d\xbf\xaa\xa5\xe0\x59\xf2\x07\x63\x96\x5a\x74\xc3\x3c\xcf\x1a\xe8\x85\x00\x4a\x35\x50\xe8\x96\x66\xd7\x8a\xc7\xc9\xa0\x5e\xa0\xba\x9a\x52\x33\xea\x72\x5b\xbd\xc2\xbb\x55\x33\x22\xc8\xa4\x17\x1c\x34\xcf\x97\x41\xdf\xa1\x96\x9a\x69\xbf\xd5\x79\x55\x6d\x1f\xbf\xc7\x4a\x80\xbe\xa1\xaa\xda\xd0\x52\x97\x02\xd6\x35\xdb\x26\x7e\xf6\x1d\x97\x96\x61\xa9\xab\xc1\xf4\xaa\x6a\x9a\x7f\x70\x74\xb7\x27\xd7\xb1\x0f\xaf\x63\x1f\xe9\x6e\x5f\xa6\xc1\xc5\xee\x7e\xb3\xd3\xf5\xbc\x2b\x58\xb1\x5e\x0f\x1d\xdf\xba\xe0\x97\x0f\x7d\x83\xbf\xb6\xfa\xd7\xcf\x6d\x7d\xb8\x43\x1f\x69\xeb\x74\x2d\x58\x41\x7b\xdb\xff\x0d\x00\x00\xff\xff\x3d\xc5\x77\x82\x0d\x11\x00\x00"

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
	"templates/views/widget.html":      templatesViewsWidgetHtml,
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
