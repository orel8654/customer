package user

import (
	"context"
	"customer/pkg/customer/pkg/prots"
)

type User struct {
	prots.UnimplementedUserServiceServer
}

func (u *User) CreateUser(ctx context.Context, req *prots.CreateUserRequest) (*prots.CreateUserResponse, error) {

}

func (u *User) GetUserList(ctx context.Context, req *prots.GetUserListRequest) (*prots.GetUserListResponse, error) {

}
