// controllers/user_controller.go

package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/seotrang/go-ent/ent"
	"github.com/seotrang/go-ent/models" // Sửa lại import từ models
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func GetUsers(c *gin.Context, client *ent.Client) {
	// Lấy danh sách người dùng từ cơ sở dữ liệu
	users, err := client.User.Query().All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
		return
	}

	// Chuyển đổi dữ liệu từ ent entity sang kiểu dữ liệu JSON mà API mong muốn
	var userList []models.User // Sử dụng models.User thay vì User
	for _, u := range users {
		userList = append(userList, models.User{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Age:   u.Age,
		})
	}

	c.JSON(http.StatusOK, userList)
}

func GetUserByID(c *gin.Context, client *ent.Client) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Lấy người dùng từ DB theo ID
	user, err := client.User.Get(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Chuyển đổi dữ liệu từ ent entity sang kiểu dữ liệu JSON mà API mong muốn
	c.JSON(http.StatusOK, models.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	})
}

func CreateUser(c *gin.Context, client *ent.Client) {
	var newUser models.User // Sử dụng models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	log.Println("user : ")
	log.Println(newUser)

	// Tạo người dùng mới trong DB
	user, err := client.User.
		Create().
		SetName(newUser.Name).
		SetEmail(newUser.Email).
		SetAge(newUser.Age).
		Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context, client *ent.Client) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser models.User // Sử dụng models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Cập nhật người dùng trong DB
	user, err := client.User.
		UpdateOneID(id).
		SetName(updatedUser.Name).
		SetEmail(updatedUser.Email).
		SetAge(updatedUser.Age).
		Save(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context, client *ent.Client) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Xóa người dùng trong DB
	err = client.User.DeleteOneID(id).Exec(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
