package TaskPool

import (
	"github.com/wagslane/go-rabbitmq"
	"log"
)

var RabbitConn *rabbitmq.Conn
var RabbitConsumer *rabbitmq.Consumer
var RabbitProducer *rabbitmq.Publisher

func InitRabbit() {
	var err error
	RabbitConn, err = rabbitmq.NewConn("amqp://guest:guest@localhost:5672", rabbitmq.WithConnectionOptionsLogging)
	if err != nil {
		log.Fatal(err.Error())
	}

	RabbitProducer, err = rabbitmq.NewPublisher(
		RabbitConn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("translation"),
		WithTranslationProducerOptions,
	)
	if err != nil {
		log.Fatal(err)
	}

	RabbitConsumer, err = rabbitmq.NewConsumer(RabbitConn, StartTaskPoolConsumer, "update_pool", WithUpdatePoolConsumerOptions)
	if err != nil {
		log.Fatal(err)
	}
}

func WithUpdatePoolConsumerOptions(options *rabbitmq.ConsumerOptions) {
	//options.ExchangeOptions.Durable = false
	options.QueueOptions.Durable = true
	options.QueueOptions.Args["x-queue-type"] = "quorum"
}

func WithTranslationProducerOptions(options *rabbitmq.PublisherOptions) {
	//options.ExchangeOptions.Durable = false
	options.ExchangeOptions.Durable = true
	options.ExchangeOptions.Args["x-single-active-consumer"] = true
	options.ConfirmMode = true
}
