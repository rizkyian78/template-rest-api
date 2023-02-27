package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rizkyian78/deployment/controller"
	"github.com/rizkyian78/deployment/queue/rabbitmq"
	"github.com/rizkyian78/deployment/utils"
	"github.com/sirupsen/logrus"
)

const (
	exchange     = "test"
	service      = "cybersource-worker"
	exchangeType = string(rabbitmq.ExchangeType_DIRECT)

	paymentIncomingQueueName = "deployment.mailer"
)

func main() {
	godotenv.Load()
	log := logrus.WithFields(logrus.Fields{
		"service": "Gateway Service",
	})

	url := os.Getenv("AMQP_URL")
	_, err := rabbitmq.NewProducer(url, "asdasd")

	pConsumer, err := rabbitmq.NewConsumer(url, exchange, exchangeType,
		paymentIncomingQueueName, paymentIncomingQueueName, service+"-tagp")
	if err != nil {
		log.WithError(err).Fatalf("failed start worker")
	}
	paymentConsumerMessage := make(chan string)
	errConsumer := make(chan error)
	go pConsumer.RetrieveMessage(paymentConsumerMessage, errConsumer)

	closer, err := utils.JaegerTracing("Gateway Service", true)
	if err != nil {
		log.WithError(err).Error("asdsad")
	}
	defer closer.Close()

	http.HandleFunc("/", controller.Ping)

	log.Infof("Listening on localhost:8181")

	log.Fatal(http.ListenAndServe(":8181", nil))
}
