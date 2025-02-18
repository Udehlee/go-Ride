package service

import (
	"github.com/Udehlee/go-Ride/db/db"
	"github.com/Udehlee/go-Ride/models"
)

type Service struct {
	db db.Repo
}

func NewService(db db.Repo) *Service {
	return &Service{
		db: db,
	}

}

func (s *Service) RequestARide(id int, name string) (models.Passenger, error) {
	// TODO: Implement equestARide functionality

}
