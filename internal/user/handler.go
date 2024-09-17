package user

import (
    "encoding/json"
    "net/http"
    "110HW/internal/quiz"
)

func HandleQuizzes(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        TakeQuiz(w, r)
    case "POST":
        SubmitAnswers(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func TakeQuiz(w http.ResponseWriter, r *http.Request) {
    quizzes, err := quiz.GetAllQuizzes()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(quizzes)
}

func SubmitAnswers(w http.ResponseWriter, r *http.Request) {
    var submission quiz.Submission
    err := json.NewDecoder(r.Body).Decode(&submission)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    score, err := quiz.EvaluateQuiz(submission)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]int{"score": score})
}
