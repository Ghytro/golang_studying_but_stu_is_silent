package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"urlshortener/util"
)

var port int = 4000
var urlStorage *util.UrlStorage = util.NewUrlStorage()
var shortUrlGenerator *util.ShortUrlGenerator = util.NewShortUrlGenerator("ABCDEFG", 4)
var mux *http.ServeMux = http.NewServeMux()

func addNewUrl(responseWriter http.ResponseWriter, request *http.Request) {
	urlParameters := request.URL.Query()

	longUrl, ok := urlParameters["url"]
	if !ok {
		responseWriter.Write([]byte("Key expected: url"))
		return
	}
	_, ok = urlParameters["expires"]
	var expiresIn int64
	if !ok {
		expiresIn = -1
	} else {
		expiresInStr := urlParameters["expires"]
		expiresInInt, _ := strconv.Atoi(expiresInStr[0])
		expiresIn = int64(expiresInInt)
	}
	shortUrl := shortUrlGenerator.NextShortUrl()
	fmt.Printf("Next short url: %s\n", shortUrl)
	if expiresIn != -1 {
		urlStorage.AddUrl(longUrl[0], shortUrl, time.Now().Unix()+expiresIn)
	} else {
		urlStorage.AddUrl(longUrl[0], shortUrl, -1)
	}
	mux.HandleFunc(fmt.Sprintf("/%s", shortUrl), func(w http.ResponseWriter, r *http.Request) {
		foundLongUrl, err := urlStorage.GetLongUrl(shortUrl)
		if err != nil || (foundLongUrl.ExpireTimestamp != -1 && foundLongUrl.ExpireTimestamp < time.Now().Unix()) {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, foundLongUrl.Url, http.StatusSeeOther)
	})
	responseWriter.Write([]byte(fmt.Sprintf("localhost:%d/%s", port, shortUrl)))
}

func main() {
	mux.HandleFunc("/create/", addNewUrl)

	log.Printf("Launching web server at port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
