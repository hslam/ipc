// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package ipc provides a way to use System V IPC.
// System V IPC includes three interprocess communication mechanisms
// that are widely available on UNIX systems: message queues, semaphore, and shared memory.
package ipc

const (
	// IPC_CREAT creates if key is nonexistent
	IPC_CREAT = 01000

	// IPC_EXCL fails if key exists.
	IPC_EXCL = 02000

	// IPC_NOWAIT returns error no wait.
	IPC_NOWAIT = 04000

	// IPC_PRIVATE is private key
	IPC_PRIVATE = 00000

	// SEM_UNDO sets up adjust on exit entry
	SEM_UNDO = 010000

	// IPC_RMID removes identifier
	IPC_RMID = 0
	// IPC_SET sets ipc_perm options.
	IPC_SET = 1
	// IPC_STAT gets ipc_perm options.
	IPC_STAT = 2
)
