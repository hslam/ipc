// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package ipc

import (
	"github.com/hslam/shm"
)

// Shmgetat calls the shmget and shmat system call.
func Shmgetat(key int, size int, shmFlg int) (uintptr, []byte, error) {
	return shm.GetAt(key, size, shmFlg)
}

// Shmdt calls the shmdt system call.
func Shmdt(b []byte) error {
	return shm.Dt(b)
}

// Shmrm removes the shm with the given id.
func Shmrm(shmid uintptr) error {
	return shm.Remove(shmid)
}
