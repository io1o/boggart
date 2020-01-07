// Code generated by go-bindata.
// sources:
// templates/layouts/widget.html
// templates/views/account.html
// templates/views/default.html
// templates/views/files.html
// templates/views/logs.html
// templates/views/user.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package xmeye

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

var _templatesLayoutsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x95\x41\x73\xd3\x30\x10\x85\xef\xf9\x15\x3b\x3a\xc1\xc1\x31\xe9\x89\x81\x24\x0c\x07\x38\xf5\xc0\x94\xe1\xdc\x59\xcb\xeb\x44\x20\x4b\xae\x24\x9b\x30\x1e\xff\x77\x46\x56\xec\x3a\x41\x2e\x6d\xda\xe6\x92\x1d\xe9\xe9\xad\x3e\xfb\x49\x6e\x5b\x70\x54\x56\x12\x1d\x01\xcb\xd0\x12\x83\x25\x74\xdd\x62\xd1\xb6\x90\x53\x21\x14\x01\xe3\x5a\x39\x52\x8e\xf9\xf1\x75\x2e\x1a\xe0\x12\xad\xdd\xb0\x0a\x77\x94\x38\xe1\x24\xb1\xed\x02\x00\x60\x3a\x59\x68\x53\x26\x3b\xa3\xeb\x0a\xaa\x5a\xca\xc4\x88\xdd\xde\x1d\x75\xe7\x5a\xa1\xaa\xda\x05\xf1\x44\xd1\xab\x24\x66\x24\x07\x5d\xe6\x54\x54\xd5\x2b\x11\xf6\x86\x8a\x0d\xfb\x84\xdc\x09\xad\x36\x5c\xab\x42\xec\x6c\x42\x87\x4a\x1b\xc7\x26\x1e\xe0\x7d\x04\xd7\xa1\xb0\x35\xe7\x64\x6d\xa8\xcb\x88\x73\xef\x2e\x46\x2e\xb4\x50\x60\xc2\xf5\x8e\x41\xcf\xbe\x61\x6d\x0b\x62\xf5\x5e\x01\xfb\xd2\xb7\x82\xd0\xb9\x36\xe8\xf7\x11\x9e\x27\xdb\xae\x53\x11\xd9\x74\x8a\x8f\x20\x91\xfa\xd5\x30\xf6\x84\xb9\x25\x37\x8b\xe2\x5b\x3f\x8d\x60\x9d\xf6\xaf\x6c\xf2\xa2\xd3\x5c\x34\xc7\x7c\x84\xf2\xf8\x77\x92\x25\x2e\x09\x4d\x21\x0e\xbe\xcd\xbf\xb3\x46\xff\x8e\x44\x8c\x6b\x99\x94\x79\xb2\xba\x02\x5f\xd9\x72\xa8\x0e\x36\x59\x5d\xcd\x44\xed\x70\x5b\xa1\x22\x79\x1e\xb3\x13\xc5\x90\xf7\x08\xad\xd7\x19\xed\x9f\x94\xc3\x2c\xe6\x34\x2a\xeb\x31\xb5\x0a\x1b\x50\xd8\x24\x0e\x33\x0b\x19\x9a\x5b\x5f\xb0\x7b\x1b\x29\x6c\xac\xd7\xe8\x24\xc5\x51\x5b\x19\xb2\xa4\x5c\x88\x95\x7f\x53\x05\xd0\x1d\x2c\x43\x4a\x80\xe5\x54\x60\x2d\xfb\x63\x3a\xb4\xf6\x53\x0d\x79\x2d\xa9\x1c\xba\x6e\x3b\x66\xeb\xc1\xe5\x3f\xb1\x41\xcb\x8d\xa8\xdc\x87\x46\x8b\xfc\xcd\xbb\xb7\x1f\xbd\x85\xb4\x04\x5d\xd7\xb6\xb0\xbc\xa1\xbb\x9a\xac\x5b\xfe\xb8\xb9\x5e\x7e\x43\xb7\x0f\xc3\xa1\x07\xdb\x8e\x21\xfa\xfe\xc7\x3a\x2a\x43\x7e\x7c\x50\xd6\xa9\x8c\x44\xe8\xc9\xa0\x21\x94\x97\x51\x0e\x6b\x2f\x41\x9c\x9e\xc7\x18\xee\xf5\x78\x58\x5e\x10\xb6\x10\x92\x2e\xa6\x1d\x17\x3f\x07\xb7\x37\x89\xf1\x7e\x0d\xee\x2f\x0c\x8c\x9c\xeb\x5a\x5d\x1c\xe3\xc9\xf2\xe7\x40\x1f\x6d\x62\xd8\x9f\x87\x0e\xff\x07\x5f\xa7\xb5\x9c\x99\x99\xdc\x38\x0e\xb3\x64\xfe\xce\x39\x59\x71\x76\xf7\x4c\x1d\xfc\x08\x14\x98\x13\x84\xc7\x05\x42\x3d\xe0\xe6\x7f\x6d\x0b\x99\xd4\xfc\x17\xf4\x5f\xf2\xc0\x03\x23\xf1\xfc\x46\xee\x2f\xf4\x47\x4e\x45\x86\xcf\x86\x66\x3f\x12\xe3\x76\xfe\x06\x00\x00\xff\xff\xeb\xdf\x05\x03\xa9\x08\x00\x00"

func templatesLayoutsWidgetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesLayoutsWidgetHtml,
		"templates/layouts/widget.html",
	)
}

