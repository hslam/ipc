// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package ipc

import (
	"github.com/hslam/ftok"
)

// Ftok uses the given pathname (which must refer to an existing, accessible file) and
// the least significant 8 bits of proj_id (which must be nonzero) to generate
// a key_t type System V IPC key.
func Ftok(pathname string, projectid uint8) (int, error) {
	return ftok.Ftok(pathname, projectid)
}
