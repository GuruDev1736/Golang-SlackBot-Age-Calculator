package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analytics <-chan *slacker.CommandEvent) {

	for event := range analytics {
		fmt.Println("Command Events ")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6468537462641-6506196249604-5kzAVpkUtSkaYDIcz4eqs73k")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06ETMCSCSW-6506180082532-26c9ea226561832684be189f890d87ce4969aa725f36183d6f8ae11133fcb365")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My yob is <year>", &slacker.CommandDefinition{
		Description: "Yob Calculator",
		// Example:"my yob is 2024",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("Error")
			}

			age := 2024 - yob
			r := fmt.Sprintf("The age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
