package dolg

import (
	"Ruslan/internal/history"
	"fmt"
	"net/http"
	"strconv"
)

func CheckDolg(dolgcount int) string {
	dolgMap := map[int]string{
		41000: "Так ты же должен 91000, пузатый",
		0:     "Шляпа, жди карнавал",
		91000: "Идем до сбера, будешь брать кредит",
	}

	if message, exists := dolgMap[dolgcount]; exists {
		return message
	}
	return fmt.Sprintf("Неправильная сумма долга: %d", dolgcount)
}

func DolgHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	sumStr := query.Get("sum")
	if sumStr == "" {
		http.Error(w, "Не названа верная сумма долга (параметр sum)", http.StatusBadRequest)
		return
	}

	sum, err := strconv.Atoi(sumStr)
	if err != nil {
		http.Error(w, "Неверная сумма. Введите число.", http.StatusBadRequest)
		return
	}

	message := CheckDolg(sum)
	fmt.Fprintln(w, message)

	history.LogHistory(fmt.Sprintf("/dolg?sum=%s", sumStr), message)
}

// верный ввод http://localhost:8080/dolg?sum=91000
