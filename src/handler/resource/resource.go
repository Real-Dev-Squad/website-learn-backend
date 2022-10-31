package resourceHandler

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	resource "github.com/Real-Dev-Squad/gopher-cloud-service/src/database"
	resourceSchema "github.com/Real-Dev-Squad/gopher-cloud-service/src/schema/resource"
	metaInfoUtil "github.com/Real-Dev-Squad/gopher-cloud-service/src/utils/metaInfo"
)

func GetResource(c *fiber.Ctx) error {

	var id string = c.Params("id")

	resource, err := resource.GetResource(id)

	if err != nil {
		log.Printf("Failed to get resource: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.ErrNotFound.Code,
			"error":   fiber.ErrNotFound.Message,
			"message": "Resource doesn't exist",
		})
	}

	return c.Status(fiber.StatusOK).JSON(resource)

}

func GetResources(c *fiber.Ctx) error {
	limit := c.Query("limit")
	startAfter := c.Query("after")

	var dataLimit int = 10
	if limit != "" {
		var err error
		dataLimit, err = strconv.Atoi(limit)
		if err != nil {
			dataLimit = 10
		}
	}
	resources, err := resource.GetResources(dataLimit, startAfter)
	if err != nil {
		log.Printf("Failed to get all resource: %v", err)
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(resources)
}

func PostResource(c *fiber.Ctx) error {

	newResource := new(resourceSchema.Resources)
	c.BodyParser(&newResource)
	newResource = processResource(*newResource)

	data, err := resource.PostResource(*newResource)
	if err != nil {
		log.Printf("Failed to create new resource: %v", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(data)

}

func processResource(resource resourceSchema.Resources) *resourceSchema.Resources {

	metaInfo, err := metaInfoUtil.GetMetaInfo(resource.Link)

	if err == nil {
		if resource.Title == "" {
			resource.Title = metaInfo.Title
		}
		if resource.Body == "" {
			resource.Body = metaInfo.Description
		}
		resource.Type = metaInfo.Type
		resource.Image = metaInfo.Image
	}
	resource.PublishedDate = time.Now()
	if resource.Filters == nil {
		resource.Filters = make([]string, 0)
	}

	return &resource
}