func templatesLayoutsWidgetHtml() (*asset, error) {
	bytes, err := templatesLayoutsWidgetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/layouts/widget.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesViewsAccountHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\xcf\x8e\x9b\x30\x10\xc6\xef\xfb\x14\x23\x2b\xed\xa9\xc0\xee\xad\x5a\x85\x54\xd5\xaa\xea\xa5\x5d\x55\x91\xf6\xd0\xa3\x83\x87\x60\x09\x6c\xd6\x1e\xb2\x1b\x21\xbf\x7b\x65\x93\x10\x87\x25\x69\x2e\x68\x6c\xff\xfc\xe7\x9b\x19\xbe\xbe\x07\x81\xa5\x54\x08\xac\xe5\x5b\x64\xe0\xdc\x1d\x00\xc0\x52\xc8\x1d\x14\x35\xb7\x36\x67\xa5\x36\x4d\xb2\x35\xba\x6b\xa1\xed\xea\x3a\x31\x72\x5b\x11\x5b\x05\x6e\xca\x4a\xd5\x76\x34\xc0\x11\x11\xa8\x9a\x6f\xb0\x3e\x72\x1b\x52\xb3\x54\x20\x39\x54\x06\xcb\x9c\x7d\xe3\x05\x49\xad\xf2\xce\xa2\x61\xd1\x4e\xf0\xbb\x65\xa1\x87\xc0\x76\x45\x81\xd6\x0e\x71\x33\x73\x5e\x38\x53\x8e\x6a\xb8\x85\x92\x27\x6d\xdd\x59\x06\x24\xa9\xc6\x9c\xf5\x3d\xc8\x87\xaf\x0a\xd8\x93\x41\x4e\x08\xc3\x85\x29\x38\xc7\x56\xcb\x4c\xce\x3c\x31\xe3\x13\x75\x59\x90\x17\x25\x25\x13\x72\x37\x0c\x0f\xe1\x10\x13\xdf\xd4\x78\x7c\xcb\x30\x08\xdf\xc4\x92\x91\x2d\x0a\x10\x9c\xf8\x30\x2f\x28\x31\x68\x5b\xad\xac\xdc\x21\x28\xfd\x66\x78\xcb\xc0\xd2\xde\x3f\xf9\x4d\x0a\xaa\x1e\x1f\xee\xef\x3f\xc5\x95\xa0\x0a\xb9\x88\xc7\x66\xf2\x4c\xaa\x56\xa3\xd8\x97\x51\xe5\x32\xa3\xea\x0a\xf8\x33\x14\xea\x16\xf2\x49\x37\x0d\x2a\xba\x89\x5d\xa3\x45\xb3\x43\x71\x13\xfc\x3d\x74\x82\x9d\x63\x97\x59\xac\xd2\xaf\x4d\x72\xb0\xd1\x62\x7f\x1a\xf7\x3d\x18\xae\xb6\x08\x0b\xf9\x05\x16\xbe\xd2\xf0\x98\x43\xea\x03\x7b\xec\xfe\x43\xee\xfc\xed\x25\xe0\xeb\x80\xa5\xcf\xbc\x41\x58\xa4\x45\x67\x0c\x2a\x02\xe7\x8e\x65\x3c\x74\xa0\x6f\x23\x54\x02\x9c\x9b\x0a\x11\x5e\x48\x74\x48\x90\x20\xae\x50\x21\xe3\xff\xc7\x7e\x63\xa3\xaf\x51\xb2\x3c\x80\xc7\x5c\x83\x73\x63\x4a\xff\xa2\x65\xb0\x18\x66\xb0\xb6\x18\xaf\x3d\xeb\xd3\x52\x50\x34\x7b\xc5\xd8\xc6\xf8\x4e\x49\x81\x8a\xd0\xcc\xfd\xce\x91\x3d\x9c\x7e\x7b\x30\xda\xb7\xf1\x25\x0f\x80\xd8\x07\xbc\xde\x74\x8d\xaf\x1d\x5a\x4a\x5f\xd6\xbf\xd2\x3f\x9c\x2a\x70\x2e\xf6\x87\xcf\xfe\xa3\x78\x83\xf9\x34\xd5\x1f\x9d\x43\x95\xfa\xdc\x42\xde\xed\x85\x27\xc0\x99\x75\x6c\xeb\x7d\x5b\x85\x4d\x63\x94\xa0\x90\xf4\xd1\x47\x7e\x84\xd9\xc5\x65\x03\x81\x39\x13\x81\x73\xdf\x38\x4d\xc5\xb9\x3f\x6f\xf7\xb1\x40\xf1\xfa\xa9\xe1\x97\x59\xb0\x92\xd5\xdd\xc8\xfd\x0b\x00\x00\xff\xff\x8a\x42\xd0\xe8\xf1\x05\x00\x00"

func templatesViewsAccountHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsAccountHtml,
		"templates/views/account.html",
	)
}

