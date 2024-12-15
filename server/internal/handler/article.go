package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
	"server/internal/entities"
	"server/internal/log"
	"server/internal/repository/postgres"
	"server/util"
	"strconv"
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
	// Создание директории для хранения статей, если она не существует
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

	// Получение ID пользователя из контекста
	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}

	// Получение данных о пользователе из базы данных
	h.logger.Debug().Msg("call postgres.DBUserGetById")
	user, err := postgres.DBUserGetById(h.db, int64(userId))
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Инициализация объекта статьи
	var article entities.Article
	article.UserId = userId
	article.Title = c.FormValue("title")
	article.Science = c.FormValue("science")
	article.Section = c.FormValue("section")

	// Формирование пути для директории пользователя
	userDirName := fmt.Sprintf("articles/%s", user.Surname+" "+user.Name+" "+user.ThirdName)
	article.Path = ""

	// Проверка существования статьи с таким же названием у пользователя
	h.logger.Debug().Msg("call postgres.DBArticleExists")
	exists, err := postgres.DBArticleExists(h.db, article.Title, article.UserId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if exists {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg("article already exists")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "article already exists"})
	}

	// Создание директории пользователя, если она не существует
	if _, err := os.Stat(userDirName); os.IsNotExist(err) {
		err = os.Mkdir(userDirName, 0755)
		if err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
				Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
			logEvent.Msg(err.Error())
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	}

	// Чтение файлов из формы
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
		// Сохранение файла в директории пользователя
		fullPath := fmt.Sprintf("./%s/%s", userDirName, file.Filename)
		if err = c.SaveFile(file, fullPath); err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
				Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
			logEvent.Msg(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		article.Path = fullPath

		// Создание записи о статье в базе данных
		h.logger.Debug().Msg("call postgres.DBArticleCreate")
		articleDB, err = postgres.DBArticleCreate(h.db, &article)
		if err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
				Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
			logEvent.Msg(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Сохранение формул, если файл в формате .tex или .docx
		var formulas []entities.GetFormulaFromArticleResponse
		ext := filepath.Ext(file.Filename)
		if ext == ".tex" || ext == ".docx" {
			formulas, err = util.ParseFormulasFromFile(c, file)
			if err != nil {
				logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
					Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
				logEvent.Msg(err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}

			for _, formula := range formulas {
				var formulaDb entities.Formula
				formulaDb.Title = ""
				formulaDb.Value = formula.Formula
				formulaDb.UserID = userId

				// Проверка существования формулы
				exists, err := postgres.DBFormulaExists(h.db, formulaDb.Value)
				if err != nil {
					logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
						Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
					logEvent.Msg(err.Error())
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
				}
				if exists {
					continue
				}

				// Сохранение формулы в базе данных
				h.logger.Debug().Msg("call postgres.DBFormulaCreate")
				_, err = postgres.DBFormulaCreate(h.db, &formulaDb)
				if err != nil {
					logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
						Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
					logEvent.Msg(err.Error())
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
				}
			}
		}
	}

	// Успешное завершение запроса
	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(articleDB)
}

// GetAllArticles
// @Tags article
// @Summary      Get all articles
// @Accept       json
// @Produce      json
// @Success 200 {object} []entities.ArticleInfo
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /article [get]
func (h *Handler) GetAllArticles(c *fiber.Ctx) error {
	// Получаем все статьи из бд
	h.logger.Debug().Msg("call postgres.DBArticleGetAll")
	articles, err := postgres.DBArticleGetAll(h.db)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Успешное завершение запроса
	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(articles)
}

// GetArticlesByUserId
// @Tags article
// @Summary      Get all articles
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success 200 {object} []entities.ArticleInfo
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /article/user_id/{id} [get]
func (h *Handler) GetArticlesByUserId(c *fiber.Ctx) error {
	// Получаем id пользователя из пути запроса
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Получаем все статьи пользователя из бд
	h.logger.Debug().Msg("call postgres.DBArticleGetAll")
	articles, err := postgres.DBArticleGetByUserId(h.db, id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Успешное завершение запроса
	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(articles)
}

// GetArticleFile
// @Tags article
// @Summary      Get articles file
// @Accept       json
// @Produce      octet-stream
// @Param        id   path      int  true  "Article ID"
// @Success 200 {file} file
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /article/file/{id} [get]
func (h *Handler) GetArticleFile(c *fiber.Ctx) error {
	// Получаем id статьи из пути запроса
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Получаем путь до файла статьи из бд
	h.logger.Debug().Msg("call postgres.DBArticleGetAll")
	path, err := postgres.DBArticleGetPath(h.db, id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Проверяем существует ли такой файл
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "File not found"})
	}

	// Успешное завершение запроса
	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.SendFile(path)
}

