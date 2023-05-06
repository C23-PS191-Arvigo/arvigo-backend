package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yusufwib/arvigo-backend/repository"
	"github.com/yusufwib/arvigo-backend/utils"
)

func RegisterLocationRoutes(e *echo.Echo) {
	v1Group := e.Group("/v1")
	authGroup := v1Group.Group("/location")
	authGroup.GET("/provinces", getProvince)
	authGroup.GET("/cities", getCity)
	authGroup.GET("/districts", getDistricts)
	authGroup.GET("/subdistricts", getSubDistricts)
	authGroup.GET("/postal_codes", getPostalCode)
}

func getProvince(c echo.Context) error {
	data, statusCode, err := repository.GetProvinces()
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, http.StatusOK)
}

func getCity(c echo.Context) error {
	provinceID := utils.StrToUint64(c.QueryParam("province_id"), 0)

	data, statusCode, err := repository.GetCities(provinceID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, http.StatusOK)
}

func getDistricts(c echo.Context) error {
	cityID := utils.StrToUint64(c.QueryParam("city_id"), 0)

	data, statusCode, err := repository.GetCities(cityID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, http.StatusOK)
}

func getSubDistricts(c echo.Context) error {
	districtID := utils.StrToUint64(c.QueryParam("district_id"), 0)

	data, statusCode, err := repository.GetCities(districtID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, http.StatusOK)
}

func getPostalCode(c echo.Context) error {
	subDistrictID := utils.StrToUint64(c.QueryParam("subdistrict_id"), 0)

	data, statusCode, err := repository.GetCities(subDistrictID)
	if err != nil {
		return utils.ResponseJSON(c, err.Error(), data, statusCode)
	}

	return utils.ResponseJSON(c, "Success", data, http.StatusOK)
}
