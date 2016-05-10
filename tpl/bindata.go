// Code generated by go-bindata.
// sources:
// tpl/mongo_collection.gogo
// tpl/mongo_foreign_key.gogo
// tpl/mongo_mongo.gogo
// tpl/mongo_orm.gogo
// tpl/mongo_search.gogo
// tpl/mssql_config.gogo
// tpl/mssql_orm.gogo
// tpl/struct.gogo
// DO NOT EDIT!

package tpl

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

var _tplMongo_collectionGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x58\x61\x6f\x9b\x48\x13\xfe\x6c\xff\x8a\x29\xaa\x5e\xe1\x88\x90\xf6\x6b\x5e\xf9\xa4\x36\xb9\x5c\x2d\x5d\xd2\x5e\xdd\x4a\x27\x45\x91\x85\xcd\xe0\x6c\x83\x59\x1b\x70\xae\x11\xe5\xbf\xdf\xcc\x0e\x60\xc0\x38\x71\x54\xb7\xba\xd3\x7d\x88\x14\x76\x67\x67\x9f\xe7\x99\xd9\xd9\x59\x67\x99\x8f\x81\x8a\x10\xac\x85\x8e\xe6\x7a\x32\xd3\x61\x88\xb3\x54\xe9\xc8\xca\xf3\xfe\xc9\xc9\x59\xf5\x0d\x97\x5e\xe4\xcd\x11\x16\x98\xde\x6a\x3f\xe9\x67\xd9\x4b\x3d\xfd\x02\xa7\x43\x70\xc9\x32\x58\x47\x33\xb0\x35\x1c\x4d\xb2\xcc\xbd\xf2\x16\x98\xe7\x97\xf3\x78\x00\x17\x2a\xf2\xdf\x47\x68\xaf\xd6\x18\x3f\x80\x8a\x52\x8c\x03\x6f\x86\x59\xee\x40\xa2\xe3\xf4\x42\x61\xe8\x27\xe0\xba\x6e\x92\xc6\x2a\x9a\x0f\xc0\x8e\x31\x59\x87\x29\x1c\x89\xff\xc2\x97\x03\x18\xc7\xfc\xa7\xc9\x67\xd6\xef\x25\x98\x24\x84\xc9\x01\xc2\xcb\x10\xea\x9b\xba\xbf\x61\x4a\xb0\xed\x41\xbf\x47\xdc\x30\x86\xc2\xd8\x3d\x0b\x75\x82\x34\xdc\xef\xad\x78\x0d\x2d\x75\x19\x9e\x60\xe3\xe1\x49\x63\xcf\x31\xe1\xb3\x57\x75\x9c\x6c\xc3\x38\x86\xb0\x72\x99\xd4\xff\x04\x2b\x6d\x14\x63\xba\x8e\xa3\x7e\xde\x17\x21\xba\x1c\xc1\xd1\x62\xae\xdd\x3f\x78\xaf\x06\xf7\xeb\x9b\x92\x3a\xf3\xda\x8c\x0f\xe1\xcf\x71\xf5\x75\xa1\x42\x52\xce\xae\x43\xe9\xa9\x00\x42\x8c\xea\x63\xf0\x0b\xbc\x62\x2f\xbd\x95\x6b\xf6\xdc\x4c\x91\xc0\xb4\xa2\x84\xd9\xcb\x59\x03\xb1\xb1\x8e\x27\xca\xb7\x06\x15\xf4\x8e\x18\x1a\xcc\x5d\x11\x0c\xd5\x42\xa5\x0e\xe8\x20\x48\x30\xe5\xa9\x1d\xc4\x6c\x43\x7d\x5c\xc6\x6c\x23\xc4\x73\x62\xd9\x19\x33\x23\x02\xa3\xa8\x51\xff\x9d\xbf\x6d\x33\x3a\x60\xaa\x6c\x53\x40\xac\xe9\x73\xa7\x96\xb6\x8c\x0e\x44\x8f\xa7\x83\x5f\xc8\x07\x15\xe0\xd5\x63\xaa\x5d\xfd\x37\x64\x23\xab\xfd\x93\xf6\xff\xcf\xcd\x58\xde\xa1\x53\xf5\x2c\x3b\x86\xd8\x8b\xa8\x1e\xbd\x24\x5e\xf8\x95\x59\x9a\xf0\x8d\xf8\x0b\x13\x2a\x49\x6c\x42\xe8\x64\xde\x1d\x25\x9f\x23\x45\xf4\xa1\x55\xac\x6a\x31\xaf\x17\xac\xb7\x0f\x34\x25\x2b\x65\xd2\xae\xbe\x49\xdb\x0b\x72\xf0\xc1\x8b\xbd\x45\x9e\xef\x5b\xb1\x24\x0f\x08\xa5\x3f\x75\x2f\x99\x71\x8d\x41\xc0\x8c\x0d\x03\xd9\x41\x14\x20\xa0\xbd\x9e\x45\x4e\xcd\x74\xe1\xd6\x3a\x85\xd6\x88\x53\xf8\xc2\xc8\xe7\x15\x79\x2d\x33\x56\x92\x17\x4d\x8a\x6e\x3d\x31\x1d\x78\xed\xc0\x2b\x07\x22\x15\xee\x2c\x97\xfb\x95\xbc\x1d\x72\x5e\xae\x93\xf4\x50\x92\x1a\x1d\x65\xc6\x81\x09\x41\xd2\xee\xf3\x5c\xf3\x1c\xc9\x2a\xf9\x5f\x6c\x31\x1c\x32\x79\x93\x83\xe5\x48\x87\x64\xf8\x57\x63\x8c\x55\xd9\x33\x80\xe2\xd4\x6d\x05\x4d\x36\xa9\x8f\x34\xa3\x58\x2e\x1b\x7b\xf7\x26\x04\x79\x4d\x6f\x63\x16\x26\xb8\x47\x22\xef\x2d\x79\x51\x94\xa4\x18\xed\x28\x4c\x1d\x97\xf4\xf5\xcd\xbf\x25\xe9\x1b\x39\xdf\x28\xc0\xad\x1a\xff\xc4\x19\x78\x13\x86\x5d\x67\x60\x83\x65\xf3\xdf\xa3\x3d\xd1\xce\x7b\xe1\xd0\x21\x68\xeb\xf2\xf3\x34\x79\x94\x3f\xaf\x79\x76\x4f\x78\x28\xae\xc7\xaf\xe5\xef\x47\x92\x7c\xe7\x25\xdb\x04\x07\x30\xd5\x3a\x3c\x40\x03\x7b\xef\xc5\x54\xc0\xd2\xba\x6f\x81\xba\x7d\xd9\x97\x65\x3b\x95\xb2\xc7\x46\x2f\xa4\xe6\x7d\xfb\x66\x7c\x34\x2a\xa0\xb9\x6a\x03\x8f\xca\x4b\xad\xe8\x40\x1a\xaf\xf1\x31\xb2\x67\x7a\x1d\xa5\x5d\x74\xcb\xc8\xd1\xe0\x21\xfa\xf6\x7a\xf1\x6f\xd3\x14\x0c\x7b\x27\xe0\xdb\x87\xd1\x79\x62\x2b\xbf\xea\xb7\xbe\x37\xf9\x14\x2d\x23\x46\x0b\xef\x0e\xed\xeb\x9b\x69\x42\xc0\xdf\x4f\xbf\xd0\xdb\x69\xe4\x9b\x2b\x96\x3b\x1f\xe5\x0f\x08\x60\xa0\x63\x98\x38\xa0\xd8\x5c\x8a\x21\xc1\x60\xf9\x29\x3c\x66\xdd\x28\x29\x57\xbe\xc3\xaf\xb6\x32\xee\x7b\x66\x83\x21\x78\xcb\x25\x95\x17\xf2\x94\x38\xd0\xd8\x44\x4c\xf9\x62\xca\xeb\xa1\x6b\x28\x5c\x9e\x3c\x53\x8e\x2d\x6e\xfc\x4f\xa5\x34\x5b\x54\x84\xe9\x7f\xc5\x05\xb8\x2e\x84\xe9\xc2\x2a\x25\x17\x3b\x95\x64\x21\x7f\xea\x0b\x8e\xa4\x7a\xd1\xa5\x95\x2f\x62\xc9\xa9\xe5\x86\xf9\xd7\x38\xbe\xd2\xe9\x05\x25\x87\xdf\x78\x00\x15\x26\x65\x16\x8d\x7c\x7b\x5b\x4d\x8a\xd6\x13\x4d\xcf\xb6\x22\x1f\x71\xa1\xef\xb1\xb3\xbc\x91\x34\x2a\x0a\xb4\x34\xf2\x67\xb7\x1c\xf9\x11\x7d\x1f\x58\x99\x22\xee\x4c\xac\x85\xe5\xd1\x48\x8a\xed\x56\x2c\xff\xb1\x41\x13\xbc\xbb\xc2\xd6\xa7\xee\x80\x76\xa3\xc6\x0d\xfd\x31\x7a\xf1\xec\x96\xdb\x84\x99\x8e\xd1\x3d\xc7\x10\x53\xb4\xad\x22\x3b\x65\xd6\xbc\x1a\xa8\xd1\x70\xa0\x39\xfe\xe9\x61\x89\x66\x58\xf9\x65\x93\x9c\x65\x45\xd3\xf1\x74\x2e\x94\x8a\x80\x5d\x68\x01\xcd\x17\x1d\xab\x28\xc9\x50\xfd\xca\x52\xf4\xb8\x26\x84\x74\x32\x0b\x0f\x04\xca\x3d\x9f\x16\xcd\x90\x63\xa8\xe1\x0a\xdc\x4f\xde\x34\x44\xb0\x2c\x1a\x24\x83\x0f\xde\xec\xce\x9b\x93\x85\x5b\xa1\xa0\x61\xe9\x14\x79\xde\x58\xcb\x18\x13\x18\xc8\x33\xea\xe4\x48\xe0\x57\x6b\x4c\x8d\x88\xfc\x4b\xed\xab\xa0\xf3\xf9\xaa\x97\xb0\x49\x60\x73\xdc\xcd\x59\xff\x31\xe7\x7c\x22\xfe\xb6\xcb\xfd\x9b\xe5\x32\x7c\xb0\xf5\xd2\x81\xe2\x6e\xdb\x84\xa3\x45\xe8\xf3\xd2\xf7\x52\x2c\xef\xfe\xb5\xf9\x6a\x1e\x4b\x83\xf6\x00\x60\x6b\x17\x70\xd7\xa6\x9b\x5f\x10\xc8\xb0\xca\x9b\x16\xce\xea\xb4\x76\x43\xb5\xbb\x8a\xc7\x81\x4e\x67\x55\x8b\x1a\x1c\xb6\x00\x6d\x68\x54\x2b\x76\x92\x89\xfc\x4d\xcc\x76\x11\x6a\xa6\x4e\xc1\xa5\x7d\x7b\x75\x64\xa5\x53\x4b\x43\x2e\x1f\xb2\x65\x75\xa1\x51\x1b\x4b\x37\x9a\x6c\x69\xde\x0b\x1f\x8d\x4b\x7a\xca\x9d\x9a\x7e\x86\x86\xf2\xc1\x36\xf0\x31\xa6\x7b\xa5\x4a\x1b\x60\x33\xe0\xa6\x28\x6d\xa1\xe8\xde\xee\x3b\x22\xde\x0d\xa2\xe6\x70\x37\x8e\xa3\x13\xf3\x58\x91\x5a\xf6\x77\x00\x00\x00\xff\xff\xb9\xbb\x35\x46\x06\x16\x00\x00")

