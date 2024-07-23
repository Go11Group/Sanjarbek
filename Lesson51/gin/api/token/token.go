package token

import (
	"fmt"
	"gin/api/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	singingKey = "awecarefvwfcwecweocjfwe oifj cwmefijw ceoifjcwef"
)

func GenToken(req *models.TokenReq) *models.Tokens {

	// header
	token := *jwt.New(jwt.SigningMethodHS256)
	// payload
	cliams := token.Claims.(jwt.MapClaims)
	cliams["userId"] = req.UserId
	cliams["userName"] = req.UserName
	cliams["email"] = req.Email
	cliams["iat"] = time.Now().Unix()
	cliams["exp"] = time.Now().Add(time.Hour).Unix()
	// signing key
	newToken, err := token.SignedString([]byte(singingKey))
	if err != nil {
		log.Print(err)
		return nil
	}

	return &models.Tokens{
		AccesToken: newToken,
	}
}

func ExtaracClaims(tokenStr string) error {

	keyFunc := func(t *jwt.Token) (interface{}, error) {
		return []byte(singingKey), nil
	}

	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return err
	}

	fmt.Println(claims)

	return nil
}
