package users

import (
	"context"
	"customer/internal/db/psql"
	"customer/internal/types"
	"customer/pkg/customer/pkg/prots"
)

type UserS struct {
	prots.UnimplementedUserServiceServer
	r *psql.Query
}

func (u *UserS) CreateUser(ctx context.Context, req *prots.CreateUserRequest) (*prots.CreateUserResponse, error) {
	user := types.UserCustomer{
		Name:       req.Name,
		OfficeUuID: req.OfficeUuid,
	}

	if err := u.r.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return new(prots.CreateUserResponse), nil
}

func (u *UserS) GetUserList(ctx context.Context, req *prots.GetUserListRequest) (*prots.GetUserListResponse, error) {
	officeID := types.OfficeID{
		OfficeUuID: req.OfficeUuid,
	}

	_, err := u.r.GetOfficeUsers(ctx, officeID)
	if err != nil {
		return nil, err
	}

	return new(prots.GetUserListResponse), nil

}

func NewUserS(repo *psql.Query) *UserS {
	return &UserS{
		r: repo,
	}
}
