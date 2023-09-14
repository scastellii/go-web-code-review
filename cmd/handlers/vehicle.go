package handlers

import (
	"app/internal/domain"
	"app/internal/vehicle/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// NewControllerVehicle returns a new instance of a vehicle controller.
func NewControllerVehicle(st service.ServiceVehicle) *ControllerVehicle {
	return &ControllerVehicle{st: st}
}

// ControllerVehicle is an struct that represents a vehicle controller.
type ControllerVehicle struct {
	// StorageVehicle is the storage of vehicles.
	st service.ServiceVehicle
}

type RequestVehicle struct {
	Id           int     `json:"id"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Registration string  `json:"registration"`
	Year         int     `json:"year"`
	Color        string  `json:"color"`
	MaxSpeed     int     `json:"max_speed"`
	FuelType     string  `json:"fuel_type"`
	Transmission string  `json:"transmission"`
	Passengers   int     `json:"passengers"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Weight       float64 `json:"weight"`
}

// GetAll returns all vehicles.
type VehicleHandler struct {
	Id           int     `json:"id"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Registration string  `json:"registration"`
	Year         int     `json:"year"`
	Color        string  `json:"color"`
	MaxSpeed     int     `json:"max_speed"`
	FuelType     string  `json:"fuel_type"`
	Transmission string  `json:"transmission"`
	Passengers   int     `json:"passengers"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Weight       float64 `json:"weight"`
}
type ResponseBodyList struct {
	Message string            `json:"message"`
	Data    []*VehicleHandler `json:"vehicles"`
	Error   bool              `json:"error"`
}

type ResponseBody struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   bool   `json:"error"`
}

func requestVehicleToVehicle(vehicle RequestVehicle) *domain.Vehicle {
	return &domain.Vehicle{
		Id: vehicle.Id,
		Attributes: domain.VehicleAttributes{
			Brand:        vehicle.Brand,
			Model:        vehicle.Model,
			Registration: vehicle.Registration,
			Year:         vehicle.Year,
			Color:        vehicle.Color,
			MaxSpeed:     vehicle.MaxSpeed,
			FuelType:     vehicle.FuelType,
			Transmission: vehicle.Transmission,
			Passengers:   vehicle.Passengers,
			Height:       vehicle.Height,
			Width:        vehicle.Width,
			Weight:       vehicle.Weight,
		},
	}
}

func vehicleToResponseVehicle(vehicle *domain.Vehicle) *VehicleHandler {
	return &VehicleHandler{
		Id:           vehicle.Id,
		Brand:        vehicle.Attributes.Brand,
		Model:        vehicle.Attributes.Model,
		Registration: vehicle.Attributes.Registration,
		Year:         vehicle.Attributes.Year,
		Color:        vehicle.Attributes.Color,
		MaxSpeed:     vehicle.Attributes.MaxSpeed,
		FuelType:     vehicle.Attributes.FuelType,
		Transmission: vehicle.Attributes.Transmission,
		Passengers:   vehicle.Attributes.Passengers,
		Height:       vehicle.Attributes.Height,
		Width:        vehicle.Attributes.Width,
		Weight:       vehicle.Attributes.Weight,
	}
}

func validateErrors(err error) (code int, body ResponseBodyList) {
	switch {
	case errors.Is(err, service.ErrServiceVehicleNotFound):
		code = http.StatusNotFound
		body = ResponseBodyList{Message: "Not found", Error: true}
		return
	case errors.Is(err, service.ErrServiceVehicleExist):
		code = http.StatusConflict
		body = ResponseBodyList{Message: "Identificador del vehículo ya existente", Error: true}
		return
	case errors.Is(err, service.ErrServiceVehicleNotFoundWithValue):
		code = http.StatusNotFound
		body = ResponseBodyList{Message: "No se encontraron vehículos con esos criterios.", Error: true}
		return
	case errors.Is(err, service.ErrServiceImposibleMaxSpeed):
		code = http.StatusBadRequest
		body = ResponseBodyList{Message: "Velocidad mal formada o fuera de rango.", Error: true}
		return
	default:
		code = http.StatusInternalServerError
		body = ResponseBodyList{Message: "Internal server error", Error: true}
		return
	}
}

