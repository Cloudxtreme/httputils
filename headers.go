// Copyright (c) 2014 Kelsey Hightower. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.
package httputils

import (
	"net/http"
)

func SetLocationHeader(w http.ResponseWriter, url string) {
	w.Header().Set("Location", url)
}
