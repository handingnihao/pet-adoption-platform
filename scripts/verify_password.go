package main

import (
	"fmt"
	"os"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("用法: go run verify_password.go <密码> <哈希值>")
		fmt.Println("或: go run verify_password.go <密码> (生成新哈希)")
		return
	}

	password := os.Args[1]
	
	if len(os.Args) == 2 {
		// 只生成哈希
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("生成密码哈希失败: %v\n", err)
			return
		}
		fmt.Printf("密码: %s\n", password)
		fmt.Printf("哈希值: %s\n", string(hash))
	} else {
		// 验证哈希
		hash := os.Args[2]
		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err == nil {
			fmt.Println("✓ 密码匹配")
		} else {
			fmt.Printf("✗ 密码不匹配: %v\n", err)
		}
	}
}
