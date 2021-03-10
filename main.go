package main

import (
	"fmt"
	"log"
	"nyaachan/nyaachan"
	"os"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	// Look for the token
	godotenv.Load("config.env")
	token, tExists := os.LookupEnv("TELEGRAM_TOKEN")
	if !tExists {
		log.Panic("Please specify a telegram token in file named config.env")
	}

	// Create a new bot from the token
	b, err := tb.NewBot(
		tb.Settings{
			Token: token,
			Poller: &tb.LongPoller{
				Timeout: 10 * time.Second,
			},
			ParseMode: tb.ModeMarkdown,
		},
	)
	// Set the bot up
	nyaachan.SetupBot(b)
	if err != nil {
		log.Panic(err)
	}

	fmt.Print("Starting The Bot..")

	// Start the bot because duh?
	b.Start()
}
