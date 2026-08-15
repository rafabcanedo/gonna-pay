package domains

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Phone    string `json:"phone"`
}

func NewUser(name, email, password, phone string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Phone:    phone,
	}
}

func NewUserWithID(id, name, email, password, phone string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Phone:    phone,
	}
}

func (u *User) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) ComparePassword(candidate string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(candidate))
}
