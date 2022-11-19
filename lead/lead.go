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

func GetLeads(c *fiber.Ctx) {
	db := database.DB

	var leads []Lead

	db.Find(&leads)

	err := c.JSON(leads)
	if err != nil {
		return
	}
}
func GetLead(c *fiber.Ctx) {
	id := c.Params("id")

	db := database.DB

	var lead Lead

	db.Find(&lead, id)

	err := c.JSON(lead)
	if err != nil {
		return
	}
}
func CreateLead(c *fiber.Ctx) {
	db := database.DB

	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(http.StatusBadRequest).JSON(err)
		return
	}

	db.Create(&lead)

	err := c.JSON(lead)
	if err != nil {
		return
	}
}
func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB

	var lead Lead

	db.Find(&lead, id)

	if lead.Name == "" {
		c.Status(http.StatusNotFound).JSON("Lead Not Found.")
		return
	}

	db.Delete(&lead)

	c.Status(http.StatusAccepted).JSON("Lead successfully deleated")
}
