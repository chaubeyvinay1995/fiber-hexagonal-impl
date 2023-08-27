package api

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"hexagonal-fiber-impl/common/constant"
	"hexagonal-fiber-impl/common/logger"
	"hexagonal-fiber-impl/common/zerologImpl"
	"net/http"
	"testing"
)

type LoginStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var LoginCredential = []LoginStruct{
	{"test1@gmail.com", "testing123"},
	{"test2@gmail.com", "testing12345"},
	{"test1gmail.com", "testing12345"},
	{"test1gmail.com", "test"},
	{"chavinay1@gmail.com", "akkistan"},
	{"chavinay2@gmail.com", "akkistan@12"},
}

type RegisterStruct struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

var RegisterCredential = []RegisterStruct{
	{"test1@gmail.com", "testing123", "testing12345"},
	{"test2@gmail.com", "testing12345", "testing"},
	{"test1gmail.com", "testing12345", ""},
	{"test1gmail.com", "test", "test"},
	{"chavinay1@gmail.com", "akkistan", "akkistan"},
	{"chavinay2@gmail.com", "akkistan@12", "akkistan@12"},
}

// TestRegisterApi function is used to test the Register Api.
func TestRegisterApi(t *testing.T) {
	var buf bytes.Buffer
	for _, test := range RegisterCredential {
		login := RegisterStruct{Email: test.Email, Password: test.Password,
			ConfirmPassword: test.ConfirmPassword}
		err := json.NewEncoder(&buf).Encode(login)
		if err != nil {
			t.Errorf("error while encoding data %s", err)
		}
		resp, err := http.Post(constant.BaseURL+"register", constant.ApplicationType, &buf)
		if err != nil {
			t.Errorf("error while http post request %s", err)
		}
		if resp.StatusCode == 200 {
			t.Log("Passed")
		} else {
			t.Log("Failed")
		}
	}
}

// TestLoginApi is used to test the Login api.
func TestLoginApi(t *testing.T) {
	var buf bytes.Buffer
	for _, test := range LoginCredential {
		login := LoginStruct{Email: test.Email, Password: test.Password}
		err := json.NewEncoder(&buf).Encode(login)
		if err != nil {
			t.Errorf("error while encoding data %s", err)
		}
		resp, err := http.Post(constant.BaseURL+"login", constant.ApplicationType, &buf)
		if err != nil {
			t.Errorf("error while http post request %s", err)
		}
		if resp.StatusCode == 200 {
			t.Log("Passed")
		} else {
			t.Log("Failed")
		}
	}
}

// BenchmarkLogrusPackage function is used to test the Benchmark functionality of
// Logrus Package
func BenchmarkLogrusPackage(b *testing.B) {
	fiberContext := fiber.Ctx{}
	for i := 0; i < b.N; i++ {
		logger.Info(&fiberContext, "Hello")
	}
}

// BenchmarkZeroLogPackage function is used to test the Benchmark functionality of
// ZerLog Package
func BenchmarkZeroLogPackage(b *testing.B) {
	fiberContext := fiber.Ctx{}
	for i := 0; i < b.N; i++ {
		zerologImpl.Info(&fiberContext, "hello")
	}
}
