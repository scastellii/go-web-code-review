package main

import (
	"app/cmd/handlers"
	"app/internal/vehicle/loader"
	"app/internal/vehicle/repository"
	"app/internal/vehicle/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// env
	godotenv.Load(".env")

	// dependencies
	ldVh := loader.NewLoaderVehicleJSON(os.Getenv("FILE_PATH_VEHICLES_JSON"))
	dbVh, err := ldVh.Load()
	if err != nil {
		panic(err)
	}

	rpVh := repository.NewRepositoryVehicleInMemory(dbVh)
	svVh := service.NewServiceVehicleDefault(rpVh)
	ctVh := handlers.NewControllerVehicle(svVh)

	// server
	rt := gin.New()
	// -> middlewares
	rt.Use(gin.Recovery())
	rt.Use(gin.Logger())
	// -> handlers
	api := rt.Group("/api/v1")
	grVh := api.Group("/vehicles")
	{
		grVh.GET("", ctVh.GetAll())
		grVh.GET("/color/:color/year/:year", ctVh.GetByColorAndYear())
		grVh.GET("/brand/:brand/between/:start_year/:end_year", ctVh.GetByBrandAndPeriod())
		grVh.GET("/average_speed/brand/:brand", ctVh.GetSpeedAverageByBrand())
		grVh.GET("/fuel_type/:type", ctVh.GetByFuelType())
		grVh.GET("/weight", ctVh.GetByWeight())

		grVh.POST("", ctVh.AddVehicle())
		grVh.POST("/batch", ctVh.AddVehicles())

		grVh.PUT("/:id/update_speed", ctVh.UpdateSpeed())

		grVh.DELETE("/:id", ctVh.DeleteVehicle())

	}

	// run
	if err := rt.Run(os.Getenv("SERVER_ADDR")); err != nil {
		panic(err)
	}
}
