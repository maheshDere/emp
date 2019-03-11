package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	api "github.com/joshsoftware/emp/empServer/api"
)

func Create(rw http.ResponseWriter, req *http.Request) {
	var data User
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		api.Error(http.StatusBadRequest, Response{Output: err.Error()}, rw)
		return
	}

	err = createService(data)
	if err != nil {
		api.Error(http.StatusInternalServerError, Response{Output: err.Error()}, rw)
		return
	}
	fmt.Println("handler :user create successful")
	api.Success(http.StatusOK, Response{Output: "User Created Successfully"}, rw)
}

func Read(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	user, err := readService(id)
	if err == errUserNotExist {
		api.Error(http.StatusNotFound, Response{Output: err.Error()}, rw)
		return
	}
	if err != nil {
		api.Error(http.StatusInternalServerError, Response{Output: err.Error()}, rw)
		return
	}
	fmt.Println("handler : read successful")
	api.Success(http.StatusOK, user, rw)
}

func Update(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var data User
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		api.Error(http.StatusBadRequest, Response{Output: err.Error()}, rw)
		return
	}

	err = updateService(id, data)
	if err != nil {
		api.Error(http.StatusInternalServerError, Response{Output: err.Error()}, rw)
		return
	}
	api.Success(http.StatusOK, Response{Output: "User Updated Successfully"}, rw)
	fmt.Println("handler: update successful")
}
func Delete(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	err := deleteService(id)
	if err != nil {
		api.Error(http.StatusInternalServerError, Response{Output: err.Error()}, rw)
		return
	}
	fmt.Println("handler : delete successful")
	api.Success(http.StatusOK, Response{Output: "User Delete Successfully"}, rw)

}

func ReadAll(rw http.ResponseWriter, req *http.Request) {
	users, err := readAllService()
	if err != nil {
		api.Error(http.StatusInternalServerError, Response{Output: err.Error()}, rw)
		return
	}
	api.Success(http.StatusOK, users, rw)
	fmt.Println("handler : all user successful")
}

func isBadRequest(err error) {
	return
}
