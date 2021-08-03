package rabbitmq

import (
	"fmt"
	"os"
	"os/signal"

	"git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/config"
	"git.chotot.org/go-common/kit/logger"
	"github.com/streadway/amqp"
)

var log = logger.GetLogger("multirr-rabbitmq")

type HandleMessage func([]byte) error

type RabbitMQConsumerConfig struct {
	Consumer *config.RabbitMQConsumerConfig
}

type RabbitMQConsumer struct {
	Config *RabbitMQConsumerConfig
}

func NewRabbitMQConsumer(config *RabbitMQConsumerConfig) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		Config: config,
	}
}

func (c *RabbitMQConsumer) Consume(handle HandleMessage) error {
	conn, err := amqp.Dial(c.Config.Consumer.Address)
	if err != nil {
		return fmt.Errorf("Failed to connect to RabbitMQ: %+v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("Failed to open a channel: %+v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		c.Config.Consumer.ConsumerName, // name
		false,                          // durable
		false,                          // delete when unused
		false,                          // exclusive
		false,                          // no-wait
		nil,                            // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to declare a queue: %+v", err)
	}

	err = ch.Qos(
		c.Config.Consumer.Route.PrefetchCount, // prefetch count
		0,                                     // prefetch size
		false,                                 // global
	)
	if err != nil {
		return fmt.Errorf("Failed to set QoS: %+v", err)
	}

	msgs, err := ch.Consume(
		q.Name,                         // queue
		c.Config.Consumer.ConsumerName, // consumer
		false,                          // auto-ack
		false,                          // exclusive
		false,                          // no-local
		false,                          // no-wait
		nil,                            // args
	)
	if err != nil {
		return fmt.Errorf("Failed to register a consumer: %+v", err)
	}

	go func() {
		for d := range msgs {
			err := handle(d.Body)
			if err != nil {
				log.Errorf("Can't handle message %+v", err)
				d.Reject(true)
				continue
			}

			d.Ack(false)
		}
	}()

	log.Info("[*] Awaiting RPC requests")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	return nil
}
