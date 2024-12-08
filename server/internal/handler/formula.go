package handler

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"server/internal/entities"
	"server/internal/log"
	"strings"
)

// GetFormulaFromArticle
// @Tags         formula
// @Summary      Get formulas from article
// @Accept       mpfd
// @Produce      json
// @Param file formData file true "Upload file"
// @Success      200 {object} []entities.GetFormulaFromArticleResponse "User  successfully logged in"
// @Failure      400 {object} entities.ErrorResponse "Invalid email or password"
// @Failure      500 {object} entities.ErrorResponse "Internal server error"
// @Router       /formula/file [post]
func (h *Handler) GetFormulaFromArticle(c *fiber.Ctx) error {
	dirName := "tmp"
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	}

	var formulas []entities.GetFormulaFromArticleResponse
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
		ext := filepath.Ext(file.Filename)

		if ext != ".tex" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid file extension"})
		}

		if err = c.SaveFile(file, fmt.Sprintf("./tmp/%s", file.Filename)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		filePath := fmt.Sprintf("tmp/%v", file.Filename)
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Ошибка при чтении файла:", err)
			return nil
		}

		mathRegex := regexp.MustCompile(`(\$.*?\$|\\\[.*?\\\]|\\mathcal\{.*?\})`)

		output := strings.ReplaceAll(string(content), "\n", "")
		output = strings.Join(strings.Fields(output), " ")

		matches := mathRegex.FindAllStringSubmatch(output, -1)

		for _, match := range matches {
			formulas = append(formulas, entities.GetFormulaFromArticleResponse{Formula: match[1]})
		}

		err = os.Remove("./tmp/" + file.Filename)
	}

	return c.Status(200).JSON(fiber.Map{"formulas": formulas})
}
