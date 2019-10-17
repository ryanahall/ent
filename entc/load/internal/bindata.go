// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package internal

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

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x51\x5d\x6b\xdb\x30\x14\x7d\xb6\x7e\xc5\x99\xe9\xa8\x5d\x52\xa5\xed\xdb\x06\x79\x28\x6d\x06\x19\x5b\x3b\x48\x61\x0f\x5d\x29\x8a\x7d\x9d\x88\x3a\x92\x77\xa5\x94\x05\xa1\xff\x3e\x24\x27\x61\x7b\xb2\xa5\x73\xee\xf9\xd0\x0d\x61\x7a\x21\xee\xec\xb0\x67\xbd\xde\x78\xdc\x5c\x5d\x7f\xba\x1c\x98\x1c\x19\x8f\x2f\xaa\xa1\x95\xb5\x6f\x58\x98\x46\xe2\xb6\xef\x91\x49\x0e\x09\xe7\x77\x6a\xa5\x78\xda\x68\x07\x67\x77\xdc\x10\x1a\xdb\x12\xb4\x43\xaf\x1b\x32\x8e\x5a\xec\x4c\x4b\x0c\xbf\x21\xdc\x0e\xaa\xd9\x10\x6e\xe4\xd5\x11\x45\x67\x77\xa6\x15\xda\x64\xfc\xdb\xe2\x6e\xfe\xb0\x9c\xa3\xd3\x3d\xe1\x70\xc7\xd6\x7a\xb4\x9a\xa9\xf1\x96\xf7\xb0\x1d\xfc\x3f\x66\x9e\x89\xa4\xb8\x98\xc6\x28\x44\x08\x68\xa9\xd3\x86\x50\x6e\x95\x36\x25\x62\x14\xd3\x29\xee\x52\x9e\x35\x19\x62\xe5\xa9\xc5\x6a\x8f\x73\x32\xbe\x39\x5d\x9d\x4b\xdc\x3f\xe2\xe1\xf1\x09\xf3\xfb\xc5\x93\x14\x83\x6a\xde\xd4\x9a\x90\x34\x84\xd0\xdb\xc1\xb2\x47\x25\x8a\xd2\xba\x52\x14\xe5\x6a\xef\x29\xfd\x84\x00\x4f\xdb\xa1\x57\x9e\x50\x8e\x2c\x97\x2d\x33\x34\xb0\x36\xbe\x43\xf9\xf1\x77\x09\xf9\xe3\xa0\x18\xa3\xa8\x73\xcc\xb3\x95\x72\x84\xcf\x33\xe4\xef\x11\x4f\xb3\xef\x8a\xe1\x9a\x0d\x6d\x95\xc3\x0c\xcf\x2f\x64\xbc\x5c\x18\x4f\xdc\xa9\x86\x42\x96\x66\x65\xd6\x84\xb3\xd7\x09\xce\x8c\xda\x66\x19\xf9\xa0\xb6\xe4\x92\x7e\x51\x84\x70\x79\xd0\x8f\x51\xa6\xc3\x29\x8a\x0b\xb1\x3c\xcc\xc4\x38\xc9\x5a\x64\x5a\x5c\xc6\x28\xa2\x10\xdd\xce\x34\xb9\x73\x55\x23\x88\x22\x05\xe9\xb5\x21\x87\xe7\x97\xe7\x97\x54\x5a\x14\x9d\x65\xbc\x4e\x0e\xf9\x92\xef\x18\xe5\x98\x37\x88\xa2\x58\x4d\x40\xcc\x09\xfb\xae\xd8\x6d\x54\xbf\xcc\x60\x35\x72\x6a\x51\x14\xba\xcb\x8c\x0f\x33\x18\xdd\xe7\x99\xa2\x53\xba\xaf\x88\x39\xc1\xa9\xc2\xe8\x3b\x83\x1a\x06\x32\x6d\x95\x8f\x13\xac\x6a\x91\x50\xeb\xe4\xd2\xb7\x76\xe7\xe5\x4f\xd6\x9e\xaa\xbc\x0f\xf9\xd5\x6a\x73\x24\x8e\x71\xab\xf2\x97\x29\xeb\xba\x3e\x75\x3b\xba\x24\x7b\xcb\xb9\xe4\xa8\x45\xcc\xa3\xd6\xd2\xb3\x36\xeb\xc4\x91\xf3\xc4\xa9\xea\x3a\x73\xe6\x7f\xb4\xaf\xae\xb3\xd2\x7f\x5b\x1f\x4b\x8d\x4b\x3f\x3c\x66\x8c\xe2\x6f\x00\x00\x00\xff\xff\xe4\x6e\x0c\x4d\x4b\x03\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 843, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x5f\x8f\xdb\x38\x0e\x7f\x8e\x3f\x05\x1b\xa0\x85\x3d\xc8\x3a\xbd\xc5\xe1\x80\x73\x91\x87\x45\x77\x16\x98\xdb\xeb\x74\xb1\x9d\xbb\x97\xa2\xe8\x3a\x36\x95\xa8\xb5\x65\x57\x52\xd2\xc9\x0e\xe6\xbb\x1f\x48\x49\xfe\x13\x67\xd2\x6d\x7b\x33\x2f\x63\x53\x24\x45\xfe\xf8\x13\x69\x65\xb9\x84\x97\x4d\x7b\xd0\x72\xb3\xb5\xf0\xe3\xf3\xbf\xfd\xf3\x87\x56\xa3\x41\x65\xe1\x97\xbc\xc0\x75\xd3\x7c\x84\x2b\x55\xa4\xf0\x53\x55\x01\x2b\x19\xa0\x75\xbd\xc7\x32\x8d\x96\x4b\xb8\xd9\x4a\x03\xa6\xd9\xe9\x02\xa1\x68\x4a\x04\x69\xa0\x92\x05\x2a\x83\x25\xec\x54\x89\x1a\xec\x16\xe1\xa7\x36\x2f\xb6\x08\x3f\xa6\xcf\xc3\x2a\x88\x66\xa7\x4a\x72\x21\x15\xab\xfc\xfb\xea\xe5\xe5\xf5\x9b\x4b\x10\xb2\xc2\x20\xd3\x4d\x63\xa1\x94\x1a\x0b\xdb\xe8\x03\x34\x02\xec\x60\x3f\xab\x11\xd3\x28\x6a\xf3\xe2\x63\xbe\x41\xa8\x9a\xbc\x8c\x22\x59\xb7\x8d\xb6\x10\x47\xb3\x39\xaa\xa2\x29\xa5\xda\x2c\x3f\x98\x46\xcd\xa3\xd9\x5c\xd4\x96\xfe\x69\x14\x15\x16\x76\x1e\x45\xb3\xf9\x46\xda\xed\x6e\x9d\x16\x4d\xbd\x14\x3e\x61\xa9\x8a\xdd\x3a\xb7\x8d\x5e\xa2\x62\xfd\x2f\xe9\x2c\x4d\xb1\xc5\x3a\x5f\x62\xb9\xc1\xaf\xd1\x17\x12\xab\x72\x1e\x25\x11\xa1\xf0\x86\x65\xa0\xd1\xe3\x6f\x20\x57\x80\xca\xa6\x7e\xc1\x6e\x73\x0b\x9f\x73\xc3\x69\x62\x09\x42\x37\x35\xe4\x50\x34\x75\x5b\x49\xc2\xda\xa0\x06\x0f\x45\x1a\xd9\x43\x8b\xc1\xa5\xb1\x7a\x57\x58\xb8\x8b\x66\xd7\x79\x8d\x00\x40\x12\xa9\x36\xf4\x04\x7f\x10\x36\xd9\x5c\xe5\x35\x2e\x9a\x5a\x5a\xac\x5b\x7b\x98\xff\x11\xcd\x5e\x36\x4a\xc8\x0d\x70\x08\xfe\xd9\xeb\x16\xfc\x36\xd6\xbe\x2c\x37\x68\x00\xe0\xed\xbb\x0b\x7a\x1c\x78\x26\x50\xcc\x58\xf9\x17\xca\xdb\xb0\x32\x3f\xf6\xca\x8c\xc8\x91\xf6\x95\x2a\xf1\x16\x0d\x69\xf3\x63\xaf\x2d\xdd\xca\x48\xfd\x9e\xc1\xfc\xad\x31\xd2\xca\x46\x41\x89\xa6\xd0\x72\x8d\x06\x72\x60\xe7\xd0\x86\x25\x4f\x31\x57\x0b\x8f\x58\x67\xd7\x63\x16\xf6\x04\x00\xa9\x2c\xc0\x72\xe9\x1d\xf1\xee\xc1\x8b\x13\x55\xd2\xd8\x34\x9a\xbd\x92\xb7\x58\x5e\x29\x32\x59\x37\x4d\x05\xcc\xf1\x52\x16\xb9\x45\x03\x52\x0c\x0c\xa8\x9e\x35\x69\xff\x20\x95\x33\x94\xea\xca\xfb\x75\x7b\xd5\x24\x1a\xef\xe5\x44\x6e\x2f\x97\xae\x43\x71\x4a\x1d\x27\xff\x06\xe6\x38\xc3\x29\x71\xdc\xdf\x80\x3e\xe7\x39\x74\xa5\x44\xd3\xab\x5d\x70\xce\xe9\xcd\xa1\x45\x5e\xf0\x66\xb4\xe1\xd8\xec\x26\x1f\x38\x7f\x68\x37\x9b\x1f\x51\xf0\x8d\xfc\x73\x10\xe3\x85\x54\xf6\x1f\x7f\x9f\x58\x19\xf9\xe7\xd1\x66\x97\x6a\x57\x9b\x4e\xed\xed\xbb\xf1\x76\x81\xc4\xa4\x34\xb6\xfb\x8f\x92\x9f\x76\xdd\x86\x5c\x67\x98\x6c\xb7\x63\xa5\xb1\xe1\xb5\xac\xaa\x7c\x5d\xe1\x59\x43\xe5\x95\xc6\xa6\xaf\x5b\x22\x67\x5e\x9d\x35\x6d\xbc\xd2\xd8\xf4\x67\x14\xf9\xae\xb2\xe7\xc3\x2d\x9d\xd2\x51\xa2\x6d\x99\x5b\x0c\xf6\x0f\x25\xca\x4a\xef\x4f\x3a\xb8\xaa\xeb\x9d\xed\x32\x7e\xc0\x81\x0c\x4a\x63\xdb\xff\xe6\x95\x2c\xa9\x6f\x72\x89\xf8\x50\x4c\x6d\xf7\x9d\xd2\x11\x23\x6c\xa3\xf3\x0d\xfe\x8a\x87\x33\x3c\x32\x4e\xe9\xfd\x47\x3c\x8c\xad\xbb\x5e\xe0\xf8\x34\x7e\x0d\xd6\xa1\x9b\x1c\x6d\x8c\x8a\xc4\xfb\xb3\x19\x9b\xa0\x74\xa2\x7f\x71\x0f\x9d\x9e\x67\x16\x7f\xc3\x71\x66\xbb\x13\xa7\xd9\x43\xf2\xe0\xf9\xa5\xa3\x7a\x42\xf1\xcc\x89\x3d\x52\x3c\x3e\xa3\xbf\xa3\x70\x9b\x8f\xf5\x34\x8a\xf7\xd3\xdd\x7f\x47\xe1\xb1\x77\x23\xa5\x57\x7e\xe0\x24\x7a\xa0\xcf\x9c\xbd\x2b\xb5\x47\x6d\xf0\x58\x55\x3a\xf1\xf1\xf6\x9f\x76\x52\x63\x79\xa4\xab\xbd\xf8\x44\xd5\x5c\xe7\x9e\x96\xcd\xc9\xbf\xa1\x6e\xce\xb0\x2f\x9c\xcf\xb4\x23\xd4\x99\x4c\xfd\x40\xee\xfa\xd9\x17\xa7\xf1\xb1\xe6\x89\x51\xec\xb2\xbc\xc6\xcf\x5c\x8f\x42\x23\x0f\xb4\x5c\x85\x8c\xc8\xb9\x4b\x8b\x9f\xdc\xec\x6d\x6d\xa3\xd3\x48\xec\x54\x11\x2c\x63\x2c\xe1\x82\x34\xd2\x9f\x3b\x8d\xc4\x17\xf9\x2e\x9a\x29\x84\x6c\x05\xcf\xe8\xf5\x2e\x9a\x11\xb5\x32\x47\x03\x2c\xd3\x9b\x7c\xb3\x20\xd9\xa1\xc5\xac\x93\x11\x1b\xa3\x19\xb3\xba\x13\xd2\x0b\x09\x1d\x62\x99\x13\xba\x17\x12\x7b\x1e\x64\x2c\xf6\x2f\x24\x0f\x35\xcf\x48\x1e\x5e\xdc\x82\xf0\xfe\x79\x41\x78\xff\xf7\xd1\x4c\x0a\xd0\x28\x28\x64\xb7\xf2\x82\x5f\x9f\xac\x40\xc9\x8a\xd2\x99\x29\x24\x31\xac\xba\xf4\x35\x8a\x84\x4d\x35\xda\x9d\x56\xa0\xb0\x47\xd6\x0d\xde\x29\xb4\xee\x73\xe1\x3c\xb6\x6c\x1b\x8b\x32\x0c\xda\x21\xba\xb1\xfb\xd2\x5a\x00\x6a\x4d\xef\x77\xd1\xcc\x70\xd0\xcf\x58\x7e\x37\xc2\x8f\xff\x44\x0f\x22\x4d\xeb\xf1\x0a\x49\x16\xa3\xe2\x84\x15\x5f\x21\x9e\xaa\xd9\x70\x81\x25\xe3\x92\x84\xa5\xbe\x2e\x61\x36\x66\x7d\x0c\x61\x10\x46\xb3\x6e\xfc\xf5\xab\x41\xc2\x51\x86\x09\x92\x75\x51\x76\x33\x25\x9a\x0d\x86\x41\xe6\x97\x7b\x09\xad\xf7\x93\x86\xd7\x2b\x54\xb1\x28\xd3\x5e\x9a\x90\x92\x9f\x82\x59\x1f\x7b\x98\x8b\xae\xe0\x9c\xdf\x70\x5e\x66\x9c\xdf\x68\x82\xf6\x9a\xdd\x98\xe8\x22\xee\x24\x1d\xb7\x8c\x60\xb0\x61\xd5\x13\x2a\xd0\x46\x56\x0b\x10\xb5\x4d\x2f\xa9\xa4\x22\x9e\xd7\xd2\x18\x3a\xc0\xdc\x3a\x24\x19\x89\x46\x7b\xe6\x3c\xfd\x34\x5f\x90\x2f\x2a\x69\xd2\xf9\xa6\x2f\xa6\x6c\x05\xfc\xa9\x44\xd9\xd2\x27\x54\xf2\xc2\xc9\x9f\xac\xe0\x39\x6f\x67\x04\xcb\x61\x05\xcf\x68\x61\xc8\x5c\x23\x16\x14\x86\xa7\xef\xab\x5c\x9b\x6d\x5e\xf9\x5b\x07\x5f\xbe\x90\x3f\x58\x07\xb7\x18\xa9\x2c\x6a\xba\x14\xd1\x53\x03\x39\xfc\xeb\xcd\xeb\x6b\x32\xe6\xee\x58\xe4\x0a\xd6\x44\x6e\x32\x2d\x9d\x0a\x39\xf0\xc6\xcd\xfa\x03\x16\xd6\xff\xf3\xbc\x1f\x6d\x1a\x9b\xb0\x37\x35\x5d\xbf\x53\x02\xf1\x1a\xde\xbe\x5b\x1f\x2c\x32\xfd\x87\x47\x80\x4f\x80\xb3\xa5\x54\xdd\x3d\x27\x0b\x77\x02\xf7\x1a\x27\xc3\xee\x42\x1f\xf2\x74\x1d\x8d\xfd\x25\x92\xdb\xcf\x6b\xe1\x77\x4e\x12\x46\x98\x4d\x1c\xc6\xb4\x61\xb6\x02\x93\xd2\x41\x76\xad\x36\xe8\xbe\xe0\xc5\x27\xa7\x0b\x8b\x5a\xb3\x0b\xd7\xb0\x3b\x37\xb9\x40\xee\xec\xc1\x47\xb7\xc7\x03\x6e\x86\xfc\xf0\xe0\x3c\xfd\x94\xc1\xd3\x3d\xd1\xc1\x1d\x70\x32\x77\x94\x20\xba\xbc\x5f\x00\x73\x42\xe7\x6a\x83\xdc\x6b\x8c\x63\x41\xea\x26\xca\x0a\xf2\xb6\x45\x55\xc6\x5e\xb0\xe8\xbb\xfa\xa0\xe1\xc4\x49\xe2\x59\xe6\xaf\x68\xc3\x04\xfc\x85\xee\x31\x53\x90\xe5\x6d\x9f\x84\x8f\x81\x1d\xfb\x05\x59\xde\x8e\xa2\xe5\x04\xc3\x45\x73\x90\xe2\x55\x08\xff\x19\x3f\x91\x07\x37\x59\x33\x60\x1f\x0e\x02\x92\xba\xd2\x66\x2c\x75\xcf\x2c\x0e\xcd\x8e\xc4\x7d\x9b\xbb\x1f\xf5\x7f\x9a\xb7\xa9\xe7\x71\x6c\x12\x7f\x9a\x7a\xbe\xf0\xa3\xf1\x07\xd9\x36\x9e\x9d\x7e\x18\x0c\x99\xee\x8f\x44\x6c\xe0\xc2\x71\x3a\x81\x09\xeb\x8e\xcf\x06\x1f\x06\x82\x86\x6f\x95\xa3\x3a\xf1\x6d\xf4\x2f\x54\xe9\xab\x0b\x24\x17\x50\x0f\xea\xe3\xee\xb3\xe4\xd0\x7f\x73\x0c\x83\xf0\xc1\xd7\xb7\x54\xa3\x69\x08\x5f\x1f\x03\x05\xc1\x51\x7c\x58\x80\xe8\x83\x70\x5b\x3b\x9f\xd4\xd6\x7c\x08\xfd\x58\x1d\xb3\x9b\xd4\x4e\x44\xf3\x0d\xe1\x70\x3c\xd4\x64\xbb\x0b\xc6\x0a\x9e\x85\x67\xe7\x94\xb9\xe7\x67\xce\x07\xa6\x55\xf8\x89\x81\x85\x56\x3b\x56\xcd\x06\xbf\x1f\x64\x20\x17\xbd\x73\xcf\xc8\x21\xb3\x3d\x47\xc1\x08\x8f\x09\xd5\xe6\x41\xf8\x1f\x87\x04\xa7\xe1\xff\x6b\xe8\xff\xdf\xa8\xf0\x10\xf2\x01\x46\xd6\xf9\x12\x80\x83\x8f\xb9\x6e\x1c\xf6\xf0\xc1\x67\x9d\xb7\x66\xf8\x6b\x93\x97\xe7\xaa\x74\xec\x0f\x82\x1a\xed\xb6\x29\xe1\xb3\xb4\x5b\xd0\x58\x34\x7b\xd4\x74\xe2\x51\x99\x9d\x46\x50\x0d\xb4\xb9\x92\x85\x01\xa9\xa0\x76\x0d\x43\xaa\x8d\x3f\xf6\x83\x72\x89\xb2\x1f\xb4\x77\xe0\x85\x09\xbc\x7d\xd7\xff\x28\x74\x9f\x40\x2c\xc2\x97\x7f\x27\x3e\x1e\x90\x25\x0a\xd4\x40\xee\xe3\xc4\xf5\x4f\x01\x7b\xae\x9a\x0b\x2e\x4e\x5e\xc0\x7e\x54\x04\xb2\x5f\x8d\x6a\xf0\xf4\x26\x64\xe7\x82\xf7\xa5\x10\xe5\x02\xf6\x7c\x00\x44\xc0\x96\xb0\x73\x5c\xa4\x8e\x1c\xca\x59\xa6\x21\x81\xc5\x11\xba\x6e\x22\x4d\xc0\x75\xe2\xef\x85\x72\x38\x66\x27\x1f\x14\x6e\x2e\x3a\xe0\x48\xf1\x31\x70\x1b\x65\x33\x82\xce\xc1\x86\x7e\x1e\x9f\x44\x6d\x68\x3c\x05\x2e\x4c\xba\x09\x74\x61\xe1\x7b\xc1\x1b\x8f\xf8\x09\x7c\xb2\xfb\x49\xb7\xbb\x1e\x3f\x22\x82\x21\xa9\x13\x18\xca\x6e\xe4\x9f\x43\x31\x64\x33\xc1\x91\xfb\xed\x14\x45\x27\xfe\x5e\x0c\x87\xe3\x77\x82\xa0\x9b\x99\x0e\xbf\x57\xfd\xe4\x7e\x14\xfc\x5c\x3a\x27\xd0\x73\x41\x9c\xc7\xce\x65\xd1\x23\xc7\xe9\x75\x1f\xd1\x16\x86\x9f\xd1\xc9\xe8\x8d\xa2\xa2\x41\x61\xd3\x5f\xa5\x2a\xe3\x84\xae\x40\x61\xfd\x37\xcb\xdf\x2c\x33\x0b\x2b\xb0\xe9\x65\x85\x75\x3c\xea\xc2\x36\xba\x8f\xfe\x17\x00\x00\xff\xff\x9f\xab\xe4\xff\xda\x1a\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 6874, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
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
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
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
