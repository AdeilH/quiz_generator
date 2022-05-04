package main

var quiz_list Quiz
var score_map map[int]int
var total int
var finito bool
var question_number int
var total_questions int
var background_image string
var quiz_file string

type Response struct {
	Reason      string `json:"Reason"`
	Description string `json:"Description"`
}

type Context struct {
	Title            string
	Question         string
	Options          []string
	Background_Image string
}

type Final_Context struct {
	Title            string
	Total            int
	Background_Image string
}

type Quiz struct {
	Title string     `json:"title"`
	Quiz  []Question `json:"quiz"`
}

type Question struct {
	Question string   `json:"Question"`
	Options  []string `json:"Options"`
	Answer   string   `json:"Answer"`
}

func initialize_vars() {
	finito = false
	score_map = make(map[int]int)
	question_number = 0
	total = 0
	quiz_list = ParseFile(quiz_file)
	total_questions = len(quiz_list.Quiz)
}
