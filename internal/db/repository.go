package db

import (
    "log"
)

func InitDB() {
    log.Println("Ma'lumotlar bazasiga ulanish sozlandi")
}

func SaveQuiz(quiz interface{}) error {
    log.Println("Quiz ma'lumotlari saqlandi")
    return nil
}
