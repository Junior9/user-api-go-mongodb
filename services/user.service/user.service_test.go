package userservice_test

import (
	m "api/models"
	"testing"
	"time"

	user_service "api/services/user.service"
)

func TestCreate(t *testing.T) {
	_, err := user_service.Add(m.User{
		Name:       "Francisco",
		Email:      "fran@api.com",
		CreatedAt:  time.Now(),
		UpdatedtAt: time.Now(),
	})

	if err != nil {
		t.Error("Error en el test:  TestCreate ")
		t.Fail()
	} else {
		t.Log("OK: TestCreate")
	}
}

func TestGets(t *testing.T) {
	_, err := user_service.Get()

	if err != nil {
		t.Error("Error en el test:  TestGet ")
		t.Fail()
	} else {
		t.Log("OK: TestGet")
	}
}

func TestGetByID(t *testing.T) {
	_, err := user_service.GetById("6260c5e5002ee39c965da955")

	if err != nil {
		t.Error("Error en el test:  TestGetByID ")
		t.Fail()
	} else {
		t.Log("OK: TestGetByID")
	}
}
