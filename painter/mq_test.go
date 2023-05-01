package painter

import (
	"testing"

	"golang.org/x/exp/shiny/screen"
)

func TestMQ_Push(t *testing.T) {
	var mq = messageQueue{}
	if len(mq.ops) != 0 {
		t.Fatal("mq have elements after creation")
	}
	mq.push(mockOp{})
	if len(mq.ops) != 1 {
		t.Fatal("element wasn't added")
	}
	mq.push(mockOp{})
	mq.push(mockOp{})
	mq.push(mockOp{})
	mq.push(mockOp{})

	if len(mq.ops) != 5 {
		t.Fatal("elements wasn't added")
	}
}

func TestMQ_Pull(t *testing.T) {
	var mq = messageQueue{}
	mq.push(mockOp{})

	firstEl := mq.ops[0]
	pulled := mq.pull()
	if pulled != firstEl {
		t.Fatal("wrong element pulled")
	}

	if len(mq.ops) != 0 {
		t.Fatal("element wasn't deleted")
	}
	mq.push(mockOp{})
	mq.push(mockOp{})
	mq.push(mockOp{})
	mq.push(mockOp{})

	secondEl := mq.ops[1]
	mq.pull()
	pulledSec := mq.pull()

	if len(mq.ops) != 2 {
		t.Fatal("elements wasn't deleted")
	}

	if secondEl != pulledSec {
		t.Fatal("wrong element pulled")
	}
}

func TestMQ_IsEmpty(t *testing.T) {
	var mq = messageQueue{}
	mq.push(mockOp{})
	mq.push(mockOp{})
	mq.push(mockOp{})
	mq.push(mockOp{})

	mq.pull()
	mq.pull()

	if mq.isEmpty() {
		t.Fatal("not empty yet")
	}
	mq.pull()
	mq.pull()
	if !mq.isEmpty() {
		t.Fatal("queue already empty")
	}
}

type mockOp struct{}

func (mo mockOp) Do(t screen.Texture) bool { return false }
