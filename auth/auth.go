package auth

import (
	"bytes"
	"crypto/sha512"
	"encoding/gob"
	"mac/models"
	"strconv"
	"time"

	"fmt"

	"gopkg.in/macaron.v1"
	"gorm.io/gorm"
)

func HandlePostLogin(ctx *macaron.Context, db *gorm.DB, user models.UserForm) (token string, result bool) {

	users := models.User{Username: user.Username}

	e := db.Take(&users)

	if e != nil {
		fmt.Println(e.Error)
	}

	if user.Username == users.Username && user.Password == users.Password {
		str := (strconv.FormatInt(time.Now().UnixNano(), 10))
		bt := []byte(str)
		ab := sha512.Sum512(bt)

		alunga := []byte{}

		// conversion
		for _, elem := range ab {

			alunga = append(alunga, elem)
		}

		buf := bytes.NewBuffer([]byte(alunga))
		gob.NewDecoder(buf).Decode(&str)
		users.Token = str

		db.Save(&users)

		if e != nil {
			fmt.Println(e.Error)
		}
		return token, true
	} else {
		return "", false
	}

}

/*

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}

		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func handleGetUser(ctx *macaron.Context, db *gorm.DB) {
	res := (ctx).Query("ciao")

	fmt.Println(res)

	var users []User
	e := db.Take(&users)

	if e != nil {
		fmt.Println(e.Error)
	}

	for u, _ := range users {
		fmt.Println(u)
	}

}*/
