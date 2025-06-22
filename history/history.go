package history

import (
	"fmt"
	"os"
	"time"
)

const infoFile = "history.txt"

func LogHistory(input, response string) {
	file, err := os.OpenFile("history.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла истории:", err)
		return
	}
	defer file.Close()

	logEntry := fmt.Sprintf("[%s]\nЗапрос: %s\nОтвет: %s\n\n", time.Now().Format("2006-01-02 15:04:05"), input, response)
	_, err = file.WriteString(logEntry)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}