func tplMongo_collectionGogoBytes() ([]byte, error) {
	return bindataRead(
		_tplMongo_collectionGogo,
		"tpl/mongo_collection.gogo",
	)
}

func tplMongo_collectionGogo() (*asset, error) {
	bytes, err := tplMongo_collectionGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/mongo_collection.gogo", size: 5638, mode: os.FileMode(420), modTime: time.Unix(1462872133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplMongo_foreign_keyGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x93\xc1\x4f\xc2\x30\x14\xc6\xcf\xec\xaf\x78\x59\x38\x0c\xb3\x94\x3b\x09\x07\x95\x20\x1c\xf0\xa2\x37\x63\x08\xb0\xd7\x59\x1d\x2d\xb4\x5d\xe2\xd2\xec\x7f\xb7\x5d\x27\xae\x0c\x12\x0f\x26\x1e\x08\x6b\xdf\xd7\xf7\xfd\xde\xb7\xce\x98\x0c\x29\xe3\x08\xf1\x5e\xf0\x5c\xac\xa9\x90\xc8\x72\xbe\xfe\xc0\x2a\xae\xeb\xc8\x98\xa1\xd8\xbe\xc3\x64\x0a\xc4\xae\xc6\x63\x5f\x7e\x03\x5b\x56\x91\xad\xca\x0d\xcf\x11\x86\x94\x61\x91\x39\x95\x53\x93\xb9\x5b\x29\xab\x1f\x18\xc3\x68\x5b\x25\x8b\x8d\x9a\xfb\xe6\xb6\x42\x4b\xbe\x83\xe4\x00\x37\xde\x80\x3c\x6e\xf6\x58\xd7\x23\x78\x40\x6d\x77\xfc\x81\x93\x3a\x19\x41\x22\x51\x95\x85\x6e\xf4\x41\xf5\xb9\x3a\xd8\x83\x29\xa0\x94\xee\x27\xe4\x08\x4c\x34\x90\xa8\x4b\xc9\xe1\xb2\x78\x95\x4b\x8b\xc8\xb3\xbb\x6a\x39\x4b\x0e\xe4\x24\x6a\x19\xa2\x3a\xba\x8a\xf7\x74\x11\xcf\x25\x74\x05\xac\xa1\xe9\x79\xc0\x14\x5c\xd7\xe5\x8c\x2c\xf0\x33\x71\x8e\x4d\x50\x78\x84\xb0\x07\x5c\x6a\xd9\xc2\x89\xbe\x65\x1b\xe0\x6d\x51\x04\xd8\x9d\xf8\x5e\x5e\xc3\x89\xce\x73\x3b\x96\x28\x2b\xf7\x1e\xb7\x4a\x70\xb2\x32\x71\xdf\x62\x39\x8b\x27\x20\x4e\xec\x75\x37\xec\x9f\xc6\xdf\x19\x5b\x96\xa4\x69\xda\x89\xf5\x0a\xf9\xbd\x28\xb9\xee\x81\xef\xdc\x2e\x30\xae\xff\x9e\xaf\x31\xfc\x2d\x9d\x1b\x26\x84\x2b\xd8\x9e\xe9\x14\x04\xa5\x0a\x1b\xc2\x14\x94\x90\xda\x5f\x7e\x20\x84\x28\x2d\x19\xcf\xff\x31\x7c\x3f\x5b\x0a\x01\x69\x97\xd2\x42\xfa\xcb\x87\x3c\xf3\x9f\xab\x7f\x38\xff\xff\x0a\x00\x00\xff\xff\x5c\x5e\xae\xae\x25\x04\x00\x00")

func tplMongo_foreign_keyGogoBytes() ([]byte, error) {
	return bindataRead(
		_tplMongo_foreign_keyGogo,
		"tpl/mongo_foreign_key.gogo",
	)
}

