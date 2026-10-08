package newsendtouser

import (
	"log"
	terminal "ssh/internal/Terminal"
	"ssh/internal/domain"
	"ssh/internal/repository"
)

type Config struct {
	Q repository.QueueMassage[*domain.TerminalOutput]
	Terminal terminal.Terminal
}

type Producer struct {
	config Config
}

func NewProducer(conn Config) *Producer {
	return &Producer{
		config: conn,
	}
}

func (p *Producer) SendToUser() {
	buf := make([]byte, 4096) 
	for {
		
		n, err := p.config.Terminal.Read(buf)
		if err != nil {
			break
		}

		data := make([]byte, n)
		copy(data, buf[:n])

		log.Printf("SendToUser: %s", string(data))

		err = p.config.Q.Push(&domain.TerminalOutput{
			Result: data,
		})
		if err != nil {
			break
		}
	}
}