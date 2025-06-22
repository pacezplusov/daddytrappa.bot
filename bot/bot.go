package bot

import (
	"context"
	"log"
	"strings"

	"Ruslan/internal/bot/handlers"

	"github.com/mymmrac/telego"
)

// StartBot запускает Telegram-бота
func StartBot(bot *telego.Bot) error {
	ctx := context.Background()
	offset := 0

	for {
		updates, err := bot.UpdatesViaLongPolling(ctx, &telego.GetUpdatesParams{
			Offset:  offset,
			Timeout: 60,
		})
		if err != nil {
			log.Printf("Ошибка при получении обновлений: %v", err)
			continue
		}

		for update := range updates {
			if update.UpdateID >= offset {
				offset = update.UpdateID + 1
			}

			if update.Message == nil {
				continue
			}

			msg := update.Message
			text := strings.ToLower(msg.Text)

			// 1. Проверка на число
			if response := handlers.HandleNumber(text); response != "" {
				_, _ = bot.SendMessage(ctx, &telego.SendMessageParams{
					ChatID: telego.ChatID{ID: msg.Chat.ID},
					Text:   response,
				})
				continue
			}

			// 2. Проверка на ключевые фразы
			if response := handlers.HandlePhrase(text); response != "" {
				_, _ = bot.SendMessage(ctx, &telego.SendMessageParams{
					ChatID: telego.ChatID{ID: msg.Chat.ID},
					Text:   response,
				})
			}
		}
	}
}
