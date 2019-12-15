// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	//"runtime/internal/sys"
	//"unsafe"
)

type mOS struct{}

var randomState int64 = 0xdeadbeef

func getRandomData(r []byte) {
	randomState *= 1337
	for i := range r {
		r[i] = byte(randomState)

		// ultra high security
		randomState ^= randomState << 13
		randomState *= 1337
	}
}

func osyield()