func templatesViewsAccountHtml() (*asset, error) {
	bytes, err := templatesViewsAccountHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/account.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesViewsDefaultHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6d\x6f\xdb\xc8\x11\xfe\x9e\x5f\x31\xe5\xb9\xa0\x04\x58\xa4\xec\xdc\x5d\x0b\xd6\x72\x71\x89\x53\xa4\xc0\x25\x31\x62\x27\x45\xfa\x25\x58\x71\x47\xd2\x3a\xe4\x2e\xb1\x3b\x94\x2c\x08\xfa\xef\xc5\x92\x94\xcc\x37\xc9\xb2\xcc\x14\x87\x43\xf4\x89\x5a\xce\xcc\x3e\xf3\xcc\xcc\xbe\x0c\x57\x2b\xe0\x38\x11\x12\xc1\x49\xd8\x14\x1d\x58\xaf\x5f\x00\x00\x5c\x4c\x94\x8e\x41\xab\x08\x47\x8e\x7d\x74\x80\x85\x24\x94\x1c\x39\xff\x74\x20\x46\x9a\x29\x3e\x72\x12\x65\xc8\x01\xc1\x47\x8e\x41\x22\x21\xa7\xc6\x01\xce\x88\x0d\x48\x4d\xa7\x56\x73\xce\x22\xc1\x19\x29\xed\x5c\x66\x56\x33\xcb\xc4\xc6\x11\x42\x18\x31\x63\x46\x4e\xfe\x87\xd3\x40\xa3\x49\x94\x34\x62\x8e\x20\xd5\x42\xb3\xc4\x01\x43\x4b\x6b\x64\x21\x38\xcd\x82\xb3\xe1\xf0\xaf\x25\x2b\xb9\xa5\x19\x32\x5e\x1f\xd3\xd5\x81\x42\x70\x33\x5f\xcc\x07\xa1\x8a\x06\x67\xce\xe5\x6a\x05\xe2\xec\xef\x12\x9c\x0f\x89\x75\xcc\x01\x0f\xd6\xeb\x0b\x9f\x66\xad\xfa\x0f\xe2\x9f\x59\x94\xe2\x2e\xe9\x0b\xbf\x3e\xbf\x95\x69\x41\x39\x56\x7c\x59\x1d\xb3\x13\x4c\xc0\x33\x4b\x43\x18\x7f\x15\x72\xa2\x36\xb1\x78\xc4\x37\xfe\x80\xed\x06\xb5\x60\x11\xc8\x34\x1e\xa3\xde\x62\xe4\x3b\xb5\xca\xb3\x79\xb9\xf2\x7b\xd5\xaa\xd4\xe2\xd8\xa3\x68\x3e\x25\x24\x62\x7c\x2a\x8c\x2b\x9c\x8b\x10\x3f\xa6\xf2\x56\xc4\x08\xeb\x35\x6c\xed\xc5\x42\xa6\x84\x66\x97\xc1\x63\x20\xde\xa8\x09\x2d\x98\x46\x98\xa3\x36\xe5\x2c\x38\x94\x33\x35\xa1\xff\x30\x8d\x9f\x73\xf5\xce\x70\xbd\x65\x9a\x5b\x5c\x4f\xc5\x63\xf5\x2c\x9e\xce\x81\x1c\x4b\xd0\x06\x50\xd7\x04\xbd\x91\xa1\x5e\x26\x74\x2c\xac\x42\xbd\x6b\x54\xaf\x52\x11\x71\x38\x24\xeb\x1b\x83\xd9\x0b\x13\x6a\x91\x10\xd0\x32\xc1\x91\xc3\x92\x24\x12\x21\xb3\x8b\x93\x7f\xc7\xe6\x2c\x7f\xe9\x5c\x72\x15\xa6\x31\x4a\xf2\x16\x5a\x10\xf6\x38\x23\xbc\x55\x37\xa4\x85\x9c\xf6\xdc\xba\xa3\x19\x22\x5b\x49\xde\xbf\x94\x8e\x19\x81\x73\x3e\x1c\xfe\x3a\x18\x9e\x0d\x86\xe7\xb7\x67\xbf\x04\xc3\x9f\x83\xe1\x2f\xff\x1d\xfe\x2d\x18\x0e\xed\xda\xef\xf6\xfb\x17\x7e\x3e\x51\x0b\xee\x8e\x56\x05\x8b\xf8\x0f\x45\x52\x06\x89\xfd\xa1\x58\xfa\x2d\x62\x3a\x06\x21\x21\x9c\x31\x29\x31\x7a\x6a\x82\x67\xfa\xff\x96\xaf\x73\xed\xce\x12\x3c\x87\xa5\x52\x7a\x16\xae\x0f\x29\x75\x0d\xec\x96\x45\xdf\x9e\x41\x97\x55\xef\x9e\xad\x0c\xd4\x33\xc8\xb2\xfa\xdf\x81\xab\xcf\x82\xa3\x7a\x06\x59\x99\x7e\xf7\x6c\xe5\xb0\x9e\x41\x57\x66\xe0\x3b\xf0\xf5\x5b\xca\xc5\x73\xf8\xca\xf4\xbb\xe7\xeb\x4a\x4c\x8f\x45\x74\x25\xa6\x5d\xa3\x79\x73\x4f\x9a\x1d\x8b\x27\x53\x7e\x1a\xa2\xd5\x0a\x50\x72\x7b\x44\x6e\x3b\x45\xdb\xdd\xe5\x6b\x98\x6a\x8d\x92\x0e\x3d\x46\x83\x56\x0b\x93\x30\x39\x72\xce\x4b\xb7\x83\xdb\xa3\xf7\x29\x2e\xe6\x9b\x5b\x87\xbd\x72\xdc\x9b\xc1\xcf\x4e\xbb\x68\x26\x2e\x64\x92\x6e\x76\x35\xc2\x7b\x72\x36\xca\xf6\xfa\x35\x08\x95\x24\xad\x22\x07\x24\x8b\xad\x80\x88\x71\x50\xf8\xe7\xc0\xdc\x5e\x49\x46\x8e\x65\xb5\xec\x78\x65\x33\xf3\x86\x67\xde\xf0\x1c\x8a\xcd\xec\x57\xbb\x8d\xe5\x57\xb7\xaa\x29\x8d\x8c\x2b\x19\x2d\xfd\x1d\x4e\xf9\x5c\xcc\x3b\xde\xf3\x3a\x61\xaf\x24\x9e\x11\x39\x98\x6a\x95\x26\x7b\x34\x8e\xe3\xbc\x03\xae\x9d\x1d\xd4\x6e\x41\xd9\x1c\x6c\xf1\x65\x30\x26\xf9\x88\x3f\x99\xfa\x38\x25\x52\x5b\x03\x63\x92\x30\x26\x39\xb0\x65\xe6\x14\x9e\xe6\x12\xa5\xe8\x4b\xb5\x28\x65\xfc\x7b\xb5\xd8\x24\x7c\x2e\xf9\x08\x5e\xdf\x02\xde\x13\x9a\xf6\x94\xd9\xf3\xea\xd0\xf2\x8f\x0c\x3e\xfd\x8a\xdc\x65\x3d\xbf\xfc\x91\x91\x3f\x32\x72\xf3\x7b\xd8\x90\x6a\x72\xd5\x56\xcf\x85\x9f\xb5\xbc\xf2\x81\x0b\xdf\x86\xf5\x32\xdf\xc2\x36\x0d\x20\x52\x9a\x4d\xb1\xd2\x01\x2a\x27\xd3\xfd\x57\x12\x14\x61\xb9\xa3\x36\x3b\x2f\x75\x34\x72\xf5\x0d\x5f\xb3\xf3\x92\x5c\x39\x83\x23\x64\x7a\x22\xee\x9d\xcb\x92\xcb\xe5\xc7\x4e\xba\x74\xf5\x0e\x5d\xa3\x3c\xf7\x77\xe6\xae\x99\xa6\x7a\x33\xab\xde\x70\x2b\xb7\xe6\x3e\x62\xcc\x84\x04\x93\xb0\x70\x77\x87\xae\xac\x70\xab\x88\x45\xbb\xe5\xab\x41\x6e\x74\xf2\xea\x5d\xbc\xd5\x0a\x34\x93\x53\x84\x13\x73\x0a\x27\x45\x1c\x21\x18\xb5\xc7\xb4\xa9\x95\x9c\xc2\x49\xc2\x34\x09\x7b\x9b\xb5\x7a\x1b\x1b\xde\xf5\x76\xf4\x09\xeb\x5d\x45\xfb\x7d\xc6\xe2\x63\x87\xb2\x59\x1a\x33\xf9\x75\xbc\x24\x34\x25\x28\x5e\xce\xeb\x8d\xa5\xc9\xfb\x24\x24\xbd\x3c\x3f\xd2\x50\xc6\xf7\xa3\x76\x0e\xa9\xad\xe6\x48\xa5\xd2\x2a\x55\x56\x3b\x2a\x16\x75\x56\x1c\x55\xcd\x11\x35\x56\x9c\x54\xcd\x9f\xa2\xc8\x5e\xd7\x8e\xec\x7b\xeb\x25\x63\xe5\x10\xc9\x57\x82\x34\xa3\xc3\x64\x3f\x62\xa8\x34\xef\xb4\x00\xc3\x53\x38\x29\x02\x9c\x15\x60\x3d\xd8\x3b\x78\xca\xeb\xa6\x10\xf6\xf6\xd4\x4c\x5d\x34\x23\xe6\x20\xc9\x82\x98\x4a\x63\x7b\x2c\xc8\xf8\x06\x43\x25\xb9\x03\x27\xfb\xcc\x88\xc9\x83\xa5\x9c\x36\x58\xaf\xb7\x76\x74\x36\x22\xe4\x34\xb7\xb2\xcd\xfb\xaa\xb9\x2a\xa9\xc7\x15\x52\xa9\xa4\x4a\x5f\x90\x6c\x74\xb6\x5f\x90\x56\x2b\x30\xc4\x48\x84\x6f\x6f\xdf\xfd\x0e\xbd\xfc\xf9\xd3\xc7\xdf\xc1\xf1\x39\x33\xb3\xb1\x62\x9a\xfb\xcc\x18\x24\xe3\xcf\x51\x72\xa5\x8d\x3f\x56\x8a\x0c\x69\x96\x0c\x38\x23\xcc\x62\x99\x88\xf0\x1b\x6a\x3f\x34\xc6\xaf\x8d\x79\xb1\x90\x5e\x68\x8c\x03\x13\x16\x19\xec\xef\xc6\x75\x67\x9e\x85\x2a\x56\x31\x4a\xf2\xef\x36\x4f\xd9\xc4\x77\x95\x79\xbb\x77\xf8\xae\xdd\xdf\xea\xb4\x79\x88\x0e\x6a\x88\x6e\xc3\x3b\x49\x65\xf6\x25\x0f\x0c\x92\x3d\x0c\xf7\xa8\x0f\xab\x4a\xae\x9d\x78\xec\x8e\xdd\xf7\x56\x8d\xc5\xdd\xce\x10\x80\x7b\xfd\xe1\xe6\xd6\x3d\x6d\xbc\x4d\x75\x14\xc0\x49\xcf\xfd\x69\xf3\x31\xd0\xed\x7b\x89\x56\x49\xcf\xcd\x3f\x1d\xba\xfd\xa6\x0e\x67\xc4\x02\x68\xce\x94\xcd\x26\x62\x0c\x80\x1a\xef\xd6\x4d\x33\x26\x0d\x43\x34\x26\x78\x70\xae\xa7\xeb\x5e\x6d\x7e\x62\x02\x3d\xed\x69\x34\x69\x44\x30\x1a\x8d\xc0\x9d\x30\x11\x21\x77\x77\x29\xd8\x9f\xc4\x05\x5c\xbf\x57\x24\x26\xcb\x16\x5e\xaa\xa8\x29\xb2\x24\xbd\xd1\x5a\xe9\x16\x96\x2a\xb2\x78\x4f\x01\x68\x2f\x46\x63\xd8\x14\x1f\x11\xce\xc9\xc7\x03\xec\xce\x04\xc7\x20\x4f\x93\xfd\x82\x76\x5b\x11\x72\x1a\x80\xbb\x4d\xc4\x97\xee\x4e\x8d\x75\xff\x1f\xad\xef\xd6\xf9\x55\x2c\x27\xb6\xf0\x05\xfe\x62\x99\x4d\x65\x5e\x82\x5d\x93\x7b\x93\xc7\xfb\x3b\xd0\x6b\x0e\xb2\xfc\xff\x26\xb8\x59\x03\x2f\x76\xe9\x95\xfa\x61\x27\xbd\xcd\x57\x90\xbe\xa7\x91\xf1\x65\xef\xa1\x3c\x1a\x35\xdf\x73\x7f\xb2\x05\xe7\xf6\xbd\xda\xb2\xd3\x12\x12\x23\xe4\x34\xc2\x2b\x46\x78\x9d\x89\x04\x40\x3a\x6d\x61\xc2\xae\x45\xcb\x57\xd9\x8d\xed\xb5\x3d\x7f\xa0\xb1\x44\x90\x1c\xec\x66\x39\x52\x21\xb3\x21\x6e\xcf\x83\x49\x76\xa5\x0d\xc0\xfd\xf2\xe5\xcb\x17\xef\xdd\x3b\xef\xea\x0a\xde\xbe\x0d\xe2\x38\x30\xa6\x49\x6b\x8d\xa4\xd3\xed\xea\x60\xd7\x66\xdd\x58\xf6\x32\xcf\x8a\x35\x31\x5f\xe8\x0b\x39\x2f\x9f\xb6\xd7\x3a\x6b\xbf\xdf\x08\x45\x2b\xb1\xf6\x4a\xeb\xf6\xbd\x30\x12\xe1\xb7\x3d\x71\x28\x83\x70\x33\x95\x9a\xf9\x17\xe5\x99\x20\xbf\xed\x16\x5f\xa2\xb6\x9b\xdf\xff\x02\x00\x00\xff\xff\xa0\xb8\x22\x60\xd5\x21\x00\x00"

func templatesViewsDefaultHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsDefaultHtml,
		"templates/views/default.html",
	)
}

