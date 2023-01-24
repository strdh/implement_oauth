package tests

import (
    "context"
    "database/sql"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "exercise/gooauth/app/models"
    _ "github.com/go-sql-driver/mysql"
    // "exercise/gooauth/utils"
)

//make unit test for user model using mock
type mockDB struct {
    mock.Mock
}

func (m *mockDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
    args = append([]interface{}{ctx, query}, args...)
    ret := m.Called(args...)
    return ret.Get(0).(sql.Result), ret.Error(1)
}

func (m *mockDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
    args = append([]interface{}{ctx, query}, args...)
    ret := m.Called(args...)
    return ret.Get(0).(*sql.Rows), ret.Error(1)
}

//test create user
func TestCreate(t *testing.T) {
    db := new(mockDB)

    model := models.NewUserModel(db)

    user := models.User{
        FullName: "Ngolo Kante",
        Email: "ngolokante@mail.com",
        Username: "ngolokante",
        Password: "ngolokante-secret",
    }

    result := new(sql.Result)
    // result.LastInsertId()
    db.On("ExecContext", mock.Anything, "INSERT INTO users (fullname, email, username, password) VALUES (?, ?, ?, ?)", user.FullName, user.Email, user.Username, user.Password).Return(result, nil)

    newUser := model.Create(context.Background(), nil, user)

    assert.Equal(t, user.FullName, newUser.FullName)
}