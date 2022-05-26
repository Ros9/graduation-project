package event_consumer

import (
	"fmt"
	"graduation-project/CityGO_bot/events"
	"graduation-project/CityGO_bot/models"
	"log"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start(commandHistory *[]models.CommandHistoryItem) error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvents(gotEvents, commandHistory); err != nil {
			log.Print(err)

			continue
		}
	}
}

func (c *Consumer) handleEvents(events []events.Event, commandHistory *[]models.CommandHistoryItem) error {
	fmt.Println("\n\n\nHandleEvents = ", commandHistory, "\n", *commandHistory, "\n", &commandHistory, "======\n\n\n")

	*commandHistory = append(*commandHistory, models.CommandHistoryItem{ChatId: 1, Text: "TEST2222222222222222222222"})

	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}

	return nil
}