func templatesViewsDefaultHtml() (*asset, error) {
	bytes, err := templatesViewsDefaultHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/default.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesViewsFilesHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6d\x6f\xe3\xb8\x11\xfe\x9e\x5f\x31\x60\xdd\x93\x7d\x88\x24\x27\xb8\xbd\x16\x3e\xdb\x45\xee\xb2\xd7\x14\x88\xdb\x43\x36\x87\x22\xbb\x58\x2c\x68\x71\x1c\x33\x2b\x91\x5a\x92\xb6\xe3\x35\xfc\xdf\x0b\x52\x96\x2c\x5b\x92\xe3\xec\xa6\x40\xbb\x38\x7f\x08\x28\x71\x66\x9e\xe1\xbc\x3c\x22\x99\xd5\x0a\x18\x4e\xb8\x40\x20\x29\xbd\x47\x02\xeb\xf5\x09\x00\x40\x9f\xf1\x39\x44\x31\xd5\x7a\x40\x16\x18\xc7\x04\xb4\x59\xc6\x38\x20\x72\x8e\x6a\x12\xcb\x45\x0f\xe8\xcc\x48\x32\x74\xd2\x4e\x63\x22\x55\x92\xab\xd8\xb1\x3f\x95\x8a\x7f\x96\xc2\xd0\x98\x80\x92\x56\xd9\xbe\x26\x90\xa0\x99\x4a\x36\x20\xf7\x68\x08\xd0\xc8\x70\x29\x06\x64\xb5\x82\xe0\x06\x3f\xcd\x50\x9b\xe0\xf7\x9b\xeb\xe0\x37\x6a\xa6\xb0\x5e\x13\xe0\x6c\x40\x26\x3c\x46\xed\x67\xda\x42\xce\x69\xcc\x19\x35\xb8\xc5\x76\xf8\x5c\xa4\x33\x03\x66\x99\xe2\x80\x4c\x39\x63\x28\x08\xcc\x69\x3c\xc3\x8d\x3e\x01\x41\x13\x1c\x90\x0c\x91\x40\x78\xbc\x7e\xae\x3a\x51\x32\xf9\x22\x45\x23\xab\x6a\xa5\x08\x47\x32\xf6\x13\xe6\xff\x40\x76\x45\x56\x2b\xe0\x67\x7f\x15\x40\x2e\xa9\x41\x50\x54\xd8\x04\x05\x79\x8a\x76\x8c\x4d\x38\xc6\x4c\xa3\x19\x56\xa6\xaa\x58\xc2\x28\x19\xfb\xf7\x4a\xce\x52\x52\x2f\xdf\xa0\xa3\x0f\x88\xef\xab\xb8\xa0\xf8\xa9\xc2\x14\x05\x83\xec\xe9\x29\xc8\xc2\x8e\x4e\xa9\xc8\x0d\x51\xc6\x7c\x29\xca\x16\x7c\xca\x98\x14\x64\xd8\xe7\x45\xb9\x51\x98\x50\x3f\xa2\x31\x0a\x46\x15\x19\xf6\x43\x3e\xec\x87\xd6\xca\x11\x60\xe5\xf4\x19\x7c\x34\xe5\x92\xb3\x65\x46\x76\x8a\x7a\x13\x8a\x22\xc3\xb6\x6e\xad\xd4\x07\x5b\x1b\xc1\xaf\x52\x25\xd4\x00\x39\xef\x76\x7f\x0c\xba\x67\x41\xf7\x1c\xce\x5e\xf5\xba\x3f\xd8\xbe\x02\x1f\x0a\x61\x23\x0f\x8a\x56\x8a\xa5\xe2\x74\xc8\xf8\xfc\x40\xee\x9a\xa7\x1b\xa6\xfa\x61\x7d\x05\xd5\x88\xd7\x14\xee\x79\x53\xe1\xfe\x32\xa5\x42\x60\xfc\x0d\x54\x6d\xce\x7f\x0b\xce\xcc\xb4\x07\x67\xdd\xee\x9f\x8f\xaa\x64\x8c\x31\x32\xae\xa2\xa2\x3c\x16\x19\x25\x14\x8f\x35\xd5\x05\x99\xda\x39\x01\x85\x9f\x66\x5c\x21\x1b\x90\x7c\x74\x04\xea\x6a\x95\x91\x05\xb4\x38\x3b\x85\x96\xe1\x26\x46\xe8\x0d\x20\xd8\x60\xea\xba\x64\xd4\xba\x2f\x53\xcb\x95\xa5\x5a\x6f\x71\x66\xeb\xd3\x26\x78\x02\xf8\x09\x5a\xb9\xcd\xcd\xcc\xc6\x73\xeb\x70\x3e\xb2\xc2\x36\x9a\xeb\xf5\xd0\x1a\xc8\xbc\x59\xaf\xfb\x61\x66\xfc\xa8\xe5\x64\xfa\x4f\x87\x3b\xcc\x40\xbf\x8d\xe6\x79\x3d\x47\x91\x11\xd3\x1f\xfd\xc3\x06\x04\x6d\x38\xfc\x2c\x1c\x59\x0b\x95\xdf\xbc\x7c\x17\x41\xb5\x01\xbe\xdf\x16\x7e\xe0\xc0\x3f\x58\x70\x20\xdf\x93\x63\x4a\x3f\xcb\xea\x45\xbc\xa1\xc3\xe3\x1b\xa0\xc6\x93\x8b\x06\x4f\x2e\x9e\xe5\x09\xb5\x1b\xaa\xaf\xf6\x65\xd4\xe0\xcb\xe8\x19\xbe\x8c\xa4\xb3\xc8\xd0\x60\x64\x5e\xc0\xa7\x9b\x06\x9f\x6e\x9e\xe1\xd3\xdf\x51\xa0\xa2\x2f\x91\xad\xab\x06\x6f\xae\x9e\x13\x21\x2a\x66\xcf\x77\xe6\x9b\x60\x44\xfb\xfb\x4e\x8c\x75\xfa\xd3\xff\x35\x05\x1e\x91\xad\xf1\xcc\x18\x29\x36\x9b\x51\x3d\x1b\x27\xdc\x14\xd4\x36\x36\x02\xc6\x46\xf8\x7a\x16\x45\xa8\x35\xd9\x96\xc6\x9b\xa9\x5c\xe4\x85\x91\x59\xf8\xdf\x4c\x77\x3f\xb4\xf4\x9c\x3d\x97\xa6\xfa\x86\x8e\x63\xcc\x97\x99\x3d\xb8\xbf\xbe\x36\x8a\xa7\xc8\x80\x51\x43\xb3\xf7\xcc\xf8\x0a\x75\x2a\x85\xe6\x73\x04\x21\x17\x8a\xee\x7f\x5b\xf6\x3e\x2d\x7d\x33\x45\xca\xca\xcf\x6a\xcf\x49\x33\xdd\x86\xf2\x9f\x34\xc1\x3c\x94\x66\x7a\x40\xf0\x67\xbc\xe7\x02\x0c\x3f\x52\xfc\xb5\x60\xc7\x0b\xbf\xe1\x9f\x8f\x13\xbc\x70\x47\x59\x5d\x27\xdb\x0f\xcb\xeb\xb4\x73\x7b\x51\x18\x4b\xb6\xdc\x3e\x97\xb6\x8c\xa7\xd0\xb2\x87\x1f\xb7\x61\x74\xa7\xa0\xf2\xd6\xa3\x26\x7a\xcc\x6d\xeb\xac\x64\xf0\x2b\x8f\xd1\x46\x30\xf3\x86\x55\x25\xfb\x3a\x52\x3c\xcd\x4f\x5b\x34\x4d\x63\x1e\x51\xbb\x86\xf0\x81\xce\x69\x36\x49\x86\x4c\x46\xb3\x04\x85\x09\x16\x8a\x1b\x6c\xdb\x13\xd3\xad\x7c\x63\x14\x17\xf7\x6d\xaf\x00\x73\x19\xb8\xe5\x09\xee\x1c\xa5\xfc\xee\x99\xdf\x3d\xbf\x75\x47\xa9\x5e\xf7\xd5\xdb\xee\x5f\x7a\xdd\xae\xa5\x59\xaf\xd3\xe9\x87\x19\xc2\xf0\xbf\xef\xdc\x6b\xc1\x5e\xd2\xb5\xd5\x0a\xa6\xb3\x84\x8a\x0f\xe3\xa5\x41\x5d\x8a\xf6\x35\x8a\x7b\x77\x51\x52\xab\x57\xf4\x14\x3e\x1a\x3f\x42\x61\x50\xd5\xb0\x50\x99\xbd\x2c\xc1\x38\xb6\x82\x62\xe4\x3f\x36\xd1\x5f\x9f\xc2\x54\xe1\x24\x3b\x18\xd4\xdd\xde\xfc\x6d\x73\xbb\xc3\xe4\x42\xc4\x92\xb2\xef\xdc\xbe\xad\xae\x5c\x9a\x58\xce\x8d\x79\xb4\x7f\xdc\xd7\xf6\xbc\x6f\x6d\xf8\xb9\x69\x02\xee\x60\xe1\x7c\xd9\x5c\x9a\x14\x33\x2d\x0b\xb0\xb9\x14\xa0\x75\xec\x55\xfd\x28\xed\x44\x73\xb7\x97\xaa\xc7\x91\x7e\x58\xea\xa6\x7e\xe8\x98\x6a\x78\xb2\x95\x3b\x29\xdd\xb2\xd9\x36\x2c\x6e\xd9\x56\x2b\xd0\x86\x1a\x1e\x5d\xdd\x8e\xae\xa1\x9d\x8d\x7f\xbf\xb9\x06\x12\x32\xaa\xa7\x63\x49\x15\x0b\xa9\xd6\x68\x74\x38\x47\xc1\xa4\xd2\xe1\x58\x4a\xa3\x8d\xa2\xa9\xbb\x9c\x70\x4d\x9b\xf2\xe8\x23\xaa\x30\xd2\x3a\xdc\x7b\x17\x24\x5c\x04\x91\xd6\x04\x26\x34\xd6\xd8\xb1\xc0\xf5\x7e\x3d\xe8\xaf\xf2\x2a\x91\xb6\x27\xc2\x87\x7c\xe4\x80\x1f\x76\x70\x5f\x7e\xc1\x0f\xf5\xeb\xdd\x85\xcd\x92\x72\x54\x6b\x17\x09\x6d\xb5\xf3\x2e\xef\x04\x0a\x29\x5b\xb6\x27\x33\xe1\x4a\x19\xda\x1d\x58\xed\x94\x4a\xab\xed\xfd\x69\x7b\x59\xe4\x75\x82\x3d\x97\xda\xab\x4a\xc1\xd9\x0f\xc1\x6f\x6e\xf2\x1f\x22\x52\x68\x71\x7a\xf0\x63\xf7\xb4\x22\x98\x70\x71\x49\x0d\xf6\x20\x0b\x6a\xbb\x13\xe8\xd9\xd8\x28\x1a\x99\xf6\xab\x53\xf0\x96\x48\x95\xd7\xa9\x51\xa3\x8f\x7b\x6a\x28\xd8\xbf\x26\x6d\x8f\xd1\x65\x9d\xbc\x8d\xc7\xf2\x67\xb7\x71\xf8\xc5\x76\x18\xea\x1e\x78\xa5\x26\xf4\xaa\x2a\xb1\x8c\x68\x8c\x3d\xa8\x2e\xce\xfe\x26\x5c\x69\x73\x49\x97\x3d\x38\xab\xaa\x3a\x01\x47\x8b\x3d\xf0\xee\xee\xee\xee\x82\xd1\x28\xb8\xbc\x84\xab\xab\x5e\x92\x78\x15\xf1\x75\xd5\x82\x8b\xae\x6e\x02\xf7\x0a\x06\xb8\x95\x8c\x2e\xb3\x8f\xa3\xd7\x83\x77\x79\x34\x4e\x8b\xb8\xbc\xaf\xf7\x6e\x6b\xe1\x0e\xb5\x41\x55\x6b\x65\x9b\x8a\xb3\x53\xb0\x91\xd5\x5e\xc9\x72\xdd\xec\x93\x68\xb7\x53\xae\x61\x81\xf8\xb1\x0e\xcd\x50\x65\x6c\x0e\xb9\x96\xff\x46\xfc\xb8\x03\xb6\x49\x6f\x31\xf5\x24\xd2\x35\xd5\xa6\x11\xa9\xec\x79\x61\xf2\xb0\x03\x0d\x3a\xcf\xf6\xcb\x45\x20\x91\xc2\x4c\x0f\x85\xc0\x09\xd4\x05\x60\x33\x71\xdc\xf2\x1b\x61\xca\x6b\xd9\x58\x3c\x84\x5d\x2b\xff\x4c\x87\xdc\xba\x6d\x3b\x1f\x5a\xf6\xa6\xdd\x2b\xab\xce\xde\x1f\xb7\xe8\x26\x8c\xf2\x1a\x32\x7b\x07\x70\xeb\xa4\x77\x7d\x79\xaa\x8b\xd7\x9d\x9f\x4e\x1a\x58\xd4\x52\x83\xe7\x30\x12\x6e\x4a\xbc\x8b\xfb\xc4\x6b\x7f\x73\xaa\x80\xa5\x30\xa8\x65\x61\x6a\x09\x6f\x87\x8b\xbd\x7d\xd8\x2a\x74\x76\x62\x7b\x57\xfa\x7f\xcf\x7b\xaf\x13\xcc\x69\xdc\xde\x2c\x9f\xa5\x59\x64\x2c\xc3\x76\x82\x8c\xc8\xda\x8e\xc7\xfc\xd1\xc8\xbf\xbc\x7c\x77\xfb\xde\x31\x59\x4f\xeb\xb7\x5e\xa7\x53\x3d\xb2\x1e\xc0\x33\xb2\x06\x0d\x05\xfb\x22\xac\x75\xe9\x39\x1f\x17\xdb\xcc\xed\x2e\xe0\x3f\x01\x00\x00\xff\xff\xba\x7d\x01\xda\x02\x1c\x00\x00"

func templatesViewsFilesHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsFilesHtml,
		"templates/views/files.html",
	)
}

