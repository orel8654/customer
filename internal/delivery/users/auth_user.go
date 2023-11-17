package users

import (
	"context"
	"customer/internal/db/psql"
	"customer/pkg/customer/pkg/prots"
)

type UserS struct {
	prots.UnimplementedUserServiceServer
	r *psql.Query
}

func (u *UserS) CreateUser(ctx context.Context, req *prots.CreateUserRequest) (*prots.CreateUserResponse, error) {
	user := psql.UserDB{
		Name:       req.Name,
		OfficeUuID: req.OfficeUuid,
	}

	if err := u.r.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return &prots.CreateUserResponse{}, nil
}

func (u *UserS) GetUserList(ctx context.Context, req *prots.GetUserListRequest) (*prots.GetUserListResponse, error) {
	officeID := psql.OfficeID{
		OfficeUuID: req.OfficeUuid,
	}

	//How to serialize data from DB to GRPC format?
	_, err := u.r.GetOfficeUsers(ctx, officeID)
	if err != nil {
		return nil, err
	}

	return nil, nil

}

func NewUserS(repo *psql.Query) *UserS {
	return &UserS{
		r: repo,
	}
}
