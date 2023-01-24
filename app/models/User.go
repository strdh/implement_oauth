package models

import (
    "context"
    // "errors"
    "database/sql"
    "exercise/gooauth/utils"
)

type User struct {
    Id int `json:"id"`
    FullName string `json:"fullname"`
    Email string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type UserModel struct {
    DB *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
    return &UserModel{
        DB: db,
    }
}

func (model *UserModel) Create(ctx context.Context, user User) User {
    tx, err := model.DB.Begin()
    utils.PanicIfError(err)
    defer utils.CommitOrRollback(tx)
    sql := "INSERT INTO users (fullname, email, username, password) VALUES (?, ?, ?, ?)"

    result, err := tx.ExecContext(ctx, sql, user.FullName, user.Email, user.Username, user.Password)
    utils.PanicIfError(err)

    id, err := result.LastInsertId()
    utils.PanicIfError(err)

    user.Id = int(id)
    return user
}

func (model *UserModel) FindByUsername(ctx context.Context, username string) (User, error) {
    sql := "SELECT * FROM users WHERE username = ?"
    rows, err := tx.QueryContext(ctx, sql, username)
    utils.PanicIfError(err)
    defer rows.Close()

    user := User{}
    if rows.Next() {
        err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.Username, &user.Password)
        utils.PanicIfError(err)
        return user, nil
    }

    return user, errors.New("User not found")
}