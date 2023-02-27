package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Producer struct {
	conn             *amqp.Connection
	channel          *amqp.Channel
	done             chan error
	confirmationChan chan amqp.Confirmation

	queueName    string
	exchange     string
	exchangeType string
	key          string
}

type Option func(*Producer)

// NewProducer will open the connection to rabbitmq server, and return the conn,
func NewProducer(amqpURL, queueName string, options ...Option) (*Producer, error) {
	p := &Producer{
		conn:    nil,
		channel: nil,
		done:    make(chan error),

		queueName: queueName,
	}

	// Apply options if there are any, can overwrite default
	for _, option := range options {
		option(p)
	}

	var err error
	p.conn, err = amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize amqp producer: %w", err)
	}

	go func() {
		fmt.Printf("closing: %s", <-p.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	p.channel, err = p.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("unable to get amqp channel: %w", err)
	}

	if _, err := p.channel.QueueDeclare(p.queueName, true, false, false, false, nil); err != nil {
		return nil, err
	}

	if err := p.channel.Confirm(false); err != nil {
		return nil, fmt.Errorf("Channel could not be put into confirm mode: %s", err)
	}

	p.confirmationChan = p.channel.NotifyPublish(make(chan amqp.Confirmation, 1))

	return p, nil
}

func WithExchange(exchange string) Option {
	return func(p *Producer) {
		p.exchange = exchange
	}
}

func (p *Producer) PublishMessage(body string) error {
	err := p.channel.Publish(
		p.exchange,  // exchange
		p.queueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(body),
			DeliveryMode: amqp.Persistent,
			Priority:     0,
		})

	if err != nil {
		return err
	}

	if confirmed := <-p.confirmationChan; confirmed.Ack {
		log.Printf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
	} else {
		log.Printf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
	}

	return nil
}

func (p *Producer) Shutdown() error {
	fmt.Println("shutdown")
	if err := p.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	fmt.Printf("AMQP shutdown OK")

	// wait for handle() to exit
	return nil
}
