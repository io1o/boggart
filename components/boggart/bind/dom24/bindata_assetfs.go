// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// DO NOT EDIT!

package dom24

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdf\x6e\xdb\x36\x17\xbf\xcf\x53\x1c\xa8\xf9\x3e\x34\x17\xb6\x9a\x7c\xdf\x45\x91\xca\x1e\x8a\x16\x01\x8c\x75\x5e\x90\xb5\xbb\x0d\x28\x89\x8e\xb9\xd2\xa4\x4a\x52\x8e\x0d\xc3\x37\xbb\xd9\x83\xec\x19\x76\xb9\x3d\x43\xfa\x46\x83\x48\xea\x8f\x53\x5b\xf1\x1f\xca\x88\x37\x0b\x68\x2b\x8b\x3c\xbf\x73\xc8\x73\x0e\x49\xfd\x8e\x3a\x9b\x41\x8c\x07\x84\x61\xf0\x22\xce\x14\x66\xca\x83\xf9\xfc\x24\x90\x91\x20\x89\x02\x35\x4d\x70\xc7\x43\x49\x42\x49\x84\x14\xe1\xcc\xff\x05\x8d\x91\x69\xf4\xba\x27\x00\x00\x83\x94\x45\x59\x0b\x8c\x11\x4d\xf1\x15\x17\x23\xa4\x5e\xb2\x33\x98\xe9\xd6\xec\x12\x58\xa5\x82\x01\xc3\xf7\xd0\x63\x8a\xb6\xfb\xe9\x28\xc4\xc2\xf6\x3c\x6b\x0f\x72\x91\x37\x5a\x62\x7e\x12\xf8\x46\x41\xf7\x24\x88\xc9\x18\x22\x8a\xa4\xec\x78\x82\xdf\x5b\x8d\xd5\xa7\x11\xa7\xad\x51\xdc\x3a\xbf\x80\xec\x4e\x8e\xf2\xbb\x89\x6c\x9d\x5f\xd8\xfe\x8f\x65\x26\xb7\x09\x62\x98\x56\x5a\xbf\xed\x91\x4f\xc6\x62\x9f\xa2\x9f\xe0\x14\x77\x3c\x85\xc2\x65\x48\x45\xcf\x94\xe6\x80\x0c\x8d\x81\xa1\x71\x4b\xa1\x50\x42\x88\xc4\x6d\x76\xe3\x95\x30\x94\xc8\x65\xba\x0a\x24\x4a\x6c\xdf\x44\x60\x89\x99\xd2\xbe\xf0\x66\x33\x20\x03\xc0\x5f\xa0\x8d\x8c\x0b\xbc\xcc\x79\xb9\xce\xec\xd9\x18\x67\x9d\x30\x8b\x61\x3e\xef\x06\x08\x86\x02\x0f\x3a\x2b\xe4\x4a\xcf\x5e\x8e\x39\x89\x5f\xbe\x3a\x7b\x93\xc9\x52\x89\x61\x3e\x9f\xcd\xa0\x7d\x83\xbf\xa4\x58\xaa\xf6\xa7\x9b\x0f\xed\x6b\xa4\x86\xe6\xb1\x01\xf7\xba\x19\xe8\xf9\x6b\x06\xde\xdb\x28\xe2\x29\x53\xd2\x83\x36\xcc\xe7\x81\x8f\xba\x81\x4f\x89\x83\xc1\x8d\xb0\xc2\x42\x6e\x3b\xc4\x52\x7a\x9b\x81\x7e\x67\x60\x3a\x06\x65\xd9\xb0\x7f\xb0\xf8\x4f\x0f\x3a\xf0\x53\xba\xa2\xa5\x12\x83\x0a\x85\xad\xd5\x51\x58\x17\x8d\x55\x84\xec\x09\x0c\x50\x8c\xc1\xcc\x15\x10\x56\x83\x96\x5d\x75\x33\x57\x2b\x58\x0a\xb7\x8d\xc4\x3a\x02\x56\x48\x20\x76\x87\xe1\x54\xcb\xc1\x65\x67\x53\x84\xe5\xeb\xc4\xa6\x92\xeb\xae\x25\x9b\xe2\x2e\x5f\x6f\x36\x47\x51\x44\x51\xbc\x05\x8a\x46\x1a\x5e\x74\x5f\xcc\x66\x76\x8a\xdb\xbd\x18\x33\x95\x25\x52\xf9\xa8\x8f\x46\x58\x87\xee\xf0\x62\x4b\x15\xdf\x2c\x77\x21\x12\x2d\x41\xee\x86\x0a\xf4\x04\xdc\x2a\xce\x69\xc8\x27\x5b\x0e\x01\xcc\x5a\x91\xe5\x78\xe9\x30\x8a\x12\x89\x5b\x94\xb0\xcf\x5e\x37\x20\x79\xc3\x00\x49\x18\xa0\x56\x34\xc4\x63\xc1\x59\x2b\x4d\xbc\x6e\xe0\x93\xee\xd3\x6b\x51\xad\xee\x95\x69\xfb\xa4\x64\x35\xc8\x28\x46\x62\x40\x26\x99\x45\x31\x19\x6f\x11\x13\x5b\x8a\xad\xb9\xbd\xad\x85\x25\xb1\x59\x1a\x0a\x3f\x68\x34\x20\x6c\xcc\x49\xb4\x6d\x88\x3e\xb6\x72\xb3\x44\x7e\x0a\xcd\xda\xd6\x1a\x62\x14\x63\xb1\x23\x30\x98\x8c\xca\xb1\x93\x94\x52\x13\xe8\x7a\x3f\xb0\x09\xf5\x36\x8e\x05\x96\x72\xa7\x9c\x2a\x94\x6d\xe7\x74\x57\xe2\x8b\x5e\xc9\xfd\xdc\x22\x6c\xc0\x1d\xba\xc8\xae\xba\xff\x2f\xf0\x23\xbe\xcd\xa2\xf9\xf8\x5a\xdc\xa1\x3d\x38\x5d\x77\x57\xa9\xb5\x1b\x19\xef\xee\x6e\x9e\x46\x93\x4a\x70\x76\x57\x1e\x26\x6e\xb0\xe4\xa9\x88\xb0\xb1\xf6\x32\xf0\x6d\x87\xca\x7a\x9d\x77\xc9\xe2\x2b\x14\xe0\x37\x64\xc9\x15\x8a\x14\x17\x53\x60\xfa\xd0\x5e\x63\x8f\xed\x68\x0e\xf7\x4e\xa6\xd8\x77\x32\xc7\x3b\xc6\x3e\xec\x39\x48\xdf\x0d\x71\xf4\x39\x4d\x0e\x23\x4c\x3f\x20\xa9\x6a\x42\x22\x6b\xb6\xe3\x79\x8f\x54\xd3\x91\xda\xc7\x93\x3a\x63\xb2\xe6\xfd\x19\xd3\x63\x0a\x8b\x31\xa2\xcb\x0c\xca\x7a\x5c\xd3\x54\x20\x0a\xde\x7f\x62\x98\x62\x24\xbc\xe2\x4e\x7a\x65\x82\x47\x99\xb9\x39\x12\x9c\x02\x23\xb4\xf8\xb3\xa2\xd3\x31\xef\x36\xbb\x0a\x87\xfd\xa4\x90\x50\x87\x91\x75\x59\xf8\xd6\x04\xba\x1e\xc9\x1e\x42\xfc\x67\x44\xd3\x27\xed\xd0\x9d\x16\x5f\x34\x3e\x31\xa2\xd6\x7e\xb3\xab\xb5\xe8\x59\x04\xea\x8e\xe2\xe6\x8d\xd9\x4e\x8d\x9e\xad\x9d\xe6\xa6\xb9\x03\xb4\x42\xe1\xd6\x6f\x9f\x0b\x98\x1a\x68\x01\x15\xf4\xdf\x2d\xa9\x04\x49\x70\x0c\x8c\xdf\x0b\x94\x38\x50\x65\xd4\x65\x27\x7e\x37\x58\x06\x4f\xb8\x03\x33\x80\xc3\x65\xa9\x1d\xf8\x6a\xd8\xa0\xa2\x4a\xf2\xba\xd5\x14\xf8\xae\x26\x28\xb3\xcb\x99\xeb\x02\x15\xf2\x78\xea\x06\xab\xe4\xab\x34\xe5\x0d\x97\x1d\x77\xf9\x5b\xbd\x9c\x86\x5a\xa0\x62\xfd\x6e\xaa\x4d\x6e\x5f\x63\x41\x78\x6c\x9c\xef\x34\x39\x62\x40\x94\xdc\xb1\x8e\x67\x5e\x87\x8b\x34\xc7\x13\xd5\x2a\x98\x45\x6b\xc4\xf7\x84\xc5\xe0\x3d\xfc\xfe\xf5\xb7\xaf\xbf\x3e\xfc\xf1\xf0\xe7\xc3\x5f\x1e\xcc\xe7\x32\x8d\x22\x2c\x65\xc9\xc8\xc6\xd9\x5c\x8b\x92\x72\x05\xcd\x83\x69\x8a\xb7\x0a\x34\x9f\x3b\x5a\x31\x8a\xc1\xac\x55\x03\x89\x79\x94\x8e\x30\x53\xed\x7b\x41\x14\x7e\x59\x2d\x82\x94\x06\xe6\xbb\xe0\xd9\x59\x51\xdb\x68\x62\x43\x2c\x2c\x77\xe6\x54\x77\xd9\x5c\xf8\xcf\x95\x5d\x6e\xf2\x39\xf0\xf5\xde\x73\xe0\x27\x88\x1d\x66\x36\xf0\x2d\x99\xb7\x17\x26\x72\x43\x91\x0d\xba\x6f\xd0\x75\xfd\x09\x5b\xaf\x67\xb9\x54\xad\x5b\x1f\x41\xb6\x44\xb6\x79\x85\xc4\x4a\xea\x1a\xc9\xa6\x28\xc7\x2a\xc9\x93\x48\xc3\x0b\xbd\x49\xda\x99\x6d\x7f\x9c\x26\xfa\xe5\xe5\x45\xf5\x61\xef\xfd\xb1\x4c\x52\xa7\xfb\x58\x26\xc9\xb1\x8e\x65\x92\xda\x32\x49\x9e\x50\xc7\x42\xc9\x86\xd0\x4d\x73\x61\x57\x82\x8f\x9e\x3b\x15\x56\x44\xcf\x55\xef\x47\x1d\x39\xb6\xe9\x1f\xc2\x2a\xc1\x9e\x5d\xfe\x91\x1f\x8c\xc3\xdf\xf1\x51\x82\xd8\xb4\xea\x74\x87\x34\x67\xc9\xe0\xf7\xfb\x96\xdc\x84\x25\xda\x7b\xfd\xfe\x91\xc2\xb4\x97\xa5\x30\x8b\x43\x93\x40\x4c\x9a\x6f\x87\x8e\x54\xe6\x73\xa5\x32\x1d\x73\x4b\x7b\xa0\x30\x17\x94\x18\xfe\xaa\x71\x35\xef\x71\x48\x54\xe3\x5a\xde\x09\x1c\xef\x41\xcd\x15\xa1\xce\x1d\xf3\xef\x60\x7b\x4f\xe3\x2c\x0e\xb2\x77\xee\x57\xae\x08\xac\x0c\x35\xd2\x7e\x77\x0d\x6b\x89\x02\x55\xae\xc2\x9a\xa0\x76\xbd\x3a\x57\xaf\x46\x88\xea\xca\x00\xda\x79\x35\xd3\x35\x59\xfd\x8d\xa2\x06\x79\xf1\x2a\x13\x6e\x39\x6e\x6f\x91\x2c\x77\x4b\x5f\xdb\x5d\xb9\x3a\xba\x9e\xbc\x46\xd3\x91\xf9\xf2\xd3\xa9\x2e\x70\x45\x97\x57\xad\x7d\x3b\xd2\x4c\xd7\x22\x6b\xfe\xdf\x17\xaf\xff\xf7\xfa\xfc\x8d\x73\xeb\xcb\x2c\xef\x00\x8a\xe3\xfc\xc7\x72\x83\x5c\xfb\xc9\x21\x31\x0e\x4e\xe9\x7f\x58\x12\xbb\xa6\x1c\xb3\x87\xd0\x65\x5c\x1d\xc3\x77\xcd\xab\xb2\x9d\xd8\xf8\xb5\xbf\x8e\x01\xac\xd7\xf9\xa6\x57\xd6\xec\x6c\xf5\xe9\xe6\x43\x23\x81\x59\xf9\x2f\x2d\x2b\x74\x16\x65\xd6\x50\x31\x08\x15\xd3\x7c\x93\xb9\x89\xb8\x79\x32\x91\x1e\x28\x24\xee\xb0\xea\x78\xb7\x21\x45\xec\xb3\xe3\x9c\x2d\xcc\x7d\xcc\x2e\xc7\xfc\x9e\x51\x8e\xe2\x6a\xdd\xd6\x9e\xae\x8b\x96\x53\x5d\xbd\x0d\xfc\x2d\x29\xe7\x5a\x7b\x7c\xe4\xdc\xfb\xcf\x36\xdc\x9f\x6d\xb9\xd6\xf1\x01\x11\x22\x4e\x65\x82\x58\xc7\xbb\x78\xbc\x11\x55\x18\x2d\x95\x7f\x7e\x79\xf8\x87\x39\x27\xbb\x93\x39\xd2\x34\xbb\x21\x1d\xfc\xd9\xc3\xc9\x4c\xdb\xcd\xf7\xc0\xa6\xba\xfb\x1c\x17\xa2\xe3\x97\x1e\xe5\x75\xfc\xd2\x63\xc3\xae\x4d\x7c\xe9\x51\xdf\xab\xc6\xb8\x15\x4d\x4b\x1e\x3f\x7a\x54\xf9\x69\x6f\xed\x3f\xa5\x39\x7f\x07\x00\x00\xff\xff\x96\xfa\xbf\x26\xc7\x3f\x00\x00"

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
	"templates/views/widget.html": templatesViewsWidgetHtml,
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
