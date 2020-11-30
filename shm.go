// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package ipc

import (
	"github.com/hslam/shm"
)

// Shmget calls the shmget system call.
func Shmget(key int, size int, shmFlg int) (int, error) {
	return shm.Get(key, size, shmFlg)
}

// Shmattach calls the shmat system call.
func Shmattach(shmid int, shmFlg int) (uintptr, error) {
	return shm.Attach(shmid, shmFlg)
}

// Shmdetach calls the shmdt system call.
func Shmdetach(addr uintptr) error {
	return shm.Detach(addr)
}

// Shmgetat calls the shmget and shmat system call.
func Shmgetat(key int, size int, shmFlg int) (int, []byte, error) {
	return shm.GetAt(key, size, shmFlg)
}

// Shmdt calls the shmdt system call with []byte b.
func Shmdt(b []byte) error {
	return shm.Dt(b)
}

// Shmrm removes the shm with the given id.
func Shmrm(shmid int) error {
	return shm.Remove(shmid)
}
