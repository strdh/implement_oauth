package models

import (
    "fmt"
    "context"
    "errors"
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

type UserKey struct {
    Id int `json:"id"`
    UserId int `json:"user_id"`
    Key string `json:"key"`
}

type UserModel struct {
    DB *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
    return &UserModel{
        DB: db,
    }
}

func (model *UserModel) Create(ctx context.Context, user User, tokenKey string) User {
    tx, err := model.DB.Begin()
    utils.PanicIfError(err)
    defer utils.CommitOrRollback(tx)

    sql1 := "INSERT INTO users (fullname, email, username, password) VALUES (?, ?, ?, ?)"
    result, err := tx.ExecContext(ctx, sql1, user.FullName, user.Email, user.Username, user.Password)
    utils.PanicIfError(err)

    id, err := result.LastInsertId()
    utils.PanicIfError(err)
    user.Id = int(id)

    sql2 := "INSERT INTO users_key (user_id, token_key) VALUES (?, ?)"
    _, err = tx.ExecContext(ctx, sql2, user.Id, tokenKey)
    utils.PanicIfError(err)

    return user
}

func (model *UserModel) FindByUsername(ctx context.Context, username string) (User, string, error) {
    tx, err := model.DB.Begin()
    utils.PanicIfError(err)
    defer utils.CommitOrRollback(tx)
    sql := "SELECT * FROM users WHERE username = ?"
    rows, err := tx.QueryContext(ctx, sql, username)
    utils.PanicIfError(err)
    defer rows.Close()

    user := User{}
    if rows.Next() {
        err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.Username, &user.Password)
        utils.PanicIfError(err)

        sql2 := "SELECT * FROM users_key WHERE user_id = ?"
        rows2, err := tx.QueryContext(ctx, sql2, user.Id)
        utils.PanicIfError(err)

        userKey := UserKey{}
        if rows2.Next() {
            err := rows2.Scan(&userKey.Id, &userKey.UserId, &userKey.Key)
            utils.PanicIfError(err)
        }

        return user, userKey.Key, nil
    }

    return user, "", errors.New(fmt.Sprintf("User with username %s not found", username))
}