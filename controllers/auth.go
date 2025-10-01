package controllers

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "strings"
    "user-service/config"
    "user-service/models"
    "user-service/utils"
)

// Signup handles user registration
func Signup(c *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    user := models.User{Username: input.Username, Email: input.Email, PasswordHash: string(hashedPassword)}

    result := config.DB.Create(&user)
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Login handles user login and JWT token generation
func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    config.DB.Where("email = ?", input.Email).First(&user)

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, _ := utils.GenerateJWT(user.ID)
    c.JSON(http.StatusOK, gin.H{"token": token, "message": "Login successful"})
}

// Dashboard returns a protected message
func Dashboard(c *gin.Context) {
    userID := c.GetUint("userID")
    c.JSON(http.StatusOK, gin.H{"message": "Welcome to your dashboard", "user_id": userID})
}

// AuthMiddleware validates JWT tokens for protected routes
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }

        tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)
        userID, err := utils.ValidateJWT(tokenStr)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Store userID in context
        c.Set("userID", userID)
        c.Next()
    }
}

