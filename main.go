package main

import (
	"fmt"
"html/template"
"log"
"net/http"
"os"
"os/exec"
)


func main() {

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/request", reqAjax)

	port := ":80"
	println("Server listen on port ", port)
	name, _ := os.Hostname()
	println(name)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe ", err)
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Print("New connection.. \n")
	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func reqAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sendedData := r.FormValue("sendedData")

		if sendedData == "Paint" {
			_ = exec.Command("mspaint").Run()
		} else if sendedData == "shutdown"{
			_ = exec.Command("shutdown", "/s").Run()
		}

		fmt.Println("My request is: ", sendedData)
	}
}
