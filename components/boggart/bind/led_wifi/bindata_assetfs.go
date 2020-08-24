// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package ledwifi

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\x51\x6b\xdb\x3c\x14\x7d\xef\xaf\xb8\x88\x3c\x7c\x1f\xcc\xf1\x9a\x42\x29\x23\x09\x0c\xd6\x3d\x6d\x30\x18\x63\x8f\x45\xb6\xae\x6b\x6d\xb2\xe4\x59\xd7\x69\x3a\xe3\xff\x3e\x64\xc9\x4e\xe6\xba\x71\x60\x2b\x2c\x2f\x91\xe4\xa3\x7b\x8f\xce\x3d\x57\x6a\x1a\x10\x98\x49\x8d\xc0\x52\xa3\x09\x35\x31\x68\xdb\x8b\xb5\x90\x3b\x48\x15\xb7\x76\xc3\x2a\xf3\xc0\xb6\x17\x00\x00\xc7\xab\xa9\x51\x51\x21\xa2\xcb\x15\xb8\x91\x2d\xfa\xd1\xde\x46\x97\xab\x80\x1f\xef\xd9\xdf\x95\x5c\xa3\x3a\xfa\xfa\x14\x41\x92\x14\x8e\x10\x1d\x2a\x5f\x6d\x3f\xdc\xbe\x83\xaf\xf2\xbd\x5c\xc7\xf9\x6a\x02\x71\xcc\x4e\x21\xaf\x32\xb9\x67\xdb\x75\x2c\xe4\x6e\x94\x70\x62\xe9\x37\x0e\xbd\x12\x13\x39\x32\x53\x15\x3d\xd0\x8d\xa3\xdc\x54\xf2\xa7\xd1\xc4\x15\x74\x73\xc5\x13\x54\x91\xc2\x8c\x18\x54\x46\xa1\x87\x31\x28\x90\x72\x23\x36\xac\x34\x96\x18\x48\xe1\x14\x2c\x0a\xae\x05\x03\x6d\x76\x5c\x49\xc1\x09\x9f\x26\x1c\x93\x93\x84\x85\x4f\x74\x5f\x99\xba\x9c\xa0\x38\xec\xea\x98\x38\xec\x86\x59\xe2\x84\xec\x50\x3a\x4d\x95\x51\x9e\x2a\x84\x42\x5e\xf5\x75\xbc\x9a\x2c\xe3\xf8\xd7\x34\x20\x2f\x6f\x34\xb0\xcf\x3e\xf4\xd2\xb9\xe6\x59\x2a\x71\x97\xea\x04\xd7\xa7\xc6\xba\xee\xf9\x5c\x9f\xc5\x67\x1c\x65\x06\x7a\x10\x68\x1e\xd7\x61\xa5\x2e\x6b\x02\x7a\x2c\x71\xc3\xd2\x1c\xd3\xef\x89\xd9\x0f\x82\x7e\xb3\x91\x7d\x90\x94\xe6\x0c\x34\x2f\xb0\xd7\xfb\xac\xc8\xe1\xe7\x0c\x11\xaa\xe4\xa4\xcd\x60\xd9\xcd\x96\x9f\xcc\x03\x56\xd0\xb6\xd0\x25\x45\xd1\x34\x80\x5a\xb8\x85\xf8\x8c\x13\xce\xe9\x0e\xd3\xdd\x70\xe6\xe7\x13\x9f\xfc\x19\x34\xc2\x7f\x0a\x35\x2c\x31\xcb\x30\x25\xfb\x3f\xbc\x7e\xce\x26\x7f\xee\x71\x9f\xe3\x65\x4c\x7e\x1b\x62\xff\x13\x2e\xb7\xa8\x30\xa5\xce\x31\xfd\x99\xbd\xeb\x46\x0a\x74\x0a\x06\x19\xc0\x6f\x9a\x8b\xdd\x34\x50\x71\x7d\x8f\xb0\x28\x8c\xc0\x57\xb0\xe8\x6e\x63\x78\xb3\x19\x4a\x78\x4a\x80\x81\xa1\x29\x49\x1a\x0d\x3b\xae\x6a\xdc\xb0\xa6\xf1\xe1\x96\x5f\xa4\xa6\x1b\x68\x5b\xe6\xdd\x81\x3f\xfc\x3a\x2c\x82\xd7\x3f\xba\x49\xdb\x06\xae\xe8\x3a\x22\x8c\xd8\x60\xfb\x6d\x5f\x93\x40\x6d\x01\x6d\xbb\x8e\x7d\xc2\xd9\xb3\xf9\x10\x33\xdd\xe0\x73\xbe\x40\x43\x9c\xc8\x7e\xec\x0c\xa5\xef\xac\x51\x52\x4c\x3e\x5d\x53\x1b\xce\xeb\x94\xe7\xbc\x57\x88\xc8\x64\x99\x45\x8a\xae\xe6\x8c\x97\xd4\x44\x46\xfb\xab\x0a\xdd\xc3\xe5\xaf\x43\x5b\x27\x85\x3c\xd8\x2e\x21\x0d\x09\xe9\xc8\xd6\x69\x8a\xd6\xb2\xa1\x62\xec\x6d\x59\xaa\x47\xdf\x44\xeb\xd8\x07\xfb\x9b\x32\xaf\x63\xa7\xc4\xc9\xa7\xfe\x68\x1a\x86\xe1\xef\x50\x9e\x5f\x01\x00\x00\xff\xff\xb4\xd7\xb5\x26\x11\x09\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xbd\x6e\xd5\x30\x1c\xc5\x4f\xf8\x58\x32\x76\x66\xf8\x33\x50\xc1\xe0\x92\x5c\x3a\x20\xdf\xeb\x7b\xf9\x50\x2b\x21\x6e\x24\x54\x02\x33\x56\x63\xd2\x88\xc4\x8e\x6c\x07\x51\xa9\x03\x74\x46\x62\x62\x46\x3c\x00\x63\xa5\x0e\x50\xf1\x08\xce\x0b\xf0\x2c\x28\xc9\x85\xaa\x67\x39\xe7\xe7\x73\x6c\xf9\xcf\xd6\x8d\xaf\x00\x70\x1d\xc0\x2d\x00\xbb\x00\x6e\x02\x58\x63\xd2\x9b\x0d\xcb\x8d\x97\x00\x56\x11\xd0\x00\xd8\x02\xf0\x23\x9a\xfc\x77\x04\x44\xb8\xd4\xb5\xcd\x9b\x78\xdc\xb6\xf5\x31\x5e\x7a\xe9\x15\x0e\x54\x6b\xac\x67\x99\x2b\xab\x82\x3d\xe9\x4a\xc7\x72\xc3\xa9\x50\xef\x1f\xbd\xab\x8e\x64\x63\x76\x6c\x17\xaf\xa5\xf3\x2c\xb7\x52\xbb\x5a\x7a\x63\x39\x3d\x1f\x2b\xca\x3a\x2b\x1b\x53\x18\x5a\x5c\xd9\x2f\xe3\xb5\xd4\x65\x27\x4b\xc5\x72\x25\x1b\x4e\xff\x99\xd3\x41\xe7\x5c\x25\x75\x9c\x3d\xcb\xf6\xd8\x6b\x65\x5d\x65\x34\xa7\x74\x27\x89\x9f\x1a\xed\x95\xf6\x2c\x3f\x6e\x15\x27\xaf\x3e\xf8\xfb\x6d\x2d\x2b\x3d\xa7\xc3\x23\x69\x9d\xf2\xe2\x55\xbe\xcf\x1e\x5e\xee\x86\xff\xbc\x55\x96\xed\xe9\x43\x53\x54\xba\xe4\x14\xbf\xa8\x3b\x2b\x6b\xb6\x6f\x6c\xe3\x38\xe9\x76\x44\x27\x1e\xcc\x69\x8a\x42\xdf\x49\x13\x21\x52\xda\xde\xa6\x21\x26\xb7\x45\x9a\xd2\x8a\x12\xe2\x23\x2f\xc5\xec\x5f\xb5\x10\xbb\x43\xbc\x3b\xce\x16\x69\x42\x27\x27\xd3\x95\xa5\x98\x25\xf7\x68\x45\x29\x71\x9a\xcd\x11\xbe\xf5\x1f\xc3\x79\xf8\x19\xce\xc2\xaf\x70\xde\x9f\xf6\x9f\x11\xbe\x87\x8b\xfe\x53\x7f\x1a\x2e\xfa\x2f\xc3\x61\x38\xc3\xdf\x00\x00\x00\xff\xff\x4e\x02\x3e\x89\xd3\x01\x00\x00"

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
