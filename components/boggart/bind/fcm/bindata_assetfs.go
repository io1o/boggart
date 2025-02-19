// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package fcm

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x4d\x8f\xda\x30\x10\xbd\xf3\x2b\x46\x3e\xb5\x95\x42\x04\x54\xab\x1e\x12\x2e\x95\xaa\x1e\x96\x4b\x3f\xce\xc8\x89\x27\xe0\x95\x63\x53\xcf\x84\xa5\x1b\xf1\xdf\x2b\xe7\x6b\xb3\x90\x52\x0e\xe5\xc2\x38\x7e\x9e\x79\x7e\x6f\x3c\x75\x0d\x0a\x0b\x6d\x11\x44\xee\x2c\xa3\x65\x01\xe7\xf3\x2c\x51\xfa\x08\xb9\x91\x44\xa9\xf0\xee\x59\xac\x67\x00\x00\xe3\xaf\xb9\x33\x51\xa9\xa2\xc5\x12\x42\x44\x65\x1f\x9d\x28\x5a\x2c\x3b\xfc\xe5\x99\xd3\xf6\x20\x2d\x9a\xd1\xee\x35\x82\x35\x1b\xbc\x40\x34\xa8\xfd\x72\x5d\xd7\xa0\x17\x9f\x2c\x88\x2f\xda\x63\x26\x09\xe1\xb3\x71\x95\x82\x0d\x12\xc9\x9d\xb6\x3b\x01\x73\x38\x9f\x93\x78\xbf\x9c\x48\x30\x26\x6f\x50\xfa\x42\x9f\xc4\x3a\x89\x95\x3e\x5e\xf0\x99\xf8\xf4\x86\x62\x2f\xd4\x44\x8d\xc2\xf9\xb2\x07\x86\x38\xda\x3b\xaf\x5f\x9c\x65\x69\xa0\x59\x1b\x99\xa1\x89\x0c\x16\x2c\xc0\x3b\x83\x2d\x4c\x40\x89\xbc\x77\x2a\x15\x07\x47\x2c\x40\xab\x54\x10\x5a\x25\xc0\xba\xa3\x34\x5a\x49\xc6\xeb\x6a\x97\xcc\x34\x63\xd9\x56\xd9\x79\x57\x1d\x26\xf8\x0d\xa7\x1a\x1a\x01\x9b\x8a\xb2\x11\x0f\xc5\xab\xb1\x96\xbd\x33\x2d\x53\xe8\x6c\x5e\xf5\x2e\xaf\x26\x4d\x9e\xfa\x0d\x66\x6d\xfa\x02\xc1\x1b\x48\xe8\x20\xed\xd0\x5a\xf8\xab\xd2\x1e\x95\x58\x7f\x48\xe2\xb0\x71\x83\x72\xdc\x10\xba\x01\xb8\x6e\xce\x87\x9e\xf5\xc3\xdd\xac\x13\xc6\x13\x4b\x8f\xf2\x8d\x8b\x9d\x26\xe0\x91\xf4\x8b\xcc\x0c\x6e\x7b\x58\x70\xf1\x99\x52\xf1\x51\x80\x95\x25\x8e\xe4\x0c\x1e\x0e\x8b\xfe\xa2\xe3\x2b\x27\x71\x9f\xe4\xd6\xad\xaf\x7a\xf1\x9e\xad\x91\x14\xc6\x6e\xc9\x19\xad\x26\x5b\x7d\xea\xc0\x7d\x0d\xf4\x37\xb1\x4b\x15\xb9\xa2\x20\xe4\x68\xf5\x2f\xa5\xb3\x8a\xd9\xd9\x51\xaf\xf3\xef\x03\xa6\x82\xaa\xac\xd4\x3c\xb4\x63\xc6\x16\x32\xb6\x11\x55\x79\x8e\x44\xe2\x75\x08\x7c\x6f\x0e\xb5\x0f\xbe\xcd\xf5\x3f\x75\x4c\xe2\x20\xc4\xcd\xc9\x30\x5a\x76\x61\xf7\x57\xd7\x80\x56\x85\x39\x3a\x1b\xcd\xd7\x27\x6a\x46\x2b\xb4\x6f\x83\x58\xb2\xce\xbf\xfe\xd8\x3c\xc2\xbb\x36\xfe\xf9\xed\x11\x44\xac\x24\xed\x33\x27\xbd\x8a\x25\x11\x32\xc5\x47\xb4\xca\x79\x8a\xbb\x39\xe0\x7c\xfc\x34\x5a\xcc\x4b\x6d\xe7\x21\x73\x21\x0d\xe1\xfb\x50\x60\xa8\xfe\x27\x00\x00\xff\xff\x81\x00\x27\xba\xdc\x05\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x3f\x6b\x14\x41\x18\xc6\x9f\x8b\x51\x61\xb1\x4a\x6d\xf1\x5a\x18\xb5\x98\xb8\x7b\x5a\xc8\xdc\xcd\x9d\x28\x09\x88\x59\x90\xb8\x5a\xd9\x0c\xd9\x71\xb3\x78\x37\xb3\xce\xcc\x8a\x81\x14\x92\x42\x10\x6d\x6d\x05\xbf\x81\x0a\x81\x80\x78\x7e\x00\x9b\xd9\x2f\x90\xcf\x22\xbb\x9b\xf3\x0f\xe4\x6d\xe6\xf9\x3d\xef\x6f\xe0\x3d\x59\x5b\xfd\x08\x00\xab\x00\x2e\x03\x18\x03\x38\x0f\xe0\x19\xfa\xa9\x00\x5c\x04\xf0\x12\xc0\x25\x00\xfb\xa7\xee\x5b\x00\xd3\x01\xf0\x1e\xc0\x1a\x80\x5f\x03\xe0\x1a\x80\x93\x41\xcf\x17\x56\x80\xc1\xa9\xdb\xce\x39\x00\x2b\x6d\x48\x95\x73\xb2\x50\xcb\x97\x9c\xd2\x1e\x8f\x95\xce\xb1\xa3\x2a\x63\x3d\x4b\x5d\x51\xe6\xec\x5e\x5d\x38\x96\x19\x4e\xb9\x7a\x75\xf7\x45\xb9\x27\xe7\x66\xc3\xd6\xd1\xb6\x74\x9e\x65\x56\x6a\x37\x93\xde\x58\x4e\x0f\xbb\x15\xa5\xb5\x95\x73\x93\x1b\x1a\xff\xe7\x4f\xa2\x6d\xa9\x8b\x5a\x16\x8a\x65\x4a\xce\x39\xfd\x61\x4e\x3b\xb5\x73\xa5\xd4\x51\xfa\x20\xdd\x64\x4f\x95\x75\xa5\xd1\x9c\x92\x8d\x38\xba\x6f\xb4\x57\xda\xb3\x6c\xbf\x52\x9c\xbc\x7a\xed\x6f\x56\x33\x59\xea\x11\xed\xee\x49\xeb\x94\x17\x4f\xb2\x2d\x76\xe7\xaf\xd7\xde\xf3\x5c\x59\xb6\xa9\x77\x4d\x5e\xea\x82\x53\xf4\x68\x56\x5b\x39\x63\x5b\xc6\xce\x1d\x27\x5d\x75\xe8\xc4\xad\x11\xf5\x51\xe8\xab\x49\x2c\x44\x42\xeb\xeb\xd4\xc6\xf8\x8a\x48\x12\x9a\x52\x4c\xbc\xe3\x89\x18\x2e\x57\x63\x71\xbb\x8d\xd7\x3b\x6d\x9c\xc4\x74\x70\xd0\x7f\x99\x88\x61\x7c\x83\xa6\x94\x10\xa7\xe1\x08\xe1\x73\x58\x84\x45\xf8\xda\xbc\x0b\x47\xe1\x47\x38\x0e\x47\x67\x54\x14\x16\xcd\x61\xf8\xd9\xbc\x09\x5f\xc2\xb7\xf0\xbd\xab\x17\x08\x9f\xfe\x29\x8f\x9b\xc3\xe6\x03\x7e\x07\x00\x00\xff\xff\xeb\x9c\x51\xab\x19\x02\x00\x00"

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
