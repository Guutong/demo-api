package user

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type UserRepositoryMock struct{}

func (usm *UserRepositoryMock) NewUser(u *User) error {
	return nil
}

func (usm *UserRepositoryMock) GetUser() ([]User, error) {
	return nil, nil
}

func (usm *UserRepositoryMock) DeleteUser(id int) error {
	return nil
}

func (usm *UserRepositoryMock) UpdateUser(id int, u *User) error {
	return nil
}

func TestNewUser_Success(t *testing.T) {
	handler := NewUserHandler(&UserRepositoryMock{})

	w := httptest.NewRecorder()
	payload := bytes.NewBufferString(`{"name": "John sss"}`)
	req := httptest.NewRequest("POST", "http://localhost:8080/users", payload)
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.NewUser(c)

	want := `{"id":0,"message":"Hello John sss"}`

	if w.Body.String() != want {
		t.Errorf("want %s but got %s", want, w.Body.String())
	}
}

func TestNewUser_Fail(t *testing.T) {
	handler := NewUserHandler(&UserRepositoryMock{})

	w := httptest.NewRecorder()
	payload := bytes.NewBufferString(`{"name": ""}`)
	req := httptest.NewRequest("POST", "http://localhost:8080/users", payload)
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.NewUser(c)

	want := `{"error":"Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`

	if w.Body.String() != want {
		t.Errorf("want %s but got %s", want, w.Body.String())
	}
}
