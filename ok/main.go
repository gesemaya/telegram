package main

import (
	"log"
	"time"

	tele "github.com/gesemaya/telegram"
)

func main() {
	pref := tele.Settings{
		Token:  "6392888294:AAHw6D0LESoKHcC3xoseU4hwvbSLNrb-p1I",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
