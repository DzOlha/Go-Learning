package __parallelize

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"runtime"
	"sync"
	"time"
)

/*
	Take an existing CPU-bound sequential program, such as the Mandelbrot program of Section 3.3
	or the 3-D surface computation of Section 3.2, and execute its main loop in parallel
	using channels for communication. How much faster does it run on a multiprocessor machine?
	What is the optimal number of goroutines to use?
*/

func RunMandelbrotSequential() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	//png.Encode(os.Stdout, img)
}

func RunMandelbrotParallel() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup

	workers := runtime.NumCPU()
	rowChannel := make(chan int, height)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for py := range rowChannel {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					img.Set(px, py, mandelbrot(z))
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(rowChannel)
	}()

	for py := 0; py < height; py++ {
		rowChannel <- py
	}

	//png.Encode(os.Stdout, img)
}

func RunMandelbrotToCompare() {
	start := time.Now()
	RunMandelbrotSequential()
	fmt.Println("Sequential:", time.Since(start))

	start = time.Now()
	RunMandelbrotParallel()
	fmt.Println("Parallel:", time.Since(start))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
