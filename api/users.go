package api

import (
	"fmt"
	"intracs_anpr_api/controller/user"
	"intracs_anpr_api/internal/response"
	"net/http"
)

type UserApi struct {
	controller *user.Controller
}

func InitUserApi(controller *user.Controller) *UserApi {
	return &UserApi{
		controller: controller,
	}
}

func (u *UserApi) Login(w http.ResponseWriter, r *http.Request) {
	response := response.NewAPIResponse()
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email != "" && password != "" {
			user, err := u.controller.Login(email, password)
			if err != nil {
				response.Update("bad request", err.Error(), nil)
				http.Error(w, response.ToJsonString(), http.StatusBadRequest)

				return
			}

			response.Update("OK", "login successfully", user)
			fmt.Fprint(w, response.ToJsonString())

			return
		}
		response.Update("Bad Request", "email or password can't null", nil)
		http.Error(w, response.ToJsonString(), http.StatusBadRequest)
		return
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}
