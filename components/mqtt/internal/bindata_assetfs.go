// Code generated by go-bindata.
// sources:
// templates/views/subscriptions.html
// locales/ru/LC_MESSAGES/mqtt.mo
// locales/ru/LC_MESSAGES/subscriptions.mo
// DO NOT EDIT!

package internal

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

var _templatesViewsSubscriptionsHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\xc1\x6e\xda\x40\x10\xbd\xf3\x15\xa3\x6d\x2a\xb5\x52\xf1\x86\x9c\xaa\xc8\x70\xe9\xa5\x07\xaa\x28\x85\x9e\xab\xf5\xee\x50\x2f\x75\x76\xdd\x9d\x0d\x04\xad\xf8\xf7\xca\x36\x84\x85\xd8\x8e\x44\x49\x4f\x98\x99\x79\xf3\xde\xcc\xb3\x27\x04\x50\xb8\xd0\x06\x81\x49\x6b\x3c\x1a\xcf\x60\xbb\x1d\x0c\x52\xa5\x57\x20\x0b\x41\x34\x66\xce\xae\x19\x68\x35\x66\x0a\x57\x5a\x22\xb1\xc9\x00\x00\x20\x2e\x79\xfa\x59\x0a\x83\xc5\x2e\xf3\x32\xeb\xb5\x2f\x30\xca\xd6\x15\xf9\xcd\x24\x04\xd0\xa3\xcf\x06\xd8\xec\x31\x23\xe9\x74\xe9\xb5\x35\xc4\x20\x81\xed\x36\xe5\xf9\xcd\x09\x22\xea\x29\x0b\x14\x6e\xa1\x9f\xd8\x24\xe5\x4a\xaf\x22\xe2\x93\xbf\x47\x3a\xf6\x33\x76\xf7\xf5\x22\x2b\x70\xe8\x90\x4a\x6b\x48\xaf\x4e\x45\xd7\xe5\x75\xcd\x11\x00\x1a\x18\x79\xa7\x4b\x54\xa0\x84\x17\x4d\x5c\xf9\xa8\x17\x18\xbb\x76\xa2\x64\x40\x7e\x53\xe0\x98\xad\xb5\xf2\xf9\xed\xe8\xfa\xfa\x7d\x0b\x4b\xc3\x94\xa3\x50\x5d\x39\xd7\x9e\xd8\x01\xf7\xfa\x1e\xd4\x50\xda\x62\x38\x62\x93\x77\x29\xf7\x79\x2f\xe6\xe0\xc7\xdc\x96\x5a\xee\x7d\x78\x05\xf5\x92\xe9\xb9\xcd\xfd\xdd\xec\xdf\x9b\x7c\x11\x45\x91\x09\xf9\x9b\x2e\xd3\xea\xd5\x36\x29\xef\x5a\x6d\x85\xe9\x31\x24\xb3\x6a\xd3\x9e\x0b\x01\x9c\x30\xbf\x10\xae\xf4\x27\xb8\xa2\xe8\x65\x87\xdb\x31\x24\x71\x80\xaa\xef\xaf\xbd\x7f\xaf\xe1\xaa\x9a\x52\x28\x05\x57\x1a\x46\xcd\x7c\x1d\x42\xa3\xfa\x23\x29\x49\x6d\xfa\x79\xd0\xfb\xbb\xd9\x79\xc0\x29\x9a\xf3\x80\xb5\x95\xbd\xd0\x10\x00\x8d\xea\x5c\x27\xef\xf0\x2b\xe5\xf5\xa7\x7b\x72\x22\x3a\x8f\xcc\xee\x71\xf7\x73\xe0\x1c\x44\xa7\xb5\x7a\x69\xd8\x5e\x48\x08\x40\x5e\x78\x2d\xbf\xce\xbf\x4d\xe1\x43\xf3\xfc\xe3\xfb\x14\x18\x57\x82\xf2\xcc\x0a\xa7\xb8\x20\x42\x4f\x7c\x85\x46\x59\x47\xfc\xf9\xa0\x50\x62\xd0\x0f\x33\xe2\x92\x9a\xe8\xbc\x89\x66\xd6\x7a\xf2\x4e\x94\xc9\x83\x36\x89\x24\x62\xb0\x10\x05\xe1\xc7\x0b\xb2\x1e\x0e\xd9\x5e\xc0\x21\xd2\x2f\xa0\x7d\x2b\x4b\xba\xe0\x4e\xf8\x92\xf8\xf2\xcf\x23\xba\x4d\x12\xad\xa5\xd2\xb2\x7c\x8b\x5d\x64\x54\x11\x76\x1a\xf0\x26\x9c\x87\x6d\x9f\x70\x47\x36\xfc\x07\xf2\xdd\xec\x9d\xde\x2f\xdb\xad\xff\x1b\x00\x00\xff\xff\x8e\xf8\xff\xd4\x6a\x08\x00\x00"

func templatesViewsSubscriptionsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsSubscriptionsHtml,
		"templates/views/subscriptions.html",
	)
}

func templatesViewsSubscriptionsHtml() (*asset, error) {
	bytes, err := templatesViewsSubscriptionsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/subscriptions.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesMqttMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\x41\x6f\xd4\x30\x10\x85\xdf\x56\xcb\x81\x1c\x39\x73\x18\x0e\x54\x20\x70\xb1\x13\x2a\x55\xde\x78\x8b\x28\xad\x84\x68\xc4\xb6\x32\x5c\x38\x59\x1b\x93\x8d\x48\xec\xc8\x76\x10\x48\xfb\x37\xf8\x47\xfc\x0e\x7e\x0b\xda\xac\x96\x15\xef\xf4\x3d\xcd\x9b\xa7\xd1\xfc\x79\x34\xff\x05\x00\x27\x00\x1e\x03\x78\x09\xe0\x01\x80\x12\x7b\xad\x00\x3c\x04\x70\x07\x60\x33\x03\xbe\x00\x98\x03\xf8\x3d\x03\x66\x38\xea\xe4\x00\xbd\x75\xe3\xbc\xba\xd3\x1a\xf7\x76\xf0\x21\xb1\x2a\x36\x6d\xcd\xde\x8e\x4d\x64\xda\x4b\xaa\xed\xf7\x37\xdf\xda\x8d\xe9\xfd\x59\x18\xb3\xd5\x47\xcd\xae\x82\x35\xa9\xf5\x8e\xbd\x33\xc9\x4a\xca\xb9\xb8\x60\xbc\x60\x79\x41\x79\x21\xcf\xcf\x5f\xf0\x82\xf3\xec\xd6\xc4\xc4\x74\x30\x2e\x76\x26\xf9\x20\xe9\xc3\xd4\x41\xd5\x18\x4c\xef\x6b\x4f\xe5\x7f\xc5\xcb\xec\xd6\xb8\x66\x34\x8d\x65\xda\x9a\x5e\xd2\x3f\x2f\xe9\x7e\x8c\xb1\x35\x2e\xab\xde\x57\xd7\xec\xb3\x0d\xb1\xf5\x4e\x92\x38\xe3\xd9\x95\x77\xc9\xba\xc4\xf4\xcf\xc1\x4a\x4a\xf6\x47\x7a\x35\x74\xa6\x75\x0b\x5a\x6f\x4c\x88\x36\xa9\x4f\xfa\x86\x5d\x1c\x73\xbb\x7b\xbe\xda\xc0\xae\xdd\xda\xd7\xad\x6b\x24\x65\xab\x6e\x0c\xa6\x63\x37\x3e\xf4\x51\x92\x1b\x26\x1b\x55\xb1\xa0\x3d\x2a\xf7\x54\x70\xa5\x04\x9d\x9e\xd2\x0e\xf9\x13\x25\x04\x5d\x12\x27\x39\xf9\xa5\xca\x0f\xa3\x52\xbd\xde\xe1\xb3\x29\x56\x0a\x4e\xdb\xed\x7e\x65\xa9\x72\xfe\x9c\x2e\x49\x90\xa4\x7c\x81\xe9\xdb\x7f\x03\x00\x00\xff\xff\xc4\x1a\x03\x04\xc9\x01\x00\x00"

func localesRuLc_messagesMqttMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesMqttMo,
		"locales/ru/LC_MESSAGES/mqtt.mo",
	)
}

