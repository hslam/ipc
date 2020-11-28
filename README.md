# ipc
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hslam/ipc)](https://pkg.go.dev/github.com/hslam/ipc)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/ipc)](https://goreportcard.com/report/github.com/hslam/ipc)
[![LICENSE](https://img.shields.io/github/license/hslam/ipc.svg?style=flat-square)](https://github.com/hslam/ipc/blob/master/LICENSE)

Package ipc provides a way to use System V IPC. System V IPC includes three interprocess communication mechanisms that are widely available on UNIX systems: message queues, semaphore, and shared memory.

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
	key, _ := ipc.Ftok("/tmp", 0x22)
	semid, _ := ipc.Semget(key)
	defer ipc.Semrm(semid)
	shmid, data, _ := ipc.Shmgetat(key, 128, ipc.IPC_CREAT|0600)
	defer ipc.Shmrm(shmid)
	defer ipc.Shmdt(data)
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
				if _, err := ipc.Semp(semid, ipc.SEM_UNDO); err != nil {
					return
				}
				copy(data, text)
				if _, err := ipc.Semv(semid, ipc.SEM_UNDO); err != nil {
					return
				}
				n := binary.PutUvarint(buf, uint64(len(text)))
				if err := ipc.Msgsnd(msgid, 1, buf[:n], 0600); err != nil {
					return
				}
			}
		} else {
			fmt.Println("Recv:")
			for {
				m, err := ipc.Msgrcv(msgid, 1, 0600)
				if err != nil {
					return
				}
				length, _ := binary.Uvarint(m)
				if _, err := ipc.Semp(semid, ipc.SEM_UNDO); err != nil {
					return
				}
				text = string(data[:length])
				if _, err := ipc.Semv(semid, ipc.SEM_UNDO); err != nil {
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
$ go run main.go -s true
Enter:
HelloWorld
```
In another terminal receive this word.
```sh
$ go run main.go -s false
Recv:
HelloWorld
```

### License
This package is licensed under a MIT license (Copyright (c) 2020 Meng Huang)


### Author
ipc was written by Meng Huang.


