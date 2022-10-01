package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	dataImagedto "tchtest/dto/dataImage"
	dto "tchtest/dto/result"
	"tchtest/models"
	"tchtest/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerImage struct {
	DataImageRepository repositories.DataImageRepository
}

func HandlerDataImage(DataImageRepository repositories.DataImageRepository) *handlerImage {
	return &handlerImage{DataImageRepository}
}

func (h *handlerImage) FindDatas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dimage, err := h.DataImageRepository.FindDatas()
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response := dto.ErrorResult{Code: http.StatusForbidden, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: dimage}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerImage) GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//var kendaraan models.DataKendaraan
	dimage, err := h.DataImageRepository.GetData(id)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response := dto.ErrorResult{Code: http.StatusForbidden, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseImage(dimage)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerImage) CreateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	request := dataImagedto.DataImageRequest{
		Image: filename,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil { //Jika ada error maka tampilkan pemberitahuan berupa error
		w.WriteHeader(http.StatusForbidden)
		response := dto.ErrorResult{Code: http.StatusForbidden, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataImage := models.Data{
		Image:  filename,
		UserID: userId,
	}

	dataImage, err = h.DataImageRepository.CreateData(dataImage)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response := dto.ErrorResult{Code: http.StatusForbidden, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseImage(dataImage)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseImage(p models.Data) models.DataImageResponse {
	return models.DataImageResponse{
		ID:    p.ID,
		Image: p.Image,
	}
}