func templatesViewsFilesHtml() (*asset, error) {
	bytes, err := templatesViewsFilesHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/files.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesViewsLogsHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x4f\x8f\xd3\x30\x10\xc5\xef\xfb\x29\x46\x66\xd1\xee\x4a\x24\x71\x56\x2c\xa0\x28\xc9\x09\x71\x43\x42\xa2\x5c\xb8\x4d\xe3\x21\x35\x4a\x6d\xcb\x1e\x5a\x45\x56\xbe\x3b\x4a\x42\xff\xd1\x16\x7a\x89\x34\xf1\xcf\xcf\x6f\xe6\x4d\x8c\xa0\xe8\x87\x36\x04\xc2\x61\x4b\x02\x86\xe1\x0e\x00\xa0\x64\x5c\x76\x04\x4d\x87\x21\x54\x62\x2e\xa6\x6f\x12\xd8\x6b\x47\x0a\x14\x32\xce\xff\x15\x27\x9e\x82\xb3\x26\xe8\x0d\x81\xb1\x5b\x8f\x4e\x40\xe0\xbe\xa3\x4a\x6c\xb5\xe2\x55\x91\x4b\xf9\x5a\xd4\x93\xf2\xac\xbe\x22\x54\xc7\xb5\x3f\x14\x7f\x80\x3a\x46\xd0\xf9\x07\x03\xe2\x95\x80\x14\x86\xa1\xcc\x78\xf5\x0f\x6a\xa1\xd7\x74\x13\xf8\x99\x42\x98\x5a\xbd\x45\xb4\x77\xb7\x81\xdf\x02\xf9\x4b\x60\x99\x1d\xb7\x36\x9e\xfd\xd5\xf8\xd2\xaa\xfe\x50\xc7\x08\x1e\x4d\x4b\x70\xaf\xdf\xc0\x7d\x67\x5b\x28\x2a\x48\x3b\xdb\x86\x5d\x30\x57\xc6\xa5\x46\x2b\xe3\x85\xf4\x8b\x0d\x9a\xb5\x35\xb3\x15\x75\x0e\x96\xa1\xf1\xda\x31\x70\xef\xa8\x12\xe8\x5c\xa7\x1b\x1c\x6f\x64\x3f\x71\x83\xf3\xa1\xa8\x95\x6d\x7e\xad\xc9\x70\xba\xf5\x9a\xe9\x51\x21\xd3\xc2\x7e\x65\xaf\x4d\xfb\xf8\xb0\x7b\x6b\x9c\x79\xfa\xc9\xfa\x35\x32\x88\x67\x29\xdf\x25\x32\x4f\xe4\xf3\x22\x7f\x29\xe4\xdb\x42\xbe\x7c\x97\xef\x0b\x29\xc7\xa5\x7a\x78\x7a\x2a\xb3\x59\xbb\xbe\x6c\x6b\xa7\xf9\x11\x19\xaf\x7a\xdf\x3f\xdc\x3b\xfa\x2f\x34\x66\x72\x06\x9d\x06\x12\x23\x90\x51\x27\xb3\xcd\x8e\x22\x29\xb3\x69\xc3\xeb\xbb\x3d\xf7\x3b\x00\x00\xff\xff\xd5\xb4\x73\xc0\x30\x03\x00\x00"

func templatesViewsLogsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsLogsHtml,
		"templates/views/logs.html",
	)
}

