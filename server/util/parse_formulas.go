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

	if ext == ".tex" || ext == ".docx" {
	} else {
		return nil, errors.New("invalid file extension")
	}

	if err := c.SaveFile(file, fmt.Sprintf("./tmp/%s", file.Filename)); err != nil {
		return nil, err
	}

	if ext == ".docx" {
		docxFile := fmt.Sprintf("./tmp/%s", file.Filename)
		texFile := fmt.Sprintf("./tmp/%s", file.Filename)
		texFile = strings.Replace(texFile, ".docx", ".tex", 1)
		cmd := exec.Command("pandoc", "-i", docxFile, "-o", texFile)

		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		err := cmd.Run()
		if err != nil {
			return nil, err
		}

		err = os.Remove("./tmp/" + file.Filename)
		file.Filename = strings.Replace(file.Filename, ".docx", ".tex", 1)
	}

	filePath := fmt.Sprintf("tmp/%v", file.Filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	mathRegex := regexp.MustCompile(`(\$.*?\$|\\\[.*?\\\]|\\mathcal\{.*?\})`)

	output := strings.ReplaceAll(string(content), "\n", "")
	output = strings.Join(strings.Fields(output), " ")

	matches := mathRegex.FindAllStringSubmatch(output, -1)

	for _, match := range matches {
		formulas = append(formulas, entities.GetFormulaFromArticleResponse{Formula: match[1]})
	}
	err = os.Remove("./tmp/" + file.Filename)

	return formulas, nil
}
