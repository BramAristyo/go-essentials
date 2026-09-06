package testing

import "fmt"

type UserRepository interface {
	GetUserById(id int) (string, error)
}

type UserService struct {
	Repo UserRepository
}

func (s *UserService) Greet(id int) (string, error) {
	name, err := s.Repo.GetUserById(id)
	if err != nil {
		return "", fmt.Errorf("failed to get user : %v", err)
	}

	return "Hello " + name, nil
}
