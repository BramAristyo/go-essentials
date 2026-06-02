package testing

import (
	"errors"
	"testing"
)

type MockUserRepository struct {
	FakeName  string
	FakeError error
}

func (m *MockUserRepository) GetUserById(id int) (string, error) {
	return m.FakeName, m.FakeError
}

func TestGreetUser_Success(t *testing.T) {
	mockRepo := &MockUserRepository{
		FakeName:  "Jane",
		FakeError: nil,
	}

	service := UserService{Repo: mockRepo}

	res, err := service.Greet(1)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if res != "Hello Jane" {
		t.Errorf("Expected result to be %q, but got %q", "Hello Jane", res)
	}
}

func TestGreetUser_Failed(t *testing.T) {
	mockRepo := &MockUserRepository{
		FakeName:  "",
		FakeError: errors.New("database connection timeout"),
	}

	service := UserService{Repo: mockRepo}

	res, err := service.Greet(1)

	if err == nil {
		t.Errorf("Expected an error to occur, but got nil")
	}

	if res != "" {
		t.Errorf("Expected result to be empty, but got: %q", res)
	}

}
