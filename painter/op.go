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

func (op updateOp) Do(t screen.Texture) bool { return true }

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
func WhiteFill(t screen.Texture) {
	backgroundColor = color.White
	drawAll(t)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	backgroundColor = color.RGBA{G: 0xff, A: 0xff}
	drawAll(t)
}

func BgRect(x1, y1, x2, y2 float32) OperationFunc {
	return func(t screen.Texture) {
		b := t.Bounds()
		xs := int(x1 * float32(b.Dx()))
		ys := int(y1 * float32(b.Dy()))
		xf := int(x2 * float32(b.Dx()))
		yf := int(y2 * float32(b.Dy()))
		rect.Min.X = xs
		rect.Min.Y = ys
		rect.Max.X = xf
		rect.Max.Y = yf
		drawAll(t)
	}
}

func AddT(x, y float32) OperationFunc {
	return func(t screen.Texture) {
		x := T{X: x, Y: y}
		Ts = append(Ts, x)
		drawAll(t)
	}
}

func MoveAll(x, y float32) OperationFunc {
	println(0)
	return func(t screen.Texture) {
		for i := range Ts {
			Ts[i].X += x
			Ts[i].Y += y
		}
	}
}

func Reset(t screen.Texture) {
	rect = image.Rect(0, 0, 0, 0)
	backgroundColor = color.Black
	Ts = []T{}
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
	hRect := image.Rect(x1-200, y1-100, x1+200, y1)
	vRect := image.Rect(x1-50, y1, x1+50, y1+300)
	t.Fill(hRect, color.RGBA{255, 255, 0, 255}, draw.Src)
	t.Fill(vRect, color.RGBA{255, 255, 0, 255}, draw.Src)
}
