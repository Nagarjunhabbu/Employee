package controller

import (
	"employee/internal/model"
	"employee/internal/service"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strconv"
)

type Controller interface {
	GetEmployees(c echo.Context) error
	CreateEmployee(c echo.Context) error
	UpdateEmployee(c echo.Context) error
	DeleteEmployee(c echo.Context) error
	ListEmployees(c echo.Context) error
}

type EmployeeController struct {
	Service service.EmployeeService
}

func (e EmployeeController) ListEmployees(c echo.Context) error {
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	pageSize := c.QueryParam("page_size")
	if pageSize == "" {
		pageSize = "20"
	}
	pageNum, _ := strconv.Atoi(page)
	pageSizeNum, _ := strconv.Atoi(pageSize)
	resp, err := e.Service.ListEmployee(c.Request().Context(), pageNum, pageSizeNum)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (e EmployeeController) GetEmployees(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	resp, err := e.Service.GetEmployee(c.Request().Context(), idInt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (e EmployeeController) CreateEmployee(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	m := model.Employee{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	m, err = e.Service.CreateEmployee(c.Request().Context(), m)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, m)
}

func (e EmployeeController) UpdateEmployee(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	m := model.Employee{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	m, err = e.Service.UpdateEmployee(c.Request().Context(), idInt, m)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, m)
}

func (e EmployeeController) DeleteEmployee(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = e.Service.DeleteEmployee(c.Request().Context(), idInt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func NewEmployeeController(employeeService service.EmployeeService) Controller {
	return EmployeeController{
		Service: employeeService,
	}
}
