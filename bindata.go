// Code generated by go-bindata.
// sources:
// templates/confirm-de.html
// templates/confirm-en.html
// templates/confirm.html
// templates/forgot-password.html
// templates/password-changed.html
// DO NOT EDIT!

package waechter

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

var _templatesConfirmDeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xbd\x6e\xdb\x30\x14\x85\x77\x3f\xc5\x41\x66\xd7\xaf\x50\x24\x6d\x87\xa2\x3f\x43\x8b\xc2\x33\x25\x1d\x49\x17\x24\x2f\x8b\xcb\xab\x00\xb5\xc0\xd7\xea\x94\xcd\x2f\x56\x58\x49\x8d\x64\xf0\xce\xef\xdc\x8f\xdf\xba\x62\xe0\x28\x4a\xdc\xb9\x78\xe2\x1d\x5a\xdb\x01\xc0\x51\x52\x8a\x25\x67\x2a\xd6\x15\x87\x5f\x95\xa6\x21\xb3\xb5\xfd\x6e\x5d\x41\x1d\x2e\x0f\x77\xaf\xf0\xbe\xe8\x28\x96\x1f\x16\xf7\xa2\xd7\x99\x4f\xef\xbe\x05\x49\x78\xa4\xc9\x28\x27\xa1\x51\x6f\xf3\x4e\xf5\x2b\x39\x04\x8d\xc4\x78\x7e\x32\x7c\x9e\x8d\xf8\xc1\x49\xaa\x9b\xd0\x16\x9d\xd0\x51\x36\xaf\x0f\x25\xff\x0e\xfa\xe7\xc8\xae\x8a\x13\xad\x1d\xf0\x20\xee\xc4\x58\xd2\x44\xc5\x4f\x21\x06\x66\x2c\x97\x71\x54\xe7\x4c\x1d\xa8\xf8\x2a\x1a\xf7\x58\xf2\x65\x1b\x5f\x8a\x7a\xc1\x68\x94\xd3\x52\xfb\x39\x24\xa7\x1e\x6e\x68\x8e\xa5\x38\xed\x6a\xf9\x5d\xfa\xd9\xdf\x14\x42\x6b\xef\x71\xa4\xbe\xdc\x16\x56\xfe\xaf\x20\x66\x7e\x7e\xca\x49\xfa\x19\xb4\xe7\x3b\x98\x43\x47\xdd\x23\x9e\xff\xaa\xf2\x35\x44\xd1\x31\xf4\x33\x64\xd2\x62\x5b\xb9\x03\x3e\xd2\x40\xab\xce\x94\x9c\xb8\xef\xfb\xb2\xa8\x43\xaa\x43\x37\x91\x8e\xba\xf8\xa9\x0b\xb6\x47\x2d\x29\xe8\x44\xd4\xad\x80\xe2\x3e\xba\x3c\x3e\xd7\xab\x49\x34\xbe\x10\x41\x63\x92\x3e\xbe\xf9\xf0\xbf\x00\x00\x00\xff\xff\x2c\x42\xfd\x61\x15\x02\x00\x00")

func templatesConfirmDeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesConfirmDeHtml,
		"templates/confirm-de.html",
	)
}

func templatesConfirmDeHtml() (*asset, error) {
	bytes, err := templatesConfirmDeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/confirm-de.html", size: 533, mode: os.FileMode(420), modTime: time.Unix(1497043907, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConfirmEnHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x3f\x4b\x03\x41\x10\xc5\xfb\x7c\x8a\x67\x04\x2b\x49\x63\x23\x58\x08\x0a\x82\x4d\x2a\xff\xd4\x9b\xbd\xd9\xdc\x90\xdd\x99\xb0\x3b\x67\x38\x96\xfd\xee\x92\x13\x8f\x53\xb4\x1c\x78\xbf\xc7\xef\x4d\xad\xe8\x28\xb0\x10\xd6\xc6\x16\x69\x8d\xd6\x56\x00\xf0\x4e\xd1\x6b\x22\xd4\x8a\xcd\x6b\xa1\x2c\x2e\x51\x6b\xd7\xab\x5a\x41\xd2\x9d\x53\xab\x05\xeb\x55\x02\xe7\xf4\x30\x98\xa9\xcc\x1d\x6f\x94\x39\x8c\xa0\xe4\x38\x2e\xc1\x5f\xa4\x91\xd8\xcc\x58\xef\xe4\x50\x10\x34\x23\xd3\x9e\x8b\x51\x66\xd9\xc3\xd9\x64\xf2\xa8\xe9\xe8\x64\xdc\xba\x44\x68\xed\x02\x4f\x1a\xa3\x9e\x60\x3d\x21\xb2\x1c\xb0\xa3\xe9\x54\x38\x6f\xfc\xe1\x8c\x30\xea\x90\xe1\xbc\xd7\x41\x6c\xf3\x8f\x7d\x50\x35\xca\xb3\xc2\x56\xed\xc7\x6c\xb4\x76\x8f\xe7\x70\xae\x42\xc7\x9d\x5c\x5d\xde\xdc\xde\xcd\x76\x7f\xab\x4d\x61\xef\x04\x85\xd3\x31\x8e\xe0\xbd\x68\x26\x58\xcf\xe5\xeb\x1f\x1b\xbc\xf4\xf4\x2d\x86\x13\xc7\x08\x51\xc3\x8e\x30\x14\xb7\x8b\x84\x41\x22\x95\x32\xf5\xcc\x63\x78\x39\xe1\x33\x00\x00\xff\xff\xe6\x8c\x65\x9f\xbb\x01\x00\x00")

func templatesConfirmEnHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesConfirmEnHtml,
		"templates/confirm-en.html",
	)
}

