package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	dto "tchtest/dto/result"
)

func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, hand, err := r.FormFile("image")

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			response := dto.ErrorResult{Code: http.StatusForbidden, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", hand.Filename)

		const MAX_UPLOAD_SIZE = 8192 * 8192 // 8MB

		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusForbidden)
			response := Result{Code: http.StatusForbidden, Message: "Max size in 8mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			response := Result{Code: http.StatusForbidden, Message: "path upload error"}
			json.NewEncoder(w).Encode(response)
			return
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// write this byte array to our temporary file
		tempFile.Write(fileBytes)

		data := tempFile.Name()
		filename := data[8:] // split uploads/

		// add filename to ctx
		ctx := context.WithValue(r.Context(), "dataFile", filename)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
