package lead

import (
	"github.com/dchung117/crm_golang_fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// define lead type gorm Model
type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

// route handles
func GetLeads(ctx *fiber.Ctx) {
	// get database connection
	db := database.DBConn

	// define slice of leads
	var leads []Lead

	// find all leads
	db.Find(&leads)

	// serve leads to client
	ctx.JSON(leads)
}

func GetLead(ctx *fiber.Ctx) {
	// get lead id
	id := ctx.Params("id")

	// get database connection
	db := database.DBConn

	// find lead w/ ID
	var lead Lead
	db.Find(&lead, id)

	// server lead back to client as JSON
	ctx.JSON(lead)
}

func NewLead(ctx *fiber.Ctx) {
	// get db connection
	db := database.DBConn

	// define new lead
	lead := new(Lead)

	// bind lead info in request body to struct, raise error
	if err := ctx.BodyParser(lead); err != nil {
		ctx.Status(503).Send(err)
		return
	}

	// post new lead to db
	db.Create(&lead)

	// sever new lead to client
	ctx.JSON(lead)

}

func DeleteLead(ctx *fiber.Ctx) {
	// get db connection
	db := database.DBConn

	// get lead ID from params
	id := ctx.Params("id")

	// define new lead
	var lead Lead

	// find lead info in db
	db.First(&lead, id)

	// send 500 error if no lead w/ ID found
	if lead.Name == "" {
		ctx.Status(500).Send("No lead found w/ given ID.")
		return
	}

	// delete lead w/ ID
	db.Delete(&lead)

	// return succesful deletion message back to client
	ctx.Send("Lead successfully deleted")
}
