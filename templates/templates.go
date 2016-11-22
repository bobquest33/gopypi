// Code generated by go-bindata.
// sources:
// index.tpl.html
// package_list.tpl.html
// DO NOT EDIT!

package templates

import (
	"github.com/elazarl/go-bindata-assetfs"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
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

var _indexTplHtml = []byte(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="gopypi - pypi server implemented in golang">
    <meta name="author" content="phonkee">

    <title>Gopypi server</title>

    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet"
          integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

    <style type="text/css">
        html {
            position: relative;
            min-height: 100%;
        }
        body {
            padding-top: 20px;
            padding-bottom: 20px;
            margin-bottom: 60px;
        }
        .header {
            padding-bottom: 20px;
            border-bottom: 1px solid #e5e5e5;
        }
        .header h3 {
            margin-top: 0;
            margin-bottom: 0;
            line-height: 40px;
        }
        .footer {
            color: #777;
            position: absolute;
            bottom: 0;
            height: 60px;
            width: 80%;
        }
    </style>
</head>

<body>
<div class="container">
    <div class="header clearfix">
        <nav class="navbar navbar-default">
            <div class="container-fluid">
                <!-- Brand and toggle get grouped for better mobile display -->
                <div class="navbar-header">
                    <a class="navbar-brand" href="https://www.github.com/phonkee/gopypi" target="_blank">Gopypi
                        <small>{{.version}}</small>
                    </a>
                </div>

                <!-- Collect the nav links, forms, and other content for toggling -->
                <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                    <ul class="nav navbar-nav navbar-right">
                        <li class="btn-xs"><a href="/admin/dashboard">Admin</a></li>
                    </ul>
                </div><!-- /.navbar-collapse -->
            </div><!-- /.container-fluid -->
        </nav>
    </div>
    <div class="jumbotron">
        <h1>Welcome!</h1>
        <p class="lead">
            Gopypi is pypi server implementation written in golang. Gopypi is flexible, safe, lightweight and blazing
            fast!
        </p>
    </div>
    <footer class="footer pull-right">
        <p>&copy; 2016 <a href="https://www.github.com/phonkee" target="_blank">phonkee</a></p>
    </footer>
</div>
</body>
</html>
`)

func indexTplHtmlBytes() ([]byte, error) {
	return _indexTplHtml, nil
}

func indexTplHtml() (*asset, error) {
	bytes, err := indexTplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.tpl.html", size: 2594, mode: os.FileMode(420), modTime: time.Unix(1479511183, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _package_listTplHtml = []byte(`<html>
    <head>
        <title>Gopypi simple Index</title>
    </head>
    <body>
        <h1>Simple Index</h1>
        {{range .Packages}}
            <a href="{{reverse "package_detail" "slug" .Name}}">{{.Name}}</a><br>
        {{end}}
    </body>
</html>`)

func package_listTplHtmlBytes() ([]byte, error) {
	return _package_listTplHtml, nil
}

func package_listTplHtml() (*asset, error) {
	bytes, err := package_listTplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "package_list.tpl.html", size: 259, mode: os.FileMode(420), modTime: time.Unix(1476313556, 0)}
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
	"index.tpl.html": indexTplHtml,
	"package_list.tpl.html": package_listTplHtml,
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
	"index.tpl.html": &bintree{indexTplHtml, map[string]*bintree{}},
	"package_list.tpl.html": &bintree{package_listTplHtml, map[string]*bintree{}},
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
