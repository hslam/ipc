// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package ipc

import (
	"github.com/hslam/sem"
	"unsafe"
)

// Sembuf represents an operation.
type Sembuf struct {
	SemNum uint16
	SemOp  int16
	SemFlg int16
}

// Semget calls the semget system call.
//
// The semget() system call returns the System V semaphore set identifier
// associated with the argument key.
//
// A new set of nsems semaphores is created if key has the value
// IPC_PRIVATE or if no existing semaphore set is associated with key
// and IPC_CREAT is specified in semflg.
//
// If semflg specifies both IPC_CREAT and IPC_EXCL and a semaphore set
// already exists for key, then semget() fails with errno set to EEXIST.
//
// The argument nsems can be 0 (a don't care) when a semaphore set is
// not being created.  Otherwise, nsems must be greater than 0 and less
// than or equal to the maximum number of semaphores per semaphore set.
//
// If successful, the return value will be the semaphore set identifier,
// otherwise, -1 is returned, with errno indicating the error.
func Semget(key int, nsems int, semflg int) (int, error) {
	return sem.Get(key, nsems, semflg)
}

// Semsetvalue calls the semctl SETVAL system call.
func Semsetvalue(semid int, semnum int, semun int) (bool, error) {
	return sem.SetValue(semid, semnum, semun)
}

// Semgetvalue calls the semctl GETVAL system call.
func Semgetvalue(semid int, semnum int) (int, error) {
	return sem.GetValue(semid, semnum)
}

// Semp calls the semop P system call.
// Flags recognized in semflg are IPC_NOWAIT and SEM_UNDO.
// If an operation specifies SEM_UNDO, it will be automatically undone when the
// process terminates.
func Semp(semid int, semnum int, semflg int) (bool, error) {
	return sem.P(semid, semnum, semflg)
}

// Semv calls the semop V system call.
// Flags recognized in semflg are IPC_NOWAIT and SEM_UNDO.
// If an operation specifies SEM_UNDO, it will be automatically undone when the
// process terminates.
func Semv(semid int, semnum int, semflg int) (bool, error) {
	return sem.V(semid, semnum, semflg)
}

// Semop calls the semop system call.
//
// semop() performs operations on selected semaphores in the set indi‐
// cated by semid.  Each of the nsops elements in the array pointed to
// by sops is a structure that specifies an operation to be performed on
// a single semaphore.  The elements of this structure are of type
// struct sembuf, containing the following members:
//
// unsigned short sem_num;  /* semaphore number */
// short          sem_op;   /* semaphore operation */
// short          sem_flg;  /* operation flags */
//
// Flags recognized in sem_flg are IPC_NOWAIT and SEM_UNDO.
// If an operation specifies SEM_UNDO, it will be automatically undone when the
// process terminates.
//
// The set of operations contained in sops is performed in array order,
// and atomically, that is, the operations are performed either as a
// complete unit, or not at all.  The behavior of the system call if not
// all operations can be performed immediately depends on the presence
// of the IPC_NOWAIT flag in the individual sem_flg fields, as noted be‐
// low.
func Semop(semid int, sops uintptr, nsops int) (bool, error) {
	return sem.Op(semid, sops, nsops)
}

// Semoperate calls the semop system call.
// Flags recognized in SemFlg are IPC_NOWAIT and SEM_UNDO.
// If an operation specifies SEM_UNDO, it will be automatically undone when the
// process terminates.
func Semoperate(semid int, sops []Sembuf) (bool, error) {
	return Semop(semid, uintptr(unsafe.Pointer(&sops[0])), len(sops))
}

// Semrm removes the semaphore with the given id.
func Semrm(semid int) error {
	return sem.Remove(semid)
}
