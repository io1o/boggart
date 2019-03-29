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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x99\xeb\x6f\xdb\x36\x10\xc0\xbf\xef\xaf\x38\x10\x2d\xfa\x9a\xfc\x90\xdd\xae\x0b\xec\xec\xc3\x1e\xc0\x80\x15\x18\xd6\x62\x1f\x56\x74\x01\x2d\x9e\x63\x36\x12\xa9\x92\x27\x3b\x99\xe1\xff\x7d\x90\x68\xc9\xf2\x43\xb6\x94\x26\xeb\x1a\x05\x88\x21\x1e\xc9\xe3\x1d\xef\xa7\x23\x45\x2e\x97\x20\x70\x2a\x15\x02\x0b\xb4\x22\x54\xc4\x60\xb5\xfa\x06\x00\x60\xb9\x04\x39\x05\xfc\x04\x1d\x1e\x90\xd4\x0a\x98\x90\x36\x0e\xf9\x4d\xd1\x62\x24\xe4\x1c\x82\x90\x5b\x3b\x66\x46\x2f\xd8\x79\x26\xdd\xad\x09\x74\xe8\x45\xc2\xeb\xfb\x90\x3e\xd9\x28\x7f\xba\xb6\x5e\xdf\x2f\xf5\xd9\xed\x77\x7d\x11\x73\x85\xe1\x4e\x8b\xfd\x56\x24\x29\xc4\x03\xad\xb2\x96\x33\xff\x3c\xf5\xa3\xff\x5a\x01\xfb\xc9\x99\x0f\x91\x16\xc8\xa0\x03\xab\xd5\xa8\x3b\xf3\x2b\x3a\x96\x1d\x08\x91\x9b\xa9\xbc\x66\xe7\xa3\xae\x90\xf3\x03\xf6\x54\x88\xb7\xcc\xcc\xa7\xb7\x62\xbc\xa9\x36\x11\x48\x31\xde\x4c\xf2\xba\x67\x5a\xe1\xcd\xb4\x91\xff\x68\x45\x3c\x84\xac\x1c\xf2\x09\x86\x5e\x88\x53\x62\x20\x38\x71\x8f\xf4\xe5\x65\x88\x63\x36\xe7\xa1\x14\x9c\xb4\x61\x10\x21\xcd\xb4\x18\xb3\x58\xdb\xaa\x51\x77\xad\xcc\x74\x5f\x1a\x9d\xc4\x47\x3a\x64\x9d\x32\x03\x36\x11\x56\x64\x74\xe8\xad\x85\x2e\xde\xc3\x3c\xdc\xc3\x52\xb4\x53\xeb\xc7\x2c\x0d\xc0\x05\xf5\xd9\x26\x34\x6f\x67\x7a\x01\xef\xfa\x2e\x2a\x30\xb2\x31\x57\x05\x58\xf8\x29\x91\x06\x05\x3b\x7f\x3e\xea\xa6\x15\x27\x2c\xeb\x66\x56\x9c\x68\xb4\x8f\xe7\xab\xdc\xdc\x57\x95\x70\x1e\xd4\x24\x55\x9c\x10\xd0\x4d\x8c\x63\x16\xcc\x30\xb8\x9a\xe8\xeb\x22\x76\x1f\xad\x67\x17\x92\x82\x19\x03\xc5\x23\xdc\x78\x9e\x45\xba\x28\x14\x6f\xda\xd3\xd8\x48\x45\xd0\x59\xd7\x3c\x03\x46\x26\xc1\xf4\x7d\x83\x4c\x37\x8a\xe5\x12\x50\x89\x54\xd0\x3d\x35\x0f\x07\x99\xac\x5b\xfd\xa5\xa0\xf0\xf7\xa0\xf0\x5b\x02\x85\x5f\x86\xc2\xaf\x84\xc2\x6f\x21\x14\x83\x3d\x28\x06\x2d\x81\x62\x50\x86\x62\x50\x09\xc5\xa0\x85\x50\x0c\xf7\xa0\x18\xb6\x04\x8a\x61\x19\x8a\x61\x25\x14\xc3\xf6\x41\xc1\x23\x9d\xa4\x5b\xac\x6d\x30\xd6\xd2\x56\xc0\x91\xfb\x5a\x00\x92\x0b\x0e\x43\xe2\x6a\xdb\x07\x4a\xac\x17\x68\x76\x39\x71\xc2\x56\x60\xb2\x76\xb5\xa0\x64\x5d\x3e\x0c\x49\x56\xd9\x3e\x46\x48\x46\xb8\x8b\x48\x26\x6b\x05\x21\xce\xd3\xcd\x3a\x93\x15\x2b\x56\x1a\x19\x61\xfb\xf0\x10\x9c\xf6\xf0\xc8\x64\xad\xc0\xc3\x79\x5a\xe0\xe1\x8a\x7b\x78\xa4\xe2\x56\x91\x91\xbe\x0a\xdb\x27\x1b\xbf\x2a\x42\x33\x77\xe7\x36\xa0\x34\x41\x90\x18\x83\x8a\x80\xb8\x91\xd3\xe9\x57\x4e\x8b\x4a\xa2\x49\xba\x70\x38\x32\x72\xef\x33\x2e\x8a\xc2\x9c\x87\x09\x8e\xd9\x72\x09\x9d\xb5\x0c\x56\x2b\x06\x31\x27\x42\xa3\xc6\xec\xef\xf7\xde\x8b\x0f\x3f\xbc\xef\x79\xdf\x7f\x78\xfe\x88\x3d\x4c\x22\xfc\x2a\x22\x1e\x3c\x0d\x7e\x99\x06\xff\x00\x0d\x7e\xfb\x68\x18\x54\xd1\x90\xed\xb4\xbe\xcd\x76\x19\xc0\x95\x78\x08\xeb\xc9\x21\x26\x06\x65\x26\x06\x07\x98\x18\xb4\x8f\x89\x61\x15\x13\x7c\x4a\x68\x20\x36\x68\x2d\x4c\x78\x70\x05\x93\x84\x48\xab\x07\x88\xc5\xb0\x8c\xc5\xf0\x00\x16\xc3\x7b\xc7\xa2\x16\x17\xa1\xba\xb0\x3a\x94\xa2\xf2\x3e\xe8\x50\xa7\xfa\x30\xd5\x9b\xf0\x1c\x2c\x3d\x9d\x5a\x24\x6f\x50\x27\x00\x8e\x9c\x75\x04\x0c\x5a\xa4\x62\x97\x37\x21\x05\x13\x52\x9e\x54\x53\x5d\xcc\xbc\x6b\xb2\xe1\xf2\x0f\xd7\xc5\xdd\x9c\x39\x65\x4d\x47\xb5\xc9\x24\x92\xfb\xc3\xda\x24\x08\xd0\xda\x62\x64\xcb\xe7\x5b\x9b\xeb\xb4\xd8\x60\xdc\xdb\xe7\x80\x51\x37\x0d\x54\xad\x2b\xbe\x1d\x51\xa9\x58\x7a\x4c\xb7\xba\xa1\xc5\xaf\xf4\xc6\xf4\xc7\xf5\xee\x24\x42\x6e\x13\x83\x11\x2a\xb2\xff\xbb\x9b\x53\xe2\x93\x10\xf3\xd6\xae\x90\xfd\x7a\x96\x8c\x8c\x31\x5b\x45\xb9\x93\x0b\xf2\x0c\xda\x58\x2b\x2b\xe7\x08\x4a\x2f\x0c\x8f\x19\x58\xba\x09\x71\xcc\x16\x52\xd0\xec\xac\xdf\xeb\x3d\x3e\x76\x5b\x4a\x33\xe4\xe2\x58\xbd\x39\x81\x26\xcd\x72\x5b\x23\xe1\xa5\x51\x2d\xef\x0d\xdf\x6c\x26\x3a\x9f\x67\x9a\x7d\x8e\xc2\x3f\xd3\xf7\xa9\x81\xaa\xd2\x6d\x39\xda\xc0\xc8\x98\x64\xbe\xd6\x1c\xef\x3f\xea\x1e\xf3\x3c\xed\x7b\x62\xde\x26\x5a\xdc\x9c\x32\xef\xc4\xdc\xba\x46\xa2\xe4\x03\x27\xb7\x91\xda\x9c\xdb\x8c\xba\x74\xc4\x8c\xfc\xcf\x7d\xc9\x66\xdf\xaf\x69\xd7\x0e\x1a\xa3\x4d\xfe\x12\xd7\x30\xa1\xc0\x11\xaf\xc9\x13\x5c\x5d\xae\x4f\x1b\x77\x34\x76\x7e\x5e\xeb\xad\x6d\x55\x39\x99\xd4\x99\x8a\x5a\x0d\xb3\xc6\x2e\xda\xeb\x34\xcd\xe3\x38\x94\x01\x4f\x83\xdf\xfd\xc8\xe7\xdc\x55\xb2\x73\xa1\x83\x24\x85\xb3\xb3\x30\x92\xf0\x69\xea\xcd\x3b\xfd\x96\x8c\x54\x97\x4f\x9f\x6c\xf9\x97\x25\xf2\xce\x2f\xda\x44\x9c\x80\xf9\xbd\xde\x2b\xaf\xd7\xf7\x7a\xfe\xbb\xfe\xcb\xb3\xde\xf0\xac\xf7\xf2\xaf\xde\x77\x67\xbd\x1e\x83\xd5\xea\xc9\xb3\x67\xa3\xae\x1b\xa1\xbe\xbd\x2e\x42\x0a\x4b\x43\x0a\x0c\x89\x43\xaf\xee\xf4\x6c\x3c\x27\xa3\xd5\xe5\xc1\x90\x35\x52\xe4\x4c\xba\xa4\xcf\x37\x29\xd7\xd6\x7f\xad\x7e\x0f\x13\xc3\x43\x60\x2f\x1e\x0b\xb0\x18\x68\x25\x58\xb9\x90\xa6\xe3\x9d\xd1\x1e\x81\x92\x61\xf1\xbf\x5b\xdb\xd0\x92\x86\xc8\x1d\x71\xa0\x6c\xff\x7f\x6a\x7e\x76\xd2\xd4\x8c\x88\xae\x43\xa2\x11\x8d\x0d\xc6\xa9\xff\xb6\xd7\xd3\xb9\x95\xf3\xf2\x35\x5b\x94\x73\x1f\x68\x05\x11\x52\x7e\xbd\x71\xda\x80\xe3\xe9\x1c\x6e\x95\x8d\xdf\xf0\x2b\x2c\x7d\xd2\x36\xca\xc4\x11\xbf\x72\xa7\x8c\x77\x97\x8a\x77\x54\xde\x7f\x2e\xde\x1e\xb4\x2a\x41\xb2\x46\x26\xdc\x02\x90\xc6\x61\xb8\x0f\x16\x7e\xe3\x96\xdc\x71\x07\xe8\xfc\xdc\xab\x11\x10\x21\xb7\xe4\x6e\xa6\x2e\xf4\x74\x7a\x71\xf7\x2b\xf5\xd1\x01\xee\x1f\x96\x3b\x5a\x8f\xab\xbc\xb8\xe5\xf2\x7c\x1f\x60\x7e\x79\x00\xd5\x67\xf2\xa7\xee\x19\xbf\x5d\xfd\x5f\x25\x7d\xea\x41\xc2\x37\xea\x1e\xf9\x70\x19\x75\xb3\xcf\xce\xbb\x3d\x4d\x70\xbe\x15\x4f\xff\x06\x00\x00\xff\xff\x0d\x2f\x72\x54\xc1\x2d\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x4b\x6c\x1b\x55\x14\x3d\x8d\xed\x26\x71\x5c\x3e\x85\x52\xa0\x85\xbe\x8a\xa6\x2d\xb4\xd3\xfa\x93\xf2\x71\xea\x14\xd1\x0f\x14\x1a\x51\xa5\xa6\x42\xac\x98\xda\xcf\x89\x15\x7b\xc6\xcc\x8c\xfb\x41\x5d\xc4\x49\x4b\x5b\xa5\x85\x0d\xa8\x15\x0b\xc2\x67\xc1\x12\x37\x8d\x5b\x93\x8f\x03\x4b\x16\x48\xf7\x89\x0d\xb0\x40\x62\x83\x40\x08\x09\x16\x6c\x10\x42\xe8\xcd\x9b\xda\x4e\xe2\x40\x36\x1d\xc9\xba\xef\xdd\x7b\xdf\x79\xe7\x9e\x7b\x9f\xfc\xc3\x5a\xff\x7b\x00\xb0\x19\xc0\x46\x00\x5b\x56\x01\xbb\x01\xec\x6c\x83\xfb\xfd\xdc\x06\xdc\x0f\xe0\x97\x36\xe0\x01\x00\xff\xb4\x01\xdb\x01\xac\xf3\x01\x1b\x00\xf4\xf9\xd4\xfe\x35\x6f\x3f\xe6\x03\x1e\x05\x30\xe1\x53\xe7\x6a\x3e\x60\x0d\x80\x6f\x7c\x40\x17\x80\x1f\x7d\x40\x08\xc0\x6f\x3e\xa0\x1b\xc0\x5f\x9e\x65\x7e\x60\x0b\x80\x43\x7e\x85\x33\xec\x07\x36\x01\xb8\xe0\x57\xdc\x3e\xf5\x03\x77\x01\xf8\xca\xaf\xf0\xbe\xf5\x03\x9d\x00\x7e\xf2\x2b\xdc\xdf\xfd\x40\x00\xc0\xdf\x7e\xc0\x0f\xa0\x2d\x00\xb4\x03\x68\xf7\xec\xdd\x9e\x5d\xef\xd9\x4d\x01\x75\x6e\x5b\x40\xe1\xc4\x02\x40\x10\x40\xc2\xdb\xbf\x10\x50\x78\x03\x01\x60\x68\x15\xf0\x6a\x00\x88\x01\xf8\x72\x35\xf0\x24\x80\x3f\x57\x03\xaf\x4b\x9d\xda\x81\x7d\x00\xae\xb4\x03\x19\x00\x5f\xb7\x03\x07\x01\xf4\x74\x00\xbd\x00\xde\xec\x00\x1e\x03\x70\xb3\x43\xe9\xf7\x7d\x07\x70\x0f\x80\x3f\x3a\x80\x47\x24\xcf\x4e\x85\xc3\x3a\xd5\xf9\x52\x27\xf0\xa2\xd4\xbd\x13\x38\x2a\x79\x05\x81\x24\x80\x73\x41\xb5\xff\x2e\x08\x6c\x03\xf0\x44\x97\xd2\x6d\xa0\x4b\xe9\x73\xba\x4b\xe9\xfd\x7e\x97\xc2\xff\xac\x0b\x58\x0b\xe0\x8b\x2e\x60\xbd\xd4\xdf\xb3\xbf\x7a\xd6\x17\x52\x76\x43\x48\xe9\x1c\x0b\x29\x3e\xc7\x42\x0a\xf7\x0d\xcf\x7f\x35\xa4\xf0\x6e\x85\x80\x55\x00\x1e\x54\x63\xe1\x9e\x95\x9f\x0f\xc0\x6a\x6f\xbd\x0e\x4a\x5b\xf9\xc9\x1e\x06\xbd\xb5\xd4\x59\x62\xdc\x07\xd5\x43\x89\x1b\xf2\x62\x1b\x3d\xbb\xc6\xb3\x9b\x3d\x2b\x67\x48\xea\x2f\x39\xdd\x0b\xd5\x53\x59\x9f\xac\x89\xa1\xf1\x49\x4d\x3b\xa0\x7a\x26\xbf\x87\x00\x3c\x2c\x75\x95\x3d\x6f\xca\x43\x77\x9a\xd9\x3c\x65\x1a\xe9\xc6\xca\xc6\x8e\x86\x77\x47\x93\x7b\xff\x90\x6e\x0c\x72\x96\xce\xda\x85\x9c\x7e\x86\xe5\xcd\x34\x67\x19\x3d\x9b\xe3\x69\x76\x2a\xeb\x0c\x31\x6e\x59\xa6\xc5\xba\x5b\x27\xda\xc5\x54\x8a\xdb\x4b\x62\x4e\x36\xbf\x32\x10\x37\xb1\x0e\x52\xb4\x2c\x6e\x38\x2c\xad\x3b\x9c\xe9\x46\x5a\x45\x4d\x83\xe5\xb9\xc3\xad\x7a\x38\xcf\x75\xbb\x68\xf1\x3c\x37\x1c\x1b\x07\x9a\x73\x71\x80\xdb\x29\x2b\x5b\x70\xb2\xa6\x81\x03\x4d\x3c\xf1\x3c\x77\x56\x50\x61\x73\xd6\xf2\x25\x1c\x36\x1c\x6e\x9d\xd4\x73\x2c\x63\x5a\x4c\xcf\x38\xdc\x62\x05\x8b\xdb\x36\x3b\xa1\xa7\x86\xd9\x89\xa2\xe3\x98\xc6\xc2\xa4\x94\xc7\xdc\xd1\xad\x6c\x26\xb3\x30\x66\x98\xce\x7f\xc6\x0b\xe6\x29\x6e\xed\x54\x74\x64\x9d\x52\x1c\x1c\xd1\x6d\x47\x45\x98\x99\xc9\x2c\xd8\x1a\xe8\xd7\x87\xb9\x4a\xeb\x6f\x28\x85\x01\x6e\x73\x07\xc7\xf4\x93\x1c\xc7\x86\xcc\x53\x2c\x19\xf1\x6c\xd4\xb3\x31\xcf\xf6\x28\xab\xe7\xcd\xa2\xe1\xa8\xb5\x0b\xe6\xae\xdc\x4b\xd4\xd2\x15\xfc\xb8\x9e\x2b\x72\x0c\xf0\x82\x69\x39\x5a\xbf\x3d\x98\x4d\x6b\xcf\x15\x07\x6d\x2d\x69\xc6\x59\x9a\x9f\x7c\x76\x38\x3b\xa4\xe7\xcd\x5d\x56\x31\x78\xf4\xe5\xa4\xb6\xdf\xe2\xba\xec\x8d\x26\xbb\x16\x67\xd1\x70\xe4\x19\x2d\x1c\xd3\xa2\x4f\xb1\x68\x2c\xbe\x67\xcf\x8e\x70\x2c\x1c\x0e\xca\x62\xb4\xa4\xa5\x1b\x76\x4e\x77\x4c\x2b\xce\x5e\x72\x31\x58\x7f\xd1\xd2\xf3\x66\xda\x64\x7b\x17\x00\xf7\x05\x8f\xe8\xc6\x60\x51\x1f\xe4\x5a\x92\xeb\xf9\x38\xab\xef\xe3\x6c\xa0\x68\xdb\x59\xdd\x08\xf6\x1f\xee\x3f\xa8\x1d\xe7\x96\x9d\x35\x8d\x38\x8b\xec\x0a\x07\xf7\x9b\x86\xc3\x0d\x47\x4b\x9e\x29\xf0\x38\x73\xf8\x69\x67\x77\x21\xa7\x67\x8d\x5e\x96\x1a\xd2\x2d\x9b\x3b\x89\x57\x92\x87\xb4\xa7\x1b\x79\x92\x4f\x86\x5b\xda\x41\x23\x65\xa6\xb3\xc6\x60\x9c\x05\x8f\xe6\x8a\x96\x9e\xd3\x0e\x99\x56\xde\x8e\x33\xa3\xe0\x6e\xed\x44\xac\x97\xa9\x65\xc2\xe8\x8e\x84\x13\x89\x08\xdb\xba\x95\xc9\x65\x78\x73\x22\x12\x61\xfb\x58\x98\xc5\xdd\x7d\x5f\x22\x7a\x3b\xb4\x37\xd1\x23\x97\xdb\xdd\xb4\xbd\x91\x30\x3b\x7b\x56\x1d\xe9\x4b\x44\xc3\x8f\xb3\x7d\x2c\xc2\xe2\x2c\xda\x2b\x5f\xb3\x28\x51\x85\xa6\xc5\x18\xcd\xd1\x14\x95\x17\x7b\xc4\xf8\x62\x8f\xfb\xd6\x17\x1d\x5a\xec\x12\xe3\x4b\x5c\xa0\x6b\x74\x8b\x66\xa9\x42\x73\xee\xaf\x4a\x15\x26\x46\xa8\x42\x37\xa9\x4a\xb3\x54\xa3\x49\x46\x55\x17\xae\x4a\xd3\x54\x16\x6f\x51\x95\xaa\x8c\x6e\x51\x99\x26\xa9\x22\x46\xc4\x45\xaa\xd2\x0c\xd5\x44\x49\x5c\x66\xa2\xc4\xa8\xe6\x7a\xae\xd3\x34\xd5\xe8\x73\xf9\x8c\xe8\x13\x51\xa2\x79\xaa\x88\x8b\x34\x47\x35\x89\xd6\x74\x9f\x18\x6f\xba\x4d\x8c\xb7\xb8\xab\x15\x41\x99\x25\x46\xe5\xf5\x34\x49\x65\x79\xfd\x9d\xa7\xb9\xf8\xce\x65\xc8\x7e\xac\xa4\x15\x97\x14\xd1\x29\x2a\x8b\x51\x2a\x33\x49\x66\xd2\xad\x74\x56\xbc\xc3\x68\x8e\xca\x4c\x94\xc4\x05\xaa\x88\x51\x71\xc1\x45\xa8\x2c\x39\x3b\x4f\x35\x89\x2c\x8b\x90\x17\xd3\x8c\xc4\x7f\xb7\x15\x20\xe8\x03\x9a\xa7\xaa\x28\x51\x59\x29\x04\xfa\xf0\xb6\xa8\x2d\x59\x4e\x50\x8d\x66\xc4\x98\x24\x70\x67\x7a\xbe\xf4\x82\x3b\xd2\xb3\x6b\x0b\x31\x5b\xe2\xcd\x4b\x04\x9a\x91\x14\xe6\xa8\x4c\x37\x5d\x01\xab\xb2\x0b\xd3\xb2\xd3\x34\x4f\xd3\x6a\xca\xfe\x17\x4a\x5c\x91\xd3\x20\x46\xe8\x86\xda\x8e\xd6\xfb\x55\xa1\x1b\x54\x93\x8e\xb2\x18\xa1\xaa\x38\x47\xe5\x95\x71\x9b\x93\xca\x37\x60\xaa\xe2\x7c\x33\x88\x52\x48\x0e\xca\x18\xcd\xca\xe7\xb1\x32\xd0\x59\xaa\x89\x4b\xb2\x34\x51\x92\x95\xee\xac\x0f\x8a\xea\x85\x3b\x3a\xee\x5c\x4a\xc0\x89\xdb\xea\xd0\x94\x3b\xea\x15\x29\xf3\x28\x4d\xd3\x8c\x78\xbb\xd1\xbd\x96\x79\x93\x4b\xb3\xea\xd3\x39\x2f\x46\xa8\xe6\xbe\xa1\x49\xaa\xd1\x94\xcb\x64\x52\xaa\x32\xb1\x78\xa8\xc5\x65\xd0\x47\x74\x5d\xe6\x8b\x12\x55\xc5\xa8\x72\xd4\xc4\x79\x31\xa2\xc6\x59\xb9\xea\x07\xc5\xb8\x44\x92\x4e\xf9\xdf\xd6\xda\x1f\x5d\xc6\x1f\x5b\xc6\xdf\xd3\xda\x5f\x57\x7e\xac\x75\xdc\x93\x71\xb9\xe8\x82\x46\x2c\x57\x43\xf3\x2b\xbe\x2a\x27\xb4\x49\xcf\x7f\x03\x00\x00\xff\xff\x71\x89\x83\xc1\xc8\x0c\x00\x00"

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
