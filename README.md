# SurrealAuth
Golang auth system with GoFiber and SurrealDb

## Primarily

In this project I wrote an authentication system using **gofiber** and **surrealdb**, you can use it by filling the requirements and use it as a reference for your projects.

## Packages I used in the project

- [color](https://github.com/fatih/color) - color library for console messages
- [go-fiber](https://github.com/gofiber/fiber/v2) - http framework
- [surrealdb](https://github.com/surrealdb/surrealdb.go) - database package
- [jwt](https://github.com/golang-jwt/jwt/v5) - json webtoken 
- [ini](https://github.com/go-ini/ini) - configuration file manager 

## Installation

To install packages
```cmd
 go mod tidy
```
SurrealDb download and launch
```cmd
macos: brew install surrealdb/tap/surreal
linux: curl  -sSf https://install.surrealdb.com |  sh
windows: iwr https://windows.surrealdb.com -useb  | iex

Start Command: surreal start --user root --pass root
``` 

## Support
<p>I would be very happy if you star 😀<p>