func tplMongo_foreign_keyGogo() (*asset, error) {
	bytes, err := tplMongo_foreign_keyGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/mongo_foreign_key.gogo", size: 1061, mode: os.FileMode(420), modTime: time.Unix(1458789996, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplMongo_mongoGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x54\x3d\x6f\xdb\x30\x10\x9d\xad\x5f\x71\x35\x32\x48\x85\xc1\x74\x0e\xa0\x21\x71\x9a\xc6\x43\x12\xa0\x41\xa6\xa2\x30\x28\xf1\xe4\x28\x95\x49\x97\x94\xf3\x51\x41\xff\xbd\x77\xa4\xac\xc6\xb2\x0d\x67\xe8\x20\x83\xe6\xbd\x77\xef\x3e\x9e\xd4\x34\x0a\x8b\x52\x23\x8c\x97\x46\x2f\xcc\xdc\xff\x8e\xdb\x36\x8a\x4e\x4f\xfd\x19\x96\x58\x3f\x1a\xe5\xa2\x67\x69\x21\x8e\x46\xa5\x76\x68\xeb\xe9\xc5\xbc\x69\xc4\xad\x5c\x62\xdb\xc2\x8f\x9f\xc5\x5a\xe7\xb1\xc9\x9e\xe0\xeb\x9f\x3b\xbb\xbc\xcb\x9e\x92\x68\xb4\x5e\x29\x59\xe3\x71\x60\x12\x45\x7c\x0b\x3d\xec\x5c\xa9\x59\x10\x91\x55\x95\xc9\xfc\x57\x9c\x67\xb0\x4b\x4c\xa0\xd9\x5b\x4d\x0a\x72\xb5\x42\xad\xe2\xdd\xd8\x04\xf2\x2c\x89\xda\x3d\x82\x0f\xa1\xd8\xe3\x82\x7b\xba\xea\x05\x77\x63\xdb\x82\xb1\x81\xcf\x7d\x28\x81\x99\x8a\x13\x70\xb5\x2d\xf5\x82\x33\x5b\xac\xd7\x56\x83\x11\xb3\x4b\x71\x8d\xaf\xf1\x41\xde\xbd\x7c\x46\x62\x52\x7f\x05\x05\x96\x0b\x23\xa6\x8f\x52\x2f\x70\x46\xff\x27\x80\xd6\xf2\x63\xac\x2f\xd7\xa1\x73\xa5\xd1\x54\x87\xa9\xe0\x2c\xfd\xd7\xf3\xcd\xc2\x8a\x6f\x58\x4f\x4d\x45\x42\x23\xf2\x00\x5a\xe8\xc0\x62\x5a\x19\x47\x0a\x11\x8d\xd7\xdd\xe2\x0b\xf3\x8c\xf0\x47\xba\x6a\x1a\xcb\x62\x70\x52\x94\x58\x29\x8e\x89\x2b\x3e\x39\xf2\x0c\x05\xcb\x02\xa4\x56\x10\xeb\x0d\xc2\xeb\xc1\x78\x76\x39\x4e\x36\x37\xd7\xd2\xdd\x60\x2d\x03\xe1\x1d\xaa\x6d\xe7\xa6\x52\x73\x32\x9c\x74\x13\x98\x07\x5d\x2a\xf2\xbc\xaa\x7a\x5c\x20\x72\xcd\x43\xaa\xc6\x97\x2d\xea\x80\x72\x55\x6a\x75\xf1\x36\xbb\x74\xb1\x11\x03\x2a\x25\x2b\x8c\x85\xf9\x84\xbd\x2e\x99\x1b\x5a\x3c\x58\x1c\x4f\x76\x34\xc8\xff\x1d\x97\xe6\x19\x59\x21\x66\x4c\xbf\x45\x4a\xde\x7e\x30\x7f\xdf\x81\xcf\xef\xb3\x84\x5d\xfb\x14\x4d\x43\x1e\x6b\xdf\x1d\xd8\xfd\x9b\x8d\xa7\xbc\x60\xf1\xb0\x62\xc7\x93\xaf\xd8\x44\x13\x30\x44\xec\x16\x47\x80\x42\x56\x0e\xa3\x6e\x47\xe2\x16\x51\xdd\xa3\xb4\xf9\x23\xa7\x24\xbc\x56\xf8\x1a\x2e\xc2\x70\x7b\x8d\x02\x42\x86\xd0\x73\x57\xeb\xe0\x05\x65\xa5\x16\x90\x04\xb6\x61\x83\xd7\xca\xc3\xa2\x8d\xd5\x0f\xbe\x17\x3e\xf7\x83\x2e\x7f\xaf\x31\xa6\xc7\xbe\x41\xa9\x6b\xb4\x85\xcc\xb1\xa1\x78\xec\x68\x28\x0a\x32\x63\xaa\xff\x6d\xf7\x7e\x9e\x67\xef\x07\x1a\x8a\x98\x80\xca\xc4\x4d\x33\x3e\x71\x58\xdf\xe9\x50\xe4\xf8\x0c\x0c\xdb\x87\x86\xc4\xac\x4f\x29\xe8\xb2\xf2\x23\xe8\x5a\xe4\xc5\xf1\x04\x29\xb1\x08\xc3\x50\x90\xa6\xf0\xc5\x63\x42\x1f\x29\xd4\x76\x8d\x1e\x39\x5c\x16\x53\x03\xe8\xe8\xf0\x77\x86\x7a\x10\xbd\x35\xeb\xa6\xb7\x26\x7d\xf2\x7a\x63\xee\xf9\xaa\x72\x01\x79\xd6\x69\xed\x68\x0c\xf7\xfc\x01\x8d\x3d\x1f\xd2\xa1\xc6\xc6\x84\x7f\x03\x00\x00\xff\xff\xdc\xdd\x87\x1b\xa4\x06\x00\x00")

func tplMongo_mongoGogoBytes() ([]byte, error) {
	return bindataRead(
		_tplMongo_mongoGogo,
		"tpl/mongo_mongo.gogo",
	)
}

