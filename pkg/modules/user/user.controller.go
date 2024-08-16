package user

import (
	"gin-base-api/pkg/middlewares"
	"gin-base-api/pkg/modules/token"
	"gin-base-api/pkg/modules/user/dto"
	"gin-base-api/pkg/utils/exception"
	"gin-base-api/pkg/utils/interfaces"
	"gin-base-api/pkg/utils/validation"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	UserService *UserService
}

var userControllerInstance *UserController

func NewController() interfaces.ControllerInterface {
	if userControllerInstance == nil {
		userControllerInstance = &UserController{
			UserService: NewService(),
		}
	}
	return userControllerInstance
}

func (userController *UserController) Setup(router *gin.RouterGroup) {

	router.GET("/me", middlewares.Auth(), func(ctx *gin.Context) {
		user, err := token.GetUser(ctx)

		if err != nil {
			exception.Handle(ctx, err)
			return
		}

		ctx.JSON(200, gin.H{"user": user})
	})

	router.GET("", middlewares.Auth(), func(ctx *gin.Context) {

		users, err := userController.UserService.GetAllUsers()
		if err != nil {
			exception.Handle(ctx, err)
			return
		}

		ctx.JSON(200, users)
	})

	router.GET("/:id", middlewares.Auth(), func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

		if err != nil {
			exception.Handle(ctx, err)
			return
		}

		user, err := userController.UserService.GetUserById(int32(id))
		if err != nil {
			exception.Handle(ctx, err)
			return
		}

		ctx.JSON(200, user)
	})

	router.POST("", func(ctx *gin.Context) {
		var createUserDto dto.CreateUserDto
		if err := ctx.ShouldBindJSON(&createUserDto); err != nil {
			exception.Handle(ctx, err)
			return
		}

		validate := validator.New()

		err := validate.Struct(createUserDto)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errorMessage := validation.GenerateValidationMessage(validationErrors)
			exception.Handle(ctx, exception.New(errorMessage, 400, validationErrors.Error()))
			return
		}

		ret, err := userController.UserService.CreateUser(&createUserDto)
		if err != nil {
			exception.Handle(ctx, err)
			return
		}
		ctx.JSON(200, ret)
	})

	router.POST("/login", func(ctx *gin.Context) {
		var loginDto dto.LoginDto

		if err := ctx.ShouldBindJSON(&loginDto); err != nil {
			exception.Handle(ctx, err)
			return
		}

		validate := validator.New()

		err := validate.Struct(loginDto)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errorMessage := validation.GenerateValidationMessage(validationErrors)
			exception.Handle(ctx, exception.New(errorMessage, 400, validationErrors.Error()))
			return
		}

		ret, err := userController.UserService.Login(&loginDto)
		if err != nil {
			exception.Handle(ctx, err)
			return
		}

		accessToken := ret["accessToken"].(*token.AccessToken)

		ctx.SetCookie("accessToken", accessToken.Token, accessToken.ExpirationTime, "/", "", false, true)

		ctx.JSON(200, gin.H{"message": "Login efetuado com sucesso", "user": ret["user"]})
	})

	router.POST("/logout", func(ctx *gin.Context) {
		ctx.SetCookie("accessToken", "", -1, "/", "", false, true)

		ctx.JSON(200, gin.H{"message": "Deslogado com sucesso"})
	})

	router.PATCH("/update-password", middlewares.Auth(), func(ctx *gin.Context) {
		var updatePasswordDto dto.UpdatePasswordDto

		if err := ctx.ShouldBindJSON(&updatePasswordDto); err != nil {
			exception.Handle(ctx, err)
			return
		}

		validate := validator.New()

		err := validate.Struct(updatePasswordDto)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errorMessage := validation.GenerateValidationMessage(validationErrors)
			exception.Handle(ctx, exception.New(errorMessage, 400, validationErrors.Error()))
			return
		}

		user, err := token.GetUser(ctx)
		if err != nil {
			exception.Handle(ctx, err)
			return
		}

		ret, err := userController.UserService.UpdatePassword(user.Id, &updatePasswordDto)

		if err != nil {
			exception.Handle(ctx, err)
			return
		}

		ctx.JSON(200, ret)
	})

	// router.PATCH("/:id", middlewares.Auth(), func(ctx *gin.Context) {
	// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	// 	if err != nil {
	// 		exception.Handle(ctx, err)
	// 		return
	// 	}

	// 	var updateUserDto dto.UpdateUserDto

	// 	if err := ctx.ShouldBindJSON(&updateUserDto); err != nil {
	// 		exception.Handle(ctx, err)
	// 		return
	// 	}

	// 	validate := validator.New()

	// 	err = validate.Struct(updateUserDto)
	// 	if err != nil {
	// 		validationErrors := err.(validator.ValidationErrors)
	// 		errorMessage := validation.GenerateValidationMessage(validationErrors)
	// 		exception.Handle(ctx, exception.New(errorMessage, 400, validationErrors.Error()))
	// 		return
	// 	}

	// 	ret, err := userController.UserService.UpdateUser(int32(id), &updateUserDto)

	// 	if err != nil {
	// 		exception.Handle(ctx, err)
	// 		return
	// 	}

	// 	ctx.JSON(200, ret)
	// })

	// router.DELETE("/:id", middlewares.Auth(), func(ctx *gin.Context) {
	// id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	// if err != nil {
	// 	exception.Handle(ctx, err)
	// 	return
	// }

	// user, err := token.GetUser(ctx)
	// if err != nil {
	// 	exception.Handle(ctx, err)
	// 	return
	// }

	// if user.Id == int32(id) {
	// 	exception.Handle(ctx, exception.New("Não é possível deletar o próprio usuário.", 400, ""))
	// 	return
	// }

	// ret, err := userController.UserService.DeleteUser(int32(id))
	// if err != nil {
	// 	exception.Handle(ctx, err)
	// 	return
	// }

	// ctx.JSON(200, ret)
	// })
}
