Atempgo
==============

Atempgo is an Automatic TEMplate Parser for GO. 

## Install

```go
go get github.com/pvormste/atempgo
```
 
 
## Features

 * Automatically parses view templates
 * Simulates template inheritance (as described at [StackOverflow](http://stackoverflow.com/questions/11467731/is-it-possible-to-have-nested-templates-in-go-using-the-standard-library-googl))
 * Generates easy keys for accessing the templates

## How does it work?

### Structure (example)
```
views/
├── SplashSite/
│   ├── index.html
│   ├── index-login.html
|   ├── index-register.html
|   └── SubSplashFolder/
|       └── start.html
|
├── nonbase/
│   ├── superspecial.html
|   └── superspecial-content.html
|
└── base.html
```

### Note:

  * base.html must be located in the root
  * Views, which doesn't extend a base-view, must be located in specific folder (See ParseOptions)
  * Inherited templates must use a delimiter specified in the ParseOptions (default "-")
  * File extensions  should be all the same (html, tmpl, ...)
  * Child template must be located in same folder as parent
  * Hash (#) is short form of: extends base -> base.index.login is written as #index.login
  * If key starts with Hash (#), you have to pass the BaseName value (default: base), else: the first parent
  * Because of go template design, you can only call the last child, not a single parent (e.g. "#index.login" not "#index")

### Config

You can overwrite the default parse options by creating a new ParseOptions struct. Example:

```go
// There also exists a DefaultParseOptions struct (see "As code" section below)
pOpt := &atempgo.ParseOptions{
	BaseName: "layout",			// default: "base"
	Delimiter: "+",				// default: "-"
    Ext: "tmpl",				// default: "html"
    NonBaseFolder: "single",	// default: "nonbase"
}
```
  
### As code

```go
import (
	"github.com/pvormste/atempgo"
    "path/filepath"
)

func init() {
	// Parameters: Path to base file (e.g. path/to/folder/with/basefile), parse options (e.g. DefaultParseOptions)
	atempgo.LoadTemplates(filepath.Join("path", "to", "folder", "with", "basefile"), atempgo.DefaultParseOptions)
}
```

```go
import (
	"net/http"
	"github.com/pvormste/atempgo"
)

// The Hash (#) means, that the view extends base view.
// So "base" must be executed. 
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	atempgo.GetTemplate("#index.login").ExecuteTemplate(w, "base", nil)
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	atempgo.GetTemplate("#index.register").ExecuteTemplate(w, "base", nil)
}

// No Hash (#) indicates, that it doesn't extends the base view.
// In this case, it must render the first named part of the key ("superspecial").
func HandleSuperspecialWithContent(w http.ResponseWriter, r *http.Request) {
	atempgo.GetTemplate("superspecial.content").ExecuteTemplate(w, "superspecial", nil)
}
```

## Roadmap

  * Code is not optimized. Just was an idea if it works. Maybe someone can rethink about it.
  * Add support to call views by path

## LICENSE (MIT)

The MIT License (MIT)

Copyright (c) 2014 Patric "pvormste" Vormstein

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.