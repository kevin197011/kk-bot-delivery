package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Parse command line arguments
	token := flag.String("token", "", "Telegram Bot Token")
	userID := flag.Int64("user", 0, "Target User ID (Chat ID)")
	filePath := flag.String("file", "", "Path to the file to send")
	flag.Parse()

	// Validate arguments
	if *token == "" {
		fmt.Println("Error: --token is required")
		flag.Usage()
		os.Exit(1)
	}
	if *userID == 0 {
		fmt.Println("Error: --user is required")
		flag.Usage()
		os.Exit(1)
	}
	if *filePath == "" {
		fmt.Println("Error: --file is required")
		flag.Usage()
		os.Exit(1)
	}

	// Validate file existence
	if _, err := os.Stat(*filePath); os.IsNotExist(err) {
		log.Fatalf("Error: File does not exist at path: %s", *filePath)
	}

	// Initialize Bot API
	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		log.Fatalf("Error initializing bot: %v", err)
	}

	// Construct file upload request
	params := tgbotapi.NewDocument(*userID, tgbotapi.FilePath(*filePath))

	// Send file
	fmt.Printf("Sending file %s to user %d...\n", *filePath, *userID)
	message, err := bot.Send(params)
	if err != nil {
		log.Fatalf("Error sending file: %v", err)
	}

	fmt.Printf("File sent successfully! Message ID: %d\n", message.MessageID)
}
