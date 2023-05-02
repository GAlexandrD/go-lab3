package painter

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/exp/shiny/screen"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool {
	drawAll(t)
	return true
}

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

var backgroundColor color.Color = color.White

var rect = image.Rect(0, 0, 0, 0)

type T struct {
	X float32
	Y float32
}

var Ts []T

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
type WhiteFill struct{}

func (op WhiteFill) Do(t screen.Texture) bool {
	backgroundColor = color.White
	return false
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
type GreenFill struct{}

func (op GreenFill) Do(t screen.Texture) bool {
	backgroundColor = color.RGBA{G: 0xff, A: 0xff}
	return false
}

type BgRect struct{ X1, Y1, X2, Y2 float32 }

func (op BgRect) Do(t screen.Texture) bool {
	b := t.Bounds()
	xs := int(op.X1 * float32(b.Dx()))
	ys := int(op.Y1 * float32(b.Dy()))
	xf := int(op.X2 * float32(b.Dx()))
	yf := int(op.Y2 * float32(b.Dy()))
	rect.Min.X = xs
	rect.Min.Y = ys
	rect.Max.X = xf
	rect.Max.Y = yf
	return false
}

type AddT struct{ X, Y float32 }

func (op AddT) Do(t screen.Texture) bool {
	x := T{X: op.X, Y: op.Y}
	Ts = append(Ts, x)
	return false
}

type MoveAll struct{ X, Y float32 }

func (op MoveAll) Do(t screen.Texture) bool {
	for i := range Ts {
		Ts[i].X += op.X
		Ts[i].Y += op.Y
	}
	return false
}

type Reset struct{}

func (op Reset) Do(t screen.Texture) bool {
	rect = image.Rect(0, 0, 0, 0)
	backgroundColor = color.Black
	Ts = []T{}
	return false
}

func drawAll(t screen.Texture) {
	t.Fill(t.Bounds(), backgroundColor, draw.Src)
	if rect.Max.X != rect.Min.X || rect.Max.Y != rect.Min.Y {
		t.Fill(rect, color.Black, draw.Src)
	}
	for i := range Ts {
		drawT(Ts[i].X, Ts[i].Y, t)
	}
}

func drawT(x, y float32, t screen.Texture) {
	w, h := float32(t.Bounds().Dx()), float32(t.Bounds().Dy())
	x1 := int(x * w)
	y1 := int(y * h)
	hRect := image.Rect(x1-100, y1-100, x1+100, y1)
	vRect := image.Rect(x1-50, y1, x1+50, y1+100)
	t.Fill(hRect, color.RGBA{255, 255, 0, 255}, draw.Src)
	t.Fill(vRect, color.RGBA{255, 255, 0, 255}, draw.Src)
}
