package delay

import (
	"delayAlert-order-management-system/client/delay"
	"delayAlert-order-management-system/storage"
)

type Service struct {
	store  *storage.Storage
	client *delay.Client
}

func New(store *storage.Storage, client *delay.Client) *Service {
	return &Service{
		store:  store,
		client: client,
	}
}
