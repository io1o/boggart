// Code generated by go-bindata.
// sources:
// templates/views/index.html
// templates/views/light.html
// locales/ru/LC_MESSAGES/light.mo
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package esphome

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

var _templatesViewsIndexHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x5f\x6f\xdb\x36\x10\x7f\xef\xa7\x20\x88\x6c\x48\x80\xca\x4a\x52\xa0\x18\x5c\x39\x43\xb1\x60\x58\x81\xae\x2d\xba\xf4\x39\x38\x49\x27\x8b\x2e\x45\x2a\x24\xe5\xd8\x35\xfc\xdd\x07\x92\x92\x2d\xdb\x92\xa3\x64\x49\x5f\x96\x87\x96\x7f\xee\xff\xef\x78\x77\xd6\x6a\x45\x52\xcc\x98\x40\x42\x13\x29\x0c\x0a\x43\xc9\x7a\xfd\x2a\x4a\xd9\x9c\x24\x1c\xb4\x9e\x50\x25\xef\xe9\xd5\x2b\x42\x08\x69\x9f\x26\x92\x07\x45\x1a\x5c\x5c\x12\xbb\xd2\x45\xb3\x5a\xe8\xe0\xe2\xb2\xa6\xdf\xe7\x59\xdc\x96\x20\x90\xb7\x6e\x0f\x29\x1a\x2b\x76\x69\x1c\x9d\x81\x98\x63\x43\xe9\x37\xee\xdf\x40\x1b\xc5\x4a\x4c\x49\x0a\x06\xfc\x79\x6a\x02\x85\xba\x94\x42\xb3\x39\x12\x21\xef\x15\x94\x94\x68\xb3\xe4\x38\xa1\xf7\x2c\x35\xf9\xf8\xe2\xfc\xfc\x97\x0e\x2d\x5e\x53\x8e\x90\x76\xdf\xf9\x7b\xd5\x7f\x59\x0b\xb8\x5a\xad\x08\xbb\xf8\x4d\x10\xfa\x39\x9e\x61\x62\xc8\x87\x6b\x4a\x46\x64\xbd\x8e\x42\x93\x3f\x82\xfb\x66\x59\xe2\x93\x18\x3f\x41\xf1\x34\xc6\x7f\x0c\x98\xa7\x71\xbe\x4f\x0c\x93\x42\x0f\xe1\x8d\xc2\xbe\x18\x5a\xbe\xde\xe8\x47\x26\x96\xe9\xb2\x5f\xec\x6a\x45\x14\x88\x29\x92\x13\xf6\x9a\x9c\xa0\x30\xcc\x2c\xc9\x78\x42\x46\x6e\xc9\x50\xdb\xe4\xee\xf7\xe5\x41\x58\x53\xeb\x6c\x2d\x77\xe4\x81\xfd\x70\xed\x9d\x3d\x92\x2f\x1d\xbc\x16\xd6\xa7\xf0\x59\x54\x1f\xc3\xe7\x80\x69\x98\x1d\xb2\xe4\x64\x30\xff\x51\x02\xb2\xf7\x74\x63\x23\x82\xa9\x92\x55\x49\x89\x92\xf6\x99\xf9\xcd\xc3\x42\x88\xc7\x8d\x65\x04\xef\x76\xe3\x43\x39\x9b\xe6\xbe\x20\x0d\x11\x12\x01\xc9\x15\x66\x13\x6a\xe3\x35\xfa\x8a\x77\x15\x6a\x33\xfa\xf6\xf5\xe3\xe8\x0b\x98\x9c\xac\xd7\xbf\x83\x4b\xd0\x89\xd7\xf1\xab\x74\xf8\x4d\xba\x11\xa5\x2d\xbf\x88\xf5\x8d\x89\x4c\xfa\x45\x22\xfd\xc9\x42\x0f\xf4\xce\x19\xc7\x1a\x81\x53\xbe\x2c\x73\x27\x64\xb3\x0a\x70\x89\x81\x2c\x51\x50\x62\x98\xb1\xc1\xdb\xbc\xa9\x6b\x34\xc0\xb8\xa6\x0e\x36\x7a\x15\x85\x6c\x98\xce\x28\x84\xe1\xc6\x1d\x84\xdf\x67\x0a\x95\x62\x70\xf0\x37\x7a\x5f\x14\x84\xd4\x3e\x6e\xf5\x1f\x60\xd8\x98\x79\x14\x0e\x6d\x64\x79\x08\xc5\xe7\x2c\x7b\x34\x0c\x1b\x7d\x8f\x84\x03\xb9\xc6\xff\x65\xe8\x4b\x0e\xcb\x8e\xd0\x8b\x9f\x17\x79\x91\x0e\x0d\xfc\x70\xea\x28\x4c\xd9\xfc\x81\x82\x7b\xb4\x24\xf7\xf7\xcb\x87\x0d\x89\xc2\x9e\xae\x19\x85\x6e\x5a\xda\x1b\xc7\x76\x4d\x6d\x6d\xeb\x65\xfd\xdf\xcb\xce\x88\x96\x8e\x43\xa9\x31\x3d\x3a\x2d\xba\x3c\xe9\x9a\x15\xf3\xcb\xd6\x00\x76\xf3\xbe\x19\x49\xf2\xcb\x0e\xda\x8a\x37\x02\x05\xcc\x89\x80\x79\x0c\x2a\x50\xb6\xf5\x10\x67\xcb\xad\x91\x92\xc7\x72\xd1\x37\x2d\x72\x76\x15\x41\xcb\x67\x67\x77\xc0\x99\xf8\x4e\xaf\xb6\xc9\x9e\x01\xc9\x20\x48\x72\x9c\x2b\x29\x02\xdb\x19\x6d\x2a\xdb\xec\x8c\x42\xde\x91\xd3\x51\x58\xf1\x8e\xd3\x76\x78\x39\x82\xca\xd8\xc2\x4a\x3a\x48\xaf\xae\xa3\x67\x1b\xb4\x1f\x3f\x50\xf7\x8f\x6d\x47\xa7\xae\x83\xc1\xb4\x1a\x36\x5d\xd6\xc3\x4f\x46\x46\xd2\xc0\xad\xaa\x84\x60\x62\x4a\xd6\xeb\x8d\xac\xfa\xc8\x0b\xdb\xd6\xdb\xed\x3d\x42\xba\xdc\xde\xba\xc7\xd5\xff\x42\x8f\x4c\xb3\x83\xbd\xfb\xa2\xe4\x54\xa1\x1e\xec\xdf\xf1\x62\xd2\x82\xba\xdc\x08\xae\xe1\x2a\x40\x4d\x99\x08\x62\x69\x8c\x2c\xc6\xe7\x03\x4a\x78\x97\xb8\x20\x06\x45\xda\x9b\x4d\x76\x74\x46\x7e\x87\xd2\xcd\x52\xb6\x19\xcd\x71\x1b\xfb\x1d\x0a\xdf\x6d\x7c\xe8\x6d\x33\xf2\x23\x65\x43\x12\x83\xa2\x04\x14\x83\x60\x0e\xbc\x42\x21\xef\x5d\xaf\x70\x2a\xef\x15\x33\x06\x85\x6b\x61\x5b\x92\x82\x89\x09\x3d\xdf\x39\x81\xc5\x96\xc9\x48\x03\xdc\xb1\xec\xa4\x74\x73\xdd\xe8\x25\xeb\x75\x5f\x8a\x1f\x84\x4c\x97\x20\x1a\x69\x9c\x09\x0c\x72\xb4\x25\x65\x7c\x79\x5e\x2e\xe8\xd5\xa1\xb5\x24\x5e\x1a\xd4\x44\x66\x64\xdf\xa8\x28\xb4\xb2\x06\xa0\x34\xa4\xc9\x1c\x25\x79\xd9\x0c\xff\x23\xc7\xe4\xbb\xae\x8a\x47\xbc\x60\x17\x87\xa4\xe6\x7b\xea\x13\x7c\xb0\xfd\x1d\xde\x64\x52\x15\x9b\xc2\x2d\x55\x11\xe4\x52\xb1\x1f\x52\x58\x40\xdc\x9e\x43\x8c\x3c\xe0\x98\x19\x92\x2a\x59\xfe\x90\x02\x9b\x24\xb5\xf7\x94\x14\x68\x72\x99\x4e\x68\x29\xb5\xa1\x84\xa5\x13\x2a\x0d\x50\x52\x4f\x60\xb4\x19\xc5\xdc\xa1\x90\x73\xe0\x2c\x05\x83\x3d\x1e\xb4\x9e\x1f\x33\x58\x78\x13\x1e\xfa\x7d\x15\x39\x1b\x2d\xed\x84\x72\x9c\x23\xa7\xdb\x1e\x25\x8c\x92\xdc\x3b\x41\xea\x2e\xfd\xa6\x69\xd2\x6f\x3a\x7b\xf4\xfe\xdf\xf6\x03\x05\x2b\x50\x56\xc6\x83\x7a\x24\xb3\x9c\xb2\x23\xd6\x1e\xce\x0d\x6f\x1b\x8b\xde\x0e\xb2\xc8\x49\x61\xa2\xac\x0c\x31\xcb\x12\x27\xd4\xe0\xc2\xd0\x1d\x14\x6b\xc7\x3d\x1e\xa6\x31\x5c\x40\x81\xad\xad\xab\x0e\xad\xd2\xe0\xcf\x5d\x71\x50\x78\x57\x31\x85\xe9\x84\x36\xab\x63\xf1\xef\x7f\x6b\x3d\x57\x51\x68\xad\x7c\xda\x38\xb6\x1d\x01\x5f\xb5\xbe\xe9\xe5\x08\xe9\xe6\x27\xdc\x6a\x45\xb4\x01\xc3\x92\xbf\x6e\xfe\xfe\x48\x4e\xfd\xfa\xdb\xd7\x8f\x84\x86\x29\xe8\x3c\x96\xa0\xd2\x10\xb4\x46\xa3\xc3\x39\x8a\x54\x2a\x1d\x6e\x3e\xa9\xe9\x91\x40\x13\xc4\x3a\x4c\xb4\x3f\xbd\xf1\xa7\xb1\x94\x46\x1b\x05\xe5\xa8\x60\x62\x94\xd8\x36\x93\x01\xd7\x78\xf6\x8c\x5a\xb7\x9f\xf2\x1a\x03\xb6\x27\x2f\x64\x40\xfd\xa6\x67\xb5\xbf\xf5\xb6\x4b\x45\x77\xe0\x67\xfa\x19\xc3\x1e\xce\x74\x38\xbb\xab\x50\x2d\x47\xad\xc8\x5b\x5b\x66\x2f\x11\xee\x58\x5b\x85\xbd\x18\xbf\x88\xce\x2d\xa0\x7b\xba\x5b\x48\xff\x04\xe5\xb5\xef\xbd\xe9\xf5\x4c\xea\xb7\xd9\x35\xdb\x4b\xae\x43\x05\x91\x4e\x14\x2b\x9b\x9a\x06\x65\xc9\x59\x02\xb6\x77\x84\x33\x98\x83\xbf\x6c\x55\xa1\x93\xd3\x54\x26\x55\x81\xc2\x9c\x8d\xdc\x10\x7b\x9a\x55\xc2\xf5\x1a\x72\x7a\x46\x56\x3b\xb5\xe5\xba\xd1\x2c\x4b\xf7\xe1\xd6\xd6\x3b\x32\xd9\x23\xb2\x7f\x25\x28\xe0\x1c\xf9\xb7\x92\x4b\x48\xf5\x98\x5c\xbc\xee\xa4\x29\x3e\x41\x81\x63\x42\x33\xa6\x8a\x7b\x50\x48\x0f\xc9\x12\x85\x60\xf0\x43\x01\x53\xbc\xc9\xab\x22\x16\xc0\xb8\x1e\x7b\x9f\x0f\xa9\x21\x49\xb0\x34\x98\xfe\xc9\x38\xea\x31\xd9\x71\xbf\x80\x24\x66\x02\xd4\xf2\xf5\x28\x66\x82\xee\xf0\xae\xdf\x6d\xb6\xeb\xb3\x77\x75\xb9\xf4\xb1\x6a\x97\xca\x7f\x03\x00\x00\xff\xff\xab\x1d\x20\x47\xfe\x18\x00\x00"

func templatesViewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsIndexHtml,
		"templates/views/index.html",
	)
}

func templatesViewsIndexHtml() (*asset, error) {
	bytes, err := templatesViewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesViewsLightHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x99\x4d\x6f\xe3\x36\x13\xc7\xef\xcf\xa7\x18\x10\xc1\x83\x16\xa8\xec\xbc\x00\x8b\xa2\x8d\x0c\x34\x45\xda\x5b\x0f\xc9\xa2\x7b\x5c\x50\xe2\xc8\x62\x4b\x91\x2a\x39\x4a\x9c\x1a\xf9\xee\x85\x48\x49\x91\x1c\x2b\x51\xd2\xd5\xd6\x49\x0e\x01\x29\x0d\x67\x86\x7f\xfe\x86\xb4\xcd\xed\x16\x04\x66\x52\x23\xb0\xd4\x68\x42\x4d\x0c\xee\xef\xff\x07\x00\xb0\xdd\x82\xcc\x60\x81\xd6\x1a\xdb\x3e\x3b\x17\xf2\x06\x52\xc5\x9d\x8b\x19\x57\x68\x09\xfc\xff\x48\x70\xbd\x46\xdb\x76\xa4\x2b\xa4\x73\x3c\x51\xc8\x56\x7e\x98\x1f\x9a\x54\x44\x46\x03\xdd\x95\x18\xb3\xd0\x61\xad\xaf\x54\x19\x87\x0c\x04\x27\xde\x0e\x6f\x02\x30\xe0\x56\xf2\x28\x97\x42\xa0\x8e\x19\xd9\x0a\xd9\xea\xff\x24\x0b\x74\x3f\x9e\x2f\x83\x9b\x55\x9b\xef\x4e\xb2\x4b\x21\x6f\xba\x77\xa8\x45\xfd\xa2\x37\x35\xae\x05\x2c\x50\x93\xa4\x3b\x58\x38\xe2\x84\xfb\xa6\x69\xcd\x6d\x7f\x16\xbd\x37\xa9\x51\x51\x21\xa2\x93\x53\xa8\x5b\xae\x68\x5b\x1b\x17\x9d\x9c\xf6\xc6\xec\x8e\xdb\x7c\x2e\xb9\x46\xb5\x63\xf1\xd8\x8a\x24\x0d\x14\x1c\x58\xe6\xa7\x2b\x3f\x61\x9f\xfe\xe2\x37\x5e\xd4\xc9\x9f\x2f\xf3\xd3\x11\xfb\x7e\xde\x0a\xb9\xcd\xe4\x86\xad\x7a\x0a\x0d\x8c\x47\x1e\x0f\xb2\x6b\x71\x19\x89\x97\x19\x5b\xb4\xc6\x75\x3b\xca\x8d\x95\x7f\x1b\x4d\x5c\x81\xef\x2b\x9e\xa0\x8a\x14\x66\xc4\xc0\x1a\x85\xc1\x8c\x41\x81\x94\x1b\x11\xb3\xd2\x38\x62\x20\x45\xad\x73\x51\x70\x2d\x18\x68\x73\xc3\x95\x14\x9c\x70\x7f\xd0\xdd\x24\x25\x61\x11\x82\xad\xad\xa9\xca\x91\x54\xbb\x91\x3e\xa3\xda\x3e\x66\x9e\x86\x07\x3a\x8d\x26\x6b\x54\x48\x19\x9a\x65\x3f\x6b\x57\xfd\x6c\x74\xd1\x77\xff\x6a\xec\x4e\xbe\xd7\xc0\xae\x83\xfb\x45\xcb\xdb\x68\x4a\x4b\x1f\xf2\x99\xbc\x1f\x23\xf9\xa1\xcd\xed\xc3\xe4\xdc\x76\x3d\x4d\x30\x87\x4e\xb4\x69\xb6\xde\x5e\xea\xb2\xa2\x66\x13\x48\x73\x4c\xff\x4c\xcc\xa6\x13\xfa\x0f\x17\xb9\x5b\x49\x69\xce\x40\xf3\x02\xdb\x75\x98\xec\xbd\xf9\xab\xa1\x69\x56\xb0\xd9\xc4\x7c\x6f\x71\xdd\xd4\x38\xf8\xc0\x28\xba\x5d\x01\x96\x13\x67\x3b\x65\x3d\x60\xbc\x82\x5e\x60\xf2\xc4\xeb\x76\x63\x0e\xa5\x7f\x5d\x95\xa5\xb1\xe4\x2e\xac\x5c\xe7\xa4\xd1\xb9\xa7\xa8\xfa\x32\xe5\x91\x74\xb1\xe6\xab\x91\x8b\x5e\x8c\x83\x2b\x94\x3e\xc5\xb6\x3e\xfb\x18\x14\x52\xc7\xec\x98\x41\xc1\x37\x31\x3b\x61\xe0\x08\xcb\x98\x2d\x8e\x4f\xd8\x60\x1b\x6c\x74\x0a\x3b\x5b\x4f\xc7\xa9\x8c\x87\xba\xe8\x2f\xc0\x0d\x57\x15\xc6\xac\x3e\x0d\x02\xe6\x03\x12\xd8\x73\x6c\xff\x2b\x0e\x9b\x43\x75\xe4\xed\x1e\x4a\xaf\xd6\xc9\xfc\x78\x5a\x14\xf3\x71\x79\x55\x3b\x7f\xaf\x40\x7a\xe5\x02\x62\x75\x73\x2a\x94\x8f\x10\xbc\x42\x31\x2f\x7b\xf0\xc5\x60\x59\x5b\x44\x3d\x1f\x2e\xbf\x06\xf7\xef\x15\x98\x46\xbd\x80\x4c\xe8\xbc\x1a\x1a\x2f\xd5\x5b\xc1\x26\x51\xd5\x8c\x1f\x10\x2f\xbc\xf7\xf7\x0a\x4d\xd0\xae\x39\xc9\xea\xf6\xab\x91\xa9\x75\x3a\xbc\x43\xee\x53\x2e\x09\x7f\xe7\x21\xb7\x99\x39\xbc\xad\x63\xcd\x07\xe2\xa7\xe0\xfe\xbd\x92\xd8\xa8\x17\x50\x0c\x9d\x57\xb3\xe8\xa5\x3a\x3c\x18\x7f\x36\xca\xd8\x8f\x58\x94\x68\x39\x55\xf6\x2b\x20\x99\xd6\x11\x23\x7a\x08\x39\x1f\x9e\x7e\x72\x30\x08\xf5\x5e\x51\x7d\xac\xea\xcb\xbe\x32\xec\x59\x95\x47\x08\xef\x61\xe5\x3f\xa6\x59\x23\x7c\xa3\x50\x77\x54\x5f\x66\x19\xa6\xe4\xbe\x85\xe3\xf9\x39\x46\x1f\x6b\x3e\x78\x2f\x1b\xff\x07\x47\xac\x43\x85\x29\x79\xea\x5a\x0d\x02\x43\x3b\x8a\xf4\x21\x85\x30\x68\x8a\xff\xed\x16\x7c\x1d\xc0\x51\xf6\x1d\x1c\x05\x9f\xf0\x43\xbc\xbb\xc6\xcf\xa9\xd2\xa5\x6b\x4a\x92\x46\xf7\x68\x6e\x7d\xde\xdf\xb3\x40\x11\xfe\xd5\x3d\x3b\x6a\x50\xbf\x6c\x4d\x5e\xf4\xe3\x52\x98\x24\x8a\x98\xb5\x2d\xd6\x51\xbc\x6a\x17\xb6\x0b\xe5\x7f\x91\x0d\xd9\x4d\x52\xe5\x89\x6a\x18\x4c\x78\x19\x82\x1f\xfe\x87\xe4\x4c\x71\x97\x47\x0a\xf5\x9a\xf2\xf9\xea\xe8\x97\x3a\x0a\x88\xca\xf2\x5a\xe9\x03\xac\xa7\xfe\x09\x40\xb8\xa1\x87\x03\x60\x74\xb7\x1f\x2a\x17\xaa\x6f\xf8\xac\xc1\xfd\xf8\x4d\x7c\x5b\x22\xcb\xb5\x93\xf5\xf2\xcc\x4e\xc3\xc7\x2e\xd4\xdb\x42\x62\x14\x85\x3d\xda\x05\x1e\xf6\xbc\x68\xa1\x70\x5f\x91\x0a\xa5\x3f\x3b\xa3\xa4\x18\xbd\x4e\xda\x37\x68\x3a\x45\x63\x9a\x17\x22\x32\x59\xe6\x90\xa2\xb3\x29\x82\x37\xb7\x8f\xfe\x6a\x00\xb5\x60\x8d\xf6\xae\x4a\x0a\xf9\xa0\x7e\x42\x1a\x12\xd2\x91\xab\xd2\x14\x9d\x63\xdd\x9e\xce\x7e\x2a\x4b\x75\x17\x40\x1a\x5e\x3b\xbe\x56\xc1\xd7\x7e\x2e\x3a\x5f\xd6\xda\x4d\xba\xb4\xdb\x79\xd4\xeb\xee\xbd\x1a\x7d\x68\xfd\x13\x00\x00\xff\xff\x9c\x9f\xf0\x8a\x16\x1e\x00\x00"

func templatesViewsLightHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsLightHtml,
		"templates/views/light.html",
	)
}

