// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package nut

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x4b\x6f\xdb\x38\x1e\xbf\xe7\x53\xfc\x97\x9b\x5d\xcb\xd8\x4a\x72\xb2\x40\xb1\x70\x25\xef\xa1\x53\xa0\x87\x4e\xa7\x68\xd2\xce\xa1\x28\x0a\x5a\xa4\x2d\xa6\x14\xa9\x21\x29\xd9\x86\xe1\xef\x3e\xa0\x5e\x96\x6d\x59\x76\x9a\xa4\xed\x00\xa3\x83\x4d\x89\xff\xf7\xeb\x47\x69\xbd\x06\x42\x67\x4c\x50\x40\x91\x14\x86\x0a\x83\x60\xb3\xb9\x58\xaf\x81\xcd\x60\x4e\xc1\xe1\x54\x80\x97\x63\xc5\xf0\x94\x53\x3d\x84\x91\xdd\x0e\x08\xcb\x21\xe2\x58\xeb\x10\x29\xb9\x40\x93\x0b\x00\x80\xf6\xd3\x48\x72\x37\x21\xee\xd5\x35\xd8\x95\x4e\xea\xd5\x52\xbb\x57\xd7\x15\xfd\x3e\xcf\xf2\x4b\x8a\x05\xe5\xad\xdd\x43\x0a\xc3\x0c\xa7\x7b\x14\x05\x55\x7c\x3d\xb1\x46\x5f\xfd\x4f\x00\x7a\x4b\xcd\x42\xaa\xaf\xf0\xe1\xdd\x0d\xdc\x4a\xc9\x35\x02\x0f\x36\x9b\xc0\x8f\xaf\x3b\x38\xdb\x56\x73\x8a\xd5\x8c\x2d\xd1\x24\xf0\x09\xcb\xf7\x0c\xe9\x78\xb4\x63\x5b\x1d\xc0\x7e\x1d\xdb\x78\xf5\x5a\x52\xc6\xef\x79\x1d\xbe\xe7\x75\xf4\x9e\x1f\x61\x2e\x04\xa4\xdb\x18\xbc\x8c\xb1\x9a\x53\x52\x7b\x9e\xf6\x70\xb5\xd4\xa6\x4a\xce\x15\xd5\x1a\x81\x36\x2b\x4e\x43\x94\x60\x35\x67\xc2\x9d\x4a\x63\x64\x32\x1e\xf5\x28\x3f\x26\xca\x9d\x62\x05\xed\x1b\x57\x67\x51\x44\xb5\x2e\x6b\xcc\x8b\x4a\x43\x61\xb3\xd9\xa3\x32\x8a\xa5\x94\x00\x8e\x0c\xcb\xe9\x7a\x0d\x54\x58\x22\xd4\x6b\x40\x71\x29\x69\x2d\xaf\x85\x4d\xb1\x42\x60\xeb\xd7\xcd\x31\xcf\x68\xc2\x44\x88\x46\x3b\x4f\xf0\x32\x44\x57\xa3\xd1\x19\x92\xb7\x4c\x42\x2e\x42\xb4\x5e\x83\xc3\x04\xa1\xcb\x56\x87\x00\x9a\x62\x63\xa8\x5a\x55\x8e\xa1\xa1\xf7\xd1\x32\x9c\x67\x7a\x15\xf5\x05\x23\x26\x1e\xdf\x53\xfc\xbf\x4e\x24\xa7\x48\x90\x4e\xb1\xa8\xb5\x70\x26\xa8\x1b\x53\x36\x8f\xcd\xf8\x7a\x94\x2e\xd1\xe4\xbe\x1a\x03\xdf\xca\x3b\x51\x13\x87\x6d\x73\xe6\x76\xdf\xd6\x63\xb6\xca\x1b\x89\x7f\xca\x3e\x59\x60\x25\x98\x98\xf7\xb5\xc5\xcf\xdf\x0d\x59\xaa\x3d\x6e\x03\xfc\xd8\x7d\xd0\x21\xf8\xc9\x3a\xa0\x4b\xd7\x0f\xa9\xfd\x63\x8f\xf7\xf0\x65\xaf\x2a\x8d\x4c\xc7\x70\xe5\x8d\x68\xf2\xe2\x7b\x40\xcf\x8d\xc1\x26\xd3\xe7\x74\x54\x3a\x09\xb4\x51\x52\xcc\x7b\xe2\xae\x4b\x69\xdb\xc8\x07\x7e\xc5\x73\x54\xf6\xf7\x1a\x1c\xef\x33\x61\x58\x42\xef\xed\xa9\xa7\x4a\xc6\x87\x38\x73\xfa\x80\xd2\xba\xad\x96\xd5\x5f\x83\xa5\x17\x3f\xf4\x9c\x67\xe9\x38\x4e\x35\x25\x0f\x3f\xf1\x7d\xac\x0d\xef\x3d\xe9\x65\xbc\x16\x2b\x70\x0e\x02\xe7\x76\xa0\x2a\xdb\xfd\x50\x58\xf4\xc5\x48\xc9\xa7\x72\x79\xac\x49\x38\x9b\x04\xb8\xe5\x79\x61\xbd\xcb\x99\xf8\x8a\x26\x01\xab\x37\x66\x18\x66\xd8\x8d\x62\x9a\x2b\x29\xdc\x2c\xb5\x87\x49\x36\x09\x7c\x3c\x09\x7c\xce\xba\x32\x99\xf1\x1f\x74\x2c\x9d\x49\x95\x54\x00\x61\x97\xa8\x40\x15\x29\x42\xf4\xff\x6a\x51\x17\x04\x82\x84\x9a\x58\x92\x10\xa5\x52\x1b\x04\x8c\x84\x28\xdf\xc6\x9c\x60\x83\x5d\x23\xe7\x73\x2b\x29\xc7\x9c\x11\x6c\xa4\x3a\x16\x46\x63\xb9\x6a\xfb\xca\x9b\xe2\xb7\xc1\x36\x62\x5c\x45\x75\x2a\x85\x66\x39\x05\x21\x17\x0a\xa7\x68\x17\x13\xae\x46\xa3\xbe\x69\x1f\x98\x98\x62\xd2\xb7\xaf\x4e\x4c\x6e\x13\x1f\x56\x57\x5d\x5c\x26\xbe\x17\x33\xcf\xbe\x8d\xf3\x17\xaa\x23\xc5\x52\x9b\x88\x73\xf8\x03\xbf\xcf\x29\xcb\x7b\x22\x24\x53\x49\x56\xc7\xf7\xd7\x6b\x50\x58\xcc\x29\x5c\xb2\x67\x70\x59\x27\x1f\xc6\x61\x7b\x62\x6f\x36\x0f\x09\x39\xb1\xbe\x37\xa2\xbd\xb7\xb8\x1a\x91\xa6\xc7\xec\x9a\xf3\x24\xf0\x97\xa3\x6e\x2b\xfd\x76\x95\x52\xef\x77\xc5\x0c\x2d\xfc\xe8\xb1\x7c\x47\x15\x13\x69\x66\xc0\xac\x52\x1a\x22\x43\x97\x06\x35\x6d\x2f\x55\xe2\xda\x5e\x53\x92\x23\x10\x38\xa1\xc5\x91\xe8\xc0\x1d\x04\xc5\x91\x69\x6f\xb3\x39\x1b\x15\xad\xd5\xc9\xe7\x9f\xe5\x23\xe5\xfa\x6c\x67\x3a\x2d\x38\x4b\x49\x09\x1f\xbd\x81\x3a\x2b\x6b\x75\xad\x6f\xcd\x68\x15\x3d\x5c\x9e\xcc\x7e\x7f\xcd\x9f\xb6\x34\xf0\x7b\xaa\x3e\xf0\x8b\xb1\xd4\x35\xb1\x6d\xb2\x1f\x1b\x7c\x23\x99\x24\x58\x90\xbf\x1c\xf6\xbe\xac\xec\xfe\x1b\x7a\xef\x03\xbd\x4f\x0c\x82\x7d\x00\xd8\x3b\x89\x77\x10\xa8\x4a\xed\x59\xe8\xd3\xe6\x7b\xb5\xa4\x51\x66\xce\x42\xbd\x6f\x45\xbc\xe3\x9d\xdf\x8b\x74\x7d\x28\xb7\x8b\x70\x51\x42\x0a\x70\xab\xdb\xf2\xd8\x18\x39\x11\xcd\x12\xd3\xa2\x84\x9c\x07\x67\x27\xa1\x2c\x98\x66\xc6\x48\x51\x21\x50\x79\xd3\x60\xd0\xd4\x08\x98\x1a\xe1\x12\xeb\x86\x2a\x96\x3a\x41\x20\x45\xc4\x59\xf4\x35\x44\xb4\xcc\xcb\xcb\x84\x38\x83\x3d\xb3\x06\xc3\x33\x5e\x9e\x0f\x13\x7c\x79\x1a\x07\x4a\x1b\x7b\x67\xf8\x89\x80\x34\x28\x61\xed\x3d\x1f\x20\x8e\x97\x48\x3f\x30\x1c\x05\x85\x4e\x40\xf8\xb6\xc1\xdf\x5a\xd6\xdf\xdf\xef\x74\xf1\xe9\x3d\x28\xdd\xab\x12\x8c\xd3\x94\xb3\x08\x5b\x6f\xfd\x3b\x9c\xe3\x72\xb3\x4a\xd5\x7a\x0d\xf6\xed\x98\x45\xaf\x6f\x7f\x7d\x03\x4e\xb9\xfe\xf0\xfe\x0d\x20\x9f\x60\x1d\x4f\x25\x56\xc4\xc7\x5a\x53\xa3\xfd\x9c\x0a\x22\x95\xf6\xed\x51\xbd\xf0\x42\x7b\x82\x1a\xff\x4e\xfb\x77\x7f\x64\x54\xad\x3c\xbb\x71\x5b\x6e\x24\x4c\x78\xd6\x9a\x19\xe6\x9a\x0e\xeb\x30\x3d\x5c\x9b\x3b\xd5\x56\x61\x4b\xd3\x54\x4a\xa3\x8d\xc2\xe9\xd3\xe9\xdc\x0e\xd1\x3d\xdd\xdb\x8d\xef\xa1\xbc\xf2\xbd\xa5\xf4\x29\x7c\x6f\xde\xbd\xac\xae\xe6\xa6\x43\x41\xa1\xe1\xd2\x21\x32\xca\x12\x2a\xcc\xd0\x53\x14\x93\x95\x33\xcb\x44\xf1\xe2\x07\xce\x10\xd6\x4d\x4d\x37\x4f\xab\x2f\xf7\xaf\xb1\x20\x9c\x2a\x47\xb5\x89\xec\xc5\x66\xe0\x28\x1b\xd8\x8c\x1b\x08\xc3\x10\x06\x33\xcc\x38\x25\x83\x7d\x42\x7b\x09\xba\x80\x77\x6f\xa5\x61\xb3\x95\x73\xb8\x6b\xaf\xe2\xf8\x31\x86\xc1\x2b\xa5\xa4\x1a\x3c\xeb\xa6\xa1\x4b\x33\x06\xe5\x25\x54\x6b\x3c\xa7\x47\x88\x56\xa9\x95\x43\x7b\xe4\xc4\x8c\xd0\x71\x19\xa0\x6e\x02\x8b\xbc\x4c\xcc\xc7\x30\x68\x12\xf7\xdf\xc1\x01\xe5\x66\xf8\x62\xe7\xd9\xa6\x3c\x90\x97\x81\xa9\x6c\x84\x7f\xd8\xc8\x64\xa2\x6c\xfc\x87\x06\xe7\xa6\xcc\xc9\x23\x84\x47\xf7\x4a\x7a\xaa\x00\x5d\x6c\x57\xcd\x72\xc1\x04\x91\x0b\x6f\x0b\x57\x10\x36\x45\xe8\x44\x09\xd9\x8f\xd8\xa5\x87\xef\xf0\xb2\x23\x50\x95\x67\xef\x7e\xbb\xb9\xed\x70\x2b\x53\x7c\x0c\x83\xfa\x5b\x47\x94\x90\x7f\x47\x09\x09\x07\xf0\x1f\x88\x12\x72\x48\x5e\xc5\x67\xbc\xd7\x06\x17\xc7\xdc\xdb\xbc\xd8\x7a\x74\xe9\x0c\xfe\xb9\x7d\x4d\x2e\x5e\x22\x3f\x31\xf2\x79\x30\xf4\xa2\xd8\x82\x75\xd3\x79\xce\xbe\x6f\x39\x56\x07\x96\xcc\x20\xdc\x95\x38\x18\x1e\x9a\x4b\x0b\x22\x13\x33\xdd\xb1\x69\xe7\x14\x84\xb0\xde\x45\xc0\x96\xc1\x35\xd1\x27\xea\xa5\x4a\xa6\xce\x80\x91\xc1\xf0\x33\x84\x40\xbd\x1c\x73\x67\xb8\x47\x7a\x22\x03\xb3\x4a\x48\xf9\x15\xa9\xcb\xda\x22\x17\x35\x59\x99\x91\x2e\x32\x6b\xd2\xb8\xf8\x7d\x84\xfc\x54\x6b\xfb\x1f\xf8\x25\xb2\xb6\x70\xfa\xcf\x00\x00\x00\xff\xff\xb1\x17\x3b\xed\x1b\x1f\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x4f\x4d\x4f\x14\x59\x14\x3d\x45\x31\x03\xd3\xc3\x4c\x94\x98\xb8\x71\xf1\x8c\x91\xe8\xa2\xb0\x0b\x5d\x68\x41\x81\x8a\x90\xa8\x74\x62\xb0\xed\x9d\x8b\x67\xf7\xb3\xa9\xd8\xfd\xaa\x53\x1f\x04\x23\x8b\xa6\x31\x9a\x88\x09\x06\x63\x8c\x26\x6a\x8c\xfe\x80\x32\xd2\x8a\x4a\x37\x7f\xe1\xbe\x3f\xe0\xc6\xad\x4b\x7f\x80\xa9\x0f\x40\x42\x2d\xea\x9e\x73\xee\xb9\xef\x9e\xfb\x7d\xb0\xf7\x19\x00\xfc\x07\xe0\x08\x80\xfb\x00\x0e\x02\xf8\x89\xf4\xcb\x6b\x40\x1f\x00\x33\xab\xe7\x34\xa0\x1f\xc0\x05\x0d\xf8\x17\xc0\x15\x0d\x38\x0c\xa0\x94\xf5\xa5\x06\xf4\x02\x98\xcf\xf8\x3d\x0d\xf8\x1b\xc0\xb2\x06\xfc\x05\xe0\x51\x36\xff\x44\x03\xfe\x01\xf0\x42\x03\x26\x34\xe0\x9d\x06\x1c\x00\xf0\xa3\x07\xf8\x1f\xc0\xaf\xac\xea\x7a\xaa\x0f\xea\xc0\x10\x80\x63\x3a\x30\x18\xef\xcd\xf4\x9b\x7a\xba\x5f\xea\xc0\x00\x80\x07\x99\xfe\x54\x07\x0e\x01\x78\x9f\xd5\x2f\x3a\xa0\x21\xcd\xd0\x97\xdd\xd6\x9f\xdd\xad\x23\xcd\x38\x80\xdd\x2f\xce\x96\x43\x7a\x63\x4f\xa6\xc5\x77\x61\x72\x8e\x7b\x55\x51\xc1\xa4\x5b\xaf\x73\xb9\x53\x7d\x5c\x12\x7e\xd9\x73\x1a\x81\xe3\xca\x3f\x31\x0b\x25\x9f\xe7\x4e\x8d\xdf\xaa\x09\x4c\x2d\x88\x72\x18\x08\xcc\xb8\xbc\x82\xd9\x50\x06\x4e\x5d\xe0\x7a\xc0\x83\xd0\x47\x89\xd7\x42\x81\x12\xf7\x9c\xc4\xba\x0d\x7c\xcc\x8a\x86\xeb\x05\x46\xc1\xaf\x3a\x15\xe3\x62\x58\xf5\x8d\xa2\x6b\xb1\x8a\x98\x3f\x7f\xc7\x99\xe3\x75\x77\xd8\x0b\x73\x33\xdc\x0f\x8c\xa2\xc7\xa5\x5f\xe3\x81\xeb\x59\xec\x6a\xd2\x62\x85\xd0\xe3\x75\xb7\xe2\xb2\xb1\x3d\xfe\xf1\xdc\x0c\x97\xd5\x90\x57\x85\x51\x14\xbc\x6e\xb1\x1d\x6e\xb1\xd9\xd0\xf7\x1d\x2e\x73\x85\xcb\x85\x29\xa3\x24\x3c\xdf\x71\xa5\xc5\xcc\xe1\x7c\x6e\xd2\x95\x81\x90\x81\x51\xbc\xdb\x10\x16\x0b\xc4\x42\x70\xaa\x51\xe3\x8e\x1c\x65\xe5\x39\xee\xf9\x22\xb0\x6f\x14\xa7\x8d\xb3\xbb\xbe\x38\xcf\x6d\xe1\x19\x53\xb2\xec\x56\x1c\x59\xb5\x58\xee\x5a\x2d\xf4\x78\xcd\x98\x76\xbd\xba\x6f\x31\xd9\x48\xa8\x6f\x9f\x1e\x65\x29\xb4\xe5\x71\x33\x6f\xdb\x26\x1b\x1a\x62\x31\xcc\x1f\xb5\x4d\x93\x4d\xb0\x3c\xb3\x12\x3e\x6e\x8f\x6c\xb7\xc6\xec\x33\x31\x3c\x91\xd8\xc6\xcc\x3c\x5b\x5c\x4c\x47\xc6\xed\x91\xfc\x49\x36\xc1\x4c\x66\xb1\x91\x51\xd0\x73\x8a\x54\x53\xad\xd2\x27\x6a\x53\x87\xba\xa0\x97\xd4\xa5\x4d\x8a\xa8\x43\xeb\x14\xed\xa1\x6a\x05\xf4\x9a\xb6\x68\x43\x2d\x25\xc2\x06\xb5\xf7\x09\x8c\x3a\xf1\x6f\x9d\xba\x6a\x49\xb5\xd4\x32\x6d\xa5\xaf\xae\xa9\x15\xda\xa2\x2e\x7d\x8b\x6d\xaa\xa5\x1e\x83\x5e\x51\x44\x1f\x55\x53\x2d\xd3\x67\xfa\x1a\xaf\x5a\x53\x4d\x6a\xd3\xa6\x5a\x65\xaa\x49\x11\x7d\xa0\xae\x6a\xc5\x3b\xdf\xaa\x16\x45\xf1\x63\x6a\x29\xce\xdb\xa1\x48\x3d\x4c\xd2\x26\x01\xde\x50\x3b\x9d\x4b\xa4\xb8\xb9\xba\x4f\x54\x2b\xd4\xc6\xef\x00\x00\x00\xff\xff\x22\xa4\x74\x8a\xce\x03\x00\x00"

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
