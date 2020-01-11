// Code generated by go-bindata.
// sources:
// templates/views/index.html
// DO NOT EDIT!

package mqtt

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

var _templatesViewsIndexHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\xdf\x4f\xe3\x38\x10\x7e\xe7\xaf\x18\x59\x1c\x82\x87\x24\x94\xa7\x13\x4a\x39\x9d\x0e\x71\x8b\xc4\xc2\x0a\xca\xf3\xca\x8d\xdd\xc6\x55\x62\x07\xcf\xb4\xa5\x8a\xf2\xbf\xaf\x9c\xf4\x47\x5a\x9a\x36\x54\xc0\xee\x4b\x3b\x19\xcf\xcc\x37\xf3\x7d\x8e\xe3\x3c\x07\x21\x07\x4a\x4b\x60\x91\xd1\x24\x35\x31\x28\x8a\xa3\xa3\x3c\x07\x35\x00\x3f\x32\x69\x66\xb4\xd4\x84\xce\x1b\x0a\x35\x81\x28\xe1\x88\x5d\x66\xcd\x94\x5d\x1d\x01\x00\xd4\xbd\x91\x49\xbc\x54\x78\x9d\x0b\x70\x16\xa6\x0b\xeb\x15\xbd\xce\xc5\x3c\x7e\x33\xe7\xf5\x67\xc6\xb5\x4c\x6a\xab\x6f\x23\x16\xbd\xad\xc7\x94\x71\xc4\xfb\x89\x5c\x44\x56\x0f\xe5\xaf\x87\x64\x55\x26\x05\x08\x4e\xbc\xf2\x0b\xf2\xac\xc4\xcc\x68\x54\x13\x09\xda\x4c\x2d\xcf\x18\x20\xcd\x12\xd9\x65\x53\x25\x28\xbe\xec\x9c\x9f\xff\xb5\x05\xa5\x42\x8a\x25\x17\xdb\xd7\xaa\x75\xdb\xbc\x38\x2f\x70\xe5\x78\xed\xfc\xad\x81\xdd\x5e\x33\xf0\xa1\x28\xc2\x80\xe2\x77\xa4\xf5\x66\x99\x3c\x28\xf1\x9e\xa7\x87\x25\x3e\x11\xa7\xc3\x32\xff\x8d\x48\x19\x8d\x6d\x72\xc3\xa0\x89\x3c\x97\xd7\x48\x7b\x48\x7d\x23\x66\xcd\x65\xf3\x1c\x2c\xd7\x43\x09\xc7\xcb\x8d\x0c\x97\xdd\xcd\x6d\xdd\x3c\xcc\x5e\x41\x85\x9b\x76\x55\xdc\xff\x5f\xd2\xed\x75\x35\xed\x8e\x9d\xd2\x98\xeb\xd4\x3d\x3c\xdb\x49\x7c\x78\x76\xa9\x73\xeb\xf4\x9d\x01\x50\x51\xaf\x06\xa0\xe5\x06\xca\x7f\x26\x4d\xb9\x16\x3d\x93\xa9\xc8\x7f\x22\xab\xf4\x10\x18\xdb\xa5\xc2\x12\xb5\x76\x20\xf4\x49\x7b\x43\x6b\xc6\x19\x03\x6b\xdc\xcb\x5b\x3d\xec\xef\x6a\xd5\x99\xb1\x70\x2a\x5f\xb6\xf0\xbf\xec\x2a\x51\xc3\x98\xd8\xd9\xbe\x30\x9c\x2a\x8a\x62\x76\xd6\x66\x86\xf5\x16\xde\xd4\xad\x24\x60\x0f\xf7\xad\x08\x59\x23\x87\x43\x6c\xe5\xa0\xcb\x9c\xac\xfe\xa3\x7c\x19\x4b\x24\xff\xf9\xf1\xce\xff\xc1\x29\x86\xa2\xf8\x87\x97\x2f\x63\x37\xaa\xf8\x3f\x59\xe2\x76\xb7\x6f\xe0\x93\x28\x15\xdd\x87\x9b\x1b\x56\xa3\x1c\x1c\xed\xc2\xbd\x4f\xb6\x34\x55\x64\x2a\xdf\x2b\xb6\xa4\x7e\xad\x65\xb5\x28\x3d\x4c\x66\x59\x5c\x16\x5b\x5a\x1e\x92\xc9\x18\x90\x22\xa7\xee\xea\x2c\x2a\xb9\x06\x33\x18\x30\x38\x86\xa2\x60\x57\x61\xa0\xde\x07\x1d\x06\xbc\x7d\x42\x9e\x83\x4c\x50\xfe\x21\x6a\xdc\xff\x2e\x31\xb2\x84\xcf\xde\x8a\xd1\x1b\x5b\x0d\x46\x7f\x9d\x12\x5a\xb4\x15\xa2\x7d\x74\x18\x08\x35\x69\x75\x9a\xb5\x28\xb8\xfb\xe8\x6c\xfe\xc6\xed\x47\x08\x83\x86\x2f\x5d\x18\x94\x57\x9b\x8d\xbb\xd3\xfa\x50\xb5\xc7\xb9\x39\xff\x5b\x61\x6e\x98\x8b\x2b\xa1\xfb\xf2\x2e\xcf\xa2\x3c\x07\x24\x4e\x2a\xfa\xd6\xfb\x7e\x07\xa7\x95\xfd\xfc\x78\x07\x2c\x10\x1c\xe3\xbe\xe1\x56\x04\x1c\x51\x12\x06\x13\xa9\x85\xb1\x18\x2c\xef\x5e\xe8\x6b\x49\x5e\x1f\x83\x08\x2b\x6f\xaf\xf2\xf6\x8d\x21\x24\xcb\x33\x3f\x55\xda\x8f\x10\x19\x0c\x78\x82\xf2\xec\x03\x51\x57\x77\xbe\x45\x03\x2b\xcf\xee\x06\xb6\xb3\x32\xc2\x0f\xe4\x24\x18\x61\x30\x7a\x19\x4b\x3b\xf3\x6b\xb4\xb8\x5e\x46\x9f\xc1\x45\x1f\x1d\x60\xa3\x00\x9f\x82\xb9\x62\x7b\x03\xbb\x26\xc3\x17\x80\xcf\x67\x6f\xd4\x7e\xd4\x20\xfd\xaf\x00\x00\x00\xff\xff\xc1\xdc\x56\x3d\x23\x0d\x00\x00"

func templatesViewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsIndexHtml,
		"templates/views/index.html",
	)
}

func templatesViewsIndexHtml() (*asset, error) {
	bytes, err := templatesViewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/views/index.html": templatesViewsIndexHtml,
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
			"index.html": &bintree{templatesViewsIndexHtml, map[string]*bintree{}},
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
