package services

import (
	"testing"
)

var s UserService

// TestLoginService function used to test the Login service method.
func TestLoginService(t *testing.T) {
	err := s.Login("a", "b")
	if err != nil {
		t.Error("Error while login")
	} else {
		t.Log("success")
	}
}
