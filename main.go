package main

import (
	"image/jpeg"
	"io/fs"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/nfnt/resize"
)

func main() {
	s := time.Now()

	godotenv.Load()

	fromFolder := os.Getenv("TARGET_FOLDER_PATH")

	pResFold := os.Getenv("PARALLEL_RESULT_FOLDER_PATH")
	// increase goroutines limit if u want
	const gorouLimit = 3
	// set false if u want run parallel resize without goroutines limit
	const limit = true
	ResizeParallel(fromFolder, pResFold, limit, gorouLimit)

	lResFold := os.Getenv("LINEAR_RESULT_FOLDER_PATH")
	ResizeLinear(fromFolder, lResFold)

	log.Printf("totalTime: %s", time.Since(s))
}

func ResizeParallel(fromFolder string, toFolder string, limit bool, gorouLimit int) {
	s := time.Now()

	files, err := os.ReadDir(fromFolder)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(files))

	if !limit {
		gorouLimit = len(files)
	}

	// goroutines limiter
	guard := make(chan struct{}, gorouLimit)

	for _, f := range files {

		guard <- struct{}{}

		go func(file fs.DirEntry, wg *sync.WaitGroup) {
			defer func() {
				wg.Done()
				<-guard
			}()

			// open image
			opened, err := os.Open(fromFolder + file.Name())
			if err != nil {
				log.Fatal(err)
			}

			// decode jpeg into image.Image
			img, err := jpeg.Decode(opened)
			if err != nil {
				log.Fatal(err)
			}
			opened.Close()

			// resize (u can select another options)
			m := resize.Thumbnail(300, 200, img, resize.Lanczos3)

			out, err := os.Create(toFolder + "r_" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			// write new image to file
			jpeg.Encode(out, m, nil)

			log.Printf("%s", file.Name())
		}(f, &wg)
	}
	wg.Wait()

	log.Printf("Parallel resizing time: %s", time.Since(s))
}

func ResizeLinear(fromFolder string, toFolder string) {
	s := time.Now()

	files, err := os.ReadDir(fromFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// open image
		opened, err := os.Open(fromFolder + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		// decode jpeg into image.Image
		img, err := jpeg.Decode(opened)
		if err != nil {
			log.Fatal(err)
		}
		opened.Close()

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Thumbnail(300, 200, img, resize.Lanczos3)

		out, err := os.Create(toFolder + "r_" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)

		log.Printf("%s", file.Name())
	}

	log.Printf("Linear resizing time: %s", time.Since(s))
}
