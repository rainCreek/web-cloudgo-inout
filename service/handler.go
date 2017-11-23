package service

import (
	"net/http"

	"github.com/unrolled/render"
  "log"
  "html/template"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", struct {
			Student      string `json:"student"`
      StudentID        string `json:"studentid"`
			Title      string `json:"title"`
		}{Student: "Rao Yuxi", StudentID: "15331262", Title: "cloudgo-inout"})
	}
}

// when request URL path = "/unknown", return status code 5xx
func unknownHandler(formatter *render.Render) http.HandlerFunc {

  return func(w http.ResponseWriter, req *http.Request) {
    if req.URL.Path == "/unknown" {
      t, err := template.ParseFiles("./templates/5xx.html")
      if err != nil {
        log.Println(err)
      }
      t.Execute(w, nil)  
    }     
  }
}

func fillinHandler(formatter *render.Render) http.HandlerFunc {

  return func(w http.ResponseWriter, req *http.Request) {
    if req.URL.Path == "/fillin" {
      t, err := template.ParseFiles("./templates/form.html")

      req.ParseForm()
      na := req.FormValue("name")
      ag := req.FormValue("age")
      gd := req.FormValue("gender")
      formatter.HTML(w, http.StatusOK, "table", struct {
        Name      string `json:"name"`
        Age        string `json:"age"`
        Gender      string `json:"gender"`
      } {Name: na, Age: ag, Gender: gd})
      if err != nil {
        log.Println(err)
      }
      t.Execute(w, nil)  
    }     
  }

}

