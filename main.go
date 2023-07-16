package main

import (
	"FiberAuthWithSurrealDb/Database"
	"FiberAuthWithSurrealDb/Routers"
	"FiberAuthWithSurrealDb/Util"
	"github.com/gofiber/fiber/v2"
	"log"
)


func main() {
	app := fiber.New()

	app.Post("/login", routers.Login)
	app.Post("/register", routers.Register)

	database.Connect()
	log.Fatal(app.Listen("127.0.0.1:" + util.GetConfig("server", "port")))

}

// func main(){	
// 	password := "password"
// 	hash:= "$2a$12$0LrbmOa7d9838w5J88gR6ebtvTg9i0X8IArI6NCFMIGHtDBSd1NNe"
// 	fmt.Println(hash)
// 	fmt.Println(checkPassword(password, hash))
// }

// func hashPassword(password string) (string, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
// 	return string(hashedPassword), err
// }

// func checkPassword(password string, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	fmt.Println(hash, password)
// 	return err == nil
// }