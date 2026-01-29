package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/schollz/progressbar/v3"
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

	// Validate file existence and get size
	fileInfo, err := os.Stat(*filePath)
	if os.IsNotExist(err) {
		log.Fatalf("Error: File does not exist at path: %s", *filePath)
	}

	// Open file
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Initialize Bot API
	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		log.Fatalf("Error initializing bot: %v", err)
	}

	fmt.Printf("Sending file %s (%d bytes) to user %d...\n", *filePath, fileInfo.Size(), *userID)

	// Create progress bar
	bar := progressbar.DefaultBytes(
		fileInfo.Size(),
		"uploading",
	)

	// Wrap file reader with progress bar using io.TeeReader
	// As data is read from 'file', it is written to 'bar', updating the progress
	reader := io.TeeReader(file, bar)

	// Construct file upload request using FileReader
	// We use filepath.Base to ensure the sent file has the correct name
	fileName := filepath.Base(*filePath)
	params := tgbotapi.NewDocument(*userID, tgbotapi.FileReader{
		Name:   fileName,
		Reader: reader,
	})

	// Send file
	message, err := bot.Send(params)
	if err != nil {
		fmt.Println() // Ensure valid newline after progress bar
		log.Fatalf("Error sending file: %v", err)
	}

	fmt.Println() // Ensure valid newline after progress bar
	fmt.Printf("File sent successfully! Message ID: %d\n", message.MessageID)
}