func localesRuLc_messagesMqttMo() (*asset, error) {
	bytes, err := localesRuLc_messagesMqttMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/mqtt.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesSubscriptionsMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\xcf\x6a\x14\x4d\x14\xc5\x4f\xfa\x4b\x3e\xb0\x45\x17\xc1\xa5\x8b\xeb\xc2\xa0\x48\xc5\xee\x1e\x03\xa1\x92\x4e\xc4\x98\x80\x38\x43\x62\xd2\xba\x72\x61\x65\xa6\x9c\x69\xd2\x53\xd5\x54\x75\x8b\xc2\x6c\x5c\xb9\x70\x27\xb8\x71\x23\xb8\xd2\xa5\xff\x70\x9c\xe0\xf8\x00\x6e\x6a\x5e\xc0\x67\x91\xee\x99\x18\x3c\xab\xdf\xb9\x75\xee\xe1\x42\xfd\x5e\x9c\x7f\x0d\x00\xff\x03\xb8\x08\xa0\x09\xe0\x2c\x80\x01\xa6\x7a\x0f\xe0\x0c\x80\x0f\x00\x16\x00\x7c\x06\xf0\x1f\x80\xef\x00\xce\x01\xf8\x31\x9b\xff\x02\xd0\x9b\x03\x26\x00\x16\x01\x70\x0f\x38\x0f\xa0\xe5\x4d\xf3\x0f\x3d\xe0\x02\x80\x47\x1e\xe0\x03\x28\x3d\x60\x6e\xd6\x55\xc9\x9b\xdd\x50\x69\x1e\xa7\x5a\x38\x81\x2d\x91\x65\x87\xa2\x7d\x64\x6b\xb2\xb8\xb7\x7b\x80\x83\xf2\xd0\xb6\x4d\x9a\x17\xa9\x56\x16\x89\xce\xd3\x36\xf6\x65\xae\x4d\xc1\x5a\xb6\x9b\x76\xd8\xad\xb2\x6b\x59\xa2\x39\x75\xe4\x93\x9b\x47\x69\x4f\xf4\xf5\xb2\x29\xfd\xbd\xdd\x84\x6d\x19\x29\xaa\x3d\x76\x5b\x14\x92\x53\x14\x84\xab\x2c\x68\xb0\xa8\x41\x51\x83\xaf\xac\x5c\x0b\x1a\x41\xe0\x37\x85\x2d\x58\x62\x84\xb2\x99\x28\xb4\xe1\x74\xb7\xee\xa0\x56\x69\x44\x5f\x77\x34\xad\xff\x53\xbc\xe1\x37\x85\xea\x96\xa2\x2b\x59\x22\x45\x9f\xd3\x5f\xcf\x69\xbf\xb4\x36\x15\xca\x6f\xdd\x69\x6d\xb3\x07\xd2\xd8\x54\x2b\x4e\xe1\x72\xe0\x6f\x69\x55\x48\x55\xb0\xe4\x59\x2e\x39\x15\xf2\x69\x71\x3d\xcf\x44\xaa\xd6\xa8\xdd\x13\xc6\xca\x22\xbe\x9f\xec\xb0\xd5\xd3\x5c\x75\xcf\x63\x69\xd8\xb6\x6a\xeb\x4e\xaa\xba\x9c\xfc\xbd\xac\x34\x22\x63\x3b\xda\xf4\x2d\x27\x95\xd7\xd6\xc6\x8d\x35\x9a\x62\xac\x2e\x87\x41\x1c\x87\xb4\xb4\x44\x15\x06\x97\xe2\x30\xa4\x4d\x0a\x88\xd7\x7e\x23\x8e\x4e\x9e\xd6\xe3\x1b\x15\x5e\xa9\x63\xeb\x61\x40\x83\xc1\x74\x65\x23\x8e\x82\xab\xb4\x49\x21\x71\x8a\xd6\xe0\xde\xb8\xb1\x3b\x76\xc7\xee\xa3\xfb\xea\x46\x6e\xec\x3e\xc1\xbd\x9a\xbc\x74\xdf\x2a\xac\x6d\xf5\x45\xee\xad\x1b\xbb\x2f\xee\xa7\x1b\x4e\x9e\x4f\x5e\xb8\xa1\x1b\xb9\x21\xdc\x3b\x37\xae\x46\x6e\x84\x3f\x01\x00\x00\xff\xff\x92\x92\x13\xd9\x80\x02\x00\x00"

func localesRuLc_messagesSubscriptionsMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesSubscriptionsMo,
		"locales/ru/LC_MESSAGES/subscriptions.mo",
	)
}

func localesRuLc_messagesSubscriptionsMo() (*asset, error) {
	bytes, err := localesRuLc_messagesSubscriptionsMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/subscriptions.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/views/subscriptions.html": templatesViewsSubscriptionsHtml,
	"locales/ru/LC_MESSAGES/mqtt.mo": localesRuLc_messagesMqttMo,
	"locales/ru/LC_MESSAGES/subscriptions.mo": localesRuLc_messagesSubscriptionsMo,
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
				"mqtt.mo": &bintree{localesRuLc_messagesMqttMo, map[string]*bintree{}},
				"subscriptions.mo": &bintree{localesRuLc_messagesSubscriptionsMo, map[string]*bintree{}},
			}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"subscriptions.html": &bintree{templatesViewsSubscriptionsHtml, map[string]*bintree{}},
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
