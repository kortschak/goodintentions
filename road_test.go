// Copyright ©2022 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goodintentions

import (
	"bytes"
	"math"
	"testing"
)

func TestGoodIntentionsBytesString(t *testing.T) {
	want := "This won't end well."
	got := Convert[[]byte, string]([]byte(want))
	if got != want {
		t.Errorf("unexpected result: got: %v want: %v", got, want)
	}
}

func TestGoodIntentionsUint64Float64(t *testing.T) {
	want := uint64(42)
	got := Convert[uint64, float64](uint64(want))
	if got != math.Float64frombits(want) {
		t.Errorf("unexpected result: got: %v want: %v", got, want)
	}
}

func TestGoodIntentionsFauxSlice(t *testing.T) {
	type fauxSlice struct {
		data *[10]byte
		len  int
		cap  int
	}
	data := [10]byte{'p', 'a', 'v', 'i', 'n', 'g'}
	want := fauxSlice{
		data: &data,
		len:  6,
		cap:  len(data),
	}
	got := Convert[fauxSlice, []byte](want)
	if !bytes.Equal(got, data[:want.len]) {
		t.Errorf("unexpected result: got: %v want: %v", got, data[:want.len])
	}
	if len(got) != want.len {
		t.Errorf("unexpected len result: got: %v want: %v", len(got), want.len)
	}
	if cap(got) != want.cap {
		t.Errorf("unexpected cap result: got: %v want: %v", cap(got), want.cap)
	}
}
