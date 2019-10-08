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

var _templatesViewsIndexHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x18\x4d\x6f\xdb\x36\xf4\xde\x5f\x41\x10\xd9\x90\x00\x95\x95\xa4\x40\x31\xb8\x72\x86\x62\xc1\xb0\x02\x59\x5b\x74\xe9\xb9\x78\x92\x9e\x2c\xba\x14\xa9\x90\x94\x63\xd7\xf0\x7f\x1f\x48\x49\x96\x6c\x4b\x8e\xec\x25\xbd\x2c\x87\x96\x7c\x7c\xdf\xdf\xd6\x6a\x45\x62\x4c\x98\x40\x42\x23\x29\x0c\x0a\x43\xc9\x7a\xfd\x2a\x88\xd9\x9c\x44\x1c\xb4\x9e\x50\x25\x1f\xe9\xcd\x2b\x42\x08\x69\x43\x23\xc9\xbd\x2c\xf6\xae\xae\x89\x3d\xe9\xac\x3e\x2d\xb4\x77\x75\x5d\xe1\xef\xd2\x2c\xbe\xe5\x20\x90\xb7\x5e\xf7\x31\x6a\x2d\xb6\x71\x1c\x9e\x81\x90\x63\x8d\x59\x5e\xdc\xbf\x9e\x36\x8a\xe5\x18\x93\x18\x0c\x94\xf0\xd8\x78\x0a\x75\x2e\x85\x66\x73\x24\x42\x3e\x2a\xc8\x29\xd1\x66\xc9\x71\x42\x1f\x59\x6c\xd2\xf1\xd5\xe5\xe5\x2f\x1d\x52\x4a\x49\x29\x42\xdc\xfd\x56\xbe\xab\xfe\xc7\x8a\xc1\xcd\x6a\x45\xd8\xd5\x6f\x82\xd0\x4f\xe1\x0c\x23\x43\x3e\xdc\x52\x32\x22\xeb\x75\xe0\x9b\xf4\x08\xea\xfb\x65\x8e\x27\x11\x7e\x84\xec\x34\xc2\x7f\x0c\x98\xd3\x28\xdf\x47\x86\x49\xa1\x87\xd0\x06\x7e\x9f\x0f\x2d\x5d\xaf\xf7\x03\x13\xca\x78\xd9\xcf\x76\xb5\x22\x0a\xc4\x14\xc9\x19\x7b\x4d\xce\x50\x18\x66\x96\x64\x3c\x21\x23\x77\x64\xa8\x6d\x72\xf7\xdb\xf2\x64\x58\x63\x6b\x6c\xc5\x77\x54\x06\xf6\xc3\x6d\x69\xec\x81\x7c\xe9\xa0\xb5\x61\x3d\x85\xce\x46\xf5\x18\x3a\x17\x98\x9a\xd8\x45\x96\x9c\x0d\xa6\x3f\x88\x40\x76\x4a\x37\x34\xc2\x9b\x2a\x59\xe4\x94\x28\x69\xcb\xac\xbc\x3c\xcd\x84\x94\x71\x63\x09\xc1\x87\x6d\xff\x50\xce\xa6\x69\xd9\x90\x86\x30\x09\x80\xa4\x0a\x93\x09\xb5\xfe\x1a\x7d\xc1\x87\x02\xb5\x19\x7d\xfd\x72\x37\xfa\x0c\x26\x25\xeb\xf5\xef\xe0\x12\x74\x52\xca\xf8\x55\xba\xf8\x4d\xba\x23\x4a\x5b\x76\x11\x6b\x1b\x13\x89\x2c\x0f\x91\x2c\x21\x0b\x3d\xd0\x3a\xa7\x1c\xab\x19\x4e\xf9\x32\x4f\x1d\x93\xcd\xc9\xc3\x25\x7a\x32\x47\x41\x89\x61\xc6\x3a\x6f\x53\x53\xb7\x68\x80\x71\x4d\x5d\xd8\xe8\x4d\xe0\xb3\x61\x32\x03\x1f\x86\x2b\xb7\xe7\xfe\x32\x53\xa8\x14\x83\x9d\xbf\x91\xfb\xa2\x41\x88\x6d\x71\xab\xff\x10\x86\x8d\x9a\x07\xc3\xa1\x8d\xcc\xf7\x43\xf1\x29\x49\x8e\x0e\xc3\x46\xde\x91\xe1\x40\xae\xf1\x7f\xe9\xfa\x9c\xc3\xb2\xc3\xf5\xe2\xe7\x79\x5e\xc4\x43\x1d\x3f\x1c\x3b\xf0\x63\x36\x7f\xa2\xe1\x1e\x6c\xc9\xfd\xf3\xf2\x69\x45\x02\xbf\x67\x6a\x06\xbe\xdb\x96\x76\xd6\xb1\x6d\x55\x5b\xd7\xea\x58\xfd\xf7\xb2\x3b\xa2\xc5\xe3\x90\x6b\x8c\x0f\x6e\x8b\x2e\x4f\xba\x76\xc5\xf4\xba\xb5\x80\xdd\xbf\xaf\x57\x92\xf4\xba\x03\xb7\xe0\x35\x43\x01\x73\x22\x60\x1e\x82\xf2\x94\x1d\x3d\xc4\xe9\xf2\xcd\x48\xc9\x43\xb9\xe8\xdb\x16\x39\xbb\x09\xa0\x65\xb3\xd3\xdb\xe3\x4c\x7c\xa7\x37\x4d\xb2\x27\x40\x12\xf0\xa2\x14\xe7\x4a\x0a\xcf\x4e\x46\x9b\xca\x36\x3b\x03\x9f\x77\xe4\x74\xe0\x17\xbc\x03\xda\x76\x2f\x47\x50\x09\x5b\x58\x4e\x7b\xe9\xd5\x05\x7a\xb6\x45\xfb\xf8\x85\xba\x7f\x6d\x3b\xb8\x75\xed\x2d\xa6\xc5\xb0\xed\xb2\x5a\x7e\x12\x32\x92\x06\xbe\xa9\x42\x08\x26\xa6\x64\xbd\xde\xf0\xaa\x40\x25\xb3\xa6\xdf\x36\xef\x08\xf1\xb2\x79\x75\xc5\xd5\x5f\xa1\xfd\xd5\xd9\xd2\x02\x95\x92\xaa\xb7\x46\x07\x7b\xe1\x0e\xb4\x21\x8e\xd7\x11\x9e\xd8\x52\xe0\x44\x3b\x0e\x75\x98\xc1\xda\x7f\x56\x72\xaa\x50\x0f\x8e\xe2\xe1\x96\xd9\x4a\xe8\x7c\xc3\xb8\x4a\xca\x0c\xd4\x94\x09\x2f\x94\xc6\xc8\x6c\x7c\x39\x60\x50\x75\xb1\xf3\x42\x50\xa4\x7d\xd9\xd4\x40\x67\x7e\x6d\x61\xba\x8d\xd1\x8e\xdc\x39\x36\x19\xb6\x85\x51\xce\xd4\xd2\xb7\x76\xe4\x96\x8b\x73\x8d\x12\x82\xa2\x04\x14\x03\x6f\x0e\xbc\x40\x21\x1f\xdd\x44\x74\x22\x1f\x15\x33\x06\x85\x1b\xd4\x0d\x4a\xc6\xc4\x84\x5e\x6e\x41\x60\xd1\x10\x19\x69\x80\x3b\x92\xad\xc2\xad\x9f\x6b\xb9\x64\xbd\xee\x2b\xe4\x3d\x97\xe9\x1c\x44\xcd\x8d\x33\x81\x5e\x8a\xb6\x71\x8e\xaf\x2f\xf3\x05\xbd\xd9\xd7\x96\x84\x4b\x83\x9a\xc8\x84\xec\x2a\x15\xf8\x96\xd7\x80\x28\x0d\x19\xa5\x07\x51\x4e\xc9\xff\xe1\x19\xfe\x47\x8a\xd1\x77\x5d\x64\xc7\x56\x67\x54\xd1\x9d\x58\xa0\x4f\x0f\xf9\xfd\x97\x44\xaa\x6c\x33\x9e\xa4\xca\xbc\x54\x2a\xf6\x43\x0a\x1b\x10\x77\xe7\x10\x22\xf7\x38\x26\x86\xc4\x4a\xe6\x3f\xa4\xc0\x3a\x49\xed\x3b\x25\x19\x9a\x54\xc6\x13\x9a\x4b\x6d\x28\x61\xf1\x84\x4a\x03\x94\x54\x7b\x26\xad\x17\x4e\x07\x14\x72\x0e\x9c\xc5\x60\xb0\xc7\x82\x56\xf9\x31\x83\x59\xa9\xc2\x53\xbf\x22\x03\xa7\xa3\xc5\x9d\x50\x8e\x73\xe4\xb4\x99\xc4\xc2\x28\xc9\x4b\x23\x48\xb5\x8b\xbc\xa9\x57\x91\x37\x9d\x9b\xc8\xee\x5f\xf3\x19\x86\x65\x28\x0b\x53\x06\xf5\x40\x66\x39\x61\x07\xb4\xdd\xdf\x8e\xde\xd6\x1a\xbd\x1d\xa4\x91\xe3\xc2\x44\x5e\x18\x62\x96\x39\x4e\xa8\xc1\x85\xa1\x5b\x51\xac\x0c\x2f\xe3\x61\x6a\xc5\x05\x64\xd8\xba\xba\xee\xd0\x6a\x0d\x25\xdc\x35\x07\x85\x0f\x05\x53\x18\x4f\x68\x7d\x3a\xe4\xff\xfe\x5a\xeb\x79\x0a\x7c\xab\xe5\x69\x4b\x67\x33\x86\x5e\xb5\xbe\x5c\xa6\x08\xf1\xe6\x87\xea\x6a\x45\xb4\x01\xc3\xa2\xbf\xee\xff\xbe\x23\xe7\xe5\xf9\xeb\x97\x3b\x42\xfd\x18\x74\x1a\x4a\x50\xb1\x0f\x5a\xa3\xd1\xfe\x1c\x45\x2c\x95\xf6\x37\x1f\x0e\xf5\x48\xa0\xf1\x42\xed\x47\xba\x84\xde\x97\xd0\x50\x4a\xa3\x8d\x82\x7c\x94\x31\x31\x8a\xec\x98\x49\x80\x6b\xbc\x78\x46\xa9\xcd\x07\xcb\x5a\x81\x06\xf2\x42\x0a\x54\x35\x3d\xab\xec\xad\xae\x5d\x22\xba\x1d\x3f\xd3\xcf\xe8\x76\x7f\xa6\xfd\xd9\x43\x81\x6a\x39\x6a\x79\xde\xea\x32\x7b\x09\x77\x87\xda\x0a\xec\x8d\xf1\x8b\xc8\x6c\x02\xba\x23\xbb\x15\xe9\x9f\x20\xbc\xb2\xbd\x37\xbd\x9e\x49\x7c\x93\x5d\xb3\x9d\xe4\xda\x17\x10\xe8\x48\xb1\xbc\xee\x69\x90\xe7\x9c\x45\x60\x67\x87\x3f\x83\x39\x94\x8f\xad\x2e\x74\x76\x1e\xcb\xa8\xc8\x50\x98\x8b\x91\x5b\xd5\xcf\x93\x42\xb8\x59\x43\xce\x2f\xc8\x6a\xab\xb7\xdc\xd6\x92\x65\xee\x3e\x4f\xdb\x7e\x47\x26\x3b\x48\xf6\x2f\x07\x05\x9c\x23\xff\x9a\x73\x09\xb1\x1e\x93\xab\xd7\x9d\x38\xd9\x47\xc8\x70\x4c\x68\xc2\x54\xf6\x08\x0a\xe9\x3e\x5a\xa4\x10\x0c\x7e\xc8\x60\x8a\xf7\x69\x91\x85\x02\x18\xd7\xe3\xd2\xe6\x7d\x6c\x88\x22\xcc\x0d\xc6\x7f\x32\x8e\x7a\x4c\xb6\xcc\xcf\x20\x0a\x99\x00\xb5\x7c\x3d\x0a\x99\xa0\x5b\xb4\xeb\x77\x9b\xeb\xfa\xe2\x5d\xd5\x2e\x4b\x5f\xb5\x5b\xe5\xbf\x01\x00\x00\xff\xff\x4e\x73\x31\x83\xe4\x19\x00\x00"

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
