package database

import (
	"FiberAuthWithSurrealDb/Util"
	"fmt"
	"github.com/surrealdb/surrealdb.go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Signin struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type Condition struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

func CreateUser(username string, password string, email string) (interface{}, error) {
	db, err := surrealdb.New(util.GetConfig("database", "url"))
	if err != nil {
		panic(err)
	}

	if _, err = db.Signin(
		Signin{
			User: util.GetConfig("database", "user"),
			Pass: util.GetConfig("database", "pass"),
		}); err != nil {
		panic(err)
	}

	if _, err = db.Use("auth", "users"); err != nil {
		panic(err)
	}

	usernameSearch, err := db.Query("SELECT * FROM user WHERE username = $username limit 1", Condition{
		Username: username,
	})
	if err != nil {
		panic(err)
	}

	emailSearch, err := db.Query("SELECT * FROM user WHERE email = $email limit 1", Condition{
		Email: email,
	})

	if err != nil {
		panic(err)
	}

	usernameSearchMap := usernameSearch.([]interface{})[0].(map[string]interface{})["result"].([]interface{})
	emailSearchMap := emailSearch.([]interface{})[0].(map[string]interface{})["result"].([]interface{})

	if len(usernameSearchMap) > 0 {
		util.Log("Database", "User already exists")
		return nil, fmt.Errorf("user_already_exists")
	}
	 
	if len(emailSearchMap) > 0 {
		util.Log("Database", "Email already exists")
		return nil, fmt.Errorf("email_already_exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		panic(err)
	}

	data, err := db.Query("INSERT INTO user (username, password, email) VALUES ($username, $password, $email);", User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	})

	if err != nil {
		panic(err)
	}
	util.Log("Database", "User created")

	return data, nil
}

func VerifyUser(username string, password string) (interface{}, error) {
	db, err := surrealdb.New(util.GetConfig("database", "url"))
	if err != nil {
		panic(err)
	}

	if _, err = db.Signin(
		Signin{
			User: util.GetConfig("database", "user"),
			Pass: util.GetConfig("database", "pass"),
		}); err != nil {
		panic(err)
	}
	if _, err = db.Use("auth", "users"); err != nil {
		panic(err)
	}

	//data, err := db.Select("user")
	data, err := db.Query("SELECT * FROM user WHERE username = $username;", map[string]interface{}{
		"username": username,
	})

	if err != nil {
		panic(err)
	}

	dataMap := data.([]interface{})[0].(map[string]interface{})
	dataData := dataMap["result"].([]interface{})

	if len(dataData) == 0 {
		util.Log("Database", "User not found")
		return nil, fmt.Errorf("user_not_found")
	}

	user, err := db.Query("SELECT * FROM user WHERE username = $username", Condition{
		Username: username,
	})
	if err != nil {
		util.Log("Database", "User not found")
		return nil, fmt.Errorf("user_not_found")
	}

	userMap := user.([]interface{})[0].(map[string]interface{})
	userData := userMap["result"].([]interface{})[0].(map[string]interface{})
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	
	if err != nil {
		panic(err)
	}

	if password !=  string(hashedPassword) {
		util.Log("Database", "Invalid Password")
		return nil, fmt.Errorf("invalid_password")
	}

	util.Log("Database", "User verified")

	//Create JWT Token

	token := util.SignToken(util.JWTData{
		Username: username,
		Email:    userData["email"].(string),
		Password: password,
	})
	return token, nil
}
