// Copyright (c) 2014 Kelsey Hightower. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.
package httputils

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

var errorResponseFormat = `{
  "error": "%s"
}`

func makeBody(error string) string {
	return fmt.Sprintf(errorResponseFormat, error)	
}

var jsonErrorTests = []struct{
	body   string
	code   int
	error  string
	header string
}{
	{makeBody("invalid request"), 500, "invalid request", "application/json"},
}

func TestJSONError(t *testing.T) {
	for _, tt := range jsonErrorTests {
		w := httptest.NewRecorder()
		err := JSONError(w, tt.error, tt.code)
		if err != nil {
			t.Errorf("unexpected error %s", err.Error())
		}
		results := map[string]struct{
			body   string
			code   int
			header string
		}{
			"want": {w.Body.String(), w.Code, w.HeaderMap["Content-Type"][0]},
			"got": {tt.body, tt.code, tt.header},
		} 
		if results["got"] != results["want"] {
			t.Errorf("want %v, got %v", results["want"], results["got"])
		}
	}
}
