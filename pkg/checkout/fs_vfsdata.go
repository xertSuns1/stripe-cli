// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package checkout

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// FS statically implements the virtual filesystem provided to vfsgen.
var FS = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/css": &vfsgen۰DirInfo{
			name:    "css",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/css/global.css": &vfsgen۰CompressedFileInfo{
			name:             "global.css",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 10323,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\xfb\x6e\x1b\xb9\xd5\xff\x5f\x4f\x41\x38\x08\x20\x05\xa2\x4c\x5d\xe3\xcc\xe2\x0b\xf2\xa5\x6d\x80\x05\x1a\xa0\x68\x80\x02\x6d\xb1\x58\x50\x33\x1c\x0d\x6b\x0e\x39\x25\x29\x5b\xca\xc2\x7d\xf6\x82\xb7\x19\xce\x4d\x96\xd3\x2c\xba\xb5\x77\x63\x0d\xe7\xf0\xf0\xf0\xc7\x73\xa7\x6e\xdf\x80\xbf\x60\x49\xf1\x9e\x11\x05\xde\xdc\x4e\x12\x29\x84\x06\xbf\x4c\x00\x80\xf0\x20\xf1\x19\x8a\x3c\x57\x44\x27\x40\x1e\xf6\x78\x8a\xe6\xc0\xff\xb7\x40\xeb\xd9\x0f\x0d\xd5\x5e\xc8\x8c\xc8\x3e\xd5\x72\x1b\x53\x31\x7a\x28\x06\x58\x6d\x62\x9a\x92\x66\x7d\x8a\xb7\x31\x45\x86\xe5\x7d\x9f\xe4\x9d\x27\xd9\x8b\xec\x0c\x53\xc1\x84\x4c\xc0\x03\x96\xd3\x86\xad\x27\x28\x08\xce\x18\xe5\x64\x80\xc8\x70\xf6\x54\x38\x4d\x09\xd7\x81\xe6\x15\x42\xbb\x5d\x8e\xa2\x15\x72\xc1\x35\xcc\x71\x49\xd9\x39\x01\x10\x57\x15\x23\x50\x9d\x95\x26\xe5\x1c\x7c\x64\x94\xdf\x7f\xc6\xe9\x17\xfb\xfc\x49\x70\x3d\x07\x0a\x73\x05\x15\x91\x34\x77\x4c\x24\xce\xe8\x51\x25\x60\x57\x9d\xdc\x40\x2e\x64\x09\x1f\x69\xa6\x8b\x04\xec\x10\x32\xc3\x4f\x93\xc9\xed\x1b\xf0\x11\x2b\x62\x4e\xe6\x8d\x3d\x95\xbd\x38\x41\x45\xbf\x52\x7e\x48\x80\x03\x1d\xee\x85\xa5\x35\x62\x59\x92\x96\x68\x6e\x77\x5d\x91\xed\x26\xed\xb3\xa2\x5f\x49\x02\x96\x5e\x8c\x16\x24\x0d\x92\x0e\x93\x47\xb2\xbf\xa7\xda\xb1\x51\xa5\x10\xba\xb0\x52\x60\xae\x29\x66\x14\x2b\x92\x19\x31\x8a\xe5\x7c\x52\xac\xe6\x93\x62\x3d\x9f\x14\x9b\xf9\xa4\xd8\xce\x27\xc5\xce\x0a\x76\x89\x7b\x89\xe5\x81\x72\xa8\x45\x95\x80\x95\x13\xc6\x0f\xed\x85\xd6\xa2\x4c\xc0\xc6\x41\x52\x2c\x9b\x4d\x3a\xe1\x57\x6f\x07\x84\x6f\x9f\xf2\xcc\xce\xdc\x34\x33\x1f\x89\x53\xc5\x2d\x42\x5d\x28\x36\x03\xdc\x1a\xed\x9d\x85\x63\xf9\x23\x3e\x8b\xa3\x36\x07\xb3\x50\x12\xd6\x56\x93\x51\x55\x31\x7c\x4e\x40\xce\x88\xe5\x63\xfe\xc2\x8c\x4a\x92\x6a\x2a\x78\x02\xa4\x78\x34\xc3\xfe\xa4\x97\x08\xbd\x76\x7b\x3d\x85\xc3\x7f\x77\x87\x9c\x04\x15\xce\x32\x8b\xf0\xe6\xce\x0d\x60\x46\x0f\x1c\xa6\x82\x6b\xc2\x75\x02\x8c\x82\x12\x69\x5e\xfc\xe3\xa8\x34\xcd\xcf\x43\xaf\x0a\xbf\x51\x7c\xd4\xc2\x2e\x44\x39\x0c\x63\x4b\x84\x1e\x8a\x06\xe9\x04\x20\x4f\xf6\x64\xf7\x64\x20\x24\xd2\xee\xaa\x73\x14\xeb\x95\x3b\x0b\x43\x55\xe1\x73\x69\x2c\x45\x1d\xcb\x12\xcb\xf3\x10\xf9\x0a\x35\xe4\x25\xa6\x7c\x6e\x3f\x79\x61\xaf\x43\x2d\x15\xec\x58\xf2\x2b\xf7\x1a\x40\x75\x78\x29\xc2\xf2\x86\xaa\x91\xc2\x2e\xec\x31\x77\xa7\xdc\x98\xe0\xac\x8d\xc1\xcf\x3f\x33\x71\x10\xce\xfe\x70\x7a\x7f\x90\xe2\xc8\x33\x48\x4b\x7c\x20\x61\xae\x21\x70\x23\xb3\x58\x94\x95\xd7\xa6\x68\x9a\x53\x33\x23\x3e\xa6\xbc\xf3\x4e\x92\x8a\x60\x9d\x00\x2e\xfc\xc7\x9e\xaa\x38\xb1\x18\x39\x60\x06\x35\x39\xe9\xbe\x61\xb5\x74\x15\x00\x43\x04\x2d\x12\x31\x54\xb1\xc2\xaf\x9d\x88\xd6\x5a\x6a\x08\xdf\xb6\x6c\xd0\x9a\xe5\x32\x3a\xf5\x9c\x12\x96\x41\x22\xa5\x90\x7d\x09\x62\xdf\xd9\x93\x81\x91\x5c\xff\xa7\x12\x18\xfb\xfb\x24\x64\x19\xac\xcf\x1e\x9c\x14\x8f\x91\xf2\x39\x9f\x06\x90\x21\x67\x78\x4f\x58\xd7\x6b\x84\x45\x87\xdc\x41\x47\x7f\xbd\xf9\xd5\x4a\x4a\xb9\x95\x74\xcf\x44\x7a\x1f\xc4\xf9\x91\x57\x47\xad\x82\x40\xd4\x3c\x39\x35\x57\x84\x91\x54\xcf\x27\x76\xe8\xef\xfa\x5c\x91\xff\xbb\x31\x78\xdc\xfc\xd4\x1e\xe3\xc7\x72\x4f\xe4\xcd\x4f\xde\xcb\xbb\x70\xba\xac\x4e\x40\x09\x46\xb3\xf8\x6c\xdd\x4b\x0b\xac\x8f\x00\x21\x98\x38\x22\xf7\x34\x6b\xb9\x90\x6d\x75\xf2\xe8\x35\xba\xb9\xf1\xba\xd9\x71\x45\x5a\x62\xae\xa8\x33\x3b\x1b\x6d\x0a\x9c\x89\x47\x80\x16\x2b\x05\x08\x56\xa4\xad\xb2\x09\x78\x2c\xa8\xb6\x83\xb0\x14\x5f\x4d\x24\x24\x58\x62\x9e\x12\xa3\xc5\x9c\xc4\x91\x63\xe8\xdd\xd0\x58\x08\xba\xeb\xd5\x7a\xb5\xcd\x82\xca\x59\xb0\x92\x5c\xa4\x47\x35\x84\x66\x78\xb3\x3f\x6a\x2d\x78\x78\x5a\xd8\xbf\x24\x6b\x62\xa7\xdd\x8d\x71\x75\xe6\xd7\xe0\x6b\x13\x89\x2d\x9a\x83\xe5\x76\x39\x07\xab\xe5\xd2\x64\x13\xeb\xd9\xdc\xbf\x36\xff\xa3\x7e\xf6\xf3\x76\x36\x9f\x00\x00\x3c\x9f\xcd\x38\x1f\xb3\x21\x71\xd4\x46\x63\x9a\x1d\x7e\x85\x94\x67\xe4\x94\x80\x77\xed\xdd\x25\x15\xc3\x29\x29\x04\xcb\x88\x1c\xdc\xe4\x38\x41\xd0\x9f\x16\xc9\x33\xde\xc1\xa9\xee\xef\x0a\x92\xde\xef\xc5\x29\x28\x6f\xea\x9f\x61\x63\x37\x95\x08\x1a\x21\x09\xc3\x9a\x3e\xb8\x73\x3a\x4a\x65\x38\x57\x82\x06\xe7\x3a\xc4\xc0\x4a\x69\xd9\x88\x0a\xa7\x54\x9f\x13\x10\x5b\x99\x74\xda\xb8\xf3\xa6\x3d\xc0\xa0\x35\x64\x3f\x74\x84\xc2\x7b\x25\xd8\xd1\x69\xa1\xf1\x2f\x7e\x81\xda\x95\xec\xda\x8a\xbe\xeb\xf9\x64\x0f\x51\xad\xca\xdf\x68\x7f\xde\xa0\x62\x0b\xc2\x8c\xc5\xa6\x73\x01\x21\xa7\xb2\xe0\x5f\x63\xdb\xfd\xad\x68\xef\xa5\x2d\xd8\x41\x92\x5d\xd8\x44\x0f\xf3\x91\x90\xd1\x8f\xb3\x47\xc9\xa6\x37\x85\xd6\x95\x4a\x6e\x6f\x95\x16\x12\x1f\xc8\xe2\x20\xc4\x81\x11\x5c\x51\xb5\x48\x45\x79\xab\xb4\xa4\x15\x81\x0a\x97\x26\x1f\xb7\x13\xd5\x2d\x4d\x05\x77\x22\x94\x58\xde\x2f\xd4\xc3\xe1\xa6\xbb\xc4\x60\xdc\xed\xc5\xec\x01\xc5\x69\x74\x10\x1a\xa8\xcd\x3f\xc1\xac\xbe\x58\xc7\x1f\x8c\xca\x85\x81\x76\xb2\xe3\x03\x48\xdf\x23\xd7\xf9\xd8\xaf\x84\xc4\x83\x14\x1c\x66\xe2\x91\x5f\x07\xc6\x1c\x0c\x82\x12\x39\x05\x23\xbe\x8d\x2f\x40\x8b\x0a\x6c\xd1\x6b\xa3\x7f\x5d\xe9\x1d\x88\x68\xb1\xdb\x92\xd2\x26\x9a\xf3\x56\x4e\xe3\x10\x4a\x70\xae\xad\xdf\x6a\x0d\x26\xb0\x54\x90\x9c\x2a\xcc\xb3\x36\x84\x41\x25\x23\xda\x42\x3c\x04\xc7\xd7\xf7\x4f\x11\x9d\x33\xb7\xdf\x96\x65\x45\x9a\x22\x2a\x83\x6d\xbf\x62\xd9\x20\xd4\xd9\x08\xe5\x0f\xd8\xb8\xa8\xe7\x32\xc1\xe8\x28\x1a\x47\xbc\xd8\xd4\x61\x40\x94\x7b\xe1\xec\xb8\xce\x63\x52\x33\x06\xfd\xd8\x8b\x12\xf5\xa7\xfe\xfc\x28\x29\x6a\x8d\x77\xcc\xa3\xe3\x53\x51\xe4\x67\x43\x42\x86\x06\xf9\xbf\x0f\x7e\x94\x4a\xa5\x61\x5a\x50\x96\x0d\xac\xf6\x1e\xc4\x3a\xd0\x90\x0e\xad\xdd\xca\xa7\x3a\x4f\xe8\x19\x29\x18\xbe\x56\x88\x86\x72\x70\xff\x00\x5d\x92\x63\x00\x9d\xcb\x41\x6b\x50\xe4\xee\x90\x49\xa6\x63\xb9\x7a\xc0\x8e\x49\x8a\xfe\x7b\xd2\xbd\x14\xc6\x21\xc5\xfa\x5e\xc2\xc5\x48\x5d\x85\x5d\x4f\xb7\x46\xb5\xeb\xca\xf5\x9e\x43\xe3\x5a\x65\xbe\x72\x39\xc1\xd9\xf9\xbb\xda\x51\x5d\xce\x75\xaa\x93\x8e\x03\x0a\xc9\x48\x77\xe6\xfb\x28\xeb\xec\x70\xe8\xfb\x97\x67\x98\xc4\xdb\x4d\xb8\xd0\xd3\x68\xb7\xb3\xd6\x76\x9d\x87\x1e\xe7\xe8\x66\x3b\x76\x22\x87\x26\x71\x9f\x81\x3a\xfb\xff\x26\xeb\xf7\x3d\x42\x5b\xf4\x28\x60\xe2\x23\xa3\xfc\xde\x3a\x70\x57\x09\x75\xd2\xae\xf1\x84\xeb\x99\x52\x72\x2c\x45\x46\xad\x3a\xd3\xe6\x00\x21\x4f\x6a\x95\xee\xbb\x81\x72\x7b\xe7\xca\xed\x5e\x9c\xbe\x98\x42\xf7\x13\xa8\xa7\x50\xf4\x35\xa1\x3f\xa7\x4c\x1b\xe9\x52\xc1\xb5\xc4\x4a\x4f\x97\xcb\xed\xeb\x59\x44\x8a\x53\x53\xc5\x58\x5a\xbb\x56\x2e\x64\x99\xb8\x8f\x0c\x6b\xf2\xd7\x29\xaa\x4e\x33\xa0\x52\xcc\xc8\x14\x2d\xde\xdd\xb9\x8e\xa9\xe7\xba\xb7\x27\xcd\x89\x52\x53\xd7\x78\xae\xd9\x66\x54\xe1\x3d\xf3\xf5\x66\x14\x65\xb7\xf1\x3e\x5b\x79\x74\xe8\x9f\x19\x01\x80\xe3\x62\xca\xd6\x23\x63\x56\x6f\xfb\x1a\xfc\x34\x99\xe0\x6b\x7b\x2e\x19\x49\x85\xc4\x0e\xc6\x50\x7b\x5e\xae\x4e\xf0\x00\x88\xad\xed\xde\xcd\x3c\x5d\x84\xe0\x20\xe1\xb6\x29\x30\x45\x46\xdc\x59\xd5\x79\x05\x66\x4c\x04\x8d\xef\xe9\x66\x74\xef\x30\xeb\x69\xd7\x35\xba\x5a\xe2\x53\xdd\x47\x5a\x21\xdf\x4e\x35\xdb\xca\x99\xc9\xf1\x42\x93\x33\x15\x19\x99\x4f\x2a\x49\xfa\x6d\xf3\x9b\x2f\x9f\xc0\x67\xc1\xc5\xcd\x1c\xdc\xfc\xf8\xf1\x33\xf8\x13\x23\xa7\x7a\xe0\x33\xe1\xcc\x7c\x28\x05\x17\xaa\xc2\x29\xe9\xf6\xb3\x42\x9b\x4a\x55\x98\x2f\xcc\x2a\xbf\xd6\x02\xa6\xcc\xb0\x89\x3e\xf8\x03\x23\x46\x8d\x40\x5c\xfc\xd7\x60\xcb\x0c\x12\xff\xfe\x97\x06\xcf\x7e\x4f\xed\xcf\x44\x55\x82\x2b\xfa\x40\xcc\x19\x9a\xf9\x1f\x4a\x92\x51\x0c\xa6\x51\x7f\xfa\xed\xca\x9a\x86\x61\xd4\xea\x7c\x5f\x6a\xda\x0e\xb4\x6d\x2d\xb1\xd2\x58\x6a\x47\xd0\xea\x76\xfb\x96\xb1\x19\x2f\x29\x0f\x2b\xaf\xc3\xe8\xd3\xc4\x2f\xde\x6b\xcf\x8e\x14\x28\x4d\xe3\xb3\x9e\x3a\xd4\xb7\x1e\xeb\x96\xd6\x93\xe2\x9e\x75\xaf\x0a\x89\xe8\xea\xfe\x72\x2f\xfa\x3c\x79\xa8\xb5\xc8\x44\x02\x54\x45\x39\x27\xf2\xb6\x92\x22\x25\x4a\x51\x7e\x00\x4a\x63\x4d\xe6\xc0\xb6\x56\xd5\x1c\x60\x4e\x4b\x6b\xc0\xf6\x38\x26\x0b\x3f\x63\x5e\x7f\x4a\xf6\x24\x17\x92\x44\x03\xa1\x82\xea\xd9\xc9\x36\x14\x5b\x8e\x30\x76\x22\xaf\x72\xfb\xd3\xd1\xb3\x95\xb7\x37\x8b\x0a\xe5\x99\x3d\x38\xf8\xce\xfc\x74\x0a\xd5\xea\x54\xdf\x30\x0c\xf7\x8a\x3c\x0c\xe1\x04\x1b\xeb\x0c\x16\xdd\x94\x60\x94\x2b\xa2\x7d\x36\xe9\x05\x08\x1d\xc4\x21\x57\xfd\xb7\x29\x72\xf7\x53\xa5\xba\xf8\x7e\xfc\xdd\xd3\x55\x60\x0e\xb7\x9b\x6a\x7d\xbe\xb9\x19\x60\xd4\xf6\xde\x8b\x4d\x77\xf3\x8b\xde\xf5\xc0\xb5\xf1\xd9\xcd\x75\x18\xd5\x6c\xac\x41\x43\xb4\xf0\xa8\xb9\x56\x58\xf3\xdc\x43\x11\x0a\x49\x5d\xaf\xdc\x71\x5b\xd6\xa4\x57\x90\xd4\x5d\xdd\xa0\xa1\x09\x60\x02\x1b\x1b\x06\x2b\x53\xec\xe5\x94\x53\x4d\x6c\x64\x01\xcb\xc5\x56\xd9\x6e\xef\xd5\xb4\x4f\x43\x47\x30\x8a\x64\x23\xd5\xcb\x91\x44\x7e\x76\xf8\x83\x62\x24\x97\x31\x92\xfd\xbd\xf7\x61\x42\xcf\xc0\x88\xbe\x15\xc3\x2b\xe1\x33\xc8\x7d\x08\x6c\xef\xc9\x39\x97\xb8\x24\xaa\x26\x37\x28\xa2\xd7\xde\x35\x0d\x58\x95\x14\xc6\xfd\x4c\x51\x46\x0e\x33\xe7\x7f\x2f\xbc\x7c\x9a\x00\xeb\xd7\x9e\x65\xb7\xde\x5d\x64\x18\xbd\x36\xee\xf1\xc3\xff\xa6\xd8\x36\xd3\x39\x2a\x2d\xca\xa6\x7b\x62\xef\xf8\xbc\xf2\xf6\x2a\x97\x45\xfb\x7d\xbf\xe7\xfc\x8a\xdc\x99\xdf\x01\x9d\xdd\x75\x6e\x88\xd7\x46\x71\x57\xdd\x7b\xaa\x6f\xb8\x4c\xb5\xd9\x06\xdc\x13\xfd\x48\x08\xff\xa1\x7d\x31\xbd\xa9\xbf\x95\xd0\x48\xde\x2e\x66\xe2\xcb\x5f\x19\x3b\x78\x33\xe7\x9f\x47\xcc\x35\xd5\x67\xa8\x88\xd6\xc3\x98\x5c\xbc\xdd\x6d\x7d\x43\x60\x94\xeb\x78\x25\x15\x45\xa1\x6d\xf7\x86\x1d\x8d\xdd\x94\x0e\x2d\xb1\xa0\x3c\x95\x36\x9b\x82\x7b\xcd\xe3\x3d\x5b\xd9\xd0\x70\x91\x39\x7e\xa9\xd0\x71\x50\x4f\x93\x91\x7b\x9d\xa0\xa8\xd6\x2b\x42\xe3\x1d\x61\xa8\x17\x2e\x4f\x10\x47\xdd\x9e\xe0\xbe\xdf\x73\xe1\x56\x2e\x6a\x40\x3b\xc5\xfe\x7f\xeb\x79\x48\x06\x72\x7f\xed\xda\xfe\xd6\x43\xe4\x98\xd0\x62\xa3\x2c\x15\x74\xb7\xdb\xf5\x2b\x98\x53\xc6\x60\x29\x32\x92\x80\xbd\xd0\x45\xfb\xa5\xa6\xa5\x49\x49\xf3\x23\xf7\x4a\xda\xba\x2f\x69\x15\x49\xbd\x1b\xdf\xde\xea\xf6\x72\xfa\x7b\x2c\x3f\xe4\xc1\xb7\xe8\x35\xf0\x75\xd8\xed\x1b\xc0\x09\xc9\x80\xc2\xd8\xee\x19\x30\x21\x2a\x90\xfc\x1e\x04\x0f\x30\x2a\x77\xc2\x75\xe1\x3a\x08\xd3\xe5\xac\xbd\x09\x98\x11\x6b\x12\xa8\xfb\x0d\x8b\x4b\x4c\x56\x23\x4c\x76\xa8\x54\x2f\xe1\xb3\x1e\xe1\xb3\x5c\xbd\x90\xd1\x66\x8c\xd1\xdd\x0b\x19\x6d\x47\x18\xad\x36\x2f\x64\xb4\x1b\x61\xb4\x46\x81\x51\x41\xb3\x8c\xf0\xe1\x0b\x86\x38\x2a\x05\xfd\x6a\x87\xa5\xf6\xe5\xe6\x58\x5b\xe1\xae\xd5\x56\xd8\x0e\x05\xa4\x9a\xd1\xf2\x12\xa3\xa8\x3f\xb1\x6c\x62\x50\x2c\xa5\xb3\xc1\xab\x85\xec\xf4\x3a\xae\x17\xaa\x2b\xc4\xbf\x03\x00\x00\xff\xff\xca\x7a\xb9\x47\x53\x28\x00\x00"),
		},
		"/html": &vfsgen۰DirInfo{
			name:    "html",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/html/cancel.html": &vfsgen۰CompressedFileInfo{
			name:             "cancel.html",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 910,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4b\x4f\xe3\x3e\x10\xbf\xf3\x29\xe6\xef\x7b\x33\x7f\xb4\x1c\x10\x72\x22\xad\x60\x0f\x9c\x76\x25\xf6\xc2\x09\xb9\xf6\xb4\x71\xf1\x23\xf2\x4c\x0b\x5d\xc4\x77\x5f\xa5\x4e\x69\x0a\x97\x3d\xc5\x9e\xdf\x63\x5e\x8e\xfe\xef\xee\xe7\xed\xef\xc7\x5f\x3f\xa0\x97\x18\xba\x0b\x3d\x7e\x20\x98\xb4\x6e\x15\x25\xd5\x5d\x00\xe8\x9e\x8c\x1b\x0f\x00\x3a\x92\x18\xb0\xbd\x29\x4c\xd2\xaa\xad\xac\x16\xd7\x0a\x70\x02\xc5\x4b\xa0\xee\x41\x8a\x1f\x08\x6e\x7b\xb2\xcf\x79\x2b\x1a\x6b\x78\xa6\x4f\x26\x52\xab\x1c\xb1\x2d\x7e\x10\x9f\x93\x02\x9b\x93\x50\x92\x56\x7d\x07\x47\x31\x43\x5e\xc1\x27\x9f\x43\x9a\x6a\x12\x7c\x7a\x86\x42\xa1\x55\xde\x8e\xe2\xbe\xd0\xaa\x55\x2b\xb3\x1b\xaf\x8d\xb7\x59\x81\xec\x07\x6a\x95\x8f\x66\x4d\xf8\xba\xa8\xb4\x63\x99\x27\x39\xcb\x3e\x10\xf7\x44\x72\x34\xe9\x45\x06\xbe\x41\xb4\x2e\x6d\xb8\xb1\x21\x6f\xdd\x2a\x98\x42\x8d\xcd\x11\xcd\xc6\xbc\x62\xf0\x4b\xc6\x94\x4b\x34\xc1\xff\x21\xbc\x6e\xfe\x6f\x2e\x4f\xf7\x26\xfa\xd4\x58\xe6\x7f\xcb\x86\x2c\x46\xbc\x45\xcb\x8c\xeb\x90\x97\x26\x9c\x6b\xeb\x80\x80\x8b\x3d\x55\xb6\xe1\x86\x0f\x93\x39\x94\xb4\xfb\x86\xaa\xd3\x58\x89\x87\x65\x61\xdd\xd6\x78\x5c\x66\xb7\x9f\x9c\x9c\xdf\x81\x0d\x86\xb9\x55\x5c\x16\x25\x67\x51\x15\xf9\x82\x45\xe3\xd3\x07\x36\x2d\x9f\xca\x8c\x50\x03\x33\xca\x17\x8b\xca\x78\x7a\x0a\x79\x9d\xc7\xea\x9c\xdf\xcd\x0c\xb1\xc2\xb3\xc8\xb9\x7a\x30\xfb\x48\x49\x16\xbc\x8d\xd1\x94\x3d\xd8\x1c\x87\x40\x42\x6e\xb1\xf3\xf4\x72\x9e\xb7\xbf\xec\x1e\xf3\xb6\x80\x10\x0b\x4c\x42\x78\x31\x0c\xd6\x24\x4b\x81\x9c\xc6\xfe\xf2\x5c\x71\xd5\x1d\x5f\x14\x3c\x10\xb3\xcf\x09\xee\xef\x6e\x40\xf3\x60\x12\x78\xd7\x2a\xae\x51\xd5\xbd\xbd\x35\x13\xe3\xfe\xee\xfd\x5d\xe3\xc8\xe8\x34\xf6\x57\xf3\x66\x66\xbd\x7d\xea\x63\x7a\xd2\xa7\x39\x9f\xb8\x1f\x47\x8d\x75\x47\x1a\xeb\xdf\xf7\x37\x00\x00\xff\xff\xf3\x28\x79\x86\x8e\x03\x00\x00"),
		},
		"/html/redirect.html": &vfsgen۰CompressedFileInfo{
			name:             "redirect.html",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 482,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\xb1\x6e\xc2\x30\x10\x86\x77\x9e\xe2\xea\x85\x30\x34\x1e\xba\x54\xe0\x44\xaa\xa0\x43\xd5\xa1\x48\xb0\x74\x34\xce\x81\xdd\x3a\x36\xb2\x2f\x48\x08\xe5\xdd\xab\x60\xa7\xd0\x4e\x77\xe7\xfb\xbf\xff\x7c\x3a\xf1\xb0\xfa\x58\x6e\x3f\xd7\xaf\xa0\xa9\xb5\xf5\x44\x0c\x01\xac\x74\x87\x8a\xa1\x63\xf5\x04\x40\x68\x94\xcd\x90\x00\x88\x16\x49\x82\xd2\x32\x44\xa4\x8a\x75\xb4\x7f\x7c\x66\xc0\x73\x93\x0c\x59\xac\x37\x14\xcc\x11\x61\xa9\x51\x7d\xfb\x8e\x04\x4f\xcf\x77\xbc\x93\x2d\x56\xac\xc1\xa8\x82\x39\x92\xf1\x8e\x81\xf2\x8e\xd0\x51\xc5\x5e\xa0\xc1\xd6\x83\xdf\xc3\x3f\x9f\xdb\x98\x84\x41\x0c\xaa\x62\x9a\xe8\x18\xe7\x9c\x7f\xc5\x32\x5e\xf5\xa5\xf2\x2d\x3f\x3d\x71\x56\x0b\x9e\x84\x7f\xa8\x54\x00\x9c\x64\x80\x04\x40\x95\x27\x15\xd3\xcb\xa5\x5c\x77\x3b\x6b\xa2\x96\x3b\x8b\xef\x78\xee\xfb\xe9\x6c\x91\x89\x6c\x1f\xb0\x31\x01\x15\x6d\xfd\xf8\xb1\xe2\x92\x15\x00\x11\x63\x34\xde\xbd\x35\x73\x18\xcc\x36\xb9\x5c\xf5\xfd\x34\x6b\xfa\x59\x49\x1a\x5d\xb1\xef\x9c\x1a\x36\x87\x22\x60\xec\x2c\xcd\xe0\xe6\xa2\xbc\x8b\xde\x62\x69\xfd\x61\xec\x2e\x7e\xf1\x94\xdd\xef\x26\xf8\x78\x1f\xb1\xf3\xcd\xb9\x16\xfc\x1a\x26\x82\xa7\x8b\xfe\x04\x00\x00\xff\xff\xb2\x15\xd3\xdd\xe2\x01\x00\x00"),
		},
		"/html/success.html": &vfsgen۰CompressedFileInfo{
			name:             "success.html",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 2617,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x4d\x6f\xdc\x36\x13\xbe\xe7\x57\xcc\x4b\xbc\xc7\x4a\xb4\x93\x1c\x02\x43\x52\x51\xd8\x45\x61\x20\x69\x83\x26\x3d\x04\xae\x61\x70\xc9\xd9\x15\x6d\x8a\x54\x39\xa3\xb5\xb7\x69\xfe\x7b\x41\x7d\xec\x6a\xbd\xde\xba\x70\x0f\x8b\x15\x39\xcf\x7c\x70\xe6\xe1\x23\x15\xff\xbb\xf8\xe5\xfc\xf3\x97\x8f\x3f\x42\xcd\x8d\xab\x5e\x15\xe9\x0f\x9c\xf2\xab\x52\xa0\x17\xd5\x2b\x80\xa2\x46\x65\xd2\x03\x40\xd1\x20\x2b\xd0\xb5\x8a\x84\x5c\x8a\x8e\x97\xd9\x3b\x01\x72\x34\xb2\x65\x87\xd5\x27\x8e\xb6\x45\x38\xaf\x51\xdf\x85\x8e\x0b\x39\x6c\xcf\xfc\xbd\x6a\xb0\x14\x06\x49\x47\xdb\xb2\x0d\x5e\x80\x0e\x9e\xd1\x73\x29\x7e\x00\x83\x4d\x80\xb0\x84\x47\x71\xfa\x34\x43\x10\x67\xfd\x1d\x44\x74\xa5\xb0\x3a\x39\xd7\x11\x97\xa5\x58\xaa\x75\x5a\xe6\x56\x07\x01\xbc\x69\xb1\x14\xb6\x51\x2b\x94\x0f\xd9\x00\x9b\xca\xdc\xb9\x13\x6f\x1c\x52\x8d\xc8\x53\x90\x9a\xb9\xa5\x33\x29\xb5\xf1\xb7\x94\x6b\x17\x3a\xb3\x74\x2a\x62\xae\x43\x23\xd5\xad\x7a\x90\xce\x2e\x48\xfa\x10\x1b\xe5\xec\x9f\x28\xdf\xe5\x27\xf9\xe9\x6e\x9d\x37\xd6\xe7\x9a\xe8\xdf\x65\x93\xc4\x8a\xad\x96\x9a\x48\xae\x5c\x58\x28\xb7\xef\x3b\x34\x08\x28\xea\x5d\x65\xb7\x94\x53\xdf\x99\xbe\xa4\xf5\x1b\x29\xaa\x42\x0e\xc0\x7e\x58\x72\x98\x56\x7a\x5c\x04\xb3\x19\x23\x19\xbb\x06\xed\x14\x51\x29\x28\x66\x31\x04\x16\x83\xe5\xc0\xd6\x28\xeb\xb7\xb6\x71\xf8\x18\x67\x80\x61\x63\x06\x39\x08\x31\x20\x6e\x6e\x5c\x58\x85\x54\x9d\xb1\xeb\x59\x40\x39\x98\x67\x3b\xfb\xde\xad\xda\x34\xe8\x39\xa3\xae\x69\x54\xdc\x80\x0e\x4d\xeb\x90\xd1\x64\x6b\x8b\xf7\xfb\x79\xeb\xd3\xea\x4b\xe8\x22\x30\x12\xc3\xe8\x08\xd4\x69\x8d\x68\xd0\x14\xb2\x3e\xdd\x87\xbf\xad\x26\x3a\xc1\x27\x24\xb2\xc1\xc3\xe5\xc5\x19\x14\xd4\x2a\x0f\xd6\x94\x82\x86\x5d\x51\x7d\xfd\x9a\x8f\x88\xcb\x8b\x6f\xdf\x0a\x99\x10\x55\x21\xeb\xb7\xf3\x93\xcc\x0e\xf6\xe8\x10\x23\x9f\xf7\x1a\x79\x5a\xfd\x8c\x0f\x0c\xc4\xd8\xd2\xd9\x7e\x6d\x45\xe7\xf6\x0a\x75\x76\xbe\x04\x28\xda\xea\x3c\xa2\x62\x04\x05\xba\x23\x0e\x0d\xe8\xe9\x20\x63\xc9\x50\xa8\x47\x1c\x9e\xd1\xc4\x04\x4d\x52\xb5\x56\x4e\x5e\x72\xf4\x22\xa9\xfb\xb8\xdf\xf7\x37\x5e\x3b\x2b\xaa\x7b\xcb\x35\x70\x8d\x70\xfe\xfe\xf2\xac\x90\xaa\x2a\x64\xfb\xb8\x9a\x88\xd5\xff\x61\x88\x7f\x50\x08\xc1\x10\x12\x7e\xef\x9d\xb2\xac\x9f\x07\x51\xd6\x45\xb7\x2b\x0e\x1f\x54\x1a\x6b\x5f\xdd\x08\x10\x5b\x0f\xad\xbc\x46\x77\xdc\x61\xb0\x6f\xf1\x06\xc4\x38\xfb\x9b\x06\xb9\x0e\xe6\x26\x5d\x7f\xba\xba\x16\xa5\x56\xd1\xcc\x60\xce\x7a\xbc\xb1\x8c\x0d\x5d\x9d\x5c\x5f\x25\x1d\xba\x16\xa5\x18\x95\xe6\x63\x9c\x17\x71\x80\x9e\xe9\x55\x72\xfa\x09\x19\x3e\x87\x15\x72\x8d\xf1\xb8\x93\x6a\x42\xe7\xf9\x5a\x94\xaf\xdf\x9c\x9c\x1c\x45\xe9\x2e\x46\xf4\x7a\x73\x2d\xca\x8e\x8e\xd7\xfb\x47\xa7\x3c\x5b\x4e\xb0\xd7\x3d\xe8\xaf\x71\x06\x99\xc1\xf5\x6e\x0e\xb1\xf3\x85\x4c\x23\x9a\x33\x4a\xee\x53\x2a\x31\x2c\x91\xca\x05\x8f\xd0\xff\x96\x90\x2e\xd2\x6e\x98\x7d\xb7\x09\xe6\x6c\x00\x0e\x40\x88\xd0\x84\x88\x60\xbd\xc1\x96\x6b\xe8\x08\x41\x2b\x42\x3a\x3b\x64\xe0\xca\x72\xdd\x2d\x86\x19\x0f\x85\x8e\x61\xb7\x3c\xcc\x82\xc7\x8c\x6d\x83\xd3\xb5\x27\x51\x1d\xb7\x25\x36\x7e\xf7\xa2\x34\x64\xfd\xca\x61\x46\xdd\x62\xf7\xd2\xa9\xfe\xc9\xfa\x1f\x52\xcd\xa2\x64\xca\x9b\x4c\x19\x93\xed\xa7\x7b\x1a\x31\xa4\x54\xde\xbc\x28\x6d\x1a\x54\xd6\x74\x8e\x6d\xeb\x30\x73\x41\x2b\x87\xf3\x66\x3e\x69\x9f\xae\xf7\x11\x7a\x7c\x50\x77\x08\xd4\x45\x4c\x93\xd7\xc1\x2f\xed\x2a\x2d\x14\xdc\xe3\xa2\x0e\xe1\x8e\x00\xbd\x69\x83\xf5\x0c\xcb\x10\x47\x19\x1d\x45\x50\x07\x83\xbb\xf4\xf9\x28\x0e\xf9\x56\xcd\x47\x49\xed\xcf\x9b\x38\xc4\xb5\xe2\x94\x46\xed\x04\x7f\x52\xf4\xfc\xb0\xc4\x42\x26\xc5\xdc\x13\xd7\x0f\x89\x95\x11\x95\xb1\x7e\xf5\xbc\xbc\x3e\xab\x96\x13\xe7\xb6\x1d\x16\x8f\xbf\x6c\xc0\x04\xdd\x25\x8c\x9a\x08\xf3\x64\x1b\x9f\x4b\x34\xf5\x52\x54\xbf\xa2\x46\xbb\xb6\x7e\xb5\xed\xef\x8b\x83\x8e\x2c\xe9\x05\xfd\xb7\x44\xef\xfe\x0e\x4f\x07\x78\x7f\x79\x18\x78\xe8\xe8\xc1\x9b\x6d\xfb\x58\xc8\xe1\x73\xa2\x90\xc3\x87\xe2\xdf\x01\x00\x00\xff\xff\x8c\xe2\x35\xcb\x39\x0a\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/css"].(os.FileInfo),
		fs["/html"].(os.FileInfo),
	}
	fs["/css"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/css/global.css"].(os.FileInfo),
	}
	fs["/html"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/html/cancel.html"].(os.FileInfo),
		fs["/html/redirect.html"].(os.FileInfo),
		fs["/html/success.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
