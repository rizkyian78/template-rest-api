package rabbitmq

import (
	"fmt"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type Consumer struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	tag          string
	done         chan error
	messagesChan chan amqp.Delivery

	exchange     string
	exchangeType string
	queueName    string
	key          string
}

// NewConsumer will open the connection to rabbitmq server, and return the conn,
func NewConsumer(amqpURL, exchange, exchangeType, queueName, key, ctag string) (*Consumer, error) {
	c := &Consumer{
		conn:    nil,
		channel: nil,
		tag:     ctag,
		done:    make(chan error),

		exchange:     exchange,
		exchangeType: exchangeType,
		key:          key,
		queueName:    queueName,
	}

	var err error
	c.conn, err = amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize amqp consumer: %w", err)
	}

	go func() {
		if err := <-c.conn.NotifyClose(make(chan *amqp.Error)); err != nil {
			fmt.Printf("AMQP connection closed: %s", err)
			time.Sleep(5 * time.Second)
			os.Exit(9)
		}
	}()

	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("unable to get amqp channel: %w", err)
	}

	go func() {
		if err := <-c.channel.NotifyClose(make(chan *amqp.Error)); err != nil {
			fmt.Printf("AMQP channel closed: %s", err)
			time.Sleep(5 * time.Second)
			os.Exit(9)
		}
	}()

	if err := c.channel.ExchangeDeclare(
		c.exchange,     // name
		c.exchangeType, // kind
		true,           // durable
		false,          // auto delete
		false,          // internal
		false,          // no wait
		nil,            // arguments
	); err != nil {
		return nil, fmt.Errorf("unable to declare exchange: %w", err)
	}

	queue, err := c.channel.QueueDeclare(
		c.queueName, // queue name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no wait
		nil,         // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("unable to declare queue: %w", err)
	}
	fmt.Printf("Established Queue (%q %d messages, %d consumers), binding to Exchange (key %q) \n",
		queue.Name, queue.Messages, queue.Consumers, c.key)

	if err := c.channel.QueueBind(
		queue.Name,
		c.key,
		c.exchange,
		false,
		nil,
	); err != nil {
		return nil, fmt.Errorf("unable to bind queue: %w", err)
	}

	return c, nil
}

func (c *Consumer) RetrieveMessage(deliveries chan<- string, errChan chan<- error) error {
	messagesChan, err := c.channel.Consume(
		c.queueName,
		c.tag,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		errChan <- fmt.Errorf("unable to consume from queue: %w", err)
	}

	for d := range messagesChan {
		fmt.Printf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		d.Ack(false)
		deliveries <- string(d.Body)
	}
	return nil
}

func (c *Consumer) Shutdown() error {
	// will close() the deliveries channel
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	// wait for handle() to exit
	return nil
}
