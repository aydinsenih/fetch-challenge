package main

import (
	"fetch/database"
	"fetch/receipt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreatedResponse struct {
	Id uuid.UUID `json:"id"`
}

type PointResponse struct {
	Points int `json:"points"`
}

func main() {
	validate := validator.New()
	e := echo.New()

	//group APIs
	g := e.Group("receipts")

	g.GET("/:id/points", func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "id could not be parsed.")
		}
		point, ok := database.Read(id)
		if !ok {
			return c.JSON(http.StatusBadRequest, "No receipt found for that id")
		}
		res := PointResponse{
			Points: point,
		}
		return c.JSON(http.StatusOK, res)
	})

	g.POST("/process", func(c echo.Context) error {
		r := receipt.Receipt{}

		//bind
		if err := c.Bind(&r); err != nil {
			return c.JSON(http.StatusBadRequest, "The receipt is invalid")
		}

		//validate
		if err := validate.Struct(&r); err != nil {
			return c.JSON(http.StatusBadRequest, "The receipt is invalid")
		}

		newID := uuid.New()

		of := CreatedResponse{
			Id: newID,
		}
		// point calculation
		point := r.Calc()

		//write points to DB
		database.Write(newID, point)

		return c.JSON(http.StatusCreated, of)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
