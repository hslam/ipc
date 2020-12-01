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

// Shmat calls the shmat system call.
func Shmat(shmid int, shmFlg int) (uintptr, error) {
	return shm.At(shmid, shmFlg)
}

// Shmdt calls the shmdt system call.
func Shmdt(addr uintptr) error {
	return shm.Dt(addr)
}

// Shmgetattach calls the shmget and shmat system call.
func Shmgetattach(key int, size int, shmFlg int) (int, []byte, error) {
	return shm.GetAttach(key, size, shmFlg)
}

// Shmdetach calls the shmdt system call with []byte b.
func Shmdetach(b []byte) error {
	return shm.Detach(b)
}

// Shmrm removes the shm with the given id.
func Shmrm(shmid int) error {
	return shm.Remove(shmid)
}
