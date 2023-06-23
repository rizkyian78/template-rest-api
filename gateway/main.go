package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rizkyian78/gateway/queue/rabbitmq"
	"github.com/rizkyian78/gateway/router"
	controller "github.com/rizkyian78/gateway/service"
	"github.com/rizkyian78/gateway/utils"
	"github.com/sirupsen/logrus"
)

const (
	mailer = "gateway.mailer"
)

func main() {

	log := logrus.WithFields(logrus.Fields{
		"service": "Gateway Service",
	})

	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Fatalf("failed start worker")
		panic(err)
	}

	url := os.Getenv("AMQP_URL")

	execProducer, error := rabbitmq.NewProducer(url, mailer)
	if error != nil {
		log.WithError(err).Fatalf("failed start cybersource worker")
	}
	if err != nil {
		log.WithError(err).Fatalf("failed start worker")
		panic(err)
	}

	closer, err := utils.JaegerTracing("Gateway Service", true, log)
	if err != nil {
		log.WithError(err)
		panic(err)
	}
	defer closer.Close()

	h, err := controller.NewHandler(execProducer)
	if err != nil {
		log.WithError(err)
		panic(err)
	}
	log.Infof("Listening on localhost:8181")

	router.Router(h).Run(":8181")

}
