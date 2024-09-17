package quiz

import (
    "errors"
    "110HW/internal/db"
)

type Quiz struct {
    ID      int      `json:"id"`
    Title   string   `json:"title"`
    Questions []Question `json:"questions"`
}

type Question struct {
    ID      int      `json:"id"`
    Text    string   `json:"text"`
    Options []Option `json:"options"`
}

type Option struct {
    ID      int    `json:"id"`
    Text    string `json:"text"`
    Correct bool   `json:"correct"`
}

var quizzes []Quiz

func CreateQuiz(q Quiz) error {
    if q.Title == "" || len(q.Questions) == 0 {
        return errors.New("invalid quiz data")
    }

    quizzes = append(quizzes, q)
    return db.SaveQuiz(q)  // Ma'lumotlar bazasiga saqlash
}

func GetAllQuizzes() ([]Quiz, error) {
    return quizzes, nil
}

type Submission struct {
    QuizID   int              `json:"quiz_id"`
    Answers  map[int]int      `json:"answers"` // question_id -> selected_option_id
}

func EvaluateQuiz(submission Submission) (int, error) {
    var score int
    for _, q := range quizzes {
        if q.ID == submission.QuizID {
            for _, question := range q.Questions {
                selectedOption := submission.Answers[question.ID]
                for _, option := range question.Options {
                    if option.ID == selectedOption && option.Correct {
                        score++
                    }
                }
            }
            return score, nil
        }
    }
    return 0, errors.New("quiz not found")
}
