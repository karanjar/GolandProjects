package cars

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Car struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Year  int64   `json:"year"`
	Price float64 `json:"price"`
}

var Cars = make(map[int]Car)

var Mu sync.Mutex

func Carhandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	entity := strings.TrimPrefix(path, "/cars")
	entity = strings.Trim(entity, "/")

	switch r.Method {
	case "POST":
		if entity == "" {
			Createcar(w, r)
		} else {
			http.Error(w, "Incorrect Bad Request", http.StatusBadRequest)
		}
	case "GET":
		if entity == "" {
			http.Error(w, "You entered bad get request", http.StatusBadRequest)
		} else {
			id, _ := strconv.Atoi(entity)
			Getcar(w, r, id)

		}
	case "DELETE":
		if entity == "" {
			http.Error(w, "You entered bad delete request", http.StatusBadRequest)
		} else {
			id, _ := strconv.Atoi(entity)
			Deletecar(w, r, id)
		}
	}

}

func Createcar(w http.ResponseWriter, r *http.Request) {
	Mu.Lock()
	defer Mu.Unlock()

	var car Car

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "incorrect json input", http.StatusBadRequest)
		return
	}

	id := rand.Intn(1000)
	car.Id = int64(id)

	Cars[int(car.Id)] = car

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)

}

func Getcar(w http.ResponseWriter, r *http.Request, id int) {
	Mu.Lock()
	defer Mu.Unlock()

	val, ok := Cars[id]
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(val)

}
func Deletecar(w http.ResponseWriter, r *http.Request, id int) {
	Mu.Lock()
	defer Mu.Unlock()
	_, ok := Cars[id]
	if !ok {
		http.Error(w, "Car Not Found", http.StatusNotFound)
		return
	}
	delete(Cars, id)

	w.WriteHeader(http.StatusNoContent)

}
