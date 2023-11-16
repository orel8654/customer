package office

import (
	"context"
	"customer/pkg/office/pkg/prots"
)

type Office struct {
	prots.UnimplementedOfficeServiceServer
}

func (o *Office) CreateOffice(ctx context.Context, req *prots.CreateOfficeRequest) (*prots.CreateOfficeResponse, error) {

}

func (o *Office) GetOfficeList(ctx context.Context, req *prots.GetOfficeListRequest) (*prots.GetOfficeListResponse, error) {

}