func templatesConfirmEnHtml() (*asset, error) {
	bytes, err := templatesConfirmEnHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/confirm-en.html", size: 443, mode: os.FileMode(420), modTime: time.Unix(1497043908, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConfirmHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xeb\x8e\x1b\xb7\x15\xfe\xbf\x4f\x71\x96\x46\x81\xa4\xd0\x68\x46\xeb\xc4\xc9\xce\x45\xb5\xe3\xba\x8d\x81\x5e\x8c\xd6\x06\x0a\x04\x46\x41\x0d\xcf\x48\xcc\x72\xc8\x29\xc9\xd1\xa5\xea\xbe\x4b\x9f\xa5\x4f\x56\x90\x9c\x91\x46\xb7\x8d\xed\x78\x81\xfc\x88\x7f\x58\xcb\xdb\xb9\x7c\xe7\xe3\x77\x38\xf9\x35\x53\xa5\xdd\x34\x08\x0b\x5b\x8b\xe9\x55\xee\x7e\x60\x5d\x0b\x69\x0a\xb2\xb0\xb6\x49\xe3\x78\xb5\x5a\x8d\x57\x4f\xc7\x4a\xcf\xe3\xc9\xed\xed\x6d\xbc\x76\x7b\x48\xd8\x94\x2e\x0b\xd2\x6a\x99\x9a\x72\x81\x35\x35\x51\xcd\x4b\xad\x8c\xaa\x6c\x54\xaa\x3a\x5d\xee\xf7\xa9\x87\xf6\xa9\xaa\xe2\x25\x76\x3f\xc4\x45\x81\x94\x4d\xaf\x00\x72\xcb\xad\xc0\x69\x1e\x87\x5f\x37\x73\x1d\x45\x3f\xf0\x0a\xae\x6b\xa3\xde\x4f\xdd\x08\xa2\xc8\x2f\xd4\x68\x29\xb8\x90\x23\xfc\x57\xcb\x97\x05\xf9\x47\xf4\xee\x45\xf4\x52\xd5\x0d\xb5\x7c\x26\x90\x40\xa9\xa4\x45\x69\x0b\xf2\xfa\x55\x81\x6c\xee\x3c\x05\x83\xf9\xf5\x0f\x28\x19\xaf\xde\x3b\x53\xa7\x86\x5e\x86\x73\xd1\xdb\x4d\x33\xb4\x62\x71\x6d\x63\x87\x45\x06\xe5\x82\x6a\x83\xb6\x78\xf7\xf6\x0f\xd1\xb7\xa4\xb7\x21\x69\x8d\x05\x59\x72\x5c\x35\x4a\xdb\xc1\xc9\x15\x67\x76\x51\x30\x5c\xf2\x12\x23\x3f\x18\x01\x97\xdc\x72\x2a\x22\x53\x52\x81\xc5\x64\x9c\x38\x33\xc6\x6e\x04\x82\xab\x4f\xe7\xae\x34\xc6\x47\xfd\x44\xb5\x56\x28\x75\x07\x14\xb6\xd0\x50\xc6\xb8\x9c\xa7\x90\x64\x70\x7f\x05\x30\xfe\x1b\x52\xf6\x67\x33\xff\x4e\xb1\x0d\x6c\xc1\x3b\x48\x61\x92\x24\xbf\xe9\xd6\x5f\xad\x2d\x6a\x49\xc5\x4b\x41\x8d\xf9\x80\x1d\xbf\x85\x2d\x08\x2e\x31\x5a\x20\x9f\x2f\x6c\xba\xdf\x38\x0b\x2e\x6a\xaa\xe7\x5c\xfa\x00\x86\xc1\x44\x2b\x9c\xdd\x71\x1b\xb9\xd0\x23\xc3\xff\x8d\x11\x65\x3f\xb6\xc6\xf6\xae\xa2\xda\x5c\x5c\x73\xd6\x2d\x9d\x09\x1c\x81\x65\xb0\x85\x99\xd2\x0c\x75\x54\x2a\x21\x68\x63\x30\xed\xff\xc8\xa0\x36\x2a\xf2\x3b\x23\x61\x1a\x5a\x62\x0a\x49\x63\x87\xd3\x7a\x38\xed\xec\xf2\x7a\xbe\xb3\xe8\x03\xed\xf2\x02\xda\x5a\x95\x1d\xa4\xda\x45\xe3\xe0\xe6\x12\x53\x90\x4a\x62\x06\x3e\x68\x86\xa5\xd2\xd4\x72\x25\xfb\x69\x97\x0f\x97\x16\x75\xa3\x84\x5f\x88\x6a\xc5\x30\x85\x19\x2f\xdb\x19\x2f\x83\xf3\x06\xb6\xc0\xb8\x69\x04\xdd\xa4\x30\x13\xaa\xbc\xcb\x76\xf8\x4d\x9e\x36\xeb\x50\xc5\x3c\xf6\xb5\x9f\x5e\x9d\x12\xfe\x21\x5e\x3c\xaf\x91\x71\x0a\x4a\x8a\x0d\x98\x52\x23\x4a\xa0\x92\xc1\x17\x35\x5d\x07\xa2\xa5\x5f\x7d\x9b\x34\xeb\x2f\x61\x7b\x05\x00\xf0\xdc\x45\xdc\x13\x74\x47\x84\xa7\x37\x49\xb3\x0e\xc1\x02\x3c\x7f\x68\xf9\x28\xce\xc3\x7b\xd4\x05\xee\xe3\xbe\xca\xd7\x4e\x5e\x00\x72\x95\xfe\xd5\x5f\xf3\xdf\xab\xb2\xad\x51\xda\xbf\xa3\xb5\x5c\xce\xcd\xd4\x7b\xcb\x55\xfa\x42\x08\xb5\x7a\xf3\x97\x3f\xc6\xbb\x99\x37\x7c\x8d\xc2\xbc\x41\xfd\x5a\x96\x8b\xe9\xed\xb3\x3c\x3e\x9e\x73\x86\xe3\xcb\x96\xf3\xd8\x7b\x3f\x1b\x9e\xb0\xe8\x42\x84\xc9\xe4\xfd\x43\xc0\x8e\xbb\x0b\x17\xcd\xb5\x6a\x9b\xa8\xe2\xeb\x0e\xc2\x80\x89\x63\x09\x5c\xf3\xda\x01\x45\xa5\xcd\x8e\xb1\x19\x38\x3e\x5b\x51\x9f\xa9\xe0\xf2\x0e\x16\x1a\xab\x20\xbd\x26\x8d\xe3\x4a\x49\x6b\xc6\x73\xa5\xe6\x02\x69\xc3\xcd\xb8\x54\xb5\x8b\xe9\x77\x15\xad\xb9\xd8\x14\xef\x66\xad\xb4\x6d\xfa\x34\x49\x46\x5f\x25\xc9\xe8\xeb\x24\x19\x7d\x93\x24\x04\x34\x8a\x82\x78\xef\x66\x81\x68\xc9\xb9\x94\x00\x2e\xa5\xeb\x17\x7d\xf1\x43\x46\xd0\x6a\xf1\xc5\xcf\x0b\xe9\xcb\x2c\x58\xdd\x61\x72\xaa\xbc\x9f\xc0\x6a\x2e\xcf\xb1\x7a\x5c\xff\xe8\x74\xa2\xad\x65\xd4\xa0\x8e\x26\x49\xb2\xe3\xae\xab\xd3\xa0\x4c\x27\x24\x8e\x43\xe7\xc9\xbd\xae\xf9\xc9\x82\xcc\x68\x79\xe7\xaa\x2e\x59\x0a\x4f\xaa\xaa\xca\x42\xdf\x60\x7c\x79\xba\xc3\xf9\x55\x3a\xed\xb6\x0d\xae\x00\xfc\x07\x5e\xbf\x7a\x3f\xed\x90\xcd\xbd\x30\x81\x56\xee\x74\xa3\xd1\xa0\xb4\x5e\x30\x48\xa7\x4a\x05\x49\x08\x94\x28\x44\x27\xa7\xbb\xb1\x93\xb2\x7e\x1c\xda\x08\x79\xe6\x0a\x4e\x05\x9f\xcb\x82\x94\xe8\xe4\x87\xf4\x81\x85\xa4\x9f\x25\xee\xc2\x92\xe9\xae\xac\xb9\xd5\xfb\x81\x1b\xb2\xfe\xc0\x50\xfa\xdc\x21\x57\x6c\xaf\xce\x7e\xe4\x24\x75\xb0\x23\xd2\xad\xc0\x14\xd7\xb4\xb4\x62\xb3\xb7\x7f\x50\xd3\x01\x4a\x9d\xc8\x25\xcd\x3a\xe8\xec\x5e\x93\xfa\x00\x1f\x80\xe5\xa7\xb0\xe8\x7c\x1c\xc6\xbb\xaf\x79\x76\x82\xd0\x1e\xe7\x69\x6e\x5d\xb9\xa7\x0e\x95\x01\x14\x5e\xe3\xfd\xa1\x34\x9c\xc9\x96\xa8\x2d\x2f\xa9\xe8\x66\xad\x6a\x32\xc6\x35\x96\xbe\x05\x08\xab\x8f\xd0\xea\x3b\xa1\x93\x4b\xe8\x12\x7c\x54\x42\x1c\x67\x70\x26\xde\xb3\x84\x38\x29\x58\xe9\xba\x7e\x41\x4e\xaf\xd1\x89\x08\x92\x07\x7c\xf5\x4d\x8e\x4b\xcf\x99\xd0\xea\x2e\x01\xe6\x3a\x5f\x36\x80\x5c\x60\x65\x0f\xea\xf7\x73\xc8\xd1\x5d\x14\x67\xe8\x03\x0a\xbf\x52\x9a\x45\x2b\x4d\x9b\x74\xa6\x91\xde\x45\x6e\x7c\xa1\xb4\x13\x57\xda\x9b\xaf\x1d\x98\x47\xfc\xfa\x0c\x5c\xbe\xf8\xe2\xe9\x16\xba\x23\x3e\xa0\x8f\xa7\x77\x87\xed\x4d\x47\x4c\xf7\x24\xa2\xc2\x16\x84\x80\x7f\x6b\xbb\x3f\xc2\x2d\x2f\x88\xbb\xad\x04\x8c\x2e\x0b\xb2\xdd\xc2\xd8\x3f\xa9\xe5\xe6\x4f\x6a\xae\xde\x69\x01\xf7\xf7\x47\x11\xa7\xfe\x31\xd4\x05\xa9\x29\xe3\xad\xf1\x31\xf6\x7c\x08\x44\xe8\x5f\x54\x7e\xf3\xf1\x7b\xca\x4f\x0e\xaa\xdf\x29\x92\x97\x8d\x7d\x39\x6f\x5c\x72\xb1\x65\xee\x3f\xed\xfe\x0b\x79\xc6\x1e\xf9\xe1\xca\xe7\xa8\xef\x34\x6f\x4e\x45\x66\xe2\x44\xf1\x48\xd6\xba\xc4\xad\x6a\xdc\x32\x18\x25\x38\x83\x27\x8c\xb1\x23\x3a\xc7\xcd\x19\x39\x78\x80\x35\x17\x0b\xfc\xf1\xda\xf8\xe9\x61\x0f\x7b\xce\x31\xae\x7d\xdb\xc8\x0e\x9a\x48\x46\xa6\xff\xfb\xef\x41\x95\x42\x75\x86\x9a\xf3\x79\x2b\xd5\x23\xe5\x24\x84\x1c\x08\xda\x0e\x8a\xb2\xd5\x46\xe9\xc0\xa7\xae\x63\x53\x4a\x83\xe9\xf0\x8c\x49\x17\x28\x96\xe8\x24\x6d\xe0\xd0\x5f\x96\x61\x76\x37\x37\x67\x54\x8b\x4c\xb7\x5b\x00\x8b\x75\x23\xa8\x45\x20\xfe\x3e\x11\xb8\xbf\xcf\x63\xc6\x97\xbf\x8c\x6c\x13\xff\xef\x20\xe1\xf0\x6e\x1b\xc1\xf7\x7d\xe2\x23\x78\xa1\x39\x15\x23\x30\x54\x9a\xc8\xa0\xe6\xd5\x50\xb0\x6f\x3e\x18\x8b\x3d\x14\xdd\x17\xf0\xa3\x83\xf1\x78\x22\x6c\xb0\xa1\x9a\x5a\xfc\x04\xc5\xbd\xac\x8e\xae\xf7\x85\xc2\xac\x16\xdc\x62\x36\x2c\xd8\x4f\x67\x09\xcb\x6e\x5c\x73\xc6\x1c\xd3\x66\x73\x6f\xac\x20\x4f\x26\xb7\x74\x72\xfb\x0d\x99\xe6\xb4\xfb\xae\x08\xfa\x2d\x2b\xae\xeb\xef\x5a\x6b\x95\x7c\x4b\xf5\x1c\xed\x50\xc4\xcf\x6a\xf1\xc9\xd7\xff\xe0\x51\xdc\x79\x39\xc8\x60\x48\xac\xef\xcf\xdc\x24\xdf\xee\xfd\x70\x15\x8c\x4a\xa5\x6b\x2a\x02\x7f\xac\xa6\xd2\x54\x4a\xd7\xc1\xf7\x5e\xa9\x32\x02\xd6\xc7\x5b\x90\x7f\xce\x04\x95\x77\xa7\xf4\xda\x67\x16\x48\x46\x7f\xed\x10\xbf\x76\x88\xc7\xe9\x10\x1f\x2b\x98\x1f\xde\x3c\xf6\x8c\xae\x94\xf2\xb5\x3c\xa3\x97\x27\x64\xf6\xeb\x17\x3f\x30\x4e\xf0\x3d\xf7\xf8\x7f\x64\xeb\xce\xc6\x55\x1e\x7b\xcb\xee\x5b\xdb\xd6\x62\xfa\xff\x00\x00\x00\xff\xff\x18\xc2\x7d\x3d\x82\x16\x00\x00")

func templatesConfirmHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesConfirmHtml,
		"templates/confirm.html",
	)
}

func templatesConfirmHtml() (*asset, error) {
	bytes, err := templatesConfirmHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/confirm.html", size: 5762, mode: os.FileMode(420), modTime: time.Unix(1497044014, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesForgotPasswordHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x6d\x93\xdb\xb6\x11\xfe\x7e\xbf\x62\x4d\x4f\xa7\x49\xe7\x28\x52\xe7\xc4\xc9\x51\xa4\x6a\xd7\x75\x5b\x7f\x68\xea\x69\xed\x99\x66\x32\xfe\x00\x01\x2b\x12\x3e\x10\x60\x80\xa5\x5e\xaa\xde\x7f\xe9\x6f\xe9\x2f\xeb\x00\xa4\x24\xea\xe5\x2e\x8e\xed\xeb\x74\x3a\xb9\x0f\x27\x81\x00\x16\xbb\xcf\x3e\xfb\x2c\xa8\xfc\x91\x30\x9c\xd6\x0d\x42\x45\xb5\x9a\x5e\xe4\xfe\x03\x56\xb5\xd2\xae\x88\x2a\xa2\x26\x4b\x92\xe5\x72\x39\x5a\x3e\x19\x19\x5b\x26\xe3\xeb\xeb\xeb\x64\xe5\xd7\x44\xdd\xa2\x6c\x51\x44\xad\xd5\x99\xe3\x15\xd6\xcc\xc5\xb5\xe4\xd6\x38\x33\xa7\x98\x9b\x3a\x5b\xec\xd7\x99\xfb\xd6\x99\xf9\x5c\x72\xec\x3f\x22\xef\x05\x32\x31\xbd\x00\xc8\x49\x92\xc2\x69\x9e\x74\x9f\xfe\xc9\xa3\x38\xfe\x41\xce\xe1\x51\xed\xcc\xbb\xa9\x1f\x41\x1c\x87\x89\x1a\x89\x81\x77\x39\xc6\x1f\x5b\xb9\x28\xa2\xbf\xc7\x6f\x9f\xc7\x2f\x4c\xdd\x30\x92\x33\x85\x11\x70\xa3\x09\x35\x15\xd1\xab\x97\x05\x8a\xd2\x9f\xd4\x19\xcc\x1f\xfd\x80\x5a\xc8\xf9\x3b\x6f\xea\xd4\xd0\x8b\x6e\x5f\xfc\x66\xdd\x0c\xad\x10\xae\x28\xf1\x58\x4c\x80\x57\xcc\x3a\xa4\xe2\xed\x9b\x3f\xc4\xdf\x46\x5b\x1b\x9a\xd5\x58\x44\x0b\x89\xcb\xc6\x58\x1a\xec\x5c\x4a\x41\x55\x21\x70\x21\x39\xc6\x61\x70\x09\x52\x4b\x92\x4c\xc5\x8e\x33\x85\xc5\x78\x94\x7a\x33\x8e\xd6\x0a\xc1\xe7\xa7\x3f\x8e\x3b\x17\xbc\x7e\x6c\x5a\x52\xc6\xdc\x00\x83\x0d\x34\x4c\x08\xa9\xcb\x0c\xd2\x09\xdc\x5e\x00\x8c\xfe\x8a\x4c\xfc\xd9\x95\xbf\x33\x62\x0d\x1b\x08\x07\x64\x30\x4e\xd3\x5f\xf5\xf3\x2f\x57\x84\x56\x33\xf5\x42\x31\xe7\x3e\x60\xc5\x6f\x60\x03\x4a\x6a\x8c\x2b\x94\x65\x45\xd9\x7e\xe1\xac\x3b\xa2\x66\xb6\x94\x3a\x38\x30\x74\x26\x5e\xe2\xec\x46\x52\xec\x5d\x8f\x9d\xfc\x07\xc6\x4c\xbc\x6f\x1d\x6d\x8f\x8a\x6b\x77\xe7\x9c\xb7\x4e\x6c\xa6\xf0\x12\x48\xc0\x06\x66\xc6\x0a\xb4\x31\x37\x4a\xb1\xc6\x61\xb6\xfd\x32\x81\xda\x99\x38\xac\x8c\x95\x6b\x18\xc7\x0c\xd2\x86\x86\x8f\xed\xf0\xb1\xb7\x2b\xeb\x72\x67\x31\x38\xda\xc7\x05\xac\x25\x33\x39\x08\xb5\xf7\xc6\xc3\x2d\x35\x66\xa0\x8d\xc6\x09\x04\xa7\x05\x72\x63\x19\x49\xa3\xb7\x8f\x7d\x3c\x52\x13\xda\xc6\xa8\x30\x11\xd7\x46\x60\x06\x33\xc9\xdb\x99\xe4\xdd\xe1\x0d\x6c\x40\x48\xd7\x28\xb6\xce\x60\xa6\x0c\xbf\x99\xec\xf0\x1b\x3f\x69\x56\x5d\x16\xf3\x24\xe4\x7e\x7a\x71\x4a\xf8\xfb\x78\xf1\xac\x46\x21\x19\x18\xad\xd6\xe0\xb8\x45\xd4\xc0\xb4\x80\x2f\x6a\xb6\xea\x88\x96\x7d\xf5\x6d\xda\xac\xbe\x84\xcd\x05\x00\xc0\x33\xef\xf1\x96\xa0\x3b\x22\x3c\xb9\x4a\x9b\x55\xe7\x2c\xc0\xb3\xfb\xa6\x8f\xfc\x3c\xac\xa3\xde\xf1\xe0\xf7\x45\xbe\xf2\xf2\x02\x90\x9b\xec\x2f\xa1\xcc\x7f\x6f\x78\x5b\xa3\xa6\xbf\x21\x91\xd4\xa5\x9b\x86\xd3\x72\x93\x3d\x57\xca\x2c\x5f\x7f\xf7\xc7\x64\xf7\xe4\xb5\x5c\xa1\x72\xaf\xd1\xbe\xd2\xbc\x9a\x5e\x3f\xcd\x93\xe3\x67\xde\x70\x72\xb7\xe5\x3c\x09\xa7\x9f\x75\x4f\x11\x7a\x17\x61\x3c\x7e\x77\x1f\xb0\xa3\xbe\xe0\xe2\xd2\x9a\xb6\x89\xe7\x72\xd5\x43\xd8\x61\xe2\x59\x02\x8f\x64\xed\x81\x62\x9a\x26\xc7\xd8\x0c\x0e\x3e\x9b\xd1\x10\xa9\x92\xfa\x06\x2a\x8b\xf3\x4e\x7a\x5d\x96\x24\x73\xa3\xc9\x8d\x4a\x63\x4a\x85\xac\x91\x6e\xc4\x4d\xed\x7d\xfa\xed\x9c\xd5\x52\xad\x8b\xb7\xb3\x56\x53\x9b\x3d\x49\xd3\xcb\xaf\xd2\xf4\xf2\xeb\x34\xbd\xfc\x26\x4d\x23\xb0\xa8\x8a\x28\x9c\xee\x2a\x44\x8a\xce\x85\x04\x70\x57\xb8\x61\x32\x24\xbf\x8b\x08\x5a\xab\xbe\xf8\x34\x97\xbe\x9c\x74\x56\x77\x98\x9c\x2a\xef\x47\xb0\x5a\xea\x73\xac\x1e\xd5\xef\xbd\x4e\xb4\xb5\x8e\x1b\xb4\xf1\x38\x4d\x77\xdc\xf5\x79\x1a\xa4\xe9\x84\xc4\x49\xd7\x79\x72\xaf\x6b\xc1\x47\x21\x17\xd3\x01\x93\xe1\x9f\xf0\xea\xe5\xbb\x69\x0f\x50\x1e\xf4\x05\xac\x51\x58\x44\x8d\x45\x87\x9a\x42\xdd\x47\xbd\xb8\x14\x51\x1a\x01\x47\xa5\x7a\x55\xdc\x8d\xbd\x22\x6d\xc7\x5d\x37\x88\x9e\xfa\xbc\x31\x25\x4b\x5d\x44\x1c\xbd\x8a\x44\x10\xdc\xea\xfb\x45\xf6\x34\xf5\x75\x17\x4d\x77\xd9\xc9\xc9\xee\x07\x7e\x28\xb6\x1b\x86\x0a\xe6\x37\xf9\x9c\x05\x91\x0d\x23\xaf\x8c\x83\x15\xb1\x6d\x15\x66\xb8\x62\x9c\xd4\x7a\x6f\xff\x20\x35\x42\x2e\xb6\xb6\x7b\xad\x4a\x9b\x55\x27\x97\x7b\x69\xd9\x3a\x78\x0f\x2c\x3f\x85\x45\x7f\xc6\xa1\xbf\xfb\xd4\x4d\x4e\x10\xda\xe3\x3c\xcd\x29\x64\xcd\xa3\x32\x80\x22\x48\x75\xd8\x94\x75\x7b\x26\x0b\xb4\x24\x39\x53\xfd\x53\x32\xcd\x44\x48\x8b\x3c\x28\xb9\x22\x7b\x84\xd6\xb6\xa1\x79\xd5\x83\x3e\xc0\x07\x25\xc4\x71\x04\x67\xfc\x3d\x4b\x88\x93\x84\x71\xdf\xbc\x8b\xe8\xb4\x1a\x4e\xb4\x2c\xba\xe7\xac\x6d\xaf\x92\x3a\x70\xa6\xeb\x58\x77\x01\xe6\x1b\xd8\x64\x00\xb9\xc2\x39\x1d\xe4\xef\x53\xc8\xd1\x17\x8a\x37\xf4\x01\x89\x5f\x1a\x2b\xe2\xa5\x65\x4d\x36\xb3\xc8\x6e\x62\x3f\xbe\x23\xb5\x63\x9f\xda\xab\xaf\x3d\x98\x47\xfc\xfa\x0c\x5c\xbe\xf3\xe2\xd2\x4f\xf4\x5b\x82\x43\x3f\x9f\xde\x3d\xb6\x57\x3d\x31\xfd\xcd\x86\x29\x2a\xa2\x08\xc2\x95\xd9\x7f\xe9\xaa\xbc\x88\x7c\xb5\x46\xe0\x2c\xdf\xb7\x17\x6e\xc4\xba\x32\xad\xc3\x11\x37\x89\xc0\xda\x24\x4c\x2c\x98\xe6\x28\x62\x87\xcc\xf2\x2a\x9e\x1b\x5b\x27\xb2\x2e\x13\x2e\x62\x65\x4a\x33\x72\x8b\xf2\x28\xb4\x2c\x5c\x7e\xfa\x68\x2c\x13\xb2\x75\x21\x98\x2d\x71\x3a\xc6\x6c\x6f\x50\x61\xf1\xf1\xfd\x29\x3c\x1c\xd0\xa4\x97\xae\xa0\x2f\xfb\xbc\x5f\x79\x14\x12\x12\xfe\x9f\xf5\xff\x3a\x40\x92\x90\xa2\xe1\xcc\xe7\x20\xc2\x34\x6f\x4e\xd5\x68\xec\xd5\xf3\x48\xff\xfa\xc0\xc9\x34\x7e\x1a\x9c\x51\x52\xc0\x63\x21\xc4\x11\xef\x93\xe6\x8c\x6e\xdc\x43\xaf\x3b\x99\xf0\xf3\x45\xf4\xe3\xdd\x1e\x36\xa7\x63\x5c\xb7\xfd\x65\x72\xd0\x6d\x26\xd1\xf4\xdf\xff\x3a\xc8\x52\x97\x9d\xa1\x38\x7d\xde\x4c\x6d\x91\xf2\x5a\x13\x1d\x28\xdf\x0e\x0a\xde\x5a\x67\x6c\xc7\x27\x6e\x94\xb1\xd9\xe3\x34\xfc\x75\xd6\xbb\x9b\x4b\xd6\xdd\x5c\x2e\xe1\x4f\xa8\x16\xe8\x65\xf0\x12\x9e\x5b\xc9\xd4\x25\x38\xa6\x5d\xec\xd0\xca\xf9\x50\xeb\xae\x9a\xd5\x41\xec\x57\x57\x67\xc4\x2f\x9a\x7e\x6f\x5a\x20\xa3\x04\xb4\x0e\xd6\xa6\x85\xb9\xb1\xa5\x21\xff\xd5\x42\xc3\x9c\xf3\x51\x8e\xe0\xd5\x3c\x4c\x5a\x64\x4a\xad\x41\x48\x11\x86\x9c\x69\xff\x72\xa9\x4b\x04\x49\xd0\x3a\xa9\x4b\xa0\x0a\x21\x5c\x15\x67\xa8\xcc\x32\xcb\x93\x70\x47\x79\x18\x4c\x1f\x4e\x06\x1d\x36\xcc\x32\xc2\x8f\xd0\xbc\xbb\x65\xc7\x77\x9f\x2e\xbf\xcb\x4a\x12\x4e\x86\x79\xff\xe9\x28\x61\xd1\x8f\x6b\x29\x84\xc2\x08\x66\x65\x30\x56\x44\x8f\xc7\xd7\x6c\x7c\xfd\xcd\x50\x14\xce\x4a\xd8\xc9\x4b\xf2\x8c\xf1\x1b\xdf\x65\xb5\xc8\x7a\x1b\x07\xfe\x0d\xd9\xb7\x63\xdd\x71\x3b\x0d\xc3\x65\x67\x54\x1b\x5b\x33\xd5\x91\x8c\x2c\xd3\xce\xeb\x73\x77\xf6\xbe\xc0\x27\xd1\xf4\x85\xd1\x73\x69\x6b\xc0\x9a\x49\x05\x4c\x08\x8b\xce\x05\x09\xfa\x6f\xe9\xe7\xff\x7a\x55\xf6\xf5\x26\xa4\xd0\xbf\x26\xb0\xf8\x63\x8b\x8e\x80\x2a\xe9\x7a\xd0\xde\xb7\x8e\x40\x96\xda\x58\x5f\x7b\x23\xf8\x7e\x58\xaf\xb0\x34\x7e\x5b\x57\x9a\xa3\x07\x29\xc1\x5f\x1a\xd0\xff\x77\x03\x62\x8c\x7d\x12\xcf\xd3\x0f\xe3\xf9\x9b\x3d\xa3\x97\xcc\x05\xe8\x6b\x16\xee\xf8\xfe\x65\x1a\x35\x01\x19\xd8\x6c\x60\xf4\x32\xac\xb9\xbd\xdd\xb5\xa2\x5d\x69\x94\xd2\x11\x5a\x60\x04\x9b\xcd\x28\xfc\xa0\xaa\xd7\xdf\xb1\x1a\xe1\xf6\xf6\xa0\x4a\x42\xf1\xd4\xe8\x1c\x2b\x71\x04\x6f\x2a\x04\xc6\xb9\x69\x35\x85\x36\xa6\x0d\xc1\x0c\xa1\x75\x28\x60\x29\xa9\x32\x2d\x85\x56\xc6\x3b\xa9\x0a\xa4\x3b\xad\xa4\x13\x91\xba\xff\x85\xfc\x24\xed\xe7\xde\x8d\x1e\xd8\xba\xb7\x71\x91\x27\xdd\x4f\x08\x79\xf8\x85\x78\xfa\x9f\x00\x00\x00\xff\xff\xf3\x11\xe2\x7e\x68\x17\x00\x00")

func templatesForgotPasswordHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesForgotPasswordHtml,
		"templates/forgot-password.html",
	)
}

func templatesForgotPasswordHtml() (*asset, error) {
	bytes, err := templatesForgotPasswordHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/forgot-password.html", size: 5992, mode: os.FileMode(420), modTime: time.Unix(1497044014, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPasswordChangedHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5b\x93\xdc\x38\x15\x7e\x9f\x5f\x71\xc6\x29\x8a\x5d\x6a\xd4\x76\x4f\x76\xc3\xc6\x6d\x37\x09\x21\x40\x1e\x58\x52\x90\x54\x41\x6d\xcd\x83\x5a\x3a\xb6\xb5\x23\x4b\x46\x92\xfb\x42\x33\xff\x85\xdf\xc2\x2f\xa3\x24\xdb\xdd\xee\xcb\x34\xd9\x5c\x5e\x80\x79\xe8\x6e\x5d\x7c\x2e\xdf\xf9\xce\x27\x79\xb2\x6b\xae\x99\xdb\x34\x08\x95\xab\xe5\xfc\x2a\xf3\x5f\xb0\xae\xa5\xb2\x79\x54\x39\xd7\xa4\x71\xbc\x5a\xad\x26\xab\xa7\x13\x6d\xca\x78\xfa\xfc\xf9\xf3\x78\xed\xf7\x44\xdd\xa6\x74\x99\x47\xad\x51\xa9\x65\x15\xd6\xd4\x92\x5a\x30\xa3\xad\x2e\x1c\x61\xba\x4e\x97\xfb\x7d\xfa\xd2\x3e\x5d\x14\x82\x61\xff\x15\xf9\x28\x90\xf2\xf9\x15\x40\xe6\x84\x93\x38\xcf\xe2\xee\xdb\xcf\x5c\x13\xf2\x83\x28\xe0\xba\xb6\xfa\x6e\xee\x47\x40\x48\x58\xa8\xd1\x51\xf0\x21\x13\xfc\x5b\x2b\x96\x79\xf4\x17\xf2\xfe\x25\x79\xa5\xeb\x86\x3a\xb1\x90\x18\x01\xd3\xca\xa1\x72\x79\xf4\xe6\x75\x8e\xbc\xf4\x9e\x3a\x83\xd9\xf5\x0f\xa8\xb8\x28\xee\xbc\xa9\x53\x43\xaf\xba\xe7\xc8\xbb\x4d\x33\xb6\xe2\x70\xed\x62\x8f\xc5\x0c\x58\x45\x8d\x45\x97\xbf\x7f\xf7\x5b\xf2\x5d\x34\xd8\x50\xb4\xc6\x3c\x5a\x0a\x5c\x35\xda\xb8\xd1\x93\x2b\xc1\x5d\x95\x73\x5c\x0a\x86\x24\x0c\x6e\x40\x28\xe1\x04\x95\xc4\x32\x2a\x31\x9f\x4e\x12\x6f\xc6\xba\x8d\x44\xf0\xf5\xe9\xdd\x31\x6b\x43\xd4\x4f\x74\xeb\xa4\xd6\xf7\x40\x61\x0b\x0d\xe5\x5c\xa8\x32\x85\x64\x06\x0f\x57\x00\x93\x3f\x21\xe5\x7f\xb0\xe5\xaf\x35\xdf\xc0\x16\x82\x83\x14\xa6\x49\xf2\xb3\x7e\xfd\xf5\xda\xa1\x51\x54\xbe\x92\xd4\xda\x0f\xd8\xf1\x0b\xd8\x82\x14\x0a\x49\x85\xa2\xac\x5c\xba\xdf\xb8\xe8\x5c\xd4\xd4\x94\x42\x85\x00\xc6\xc1\x90\x15\x2e\xee\x85\x23\x3e\x74\x62\xc5\xdf\x91\x50\xfe\x63\x6b\xdd\xe0\x8a\xd4\xf6\xd1\x35\x6f\xdd\xd1\x85\xc4\x1b\x70\x1c\xb6\xb0\xd0\x86\xa3\x21\x4c\x4b\x49\x1b\x8b\xe9\xf0\x63\x06\xb5\xd5\x24\xec\x24\xd2\x36\x94\x61\x0a\x49\xe3\xc6\xd3\x66\x3c\xed\xed\x8a\xba\xdc\x59\x0c\x81\xf6\x79\x01\x6d\x9d\x9e\x1d\xa4\xda\x47\xe3\xe1\x16\x0a\x53\x50\x5a\xe1\x0c\x42\xd0\x1c\x99\x36\xd4\x09\xad\x86\x69\x9f\x8f\x50\x0e\x4d\xa3\x65\x58\x20\xb5\xe6\x98\xc2\x42\xb0\x76\x21\x58\xe7\xbc\x81\x2d\x70\x61\x1b\x49\x37\x29\x2c\xa4\x66\xf7\xb3\x1d\x7e\xd3\xa7\xcd\xba\xab\x62\x16\x87\xda\xcf\xaf\x4e\x09\x7f\x89\x17\x2f\x6a\xe4\x82\x82\x56\x72\x03\x96\x19\x44\x05\x54\x71\xf8\xaa\xa6\xeb\x8e\x68\xe9\x37\xdf\x25\xcd\xfa\x6b\xd8\x5e\x01\x00\xbc\xf0\x11\x0f\x04\xdd\x11\xe1\xe9\x6d\xd2\xac\xbb\x60\x01\x5e\x5c\x5a\x3e\x8a\xf3\xb0\x8f\xfa\xc0\x43\xdc\x57\xd9\xda\xcb\x0b\x40\xa6\xd3\x3f\x86\x36\xff\x8d\x66\x6d\x8d\xca\xfd\x19\x9d\x13\xaa\xb4\xf3\xe0\x2d\xd3\xe9\x4b\x29\xf5\xea\xed\xf7\xbf\x8b\x77\x33\x6f\xc5\x1a\xa5\x7d\x8b\xe6\x8d\x62\xd5\xfc\xf9\xb3\x2c\x3e\x9e\xf3\x86\xe3\xc7\x2d\x67\x71\xf0\x7e\x36\x3c\xe9\xd0\x87\x08\xd3\xe9\xdd\x25\x60\x27\x7d\xc3\x91\xd2\xe8\xb6\x21\x85\x58\xf7\x10\x76\x98\x78\x96\xc0\xb5\xa8\x3d\x50\x54\xb9\xd9\x31\x36\x23\xc7\x67\x2b\x1a\x32\x95\x42\xdd\x43\x65\xb0\xe8\xa4\xd7\xa6\x71\x5c\x68\xe5\xec\xa4\xd4\xba\x94\x48\x1b\x61\x27\x4c\xd7\x3e\xa6\x5f\x15\xb4\x16\x72\x93\xbf\x5f\xb4\xca\xb5\xe9\xd3\x24\xb9\xf9\x26\x49\x6e\xbe\x4d\x92\x9b\x5f\x26\x49\x04\x06\x65\x1e\x05\xef\xb6\x42\x74\xd1\xb9\x94\x00\x1e\x4b\x37\x2c\x86\xe2\x77\x19\x41\x6b\xe4\x57\x9f\x16\xd2\xd7\xb3\xce\xea\x0e\x93\x53\xe5\xfd\x08\x56\x0b\x75\x8e\xd5\x93\xfa\x47\xaf\x13\x6d\xad\x48\x83\x86\x4c\x93\x64\xc7\x5d\x5f\xa7\x51\x99\x4e\x48\x1c\x77\x27\x4f\xe6\x75\x2d\xc4\xc8\xc5\x72\x3e\x62\x32\xfc\x03\xde\xbc\xbe\x9b\xf7\x00\x65\x41\x5f\xc0\x68\x89\x79\xd4\x18\xb4\xa8\x5c\xe8\xfb\xa8\x17\x97\x3c\x4a\x22\x60\x28\x65\xaf\x8a\xbb\xb1\x57\xa4\x61\xdc\x9d\x06\xd1\x33\x5f\x37\x2a\x45\xa9\xf2\x88\xa1\x57\x91\x08\x42\x58\xfd\x79\x91\x3e\x4b\x7c\xdf\x45\xf3\x5d\x75\x32\x67\xf6\x03\x3f\xe4\xc3\x03\x63\x05\xf3\x0f\xf9\x9a\x05\x91\x0d\x23\xaf\x8c\xa3\x1d\xc4\xb4\x12\x53\x5c\x53\xe6\xe4\x66\x6f\xff\xa0\x34\x5c\x2c\x07\xdb\xbd\x56\x25\xcd\xba\x93\xcb\xbd\xb4\x0c\x01\x5e\x80\xe5\x3f\x61\xd1\xfb\x38\x8c\x77\x5f\xba\xd9\x09\x42\x7b\x9c\xe7\x99\x0b\x55\xf3\xa8\x8c\xa0\x08\x52\x1d\x1e\x4a\xbb\x67\x66\x4b\x34\x4e\x30\x2a\xfb\x59\xa7\x9b\x19\x17\x06\x59\x50\x72\xe9\xcc\x11\x5a\xc3\x81\xe6\x55\x0f\xfa\x04\xbf\x28\x21\x8e\x33\x38\x13\xef\x59\x42\x9c\x14\x8c\xf9\xc3\x3b\x8f\x4e\xbb\xe1\x44\xcb\xa2\x0b\xbe\x86\xb3\x4a\xa8\xc0\x99\xee\xc4\x7a\x0c\x30\x7f\x80\xcd\x46\x90\x4b\x2c\xdc\x41\xfd\x3e\x85\x1c\x7d\xa3\x78\x43\x1f\x50\xf8\x95\x36\x9c\xac\x0c\x6d\xd2\x85\x41\x7a\x4f\xfc\xf8\x91\xd2\x4e\x7d\x69\x6f\xbf\xf5\x60\x1e\xf1\xeb\x33\x70\xf9\xd1\x8b\x4b\xbf\xd0\x3f\x12\x02\xfa\xe9\xf4\xee\xb1\xbd\xed\x89\xe9\x6f\x36\x54\xba\x3c\x8a\x20\x5c\x99\xfd\x8f\xae\xcb\xf3\xc8\x77\x6b\x04\xd6\xb0\xfd\xf1\xc2\x34\xdf\x54\xba\xb5\x38\x61\x3a\xe6\x58\xeb\x98\xf2\x25\x55\x0c\x39\xb1\x48\x0d\xab\x48\xa1\x4d\x1d\x8b\xba\x8c\x19\x27\x52\x97\x7a\x62\x97\xe5\x51\x6a\x69\xb8\xfc\xf4\xd9\x18\xca\x45\x6b\x43\x32\x03\x71\x3a\xc6\x0c\x37\xa8\xb0\xf9\xf8\xfe\x14\x26\x47\x34\xe9\xa5\x2b\xe8\xcb\xbe\xee\xb7\x1e\x85\xd8\x71\xff\x61\xfc\x47\x07\x48\x1c\x4a\x34\x5e\xf9\x1c\x44\x98\x67\xcd\xa9\x1a\x4d\xbd\x7a\x1e\xe9\x5f\x9f\xb8\xd3\x8d\x5f\x06\xab\xa5\xe0\xf0\x84\x73\x7e\xc4\xfb\xb8\x39\xa3\x1b\x17\xe8\xf5\x28\x13\x7e\xba\x88\x7e\x7c\xd8\xe3\xc3\xe9\x18\xd7\xe1\x7c\x99\x1d\x9c\x36\xb3\x68\xfe\xaf\x7f\x1e\x54\xa9\xab\xce\x58\x9c\x3e\x6f\xa5\x06\xa4\xbc\xd6\x44\x07\xca\xb7\x83\x82\xb5\xc6\x6a\xd3\xf1\x89\x69\xa9\x4d\xfa\x24\x09\x7f\x9d\xf5\xee\xe6\x92\x76\x37\x97\x1b\xf8\x3d\xca\x25\x7a\x19\xbc\x81\x97\x46\x50\x79\x03\x96\x2a\x4b\x2c\x1a\x51\x8c\xb5\xee\xb6\x59\x1f\xe4\x7e\x7b\x7b\x46\xfc\xa2\xf9\x5f\x75\x0b\xfe\xcd\xc6\xbf\x25\xaa\x12\x39\x6c\x74\x6b\xa0\xa1\xd6\xfa\xfc\xa0\xd0\xa6\x9b\xa1\x8c\xe9\x56\x39\xa0\x0e\xb6\x5b\x98\x84\x37\x57\xb5\xf9\x9e\xd6\x08\x0f\x0f\xd7\x59\x1c\xae\x22\xff\x5b\xd0\xbd\x29\x3c\x34\xc0\x05\x57\x3f\x1f\xf0\x3b\x82\xcf\x8f\x2c\xca\x62\x02\x6f\x25\x52\x8b\xc0\x2a\x64\xf7\xb0\xaa\x34\x54\xd4\x7a\x50\xd1\x5a\x70\xba\x7b\x0c\x6b\x2a\xe4\x0e\x69\x51\x87\x8b\xa5\x43\xb9\x09\x17\xca\xb3\x0e\xce\x96\x63\xf2\x45\xca\xf1\x7f\xcd\xf9\xef\xd6\x1c\x4a\xe9\x27\x75\x4d\xf2\x61\x5d\xf3\xae\x12\xb6\x67\xfa\xca\xb7\x40\xeb\x74\x4d\xc3\xb5\xce\xbf\x3f\xa1\x72\xbe\x1d\x3c\xa7\x5f\x87\x3d\x0f\x0f\x13\x38\x6c\x34\x83\xa5\xb0\x0e\x4d\xc7\xfd\x23\xea\x77\x5a\x26\x4a\xa5\x0d\x82\xf3\xae\x6a\xb4\x96\x96\x38\x81\x77\x15\xee\x7a\x8b\x51\x05\x4a\x3b\x58\x20\xb4\x16\x39\xac\x84\xab\x74\xeb\xc0\x55\x08\x4c\xab\x42\x98\x3a\x90\xee\xb4\x93\x4e\xce\xf5\xcb\xef\x60\x27\x65\x3f\x77\x1d\xfe\xc2\xd6\xbd\x8d\xab\x2c\xee\xde\x1a\xb3\xf0\x4f\xc1\xf9\xbf\x03\x00\x00\xff\xff\xfd\x3e\xb0\xdb\x5b\x15\x00\x00")

func templatesPasswordChangedHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesPasswordChangedHtml,
		"templates/password-changed.html",
	)
}

func templatesPasswordChangedHtml() (*asset, error) {
	bytes, err := templatesPasswordChangedHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/password-changed.html", size: 5467, mode: os.FileMode(420), modTime: time.Unix(1497044014, 0)}
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
	"templates/confirm-de.html": templatesConfirmDeHtml,
	"templates/confirm-en.html": templatesConfirmEnHtml,
	"templates/confirm.html": templatesConfirmHtml,
	"templates/forgot-password.html": templatesForgotPasswordHtml,
	"templates/password-changed.html": templatesPasswordChangedHtml,
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
		"confirm-de.html": &bintree{templatesConfirmDeHtml, map[string]*bintree{}},
		"confirm-en.html": &bintree{templatesConfirmEnHtml, map[string]*bintree{}},
		"confirm.html": &bintree{templatesConfirmHtml, map[string]*bintree{}},
		"forgot-password.html": &bintree{templatesForgotPasswordHtml, map[string]*bintree{}},
		"password-changed.html": &bintree{templatesPasswordChangedHtml, map[string]*bintree{}},
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

