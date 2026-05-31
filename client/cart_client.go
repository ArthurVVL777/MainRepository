package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"
)

type CartClient struct {
	httpClient *http.Client
}

func NewCartClient() *CartClient {
	return &CartClient{httpClient: &http.Client{Timeout: 10 * time.Second}}
}

type CreateCartResponse struct {
	CartId int `json:"cartId"`
}

type AddItemRequest struct {
	Device map[string]string `json:"device"`
	CartId int               `json:"cartId"`
	Item   struct {
		DishId   string `json:"dishId"`
		Quantity int    `json:"quantity"`
	} `json:"item"`
}

func (c *CartClient) CreateCart(t testing.TB, payload interface{}) (*CreateCartResponse, int) {
	url := "https://api.stage.digital.uni.rest/api/mobile/bff/api/v1/cart/create"

	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Не удалось маршализовать payload: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("Не удалось создать запрос CreateCart: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		t.Fatalf("Ошибка запроса CreateCart: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Не удалось прочитать тело: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("CreateCart вернул ошибку: %d\nТело: %s", resp.StatusCode, string(body))
	}

	var res CreateCartResponse
	if err := json.Unmarshal(body, &res); err != nil {
		t.Logf("Не удалось распарсить ответ: %v", err)
	}

	return &res, resp.StatusCode
}

func (c *CartClient) AddItem(t testing.TB, payload AddItemRequest) int {
	url := "https://api.stage.digital.uni.rest/api/mobile/bff/api/v1/cart/add_item"

	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Не удалось маршализовать AddItem: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("Не удалось создать запрос AddItem: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		t.Fatalf("Ошибка запроса AddItem: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body) // можно не фаталить, если дальше проверяем статус

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("AddItem ошибка: %d\n%s", resp.StatusCode, string(body))
	}

	return resp.StatusCode
}
