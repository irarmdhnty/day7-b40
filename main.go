package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/project",project).Methods("GET")
	r.HandleFunc("/add-project", addProject).Methods("POST")
	r.HandleFunc("/detail", detail).Methods("GET")

	fmt.Println("server on in port 8000")
	http.ListenAndServe("localhost:8000", r )
}

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	
	tmpl, err := template.ParseFiles("views/index.html")

	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.Execute(w,"")
}

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/contact.html")
	if err != nil{
		w.Write([]byte(err.Error()))
	}

	tmpl.Execute(w, "")
}

func project(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("views/addProject.html")
	if err != nil{
		w.Write([]byte(err.Error()))
	}

	tmpl.Execute(w, "")
}

func addProject(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Project Name :", r.PostForm.Get("inputName"))
	fmt.Println("Start date :", r.PostForm.Get("startDate"))
	fmt.Println("End Date :", r.PostForm.Get("endDate"))
	fmt.Println("Desc :", r.PostForm.Get("desc"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func detail(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("views/detail.html")
	if err != nil{
		w.Write([]byte(err.Error()))
	}

	data := map[string]interface{}{
		"Date": "20 Apr 2022 - 27 jul 2022",
		"Duration": "2 Month",
		"Desc": "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Pariatur officia distinctio esse commodi, quasi saepe quidem molestias. Hic deleniti aliquam pariatur eligendi vitae deserunt placeat suscipit, voluptate, rerum dolore corrupti!",
	}

	tmpl.Execute(w, data)
}