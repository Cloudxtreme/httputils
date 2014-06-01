// Copyright (c) 2014 Kelsey Hightower. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.
package httputils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// ErrorResponse represents a JSON error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// JSONError replies to the request with the specified error message and HTTP code.
// The response body will contain a JSON representation of the error message.
//
// {
//    "error": "error"
// }
func JSONError(w http.ResponseWriter, error string, code int) error {
	e := ErrorResponse{Error: error}
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
	return nil
}

// JSONWrite replies to the request with the JSON-encoded representation of the
// value pointed to by v and the specified HTTP code.
// The "Content-Type" HTTP header will be set to "application/json".
func JSONWrite(w http.ResponseWriter, v interface{}, code int) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
	return nil
}

// UnmarshalJSONBody reads from r until an error or EOF then parses the
// JSON-encoded data and stores the result int the value pointed to by v.
func UnmarshalJSONBody(r io.ReadCloser, v interface{}) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
