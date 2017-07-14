package teststorage

import (
	_ "cloud.google.com/go/storage"
	_ "golang.org/x/oauth2"
	_ "google.golang.org/appengine"
	"net/http"
	"google.golang.org/appengine"
	"context"
	"cloud.google.com/go/storage"
	"io"
)



func main() {
	http.HandleFunc("/", serveFile)
	appengine.Main()
}

func serveFile (w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, _ := storage.NewClient(ctx)
	rc, err := client.Bucket("teststatics").Object("html/main.html").NewReader(ctx)
	if err != nil {
		return
	}
	defer rc.Close()
	io.Copy(w, rc)

}