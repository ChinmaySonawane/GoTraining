package website

import (
	"net/http"
)

//struct for website to hold result
type Website struct {
	Url    string
	Status bool
}

//constructor for Website
func NewWebsite(url string) *Website {
	return &Website{url, false}
}

//CheckStatus is func call on Website take chan *Website to pass status
func (w *Website) CheckStatus(done chan *Website) {
	_, err := http.Get(w.Url)
	if err != nil {
		w.Status = false
		done <- w //passing ptr to Website struct in chan
		return
	}
	done <- w //passing ptr to Website struct in chan
	w.Status = true
}
