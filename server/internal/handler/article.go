package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"server/internal/log"
)

// CreateArticle
// @Tags         article
// @Summary      Create article
// @Accept       mpfd
// @Produce      json
// @Param file formData file true "Upload file"
// @Success      200 {object} entities.Message
// @Failure      400 {object} entities.ErrorResponse "Invalid email or password"
// @Failure      500 {object} entities.ErrorResponse "Internal server error"
// @Router       /auth/article [post]
// @Security ApiKeyAuth
func (h *Handler) CreateArticle(c *fiber.Ctx) error {
	dirName := "articles"
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	}

	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	files := form.File["file"]
	if len(files) == 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg("empty file")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "empty file"})
	}

	for _, file := range files {
		if err = c.SaveFile(file, fmt.Sprintf("./articles/%s", file.Filename)); err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
				Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
			logEvent.Msg(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}
