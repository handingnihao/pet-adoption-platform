package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "Test@123456"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("生成密码哈希失败: %v\n", err)
		return
	}
	fmt.Printf("密码: %s\n", password)
	fmt.Printf("哈希值: %s\n", string(hash))
	
	// 验证
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err == nil {
		fmt.Println("✓ 验证成功")
	} else {
		fmt.Println("✗ 验证失败")
	}
}
