package http_v1

import (
	"inforce-task/internal/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/nordew/go-errx"
)

func (h *Handler) initReribleRoutes(app *fiber.App) {
	app.Get("/ownership", responseWrapper(h.GetOwnership))
	app.Post("/traits", responseWrapper(h.GetTraits))
}

// @Summary Get Ownership
// @Tags rarible
// @Description Returns NFT ownerships for a given wallet address.
// @Accept json
// @Produce json
// @Param id query string true "Wallet Address"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /ownership [get]
func (h *Handler) GetOwnership(c *fiber.Ctx) (map[string]any, error) {
	id := c.Query("id")

	data, err := h.svc.GetOwnerships(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// @Summary Get Collection Traits
// @Tags rarible
// @Description Returns traits of items in a specific NFT collection.
// @Accept json
// @Produce json
// @Param input body dto.CollectionTrait true "Collection and trait filters"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /traits [post]
func (h *Handler) GetTraits(c *fiber.Ctx) (map[string]any, error) {
	var dto dto.CollectionTrait
	if err := c.BodyParser(&dto); err != nil {
		return nil, errx.NewBadRequest().WithDescriptionAndCause("failed to parse request body", err)
	}

	data, err := h.svc.GetTraits(dto)
	if err != nil {
		return nil, err
	}

	return data, nil
}
