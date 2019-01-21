package urlStatus

import (
	"fmt"
	"net/http"
	"os"
)

/*var urlmap

func init()  {

	urlmap = make(map[string]bool)
}
*/
func CheckUrl(url string, c chan bool) {
	response, err := http.Get(url)
	if err != nil {
		c <- false
		fmt.Println("Error for:", url, "\n\t", err)
		return
	}

	c <- true
	fmt.Println("Result for:", url, "\n\t", response.Status)
}

func CheckUrls(urls []string, no int, done chan bool) {
	//urlmap := make(map[string]bool)
	var c = make(chan bool)
	//var done = make(chan bool)
	go func(no int) {

		for range urls {
			i := 0
			status := <-c
			if !status {
				//		fmt.Println("Status :", status)
				i++
			}
			fmt.Println("fail count:", i)
			if i >= no {
				done <- true
				os.Exit(0)
			}
		}
		defer func() {
			done <- true
		}()
	}(no)

	for _, u := range urls {
		go CheckUrl(u, c)
	}

	//return urlmap
}
