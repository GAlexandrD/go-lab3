package painter

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"golang.org/x/exp/shiny/screen"
)

func TestLoop_TextureUpdate(t *testing.T) {
	var (
		l  Loop
		tr testReciever
	)
	l.Receiver = &tr
	l.Start(mockScreen{})
	l.Post(UpdateOp)
	l.Post(UpdateOp)

	if tr.LastTexture != nil {
		t.Fatal("Reciever got texture too early")
	}
	l.StopAndWait()

	_, ok := tr.LastTexture.(*mockTexture)
	if !ok {
		t.Fatal("reciever still has not texture")
	}

	if tr.callsCount != 2 {
		t.Fatal("Texture wasn't updated second time")
	}
}
func TestLoop_Post_Single(t *testing.T) {
	var (
		l  Loop
		tr testReciever
	)
	l.Receiver = &tr
	isDone := false

	l.Start(mockScreen{})
	l.Post(OperationFunc(mockOperation(&isDone)))
	l.StopAndWait()

	if len(l.mq.ops) != 0 {
		t.Fatal("message queue still have operations after StopAndWait call")
	}

	if !isDone {
		t.Fatal("Operation wasn't executed")
	}

	if tr.callsCount != 0 {
		t.Fatal("Unexpected update")
	}
}

func TestLoop_Post_Multiple(t *testing.T) {
	var (
		l  Loop
		tr testReciever
	)
	l.Receiver = &tr
	isDone := [3]bool{}

	l.Start(mockScreen{})
	l.Post(OperationFunc(mockOperation(&isDone[0])))
	l.Post(UpdateOp)
	l.Post(OperationFunc(mockOperation(&isDone[1])))
	l.Post(OperationFunc(mockOperation(&isDone[2])))
	l.Post(UpdateOp)
	l.StopAndWait()

	if len(l.mq.ops) != 0 {
		t.Fatal("message queue still have operations after StopAndWait call")
	}

	for i := range isDone {
		if !isDone[i] {
			t.Fatal("Operation wasn't executed")
		}
	}

	if tr.callsCount != 2 {
		t.Fatal("Texture wasn't updated")
	}
}

type testReciever struct {
	LastTexture screen.Texture
	callsCount  int
}

func (tr *testReciever) Update(t screen.Texture) {
	tr.LastTexture = t
	tr.callsCount++
}

func mockOperation(flag *bool) OperationFunc{
	return func(t screen.Texture) {
		*flag = true
	}
}

type mockScreen struct{}

func (m mockScreen) NewBuffer(p image.Point) (screen.Buffer, error)   { panic("not implemented") }
func (m mockScreen) NewTexture(p image.Point) (screen.Texture, error) { return new(mockTexture), nil }
func (m mockScreen) NewWindow(w *screen.NewWindowOptions) (screen.Window, error) {
	panic("not implemented")
}

type mockTexture struct{}

func (m *mockTexture) Release()                {}
func (m *mockTexture) Size() image.Point       { return size }
func (m *mockTexture) Bounds() image.Rectangle { return image.Rectangle{Max: size} }
func (m *mockTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	panic("not implemented")
}
func (m *mockTexture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {}