func (c *ControllerVehicle) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// ...

		// process
		vehicles, err := c.st.GetAll()
		if err != nil {
			code, body := validateErrors(err)
			ctx.JSON(code, body)
			return
		}

		// response
		code := http.StatusOK
		body := ResponseBodyList{Message: "Success", Data: make([]*VehicleHandler, 0, len(vehicles)), Error: false}
		for _, vehicle := range vehicles {
			body.Data = append(body.Data, vehicleToResponseVehicle(vehicle))
		}
		ctx.JSON(code, body)
	}
}

func (c *ControllerVehicle) AddVehicle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		var requestVehicle RequestVehicle
		err := ctx.ShouldBindJSON(&requestVehicle)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, ResponseBodyList{
				Message: "Bad Request: Datos del vehículo mal formados o incompletos.",
				Data:    nil,
				Error:   true,
			})
			return
		}
		// process
		vehicle, err := c.st.AddVehicle(requestVehicleToVehicle(requestVehicle))
		if err != nil {
			code, body := validateErrors(err)
			ctx.JSON(code, body)
			return
		}
		// response
		code := http.StatusOK
		body := ResponseBody{
			Message: "Success",
			Data:    vehicleToResponseVehicle(vehicle),
			Error:   false,
		}
		ctx.JSON(code, body)
	}
}

func (c *ControllerVehicle) AddVehicles() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		var requestVehicle []RequestVehicle
		err := ctx.ShouldBindJSON(&requestVehicle)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, ResponseBodyList{
				Message: "Bad Request: Datos del vehículo mal formados o incompletos.",
				Data:    nil,
				Error:   true,
			})
			return
		}
		// process
		var vehicles []*domain.Vehicle
		for _, vehicle := range requestVehicle {
			vehicles = append(vehicles, requestVehicleToVehicle(vehicle))
		}
		addedVehicles, err := c.st.AddVehicles(vehicles)
		if err != nil {
			code, body := validateErrors(err)
			ctx.JSON(code, body)
			return
		}
		// response
		var responseData []*VehicleHandler
		for _, vehicle := range addedVehicles {
			responseData = append(responseData, vehicleToResponseVehicle(vehicle))
		}
		code := http.StatusCreated
		body := ResponseBodyList{
			Message: "Vehículos creados exitosamente.",
			Data:    responseData,
			Error:   false,
		}
		ctx.JSON(code, body)
	}
}

func (c *ControllerVehicle) GetByColorAndYear() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		color := ctx.Param("color")
		year := ctx.Param("year")
		intYear, _ := strconv.Atoi(year)

		// process
		vehicles, err := c.st.GetByColorAndYear(color, intYear)
		if err != nil {
			code, body := validateErrors(err)
			ctx.JSON(code, body)
			return
		}

		// response
		code := http.StatusOK
		body := ResponseBodyList{Message: "Success",
			Data:  make([]*VehicleHandler, 0, len(vehicles)),
			Error: false,
		}
		for _, vehicle := range vehicles {
			body.Data = append(body.Data, vehicleToResponseVehicle(vehicle))
		}
		ctx.JSON(code, body)
	}
}

func (c *ControllerVehicle) GetByBrandAndPeriod() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		brand := ctx.Param("brand")
		start := ctx.Param("start_year")
		end := ctx.Param("end_year")
		intStart, _ := strconv.Atoi(start)
		intEnd, _ := strconv.Atoi(end)

		// process
		vehicles, err := c.st.GetByBrandAndPeriod(brand, intStart, intEnd)
		if err != nil {
			code, body := validateErrors(err)
			ctx.JSON(code, body)
			return
		}

		// response
		code := http.StatusOK
		body := ResponseBodyList{Message: "Success", Data: make([]*VehicleHandler, 0, len(vehicles)), Error: false}
		for _, vehicle := range vehicles {
			body.Data = append(body.Data, vehicleToResponseVehicle(vehicle))
		}
		ctx.JSON(code, body)
	}
}

