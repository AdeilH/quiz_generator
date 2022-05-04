package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func previousquestion(rw http.ResponseWriter, r *http.Request) {
	if question_number-1 < 0 {
		reason := &Response{
			Reason:      "You are at the first question",
			Description: "You are at the first question",
		}
		byteArray, err := json.Marshal(reason)
		if err != nil {
			fmt.Println(err)
		}
		rw.Write(byteArray)
		return
	} else {
		reason := &Response{
			Reason:      "Moving",
			Description: "Moving",
		}
		byteArray, err := json.Marshal(reason)
		if err != nil {
			fmt.Println(err)
		}
		rw.Write(byteArray)
		question_number = question_number - 1
	}
}
func nextquestion(rw http.ResponseWriter, r *http.Request) {
	if question_number >= total_questions-1 {
		reason := &Response{
			Reason:      "You are at the last question",
			Description: "You are at the last question",
		}
		byteArray, err := json.Marshal(reason)
		if err != nil {
			fmt.Println(err)
		}
		rw.Write(byteArray)
		// return
		for _, value := range score_map {
			total = total + value
		}
		finito = true
	} else {
		reason := &Response{
			Reason:      "Moving",
			Description: "Moving",
		}
		byteArray, err := json.Marshal(reason)
		if err != nil {
			fmt.Println(err)
		}
		rw.Write(byteArray)
		question_number = question_number + 1
	}
}
func retryquiz(rw http.ResponseWriter, r *http.Request) {
	reason := &Response{
		Reason:      "Restarting",
		Description: "Restarting the quiz",
	}
	byteArray, err := json.Marshal(reason)
	if err != nil {
		fmt.Println(err)
	}
	rw.Write(byteArray)
	finito = false
	score_map = make(map[int]int)
	question_number = 0
	total = 0
}
func endquiz(rw http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func startquiz(w http.ResponseWriter, req *http.Request) {
	if finito {
		tmpl := template.Must(template.ParseFiles("templates/final_score.html"))
		final_context := Final_Context{
			Title:            quiz_list.Title,
			Total:            total,
			Background_Image: background_image,
		}
		tmpl.Execute(w, final_context)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/quiz_template.html"))
	w.Header().Add("Content Type", "text/html")
	context := Context{
		Title:            quiz_list.Title,
		Question:         quiz_list.Quiz[question_number].Question,
		Options:          quiz_list.Quiz[question_number].Options,
		Background_Image: background_image,
	}
	tmpl.Execute(w, context)
	if req.Method == http.MethodPost {
		req.ParseForm()
		if req.Form.Get("option") == quiz_list.Quiz[question_number].Answer {
			score_map[question_number] = 1
		} else {
			score_map[question_number] = 0
		}
	}

}

func register_handles() {
	fs := http.FileServer(http.Dir("./background_image"))
	http.Handle("/background_image/", http.StripPrefix("/background_image/", fs))
	http.HandleFunc("/", startquiz)
	http.HandleFunc("/gotoprev", previousquestion)
	http.HandleFunc("/gotonext", nextquestion)
	http.HandleFunc("/gotostart", retryquiz)
	http.HandleFunc("/gotoend", endquiz)
}

func handle_flags() {
	flag.StringVar(&quiz_file, "f", "quiz_json/test_quiz.json", "Specify Quiz File. Default is test_quiz.json")
	flag.StringVar(&background_image, "i", "background_image/dr_who.jpg", "Specify Background Image.")

	flag.Parse()
}
