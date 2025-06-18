package application

import (
	"beginning_http/model"
	"encoding/json"
	"net/http"
)

type Api struct {
	Addr string
}

// Assuming the User type is defined in a package named 'models'.
// You will need to add an import statement for the 'models' package
// at the top of the file, like: import "your_module_path/models"
// replacing 'your_module_path' with the actual module path.
var users = []model.User{
	{FirstName: "John", LastName: "Doe", Age: 30},
	{FirstName: "Jane", LastName: "Smith", Age: 25},
	{FirstName: "Peter", LastName: "Jones", Age: 42},
	{FirstName: "Alice", LastName: "Williams", Age: 28},
	{FirstName: "Bob", LastName: "Brown", Age: 35},
}

func (s *Api) GetHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page"))
}

func (s *Api) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	val, err := json.Marshal(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(val)
}

func (s *Api) CreateUsers(w http.ResponseWriter, r *http.Request) {
    var payload model.User

    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    u := model.User{
        FirstName: payload.FirstName,
        LastName:  payload.LastName,
        Age:       payload.Age,
    }
    users = append(users, u)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(users) 
}
