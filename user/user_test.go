package user

import (
	"errors"
	"strconv"
	"testing"
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

type TestContext struct {
	bindError error
	response  map[string]interface{}
	request   map[string]interface{}
}

func (tc *TestContext) BindJSON(v interface{}) error {
	if tc.bindError != nil {
		return tc.bindError
	}

	*v.(*User) = tc.request["user"].(User)
	return nil
}

func (tc *TestContext) JSON(code int, v interface{}) {
	tc.response = v.(map[string]interface{})
}

func (tc *TestContext) Param(key string) string {
	return tc.request["param"].(map[string]string)[key]
}

func (tc *TestContext) ParamInt(key string) (int, error) {
	value := tc.request["param"].(map[string]string)[key]
	return strconv.Atoi(value)
}

func (tc *TestContext) Name() string {
	return "MockName"
}

func TestNewUser_Success(t *testing.T) {
	c := &TestContext{
		request: map[string]interface{}{
			"user": User{
				Name: "John sss",
			},
		},
		bindError: nil,
	}
	handler := NewUserHandler(&UserRepositoryMock{})
	handler.NewUser(c)

	want := "Hello John sss"

	if c.response["message"] != want {
		t.Errorf("want %s but got %s", want, c.response["message"])
	}
}

func TestNewUser_Fail(t *testing.T) {
	c := &TestContext{
		request: map[string]interface{}{
			"user": User{
				Name: "",
			},
		},
		bindError: errors.New("Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag"),
	}
	want := "Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag"

	handler := NewUserHandler(&UserRepositoryMock{})
	handler.NewUser(c)

	if c.response["error"] != want {
		t.Errorf("want %s but got %s", want, c.response["error"])
	}
}

// func TestNewUser_Success(t *testing.T) {
// 	handler := NewUserHandler(&UserRepositoryMock{})

// 	w := httptest.NewRecorder()
// 	payload := bytes.NewBufferString(`{"name": "John sss"}`)
// 	req := httptest.NewRequest("POST", "http://localhost:8080/users", payload)
// 	req.Header.Set("Content-Type", "application/json")

// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = req

// 	handler.NewUser(c)

// 	want := `{"id":0,"message":"Hello John sss"}`

// 	if w.Body.String() != want {
// 		t.Errorf("want %s but got %s", want, w.Body.String())
// 	}
// }

// func TestNewUser_Fail(t *testing.T) {
// 	handler := NewUserHandler(&UserRepositoryMock{})

// 	w := httptest.NewRecorder()
// 	payload := bytes.NewBufferString(`{"name": ""}`)
// 	req := httptest.NewRequest("POST", "http://localhost:8080/users", payload)
// 	req.Header.Set("Content-Type", "application/json")

// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = req

// 	handler.NewUser(c)

// 	want := `{"error":"Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`

// 	if w.Body.String() != want {
// 		t.Errorf("want %s but got %s", want, w.Body.String())
// 	}
// }
