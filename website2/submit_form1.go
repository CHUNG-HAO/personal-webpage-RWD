package main

import (
	"fmt"
	"net/http"
	"net/smtp"
)

func main() {
	http.HandleFunc("/submit_form", submitFormHandler)
	http.ListenAndServe(":8080", nil)
}

func submitFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 解析表單數據
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "無法解析表單數據", http.StatusBadRequest)
			return
		}

		name := r.Form.Get("name")
		email := r.Form.Get("email")
		phone := r.Form.Get("phone")
		need := r.Form.Get("need")

		// 設置郵件內容
		message := fmt.Sprintf("姓名: %s\n郵件: %s\n電話: %s\n需求: %s\n", name, email, phone, need)

		// 發送郵件
		auth := smtp.PlainAuth("", "chunghao777@gmail.com", "avtc911024", "smtp.gmail.com")
		err = smtp.SendMail("smtp.gmail.com:587", auth, "chunghao777@gmail.com", []string{"chunghao777@gmail.com"}, []byte(message))
		if err != nil {
			http.Error(w, "無法發送郵件", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "郵件已成功發送。")
	} else {
		http.Error(w, "僅支持POST請求", http.StatusMethodNotAllowed)
	}
}
