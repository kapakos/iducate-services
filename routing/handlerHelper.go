package routing

import (
	"github.com/kapakos/iducate-services/model"
	"net/http"
	"encoding/json"
)

func SendCourses(url string, m *model.CourseraCourseCollection, w http.ResponseWriter) {
response, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		defer response.Body.Close()
		err := json.NewDecoder(response.Body).Decode(&m)
		if err != nil {
			panic(err)
		}

		coursesJson, error := json.Marshal(m.Courses)
		if error != nil {
			panic(error)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		w.Write(coursesJson)
	}
}