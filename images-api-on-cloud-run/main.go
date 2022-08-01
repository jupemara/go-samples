package main

import (
	"log"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/image"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", handler)
	log.Println("starting the http server")
	if err := http.ListenAndServe("0.0.0.0:8080", m); err != nil {
		log.Fatalf("could not start the http server: %s", err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	body := "URL: "
	ctx := appengine.NewContext(r)
	bk, err := blobstore.BlobKeyForFile(ctx, "/gs/images-api-sample/test.png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	url, err := image.ServingURL(
		ctx,
		bk,
		&image.ServingURLOptions{
			Secure: true,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body = body + url.String()
	w.Write([]byte(body))
}
