package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	token := strings.TrimSpace(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatalf("failed to create Telegram bot: %v", err)
	}

	updates, err := bot.UpdatesViaLongPolling(&telego.GetUpdatesParams{Timeout: 10})
	if err != nil {
		log.Fatalf("failed to start long polling: %v", err)
	}
	defer bot.StopLongPolling()

	log.Println("Bot is up and running. Waiting for messages...")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID

		if update.Message.IsCommand() {
			respondToCommand(bot, chatID, update.Message.Command())
			continue
		}

		if text := strings.TrimSpace(update.Message.Text); text != "" {
			reply := fmt.Sprintf("سلام! پیام شما دریافت شد: %s", text)
			if _, err := bot.SendMessage(tu.Message(tu.ID(chatID), reply)); err != nil {
				log.Printf("failed to send echo response: %v", err)
			}
		}
	}
}

func respondToCommand(bot *telego.Bot, chatID int64, command string) {
	switch command {
	case "start":
		send(bot, chatID, "سلام! من یک ربات ساده‌ام و پیام‌های شما را پاسخ می‌دهم.")
	case "help":
		send(bot, chatID, "فقط کافی است هر پیامی بفرستید تا پاسخی دریافت کنید.")
	default:
		send(bot, chatID, "این دستور را نمی‌شناسم، اما می‌توانم به پیام‌های معمولی پاسخ دهم.")
	}
}

func send(bot *telego.Bot, chatID int64, text string) {
	if _, err := bot.SendMessage(tu.Message(tu.ID(chatID), text).WithParseMode(telego.ModeHTML)); err != nil {
		log.Printf("failed to send message: %v", err)
	}

	time.Sleep(200 * time.Millisecond)
}
