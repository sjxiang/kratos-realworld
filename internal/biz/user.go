package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)


type Profile struct {
}


type User struct {
	Username string 
	Token string
	Email string
	Bio string
	Image string

	// Password string
	PasswordHash string  
}


type UserLogin struct {
	Username string 
	Token string
	Email string
	Bio string
	Image string
}


func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func verifyPassword(hashed string, input string) bool {
	if err :=  bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil {
		return false
	}
	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProfileRepo interface {
	
}

type UserUsecase struct {
	ur UserRepo
	pr ProfileRepo
	
	log  *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	
	return &UserUsecase{
			ur: ur, 
			pr: pr,
			log: log.NewHelper(logger),
		}
}



func (uu *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Username: username,
		Email: email,
		PasswordHash: hashPassword(password),

	}
	if err := uu.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}

	return &UserLogin{
		Email: email,
		Username: username,
		Token: "abc",
	}, nil
}



func (uu *UserUsecase) Login(ctx context.Context, email string, password string) (*UserLogin, error) {
	u, err := uu.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if verifyPassword(u.PasswordHash, password) {
		return nil, errors.New("login fail")
	}

	return &UserLogin{
		Email: u.Email,
		Username: u.Username,
		Token: "abc",
		Bio: u.Bio,
		Image: u.Image,
	}, nil
}