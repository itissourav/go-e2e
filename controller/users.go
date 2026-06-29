package controller

import (
	"database/sql"
	"go-e2e/db"
	"go-e2e/models.go"
	"go-e2e/utils"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	DB *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{
		DB: db,
	}
}

func (u *UserController) SignUp(c *gin.Context) {

	var req models.SignupReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if req.Firstname == "" ||
		req.Lastname == "" ||
		req.Email == "" ||
		req.Password == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing input parameters",
		})
		return
	}

	exists, err := db.UserExists(u.DB, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to check existing user",
		})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already registered",
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.SignupReq{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Password:  string(passwordHash),
	}

	err = db.CreateUser(u.DB, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Signup successful",
	})
}

func (u *UserController) ListUsers(c *gin.Context) {

	users, err := db.ListUsers(u.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (u *UserController) Login(c *gin.Context) {

	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email and Password are required",
		})
		return
	}
	log.Println(req, req.Email)

	user, err := db.GetUserByEmail(u.DB, req.Email)
	// log.Println(user.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	jwtToken, err := utils.GenerateJWT(user.Id, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Id, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate refresh token",
		})
		return
	}

	resp := models.LoginResponse{
		Id:           user.Id,
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Email:        user.Email,
		JwtToken:     jwtToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, resp)
}
