package repo

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, password string) (*User, error)
	Get(userId int) (*User, error)
	List() ([]*User, error)
	Delete(userId int) error
	Update(user User) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

// constructor or constructor function
func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user User) (*User, error) {
	query := `INSERT INTO users (
		first_name, 
		last_name, 
		email, 
		password, 
		is_shop_owner
	)
	VALUES (
		:first_name,
		:last_name,
		:email,
		:password,
		:is_shop_owner
	) 
	RETURNING id`

	var userID int
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&userID)
	}

	user.ID = userID

	return &user, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner 
	FROM users 
	WHERE email = $1 AND password = $2 
	LIMIT 1`

	var user User
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Get(userId int) (*User, error) {
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner 
	FROM users 
	WHERE id = $1 
	LIMIT 1`

	var user User
	err := r.db.Get(&user, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) List() ([]*User, error) {
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner 
	FROM users`

	var users []*User
	err := r.db.Select(&users, query)
	if err != nil {
		log.Printf("Failed to list users: %v", err)
		return nil, err
	}

	return users, nil
}
func (r *userRepo) Delete(userId int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(query, userId)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return err
	}

	return nil
}
func (r *userRepo) Update(user User) (*User, error) {
	query := `UPDATE users 
	SET first_name = :first_name, 
	last_name = :last_name, 
	email = :email, 
	password = :password, 
	is_shop_owner = :is_shop_owner 
	WHERE id = :id`

	_, err := r.db.NamedExec(query, user)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return nil, err
	}

	return &user, nil
}
