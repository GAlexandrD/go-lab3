package scripts

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
