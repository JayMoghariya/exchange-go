package utils

import (
	"errors"
	"fmt"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func IsStrongPassword(password string) bool {
	// A strong password should be at least 8 characters long and contain a mix of letters, numbers, and special characters.
	if len(password) < 8 {
		return false
	}

	hasLetter := false
	hasNumber := false
	hasSpecialChar := false

	for _, char := range password {
		switch {
		case char >= 'a' && char <= 'z', char >= 'A' && char <= 'Z':
			hasLetter = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case char == '@' || char == '#' || char == '$' || char == '%' || char == '&':
			hasSpecialChar = true
		}
	}

	return hasLetter && hasNumber && hasSpecialChar
}

// HashPassword hashes a plain password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword compares a hashed password with a plain password.
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// GenerateJWT generates a JWT token for a user.
func GenerateJWT(userID uint, username, role string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	jwtExpiration := os.Getenv("JWT_EXPIRATION")
	if jwtExpiration == "" {
		jwtExpiration = "72h" // Default expiration time
	}
	if secret == "" {
		return "", errors.New("JWT_SECRET not set in environment")
	}

	// Parse the expiration duration string (e.g., "72h", "30m")
	duration, err := time.ParseDuration(jwtExpiration)
	if err != nil {
		return "", errors.New("invalid jwt_expiration format")
	}

	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseJWT parses and validates a JWT token, returning the claims if valid.
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET not set in environment")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expVal, ok := claims["exp"]
		if !ok {
			return nil, errors.New("token missing exp claim")
		}
		var exp int64
		switch v := expVal.(type) {
		case float64:
			exp = int64(v)
		case int64:
			exp = v
		default:
			return nil, errors.New("invalid exp claim type")
		}
		if time.Now().Unix() > exp {
			return nil, errors.New("token has expired")
		}
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
