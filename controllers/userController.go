package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rudiath95/blog-rest-API/ini"
	"github.com/rudiath95/blog-rest-API/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c *gin.Context) {
	//check admin status
	user, _ := c.Get("user")
	checkAdmin := user.(models.User).AdminPower

	if checkAdmin == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "StatusUnauthorized",
		})
		return
	}

	//Call DB
	var post []models.User
	ini.DB.Find(&post)

	//Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetUserInfo(c *gin.Context) {
	//check user status
	user, _ := c.Get("user")
	id := user.(models.User).ID

	// Start Association Mode
	var post models.UserInfo

	ini.DB.Raw("SELECT i.id,i.user_refer,i.email,i.first_name,i.last_name,u.admin_power FROM user_infos i INNER JOIN users u ON user_refer = ?", id).Preload("User").Find(&post)

	//Return
	c.JSON(200, gin.H{
		"user": post,
	})
}

func UpdateGetUserInfo(c *gin.Context) {
	//check user status
	user, _ := c.Get("user")
	id := user.(models.User).ID

	var body struct {
		Email     string
		FirstName string
		LastName  string
	}

	c.Bind(&body)
	// Start Association Mode
	var post models.UserInfo

	ini.DB.Raw("UPDATE user_infos SET email = ?,first_name = ?,last_name = ? WHERE user_refer = ?",
		body.Email, body.FirstName, body.LastName, id).Preload("User").Find(&post)

	//Return
	c.JSON(200, gin.H{
		"successfully updated user with id": id,
	})
}

func SignUp(c *gin.Context) {
	//Get the username/pass off req body
	var body struct {
		Username string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//HashPassword
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	//Create the user
	user := models.User{
		Username: body.Username,
		Password: string(hash),
	}
	result := ini.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user or Username already used",
		})
		return
	}

	//Respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	//Get the email/pass off req body
	var body struct {
		Username string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Look up requested user
	var user models.User
	ini.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	//Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	//Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//Sign and get the token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to Create Token",
		})
		return
	}

	//sent it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30 /*path*/, "" /*domain_name*/, "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func EditUser(c *gin.Context) {
	//get id
	id := c.Param("id")

	//check admin status
	user, _ := c.Get("user")
	checkAdmin := user.(models.User).AdminPower

	if checkAdmin == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "StatusUnauthorized",
		})
		return
	}

	var body struct {
		//	Username   string
		AdminPower bool
	}

	c.Bind(&body)
	var post models.User
	ini.DB.Find(&post, id)

	//Update
	ini.DB.Model(&post).Updates(map[string]interface{}{
		//	"Username":   body.Username,
		"AdminPower": body.AdminPower,
	})

	//Return
	c.JSON(200, gin.H{
		"Edit": post,
	})
}

func DeleteUser(c *gin.Context) {
	//get id
	id := c.Param("id")

	//check admin status
	user, _ := c.Get("user")
	checkAdmin := user.(models.User).AdminPower

	if checkAdmin == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "StatusUnauthorized",
		})
		return
	}

	//Delete
	ini.DB.Delete(&models.User{}, id)

	//Return
	c.JSON(200, gin.H{
		"Delete User With id": id,
	})
}