func templatesViewsLogsHtml() (*asset, error) {
	bytes, err := templatesViewsLogsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/logs.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesViewsUserHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\x41\x6b\xdc\x3a\x10\xbe\xe7\x57\x0c\x22\x87\xf7\x1e\xcf\x36\x49\x20\x94\xb2\xde\x4b\x0f\xed\x21\x29\xa5\x34\xe7\xa2\xb5\xc6\x1b\x05\x59\x72\x34\xf2\x26\xad\xf1\x7f\x2f\x92\x6c\xaf\xe3\xdd\x2d\x49\xba\xd0\xdc\x46\xb6\xf4\xcd\xf7\xcd\x37\x23\xd4\xb6\x20\xb0\x94\x1a\x81\xd5\x7c\x8d\x0c\xba\xee\x04\x00\x60\x51\x1a\x5b\x41\xa1\x38\x51\xce\x7c\x9c\xdc\x1a\x2b\x7f\x1a\xed\xb8\x82\xb0\x56\x7c\x85\x2a\x51\x58\x3a\x06\xd6\x28\x8c\xdb\x18\x54\xe8\x6e\x8d\xc8\x59\x6d\xc8\x31\x90\x22\x67\xc4\x37\x98\x34\x84\x96\x81\x36\x1b\xae\xa4\xe0\x0e\x97\x21\x4d\x48\x25\xe4\x66\xc8\x24\x1d\x56\x11\x7e\x6d\x4d\x53\xb3\xed\xae\xb0\x33\xe4\xf4\xff\x73\xa6\x79\x85\x6c\x38\x56\x18\xed\xac\x51\x91\x13\x14\x46\x25\x95\x48\x2e\x42\x40\x55\x1f\x3c\x52\x72\x76\x3e\x03\x6c\x5b\x90\x67\xef\x34\xb0\xcf\x01\x2d\x85\xae\x83\x05\xd5\x5c\x0f\xc0\x16\xef\x1b\x69\x51\xb0\xe5\x7f\x8b\xcc\xff\x98\x11\xca\x42\xc6\xd9\xc7\x89\x9e\x9e\xca\xe5\x40\xe5\xf2\x20\x95\x70\x52\xea\xba\x71\xe0\x7e\xd4\x98\x33\x87\x8f\x8e\x3d\x71\xa0\x57\xc9\xc0\x6b\x1f\x2a\xe0\x0b\x1c\xa3\x81\xeb\x84\x35\x6c\xb8\x6a\x30\x67\x6d\x0b\xa9\x37\x20\xf5\x3a\xa1\xeb\xe6\x75\xcd\x84\xdc\x4c\x0c\x99\x2d\x5f\xe5\x4f\x85\x95\x39\x9e\x3f\xd7\x01\x2d\x1d\xba\xf3\x4d\xd4\x3f\x2a\xf4\xf5\x8f\xd1\xbc\xd6\x9e\xf3\x4b\x6b\xed\x05\x97\x80\xf7\x53\xbb\x18\x9b\xca\x7e\x9d\x1b\x35\x27\x7a\x30\x56\x1c\xcf\x91\x2f\x23\xe2\x9b\x9a\x9a\x1d\xa1\xfb\x9c\xdb\x6e\xf2\xee\x6d\x57\x7b\x26\xe8\x85\xe6\xa1\x16\x7f\x6e\x56\xfc\x7f\x34\xa7\x3e\x46\xb8\xbf\x6e\x13\xa1\xc2\xc2\x85\x92\xf7\x0a\xa3\x1b\x4f\xe5\x4e\xfd\x82\x78\xe4\xfc\x19\xce\xf4\x8a\x2d\xd7\x6b\x84\x53\xf9\x3f\x9c\x06\x58\x78\x9f\x43\x1a\x22\x9a\xdf\x1d\x23\x2f\x53\x3b\x69\xf4\x64\x7c\xe3\xd1\xf1\xae\x1c\x67\x72\xfa\xfd\x34\x0e\x68\xa8\x2d\x74\x5d\x24\xea\xf9\x0d\x11\x1b\xbb\x61\xb9\x03\xb9\xc8\x62\xce\xbd\x12\x66\x2d\xb4\x75\x23\x22\xbf\xf6\xe6\x56\xfa\x3b\x19\x25\x05\x5b\xfe\x66\xd3\xe1\xf6\x3c\x64\x7b\x25\x12\x53\x96\x84\x2e\xb9\xd8\xd3\x7f\x25\x68\x3c\x7c\x95\x8d\xe0\xab\xc6\x39\xa3\xe3\x5b\x01\xb5\x60\xfd\x2c\x53\xb3\xaa\xe4\xf6\x0e\x5e\x39\x0d\x2b\xa7\x93\x07\x6e\xb5\xd4\x6b\xb6\x1c\x3b\xfc\xa6\xf6\x2f\x8a\xd8\xe2\x8b\x2c\xa2\xed\x90\x41\x45\x78\x84\xec\xd4\x14\x05\x12\x4d\xb2\x7f\xb0\xf8\x8c\xec\xbb\xae\x1e\xf6\x6e\x91\x79\x23\x96\x27\xdb\x73\x27\x93\x77\xda\x1d\x8d\x55\x6c\x5b\x20\xc7\x9d\x2c\x3e\x7d\xbb\xbe\x82\x7f\x62\x7c\xf3\xf5\x0a\x58\x26\x38\xdd\xae\x0c\xb7\x22\xe3\x44\xe8\x28\xdb\xa0\x16\xc6\x52\xd6\x3f\xc0\x8c\xcd\xee\x26\x8b\xb4\x92\x3a\xf5\xc8\x25\x57\x84\xff\xfa\x04\x63\xf6\x5f\x01\x00\x00\xff\xff\x8c\x52\x21\x9f\x24\x0a\x00\x00"

func templatesViewsUserHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsUserHtml,
		"templates/views/user.html",
	)
}