func templatesViewsLightHtml() (*asset, error) {
	bytes, err := templatesViewsLightHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/light.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesLightMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x50\xcf\x8b\x1c\x45\x18\x7d\x9d\x19\x7f\x64\x8c\xbf\xe2\xc1\x4b\x0e\xe5\xc1\xa8\x48\xc7\xe9\x8d\x88\xf4\x6e\xef\xea\x86\x5d\x11\x77\x24\x8c\xa3\x1e\xa5\x32\x5d\x3b\xd3\xa6\xa7\x6a\xa8\xee\x96\x2c\x04\xdc\x9d\xac\x3f\x20\x81\x45\x44\x05\x8d\xc8\x22\x88\x17\x71\x76\x49\xeb\x66\xb3\x99\x3d\x79\xf3\xf0\xd5\x41\xf0\x20\x5e\x04\xff\x03\xf1\x24\x48\x75\xf7\xb8\x78\x4b\x5d\xea\x7b\xf5\xbd\xf7\xbe\xf7\xd5\xef\x27\xeb\x9f\x00\xc0\x43\x00\x4e\x01\xb8\x0e\xe0\x51\x5b\x3b\x28\xce\x25\x07\xb8\x0b\xc0\x9a\x03\xd4\x01\xbc\xeb\x00\x0d\x00\x9b\x0e\xf0\x30\x80\x2d\x07\xb8\x1b\xc0\xb6\x03\x3c\x00\xe0\x5b\x07\xb8\x07\xc0\x8f\x95\xee\xb0\xea\xff\x54\xe9\x7f\xae\xb0\x71\x80\x1a\x80\x5f\x2b\xfc\x9b\x03\x3c\x02\xe0\x8f\x4a\xf7\xb7\x03\x2c\x38\xc0\x3f\x0e\x70\x12\x40\x50\x2b\xfd\xcf\x57\xf7\x5b\x35\xe0\x09\x00\xb2\x06\x9c\x00\xf0\x69\x0d\xf0\x00\x7c\x5d\x2b\x77\xf9\xa5\xe2\xfd\x59\xf5\xff\xaa\x95\x73\x8e\xd5\x4b\xbf\xe3\xf5\xb2\x7f\xaa\x5e\xee\xf3\x74\x1d\x38\x0b\xe0\xb9\x0a\xbf\x5d\x07\xec\x17\xd8\x2c\x8d\xf2\x2b\x70\xac\xba\xad\xce\xee\x62\xf7\x7c\xb0\x7a\x3b\x81\xa3\x63\xf7\xba\x0f\xc0\xfd\xd5\xbf\xda\x73\x1c\xe5\xfc\x7b\xa7\xa4\x17\x87\xc3\x78\x0d\x8b\x71\x26\xb0\xa8\xa3\x5e\x3f\x95\x22\x49\x70\x4e\xc5\x4a\xb3\x54\x0c\x86\x42\xf3\x34\xd3\x02\x4b\xab\xab\xa2\x9b\x62\x39\xe6\x49\x9f\x85\x99\xe6\x69\xa4\x24\x96\xe3\xa8\x7b\x51\x68\xbc\xa4\x85\x90\x58\xe1\x83\x0b\x21\xc7\xab\x4a\x0a\xb4\xb9\x0c\xd5\x00\x6d\x11\xe2\xb5\x54\xab\x0b\x02\x1d\xcd\x65\x12\x59\xd9\x91\xfe\xcd\x7e\x94\x0a\xb4\xc5\x50\xe9\xd4\x6d\x25\xbd\x28\x74\x17\xb3\x5e\xe2\x76\x94\xcf\x42\xf1\xce\x0b\x17\xa3\x3e\x1f\xa8\x33\x3a\x6b\xac\xf0\x24\x75\x0b\x8b\x98\xa7\x4a\xfb\xec\x95\xa2\xc5\x5a\x99\xe6\x03\x15\x2a\x36\xf7\x3f\xfe\x7c\x63\x85\xcb\x5e\xc6\x7b\xc2\xed\x08\x3e\xf0\xd9\x7f\xd8\x67\xed\x2c\x49\x22\x2e\x1b\xad\x97\x5b\x4b\xee\x1b\x42\x27\x91\x92\x3e\xf3\xce\x34\x1b\xe7\x94\x4c\x85\x4c\xdd\xce\xda\x50\xf8\x2c\x15\x97\xd2\x67\x86\x31\x8f\xe4\x2c\xeb\xf6\xb9\x4e\x44\x1a\xbc\xde\x59\x76\x9f\x3f\xe2\xd9\x3c\xab\x42\xbb\x4b\xb2\xab\xc2\x48\xf6\x7c\xd6\x38\x1f\x67\x9a\xc7\xee\xb2\xd2\x83\xc4\x67\x72\x58\xc0\x24\x38\x3b\xcb\xca\x32\x90\x8f\x7b\xcd\x20\xf0\xd8\xe9\xd3\xcc\x96\xcd\xc7\x02\xcf\x63\x0b\xac\xc9\xfc\x02\xcf\x07\x33\xd3\xd6\x5c\xf0\xac\x2d\x9f\x2c\x68\x73\x5e\x93\x5d\xbe\x5c\x4a\xe6\x83\x99\xe6\x53\x6c\x81\x79\xcc\x67\x33\xb3\xa0\xaf\xcc\x3a\xed\xd1\x01\xe5\x74\x9b\xf6\xcc\xc8\x5c\x03\x7d\x4c\x13\xba\x65\xae\xd0\x0e\x4d\xe8\x26\xe8\x7b\xb3\x4e\xfb\x34\x31\x1b\x65\xf7\x1b\xda\xa5\xdc\x8c\x68\x42\xbb\x34\x36\x5b\xcc\x8c\x28\xa7\x03\x3a\xa4\xdc\xac\xd3\xd8\x8c\xcc\x15\x7b\x83\xbe\x33\x9b\x66\x93\x72\xda\x37\xa3\x72\xca\x84\x6e\x58\x63\xfa\xc1\xce\xa1\x9c\x6e\x99\x6b\x74\x7b\xea\xcb\x68\xd7\x6c\xd0\xa1\xb9\x6a\x3e\xa4\x7d\xda\x03\x5d\xb7\x7e\xe6\x7d\x1a\xdb\x5c\x94\x83\x3e\xb3\x12\x1b\xd3\x5c\xb5\xa9\xbe\x30\x5b\x74\x40\x3b\x74\xc3\xce\xfa\xd2\x26\x02\x6d\xdb\xd8\xe6\x03\x1a\xd3\xcd\x29\xed\xf3\x22\xd4\xc6\x14\x6e\x9b\x51\x91\x64\xe7\x8e\x22\x95\x3b\xe5\xe6\xbd\x82\x38\x06\x7d\x54\x70\xac\xd3\xbf\x01\x00\x00\xff\xff\x8e\x34\x3d\xdf\x75\x04\x00\x00"

func localesRuLc_messagesLightMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesLightMo,
		"locales/ru/LC_MESSAGES/light.mo",
	)
}

