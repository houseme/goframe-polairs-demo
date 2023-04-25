// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/houseme/goframe-polaris-demo/api/pbentity"
)

type (
	IUser interface {
		GetById(ctx context.Context, uid uint64) (*pbentity.User, error)
		DeleteById(ctx context.Context, uid uint64) error
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
