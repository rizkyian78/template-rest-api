package rabbitmq

type ExchangeType string

const (
	ExchangeType_DIRECT ExchangeType = "direct"
	ExchangeType_FANOUT ExchangeType = "fanout"
	ExchangeType_HEADER ExchangeType = "headers"
	ExchangeType_TOPIC  ExchangeType = "topic"
)
