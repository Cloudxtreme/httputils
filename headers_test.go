// Copyright (c) 2014 Kelsey Hightower. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.
package httputils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var locationTests = []struct {
	code     int
	location string
}{
	{200, "http://example.com/somepath"},
}

func TestSetLocationHeader(t *testing.T) {
	for _, tt := range locationTests {
		handler := func(w http.ResponseWriter, r *http.Request) {
			SetLocationHeader(w, tt.location)
		}
		req, err := http.NewRequest("GET", "http://example.com", nil)
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
		w := httptest.NewRecorder()
		handler(w, req)
		if w.Code != tt.code {
			t.Errorf("StatusCode => %d, want %d", w.Code, tt.code)
		}
		if w.HeaderMap["Location"][0] != tt.location {
			t.Errorf("Header['Location'] => %s, want %s", w.HeaderMap["Location"], tt.location)
		}
	}
}
