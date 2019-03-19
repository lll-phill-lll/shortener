package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lll-phill-lll/shortener/api"
	"github.com/lll-phill-lll/shortener/logger"
	"github.com/lll-phill-lll/shortener/pkg/task"
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

func (serv *Impl) hash(w http.ResponseWriter, r *http.Request) {
	u := mux.Vars(r)
	hash := u["hash"]
	task, err := serv.DB.Load(hash)
	logger.Info.Println(r.URL.Path)
	if err != nil {
		logger.Error.Println(err)
		_, err := fmt.Fprintln(w, "Not found")
		if err != nil {
			logger.Error.Println(err)
		}

	} else {
		http.Redirect(w, r, task.URL, http.StatusSeeOther)
	}
}

func (serv *Impl) short(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	h := RandStringBytesMaskImprSrc(4)
	task := task.Task{URL: req.URL, Hash: h, HostURL: serv.HostURL}
	err = serv.DB.Save(task)
	if err != nil {
		logger.Error.Println(err.Error())
	}
	toDend, _ := json.Marshal(api.Response{HashedURL: task.GetHashedURL()})
	_, err = fmt.Fprintln(w, string(toDend))
	if err != nil {
		logger.Error.Println(err)
	}
}
