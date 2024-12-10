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
	"server/internal/repository/postgres"
	"strconv"
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

		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(dir)

		docxFile := "C:/Users/danii/OneDrive/Desktop/Go/MIREATOM/server/tmp/formula.docx"

		cmd := exec.Command("pandoc", "-i", docxFile, "-o", "C:/Users/danii/OneDrive/Desktop/Go/MIREATOM/server/tmp/formula.tex")

		// Выполните команду и сохраните вывод
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr

		// Запустите команду
		err = cmd.Run()
		if err != nil {
			fmt.Printf("Ошибка выполнения команды: %s\n", err)
			fmt.Printf("Вывод ошибки: %s\n", stderr.String())
		}

		// Выведите вывод команды
		fmt.Println("Вывод команды:")
		fmt.Println(out.String())

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

// CreateFormula
// @Tags formula
// @Summary      Create formula
// @Accept       json
// @Produce      json
// @Param data body entities.CreateFormulaRequest true "formula data"
// @Success 200 {object} entities.CreateFormulaResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /formula [post]
func (h *Handler) CreateFormula(c *fiber.Ctx) error {
	var formula entities.Formula
	err := c.BodyParser(&formula)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call postgres.DBFormulaCreate")
	formulaDB, err := postgres.DBFormulaCreate(h.db, &formula)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(entities.CreateFormulaResponse{ID: formulaDB.ID})
}

// GetFormulaById
// @Tags formula
// @Summary      Get formula by id
// @Accept       json
// @Produce      json
// @Param id path string true "formula id"
// @Success 200 {object} entities.Formula
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /formula/id/{id} [get]
func (h *Handler) GetFormulaById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call postgres.DBFormulaGetByID")
	formulaDB, err := postgres.DBFormulaGetByID(h.db, int64(id))
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(formulaDB)
}

// GetFormulaByUserId
// @Tags formula
// @Summary      Get formula by user id
// @Accept       json
// @Produce      json
// @Param id path string true "user id"
// @Success 200 {object} entities.Formula
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /formula/user_id/{id} [get]
func (h *Handler) GetFormulaByUserId(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call postgres.DBFormulaGetByUserID")
	formulaDB, err := postgres.DBFormulaGetByUserID(h.db, int64(id))
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(formulaDB)
}

// UpdateFormula
// @Tags formula
// @Summary      Update formula
// @Accept       json
// @Produce      json
// @Param data body entities.UpdateFormulaRequest true "formula data"
// @Success 200 {object} entities.Message
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /formula [put]
func (h *Handler) UpdateFormula(c *fiber.Ctx) error {
	var formula entities.UpdateFormulaRequest
	err := c.BodyParser(&formula)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	h.logger.Debug().Msg("call postgres.DBFormulaUpdate")
	err = postgres.DBFormulaUpdate(h.db, &formula)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

// DeleteFormula
// @Tags formula
// @Summary      Delete formula
// @Accept       json
// @Produce      json
// @Param id path string true "formula id"
// @Success 200 {object} entities.Message
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /formula/id/{id} [delete]
func (h *Handler) DeleteFormula(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call postgres.DBFormulaDelete")
	err = postgres.DBFormulaDelete(h.db, int64(id))
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

// GetFormulasHistory
// @Tags formula
// @Summary      Retrieve formula history by user ID
// @Description  Returns a paginated list of formulas for a specific user
// @Accept       json
// @Produce      json
// @Param        id     path      int  true  "User ID"
// @Param        number path      int  true  "Page number"
// @Success      200    {array}   entities.Formula        "List of formulas"
// @Failure      400    {object}  entities.ErrorResponse  "Invalid ID or page number"
// @Failure      500    {object}  entities.ErrorResponse  "Internal server error"
// @Router       /formula/history/user/{id}/page/{number} [get]
func (h *Handler) GetFormulasHistory(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	numberStr := c.Params("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call postgres.DBFormulaGetAll")
	formulas, err := postgres.DBFormulaHistoryGet(h.db, int64(id), int64(number))
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(formulas)
}
