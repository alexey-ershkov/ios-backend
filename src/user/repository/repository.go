package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"ios-backend/src/user/models"
)

type PostgresUserRepository struct {
	conn *sqlx.DB
}

func NewPostgresUserRepository(conn *sqlx.DB) PostgresUserRepository {
	return PostgresUserRepository{conn}
}

func (p PostgresUserRepository) Add(c context.Context, newUser models.User) (models.SafeUser, error) {
	var dbUser models.SafeUser
	query := `INSERT into users(
                  nickname, 
                  email, 
                  password,
                  photo) 
                  VALUES ($1,$2,$3,$4) 
                  RETURNING UserID,NickName, Email,Photo`

	err := p.conn.GetContext(c, &dbUser, query, newUser.NickName, newUser.Email, newUser.Password, newUser.Photo)
	return dbUser, err
}

func (p PostgresUserRepository) GetByID(c context.Context, id int) (models.SafeUser, error) {
	var dbUser models.User
	query := `SELECT * from users where UserID=$1`
	err := p.conn.GetContext(c, &dbUser, query, id)
	return models.SafeUser{
		NickName: dbUser.NickName,
		Email:    dbUser.Email,
		Photo:    dbUser.Photo,
		UserID:   dbUser.UserID,
	}, err
}

func (p PostgresUserRepository) GetByEmailAndPassword(ctx context.Context, email string, password string) (models.SafeUser, error) {
	var dbUser models.User
	query := `SELECT * from users where Email=$1 and password=$2`
	err := p.conn.GetContext(ctx, &dbUser, query, email, password)
	fmt.Println(dbUser, err)
	return models.SafeUser{
		NickName: dbUser.NickName,
		Email:    dbUser.Email,
		Photo:    dbUser.Photo,
		UserID:   dbUser.UserID,
	}, err
}
