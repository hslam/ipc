// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package ipc

import (
	"github.com/hslam/sem"
)

// Semget calls the semget system call.
func Semget(key int) (uintptr, error) {
	return sem.Get(key)
}

// Semp calls the semop P system call.
func Semp(semid uintptr, flg int16) (bool, error) {
	return sem.P(semid, flg)
}

// Semv calls the semop V system call.
func Semv(semid uintptr, flg int16) (bool, error) {
	return sem.V(semid, flg)
}

// Semgetvalue calls the semctl GETVAL system call.
func Semgetvalue(semid uintptr) (int, error) {
	return sem.GetValue(semid)
}

// Semrm removes the semaphore with the given id.
func Semrm(semid uintptr) error {
	return sem.Remove(semid)
}
