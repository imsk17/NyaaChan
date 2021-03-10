package nyaachan

import (
	"fmt"
	"net/url"
	"nyaachan/nyaachan/scraper"
	"strconv"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	limit   = 15
	baseURL = "https://nyaa.si"
)

// BotCommands is a struct that holds all our bot commands.
type BotCommands struct {
	Bot *tb.Bot
}

// SetupBot is used to set the bot of for
// commands it can handle.
func SetupBot(b *tb.Bot) {
	bot := BotCommands{
		Bot: b,
	}
	b.Handle("/start", bot.start)
	b.Handle("/anime", bot.anime)
	b.Handle("/latest", bot.latest)
	b.Handle("/help", bot.help)
}

// Start is the handler for /start handler. It tells the user about the bot.
func (b *BotCommands) start(m *tb.Message) {
	b.Bot.Reply(m, `*Hi, I am Nyaa Chan UwU.*

	I can search any anime for you on [Nyaa](https://nyaa.si). To know how to use this bot, Use the /help command. My Master is @imsk17. I Suggest You to not disturb him for no reason.`)
}

// Anime is the handler for /anime handler. It Searches nyaa.si for whatever search term is provided by the user.
func (b *BotCommands) anime(m *tb.Message) {
	query := strings.SplitN(m.Text, " ", 2)[1]
	result := scraper.FindAnime(fmt.Sprintf("https://nyaa.si/?q=%s&f=0&c=1_0", url.PathEscape(query)))
	for i := 0; i < len(result); i += limit {
		batch := result[i:min(i+limit, len(result))]
		res := markDownify(batch)
		b.Bot.Reply(
			m, res,
		)
	}
}

// Latest replies with the latest anime from nyaa
func (b *BotCommands) latest(m *tb.Message) {
	var page int
	split := strings.Split(m.Text, " ")
	if len(split) == 2 {
		page, _ = strconv.Atoi(split[1])
	} else {
		page = 1
	}
	result := scraper.FindAnime(fmt.Sprintf("https://nyaa.si/?p=%v", page))
	b.Bot.Reply(
		m, fmt.Sprintf("Found %v Results ...", len(result)),
	)
	for i := 0; i < len(result); i += limit {
		batch := result[i:min(i+limit, len(result))]
		res := markDownify(batch)
		b.Bot.Reply(
			m, res,
		)
	}
}

// Help command for telling the user how to use the bot.
func (b *BotCommands) help(m *tb.Message) {
	b.Bot.Reply(m, `
	Hello, So You Want To Search [Nyaa](https://nyaa.si/) huh? Let me help you with that. So, Here are the commands that you can use right now -

	/latest _<page no which should be a number>_
	*This command fetches you the latest animes available on nyaa on the page as certified.*
	Note - If no digit is specified, I fetch the 1 Page Results For You ~ UwU ~.

	/anime _<search term which can contain whitespaces>_
	*This command searches for the search term and fetches all the first page results for you.*
	Note - Remember, The term should atleast match for what you are looking for.

	`)
}
