package handlers

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    phrases := []string{
        "ну вы и дауны",
        "это кумарит",
        "го хапать это ебет ушки",
        "черный ворк подкормил",
        "рахатни мне сладкий чиназес",
    }

    rand.Seed(time.Now().UnixNano()) // инициализация генератора случайных чисел

    ticker := time.NewTicker(1 * time.Hour)
    defer ticker.Stop()

    // выводим сразу при старте
    fmt.Println(randomPhrase(phrases))

    for {
        <-ticker.C
        fmt.Println(randomPhrase(phrases))
    }
}

func randomPhrase(phrases []string) string {
    return phrases[rand.Intn(len(phrases))]
}