func tplMongo_mongoGogo() (*asset, error) {
	bytes, err := tplMongo_mongoGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/mongo_mongo.gogo", size: 1700, mode: os.FileMode(420), modTime: time.Unix(1458789996, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplMongo_ormGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x55\xcd\x6e\xdb\x3c\x10\x3c\x4b\x4f\xc1\x8f\xc8\x41\x02\xfc\xd1\x40\x7b\x73\x91\x4b\xf3\x53\x18\x4d\xe3\x20\x4e\x4f\x41\x11\xe8\x67\xad\x30\x91\x48\x95\xa4\x92\x3a\x82\xde\xbd\x4b\x52\x96\xe5\xd8\x2e\xda\x8b\x23\x91\xb3\xa3\xd9\x99\x25\xd3\xb6\x39\xac\xb8\x00\x42\x2b\x29\x0a\xf9\x20\x55\x45\xbb\xae\x4e\xb2\xe7\xa4\x00\xd2\xb6\xec\xc6\x3f\x76\x5d\xd8\xb6\x27\x32\x7d\x22\xb3\x53\xc2\xf0\x8d\x57\xb5\x54\x86\x44\x61\x30\x9d\x7e\x54\x39\xa9\x13\x65\xd6\xa4\xe4\xa9\x0e\x03\x5a\xc8\xfa\xb9\x60\x5c\x4c\xab\x42\xb2\x97\x0f\x74\x7f\x69\x9a\x6a\x29\x68\x68\xab\x17\xaf\x62\xa8\xe3\xe6\xb1\x49\x59\x26\xab\x29\xbc\xa5\xcd\x1a\x7f\x51\xd0\x34\x4f\x91\x81\x91\x63\xdb\x56\x33\x32\xb5\xad\x4a\x04\x8a\x66\x5f\xc0\x2c\x54\x35\x77\x02\x35\x4a\x0d\x28\xf6\xd1\x75\xd4\x42\x40\xe4\xb8\x12\x87\xe1\xaa\x11\x19\xe1\x82\x9b\x28\x26\xad\xdd\xe1\x2b\xc2\xae\x01\xf2\xb9\xc8\xe1\x97\xad\xca\x53\xb6\x44\x26\x71\x89\x28\xfd\x38\xb7\x50\x8b\x47\xae\xeb\xa4\x42\x43\x1c\x30\xde\x92\x8e\x38\xbe\x25\x75\xcd\x45\xf1\x37\x2c\x3d\x74\xcc\x73\x0b\x05\xd7\x06\xd4\xc5\x1b\xb6\xb1\x48\x9f\x3e\xaf\xe7\xe7\x11\x1d\x67\x41\x27\x84\x0e\x0c\xf8\x22\xe0\x75\x78\xc5\x0f\xe5\xb6\x22\xde\x27\xba\x85\x4a\xbe\xc0\x9f\xa9\xb6\xc2\x0a\xc5\x7c\x41\xcf\xb6\x6d\x6f\x09\x89\xca\x1e\xc7\x52\x9d\x17\xa0\xfe\x81\xda\x73\xdc\x42\xda\xf0\x32\x1f\x75\xdf\x85\xe1\x21\x1f\x87\xb8\xde\x1b\xe7\xe3\x93\xb5\xb1\x83\x89\xad\xf3\x0c\x34\xeb\xf7\x16\xb5\xe1\x52\x68\xdc\x0f\x6e\x94\xac\x41\x19\x0e\x7a\x46\xaa\xa4\xbe\xd7\x46\x21\xe0\x07\x17\x28\x7e\x95\x64\xd0\x76\x16\xb5\x33\x43\x4b\x07\xb9\xe4\x25\x42\x2e\x39\x94\xb9\x9b\xa5\x20\x18\xf5\xb4\xc3\xe5\xff\xb4\x94\x5b\x2f\x70\x8b\x0a\x69\x1e\x12\x91\x94\xeb\x37\xc8\xad\x17\x66\x5d\x83\x5d\xf7\x40\xda\x4d\xfc\x17\xfb\xd4\x03\xfb\x8e\xdd\x07\x9b\x26\xce\x14\x24\x06\x9c\xb3\xce\x57\xef\x58\x3f\x9e\x34\xde\x02\x6f\x1a\xb3\x31\x63\x1f\xe7\x33\xf0\x6b\x77\x28\x60\x37\x89\xb6\x9b\x10\xf4\x2e\x0e\xed\xf1\xf6\x42\xc2\xfd\xa3\xb0\x6f\xbe\x57\xe5\xac\xd7\xa0\x35\xba\x3c\x21\x99\x2c\x4b\xc8\xac\xe3\x36\x8a\x9d\xb0\xd1\xcd\x33\x59\x46\xa8\x19\x6f\x1b\x50\xa4\xaf\x61\x67\xa5\xd4\x10\xc5\xdb\xb3\x7b\xe2\xcc\xb3\xf5\xf6\xae\x61\x7e\xac\xac\xef\x0a\x2e\x84\x6e\x14\xe0\x25\xe4\x20\x3d\xf9\x0c\x5d\x58\x11\x50\xca\x96\x6c\x05\x30\x0f\xf6\x2a\xed\x85\xe3\x9e\x6c\xc2\x5f\x61\x3d\x23\xf7\x43\x56\x03\x1d\x2a\x74\x19\x5f\xe1\x38\xa3\x45\x2e\x9b\xb6\xfd\x9f\x20\x7b\x8f\x98\xeb\xef\x82\xff\x6c\x80\xb8\xb0\xfc\xf3\x8c\x18\xd5\xc0\x06\xbb\x09\xf2\x7d\xdd\x12\xef\x46\xdd\xd7\xf9\xe7\x83\x75\x5d\xfc\xc9\x75\xf2\xdf\x29\x11\xbc\xb4\xd6\x06\x35\xaa\x34\xa5\x88\x28\xae\x4b\x85\xc0\xa1\x2b\xe2\x6f\xe3\xde\x06\xf2\xce\x16\x8c\x18\x2b\xd8\x85\xad\x8a\x62\xf4\x37\xb0\xc4\x3b\x0e\x9d\xe3\x81\x18\x66\x6b\xb7\xd8\xe2\x7b\x57\x47\x5a\x82\x3a\x11\x3c\x8b\x70\xd5\xee\xdb\x5e\x0a\x69\x24\x39\x12\x4c\x68\x67\x79\x74\xa8\x87\xe9\x72\xb3\x74\xe8\xbe\x8a\x78\x4e\x7c\x2c\x31\x89\x14\xe8\xa6\x34\x64\x73\x71\xb9\x76\x88\x33\xc1\xcd\x9c\x02\xd3\x28\xb1\x3b\x63\x23\x9e\xd8\xdf\x22\x84\x5d\xc9\x24\xbf\xab\xcb\xcd\x3f\x37\xf7\x4b\x49\x77\x78\x77\x25\x15\xf0\x42\x3c\x3c\xc3\xfa\x28\x66\x6b\xe0\x51\x88\x76\x47\xcd\x6d\x6f\x9a\xfe\x1d\x00\x00\xff\xff\x17\x51\x3d\xef\x65\x07\x00\x00")

func tplMongo_ormGogoBytes() ([]byte, error) {
	return bindataRead(
		_tplMongo_ormGogo,
		"tpl/mongo_orm.gogo",
	)
}

