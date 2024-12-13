package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/log"
	"server/internal/repository/postgres"
)

// GetFormulaCommits
// @Tags         formula
// @Summary      Retrieve commits for a specific formula
// @Description  Fetches all commits associated with a given formula ID.
// @Accept       json
// @Produce      json
// @Param        id path int true "Formula ID"
// @Success      200 {array} entities.FormulaHistory "List of commits for the specified formula"
// @Failure      400 {object} entities.ErrorResponse "Invalid formula ID"
// @Failure      500 {object} entities.ErrorResponse "Internal server error"
// @Router       /auth/formula/id/{id}/commits [get]
// @Security ApiKeyAuth
func (h *Handler) GetFormulaCommits(c *fiber.Ctx) error {
	formulaID, err := c.ParamsInt("id")
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	h.logger.Debug().Msg("call postgres.DBGetFormulaCommits")
	commits, err := postgres.DBGetFormulaCommits(h.db, int64(formulaID))
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(commits)

}
