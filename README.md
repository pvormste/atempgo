Atempgo
==============

Atempgo is an Automatic TEMplate Parser for GO. 

## Install

```go
go get github.com/pvormste/atempgo
```
 
 
## Features

 * Automatically parses view templates
 * Simulates template inheritance
 * Generates easy keys for accessing the templates

## How does it work?

### Structure
```
views/
├── SplashSite/
│   ├── index.html
│   ├── index-login.html
|   ├── index-login-special.html
|   └── SubSplashFolder/
|       ├── start.html
|       └── start-special.html
|
└── base.html
```

### Note:

  * base.html must be located in the root
  * Inherited templates must use dashes ("-")
  * File extensions  should be all the same (html, tmpl, ...)
  * Child template must be located in same folder as parent

### Config

You can overwrite the default parse options by creating a new ParseOptions struct. Example:

```go
// There also exists a DefaultParseOptions struct (see "As code" section below)
pOpt = &atempgo.ParseOptions{
	BaseName: "layout",	// default: "base"
    Ext: "tmpl",		// default "html"
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

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	atempgo.GetTemplate("index").ExecuteTemplate(w, "base", nil)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	atempgo.GetTemplate("index.login").ExecuteTemplate(w, "base", nil)
}

func HandleSpecialLogin(w http.ResponseWriter, r *http.Request) {
	atempgo.GetTemplate("login.special").ExecuteTemplate(w, "base", nil)
}
```

## Roadmap

  * Code is not optimized. Just was an idea if it works. Maybe someone can rethink about it.