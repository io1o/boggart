// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package miio

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x6d\x6f\xdb\xb6\xb7\x7f\xff\xff\x14\x9c\xd6\xc1\xf6\x5d\x25\x3b\xe9\xc3\x06\xcf\xce\x45\x9b\xb4\x37\x01\x92\x2e\xb7\xc9\x06\x14\xc3\x50\xd0\x22\x6d\xb1\x91\x48\x95\xa4\x9c\x64\x81\xbf\xfb\x1f\x24\x25\x59\x72\x24\xc7\x92\x9d\x58\x5d\xad\x17\xad\x4c\x1d\x1e\x9e\xdf\x79\x20\x0f\x49\x51\xb9\xbb\x03\x08\x8f\x09\xc5\xc0\x72\x19\x95\x98\x4a\x0b\xcc\x66\xff\x19\x20\x32\x05\xae\x0f\x85\x18\x5a\x9c\x5d\x5b\x07\xff\x01\x00\x80\x6c\xa9\xcb\x7c\x3b\x40\xf6\xde\x3e\x50\x77\x22\x48\xee\x6e\x84\xbd\xb7\x1f\xd3\x2f\xd6\xb9\xf9\x1c\x42\x8a\xfd\xcc\xd3\xfb\x14\x89\x14\x79\x9a\x94\x8e\x33\x1f\x0f\x2d\x09\x47\x45\x9c\x52\xca\xc8\x4f\x18\x52\x38\x05\x14\x4e\x6d\x09\x47\x02\x8c\x20\xff\xac\x6e\xac\x39\x1b\x9f\x88\xa2\xb6\x52\x4e\x3e\x89\x69\x43\x8e\x05\xa6\x12\x4a\xc2\xa8\x75\x77\x07\xc8\x18\xe0\xaf\xc0\x81\xae\x2a\x00\x96\xd2\x5a\xd2\xa6\x2a\x9b\x62\x45\x84\x29\x02\xb3\xd9\xc1\x00\x02\x8f\xe3\xf1\xb0\xa4\xde\x17\x38\x85\xc2\xe5\x24\x94\xfd\x29\x23\xa8\xdd\xeb\xfc\xa6\xea\xfa\x02\x83\xd9\xec\xee\x0e\x38\x1f\xf1\xd7\x08\x0b\xe9\xfc\xf1\xf1\xd4\x39\x87\xd2\x33\xc5\x86\xb9\x75\xa0\x98\xee\xfd\x4a\x81\x75\x21\xa1\x8c\x84\x05\x1c\x30\x9b\x0d\xba\xf0\x60\xd0\xf5\xc9\x06\xa0\x09\x2c\x25\xa1\x13\x51\x17\x62\xb6\x7e\x1d\xa8\xff\x6b\x18\x0d\x13\x3e\x85\xd0\xd3\x36\x36\x0c\xde\x23\x42\x32\x7e\x5b\x17\x7b\xa6\xfa\x3a\xd0\x63\x36\x45\xc8\x8f\x93\x16\x1e\x06\x3e\xe8\x46\x7e\xc9\x93\x4c\x08\x4a\x38\xb2\xcb\x83\x30\x57\x63\x21\x18\xb3\x1c\x54\x09\x18\x43\x84\x81\x51\x17\x20\x74\x09\xb7\xe5\x4e\x53\x5a\x4d\x0b\x32\x66\x3c\x48\x5a\x56\xf7\xb6\xc7\x38\xf9\x87\x51\x09\x7d\xa0\x7f\xfb\x70\x84\x7d\xdb\xc7\x63\x99\xc4\xbd\x2a\xb6\x00\x82\x12\xda\x92\x4d\x26\xaa\x68\x0a\x7d\x82\xa0\x64\xdc\x02\x04\x0d\x33\xcd\xc7\x16\x50\xb6\x7d\xb6\x8a\x77\x5a\x20\xc0\xd2\x63\x68\x68\x85\x6c\x69\xd7\x52\xa4\x7b\xe8\xba\x8c\x23\xe5\x89\x39\x29\xec\x4c\x79\xbe\xe7\x02\x90\x13\x68\x07\x91\x2f\x89\xc0\x3e\x76\x55\xb9\x7a\xcc\x23\xbc\x42\xd3\x8b\xcd\x2f\xeb\x53\x0b\xeb\xc2\x5c\x4d\xdb\xc3\x10\x11\x3a\xc9\x08\x59\x86\xc3\x76\x7d\x0c\xa9\xa6\xcd\x99\xc1\x65\xbe\x0f\x43\x81\xe3\xe2\x10\x72\x4c\xe5\xd0\xfa\xb1\x48\x13\x26\xe6\xe6\x8f\x92\xba\x19\xde\x5a\x3b\xf8\x26\x84\x14\x61\x14\xab\xc5\x14\x2a\x0f\xe7\xcc\x17\x19\xe1\xee\xd7\x5f\x5d\x13\x5a\x1b\xde\xcb\xbc\x3a\x24\x91\x3e\xce\x84\xea\x61\x2a\x97\x89\x55\xef\x65\x05\x55\x77\x61\x05\x62\x65\xd3\x9c\xe2\x0b\x74\x93\x13\x35\x79\x0e\xd2\x1b\x92\xf5\xb5\x38\xbe\xb5\xe6\x74\x38\xf9\x18\x8d\x6e\x97\x1b\xb6\xa2\xf2\x16\xdd\xd0\x1e\x31\x74\x5b\x91\x09\x48\xbb\x12\xc7\x85\x3c\xc4\xf2\x8c\x21\xfc\x50\x07\xf2\x90\x34\x44\xe2\xc0\x74\x24\x13\xce\xa2\xb0\x86\x48\x9a\xa3\x56\x9b\xe2\x33\xb4\x8c\x6c\x76\xc0\x10\xb6\x31\x55\x11\x8b\xac\x79\x5a\xa5\xfd\xd2\x68\x19\xc4\x49\xd6\x8b\x24\xc7\x7a\x51\x98\x62\x55\xb9\xe6\xce\xa8\x85\x00\x4a\x08\x90\x0a\xe1\xd4\x51\x96\x86\xd7\xd5\x02\xd7\xd4\xcd\xfd\xb4\xf2\x75\x82\xf8\xf5\xda\x88\x75\x0b\x84\x86\x91\x04\xf2\x36\x54\x7d\x8c\x87\xdd\xab\x11\xbb\x49\x75\xfe\x45\xd8\xe2\x9a\x48\xd7\xb3\x00\x85\x01\x2e\x31\x90\x8a\xa7\xc2\x07\xe9\xe0\xd5\x0e\x39\xa1\x32\xeb\x7b\xce\x3b\x43\xd4\x01\x71\xdf\xa3\xb2\x08\xd5\x3c\x46\xe9\x58\x0e\xba\x35\xb5\xd6\x45\x64\x5a\xbd\x6a\xdd\x6a\x4f\x16\x11\x84\x4a\x3c\xe1\xd0\x7f\xb2\x90\x00\x65\x61\x31\x97\x44\xc5\x05\x18\x88\x10\xd2\x74\x52\x84\xbf\x46\x84\x63\x64\x1d\xfc\xcf\xa0\xab\x1e\xd4\x35\xe2\xb7\x13\x37\x12\xdf\x48\x2b\x97\x69\xc5\x96\x29\x0a\x9b\xb9\xee\x16\xe3\x66\xfe\x64\x0a\xfd\x08\xeb\xcc\x2a\x1b\x32\x87\x11\x57\x83\xfe\x49\x4c\xa6\xb2\x5d\x90\x68\x3b\xa3\xf7\x5d\xd8\xe4\xd5\xea\x91\x89\xb7\xdd\x51\xc4\x48\xb0\x0b\x95\xaa\xa1\x62\xf4\xb6\x18\x26\xa6\x74\x69\x88\x1c\x93\x89\xb7\x0b\x8f\x05\x8e\x25\xe1\xe1\xb3\xeb\xed\x46\x87\x16\x60\x17\x1c\x55\x83\x43\xab\x6d\x31\x36\x74\xe1\xd2\xd0\x38\x65\xd7\xbb\xc8\x58\xe0\x58\x12\x19\x42\x42\x5f\xcd\x55\x03\xbc\xdd\x00\xd1\x72\x00\x23\xc7\x2e\x4e\xaa\xc6\x49\xd6\x8a\x8b\xe1\x92\x7d\x56\x1c\x35\x17\x8a\xe2\x92\x04\xf8\xdb\x0e\x9a\x74\x4e\x55\x77\xd5\x60\x0c\xe9\x39\xbb\xc6\xbc\x79\x6b\x06\x63\x48\xed\x50\x89\xf6\xf4\x41\xfa\x1e\x52\x10\x37\xbd\x8b\x4b\x3c\xb4\x38\xa4\x13\x6c\x81\x80\xd0\xa1\xd5\xb3\x40\x00\x6f\x86\xd6\x5e\xaf\xb7\x34\x54\x33\xe6\x53\xd1\x99\xf9\x99\x09\xc8\x8c\xf7\x7d\xc3\x41\xf8\xef\x8f\x03\xa0\xf7\x8c\xe4\x77\x1c\x0e\x66\xaf\x61\xd1\x95\x0b\xfc\x1f\x18\xca\xfd\x42\x7f\xae\x2f\x80\x16\x82\x85\x7a\xab\x28\x0e\xa0\xec\x3a\x7b\x24\x24\x0b\x92\x55\x76\x43\xb6\xd1\xb6\x5e\xfc\x9a\xd9\xea\x4b\xa3\xf6\xc5\xaf\xca\x1d\x0c\x60\xac\x17\xdf\xcd\x5d\x66\xa7\x30\x15\xf1\xff\x23\x92\x38\xd0\xa3\x48\xf8\xba\x57\x24\xe1\xeb\xde\xea\x12\xbe\x85\x3e\xa4\x6e\xb2\x28\xfc\x28\x42\xfe\xf2\xaa\x48\xc8\x5f\x5e\xad\x2e\xe4\x65\xc4\x47\xec\x11\x25\x54\xdd\x7a\x81\x88\x7b\xbd\x0a\x8a\x3c\x83\x37\x8f\x2a\x61\xa1\x12\xf7\x7a\x15\xb4\x78\xc6\x46\x9b\x92\x70\xd0\x35\x0d\x7d\x0b\x03\x55\xbd\x6c\xb1\x62\x63\x15\xc8\x2b\x90\xc6\xe9\x2a\x5a\x5d\xfc\x2d\x6d\xf3\x22\x8a\x1e\x63\x87\x57\xb3\xad\xb7\xb9\xab\xaa\x6e\x78\x5f\xf7\x88\x01\xca\x24\x40\x44\xc8\x88\x8f\xb6\xb9\xbb\xab\xf5\xb2\x7c\x63\xb7\xde\xae\x6e\x0d\xad\x6d\x6a\x43\xf7\x71\xf3\x49\x44\xd1\xf6\xf6\x60\xdf\xed\xf6\x5d\x73\xfb\xae\x39\x63\x28\x0f\xcf\x15\xdc\xdb\x67\x45\x14\xed\x36\x58\x97\x70\x5c\x70\xf3\xed\x2c\xf1\x5d\xa6\x0b\x7a\xff\x7a\x07\x37\x0b\x76\x89\xdf\xe6\xd4\x9d\x5f\x13\xc8\xcc\xf9\x11\x45\x9f\x85\x84\x5c\x3a\xef\x19\x0f\xa0\x04\xd6\x7e\xaf\xf7\xda\xe9\xed\x39\xbd\x7d\xb0\xf7\xaa\xdf\x7b\xa9\x7d\xda\x06\x09\x31\xa6\x68\x29\xe9\x37\xb0\x4a\xd0\x94\x14\xaa\x42\xf6\xb7\xa5\xf4\x89\x49\x4f\x4d\xac\x37\x9f\x40\xc5\x8c\xeb\xa5\x50\xa6\xf2\x86\x93\xa8\xdf\x8d\x44\xdb\xcb\x9d\x62\x95\x3c\x46\xf6\x54\x4b\x61\x1b\x7e\x21\x6e\xca\xfc\x28\x68\xe0\xcb\x70\x46\xae\xa7\x1f\x97\xfe\x8c\xdb\xfd\x2e\x46\xa6\x3a\x4b\xd6\x89\x61\x54\xb4\x24\xf7\x99\x81\x2b\xf5\xa7\x6f\x78\xa9\x7a\xed\xfd\x22\x35\xc6\xff\xc3\x68\x03\xc3\x2a\x91\x6c\x3b\x09\x9f\x69\xf9\xbb\x08\xad\x07\x77\x69\xe7\x96\x50\x91\x34\xff\x95\x89\xa5\x8c\x1b\x7d\x77\xd1\xb4\xf5\x64\x70\x05\xb2\x41\x57\x19\x76\xe9\xf1\x15\x7d\x8e\xa7\xfc\xf0\xcf\x52\xf6\x71\x67\x22\xa2\x20\x80\xfc\x76\x05\xfd\x2d\x9c\xce\x03\x92\x85\x9f\x25\xf1\xb1\xa8\x71\xe6\x03\x52\x12\x40\x89\x11\x18\xfb\x24\x3c\xa1\x9f\x74\x84\xf8\x13\xfb\x65\xd2\x47\xbc\x4c\x82\xe7\x65\xad\xe0\xc9\x1d\x2d\x22\xbe\x7e\x1b\x41\xae\x2a\x68\x11\x13\xe2\x32\x6a\x1d\x0c\x48\x1a\x73\x10\x8c\xa1\xed\xfa\xcc\xbd\x52\x37\xaf\x6e\xac\x83\x41\x97\x1c\xd4\x99\x97\xe4\x7a\x8e\x88\x4a\x9d\x9e\x26\x76\x71\x2e\x99\x84\xc9\x4b\x12\x75\xb8\x7b\x2f\x32\x1b\x17\x8a\x57\xfc\xb2\x8b\x4e\x78\x5f\x6c\x7b\x3e\xf4\x6f\xf4\x09\xce\xc2\x27\x70\x89\x37\x1c\xc3\x8d\xb9\x04\xe4\x18\xee\x5c\xe2\xd1\x5c\xc2\x83\xc2\x93\x70\xf2\x04\x5e\xa1\x0f\x7b\x45\xa1\xd8\x98\x67\xb8\x31\xc3\x26\x78\xc7\x0a\x64\xab\x65\x03\xc9\x46\x16\x96\x90\xf8\x62\x95\xdc\xa1\xf8\x64\xfa\x83\xb5\xf4\xc1\xc8\xcc\x01\x55\x1f\x03\xfd\xaf\x2d\x24\x27\x21\x46\x7a\x31\xc5\x94\x23\x69\x73\x2c\x42\x46\x05\x99\x62\x40\xd9\x35\x87\xa1\x05\x84\xbc\x55\x13\xff\x6b\x82\xa4\xd7\xdf\xeb\xf5\x7e\xaa\xe2\xdc\xd2\xc3\x10\x55\xa1\xe7\x15\xdd\x45\x7a\xb9\x63\xe0\x5c\x66\x5e\xa4\x1c\x74\xa5\xb7\x06\xbb\x77\x14\x6d\x8e\xd9\x51\xc4\xcd\x49\xeb\x4d\x30\x7b\xa3\xfb\xca\xea\x8c\x06\xdd\x2a\xfa\x55\xbc\x2b\x5a\x6f\xc4\xd0\xed\xea\xf4\x77\x77\x40\xcf\xcf\xc1\x33\x13\x06\xa0\x3f\xac\x14\x11\xf3\x76\x79\xe2\xe0\x26\xac\x62\x7e\xce\x21\x0b\x42\x1f\xab\xce\x7b\x36\x13\x91\xeb\x62\x21\xe6\x27\xcf\x91\x6a\x9a\x67\x4e\x96\x57\xb4\x07\x3a\x18\x98\x53\xed\xf1\x7c\x08\x86\xa1\x4f\x5c\x6d\xe4\xee\xfc\xc8\xbb\x75\x80\x98\x1b\x05\x98\x4a\xe7\x9a\x13\x89\xdb\x08\x4a\x7c\xc9\x2e\x24\x27\x74\xd2\x6e\xdd\xdd\xa5\xc2\x6a\xef\x55\x59\x56\x6e\x95\xdb\xee\xed\xd9\xbd\xfd\x4b\xbd\xca\xdd\xef\xbd\xb2\x7b\xbf\xf4\x7b\x3d\xd5\x13\xb6\x3a\x9d\x41\xd7\xb4\x71\x30\xe8\xca\x0a\x66\x7a\x1c\xf1\xdf\x51\xf4\x94\xc2\x67\x5a\x4e\x4e\x17\x27\x21\x66\xc2\x62\x3d\x9e\x69\x6a\x53\x85\x4f\xb5\xf0\xaa\x3e\x6b\x1c\x74\x2b\x04\xd8\xa0\xab\x3b\xf5\x27\x1a\xd2\xe6\x61\xb5\xca\xa8\x47\xe8\x98\x3d\xf8\x79\x83\xad\x0c\x5c\x55\x06\xac\x4a\x83\xd5\x40\x7a\x09\x96\x00\xd9\x2a\x51\xdc\xcf\xec\x01\xfc\x09\x39\x51\x48\x6a\x75\xea\x0f\x70\xf6\xa3\xca\x6c\x57\x73\xe4\x95\xc7\x87\x2a\xe3\x42\x45\xa5\xa2\xcc\xcb\x5c\x6f\x0e\x53\x9c\x95\x06\x2d\xcd\x43\xfb\xa4\x73\xf6\xe6\xb0\x12\x83\xd5\x23\x7e\x0d\x5c\x0c\x61\x7f\x7d\x64\x8a\x4b\xe3\xb0\x1d\x43\x8e\xae\x21\xc7\x60\x8a\xb9\xc8\x66\x47\x75\x61\x26\x0c\xff\x34\xfc\x1a\x07\xf8\x3d\xe1\xc1\x46\x01\x27\x0c\x9b\x0a\xf8\x92\x5d\xe1\xf5\x51\x6a\x2e\x8d\xc3\x76\x4a\xc6\x38\x3f\x43\xa8\x8b\x4f\x71\x4a\x17\xd8\x9a\x04\xf1\x8d\x4e\x9a\x41\xc8\x08\x95\xe0\xed\xc5\xc5\xc9\xd1\xda\x58\x0d\xcb\x73\xc5\xd1\xd1\x1c\x9b\x0d\xfa\xe3\xc5\xc5\xc9\x46\x31\x2b\x86\xcd\x86\xbc\x71\x33\x37\xd2\xca\x1f\xb0\xbc\x66\xfc\x0a\xf8\xcc\x85\x3e\x38\x39\x5f\x1b\x71\xcc\xd0\x39\x55\x0c\x4f\xce\x1b\x0b\x78\x02\x25\xbe\x86\xb7\x1b\xc3\xfb\x7f\x86\x5f\x63\xf1\x06\x50\x5c\x6d\x0c\xec\x19\x14\x57\x8f\x80\x74\xa5\xa9\xdd\x4a\x53\xba\xf9\x5c\x6d\x95\x99\xd8\x35\x19\x93\xdd\x4c\x6c\x37\x13\x2b\xa4\xad\x1b\x78\x17\x12\xca\xb5\x72\x22\xe5\x95\x8e\xe6\xd2\xb8\x4e\xe5\x4d\x24\x3d\x30\x86\xc4\x07\x66\xff\x63\x5d\x94\x8a\xdf\x7b\x48\xfc\x43\xc5\xad\x71\x68\x0f\x19\xa5\xd8\x95\x20\x5e\x3c\xdd\x10\xe6\x98\xeb\x85\x61\xda\x6c\xe4\x1b\x34\x75\xcc\xb2\xb9\xd6\x3e\x3a\x3e\x3c\xdf\x24\x60\xc5\xef\x31\xd1\x6e\x6b\xd0\x14\xfa\x73\xc3\xbb\x61\x73\x37\x6c\x16\xd2\xd6\x5e\xe8\xc3\x42\xc0\xc9\x46\x96\x86\x8c\x87\x3a\x31\xc7\xa6\xae\x0d\x25\x80\x05\xfe\x1a\x61\xea\xae\x95\x32\xe4\x11\x5f\xc4\x1c\x1b\x07\x79\xed\xd4\x28\xc6\xd9\xcc\xe4\xe8\x2d\x94\x12\xf3\xb5\x66\x96\x31\xbe\x98\x13\x98\xcd\xc0\x4f\xcd\xc2\xa8\x77\x1b\xd7\x5e\xf5\x8b\x61\x6a\x66\x8d\x5c\xf8\x33\x30\x61\xf6\x45\x83\x35\x61\xc6\x5b\xaa\x20\x08\x06\x22\x0a\x0f\xf6\x07\x5d\xf5\x5f\xb3\x50\xbf\xe3\x9c\xf1\x0d\x00\xd6\x7c\x1a\x67\xd3\x33\x18\x82\xf8\xab\xff\x9b\xe8\x6a\x61\x78\x6e\x98\x35\x0e\xe8\x09\x05\xee\xc2\x27\xc7\xd7\x00\x7a\x42\x93\x37\x0c\x9a\x08\x94\x63\x19\xf1\x8d\x21\xfd\x98\x70\x6b\x22\xd4\x31\xc7\xc2\x03\x62\x43\x43\xe8\x09\x7d\xaf\xf8\x35\x73\x24\x3d\x85\x23\x20\x72\x7f\xd5\x64\x0d\xa4\xa7\x70\x74\x91\x4c\x59\x9a\x05\x73\xe1\x0b\x6b\x6b\xa1\x7c\x3f\xff\x88\x58\xb3\x40\x1e\x7d\x38\xca\x7f\x68\x7e\x2d\x98\x47\x1f\x8e\xe2\xa3\xff\xdf\xc8\x8c\xba\x3e\xc5\x92\x37\x93\x4a\x1e\x15\x14\x2f\x14\x65\x7e\xc6\xb7\xf1\x7f\x99\x25\x80\xcc\x9f\x86\x52\xf3\xc3\xf4\x18\xcb\x8a\x7f\xac\xe5\xee\x4e\x87\x2e\x71\x8f\x2f\xcf\x4e\x41\xdb\xdc\xff\xf1\xf1\x14\x58\x5d\x04\x85\x37\x62\x90\xa3\x2e\x14\x02\x4b\xd1\x9d\x62\x8a\x18\x17\xdd\x11\x63\x52\x48\x0e\x43\x1b\x41\x89\xf5\x5b\x91\x21\x71\xaf\x30\xef\xba\x42\x74\x17\xca\x9c\x80\x50\xc7\x15\xc2\x02\x63\xe8\x0b\xdc\xc9\xc8\x17\x43\x28\x06\xf3\x45\x3c\x05\x94\xf4\x0f\xca\x74\xbf\x64\x7e\x68\x99\xbf\xdc\x13\xb9\x6e\x23\x01\x0b\x30\x95\xaa\x05\x73\xb7\x61\xf6\xe5\xe6\xf8\x52\x6c\x8d\x7c\xd3\x73\x6f\x5b\xe9\x15\xcb\x9c\xbf\x3e\x6b\x27\x6f\x5c\x76\x1c\x8e\x21\xba\x6d\x8f\x23\x6a\x6c\xd4\xee\x80\xbb\x7b\x2e\x3f\x85\xbc\x34\x80\xc6\x81\x3c\xa1\x60\x08\x5a\x9f\x3e\x7d\xfa\xe4\x9c\x9d\x39\x47\x47\xe0\xf8\xb8\x1f\x04\xad\xe7\xcb\xea\xfc\x1e\x49\x55\xc9\x50\xde\x23\xfc\xed\xfe\x1a\xd9\xb3\x76\xeb\xc7\xe4\xcb\x0f\xad\x8e\xb3\xa0\xa0\xf6\x7d\xa1\xd5\xa5\x88\xcf\x35\x01\xe8\x03\xc9\x23\x5c\x2c\xd3\x9c\x6c\xff\xe5\x31\x8b\x56\x25\x3e\xa1\x2e\xc7\x4a\x8b\xa0\x0f\xf6\x1e\x22\xbe\xc0\x2e\xa3\x48\x80\xbe\xb1\x61\x31\xb9\xb2\xde\xed\xdb\x48\x4a\x46\x0f\x7d\xe5\x2f\xa2\x0f\x5a\x23\x49\xed\x78\x71\xbc\x44\xa7\x7a\x7b\x19\x83\x7e\x81\xe9\x92\x6b\x6c\x5e\x99\xed\x1b\x7b\x15\x92\xcd\x4a\x64\x8a\x24\xfb\x23\x54\xfa\x3e\xa1\x61\x24\x63\xf9\xef\x91\xce\x3a\x0e\xa3\xed\x96\xf0\xd8\xf5\xa2\x71\x5a\xcf\x41\xe2\x5f\x6d\x3c\x7d\x0e\x4c\x69\x91\xa3\xa9\x2b\xf6\x78\x97\x51\x09\x09\xc5\xdc\x19\x13\x8a\xda\x2d\x47\x41\xa4\x08\x72\x5b\x8f\x0c\xad\x8e\xe3\x11\x84\xdb\x9d\xdf\xca\x24\xd1\xba\x5c\x4f\x94\x67\x6d\xe9\x11\xd1\x71\xa6\xd0\x6f\xc7\x62\xe9\xcf\x89\x1c\x41\x89\x1d\xa3\xd1\xb6\xf1\xe5\x0e\xf8\xb9\x05\x6c\xd0\xfa\x39\x26\xc3\x14\x15\x10\x75\x1c\xd7\x53\xb2\x14\x4b\xfd\xb0\xd3\x2b\x39\x0a\x05\x35\x1d\x94\x7e\x57\x7a\xc5\x8f\x9e\x28\x45\x28\x4f\xe8\xdc\xc3\x51\xd8\x80\x06\x57\xf2\x6c\xb1\xf1\x87\x3e\xa2\x52\xd6\x74\x99\x02\x92\x51\x03\xe8\x93\xc1\x7f\x11\xf4\xf7\x73\x30\x2f\x34\x9f\xc4\xcb\x96\x48\x7c\x23\x21\xc7\x50\x51\xb6\x52\x95\xa7\x66\x2f\x33\xf6\xb2\x4e\x4e\x5d\x18\x0c\x13\x87\x28\xef\xd9\xd4\x85\xa0\x84\x60\x08\xee\x4a\xe2\x29\xb9\xae\xc0\x10\x60\x27\xe4\x2c\x6c\xb7\x08\x6a\x75\x0a\x89\x0b\x54\xa2\x2e\x32\x06\xed\xa4\xae\xea\xf7\x5b\x1d\x30\x1c\x0e\x41\x2b\xf9\x0a\x54\xab\x0c\x64\x22\xdf\x5f\x57\x7f\x67\x9a\x8f\x3f\xea\xd4\xea\x80\x1f\x86\xa0\xd5\xba\xef\x9c\xea\x9a\x99\xb7\xaf\x57\xe3\xab\x3c\xb5\x18\x51\xc9\xf6\xc7\x33\x07\x7e\x81\x37\x25\xdd\xb8\xba\x14\xcc\x7e\xce\x1f\x5a\x9d\x58\x7c\xf3\x37\xe1\x5a\x4b\xcc\x12\x71\xbf\xa4\xae\x49\x4d\x96\xd5\x55\xb0\xfa\xfa\xdf\x72\x9a\xb8\x7b\xee\x83\xf9\x38\x5a\xda\xa9\x24\x97\xb2\x21\x77\x38\x16\x91\x2f\x8d\xf5\xc6\x90\xf8\xda\x0a\xcb\x2b\xaa\x8b\xe2\x6b\x70\xfe\x81\x49\x32\xbe\x5d\xa2\xb3\xec\xa5\x3f\x09\xd3\x07\x2d\xbd\x6c\xb4\x64\x6c\xce\xd5\xc1\x37\xb2\x0f\xb8\x13\x98\x15\xef\x15\x2b\x69\x43\xb5\x70\x85\x76\x84\xbc\xf5\x09\x9d\xa8\xb1\x2e\x49\x8a\x5e\xdc\x4f\x0a\x16\xaf\x59\x41\x2f\x9a\x7b\x9e\x9e\x1b\x6f\xa7\x10\xc0\x0f\x4a\xd1\x11\x35\xe9\xea\x63\xeb\xfa\x62\xe9\xa8\x7d\xaf\xd6\x1a\xda\x5e\x9e\x1f\x2c\x5e\x8f\xa4\xef\xd2\xa7\xc5\x4f\x8a\xf8\x2d\x96\x65\x7f\xa7\xa7\x79\xca\x26\x22\xff\x0d\x00\x00\xff\xff\x5f\x69\x69\x95\x6f\x77\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xcd\x4a\xc3\x40\x14\x85\x6f\xd1\x55\x96\xae\x5d\x5c\x41\x8b\x82\x53\x33\xd5\x45\x9d\x76\x5a\x51\x5a\x10\x1b\x90\x12\xdd\x5f\x9a\x31\x1d\x4c\x66\xc2\xfc\x88\x42\x5f\xc3\xd7\xf3\x59\xa4\xa9\x56\x3c\xab\xf3\x71\xbe\x0b\xf7\xeb\x60\xff\x13\x00\xa0\x03\x00\x87\x00\x70\x0c\x00\x7b\x00\x70\x0e\xdb\x0c\x00\x60\xd2\x01\xb8\xfe\x71\x76\x59\xa8\xc6\xba\xc0\x32\x5f\xea\x82\xdd\xc6\xd2\xb3\xdc\x0a\x2c\xd4\xdb\xcd\xab\x5e\x51\x6d\x7b\x2e\x26\x73\xf2\x81\xe5\x8e\x8c\xaf\x28\x58\x27\xf0\xa1\x9d\x30\x8b\x8e\x6a\x5b\x58\x1c\xfd\xf3\xc7\xc9\x9c\x4c\x19\xa9\x54\x2c\x57\x54\x0b\xdc\xb1\xc0\x45\xf4\x5e\x93\x49\xb2\xfb\x6c\xca\x9e\x95\xf3\xda\x1a\x81\xbc\x97\x26\x77\xd6\x04\x65\x02\xcb\x3f\x1a\x25\x30\xa8\xf7\x70\xd1\x54\xa4\xcd\x10\x97\x2b\x72\x5e\x05\xf9\x94\xcf\xd8\xe0\xcf\xdb\xfc\xf3\xa2\x1c\x9b\x9a\xa5\x2d\xb4\x29\x05\x26\x8f\x55\x74\x54\xb1\x99\x75\xb5\x17\x68\x9a\x16\xbd\xbc\x1c\xe2\xb6\x4a\x73\xc2\x53\x29\x39\x76\xbb\xb8\xa9\xe9\x91\xe4\x1c\x27\x98\xa2\x68\x79\x2c\xfb\xbf\xd3\x48\x5e\x6d\xea\x69\xab\x8d\x78\x8a\xeb\xf5\xf6\x64\x2c\xfb\xe9\x19\x4e\x90\xa3\xc0\xfe\x10\xbe\x03\x00\x00\xff\xff\xcf\x65\x2a\xf0\x79\x01\x00\x00"

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