// UpdateArticle
// @Tags article
// @Summary      Update article
// @Accept       json
// @Produce      json
// @Param data body entities.UpdateArticleRequest true "formula data"
// @Success 200 {object} entities.Message
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/article [put]
// @Security ApiKeyAuth
func (h *Handler) UpdateArticle(c *fiber.Ctx) error {
	// Получаем статью в теле запроса
	var article entities.UpdateArticleRequest
	err := c.BodyParser(&article)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Err(err).Msg("invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	// Обновляем статью в бд
	h.logger.Debug().Msg("call postgres.DBArticleUpdatePath")
	err = postgres.DBArticleUpdate(h.db, &article)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Успешное завершение запроса
	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

// UpdateArticleFile
// @Tags article
// @Summary      Update article
// @Accept       json
// @Produce      json
// @Param        id    formData string true  "Article id"
// @Param        file     formData file   true  "Upload file"
// @Success 200 {object} entities.Message
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/article/file [put]
// @Security ApiKeyAuth
func (h *Handler) UpdateArticleFile(c *fiber.Ctx) error {
	// Получаем id пользователя из контекста
	userId, ok := c.Locals("id").(int)
	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}
	articleId, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid article_id",
		})
	}

	// Получаем информацию о пользователе из бд
	h.logger.Debug().Msg("call postgres.DBUserGetById")
	user, err := postgres.DBUserGetById(h.db, int64(userId))
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	userDirName := fmt.Sprintf("articles/%s", user.Surname+" "+user.Name+" "+user.ThirdName)

	// Получаем путь до статьи из бд
	h.logger.Debug().Msg("call postgres.DBArticleGetPath")
	path, err := postgres.DBArticleGetPath(h.db, articleId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	// Удаляем старый файл статьи
	err = os.Remove(path)
	if err != nil {
		if os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "File not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete file"})
	}

	// Получаем файл из запроса
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
		// Сохраняем новый файл
		fullPath := fmt.Sprintf("./%s/%s", userDirName, file.Filename)
		if err = c.SaveFile(file, fullPath); err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
				Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
			logEvent.Msg(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Сохрнаняем новый путь до файла в бд
		h.logger.Debug().Msg("call postgres.DBArticleUpdatePath")
		err = postgres.DBArticleUpdatePath(h.db, fullPath, articleId)
		if err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
				Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
			logEvent.Msg(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	// Успешное завершение запроса
	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

// DeleteArticle
// @Tags article
// @Summary      Delete article
// @Accept       json
// @Produce      json
// @Param id path string true "article id"
// @Success 200 {object} entities.Message
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /auth/article/id/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) DeleteArticle(c *fiber.Ctx) error {
	// Получаем id статьи из пути запроса
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Получаем путь до статьи из бд
	path, err := postgres.DBArticleGetPath(h.db, id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	// Удаляем файл статьи
	err = os.Remove(path)
	if err != nil {
		if os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "File not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete file"})
	}

	// Удаляем статью из бд
	h.logger.Debug().Msg("call postgres.DBArticleDelete")
	err = postgres.DBArticleDelete(h.db, int64(id))
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Успешное завершение запроса
	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}
