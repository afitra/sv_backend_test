package models

import "github.com/labstack/echo"

// EchoGroup to store routes group
type EchoGroup struct {
	API   *echo.Group
	Token *echo.Group
}
