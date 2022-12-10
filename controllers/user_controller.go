package controllers

import (
	"final_project_4/helpers"
	"final_project_4/models"
	"final_project_4/service"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct { // implementasi Controller
	usa service.UserServiceApi
}

func NewUserController(usa service.UserServiceApi) *UserController {
	return &UserController{usa: usa}
}

var (
	appJSON = "application/json"
)

func Valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (uc *UserController) UserRegisterControllers(c *gin.Context) {
	var input models.User

	ContentType := helpers.GetContentType(c)
	var err error = nil

	if ContentType == appJSON {
		err = c.ShouldBindJSON(&input)
	} else {
		err = c.ShouldBind(&input)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Failed Create User",
		})
		return
	}
	user, err := uc.usa.UserRegisterService(input)
	validateEmail := Valid(user.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed Create User / Email Sudah Terdaftar",
		})
		return
	} else if user.Full_name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your full_name is required",
		})
	} else if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your email is required",
		})
	} else if validateEmail == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email Format",
		})
	} else if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your password is required",
		})
	} else if len(user.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password Minimal 6 Karakter",
		})
	} else if user.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your role is Required",
		})
	} else if user.Role != "admin" && user.Role != "customer" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Role Hanya boleh diisi dengan admin atau customer",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":         user.ID,
			"full_name":  user.Full_name,
			"email":      user.Email,
			"password":   user.Password,
			"balance":    user.Balance,
			"created_at": user.CreatedAt,
		})
	}
}

func (uc *UserController) UserLoginControllers(c *gin.Context) {
	var input models.LoginInput

	contentType := helpers.GetContentType(c)
	var err error = nil
	if contentType == appJSON {
		err = c.ShouldBindJSON(&input)
	} else {
		err = c.ShouldBind(&input)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	User, err := uc.usa.UserLoginService(input)
	token := helpers.GenerateToken(User.ID, User.Email, User.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email / password",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": token,
	})
}

func (uc *UserController) UpdateUserControllers(c *gin.Context) {
	var input models.User

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	contentType := helpers.GetContentType(c)
	var err error = nil
	if contentType == appJSON {
		err = c.ShouldBindJSON(&input)
	} else {
		err = c.ShouldBind(&input)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	user, err := uc.usa.UpdateUserService(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":         userID,
			"Full_name":  input.Full_name,
			"email":      input.Email,
			"updated_at": user.UpdatedAt,
		})
	}
}

func (uc *UserController) TopUpBalanceControllers(c *gin.Context) {
	var input models.User

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	contentType := helpers.GetContentType(c)
	var err error = nil
	if contentType == appJSON {
		err = c.ShouldBindJSON(&input)
	} else {
		err = c.ShouldBind(&input)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	UpdateBalance, _, err := uc.usa.TopUpBalanceService(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	} else if input.Balance == 0 {
		c.JSON(http.StatusOK, gin.H{
			"error": "Balance Tidak Boleh Kurang dari angka 0",
		})
	} else if UpdateBalance.Balance > 100000000 {
		c.JSON(http.StatusOK, gin.H{
			"error": "Balance Tidak Boleh Melebehi angka 100.000.000",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Your Balance Has been Successfully updated to",
			"Rp":      UpdateBalance.Balance,
		})
	}
}

func (uc *UserController) DeleteUserControllers(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var err error = nil

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	err = uc.usa.DeleteUserService(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "Your Account Has Been Successfully Deleted",
		})
	}
}
