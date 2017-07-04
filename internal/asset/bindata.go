// Code generated by go-bindata.
// sources:
// assets/text/adjectives.txt
// assets/text/adverbs.txt
// DO NOT EDIT!

package assets

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

var _assetsTextAdjectivesTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x56\x41\xba\x23\xab\x0a\x9e\xb3\x8a\xb7\x81\xb7\x28\x4a\x29\xa5\xa3\x62\x23\x26\xc7\xb3\xfa\xfb\x41\xa5\x27\x40\x28\x40\xc0\x1f\x0c\x5e\xba\xa7\x01\x26\xce\x9c\x00\xb3\x28\x5e\x8d\x00\xf3\x9b\x86\x6d\x95\xbd\x00\x4b\x51\x5a\x8b\xdf\x04\x58\xd8\xd0\x28\x03\x36\x52\x03\x6c\x22\x37\x60\xe7\xc7\xa7\xef\xe5\x9f\xc6\x90\x13\xdc\xd6\x01\x1c\x3f\x1c\x41\xe6\xc4\xd6\x78\x14\x97\xc8\xf8\xf7\x2b\x2a\x55\x1a\x4f\x6c\x55\x29\x38\x0c\x70\x55\xec\x1e\x60\x99\x0c\x5e\x35\x2c\xcd\x14\x93\x85\xdd\x9b\x14\x0b\xc1\x85\x66\x07\x2e\xa2\xdb\xe9\x87\x5b\x26\xa5\x0c\x17\x9b\x3b\x5c\x6c\x46\x0a\x57\xc3\x91\xe1\x6a\xfb\x09\x73\x49\x98\x28\xbe\x09\x2e\xe5\x52\x0d\x2e\x15\xcc\x70\xed\xf6\x3a\x70\x6d\x6d\x07\x52\x45\xed\x6e\x9d\x2a\xd1\xeb\x04\xd3\x7b\x37\x48\x75\x5f\xd7\x81\xd4\x08\x47\x50\x85\xd4\x64\x67\x57\x6d\x6a\xb4\x96\x0b\x7d\x1d\x48\xd2\xe4\x71\x91\x26\x6b\xa1\x0b\xfd\xc2\xc8\x3f\x49\xbf\x45\x2d\x5a\x96\x64\x64\xea\x83\xf2\x23\xad\x44\x23\xc7\xc9\x32\xee\x68\x66\x92\x61\xd4\x67\xfb\xe7\x3a\xde\x3c\xd2\xd7\xe2\x2d\x6d\x5b\xd8\xc8\x24\xfd\x67\xa1\xc3\x8f\x5f\xe6\x85\xc8\xf6\x4e\x79\xfb\x93\x62\xa4\xae\x44\xd3\x99\xc8\xcb\x3d\x75\x53\x83\xb4\xfb\x45\xba\xa4\x13\xa4\xad\x6f\xd7\x9f\xc1\x09\x1b\x64\x1c\x85\x02\x03\x19\x9f\x06\x66\x4a\xe8\x97\x9b\x29\x11\x9b\x57\x98\x89\x26\x64\xba\x29\x70\x91\xe9\x66\xbf\xc3\x4c\x8d\x53\x5c\xbc\x4b\xa5\x7e\x4d\xa7\x77\x3e\x87\x40\x6b\x85\xb4\x26\xa7\xe8\x45\x26\x23\xef\xbb\x6b\xb9\xe1\xe4\xfc\x44\xe4\xce\x63\x47\x71\x99\x57\xd9\xeb\x51\x2e\xe3\x91\x2c\x04\xc5\xed\x17\x19\xe2\x37\x2a\x2f\xdb\x7a\x85\xf4\xfb\x7b\x20\x2b\x5e\x4e\x9e\xe0\xbb\x35\x20\x2c\xa4\x40\x69\x19\x1a\x27\xa0\x16\x67\x51\xa3\x40\x20\x75\x4c\xfc\x68\xfa\x85\xaa\x18\x51\x69\xa4\x8a\x23\xd0\x45\xe3\xe9\xed\x23\x93\x16\x8a\x28\x43\xb4\x7b\xcd\x34\xac\xee\xc5\xb8\x1e\xed\x3b\x1a\x41\x3f\xb8\xe2\x9e\x3c\xd4\x4f\xe2\x87\x57\x6e\xf8\x4f\x67\xdf\x49\xa0\x9f\x7d\x91\x7a\x22\x37\x8e\x74\x9c\xda\x13\xeb\x66\xd2\x44\x70\x73\xb3\x7a\xe0\x6e\x68\x70\x37\x99\xd3\xe5\xed\x80\xf7\x84\x6e\x91\xc6\xab\xc2\xed\x21\xdc\x49\x29\x7e\x31\x8d\xdc\x8e\x0b\xa5\x1a\x79\x2b\x6e\x95\x88\xa3\xdb\xbb\x18\x95\xdd\x7b\x8c\x03\xf7\xf6\xb6\x15\x74\x6c\x17\x1a\xd6\x08\x4a\xc5\x00\x55\xe1\xec\x4a\x2e\x4f\xf0\xd2\xb0\x4b\x40\xa4\x34\xc2\x18\x9b\xd2\x44\xa3\xe2\x22\xfa\xa0\xaf\x28\x26\x72\x04\x14\x25\x5c\x07\x8a\x32\xbd\xc3\x54\xd9\x87\xb8\xa8\x18\xad\xbf\x9b\xa0\x68\x8c\x58\xd1\xdd\xe7\x81\x8a\x23\x07\x2e\x2b\xce\xf8\xa9\x74\x7d\x6f\xb1\x12\x46\x0f\x2a\xb5\xe9\x91\x9d\xc7\x04\x56\x2e\x15\xaa\xb4\x26\x1f\xa8\xd2\xa9\x1d\xa8\xa2\xca\x37\x27\xa8\xbb\x10\xd4\x3d\x8a\x1e\xa8\x5b\x0d\x38\x1d\xe0\x4c\xd8\x80\x7b\xa7\xb1\x08\xb8\x07\x34\x59\x46\x80\x92\x87\x29\x97\x4d\x19\xd8\x6f\x09\x58\x95\x9f\xd1\x65\x4b\xf5\xc0\x1f\xc2\xe6\x15\xfe\x89\x75\x73\xe0\x8f\xb4\xe6\xf4\x84\x72\x7b\xfc\x3f\x51\xca\x8b\x47\x86\x86\xe9\x85\x19\x79\xc5\x74\x35\xd4\x42\xd0\xf0\xf7\x40\x23\xab\xae\x61\xf3\x56\x37\x7e\x7b\xda\x8d\xdf\x9c\xa1\xc9\x88\x1f\x22\x8b\xa0\xc9\xf3\x65\xa7\x97\xd3\xcc\x29\x5a\xdf\x31\x55\x81\x8e\x65\x78\x99\x34\x0c\x3a\xf6\x2e\x56\xa1\xe3\x60\xf4\xc3\x3a\x3e\x1b\xbc\x53\xc3\x91\xaa\xb4\xe3\xa2\xe3\xae\xf3\x60\xb4\xad\x04\x31\x68\xce\x96\xe1\x8b\x46\x08\x07\xba\x48\x76\xaa\xc6\x37\xbb\xbd\x18\xcb\x88\x5e\xf7\xed\x58\xe8\x67\x39\xf2\x3c\x91\xe1\x4b\xfc\x03\x03\xdd\x71\xc4\x5c\x1e\x18\xa4\xef\xf8\x28\x3e\x40\xcd\x71\x3d\x64\x2c\x07\xbb\x67\x36\xb6\x29\xdb\xe3\xbe\x1d\x0d\x72\x51\x66\x2f\x42\x2e\x6f\x81\x7f\x90\x6b\xc8\xf3\x8e\x48\xce\x20\x2d\xff\xff\xf6\x85\x24\x8e\x04\xd9\xf6\x6f\xcb\x4d\x1c\xec\x9d\x99\xa4\x37\x25\x73\x3e\x1b\xfd\x50\x06\x7f\x72\x8c\x9c\xd9\x81\xd9\x90\x07\xcc\x46\xb8\x3c\x99\x29\xec\xe3\x3d\xa5\xcf\x08\xa2\xf4\x6c\xae\xa9\x9c\x5e\xed\xc0\x54\xd9\x19\xe6\x1e\x3e\x09\xce\x0f\xfc\xdd\xc8\xc3\xe0\xef\xe6\xdf\xdf\x28\x42\xe3\x31\x72\x74\xaf\xed\x4f\x8c\x52\x63\x7a\x87\x30\x77\x8b\xce\x2b\xad\x29\xcf\x80\x2b\x4f\x02\x95\x6b\x2f\x03\x15\x33\x1a\xce\xf6\xc8\xa0\xb2\x4b\x75\x3a\x32\x2c\x6c\x76\x60\xa1\xa6\x67\xf8\x57\xf2\x6c\x57\x42\x3d\x4e\x1d\x70\x94\x61\x25\xc5\xcf\x38\xb0\xa8\xdd\x3e\xf5\xab\x62\x29\xc7\xd9\x2b\x68\x0c\xc2\xaa\xa8\x13\x7c\x85\xbb\x4e\xd4\x60\xb1\x3f\x78\x8b\x1d\xae\xeb\xc5\x3e\xf4\xab\x71\x0f\x3a\xa7\x43\x79\x75\x6c\xcd\xa9\xba\xb6\x73\x3c\xde\xab\x4b\x44\xef\xe2\xe8\x5a\x7d\x17\x58\x8f\x46\x1a\x67\x58\xa2\x04\x4b\xb6\xc2\x9a\xa8\xaf\xc7\x65\xfa\x14\xac\xd9\xfc\x6d\xcb\xb0\xa6\x58\x80\x67\xfd\xdd\xe8\xd6\x86\xcd\x29\x61\x3e\xce\x68\x82\x57\xeb\xc9\x99\xc4\xd1\x26\xdb\xc0\x17\x54\xbc\xd7\x2e\x8c\xe2\x1e\x2a\x1e\xdd\xf6\x18\x71\xcc\xbe\x96\xf9\x4e\xc2\x06\x6b\xa7\x44\x6b\xf9\x5e\x70\x71\x37\xbf\xb8\xb5\x1d\x16\x9c\x1e\x83\xe9\x88\x55\x58\x1f\x1c\x7e\xd4\x87\xc8\xc0\x50\x9d\x38\x76\x8d\x68\x38\x1d\x99\xd4\xd9\x22\x30\x52\x65\x9f\x7c\xa3\xb0\xa8\x38\x5e\x7e\x82\x55\x4e\x2f\xb0\xea\x37\x67\x8f\x22\xc4\x28\xd2\x22\x67\xe3\xd8\x43\xa6\x51\x97\xc3\x9d\xc0\x54\xf6\xd5\x28\x83\x7d\x62\x77\xfc\x6f\xc6\x2b\xb0\x07\xbd\x69\xc0\x1e\xcb\x3d\xdb\x81\x3d\x17\x19\xec\xf9\x44\x7a\xe3\x32\x78\x07\xa4\xdf\x9c\xec\xbb\x68\xdf\xac\xb6\x1f\xe1\x8d\xe9\xab\xf2\xe5\xf1\x41\x3f\xf0\x43\x0e\x99\x4f\xe5\xfe\x4c\xdc\xa7\xca\x9c\xde\xb3\x0f\x27\xff\x1b\xf0\x89\x25\xfc\x91\xeb\x6a\xce\xbc\x66\xaf\xe3\xe3\x6b\x93\x32\x9c\xdd\xfb\x81\x5f\x1c\x07\x7e\xbf\xeb\xee\x97\x7d\x1f\xff\x17\x00\x00\xff\xff\x92\x9b\x72\x91\x3c\x0a\x00\x00")

