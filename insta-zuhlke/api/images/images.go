package images

import (
	"encoding/json"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/rs/xid"
)

func Upsert(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	go storeImage(file)
}

func storeImage(file multipart.File) {
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
		return
	}

	out, err := os.Create("static/" + xid.New().String() + ".jpg")
	if err != nil {
		log.Println(err)
		return
	}

	defer out.Close()

	err = jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

type ImageListing struct {
	Items []string `json:"items"`
}

func List(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir("./static")
	fileNames := make([]string, 0, len(files))

	for _, f := range files {
		name := f.Name()
		if !strings.HasPrefix(name, ".") {
			fileNames = append(fileNames, r.URL.String()+"/"+f.Name())
		}
	}

	response := ImageListing{Items: fileNames}

	json, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
