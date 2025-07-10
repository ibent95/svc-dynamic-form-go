package services

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"svc-dynamic-form-go/configs"
)

func StartApp() {

    log.Println("🌐 Service svc-dynamic-form-go")

    // Load environment variables
    configs.LoadEnv()
    log.Printf("ENV: DRIVER=%s, URL=%s", configs.Env.Driver, configs.Env.URL)

    // Inisialisasi koneksi DB
    configs.InitDB()

    e := echo.New()
    e.HideBanner = true
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    opts := configs.Build()

    configs.SetupRoutes(e, opts)

    address := configs.Env.AppHost + ":" + configs.Env.AppPort

    log.Println("🚀 Golang service running at http://" + address)
    e.Logger.Fatal(e.Start(address))

}
