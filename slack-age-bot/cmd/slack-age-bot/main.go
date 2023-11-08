package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func diffMonths(now time.Time, then time.Time) int {
	diffYears := now.Year() - then.Year()
	if diffYears == 0 {
		return int(now.Month() - then.Month())
	}

	if diffYears == 1 {
		return int(12-then.Month()) + int(now.Month())
	}

	yearsInMonths := (now.Year() - then.Year() - 1) * 12
	return yearsInMonths + int(12-then.Month()) + int(now.Month())
}

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5337215290385-5334022753106-sfipNNFD7n7EiPjgsGcu3YUg")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A059TS2A18T-6162726263426-14ddb2903a7212bcf199a1f9b030dd36c2ea56b85d61f1e0d46653522ca4b5bb")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my dob is <dob> ", &slacker.CommandDefinition{
		Description: "dob calculator",
		Examples:    []string{"my dob is 2000-01-02"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			dob := request.Param("dob")
			mydate, error := time.Parse("2006-01-02", dob)
			if error != nil {
				println("error")
			}
			today := time.Now()
			difference := diffMonths(today, mydate)
			r := fmt.Sprintf("Hello user,\nyou are %d years old!", difference/12)
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
