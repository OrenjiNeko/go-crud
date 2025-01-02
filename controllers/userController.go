package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//get the email/password from body
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body)!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		});

		return
	}

	//hash the password
	hash, err :=bcrypt.GenerateFromPassword([]byte(body.Password), 10)	

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",	
		});

		return
	}

	//create the password
	user:= models.User{Email: body.Email, Password: string(hash)}
	result:=initializers.DB.Create(&user)

	if result.Error !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		});

		return	
	}

	//respond
	c.JSON(200, gin.H{
		"message": "User created successfully",
		"user": user,
	})
}

func Login(c *gin.Context) {
	//get the email/password from body
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body)!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		});

		return
	}

	//query user
	var user = models.User{}
	initializers.DB.First(&user, "email = ?", body.Email)

	if(user.ID ==0){
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		});

		return
	}

	//compare sent in password with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		});

		return
	}

	//generate jwt token
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})	

	//sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate token",
		});
		return 
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func GetUser(c *gin.Context) {	
	user,_:=c.Get("user")
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Authorized",
		"user": user,
	})
}