package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lll-phill-lll/shortener/api"
	"github.com/lll-phill-lll/shortener/logger"
	"github.com/lll-phill-lll/shortener/pkg/storage"
	"math/rand"
	"net/http"
	"time"
)

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func Hash(w http.ResponseWriter, r *http.Request) {
	u := mux.Vars(r)
	url, in := storage.DB[u["hash"]]
	if in {
		http.Redirect(w, r, url, http.StatusSeeOther)

	} else {
		_, err := fmt.Fprintln(w, "Not found")
		logger.Error.Println(err)
	}
}

func Short(w http.ResponseWriter, r *http.Request) {
	logger.Debug.Println("short")
	if r.Method != "POST" {
		_, err := fmt.Fprintln(w, "Should be post")
		if err != nil {
			logger.Error.Println(err)
			return
		}
	}
	req := api.Request{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		logger.Error.Println(err)
	}

	h := RandStringBytesMaskImprSrc(4)
	storage.DB[h] = req.Url
	_, err = fmt.Fprintln(w, "http://localhost:8080/"+h)
	if err != nil {
		logger.Error.Println(err)
	}

}
