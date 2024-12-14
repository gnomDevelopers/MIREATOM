package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"path/filepath"
	"server/internal/config"
	"server/internal/entities"
	"server/internal/log"
	"server/internal/repository/postgres"
	"server/util"
	"strconv"
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
		formulas, err = util.ParseFormulasFromFile(c, file)
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
// @Router       /auth/formula [post]
// @Security ApiKeyAuth
func (h *Handler) CreateFormula(c *fiber.Ctx) error {
	var req entities.CreateFormulaRequest
	err := c.BodyParser(&req)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}

	h.logger.Debug().Msg("call postgres.DBFormulaExists")
	exists, err := postgres.DBFormulaExists(h.db, req.Value)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if exists {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("formula already exists")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "formula already exists"})
	}

	formula := &entities.Formula{
		UserID: userId,
		Title:  req.Title,
		Value:  req.Value,
	}

	h.logger.Debug().Msg("call postgres.DBFormulaCreate")
	formulaDB, err := postgres.DBFormulaCreate(h.db, formula)
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
// @Router       /auth/formula [put]
// @Security ApiKeyAuth
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
// @Router       /auth/formula/id/{id} [delete]
// @Security ApiKeyAuth
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

// FormulaRecognize
// @Tags formula
// @Summary      Recognizes the formula from the image
// @Accept       multipart/form-data
// @Produce      json
// @Param file formData file true "Photo file"
// @Success 200 {object} entities.RecognizedFormula
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /formula/recognize [post]
func (h *Handler) FormulaRecognize(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("failed to retrieve file")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to retrieve file"})
	}

	if file.Header.Get("Content-Type") != "image/jpeg" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("only JPEG images are allowed")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "only JPEG images are allowed"})
	}

	saveDir := "./tmp"
	savePath := filepath.Join(saveDir, file.Filename)

	if err := c.SaveFile(file, savePath); err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg("failed to save file")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save file"})
	}

	payload := map[string]string{
		"type":    "1",
		"content": savePath,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg("failed to prepare request payload")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to prepare request payload"})
	}

	resp, err := http.Post(config.LlamaAPI, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg("failed to send request to external API")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to send request to external API"})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg("external API returned an error")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "external API returned an error"})
	}

	var apiResponse struct {
		Formula string `json:"formula"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg("failed to parse external API response")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to parse external API response"})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"formula": apiResponse.Formula})
}
