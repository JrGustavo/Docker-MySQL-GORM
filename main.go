package Docker_Go

import (
	"github.com/fuenrob/server-go/utils"
	"github.com/labstack/echo/v4"
)

func main() {

	utils.InitConnection()

	e := echo.New()

	e.Logger.Fatal(e.Start(":9999"))
}