package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	item struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	items = map[int]*item{}
	seq   = 1
)

// Stores password
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func initiateData() {
	items[0] = &item{ID: 0, Name: "kecap"}
	items[1] = &item{ID: 1, Name: "susu"}
	items[2] = &item{ID: 2, Name: "kopi"}
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Handles wrong password
	expectedPassword, ok := users[username]

	if !ok || expectedPassword != password {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set token claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Sign token with secret
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// Example of extracting JWT token content
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func getItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, items[id])
}

func createItem(c echo.Context) error {
	dataLength := len(items)
	i := &item{
		ID: dataLength + 1,
	}
	if err := c.Bind(i); err != nil {
		return err
	}
	items[i.ID] = i
	seq++
	return c.JSON(http.StatusOK, i)
}

func main() {
	initiateData()

	e := echo.New()

	// Login route
	e.POST("/login", login)

	// Route that requires no authorization
	e.GET("/items/:id", getItem)

	// Route that requires authorization
	r := e.Group("/member")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)
	r.POST("/items", createItem)

	e.Logger.Fatal(e.Start(":9386"))
}
