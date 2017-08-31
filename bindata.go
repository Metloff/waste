// Code generated by go-bindata.
// sources:
// assets/app.js
// assets/statistic.html
// DO NOT EDIT!

package main

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

var _assetsAppJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x57\xdb\x6e\xdb\x36\x18\xbe\xae\x9f\xe2\x87\x6f\x24\x2f\x9a\x2a\x29\x71\xd2\xc8\x75\x81\x34\x5d\xb1\x16\x4d\x0a\xac\x05\x86\xcd\x30\x0a\x5a\xa2\x1d\xa6\x34\x25\x50\x54\x62\x63\x30\xb0\xee\xa6\x18\xf6\x0e\x1b\xf6\x06\xc5\xd6\x62\xdd\xb0\x65\xaf\x20\xbf\xd1\xf0\xcb\x92\x75\x8a\x5d\x8f\x17\xa4\x25\x7d\xff\xe9\xe3\xe9\xf3\x15\x91\x30\x0d\x84\xba\x88\xa0\x0f\x03\xed\x29\x11\x9a\x01\xda\x63\x3a\xc2\xe1\x8c\x48\x1c\x4e\x42\xb9\x7a\x9a\xe3\xf0\x34\x16\x74\x35\xf2\xf4\xf9\x24\x9e\xe0\xf0\x82\x86\x0a\xc7\xe7\x5e\x3a\x9c\x07\x57\x38\x3c\xa2\x9e\x36\xec\xb5\x8a\x28\xe7\xf1\x74\x44\x65\x1a\xcb\xb2\x35\x43\xb3\x1c\xec\xf6\xb1\x3b\xc0\xae\x8b\xdd\x21\x76\x47\xd8\xdd\xc3\xee\x58\x33\x34\xdb\xc2\x0e\x2d\x6c\x27\xf7\xc8\xc9\x88\xf2\xe8\x21\x49\xdd\x65\xef\x7c\xa2\x48\xf5\xcd\x0a\xf5\x8c\x09\x5a\x83\x95\x5e\xa5\xef\x42\x49\x43\x22\xe9\x23\xa2\x08\xf4\x61\x1c\x0b\x4f\xb1\x40\xe8\x1d\xf8\xae\x05\x00\x30\x0e\xa4\x8e\x30\x06\x7d\xb0\x7a\xc0\xe0\x3e\x5c\x33\xe1\x07\xd7\x26\x31\x4f\x63\x79\x86\xe5\x99\x9c\x8a\x89\xba\xe8\x01\xdb\xdb\xcb\xed\xb0\xad\x33\x35\xc3\x38\xba\xd0\x1b\x76\x03\x36\x34\x4f\x89\xa2\x93\x40\xce\x3b\xbd\xb5\x59\x56\xcc\x16\xa3\x93\x69\x10\x0b\x95\x99\x2c\x00\x5a\x8d\x54\xed\x34\xd5\x3e\xd8\x4e\x9e\x55\xee\x7d\x1d\x26\x87\x5f\xae\x2a\xbb\x2c\x57\xf6\x0d\x25\xf2\x85\x22\x6a\x5d\xd9\x65\xb5\x32\x6c\x5e\x20\xa2\x80\x53\x93\x07\x13\xbd\x61\x38\xb8\x1c\x9a\x69\xba\xa5\xba\xb0\xb1\xf1\x16\x2c\xf4\xfb\xc0\xea\x61\x32\x42\x70\xda\x6a\x8c\x94\xed\x2b\x8c\x94\x9a\xa4\x2a\x96\x02\x2a\xaf\x17\xad\xcd\xce\xad\x4e\xaf\x44\xd1\x0a\xb9\x68\x2d\xb2\xb5\xe2\xa9\x19\xf4\xc1\x0f\xbc\x78\x4a\x85\x32\x27\x54\x7d\xc1\x29\xfe\x7c\x38\x7f\xe2\xeb\xed\xe9\xfc\xf4\x82\x48\xd5\xee\xf4\x5a\xa5\x55\xa5\x77\xb2\x9d\xb0\xfa\x0a\x7d\x10\xf4\x1a\xd2\xdf\xba\xa7\x66\x46\x56\xaf\x9a\x87\xd4\x05\x6d\x84\x9b\xaf\x95\x27\xe6\x36\x56\x93\x5b\xac\x2a\xa3\xb2\x62\x22\xaa\x22\x17\x06\x55\xf2\x52\xac\x0b\x5a\xf2\x4b\x72\xb3\xfc\x61\xf9\x7d\xf2\x6e\xf9\x36\xf9\x90\xfc\x93\xdc\x40\xf2\x3e\xfd\xf1\x21\xf9\x5d\x33\x1a\x84\xb8\xf9\x22\xac\x7e\x1a\x11\xef\xf5\x44\x06\xb1\xf0\x4f\x03\x1e\x48\x17\x06\x75\xbe\x41\x93\x93\x11\xd1\x9d\x43\x03\xec\xa3\x63\x03\xec\x83\x7b\x06\x58\xe6\x41\xa7\x16\xa4\x80\x76\x0f\x0c\xb0\x0f\x1d\x03\x9c\xfd\x2e\x42\x9d\xcd\x50\xa7\xdb\x35\xc0\xb1\x0e\x0d\xb8\x77\xf8\x09\xe8\x51\xd7\x00\xfb\xd8\xc9\xba\xad\x50\xbb\xbb\x6f\x80\x6d\x61\x06\xdd\x9d\x32\xb0\xbb\xc7\x06\x1c\x1e\x64\xd0\x0a\x72\x58\xe3\x2b\x90\x3e\x95\xff\x93\xab\xa3\x5d\xb9\xb2\x77\x65\x6a\x0b\xb0\xca\xd3\x16\x60\x8d\xa5\x4f\xc5\x5e\x73\x64\xef\xc2\xd0\xd7\xcc\x57\x17\x2e\xd8\xc5\xc6\x1b\xae\x76\xde\x0a\x1c\x84\x78\x24\x47\xe5\xcd\x20\x69\x14\x06\x22\x62\x57\xd4\x05\x25\x63\x5a\x78\x8d\x3c\xc2\x69\x05\x8b\x6d\x7e\x32\xa3\xe9\xf6\x00\xc5\xbc\xd7\xf8\x19\x46\x74\xc2\xc4\x89\xfa\x96\xca\xc0\x45\x17\xb0\xc8\xc3\x96\x42\x63\xe3\x74\x42\x85\x5f\xf7\x18\x06\x11\xc3\xb4\x5c\xd0\x54\x10\x96\xd8\x28\x59\x2a\xa6\x38\xad\x1b\xfa\x2c\x0a\x39\x99\xd7\xf3\x4e\xf1\x74\xa6\x5c\xd0\xd2\xb3\xc1\xbc\x8c\x00\xaf\xb4\xf4\x41\xbb\xcd\x7b\xeb\xce\x9d\x40\x9c\x72\xe6\xbd\x76\x8b\x7b\x8b\x1a\x59\xbe\x4f\x14\x9d\xd6\xcf\xd2\xf4\x0c\xcb\x2e\x9c\x57\xd7\x24\x52\x78\x13\x36\x6f\x98\xc2\xc1\xc0\x1a\x9a\xaf\x98\xf0\xe9\x6c\xd8\x70\x14\xf3\x2d\x67\xa1\x96\x87\xc1\x03\x5a\xab\x9d\xcb\x59\x1a\xe7\x64\x4a\x77\x71\x81\x38\x74\x51\xf1\x71\xf7\x2e\x24\x3f\x2f\xdf\x26\x1f\x97\x3f\x26\xef\x92\x0f\xc9\xdf\xb0\x7c\x93\xfc\x9b\x7c\x5c\xbe\x49\x6e\x92\xbf\x72\x86\xca\x16\x31\x37\x99\x10\x54\x7e\xf9\xf2\xec\x19\xf4\x41\xd3\xaa\x49\x65\x09\x6d\xc5\x14\xad\x91\xcb\xaf\xc9\x4d\xf2\x47\xf2\x3e\xcb\x05\x4f\xd8\xe4\xb7\xe5\x4f\xc9\x9f\x95\xb4\x2a\x56\xe3\x40\x42\x5d\x5a\x54\x27\xc7\x7c\xfa\xe2\xf9\xf9\x26\x71\x51\x26\x93\xb3\x32\x8f\x9e\xa4\x44\xd1\x8c\x4a\x5d\xe3\xac\xce\x3f\x36\xce\x2a\x95\xb6\xef\xfb\xec\x0a\x3c\x4e\xa2\xa8\xaf\x45\x8a\xa8\xaf\x82\x6b\xed\x41\x1b\xf6\x6e\x4b\x09\x25\xc8\xd8\x81\x3d\x68\xbb\xf0\xf9\x36\x8c\x8d\x98\xfb\x77\x7d\x76\xf5\x00\xa0\x1c\x61\x22\xc9\x5c\x7b\xd0\xde\x43\x11\x71\x45\xa5\x7a\x44\x14\xd5\x37\x79\xd9\xef\x14\x6e\xda\xcd\x4a\x62\x6e\x92\x30\xa4\xc2\x3f\xbd\x60\xdc\xd7\x39\xab\x55\xbb\xf8\xe4\x34\x37\xdf\x35\x4a\xca\x65\x5a\xaf\x36\xf5\x65\x11\x54\x35\x41\x21\xb1\x79\xd9\x64\xdb\x78\xd1\x5a\x74\x72\x6d\x51\x30\x51\xd6\xa1\xb1\x60\xb3\x97\x6c\x4a\xf3\xa9\xcf\x64\x2c\xcd\x64\x44\x4a\x5c\x8e\xf9\xcc\xb6\x2c\x2b\xab\x1e\x71\x73\x9a\x4a\x62\x84\xe3\xe6\x7a\x1c\x73\x8e\x9a\x49\x2f\x41\xa6\x2b\xdd\x55\x91\xe9\x83\xdc\x20\x3d\x0d\xf4\xce\xb0\x57\x8a\x3c\x2f\x39\x4c\x83\x67\xce\x32\xa1\x85\x80\x3d\x68\xa7\xab\x62\xe5\x3b\x7f\xc2\x64\x52\x21\x55\xa6\xac\x50\xe9\x9d\x5e\x6b\xbb\xa0\xb4\x0a\x41\x99\x6b\x31\x67\x07\x31\xe6\xb4\xd7\xf2\x0b\xc3\xdc\x2a\xc1\x9c\x9a\x06\xe3\x0c\xff\xeb\x6c\x15\x61\x2b\xbe\x6e\x51\x60\xb7\x0b\xb0\xf6\xd9\x1c\x1e\x33\x19\x29\x9c\x30\x04\xb6\x37\x69\x2e\x4c\xb2\xfa\x6d\xcc\x38\x77\x61\x4c\x78\x44\xb7\xa8\x0b\xbc\x7b\x2b\xb7\x79\xfd\x82\xc6\xa2\x5e\x52\x11\xa5\x57\x96\x65\xda\xc5\xb6\x18\x96\xee\x94\xe2\xaa\x6d\x5c\xb1\xe9\x5a\xfd\x2f\x00\x00\xff\xff\x67\x4b\x8e\xb8\x39\x0e\x00\x00")

func assetsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsAppJs,
		"assets/app.js",
	)
}

func assetsAppJs() (*asset, error) {
	bytes, err := assetsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/app.js", size: 3641, mode: os.FileMode(420), modTime: time.Unix(1504161519, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsStatisticHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x5a\xe9\x8e\xdb\x46\xb6\xfe\xdf\x4f\x71\x42\xff\x50\xb7\x5b\x24\x45\x2d\xdd\xb6\x5a\x12\xd0\x71\xe2\xeb\x38\xb1\x73\x61\x3b\xb9\xf0\x35\x1a\x41\x89\x2c\x49\xd5\x2e\x15\x89\x62\x51\x4b\x82\x06\x62\xe3\xe2\x06\x83\xfc\x98\x27\x98\x19\x64\x9e\xc0\x33\xb1\x27\x9d\xc5\xce\x2b\x50\x6f\x34\xa8\x22\xb5\x90\x2c\xaa\xd5\x49\xfe\x0c\x7f\x14\xb7\xb3\xd5\x77\xaa\x4e\x9d\xc3\x62\x67\x24\xc6\xb4\x07\xb0\x07\x00\xd0\x19\x61\xe4\xf5\xd4\xa5\xba\x15\x44\x50\xdc\x8b\xbf\x5b\xbc\x8c\x5f\x2d\x5e\xc6\x97\x8b\x17\xb2\x8d\x7f\x8a\x5f\x75\xec\xe4\xdd\x9a\x36\x74\x39\x09\x04\x84\xdc\xed\x1a\x23\x21\x82\xb0\x6d\xdb\xae\xc7\xce\x43\xcb\xa5\x7e\xe4\x0d\x28\xe2\xd8\x72\xfd\xb1\x8d\xce\xd1\xcc\xa6\xa4\x1f\xda\x77\x46\x88\x0b\xeb\x3c\xb4\xeb\x56\xcb\xaa\xa5\xb7\xfd\x88\x79\x14\x5b\x63\xc2\xac\xf3\xd0\xe8\x75\xec\x44\xf0\x86\x26\x4a\xd8\x73\xe0\x98\x76\x8d\x50\xcc\x29\x0e\x47\x18\x0b\x03\x46\x1c\x0f\xd6\x9a\xc7\x68\xe6\x7a\xcc\xea\xfb\xbe\x08\x05\x47\x81\xbc\x91\xca\x57\x0f\xec\x86\xd5\xb0\x8e\x6d\x37\x0c\xd7\xcf\x94\x52\x37\x0c\x0d\x20\x4c\xe0\x21\x27\x62\xde\x35\xc2\x11\x6a\xdc\x6a\x9a\xef\x7f\xfe\x94\x90\xc7\x1f\xdd\xc5\x1f\x3b\xde\x7f\x8d\xef\x3f\x3a\x7d\x3e\x77\xa3\x7b\xa7\xf7\x1e\x0d\x1b\xf5\x4f\xc7\x9f\xb9\xd3\xe9\xb1\xcf\x1a\x8f\x9e\x7a\xc3\xe6\xe7\xe8\xf0\xbf\xc7\x8f\x9f\x84\x5f\xda\x1f\x1f\xdd\x9a\xf4\xbd\x0f\xcf\x47\xcd\xc8\x00\x97\xfb\x61\xe8\x73\x32\x24\xac\x6b\x20\xe6\xb3\xf9\xd8\x8f\x42\x63\xa3\x63\x63\x2c\x10\x30\x34\xc6\x5d\x63\x42\xf0\x34\xf0\xb9\x30\xc0\xf5\x99\xc0\x4c\x74\x8d\x29\xf1\xc4\xa8\xeb\xe1\x09\x71\xb1\xa9\x6e\xaa\x40\x18\x11\x04\x51\x33\x74\x11\xc5\x5d\x47\x09\xeb\xd8\x4b\x2f\x76\xfa\xbe\x37\x07\x97\xa2\x30\xec\x1a\x43\x8e\xe6\x66\x7f\x68\x80\x02\xad\x6b\x8c\x11\x1f\x12\xd6\x86\x5a\x6a\x41\xc7\x23\x93\x25\xad\xd4\x89\x08\xc3\xdc\x1c\xd0\x88\x78\x9b\x36\x6e\x50\x05\x68\x88\x4d\xa9\x0b\xb8\x3f\xdd\xa0\x49\x46\x52\x3d\xfb\x40\x3d\x44\xa9\x97\x6e\x18\xbd\xa7\x7e\xc4\xe1\xb1\x40\x82\x84\x82\xb8\x1d\x1b\xf5\x20\x2b\xc0\x1e\xd5\x7b\xf2\x62\xad\xda\xf6\xc8\xa4\xb7\xa7\x35\x65\xca\x51\x80\x79\xc6\x86\xcd\xd7\xca\x3e\xc8\x29\xc8\xf4\x97\x9a\x63\xcf\xbc\x65\xe4\x6d\xde\x24\x22\x7d\x7f\x66\xe4\xac\xd4\xd1\x98\xa9\xc7\x0a\xc2\x14\xb1\x8b\xd8\x04\x85\x40\xbc\xae\x31\x9e\xab\x11\x6f\xc0\x18\xcd\x12\x87\x76\x8d\x46\xad\x96\xdc\x8f\x30\x19\x8e\x44\xf2\xa0\xd7\xb1\x13\xb6\x2c\x20\x39\x60\xae\x7c\x98\x07\xb0\x04\x85\xe6\x0e\x28\x5c\x0d\x82\x0a\x0f\x7a\x08\x46\xad\xde\x13\x8e\x58\x88\x5c\x41\x7c\x16\x76\xec\x51\x4b\x4b\x17\x06\x88\x2d\x65\x52\xd4\xc7\x14\x54\x6b\x7a\x78\x80\x22\x2a\x8c\xde\x67\x4f\xee\x1c\x36\x3a\xb6\xa4\xd3\x58\xa4\x43\xa5\xdc\x5b\xc0\x7c\x33\x40\x9e\x47\xd8\x50\x6f\xb5\xe4\x23\x5e\xb7\xe2\x22\x81\x87\x3e\x9f\x3f\x44\x63\x5c\xe9\x9d\x3f\x1f\x9d\x87\xa3\x12\x5d\x00\x9d\xf7\x4c\xb3\x13\xd1\xa5\x42\x86\x26\xc0\xd0\xc4\x14\xa8\x1f\x1a\x3d\xd3\xd4\xb1\xa4\x4c\x94\x2c\x99\x24\x4c\x13\x6c\xf4\x36\x66\xcf\x72\x0e\xf7\x7d\xee\x61\x6e\x52\x3c\x10\xa6\xeb\x53\x9f\xb7\x41\x48\x64\x03\xc4\x65\x97\xde\x23\x63\x19\x43\x10\x13\x27\x90\x92\x72\xe4\x91\x28\x6c\x43\x4d\xf7\x52\xf8\xc1\x15\x62\x8c\x5e\xfc\xb7\xf8\xdd\xe2\x45\xfc\x73\xfc\x26\x7e\x1d\xbf\x8d\x2f\xe3\x37\xb0\x78\xb9\xf8\x5a\xae\x0f\x8b\x6f\xe5\x34\xee\xd8\x94\xfc\x8e\xae\xed\xa8\x40\x2f\x9d\x92\xac\xa8\xbf\x27\x7c\xdb\xb8\x3a\x76\x44\x4b\xcc\x95\x8e\xdb\xf4\xb8\x8c\x56\x95\x72\xcd\xf1\x5f\xe2\x57\xf1\xbf\xe2\x5f\xe2\xcb\xc5\xcb\xf8\x0d\xc4\x6f\xe3\x57\x10\x7f\xaf\x2c\xf8\x3f\xb9\x68\x42\xfc\x3a\xfe\x79\xf1\x67\x88\x7f\x8d\xdf\xc5\xaf\x17\x5f\xc7\xef\xe2\x7f\xc4\x6f\xe3\x77\xf1\x8f\xf2\xcd\x1b\xb9\xc6\xc6\x3f\xc7\x97\xf1\x0f\xf1\xab\xc5\xff\xc7\x97\xf1\xe5\x56\x93\x7f\x6f\x14\xc8\xfa\x84\xfb\x53\x70\xa9\x1f\x62\xd3\xec\x6d\x12\xef\x6d\xe3\xcb\x07\xd8\xab\x02\x8b\x53\xff\x43\x22\xcb\x35\xc3\x6b\xfd\x9a\xf1\xf5\x0f\x04\x36\xfb\x60\xe3\x36\xf7\x26\x9f\x44\x65\xe5\x4e\x09\xf3\xfc\xa9\x85\xa0\x0b\x5f\x7d\x65\x5d\x5c\x6c\x08\x2c\xa4\x46\x45\xfe\x09\xe2\x30\xf6\x99\x18\x85\xd0\x85\x67\x95\xfb\x88\x55\xaa\x50\xb9\x8b\xfb\xf2\xf4\x00\x71\x79\x3a\x0d\x78\x72\x37\x97\xa7\xfb\x11\xc3\xc9\x99\xaa\xfb\xd3\x68\x28\x4f\x8f\x71\x20\xe4\xf9\x53\x57\x9d\x1e\xfa\x13\x79\xfa\x00\xbb\x95\xb3\x13\xad\xbe\x87\xd1\xb8\x8f\xb9\xd2\x5a\x73\x2a\xd5\x4a\xad\x2e\x9b\x86\x6c\x9a\xb2\x69\xc9\xe6\x48\x36\xc7\xb2\xb9\x25\x9b\xdb\x95\x6a\xc5\xa9\xc9\x46\x72\x38\xf5\xa2\x6c\xb5\x04\x84\xef\x23\x25\xb8\xf0\xd6\x43\x02\x95\xbd\x4b\x38\x3f\x21\x0c\x97\xb2\x6e\xbc\xcc\xbd\x0d\x38\x96\xd1\xf0\x03\x24\xa4\x1f\x06\x11\x53\x6b\xd7\xfe\x01\x7c\x95\x1b\x18\x03\x9f\xef\x4b\x06\x02\x5d\xa8\x9d\x00\x81\xce\xca\x81\xd6\x9d\x88\x3f\x90\xd0\x58\x14\xb3\xa1\x18\x9d\x00\x39\x3c\x2c\x4a\x90\xc7\xaa\x97\x56\x10\x85\xa3\xfd\x82\x84\x67\xe4\xcc\xba\x93\x86\xa5\x83\x13\x8d\x80\x14\x88\x2d\xec\xa7\x63\x3f\x62\xa2\xc0\x7c\x01\xb0\xb7\xa5\x4b\x8e\xea\x52\x17\x9c\xfa\xd2\xfa\x25\x95\xc6\x88\x25\xe3\x79\x82\xc5\xf9\x26\x16\x4f\x31\xe2\x32\xa0\xae\xb0\x38\x2f\xc3\x42\x1e\xae\xcf\x42\x9f\x62\x8b\xfa\xc3\xfd\x82\x88\x67\xe7\x67\x96\xea\x96\x16\x09\x79\x90\xc1\x16\x2e\xe8\x76\x81\x94\xab\x86\xd5\xc8\xc8\x81\xb9\x29\xa8\x04\xcc\xf5\xc1\xb1\x88\x38\xd3\x81\x24\x8f\x8b\x92\xe7\x59\xcd\xb5\x83\x13\x2d\xcc\x79\xee\xcd\xfb\x8b\xc2\x48\x76\xc5\x0c\xba\xe0\xf9\x6e\x34\xc6\x4c\x58\x43\x2c\x3e\xa4\x58\x5e\xbe\x3f\xff\xc8\xdb\x5f\x25\xa5\x99\xbe\x6c\x8c\xfe\xfd\x83\xc2\x6c\x4f\x38\xa0\x0b\x0c\x4f\x41\x5d\xef\xbb\x62\x56\x2d\x20\x2a\xe6\x01\x6e\x43\xa5\x2f\x83\xce\x5e\xb1\xa3\xed\x2d\x33\xa1\xbd\x9e\x11\x79\xd6\x25\x7b\x88\x45\xd8\x86\x67\x65\x6e\x54\xfc\x6d\xa8\xa8\xf4\x22\xc9\x27\xbe\x89\xdf\xc8\xf5\x57\xad\xbe\xf1\xdb\xf8\x4d\xfc\x7d\xc1\xae\xac\x7d\xe9\xa4\x2a\x23\xea\x23\xf7\xf9\x90\xfb\x11\xf3\xee\x24\x49\xd4\xb3\xd2\xe1\x00\x50\xe1\xc3\x3e\xda\xaf\x1f\x55\xc1\x39\xbe\x5d\x05\xa7\x79\xab\x0a\x35\xab\x79\x50\x6a\xc2\x9a\xa9\xd5\xac\x82\x73\x54\xaf\x42\xbd\xd1\x92\x4c\xf5\x5d\x98\xea\xad\x56\x15\xea\xb5\xa3\x2a\xdc\x3a\xda\x99\xe9\xb8\x55\x05\xe7\x76\x3d\x6d\x76\x64\x72\x5a\x8d\x2a\x38\x35\x69\x5f\xeb\x9a\xf6\x39\xad\xdb\x55\x38\x6a\xa6\x4c\x25\x3c\x67\xa5\xf8\xab\x2c\xf6\x37\x63\x7f\x7c\x7d\xec\x9d\xeb\x23\xbf\x13\x4b\x16\xf7\x9d\x58\x72\xa8\xef\x6e\xd9\x0a\x73\xe7\xb7\x22\xfe\x3f\x32\xb5\x6a\x83\xa3\x0b\x4d\x67\xf9\xd8\x94\x17\xe5\x07\xaa\x08\xd4\x4f\x7e\x8e\xc3\xc0\x67\x21\x99\x60\x59\x91\x44\x58\x67\x87\xfa\xf4\x51\xc2\x2f\x8f\xf9\xe9\x0c\xab\xc0\x00\x82\xb8\xcf\x25\x21\xf4\xf1\x90\xb0\x53\xf1\xbf\x98\xfb\x6d\x29\x16\x2e\x8a\x86\x6a\x8d\x95\x07\xc5\x43\xcc\xbc\x72\x7d\x81\x1f\x12\xd9\xa5\x36\x54\x84\x1f\x68\xbd\xa0\x95\xab\x0a\xe6\x72\xb1\x1e\x09\x03\x8a\xe6\xe5\x38\x28\x19\x78\x26\xda\x50\x59\x7e\x57\x03\x99\x0a\xa9\x1b\x9d\x67\xb5\x56\xe4\xbd\xc3\xee\x50\xe2\x3e\x6f\xaf\x33\x1e\x5c\x4d\x11\xf8\x48\xe0\x71\xf9\xaa\xa9\x56\x9a\x34\x41\xf9\x62\x8a\x42\x21\xf3\xaa\x62\x1e\xb2\x16\xf5\xac\x76\x66\x7d\x41\x98\x87\x67\x3a\x4f\x2c\x45\x46\x74\xcb\xda\x95\x2d\xd4\x4a\x17\xe3\xd4\x34\x59\xbf\xef\x22\x4c\xd5\xf9\x07\x27\xf9\xa4\x68\x79\xd8\x36\xc4\x7f\x5d\x7c\x13\x5f\x2e\xfe\x14\xbf\x8a\xdf\xc4\xbf\xc0\xe2\x45\xfc\x6b\x7c\xb9\x78\x11\xbf\x8b\x7f\xd2\xa3\xba\x3c\x22\x6a\x11\xc6\x30\xbf\xf7\xe4\xc1\x27\xd0\x85\x4a\xa5\xcc\xe4\xd4\xdc\x1d\xa9\xd7\xc7\x16\x9b\xbf\x8b\xdf\xc5\x3f\xc4\xaf\x53\x9b\x55\x1d\xfa\xcf\xc5\xb7\xf1\x8f\x19\xf3\x4b\xf8\x07\x3e\x87\x7c\x92\x9b\x75\xb6\x75\xff\xf1\xa7\x0f\xaf\x4e\x73\x97\x87\x4a\xcf\xc9\xa6\x37\x5c\x8e\x91\xc0\xa9\x43\xf6\x2b\x94\x94\xfb\x53\x1e\x94\x64\xb0\x31\x36\x8a\xc6\x4a\x28\x90\x78\xe4\x4f\x2b\x3d\x03\x0e\x75\x66\xca\x64\x78\x50\x87\x43\x30\xda\x60\x6e\xa3\x71\x24\x4d\x52\xba\x65\xca\xd2\xca\x90\xa3\x79\xa5\x67\x1c\xca\x34\x75\x82\xb9\xf8\x00\x09\xbc\x5f\x26\xa5\x71\xb0\x16\x63\x6c\xeb\x53\x44\x2d\x14\x04\x98\x79\x77\x46\x84\x7a\xfb\x94\x94\x22\x50\x96\x3e\xea\x86\x4d\xf1\x59\xa1\xc3\xcb\xc2\x42\x9f\x6d\x42\x2e\x1d\xcf\x32\xcb\x14\xb5\x70\xec\x12\x7c\x32\x49\xeb\x41\x31\x6b\x5d\x23\xbb\x59\x7f\x45\x8c\xcc\x9e\x90\x31\x2e\x0e\xaf\xb4\xa4\xc3\x69\x5a\xaa\x5c\xb2\xa4\xbe\xe9\xd4\x6a\xb5\x02\x9a\x92\x63\x8e\x55\xf1\x28\x19\x65\x40\xb8\x1b\x51\x2a\xd3\xfc\x7d\x2d\xf1\x38\xa9\x1e\x32\xe5\xee\xb3\x25\xab\x8a\x6f\xfb\x07\x67\x3a\x46\x0f\xcd\x37\x94\x28\xd3\x0a\x0a\xd2\x82\x41\x92\x1e\x82\xa1\xc6\x65\xa2\x6f\x79\x27\x4d\x2d\x4d\xf4\x37\x1d\xb4\xae\x7b\x33\x4a\xb6\x97\x54\xb5\x75\x49\x55\xac\x1f\xea\x3b\x14\x10\x75\x43\x53\x28\x48\x23\xb4\xc5\x42\xbd\xb4\x5a\xa0\x84\xe1\xdf\x50\x2e\x24\x3e\xd9\x5a\x2b\x5c\x55\x2a\x18\x0f\xe6\x70\x97\xf0\x50\xc8\xc1\x23\x59\x8c\xab\xab\x03\xd9\xc1\x32\xaa\x01\xa1\xb4\x0d\x03\x44\xc3\x52\x92\x4c\x06\x2b\x73\xb4\x4c\x26\x58\x9e\xd2\x49\x90\x9e\x60\x16\xaa\xb4\xa3\x66\x39\xba\x80\x70\xa6\x5d\xef\xd7\xe9\x57\x21\xd9\xca\xcf\xc8\xe5\x75\xf1\xb3\xd3\x5e\xee\x0c\x1d\xf5\x7d\x3a\xfb\x19\x4b\xed\x45\x15\x21\x1f\xf8\x4c\x98\x03\x34\x26\x74\xde\x06\xc3\x0f\x30\x83\x10\xb1\xd0\xa8\x82\x71\x0f\xd3\x09\x16\xc4\x45\xf0\x10\x47\xd8\xa8\xc2\xea\x41\x15\x4e\x39\x41\xb4\xaa\x48\xcd\x10\x73\x32\xc8\xce\x9f\x2c\x00\x23\xa7\x0a\xa3\x7a\x99\xf2\xa9\xfa\x12\xd8\x06\xa7\x56\x2b\x06\x58\x45\x11\x92\x2f\x71\x1b\xea\xcd\x60\x56\x24\x48\xf7\x0c\x4c\xe1\x07\x6d\xa8\xd7\xf2\x24\x59\x43\x2c\xd2\xf7\x67\x1a\x3b\x5c\x8a\x11\x6f\x43\xdf\x17\xa3\xa2\x86\x64\xa3\xce\xec\xfb\x42\xf8\xe3\x36\xd4\x5b\x3a\x33\x52\x22\x65\x85\xa6\x1b\xa9\x95\x85\x77\x1a\xf3\x56\x1b\x22\x45\x33\xd7\x05\xee\x72\x9b\x60\x3a\x22\x02\x17\xd5\xa5\x6f\x09\x1b\x61\x4e\xc4\x16\x73\x9c\x56\x30\x53\xa8\xad\x9b\x22\x75\xba\x3b\x91\x0a\xbd\x81\x8f\x31\xc2\x6e\x29\x19\x19\xa3\x21\x6e\x03\xf3\x99\xc6\xb0\x94\x46\x0d\xcf\x36\x84\x3e\x25\x5e\xda\x6e\xa5\x9f\xa6\x15\x4e\x30\x83\x5d\x3c\xbc\x82\xd0\x5a\x6e\xf6\x14\xc0\x4c\x05\x2f\x9d\x2a\x25\x27\x76\x68\xbb\x77\x91\xcd\x3e\xed\x9b\x65\x6a\x28\x29\xd1\xa4\x03\xe4\xe2\xa6\xbd\xb7\x93\xf1\x94\x80\x95\x6c\xd5\x00\x2a\xca\x2f\x0c\x0a\x8e\xbd\x3c\x92\xf6\xcd\xa5\x19\x57\xf4\x34\x87\x8c\x66\x3b\xaa\x84\xbe\x64\x0f\xac\x84\x5a\xbf\xd3\x95\xcb\x5d\xf4\xac\xab\x1d\xb4\x93\x9b\xf6\x36\x27\x65\xc1\xbc\xb1\x59\x9b\xe8\xdc\x64\xdf\xd4\x03\x99\x53\xb2\xf6\xe7\xae\x40\x5e\x0f\x18\x2e\x63\xe1\xef\xc6\xf1\x2a\xe4\x72\xef\x57\xe1\xe0\xb8\x38\xb9\x32\xb7\x59\x14\xf3\x10\x52\x12\x8a\x64\x6a\x9b\x49\xf2\xa0\x9b\xd3\xcb\x78\x2d\x61\x49\xe2\x4f\xae\xb4\xbb\xd0\x6b\x54\xe5\x62\x21\x47\xc1\x33\x61\x22\x4a\x86\xac\x0d\x2e\x66\x02\xf3\xbc\xba\xcc\x0a\x73\x54\x5c\x61\x56\x5d\xd7\x44\xf5\xff\x38\x4f\x6b\x67\x80\xfa\xae\xb1\x43\xd4\xd0\x2e\x25\x3b\x85\xfe\xab\x03\xff\xf5\xc2\x7e\x36\xe8\x37\x92\xa0\x5f\x1c\xb5\xdb\x17\xb9\xdc\xaa\x5d\xee\x78\xa7\x19\xcc\x92\x75\xb0\x30\xfa\x01\xc6\x84\xa5\x3b\x95\x6d\x68\xde\xba\x62\x76\x6c\xe2\x3d\x6a\x15\x20\x5f\x7d\x41\x22\x4c\xe6\x8b\x66\x9f\xfa\xee\x73\xed\x78\x4d\xf2\x1d\x47\x93\xef\xac\x7e\x1a\x82\x9a\xce\xdc\xb2\x34\x23\x9d\x29\xfe\x04\xf3\x01\xf5\xa7\x6d\xc0\x94\x92\x20\x24\x61\x41\x3d\xf5\x91\x68\x83\x1c\xc9\x57\xaf\xb2\x49\x4f\xad\xe4\xd7\x90\xc2\x3e\x5c\x99\xa4\x95\x6b\x92\x08\x50\xe8\xe4\x16\x50\xad\xcc\xff\x27\x3b\x8c\xe9\x1b\x0e\xea\x37\x6e\x37\x4b\x46\x8e\x66\xc4\xe7\xba\x99\xfc\xd6\x54\x50\xb4\x82\x59\x25\x4d\x8e\x6c\x9a\x57\xe5\x25\xeb\xdf\x5b\xca\xc5\x65\x7e\x0a\xd9\x26\x2c\xfd\x97\x6c\xa7\x14\xf1\xc6\xa0\x31\x68\x0c\x9a\x57\xca\xd3\xa5\xc5\x89\x04\xf9\x76\x6b\x66\x1e\x8e\x11\xa5\xc5\xd8\x5b\x32\xe4\xf3\x44\x5b\xbd\x9f\x7e\x3c\xd2\x18\xa7\x95\xbd\xa5\x93\xc5\xcc\x69\xd5\xc1\x1b\x8d\xc6\x31\xea\x1f\x17\xed\x57\xb3\xc6\xc3\xae\xcf\x51\xf2\x49\x59\x93\xc4\x65\xed\x5d\xff\x93\xb7\x9b\x6f\xd4\xb1\xad\x5e\x58\x0e\xb0\xd5\x50\x2b\x12\xdb\x37\x33\x33\xca\x74\xb6\x55\x28\x3c\x09\x66\x09\x51\x21\x83\x4a\x2f\x36\x6b\xcd\x65\x25\xd9\xb1\x65\x05\xd9\xdb\xeb\xd8\xea\x8f\xd5\x7f\x07\x00\x00\xff\xff\x4a\x4c\xef\x6e\xb8\x2a\x00\x00")

func assetsStatisticHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsStatisticHtml,
		"assets/statistic.html",
	)
}

func assetsStatisticHtml() (*asset, error) {
	bytes, err := assetsStatisticHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/statistic.html", size: 10936, mode: os.FileMode(420), modTime: time.Unix(1504161558, 0)}
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
	"assets/app.js": assetsAppJs,
	"assets/statistic.html": assetsStatisticHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"app.js": &bintree{assetsAppJs, map[string]*bintree{}},
		"statistic.html": &bintree{assetsStatisticHtml, map[string]*bintree{}},
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
