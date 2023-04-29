package painter

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"golang.org/x/exp/shiny/screen"
)

func TestLoop_Post(t *testing.T) {
	var ( 
		l Loop
		tr testReciever
	)

	l.Receiver = &tr

	l.Start(mockScreen{})
	l.Post(OperationFunc(WhiteFill))
	l.Post(OperationFunc(GreenFill))
	l.Post(UpdateOp)

	if tr.LastTexture != nil {
		t.Fatal("Reciever got texture too early")
	}
	
	l.StopAndWait()
	tx, ok := tr.LastTexture.(*mockTexture)
	if !ok {
		t.Fatal("reciever still has not texture")
	}

	if tx.FillCnt != 2 {
		t.Fatal("wrong amount of fill calls")
	}
}

type testReciever struct {
	LastTexture screen.Texture
}

func (tr *testReciever) Update(t screen.Texture) {
	tr.LastTexture = t
}



type mockScreen struct {}

func (m mockScreen) NewBuffer(p image.Point) (screen.Buffer, error) {
	panic("not implemented")
}

func (m mockScreen) NewTexture(p image.Point) (screen.Texture, error) {
	return new(mockTexture), nil
}

func (m mockScreen) NewWindow(w *screen.NewWindowOptions) (screen.Window, error) {
	panic("not implemented")
}

type mockTexture struct{
	FillCnt int
}

func (m *mockTexture) Release() {}

func (m *mockTexture) Size() image.Point {
	return size
}

func (m *mockTexture) Bounds() image.Rectangle {
	return image.Rectangle{Max: size}
}

func (m *mockTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	panic("not implemented") 
}


func (m *mockTexture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	m.FillCnt++
}