package examples

import (
	"log"
	"os"
	"sync"
)

// ImageFile reads an image from infile and writes
// a thumbnail-size version of it in the same directory.
// It returns the generated file name, e.g., "foo.thumb.jpg".
func ImageFile(infile string) (string, error) {
	return "", nil
}

// makeThumbnails3 makes thumbnails of the specified files in parallel.
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})

	for _, f := range filenames {
		go func() {
			ImageFile(f)
			ch <- struct{}{}
		}()
	}
	for range filenames {
		<-ch
	}
}

// makeThumbnails4 makes thumbnails of the specified files in parallel.
// It returns an error if any step failed
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func() {
			_, err := ImageFile(f)
			errors <- err
		}()
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: goroutine leak!
		}
	}
	return nil
}

// makeThumbnails5 makes thumbnails of the specified files in parallel.
// It returns the generated file names in an arbitrary order, or an error if any step failed
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func() {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}()
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

// makeThumbnails6 makes thumbnails of the specified files in parallel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames []string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines

	for _, f := range filenames {
		wg.Add(1)

		// worker
		go func() {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore errors
			sizes <- info.Size()
		}()
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes) // close channel to let the following for-range loop know when to stop
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
