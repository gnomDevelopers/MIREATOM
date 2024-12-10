package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"server/internal/entities"
	"server/internal/log"
	"server/internal/repository/postgres"
)

// CreateArticle
// @Tags         article
// @Summary      Create article
// @Accept       mpfd
// @Produce      json
// @Param        title    formData string true  "Article title"
// @Param        science  formData string true  "Science field"
// @Param        section  formData string true  "Article section"
// @Param        file     formData file   true  "Upload file"
// @Success      200 {object} entities.Message
// @Failure      400 {object} entities.ErrorResponse "Invalid email or password"
// @Failure      500 {object} entities.ErrorResponse "Internal server error"
// @Router       /auth/article [post]
// @Security ApiKeyAuth
func (h *Handler) CreateArticle(c *fiber.Ctx) error {
	dirName := "articles"
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0755)
		if err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
				Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
			logEvent.Msg(err.Error())
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	}

	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}

	h.logger.Debug().Msg("call postgres.DBUserGetById")
	user, err := postgres.DBUserGetById(h.db, int64(userId))
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var article entities.Article

	article.UserId = userId
	article.Title = c.FormValue("title")
	article.Science = c.FormValue("science")
	article.Section = c.FormValue("section")

	userDirName := fmt.Sprintf("articles/%s", user.Surname+" "+user.Name+" "+user.ThirdName)
	article.Path = ""
	if _, err := os.Stat(userDirName); os.IsNotExist(err) {
		err = os.Mkdir(userDirName, 0755)
		if err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
				Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
			logEvent.Msg(err.Error())
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
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

	var articleDB *entities.Article
	for _, file := range files {
		fullPath := fmt.Sprintf("./%s/%s", userDirName, file.Filename)
		if err = c.SaveFile(file, fullPath); err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
				Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
			logEvent.Msg(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		article.Path = fullPath

		h.logger.Debug().Msg("call postgres.DBArticleCreate")
		articleDB, err = postgres.DBArticleCreate(h.db, &article)
		if err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
				Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
			logEvent.Msg(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(articleDB)
}
