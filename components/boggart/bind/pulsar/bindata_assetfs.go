// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package pulsar

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\xff\x6e\xe3\x36\x12\xfe\x7f\x9f\x62\xa0\xdb\x9e\x6d\xdc\x5a\xf2\xaf\xf6\x0a\xaf\x9d\xc3\x75\xd3\x76\x17\xd8\x74\x83\x6d\xda\x43\xae\x28\x02\x5a\xa2\x6d\xed\x4a\xa4\x4a\x52\xc9\xe6\x02\xbf\xd3\x3d\xc3\x3d\xd9\x81\xa4\x24\xcb\xb6\xec\x48\x8a\x95\x28\xae\x09\xb4\xab\x48\xe4\xf0\xe3\xcc\x37\xc3\x91\x4c\x4a\x77\x77\xe0\xe0\xa9\x4b\x30\x18\x36\x25\x02\x13\x61\xc0\x62\xf1\x62\xc4\x6d\xe6\x06\x02\xc4\x6d\x80\xc7\x06\x0a\x02\xcf\xb5\x91\x70\x29\xb1\x3e\xa1\x6b\xa4\x2f\x1a\x27\x2f\x00\x00\xa6\x21\xb1\xe5\x15\xb8\x46\x5e\x88\x7f\xa0\xcc\x47\xa2\x49\x5e\x01\x0d\xe4\x59\xde\x82\x3b\x55\x4d\x16\x77\x0a\x4d\x29\x91\x4e\xe3\xab\x30\x1e\x8f\xa1\x11\x12\x8d\xc1\x69\xa4\x6b\xcb\x92\x54\x5b\x3b\x2f\x8b\xef\x12\xd7\x0f\xfd\x1f\x18\x52\xfd\x9f\xba\x33\x57\xf0\x21\xf4\x5f\xad\xd4\x5c\xbc\x58\x1e\x25\x87\x0c\x8b\x90\x11\x20\xf8\x06\xde\x11\xe1\x99\x3f\x85\xfe\x04\xb3\x08\x7b\x82\x66\x39\x06\x73\x1a\x0d\xab\xf5\xfa\x85\x16\x35\xb2\xb4\x12\x4e\x5e\x8c\x1c\xf7\x1a\x6c\x0f\x71\x3e\x36\x18\xbd\x89\xb4\x92\x3e\x6b\x53\xaf\xed\x3b\xed\x6e\x0f\xe4\x11\xf7\xe3\xa3\x2f\xbc\xdd\xed\x45\xf5\xd7\xdb\x7c\xb9\x0a\x10\xc1\x5e\xea\xea\x66\x8d\xd8\x60\x27\x1b\xaa\x51\xf5\x18\xf5\xf0\xd8\x10\x68\x92\x25\x29\xa9\x19\x7a\xb1\x40\x82\xae\x81\xa0\xeb\xb6\x40\x13\x0e\x13\xc4\xae\xe4\x81\xb1\x14\xe3\xb9\x3c\xab\xaf\x44\x92\xe7\x46\x75\x03\x86\x39\x26\x42\xf1\xc5\xb8\xbb\x93\x66\xc7\x7f\x80\xa9\xcd\x04\x86\x24\x58\xdc\xa7\x3c\x77\x8d\x65\x25\x4c\x1c\x58\x2c\x4e\x46\x08\xe6\x0c\x4f\xc7\x5b\xda\x2d\xd9\x37\xbc\xa6\xae\xd3\xec\xb4\x5e\xcb\xb6\x1e\xc7\xb0\x58\xdc\xdd\x81\xf9\x11\xff\x11\x62\x2e\xcc\x5f\x3e\xbe\x37\xcf\x91\x98\xeb\xd3\x5a\xb8\x71\x22\x85\x76\xbf\x25\x60\xbc\x09\x19\xc3\x44\x80\x8f\x11\x0f\x19\xf6\x31\x11\xdc\x00\x13\x16\x8b\x91\x85\x4e\x46\x96\xe7\xee\x61\xa0\x88\xd9\x73\x39\xba\x92\xe3\x4d\x35\x2f\x33\xec\x7f\x68\x39\xe3\x48\x4c\x96\x16\xfe\x19\xf7\x70\xff\xc0\x47\x56\xe8\x6d\xb9\x92\xe2\xa4\x40\x93\xf6\x76\x56\xee\x62\x67\x5a\x82\x3c\x03\x53\xe4\x60\xd0\xea\x02\x97\xec\x90\x26\xcb\x4e\xe5\xed\x6c\x99\x46\x7f\x83\x3d\xcf\x00\x2e\x6e\x25\x32\x7a\x8d\xd9\xd4\xa3\x37\x43\x40\xa1\xa0\xf7\xf4\xaf\x24\xc9\x18\x11\x8b\x92\xc7\xed\x39\x65\xee\x7f\x28\x11\xc8\x8b\xbd\x48\x9e\x36\xc0\xc7\x62\x4e\x9d\xb1\x31\xc3\xc2\x80\xc8\x4a\xc6\x16\x2b\x1a\xe0\x3a\xe3\x78\x34\x6d\xdd\x9e\xd0\x6b\xe4\xb9\x0e\x12\xf8\x7e\x54\x0a\x99\x4b\x82\x30\x8e\xe5\x73\xd7\x71\x30\x31\x74\xb4\x1e\x2f\xf5\x44\x90\x8f\x35\x3f\x29\x31\xc0\x7a\xb8\xe4\x58\xe4\x94\x51\x7f\xaf\x02\x05\xcd\x2f\x6e\x33\x0a\x0f\x72\x98\x32\x2e\x89\xa3\x9c\x22\x81\x81\x21\x32\x8b\x7c\x25\xb7\x84\xd1\xd4\xc5\x9e\xc3\xb1\xc8\xdf\xe9\x26\x6e\x22\x18\xf5\xda\x33\x46\xc3\xa0\x00\xf8\x1d\xb2\x78\x09\x31\xeb\xa2\x94\xa1\xda\x01\xc3\x81\x0c\x2b\xfa\xaf\xb2\x10\x13\xf9\x3c\x40\x24\x09\x95\x8e\xd3\xa6\x24\x2d\xb9\x8d\x1c\x87\x12\xe3\x64\xe4\x26\x6e\x86\x60\x8a\xda\x36\xf2\x30\x71\x10\x33\x4e\x46\x96\x7b\x32\xb2\xa4\x94\x07\x80\x48\x53\x50\xe0\x2f\x62\xd5\x05\xa5\xdb\x19\x2b\x6e\x1e\x29\x35\xe1\xa9\xf4\x64\x59\xeb\x4a\x32\xdf\xd4\x49\x05\x18\xbd\x4e\xe7\x1b\xb3\xd3\x35\x3b\x3d\xe8\x7e\x3d\xec\x0c\xd4\xac\xd0\x86\xa4\xb2\xa0\x3b\xab\xe6\xa6\xfc\xc6\x70\x2c\xc7\xbd\x2e\xc1\x9a\xe2\xcd\x0a\x36\x19\x59\xc5\x7c\xa3\x80\xf8\x7d\xb9\xfd\x39\x66\x2e\x75\x8e\x2e\x7f\xaf\xcb\xc7\x93\xe6\x8d\xeb\x88\xf9\x10\xba\x9d\xce\x57\x0f\x0a\x03\xd8\xc3\xb6\x50\x6e\x17\x44\x26\xd0\xa1\x3f\xfe\x2b\xc3\xfd\x40\x37\xea\x19\xc0\xf0\x1f\xa1\xcb\xb0\x33\x36\xe2\xa3\x07\x60\x51\x78\xf4\x6d\x40\xec\xdf\x3e\x25\x62\xee\xdd\xa6\x32\x36\x0d\x0b\x92\x2b\xd2\xb5\x35\x1c\x89\x22\x3e\x4a\x65\x7d\x09\xc3\xce\xe2\x16\x3a\x03\xd3\x1d\xed\x15\xad\x83\xdc\x4c\xac\xfa\x7c\x7e\xa4\xa7\xba\x7e\x65\x38\xe7\x34\x64\x99\x40\xa3\x0b\xf9\x91\xbe\x8d\x1a\xec\x07\xea\xc8\xd2\x7d\x1d\xa3\xef\x03\xa2\x6f\xaf\x80\x07\xfe\x95\x4c\x78\xf0\xfa\x18\x6e\xab\xc9\xb0\x26\xa1\x10\x94\x44\xd9\x0d\x0f\x27\xbe\x2b\x92\x70\x3a\x11\x04\x26\x82\xb4\x79\x68\xdb\x98\xf3\xd4\x8d\xe2\xcf\x73\x7a\x13\x3b\x94\x96\x70\x74\x87\x3c\xe2\x47\x96\x9c\xa2\x76\xd7\xcb\x21\x6a\x24\xd0\xc4\xc3\xa9\x5b\x64\x0f\x83\xfa\x7f\x9b\x0b\xe6\x06\xd8\x01\x07\x09\xa4\xcf\x3b\xa2\xcd\x30\x0f\x28\xe1\xf2\xce\x99\xd0\x1b\x86\xd6\xe7\xe7\x9c\xd3\xf3\x48\xcc\x31\x72\xf2\xd4\x63\x39\x95\x26\xe6\x27\x2b\xb7\x54\x31\xa5\xc4\xbc\x84\x80\xef\x09\x66\xb3\x38\xcc\x4f\x18\x58\xcb\x4b\x3f\xda\xf2\x96\xbb\x98\x6c\x3d\xf1\xcc\x04\x98\xdf\xb9\xc4\x31\xdf\x50\x32\x75\x67\xe6\x3b\xe9\x75\xfc\x0d\x0d\x89\x80\x8e\x94\x97\x46\x70\x1e\x7a\x1c\x6b\xc7\x84\xaf\x54\x9a\x48\x5c\x2f\xf9\xaf\x9b\xc6\xc5\xd0\x0d\x34\x75\x2b\x7f\xc4\xc3\xe0\xa4\x3f\xb2\xe4\x3f\x60\x41\x20\xa5\x70\x03\xcc\x56\x8c\x37\x99\xd3\xf6\x04\xbc\x5b\x0c\x78\xaf\x36\xc0\x7b\xc5\x80\xf7\x6b\x03\xbc\x5f\x0c\xf8\xe0\x11\x80\x8f\xac\x3c\x6e\x2a\x85\xe6\x74\xfa\x09\x75\x6e\xef\xaf\xa7\x06\x44\x66\x18\x5e\xba\xaf\xe0\x25\x17\x48\xc0\x70\x0c\xa6\x3c\xe0\xb9\x60\xe7\x0f\x2e\x39\x50\xa7\x50\xe9\x2c\xf3\x65\x56\xee\x5e\x68\xa2\x89\x4d\xac\xc6\x66\xca\xa8\x66\xaa\xa4\xde\xfc\x59\x30\x97\xcc\xe0\xa5\xb2\xb0\xe1\x20\x3e\x9f\x50\xc4\x1c\x95\xc8\xde\xdd\xa5\xeb\x5f\x62\xc4\x8a\x74\x1b\x3f\xfa\x5d\x1f\x42\x92\xd2\x17\x1d\x40\x0a\xcb\x29\xba\x8d\x00\x3e\xd5\xa8\x0a\xc2\xcf\xf7\x4b\x95\x43\xed\xd0\xc7\x44\x98\x37\xcc\x15\xb8\xe9\x20\x81\x2f\xa8\x1e\x4b\xb3\xb1\x8a\x3b\xfd\x14\xa6\xdd\xe9\xb6\x3b\xbd\x0b\xf5\x14\x66\xd8\xf9\xba\xdd\xf9\xfb\xb0\xd3\x91\x63\x6d\xb4\x5a\xcb\x5f\x80\x8a\x8c\x30\x7f\x94\x19\x59\x79\x19\x5d\x88\xfa\xa5\xf4\x95\xfe\x69\x2f\xd1\x96\x9e\x80\x61\xb1\x28\xa9\x0b\x1d\x43\xd3\xb2\x2e\x98\xd4\x4f\xc7\xec\x14\x27\x41\xea\xc1\xa1\x87\x26\xd8\x03\xf5\xff\xb6\x23\x63\x0f\x2b\x93\xd3\xef\x59\x4d\xa7\xd8\x13\x68\x55\x57\x60\x3d\x39\x2a\xad\xf0\x72\x16\x04\x7d\x73\x5c\xe8\x59\x6b\x2a\x72\x79\x55\x9b\x3e\xb9\x8b\x79\x72\x2d\x1f\x6d\x0f\x9b\xb3\x56\xa5\xb6\x77\xc9\x34\xcf\x4f\x76\x9b\xf2\x8e\x86\xcf\x81\xa7\xbc\xe1\xd7\xe2\xbd\xd6\xcf\x5e\xe3\x3d\x9e\xa2\xd0\xdb\xf5\xe3\xf3\x76\x91\x95\xdb\xbe\x72\x2d\x57\x91\x5c\x2c\x27\xea\x9d\x37\xc6\x75\x4e\x54\xd4\xcd\x57\xf7\x57\xea\x85\x3e\x2e\xed\x8d\xfb\x86\x03\x8b\xc5\xab\xbb\xec\x25\x4c\x9d\x7d\x64\x54\xba\x97\x83\xcc\xa8\xf4\xd0\xea\x16\x5c\xd3\x0a\x7f\xda\x8c\xaa\x2a\xd3\xd7\x21\xa3\x3a\xda\x7e\x59\xb2\x32\xaa\xaa\x6c\xff\xe4\x19\xd5\xd1\xf0\xcb\x92\x95\x51\xa5\xf5\x73\x68\x19\xd5\x56\xdb\x3f\xd7\x8c\xaa\xcc\xc3\xe6\xed\xf9\x57\xf7\x39\xe4\x5f\xbd\x7a\xe5\x5f\xbd\x47\xc9\xbf\x7a\x87\x9b\x7f\xf5\x6a\x19\x8a\x7b\xb5\xc9\xbf\x2a\x30\x7d\x6d\xf2\xaf\xa3\xed\x55\xd9\x9a\x7f\x55\x60\xfb\x7a\xe4\x5f\x47\xc3\xab\xb2\x35\xff\xea\x1d\x70\xfe\x95\x65\xfb\x63\xfe\x15\x2f\x53\xa8\x7d\xfe\xd5\xaf\x57\xfe\xd5\x7f\x94\xfc\xab\x7f\xb8\xf9\x57\xbf\x96\xa1\xb8\x5f\x9b\xfc\xab\x02\xd3\xd7\x26\xff\x3a\xda\x5e\x95\xad\xf9\x57\x05\xb6\xaf\x47\xfe\x75\x34\xbc\x2a\x5b\xf3\xaf\xfe\x01\xe7\x5f\x59\xb6\x3f\xe6\x5f\xf1\x6a\xcb\xda\xe7\x5f\x83\x7a\xe5\x5f\x83\x47\xc9\xbf\x06\x87\x9b\x7f\x0d\x6a\x19\x8a\x07\xb5\xc9\xbf\x2a\x30\x7d\x6d\xf2\xaf\xa3\xed\x55\xd9\x9a\x7f\x55\x60\xfb\x7a\xe4\x5f\x47\xc3\xab\xb2\x35\xff\x1a\x1c\x70\xfe\x95\x65\xfb\x03\xcf\xbf\xf2\xed\x19\x29\x24\xef\xfe\x5d\x23\x23\x4b\x6d\x64\xdb\x5d\x29\xe7\x66\x81\x7a\x6f\x9c\x83\x22\xfb\x5b\x60\x7d\xff\xdb\xd9\xf2\xd5\x45\x45\xf7\xba\x6d\xc8\xfa\x55\xb2\xbe\xf0\x6e\xbc\xa7\xd9\x51\x04\xc5\xb5\xe6\xac\x6e\x3b\x04\x44\x1c\x10\xae\xbf\x1c\x71\xe1\xad\x43\xea\x9d\x1d\x52\x84\xf9\x3d\x63\xb4\xd0\x3e\x97\x08\x52\x42\x4a\xfc\x45\x24\x49\x6d\xfc\x36\x90\xa5\xe4\x44\x7e\x61\x94\xa5\x36\xd3\x14\xe9\x23\x69\x54\x26\xce\x6e\xec\xc0\x59\x8e\x5b\x91\xb1\xc2\x6d\x38\x71\xd1\x96\x24\x38\xd5\x75\x34\x7d\x15\x55\x9b\xd6\x82\x60\x94\xcc\x32\xcd\x5a\x58\x18\xac\xec\x32\xdc\x07\xbc\xb4\xd4\xee\xb7\xe4\xdc\x0b\x19\xf2\xc0\xf8\xdb\x57\x0e\x70\x6c\x53\xe2\x18\xe9\x3f\xb8\xb1\xd1\xeb\xcb\x95\x1d\x8b\xeb\x57\x4b\x22\x2a\x49\xd3\x1d\x03\x4a\x8f\xe7\x49\x86\x93\x7f\xc6\x4e\x97\x91\xa5\xe9\x53\x8a\xc5\x25\xfa\x2c\x1e\x4d\x0a\x65\x22\xb9\xf7\x6c\x96\x0d\xe3\x17\xd8\x0f\x30\x43\x22\x64\xea\xed\x70\xe5\xe3\xb8\x58\x0a\xba\x72\xc9\xfe\xa3\x79\x96\xfc\xc7\x8f\xe9\xfb\x48\x85\xd7\x87\xa2\xc2\xf4\x8e\xa7\x48\xbd\x95\x9b\xa4\xc4\x72\xff\xfb\xef\x1b\xa3\xd4\xd0\xeb\x4b\x40\x1a\x8a\x3d\x31\x90\x86\xa2\x5a\x0a\x26\x1d\x1c\x00\x07\xe5\x58\x8e\x24\x8c\x49\xe8\xc8\x89\x6b\x4f\x34\x54\xb2\xaa\x25\x62\xaa\x8b\x03\xa0\xa2\x1e\xcd\x9f\x9b\x8c\xe7\xf4\x06\xb3\x07\x10\x30\x90\xed\xf7\x4f\xba\x94\xd8\x67\x4a\x34\x3d\x82\xfb\xc8\x35\xc8\x26\xd7\x8f\x36\xf2\xac\xf9\x21\xf0\x6b\xe5\x7d\x3f\x65\x08\x86\x95\x80\xfd\x33\x2c\x2d\xf7\x99\x52\x2c\x1a\x42\xcc\xb1\xad\x4c\x3a\x04\x1e\xbd\x41\x01\xb2\x5d\xf1\x10\x26\xd9\x91\x88\xfd\x73\x69\x55\xf2\x33\x65\x53\x32\x88\x2d\x7c\xda\xf6\x86\x23\xa3\x75\x08\xf4\xa2\x84\x87\x7e\xa0\x5f\xa3\x5d\x9e\x61\x4b\x29\x15\x90\x6c\x5d\xf8\x73\xe5\x59\x6a\x1c\x05\xa9\x66\xcd\x6b\x44\xb6\x7c\xaf\x9e\xab\x3a\x7d\xbb\xff\x65\x76\xa5\xb2\xba\x90\x7b\xf8\x4a\x89\xbd\xea\x56\x90\xdd\x6d\x8a\x7f\xa6\x64\x5e\x1d\xc9\x7d\xd9\x5e\x27\x7b\x8e\x8e\x5f\x10\x57\x23\x66\xef\xfd\x4d\x86\x4f\xea\x07\xbd\x7d\xf8\x41\xaf\x5a\x3f\xe8\x1d\x8c\x1f\xf4\x8e\x7e\xf0\xe0\x1d\x27\xd5\xf8\x41\x7f\x1f\x7e\xd0\xaf\xd6\x0f\xfa\x07\xe3\x07\xfd\xa3\x1f\x3c\x78\xe5\x6f\x35\x7e\x30\xd8\x87\x1f\x0c\xaa\xf5\x83\xc1\xc1\xf8\xc1\xe0\xcf\xe9\x07\xe5\xc9\xfb\x41\x3d\xa1\x76\xc9\xec\xa1\xcb\x5d\x68\x2c\xe8\xaa\x9a\x45\x2f\x59\xf2\x1f\x9f\xb0\x19\x48\xf4\x62\x94\xe8\xad\xb7\x75\xa1\x4f\xd1\x45\x63\x60\x53\x8f\x07\x88\x8c\x8d\x9e\xb1\x62\x02\x1b\x13\x11\x99\x20\x5a\x1e\xe5\xa2\x19\xa1\x5c\xb8\x76\x35\xab\xc1\x20\xb5\x82\x2a\xe9\xea\x6a\x82\x84\xc0\xac\xe0\xe3\xb4\xa2\x5e\xb1\x55\x09\x9a\x87\xb0\xae\x90\xad\x00\x4b\xf0\xb2\x58\x64\x28\xc0\xde\xf2\xa1\xe1\x3b\x3d\x24\x70\x39\x78\xcb\x8f\x3b\x94\x59\x09\xb7\xa9\xa6\x38\x4a\xef\x2b\x3c\x68\xc4\x97\x98\x97\x84\x59\x32\x1e\xac\xc0\xd9\xfc\x28\xc6\x4f\xb4\x2c\x9c\x1a\xc4\x84\xf4\x0f\x2b\xdf\x9f\x7f\xfc\x70\x06\x0c\x23\x66\xa9\x39\x18\xb0\x24\xf8\x7e\x28\x81\x71\xc0\xa8\x7f\x64\xc4\x3d\x70\xea\xc5\x88\x8f\x98\x63\x01\xb6\x4c\xb1\x31\x2b\xab\xe2\x35\x22\xc4\xd2\x8e\x54\xb8\x07\x4e\xbd\xa8\x70\x31\xc7\xcc\xa7\x3e\x16\x98\xc1\x3b\x02\x13\x46\x3f\x3f\x68\xd9\xf4\x92\x10\x62\x29\x39\xbd\x9c\xec\x48\x8b\x6c\x38\xf5\xa5\xc5\x87\x50\x54\xc5\x8b\xf4\x12\xaf\x23\x31\xb2\xe1\xd4\x8b\x18\x3f\xe1\x19\x52\xdf\x6c\x4e\x2d\x8f\xe2\x0f\x5e\x9c\x96\xa2\x07\x89\x3a\x58\x5b\x73\x75\x64\x47\x36\x9c\xa7\x7d\x7a\xb1\xdf\x0d\x5f\xbb\xfb\xdc\xf1\x1d\xb5\x2d\x97\x32\x4e\xaf\x9d\x4a\xfd\x19\x1d\x46\xff\x2c\xe1\xc8\x43\x07\x4f\x5d\x82\xc1\x98\x63\xe4\x24\xdf\xdf\xc9\xf7\x4d\xf2\xbb\x3b\xe0\x02\x09\xd7\x7e\x7b\x71\xf6\x1e\x9a\xfa\xf8\x97\x8f\xef\xc1\xb0\x92\x8f\xeb\x58\x88\x73\x2c\xb8\x75\x8d\x89\x43\x19\xb7\x26\x94\x0a\x2e\x18\x0a\xd4\xc7\x88\xd5\xe7\x95\x02\xd7\xfe\x8c\x99\x65\x73\x6e\xad\x9d\x33\x7d\x97\x98\x36\xe7\x06\x4c\x91\xc7\x71\x2b\x05\x2f\x1a\x41\xf6\x58\x3e\xf1\x47\x18\x89\x8c\xf3\x44\x58\x9f\xe2\x23\x05\xf6\xd3\x06\xd6\xfd\x2b\xea\x53\xb6\x9e\x56\xbb\x5e\xd2\x20\xd7\x83\xcc\x15\x22\xbd\x6c\xc6\xcf\x35\x5b\x26\xc3\xc8\xb9\x6d\x4e\x43\xa2\xb5\xd7\x6c\xc1\xdd\x06\x17\x5f\x36\x1b\x7f\x49\x7f\x62\xba\xd1\x32\xd7\x00\x36\x37\x1b\xc9\x22\x5c\x1f\x9f\xab\x0a\xef\x88\xad\xb7\x03\x0e\xe1\x9b\xce\xab\xcc\xca\xbe\x4b\x4e\x91\xc0\x43\xd0\xea\x6e\xb6\x4c\x1e\x4e\x04\x43\xb6\x68\x7e\xfd\x0a\x1a\xb7\x18\xb1\x46\x6b\x4b\x53\xf4\x65\xad\x29\x26\xce\x87\x69\xb3\xe1\xa0\xdb\x6d\x6d\xa4\xa6\x6e\xbf\x53\xdf\x8b\x7c\x23\xe3\x1b\xe6\x43\x68\xa4\xbe\x30\xd9\xc8\x6e\xe6\x51\x1b\x79\x78\x98\xa1\xa5\xb8\x4c\xd5\xf3\xe1\x21\x34\x2e\x2f\x2f\x2f\xcd\xb3\x33\xf3\xf4\x14\xde\xbe\x1d\xfa\x7e\x23\xb3\xc9\x22\xbb\x1f\xa5\x5c\xbe\xab\x9f\xc6\x32\xe1\xa1\x0e\x52\x4b\xc6\x8c\x35\xb3\x18\x60\x4c\xe8\x6c\x86\x98\x50\x1b\xd2\x86\xf0\xdb\x52\xb7\x02\x31\xb1\x54\x51\xa2\xb9\xdf\xb3\xe1\xac\x76\x78\x89\xb9\xc0\xac\x44\xa7\xb1\x41\xbb\xaf\x40\x76\xcc\x1b\xdb\x81\x64\x56\x4e\x9b\x35\x17\xd2\x8b\xb9\xcb\xe1\x06\xe3\xcf\x65\xd5\xe3\x72\xfa\x2f\x8c\x3f\xaf\x20\x8b\x50\x24\x97\x72\x21\x79\x8f\xb8\x28\x85\x24\xad\x86\xa4\xcb\xdd\x00\xb7\xb4\x29\x85\x5b\x69\x50\x7d\x80\xae\xac\x0a\x55\xe3\x2c\x05\x46\x17\xf2\xab\xaf\x14\x8c\xb4\x2e\xa2\x1e\x77\x61\xcb\xac\x5f\x02\xb0\xd2\x9b\x0c\x5a\x65\xd5\x16\x05\xbc\x0d\xad\xe9\xf3\xf9\x95\x56\x06\x43\x5a\x07\xba\xbf\x1d\xb8\xb2\x6a\xaf\x62\xcd\x1b\xf8\x16\xad\xd7\x2f\x76\xce\x3e\x32\xb6\x36\x54\x8f\xbe\x2b\x52\xb3\x16\xce\x9a\xb6\x64\xb9\x46\x0c\x9c\x00\xc6\x5b\xe6\x30\x24\x63\xc9\x8a\x56\x1a\x59\x10\xb2\x60\xe8\x9f\x5d\x7f\xd3\xdf\x7e\x9f\x32\xea\x1b\xbf\x37\x5a\xe6\x35\xf2\x9a\x91\x62\x9c\x40\xeb\x4c\xce\x4c\x2d\x53\xcf\x0a\x4d\x35\x29\xb4\xcf\xce\xda\xa7\xa7\xbf\x5d\xfc\xae\xa6\x85\x21\xe7\xff\x6e\xb4\x5a\xd9\x5f\x95\xde\xd9\xab\xa0\x19\x7d\x62\xe2\x94\xee\x71\xb1\x76\x2e\xfd\xf7\xea\x5e\xe6\x8c\x0c\xed\xff\x01\x00\x00\xff\xff\xa6\x3b\x01\xb6\xae\x8b\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x53\x4d\x6c\x1b\x45\x14\xfe\xe2\xad\x6d\x30\x29\x3f\xe5\xa7\x50\x68\x35\x50\xa5\x50\x2a\xd7\x76\x52\x55\xc8\x89\x13\x20\x4d\x4b\x45\x2d\x42\x6b\x7a\x66\xb1\xa7\xf6\x0a\x7b\xd7\xda\x9f\x06\x4b\x39\xc4\x89\x4a\x28\x41\x2a\xa2\x20\x38\x54\xa4\xed\x01\x04\x12\x92\x93\xd6\x60\x92\xd8\x91\x10\xe2\x86\xf4\xe6\xc0\x91\xde\xe0\x80\xc4\x0d\xc1\x05\x09\x34\x3b\x9b\xb8\x76\x05\x73\xd8\x79\xef\x7b\x6f\xbe\xf7\xbd\x37\xb3\xb7\x76\xed\xf8\x08\x00\xf6\x01\x78\x02\xc0\xe0\x00\x70\x10\xc0\xdf\x03\xf0\xd7\xd5\x10\xf0\x10\x80\x6b\x21\xe0\x11\x00\x5f\x87\x80\x28\x80\x1f\x43\xc0\x5d\x00\x7e\x0a\x01\xf7\x00\xb8\x15\xe4\xfd\x16\x02\xc2\x00\xfe\x09\x01\x3b\x00\x84\x35\x60\x27\x80\x98\x06\xc4\x00\x3c\xaa\x01\x11\x00\xfb\x35\x15\x7f\x36\xf0\x13\xc1\x7e\x54\x53\x7c\x19\x4d\xd5\x39\xa9\x01\xf7\x02\x78\x35\x88\xbf\xae\x29\xfe\x62\x80\x57\x02\x9e\x59\x0d\x78\x00\x40\x3d\xc0\x2f\x6b\xc0\x7d\x00\xae\x07\xf9\xab\x81\x8e\x96\xa6\xfa\xf8\x21\x88\xff\x1c\xf0\xfe\xae\x01\x13\x03\xc0\x1f\x1a\x30\x02\xe0\x48\x18\x38\x0a\xa0\x14\x56\xba\x97\x83\xfd\xcb\x30\x30\x08\xa0\x19\x06\xf6\x03\xf8\x2e\x0c\xdc\x0f\xe0\xd7\xb0\x9a\xc7\x5f\x61\xc5\x2f\x49\x77\x03\xd8\x1d\x51\x7a\x12\x11\x15\x1f\x8f\xa8\xfe\x4e\x04\xf8\x99\x88\x9a\xdb\x1b\x11\xc5\x33\x13\x01\xf6\x02\x78\x27\xa2\xea\x7c\x16\xe0\x37\x22\x00\x03\xf0\x7d\xe0\xff\x12\x51\xf5\xff\x0c\xf6\x3d\x51\xe0\x69\xc9\x1f\x55\xf1\x73\x51\xc5\x3f\x13\x05\x9e\x02\xf0\x76\x54\xcd\xe7\x8b\x28\xb0\x0b\xc0\x7a\x14\x90\x57\x2c\x35\xca\xf9\x0d\x2a\xc9\xfe\x0a\x41\x71\x48\xbd\x7b\xa0\x78\x1e\x03\xf0\xb8\xbc\x3f\xa8\x79\xee\x43\x77\x45\x83\x5d\xf6\x2d\x7b\x91\x73\x7a\x10\xaa\x4f\xb9\x64\xbd\xbd\x81\xbd\x33\x78\x67\x77\xdf\x76\xfe\x61\xf4\x2e\xa9\x53\x83\xba\x1f\x0c\x15\x98\xc3\xf3\x96\x59\xe8\x5a\x0e\x0e\x75\xd1\x43\xb7\xc1\x2f\xd8\xf9\x92\x71\x9e\x63\x52\xaf\xea\x79\xc3\xad\x61\xd2\x32\x1d\xaf\x52\x75\x0d\xcb\xc4\xa4\x67\xdb\xdc\x74\x59\x85\xeb\x8e\x67\xf3\x0a\x37\x5d\x07\xc7\x74\xa3\x5c\xc3\x31\xdd\xe5\xfe\x87\xe9\x66\x81\xb9\x46\x25\xf0\x6c\xdd\x2c\x72\x4c\x99\xdc\x2e\xd6\x70\x22\xaf\x97\xfd\x4f\xa2\x84\x97\x2c\xcf\x2e\xd7\x90\xed\x52\x21\x6b\x99\x6e\xa9\x5c\xc3\x2b\x55\x6e\xeb\xae\x61\x16\x15\xcf\x34\xb7\x0d\xab\x80\x69\x6b\x86\xdb\x98\xf6\xca\x0e\x67\x86\x59\xf5\x5c\x36\x54\xc0\x99\x92\x35\x83\x1c\xaf\xf8\x27\x3c\x9b\xb3\x02\x2f\xbb\x7a\x0f\x62\x98\x3d\xae\xe5\xb9\x38\xab\x97\x3d\x8e\xca\x98\xe3\x55\xc7\x47\xc6\x12\x72\xeb\xf5\x58\x82\x55\x65\x21\xa7\x17\x4e\x94\x10\xc0\xa7\x79\xd5\xb2\xdd\x78\xd6\x29\x1a\x85\xf8\x8b\x5e\xd1\x89\xe7\xac\x34\x2b\xf0\xf3\xcf\xbf\x69\x94\xf4\x8a\x75\xd8\xf6\x62\xa7\x74\xc7\x8d\xe7\x6c\xdd\x74\xca\xba\x6b\xd9\x69\xf6\xb2\x1f\x62\x59\xcf\xd6\x2b\x56\xc1\x62\x63\x3d\xf9\xe3\xb1\x53\xba\x59\xf4\xf4\x22\x8f\xe7\xb8\x5e\x49\xb3\x6d\x3f\xcd\x4e\x7b\x8e\x63\xe8\x66\x2c\x7b\x32\x3b\x15\x3f\xcb\x6d\xc7\xb0\xcc\x34\x4b\x1d\x4e\xc6\x26\x2d\xd3\xe5\xa6\x1b\xcf\xd5\xaa\x3c\xcd\x5c\xfe\x96\x9b\xa8\x96\x75\xc3\x1c\x65\xf9\x92\x6e\x3b\xdc\xcd\xbc\x96\x3b\x1e\x7f\xae\x9b\x27\xf5\x9c\xe3\x76\x7c\xca\xcc\x5b\x05\xc3\x2c\xa6\x59\x6c\xba\xec\xd9\x7a\x39\x7e\xdc\xb2\x2b\x4e\x9a\x99\x55\xdf\x75\x32\x23\xa3\x4c\x99\x19\x73\x28\x95\xcc\x64\x52\xec\xc0\x01\x26\xcd\xe4\x93\x99\x54\x8a\x4d\xb0\x24\x4b\xfb\xfe\x78\x66\x78\x2b\x34\x96\x39\x22\xcd\x67\xfc\xb4\xb1\x54\x92\xcd\xce\xaa\x23\xe3\x99\xe1\xe4\x41\x36\xc1\x52\x2c\xcd\x86\x47\xe5\x43\x14\x75\x6a\xd2\x9a\x58\xa0\x36\xdd\xa4\x46\x3f\x22\x96\xfa\x11\xff\x99\xf6\x1d\xea\x87\xc4\xd2\x1d\x10\xe8\x7d\x31\x27\x2e\x50\x8b\x56\x41\x9f\xd2\x8a\x78\x97\x9a\xb4\x01\xba\x4a\x0d\x51\x17\x17\xa8\x23\x53\xae\xab\x7c\x71\x91\x5a\xd4\x64\xb4\x49\x1d\x5a\xa3\x06\x7d\x43\x0d\x31\x4f\x4d\x5a\xa7\x16\xe8\x9a\x58\x10\xf3\xd4\x11\x8b\xd4\xa6\x86\xb8\x04\xfa\xd0\x8f\x36\xb6\x0d\x46\x2d\x46\xab\x62\x4e\xf2\xab\x78\x8b\x1a\xb4\xe9\xf3\x74\xa8\xcd\xa4\x62\x31\x0f\xfa\x8a\xda\xd4\x14\x73\x74\x83\x5a\x32\xed\xb2\x5f\x6a\x7d\xdb\x48\x88\x45\xd0\xe7\x52\x1d\x75\x68\x55\x95\x5a\xee\x17\x24\xde\x03\x5d\xa1\xa6\xa8\x8b\x4b\x5d\x41\x1f\x6c\xd5\x66\x3e\x32\x47\x0d\x5a\xa1\x8e\x98\xa7\x35\xa9\x7f\xd9\x2f\xda\x52\x0d\x5f\xa1\x8e\xb8\x48\x6d\xea\x88\xba\x98\x97\x64\x9f\xd0\x06\x6d\x8a\x05\xc9\x2c\xea\xd4\x16\x4b\xf4\xad\x6c\xc6\x9f\x8f\xfc\xd1\x7a\x25\xc8\x13\x8a\xaf\xe9\x37\x78\x93\xf9\xb2\x36\x68\xd3\x07\x65\xc6\x82\x98\x53\x73\xed\x03\xe5\x9c\xe4\x7c\xe5\x30\x16\xa5\xae\xff\xc8\xe9\xd0\x8a\x02\xb6\xba\xfb\xd8\x37\x16\xa9\x49\x6d\x79\x4b\xa0\x8d\xff\xfd\x7f\xa9\xd5\xd3\x51\x47\x5e\x7f\xcf\x09\x7f\xce\x77\x26\xfd\x1b\x00\x00\xff\xff\x28\x66\x1a\xe2\xce\x07\x00\x00"

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
