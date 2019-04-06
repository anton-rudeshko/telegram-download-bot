package main

import (
	"log"
	"os"
	"path"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func parseArgs(args []string) string {
	if len(args) < 2 {
		log.Fatalf("Please provide path to config")
	}

	filepath := args[1]

	return filepath
}

func main() {
	log.Println("Starting up...")

	filepath := parseArgs(os.Args)
	config := ReadConfig(filepath)

	httpClient := MakeHttpClient(config.ProxyUrl)
	bot, err := tgbotapi.NewBotAPIWithClient(config.Token, httpClient)
	if err != nil {
		log.Fatalf("Could not init bot api: %s", err)
	}
	bot.Debug = config.Debug

	log.Printf("Authorized as %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = config.Timeout

	updates, err := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		msg := update.Message

		if needIgnoreUpdate(config, update) {
			continue
		}

		fileUrl, _ := bot.GetFileDirectURL(msg.Document.FileID)
		fileName := path.Join(config.Location, msg.Document.FileName)
		log.Printf(`Downloading "%s" as "%s"`, fileUrl, fileName)

		if err = DownloadFile(httpClient, fileUrl, fileName); err != nil {
			log.Printf("Failed to download file: %s", err)

			continue
		}

		response := tgbotapi.NewMessage(msg.Chat.ID, config.SuccessText)
		response.ReplyToMessageID = msg.MessageID

		if _, err := bot.Send(response); err != nil {
			log.Printf("Failed to send message: %s", err)
		}
	}
}

func needIgnoreUpdate(cfg Config, update tgbotapi.Update) bool {
	uid := update.UpdateID
	msg := update.Message

	if msg == nil {
		log.Printf("[uid:%d] Ignoring non-message update", uid)

		return true
	}

	if msg.From == nil {
		log.Printf(`[uid:%d] Ignoring message without sender`, uid)

		return true
	}

	if msg.Document == nil {
		log.Printf(`[uid:%d] Ignoring non-document update`, uid)

		return true
	}

	if !ContainsInt(cfg.AllowedUserIds, msg.From.ID) {
		log.Printf("[uid:%d] Ignoring document from user %d: not allowed", uid, msg.From.ID)

		return true
	}

	if len(cfg.MimeWhitelist) > 0 && !ContainsString(cfg.MimeWhitelist, msg.Document.MimeType) {
		log.Printf(`[uid:%d] Ignoring document with mime "%s": not allowed`, uid, msg.Document.MimeType)

		return true
	}

	return false
}
