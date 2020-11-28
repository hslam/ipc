// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package ipc

import (
	"github.com/hslam/msg"
)

// Msgget calls the msgget system call.
func Msgget(key int, msgflg int) (uintptr, error) {
	return msg.Get(key, msgflg)
}

// Msgsnd calls the msgsnd system call.
func Msgsnd(msgid uintptr, msgType uint, msgText []byte, flags uint) error {
	return msg.Snd(msgid, msgType, msgText, flags)
}

// Msgrcv calls the msgrcv system call.
func Msgrcv(msgid uintptr, msgType uint, flags uint) ([]byte, error) {
	return msg.Rcv(msgid, msgType, flags)
}

// Msgrm removes the shm with the given id.
func Msgrm(msgid uintptr) error {
	return msg.Remove(msgid)
}
