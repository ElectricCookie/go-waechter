// Code generated by go-bindata.
// sources:
// localeDefault/templates/confirm-de.html
// localeDefault/templates/confirm-en.html
// localeDefault/templates/confirm.html
// localeDefault/templates/forgot-password-de.html
// localeDefault/templates/forgot-password-en.html
// localeDefault/templates/forgot-password.html
// localeDefault/templates/password-changed-de.html
// localeDefault/templates/password-changed-en.html
// localeDefault/templates/password-changed.html
// DO NOT EDIT!

package localeDefault

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

var _localedefaultTemplatesConfirmDeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x4d\x6a\xdc\x40\x10\x46\xf7\x06\xdf\xe1\xc3\xeb\xf1\x5c\x21\x38\x3f\x8b\x90\x9f\x45\x42\x98\x75\xab\xf5\x49\x2a\xba\xbb\x3a\x54\x97\x0c\x19\xa1\x6b\x65\xe5\xdd\x5c\x2c\x8c\x3c\x86\x24\x83\x97\x0d\x5d\xaf\x5e\xbd\x65\x41\xcf\x41\x94\xb8\x73\xf1\xcc\xfb\x58\x75\x10\x2b\x77\x58\xd7\xdb\x1b\x00\x38\x48\xce\xa9\x96\x42\xc5\xb2\x60\xff\xa3\xd1\x34\x14\x62\x5d\x77\xb7\x37\xcb\x02\x6a\xbf\xfd\xdd\x1e\x2f\xac\x6e\x76\xaf\x7a\x0d\xfb\x70\xff\x25\x48\xc6\x23\x4d\x06\x39\x0a\x8d\xfa\x3a\x24\x56\x75\xaa\x5f\x53\xfa\xa0\x89\x18\x4e\x4f\x86\x8f\x93\x11\xdf\x38\x4a\x73\x13\xda\xac\x23\x3a\xca\x66\xfa\xae\x96\x9f\x41\x7f\x1d\xd8\x35\xf1\xb3\xef\x1e\x6f\xc5\x9d\x18\x6a\x1e\xa9\xf8\x2e\x44\xcf\x82\xf9\xbc\x05\xcd\x39\x51\x7b\x2a\x3e\x8b\xa6\x1d\xe6\x72\x66\xe3\x53\x55\xaf\x18\x8c\x72\x9c\x5b\x9c\x42\x76\xea\xfe\x75\xe5\xa1\x56\xa7\x5d\x1b\x7f\x95\x38\xf9\xff\xfd\xde\xe0\x40\xbd\x78\x08\x1b\x5f\xea\x88\x99\x9f\x9e\x4a\x96\x38\x81\xf6\xbc\x13\x53\xe8\xa8\x3b\xa4\xd3\x6f\x55\xfe\x3d\x44\xd1\x21\xc4\x09\x32\x6a\xb5\xad\xe8\x1e\xef\x69\xa0\x35\x67\xce\x4e\x3c\xc4\x58\x67\x75\x48\x73\xe8\x26\xd2\x51\x67\x3f\x76\xc1\x76\x68\x35\x07\x1d\x79\xa9\xa1\x78\x48\x2e\x8f\xcf\x25\x5b\x16\x4d\x97\x89\xa0\x29\x4b\x4c\xff\x1e\xff\x27\x00\x00\xff\xff\x71\x2a\x30\xe7\x3d\x02\x00\x00")

func localedefaultTemplatesConfirmDeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_localedefaultTemplatesConfirmDeHtml,
		"localeDefault/templates/confirm-de.html",
	)
}

