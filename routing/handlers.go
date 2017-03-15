package routing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"io"
	"../model"
	"../data"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	/*claims, ok := r.Context().Value(middleware.MyKey).(model.Claims)
	if !ok {
		http.NotFound(w, r)
		return
	}
*/
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data.Users); err != nil {
		panic(err)
	}
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	userId := vars["userId"]

	fmt.Fprintln(w, "User show:", userId)
}

func UserCreate( w http.ResponseWriter, r * http.Request) {
	var user model.User

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	u := data.RepoCreateUser(user)
	w.Header().Set("Content-Type", "applocation/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

func Status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Api is up and running"))
}

func UdacityCourses (w http.ResponseWriter, r *http.Request) {
	url := "https://www.udacity.com/public-api/v0/courses"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		defer response.Body.Close()
		courseList := model.UdacityCourseCollection{}
		err := json.NewDecoder(response.Body).Decode(&courseList)
		if err != nil {
			panic(err)
		}

		coursesJson, error := json.Marshal(courseList.Courses)
		if error != nil {
			panic(error)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		w.Write(coursesJson)
	}
}