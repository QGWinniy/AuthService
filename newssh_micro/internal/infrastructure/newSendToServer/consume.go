package newsendtoserver

import (
	"log"
	terminal "ssh/internal/Terminal"
	"ssh/internal/domain"
	"ssh/internal/repository"
)

type Config struct {
	Q repository.QueueMassage[*domain.TerminalInput]
	Terminal terminal.Terminal
}

type Consumer struct {
	config Config
}

func NewConsumer(conf Config) *Consumer {
	return &Consumer{
		config: conf,
	}
}

func (c *Consumer) SendToServer() {
	ch := c.config.Q.Pop()
	for {
		terminalInput, ok := <-ch
		if !ok {
			return
		}
		log.Printf("consumer SendToServer terminalInput: %q", terminalInput.Command)
		c.config.Terminal.Write(terminalInput.Command)
	}
}