package admin

import (
    "encoding/json"
    "net/http"
    "110HW/internal/quiz"
)

func HandleQuizzes(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "POST":
        CreateQuiz(w, r)
    case "GET":
        GetQuizzes(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func CreateQuiz(w http.ResponseWriter, r *http.Request) {
    var q quiz.Quiz
    err := json.NewDecoder(r.Body).Decode(&q)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }
    
    err = quiz.CreateQuiz(q)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(q)
}

func GetQuizzes(w http.ResponseWriter, r *http.Request) {
    quizzes, err := quiz.GetAllQuizzes()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(quizzes)
}