func tplMongo_ormGogo() (*asset, error) {
	bytes, err := tplMongo_ormGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/mongo_orm.gogo", size: 1893, mode: os.FileMode(420), modTime: time.Unix(1458790048, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplMongo_searchGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x57\x5f\x8b\xdb\x46\x10\x7f\x3e\x7d\x8a\x41\x84\x20\xa5\xc2\xbe\xbe\x9a\xfa\x21\xf4\x7a\xed\x51\x2e\x4d\x9b\x40\x0b\x21\x38\x7b\xd6\xc8\xdd\x54\xd2\xfa\x76\x57\x6d\x5c\xe3\xef\xde\x99\x5d\xad\x2d\xd9\x56\x63\xc1\x11\x02\x07\xd6\xce\xbf\xdf\xec\xcc\xec\xcc\xdc\x76\x9b\x63\x21\x6b\x84\xb8\x52\xf5\x4a\x2d\x0c\x0a\xbd\xfc\x33\xde\xed\xa2\xed\xf6\x99\x7a\xf8\x08\xb3\x39\x4c\xe8\x34\x9d\xbe\x71\x1c\x22\xcb\x02\x98\x33\x79\x85\x98\x7b\x22\xf1\xed\x66\x8d\xe0\x55\x26\xaf\x44\x85\xbb\x9d\x67\xdd\xd5\x39\x7e\xfa\x85\xec\x18\xab\x9b\xa5\x85\x6d\x74\xb5\xdd\x6a\x51\xaf\x10\x9e\x15\x12\xcb\x9c\x01\x9c\xd6\x8f\x68\xbd\xca\x2d\x93\x0d\xd9\x24\x49\x2f\xd3\x5a\x84\xfd\xf9\x2d\xa1\xd1\xf9\xc3\x47\xa3\xea\x59\x7c\x24\x16\x7f\x60\x4d\xac\x73\x6f\x62\x08\xec\x56\x96\x16\xf5\xcb\x3a\x77\x2e\x3e\x25\xe8\x2e\x8a\x8a\xa6\x5e\x0e\x86\x83\x09\x49\x0a\x09\x85\x44\xd6\xab\x0c\xfc\x6f\xca\xb1\xd1\x68\x1b\x5d\x43\xdc\xaa\x76\x94\x08\x22\x3b\xa2\x7b\x87\xe2\x3d\x5e\xa2\xe0\x45\x0f\x33\x05\xa7\xea\xa5\x19\xd1\x61\xe4\xc2\x0a\x0e\xc4\xff\x64\x8b\xa4\x46\xa4\xe9\xca\xd5\x04\x3e\x42\x27\x50\x10\xfb\x5b\xc5\xad\x40\x2f\x58\x33\x50\x3e\x03\x44\x7a\x69\xde\x38\xc1\xe4\x24\xa2\x69\xe6\x34\xb1\x34\x38\x64\xe4\x88\xd4\x2a\xf8\xcc\xf7\xbe\xc6\x16\xc1\x58\x30\xfa\x93\xce\xb0\x9a\xdc\xdd\x4c\x7e\xc2\x4f\x49\x1a\x5d\x2d\x95\xc6\x89\x33\x9c\x8c\x4a\x68\x06\x32\xcf\xa0\x96\x65\x06\x9c\xab\x94\x13\xfc\xb7\xd0\xe7\x12\x46\xfe\x63\x6e\x60\x0e\xef\xde\xfb\x78\x8f\x7a\x5f\x27\x21\xcf\x7a\x45\x3c\x80\xf9\x16\x75\xf5\x33\x6e\x18\xb5\x12\xeb\x77\x1e\xf7\xfd\x83\x52\x65\x17\xfc\x2f\xdc\x74\xa1\x83\x52\x80\x25\x36\xe1\xcd\x80\xba\x02\x5e\x84\x7a\x23\x2c\x8e\x46\x25\x25\x11\x50\x0f\x8d\xeb\x8e\x4a\x4e\x68\x4b\x0c\xf0\x7e\x0c\xb9\xd4\x56\x5e\x97\x59\x08\x22\x66\xbd\x06\xf3\xd9\x47\xff\x5a\x68\x51\x19\x7a\x56\xc9\x3f\x4a\xe7\x10\x9e\xfd\xda\x91\xbb\xb7\x09\x8d\xa0\x43\x92\x35\xd5\x67\x21\x96\xb8\xdd\xf1\xdb\xa5\x2b\x38\x1b\xf3\x39\xc4\x31\x3c\x7f\x0e\x25\xd6\x89\x37\x94\x32\xf1\x9a\x85\x42\x17\xa1\x0a\x72\x95\xd9\x1e\xff\xd8\x57\x40\x70\x25\x1b\x2e\xa9\xe0\xde\x59\x91\x90\xcc\xb3\xcc\x90\xa8\x34\x3a\x34\xa6\x0a\x5e\x2c\x7a\xa2\xf7\x2b\x9d\x82\x97\xff\x5d\xda\x36\x42\x97\x85\x27\x83\x52\x56\xd2\x02\x45\x26\x03\x55\x14\x06\xdd\x37\x75\x37\xf7\xde\xbc\xd1\xdf\xd0\x34\x25\xf1\x51\x6b\xa5\x5d\xd7\x33\xe1\xe6\x03\xad\xaf\x9f\xa4\x00\x4f\xaf\x58\xe8\x95\x61\x9d\xa1\xa4\xc4\x85\x56\x55\x0c\xb3\xd6\x17\x7a\xd3\x46\xfe\x8b\x54\x2a\xce\xcd\x43\xf8\x7b\xde\x3d\x36\x68\xec\xd8\xae\xc0\x9e\xd0\xb4\x08\x17\x71\x7d\xe1\xdc\x38\xbe\x27\x24\xaa\xc8\xcb\x42\xcf\xc2\x97\x85\xbf\x73\xeb\xa7\xc8\x01\x97\x23\x83\x7f\xd1\x92\xfc\xba\xb3\x19\x7a\xc9\x05\xa9\xeb\x27\xeb\x69\xb2\xf1\x6b\x83\x7a\x93\xc4\x8f\xfc\xb3\x68\xc7\x77\x76\xe0\xb4\x73\xfa\x33\x59\x4a\xd3\x7d\x8c\x8e\x7d\x3f\x8a\xd7\x1e\xbf\xf5\x3f\xf8\x9e\x5e\xd6\x37\x5e\x37\x7a\x85\x7e\xa3\x71\xd7\xbc\xc1\x12\x2d\x8e\xcd\x42\x1c\xbb\x49\x4b\x4e\x4f\xa7\x50\x28\x0d\xb9\x33\x03\x8a\x26\xa7\x64\xe5\xf3\xd6\xc3\xbc\xe4\xf2\xa8\xd6\x25\x76\x2d\x5d\xe4\x7d\x08\x82\x0b\x36\x9c\x3e\xad\x31\x89\xfc\xda\x6a\xda\xdd\xe9\xe2\x38\x3c\x34\xb2\xcc\x7d\x1e\x07\x0a\xa6\xcd\xf4\x51\xaf\xbb\x17\xeb\x35\x5d\x95\x67\xb0\xac\xa5\xed\xeb\x7a\x9e\xd3\x69\xe7\xb4\x55\x56\x94\x27\xbd\x9f\x41\xbe\x57\x4d\x6d\x13\x5f\x04\xc6\xe2\x9a\x85\xbe\xbd\xbe\xe6\x03\x6d\x09\x7c\xa2\x6f\x2e\x0d\x7f\xfe\x0e\xbc\x29\x1e\xb4\x64\x88\xee\xbb\x38\x6b\xf6\x96\xca\x27\x71\x4b\x1c\x1b\xcd\xbc\x36\x41\x38\x53\x0b\xca\x2f\x6b\xf9\xb5\x85\xcd\x38\x7b\x57\x6a\xd2\xdb\xd8\x89\xc4\xcb\xa8\x07\xfe\x66\xee\xa6\x3d\x0b\xa7\x3c\xd7\x4f\xfa\xc5\x99\xf5\xdf\x78\x4b\x3f\xd4\xe2\x81\x1e\x27\x45\x99\x17\x26\x82\x1a\xf8\x2f\x2e\x54\x02\xaf\x40\xd1\x61\x01\x6a\xa9\x6e\xf7\x89\x3a\x0b\xcf\x74\x4a\x9f\x6d\x0b\x8b\xf6\x8c\xff\x02\x00\x00\xff\xff\x3d\xa2\x79\x03\x53\x0e\x00\x00")

func tplMongo_searchGogoBytes() ([]byte, error) {
	return bindataRead(
		_tplMongo_searchGogo,
		"tpl/mongo_search.gogo",
	)
}

func tplMongo_searchGogo() (*asset, error) {
	bytes, err := tplMongo_searchGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/mongo_search.gogo", size: 3667, mode: os.FileMode(420), modTime: time.Unix(1458789996, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplMssql_configGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x52\xcd\x6e\xd4\x30\x10\x3e\xc7\x4f\x31\x58\x80\x92\xaa\xb8\xf7\x95\xf6\xd2\x16\x21\x0e\x05\xa4\x15\x27\x84\x56\x4e\x3c\x09\x6e\x13\x3b\x8c\x9d\xaa\x60\xe5\xdd\x19\x67\xcb\x92\xae\x4a\x73\x48\x3c\x9e\xef\x67\xbe\x51\x52\x32\xd8\x5a\x87\x20\x87\x10\x7e\xf6\xfb\xc6\xbb\xd6\x76\x72\x9e\x47\xdd\xdc\xe9\x0e\x21\x25\xf5\xe5\x70\x9c\x67\x91\xd2\x6b\x5f\xdf\xc2\x66\x0b\x8a\x2b\x3b\x8c\x9e\x22\x94\xa2\x90\xed\x10\x25\x7f\x8c\x8e\xba\xd6\x01\x2f\x58\x4a\x0a\x51\xec\x41\x76\x36\xfe\x98\x6a\xd5\xf8\xe1\xc2\xa0\xb3\x01\xdd\x1d\x1f\x3b\xff\x6e\xf1\x33\x75\xa6\xad\x30\xf8\xbb\x9e\x7e\xf1\xdb\x13\xe3\x4f\x9b\xb7\x83\xb7\xe4\x5d\x56\x7f\x90\xa2\x12\xe2\x5e\x53\xb6\xdf\xf3\xc5\x0e\xe9\x1e\x09\xce\x4c\xad\x76\x7f\x2b\xee\x98\x1a\x1e\x9f\x33\x06\xa9\xeb\xcb\x4c\x6b\x27\xd7\xc0\x4d\x58\x58\xf1\xeb\x58\xe6\xb1\x77\x7e\xa2\x06\x3f\xe9\x01\x21\x44\xb2\xae\xab\x20\x89\x82\xd7\xe1\xce\x01\x89\x72\xe6\x6c\xab\xae\xf8\x06\x9b\x58\x1e\xf6\x25\xcf\xe1\x29\xb9\x12\x85\x6d\x17\xc2\xab\x2d\x38\xdb\x67\x91\x62\xd4\xce\x36\x25\xef\x48\xed\x46\x96\x8e\x6d\x29\xbf\xf1\x9c\x1f\x30\x1e\x47\xfd\x0e\x7e\x44\x97\x2d\xa0\xd5\xb6\xdf\xbc\x09\x72\xf1\x55\xef\x89\x3c\x95\x55\xc5\xc2\xb3\x78\x12\x75\x0b\x6f\xd7\x61\xd3\xf5\xe5\x06\xf2\xbc\xf3\x21\xf6\x76\x29\x72\xe2\xf9\x24\xf1\x8d\x7e\xf8\xcc\x66\x39\x49\x28\x87\x55\x01\x3c\xdb\x12\xfb\x9f\x8b\x7a\x09\x5f\x3d\x27\xfd\xd1\xf4\x78\x84\x1e\x8b\xff\x4b\x3f\x8f\x3f\x91\xbe\xea\x7d\xc0\xb2\xca\x1b\xf1\x94\x65\x08\xe3\x44\x0e\x56\x6a\x8f\x90\xcc\x4b\x09\x9d\xe1\x1f\xf4\x4f\x00\x00\x00\xff\xff\xad\x03\x7f\x1c\xdd\x02\x00\x00")

