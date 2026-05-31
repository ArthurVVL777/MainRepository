package tests

import (
	"Roctics/client"
	"net/http"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type CartSuite struct {
	suite.Suite
	cartApi *client.CartClient
	device  map[string]string
}

func (s *CartSuite) BeforeAll(t provider.T) {
	s.cartApi = client.NewCartClient()

	s.device = map[string]string{
		"platform":   "android",
		"version":    "11.0.0.sp27gpv8",
		"deviceName": "CPH2609",
		"osVersion":  "16",
		"storeId":    "74321670",
		"deviceId":   "bf1235c7-a0ca-403f-829b-2cb16dc74a911",
		"deviceType": "mobile",
	}
}

func (s *CartSuite) TestAddSoloItem(t provider.T) {
	t.AllureID("13683")
	t.Title("Добавление соло-блюда в корзину")
	t.Tags("Regression", "Cart")

	var cartId int

	t.WithNewStep("Создание корзины", func(sCtx provider.StepCtx) {
		payload := map[string]interface{}{"device": s.device}
		res, status := s.cartApi.CreateCart(t.RealT(), payload)
		sCtx.Require().Equal(http.StatusOK, status, "Не удалось создать корзину")
		cartId = res.CartId
	})

	t.WithNewStep("Добавление товара", func(sCtx provider.StepCtx) {
		payload := client.AddItemRequest{
			Device: s.device,
			CartId: cartId,
			Item: struct {
				DishId   string `json:"dishId"`
				Quantity int    `json:"quantity"`
			}{
				DishId:   "193584",
				Quantity: 1,
			},
		}

		status := s.cartApi.AddItem(t.RealT(), payload)
		sCtx.Assert().Equal(http.StatusOK, status)
	})
}

func TestCartSuite(t *testing.T) {
	suite.RunSuite(t, new(CartSuite))
}
