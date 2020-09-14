// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// DO NOT EDIT!

package openweathermap

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x57\x5d\x6f\xe2\x3a\x13\xbe\xdf\x5f\x31\xf2\xcb\x05\x48\x4b\x80\x04\x02\x2f\x07\x90\xaa\xd2\xdd\xed\x11\x6d\x8f\xda\xae\xaa\x5e\x21\x37\x31\xe0\x3d\xc1\x41\xb6\xf9\x52\x94\xff\x7e\x64\x12\xbe\xc2\x10\xe8\x6a\x7b\x43\x3a\x19\xcf\x3c\xf3\xcc\xf8\xb1\x13\x45\xe0\xb3\x11\x17\x0c\x88\x17\x0a\xcd\x84\x26\x10\xc7\x5f\xa2\x08\xf8\x08\xa8\xf0\xc1\xf2\x29\x0f\xd6\x60\x79\x73\x29\x99\xd0\xe6\x65\xc7\xe7\x0b\xf0\x02\xaa\x54\x97\xc8\x70\x49\x7a\x5f\x00\x00\x0e\xad\x5e\x18\x94\xa7\x7e\xb9\x66\x83\x79\x52\xd3\xed\xd3\x4a\x95\x6b\x76\xea\x9f\x5d\xb3\x1a\xce\xa8\x60\xc1\xc1\xdb\x53\x0f\xcd\x75\xc0\x32\x1e\x1b\xaf\x89\xdd\x33\x90\x6b\x2d\x01\xa4\x6f\x00\x13\xb0\x20\x8e\x3b\x95\x89\x8d\x78\x1f\x22\x0d\x18\x95\x23\xbe\x22\xbd\x4e\xc5\xe7\x8b\x4c\x72\xc4\x74\x84\x67\x4b\x59\x7e\x8e\x3d\x47\xb9\x48\xb6\x4c\x9d\xf1\xcd\xfa\x6b\x36\x9d\x31\x49\xf5\x5c\x32\xd2\xeb\x7c\xec\xea\x2f\x6e\x5b\x65\xf5\xb5\xf5\xca\xa7\xcc\xfa\x16\xca\x29\xd5\x40\x1e\x42\xe1\xd3\x35\x29\xa5\xd4\x7c\xf4\xbe\x42\x14\xc1\x59\xf7\x5a\xa3\x5d\xad\x93\x8d\xeb\x09\x0d\x39\x0c\xe5\x99\x7f\x97\x94\xfa\x95\x9c\x2c\x19\xd5\x13\x26\xcb\xdc\x0b\x45\xce\x92\xcd\x32\x8f\x8a\x05\x55\x30\x61\x7c\x3c\xd1\x5d\xd2\xaa\x13\x58\x72\x5f\x4f\x92\x47\xee\x77\x89\xfa\x77\xed\x85\x42\x0d\x53\x82\xcc\x88\x24\x8b\x72\xc0\x7c\x9a\xaa\x33\x15\xb7\x3e\x59\xb1\x66\x2b\x6c\x0c\x8f\x96\x25\x9b\xc4\x0a\x42\x8f\x6a\x1e\x8a\xa1\xa0\x53\x66\xda\xfb\x21\x7b\x1d\x6e\x5e\x49\x2a\xc6\x0c\x0a\x4b\x68\x77\xf7\x63\xf1\x96\x64\x80\x38\x86\xed\x88\x15\x96\x56\x9f\x29\x4f\xf2\x99\x89\x03\x05\x88\xe3\x28\x02\x26\xfc\xcd\xb0\xf0\x1e\xbe\xed\x7e\x9f\xa3\x2b\xc6\xe9\xd2\xee\x39\xc7\x19\xcc\xe6\x41\x50\x96\x66\x06\xf2\x18\x9f\x38\xdb\xc5\x3e\x1b\x4b\xc6\x14\xe9\x1d\x6e\x9d\x57\x36\x9d\x25\x72\xe3\xfc\xf1\xc2\xf2\x04\x2a\xeb\x2c\xc3\x25\x6c\xab\xf3\xe9\x5a\x21\x25\xed\xdb\xcc\xbf\x42\xc1\xa7\xeb\x4d\xb3\x13\x91\x8f\x63\x14\x7c\x72\x18\x04\x1a\x0a\x1c\xdc\x73\x4e\x48\x33\xae\x55\xb2\x4d\xf2\x72\x8a\x9b\x80\xd2\xeb\x80\x19\x7d\x5b\xe9\x32\x0d\xf8\x58\xb4\x3d\x26\x34\x93\x97\xc7\x7b\x1f\x70\xbd\xe9\x8f\x29\xef\x44\xd6\xaa\x36\x39\x9c\xe5\x22\xea\xf4\x37\x15\x73\x2a\x8d\x56\x16\xce\x1e\x23\xc7\xc9\xd1\x01\xd9\xc4\x36\xc3\x61\x3d\x70\x91\x3b\x20\xd7\x07\xa2\xab\xeb\x02\xa5\xfa\x76\x28\x64\x59\x4a\xfa\xd4\xf4\x7c\xa7\x7b\x8e\x4d\x76\x72\xe8\xd8\xd7\xc8\x5d\x82\xb8\xb1\x43\xf8\xc6\x85\xff\x32\x63\xcc\x08\x01\x74\x78\x6f\x5a\x51\xa9\x1e\x34\xfe\xb4\x66\xee\x04\x07\x9b\xf0\x33\x6f\xe0\xd3\x7b\xeb\xe2\x9d\xe0\xe0\xdf\xf4\x31\xfd\xd9\xa3\xd8\x3f\x7d\x39\xb8\x6d\xfd\x52\x97\x2f\x5a\x51\x04\x4a\x53\xcd\xbd\x1f\xaf\x0f\x03\x28\x26\xcf\x3f\x9f\x07\x40\x2a\x1f\xe1\x78\x4c\xa5\xae\x50\xa5\x98\x56\x95\x05\x13\x7e\x28\x2b\x69\xa3\xb7\xbf\x96\x49\x32\xa2\x81\x62\xa5\x4d\xfa\x4e\xa2\xd9\xa0\xd7\x33\xd6\x25\x74\x36\x0b\x78\x72\x14\x54\x7e\xd1\x05\x4d\x5e\xa6\xbb\x6c\x34\x17\xde\x46\xdc\xbd\x50\x2c\x98\xd4\x2f\x49\xc4\x62\xba\x4b\xef\xfd\x12\x44\x3b\x12\xd4\x92\x6b\x6f\x02\xf8\x4b\xf3\x57\xa9\x80\x64\x7a\x2e\x05\xa4\x71\xac\xdb\xc1\xdd\xcd\xf3\xf0\xf1\xfe\xfb\x8f\xd7\xbf\x2e\xb8\xfe\x73\xf3\xfc\x3a\x78\x1f\xde\x0e\x9e\x7e\xf6\xdf\x87\xfd\x9b\xf7\xcf\x2d\x48\x73\x1c\xad\xf1\xa8\x62\xd0\xaa\xd6\xda\x98\xd5\x46\xad\x0e\x6a\xad\xb7\x4f\x66\xe6\xa4\x50\x83\x02\xcb\x6f\x57\xab\x48\x4c\x1b\x45\x65\xa3\xa8\x6c\x07\x8d\xe0\xa0\x11\x1c\x2c\x82\x83\x62\x70\x50\x0c\x0e\x8a\xc1\xa9\xa1\x11\x6a\x68\x84\x1a\x1e\x01\xe3\xd6\xa9\xd5\x31\xab\x8d\xc5\x6d\xa0\x55\x34\xd0\x2a\x1a\x68\x15\x0d\xb4\xbf\x8d\x6c\x7f\x13\x2b\x5a\x5b\xc3\x46\x31\xe0\x78\x6d\x14\x43\xb6\x6f\x70\x3a\x4d\xcf\x37\xf7\x8f\xd8\x2c\xb9\x28\x26\x17\xe5\xdb\xcd\xf2\x8d\xe4\x79\x19\xdc\xdd\xa1\x9b\xc6\x45\xa9\x76\x51\xaa\x5d\x94\x6a\xb7\xd6\x40\xad\x2e\x66\x45\x49\x75\x51\x52\xdd\x2c\xa9\x58\x59\x8f\x4f\x6f\x58\x55\x4d\x74\xcb\x34\x5b\x97\x1b\xf2\x76\xff\xd8\x47\x23\xa2\x8c\x34\xd1\x36\x35\xd1\x7a\x9a\x75\xd4\xda\x40\xad\x2e\x6e\xc5\xf8\x6f\x36\x2f\x57\xf5\xed\xe9\x3b\xae\x98\x99\x76\xf8\x6c\x44\xe7\x81\xbe\x42\x04\x8d\xda\x1f\x49\x77\x72\x3a\xc7\x49\x96\x42\xd1\x0f\xbd\xf9\x94\x09\x5d\xb2\x24\xa3\xfe\xba\xb8\x3b\x7f\x8a\x87\x67\xc9\x82\x4a\x30\xdf\x76\x0a\xba\x20\xd8\x72\x1b\xbf\x18\x99\x2b\x67\x28\x49\x1b\xc8\xff\x9a\x4e\xab\xf9\xff\x5b\x12\x97\xf6\xb9\x36\x4b\x2c\xc5\x74\xf1\xe4\x9b\xee\x6b\xf6\x80\x8b\x22\x28\x72\xe1\xb3\xd5\xe9\x17\x50\xb5\x64\xdd\xf7\x21\x8e\x4b\xa5\x03\x72\xf6\xf7\xe9\xdc\xbb\x34\x02\xe1\xcc\x6d\x2c\x07\x51\x72\xc7\x42\xd0\x1c\x80\xd9\x5e\x35\x8e\x13\xcf\x02\xba\x2e\xa6\x8e\x86\x99\x4e\x25\x39\xed\xd1\x7b\xca\x7f\x01\x00\x00\xff\xff\x80\x5e\x1b\x52\x14\x12\x00\x00"

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
