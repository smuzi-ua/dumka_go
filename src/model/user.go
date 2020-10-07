package model

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"os"
	"time"
)

// todo add picture
type User struct {
	Id            int        `json:"id,omitempty" gorm:"default"`
	SchoolId      int        `json:"school_id,omitempty"`
	Name          string     `json:"name,omitempty"`
	Nickname      string     `json:"nickname,omitempty"`
	Code          int        `json:"code,omitempty" gorm:"default"`
	CodeGenerated *time.Time `json:"code_generated,omitempty" gorm:"default"`
	Type          string     `json:"type,omitempty" gorm:"default"`
	BannedTo      *time.Time `json:"banned_to,omitempty" gorm:"default"`
	Verified      bool       `json:"verified,omitempty" gorm:"default"`
	Date          *time.Time `json:"date,omitempty" gorm:"default"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) UpdateCode() {
	u.Code = rand.Intn(900000) + 100000
	ts := time.Now()
	u.CodeGenerated = &ts
}

func (u *User) GenerateAuthToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": u.Id,
		"iot":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("GWT_SECRET_KEY")))

	if err != nil {
		fmt.Println(err)
		//panic(err)
	}

	return tokenString

}

func (u *User) VerifyAuthToken(tokenString string) int {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("GWT_SECRET_KEY")), nil
	})

	if err != nil {
		return 0
	}

	if data, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return data["user"].(int)
	} else {
		return 0
	}
}
