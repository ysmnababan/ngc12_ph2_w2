package main

import (
	"ngc12/config"
	"ngc12/custommiddleware"
	"ngc12/handler"
	"ngc12/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := config.Connect()

	h := &handler.Repo{DB: db}

	e := echo.New()
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.Use(middleware.Recover())

	// cara 1 logger
	// e.Use(middleware.Logger())

	// cara 2 logger
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))

	// cara 3 logger
	// e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	// 	LogStatus: true,
	// 	LogURI:    true,
	// 	BeforeNextFunc: func(c echo.Context) {
	// 		c.Set("customValueFromContext", 42)
	// 	},
	// 	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
	// 		value, _ := c.Get("customValueFromContext").(int)
	// 		fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
	// 		return nil
	// 	},
	// }))

	// cara 4 logger pake logrus
	// log := logrus.New()
	// e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	// 	LogURI:    true,
	// 	LogStatus: true,
	// 	LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
	// 		log.WithFields(logrus.Fields{
	// 			"URI":    values.URI,
	// 			"status": values.Status,
	// 		}).Info("request")

	// 		return nil
	// 	},
	// }))

	// cara 5 logger pake logrus dengan function
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.Logging(c).Info("incoming request")
			return next(c)
		}
	})

	buy := e.Group("")
	buy.Use(custommiddleware.Auth)
	{
		buy.GET("/filter-product", h.GetProductsFilter)
		buy.GET("/products", h.GetProducts)
		buy.POST("/transactions", h.BuyProduct)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