func localedefaultTemplatesConfirmDeHtml() (*asset, error) {
	bytes, err := localedefaultTemplatesConfirmDeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "localeDefault/templates/confirm-de.html", size: 573, mode: os.FileMode(438), modTime: time.Unix(1501757702, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localedefaultTemplatesConfirmEnHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x3f\x4b\x03\x41\x10\xc5\xfb\x40\xbe\xc3\x33\x82\x95\xa6\xb1\x11\x2c\x2c\x04\xc1\x26\x95\x7f\xea\xbd\xbb\xb9\xdc\x90\xdd\x99\xb0\x3b\x67\x38\x96\xfd\xee\xe2\x46\xc5\x78\x5a\x0e\xbc\xf7\xe3\xf7\x26\x67\x74\xd4\xb3\x10\x56\xc6\xe6\xe9\xaa\x55\xe9\x39\x86\x15\x4a\x59\x2e\x00\xe0\x95\x7c\xab\x81\x90\x33\xd6\xcf\x89\xa2\xb8\x40\xa5\x5c\x2e\x17\x39\x83\xa4\xab\xb9\x7a\x7c\x71\x9a\xd1\x4c\x65\x0e\x7a\xa1\xc8\xfd\x04\x0a\x8e\xfd\xaf\xf6\x49\xbf\x55\x31\x12\x9b\x03\x6c\x70\xb2\x4b\xe8\x35\x22\xd2\x96\x93\x51\x64\xd9\xc2\x59\x75\xbb\xd7\xb0\x77\x32\x6d\x5c\x20\x94\x72\x86\x07\xf5\x5e\x0f\xb0\x81\xe0\x59\x76\x68\xa8\x9e\x8a\xb7\xa3\xc7\xa4\x63\x3c\xca\xc0\x75\x5d\xa4\x94\xd6\xff\x6f\xea\x55\x8d\xe2\x5c\x69\xa3\x76\xf2\x18\x94\x72\x87\xc7\xfe\x03\x8e\x8e\x3b\xb9\x38\xbf\xbe\xb9\xfd\xb6\xfd\x5b\xb5\x86\x5b\x27\x48\x1c\xf6\x7e\x02\x6f\x45\x23\xc1\x06\x4e\x47\xbf\x35\x9e\x06\x82\x6b\x5b\x1d\xc5\x70\x60\xef\x21\x6a\x68\x08\x63\x72\x8d\x27\x8c\xe2\x29\xa5\xca\xf9\x1c\xc7\xf6\x73\xcc\x7b\x00\x00\x00\xff\xff\x82\x52\x73\x0b\xe5\x01\x00\x00")

func localedefaultTemplatesConfirmEnHtmlBytes() ([]byte, error) {
	return bindataRead(
		_localedefaultTemplatesConfirmEnHtml,
		"localeDefault/templates/confirm-en.html",
	)
}

func localedefaultTemplatesConfirmEnHtml() (*asset, error) {
	bytes, err := localedefaultTemplatesConfirmEnHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "localeDefault/templates/confirm-en.html", size: 485, mode: os.FileMode(438), modTime: time.Unix(1501757704, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localedefaultTemplatesConfirmHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xeb\x8e\x1b\xb7\x15\xfe\xbf\x4f\x71\x96\x46\x80\xa4\xd8\xd1\x8c\xd6\x89\x93\x9d\x5b\xed\xba\x6e\x63\x20\x6d\x8d\x34\x06\x0a\x04\x46\x41\x0d\xcf\x48\xcc\x72\xc8\x29\xc9\xd1\xa5\xea\xbe\x4b\x9f\xa5\x4f\x56\x90\x1c\x49\xa3\x6b\xd6\x8e\xdd\xfe\x89\x7f\x58\xcb\xdb\xb9\x7c\xe7\xe3\x77\x38\xf9\x35\x53\x95\x5d\xb5\x08\x33\xdb\x88\xf2\x2a\x77\x3f\xb0\x6c\x84\x34\x05\x99\x59\xdb\xa6\x71\xbc\x58\x2c\x46\x8b\xa7\x23\xa5\xa7\xf1\xf8\xee\xee\x2e\x5e\xba\x3d\x24\x6c\x4a\xe7\x05\xe9\xb4\x4c\x4d\x35\xc3\x86\x9a\xa8\xe1\x95\x56\x46\xd5\x36\xaa\x54\x93\xce\x77\xfb\xd4\xa5\x7d\xaa\xae\x79\x85\xfd\x0f\x71\x51\x20\x65\xe5\x15\x40\x6e\xb9\x15\x58\xe6\x71\xf8\x75\x33\xd7\x51\xf4\x23\xaf\xe1\xba\x31\xea\x5d\xe9\x46\x10\x45\x7e\xa1\x41\x4b\xc1\x85\x1c\xe1\x3f\x3a\x3e\x2f\xc8\xdf\xa2\xb7\x2f\xa2\x97\xaa\x69\xa9\xe5\x13\x81\x04\x2a\x25\x2d\x4a\x5b\x90\xd7\xaf\x0a\x64\x53\xe7\x29\x18\xcc\xaf\x7f\x44\xc9\x78\xfd\xce\x99\x3a\x36\xf4\x32\x9c\x8b\x7e\x58\xb5\x43\x2b\x16\x97\x36\x76\x58\x64\x50\xcd\xa8\x36\x68\x8b\xb7\x3f\xfc\x21\xfa\x86\x6c\x6c\x48\xda\x60\x41\xe6\x1c\x17\xad\xd2\x76\x70\x72\xc1\x99\x9d\x15\x0c\xe7\xbc\xc2\xc8\x0f\x6e\x80\x4b\x6e\x39\x15\x91\xa9\xa8\xc0\x62\x3c\x4a\x9c\x19\x63\x57\x02\xc1\xd5\xa7\x77\x57\x19\xe3\xa3\x7e\xa2\x3a\x2b\x94\xba\x07\x0a\x6b\x68\x29\x63\x5c\x4e\x53\x48\x32\x78\xb8\x02\x18\x7d\x8f\x94\xfd\xc9\x4c\x7f\xa7\xd8\x0a\xd6\xe0\x1d\xa4\x30\x4e\x92\xcf\xfa\xf5\x57\x4b\x8b\x5a\x52\xf1\x52\x50\x63\x1e\xb1\xe3\x37\xb0\x06\xc1\x25\x46\x33\xe4\xd3\x99\x4d\x77\x1b\x27\xc1\x45\x43\xf5\x94\x4b\x1f\xc0\x30\x98\x68\x81\x93\x7b\x6e\x23\x17\x7a\x64\xf8\x3f\x31\xa2\xec\xa7\xce\xd8\x8d\xab\xa8\x31\x67\xd7\x9c\x75\x4b\x27\x02\x6f\xc0\x32\x58\xc3\x44\x69\x86\x3a\xaa\x94\x10\xb4\x35\x98\x6e\xfe\xc8\xa0\x31\x2a\xf2\x3b\x23\x61\x5a\x5a\x61\x0a\x49\x6b\x87\xd3\x7a\x38\xed\xec\xf2\x66\xba\xb5\xe8\x03\xed\xf3\x02\xda\x59\x95\xed\xa5\xda\x47\xe3\xe0\xe6\x12\x53\x90\x4a\x62\x06\x3e\x68\x86\x95\xd2\xd4\x72\x25\x37\xd3\x2e\x1f\x2e\x2d\xea\x56\x09\xbf\x10\x35\x8a\x61\x0a\x13\x5e\x75\x13\x5e\x05\xe7\x2d\xac\x81\x71\xd3\x0a\xba\x4a\x61\x22\x54\x75\x9f\x6d\xf1\x1b\x3f\x6d\x97\xa1\x8a\x79\xec\x6b\x5f\x5e\x1d\x13\xfe\x12\x2f\x9e\x37\xc8\x38\x05\x25\xc5\x0a\x4c\xa5\x11\x25\x50\xc9\xe0\xf3\x86\x2e\x03\xd1\xd2\x2f\xbf\x49\xda\xe5\x17\xb0\xbe\x02\x00\x78\xee\x22\xde\x10\x74\x4b\x84\xa7\xb7\x49\xbb\x0c\xc1\x02\x3c\xbf\xb4\x7c\x10\xe7\xfe\x3d\xea\x03\xf7\x71\x5f\xe5\x4b\x27\x2f\x00\xb9\x4a\xff\xe2\xaf\xf9\xef\x55\xd5\x35\x28\xed\x5f\xd1\x5a\x2e\xa7\xa6\xf4\xde\x72\x95\xbe\x10\x42\x2d\xde\xfc\xf9\x8f\xf1\x76\xe6\x0d\x5f\xa2\x30\x6f\x50\xbf\x96\xd5\xac\xbc\x7b\x96\xc7\x87\x73\xce\x70\x7c\xde\x72\x1e\x7b\xef\x27\xc3\x13\x16\x5d\x88\x30\x1e\xbf\xbb\x04\xec\xa8\xbf\x70\xd1\x54\xab\xae\x8d\x6a\xbe\xec\x21\x0c\x98\x38\x96\xc0\x35\x6f\x1c\x50\x54\xda\xec\x10\x9b\x81\xe3\x93\x15\xf5\x99\x0a\x2e\xef\x61\xa6\xb1\x0e\xd2\x6b\xd2\x38\xae\x95\xb4\x66\x34\x55\x6a\x2a\x90\xb6\xdc\x8c\x2a\xd5\xb8\x98\x7e\x5b\xd3\x86\x8b\x55\xf1\x76\xd2\x49\xdb\xa5\x4f\x93\xe4\xe6\xcb\x24\xb9\xf9\x2a\x49\x6e\xbe\x4e\x12\x02\x1a\x45\x41\xbc\x77\x33\x43\xb4\xe4\x54\x4a\x00\xe7\xd2\xf5\x8b\xbe\xf8\x21\x23\xe8\xb4\xf8\xfc\x97\x85\xf4\x45\x16\xac\x6e\x31\x39\x56\xde\x0f\x60\x35\x97\xa7\x58\x3d\x6a\x7e\x72\x3a\xd1\x35\x32\x6a\x51\x47\xe3\x24\xd9\x72\xd7\xd5\x69\x50\xa6\x23\x12\xc7\xa1\xf3\xe4\x5e\xd7\xfc\x64\x41\x26\xb4\xba\x77\x55\x97\x2c\x85\x27\x75\x5d\x67\x3e\x26\x97\x00\xe3\x73\xa8\x9c\x42\x16\xc4\xbb\x94\x96\x72\x89\x9a\x1c\x9f\x74\xf1\x28\x9d\xf6\xc7\x07\x57\x03\xfe\x05\xaf\x5f\xbd\x2b\x7b\xc4\x73\x2f\x58\xa0\x95\x3b\xdd\x6a\x34\x28\xad\x17\x12\xd2\xab\x55\x41\x12\x02\x15\x0a\xd1\xcb\xec\x76\xec\x24\x6e\x33\x0e\xed\x85\x3c\x73\x44\xa0\x82\x4f\x65\x41\x2a\x74\xb2\xb4\x0d\x2c\x80\xf1\x2c\x71\x17\x99\x94\xdb\x72\xe7\x56\xef\x06\x6e\xc8\x36\x07\x86\x92\xe8\x0e\x39\x12\x78\xd5\xf6\x23\x27\xb5\x83\x1d\x91\xee\x04\xa6\xb8\xa4\x95\x15\xab\x9d\xfd\xbd\x5a\x3b\xe8\x7a\xdb\xbd\xf8\x25\xed\x32\xe8\xef\x4e\xab\x36\x01\x5e\x80\xe5\xe7\xb0\xe8\x7d\xec\xc7\xbb\xe3\x42\x76\x84\xd0\x0e\xe7\x32\xb7\x8e\x06\xa5\x43\x65\x00\x85\xd7\x7e\x7f\x28\x0d\x67\xb2\x39\x6a\xcb\x2b\x2a\xfa\x59\xab\xda\x8c\x71\x8d\x95\x6f\x0d\xc2\xea\x03\xb4\x36\x1d\xd2\xc9\x28\xf4\x09\x7e\x52\x42\x3c\xa6\xc0\x27\x72\x38\x49\x92\xa3\x22\x0e\xf9\xbf\x77\xe5\x8e\x04\x93\x5c\xf0\xb5\x69\x88\x5c\x7a\x1e\x85\xb6\x78\x0e\x44\xd7\x25\xb3\x41\x19\x04\xd6\x76\xaf\xa6\xbf\x84\x30\xfd\xe5\x71\x86\x1e\x41\x86\x85\xd2\x2c\x5a\x68\xda\xa6\x13\x8d\xf4\x3e\x72\xe3\x33\xe5\x1e\xbb\x72\xdf\x7e\xe5\xc0\x3c\xe0\xdc\x47\xe0\xf7\xd9\xd7\x51\xbf\xd0\x1f\xf1\x01\xbd\x3f\xe5\x7b\x6c\x6f\x7b\xb2\xba\xe7\x13\x15\xb6\x20\x04\xfc\xbb\xdc\xfd\x11\x6e\x7e\x41\xdc\x0d\x26\x60\x74\x55\x90\xf5\x1a\x46\xdf\xa9\xa9\x7a\xfb\xfd\x77\xf0\xf0\x70\x10\x6a\xea\x5f\x4c\x7d\x74\x9a\x32\xde\x19\x1f\xdc\x86\x08\x81\x01\x07\x35\xdf\xbc\xc2\xfc\xd9\xc3\x37\x98\x9f\x1c\xb0\xa0\x57\x2b\x2f\x29\xbb\xb2\xde\xba\x24\x63\xcb\xdc\x7f\xda\xfd\x17\xf2\x8d\x7d\x05\x86\x2b\x1f\xa3\xce\x65\xde\x1e\x0b\xd0\xd8\x09\xe6\x81\xe4\xf5\x38\x58\xd5\xba\x65\x30\x4a\x70\x06\x4f\x18\x63\x07\xb4\x8e\xdb\x13\x52\x71\x81\x3d\x67\x0b\xfd\xfe\xba\xf9\xe1\x61\x0f\xfb\xd1\x21\xae\x9b\x96\x92\xed\x35\x98\x8c\x94\xff\xf9\xf7\x5e\x95\x42\x75\x86\xda\xf3\x71\x2b\xb5\x41\xca\x49\x09\xd9\xeb\x4e\x55\xa7\x8d\xd2\x81\x46\x7d\x13\xa7\x94\x06\x8b\xe1\xc5\x93\xce\x50\xcc\xd1\x29\xda\xc0\x8f\xbf\x2b\xc3\xa4\x6e\x6f\x4f\x88\x16\x29\xd7\x6b\x00\x8b\x4d\x2b\xa8\x45\x20\xfe\x3a\xb9\x87\x44\xcd\x75\x43\x60\x04\x0f\x0f\x79\xcc\xf8\xfc\xff\x9a\x6e\xe2\xff\xed\x65\x1c\xde\x78\x37\xf0\xed\x26\xf3\x1b\x78\xa1\x39\x15\x37\x60\xa8\x34\x91\x41\xcd\xeb\xe1\xe5\xbd\x7d\x34\x18\x3b\x2c\xfa\xaf\xe5\xff\x21\x1a\x9f\x4e\x8e\x0d\xb6\x54\x53\x8b\x1f\xa0\xbd\xe7\xe5\xd2\x29\x62\x28\xd1\x62\xc6\x2d\x66\xc3\xd2\xfd\x7c\x96\x30\xef\xc7\x0d\x67\x4c\x20\x81\xc9\xd4\x1b\x2b\xc8\x93\xf1\x1d\x1d\xdf\x7d\x4d\xca\x9c\xf6\x5f\x23\x4e\xc9\x5f\x86\x32\xbc\x60\x4c\xa3\x31\x43\x41\x3f\x29\xc4\x83\x47\x73\x6f\x6f\x2f\xd6\x21\x99\xbe\x3d\x71\x7d\xbc\xdc\xfb\xe1\x22\x30\x46\x2a\xdd\x50\xb1\x47\xa2\xf1\x6d\xf2\x59\x20\x91\xd5\x54\x9a\x5a\xe9\x26\xf8\xde\xc9\x54\x46\xc0\x52\x3d\x45\x5b\x90\xbf\x4f\x04\x95\xf7\x07\x1c\x9b\x74\xd6\x2a\x79\x48\x31\xfa\x6b\x87\xf8\xb5\x43\x7c\xd4\x0e\xf1\xbe\x7a\xf9\xf8\xe6\xb1\xe3\x72\xad\x94\xf5\x9a\x73\x49\x2e\x8f\xd8\xec\xd7\xcf\x7e\x7d\x1c\x01\x7c\xea\x2b\xe0\x13\x5b\x77\x36\xae\xf2\xd8\x5b\x76\x1f\xe8\xb6\x11\xe5\x7f\x03\x00\x00\xff\xff\xbd\x6a\x83\xfa\xb7\x16\x00\x00")

func localedefaultTemplatesConfirmHtmlBytes() ([]byte, error) {
	return bindataRead(
		_localedefaultTemplatesConfirmHtml,
		"localeDefault/templates/confirm.html",
	)
}

func localedefaultTemplatesConfirmHtml() (*asset, error) {
	bytes, err := localedefaultTemplatesConfirmHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "localeDefault/templates/confirm.html", size: 5815, mode: os.FileMode(438), modTime: time.Unix(1501757950, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localedefaultTemplatesForgotPasswordDeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\x6a\xeb\x30\x14\x44\xf7\x81\xfc\xc3\x90\x85\x57\x79\xfe\x81\xb7\x2a\x34\x85\x42\x0b\x85\x74\xd5\x9d\x12\x4d\x6c\x61\xfb\x0a\xa4\x2b\x52\x2c\xfc\xef\x45\x6e\x1d\x9c\x45\x96\x12\x97\x33\x33\x27\x67\x58\x5e\x9c\x10\xbb\xb3\x17\xa5\xe8\x0e\xd3\xb4\xdd\x00\xc0\xd1\x11\xad\x39\x51\xc0\x72\x70\xf8\xf7\x6e\x5c\x8f\x31\x05\x7c\xa5\x50\xa5\x34\xf4\xff\xcf\x5d\xa4\x8e\x49\x1a\xbc\xb6\x81\x11\x1f\x26\xc6\xab\x0f\x1a\x61\xa4\xe1\xc5\x07\xcb\xa0\x35\x5e\x7c\xdf\x50\x66\xa0\xe5\x80\x54\x82\x10\x95\x2d\xa5\xbc\xdf\x9c\x74\x7b\xa4\xa1\x40\x6e\x08\x8c\x09\x95\x29\x21\x62\x19\xa4\xde\x6e\x72\x06\xc5\xce\xf5\xe6\xc7\x52\xfc\x94\x54\xbd\x7c\xf2\x7b\xd5\x7d\x45\xb9\xab\x4a\x79\xcc\xb9\x78\xaf\x0c\xf7\x9c\x67\xc7\x78\x9b\x7e\x4d\xc1\x12\x46\x90\x33\xea\xc3\x50\xbe\xa6\x09\x0d\x63\x99\xa1\x35\x8e\xbe\xef\xf5\x6f\x67\x37\x3b\xab\x9e\x96\x01\xc5\x91\x7d\x60\xe8\xd7\xf2\x1e\x5d\xe5\xe7\x73\x59\x5c\xad\xc3\x5d\x23\x3e\x38\x06\x4a\x8d\xd5\x86\x9f\x00\x00\x00\xff\xff\x3c\x41\xec\x35\xc1\x01\x00\x00")

func localedefaultTemplatesForgotPasswordDeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_localedefaultTemplatesForgotPasswordDeHtml,
		"localeDefault/templates/forgot-password-de.html",
	)
}

func localedefaultTemplatesForgotPasswordDeHtml() (*asset, error) {
	bytes, err := localedefaultTemplatesForgotPasswordDeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "localeDefault/templates/forgot-password-de.html", size: 449, mode: os.FileMode(438), modTime: time.Unix(1501757707, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localedefaultTemplatesForgotPasswordEnHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x31\x4f\xc3\x30\x10\x85\xf7\x4a\xfd\x0f\x4f\x45\x62\x42\x5d\x98\x80\x99\x81\x15\x75\x61\x74\xe3\x8b\x63\xe1\xdc\x15\xdf\x59\x21\xb2\xf2\xdf\x51\x2c\x88\xda\x81\xcd\xcf\xb6\xde\xfb\xbe\x5a\xe1\xa9\x8f\x4c\x38\x74\xc2\x46\x6c\x07\x2c\xcb\x7e\x07\x00\x1f\x52\x60\x92\x3c\x8a\x62\x96\x82\x5e\x72\x10\x5b\x8f\x19\x17\xa7\x3a\x49\xf6\x47\xbc\xf5\xed\x31\x93\x4b\x69\x86\x8f\xbe\xc5\xce\x31\xba\xc1\x71\x20\x44\x43\xd1\xc8\x01\x36\x10\x52\xe4\x4f\x9c\x29\xc9\xf4\xbc\xdf\xd5\x0a\x62\xdf\xf6\x5a\xf8\x23\x39\x17\x33\xe1\x13\x7d\x5f\xc1\xbc\x93\x9a\x64\xc2\x38\x6f\xdb\xff\x17\xf4\x22\x46\xf9\xb6\xe0\x34\x44\x05\x8d\x2e\x26\x4c\x4e\xa1\xc4\x06\x13\xd4\x8a\xe3\x6b\xbb\x5d\x96\x4d\xc6\x47\xcf\xf7\x77\x8f\x4f\x2f\x86\x4c\x5f\x85\xb4\x7d\xfd\xf5\xb9\xf1\x7f\xd8\x6c\x35\x8e\x97\x34\x23\x06\x5e\x31\x6d\x5d\x1b\x49\xd5\x05\x3a\x5e\x81\xfe\x04\x00\x00\xff\xff\x0d\xc7\x30\xbe\x70\x01\x00\x00")

func localedefaultTemplatesForgotPasswordEnHtmlBytes() ([]byte, error) {
	return bindataRead(
		_localedefaultTemplatesForgotPasswordEnHtml,
		"localeDefault/templates/forgot-password-en.html",
	)
}

func localedefaultTemplatesForgotPasswordEnHtml() (*asset, error) {
	bytes, err := localedefaultTemplatesForgotPasswordEnHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "localeDefault/templates/forgot-password-en.html", size: 368, mode: os.FileMode(438), modTime: time.Unix(1501757709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localedefaultTemplatesForgotPasswordHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xd9\x6e\x1b\xbd\x15\xbe\xf7\x53\x1c\xd3\xf8\x81\x3f\x85\x46\x33\x92\xb3\x79\x16\x35\x69\x9a\xb6\x01\xba\x04\x6d\x0c\x14\x08\x7c\x41\x0d\xcf\x48\x8c\x39\xe4\x94\xe4\x68\xa9\xea\x77\xe9\xb3\xf4\xc9\x0a\x72\x46\xd2\x68\x73\x12\x3b\xbe\x8b\x2f\x2c\x71\x3b\xfc\xce\x77\x3e\x7e\xa4\xd2\x73\xa6\x72\xbb\xac\x10\xa6\xb6\x14\xa3\xb3\xd4\x7d\xc0\xa2\x14\xd2\x64\x64\x6a\x6d\x15\x87\xe1\x7c\x3e\xef\xcf\x2f\xfb\x4a\x4f\xc2\xc1\xd5\xd5\x55\xb8\x70\x73\x48\x33\x29\x9e\x65\xa4\xd6\x32\x36\xf9\x14\x4b\x6a\x82\x92\xe7\x5a\x19\x55\xd8\x20\x57\x65\x3c\xdb\xce\x53\xf7\xcd\x53\x45\xc1\x73\x6c\x3f\x88\x43\x81\x94\x8d\xce\x00\x52\xcb\xad\xc0\x51\x1a\x36\x9f\xae\xe7\x3c\x08\x3e\xf3\x02\xce\x4b\xa3\x6e\x46\xae\x05\x41\xe0\x07\x4a\xb4\x14\x1c\xe4\x00\xff\x55\xf3\x59\x46\xfe\x19\x5c\xbf\x0d\xde\xa9\xb2\xa2\x96\x8f\x05\x12\xc8\x95\xb4\x28\x6d\x46\x3e\xbc\xcf\x90\x4d\xdc\x4e\x4d\xc0\xf4\xfc\x33\x4a\xc6\x8b\x1b\x17\xea\x30\xd0\xbb\x66\x5d\xf0\x69\x59\x75\xa3\x58\x5c\xd8\xd0\x71\x91\x40\x3e\xa5\xda\xa0\xcd\xae\x3f\xfd\x21\x78\x4d\xd6\x31\x24\x2d\x31\x23\x33\x8e\xf3\x4a\x69\xdb\x59\x39\xe7\xcc\x4e\x33\x86\x33\x9e\x63\xe0\x1b\x3d\xe0\x92\x5b\x4e\x45\x60\x72\x2a\x30\x1b\xf4\x23\x17\xc6\xd8\xa5\x40\x70\xf5\x69\xb7\xcb\x8d\xf1\xa8\x2f\x54\x6d\x85\x52\xb7\x40\x61\x05\x15\x65\x8c\xcb\x49\x0c\x51\x02\x77\x67\x00\xfd\xbf\x23\x65\x7f\x31\x93\xdf\x29\xb6\x84\x15\xf8\x0d\x62\x18\x44\xd1\x2f\xed\xf8\xfb\x85\x45\x2d\xa9\x78\x27\xa8\x31\xdf\x30\xe3\x37\xb0\x02\xc1\x25\x06\x53\xe4\x93\xa9\x8d\xb7\x13\xc7\xcd\x16\x25\xd5\x13\x2e\x3d\x80\x2e\x98\x60\x8e\xe3\x5b\x6e\x03\x07\x3d\x30\xfc\xdf\x18\x50\xf6\xa5\x36\x76\xbd\x55\x50\x9a\x93\x63\x2e\xba\xa5\x63\x81\x3d\xb0\x0c\x56\x30\x56\x9a\xa1\x0e\x72\x25\x04\xad\x0c\xc6\xeb\x2f\x09\x94\x46\x05\x7e\x66\x20\x4c\x45\x73\x8c\x21\xaa\x6c\xb7\x5b\x77\xbb\x5d\x5c\x5e\x4e\x36\x11\x3d\xd0\x36\x2f\xa0\xb5\x55\xc9\x4e\xaa\x2d\x1a\x47\x37\x97\x18\x83\x54\x12\x13\xf0\xa0\x19\xe6\x4a\x53\xcb\x95\x5c\x77\xbb\x7c\xb8\xb4\xa8\x2b\x25\xfc\x40\x50\x2a\x86\x31\x8c\x79\x5e\x8f\x79\xde\x6c\x5e\xc1\x0a\x18\x37\x95\xa0\xcb\x18\xc6\x42\xe5\xb7\xc9\x86\xbf\xc1\x65\xb5\x68\xaa\x98\x86\xbe\xf6\xa3\xb3\x43\xc1\xdf\xa7\x8b\x37\x25\x32\x4e\x41\x49\xb1\x04\x93\x6b\x44\x09\x54\x32\xf8\xb5\xa4\x8b\x46\x68\xf1\xf3\xd7\x51\xb5\x78\x06\xab\x33\x00\x80\x37\x0e\xf1\x5a\xa0\x1b\x21\x5c\x0e\xa3\x6a\xd1\x80\x05\x78\x73\xdf\xf0\x1e\xce\xdd\x73\xd4\x02\xf7\xb8\xcf\xd2\x85\xb3\x17\x80\x54\xc5\x7f\xf3\xc7\xfc\xf7\x2a\xaf\x4b\x94\xf6\x1f\x68\x2d\x97\x13\x33\xf2\xbb\xa5\x2a\x7e\x2b\x84\x9a\x7f\xfc\xeb\x1f\xc3\x4d\xcf\x47\xbe\x40\x61\x3e\xa2\xfe\x20\xf3\xe9\xe8\xea\x65\x1a\xee\xf7\xb9\xc0\xe1\xe9\xc8\x69\xe8\x77\x3f\x0a\x4f\x58\x74\x10\x61\x30\xb8\xb9\x8f\xd8\x7e\x7b\xe0\x82\x89\x56\x75\x15\x14\x7c\xd1\x52\xd8\x70\xe2\x54\x02\xe7\xbc\x74\x44\x51\x69\x93\x7d\x6e\x3a\x1b\x1f\xad\xa8\xcf\x54\x70\x79\x0b\x53\x8d\x45\x63\xbd\x26\x0e\xc3\x42\x49\x6b\xfa\x13\xa5\x26\x02\x69\xc5\x4d\x3f\x57\xa5\xc3\xf4\xdb\x82\x96\x5c\x2c\xb3\xeb\x71\x2d\x6d\x1d\x5f\x46\x51\xef\x79\x14\xf5\x5e\x44\x51\xef\x55\x14\x11\xd0\x28\x32\xe2\x77\x37\x53\x44\x4b\x8e\xa5\x04\x70\x2a\x5d\x3f\xe8\x8b\xdf\x64\x04\xb5\x16\xbf\x3e\x0e\xd2\xb3\xa4\x89\xba\xe1\xe4\xd0\x79\x1f\xa0\x6a\x2e\x8f\xa9\xba\x5f\x7e\x71\x3e\x51\x97\x32\xa8\x50\x07\x83\x28\xda\x68\xd7\xd5\xa9\x53\xa6\x03\x11\x87\xcd\xcd\x93\x7a\x5f\xf3\x9d\x19\x19\xd3\xfc\xd6\x55\x5d\xb2\x18\x2e\x8a\xa2\x48\x3c\x26\x97\x00\xe3\x33\xc8\x9d\x43\x66\xc4\x6f\x29\x2d\xe5\x12\x35\x39\x5c\xe9\xf0\x28\x1d\xb7\xcb\x3b\x47\x03\xfe\x03\x1f\xde\xdf\x8c\x5a\xc6\x53\x6f\x58\xa0\x95\x5b\x5d\x69\x34\x28\xad\x37\x12\xd2\xba\x55\x46\x22\x02\x39\x0a\xd1\xda\xec\xa6\xed\x2c\x6e\xdd\x6e\xae\x17\xf2\xd2\x09\x81\x0a\x3e\x91\x19\xc9\xd1\xd9\xd2\x06\x58\x43\xc6\xcb\xc8\x1d\x64\x32\xda\x94\x3b\xb5\x7a\xdb\x70\x4d\xb6\x5e\xd0\xb5\x44\xb7\xc8\x89\xc0\xbb\xb6\x6f\x39\xab\xed\xcc\x08\x74\x2d\x30\xc6\x05\xcd\xad\x58\x6e\xe3\xef\xd4\xda\x51\xd7\xc6\x6e\xcd\x2f\xaa\x16\x8d\xff\x6e\xbd\x6a\x0d\xf0\x1e\x5a\xbe\xc6\x45\xbb\xc7\x2e\xde\xad\x16\x92\x03\x86\xb6\x3c\x8f\x52\xeb\x64\x30\x72\xac\x74\xa8\xf0\xde\xef\x17\xc5\xcd\x9a\x64\x86\xda\xf2\x9c\x8a\xb6\xd7\xaa\x2a\x61\x5c\x63\xee\xaf\x06\x61\xf5\x1e\x5b\xeb\x1b\xd2\xd9\x28\xb4\x09\x3e\xa9\x20\xbe\xa5\xc0\x47\x72\x38\x2a\x92\x83\x22\x76\xf5\xbf\x73\xe4\x0e\x0c\x93\xdc\xb3\xd7\xfa\x42\xe4\xd2\xeb\xa8\xb9\x16\x4f\x91\xe8\x6e\xc9\xa4\x53\x06\x81\x85\xdd\xa9\xe9\x63\x04\xd3\x1e\x1e\x17\xe8\x1b\xc4\x30\x57\x9a\x05\x73\x4d\xab\x78\xac\x91\xde\x06\xae\x7d\xa2\xdc\x03\x57\xee\xe1\x0b\x47\xe6\x9e\xe6\x7e\x80\xbe\x4f\xbe\x8e\xda\x81\x76\x89\x07\xf4\xfd\x92\x6f\xb9\x1d\xb6\x62\x75\xcf\x27\x2a\x6c\x46\x08\xf8\x77\xb9\xfb\xd2\x9c\xfc\x8c\xb8\x13\x4c\xc0\xe8\x3c\x23\xab\x15\xf4\xfd\xf3\x5b\x2e\xff\xac\x26\xea\x5a\x0b\xb8\xbb\xdb\x43\x1c\xfb\x87\x53\x0b\x52\x53\xc6\x6b\xe3\x31\xae\xf5\xd0\x08\x61\xaf\xf4\xeb\xc7\x98\x5f\xbb\xff\x14\xf3\x9d\x1d\x31\xb4\xa6\xe5\x9d\x65\x5b\xdd\xa1\xcb\x35\xb4\xcc\xfd\xd3\xee\x5f\x93\x76\xe8\x0b\xd1\x1d\xf9\x11\xe5\x1e\xa5\xd5\xa1\x0f\x0d\x9c\x6f\xee\x39\x5f\xcb\x83\x55\x95\x1b\x06\xa3\x04\x67\x70\xc1\x18\xdb\x53\x77\x58\x1d\x71\x8c\x7b\x44\x74\xb2\xde\xdf\x6f\x9f\x0f\x87\xdd\xbd\x96\xf6\x79\x5d\xdf\x2c\xc9\xce\x3d\x93\x90\xd1\xff\xfe\xbb\x53\xa5\xa6\x3a\x5d\x0b\xfa\xb1\x95\x5a\x33\xe5\x1c\x85\xec\x5c\x52\x79\xad\x8d\xd2\x8d\x8c\xda\xbb\x3c\xf2\x7f\x4d\xd0\xe6\xed\x13\x37\x6f\x9f\x1e\xfc\x09\xc5\x0c\x9d\xc7\xf5\xe0\xad\xe6\x54\xf4\xc0\x50\x69\x02\x83\x9a\x17\x5d\x35\x0f\xab\xc5\x4e\xca\xc3\xe1\x11\x67\x23\xa3\xd5\x0a\x2c\x96\x95\xa0\x16\x81\xb4\xbf\x22\x09\xdc\xdd\xa5\x21\xe3\xb3\xa7\xe2\xe0\xe9\xcc\xc9\x60\x45\x35\xb5\xf8\x00\x27\x3a\xed\x1a\xce\x18\x9a\xc2\xcc\xa7\xdc\x62\xd2\x2d\xd8\xd7\xb3\x84\x59\xdb\x2e\x39\x63\x02\x09\x8c\x27\x3e\x58\x46\x2e\x06\x57\x74\x70\xf5\xaa\x7b\x88\x8f\x5a\x4e\xe7\x95\xd8\x2e\xd9\x81\xd3\x55\xc9\x46\x1d\xfb\xc6\xe6\x9b\xf3\x46\x0a\x52\xe9\x92\x8a\x1d\x75\x0c\x86\xd1\x2f\x8d\x3a\xac\xa6\xd2\x14\x4a\x97\xcd\xde\xdb\x03\xb9\x2f\x96\x71\x6d\xad\x92\x9f\x70\xd1\xea\xa5\xfa\xe9\x7a\x3f\x5d\xef\x51\xae\x47\x29\x7d\x94\xe5\x45\x0f\xb0\xbc\x42\x29\x8b\x7a\xab\xe2\x3d\xd7\x3b\xd0\xb1\x1f\x3f\xf9\xa4\x3e\xa0\xf6\xd8\xd3\xf6\x89\xa3\xbb\x18\x67\x69\xe8\x23\xbb\x5f\x9d\xb6\x14\xa3\xff\x07\x00\x00\xff\xff\x36\x7f\x95\xe3\x8c\x15\x00\x00")

func localedefaultTemplatesForgotPasswordHtmlBytes() ([]byte, error) {
	return bindataRead(
		_localedefaultTemplatesForgotPasswordHtml,
		"localeDefault/templates/forgot-password.html",
	)
}

func localedefaultTemplatesForgotPasswordHtml() (*asset, error) {
	bytes, err := localedefaultTemplatesForgotPasswordHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "localeDefault/templates/forgot-password.html", size: 5516, mode: os.FileMode(438), modTime: time.Unix(1501757950, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localedefaultTemplatesPasswordChangedDeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\xc1\x4a\xf4\x30\x14\x85\xf7\x03\xf3\x0e\xe7\x9f\x45\x57\x7f\xe7\x05\x5c\x89\xb8\x10\x51\x04\x17\x82\xbb\xb4\x39\x4d\x03\xcd\x4d\x49\x6f\x1c\x34\xf4\x6d\x7c\x13\x5f\x4c\x32\x65\xc0\x8d\xe0\x2e\x81\x7b\xce\xf9\xbe\x52\x60\x39\x78\x21\x0e\xea\x75\x62\x3b\x9b\x65\x39\xc5\x64\xdb\x7e\x34\xe2\x68\x0f\x58\xd7\xfd\xee\xd9\x13\xa3\xe9\x28\xb8\x1b\x13\x9e\xb6\x1b\x45\x47\x8f\x52\x70\xbc\x89\x61\x36\xf2\xfe\x68\x02\xb1\xae\x70\xfc\xfa\x14\xcb\xa4\xff\xf6\xbb\x52\x40\xb1\xe7\x92\xf3\xe7\xb2\xd6\x47\x51\x8a\xfe\xb2\xf7\x42\x11\xd4\x51\xeb\xb9\x10\xcd\x75\x0e\xd3\x55\xad\xcc\xe2\x20\xbe\x1f\x15\x6f\x31\x39\x4a\x0c\x81\xb2\xa1\xfd\xc7\x12\xa7\x49\xb9\x05\x97\x38\x54\xc2\x26\xd7\x64\xc7\x34\xa7\xed\x39\x50\x70\x62\xc2\x6b\x76\x46\x1c\x3e\x72\x35\x62\xc0\x6d\xfb\x60\xfc\x84\xfb\x28\x1a\x31\x1a\x45\x16\xfb\x27\x59\x26\x61\x56\x34\xe6\x82\x28\xc7\x1f\xd6\xdf\x01\x00\x00\xff\xff\x8a\x82\x33\xdf\x60\x01\x00\x00")

func localedefaultTemplatesPasswordChangedDeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_localedefaultTemplatesPasswordChangedDeHtml,
		"localeDefault/templates/password-changed-de.html",
	)
}

func localedefaultTemplatesPasswordChangedDeHtml() (*asset, error) {
	bytes, err := localedefaultTemplatesPasswordChangedDeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "localeDefault/templates/password-changed-de.html", size: 352, mode: os.FileMode(438), modTime: time.Unix(1501757711, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localedefaultTemplatesPasswordChangedEnHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x31\x4b\xc4\x40\x10\x85\xfb\x83\xfb\x0f\xcf\x13\xec\x2e\x8d\x95\x58\x5a\xd9\x88\xad\xe5\xb0\x3b\x31\xab\xbb\x33\x47\x66\xc2\x11\x96\xfd\xef\x92\x3b\x23\x22\xb1\x7c\x8f\x99\xf7\xbe\x57\x2b\x22\xf7\x49\x18\x07\x4f\x9e\xf9\x78\x22\xb3\xb3\x8e\xf1\x18\x06\x92\x77\x8e\x07\xb4\xb6\xdf\xbd\xe9\x84\x8f\xc9\x1c\xdf\x2e\x66\x9d\x46\xac\xb7\xe8\x75\xbc\x3a\x14\x82\x4e\xe2\x20\x47\xad\xe8\x9e\xb4\x9c\x48\xe6\x17\x2a\x8c\xd6\x6e\xf6\xbb\x5a\xc1\x12\x2f\x91\x17\xb1\x76\x07\x15\x67\xf1\x7f\xda\x9f\xfb\x25\x1d\x31\x45\xb9\xbb\xbd\x7f\x78\x5c\x31\xfe\x50\x2c\xca\x38\xf7\x1d\x5e\x33\x93\x31\xc2\xc0\xe1\x13\xe7\x41\x31\x90\x2d\x6c\x6c\x06\xd7\xeb\x1b\x17\x4a\xf9\x07\x38\x95\xc2\x31\x91\x73\x9e\x41\x12\xb7\x0b\x36\x57\x75\xf8\x35\xeb\x2b\x00\x00\xff\xff\xac\x23\xab\xd5\x4f\x01\x00\x00")

func localedefaultTemplatesPasswordChangedEnHtmlBytes() ([]byte, error) {
	return bindataRead(
		_localedefaultTemplatesPasswordChangedEnHtml,
		"localeDefault/templates/password-changed-en.html",
	)
}

func localedefaultTemplatesPasswordChangedEnHtml() (*asset, error) {
	bytes, err := localedefaultTemplatesPasswordChangedEnHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "localeDefault/templates/password-changed-en.html", size: 335, mode: os.FileMode(438), modTime: time.Unix(1501757713, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localedefaultTemplatesPasswordChangedHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xdb\x92\x1b\x3b\x15\x7d\x9f\xaf\xd8\xa3\x14\x55\xe7\x50\x96\xbb\x3d\x73\x4e\x38\xd3\xee\x36\x09\x21\x40\x1e\x80\x14\x9c\x54\x51\x95\x9a\x07\x59\xda\x6e\x2b\xa3\x96\x1a\x49\xed\x0b\x26\xff\xc2\xb7\xf0\x65\x94\xd4\x6d\xbb\x7d\x1b\x72\x7f\x81\x3c\x8c\xad\xdb\xbe\xac\xbd\xb4\xb6\x9c\xfc\x5a\x18\xee\xd7\x35\xc2\xdc\x57\x6a\x72\x95\x87\x0f\x58\x55\x4a\xbb\x82\xcc\xbd\xaf\xb3\x24\x59\x2e\x97\xc3\xe5\xed\xd0\xd8\x32\x19\xdd\xdd\xdd\x25\xab\xb0\x87\xb4\x9b\xb2\x45\x41\x1a\xab\x33\xc7\xe7\x58\x31\x47\x2b\xc9\xad\x71\x66\xe6\x29\x37\x55\xb6\xd8\xef\x33\x8f\xed\x33\xb3\x99\xe4\xd8\x7d\x90\x10\x05\x32\x31\xb9\x02\xc8\xbd\xf4\x0a\x27\x79\xd2\x7e\x86\x99\x6b\x4a\xdf\xca\x19\x5c\x57\xce\xdc\x4f\xc2\x08\x28\x8d\x0b\x15\x7a\x06\x21\x64\x8a\x7f\x6f\xe4\xa2\x20\x7f\xa3\x6f\x9e\xd3\x17\xa6\xaa\x99\x97\x53\x85\x04\xb8\xd1\x1e\xb5\x2f\xc8\xab\x97\x05\x8a\x32\x78\x6a\x0d\xe6\xd7\x6f\x51\x0b\x39\xbb\x0f\xa6\x4e\x0d\xbd\x68\xcf\xd1\x9f\xd7\x75\xdf\x8a\xc7\x95\x4f\x02\x16\x63\xe0\x73\x66\x1d\xfa\xe2\xcd\xcf\xbf\xa3\x3f\x91\xad\x0d\xcd\x2a\x2c\xc8\x42\xe2\xb2\x36\xd6\xf7\x4e\x2e\xa5\xf0\xf3\x42\xe0\x42\x72\xa4\x71\x30\x00\xa9\xa5\x97\x4c\x51\xc7\x99\xc2\x62\x34\x4c\x83\x19\xe7\xd7\x0a\x21\xd4\xa7\x73\xc7\x9d\x8b\x51\x3f\x31\x8d\x57\xc6\x3c\x00\x83\x0d\xd4\x4c\x08\xa9\xcb\x0c\xd2\x31\xbc\xbf\x02\x18\xfe\x05\x99\xf8\xa3\x2b\x7f\x63\xc4\x1a\x36\x10\x1d\x64\x30\x4a\xd3\x5f\x74\xeb\x2f\x57\x1e\xad\x66\xea\x85\x62\xce\x7d\xc0\x8e\x5f\xc2\x06\x94\xd4\x48\xe7\x28\xcb\xb9\xcf\xf6\x1b\xa7\xad\x8b\x8a\xd9\x52\xea\x18\x40\x3f\x18\xba\xc4\xe9\x83\xf4\x34\x84\x4e\x9d\xfc\x07\x52\x26\xde\x35\xce\x6f\x5d\xd1\xca\x5d\x5c\x0b\xd6\x3d\x9b\x2a\x1c\x80\x17\xb0\x81\xa9\xb1\x02\x2d\xe5\x46\x29\x56\x3b\xcc\xb6\x5f\xc6\x50\x39\x43\xe3\x4e\xaa\x5c\xcd\x38\x66\x90\xd6\xbe\x3f\x6d\xfb\xd3\xc1\xae\xac\xca\x9d\xc5\x18\x68\x97\x17\xb0\xc6\x9b\xf1\x41\xaa\x5d\x34\x01\x6e\xa9\x31\x03\x6d\x34\x8e\x21\x06\x2d\x90\x1b\xcb\xbc\x34\x7a\x3b\x1d\xf2\x91\xda\xa3\xad\x8d\x8a\x0b\xb4\x32\x02\x33\x98\x4a\xde\x4c\x25\x6f\x9d\xd7\xb0\x01\x21\x5d\xad\xd8\x3a\x83\xa9\x32\xfc\x61\xbc\xc3\x6f\x74\x5b\xaf\xda\x2a\xe6\x49\xac\xfd\xe4\xea\x94\xf0\x8f\xf1\xe2\x59\x85\x42\x32\x30\x5a\xad\xc1\x71\x8b\xa8\x81\x69\x01\xdf\x55\x6c\xd5\x12\x2d\xfb\xe1\xa7\xb4\x5e\x7d\x0f\x9b\x2b\x00\x80\x67\x21\xe2\x2d\x41\x77\x44\xb8\xbd\x49\xeb\x55\x1b\x2c\xc0\xb3\xc7\x96\x8f\xe2\x3c\xbc\x47\x5d\xe0\x31\xee\xab\x7c\x15\xe4\x05\x20\x37\xd9\x9f\xe3\x35\xff\xad\xe1\x4d\x85\xda\xff\x15\xbd\x97\xba\x74\x93\xe8\x2d\x37\xd9\x73\xa5\xcc\xf2\xf5\x9f\x7e\x9f\xec\x66\x5e\xcb\x15\x2a\xf7\x1a\xed\x2b\xcd\xe7\x93\xbb\xa7\x79\x72\x3c\x17\x0c\x27\x97\x2d\xe7\x49\xf4\x7e\x36\x3c\xe5\x31\x84\x08\xa3\xd1\xfd\x63\xc0\x0e\xbb\x0b\x47\x4b\x6b\x9a\x9a\xce\xe4\xaa\x83\xb0\xc5\x24\xb0\x04\xae\x65\x15\x80\x62\xda\x8f\x8f\xb1\xe9\x39\x3e\x5b\xd1\x98\xa9\x92\xfa\x01\xe6\x16\x67\xad\xf4\xba\x2c\x49\x66\x46\x7b\x37\x2c\x8d\x29\x15\xb2\x5a\xba\x21\x37\x55\x88\xe9\xd7\x33\x56\x49\xb5\x2e\xde\x4c\x1b\xed\x9b\xec\x36\x4d\x07\x3f\xa4\xe9\xe0\xc7\x34\x1d\xfc\x2a\x4d\x09\x58\x54\x05\x89\xde\xdd\x1c\xd1\x93\x73\x29\x01\x5c\x4a\x37\x2e\xc6\xe2\xb7\x19\x41\x63\xd5\x77\x9f\x17\xd2\xf7\xe3\xd6\xea\x0e\x93\x53\xe5\xfd\x04\x56\x4b\x7d\x8e\xd5\xc3\xea\x5d\xd0\x89\xa6\xd2\xb4\x46\x4b\x47\x69\xba\xe3\x6e\xa8\x53\xaf\x4c\x27\x24\x4e\xda\xce\x93\x07\x5d\x0b\x9e\x43\x98\x42\x2e\x80\x07\x1d\x2c\x48\x34\xac\x3d\x93\x1a\x2d\x99\xf4\x18\x0e\xff\x84\x57\x2f\xef\x27\x1d\x70\x79\xd4\x1d\xb0\x46\x61\x41\x6a\x8b\x0e\xb5\x8f\x7a\x40\x3a\xd1\x29\x48\x4a\x80\xa3\x52\x9d\x5a\xee\xc6\x41\xa9\xb6\xe3\xb6\x4b\x90\xa7\xa1\x9e\x4c\xc9\x52\x17\x84\x63\x50\x17\x02\x31\xdc\xae\x8f\x64\x4f\xd3\x70\x1f\xc9\x64\x57\xb5\xdc\xdb\xfd\x20\x0c\xc5\xf6\x40\x5f\xd9\xc2\xa1\x50\xcb\x28\xbe\x71\x14\x14\xb3\xb7\x83\xda\x46\x61\x86\x2b\xc6\xbd\x5a\xef\xed\x1f\x94\x2c\x60\xd3\xd9\xee\x34\x2c\xad\x57\xad\x8c\xee\x25\x67\x1b\xe0\x23\xb0\xfc\x37\x2c\x3a\x1f\x87\xf1\xee\x4b\x3a\x3e\x41\x68\x8f\xf3\x24\xf7\xb1\x9a\x01\x95\x1e\x14\x51\xc2\xe3\xa1\xac\x3d\x33\x5e\xa0\xf5\x92\x33\xd5\xcd\x7a\x53\x8f\x85\xb4\xc8\xa3\xc2\x2b\x6f\x8f\xd0\xda\x36\xba\xa0\x86\xd0\x25\xf8\x55\x09\xf1\x21\x05\x3e\x93\xc3\x59\x92\x9c\x14\xb1\x4f\xf0\x83\x9b\x73\xa2\x7b\xe4\x11\x5f\xdb\xbe\x26\x75\xe4\x51\xdb\xdd\x2e\x81\x18\x9a\xdd\xb8\x57\x06\x85\x33\x7f\x50\xd3\xcf\x21\x4c\x77\x79\x82\xa1\x0f\x20\xc3\xd2\x58\x41\x97\x96\xd5\xd9\xd4\x22\x7b\xa0\x61\x7c\xa1\xdc\xa3\x50\xee\x9b\x1f\x03\x98\x47\x9c\xfb\x02\xfc\xbe\xf8\xc8\xe9\x16\xba\x23\x31\xa0\x8f\xa7\x7c\x87\xed\x4d\x47\xd6\xf0\x0a\x62\xca\x17\x84\x40\x7c\x5e\x87\x2f\xed\xcd\x2f\x48\xb8\xc1\x04\x9c\xe5\xfb\x56\xc4\x8d\x58\xcf\x4d\xe3\x70\xc8\x4d\x22\xb0\x32\x09\x13\x0b\xa6\x39\x0a\xea\x90\x59\x3e\xa7\x33\x63\xab\x44\x56\x65\xc2\x05\x55\xa6\x34\x43\xb7\x28\x8f\x52\xcb\xe2\x43\xa9\xcb\xc6\x32\x21\x1b\x17\x93\xd9\x12\xa7\x65\xcc\x11\x47\xb6\x8f\xaf\x78\xf6\xf8\xe9\x15\x27\x7b\xac\xe9\xd4\x2d\x4a\xd0\x9e\x06\x37\x01\x94\xc4\x8b\xf0\xc7\x86\x3f\x2d\x3e\x49\xac\x58\x7f\xe5\x4b\xf0\x62\x92\xd7\xa7\x82\x35\x0a\x02\x7b\x24\x91\x1d\x0e\xde\xd4\x61\x19\x9c\x51\x52\xc0\x13\x21\xc4\xd1\x35\x48\xea\x33\xd2\xf2\x08\xdb\x2e\x12\xe3\xe3\x75\xf6\xd3\xc3\xee\xf7\xaf\x63\x5c\xb7\x2d\x68\x7c\xd0\x90\xc6\x64\xf2\xef\x7f\x1d\x54\xa9\xad\x4e\x5f\xab\xbe\x6c\xa5\xb6\x48\x05\xe9\x21\x07\xdd\x8c\x37\xd6\x19\xdb\xd2\x88\x1b\x65\x6c\xf6\x24\x8d\xff\x5a\xa3\xed\x5b\x27\x6b\xdf\x3a\x03\xf8\x03\xaa\x05\x06\x31\x1c\xc0\x73\x2b\x99\x1a\x80\x63\xda\x51\x87\x56\xce\xfa\x6c\xbe\xa9\x57\x07\x29\xdf\xdc\x9c\x91\x40\x32\xd9\x6c\xc0\x63\x55\x2b\xe6\x11\x48\xbc\x9b\xb4\x66\xce\xc5\x1c\xf9\x9c\xe9\x12\x05\x81\x21\xbc\x7f\x9f\x27\x42\x2e\xfe\x17\x41\xe9\x7e\x4a\x7f\x4b\x58\xfe\x7f\xab\xbf\xc5\xad\x3e\x51\xe6\x58\xc9\x8b\xef\xaa\x13\xb3\xe7\xde\x37\x5f\xd9\x7a\xb0\x71\x95\x27\xed\x4f\x86\x3c\xfe\x8f\xd0\xe4\x3f\x01\x00\x00\xff\xff\xc6\x5b\xbd\xd2\x58\x13\x00\x00")

func localedefaultTemplatesPasswordChangedHtmlBytes() ([]byte, error) {
	return bindataRead(
		_localedefaultTemplatesPasswordChangedHtml,
		"localeDefault/templates/password-changed.html",
	)
}

func localedefaultTemplatesPasswordChangedHtml() (*asset, error) {
	bytes, err := localedefaultTemplatesPasswordChangedHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "localeDefault/templates/password-changed.html", size: 4952, mode: os.FileMode(438), modTime: time.Unix(1501757950, 0)}
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
	"localeDefault/templates/confirm-de.html": localedefaultTemplatesConfirmDeHtml,
	"localeDefault/templates/confirm-en.html": localedefaultTemplatesConfirmEnHtml,
	"localeDefault/templates/confirm.html": localedefaultTemplatesConfirmHtml,
	"localeDefault/templates/forgot-password-de.html": localedefaultTemplatesForgotPasswordDeHtml,
	"localeDefault/templates/forgot-password-en.html": localedefaultTemplatesForgotPasswordEnHtml,
	"localeDefault/templates/forgot-password.html": localedefaultTemplatesForgotPasswordHtml,
	"localeDefault/templates/password-changed-de.html": localedefaultTemplatesPasswordChangedDeHtml,
	"localeDefault/templates/password-changed-en.html": localedefaultTemplatesPasswordChangedEnHtml,
	"localeDefault/templates/password-changed.html": localedefaultTemplatesPasswordChangedHtml,
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
	"localeDefault": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"confirm-de.html": &bintree{localedefaultTemplatesConfirmDeHtml, map[string]*bintree{}},
			"confirm-en.html": &bintree{localedefaultTemplatesConfirmEnHtml, map[string]*bintree{}},
			"confirm.html": &bintree{localedefaultTemplatesConfirmHtml, map[string]*bintree{}},
			"forgot-password-de.html": &bintree{localedefaultTemplatesForgotPasswordDeHtml, map[string]*bintree{}},
			"forgot-password-en.html": &bintree{localedefaultTemplatesForgotPasswordEnHtml, map[string]*bintree{}},
			"forgot-password.html": &bintree{localedefaultTemplatesForgotPasswordHtml, map[string]*bintree{}},
			"password-changed-de.html": &bintree{localedefaultTemplatesPasswordChangedDeHtml, map[string]*bintree{}},
			"password-changed-en.html": &bintree{localedefaultTemplatesPasswordChangedEnHtml, map[string]*bintree{}},
			"password-changed.html": &bintree{localedefaultTemplatesPasswordChangedHtml, map[string]*bintree{}},
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

