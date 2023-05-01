package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"
)

func SendCommandToPainter(command string) error {
	// Створюємо запит
	req, err := http.NewRequest("POST", "http://localhost:17000", nil)
	if err != nil {
		return err
	}

	// Додаємо параметри до запиту
	req.Body = ioutil.NopCloser(strings.NewReader(command))

	// Відправляємо запит
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Перевіряємо код відповіді
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func MoveOnCircle(a int, r float64) {
	x := r * math.Cos(float64(a))
	y := r * math.Sin(float64(a))
	SendCommandToPainter(fmt.Sprintf("move %f %f", x, y))
	SendCommandToPainter("update")
}

func main() {
	SendCommandToPainter("green")
	SendCommandToPainter("figure 0.4 0.2")
	var a = 0
	var r float64 = 0.2
	for {
		MoveOnCircle(a, r)
		time.Sleep(200 * time.Millisecond)
		a++
	}
}
