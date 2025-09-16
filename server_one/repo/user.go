package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
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
	users []*User
}

// constructor or constructor function
func NewUserRepo() UserRepo {
	repo := &userRepo{}

	return repo
}

func (r *userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}

	user.ID = len(r.users) + 1
	r.users = append(r.users, &user)
	return &user, nil
}
func (r *userRepo) Get(userId int) (*User, error) {
	for _, user := range r.users {
		if user.ID == userId {
			return user, nil
		}
	}

	return nil, nil
}
func (r *userRepo) Find(email, password string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}

	return nil, nil
}
func (r *userRepo) List() ([]*User, error) {
	return r.users, nil
}
func (r *userRepo) Delete(userId int) error {
	var tempList []*User

	for _, p := range r.users {
		if p.ID != userId {
			tempList = append(tempList, p)
		}
	}

	r.users = tempList

	return nil
}
func (r *userRepo) Update(user User) (*User, error) {
	for i, p := range r.users {
		if p.ID == user.ID {
			r.users[i] = &user
		}
	}

	return &user, nil
}
