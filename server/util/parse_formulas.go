package util

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"server/internal/entities"
	"strings"
)

func ParseFormulasFromFile(c *fiber.Ctx, file *multipart.FileHeader) ([]entities.GetFormulaFromArticleResponse, error) {
	var formulas []entities.GetFormulaFromArticleResponse
	ext := filepath.Ext(file.Filename)

	// Проверяем, что файл имеет допустимое расширение (.tex или .docx)
	if ext == ".tex" || ext == ".docx" {
	} else {
		return nil, errors.New("invalid file extension")
	}

	// Сохраняем файл во временную директорию
	if err := c.SaveFile(file, fmt.Sprintf("./tmp/%s", file.Filename)); err != nil {
		return nil, err // Возвращаем ошибку, если файл не удалось сохранить
	}

	// Если файл в формате .docx, конвертируем его в .tex с использованием pandoc
	if ext == ".docx" {
		docxFile := fmt.Sprintf("./tmp/%s", file.Filename)     // Путь к исходному файлу .docx
		texFile := fmt.Sprintf("./tmp/%s", file.Filename)      // Путь к будущему файлу .tex
		texFile = strings.Replace(texFile, ".docx", ".tex", 1) // Меняем расширение на .tex

		// Команда для конвертации .docx в .tex с использованием pandoc
		cmd := exec.Command("pandoc", "-i", docxFile, "-o", texFile)

		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		err := cmd.Run()
		if err != nil {
			return nil, err
		}

		// Удаляем исходный .docx файл после конвертации
		err = os.Remove("./tmp/" + file.Filename)

		// Обновляем имя файла на .tex
		file.Filename = strings.Replace(file.Filename, ".docx", ".tex", 1)
	}

	filePath := fmt.Sprintf("tmp/%v", file.Filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Регулярное выражение для поиска математических формул
	mathRegex := regexp.MustCompile(`(\$.*?\$|\\\[.*?\\\]|\\mathcal\{.*?\})`)

	// Удаляем переводы строк и объединяем текст в одну строку
	output := strings.ReplaceAll(string(content), "\n", "")
	output = strings.Join(strings.Fields(output), " ")

	matches := mathRegex.FindAllStringSubmatch(output, -1)

	// Добавляем найденные формулы в массив результатов
	for _, match := range matches {
		formulas = append(formulas, entities.GetFormulaFromArticleResponse{Formula: match[1]})
	}

	// Удаляем временный файл
	err = os.Remove("./tmp/" + file.Filename)

	return formulas, nil
}
