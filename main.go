package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

func main() {
	http.HandleFunc("/", handleFormSubmit)
	fmt.Println("已開啟伺服器")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleFormSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析表单数据
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	need := r.FormValue("need")

	// 邮件接收者
	to := []string{""}

	// 邮件内容
	message := "姓名: " + name + "\n"
	message += "郵件: " + email + "\n"
	message += "電話: " + phone + "\n"
	message += "需求: " + need + "\n"

	// 发送邮件
	from := ""
	password := ""

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, from, to, []byte(message))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Failed to send email:", err)
		return
	}

	fmt.Fprintf(w, "表单已成功提交并发送到您的邮箱。")
}
