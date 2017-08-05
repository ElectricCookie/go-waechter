// Code generated by go-bindata.
// sources:
// testEmail/test-email-template.html
// DO NOT EDIT!

package testEmail

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
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
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

var _testemailTestEmailTemplateHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\x5d\x4e\x33\x31\x0c\x7c\xaf\xd4\x3b\x58\x79\xff\xd6\x9f\xd4\x07\xa4\xb2\x1b\x09\x21\xb8\x00\xbd\x40\x1a\x07\x12\xc8\xcf\x2a\x31\x15\xa5\xda\xbb\xa3\xfd\xa1\xdd\x42\x51\xdf\x32\xe3\xf1\x64\x2c\xbb\xb6\x1c\xbc\x5c\x2e\x96\x0b\x00\x80\xda\x1a\x45\x72\x7c\x0f\x98\x1d\x7b\x23\x37\xa6\x30\x3c\x04\xe5\x3c\xdc\x91\x6a\xd9\xe4\x1a\xc7\xca\x4c\xea\x5d\x7c\x83\x6c\x7c\x23\x0a\xef\xbd\x29\xd6\x18\x16\x60\xb3\x79\x6e\x84\x65\x6e\xcb\x1a\x31\xa8\x0f\x4d\xb1\xda\xa6\xc4\x85\xb3\x6a\x7b\xa0\x53\xc0\x23\x81\xab\x6a\x55\xdd\xa0\x2e\xe5\xc4\x55\xc1\xc5\x4a\x97\x22\xa6\xdf\x6a\x9c\x52\x4e\x70\x9b\x68\x7f\x44\x03\x43\x6e\x07\xda\xab\x52\x1a\xa1\x53\x64\xe5\xa2\xc9\x02\x86\x58\x8d\x68\x15\x91\x8b\x2f\x6b\x58\x65\x13\xe0\xff\xad\x98\x0d\x31\x76\xfb\xef\x66\xf2\xff\x6c\xca\xee\xb3\xf7\xf0\x42\xc2\xb9\x70\x14\xb3\x7c\xcc\x29\xd4\x48\x2c\x2f\x95\x49\x1e\x0e\x50\xf5\x12\xe8\xba\x1a\x89\xce\x92\xce\x6d\x36\x69\x32\xf9\xdb\x66\x93\xae\x99\x3c\xbd\x6f\x5f\x8d\xe6\x2b\x71\x26\xd5\xc9\xec\x4c\x84\xe4\x7f\x52\x36\xa3\x04\xb8\x30\x7f\xef\x76\x9f\x22\x9b\xc8\x5d\xf7\xab\x09\x70\x7e\x20\x48\x6e\x77\xda\x1a\x1e\xd7\x56\xe3\x70\x83\x5f\x01\x00\x00\xff\xff\xae\x58\x5d\x93\x8a\x02\x00\x00")

func testemailTestEmailTemplateHtmlBytes() ([]byte, error) {
	return bindataRead(
		_testemailTestEmailTemplateHtml,
		"testEmail/test-email-template.html",
	)
}

func testemailTestEmailTemplateHtml() (*asset, error) {
	bytes, err := testemailTestEmailTemplateHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testEmail/test-email-template.html", size: 650, mode: os.FileMode(438), modTime: time.Unix(1501768215, 0)}
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
	"testEmail/test-email-template.html": testemailTestEmailTemplateHtml,
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
	"testEmail": &bintree{nil, map[string]*bintree{
		"test-email-template.html": &bintree{testemailTestEmailTemplateHtml, map[string]*bintree{}},
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
