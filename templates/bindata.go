// Package templates Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// readme.tpl
package templates

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _readmeTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x94\x51\x6f\xda\x30\x10\xc7\xdf\xfd\x29\x4e\x62\x0f\x05\x95\xf0\x4e\xd5\x49\x55\xbb\x69\xd3\xd8\x54\xb5\xb4\xd2\x84\x90\x72\x4d\x2e\xc4\xab\x63\x7b\xb6\x43\x87\x22\xbe\xfb\x64\x3b\x84\x40\xdb\x6d\x8f\xeb\x1b\x3e\xe7\xee\x7c\xff\xdf\xff\x18\x40\xd3\x40\x72\x6d\xd4\x0f\xca\x5c\x72\x5b\x2a\xe3\xbe\x61\x45\xb0\xdd\x32\xb6\xe8\x5f\xb5\xd1\xe5\x49\x3f\x78\x77\x33\x83\xed\x76\x08\xe3\x83\x2a\x57\x64\x33\xc3\xb5\xe3\x4a\x86\x3a\x83\x01\xcc\x67\x67\x57\x37\x67\x8c\xa5\x69\x9a\x29\x69\x95\x20\xf6\x0e\x4a\x12\x15\x18\xd2\x0a\x30\xcf\x43\x85\x1b\xd2\xca\x72\xa7\xcc\x66\xd7\xef\x38\x1c\x3b\x1e\x24\xd7\x3a\x47\xd7\xd5\xe3\xd2\x3a\x14\xa2\xcd\x13\x84\x96\x5e\xab\xd5\x86\x27\x3e\x7c\x59\xa2\xe9\x86\x84\xb1\x7c\x96\x6f\x35\x66\xfe\xae\x69\xe0\x89\xbb\x72\x97\x71\x4f\xc6\xc6\x39\x61\x3c\x5e\xc7\xc3\x79\xd3\x24\xe1\x43\x92\xb9\x7f\x6c\x9a\xa6\x41\x84\xcf\xd2\x19\x95\xd7\x99\x17\x86\xb1\x79\xc9\x2d\x64\xbe\x08\xe4\xa4\x85\xda\xd8\x03\x11\x2f\xb4\xf6\x45\x95\x04\x84\xc5\x97\xfa\x81\x8c\x24\x47\x76\x79\x52\x3a\xa7\xa7\x93\xc9\x63\x17\x4a\xb8\x1a\x42\x26\x6a\xeb\xc8\x40\x6d\xb9\x5c\x81\x2b\x09\x16\x9f\x48\x54\xf1\x73\x3b\x9d\x4c\xbc\x38\x89\x2d\x87\xa0\x31\x7b\xc4\x15\x41\x85\x12\x57\x64\x92\xf0\xb4\x6b\x43\x86\x7e\xd6\xdc\x72\x47\x96\x35\x0d\x18\x94\x2b\xf2\xaf\xe9\xc5\xfd\x2c\x91\xb4\xff\xd5\x34\xe3\xdd\x80\x71\xb8\xa0\xfb\xae\x7b\x50\x87\xb1\xb9\xea\x80\xf8\x68\x1c\x37\xe8\xe7\x8f\x26\xea\x0b\xd2\xcb\x9e\xbe\x40\x2c\x9d\xbe\x64\x99\xff\x15\xf1\xdc\x4f\xa8\xaa\x0a\x65\xfe\x17\xa4\x7e\xf8\x3d\xd3\x0e\x1e\x8f\x37\x39\x15\x58\x0b\x07\x99\x92\x05\x5f\xd5\x06\xbd\x61\x12\xf0\xe5\x17\x07\xb1\xe5\xc9\xe0\xe0\x3c\x04\x4b\xc1\x5d\x20\xb8\x75\x36\x14\xd3\x68\xb0\x22\x47\xc6\x1f\xd1\x41\x86\x12\x1e\xa8\xab\x4d\x39\xe4\xb5\xf1\xd4\x5a\x55\x63\x33\xc6\xde\xc3\x68\x34\xe7\x7a\x34\x9a\xc2\x8c\x5b\x07\x5e\xf0\x96\x97\x6d\x5d\x96\x06\x1a\xbe\x55\xf4\xf7\x9d\xe4\xaf\x99\xa0\xde\x5d\x4d\x72\x12\xe4\x28\x5c\xff\x3b\xf1\x36\xe9\x25\xe0\x7f\xe0\xf8\x9c\x8a\xa1\x4a\xad\xc9\xc2\xce\x8f\x7d\x06\xaa\xd2\x4a\x92\x74\x16\xd0\x5a\x95\x71\x74\x94\xef\xad\x1a\x9d\x1b\xc9\x8a\x90\xd0\x33\x70\xdc\xa1\xcb\x3e\x8a\xd8\xb7\x50\x42\xa8\xa7\x20\x07\x3e\x08\xea\x61\xe9\xb8\xf9\x70\x8f\x91\x2a\xf6\xd2\x1c\x18\x36\xed\x3d\xc1\x95\xc4\x4d\x67\x93\x35\x8a\x9a\x6c\xc2\xd8\x3e\xe7\x3e\x84\xc2\x6a\xde\x6a\xca\x78\xb1\x01\xc2\xac\xdc\x37\xea\xfd\x4f\xa4\xe3\xb1\x25\x07\x8f\xb4\x39\x0f\x95\x16\xa7\xdd\xcf\x65\x0a\x68\x56\x75\x45\xd2\x81\x53\x2d\xef\x16\x64\x9a\xc0\x47\x65\x80\x7e\x61\xa5\x05\xbd\x91\x3d\x85\x38\xeb\xb1\x50\x1f\xe2\x10\x9d\x65\x2e\x84\x23\x23\xd1\xf1\x35\x89\xcd\x29\x20\x7c\xbf\xf8\x3a\x83\x82\x0b\x8a\x3b\x64\x83\xa6\xbc\x35\x41\xd4\x1f\x0a\x65\x8e\x17\xae\xdd\x35\x6d\xd4\x9a\xe7\xde\x4d\x25\x17\xc4\x8e\x96\x24\x60\x7d\xa3\x62\xb6\xa3\xb7\x0e\xdc\x60\x25\x82\x80\xbf\x03\x00\x00\xff\xff\x70\x83\x16\x13\x55\x08\x00\x00")

func readmeTplBytes() ([]byte, error) {
	return bindataRead(
		_readmeTpl,
		"readme.tpl",
	)
}

func readmeTpl() (*asset, error) {
	bytes, err := readmeTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "readme.tpl", size: 2133, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"readme.tpl": readmeTpl,
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
	"readme.tpl": {readmeTpl, map[string]*bintree{}},
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
