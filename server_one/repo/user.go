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
	// Get(userId int) (*User, error)
	// List() ([]*User, error)
	// Delete(userId int) error
	// Update(user User) (*User, error)
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

// func (r *userRepo) Get(userId int) (*User, error) {
// 	for _, user := range r.users {
// 		if user.ID == userId {
// 			return user, nil
// 		}
// 	}

//		return nil, nil
//	}

// func (r *userRepo) List() ([]*User, error) {
// 	return r.users, nil
// }
// func (r *userRepo) Delete(userId int) error {
// 	var tempList []*User

// 	for _, p := range r.users {
// 		if p.ID != userId {
// 			tempList = append(tempList, p)
// 		}
// 	}

// 	r.users = tempList

// 	return nil
// }
// func (r *userRepo) Update(user User) (*User, error) {
// 	for i, p := range r.users {
// 		if p.ID == user.ID {
// 			r.users[i] = &user
// 		}
// 	}

// 	return &user, nil
// }
