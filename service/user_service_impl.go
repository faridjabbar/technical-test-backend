package service

import (
	"technical-test-backend/exception"
	"technical-test-backend/helper"
	"technical-test-backend/model/domain"
	"technical-test-backend/model/web"
	"technical-test-backend/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(
	user repository.UserRepository,
	db *gorm.DB,
	validate *validator.Validate,
) UserService {
	return &UserServiceImpl{
		UserRepository: user,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(request *web.UserCreateRequest, c *gin.Context) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	db := service.DB.Begin()
	defer helper.CommitOrRollback(db)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user := &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Role:     request.Role,
		Password: string(hashedPassword),
	}
	user = service.UserRepository.Create(db, user)

	return user.ToUserResponse()
}

func (service *UserServiceImpl) Login(request *web.UserLoginRequest, c *gin.Context) string {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	db := service.DB.Begin()
	defer helper.CommitOrRollback(db)

	user := service.UserRepository.FindByEmail(db, &request.Email)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		helper.PanicIfError(exception.ErrUnauthorized)
	}

	secret := []byte("Nqjdl1GwhDeV6JjUkqM7nU7flT4O7D5aZPdSpnOgPbY=") // Ganti dengan secret Anda

	claims := jwt.MapClaims{
		"access_uuid": "abc12345-def6-7890-gh12-ijklmnopqrst",
		"authorized":  true,
		"exp":         time.Now().Add(time.Hour * 24 * 7).Unix(), // expire 7 hari
		"id":          user.ID,
		"name":        user.Name,
		"role":        user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	return signedToken

}