func tplMssql_configGogoBytes() ([]byte, error) {
	return bindataRead(
		_tplMssql_configGogo,
		"tpl/mssql_config.gogo",
	)
}

func tplMssql_configGogo() (*asset, error) {
	bytes, err := tplMssql_configGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/mssql_config.gogo", size: 733, mode: os.FileMode(420), modTime: time.Unix(1462872133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplMssql_ormGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x58\x6d\x4f\xdb\xca\x12\xfe\x6c\xff\x8a\xe9\xaa\x45\x36\xd7\x35\xf7\x7e\x45\x8a\x2a\x5e\x42\xc9\x3d\x34\xa1\x24\xb4\xa7\x42\x08\x9c\x78\x03\x6e\xfd\x92\xae\x1d\x20\x8a\xf2\xdf\xcf\xcc\xee\xda\xb1\x4d\x5e\x9c\xa3\xf6\x48\x47\x48\xc1\xde\xdd\x99\x9d\x79\xe6\x99\xd9\x59\xcf\xe7\x3e\x1f\x07\x31\x07\x16\xa5\xe9\xcf\xf0\x2e\x11\x11\x5b\x2c\x26\xde\xe8\x87\xf7\xc0\x61\x3e\x77\x2f\xd5\xe3\x62\x61\xce\xe7\x6f\x93\xe1\x77\x38\x6c\x81\xab\xde\x02\xff\x2c\xe0\xa1\xdf\xf5\x22\x4e\xa3\x13\x11\xc4\xd9\x18\xd8\xbb\xb4\xe3\x33\x70\x69\x58\xae\x3b\xd8\x87\x5e\x1c\xce\xe0\x81\xc7\x5c\x78\x19\x87\x34\x99\x8a\x11\x87\x71\x10\xe2\x4f\x22\x20\xf3\x86\xf8\xf4\x1c\x64\x8f\xa4\x23\xf2\xc4\x0c\x7e\xf0\x19\xec\x1f\x48\x71\xe1\xc5\x68\xc9\xdb\x20\xf6\xf9\x8b\x03\x6f\xc7\xb4\x25\x6d\x47\xc6\xb8\xd2\x80\x54\xae\x0b\xc6\xc0\x7f\xea\x79\xb9\x39\x94\x0d\xac\x58\x2c\xc5\xe5\x42\x35\xec\x85\xa1\x52\x44\x13\xdf\x93\x20\x56\xca\x3f\xf2\xac\x90\x4f\x81\x39\xc0\xd4\xf2\x38\x89\x3b\x7e\x69\x46\xfa\xce\xfb\xd3\xf1\x38\x78\xf9\x7f\x59\xba\x5b\x5b\xc8\x68\x9c\x01\x63\xb9\x32\xd3\x0c\xa2\x49\x22\x32\xb0\x4c\x83\xf9\x1e\x02\xe1\xa5\xfc\x00\xe3\xc0\xf0\x7d\x1c\x65\xf4\x2f\xcd\x10\xd7\x87\x94\x1e\xb3\x20\xe2\xf8\xbf\xc0\x44\x6f\xd3\x13\x51\x47\xaa\x21\x20\x0c\x86\x41\x5b\x2c\xe4\x32\x1e\x93\x83\xb6\x69\x1e\x1c\xc0\xd1\x53\x12\xf8\x30\x4a\xa2\x09\xc1\xce\x85\x48\x04\x0d\x9f\x21\xae\x10\x27\xc0\xbd\x74\x06\xcf\xde\x0c\xb2\x04\x7c\x9e\x71\x11\x11\x29\x10\x53\xb5\x29\xe4\x8c\x48\x1f\x93\x29\x02\x38\xc4\x39\xb9\x25\xf7\xcd\x27\x4f\xc0\x1d\xd0\x32\x77\x80\x3f\xa6\x39\x9e\xc6\x23\xb0\x22\xd8\xbf\x53\x94\xd1\x54\xf8\xf4\x20\x6c\xf8\x39\xe5\x62\x66\xc9\x5f\x50\x9e\x39\xe0\x89\x87\x14\x5c\xd7\x45\xfa\x70\x31\xf6\x46\x7c\xbe\xb0\xc1\xba\xb9\xdd\xaf\x88\x3b\xca\x68\x1b\xe6\xa6\x21\x92\xe7\x54\xbe\x13\xf6\x77\xfe\xd0\xfd\xbc\x54\xab\xf4\xa1\x3a\xdb\x34\x88\x13\xb8\xe8\x4d\x0b\xe2\x20\x24\x41\x43\xf0\x6c\x2a\x62\x7a\x95\xf2\xa6\x81\x88\x61\x06\x70\x01\xa4\xd3\x3d\x09\x93\x94\x5b\x88\x58\x09\xe5\x2d\xcc\xc3\x85\xef\x09\x28\xcd\xbc\x4e\xda\x9d\x86\x21\x31\xfa\x12\xb9\x1c\x64\xc1\x13\x71\xcf\x30\x08\x25\xf4\x67\x49\xcf\xc5\x02\x30\xd0\x6e\x31\x46\x7c\x41\xc9\xfe\xe7\x8b\xc1\x6c\xc2\x73\xc5\x18\x42\x78\xaf\x5e\x54\x34\xd1\x32\x52\x25\x78\x3a\x0d\xb3\x14\xea\x28\x99\x06\x65\x94\xf4\xa5\xcb\x5f\x32\x4b\xc2\x55\x92\x80\xda\x6a\x43\x83\x28\x25\xfa\x23\x2f\x46\x26\xca\x8d\x1b\x38\xaf\x0c\xdb\xee\xbf\x5e\x67\xec\xd5\xfc\x77\xb4\x30\x0f\xd3\x7c\xcd\x9e\xa5\xcc\x74\x6b\x4b\xed\x62\x6d\x8e\x47\xed\x0d\x83\xbd\x22\xda\xaf\xc3\x8d\xf1\x96\xb2\x0d\x63\xdb\x30\xb8\xc6\x6a\xab\xa1\x45\x73\x55\x15\xa5\x18\x77\x39\xf7\x4f\xbc\x34\xcb\xf1\x31\xca\x5c\xf8\x98\x28\x1a\x58\x35\x9d\xa5\x3d\x96\x9a\xbe\x78\xe1\x94\x50\xca\xb7\x2b\x21\x6a\xec\x20\x5f\x88\x97\x40\xae\x3d\x2f\x24\x7e\x39\xfb\x5a\xe0\x4d\x26\x38\xaa\xa3\x86\x39\xb9\xa7\x9e\x6c\x99\x58\x1a\xfd\x62\x12\xc3\x60\x12\x83\xb7\x16\x88\x5e\xcc\x1b\xd5\x88\x8d\x15\xa2\x52\x1b\xae\x92\xe7\x57\xe5\xe1\xdf\x91\xe3\xf5\x8c\x5d\x26\x6c\x9e\xaf\x3b\xa5\x6b\xd3\x6c\x5d\x91\xac\x55\x66\x6d\x4a\xd5\x9a\x57\xa5\x97\xa6\x45\xf9\x97\xc7\x66\x43\x86\x36\x4e\xd0\x5f\x91\x9f\xb5\xf4\xdc\x2d\x3b\x5f\xb3\xa5\x46\x9c\x1c\x4a\x9d\x86\xcd\x52\xae\xef\x3d\x71\x8b\x5a\xbb\x6a\x3a\x61\x7e\x11\x79\xaf\xb4\xa6\x65\x6e\x21\x56\xb4\xaa\xda\x01\x12\x96\x2d\xf8\x6f\x39\x9e\x91\x9b\xa2\xe6\x4e\x9c\x72\x91\x91\xfe\x4a\x51\x50\x93\xd7\x13\xec\x7a\xb8\x9a\x5c\x6c\xb6\xb2\xaa\xab\xa9\xad\xaa\x8a\x20\x69\x58\xa7\xdb\x6f\x5f\x0d\xa0\xd3\x1d\xf4\xe0\xc6\x1f\x26\xb7\xee\x4d\x45\xc5\x2d\xe8\x54\xfa\xbe\xb1\x85\xc3\xc6\x4d\x33\x19\xbe\x1c\x5d\x5c\xb7\xfb\xb9\x98\xe0\x13\xee\x65\xb2\xff\xb3\x42\xbe\x56\x83\x0d\xec\x03\x2b\xab\x61\xa6\x21\x96\x66\xcb\xa2\x85\xae\xf4\xb9\x78\xe2\xc2\x6d\xbf\xf0\x51\x5e\xb7\x5e\x77\x9e\x92\x51\x6b\xd3\xa9\xa4\x55\x65\x54\x88\x54\x56\x08\x76\xfc\x62\x33\x9d\x17\x17\xa5\x39\x6b\x27\xad\xab\xa9\x00\xcb\x21\x57\x27\x4a\x79\x7b\xdb\x34\x57\x29\x6c\x40\x81\x25\x63\x76\xa7\xc0\xf5\xe5\xe9\xd1\xa0\xbd\x26\xfa\xfd\xf6\x40\x05\xb2\x61\x2f\x4f\x7f\x2d\x19\x4a\x06\xe8\xf1\xd7\xf3\xf6\x55\xbb\x5e\xfd\x71\xbe\xf0\x73\x87\xa8\x3a\x79\x7e\x55\xea\xea\x36\x74\xa8\x95\x3f\x9e\x75\x4e\x2d\xec\xf3\x5f\xa1\xbf\xf1\xb0\x5c\x22\xd4\x6f\x5f\xb4\x4f\x06\x50\xbe\x13\xa1\x6f\x67\x57\xbd\x4f\x6b\x50\x2b\xdc\xae\xc4\xbf\xec\x77\xe4\x56\x8f\x74\x07\x02\x5f\xfa\x52\x2d\xf0\x45\x61\xef\xd0\x1b\x2f\xee\x74\x6a\x16\xcb\xfa\x75\x1c\xa0\x3c\x42\xbd\x15\x04\xdc\xea\x78\x46\x26\x49\x49\x35\x69\x15\xef\x74\xab\x43\x05\x97\x9e\xf0\xa2\xdf\x8c\xcb\x4a\x3e\x29\x2b\x56\x32\x09\x8e\xba\xa7\xb2\x2c\x6c\x42\x6f\xa5\x23\x45\x39\x20\xd4\xe8\x78\x69\x00\x53\x63\x8c\x90\x8e\xe3\x71\xca\x33\xc0\xe6\xcb\x81\x90\x4e\x55\xf5\x98\xe2\x05\x50\xf7\x16\xd8\x50\xa9\x56\x0d\x01\x7d\x75\x29\x71\xa0\x04\x6a\x22\x7c\x2e\x8e\x15\xac\xbd\xab\xd3\xf6\x15\x1c\x7f\x83\x77\x74\xbf\xc5\x68\x63\xe1\xb4\x96\x5a\x6d\xaa\x3f\xea\x48\xc9\xa5\x5a\x80\xd7\x62\xb7\xaf\xbe\x35\x58\x7a\xd4\xd1\x6d\x62\xea\x12\xc4\x25\x05\x0e\x65\xa7\x4d\xe7\x8e\x3a\x73\x1b\x68\x62\x75\x2e\x33\x5b\xd5\xb9\x82\x0c\x65\xb1\x7f\x82\x18\xf4\x9b\x13\xc3\x34\xde\xa5\x00\xbd\xb3\x33\xac\x56\xf0\x01\xb0\xab\x4d\xe1\xac\x3d\x38\x39\x87\x6e\xfb\xcf\x62\xa4\xd7\xbd\xf8\x86\x27\x8c\xf6\xa9\x54\x6b\x35\x9b\x9a\x50\x29\x0f\xba\x0e\xb8\x66\x56\xac\xbe\x98\xe8\x6e\xa3\x41\x22\x5a\xcf\x8f\x5c\xf0\xbf\xdd\xc6\x17\xa0\x47\xee\x03\xcf\x64\x27\x8f\x65\xd4\xca\xc4\x94\x3b\x20\x55\xdb\x1b\x52\xa5\xe8\xf4\xb7\x75\x40\x64\x6c\x23\x4b\x37\x7d\x94\x58\x63\xeb\xd8\x43\xe6\xad\x33\x76\x95\xa5\xdb\x0c\x3d\x0a\x43\xbc\xd6\x5b\xeb\xbe\x00\xa8\x93\xbd\x74\x13\xca\x37\x94\x4e\x32\xd6\x68\x8f\xaf\x41\xf6\xd8\x93\xf1\xaf\xc1\xb2\xa6\x12\xfc\x26\xb4\xf2\x65\x2d\x75\x2d\x84\xff\x60\x2e\x34\x22\x3f\x4a\x4a\x93\x8a\x0b\x2a\xbd\xe5\xd6\xdb\xab\x27\x35\xcf\x9b\x84\x68\x23\x7e\x15\xa2\x26\x13\xe4\x23\x0c\x93\x24\xd4\x4e\x41\x5e\x26\xd5\xff\x0a\x14\xf7\xba\x9a\xdc\xcb\x5a\xa8\x65\xa9\x66\xd5\x51\xb8\x87\x41\xef\x12\xfe\x47\x0b\x17\xaf\x31\xba\xdf\xa5\x1c\x75\x06\xe7\x56\xb7\x77\xd1\x3b\xf9\xc3\x46\x75\xa6\xa1\x8c\x6c\x15\x15\x75\x80\xb7\x28\xc5\x01\x6a\x59\x99\x6a\x0d\xd5\x22\x2c\xce\x8c\x49\xfb\xa6\x13\x35\x72\x58\x92\x4b\xae\x11\x5c\x61\x15\x91\x24\xb9\x37\xf9\xec\xb9\x97\x5e\x0a\x8e\xa5\xcf\xd2\xa2\xa8\x5c\x56\x46\x66\xc3\xde\xde\xe6\x75\xf9\xa1\xc1\xd4\xb7\xad\xc2\x62\xa6\x9b\x11\x86\x10\xc8\xb1\xfc\x6b\x4f\x1d\x1f\x3d\xb9\xbc\x8b\xc8\x89\x6d\x69\x71\xca\xc3\x5a\x2e\x4c\xa8\x5c\xae\xa0\xfd\xb6\x36\xf4\x14\xa3\x8c\x6d\xe8\xfa\xa8\xb0\x95\x20\xe7\x6e\x54\xce\x9f\xb2\xae\x8a\x92\x3a\x18\x95\xcb\x17\x7d\x16\x29\x37\xa1\xca\x93\x9c\xdd\x07\x07\x44\xf6\x69\xc4\xe3\x0c\xf8\x8b\x17\x4d\x42\x7e\x48\x83\x98\x3a\x87\xcc\x6b\x7d\x70\x60\x48\x0d\x1e\x8e\x48\xcd\x87\x6c\x84\x87\x94\x87\x77\x51\x5f\x0f\x2b\x75\x87\x37\xb7\x25\x5c\xe6\xcc\xa3\x4b\xcf\x90\x7e\x46\xf4\xe3\xb3\x05\x6e\xb8\x11\x72\xdd\xeb\xcb\x53\xe8\x97\x40\x5f\x81\x6e\xdb\x6d\x80\x9a\x12\x07\x54\xbd\x68\x1a\x8e\x06\x3a\x75\x60\x72\xe5\xce\xce\xe1\xa9\x1f\xc3\xf9\xff\xbf\x02\x00\x00\xff\xff\x49\xf6\x34\x69\xcb\x19\x00\x00")

