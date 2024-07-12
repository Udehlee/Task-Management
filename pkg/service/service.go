package service

import "github.com/Udehlee/Task-Management/pkg/store"

// type Service struct {
// 	Store store.Store
// }

// func NewService(store store.Store) *Service {

// 	return &Service{
// 		Store: store,
// 	}
// }

// type Service struct {
// 	Store store.Store
// }

// func NewService(store store.Store) Service {
// 	return Service{
// 		Store: store,
// 	}
// }

type Service struct {
	Store store.Store
}

func NewService(db store.Store) Service {
	return Service{
		Store: db,
	}
}
