package handlers

import (
	"fmt"
)

var numberResponses = map[int]string{
	41: "Да все кент - бабки",
	91: "Я не даун",
	0:  "Да все марк не кумарь меня",
}

// HandleNumber проверяет, является ли текст числом и возвращает соответствующий ответ
func HandleNumber(text string) string {
	var num int
	if _, err := fmt.Sscanf(text, "%d", &num); err != nil {
		return ""
	}

	if resp, ok := numberResponses[num]; ok {
		return resp
	}

	return "Certified shlyappa bot"
}
