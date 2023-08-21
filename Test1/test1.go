package main

// import thư viện
import (
	"fmt"
)

// tạo interface Fetcher với phương thức Fetch
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// Tạo phương thức Fetch cho lớp fakeResult 
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	// gán ok=url; 
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

// Định nghĩa cấu trúc class fakeResult
type fakeResult struct {
	body string
	urls []string
}

// Khai báo kiểu dữ liệu map với key: string, value: fakeResult (kiểu con trỏ)
type fakeFetcher map[string]*fakeResult


// Tạo hàm Crawl Đệ Quy
func Crawl(url string, depth int, fetch Fetcher) {
	// fmt.Println(url)
	if depth <= 0 {
		return
	}

	body, urls, err := fetch.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	// đã tìm thấy đường dẫn
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		Crawl(u, depth - 1, fetcher)
	}

	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	fmt.Println("Pause");
}


// Tạo 1 mảng gồm 4 phần tử với key: string, value: fakeResult
// &fakeResult: tạo đối tượng tham chiếu tới kiểu con trỏ fakeFetcher
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}