func (c *ControllerVehicle) UpdateSpeed() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		id := ctx.Param("id")
		intId, _ := strconv.Atoi(id)
		var requestVehicle RequestVehicle
		err := ctx.ShouldBindJSON(&requestVehicle)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, ResponseBodyList{
				Message: "Bad Request: Velocidad mal formada o fuera de rango.",
				Data:    nil,
				Error:   true,
			})
			return
		}

		// process
		vehicle := requestVehicleToVehicle(requestVehicle)
		vehicle.Id = intId
		updateVehicle, err := c.st.UpdateSpeed(vehicle)
		if err != nil {
			code, body := validateErrors(err)
			ctx.JSON(code, body)
			return
		}

		// response
		code := http.StatusOK
		body := ResponseBody{
			Message: "Velocidad del vehículo actualizada exitosamente",
			Data:    vehicleToResponseVehicle(updateVehicle),
			Error:   false}
		ctx.JSON(code, body)
	}
}

func (c *ControllerVehicle) GetSpeedAverageByBrand() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		brand := ctx.Param("brand")

		// process
		average, err := c.st.GetSpeedAverageByBrand(brand)
		if err != nil {
			code, body := validateErrors(err)
			ctx.JSON(code, body)
			return
		}

		// response
		code := http.StatusOK
		body := ResponseBody{
			Message: "Success",
			Data:    average,
			Error:   false,
		}
		ctx.JSON(code, body)
	}
}

func (c *ControllerVehicle) GetByFuelType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		fuelType := ctx.Param("type")

		// process
		vehicles, err := c.st.GetByFuelType(fuelType)
		if err != nil {
			code, body := validateErrors(err)
			ctx.JSON(code, body)
			return
		}

		// response
		var resData []*VehicleHandler
		for _, vehicle := range vehicles {
			resData = append(resData, vehicleToResponseVehicle(vehicle))
		}
		code := http.StatusOK
		body := ResponseBodyList{
			Message: "Success",
			Data:    resData,
			Error:   false,
		}
		ctx.JSON(code, body)
	}
}

func (c *ControllerVehicle) DeleteVehicle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		id := ctx.Param("id")
		intId, _ := strconv.Atoi(id)
		// process
		vehicle, err := c.st.DeleteVehicle(intId)
		if err != nil {
			code, body := validateErrors(err)
			ctx.JSON(code, body)
			return
		}
		// response
		code := http.StatusNoContent
		body := ResponseBody{
			Message: "Vehículo eliminado exitosamente.",
			Data:    vehicleToResponseVehicle(vehicle), //si hay un 204 no muestra por mas que lo ponga
			Error:   false,
		}
		ctx.JSON(code, body)
	}
}

func (c *ControllerVehicle) GetByWeight() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		weightMin := ctx.Query("weight_min")
		weightMax := ctx.Query("weight_max")
		weightMaxInt, _ := strconv.Atoi(weightMax)
		weightMinInt, _ := strconv.Atoi(weightMin)

		// process
		vehicles, err := c.st.GetByWeight(float64(weightMinInt), float64(weightMaxInt))
		if err != nil {
			code, body := validateErrors(err)
			ctx.JSON(code, body)
			return
		}
		// response
		var resData []*VehicleHandler
		for _, vehicle := range vehicles {
			resData = append(resData, vehicleToResponseVehicle(vehicle))
		}
		code := http.StatusOK
		body := ResponseBodyList{
			Message: "Success",
			Data:    resData,
			Error:   false,
		}
		ctx.JSON(code, body)
	}
}
