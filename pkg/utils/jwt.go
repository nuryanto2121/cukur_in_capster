package util

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// var screet = setting.FileConfigSetting.App.JwtSecret
// var jwtSecret = []byte("secreet")
// var jwtConf *middleware.JWTConfig

// Claims :
type Claims struct {
	CapsterID   string `json:"capster_id,omitempty"`
	CapsterName string `json:"capster_name,omitempty"`
	BarberID    string `json:"barber_id,omitempty"`
	// CompanyID int    `json:"company_id,omitempty"`
	jwt.StandardClaims
}

// GenerateToken :
func GenerateToken(id int, user_name string, barber_id int) (string, error) {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config.json': %v", err)
	}

	var screet = viper.GetString(`jwt_secret`)
	expired_time := viper.GetInt(`expire_jwt`)
	issuer := viper.GetString(`app.issuer`)
	var jwtSecret = []byte(screet)
	// Set custom claims
	// Ids,_ :=strconv.I(id)
	claims := &Claims{
		CapsterID:   strconv.Itoa(id),
		CapsterName: user_name,
		BarberID:    strconv.Itoa(barber_id),
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expired_time)).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// ParseToken :
func ParseToken(token string) (*Claims, error) {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config.json': %v", err)
	}

	var screet = viper.GetString(`jwt_secret`)
	var jwtSecret = []byte(screet)

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GetEmailToken :
func GetEmailToken(email string) string {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config.json': %v", err)
	}

	var screet = viper.GetString(`jwt_secret`)
	// expired_time := viper.GetInt(`expire_jwt`)
	var jwtSecret = []byte(screet)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}

// ParseEmailToken :
func ParseEmailToken(token string) (string, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// tkn, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if err, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
			return err, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}
	claims, _ := tkn.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%s", claims["email"]), nil
}
