// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

// go test -gcflags="" -c ./pkg/grpc/exec/ -o go-tests/grpc-exec.test
// sudo ./go-tests/grpc-exec.test  [ -test.run TestGrpcExec ]

package exec

import (
	"testing"
)

func TestGrpcExecOutOfOrder(t *testing.T) {
	t.Skip()
	GrpcExecOutOfOrder[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcExecInOrder(t *testing.T) {
	t.Skip()
	GrpcExecInOrder[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcExecMisingParent(t *testing.T) {
	t.Skip()
	GrpcExecMisingParent[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcMissingExec(t *testing.T) {
	t.Skip()
	GrpcMissingExec[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcExecParentOutOfOrder(t *testing.T) {
	t.Skip()
	GrpcExecParentOutOfOrder[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcExecCloneInOrder(t *testing.T) {
	t.Skip()
	GrpcExecCloneInOrder[*MsgExecveEventUnix, *MsgCloneEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcExecCloneOutOfOrder(t *testing.T) {
	t.Skip()
	GrpcExecCloneOutOfOrder[*MsgExecveEventUnix, *MsgCloneEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcParentInOrder(t *testing.T) {
	t.Skip()
	GrpcParentInOrder[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcExecPodInfoInOrder(t *testing.T) {
	t.Skip()
	GrpcExecPodInfoInOrder[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcExecPodInfoOutOfOrder(t *testing.T) {
	t.Skip()
	GrpcExecPodInfoOutOfOrder[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcExecPodInfoInOrderAfter(t *testing.T) {
	t.Skip()
	GrpcExecPodInfoInOrderAfter[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcExecPodInfoOutOfOrderAfter(t *testing.T) {
	t.Skip()
	GrpcExecPodInfoOutOfOrderAfter[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcExecPodInfoDelayedOutOfOrder(t *testing.T) {
	t.Skip()
	GrpcExecPodInfoDelayedOutOfOrder[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcExecPodInfoDelayedInOrder(t *testing.T) {
	t.Skip()
	GrpcExecPodInfoDelayedInOrder[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}

func TestGrpcDelayedExecK8sOutOfOrder(t *testing.T) {
	t.Skip()
	GrpcDelayedExecK8sOutOfOrder[*MsgExecveEventUnix, *MsgExitEventUnix](t)
}
