// Copyright 2014 Patric "pvormste" Vormstein.
// All rights reserved.
package atempgo

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ParseOptions struct {
	BasePath      string
	BaseName      string
	Delimiter     string
	Ext           string
	NonBaseFolder string
}

func (pOpt *ParseOptions) getPathToBase() string {
	return filepath.Join(pOpt.BasePath, pOpt.BaseName) + "." + pOpt.Ext
}

// This map contains all templates
var templates map[string]*template.Template

// ## Private

// This function checks the dir for files and subfolders.
// It will create inherited templates based on its naming.
// Every template inherits from a base template.
// So: index.html will inherit from base.html
// index-login.html inherits from index.html and so on.
// The templates can be called by templates["#index"] or templates["#index.login"]
func checkDir(relativePath string, pOpt *ParseOptions, isNonBase bool) {
	var files []os.FileInfo
	var err error

	files, err = ioutil.ReadDir(relativePath)

	// If there is an error: PANIC!
	if err != nil {
		panic(err)
	}

	// Iterate all found files and folders
	for _, file := range files {
		// Extract the filename without extension
		filename := strings.Split(file.Name(), ".")[0]

		// Check for inheritance marked by "-"
		if strings.Contains(filename, "-") && !file.IsDir() {
			// Split the filename
			partialTmpl := strings.Split(filename, pOpt.Delimiter)
			rebuildTmpl := make([]string, len(partialTmpl))

			// Rename the strings correctly
			// Like: views/index.html && views/index-login.html
			// So index-login.html inherits from index.html
			// and index.html inherits from base.html
			for i, _ := range partialTmpl {
				var parent string
				for j := 0; j <= i; j++ {
					if j == i {
						parent += partialTmpl[j]
					} else {
						parent += partialTmpl[j] + pOpt.Delimiter
					}
				}

				rebuildTmpl[i] = createPathToView(relativePath, parent, true, pOpt)
			}

			// Saving the templates like e.g.:
			// index.html inherits from base.html: templates["#index"]
			// index-login.html inherits from index.html ...: templates["#index.login"]
			// index-login-special.html inherits from index-login.html ...: templates["#login.special"]
			// In NonBase, the '#' dissapears (doesn't extends base), like: templates["superspecial"] for views/nonbase/superspecial.html
			if !isNonBase {
				templates["#"+partialTmpl[len(partialTmpl)-2]+"."+partialTmpl[len(partialTmpl)-1]] = createInheritedTemplate(pOpt, true, rebuildTmpl...)
			} else if isNonBase && len(partialTmpl) == 2 {
				templates[partialTmpl[0]+"."+partialTmpl[1]] = createInheritedTemplate(pOpt, false, rebuildTmpl...)
			} else if isNonBase && len(partialTmpl) > 2 {
				templates[partialTmpl[0]+"."+partialTmpl[len(partialTmpl)-2]+"."+partialTmpl[len(partialTmpl)-1]] = createInheritedTemplate(pOpt, false, rebuildTmpl...)
			}

		} else if !file.IsDir() {
			// Add template with inheritance
			if filename != pOpt.BaseName {
				if !isNonBase {
					templates["#"+filename] = createInheritedTemplate(pOpt, true, createPathToView(relativePath, file.Name(), false, pOpt))
				} else {
					templates[filename] = createInheritedTemplate(pOpt, false, createPathToView(relativePath, file.Name(), false, pOpt))
				}
			}

		} else if file.IsDir() && file.Name() == pOpt.NonBaseFolder {
			// Check if entering NonBase Folder
			checkDir(filepath.Join(relativePath, file.Name()), pOpt, true)

		} else {
			// Check subfolder the same way.
			checkDir(filepath.Join(relativePath, file.Name()), pOpt, isNonBase)
		}

	}
}

// This function parses the views and save them as inherited
func createInheritedTemplate(pOpt *ParseOptions, useBase bool, children ...string) *template.Template {
	if useBase {
		temp := make([]string, len(children)+1)
		temp[0] = pOpt.getPathToBase()

		for i, child := range children {
			temp[i+1] = child
		}

		return template.Must(template.ParseFiles(temp...))
	}

	return template.Must(template.ParseFiles(children...))
}

// Creates full paths to the views
func createPathToView(relativePath string, filename string, withExt bool, pOpt *ParseOptions) string {
	var fullpath string

	fullpath = filepath.Join(relativePath, filename)

	if withExt {
		fullpath += "." + pOpt.Ext
	}

	return fullpath
}

// ## Public

// Default Options
var DefaultParseOptions = &ParseOptions{BaseName: "base", Delimiter: "-", Ext: "html", NonBaseFolder: "nonbase"}

// This function checks the view directory and parses the templates.
// relativePath: Relative path from exectuable to the view directory
// baseViewName: How the base view file is named. Typically "base".
// ext: Extension naming of the views. Typically "html" or "tmpl"
func LoadTemplates(relativePath string, pOpt *ParseOptions) {
	// Initializes the template map
	templates = make(map[string]*template.Template)

	// Save Path to Base file
	pOpt.BasePath = relativePath

	// Check if every option is set
	if pOpt.BaseName == "" {
		pOpt.BaseName = DefaultParseOptions.BaseName
	}

	if pOpt.Delimiter == "" {
		pOpt.Delimiter = DefaultParseOptions.Delimiter
	}

	if pOpt.Ext == "" {
		pOpt.Ext = DefaultParseOptions.Ext
	}

	if pOpt.NonBaseFolder == "" {
		pOpt.NonBaseFolder = DefaultParseOptions.NonBaseFolder
	}

	// Start checking the main dir of the views
	checkDir(relativePath, pOpt, false)
}

// This function returns the actual template with key "name".
// Naming is 'templates["#singleInherited"]' or 'templates["#multi.inherited"]'
func GetTemplate(name string) *template.Template {
	return templates[name]
}
