package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	dto "waysbuck/dto/result"
	topingdto "waysbuck/dto/toping"
	"waysbuck/models"
	"waysbuck/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerToping struct {
	TopingRepository repositories.TopingRepository
}

var toping_file = "http://localhost:5000/uploads/"

func HandlerToping(TopingRepository repositories.TopingRepository) *handlerToping {
	return &handlerToping{TopingRepository}
}

// GET ALL
func (h *handlerToping) FindTopings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	topings, err := h.TopingRepository.FindTopings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range topings {
		topings[i].Image = toping_file + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccesFindTopings{Code: http.StatusOK, Data: topings}
	json.NewEncoder(w).Encode(response)
}

// GET DETAIL
func (h *handlerToping) GetToping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var toping models.Toping
	toping, err := h.TopingRepository.GetToping(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	toping.Image = toping_file + toping.Image
	w.WriteHeader(http.StatusOK)
	response := dto.SuccesGetToping{Code: http.StatusOK, Data: toping}
	json.NewEncoder(w).Encode(response)
}

// CREATE
func (h *handlerToping) CreateToping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Get dataFile from midleware and store to filename variable here ...
	dataContex := r.Context().Value("dataFile") //GET THE VALUE
	filename := dataContex.(string)             //CONVERT TO STRING

	price, _ := strconv.Atoi(r.FormValue("price"))
	request := topingdto.TopingRequest{
		Title: r.FormValue("title"),
		Price: price,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	toping := models.Toping{
		Title:     request.Title,
		Price:     request.Price,
		Image:     filename,
		ProductID: userId,
	}

	// err := mysql.DB.Create(&product).Error
	toping, err = h.TopingRepository.CreateToping(toping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	toping, _ = h.TopingRepository.GetToping(toping.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccesGetToping{Code: http.StatusOK, Data: toping}
	json.NewEncoder(w).Encode(response)
}

func convertResponseToping(u models.Toping) models.TopingResponse {
	return models.TopingResponse{
		ID:    u.ID,
		Title: u.Title,
		Price: u.Price,
		Image: u.Image,
	}
}
