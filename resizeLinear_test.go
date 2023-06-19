package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func BenchmarkResizeLinear(b *testing.B) {
	b.SetBytes(2)

	godotenv.Load()

	fromFolder := os.Getenv("TARGET_FOLDER_PATH")

	lResFold := os.Getenv("LINEAR_RESULT_FOLDER_PATH")

	ResizeLinear(fromFolder, lResFold)
}

func TestResizeLinear(t *testing.T) {
	godotenv.Load()

	fromFolder := os.Getenv("TARGET_FOLDER_PATH")

	lResFold := os.Getenv("LINEAR_RESULT_FOLDER_PATH")

	ResizeLinear(fromFolder, lResFold)
}
