package user

import (
	"errors"
	"gin-base-api/pkg/db"
	"gin-base-api/pkg/models"
	"gin-base-api/pkg/modules/token"
	"gin-base-api/pkg/modules/user/dto"
	"gin-base-api/pkg/utils/exception"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db           *gorm.DB
	tokenService *token.TokenService
}

var userServiceInstance *UserService

func NewService() *UserService {
	if userServiceInstance == nil {
		userServiceInstance = &UserService{
			db:           db.New(),
			tokenService: token.NewService(),
		}
	}
	return userServiceInstance
}

func (userService *UserService) GetAllUsers() (gin.H, error) {
	users := []map[string]interface{}{}

	result := userService.db.Model(&models.User{}).Select("id", "name", "email").Scan(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return gin.H{"users": users}, nil
}

func (userService *UserService) GetUserById(id int32) (gin.H, error) {
	user := map[string]interface{}{}

	result := userService.db.Model(&models.User{}).Where("id = ?", id).Select("id", "name", "email").Scan(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return gin.H{"user": user}, nil
}

func (userService *UserService) CreateUser(createUserDto *dto.CreateUserDto) (gin.H, error) {
	user := createUserDto.ToUser()

	if err := encryptPassword(user); err != nil {
		return nil, err
	}

	result := userService.db.Create(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, exception.New("Já existe um usuário com este e-mail.", 409, result.Error.Error())
		}
		return nil, result.Error
	}

	return gin.H{"message": "Usuário criado com sucesso!"}, nil
}

func (userService *UserService) Login(loginDto *dto.LoginDto) (gin.H, error) {
	user := &models.User{}

	result := userService.db.Where("email = ?", loginDto.Email).First(user)

	if result.Error != nil {
		return nil, exception.New("Email ou senha incorretos.", 400, "")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		return nil, exception.New("Email ou senha incorretos.", 400, "")
	}

	accessToken, err := userService.tokenService.NewAccessToken(&token.UserClaims{Email: user.Email, Id: user.Id, Name: user.Name})
	if err != nil {
		return nil, exception.New("Email ou senha incorretos.", 400, "")
	}

	return gin.H{
		"accessToken": accessToken,
		"user":        gin.H{"name": user.Name, "email": user.Email, "id": user.Id},
	}, nil
}

func (userService *UserService) UpdateUser(id int32, updateUserDto *dto.UpdateUserDto) (gin.H, error) {
	user := updateUserDto.ToUser()

	result := userService.db.Model(&models.User{}).Select("name", "email").Where("id = ?", id).Updates(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, exception.New("Já existe um usuário com este e-mail.", 409, result.Error.Error())
		}
		return nil, result.Error
	}

	return gin.H{"message": "Usuário atualizado com sucesso!"}, nil
}

func (userService *UserService) UpdatePassword(id int32, updatePasswordDto *dto.UpdatePasswordDto) (gin.H, error) {
	user := &models.User{}

	result := userService.db.Where("id = ?", id).First(user)

	if result.Error != nil {
		return nil, result.Error
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(updatePasswordDto.OldPassword)); err != nil {
		return nil, exception.New("Senha antiga incorreta.", 400, err.Error())
	}

	user.Password = updatePasswordDto.NewPassword
	if err := encryptPassword(user); err != nil {
		return nil, err
	}

	result = userService.db.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return gin.H{"message": "Senha atualizada com sucesso!"}, nil
}

func (userService *UserService) DeleteUser(id int32) (gin.H, error) {
	result := userService.db.Delete(&models.User{Id: id})
	if result.Error != nil {
		return nil, result.Error
	}

	return gin.H{"message": "Usuário deletado com sucesso!"}, nil
}

func encryptPassword(user *models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}