func tplMssql_ormGogoBytes() ([]byte, error) {
	return bindataRead(
		_tplMssql_ormGogo,
		"tpl/mssql_orm.gogo",
	)
}

func tplMssql_ormGogo() (*asset, error) {
	bytes, err := tplMssql_ormGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/mssql_orm.gogo", size: 6603, mode: os.FileMode(420), modTime: time.Unix(1462872133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplStructGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x53\xc1\x6e\xdb\x30\x0c\x3d\xdb\x5f\x41\x08\x3d\xd8\xc5\xea\x02\x3b\x06\xc8\x65\x0b\xd6\x66\xc0\xba\x01\x0b\x76\x6d\xe4\x98\x36\x94\xda\x92\x6b\xc9\x29\x0a\x23\xff\x3e\x4a\x8a\x6c\x27\x5b\xb7\xe6\x12\xfa\x49\x22\xdf\x23\x1f\x87\xa1\xc0\x52\x48\x04\xa6\x4d\xd7\xef\x0c\x3b\x1e\x5b\xbe\x7b\xe2\x15\xc2\x30\x64\x3f\x7c\x78\x3c\xc6\xc3\x70\xa5\xf2\x3d\x2c\x96\x90\xb9\x2f\x51\x02\x3e\x43\xb6\xca\x81\x35\x4a\x56\x8a\xde\xc5\xa2\x69\x55\x67\x80\x55\xaa\x7d\xaa\x32\x21\x6f\x9b\x4a\x65\x87\x8f\xb7\xb9\x56\x92\xd1\x1b\x94\x45\x78\xcb\x65\x01\xc9\x98\x40\xeb\xe7\x9a\xa5\xe1\xcc\xe2\x15\x9a\x7b\x7e\xc0\x8d\x68\x10\x88\x17\xa6\x30\xcb\x6f\x08\xb5\xf9\x6e\x20\x64\xf4\xff\xb1\x79\x6d\x1d\xed\x07\xde\x10\x67\xf0\x8a\x60\x70\x57\x3b\x2e\x49\xd3\x55\x29\xb0\x2e\x9c\x8c\x2f\x36\xd2\x33\x31\x56\xe0\x8c\x8f\x17\xed\xee\x87\x84\x30\x02\x77\x68\xee\xd4\x86\xca\x11\xba\x2d\xf2\x05\xbb\xb8\xca\xb6\x9e\x5f\xad\x6d\xf3\x22\x1b\xcf\x54\x4f\x95\x5c\xeb\x52\x0f\x4e\xef\x81\xad\x57\xae\x1f\xd1\x7a\x05\xe1\x67\xbb\x98\x7d\xcf\xf7\xb8\x33\xeb\x02\xb6\x7b\xfa\x5c\x30\x51\x30\x77\xb0\x60\x8f\xa2\xf8\xa0\x1a\x61\xb0\x69\xcd\x2b\x95\x8f\xce\xeb\xbf\x47\xc8\x08\x7e\xfa\xba\xe1\x55\xe0\x4d\xbd\x85\x1b\xd7\x8c\xb7\xe3\xa0\x8d\x7c\xe4\x74\x59\x49\x21\x9e\xa6\x1b\x09\xfd\x80\x2f\x90\x2b\x55\xcf\xa6\x47\x73\x2b\x7b\xb9\x83\xa4\x85\xeb\x71\x76\x29\x10\x31\x1b\xfe\x24\x37\x62\x92\xda\x59\x0a\x59\xd1\x2c\xa3\x0e\x4d\xdf\x49\x60\x73\x7b\xb2\x7f\x64\xf9\x5c\x73\xad\xed\xf7\x5b\x59\x4e\x13\x8b\x83\x81\x1e\x47\xec\x5b\xd5\x4d\x26\xa2\xe3\x03\xef\xe0\xec\xf0\xfa\xec\x6e\xa0\xd0\x5c\xe0\x29\x90\xec\x11\x20\x1a\x13\x43\x47\xe5\xc0\x6b\xeb\x48\x89\x2f\xc9\x44\x3d\xf4\xf5\x8f\x2d\x73\xf7\x33\xdf\xca\xa5\x5b\x8e\x13\x44\x5e\x59\x7a\x97\xd0\x51\x30\x4a\x92\xfa\xf5\x70\xa3\xfa\xcf\x22\x04\x9b\x9e\x5c\x70\xcf\xf5\x0a\x4b\xde\xd7\xe6\x17\xaf\x7b\x1c\x4b\x5f\x9a\x69\x09\x93\xfb\xcf\x1f\xb0\xbf\x3a\xc8\x25\xf2\xed\xb7\xf9\xe2\x69\x81\x7f\x07\x00\x00\xff\xff\x6a\xe6\x5b\x6b\x8d\x04\x00\x00")

