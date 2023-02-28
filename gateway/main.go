package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	joonix "github.com/joonix/log"
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

	pConsumer, err := rabbitmq.NewConsumer(url, exchange, exchangeType,
		paymentIncomingQueueName, paymentIncomingQueueName, service+"-tagp")
	if err != nil {
		log.WithError(err).Fatalf("failed start worker")
		panic(err)
	}
	paymentConsumerMessage := make(chan string)
	errConsumer := make(chan error)

	go pConsumer.RetrieveMessage(paymentConsumerMessage, errConsumer)

	closer, err := utils.JaegerTracing("Gateway Service", true, log)
	if err != nil {
		log.WithError(err)
		panic(err)
	}
	defer closer.Close()

	http.HandleFunc("/", controller.Ping)

	log.Infof("Listening on localhost:8181")

	log.Fatal(http.ListenAndServe(":8181", nil))
}
