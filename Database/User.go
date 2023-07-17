package database

import (
	"FiberAuthWithSurrealDb/Util"
	"fmt"
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
	Column   string `json:"column,omitempty"`
}

func CreateUser(username string, password string, email string) (interface{}, error) {
	var err error
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

	_, err = validateUsernameAndEmail(username, email)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	hashedPassword, err := hashPassword(password)
	
	if err != nil {
		panic(err)
	}

	data, err := db.Query("INSERT INTO user (username, password, email) VALUES ($username, $password, $email);", User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
	})

	if err != nil {
		panic(err)
	}
	util.Log("Database", "User created")

	return data, nil
}

func VerifyUser(username string, password string) (interface{}, error) {
	var err error
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

	if err != nil {
		panic(err)
	}

	isPasswordValid := checkPassword(password, userData["password"].(string))

	if !isPasswordValid {
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

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashedPassword), err
}

func checkPassword(password, hash string) bool {
	sex := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return sex == nil
}

func validateUsernameAndEmail(username, email string) (interface{}, error) {
    usernameSearch := queryUser("username", username)
    emailSearch := queryUser("email", email)

    if len(usernameSearch) > 0 {
        util.Log("Database", "User already exists")
        return nil, fmt.Errorf("user_already_exists")
    }

    if len(emailSearch) > 0 {
        util.Log("Database", "Email already exists")
        return nil, fmt.Errorf("email_already_exists")
    }

    return nil, nil
}

func queryUser(column, value string) []interface{} {
    query := fmt.Sprintf("SELECT * FROM user WHERE %s = $%s limit 1", column, column)
    result, err := db.Query(query, Condition{
		Column: value,
	})
    if err != nil {
        panic(err)
    }
    return result.([]interface{})[0].(map[string]interface{})["result"].([]interface{})
}