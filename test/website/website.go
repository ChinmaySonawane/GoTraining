package website

import (
	"net/http"
)

type Website struct {
	Url    string
	Status bool
}

func NewWebsite(url string) *Website {
	return &Website{url, false}
}

func (w *Website) CheckStatus(done chan *Website) {
	_, err := http.Get(w.Url)
	if err != nil {
		w.Status = false
		done <- w
		return
	}
	done <- w
	w.Status = true
}