func tplStructGogoBytes() ([]byte, error) {
	return bindataRead(
		_tplStructGogo,
		"tpl/struct.gogo",
	)
}

func tplStructGogo() (*asset, error) {
	bytes, err := tplStructGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/struct.gogo", size: 1165, mode: os.FileMode(420), modTime: time.Unix(1458790048, 0)}
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
	"tpl/mongo_collection.gogo": tplMongo_collectionGogo,
	"tpl/mongo_foreign_key.gogo": tplMongo_foreign_keyGogo,
	"tpl/mongo_mongo.gogo": tplMongo_mongoGogo,
	"tpl/mongo_orm.gogo": tplMongo_ormGogo,
	"tpl/mongo_search.gogo": tplMongo_searchGogo,
	"tpl/mssql_config.gogo": tplMssql_configGogo,
	"tpl/mssql_orm.gogo": tplMssql_ormGogo,
	"tpl/struct.gogo": tplStructGogo,
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
	"tpl": &bintree{nil, map[string]*bintree{
		"mongo_collection.gogo": &bintree{tplMongo_collectionGogo, map[string]*bintree{}},
		"mongo_foreign_key.gogo": &bintree{tplMongo_foreign_keyGogo, map[string]*bintree{}},
		"mongo_mongo.gogo": &bintree{tplMongo_mongoGogo, map[string]*bintree{}},
		"mongo_orm.gogo": &bintree{tplMongo_ormGogo, map[string]*bintree{}},
		"mongo_search.gogo": &bintree{tplMongo_searchGogo, map[string]*bintree{}},
		"mssql_config.gogo": &bintree{tplMssql_configGogo, map[string]*bintree{}},
		"mssql_orm.gogo": &bintree{tplMssql_ormGogo, map[string]*bintree{}},
		"struct.gogo": &bintree{tplStructGogo, map[string]*bintree{}},
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

