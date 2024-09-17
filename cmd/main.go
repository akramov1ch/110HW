package main

import (
    "log"
    "net/http"
    "110HW/internal/admin"
    "110HW/internal/user"
    "110HW/internal/db"
)

func main() {
    db.InitDB()

    http.HandleFunc("/admin/quiz", admin.HandleQuizzes)   // Admin API
    http.HandleFunc("/user/quiz", user.HandleQuizzes)     // Foydalanuvchi API

    log.Println("Server 8080 portda ishlayapti...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
