package store

import "context"

type Storage struct {
	Users interface {
		Create(context.Context, *User) error
	}
}
