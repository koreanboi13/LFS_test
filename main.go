package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

// Намеренная ошибка 1: hardcoded password (gosec G101)
const adminPassword = "supersecret123"

// Намеренная ошибка 2: MD5 используется для хэширования паролей (gosec G401)
func hashPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Намеренная ошибка 3: игнорирование ошибки (errcheck)
func writeToFile(filename, content string) {
	f, _ := os.Create(filename)
	f.WriteString(content)
	f.Close()
}

func main() {
	//test
	hash := hashPassword(adminPassword)
	fmt.Println("Hash:", hash)
	writeToFile("/tmp/output.txt", "hello")
}
