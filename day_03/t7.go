// 7. Реализовать программу «Викторина» (Кто хочет стать миллионером) - вопросы-варианты ответов, используя структуры

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	MAX_GAME_ID   = 100
	MAX_CLIENT_ID = 10000
)

type Client struct {
	Id   int  `json:"id"`
	Game Game `json:"game"`
}

type Game struct {
	Id              int      `json:"id"`
	LastQuestionId  int      `json:"last_question_id"`
	CurrentQuestion Question `json:"current_question"`
	AnswerId        int      `json:"answer_id"`
	weight          int
}

type Answer struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	weight int
}

type Question struct {
	Id      int      `json:"id"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

type Clients = map[int]Client

var clients = make(Clients)

var questions = []Question{
	{
		Id:   1,
		Text: "Каким отношением описывается наследование в ООП",
		Answers: []Answer{
			{
				Id:     1,
				Text:   "Содержит",
				weight: 0,
			},
			{
				Id:     2,
				Text:   "Является",
				weight: 10,
			},
			{
				Id:     2,
				Text:   "Не содержит",
				weight: 0,
			},
		},
	},
}

// http://127.0.0.1:8080?client_id=1&question_id=2&answer_id=3
func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if q.Get("client_id") == "" {
		newGame(w)
		return
	}

	cId, _ := strconv.Atoi(q.Get("client_id"))
	aId, _ := strconv.Atoi(q.Get("answer_id"))

	c := clients[cId]
	c.Game.AnswerId = aId

	checkAnswer(c, w)
}

func newGame(w http.ResponseWriter) {
	gameId := rand.Intn(MAX_GAME_ID)
	cId := rand.Intn(MAX_CLIENT_ID)

	game := Game{
		Id:              gameId,
		LastQuestionId:  len(questions),
		CurrentQuestion: questions[0],
	}

	c := Client{
		Id:   cId,
		Game: game,
	}

	clients[cId] = c

	writeAnswer(c, w)
}

func checkAnswer(c Client, w http.ResponseWriter) {
	weight := questions[c.Game.CurrentQuestion.Id-1].Answers[c.Game.AnswerId-1].weight
	c.Game.weight = c.Game.weight + weight

	if c.Game.LastQuestionId < c.Game.CurrentQuestion.Id {
		c.Game.CurrentQuestion = questions[c.Game.CurrentQuestion.Id+1]
		clients[c.Id] = c
		writeAnswer(c, w)
		return
	}

	writeAnswer(fmt.Sprintf("your scores: %d", c.Game.weight), w)
}

func writeAnswer(data any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
