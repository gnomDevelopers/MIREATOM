package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"math/rand"
	"time"
)

func GenerateHash(content string) string {
	hash := sha256.Sum256([]byte(content))
	return hex.EncodeToString(hash[:])
}

func GenerateName() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%x", rand.Int63())
}

func CalculateDiff(oldContent, newContent string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(oldContent, newContent, false)
	return dmp.DiffPrettyText(diffs)
}
