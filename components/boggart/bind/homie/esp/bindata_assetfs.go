// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package esp

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x6b\x73\xdb\xb6\xd2\xfe\x9e\x5f\xb1\x45\x9d\x91\xd4\x9a\x94\x6c\xb7\x79\xdf\xa3\x48\xce\xa4\xcd\xe9\x34\x67\xd2\x24\xd3\xb8\xed\x87\x34\xc7\x03\x11\x90\x04\x07\x04\x58\x00\x94\xac\x78\xf4\xdf\xcf\x80\xe0\x4d\x12\x29\x51\xb6\x9c\xcb\x4c\x34\x13\x87\x04\x17\xbb\xc0\x5e\x9e\x5d\xac\xa8\x9b\x1b\x20\x74\xcc\x04\x05\x14\x48\x61\xa8\x30\x08\x96\xcb\x07\x03\xc2\x66\x10\x70\xac\xf5\x10\x45\x78\x42\x3d\xc3\x0c\xa7\xe8\xfc\x01\x00\x40\xf9\x61\x32\x7e\xc9\xe9\xd8\xa4\x0f\x13\x82\xe9\xd9\xf9\xcd\x0d\xb0\x93\xff\x17\x80\x9e\xd1\x19\x0b\x28\x3c\xd4\x08\x8e\x40\x30\x9e\xff\xf3\x05\x0e\x29\x2c\x97\x83\xee\xf4\x2c\x65\xdc\x25\x6c\x76\xfe\xa0\x46\x88\x62\x93\xe9\x8a\x94\x12\xc5\x58\xaa\xd0\x9b\x28\x19\x47\x10\xc5\x9c\x7b\xeb\xb4\xeb\xf4\x4c\x44\xb1\x71\x13\xd6\xa8\x12\x4a\x8e\x47\x94\x67\xb4\x23\x23\x6a\x29\xed\xc7\xee\x74\x0c\xfe\x4f\x4c\x10\xff\x37\x6a\xb0\xff\xc6\x60\x13\x6b\xff\xb9\x76\x17\xaf\x04\xb7\xfa\x5d\x2e\x2b\x67\x0f\x46\xb1\x31\x52\x80\x59\x44\x74\x88\xdc\x0d\x2a\x89\x06\x2b\x9e\x05\xd2\x5d\xcc\xb1\x12\x4c\x4c\x92\x6b\x1d\xa2\x4a\x8e\xd9\x87\x60\x83\x3d\x23\x27\x13\x4e\x87\x28\x94\x04\x73\x94\x8e\x61\x35\xa1\x66\x88\xbe\x75\x83\xbb\x99\x24\x74\xce\x05\x86\x28\x37\xec\xcf\x52\x8c\x99\x0a\x41\x51\x6d\xb0\x32\x40\x12\x43\x23\xf0\x61\xb9\x6c\xcc\x74\x24\xc9\xa2\xc4\xf3\x2f\xc6\x39\x8c\xe8\x1a\x4f\x1f\x9e\x2a\x0a\x0b\x19\x83\x8e\x15\x7d\xb2\xa7\x88\x00\x73\x3e\xc2\xc1\xfb\x21\xc2\x81\x61\x52\xb4\x5b\x29\xf7\x56\xe7\x71\x8d\x49\x13\xc3\xb0\xdc\xb9\xb0\x86\x31\xf6\x22\x39\xa7\xca\x93\xe3\x31\x82\x75\x55\xfc\xee\x18\xa6\x0b\x3b\x1f\x74\x59\x35\xdf\x41\xd7\x19\xb8\xd6\x91\xa8\x20\x87\x71\x14\x82\xc5\x84\xaa\xcf\xcd\x4f\xe8\x61\xbd\x64\xce\x22\x0a\x81\x65\x3f\xb9\x3f\x4f\xa1\xfb\xfa\x09\x55\x58\x53\xb5\xe9\x24\x7f\xb1\x88\xde\xc1\x43\x06\xdd\x04\x95\xd6\x30\xcd\x21\x66\xc5\x6d\x7a\x99\x41\x6a\x19\xfc\x02\x4e\xb1\x1a\xb3\x6b\xbb\x8c\xcd\xa7\x4a\xce\x2b\x70\x3e\x90\xdc\x0b\x89\x77\x72\x0a\xf6\x4a\x87\xd9\xd5\xb5\xf6\x4e\x4e\x6b\x40\xf9\xfa\x32\xc2\x82\xf2\x2d\x30\x7c\x7d\x59\xce\x2b\x2b\x54\xd3\xd3\x8d\x14\x82\x8d\x51\x6c\x14\x1b\xaa\x9d\x1a\x07\xdd\xe9\x69\xc5\xcc\x38\x47\x6e\x81\x67\x20\xf0\x6c\x84\x95\x4b\x08\x90\xac\xe7\xd2\x48\xc9\x47\xf2\xba\xc6\xa6\x03\xce\xce\x07\xb8\xb4\x6f\x8e\x23\x4d\x3d\xce\xc4\x7b\x74\x5e\x36\xb6\xb5\x75\x30\xa5\x33\x25\x85\x67\x93\x83\xb5\xe9\xa0\x8b\xcf\x07\x5d\x5e\x61\xdc\x41\x37\xe6\x15\xa3\xdb\xcc\xb2\xc5\xd0\x9b\x8a\xcc\xb2\x77\x85\x8c\x28\xcf\xa3\xf4\xda\x78\x61\x6c\x28\x81\xb1\x14\xc6\x3b\x39\x83\xd0\x1b\x79\x67\xbd\xed\x99\x8d\x63\x6d\x2e\xe3\x88\x60\x43\x2f\x09\xe5\x06\xd7\x41\x54\x3a\x47\xe1\x39\xb4\x9d\xe5\x5e\x60\x6d\x40\xba\x04\x68\x58\x68\xcb\x00\x68\x3f\xd4\x1d\x6b\xc0\x72\x2d\xd0\x8e\x14\x13\x06\xd0\x40\x07\x8a\x45\xc6\x81\xdc\xdf\x08\x47\x11\x67\x01\xb6\xb1\xd8\xbd\xc2\x33\xec\x9e\xfe\x8d\xce\x89\x0c\xe2\x90\x0a\xe3\xcf\x15\x33\xb4\x6d\x97\x76\x21\xdf\x18\xc5\xc4\xa4\xdd\x42\xd0\x2e\xaf\xd9\xff\x45\xaa\x10\x1b\x40\xa7\xbd\xde\x23\xaf\x77\xe2\xf5\x4e\x2f\x4e\x7e\xec\xf7\x7e\xe8\xf7\x7e\xf4\x7a\xff\xd7\xef\xf5\x50\x07\x50\xab\xd3\x19\x74\x1d\xff\x73\xd4\xa9\xda\x74\xa7\x6e\xdb\x16\xb7\xb9\xae\xcd\xf0\x65\xad\x40\xad\x5a\x3e\x7b\x8d\x6c\xdd\x7e\x75\xda\x1a\x74\xa3\xb4\x9a\x5b\x19\x35\x78\xc4\x69\xee\x93\xc9\x4d\xf2\xd7\xd3\x46\xb1\x88\x92\x04\x92\xdd\x38\x31\x9e\xa2\x3a\x92\x42\xb3\x19\x05\x21\xe7\x0a\x47\x08\xb4\x59\x58\x64\x9d\x33\x62\xa6\xfd\x93\x5e\xef\x61\x5d\x28\x9b\x29\xc5\xa4\xee\x99\xda\x82\xe9\x66\x9a\xad\x2f\x24\x9e\xc5\xb9\x53\x54\xc0\xd1\xd3\x0c\x87\x32\x18\x32\xd3\xad\xac\x8a\x99\x7f\x62\x1e\xef\x9c\x35\xe8\xd6\x2d\xcd\xce\xd9\xb2\x21\x9b\x25\xeb\x17\x92\xb8\xa0\x98\x50\x38\xb2\x75\xf7\x31\x1c\xcd\xec\x62\xa0\x3f\x04\xdf\xe5\x4e\x7d\x59\x00\xec\x36\x5f\xde\xaa\x39\x47\x40\xec\x96\x8f\xf2\xfa\xde\xd4\x2c\xb9\x3c\x61\x2b\x01\xe4\x60\x44\xff\x49\xf9\x22\x16\x46\x9c\x5a\x8f\x4f\x82\xc1\x77\x65\x00\xda\xb6\xf0\x15\x91\x91\xa2\x99\x1f\x45\x98\x10\x26\x26\xfd\x1e\x3a\x1f\x04\x92\xe4\xae\x79\xa5\xa5\x40\xe7\x2b\x01\x58\x13\x7f\x1b\xe1\xf7\x9f\x37\xaf\x5e\xfa\x3a\x89\x3e\x36\x5e\xb8\xdb\x08\x2b\x4d\xdb\x29\x12\xa4\xda\x5f\x2e\x3b\xc7\x20\x62\xce\x8f\xe1\xb4\x14\x6e\x83\xae\x5d\xc7\xf9\xa0\x1b\x29\xda\x48\x33\xbb\xf0\x67\x8d\x3c\x97\xde\x88\x77\x7d\x4d\x9a\x6b\x73\xab\x89\xeb\x1d\x7a\xb7\x84\x41\xb7\xc6\xad\x07\xdd\x04\x21\x6e\x59\x0e\xa5\x99\x4d\x48\xe2\x7c\xfd\x33\x2f\x80\x5e\xda\x75\x7e\x2d\x7a\xf6\x29\x7a\x3e\x8b\x24\x03\xcd\xe0\xb2\x94\x21\x9e\x3f\x6b\x92\x54\x60\x05\x0f\xfd\x48\x49\x23\x03\xc9\x2f\x43\x7c\x25\x15\xa0\x33\x0b\x83\x2b\x6c\x5f\xe2\xb0\x9c\x77\x1a\x46\x75\x99\xc3\xc5\x22\x6a\x94\xef\x36\x26\x3e\x55\x0a\x2f\xf6\x98\x99\x29\x9e\x30\x1d\x71\xbc\xe8\x83\x90\x82\x5a\xf7\xd8\x9a\x65\x0f\x9f\x33\x4b\xf9\xf2\x18\x8e\x2c\x52\x24\xb9\x32\x87\x8c\xfa\x2d\x34\x4d\x8f\x92\x50\xff\xf9\x33\xdf\x15\x69\x8d\x12\xe5\x6e\x9b\x97\x58\x5b\x9b\xaf\x31\x6f\x6a\xf7\x12\x17\x6b\xf7\x7d\x96\xb8\x32\x39\xb1\xfd\xbe\xb3\x9b\xe4\xa4\xa3\x48\xc9\x88\x2a\xc3\xa8\xb6\x56\x71\xd2\x5e\x17\x63\xcd\x32\x1b\x1b\xaf\x30\x6a\x30\x69\x37\xac\xa4\xf0\xb1\x7b\x17\xd0\x00\x3b\x36\xe9\x77\xf8\x56\x85\x80\x22\x12\x53\xfd\x2c\x10\x1c\x35\x0a\xc6\xad\xcc\xd2\x52\x76\x7f\x4e\xdb\xab\x81\x35\xca\xe6\xea\xd9\x55\xfe\xae\x7f\x8a\xf0\x8e\x8e\x73\x2f\x58\x24\xce\xb4\x9f\x4b\x14\x0b\xd8\xdb\x36\x2e\x52\x32\xd9\x55\x01\x7b\x37\x86\x89\x8d\x0a\x8e\xb0\xf2\xf0\x0f\xc1\xcc\x6d\xa5\x35\xb7\x21\x34\xae\x21\x4b\xbc\x9b\x59\xb2\xb2\x02\xbc\xbd\xfc\xdb\xd6\xb0\x9f\xa2\x7e\x2d\x64\x66\x95\xac\xa6\xc6\x30\x31\xb9\xff\x62\x16\xb2\xe2\x90\xdc\xbd\xac\xfd\x39\x3d\xb3\x7d\xad\x6b\x9b\xd7\xb5\x63\xa9\x42\x50\xd2\x56\x48\xf6\x12\x81\x6b\x55\x0f\xd1\x93\xf4\x22\xf3\x04\x04\x21\x35\x53\x49\x86\x28\x92\xda\x20\x60\x64\x88\x8a\x67\x2b\x4d\xff\x19\xe6\x8c\x60\x23\x55\x6d\x89\xfb\xb1\x8a\x69\x38\x50\x41\xbd\xd6\xbd\x39\x29\x75\x6f\x5e\x45\x56\x4d\xb7\x2a\x65\x1b\xb5\x6f\x60\x27\x3e\xee\xcc\x6b\xcd\x5b\x39\xec\x18\x8e\x64\xb2\xa1\xa4\x38\x2d\xa3\x40\x3d\xf3\x86\xf5\xa9\xe3\x9b\x64\xa5\xc3\xd4\x6e\x45\xf5\x9a\xb2\xb6\xb5\x25\xa0\x91\x94\xbc\x51\xdf\x66\x90\x7c\x6b\x9b\x36\x62\x82\x29\x0d\xde\xdb\xc0\x2f\xba\x35\x9e\x9e\x33\x13\x4c\x11\x08\x1c\xba\x2f\x5e\xd6\xf6\xe0\x82\xa0\x6a\x3c\x5f\x5a\xda\x76\xcd\x08\x12\x8b\x77\x00\x19\x65\x0d\xbf\x5c\x42\x22\x96\x92\x1c\x80\xa1\xbb\x73\xcf\x49\x6f\x86\x8d\x41\x2a\x68\x6f\x6c\x9e\x09\x83\x3a\xd5\xe3\x8f\x7e\x40\xb5\x3d\xd7\x5a\xb5\x88\x38\x1c\x51\x85\x56\xbe\x17\xb7\x68\xa2\xac\x8e\xeb\xf5\x92\xf4\x83\x56\x1e\xfd\x99\x76\x88\xea\x75\x16\x61\x63\xa8\x12\x43\xf4\xdf\xb7\xde\xf7\xef\x9e\xbc\xed\x79\xff\x7a\xf7\xdd\x11\xba\xa3\x42\xe2\x3a\x8d\xc4\x5f\x96\x4a\x6e\xa5\x91\x8d\x4d\x8f\xb9\xc4\x76\xd7\x5f\xc8\xa6\x0b\x3f\x68\xff\xed\xbb\x8b\xce\x93\x7d\x14\x80\x05\xa9\x30\xbd\x6b\xa4\x5a\xaf\x98\x50\x68\x73\x2a\xd6\xe3\xf3\xa4\xd7\x6b\xe6\x17\x86\x5e\x1b\xac\x28\xae\xd2\x0b\x28\xaa\xd9\x07\x9b\xc4\x2e\x33\x32\x04\x4a\xce\xf5\x10\xfd\x70\x0b\x4c\x39\xaf\x50\xe2\xa0\x9b\x71\x6e\xa6\x91\x7d\xad\x6e\xb9\x7f\x0c\x9b\x37\x30\xe8\xa7\x6c\x1a\xc3\xae\x63\x44\xed\xd1\x61\xd0\xb5\x5a\x3b\x54\x51\x2e\x0d\xbe\xa4\xc2\x4a\x22\x5f\x54\x5d\xfe\xea\xe2\xe9\xd7\xa2\xfc\xa0\xcd\xe6\xfd\x9b\xca\xdb\x2b\xc0\xfd\x9a\xca\xee\x3d\xb4\xe6\x25\x6f\x52\xff\x65\x0e\xac\x62\x21\x5c\x8f\x20\xe7\x97\x0e\x39\x86\x05\x56\x15\xcf\x29\x26\x8b\xe2\x69\x12\x14\x77\x89\xf6\xfd\x76\xfb\x5a\xc9\x89\xa2\x7a\xaf\xfd\xee\x86\xd9\xf2\x0b\x91\xb9\x80\xd4\x9c\x21\x56\x13\x26\xbc\x91\x34\x46\x86\xfd\xba\xb7\x29\x9a\xb0\xf4\x46\x58\x41\xf9\x26\xf7\xa0\x4a\x8b\xac\x50\x32\x31\x96\xc9\x69\x70\x46\x0b\x9b\xac\x50\xb8\x97\xc2\x9c\x49\x2c\xba\xbb\x63\x64\x46\x32\xc2\x0a\x01\x56\x0c\x7b\x49\x46\x10\x72\x9e\x60\x7f\x22\x72\xae\x98\x31\x54\x24\x39\xa1\x20\x09\x99\x18\xa2\xde\xca\x08\xbe\x2e\x26\x19\x69\x30\x4f\xa6\xac\xb8\x7d\xf6\x38\x93\x0b\xcb\xe5\xb6\xe3\xe0\x86\xda\x74\x84\x45\xc6\x91\x33\x41\xbd\x29\xb5\xd0\xd3\x3f\xed\x45\x75\x88\x53\xf7\xc9\x7d\xe6\x21\x81\xd1\xc2\x50\x0d\x72\x0c\x0f\xc9\xfa\xbb\x18\x2b\x2a\x58\xd9\xda\x1e\x1d\x33\xbb\xec\xa6\x7d\xcf\x0d\x24\xba\x25\xd9\xc7\x0b\xba\x9f\xed\xb9\x48\xc7\xe1\x9e\x20\x93\x28\x33\x48\xe7\xde\x01\x25\x76\xb7\xd9\x36\x9e\xe4\x47\xbe\x95\x98\x1a\xe3\x9a\xba\xcb\x35\x5e\xca\xa5\xd5\x54\x2a\xf6\x41\x0a\xeb\x08\xc9\x7d\xf2\x7a\x9e\xc7\xe9\xd8\x00\x51\x32\xfa\x20\x05\x45\x2b\x9d\x9a\xcd\x7e\x8c\x34\x78\xb3\x81\x93\x0c\x0a\x99\x36\x65\x6a\xfa\x9b\x2b\x2f\x33\x1b\x1a\x42\xf1\x06\xf4\xb6\xd6\x8a\x7b\xb1\x79\x2c\xd5\x10\x71\x3a\xa3\x1c\x15\x59\x38\xa9\x13\xbd\xf4\xcd\x67\x57\x87\x9c\x65\x65\xc8\x59\x65\x15\x52\xa5\x51\xf7\x5d\x21\x0b\xa9\x8c\xd3\xd7\x61\xb7\x58\xb3\xe2\x85\xc6\xba\x3d\xa6\x2b\x7a\x94\xad\xe8\x51\xa3\x15\xc1\x1e\x05\xb2\xb5\x87\xc9\x16\xee\xaa\xe5\xfc\xb6\x28\x91\x5d\xf4\xbb\xf1\x04\xda\x14\xfd\x27\x66\x8a\x92\x21\xca\xae\xb6\xe9\xbf\x3e\x64\x6b\x1e\x55\x15\xa4\x50\x5b\x00\xdf\xa2\x54\xfd\x62\x8a\xd2\x9f\x94\xc4\x24\xc0\xda\x40\x48\xb5\xc6\x13\xfa\xb5\x44\xdd\xbb\x6f\xdc\x0c\xbe\x76\xa1\xd6\x28\xb3\xc4\x26\x76\x95\x1e\x7d\x89\x08\x06\x65\x14\x7b\xe1\xd8\x5b\x27\x4b\x8b\x8e\x2c\x4c\xf2\x40\xff\x6e\x57\x5a\xff\x0c\x31\x2e\xd5\x9a\x43\xb8\xf4\x26\xe2\x38\xa0\x53\xc9\x09\x55\x43\x84\x39\x55\x05\xe6\xdd\x3b\xc2\xad\x6b\xe1\x76\xbe\x90\x63\xc2\x7d\x79\xc3\x6f\x65\xd0\xf9\xe4\xf6\xbe\x6b\x1f\x2b\x57\x97\xf5\x88\xec\xe6\xbc\x49\x7b\xea\xee\x06\xe6\xe2\x52\x4b\xce\x48\x25\xc4\x55\x4d\x68\xe6\x0c\x75\x8a\x0d\x89\x27\xc7\x63\x4d\x8d\x77\xb6\x4b\xab\x2b\xbf\xb0\xd1\xf1\x28\x64\x66\xe3\x17\x36\x3a\x0e\x02\x7b\xf6\x2b\x9d\xae\xa9\x20\x59\x2e\xda\xf6\xf3\x9e\xdb\x69\xef\x70\x1d\xa9\xec\x77\x7e\x53\x8a\x49\xde\x50\xbe\xb9\x01\x6d\xb0\x61\xc1\xaf\x17\xbf\xbd\x80\xb6\xbb\xfe\xe3\xf7\x17\x80\xba\x04\xeb\xe9\x48\x62\x45\xba\x58\x6b\x6a\x74\x77\x46\x05\x91\x4a\x77\xf3\x2f\xfb\xb4\x2f\xa8\xf1\x46\xba\x1b\x68\x37\x7a\xe1\x46\x47\x52\x1a\x6d\x14\x8e\xfc\x90\x09\x3f\xb0\x27\xe5\xa4\xae\xee\x1c\x50\x6a\xf1\x25\x63\xb6\x80\x62\xe4\x7e\x16\x30\x65\x93\x29\xb7\x15\xc4\x95\x93\x67\x64\x28\x95\x92\xf3\x83\x6e\x32\x3d\x38\xa4\x22\xb2\xdb\x2a\x11\xce\xde\xc9\x39\xb8\x04\xfc\x76\x5a\xc9\xcd\xbf\xcd\xbf\x11\x4c\xce\x47\x90\xe4\x89\xe3\x8d\xe1\x1c\x4e\x6e\x56\xfc\xcc\x9d\xd7\xe1\xa4\xd7\x7b\xf8\x38\x7f\xb0\x4c\x1d\x2d\x91\x5c\xe7\x63\x57\xfa\x80\x1e\xd6\xbd\xd2\xdd\xab\x7f\x62\xaa\x16\x7e\xc9\xc9\xac\x4a\xae\xee\xc3\xb3\x46\xda\x0a\xac\x75\xe7\x7b\x91\x59\xf8\xee\x9a\xec\x92\x53\x7f\x04\xe1\xe9\xde\x6b\x23\xe9\x30\xe2\xcb\x81\x74\x55\xba\xf5\x23\x1c\xbc\x3f\xa0\x9c\xfc\xb5\x06\x2b\x25\xbf\x39\xa4\x1e\x8b\x68\xbd\x5a\x0b\xd6\x4d\x01\xcd\x7e\x47\x90\x87\xd9\x51\x3b\xfb\x49\x41\xc7\x4f\xfa\xa9\xed\x71\x2c\x92\x22\x1b\xda\x9d\xb5\x38\x7d\x96\x49\x76\x5f\xd3\x68\x7b\x48\x85\xe1\x1a\x91\xfd\x44\x58\x61\xce\x29\xff\x23\xe2\x12\x13\xdd\x87\x93\xe3\x4a\x9a\xf0\x25\x0e\x69\x1f\xd0\x98\xa9\x70\x8e\x15\x45\x9b\x64\x81\xa2\xd8\xd0\xe7\x21\x9e\xd0\x8b\x69\x1c\x8e\x04\x66\x5c\xf7\xdd\x9e\x37\xa9\x71\x10\xd0\xc8\x50\xf2\x0b\xe3\x54\xf7\x61\x65\xfb\x21\x0e\x46\x4c\x60\xb5\x38\xf6\x47\x4c\xac\xfe\x58\x73\xf9\x78\xb5\x71\x33\xe5\x57\xda\x67\x82\x99\x5f\x33\xa7\x61\x62\xf2\x4a\xbc\x90\x98\xb4\x3b\x6b\xb4\xb9\xbe\xd2\x94\xfd\x2b\x16\x84\x53\xd5\x56\xeb\xea\xb3\x1f\x36\x86\xb6\xb2\xd1\x16\x73\x03\xc3\xe1\x10\x5a\x63\xcc\x38\x25\xad\x2a\x62\xfb\x11\x74\x0e\xaf\x5f\x4a\xc3\xc6\x8b\x76\x35\x85\xfd\x24\x27\xdb\x3e\xb4\xfe\xad\x94\x54\xad\x4d\xbd\xe4\x74\xf4\xda\xf4\x41\xf9\x69\x29\xb6\x85\x70\x11\x59\x7e\x74\x07\xbf\x29\x23\xb4\xd6\x18\xd9\xc7\x22\x38\x13\x93\x3e\xb4\xf2\x28\x3f\x6b\x55\x52\x2f\x3b\x8f\x37\xc6\x97\xf9\x97\xb6\xed\x7c\xdd\xf0\x8d\xd5\x5c\x2c\x5c\x1a\x38\x94\xf2\xde\x38\xfb\x1d\x50\x7d\x7a\x27\xc7\x8f\xa1\xc0\x55\x4f\x7f\xb0\x96\x7a\x05\x91\x73\xdf\x1d\xab\x61\x98\x3b\x73\x3b\x08\x49\x95\x56\x8f\x7c\x7c\x85\xaf\x6b\x94\x99\xee\xfa\xf5\xab\x37\x17\x35\x5b\x8e\x15\xef\x43\x2b\x3b\xc5\xb7\xe0\x7b\x08\x42\x52\x4d\x9a\xea\xae\xbf\x16\x56\x9b\xdb\x5b\xdb\xf2\x7a\x28\x1f\xb5\x5b\xdf\xe6\xed\x82\x56\xc7\x77\xf5\x76\x8e\x72\x6d\x3a\xb3\xd8\xb7\xb9\xa1\x64\xdc\x8f\x54\xf2\xff\x33\x3a\xc6\x31\x37\xed\x0a\xf5\xce\xb0\x02\x0a\x43\x38\x6a\x9b\x29\xd3\xeb\xe0\x70\x28\x9d\x51\x3f\x52\x32\x6a\xb7\x9c\xe6\x5a\x9d\x6a\x52\x9b\x70\x2d\xad\xa6\x8a\x61\xce\x3e\xd0\x76\x0d\xe1\xad\x95\xbb\xbe\xbf\x8a\xd7\x42\x57\xb6\xde\x6e\x15\x35\x60\x52\x14\xbe\x65\xe4\xdd\x4a\x61\x98\x96\x84\x76\xbc\xd5\xf1\x83\x29\x16\x13\x5a\x58\xa7\xca\x07\x67\x78\x73\xa1\xf6\x53\xb2\x42\xbd\x7a\x6c\xbe\x5a\x56\x3f\x7e\x0f\xc3\x5c\xcf\x8c\xb4\x3a\x1b\x44\x15\xb6\xb5\xb0\x94\xcd\xb1\xa6\x6c\x75\x1c\xaa\x67\xaf\x8d\xd5\x42\x93\x5d\xcb\xdb\xf7\xef\x4a\x22\xd3\x57\xbe\x5a\x1d\xf8\x66\x08\xad\x56\x2d\x12\xee\xe6\x37\xc3\xbc\xbd\xb9\xfa\xe5\xfd\x78\x66\xd9\xc4\xad\xce\x1e\x6e\x6a\xff\x1e\xde\x3b\xcb\xf7\xa5\x33\xc3\xda\xb0\xf5\xd9\x24\xb7\x55\x35\x59\x76\x67\x8d\x26\xe9\xd6\xe5\x8a\xd6\xcd\x4d\x21\xa9\x8e\x72\x77\xae\xdd\x99\x26\x9a\xa5\x88\x7a\x05\xad\x53\xe4\x3f\x11\x2d\x0e\x5e\xff\x0b\x00\x00\xff\xff\xbb\x6d\x79\x96\xc3\x47\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x5b\x6c\x54\x55\x17\xfe\x3a\x67\x3a\xd3\x7f\x7e\x2f\x88\x88\x78\x65\xfb\x50\x84\x90\xa9\x33\x55\x89\x19\x3a\x2d\x85\x42\x24\x52\x24\x65\xc0\x07\x1f\xcc\xa1\xb3\x19\x4e\x98\x39\x67\x72\x2e\x68\x13\x62\xda\xa9\x51\x08\x24\x5c\x12\xd4\x78\x81\x00\x26\x9a\x68\xcc\xf4\x32\x30\x40\x7b\x1a\x7d\xc2\xc4\x87\xb5\x1f\x34\x1a\x12\x1e\x7c\x30\xbe\xfb\x60\xe2\x93\xd9\x67\xef\x99\x22\x6d\x13\x27\x99\xac\xbd\xd7\x5e\x97\x6f\x7d\xeb\x3b\x77\x57\xc7\x2f\x00\x00\x03\xf0\xb4\xfc\x77\x00\x9b\xa5\x8d\x21\xfa\xdd\x8e\x01\x0f\x01\xf8\x21\x06\x74\x02\xf8\x25\x06\xfc\x0f\xc0\xdd\x18\xf0\x08\x80\x3f\x62\x40\x17\x80\xbf\x63\x40\x02\x80\x61\x00\x6b\x00\xa4\x0c\x60\x2d\x80\xf5\x86\x8a\xdf\x62\xa8\xf8\x6d\x06\x20\x4b\x8f\x18\xc0\xa3\x00\x0e\x18\xc0\x3a\x00\x96\xa1\xea\xd7\x0c\x20\x09\xe0\x84\x01\xc4\x01\x9c\xd5\xfe\x0b\x86\xaa\xff\xa9\xa1\xfa\x5d\xd1\xf6\x6b\x1d\x3f\xab\xe3\xe7\x74\xdc\x77\xda\x7f\x5b\xfb\x49\xd7\xf9\xc9\x00\xba\x01\xdc\x31\x80\x4d\x00\x3a\xe2\xea\x7d\x73\x5c\xbd\x67\xe2\x2a\x6f\x4b\x1c\x18\xe8\x00\xfa\xe3\x0a\xe7\xbb\x9d\xc0\x03\x00\xce\x77\x02\xab\x00\x5c\xea\x04\x5e\x06\x50\xd7\xf7\xdf\x3b\xd5\x1c\x7f\x75\x02\x7d\x00\x56\x25\x80\x57\x01\x0c\x27\x80\xc7\x01\x7c\x98\x00\x9e\x07\xf0\x6d\x42\xcd\xff\x73\x42\xf5\xff\x35\x01\xf4\x48\xde\x92\x8a\xe7\xde\x24\xb0\x1a\xc0\x60\x52\xcd\x51\x48\xaa\x39\xdf\x4c\x02\x29\x00\xc5\xa4\xea\xe7\x68\x3b\x9e\x04\x9e\x04\x70\x4e\xe7\x7d\x93\x54\x38\x17\x74\xbd\x1f\x75\x9d\x3b\x3a\xfe\xb7\x24\xb0\x0f\xc0\x9f\x49\xe0\x2d\x00\x43\x5d\xca\xff\x45\x97\xca\x6b\xe8\xfb\xf7\x5d\x40\x87\x92\x40\x34\x3f\xb4\x3e\x9e\xd5\xe7\xb5\xda\xae\xd3\xb6\x4b\xdb\x55\xda\xca\x5d\x4b\x5c\x12\xf3\x53\x50\xd8\xb4\xa4\xf0\x30\x80\x67\xb0\xf8\x93\xf8\xfe\x0f\x85\x57\x72\xb5\x46\xeb\x51\xe2\x59\x0f\xa5\x1f\x43\xc7\x3e\x06\xb5\x2f\xb9\x2b\xb9\xa7\x07\x01\x3c\x21\x1f\xba\x8b\xec\xd0\x98\xcf\x3d\xe6\x1c\x66\xdd\x45\x0c\xba\xae\x39\x86\x41\xdf\x77\xad\x43\x81\xcf\xb1\xdd\x75\xcc\xe2\xa8\xe9\xf9\xac\xc2\x3d\xcf\x2c\x71\xec\x38\xc2\x47\x8f\x7a\x41\x05\x3b\x1c\xfb\xb0\x55\x52\xc6\xad\x30\x97\x7b\xdc\x67\x45\x7e\xcc\x1a\xe5\xf7\x3a\x7d\xd3\x6d\xbb\x87\x22\xc3\xba\xbd\xd6\xc9\x6c\x35\xf2\xb0\x7b\x08\x7b\x64\x1f\xc7\x2e\x5b\x36\x67\xbe\x55\x89\x02\x97\xf1\xb1\x8d\xdd\xde\x26\xec\xe1\xc7\x78\x19\xc3\x1a\xd5\x5e\xb3\xc2\xb1\xd7\x29\x72\x0f\xaf\x57\x7d\xcb\xb1\xb1\xcf\x75\x4a\x2e\xf7\x3c\x79\xa8\x72\xd7\x1f\xc3\x88\x02\x83\xfd\xdc\x2e\x62\xbf\x6f\xfa\x81\x87\x82\x55\xe1\x4e\xe0\xa3\x30\x56\xe5\x38\x68\x96\x03\x8e\x37\xac\x72\x99\x1d\xe2\xf7\x61\xef\x61\x83\x2e\x67\x63\x4e\xc0\xbc\xc0\xe5\x03\xed\xa8\xb7\xad\x2a\x67\xa3\x11\x15\x2b\x46\x56\x39\x5c\x6e\x16\xc7\xe0\x06\xb6\x6d\xd9\x25\x8c\xf0\xaa\xe3\xfa\xe9\x61\xaf\x64\x15\xd3\xdb\x83\x92\x97\x2e\x38\x39\x99\xbe\xed\xa8\x75\xc4\xac\x38\x3d\x6e\x90\x92\x83\xa7\x0b\xae\x69\x7b\x65\xd3\x77\xdc\x1c\x7b\x2d\x7a\x62\xc3\x81\x6b\x56\x9c\xa2\xc3\xfa\xfe\x15\xdf\x9f\xda\x63\xda\xa5\xc0\x2c\xf1\x74\x81\x9b\x95\x1c\x6b\xdf\x73\x6c\x24\xf0\x3c\xcb\xb4\x53\xc3\xbb\x87\x77\xa6\x0f\x72\xd7\xb3\x1c\x3b\xc7\xb2\x3d\x99\xd4\x0e\xc7\xf6\xb9\xed\xa7\xe5\xf8\x39\xe6\xf3\x77\xfc\x17\xaa\x65\xd3\xb2\xb7\xb2\xd1\x23\xa6\xeb\x71\x3f\x7f\xa0\xb0\x2b\xfd\xca\x62\x9c\xc4\x73\x98\xbb\xe9\x9d\xf6\xa8\x53\xb4\xec\x52\x8e\xa5\xf6\x95\x03\xd7\x2c\xa7\x77\x39\x6e\xc5\xcb\x31\xbb\x1a\x5d\xbd\xfc\x8b\x5b\x99\x3a\xe6\xed\xee\x6c\x26\x9f\xcf\xb2\x0d\x1b\x98\x3c\x66\x9e\xcb\x67\xb3\x6c\x80\x65\x58\x2e\xba\xf7\xe7\x7b\x5b\x4f\x7d\xf9\x97\xe4\x71\x63\x14\xd6\x97\xcd\xb0\xe3\xc7\x55\x4a\x7f\xbe\x37\xb3\x89\x0d\xb0\x2c\xcb\xb1\xde\xad\x52\xb4\x34\x45\x75\xba\x21\x6a\x8c\x9a\x74\x5d\x2a\x97\x3e\xa3\xba\x98\x10\x13\xd4\xa4\x69\xd0\x59\x51\x13\x35\x31\x4e\x4d\x9a\x12\x93\xa2\x06\xfa\x8a\x9a\x62\x9c\x42\xba\x49\x21\x4d\x53\x43\x9c\xa4\xba\xa8\x51\x83\x6e\x89\xd3\x34\x4f\x21\x35\x98\x98\xa0\x90\x42\x9a\x12\x27\xa9\x41\xf3\xd4\xa4\x06\xe8\x4b\x6a\xd0\x4d\x31\x21\x26\x69\x8e\xe6\xa8\x0e\xfa\x84\x42\x9a\x17\xef\x51\x93\x66\xc4\xa4\x18\xa7\xba\x78\x9f\x9a\xe2\x0c\xe8\x12\x85\x34\x2b\x6a\x51\xed\x71\xba\x46\xb3\xad\x22\xb2\xee\x94\x6c\x2d\x26\xa8\xce\xc4\xa4\x98\x88\x80\x85\x74\x43\x9e\x68\x5a\x56\x5d\x39\x99\x16\x22\x57\x83\xae\x53\x9d\x66\xc4\xb8\x98\xa4\xeb\x74\x93\x9a\x2b\x14\xba\xba\xc4\x19\xca\x8f\xe8\x7e\x3a\xc4\xa9\xe5\xf3\x77\x0f\x81\xce\x47\xed\xe6\xc4\x19\xd9\x5b\x82\xbe\x45\x0d\x9a\xa5\x79\x6a\xd0\x0c\x85\xad\x4f\x51\x16\xfd\xaf\x91\xea\x7b\xa5\xab\x51\xb3\x69\x39\x9a\x38\x0d\xba\xbc\x94\xed\x8f\x65\x31\xd0\xe7\x11\x1b\xa7\x40\x17\x69\xa1\x4d\x6f\x94\x3c\x23\x3b\xca\x25\xcb\xf4\xe9\x7b\xa7\x94\x1c\x2e\x25\xaa\x29\x6a\xb2\xd5\x45\x51\xa3\x05\xb9\x2c\x9a\x6e\xb9\x2e\x8b\x9a\x14\x80\x64\x01\x74\x45\x4a\x49\xee\x57\x49\xe5\x0a\x35\x69\x01\xf4\x11\xcd\x53\x5d\x7c\xb0\x08\x6f\x39\x72\x25\x9f\x72\x5d\x52\x88\x4b\x57\x75\x2d\x4a\x0e\x7b\x18\x9d\x8f\x28\xd7\x0b\x96\x0c\x9c\x1a\x00\x9d\x6b\xe7\xb6\x35\x72\x22\xca\xa8\xb3\x48\xa7\xcb\x08\x6d\xd9\xbd\xad\x54\xff\x72\x5b\x78\x7a\xea\x19\x0a\x45\x2d\x5a\x42\x88\x08\xe6\x82\x98\xd4\x0b\x08\xf1\x4f\x00\x00\x00\xff\xff\xed\x76\xe8\xc1\xd3\x08\x00\x00"

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
