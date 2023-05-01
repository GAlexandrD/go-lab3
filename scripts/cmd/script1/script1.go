package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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



func main() {
	SendCommandToPainter(`white
	bgrect 0.25 0.25 0.75 0.75
	figure 0.5 0.5
	green
	figure 0.6 0.6
	update`)
}