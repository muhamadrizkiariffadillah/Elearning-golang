package membership

import (
	"errors"
	"time"
)

type Service interface {
	CreateMembership(userId int) (Membership, error)
	UpdateMembership(userId, duration int64) (Membership, error)
}

type service struct {
	repo Repository
}

func Services(repo Repository) *service {
	return &service{repo}
}
func (s *service) CreateMembership(userId int) (Membership, error) {

	membership := Membership{
		UserId:  userId,
		StartAt: time.Now(),
		EndAt:   time.Now(),
	}

	newMembership, err := s.repo.Create(membership)

	if err != nil {
		return Membership{}, err
	}

	return newMembership, nil
}

func (s *service) UpdateMembership(userId, duration int64) (Membership, error) {

	membership, err := s.repo.FindByUserId(int(userId))

	if err != nil {
		return Membership{}, errors.New("error searching by user id")
	}

	if membership.Id == 0 {
		return Membership{}, errors.New("membership is not found")
	}

	membership.EndAt = time.Now().Add(time.Duration(duration) * time.Hour)

	updatedMembership, err := s.repo.Update(membership)

	if err != nil {
		return Membership{}, errors.New("error to update membership")
	}

	return updatedMembership, nil
}