func assetsTextAdjectivesTxtBytes() ([]byte, error) {
	return bindataRead(
		_assetsTextAdjectivesTxt,
		"assets/text/adjectives.txt",
	)
}

func assetsTextAdjectivesTxt() (*asset, error) {
	bytes, err := assetsTextAdjectivesTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/text/adjectives.txt", size: 2620, mode: os.FileMode(420), modTime: time.Unix(1499083390, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTextAdverbsTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x56\xdf\x9e\xbf\x20\x0a\xbd\xef\x5d\xf6\xa1\xf8\x2a\x15\x93\x82\x83\x5a\xd3\x3c\xfd\x7e\x04\x6b\x7e\x57\xbb\x37\x71\x4a\x54\xfe\x1c\x20\xf8\xb0\x68\x86\x94\xee\x05\x3e\x15\xb9\x65\xe2\x88\x71\xbc\x86\x40\x11\xb9\xf9\x5a\x68\xdd\x41\x3c\x91\x5b\x57\xe9\x75\xbc\xad\x0d\xf5\x02\x8d\x75\x81\x94\xa5\xb6\x05\xd2\x05\x77\x5d\x80\x79\xea\xf3\x0f\x4d\x5d\x55\xd9\x80\xdb\x80\xd7\x31\x36\xa5\x7b\xf9\x40\xdd\xd7\x3e\x14\x3f\x08\xbd\xd1\xc4\xd4\x1a\xea\x00\x09\xe1\x30\x49\x1c\x5d\xd6\x3a\x75\x04\x6a\x7b\x60\xb2\x45\x85\x13\x4d\x12\xae\x2e\xb7\xbd\x39\xa8\x76\x8a\x0a\x98\x62\xaf\x94\xee\x25\x40\xca\x26\x14\xfd\x9c\x81\x12\xd6\x6a\xb0\xb7\x69\x77\x40\x6d\x40\x3c\xd0\x8e\xa8\x53\x35\x21\xa8\xcb\x13\x1d\x48\x1d\x97\x07\x81\x1f\xe2\xcd\x50\x92\x47\x5b\x72\x16\x3b\x41\xb8\xd1\x8c\x4c\x10\x71\xa1\x8a\xa1\x19\xea\x0a\x1b\xce\x5b\x55\xdc\x10\xed\x68\x6a\x5d\xa7\x3d\x11\xc8\x9f\xdc\x0c\xb8\x21\x11\x03\xd2\xe9\x37\x47\xc4\x62\x62\x25\x8f\x77\xc4\x44\x1f\x54\x68\x38\x5f\xb6\x7d\x86\x2e\x52\xa2\x0d\x5d\x89\x46\x34\xa2\xf4\xcf\xb3\xa6\x08\x79\x5c\x81\x60\xf1\xc2\x84\x33\x7f\xc8\xa8\x1b\x36\x0a\xe6\x08\x0e\x06\xb9\x6d\xc8\x6d\xef\x95\xa0\xbe\x6b\xdf\xee\x2c\xd6\x82\x81\x1c\x9e\xc8\xf6\x98\x78\x12\x0b\x7f\xc0\xa2\x80\x3f\x81\x9a\x11\x10\x7f\x9a\x62\x1e\x26\xaf\x40\xea\xa2\x4d\xb6\xac\x30\x6f\x5c\x41\x97\x15\x6a\x5b\x56\x70\xa6\xae\xa8\x12\x66\xa8\x56\xd4\xd3\x9d\x5b\x09\x35\xd8\x51\x62\x44\x5a\x45\x12\xd5\xdd\x90\xb6\xce\x1e\x9a\x55\x81\x8f\x29\x1f\x0f\x56\x45\x5f\x42\xc6\xbf\x6f\x7f\x11\x7c\x9e\x4f\x7a\xb6\x11\x1c\xd3\x32\xf4\x7e\x1c\x56\x6c\xc9\xf8\xb7\x25\x9c\x8c\xdb\x14\xc2\x1f\x6c\x2f\x44\x68\x2e\x31\x8e\xc8\xef\x50\x8a\xcb\x6a\x39\xdf\x11\x52\xdb\x27\x3a\x5d\xa6\xe2\x9b\x07\x9a\x24\xde\x69\x1b\x1e\xee\xc2\x58\x9b\x81\xf2\xf0\x7b\x97\x3e\x22\xba\x77\xde\x74\xec\xa7\x9c\x31\x92\x47\x81\x98\x25\xb8\xc1\xc4\xdf\x9d\x2a\x35\x3a\x7d\xa1\x36\x98\xdf\x1b\x72\xc5\x07\x3d\x9f\x14\x6b\x73\x0a\x12\xcf\x02\x27\x55\x6a\xf0\x49\xf7\xf2\x05\xdb\x66\x79\xfd\x42\x48\x1e\x96\x2f\x39\x9d\x13\x5f\x72\xbb\xf5\x5f\x72\xcf\xa5\xfe\xa1\xe4\x97\x7d\xf5\xb8\xe5\xa7\x11\x7d\x75\x73\xe5\x40\xe3\xcf\x41\x31\xfa\x85\x07\x71\xdc\x11\xd4\xa9\x73\x78\xbb\x38\x18\x4e\xcf\xf2\xc1\x72\x4d\x45\x96\x2b\x61\xdc\xd0\x8c\x3a\x44\x8e\xe1\x7f\x82\x5f\x13\x58\x97\xba\xa4\xd9\x37\x12\x1d\x68\x22\x17\x13\x16\x84\x24\xab\xe5\x20\x09\x6f\x7e\x62\x12\x2f\xfd\x24\x3d\x9a\x38\x9f\xef\xb7\x99\x9c\x2d\xe9\x19\xbe\xf0\x2d\x8b\x8c\xc0\xc4\x9b\xfb\x9c\x31\xec\xc0\xef\x8a\x5a\x42\x32\x55\x54\x33\x31\x4b\x38\xfc\xc0\x2c\xdc\x76\x93\x8a\xe3\xe1\x11\x19\x6d\x77\x88\xbb\x36\x7c\x48\xc8\xd0\xba\x93\x90\xbd\x3d\xb0\x33\x8a\x51\xcf\xa9\x31\xfa\xd6\xc2\x64\x45\xc1\x42\x56\xe1\x2c\x6d\x91\x0f\x46\xf2\x94\xca\x87\xe5\xe9\xde\x12\x87\x13\xb2\xae\xc8\xd5\x03\x21\xeb\x4a\xb3\xa4\x65\x6d\xc8\x8b\x75\x38\x29\xe8\xa2\x51\xa6\xd7\x5f\x39\x51\x83\xf0\x6a\x03\x25\xdd\x4b\x01\x62\x77\xbe\x80\x36\x3f\xa4\x40\x9b\xf7\x16\xd4\xd5\x5b\x62\xd9\xef\x3a\x8f\x28\x09\x26\x47\x8a\x24\x32\xa2\x16\x11\x35\xf1\x32\xb4\xc8\x60\xe3\x3c\x4f\xae\xa7\x5b\x17\x95\x5c\x9a\x83\x62\xed\xba\x74\x7e\x26\xda\x77\x1f\xad\xd4\xc0\x6c\x74\xdf\x1d\x4d\xe9\xbb\x8f\x8c\x09\x5b\x1a\xbe\x3b\x85\x03\xd5\xa5\xbf\xa3\xef\x22\x3d\x7c\x17\xfd\xfe\x4e\x63\x15\x38\x4a\x36\x50\x28\x9a\xd4\x61\x9f\x22\x58\x41\x2b\xba\x1a\x42\xad\x5d\x3d\xbb\x8a\xe1\x98\xd5\xa9\xb8\xf5\x64\x79\x53\x4c\x3d\xcc\xa2\x53\x2c\x08\xce\x6f\xc5\xa2\x02\x61\x36\xc3\x51\x76\x13\x0d\xe6\xce\x09\xa2\xb4\xf9\xd5\xd2\xad\x0d\x68\x8f\xc3\x84\x0a\xab\x89\x00\xde\x10\x07\x18\x26\x55\x04\x0d\xbb\x9b\x52\x31\x7a\x2b\xa8\x88\xf9\xf9\x94\xa2\xe4\x21\x56\xaf\xa8\x8a\x05\xf4\xd1\x7a\x78\x57\x77\xb0\x58\xd4\x1d\xb4\x98\x44\x2c\x53\x7f\x57\x4a\x26\xef\xf1\xa4\xe4\xb9\xae\x69\x68\x18\x90\x6b\x88\x2c\x62\x24\xaf\xb2\xda\xba\x24\xcc\xec\xc0\xdc\xa9\x92\xb1\x51\x1e\x85\x2a\xc2\x4b\x2d\xb3\x47\xd6\xf6\x36\xc5\x51\x08\xb6\xa5\x29\x19\x8d\x6a\x0f\x01\x9f\x1f\x87\xda\x63\x34\x8e\xd6\x5e\x8a\x54\x8b\x67\xed\x5a\x94\xea\x74\xb5\xd7\x42\xcf\x08\xa9\x17\x5a\x9a\xeb\x45\x6e\xcf\x9d\x0b\xb4\xfd\x1d\x05\x0d\x39\x1a\x59\x9e\x86\xd8\x46\x01\x0f\xc2\xb4\x1d\xf8\xf0\x2b\xdb\x2e\x4f\x16\xda\x3e\xc0\x4c\x57\x9b\x8d\xa6\x49\x16\x55\xb9\x96\x26\xb2\xd8\xd0\xe3\xe8\xd7\x37\xa5\x9e\xcb\xee\x04\x68\xda\xfd\xf9\x4c\xc1\x7f\xc6\xd0\x5f\x3e\xff\x6f\x1e\xff\x77\x02\x2d\x61\x4b\x4f\x8d\xb2\x2f\x76\x86\xf1\x97\x66\x71\xea\x0c\x21\x48\x67\xef\xe8\x9d\x3f\x08\x3a\x21\xb6\x7d\x86\xa4\x33\xfe\x14\x0c\x6d\xee\xf8\x77\xc0\x76\xa6\x5c\x14\xeb\x6c\x21\x9d\xff\x1a\x55\x67\xc6\x91\x24\x67\x63\x2f\x1f\x84\xb6\xf4\x62\x2e\x2e\xbd\x54\x8a\xf8\x9f\x28\x17\x2f\xbd\x8c\xd1\xb2\x74\x9d\x43\xb5\xd7\x39\x36\x7b\x7d\xa6\x5b\xaf\x5e\xda\x7d\xfe\x46\x9e\x10\x3c\x82\x27\x6c\x1d\x4d\xda\x0f\xdd\x09\x89\x9e\x05\x6b\xa2\x27\xea\xc7\x76\x9e\xa8\xf7\x72\xbe\x2c\x38\x29\x34\xd1\xf7\x45\x26\x79\x4f\x3a\xe1\x55\x91\x34\xe2\x62\xd6\x5f\xa0\xa3\xf8\x2f\xff\x79\xbd\x70\x7e\xc5\x94\x96\xcb\xd8\x74\xed\xf6\xef\x77\x91\xfd\xb8\x5e\x94\x92\xbb\x70\x91\x71\xe8\x92\xe9\xd2\x25\x83\x5d\x0f\x56\x25\x8b\xe9\xa5\x62\xb9\xbc\xe1\x62\xcf\xea\xed\x5d\x7e\x88\xf7\xcb\xa8\x82\x08\xf7\x72\x13\xa6\x39\x24\x6f\x79\xa9\xf3\xfb\x0e\xe1\xdf\xb7\x7b\x0c\x34\x0c\xfd\x6f\x00\x00\x00\xff\xff\x16\x57\x5c\x0d\x14\x0c\x00\x00")

func assetsTextAdverbsTxtBytes() ([]byte, error) {
	return bindataRead(
		_assetsTextAdverbsTxt,
		"assets/text/adverbs.txt",
	)
}

func assetsTextAdverbsTxt() (*asset, error) {
	bytes, err := assetsTextAdverbsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/text/adverbs.txt", size: 3092, mode: os.FileMode(420), modTime: time.Unix(1499083405, 0)}
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
	"assets/text/adjectives.txt": assetsTextAdjectivesTxt,
	"assets/text/adverbs.txt": assetsTextAdverbsTxt,
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
		"text": &bintree{nil, map[string]*bintree{
			"adjectives.txt": &bintree{assetsTextAdjectivesTxt, map[string]*bintree{}},
			"adverbs.txt": &bintree{assetsTextAdverbsTxt, map[string]*bintree{}},
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

