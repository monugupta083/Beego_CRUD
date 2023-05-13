package controllers

import (
	"myproject/models"
	"myproject/persistence"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

type UserInterface interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

func (uc *UserController) Index(c *gin.Context) {
	users := persistence.FindAllUsers()
	c.JSON(200, gin.H{"data": users})
}

func (uc *UserController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := persistence.FindUserById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, gin.H{"data": user})
}

func (uc *UserController) Create(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	if user.Name == "" || user.Email == "" || user.Password == "" {
		c.JSON(422, gin.H{"error": "Required fields are missing"})
		return
	}
	err := persistence.CreateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to create user"})
		return
	}
	c.JSON(201, gin.H{"data": user})
}

func (uc *UserController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := persistence.FindUserById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.BindJSON(&user)
	if user.Name == "" || user.Email == "" || user.Password == "" {
		c.JSON(422, gin.H{"error": "Required fields are missing"})
		return
	}
	err = persistence.UpdateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to update user"})
		return
	}
	c.JSON(200, gin.H{"data": user})
}

func (uc *UserController) Destroy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := persistence.DeleteUser(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to delete user"})
		return
	}
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
