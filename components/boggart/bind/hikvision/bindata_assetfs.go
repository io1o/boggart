// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package hikvision

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\xff\x73\xdb\xb8\xb1\xff\x3d\x7f\xc5\x1e\x2f\xaf\x92\xfa\x4c\xc9\x76\xae\xed\x1b\xc7\xf2\x9b\xe4\x9c\xf4\x3c\xe3\xba\x19\xc7\xbe\xf7\x43\xa7\xe3\x81\xc8\x95\x88\x18\x04\x18\x00\x94\xad\x66\xf4\xbf\xbf\x01\x40\x52\xa4\x44\x49\xd4\x17\xbb\x4e\xa6\xbe\x99\x0b\x05\x62\x17\xbb\x1f\x2c\x16\x8b\x2f\xcb\x6f\xdf\x20\xc4\x21\xe5\x08\x5e\x20\xb8\x46\xae\x3d\x98\x4e\x5f\x9d\x86\x74\x0c\x01\x23\x4a\xf5\x3d\x29\x1e\xbc\xb3\x57\x00\x00\xe5\xd2\x40\x30\x3f\x0e\xfd\xa3\x63\x30\x4f\x2a\xce\x9f\x1e\x95\x7f\x74\x9c\xd5\x9f\xa7\x79\xbc\x4b\x08\x47\x56\x7a\xbb\x58\x23\x97\xa2\x5a\xa7\xa8\x27\x05\xc3\xbe\xa7\xc9\xa0\x8e\x53\x51\x33\x65\x39\x43\x4e\xc6\xc0\xc9\xd8\xd7\x64\xa0\x60\x40\xe4\x9d\x79\xf0\x66\x6c\x18\x55\x75\x6d\x15\x9c\x18\xcd\xea\x26\x12\x15\x72\x4d\x34\x15\xdc\xfb\xf6\x0d\xe8\x10\xf0\x2b\x74\x49\x60\x0a\xc0\x33\xa8\xe5\x6d\x9a\xb2\x31\x9a\x4a\xc8\x43\x98\x4e\xcf\x4e\x09\x44\x12\x87\xfd\x25\x74\x5f\xc8\x98\xa8\x40\xd2\x44\x9f\x8c\x05\x0d\xdb\x87\x9d\xb7\x86\x96\x29\x84\xe9\xf4\xdb\x37\xe8\x5e\xe3\xd7\x14\x95\xee\xde\x5e\x5f\x76\x3f\x11\x1d\xb9\x62\xc7\xdc\x3b\x33\x4c\x8f\xfe\x87\x83\xf7\x49\xe2\x98\xe2\x83\x07\x5d\x98\x4e\x4f\x7b\xe4\xec\xb4\xc7\xe8\x1e\x74\xa3\x31\x19\xe1\xb6\x0a\x16\xc4\xdb\x68\xf9\xbf\x8e\x4b\xdf\x32\xa9\x53\xf9\xc2\x71\xdf\xb3\xc2\x6a\xa2\x34\xc6\xdb\x6a\x3c\xa3\xde\x45\x65\xc7\xa5\x4e\xe7\xcf\x19\xff\x66\x4a\x3b\xf9\xba\x38\x46\xae\xd5\x1d\x72\x32\x60\x68\x18\x36\x06\x83\x0b\x4d\x87\x34\x70\xaf\xb7\x84\x64\x9e\xc7\x2e\xc0\x94\x79\xd5\xc1\x73\x55\x69\xab\x02\x52\x51\xbd\xde\x6b\xf4\x52\xb6\xc4\x9f\x94\x3c\x94\x26\x03\x7f\xb9\x8f\xaa\x50\xcc\xf9\xaa\x32\x07\x53\x02\x43\x12\x22\x38\x18\x81\xf2\x15\xdc\x56\x8c\xaa\xa5\x34\x19\x9d\x24\x7c\x84\xf0\x9a\x1e\xc0\xeb\x20\x22\x9c\x23\x83\x93\x3e\x74\xb3\x67\xb5\x8e\xc3\xbc\xf6\x99\xd7\xff\x53\xee\xf4\xff\x54\xeb\xf3\x57\x72\xa3\xf1\x08\x94\x0c\xac\x99\xbc\x5e\xd5\xd1\x89\x73\x67\x7f\xc8\x64\xed\x1f\x1d\x1a\x8a\xec\x57\xf7\xe2\xdc\x74\x3b\x3c\xd0\x50\x47\x7d\xef\xe8\xf0\xf0\xbf\x3c\x88\x90\x8e\x22\x9d\xff\xa2\xa1\xb5\x6d\xe7\x12\x7b\xeb\xa5\x3b\xed\x85\x74\xdc\xa0\xda\x22\x1c\x7f\xc9\xe1\xf8\xcb\xe6\x70\x0c\x85\x8c\x33\x4b\x31\x8f\x1e\x84\x44\x13\x5f\x8b\xd1\xc8\x14\x8d\x09\xa3\x21\xd1\x42\x3a\x75\x6c\xb7\xfb\x0a\xb5\xa6\x7c\xa4\x3c\xc8\x90\x5a\x0b\x65\x66\x2e\x31\xea\x48\x18\x54\xc4\xca\x49\x6f\xfe\xcf\x99\xdf\x0c\x79\x19\xa4\xfa\x23\x65\x1a\xe5\xcd\x24\xc1\x26\x26\x54\x87\x1d\x09\x02\x21\x43\x3b\x48\x67\xaa\x95\x0a\xab\x53\x34\x10\x49\x89\x1f\xa7\x4c\x53\x85\x0c\x03\x53\x6e\x5e\xcb\x14\x37\xd0\x64\x5e\x86\x55\x41\xc4\x4a\x1e\xa4\xc2\xc1\x8f\x90\x84\x94\x8f\x4a\x42\xd7\x2a\xe5\xab\x07\xaa\x83\x68\xae\x8f\x03\xc1\x18\x49\x14\x66\xc5\x09\x91\xc8\x75\xdf\xfb\x79\x01\x11\xe7\x5d\xb3\xf2\x9c\xaa\xe0\x69\x01\xc2\xc7\x84\xf0\x10\xc3\x0c\x19\x57\x68\xdc\x95\x14\x4c\xe5\xf2\xcc\x93\x6e\xae\xbf\xc5\x20\xfa\xa5\x0a\x82\xa6\x9a\x61\xc9\x11\x9f\x93\x49\xef\xca\x0c\x48\xc8\x45\x7c\x6d\x9d\x71\xf4\xcb\x16\x80\xf7\xc8\x16\x44\xa6\xa7\x67\xdd\xb0\x80\x58\x45\xfa\xfc\x2d\x14\x0f\xb4\x6c\x84\x99\x0b\xb7\x78\x32\x32\x40\xc6\x30\x1c\x4c\x96\xf6\xf0\x96\x90\xce\x9b\xa6\x3f\x10\xe1\x64\x4b\x66\x96\xa1\x15\x15\x86\x42\xf6\x3d\x2a\xfd\x20\xd5\xfe\xd0\x0e\x5c\x5f\x4f\x92\x4d\x07\xce\xfc\xdf\x9a\x8e\x86\x53\x95\x10\x5e\x2c\x22\xf0\x6b\x4a\x25\x86\xde\xd9\x1f\x4f\x7b\xe6\xc5\x0e\x4a\xf5\xac\x56\x3b\x30\x70\x1e\xc4\x99\xc6\x22\x2a\xb9\xc8\xc6\x1d\xe7\x63\x07\x1c\xc9\xb1\x07\xb9\x22\x65\x95\x76\x82\xf1\x54\x24\x76\x4e\x1f\x13\x96\x62\xdf\x0b\xc9\x64\x16\x40\x2d\xf5\xb9\xb6\x9a\x01\xd9\xc9\x65\xc4\xc9\x9f\x4a\x11\x59\xb9\x87\xf2\xd1\xe7\x5a\xdb\xab\xc8\xdc\xf4\x7d\x13\xa1\x5d\xc5\xe6\x62\x5f\xb9\xfa\x4f\x26\x38\x49\xb5\x68\x22\xb7\xad\xd7\x5c\xec\x77\xb6\xfa\x93\x49\xad\x82\x08\xc3\x94\x61\x13\xc9\x8b\xba\xcd\xa5\xff\x5c\x90\xec\x47\x83\xd3\x9e\x6b\x6d\x4b\x97\xd8\x2c\x28\xdb\x03\xd9\x36\x24\xff\xc6\x50\x62\xc8\x68\xb2\xdf\x40\xc2\x71\xdc\x22\x8c\xb0\x84\x4f\x14\x44\x7c\xb4\x42\xbd\x90\xc0\xc1\x21\xb4\x3a\x6c\xd8\x22\x66\xd8\x05\xbf\xa7\x8c\x18\x76\x90\x2b\xff\xab\xe9\xc7\xef\x23\x2e\xa8\x74\xf5\xbf\x29\x12\xa0\xca\x6e\xd6\x64\x7e\x9e\x0b\x5d\x72\xf4\xc6\x84\x0c\xa6\x1f\x8a\x0d\x9d\xa6\xa1\x40\xce\xf5\xe9\x26\xa8\xcb\x0f\x1f\x6f\xae\x2f\xfe\xfa\xdb\x4d\xdd\x0c\x95\x0b\xfe\x59\x4f\x18\x42\xa9\x6e\x73\x15\x2e\x71\xa8\xa1\x07\xf2\x89\xa3\x83\xdb\x4f\xe7\x7f\xff\xbf\xab\x06\x3a\x64\x15\x9b\x2b\x70\x9b\x40\x0f\x42\xf1\xc0\x9f\x50\xfa\x5f\x3f\x5c\xdd\x7c\xb8\x6e\x20\x7d\x56\xb1\xb9\xf4\xbf\x22\xd7\x28\xf7\x2a\xfa\x4f\xbe\x3f\x27\xfe\xbb\xdb\x9b\xbf\x37\x10\xde\x56\xdb\x29\x32\xf3\xfd\xff\x84\x36\x75\x24\x3f\xf9\xfe\x4b\x89\x85\x8a\xb9\x35\x5c\x1b\x15\x09\x1d\x19\xd3\xdc\x67\x58\x94\xb1\x9c\x8b\x8b\x86\x84\xa9\xb5\x81\x91\x23\xdd\x6f\x64\xf4\xab\x63\x4e\x07\x0c\xe1\xaf\x52\xa4\x09\x5c\x68\x8c\x15\xfc\x7c\xfc\x02\x02\xa4\x0c\xab\xbd\x47\x48\x3b\x01\xb9\xf7\x10\x29\x39\x3b\x55\x5a\x0a\x3e\xaa\x74\x86\xe9\x06\x38\xb6\x36\x76\xda\xcb\xde\xef\xe0\x59\x92\xed\x89\xdf\x71\x1a\x43\x62\x40\xd5\xa9\x84\x80\xd1\x20\x42\x90\x98\x48\x8c\x90\x87\x28\xa9\x3e\x00\x34\x75\x90\xa6\x2a\x16\x21\x44\x74\x14\x01\xa3\x43\x04\x12\x04\xa9\x22\x71\xaa\x40\xa3\x94\x13\x90\x34\x88\x88\x0c\x95\xe0\x40\x42\x50\x5f\x53\x1a\x76\xe1\x0d\x3c\x08\x36\x84\x58\x08\x0e\x62\x38\xa4\x01\x25\x40\x52\x8d\x07\xc0\x05\x87\x20\x4d\xa8\x01\x41\x83\xba\x27\x1a\x07\x82\xc8\x10\x42\xc1\x84\x84\x81\x4c\x79\x10\x75\xe1\xa3\x10\x21\x68\x99\x06\xf7\xf0\x35\xa5\x5c\x10\xe0\xa8\x02\x9a\x72\x0d\x8c\x0c\x84\x4c\x0b\xd1\xba\xf0\xde\xd2\x54\xda\xd4\x18\x27\x42\x1e\xfc\x78\x4e\x77\x93\xb9\x68\x43\xf6\xab\x8f\xbe\x16\x99\x9b\x90\x77\x2f\x47\x26\xeb\x1b\xce\x4f\xfe\x96\x1e\xa3\xae\xe4\x5f\xbd\x3d\x90\x39\xc9\xf5\x82\x47\xc7\xb3\xa8\xe0\x3d\x51\x34\x00\xca\x8d\xce\x95\x63\xc3\xe8\x78\xc3\xe3\x20\x86\x44\x0e\xe9\xa3\x77\xd6\x00\x98\xac\xca\xea\x3a\xf6\x98\xa3\x74\x70\xc8\x10\xec\xff\x7d\xa5\x25\x4d\x30\xb4\xee\xc6\x95\x87\xda\x97\xa8\x12\xc1\x15\x1d\x23\x70\xf1\x20\x49\xe2\x81\x32\x71\x52\xdf\xb3\xc7\x64\x27\xf6\x5c\xac\x81\x4a\xda\x4c\xbc\x4d\xea\xc9\x86\x87\x5c\x3a\x9a\x3b\x31\xfb\xa5\x7c\x6d\x81\x48\x12\xa3\x0d\x29\x1d\xec\x3a\x6a\xcc\x76\xc6\xe5\x77\x13\x37\x6e\xc2\xe1\xb4\xd7\x44\x7c\xc3\xab\x21\x18\x66\x5a\x69\x2a\x78\x43\xe0\x5c\xe5\xb0\xb4\x7a\xc3\x31\x0d\x10\x38\x89\x67\xba\x36\x10\x6e\x8e\x57\xd7\xd8\x7a\xd7\xf1\xba\x22\x31\x6e\xc4\xa7\x19\x6e\xb0\x17\x35\xd3\x78\x50\x32\x8b\x6d\x15\xbd\x31\x41\xba\x8b\xd1\xec\xa9\xf1\xcb\xd2\xf5\x6f\x22\x34\x31\xd0\x8e\x3a\x5a\x2e\x2f\x4e\xb7\xcf\x28\x29\x61\xfb\xea\x47\xc7\xed\xca\x32\x7b\x71\xaa\x7e\xa4\x32\x7e\x20\x12\x61\x8c\x52\x95\xe6\x90\xad\xb5\xcd\x19\xfe\xee\xf8\x99\x45\xee\xc2\xbb\x6b\x64\x48\x14\x86\xe7\x44\x3f\xd5\x20\xce\x36\x9f\xb0\x5e\xaa\x0b\x3e\x14\xee\x12\xdb\x33\xc3\x6b\xa7\xeb\x7d\x63\x6c\xb5\x79\x32\x18\x37\x89\xbf\xb6\xc6\xe9\x03\x0f\x44\x88\x72\x6f\x56\x98\xf1\xab\x33\xc2\xec\xd5\xf3\xdb\xe0\x39\x56\x64\x7a\x0e\xf3\xcb\x9a\xdc\x1b\xac\x73\x2a\x94\x61\xcd\x5e\x3d\x0f\xac\x1b\xd8\xe4\x5c\x27\xbc\x17\x42\x3f\x67\x0f\x98\xf6\xf6\x06\x7f\x59\xf8\x32\xf6\xa6\xfc\xc5\x01\xbf\x73\x18\x75\x71\xbe\xa7\x58\xf1\xc9\xa2\xa7\x85\x01\x6e\x5a\x3b\x47\x77\x45\xf4\xf9\xc6\xb8\x45\x2b\x9c\x35\xbb\x27\xd8\xca\x8a\xbc\x04\x83\xaa\x45\xfb\x52\xb8\x2b\xb3\xcf\x09\x35\x13\xd5\x6b\xba\x3b\xe2\x5c\xa8\xf0\x02\x41\x76\x97\xb6\x7f\x15\x5c\x93\x40\x3f\x0b\xc6\xae\x45\x08\x5c\x93\xbb\x07\xdf\x15\x05\x5e\x02\xc2\x3b\x2c\xb9\x48\x00\x24\x0c\x25\x2a\xb5\xfb\xc2\x8b\x04\xef\x1c\xab\xe7\xf1\x8d\xbf\x11\x19\x96\x97\x05\xcf\x61\x4a\x79\x9b\x7b\x9b\x7f\xe7\x95\xf8\xbe\xcd\x29\xf3\x66\xee\xf2\xe0\x5e\x3c\x59\x76\xaf\xfa\x65\xad\x70\x3f\xa7\x49\x22\xa4\x86\x01\x62\xb2\xbb\x3b\x71\xcc\xde\x23\x26\x2f\x56\xd1\x31\x0d\x51\x00\x13\x7b\xf0\x12\x19\xcb\xdf\x0d\xc7\x4b\xf1\x24\xbe\xe2\xb4\xd7\x60\x07\xf2\xb4\x67\x37\x8b\xd7\x6d\x3b\xef\xbe\xa7\x7e\x9b\x84\x26\x7e\x1e\x66\x0b\xec\x17\xb3\xa3\x9e\xa5\x43\xa5\xc9\x48\x92\x10\xbb\xb7\xf6\x5f\xca\x47\x9b\x26\xc2\x24\x52\x8c\xdc\xfc\x91\xed\xb0\xc7\x44\x8e\x28\xf7\x07\x42\x6b\x11\x9f\x1c\x36\x4d\xff\xa8\x61\xe9\x0f\x88\x84\xf2\x8f\x62\xb3\xbf\x52\x68\xec\x2a\xcb\x1e\xf2\x8a\x4c\x2e\xf7\x7e\x40\xf2\x03\x6b\x7b\x9b\x82\x8b\x07\x9b\x22\x52\x68\xfd\x09\x65\x80\x5c\xdb\x04\x9a\x59\xb5\x98\xf2\xbe\x77\x58\x29\x21\x8f\x36\x99\x66\xee\x18\xa1\x9e\x55\x93\x93\x85\x42\x6d\x7b\x05\x2c\x63\xca\x28\x47\xdf\xa5\xee\x9c\x1c\x1f\x26\x8f\xf6\x68\xa0\xae\x81\x4d\x2e\x88\x35\x4d\xe6\x69\x76\x80\xe5\x12\xd3\xd6\x73\xb3\xc9\x3c\xe5\x4b\x63\x91\x90\xf4\x5f\x26\x66\xb2\x57\xea\x62\x77\xc8\xed\x33\x1c\x6a\x08\xa5\x48\xfe\x25\x78\xd1\x79\x2e\xfb\xa7\x92\xa2\xe3\x6e\xa2\x15\x43\x28\xcf\xf6\xa9\xa6\x08\x7a\xc0\x45\x96\x26\x84\x67\x4d\x0e\xee\x36\x3a\x91\x23\x3c\x5c\x48\x1e\x6c\x2f\x4d\xed\xeb\xac\x3d\xad\xdb\x28\xdd\xa9\xcc\x7b\x8b\xac\xa7\x6a\xd6\x61\x15\xd9\x66\x9d\xd4\xc4\xeb\x95\x46\xb0\xa5\x1e\x49\x91\x36\xbd\x36\x59\xbe\x6d\x99\xca\x59\x8a\x60\x76\x5a\xe1\x24\x81\xec\xcc\xec\x4d\x9e\x65\xf6\x66\xe3\x2c\x33\x28\x5f\xc5\xbc\xbd\xbe\x74\xee\xb8\xe1\x40\xda\xe0\x56\x65\x4d\x6e\xdc\x9f\x73\xa9\xff\xbc\x95\xd4\xa7\x94\x27\xa9\xb6\xb1\x55\xdf\xd3\xf8\xa8\x6b\x2f\x65\x7a\xf6\x44\x2c\xc3\xd0\x58\x8e\x7d\xc8\xaf\x65\x6e\x74\x06\x9f\x30\x12\x60\x24\x58\x88\xd2\x39\xcd\x15\x06\x66\xc7\x85\x9b\x51\x46\x08\x6d\x86\xdc\x8c\x8b\x00\x95\xba\xbb\xc7\x49\x07\x0e\x61\x3a\xfd\xc3\x3d\x4e\xfa\x86\xcf\xec\x45\x25\x7d\x7b\x23\xe1\xb2\xcb\x71\x86\x5d\xd9\xb8\xbb\xa9\x64\xd6\x99\x37\xc8\x72\x84\xbd\x3a\x47\xd8\xe7\x08\xc8\x96\x64\xbe\x3b\x87\x7f\x9e\xc1\x90\xaf\xdd\xf2\x36\xbf\x9b\x71\x51\xba\xab\x3c\x8f\x9b\x1b\x0b\x4b\xd0\x7c\xa2\xbb\xcc\x73\xd7\x37\x69\x52\xac\xaf\x67\xc9\xcb\x15\x93\xcd\xde\xdf\x39\xf1\xa0\x4c\xd1\xf8\x3a\xe7\xfb\x09\x5c\x7c\x9a\xad\xe4\xb7\xbf\x90\x3a\x27\x7d\x24\x94\x76\x47\xec\x0d\x85\x9f\x11\x6c\x22\xfb\x8c\x6a\x3b\xc9\x37\xbb\x7e\xfa\x42\x07\xfd\x0c\x84\xe7\x18\xee\xbf\x15\xad\x7d\x37\x03\x7d\xa3\x09\x70\x86\xa6\x71\x0c\xc5\xaf\x9d\x66\xc0\x40\xf0\x21\x1d\x81\x17\x12\x15\xd9\xdb\x7c\xdd\xc8\x86\x50\xfb\x9a\xbc\x72\x29\x7f\x80\x19\xac\x94\x33\xf2\xa4\x66\x5c\xf6\x7a\x3f\xa6\x21\x1b\x24\xed\x05\xe3\x64\x07\xe3\x3d\xec\xda\xff\xf6\x63\xa6\x34\xb9\x23\xc5\x46\xef\xf7\x6e\xa8\x89\x90\xcf\x14\x60\x7d\xb2\x2d\xfd\x98\x46\xea\x50\xb4\xdf\x19\xb1\x4f\x7b\x59\x70\x2c\xba\x5b\xc7\x7c\x5f\xee\xd6\x6e\x6d\xbe\x08\x0b\x66\xfc\x4e\x09\x46\xc3\x46\x9b\x7b\xf3\xc4\x9b\x9b\xff\x32\x63\x89\x43\x5f\x0c\x87\x0a\xb5\xff\x66\x13\x4b\x19\xa4\x5a\x0b\x9e\x99\x8a\x4a\x07\x31\x9d\x19\xcb\x40\x73\x18\x68\xee\xab\xd4\x2e\xfa\x4a\x97\x5d\xdf\x25\x09\x9b\xe4\xbb\xa1\x8e\xc5\x73\x77\xc3\xba\x3d\xa2\x86\x1b\x5e\xf5\x9d\x01\x49\xca\x98\xef\x52\xf3\x36\xeb\x50\x3b\xf2\xb6\x71\x68\x33\xcc\x37\xa2\x86\x85\x4e\x74\x3f\x16\x3a\x91\x06\x22\x7b\xe0\x43\xe1\xba\x35\xf6\x40\xf0\x80\xd1\xe0\xbe\xf8\xc4\xd0\x35\x0e\x25\xaa\xa8\xdd\xd9\x74\xe1\x44\x0b\x10\x89\x82\x21\xf1\xd5\x84\x07\x1e\xd8\x4d\x77\xf7\x29\x2d\x6b\x37\x19\x7b\x67\x39\x66\xc0\xac\xf8\xe4\xd7\x42\x13\x1b\x19\x9a\xa5\x28\x7d\xca\xab\xc9\x27\x9a\x42\xf1\xc0\x99\x20\x61\xff\xa8\x19\x7a\x3b\x42\x94\x37\xb7\x08\xd3\x79\xf1\x66\x2b\x9c\x1a\xa6\x42\x35\x9e\xa4\x9a\x1f\x5a\xac\xac\x52\xfe\x64\x56\x83\xee\xd8\xcb\x57\xb1\xd6\xef\x12\xaf\x90\x7c\xc9\xab\x9a\xe2\xb9\xa2\xd2\xcf\xec\x31\xfb\x67\x26\xce\xab\xd2\x27\x2b\x23\x34\x5d\x9d\xc9\xb8\xea\x4b\x7c\xaf\x4a\x6a\x29\x4d\x34\x0d\x7e\xbb\xf9\xdb\x25\xb4\xdd\xf3\xed\xf5\x25\x78\xbd\x62\xc2\xed\x11\xa5\x50\xab\xde\x18\x79\x28\xa4\xea\xe5\x3b\xf5\x5f\x54\x2f\x50\xb3\x9f\xdd\x98\xf2\x6e\x60\x42\x70\x9b\x0c\xd8\x29\xc9\x91\x89\x5a\x2f\xf4\x17\xb5\x42\xe4\xc5\x0f\xbb\x6d\x23\x71\xb1\x8f\xde\xfb\x52\xfa\x61\x05\xfe\x52\x2b\x6f\xb3\xfc\x9b\x1d\xc1\xfb\x32\x87\x5d\x55\x94\x99\x0d\xb8\xdb\x52\x99\x53\x26\x49\xc2\xb2\xd0\xa5\x37\xfb\x6c\xe0\x9c\x07\x79\xdd\x0e\x45\x90\xc6\xc8\x75\xa7\x2b\x91\x84\x93\xf6\x30\xe5\x4e\x8f\x76\x07\xbe\x2d\x18\xe2\x79\x2e\x86\xdb\x72\x51\xdd\xfc\x90\x05\xfa\x35\xb5\xcd\x5f\x42\x24\x61\x0c\xd9\x6d\x62\xdc\x8b\x3a\x81\xa3\xfa\xfc\x33\x53\x2f\xbe\x22\x31\x9e\xc0\xec\xe4\xa6\xbe\x6a\x20\x91\x68\xb4\x19\xcd\x37\x51\x1a\x0f\x38\xa1\x4c\x9d\x38\x44\xea\x29\x48\x10\x60\xa2\x31\xfc\x48\x19\xaa\x13\xa8\x60\x23\x02\x8d\xda\x57\x5a\x22\x89\x0f\xba\x21\x19\x2f\x46\x8b\xd3\xb7\x95\xa2\x69\xe7\x6d\x69\xdc\x39\x5c\xcf\xe6\x2c\xf8\xd5\x16\x1d\xf2\x40\x79\x28\x1e\xba\xd5\x79\x11\xfa\x90\x77\xc9\x42\x8f\xbc\x6e\xb7\x7e\xce\x6a\xb7\x3a\x5d\xa2\xb5\x6c\xb7\x94\x0c\x5a\x07\xd0\x6a\x38\xff\xdc\xe9\x7e\xeb\xbf\xdb\x1c\x1f\xe0\x9c\x68\x6c\x77\x3a\xdd\x11\xea\x1b\x1a\x9b\xc7\x99\x92\xd3\xb7\x33\x23\xab\xb6\xb9\xce\x66\x32\x95\x14\xea\x0b\xae\x51\x8e\x09\x6b\xd7\x6a\x79\x60\xaf\x92\x66\x65\x77\xd2\x15\xde\xd1\x8c\x06\xa6\x53\xf8\x23\x1c\x1d\x1e\x1e\x96\x85\xea\x54\xa4\x6a\x6e\xc7\x46\x83\xea\xf7\x00\xc1\x86\x50\xff\xa0\xe1\x3f\x0f\x60\xfe\x95\xdb\x29\x5c\x2c\x37\x6b\x1c\x22\x91\x18\xaa\x56\xc7\x7e\x17\x72\x84\xed\xa5\x5d\x05\x76\x61\x21\x97\xce\x09\x66\x00\xbd\x6e\xeb\x88\xaa\xce\xf2\xfc\xcc\x90\x68\x62\x06\xda\x74\x79\x95\x7b\xe8\x03\x76\x13\x29\x92\x76\x8b\x86\xad\xce\x42\xc5\xb7\x8b\xc7\xfc\x74\x08\xed\x9c\xc6\x98\x6a\xab\x03\xfd\x7e\x1f\x5a\x41\x84\xc1\xfd\x40\x3c\xb6\xea\x94\xc9\xe5\xf9\xc7\xfd\x3f\x4b\x4d\x5a\x12\x0c\x5b\x1d\xf8\xa9\x0f\xad\xd6\xdb\xc5\xa1\xe4\x5c\xe6\x7a\x7e\xc6\x54\x16\xa5\x9f\x2e\x4a\xff\xba\x4b\xbe\x90\xc7\x76\x3d\x47\xa3\xce\x49\x4d\x8f\xb7\x3a\x99\xc0\xee\x70\xb3\xb5\x04\xf4\x54\xb2\x95\xd4\x6e\x34\x2d\xa3\x36\xea\x9c\xd8\xff\xd7\xbf\xcf\x16\x39\x27\x30\xb3\x55\xb9\x0c\x6a\xc8\xfa\x49\x76\x25\xaa\x94\x69\xd7\x43\x43\x42\x99\x45\x7b\x39\x91\xf9\x33\xe3\xfb\x93\xfd\xbc\xeb\x64\x09\x4e\x15\xcc\x4c\x68\x78\x02\xad\x0f\x52\x0a\xd9\x5a\x9f\x2d\x6c\x46\xc2\x09\xc8\x6e\x8c\x4a\x91\xd1\x12\x0f\x5c\x21\xb0\x9d\xd2\xc2\x86\xfc\x95\x9e\x30\xca\x47\x27\xd0\x1a\x08\xa1\x95\x96\x24\x79\xd3\x5a\x49\x55\x76\xd0\x0b\xef\x8a\x39\xbb\x5d\x88\x0c\x3f\x19\x30\x53\xee\x02\x8d\xa7\xc2\xf3\xb3\xeb\xee\x27\x44\x54\x35\x6e\x61\xcf\x98\xd6\xbe\x59\x2c\x9d\xe7\x31\x9d\x73\xe7\x50\x99\x50\x8b\xc9\xf4\xff\x03\x00\x00\xff\xff\xeb\x39\xf1\xc9\x6b\x5d\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x52\x4d\x6f\x1b\x55\x14\x3d\xe3\x69\x29\x18\x84\xa0\x0b\x24\x24\x16\xb7\x12\x8d\xa0\xd2\x34\x76\x60\x81\xc6\x99\x84\x8f\xb4\xa8\xa2\x46\x55\x6a\xd8\x3f\xd9\xcf\xf6\x80\x3d\xcf\x9a\x8f\x98\x48\x41\xb2\x13\x20\x2c\xa2\x88\x05\x91\xd8\x40\x84\xd8\xb2\x70\x22\x4c\x86\x24\x1e\xfe\x01\xba\x6f\xc1\x0e\x65\x0b\x3f\x03\x3d\xcf\x24\xd0\x2b\x59\xf7\x9c\xf7\xce\xb9\xf7\x3c\x6b\x2e\x6e\x5e\x3b\x00\x80\xe7\x00\xbc\x02\x60\x0b\xc0\x8b\x00\xfe\x42\x5e\x64\x01\xd7\x00\xdc\xb2\x00\x1b\xc0\x82\x05\xbc\x00\xe0\x8e\x05\x3c\x0d\xa0\x66\x01\xaf\x02\x58\xb3\x80\xeb\x00\xfc\xa2\x2b\x0b\xb8\x01\x20\x29\xfa\xe7\x85\xfe\x4b\x0b\x78\x0a\xc0\xbe\x05\xac\x5a\xc0\x41\x71\xfe\x67\x29\xef\x17\x25\xe0\x0e\x80\xbf\x4b\xc0\xf3\x00\x6e\xda\xc0\x27\x66\xbf\x0d\xbc\x04\x60\xcf\xce\x75\x3f\xd9\x79\x8e\x9f\x8b\x3e\xb3\x81\x97\x01\xfc\x61\xe7\xbe\x7f\x6c\xc0\x42\x9e\xdd\x94\xf1\x5c\x2f\xf0\xb3\xc8\xdf\x72\xa3\xe0\x65\xe4\x99\x9e\x01\x50\x2a\xfe\x8b\xab\x7a\x27\x89\x15\xd6\xc4\xa6\xf9\x2d\x7e\xe8\x77\xba\x31\x45\x43\x3f\x6e\x76\xb1\xa6\x86\x41\x4f\x89\x16\xde\x97\x31\xb5\xe4\x86\xdf\x94\xe4\x07\x6d\x45\x6d\xe1\xf7\x64\x8b\x86\x7e\xdc\x25\x19\x86\x2a\xa4\xdb\x11\x1e\xf4\x45\x47\x62\x3e\x00\x8f\x42\xb9\xe1\xcb\x21\xd6\x65\x3b\x94\x51\x17\x8f\x9b\x5d\xd9\x4a\x7a\x12\x8f\x37\xa3\x58\xf6\xb1\x2e\x07\x2a\x8c\x9d\x7a\xd4\xf1\x5b\xce\xbb\x49\x27\x72\x1a\xca\x35\x2b\xde\xfe\xd4\xef\x8a\xbe\xba\x1b\x26\xe5\x87\x22\x8a\x9d\x46\x28\x82\xa8\x27\x62\x15\xba\xf4\xc1\xfc\x8a\xea\x49\x28\xfa\xaa\xa5\x68\xf9\x09\xfd\x4a\xf9\xa1\x08\x3a\x89\xe8\x48\xa7\x21\x45\xdf\xa5\x2b\xee\xd2\x7a\x12\x45\xbe\x08\xca\xf5\x07\xf5\x7b\xce\xc7\x32\x8c\x7c\x15\xb8\x54\xbd\x5b\x29\xbf\xa7\x82\x58\x06\xb1\xd3\xd8\x1c\x48\x97\x62\xf9\x59\xbc\x38\xe8\x09\x3f\xa8\x51\xb3\x2b\xc2\x48\xc6\xde\x47\x8d\xfb\xce\x5b\xff\xe9\x4c\x9e\xb6\x0c\x9d\x7b\x41\x53\xb5\xfc\xa0\xe3\x52\xf9\x51\x2f\x09\x45\xcf\xb9\xaf\xc2\x7e\xe4\x52\x30\x98\xd3\xc8\x7b\xa3\x46\x39\xf4\x82\xdb\xd5\x8a\xe7\x55\x69\x61\x81\x0c\xac\xdc\xf2\xaa\x55\x5a\xa5\x0a\xb9\x73\xbe\xe2\x2d\x5d\x5e\x2d\x7b\x6f\x1a\xf8\xda\x5c\xb6\x5c\xad\xd0\xd6\x56\x6e\x59\xf1\x96\x2a\xaf\xd3\x2a\x55\xc9\xa5\xa5\x1a\xf8\x1b\x3e\xd6\xdb\x9c\x81\xbf\xe5\x29\xcf\xf4\x1e\xf8\x90\xa7\x7a\xc4\x53\x3e\xe5\x33\xbd\xaf\x77\xcd\x31\xa7\x3c\x25\xfe\x25\x57\x2c\xf2\x8c\x33\xbd\x6b\xa4\x3f\xf2\x29\x4f\xf4\x2e\x4f\xf4\x76\xee\xcc\xf8\x4c\xef\xfc\xdf\x93\xf2\x4c\x7f\xc1\x99\x1e\xf1\x39\x4f\xf4\x57\x9c\x72\x4a\x9c\xf1\x11\xe9\x1d\x3d\xd6\xdb\x7a\xc4\x19\xff\x66\x10\x1f\x1b\xf9\x09\x4f\x0c\xd0\x23\xfd\x35\xa7\x7c\xc6\x99\x1e\xeb\x3d\xe2\x6c\x4e\x8f\xf8\xd4\xa8\xcd\x27\xc2\xdf\xf1\x89\x19\xa3\x47\x3c\xe1\x5f\x2f\xd7\x81\xbf\xbf\x4c\x76\x68\x26\xeb\x31\x9f\x73\x66\xb6\x80\x7f\xe0\x23\x93\x9b\x8f\x39\xbd\x0a\x4b\xc6\xae\xc7\xfc\x3b\xa7\x7a\xcc\x13\x33\x44\xef\x9b\x57\xa5\xf3\x44\x53\x93\x19\xff\x06\x00\x00\xff\xff\xe4\x4d\x6a\x9d\xf7\x03\x00\x00"

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
