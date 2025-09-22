package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/config"
)

type DishResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	CreatedAt    string `json:"created_at"`
	RestaurantID uint   `json:"restaurant_id,omitempty"`
}

type MenuClient interface {
	GetDishByID(dishID uint) (DishResponse, error)
}

type menuClient struct {
	baseURL    string
	endpoint   string
	httpClient *http.Client
}

func NewMenuClient(cfg *config.Config) MenuClient {
	return &menuClient{
		baseURL:  cfg.MenuService.BaseURL,
		endpoint: cfg.MenuService.Endpoint.GetDishByID,
		httpClient: &http.Client{
			Timeout: time.Duration(cfg.MenuService.ConnectTimeout) * time.Second,
		},
	}
}

func (m *menuClient) GetDishByID(dishID uint) (DishResponse, error) {
	url := fmt.Sprintf("%s%s", m.baseURL, strings.Replace(m.endpoint, ":id", fmt.Sprintf("%d", dishID), 1))

	resp, err := m.httpClient.Get(url)
	if err != nil {
		return DishResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return DishResponse{}, fmt.Errorf("failed to fetch dish %d, status: %d", dishID, resp.StatusCode)
	}

	var dish DishResponse
	if err := json.NewDecoder(resp.Body).Decode(&dish); err != nil {
		return DishResponse{}, err
	}

	return dish, nil
}
