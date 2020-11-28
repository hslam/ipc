// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package ipc

import (
	"strings"
	"testing"
	"time"
)

func TestIPC(t *testing.T) {
	context := strings.Repeat("1", 64)
	done := make(chan struct{})
	go func() {
		key, _ := Ftok("/tmp", 0x22)
		semid, _ := Semget(key)
		defer Semrm(semid)
		shmid, data, _ := Shmgetat(key, 128, IPC_CREAT|0600)
		defer Shmrm(shmid)
		defer Shmdt(data)
		msgid, _ := Msgget(key, IPC_CREAT|0600)
		defer Msgrm(msgid)

		if _, err := Semp(semid, SEM_UNDO); err != nil {
			return
		}
		copy(data, context)
		if _, err := Semv(semid, SEM_UNDO); err != nil {
			return
		}
		if err := Msgsnd(msgid, 1, []byte{byte(len(context))}, 0600); err != nil {
			return
		}

		time.Sleep(time.Millisecond * 200)
		close(done)
	}()
	time.Sleep(time.Millisecond * 100)

	key, _ := Ftok("/tmp", 0x22)
	semid, _ := Semget(key)
	defer Semrm(semid)
	shmid, data, _ := Shmgetat(key, 128, IPC_CREAT|0600)
	defer Shmrm(shmid)
	defer Shmdt(data)
	msgid, _ := Msgget(key, IPC_CREAT|0600)
	defer Msgrm(msgid)

	m, err := Msgrcv(msgid, 1, 0600)
	if err != nil {
		return
	}
	length := int(m[0])
	if ret, err := Semgetvalue(semid); err != nil {
		t.Error(err)
	} else if ret < 1 {
		t.Error()
	}
	if _, err := Semp(semid, SEM_UNDO); err != nil {
		return
	}
	text := string(data[:length])
	if _, err := Semv(semid, SEM_UNDO); err != nil {
		return
	}

	if context != string(text) {
		t.Error(context, string(text))
	}
	<-done
}
