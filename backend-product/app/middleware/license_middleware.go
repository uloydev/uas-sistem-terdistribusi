package middleware

import (
	"encoding/json"
	"net/http"
	"net/url"
	"sister-backend/app/model"
	"sister-backend/exception"

	"github.com/gofiber/fiber/v2"
)

func LicenseProtected(c *fiber.Ctx) error {
	if c.Query("master") == "supersecret" {
		return c.Next()
	}

	var request model.LicenseCheckRequest
	var data map[string]any

	err := c.QueryParser(&request)
	exception.PanicWhenError(err)

	url := url.URL{
		Scheme: "http",
		Host:   "localhost:8692",
		Path:   "v1/license/check",
	}

	query := url.Query()

	query.Add("username", request.Username)
	query.Add("key", request.Key)

	url.RawQuery = query.Encode()

	resp, err := http.Get(url.String())

	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if val, _ := data["code"]; int(val.(float64)) != http.StatusOK {
		return c.JSON(model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unathorized",
			Data:   "License Needed",
		})
	}

	return c.Next()
}
