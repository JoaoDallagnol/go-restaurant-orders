package publishers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/config"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/constants"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderServicePublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	cfg     *config.Config
}

func NewOrderServicePublisher(cfg *config.Config) (*OrderServicePublisher, error) {
	conn, err := amqp.Dial(cfg.RabbitMQ.Url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open RabbitMQ channel: %w", err)
	}

	err = ch.ExchangeDeclare(
		cfg.RabbitMQ.ExchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	log.Printf("✅ Connected to RabbitMQ (Exchange: %s)", cfg.RabbitMQ.ExchangeName)
	return &OrderServicePublisher{conn: conn, channel: ch, cfg: cfg}, nil
}

func (p *OrderServicePublisher) Publish(orderID uint, status constants.OrderStatus) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := map[string]interface{}{
		"order_id": orderID,
		"status":   status,
		"sent_at":  time.Now().Format(time.RFC3339),
	}

	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to serialize message: %w", err)
	}

	err = p.channel.PublishWithContext(ctx,
		p.cfg.RabbitMQ.ExchangeName,
		p.cfg.RabbitMQ.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("📨 Message published — order_id=%d status=%s", orderID, status)
	return nil
}

func (p *OrderServicePublisher) Close() {
	if p.channel != nil {
		_ = p.channel.Close()
	}
	if p.conn != nil {
		_ = p.conn.Close()
	}
	log.Println("🔌 RabbitMQ connection closed")
}
