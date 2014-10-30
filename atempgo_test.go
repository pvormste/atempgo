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

var tests = []Test{
	// Default Configuration
	{"#test", "base", "<h1>Base.html</h1><h2>Test.html</h2>"},
	{"#index.login", "base", "<h1>Base.html</h1><h2>Index.html</h2><h3>Index-login.html</h3>"},
	{"#index.register", "base", "<h1>Base.html</h1><h2>Index.html</h2><h3>Index-register.html</h3>"},
	{"#start", "base", "<h1>Base.html</h1><h2>Start.html</h2>"},
	{"superspecial.content", "superspecial", "<h1>Superspecial.html</h1><h2>Superspecial-content.html</h2>"},
}

func TestAtempgo(t *testing.T) {
	LoadTemplates("test_views", DefaultParseOptions)
	for _, test := range tests {
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
