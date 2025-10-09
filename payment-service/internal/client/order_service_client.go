package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/config"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/constants"
	"github.com/shopspring/decimal"
)

type OrderResponse struct {
	ID         uint                  `json:"id"`
	ClientID   uint                  `json:"clientId"`
	Total      decimal.Decimal       `json:"total"`
	Status     constants.OrderStatus `json:"status"`
	CreatedAt  string                `json:"created_at"`
	OrderItems []OrderItemResponse   `json:"items"`
}

type OrderItemResponse struct {
	ID           uint            `json:"id"`
	OrderID      uint            `json:"orderId"`
	DishID       uint            `json:"dishId"`
	RestaurantID uint            `json:"restaurantId"`
	Quantity     int             `json:"quantity"`
	Price        decimal.Decimal `json:"price"`
	CreatedAt    string          `json:"created_at"`
}

type OrderClient interface {
	GetOrderById(orderId uint) (OrderResponse, error)
}

type orderClient struct {
	baseURL    string
	endpoint   string
	httpClient *http.Client
}

func NewOrderClient(cfg *config.Config) OrderClient {
	return &orderClient{
		baseURL:  cfg.OrderService.BaseURL,
		endpoint: cfg.OrderService.Endpoint.GetOrderByID,
		httpClient: &http.Client{
			Timeout: time.Duration(cfg.OrderService.ConnectTimeout) * time.Second,
		},
	}
}

func (o *orderClient) GetOrderById(orderId uint) (OrderResponse, error) {
	url := fmt.Sprintf("%s%s", o.baseURL, strings.Replace(o.endpoint, ":id", fmt.Sprintf("%d", orderId), 1))

	resp, err := o.httpClient.Get(url)
	if err != nil {
		return OrderResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return OrderResponse{}, fmt.Errorf("failed to fetch order %d, status: %d", orderId, resp.StatusCode)
	}

	var order OrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&order); err != nil {
		return OrderResponse{}, err
	}

	return order, nil
}
