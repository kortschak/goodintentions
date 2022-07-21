// Copyright ©2022 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goodintentions provides a well intended function for
// converting arbitrary types.
package goodintentions

import "unsafe"

// Convert converts the value from the From type to the To type with the best
// of intentions.
func Convert[From any, To any](value From) To { return *((*To)(unsafe.Pointer(&value))) }
