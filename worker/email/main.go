package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	joonix "github.com/joonix/log"
	"github.com/rizkyian78/worker/queue/rabbitmq"
	"github.com/sirupsen/logrus"
)

var (
	exchange     = "mpg"
	service      = "cybersource-worker"
	exchangeType = string(rabbitmq.ExchangeType_DIRECT)

	paymentIncomingQueueName = "gateway.mailer"
	executedQueueName        = "mpg.payment.cybersource.executed"
)

func main() {
	logger := logrus.New()

	logger.SetFormatter(joonix.NewFormatter())

	logger.SetOutput(os.Stdout)

	log := logrus.WithFields(logrus.Fields{
		"service": "Gateway Service",
	})

	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Fatalf("failed start worker")
		panic(err)
	}

	url := os.Getenv("AMQP_URL")

	_, errors := rabbitmq.NewProducer(url, executedQueueName)
	if errors != nil {
		log.WithError(err).Fatalf("failed start cybersource worker")
	}

	pConsumer, err := rabbitmq.NewConsumer(url, exchange, exchangeType,
		paymentIncomingQueueName, paymentIncomingQueueName, service+"-tagp")
	if err != nil {
		log.WithError(err).Fatalf("failed start worker")
	}

	closeChan := make(chan os.Signal)
	go signal.Notify(closeChan, os.Interrupt)

	paymentConsumerMessage := make(chan string)

	errConsumer := make(chan error)

	go pConsumer.RetrieveMessage(paymentConsumerMessage, errConsumer)

loopMessage:

	for {
		select {

		case msg := <-paymentConsumerMessage:
			ctx := context.Background()

			fmt.Println(msg, ctx, "masok sini")

		case err := <-errConsumer:
			log.WithError(err).Error("failed to consume msg")

		case <-closeChan:
			log.Warning("catch interuption signal")
			time.Sleep(time.Second)
			break loopMessage

		}
	}

	if err := pConsumer.Shutdown(); err != nil {
		log.WithError(err).Fatal("error during shutting down payment consumer")
	}

	log.Warning("cybersource worker shutted down")
}
