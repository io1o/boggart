// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package mercury

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x57\x4d\x6f\xdb\x38\x10\xbd\xe7\x57\x0c\x88\x0d\x92\x60\x21\x4b\x36\x36\xbb\x81\xa1\xf8\xb2\x6d\x4f\x0d\x50\xa0\x41\x0f\xbd\x18\x8c\x38\xb2\xd9\x50\xa4\x40\x8e\x3f\x02\xc3\xff\xbd\x90\x65\xd9\xb2\x6a\x29\x52\xe2\xb4\x68\x03\xc4\xe6\xc7\xcc\x9b\xc7\x99\xa7\xa1\xb5\x5a\x81\xc0\x58\x6a\x04\x16\x19\x4d\xa8\x89\xc1\x7a\x7d\x16\x0a\x39\x87\x48\x71\xe7\x6e\x99\x35\x0b\x36\x3a\x03\x00\x28\xaf\x46\x46\x79\x89\xf0\xfa\x03\xc8\x46\x2e\x29\x46\x4b\xe7\xf5\x07\x5b\xfb\xaa\xcf\x72\x9c\x72\x8d\xaa\xb4\xfb\xa3\x05\x49\x52\x58\xb1\xd8\x58\x4d\x07\xa3\xd5\x0a\x64\xff\x46\x03\xfb\x7f\x66\x2d\x6a\x82\x04\xb9\x9b\x59\x4c\x50\x93\x63\xd0\x83\xf5\x3a\xf4\xa7\x83\x23\xce\x65\xe2\x0a\xb9\x8d\xe5\x92\x8d\x42\x5f\xc8\x79\x85\xcb\x91\xa5\x03\x7a\x45\x92\x8e\xc4\x20\xfe\xa0\xb0\xb0\xcc\x27\x9b\x4f\xcf\x91\x95\x29\x0a\x10\x9c\x78\xbe\x2e\xc8\xb3\xe8\x52\xa3\x9d\x9c\x23\x68\xb3\xb0\x3c\x65\xe0\xe8\x49\xe1\x2d\x5b\x48\x41\xd3\x61\x3f\x08\xce\x8f\x44\xc9\x23\x4d\x91\x8b\xba\x3d\x7b\x7c\x63\xeb\x58\xf0\x4b\x84\x97\x55\x6b\xc0\xf6\x49\xbd\xdb\x27\xb3\xc8\x25\x4d\x5f\x0a\xf6\x85\xab\x19\xb6\x84\xd9\x7b\xbd\x43\x17\x59\x99\x92\x34\xfa\x39\xdf\xd0\xaf\x3b\x69\xe6\xd3\x90\x9f\x07\x23\x9e\x9a\xe8\x34\xe4\x2f\x37\x10\x25\xbe\x9c\x10\xb8\x16\x40\x32\xd9\x9f\xb6\x26\x74\xf1\x97\x79\xc7\xd0\x13\x9c\x30\x73\xeb\xa1\xb5\xc6\x66\x0f\x5d\x93\xd3\x36\xf4\x4e\x5e\xb8\x24\x4f\x70\x3d\x41\xbb\x49\x7a\x05\xad\xf7\x7e\x8b\xd9\x8a\x0d\x2a\x87\x2d\xe3\x37\x63\xed\x0c\xf3\x2a\x02\x3d\xa5\x78\xcb\x78\x9a\x2a\x19\xf1\xac\xa8\xfe\x37\x3e\xe7\xf9\x26\x1b\x09\x13\xcd\x32\xb1\xf5\x16\x56\x12\x5e\x66\x27\xb8\x37\x9f\xc9\x4a\x3d\xb9\xbc\x38\x38\xd3\x3c\x13\x53\xef\x83\xb1\x09\x27\x60\x83\x20\xf8\xd7\x0b\xfa\x5e\x30\xb8\xef\x5f\x0f\x83\x7f\x86\xc1\xf5\xd7\xe0\xbf\x61\x10\x64\x9d\xeb\xe2\xea\x2a\xf4\xf3\x08\xed\xb8\xe6\xd5\xd0\x58\x0a\x27\x50\x11\x87\xa0\x4d\x4a\xf6\x27\x26\x6b\xf4\xe4\x68\x79\x5a\x83\xe4\x54\x26\xf4\x3a\x2a\x05\x52\xff\x46\x7f\x52\x33\xcb\x15\xb0\xbf\xcf\x05\x38\x8c\x8c\x16\xac\x3c\xc9\x5a\x66\x25\xd2\x5f\xa0\xa5\xda\xfd\x57\x77\x3b\xb0\xe8\x20\xab\x06\xe2\x65\xde\x3f\x8d\xb6\x16\x9d\x2a\xef\xe7\xa5\x6f\xad\xb6\x96\xf8\xed\x9e\xdc\xe7\xb1\x0e\xfa\x55\x71\x6f\x8a\x72\xdf\x02\xa3\x21\x41\x42\xdb\xa6\x81\xd5\xb7\x5d\xe8\xdc\x3d\xef\xf8\x23\x6e\xa8\x74\xec\x9c\x09\x7f\xc4\x71\xe6\x77\x9a\xd6\x59\x81\x7b\xbb\xde\x79\x18\xac\xae\xa9\xb1\xd6\xa1\x3b\x16\xbf\x53\xba\x4f\x59\xe7\x8f\xdc\x11\xa4\x66\x81\x16\x4c\x1c\x77\x2c\xb6\xe2\x8e\xc6\x1b\xe7\xb1\x89\xe3\xf1\x69\x6f\xcd\x46\xf0\xb7\x13\xc2\x89\xee\xc7\x3a\xf6\x2f\xbc\x2e\x4f\x29\xba\x5f\x23\x2e\xfd\x0a\x6d\xe9\x37\x94\x56\x15\xfb\xb7\x52\x96\xfe\x23\x84\x15\xfa\x35\x3f\xfc\x43\x7f\xf3\x4a\xd6\xf8\x1a\x58\x9a\x6e\x87\xdb\xaf\x1d\xef\xef\x01\x00\x00\xff\xff\x63\xd4\x20\x28\x47\x0f\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xdf\x6b\x1c\x55\x14\xfe\xc6\x9d\xaa\x5d\x7f\xa0\x22\x3e\x89\x9e\x20\xa9\x4a\x98\x3a\xb3\x6b\xab\x4e\x32\xa9\x98\xb6\x20\x76\xb1\x94\xb5\xf8\x7a\xc9\xdc\x6c\x86\xee\xdc\x59\xef\xcc\xf8\x03\x82\x74\xe3\x43\x91\x8a\x7d\x51\x7c\x10\x54\x7c\x12\x9f\x96\x90\xa1\xcb\xc6\xdd\x80\x7f\x80\x9c\x0b\x3e\x8a\x4f\xfe\x21\x32\xb3\x6b\x36\xd9\x48\xef\xcb\x9c\xef\x3b\xdf\xf9\xbe\x73\x87\xfb\xf7\x33\xf6\xb7\x00\xf0\x38\x80\xe7\x01\xec\x00\x78\x1a\xc0\x5f\x98\x1e\xb2\x80\x67\x01\x2c\x59\xc0\x73\x00\x2e\x5a\xc0\x0b\x00\x5a\x33\xbe\x6b\x01\x4f\x00\xe8\x5b\xc0\x63\x00\xee\x59\xc0\x93\x00\xbe\x9f\xf1\xbf\x5a\xc0\x59\x00\xc3\x59\xff\x77\x0b\x38\x03\xe0\x0f\x0b\xd8\xb6\x80\x3f\x2d\xa0\x09\x20\xa8\x01\x17\x01\x7c\x54\x03\x56\x01\xfc\x52\x03\x5e\x02\xf0\x4f\x6d\x9a\xfb\xa8\x0d\x3c\x05\xe0\x45\x1b\x78\x19\x80\x6b\x03\xcb\x00\x3e\xb4\x81\x25\x00\x9f\xdb\xd3\x7d\x7e\x9a\xe9\x0a\x1b\xb0\x00\x3c\x02\xa0\x36\xbb\x8b\x8d\xe9\x0e\x67\x67\xf8\xcc\xec\xfb\x30\xe6\xa7\xfc\x0f\x0f\x95\x79\x00\xea\x25\xb1\x1c\x52\x2a\x37\x13\x15\xce\xab\x14\x2b\x73\x76\xe5\x18\xbd\x91\x6b\x2d\x55\x46\xa1\xc8\x24\x09\x15\x52\x16\xc5\x92\x12\x45\xb1\xcc\xa4\x3e\x6a\xc7\x52\xa4\xb9\x96\xb1\x54\x59\x8a\xcb\xc7\xb5\xb8\x2c\xd3\x4d\x1d\xf5\xb2\x28\x51\xb8\x26\xd2\x8c\x7a\xc9\x27\x52\x53\xb2\xb5\x75\x02\x2a\xb4\xc4\x2d\x59\xe5\xa0\x35\x77\xc3\x4d\xd1\xcd\x25\x6e\xc8\x5e\xa2\x33\xa7\x95\x76\xa2\xd0\x79\x27\xef\xa4\x4e\x3b\xf1\x29\x94\x1f\xbf\x7d\x2b\xda\x16\x71\x72\x5e\xe7\xf5\xeb\xef\xb7\x9d\x0d\x2d\x45\x99\xe4\x94\x3b\xf8\xd4\x70\xbd\xb7\x1c\xb7\xe9\x34\xde\xa0\x46\xd3\xbf\x70\x61\xc5\x6d\xba\x6e\xbd\x8c\x75\xda\x5a\xa8\xb4\x2b\xb2\x44\xfb\xf4\x5e\xe5\x41\xad\x5c\x8b\x38\x09\x13\x5a\x3b\x61\xbc\x5e\xbf\x26\x54\x27\x17\x1d\xe9\xb4\xa5\x88\x7d\x3a\xc2\x3e\xdd\xc8\xd3\x34\x12\xaa\xde\x7a\xb7\x75\xc5\xb9\x29\x75\x1a\x25\xca\x27\xef\xbc\x5b\xdf\x48\x54\x26\x55\xe6\xb4\x3f\xeb\x49\x9f\x32\xf9\x69\xf6\x5a\xaf\x2b\x22\xb5\x4a\x9b\xdb\x42\xa7\x32\x0b\x3e\x68\x5f\x75\xde\x9c\xeb\xca\x7d\xb6\xa4\x76\xae\xa8\xcd\x24\x8c\x54\xc7\xa7\xfa\xf5\x6e\xae\x45\xd7\xb9\x9a\xe8\x38\xf5\x49\xf5\x2a\x98\x06\xcd\x55\x9a\x96\x81\x5a\xf6\xdc\x20\xf0\xe8\xdc\x39\x2a\x4b\x77\x29\xf0\x3c\xba\x44\x2e\xf9\x15\x5e\x0f\x1a\xff\xb5\xd6\x82\xd7\xcb\xf2\x95\x4a\xb6\xe6\xb9\xb4\xb3\x33\x1d\x59\x0f\x1a\xee\xab\x74\x89\x3c\xf2\xa9\xb1\x5a\xbe\x08\xd3\xe7\x82\x47\xe6\x0b\x1e\xf3\x3e\x0f\x16\x19\x73\x77\x91\xa9\xde\xcb\xc2\xd0\x22\x65\xee\x9e\xa2\xc0\x3f\x4f\x81\xf9\x92\x87\x5c\x50\x39\x67\x76\x79\x40\x3c\x24\xde\x33\xb7\xb9\xe0\xdf\xcc\x3d\xe2\x31\x0f\xc8\xf4\xcd\x1d\x2e\xcc\xae\xb9\xc3\x43\x1e\x71\x71\x6a\xf6\x90\x27\x3c\xe2\x01\xdf\xaf\x3c\x0a\x3e\xe0\x21\xf8\x9b\xff\x33\x04\xff\xc0\x87\x3c\x34\x7d\x1e\xf0\xb8\x1c\x06\xff\xc8\x13\xd3\xe7\x03\x2e\x78\x9f\xc7\x5c\x94\x7e\x13\xb3\xcb\x23\x3e\x30\x5f\x97\xb9\x0f\xd0\xed\x9d\x56\x1d\xa5\x1e\x9a\xdb\x3c\xe1\x21\xdf\xe7\x3d\x9e\xf0\xbe\xe9\x9b\x5d\xde\xe3\x41\xe5\x73\x72\x59\xf3\x15\xf8\xbb\xf2\xa2\xc7\x7c\xfe\x0d\x00\x00\xff\xff\x2c\x66\xbc\x90\xc6\x04\x00\x00"

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
	"templates/views/widget.html": templatesViewsWidgetHtml,
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
