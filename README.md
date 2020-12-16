# ipc
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hslam/ipc)](https://pkg.go.dev/github.com/hslam/ipc)
[![Build Status](https://api.travis-ci.com/hslam/ipc.svg?branch=master)](https://travis-ci.com/hslam/ipc)
[![codecov](https://codecov.io/gh/hslam/ipc/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/ipc)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/ipc)](https://goreportcard.com/report/github.com/hslam/ipc)
[![LICENSE](https://img.shields.io/github/license/hslam/ipc.svg?style=flat-square)](https://github.com/hslam/ipc/blob/master/LICENSE)

Package ipc provides a way to use System V IPC. System V IPC includes three interprocess communication mechanisms that are widely available on UNIX systems: message queues, semaphore, and shared memory.

## Features
* [Ftok](https://github.com/hslam/ftok "ftok")
* [Message queues](https://github.com/hslam/msg "msg")
* [Semaphores](https://github.com/hslam/sem "sem")
* [Shared memory](https://github.com/hslam/shm "shm")

## Get started

### Install
```
go get github.com/hslam/ipc
```
### Import
```
import "github.com/hslam/ipc"
```
### Usage
####  Example
```go
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/hslam/ipc"
	"os"
	"os/signal"
	"syscall"
)

var send = flag.Bool("s", true, "send")

func main() {
	flag.Parse()
	key, _ := ipc.Ftok("/tmp", 0x22)
	semnum := 0
	semid, err := ipc.Semget(key, 1, 0666)
	if err != nil {
		semid, err = ipc.Semget(key, 1, ipc.IPC_CREAT|ipc.IPC_EXCL|0666)
		if err != nil {
			panic(err)
		}
		_, err := ipc.Semsetvalue(semid, semnum, 1)
		if err != nil {
			panic(err)
		}
	}
	defer ipc.Semrm(semid)
	shmid, data, _ := ipc.Shmgetattach(key, 128, ipc.IPC_CREAT|0600)
	defer ipc.Shmrm(shmid)
	defer ipc.Shmdetach(data)
	msgid, _ := ipc.Msgget(key, ipc.IPC_CREAT|0600)
	defer ipc.Msgrm(msgid)
	var text string
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer close(quit)
		if *send {
			fmt.Println("Enter:")
			buf := make([]byte, 10)
			for {
				fmt.Scanln(&text)
				if _, err := ipc.Semp(semid, semnum, ipc.SEM_UNDO); err != nil {
					return
				}
				copy(data, text)
				if _, err := ipc.Semv(semid, semnum, ipc.SEM_UNDO); err != nil {
					return
				}
				n := binary.PutUvarint(buf, uint64(len(text)))
				if err := ipc.Msgsend(msgid, 1, buf[:n], 0600); err != nil {
					return
				}
			}
		} else {
			fmt.Println("Recv:")
			for {
				m, err := ipc.Msgreceive(msgid, 1, 0600)
				if err != nil {
					return
				}
				length, _ := binary.Uvarint(m)
				if _, err := ipc.Semp(semid, semnum, ipc.SEM_UNDO); err != nil {
					return
				}
				text = string(data[:length])
				if _, err := ipc.Semv(semid, semnum, ipc.SEM_UNDO); err != nil {
					return
				}
				fmt.Println(text)
			}
		}
	}()
	<-quit
}
```

#### Output
Enter a word.
```sh
$ go run main.go -s=true
Enter:
HelloWorld
```
In another terminal receive this word.
```sh
$ go run main.go -s=false
Recv:
HelloWorld
```

### License
This package is licensed under a MIT license (Copyright (c) 2020 Meng Huang)


### Author
ipc was written by Meng Huang.


