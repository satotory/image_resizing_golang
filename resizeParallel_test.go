package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func BenchmarkResizeParallel(b *testing.B) {
	b.SetBytes(2)

	godotenv.Load()

	fromFolder := os.Getenv("TARGET_FOLDER_PATH")

	pResFold := os.Getenv("PARALLEL_RESULT_FOLDER_PATH")
	// increase goroutines limit if u want
	const gorouLimit = 3
	// set false if u want run parallel resize without goroutines limit
	const limit = true

	for i := 0; i < b.N; i++ {
		ResizeParallel(fromFolder, pResFold, limit, gorouLimit)
	}
}

func TestResizeParallel(t *testing.T) {
	godotenv.Load()

	fromFolder := os.Getenv("TARGET_FOLDER_PATH")

	pResFold := os.Getenv("PARALLEL_RESULT_FOLDER_PATH")
	// increase goroutines limit if u want
	const gorouLimit = 3
	// set false if u want run parallel resize without goroutines limit
	const limit = true

	ResizeParallel(fromFolder, pResFold, limit, gorouLimit)
}
