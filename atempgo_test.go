// Copyright 2014 Patric "pvormste" Vormstein.
// All rights reserved.
package atempgo

import (
	"bytes"
	"testing"
)

type Test struct {
	key      string
	tmpl     string
	expected string
}

var testsDefault = []Test{
	// Default Configuration
	{"#test", "base", "<h1>Base.html</h1><h2>Test.html</h2>"},
	{"#index.login", "base", "<h1>Base.html</h1><h2>Index.html</h2><h3>Index-login.html</h3>"},
	{"#index.register", "base", "<h1>Base.html</h1><h2>Index.html</h2><h3>Index-register.html</h3>"},
	{"#start", "base", "<h1>Base.html</h1><h2>Start.html</h2>"},
	{"superspecial.content", "superspecial", "<h1>Superspecial.html</h1><h2>Superspecial-content.html</h2>"},
}

var testsNonDefault = []Test{
	// Non-Default Configuration
	{"#test", "layout", "<h1>Layout.tmpl</h1><h2>Test.tmpl</h2>"},
	{"#index.login", "layout", "<h1>Layout.tmpl</h1><h2>Index.tmpl</h2><h3>Index+login.tmpl</h3>"},
	{"#index.register", "layout", "<h1>Layout.tmpl</h1><h2>Index.tmpl</h2><h3>Index+register.tmpl</h3>"},
	{"#start", "layout", "<h1>Layout.tmpl</h1><h2>Start.tmpl</h2>"},
	{"superspecial.content", "superspecial", "<h1>Superspecial.tmpl</h1><h2>Superspecial+content.html</h2>"},
}

func TestAtempgoDefault(t *testing.T) {
	LoadTemplates("test_views_default", DefaultParseOptions)
	for _, test := range testsDefault {
		var output bytes.Buffer
		err := GetTemplate(test.key).ExecuteTemplate(&output, test.tmpl, nil)

		if err != nil {
			t.Fatalf("%s failed while ExecuteTemplate", test.key)
		}

		str := output.String()

		if str != test.expected {
			t.Fatalf("%s failed. Expected: %s | Output: %s", test.key, test.expected, str)
		}
	}
}

func TestAtempgoNonDefault(t *testing.T) {
	pOptions := &ParseOptions{
		BaseName:      "layout",
		Delimiter:     "+",
		Ext:           "tmpl",
		NonBaseFolder: "single",
	}

	LoadTemplates("test_views_nondefault", pOptions)
	for _, test := range testsNonDefault {
		var output bytes.Buffer
		err := GetTemplate(test.key).ExecuteTemplate(&output, test.tmpl, nil)

		if err != nil {
			t.Fatalf("%s failed while ExecuteTemplate", test.key)
		}

		str := output.String()

		if str != test.expected {
			t.Fatalf("%s failed. Expected: %s | Output: %s", test.key, test.expected, str)
		}
	}
}
