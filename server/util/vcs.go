package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"math/rand"
	"time"
)

type Change struct {
	Type    string `json:"type"`    // Тип изменения: "insert", "delete", "equal"
	Content string `json:"content"` // Содержимое изменения
	Start   int    `json:"start"`   // Начальная позиция (только для insert/delete)
	End     int    `json:"end"`     // Конечная позиция (только для insert/delete)
}

func GenerateHash(content string) string {
	hash := sha256.Sum256([]byte(content))
	return hex.EncodeToString(hash[:])
}

func GenerateName() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%x", rand.Int63())
}

func CompareStrings(oldContent, newContent string) (string, error) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(oldContent, newContent, false)

	var changes []Change
	oldIndex := 0
	newIndex := 0

	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			changes = append(changes, Change{
				Type:    "insert",
				Content: diff.Text,
				Start:   newIndex,
				End:     newIndex + len(diff.Text),
			})
			newIndex += len(diff.Text)
		case diffmatchpatch.DiffDelete:
			changes = append(changes, Change{
				Type:    "delete",
				Content: diff.Text,
				Start:   oldIndex,
				End:     oldIndex + len(diff.Text),
			})
			oldIndex += len(diff.Text)
		case diffmatchpatch.DiffEqual:
			changes = append(changes, Change{
				Type:    "equal",
				Content: diff.Text,
				Start:   oldIndex,
				End:     oldIndex + len(diff.Text),
			})
			oldIndex += len(diff.Text)
			newIndex += len(diff.Text)
		}
	}

	// Преобразуем массив изменений в строку JSON
	jsonData, err := json.Marshal(changes)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