func localesRuLc_messagesLightMo() (*asset, error) {
	bytes, err := localesRuLc_messagesLightMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/light.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x4d\x6b\x13\x41\x18\xc7\xff\xbb\xa9\x6f\xd1\x83\x16\x6f\x7a\x78\x7a\xb0\xd8\xc3\xd6\x6c\x55\x90\x6d\xb7\xf5\x25\x2d\x96\x36\x56\x6b\xec\x41\xf0\x30\x26\x93\x74\x35\x99\x09\xbb\x13\xb1\xd0\x8b\xbd\x08\x2a\x16\x8a\x82\x08\x1e\xf4\x03\x88\x16\x02\xb1\xb5\xf1\xe2\xcd\xcb\xb3\x77\xf5\x33\xf8\x11\x64\xb3\xeb\x5b\x71\x2e\xfb\xfc\x5f\x66\x9f\xdf\x7c\x1f\x1c\x78\x0e\x00\x87\x00\x1c\x07\xb0\x0a\xe0\x08\x80\xaf\x48\x0f\x59\xc0\x3e\x00\x43\xd9\x77\xc4\x02\x72\x00\x5c\x0b\x18\x00\x70\xd6\x02\x0e\x00\xf0\x32\xbf\x68\x01\x36\x80\xcb\x16\xb0\x07\xc0\x5c\xd6\xbb\x96\xe5\x4b\x59\x7e\xd3\x02\xa6\x2c\xe0\x96\x05\x1c\x06\xf0\xca\x4e\x19\xde\xda\xc0\x7e\x00\xdb\x36\xb0\x17\xc0\x67\x3b\xe5\xf9\x62\x03\x83\x00\xbe\xd9\x69\xff\x47\xa6\xad\x5c\xda\x3b\x9a\x4b\xf5\xb1\x5c\x9a\x8f\xe6\x00\x0b\xff\x3f\xc9\x9e\x64\x47\xf2\x9e\x3c\x52\xbe\x84\xf5\x20\x52\xc6\xe4\x7f\xc9\x9b\x12\x4e\x5c\xa8\x98\x40\xab\x08\x45\x69\x44\xd0\x88\x30\x27\x57\x70\x45\x34\x25\x16\x6e\xdf\x91\x15\x43\xb3\x45\x2c\xd4\x6a\x58\x50\xb8\x6e\x84\x91\x28\xaf\xb4\x24\x74\xad\x06\xad\xb0\x28\x5b\x3a\x34\x4e\x29\xaa\x07\x55\xe7\x62\xbb\x1e\x39\x65\xed\x51\x55\xde\x3b\x7f\x37\x58\x16\x4d\x3d\x1a\xb6\xf3\xf3\x22\x32\x4e\x39\x14\x2a\x6a\x08\xa3\x43\x8f\xe6\xfa\x11\x95\xda\xa1\x68\xea\xaa\xa6\x89\x7f\xfa\x93\xf9\x79\xa1\xea\x6d\x51\x97\x4e\x59\x8a\xa6\x47\xbf\xb5\x47\x8b\xed\x28\x0a\x84\xca\x97\x66\x4b\xd3\xce\x92\x0c\xa3\x40\x2b\x8f\xdc\xd1\x42\xfe\x92\x56\x46\x2a\xe3\x24\x70\x1e\x19\x79\xdf\x9c\x6a\x35\x44\xa0\xc6\xa9\xb2\x2c\xc2\x48\x1a\xff\x46\x79\xc6\x39\xf7\xa7\x97\xf0\xd4\x64\xe8\x4c\xab\x8a\xae\x06\xaa\xee\x51\xfe\x6a\xa3\x1d\x8a\x86\x33\xa3\xc3\x66\xe4\x91\x6a\xf5\x65\xe4\x9f\x1e\xa7\x74\xf4\xd5\x09\xb7\xe0\xfb\x2e\x0d\x0f\x53\x32\x16\x86\x7c\xd7\xa5\x29\x2a\x90\xd7\xd7\x93\xfe\xd8\xaf\x68\xc2\x3f\x93\x8c\x27\xfb\xb5\x09\xb7\x40\xab\xab\xe9\x95\x49\x7f\xac\x30\x42\x53\xe4\x92\x47\x63\xe3\xe0\x67\xdc\xe1\x0f\xf1\x83\x78\x8d\x37\xb9\x1b\xaf\xf7\x8d\x78\x8d\xdf\xf1\x36\x77\xc1\x2f\x79\x3b\x7e\x1a\x3f\x04\xbf\xe0\x8f\xf1\x3a\x66\x8b\xc4\x3d\x7e\x1f\x3f\xe2\x0e\x6f\x25\x2d\xf0\x46\xfc\x98\xb7\xd2\x16\x77\xe3\xb5\xf8\x09\x78\x63\xb7\xf1\x9a\x7b\xfd\x0d\xbd\x78\x9d\x77\xb8\xcb\x1d\xf0\x1b\xee\xf2\x27\xf0\xe6\x5f\xb7\x3b\xbc\xc3\x3d\xf0\xe6\x2e\xe3\x67\x00\x00\x00\xff\xff\x4b\xef\x56\x03\x3f\x03\x00\x00"

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
	"templates/views/index.html": templatesViewsIndexHtml,
	"templates/views/light.html": templatesViewsLightHtml,
	"locales/ru/LC_MESSAGES/light.mo": localesRuLc_messagesLightMo,
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
				"light.mo": &bintree{localesRuLc_messagesLightMo, map[string]*bintree{}},
				"widget.mo": &bintree{localesRuLc_messagesWidgetMo, map[string]*bintree{}},
			}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{templatesViewsIndexHtml, map[string]*bintree{}},
			"light.html": &bintree{templatesViewsLightHtml, map[string]*bintree{}},
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
