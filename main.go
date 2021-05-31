package main
 
import (
    "os"
    "net/http"    
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

type Response struct {
    Status  int
}

func main() {
  e := echo.New()

  file, err := os.Create(`/app/log/rest_body.log`)
  if err != nil {
      return
  }
  defer file.Close()

  e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
     file.Write(reqBody) 
  }))

  e.POST("/", func(c echo.Context) error {
    return c.JSON(http.StatusOK, Response{
      Status:  http.StatusOK,
    })
  })

  e.Logger.Fatal(e.Start(":8080"))
}
