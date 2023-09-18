package main

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	// Cấu hình Sentry.
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "YOUR_SENTRY_DSN",
	})
	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		return
	}
	defer sentry.Flush(2 * time.Second) // Đảm bảo các sự kiện còn lại được gửi đi trước khi thoát.

	// Tạo một UserService.
	userService := NewUserService()

	// Sử dụng use case để tạo người dùng và xử lý lỗi nếu có.
	err = userService.CreateUser("John Doe")
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
	}
}