func templatesViewsUserHtml() (*asset, error) {
	bytes, err := templatesViewsUserHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/user.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xdf\x4b\xe4\x56\x14\xfe\x62\x6c\x6b\xa7\x3f\x68\xa5\x0f\xa5\xf4\xe1\xf6\x41\x69\x1f\x62\x27\xb6\x05\x89\x46\x4b\xad\x42\xa9\xd3\x8a\x1d\x85\x3e\x86\x99\x6b\x4c\x3b\x93\x0c\xb9\x89\x55\xb0\x74\x1c\xa5\xb6\x20\x2d\x08\xcb\xc2\x2e\xec\x32\x3e\xef\xc2\xec\xca\x60\xf0\x47\xfc\x17\xce\x7d\xda\xb7\x85\x7d\xdb\x87\xfd\x23\x96\x9b\x64\x5c\x16\x16\xf6\x42\x72\xbf\x73\xce\x77\xbe\xf3\x9d\x90\xc7\xa3\xc3\x37\x00\xe0\x3d\x00\x9f\x02\xd8\x07\xf0\x21\x80\x67\xc8\x4f\x59\x03\x46\x01\x98\x1a\xf0\x11\x80\xef\x35\xe0\x1d\x00\xbf\x6a\xc0\x30\x00\xae\x01\x6f\x01\xf8\x4d\x03\x74\x00\xa1\x06\xbc\x09\x60\xb3\xa8\xff\x59\xdc\x9d\xe2\x3e\xd0\x80\x37\x00\x1c\x6a\xc0\xfb\x00\x8e\x34\x60\x4e\x03\xba\x1a\x30\x05\xe0\xc9\x10\xf0\x0d\x80\x4f\x74\x60\x0c\xc0\xb2\x0e\x8c\x28\x3d\x3d\xf7\xf1\x97\x5e\xf4\xe9\x40\x09\x40\xb7\xb8\xef\xe9\xf9\xdc\x53\x1d\xf8\x18\x40\xaa\x03\x1f\x00\x78\x54\xe4\x9f\xea\x80\x86\xdc\x83\x3a\x6a\x87\x91\x02\xab\xfa\xbb\x05\x56\x5a\x6a\x1f\xbd\x88\x87\x90\xfb\x55\xdf\xe7\x6d\x95\x98\x0f\xfc\x75\xcf\x65\x82\x47\x4c\xc4\xb5\x1a\x17\x02\x0b\x5b\xad\x20\x8c\x58\x2d\xab\xc4\xa1\x13\x79\x81\x3f\x48\x36\x02\x57\x60\x49\xbd\x2a\x5c\x08\xc7\xe5\xf8\x29\xf8\x03\x3f\xb7\x32\x4e\xd5\x6b\x72\x54\xb7\x5b\x1c\xab\x82\x87\x58\x73\x1a\x31\x07\x75\x29\x91\xbb\xb2\x43\x7d\xba\xa0\x1e\x56\xb8\xd2\x31\x2a\xc2\xf5\xea\xc6\x77\xb1\x2b\x8c\x6a\x60\xb1\x3a\xdf\xfc\xf6\x77\x6f\xc3\x69\x06\x13\x61\x5c\x5a\x72\x44\x64\x54\x43\xc7\x17\x0d\x27\x0a\x42\x8b\xfd\x98\x95\x58\x25\x0e\x9d\x66\x50\x0f\xd8\xcc\x4b\xfc\xd9\xd2\x92\xe3\xbb\xb1\xe3\x72\xa3\xca\x9d\xa6\xc5\xae\x63\x8b\xad\xc4\x42\x78\x8e\x5f\xaa\xfc\x50\x59\x30\xd6\x78\x28\xbc\xc0\xb7\x98\x39\x51\x2e\xcd\x07\x7e\xc4\xfd\xc8\x50\x7e\x2d\x16\xf1\xad\xe8\xcb\x56\xc3\xf1\xfc\x69\x56\xdb\x70\x42\xc1\x23\x7b\xb5\xba\x68\x4c\xbd\xe0\x29\x3f\xeb\x3c\x34\x16\xfc\x5a\x50\xf7\x7c\xd7\x62\xa5\xe5\x46\x1c\x3a\x0d\x63\x31\x08\x9b\xc2\x62\x7e\x2b\x0b\x85\xfd\xd5\x34\xcb\xa1\xed\x8f\x99\x65\xdb\x36\xd9\xf8\x38\x53\xb0\xfc\x99\x6d\x9a\x6c\x8e\x95\x99\x95\xc5\xb3\xf6\xe4\xa0\x34\x63\x7f\xad\xe0\xe7\x19\x6d\xc6\x2c\xb3\x9d\x9d\xbc\x65\xd6\x9e\x2c\x7f\xc1\xe6\x98\xc9\x2c\x36\x39\x0d\xba\x45\x29\x5d\xca\x7d\x4a\xe8\x44\xee\xc9\x36\xf5\xe4\xdf\x94\xc8\xff\x99\xdc\x93\xbb\x74\x45\x7d\xf9\x0f\x5d\x52\xca\x28\xa1\x53\xba\xa0\x3e\x5d\x66\x4f\x0f\x74\x9f\xce\x32\x46\x2a\xdb\xb2\x43\x89\x6c\x53\x4a\x0f\xa9\x27\x3b\xf2\x90\xd1\xd9\x2b\x55\xff\x7b\x4d\xd7\x39\xa5\x74\x42\x09\xe8\xf6\x00\x74\x29\xa5\x94\x1e\xc8\x7f\xb3\xa9\x09\xf5\x41\xc7\xd4\xa7\x33\xb9\x97\xa5\xfa\xa0\x3b\x74\x95\x3b\x06\x1d\xc9\xb6\xfa\x2b\x14\x3c\xa6\x84\xae\x40\x77\x29\xa5\x73\x79\x48\xa7\x83\x29\xd4\x57\x31\xe8\xa6\xda\x41\x1e\x5c\xab\xfe\xb2\x2d\x22\xde\xc4\xf3\x00\x00\x00\xff\xff\xb6\x0e\x85\xf6\xef\x03\x00\x00"

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
	"templates/layouts/widget.html":    templatesLayoutsWidgetHtml,
	"templates/views/account.html":     templatesViewsAccountHtml,
	"templates/views/default.html":     templatesViewsDefaultHtml,
	"templates/views/files.html":       templatesViewsFilesHtml,
	"templates/views/logs.html":        templatesViewsLogsHtml,
	"templates/views/user.html":        templatesViewsUserHtml,
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
		"layouts": &bintree{nil, map[string]*bintree{
			"widget.html": &bintree{templatesLayoutsWidgetHtml, map[string]*bintree{}},
		}},
		"views": &bintree{nil, map[string]*bintree{
			"account.html": &bintree{templatesViewsAccountHtml, map[string]*bintree{}},
			"default.html": &bintree{templatesViewsDefaultHtml, map[string]*bintree{}},
			"files.html":   &bintree{templatesViewsFilesHtml, map[string]*bintree{}},
			"logs.html":    &bintree{templatesViewsLogsHtml, map[string]*bintree{}},
			"user.html":    &bintree{templatesViewsUserHtml, map[string]*bintree{}},
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
