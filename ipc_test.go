// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package ipc

import (
	"github.com/hslam/sem"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func TestIPC(t *testing.T) {
	context := strings.Repeat("1", 64)
	done := make(chan struct{})
	semnum := 0
	go func() {
		key, _ := Ftok("/tmp", 0x22)
		semid, err := Semget(key, 1, 0600)
		if err != nil {
			semid, err = Semget(key, 1, IPC_CREAT|IPC_EXCL|0600)
			if err != nil {
				panic(err)
			}
			_, err := Semsetvalue(semid, semnum, 1)
			if err != nil {
				panic(err)
			}
		}
		defer Semrm(semid)
		shmid, data, _ := Shmgetattach(key, 128, IPC_CREAT|0600)
		defer Shmrm(shmid)
		defer Shmdetach(data)
		msgid, _ := Msgget(key, IPC_CREAT|0600)
		defer Msgrm(msgid)
		if _, err := Semp(semid, semnum, SEM_UNDO); err != nil {
			return
		}
		copy(data, context)
		if _, err := Semv(semid, semnum, SEM_UNDO); err != nil {
			return
		}
		if err := Msgsend(msgid, 1, []byte{byte(len(context))}, 0600); err != nil {
			return
		}

		time.Sleep(time.Millisecond * 200)
		close(done)
	}()
	time.Sleep(time.Millisecond * 100)

	key, _ := Ftok("/tmp", 0x22)
	semid, err := Semget(key, 1, 0600)
	if err != nil {
		semid, err = Semget(key, 1, IPC_CREAT|IPC_EXCL|0600)
		if err != nil {
			panic(err)
		}
		_, err := Semsetvalue(semid, semnum, 1)
		if err != nil {
			panic(err)
		}
	}
	defer Semrm(semid)
	shmid, data, _ := Shmgetattach(key, 128, 0600)
	defer Shmrm(shmid)
	defer Shmdetach(data)
	msgid, _ := Msgget(key, 0600)
	defer Msgrm(msgid)

	m, err := Msgreceive(msgid, 1, 0600)
	if err != nil {
		return
	}
	length := int(m[0])
	if ret, err := Semgetvalue(semid, semnum); err != nil {
		t.Error(err)
	} else if ret < 1 {
		t.Error()
	}
	if _, err := Semp(semid, semnum, SEM_UNDO); err != nil {
		return
	}
	text := string(data[:length])
	if _, err := Semv(semid, semnum, SEM_UNDO); err != nil {
		return
	}

	if context != string(text) {
		t.Error(context, string(text))
	}
	<-done
}

func TestMore(t *testing.T) {
	context := strings.Repeat("1", 64)
	done := make(chan struct{})
	semnum := 0
	msgType := uint(1)
	go func() {
		key, _ := Ftok("/tmp", 0x22)
		semid, err := Semget(key, 1, 0600)
		if err != nil {
			semid, err = Semget(key, 1, IPC_CREAT|IPC_EXCL|0600)
			if err != nil {
				panic(err)
			}
			_, err := Semsetvalue(semid, semnum, 1)
			if err != nil {
				panic(err)
			}
		}
		defer Semrm(semid)
		shmid, data, _ := Shmgetattach(key, 128, IPC_CREAT|0600)
		defer Shmrm(shmid)
		defer Shmdetach(data)
		msgid, _ := Msgget(key, IPC_CREAT|0600)
		defer Msgrm(msgid)
		if _, err := Semp(semid, semnum, SEM_UNDO); err != nil {
			return
		}
		copy(data, context)
		if _, err := Semv(semid, semnum, SEM_UNDO); err != nil {
			return
		}
		st := struct {
			Type uint
			Text [8192]byte
		}{}
		st.Type = msgType
		n := copy(st.Text[:], []byte{byte(len(context))})
		if err := Msgsnd(msgid, uintptr(unsafe.Pointer(&st)), n, 0600); err != nil {
			return
		}
		time.Sleep(time.Millisecond * 200)
		close(done)
	}()
	time.Sleep(time.Millisecond * 100)

	key, _ := Ftok("/tmp", 0x22)
	semid, err := Semget(key, 1, 0600)
	if err != nil {
		semid, err = Semget(key, 1, IPC_CREAT|IPC_EXCL|0600)
		if err != nil {
			panic(err)
		}
		_, err := Semsetvalue(semid, semnum, 1)
		if err != nil {
			panic(err)
		}
	}
	defer Semrm(semid)
	size := 128
	shmid, _ := Shmget(key, size, 0600)
	defer Shmrm(shmid)
	shmaddr, _ := Shmat(shmid, 0600)
	var sl = struct {
		addr uintptr
		len  int
		cap  int
	}{shmaddr, size, size}
	data := *(*[]byte)(unsafe.Pointer(&sl))
	defer Shmdt(shmaddr)
	msgid, _ := Msgget(key, 0600)
	defer Msgrm(msgid)
	st := struct {
		Type uint
		Text [8192]byte
	}{}
	st.Type = msgType
	sz, err := Msgrcv(msgid, uintptr(unsafe.Pointer(&st)), 8192, msgType, 0600)
	if err != nil {
		return
	}
	if sz < 1 {
		t.Error()
	}
	length := int(st.Text[0])
	if ret, err := Semgetvalue(semid, semnum); err != nil {
		t.Error(err)
	} else if ret < 1 {
		t.Error()
	}
	var sops [1]sem.Sembuf
	sops[0] = sem.Sembuf{SemNum: uint16(semnum), SemFlg: SEM_UNDO}
	sops[0].SemOp = -1
	if _, err := Semop(semid, sops[:]); err != nil {
		return
	}
	text := string(data[:length])
	sops[0].SemOp = 1
	if _, err := Semop(semid, sops[:]); err != nil {
		return
	}

	if context != string(text) {
		t.Error(context, string(text))
	}
	<-done
}
