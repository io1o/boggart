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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\xfd\x72\xdb\x36\x12\xff\x3f\x4f\x81\x63\x93\x89\x3d\x57\x92\xfa\x72\x93\x30\xb2\x6f\x9c\xd8\x6e\xda\x69\xda\x4c\x93\x6b\x27\x77\x73\x93\x81\x48\x50\x42\x03\x02\x2c\x08\xfa\xa3\x1e\xbd\xd3\x3d\xc3\x3d\xd9\x0d\xc0\x0f\x51\x14\x45\x81\x94\xd4\xc8\xaa\x38\xd3\x98\x22\x17\xbb\x8b\xdd\x1f\x17\x3f\xb0\x20\x79\x7f\x0f\x3c\xe4\x63\x8a\x80\xe1\x32\x2a\x10\x15\x06\x98\x4e\x1f\x0d\x23\x97\xe3\x50\x00\x71\x17\xa2\x53\x03\x86\x21\xc1\x2e\x14\x98\x51\xfb\x37\x78\x0d\x93\x93\xc6\xd9\x23\x00\x00\xf0\x63\xea\xca\x33\xe0\x1a\x92\x18\x5d\x31\x1e\x40\x71\x44\x8f\xc1\xbd\x3a\x2b\x37\x8e\x44\xcc\x29\xa0\xe8\x06\x7c\x47\x05\xb1\x7e\x8c\x83\x11\xe2\xa9\xe4\xb1\xe5\x67\x4d\x5e\xaa\x16\xd3\x47\x43\x3b\x31\x70\xf6\x68\xe8\xe1\x6b\xe0\x12\x18\x45\xa7\x06\x67\x37\xa9\xc5\xe2\x51\x97\x11\x33\xf0\xcc\x6e\x0f\xc8\xbd\x28\xc8\xf6\x6e\x23\xb3\xdb\x4b\xe5\xcb\x6d\x6e\x3f\x85\x90\x22\x52\x38\xbb\x28\x91\x05\x63\x5e\x26\x97\xe3\x8c\xa0\x53\x43\xc0\x51\x95\xa6\x5c\x32\x26\x99\x42\x0a\xaf\x01\x85\xd7\xa6\x80\xa3\x08\x8c\x20\xff\x24\x77\x8c\x99\x1a\x82\xa3\x2a\x5b\xb9\x26\x82\x53\xd9\x90\xa3\x08\x51\xa1\x72\x61\xdc\xdf\x03\xec\x03\xf4\x3b\xb0\x60\x92\x02\x43\x26\x2f\xb3\x29\x8f\x5d\x23\x29\x84\xa8\x07\xa6\xd3\xb3\x21\x04\x13\x8e\xfc\xd3\x25\xed\x66\x99\x75\xae\x19\xf6\x8e\x3a\xc7\x2f\x65\x5b\x12\x21\x30\x9d\xde\xdf\x03\xeb\x67\xf4\x7b\x8c\x22\x61\xfd\xf3\xe7\x1f\xac\x77\x50\x4c\x92\xc3\x89\x72\xe3\x4c\x2a\xed\x3e\xa7\xc0\x78\x1d\x73\x8e\xa8\x00\x01\x82\x51\xcc\x51\x80\xa8\x88\x0c\x60\x81\xe9\x74\x68\xc3\xb3\xa1\x4d\xf0\x06\x3a\x0a\xb9\x3b\x91\xbd\x6b\xd9\xdf\x42\xf3\x36\xdd\xfe\x47\xa2\xe7\x34\x55\x53\x15\x85\xf3\xcc\xc2\xea\x8e\x0f\xed\x98\x2c\x39\x53\xc0\xa4\x80\x23\x73\x39\x2a\xeb\xd0\x59\xd4\x20\x8f\x00\x1f\x7a\x08\x24\xe1\x02\x98\xd6\x68\x93\x5b\x6d\xf0\x86\xf2\xe2\xcd\xf4\xcb\x7d\x73\xc2\x38\xfe\x83\x51\x01\x97\x5d\x15\x55\x7d\xbb\x41\x84\x18\x20\x12\x77\xd2\x6f\x76\x8d\xb8\x4f\xd8\x8d\x03\x60\x2c\xd8\x0a\x2d\x4a\x53\x7b\x27\xaa\x9c\x49\x4b\xca\x89\x66\xd3\x2c\x46\x2a\xeb\x17\x50\x20\xf0\x33\xa4\xe3\x34\xf1\xda\x1a\x86\x3e\x46\xc4\x8b\x90\xd0\x37\xba\xe8\x37\x15\x9c\x11\x73\xcc\x59\x1c\x36\x70\xbe\x46\x57\xd4\x42\x4d\x59\x15\xa6\x61\x2c\xcc\x90\xa3\x50\x5e\x23\xc9\xaf\xb6\x2e\xe6\xfa\xa3\x10\xd2\xfc\xba\xf7\x3c\x93\xd1\xa2\x66\x13\x7a\x1e\xa3\xc6\xd9\x10\xe7\xa8\x80\xc0\x87\xa6\x0b\x09\xa2\x1e\xe4\xc6\xd9\xd0\xc6\x67\x43\x5b\x6a\x59\xc3\x09\x65\x31\x1d\x22\x05\xba\x15\x06\xa0\x30\x40\xa7\x86\x07\x05\x32\x00\xf6\xb2\xbd\x22\x32\xd3\xc0\x1a\xc9\x68\xa9\x4a\x93\x25\xa5\x3e\xf9\x9c\x05\x56\x32\x26\x02\xa3\xd7\xe9\x7c\x63\x75\xba\x56\xa7\x07\xba\x27\x4e\x67\xe0\x74\x4e\x54\xa5\x33\x41\x2e\x2f\xd8\x2a\x69\x03\xd8\x2d\xb3\x67\x7b\xf8\xba\x05\x7e\x9a\x37\x6b\xd8\x64\x68\x37\xbb\x4a\x1a\xa8\xaf\x28\x00\x83\x36\x05\xe0\x1d\xe2\x98\x79\x87\x8b\x7f\xab\x17\x3f\x22\xc8\x15\xea\x02\x0b\xd3\x70\x57\x5c\x62\x20\x11\xeb\x19\x80\xa3\xdf\x63\xcc\x91\x77\x6a\x64\x7b\x6b\x58\x57\x1e\xb0\x70\xc6\x78\x4f\x8d\x80\x51\x31\x21\x77\x05\x9a\x91\xb8\x05\xf2\x33\xf2\xda\x4d\xdc\x91\x5e\x64\x7b\x05\xaa\x92\xe3\xe7\x6d\xd6\x22\xa1\x0d\x89\xa1\x8d\x7a\xeb\x41\x5c\xe9\x6b\x72\x5c\xdf\xd3\x8b\x44\x7e\x6b\x7e\x4e\x58\xcc\x2b\x1d\x4d\x4f\xe8\x7b\xfa\x26\x6d\xb0\x19\x57\x87\x76\x62\xeb\x50\x5b\x75\xd4\x6b\x88\x0d\x6d\x79\xd9\xae\x90\x91\x73\x24\x54\x20\xb2\x04\x01\xf5\xaf\x19\x09\x8e\x43\xe4\x01\x0f\x0a\x98\x1c\xf7\x84\xc9\x51\x14\x32\x1a\x49\x7e\x4b\xd9\x0d\x87\x61\x4e\x2e\x6f\xb0\x27\x26\x4e\xb7\xd3\x79\xa2\x43\x2c\xc5\x04\x41\x4f\x47\x8e\x6b\x06\x4d\x4c\xce\xe6\xb8\x62\x06\x4a\x31\x69\xa1\xe0\x92\x22\x3e\xce\x70\x3d\xe2\xc0\x9e\x9d\xfa\xd6\x85\x64\x2d\xdd\xef\x62\x39\x01\x4a\x18\xce\x13\x35\x9e\x51\x4c\xf2\xff\xba\x45\x8b\x1c\xde\x80\xa3\xa4\x55\x30\x8c\xe2\xf0\xac\x3f\xb4\xe5\x1f\x60\x83\x50\x6a\x91\x93\xbf\xe3\x6d\x79\xd2\xdb\x19\x4f\xfa\x3b\xe3\xc9\x60\x7b\x9e\x0c\x6d\x1d\xa8\x4b\x5d\x9a\x17\xce\x88\x79\x77\xab\xe5\x54\x3f\xe8\x18\x81\xc7\xf8\x6b\xf0\x38\x12\x50\x00\xe7\x14\x58\x72\x27\xd2\x21\x5a\x0d\x2e\x50\x0d\xaf\x0b\x5e\x25\x43\xd3\xe3\xaa\x01\xbf\x51\xb1\xce\x32\xab\xfa\x66\xc9\xca\x60\x29\x26\x60\xbd\x17\x1c\xd3\x31\x78\x2c\x47\xbb\xfb\xfb\xe2\xf9\x8f\x08\xf2\x26\x66\xb2\x9b\x1a\x65\x97\xf3\x71\xbf\xa9\xc3\x05\x5f\x2e\xe0\x5d\xea\xe0\x9f\xd5\x8b\x86\xee\xea\xdd\xcf\xf4\x98\x1b\x07\x88\x0a\xeb\x86\x63\x81\x8e\xe4\x3c\xeb\x03\x4b\x7c\x3f\x7a\x3a\xef\x77\x71\xee\x65\x76\xba\x66\xa7\xf7\x21\x9b\x7b\xfd\xab\xf3\xcc\xe9\x74\x64\x44\x9f\x1e\x1f\xcf\xee\x65\x36\xe9\xa1\x22\x30\x9a\xc3\xb0\x2e\x62\x1b\x41\xbb\x55\xbc\x8a\x37\x80\xf3\x68\x25\x83\x14\x98\x4e\x5b\xc6\x02\xfb\x60\x2c\xe6\x74\x7d\xe0\x32\x3e\x1d\xab\xd3\x1c\x04\x85\xbb\x06\x04\x8e\x10\x01\xea\x5f\xd3\x93\xb5\x85\xb7\x99\x35\x6d\x38\x4c\x17\x88\x08\x38\x1f\x2b\x60\x7f\x71\xaf\x92\x80\xb7\xcb\x20\x48\x18\x73\xa3\x1b\x2d\x85\x4a\x45\xb6\x9d\xfa\x28\x76\x5d\x14\xb5\x99\xea\x1e\x72\xaf\xe5\x4f\xfb\xdc\xcb\x51\x6a\xab\xb9\xc7\xd4\xd7\xb9\xbd\xbc\xa8\xef\x90\x78\x0d\x7f\xda\x27\xbe\x54\xef\x93\xf8\x6c\xb4\xde\x23\x1f\xc6\xa4\xee\x7f\xa3\x2c\x57\xb9\xf5\xdc\x6f\x3d\xca\x7b\x44\x2e\xd4\x3c\xa8\xfb\x0b\x23\x71\x80\xca\x57\x10\x28\x49\x35\x64\x99\x45\x1c\x26\x0a\xf6\x92\x77\x24\x5d\xdb\xb5\x12\x54\x0c\xf8\x97\xe5\x1d\xdb\x4a\xfd\x2e\xf0\x8e\x43\xee\x67\x5b\x15\xef\xd8\x56\xee\xbf\x38\xef\x38\x24\x7e\xb6\x55\xf1\x8e\x62\x7c\xf6\x8d\x77\x2c\xcd\xfd\x81\x77\x34\x0c\x62\x4f\x8b\x77\xf4\xd6\xe5\x1d\xbd\xfd\xe5\x1d\xbd\x9d\x2c\x41\xbd\x9d\xe1\x1d\x5b\x48\xfd\xce\xf0\x8e\x43\xee\xd5\xb6\x94\x77\x6c\x21\xf7\xbb\xc1\x3b\x0e\x89\x57\xdb\x52\xde\xd1\xdb\x63\xde\x51\x95\xfb\x03\xef\x68\x18\xc4\xbe\x16\xef\xe8\xaf\xcb\x3b\xfa\xfb\xcb\x3b\xfa\x3b\x59\x82\xfa\x3b\xc3\x3b\xb6\x90\xfa\x9d\xe1\x1d\x87\xdc\xab\x6d\x29\xef\xd8\x42\xee\x77\x83\x77\x1c\x12\xaf\xb6\xa5\xbc\xa3\xbf\xc7\xbc\xa3\x2a\xf7\x07\xde\xd1\x30\x88\x03\x2d\xde\x31\x58\x97\x77\x0c\xf6\x97\x77\x0c\x76\xb2\x04\x0d\x76\x86\x77\x6c\x21\xf5\x3b\xc3\x3b\x0e\xb9\x57\xdb\x52\xde\xb1\x85\xdc\xef\x06\xef\x38\x24\x5e\x6d\x4b\x79\xc7\x60\x8f\x79\x47\x55\xee\x1f\x1c\xef\xd0\x5b\xf3\xad\x6f\x79\x68\x6b\xac\xfa\x1e\xda\xea\x61\x8e\x95\x8f\x26\xeb\x2c\x06\xde\xed\x87\x47\x40\x93\xf5\xe9\xa0\xfc\x24\xc0\xdb\xd9\x43\xf6\x4d\x9f\xf7\x58\xd0\xf5\x8b\xc4\x71\xe3\xa7\x46\xbe\xcc\x13\x01\xa0\x79\xd4\xbc\xf9\x47\x6f\x00\xa4\x1e\x10\x38\x98\xf5\xb8\xf1\xd2\x7f\xf5\x24\xae\x54\x61\x5d\x72\xce\x1a\xad\x63\x4f\x5d\xca\x41\x89\x6e\x45\x4e\x53\xb3\x67\x7c\x67\x9a\x73\xfd\x8d\xbd\x6c\xb5\x58\xbe\x89\x8d\xbc\xd1\x46\x56\xd8\xcf\xfa\xad\xc0\xb8\xc5\x65\xf6\xd9\x96\x64\x92\xa2\x82\xe9\x74\x40\x6a\x1a\xb6\x24\x0a\x82\x33\x3a\xae\x4c\x6b\x63\x65\xa0\x38\x43\xda\x88\x7b\x45\xad\xdd\xe7\xf4\x1d\x89\x39\x24\xc0\xf8\xfb\x13\x0f\x44\xc8\x65\xd4\x33\x8a\x3f\x22\x63\xc1\xea\xe3\xb9\x07\x8d\xca\x67\x5b\x7a\xd4\x12\xa6\x35\x1d\x2a\xf6\xe7\x8b\x74\x47\x7f\x0c\x2e\x6e\x43\x3b\x81\x4f\x2b\x14\xb7\xb0\xd9\xbc\x9a\x34\xe2\x16\xda\xcf\x5c\xb5\x2d\xe3\x1f\x50\x10\x22\x0e\x45\xcc\xd5\x7b\x4c\xda\xd7\x71\x31\x53\xf4\x09\xd3\xcd\x57\xf3\x2a\xfd\x7f\x7e\x4d\xdf\x04\xb9\x2d\x77\x45\x95\xe9\xd2\xe4\x26\xcf\xcf\xff\xfe\xfb\xda\x68\xd5\xc1\xdd\x85\x19\x8b\xc5\x86\x70\xc6\x62\xb1\x5d\xa0\xe5\x06\xf6\x00\x69\xb2\x2f\x7f\x35\xa8\x79\x72\x10\xda\x10\xd8\x94\xae\xed\xc2\xad\x60\x62\x0f\x00\x97\xf4\xe6\xaf\x00\xb9\x77\xec\x06\xf1\x35\x60\x16\xca\xf6\x9b\x87\x56\x41\xed\x03\x85\x53\xd2\x83\x5a\x08\x7d\xeb\x42\x62\x4f\xf6\x01\x45\x73\xef\xa2\x68\x03\x23\xa4\x14\x6c\x1e\x47\x45\xbd\x0f\x14\x48\x69\x17\x56\x22\x69\x1f\x70\xf4\x1a\x86\xd0\xc5\x62\x1d\x24\xb9\xa9\x8a\xcd\x63\x69\x5e\xf3\x03\x45\x53\xde\x89\x25\x78\x5a\xf6\xc2\x10\xe3\x78\x1f\xe0\xc5\x68\x14\x07\xea\x0d\x50\xeb\x20\x6c\xa6\x65\x0b\x20\x2b\x2b\x7f\xa8\x38\x2b\xf4\xa3\x21\xd4\xec\xc9\x5e\x80\x4d\xe3\x1d\x4a\xad\x08\x57\x1c\x11\x39\xf7\x0e\x63\xf1\xa9\xbb\x05\xe2\xb5\xa8\xfe\x81\x22\x70\xbe\x27\xb5\xc3\x67\xf6\x2a\xa4\x7d\x07\x5d\x6f\x13\xa0\xeb\x6d\x17\x74\xbd\xbd\x01\x5d\xef\x00\xba\xfc\xcd\x68\xeb\x82\xae\xbf\x5d\xd0\xf5\xf7\x06\x74\xfd\x03\xe8\xf2\x97\xe0\xad\x0b\xba\xc1\x76\x41\x37\xd8\x1b\xd0\x0d\x1e\x0e\xe8\x36\xbb\xfc\xa4\xde\xbb\x9a\xb7\x9f\x2e\x39\x55\x71\xb8\x74\xa8\xf0\x33\xdd\x4d\xff\xcc\xdc\x79\x54\xf8\xaa\xc7\x04\x41\x2f\x7f\xbb\xdf\x90\x60\xfa\x39\xfd\x28\xc2\x44\x88\x30\x72\x6c\xdb\x65\x84\x71\x82\x47\x96\xcb\x02\x3b\x64\xe4\x6e\xcc\xa8\x7d\x8d\xa8\xc7\x78\x64\x8f\x18\x13\x91\xe0\x30\x34\x3d\x28\x90\x7a\x07\x63\x88\xdd\xcf\x88\xdb\xa5\xdf\x96\x1b\x45\x06\xe0\x88\x9c\x1a\x6a\xb1\x4c\x34\x41\x48\x18\x67\x8f\x12\xb3\xea\x50\xe1\x4d\xe9\xb6\x3b\xb7\x24\xd3\x2a\x29\x03\x96\xfa\x11\x01\x82\x0b\x9f\x11\x91\x9b\x72\xd6\x01\x5f\x3d\xeb\x3f\x7f\xf6\xe2\x75\x7e\x6a\xaa\xa3\xca\x4a\xbe\x78\xf0\xb5\x8e\xac\x33\x61\xd7\x88\x97\x8c\x8f\xa0\xfb\x79\xcc\x59\x4c\x3d\x07\x7c\x75\xd2\xff\xe6\xfc\xd9\xd5\xcb\x79\x01\xc6\x3d\xc4\x1d\xd0\x0d\x6f\x41\xc4\x08\xf6\xaa\xc5\xb2\x4e\xf8\xbe\xaf\xd5\x83\xe4\xb5\xd9\x01\xa6\xe5\x68\xcc\x1c\x32\x33\x9d\x08\xa1\x95\x3e\xb9\xae\x5b\x96\xb9\x35\xa3\x09\xf4\xd8\x8d\x03\x28\xa3\x08\xfc\x0d\x07\x21\xe3\x02\x52\xd1\xd0\xc1\x34\xc6\x65\x3f\x2b\x7d\xd0\x51\x9d\xbc\x6a\xd9\x52\xaf\xd5\x4c\xf6\x6b\xd2\x97\x0a\xdf\x21\xc8\x75\x65\x27\x2c\xd6\x96\x0d\x30\x8d\x05\xd2\x95\x4e\xd7\x1c\x68\x4a\xc3\x20\x0c\xd2\xd7\x9a\xcf\x87\xce\x67\x54\x98\x11\xfe\x03\x39\xa0\xdb\x0b\x6f\xe7\xf3\x16\x42\xcf\xc3\x74\xac\x02\x3b\x7f\x66\x82\xf0\x78\x22\x92\x2f\x68\xcc\x9f\x09\x20\x1f\x63\xea\x80\x4e\x09\x93\x31\x8f\x24\x80\xd2\xa5\x93\xd5\xda\xfa\x9d\xb2\x9d\x8a\xcc\x9e\x5f\xbc\xea\xbd\x3a\x99\x17\x23\x98\x22\x73\x95\x16\x93\x43\x0f\xc7\x91\x03\x3a\xe1\x6d\x43\x04\x2e\xc2\xa4\x14\xc6\xa4\xd7\x26\x4f\x3c\xe8\x3d\xd1\x51\x2a\x3c\x4b\x36\x51\xef\x9d\xad\x29\x04\x97\x83\xcb\x67\x97\xe7\xd5\x57\xb8\x7e\x99\x12\xde\xea\xf2\x94\xcb\xac\xa8\x4e\x79\x31\xd8\x40\xf1\x11\x13\x0b\x5e\x43\x4c\xe4\x68\xb8\xba\x28\x2e\x54\x9f\xcc\x5a\x7f\x30\x78\x71\x72\xa9\x61\xd0\x19\x21\x9f\xf1\x9a\x28\x38\xd0\x17\x0b\x4e\xa4\xdf\xc5\x49\xaa\x97\x56\xc9\xca\x3e\x05\x62\x45\x98\x8e\x49\x39\xc1\xf9\x45\x02\x3a\x60\x10\xde\x82\x4e\x2b\x9d\xf9\x01\x33\x59\x5b\x3a\x6f\x23\x59\x4c\x0a\x7a\xbd\xc1\xd2\xcb\x3a\x37\xdf\xb4\x20\xaf\xf2\x44\x2d\xc6\x04\x82\x3b\x3e\xe6\x91\x30\xdd\x09\x26\x1e\x10\x93\x92\x87\xb9\x1f\xcf\xc3\x5b\x70\x12\xde\x6e\xd2\x74\xd9\x56\xb9\x04\x68\xd8\xb2\x92\x3f\x9f\xba\x4b\x38\x82\xef\xfb\x2f\x97\x43\x55\x1b\x91\x33\x33\xf5\xf9\xdc\xb4\xf6\x2c\x47\x75\x66\x7a\xdd\x7e\x7f\x70\xb2\x11\x33\x73\x50\xa8\x31\xd9\x3d\x7f\xf5\x4a\xab\xa4\xd5\x98\xf4\x2c\xe6\xfb\xab\xc3\x57\x5d\x4a\x5e\xbc\x78\xb1\xae\xf5\xfa\x8a\xd6\xb8\x66\x65\xb6\x7a\xba\x39\xcc\x2c\x34\x8e\xe5\x72\x0b\x35\xe9\x6b\x3c\x18\x65\xd6\xfa\xeb\x58\xab\xbf\xfa\x1a\xf7\x7c\xb0\x3d\x5f\x1a\xe7\xb9\xc2\x17\x6f\xf9\x80\x55\xd7\x6a\xd5\x75\x70\xf9\xfa\xaa\x73\xd5\x5d\x49\xb9\x16\xfb\xd7\xbe\x9a\x56\xba\x59\xcd\xea\x6b\x83\xa8\x33\x44\x10\xe4\x2f\xf0\x5d\xc2\xa0\x70\x80\x3a\xb3\x74\xd0\x1b\xda\x6a\x42\xb9\x6c\xca\xfb\x5b\x94\x4f\x78\x6b\x3f\x5e\x37\x9b\x52\xa7\xb7\x43\x22\xee\x6a\xce\x8b\x03\x16\x20\x2a\xec\x00\xd3\x74\x57\xce\x0e\x2c\x69\xb8\xfe\x96\xca\xe2\x3a\xf5\xe6\xb6\xf5\xe7\xe4\x4d\xfc\x59\x70\x48\xe3\xd3\x9f\xd9\xf6\xf8\x28\xbb\x55\x74\x6c\x71\x04\xbd\xbb\xa3\xfc\xa3\xa0\x47\xc7\xa5\x04\x27\xf2\x4f\xbf\x92\xbe\x3e\x3d\x2e\x83\xe3\x68\x51\x58\x6e\x02\x07\xe8\x5d\x42\xff\x80\xe0\x71\x81\x1e\x56\x4b\xf5\x06\x6f\x58\xac\x29\xfb\x1d\x75\x93\xe7\x67\x1c\xf0\x4d\xa7\x5a\x38\x80\xb7\x17\x50\x20\x07\x24\xb9\x3e\x3a\xae\x16\x23\x4c\x22\xdb\xa9\xe8\x70\xb6\x25\x9f\x3a\x75\xc0\xd3\x8f\x1f\x3f\x7e\xb4\xde\xbe\xb5\x2e\x2e\xc0\x9b\x37\x4e\x10\x38\x51\xf4\xf4\xeb\x47\x4b\xdb\xa9\x34\xdc\xfd\x00\x47\x88\x18\x0e\x30\xce\xe5\x2f\xa3\xda\x09\x25\xee\x42\xea\x22\x92\xcb\xbf\x56\x3f\xeb\x1a\xf8\x9c\x05\xb9\xf8\x15\x67\x41\x9d\xb0\x60\xb9\xe8\x07\x56\xeb\x46\x1c\x09\x16\xa8\xcf\x21\xce\x7c\x51\xc7\xea\x5a\xdd\x20\xf4\x39\x17\xff\xb5\x4e\xd2\x83\x77\xd1\x4f\xfe\xaf\x08\x7d\x36\x1c\xf0\xef\xa5\x72\x4a\xf6\x7d\x5c\xa3\x49\x49\xbc\xad\xeb\x8b\x92\xf8\xb0\x52\xc7\xaf\x68\xa5\x8e\xc9\x2a\x89\x2b\xbe\x4a\xe2\x3d\x34\x96\x0a\xfc\xa7\x26\x5e\x6a\x7e\xfc\x23\x0c\x50\xb4\x3a\x5e\xf9\xbd\xdb\xef\x21\x8d\x21\x4f\x17\xe4\xac\xf2\x2c\x6f\x75\x85\x46\xbc\x45\xb3\xb7\xb2\x3c\x37\x6c\x73\x1e\x72\x4c\x1a\xdb\x69\xea\xd9\xf7\x31\x45\x8d\x9b\x90\xa6\x56\xce\xe3\x71\x1c\x89\x86\x8d\xde\xa3\x50\xa0\x60\x94\xad\xe2\xd4\x6e\xf7\x93\x2b\x58\xf3\x56\x3f\xb2\xeb\x36\xc6\x2e\x90\x5b\x6c\xd6\x0a\xbf\x8a\xe9\x5d\xc0\x3b\xc3\x01\xdd\x4a\xa9\xc5\x3b\xf1\xd3\xe3\x79\x66\x54\xfc\x3d\x3f\x18\xcf\xb8\x44\xbe\xf7\xff\x00\x00\x00\xff\xff\x60\x1b\x65\xd1\x1d\x7b\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x53\xcd\x8f\x53\xd5\x1b\x7e\xe0\x96\xb6\xf4\xf7\x13\x05\x11\xbf\xe1\xa0\x0e\x6a\xc8\xa5\xed\x0c\xa0\x76\xa6\x83\x38\x80\xe1\xa3\x42\xa0\x12\x77\x78\x69\x8f\xed\x8d\xed\xbd\x37\xf7\x03\x9c\x48\xe2\x7c\x84\x8c\x88\x11\x12\x35\xb8\x40\x40\xa3\x3b\x8c\x65\x98\xc6\x01\xa6\x9d\x8d\x0b\x17\x2e\xde\x13\xdd\x19\x8d\x1b\x37\xfa\x07\x18\x5d\x99\x73\xcf\x81\x99\x0e\xd1\x2e\xee\xfb\xbe\xcf\x79\x3f\x9e\xf7\x39\xa7\xbf\xac\x49\x7c\x0c\x00\x4f\x02\x78\x0c\xc0\xd6\x65\x40\x16\xc0\x81\xe5\x88\x7f\x69\x03\x58\x0b\x60\xa5\x01\xac\x03\xb0\xde\x00\x56\x00\xd8\x66\x00\x29\x00\x83\x06\x90\x04\x30\x62\x00\x69\x00\xfb\x0c\xe0\x7f\x00\x8e\xe8\xba\x63\x06\x90\x00\x10\x1a\xc0\x3d\x00\x46\x75\xde\x94\xae\x3b\xaf\xe3\x0b\x3a\xef\x8a\xc6\xbf\xd4\xfd\xaf\x6a\x7c\x46\xdb\x1b\x7a\x7e\xd7\x00\x0c\x00\xdf\xea\x79\xdf\xe9\x3e\x3f\xea\xba\x9f\x75\xde\x6f\x06\xb0\x0a\xc0\x1f\x06\xb0\x12\xc0\xdf\x06\xb0\x1a\x40\x22\xa1\xf0\x75\x09\xe0\x5e\x00\x7d\x09\xbd\x57\x42\xf1\x1c\x4c\xa8\x7d\xf7\xea\xf3\xd7\x13\x8a\x97\x97\x00\xea\xcb\x80\x93\x09\x60\x00\xc0\xaf\x2b\x80\xed\x00\xd6\x26\x81\xff\x03\xd8\x9f\x04\x32\x00\x5e\xd3\x71\x45\xc7\x9e\x8e\xdf\x4e\x2a\xad\xdf\x4d\x2a\xbe\x57\x93\x6a\x4e\x3b\xa9\xf8\x7c\xaf\xed\x4f\xda\xfe\xae\xf3\xfe\x4a\xaa\x3d\x97\xa7\x54\x9f\x55\x29\x85\x3f\xaa\xed\x53\xda\x9a\x29\xc5\x73\x7b\x4a\xe9\x5f\xd4\xf9\x87\x53\xaa\xdf\xb1\x14\x70\x1f\x00\x27\x05\x30\x00\xef\xe8\xf8\x4a\x0a\xd8\x08\xe0\x5a\x4a\xf1\xfb\x21\x05\x3c\x0d\xe0\x4f\x7d\xbe\x21\xad\xea\x73\x69\xe0\x09\x00\x3b\xd2\x4a\xc7\xe3\x69\x60\x8d\xbc\xdf\x34\xb0\x0c\xbd\x3f\xf9\x9e\xe4\x1d\x49\xfd\x1e\x06\xf0\xc8\xa2\xb3\x07\xa0\xf6\x91\x5a\x4b\x1e\x52\x03\xc9\xfb\x71\x7d\x2e\x39\xac\x5e\x94\x7f\xbf\xdc\x55\xfb\xf2\x1d\xc8\x9d\xe4\xbe\x1b\xa0\xe6\x3f\x04\xc5\x4b\xee\x20\xf5\x7e\x50\xe7\xae\xd5\x76\xa5\xb6\xf2\x8e\xe5\x3e\xf2\x79\xcb\x77\x22\x77\x5a\x2f\x0f\xfa\xaa\x2c\xe0\x15\xd7\xa9\x2e\x78\x01\x36\x2f\xa0\x9b\x17\xc1\x3b\x3d\xdf\x6e\x60\xa7\x5f\xa9\xdb\x27\x38\x76\x46\xb5\x28\x08\x31\x62\x79\x56\xc5\x0e\x47\x31\xe2\x3a\x41\xd4\xf4\x42\xdb\x75\x30\x12\xf9\x3e\x77\x42\xd6\xe4\x56\x10\xf9\xbc\xc9\x9d\x30\xc0\x2e\x2b\xe4\xf1\x87\x59\x4e\x95\x85\x76\x93\x63\x17\xaf\xf0\xe6\x71\xee\x63\xb7\xc3\xfd\xda\x28\xf6\xf0\xe3\x7e\x64\xf9\xa3\x78\xb9\x62\x35\xe2\x4f\xb6\x8e\x7d\x96\x13\x63\xfb\xa2\x86\xfc\x38\x1c\x25\xcb\xaf\xd4\x51\xb2\x46\x51\x5a\x18\x80\x57\xdc\x13\xaa\xd9\xc1\x4a\xe8\x4a\x7b\xc8\x3d\x29\xbf\x51\x23\xe0\xcc\x76\xbc\x28\x64\x7d\x55\x1c\xe1\x5e\xa8\xd2\xca\xbc\xe9\x71\xdf\x0a\x23\x9f\xb3\x2a\x6f\x84\x56\x0f\x62\x3b\x3d\xa1\x1b\x85\x38\x6a\x35\x22\x8e\xe6\x50\x10\x79\xc3\x03\x43\x59\x69\x7a\x23\x96\x65\x9e\x9c\x16\xf4\xc2\xd9\x3a\x34\x7c\x98\x7b\xae\x1f\x9a\xa5\xa0\x66\x57\xcd\x97\xa2\x5a\x60\x96\xdd\x02\xab\xf2\x13\x2f\xbe\x69\xd7\xad\xa6\xbb\xc5\x8f\x32\x87\x0e\x96\xcd\x11\x9f\x5b\x52\x48\x53\xca\x55\x60\xfd\xb9\xfc\x0b\x66\x6e\xc0\xec\x7f\x8e\xf5\x0f\x14\xb6\x6d\xdb\x9c\x1b\xc8\xe5\x32\x07\xac\x20\x34\xcb\xbe\xe5\x04\x0d\x2b\x74\xfd\x02\xdb\x1f\xf7\x60\xa5\xc8\xb7\x9a\x6e\xd5\x65\x43\x3d\x8d\x87\x33\x07\x2c\xa7\x16\x59\x35\x6e\x96\xb9\xd5\x2c\xb0\x3b\x71\x81\x1d\x8e\x82\xc0\xb6\x9c\x4c\x69\x6f\x69\xb7\x79\x94\xfb\x81\xed\x3a\x05\x96\xdf\x92\xcb\x8c\xb8\x4e\xc8\x9d\xd0\x2c\x8f\x7a\xbc\xc0\x42\xfe\x56\x98\xf5\x1a\x96\xed\x0c\xb2\x4a\xdd\xf2\x03\x1e\x16\x5f\x2d\xef\x31\x9f\x5f\xc8\x93\x7c\xde\xe0\xbe\xb9\xdb\xa9\xb8\x55\xdb\xa9\x15\x58\xe6\x50\x23\xf2\xad\x86\xb9\xc7\xf5\x9b\x41\x81\x39\x5e\x1c\x06\xc5\x81\x41\xa6\xdc\xa2\xd3\x97\xcf\x15\x8b\x79\xb6\x69\x13\x93\x6e\x6e\x63\x31\x9f\x67\x3b\x58\x8e\x15\xe2\x78\xb8\xd8\x7f\xfb\x68\xa8\xb8\x55\xba\xcf\xc4\x69\x43\xf9\x1c\x3b\x75\x4a\x95\x0c\x17\xfb\x73\xcf\xb2\x1d\x2c\xcf\x0a\xac\x7f\x50\x3e\x66\x31\x4e\x6d\xba\x29\x26\xa9\x43\x33\xd4\x5a\x8a\x88\xb3\x4b\x91\xf8\xa9\x2f\x29\x5a\x0a\x89\xb3\x77\x41\xa0\xf3\x34\x2f\xc6\xa8\x4d\xb7\xc4\xfb\xa0\xf3\x62\x4c\x9c\xa6\x59\x9a\x96\xf8\x34\x5d\x17\x93\x62\x5c\x4c\x80\x2e\xd1\x35\xf1\x1e\xb5\x69\x0e\x74\x85\x5a\x62\x5c\x9c\xa6\xae\x2c\xfe\x5c\x75\x12\x67\x68\x96\xda\x8c\xe6\xa9\x4b\x37\xa9\x45\xdf\x50\x4b\x4c\xc8\x9e\x34\x0b\xfa\x28\x0e\x5a\x77\x1c\x46\xb3\x8c\xa6\xe3\x99\x73\xe2\x9c\x84\xdb\x71\xd1\x35\x31\x26\x29\x7c\x45\x1d\x6a\x8b\x31\xba\x4e\xb3\xf2\xf4\x0b\x6a\xc7\xc9\x2d\x45\xf0\xc3\x38\xf5\xd6\x1d\x27\x2b\xa6\x40\x5f\x53\x87\xa6\xa9\xa5\xea\x3f\x11\x1f\xa8\x54\xe9\x74\xa4\x73\x31\x3e\x9a\x88\x1d\xba\x01\xba\xbc\x94\xa5\x4c\xfa\x94\xba\xe2\xdc\x6d\x0e\x97\xe8\xa6\x98\x58\x08\x2f\x52\x57\x9c\xa1\x0e\x75\xa5\x18\x71\x6b\x9a\xa3\x79\x31\x29\x2b\xc5\x38\x75\xc4\x59\xba\x21\x57\x8a\x45\x91\xff\x52\xfa\x8c\xda\xd4\x59\xdc\xe2\x72\xbc\x53\x9b\xe6\xa9\x45\x33\x2c\x1e\x3b\x47\xf3\x31\xd8\x12\x13\x62\x52\x49\xd9\x03\x49\xab\x24\x9d\xa1\x96\x98\x92\x52\xfe\x4b\x4e\x57\x4e\x89\x97\xe9\x50\x4b\x8a\x76\x21\x76\xa6\x24\x09\x79\x31\xa0\xb9\xff\xfc\xd7\xd3\x6c\xcf\x3e\x5d\x79\xfd\x3d\x15\xb1\xca\x77\x27\xfd\x13\x00\x00\xff\xff\x78\x10\x91\x8b\x87\x08\x00\x00"

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
