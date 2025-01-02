package controllers

import (
	"AirBnB/internal/database"
	"AirBnB/internal/utils"
	"fmt"
	"html"
	"mime/multipart"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	db       database.Service
	validate *validator.Validate
}

func NewAuthController(db database.Service) *AuthController {
	return &AuthController{
		db:       db,
		validate: validator.New(),
	}
}

type RegisterRequest struct {
	Email        string                `form:"email" validate:"required,email,max=255"`
	Password     string                `form:"password" validate:"required,min=8,max=255"`
	FirstName    string                `form:"firstName" validate:"required,max=255"`
	LastName     string                `form:"lastName" validate:"required,max=255"`
	Country      string                `form:"country" validate:"required,max=50"`
	Phone        string                `form:"phone" validate:"required,max=20,min=10,e164"`
	ProfileImage *multipart.FileHeader `form:"profileImage" validate:""`
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	var req RegisterRequest

	// Parse multipart form
	if _, err := c.MultipartForm(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid form data",
			"details": err.Error(),
		})
	}

	// Parse form fields
	req.Email = c.FormValue("email")
	req.Password = c.FormValue("password")
	req.FirstName = c.FormValue("firstName")
	req.LastName = c.FormValue("lastName")
	req.Country = c.FormValue("country")
	req.Phone = c.FormValue("phone")

	// Get profile image
	file, err := c.FormFile("profileImage")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Profile image is required",
			"details": err.Error(),
			"status":  fiber.StatusUnauthorized,
		})
	}
	req.ProfileImage = file

	// Validate request
	if err := ac.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": utils.FormatValidationErrors(err),
			"status":  fiber.StatusUnauthorized,
		})
	}

	// Check file size (10MB limit)
	const maxFileSize = 10 * 1024 * 1024
	if req.ProfileImage.Size > maxFileSize {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Profile image size must not exceed 10 MB",
			"status": fiber.StatusUnauthorized,
		})
	}

	// Check for existing user
	if existingUser, _ := ac.db.GetUserByEmail(req.Email); existingUser != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error":  "Email already registered",
			"status": fiber.StatusUnauthorized,
		})
	}

	// Save profile image
	filename := utils.GenerateUniqueFilename(req.ProfileImage.Filename)
	filePath := fmt.Sprintf("uploads/profile_images/%s", filename)
	if err := c.SaveFile(req.ProfileImage, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to save profile image",
			"status": fiber.StatusUnauthorized,
		})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to process password",
			"status": fiber.StatusUnauthorized,
		})
	}

	token, err := utils.GenerateVerificationToken()
	if err == nil {
		utils.SendVerificationEmail(req.Email, token)
	}

	// Create user
	newUser := database.User{
		Email:             html.EscapeString(req.Email),
		Password:          string(hashedPassword),
		FirstName:         html.EscapeString(req.FirstName),
		LastName:          html.EscapeString(req.LastName),
		Country:           html.EscapeString(req.Country),
		Phone:             html.EscapeString(req.Phone),
		Profile_Url:       fmt.Sprintf("/uploads/profile_images/%s", filename),
		VerificationToken: string(token),
	}

	if err := ac.db.CreateNewUser(&newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to create user",
			"status": fiber.StatusUnauthorized,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully. Please check your email for verification.",
		"status":  fiber.StatusUnauthorized,
	})
}

// LoginRequest defines the expected JSON structure for user login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"` // Must be a valid email
	Password string `json:"password" validate:"required"`    // Cannot be empty
}

// Login authenticates a user and returns a JWT token
// POST /auth/login
func (ac *AuthController) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	// Get user by email and password
	user, err := ac.db.GetUserByEmailAndPassword(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error during login",
		})
	}

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email, user.FirstName, user.LastName, user.Country, user.Phone, user.VerificationToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error generating token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
		"user": fiber.Map{
			"email":      user.Email,
			"firstName":  user.FirstName,
			"lastName":   user.LastName,
			"country":    user.Country,
			"phone":      user.Phone,
			"profileUrl": user.Profile_Url,
		},
	})
}

// TODO: Verify credentials against database

// TODO: Generate JWT token for authenticated user

// ForgotPasswordRequest defines the expected JSON structure for password reset

// ForgotPassword initiates the password reset process
// POST /auth/forgot-password
func (ac *AuthController) ForgotPassword(c *fiber.Ctx) error {
	// Get the Authorization header - Fiber uses c.Get() not r.Header.Get()
	authHeader := c.Get("Authorization")

	// Extract token and handle error properly with Fiber context
	claims, errResponse := utils.ExtractTokenFromHeader(authHeader)
	if errResponse != nil {
		return c.Status((*errResponse)["status"].(int)).JSON(errResponse)
	}

	// Use the claims data
	userEmail := claims.Email
	token := claims.Verification_token

	err := utils.SendVerificationEmail(userEmail, token)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Failed to send email",
			"detail": err.Error(),
		})
	}

	// TODO: Implement password reset logic here using the claims data

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password reset instructions sent to email",
	})
}

func (ac *AuthController) VerifyEmail(c *fiber.Ctx) error {
	token := strings.TrimSpace(html.EscapeString(c.Query("token")))
	if len(token) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Verification token is required",
		})
	}

	err := ac.db.VerifyUserEmail(token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Failed to verify email",
			"detail": err.Error(),
		})
	}

	// Get verified user data
	user, err := ac.db.GetUserByVerificationToken(token)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	authToken, err := utils.GenerateToken(
		user.ID,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Country,
		user.Phone,
		user.VerificationToken,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate authentication token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Email verified successfully",
		"token":   authToken,
		"user": fiber.Map{
			"email":      user.Email,
			"firstName":  user.FirstName,
			"lastName":   user.LastName,
			"country":    user.Country,
			"phone":      user.Phone,
			"profileUrl": user.Profile_Url,
		},
	})
}
