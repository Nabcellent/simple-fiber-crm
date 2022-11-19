package lead

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"simple-fiber-crm/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Phone   int    `json:"phone"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DB

	var leads []Lead

	db.Find(&leads)

	return c.JSON(leads)
}
func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var lead Lead

	db.Find(&lead, id)

	return c.JSON(lead)
}
func CreateLead(c *fiber.Ctx) error {
	db := database.DB

	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	db.Create(&lead)

	return c.JSON(lead)
}
func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var lead Lead

	db.Find(&lead, id)

	if lead.Name == "" {
		return c.Status(http.StatusNotFound).JSON("Lead Not Found.")
	}

	db.Delete(&lead)

	return c.Status(http.StatusAccepted).JSON("Lead successfully deleated")
}
