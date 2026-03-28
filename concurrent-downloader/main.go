package main

import(
	"os"
	"path/filepath"
	"net/http"
	"time"
	"fmt"
	"io"
	"log"
	"sync"
)

func DownloadFile(url , destDir string) error {
	filename := filepath.Base(url)
	filePath := filepath.Join(destDir , filename)
	out , err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Println("Downloading",url)
	start := time.Now()

	resp , err := http.Get(url)
	if err != nil{
		out.Close()
		_ = os.Remove(filePath)
		return err
	}



	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK {
		out.Close()
		_ = os.Remove(filePath)
		return fmt.Errorf("Bad status %s",resp.Status)
	}

	_ , err = io.Copy(out , resp.Body)
if err != nil{
		return err
	}

fmt.Printf("Downloading %s took %s \n",filename, time.Since(start))
return nil
}

type Result struct {
	URL string
	FileName string
	Size int64
	Duration time.Duration
	Error error
}


// func SequentialDownloader(urls []string , destDir string) error {
// 	start := time.Now()
// 	if err := os.MkdirAll(destDir, 0755); err != nil {
// 		return err
// 	}
// 	for _, url := range urls {
// 		filename := filepath.Base(url) // Extract filename from URL
// 		err := DownloadFile(url, destDir) // Declare err within the loop
// 		if err != nil {
// 			fmt.Println("Error downloading", filename, ":", err)
// 			continue
// 		}
// 	}
// 	fmt.Println("Downloading all files took", time.Since(start), "time.")
// 	return nil
// }

func ConcurrentDownloader(urls []string, destDir string, maxConcurrent int) error {
	err := os.MkdirAll(destDir, 0755)
	if err != nil {
		return err
	}

	results := make(chan Result)

	var wg sync.WaitGroup

	limiter := make(chan struct{}, maxConcurrent)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			limiter <- struct{}{}
			defer func() {<-limiter}()

			start := time.Now()
			fileName := filepath.Base(url) // Fixed filePath to filepath
			filePath := filepath.Join(destDir, fileName)

			out, err := os.Create(filePath)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			defer out.Close()

			resp, err := http.Get(url)
			if err != nil {
				out.Close()
				_ = os.Remove(filePath)
				results <- Result{URL: url, Error: err}
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				out.Close()
				_ = os.Remove(filePath)
				results <- Result{URL: url, Error: fmt.Errorf("Bad Status: %s", resp.Status)}
				return
			}

			size, err := io.Copy(out, resp.Body)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}

			timeSince := time.Since(start)
			results <- Result{URL: url, FileName: fileName, Size: size, Duration: timeSince, Error: nil}
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var totalSize int64
	var errors []error
	start := time.Now()

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error downloading %s: %s\n", result.URL, result.Error.Error())
			errors = append(errors, result.Error)
		} else {
			totalSize += result.Size
			fmt.Println("Downloaded", result.FileName)
		}
	}

	fmt.Printf("All downloads completed in %s\n", time.Since(start))
	if len(errors) > 0 {
		return fmt.Errorf("errors downloading: %v", errors)
	}
	return nil
}




func main(){
urls:= []string{
	"https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
	// "https://file-examples.com/wp-content/uploads/2017/02/file_example_CSV_5000.csv",
	// "https://file-examples.com/wp-content/uploads/2017/10/file_example_JPG_100kB.jpg",
	"https://www.arvindguptatoys.com/arvindgupta/funny-up-ruskin.pdf",
	"https://tourism.gov.in/sites/default/files/2019-04/dummy-pdf_2.pdf",
}
// url:= "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf"

// err:= DownloadFile(url,"./")
err:=ConcurrentDownloader(urls,"./",3)
if err != nil{
		fmt.Println(err)
		return
	}
log.Println("Done")
}