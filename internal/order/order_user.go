package order

import (
	"context"
	"customer/pkg/order/pkg/prots"
)

type Order struct {
	prots.UnimplementedOrderServiceServer
}

func (o *Order) CreateOrder(ctx context.Context, req *prots.CreateOrderRequest) (*prots.CreateOrderResponse, error) {

}

func (o *Order) GetActualMenu(ctx context.Context, req *prots.GetActualMenuRequest) (*prots.GetActualMenuResponse, error) {

}
