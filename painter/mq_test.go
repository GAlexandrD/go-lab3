package painter

import (
	"testing"
)

func TestMQ_Push(t *testing.T) {
	var mq = messageQueue{}
	if len(mq.ops) != 0 {
		t.Fatal("mq have elements after creation")
	}
	mq.push(&mockOperation{})
	if len(mq.ops) != 1 {
		t.Fatal("element wasn't added")
	}
	mq.push(&mockOperation{})
	mq.push(&mockOperation{})
	mq.push(&mockOperation{})
	mq.push(&mockOperation{})

	if len(mq.ops) != 5 {
		t.Fatal("elements wasn't added")
	}
}

func TestMQ_Pull(t *testing.T) {
	var mq = messageQueue{}
	mockOps := []mockOperation{{}, {}, {}, {}, {}, {}}
	mq.push(&mockOps[0])

	pulled := mq.pull()
	pulled.Do(&mockTexture{})
	if !mockOps[0].isDone {
		t.Fatal("wrong element pulled")
	}

	if len(mq.ops) != 0 {
		t.Fatal("element wasn't deleted")
	}
	mq.push(&mockOps[1])
	mq.push(&mockOps[2])
	mq.push(&mockOps[3])
	mq.push(&mockOps[4])

	mq.pull()
	pulledSec := mq.pull()
	pulledSec.Do(&mockTexture{})

	if len(mq.ops) != 2 {
		t.Fatal("elements wasn't deleted")
	}

	if !mockOps[2].isDone {
		t.Fatal("wrong element pulled")
	}
}

func TestMQ_IsEmpty(t *testing.T) {
	var mq = messageQueue{}
	mq.push(&mockOperation{})
	mq.push(&mockOperation{})
	mq.push(&mockOperation{})
	mq.push(&mockOperation{})

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

