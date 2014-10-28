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
|   └── index-login-special.html
|
└── base.html
```

### Note:

  * base.html must be located in the root
  * Inherited templates must use dashes ("-")
  * File extensions  should be all the same (html, tmpl, ...)
  * **SPECIAL NOTE:** It only supports a subfolder depth of one at the moment
  
### As code

```go
import (
	"github.com/pvormste/atempgo"
)

func init() {
	// Parameters: Path to base file, Name of base file, template extensions
	atempgo.LoadTemplates("path/to/folde/with/basefile/", "base", "tmpl")
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
	atempgo.GetTemplate("index->login").ExecuteTemplate(w, "base", nil)
}

func HandleSpecialLogin(w http.ResponseWriter, r *http.Request) {
	atempgo.GetTemplate("login->special").ExecuteTemplate(w, "base", nil)
}
```

## Roadmap

  * Code is not optimized. Just was an idea if it works. Maybe someone can rethink about it.
  * Support of "unlimited" subfolder depth