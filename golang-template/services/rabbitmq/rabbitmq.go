package rabbitmq

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/vietkytech/golang-template/golang-template/config"
	"git.chotot.org/go-common/kit/logger"
	"github.com/streadway/amqp"
)

var log = logger.GetLogger("multirr-rabbitmq")

type HandleMessage func([]byte) error

type RabbitMQConsumerConfig struct {
	Consumer *config.RabbitMQConfig
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
		true,                           // durable
		false,                          // delete when unused
		false,                          // exclusive
		false,                          // no-wait
		amqp.Table{
			"x-dead-letter-exchange":    c.Config.Consumer.Route.Exchange,
			"x-dead-letter-routing-key": fmt.Sprintf("%s.dlx", c.Config.Consumer.ConsumerName),
		}, // arguments
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

	err = ch.QueueBind(
		q.Name,                             // queue name
		c.Config.Consumer.Route.RoutingKey, // routing key
		c.Config.Consumer.Route.Exchange,   // exchange
		false,
		nil)
	if err != nil {
		return fmt.Errorf("Failed to bind queue: %+v", err)
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
				d.Reject(false)
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
