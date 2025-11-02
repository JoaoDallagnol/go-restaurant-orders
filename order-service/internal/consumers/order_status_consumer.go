package consumers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/config"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/constants"
	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/repository"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderStatusConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	repo    repository.OrderRepository
	cfg     *config.Config
}

func NewOrderStatusConsumer(cfg *config.Config, repo repository.OrderRepository) (*OrderStatusConsumer, error) {
	conn, err := amqp.Dial(cfg.RabbitMQ.Url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	if err := ch.ExchangeDeclare(
		cfg.RabbitMQ.ExchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	q, err := ch.QueueDeclare(
		cfg.RabbitMQ.QueueName, // ex: "order_status_queue"
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	if err := ch.QueueBind(
		q.Name,
		cfg.RabbitMQ.RoutingKey,
		cfg.RabbitMQ.ExchangeName,
		false,
		nil,
	); err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	log.Printf("✅ RabbitMQ consumer ready — listening on queue: %s", q.Name)
	return &OrderStatusConsumer{conn: conn, channel: ch, repo: repo, cfg: cfg}, nil
}

func (c *OrderStatusConsumer) Start() error {
	msgs, err := c.channel.Consume(
		c.cfg.RabbitMQ.QueueName,
		"",
		true,  // auto-ack
		false, // not exclusive
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to start consuming: %w", err)
	}

	go func() {
		for msg := range msgs {
			var payload struct {
				OrderID uint                  `json:"order_id"`
				Status  constants.OrderStatus `json:"status"`
				SentAt  string                `json:"sent_at"`
			}

			if err := json.Unmarshal(msg.Body, &payload); err != nil {
				log.Printf("❌ Failed to parse message: %v", err)
				continue
			}

			log.Printf("📥 Received message: order_id=%d status=%s", payload.OrderID, payload.Status)

			order, _ := c.repo.GetOrderByID(payload.OrderID)
			order.Status = payload.Status
			err = c.repo.UpdateOrder(order)
			if err != nil {
				log.Printf("❌ Failed to update order %d: %v", payload.OrderID, err)
			} else {
				log.Printf("✅ Order %d updated to status %s", payload.OrderID, payload.Status)
			}
		}
	}()

	log.Println("👂 Consumer started, waiting for messages...")
	select {} // block forever
}

func (c *OrderStatusConsumer) Close() {
	if c.channel != nil {
		_ = c.channel.Close()
	}
	if c.conn != nil {
		_ = c.conn.Close()
	}
	log.Println("🔌 RabbitMQ connection closed")
}
