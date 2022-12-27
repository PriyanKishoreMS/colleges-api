package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/time/rate"
	"github.com/PriyanKishoreMS/colleges-list-api/config"
	"github.com/PriyanKishoreMS/colleges-list-api/entities"
	"github.com/gofiber/fiber/v2"
)

type APIhandler struct{
	rateLimiter *rate.Limiter
}
func NewAPIhandler() *APIhandler {
	return &APIhandler{
		//bucket size 10 and refill rate is 3/sec
		rateLimiter: rate.NewLimiter(rate.Every(time.Second)*3,10),
	}
}

func(handlerObj *APIhandler) GetAllStates(c *fiber.Ctx) error {

	err := handlerObj.rateLimiter.Wait(c.Context())
	if  err != nil{
		return err
	}

	var states []string

	// sort alphabetically
	result := config.Db.Model(&entities.College{}).Distinct("state").Order("state").Find(&states)

	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "States not found",
		})
	}

	return c.Status(http.StatusOK).JSON(states)
}

func(handlerObj *APIhandler) GetDistrictsByState(c *fiber.Ctx) error {

	err := handlerObj.rateLimiter.Wait(c.Context())
	if  err != nil{
		return err
	}

	state := c.Params("state")
	var districts []string

	state = strings.ReplaceAll(state, "%20", " ")
	result := config.Db.Model(&entities.College{}).Distinct("district").Where("state = ?", state).Order("district").Find(&districts)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Districts not found",
		})
	}

	return c.Status(http.StatusOK).JSON(districts)
}

func(handlerObj *APIhandler) GetAllCollegesInState(c *fiber.Ctx) error {

	err := handlerObj.rateLimiter.Wait(c.Context())
	if  err != nil{
		return err
	}

	state := c.Params("state")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search")

	if page < 1 {
		page = 1
	}

	var colleges []entities.College
	state = strings.ReplaceAll(state, "%20", " ")

	result := config.Db.Where("state = ?", state).Limit(limit).Order("name").Offset((page - 1) * limit).Find(&colleges)

	var total int64
	config.Db.Model(&entities.College{}).Where("state = ?", state).Count(&total)
	totalPages := int(total) / limit

	if search != "" {
		result = config.Db.Where("state = ? AND name LIKE ?", state, search+"%").Limit(limit).Order("name").Offset((page - 1) * limit).Find(&colleges)

		config.Db.Model(&entities.College{}).Where("state = ? AND name LIKE ?", state, search+"%").Count(&total)
		totalPages = int(total) / limit

	}
	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "College not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"count":       total,
		"currentPage": page,
		"pages":       totalPages + 1,
		"colleges":    colleges,
	})
}

func (handlerObj *APIhandler) GetAllCollegesInDistrict(c *fiber.Ctx) error {

	err := handlerObj.rateLimiter.Wait(c.Context())
	if  err != nil{
		return err
	}	

	state := c.Params("state")
	district := c.Params("district")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search")

	if page < 1 {
		page = 1
	}

	var colleges []entities.College
	state = strings.ReplaceAll(state, "%20", " ")
	district = strings.ReplaceAll(district, "%20", " ")

	result := config.Db.Where("state = ? AND district = ?", state, district).Limit(limit).Order("name").Offset((page - 1) * limit).Find(&colleges)

	var total int64
	config.Db.Model(&entities.College{}).Where("state = ? AND district = ?", state, district).Count(&total)
	totalPages := int(total) / limit

	if search != "" {
		result = config.Db.Where("state = ? AND district = ? AND name LIKE ?", state, district, search+"%").Limit(limit).Order("name").Offset((page - 1) * limit).Find(&colleges)

		config.Db.Model(&entities.College{}).Where("state = ? AND district = ? AND name LIKE ?", state, district, search+"%").Count(&total)
		totalPages = int(total) / limit
	}

	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "College not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"count":       total,
		"currentPage": page,
		"pages":       totalPages + 1,
		"colleges":    colleges,
	})

}

func (handlerObj *APIhandler) SearchCollege(c *fiber.Ctx) error {

	err := handlerObj.rateLimiter.Wait(c.Context())
	if  err != nil{
		return err
	}	

	search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}

	var colleges []entities.College

	result := config.Db.Where("name LIKE ?", search+"%").Order("name").Limit(limit).Offset((page - 1) * limit).Find(&colleges)

	var total int64
	config.Db.Model(&entities.College{}).Where("name LIKE ?", search+"%").Count(&total)
	totalPages := int(total) / limit

	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "College not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"count":       total,
		"currentPage": page,
		"pages":       totalPages + 1,
		"colleges":    colleges,
	})
}
