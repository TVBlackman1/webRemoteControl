package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
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
		sentData := r.FormValue("sentData")
		fmt.Println("My request is: ", sentData)

		if sentData == "Paint" {
			_ = exec.Command("mspaint").Run()
		} else if sentData == "shutdown"{
			_ = exec.Command("shutdown", "/s", "/t", "0").Run()
		} else if sentData == "em-shutdown" {
			_ = exec.Command("shutdown", "/s", "/t", "30").Run()
			time.Sleep(20 * time.Second)
			_ = exec.Command("shutdown", "-a").Run()
		} else if sentData == "enter" {
			dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
			file := dir + "\\scripts\\start_dota.exe"
			er := exec.Command(file).Run()
			if er != nil {
				log.Fatal("ListenAndServe ", er, er.Error())
			}
		}
	}
}
