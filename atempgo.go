// Copyright 2014 Patric "pvormste" Vormstein.
// All rights reserved.
package atempgo

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

// This map contains all templates
var templates map[string]*template.Template

// ## Private

// This function checks the dir for files and subfolders.
// It will create inherited templates based on its naming.
// Every template inherits from a base template.
// So: index.html will inherit from base.html
// index-login.html inherits from index.html and so on.
// The templates can be called by templates["index"] or templates["index->login"]
func checkDir(dir string, relPathViews string, baseViewName string, ext string) {
	var files []os.FileInfo
	var err error

	// Read all files and dirs from the actual folder
	if dir != "" {
		files, err = ioutil.ReadDir(relPathViews + "/" + dir)
	} else {
		files, err = ioutil.ReadDir(relPathViews)
	}

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
			partialTmpl := strings.Split(filename, "-")
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
						parent += partialTmpl[j] + "-"
					}
				}

				rebuildTmpl[i] = createPathToView(dir, relPathViews, parent, true, ext)
			}

			// Saving the templates like e.g.:
			// index.html inherits from base.html: templates["index"]
			// index-login.html inherits from index.html ...: templates["index->login"]
			// index-login-special.html inherits from index-login.html ...: templates["login->special"]
			templates[partialTmpl[len(partialTmpl)-2]+"->"+partialTmpl[len(partialTmpl)-1]] = createInheritedTemplate(createPathToView("", relPathViews, baseViewName, true, ext), rebuildTmpl...)
		} else if !file.IsDir() {
			// Add template with inheritance
			if filename != baseViewName {
				templates[filename] = createInheritedTemplate(createPathToView("", relPathViews, baseViewName, true, ext), createPathToView(dir, relPathViews, file.Name(), false, ""))
			}

		} else {
			// Check subfolder the same way.
			// NOTE: At the moment, there is only one subfolder supported
			checkDir(file.Name(), relPathViews, baseViewName, ext)
		}

	}
}

// This function parses the views and save them as inherited
func createInheritedTemplate(base string, children ...string) *template.Template {
	temp := make([]string, len(children)+1)
	temp[0] = base

	for i, child := range children {
		temp[i+1] = child
	}

	return template.Must(template.ParseFiles(temp...))
}

// Creates full paths to the views
// NOTE: Only supports one subfolder depth of one
func createPathToView(dir string, relPathViews string, filename string, withExt bool, ext string) string {
	var fullpath string

	if dir != "" {
		fullpath = relPathViews + "/" + dir + "/" + filename
	} else {
		relPathViews + "/" + filename
	}

	if withExt {
		fullpath += "." + ext
	}

	return fullpath
}

// ## Public

// This function checks the view directory and parses the templates.
// relPathViews: Relative path from exectuable to the view directory
// baseViewName: How the base view file is named. Typically "base".
// ext: Extension naming of the views. Typically "html" or "tmpl"
func LoadTemplates(relPathViews string, baseViewName string, ext string) {
	// Initializes the template map
	templates = make(map[string]*template.Template)

	// Start checking the main dir of the views
	checkDir("", relPathViews, baseViewName, ext)
}

// This function returns the actual template with key "name".
// Naming is 'templates["singleInherited"]' or 'templates["multi->inherited"]'
func GetTemplate(name string) *template.Template {
	return templates[name]
}
