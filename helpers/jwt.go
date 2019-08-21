package helpers

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ipoool/laporcuranmor-api/models"
	config "github.com/spf13/viper"
)

// JWT - To encrypt/decrypt token
type JWT struct {
}

// GetToken - Get token
func (h *JWT) GetToken(currentUser models.UserTokenModel) (tokenString string, err error) {
	// generate Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":           currentUser.ID,
		"email":        currentUser.Email,
		"display_name": currentUser.DisplayName,
		"expired":      TimeToString(time.Now().AddDate(0, 0, 1)),
	})
	tokenString, err = token.SignedString([]byte(config.GetString("app.secret_key")))
	if err != nil {
		return
	}
	return
}

// CheckToken - Verify the JWT Token
func (h *JWT) CheckToken(tokenString string) (err error) {
	claim, err := h.ClaimToken(tokenString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	claimExpired := claim["expired"].(string)
	expired, err := StringToTime(claimExpired)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	diff := time.Now().Sub(expired)
	days := int(diff.Hours() / 24)

	if days > 0 {
		err = errors.New("token expired")
		fmt.Println(err.Error())
		return
	}
	return
}

// ClaimToken - Get data token
func (h *JWT) ClaimToken(tokenString string) (claim jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetString("app.secret_key")), nil
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed Verify Token: %s", err.Error()))
		err = errors.New("Invalid token")
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		fmt.Println(fmt.Sprintf("Failed Claim Token: %s", err.Error()))
		err = errors.New("Invalid token")
		return
	}
	return
}
