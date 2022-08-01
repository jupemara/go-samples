package main

import (
	"net/http"

	"google.golang.org/appengine/v2"
	"google.golang.org/appengine/v2/blobstore"
	"google.golang.org/appengine/v2/image"
)

func main() {
	http.HandleFunc("/", handler)
	appengine.Main()
}

func handler(w http.ResponseWriter, r *http.Request) {
	body := "URL: "
	ctx := appengine.NewContext(r)
	bk, err := blobstore.BlobKeyForFile(ctx, "/gs/images-api-samples/test.png")
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
