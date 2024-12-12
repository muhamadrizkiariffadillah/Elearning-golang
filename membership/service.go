package membership

import "time"

type Service interface {
	CreateMembership(userId int) (Membership, error)
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
