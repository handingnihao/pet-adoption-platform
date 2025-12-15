package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库配置 - 请根据实际情况修改
const (
	DBUser     = "root"
	DBPassword = "root"
	DBHost     = "127.0.0.1"
	DBPort     = "3306"
	DBName     = "pet_adoption"
)

type Pet struct {
	ID         uint64
	Name       string
	Type       string
	CoverPhoto sql.NullString
	Photos     sql.NullString
}

func main() {
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		DBUser, DBPassword, DBHost, DBPort, DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("❌ 数据库连接失败:", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("❌ 数据库连接失败:", err)
		return
	}
	fmt.Println("✅ 数据库连接成功")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n========== 宠物图片管理工具 ==========")
		fmt.Println("1. 查看所有宠物")
		fmt.Println("2. 更新宠物封面图片")
		fmt.Println("3. 添加宠物相册图片")
		fmt.Println("4. 清空宠物相册")
		fmt.Println("5. 批量更新图片（从文件读取）")
		fmt.Println("0. 退出")
		fmt.Print("\n请选择操作: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			listPets(db)
		case "2":
			updateCoverPhoto(db, reader)
		case "3":
			addPhotos(db, reader)
		case "4":
			clearPhotos(db, reader)
		case "5":
			batchUpdateFromFile(db, reader)
		case "0":
			fmt.Println("👋 再见!")
			return
		default:
			fmt.Println("❌ 无效选择")
		}
	}
}

// 查看所有宠物
func listPets(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, type, cover_photo, photos FROM pets ORDER BY id")
	if err != nil {
		fmt.Println("❌ 查询失败:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n----- 宠物列表 -----")
	fmt.Printf("%-5s %-15s %-8s %-40s %s\n", "ID", "名称", "类型", "封面图片", "相册数量")
	fmt.Println(strings.Repeat("-", 100))

	for rows.Next() {
		var pet Pet
		if err := rows.Scan(&pet.ID, &pet.Name, &pet.Type, &pet.CoverPhoto, &pet.Photos); err != nil {
			continue
		}

		coverPhoto := "无"
		if pet.CoverPhoto.Valid && pet.CoverPhoto.String != "" {
			if len(pet.CoverPhoto.String) > 35 {
				coverPhoto = pet.CoverPhoto.String[:35] + "..."
			} else {
				coverPhoto = pet.CoverPhoto.String
			}
		}

		photosCount := 0
		if pet.Photos.Valid && pet.Photos.String != "" {
			var photos []string
			json.Unmarshal([]byte(pet.Photos.String), &photos)
			photosCount = len(photos)
		}

		fmt.Printf("%-5d %-15s %-8s %-40s %d张\n", pet.ID, pet.Name, pet.Type, coverPhoto, photosCount)
	}
}

// 更新封面图片
func updateCoverPhoto(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("\n请输入宠物ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64)
	if err != nil {
		fmt.Println("❌ 无效的ID")
		return
	}

	// 显示当前信息
	var pet Pet
	err = db.QueryRow("SELECT id, name, cover_photo FROM pets WHERE id = ?", id).
		Scan(&pet.ID, &pet.Name, &pet.CoverPhoto)
	if err != nil {
		fmt.Println("❌ 未找到该宠物")
		return
	}

	fmt.Printf("当前宠物: %s\n", pet.Name)
	if pet.CoverPhoto.Valid {
		fmt.Printf("当前封面: %s\n", pet.CoverPhoto.String)
	} else {
		fmt.Println("当前封面: 无")
	}

	fmt.Print("请输入新的封面图片URL (留空保持不变): ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	if url == "" {
		fmt.Println("⚠️  未修改")
		return
	}

	_, err = db.Exec("UPDATE pets SET cover_photo = ? WHERE id = ?", url, id)
	if err != nil {
		fmt.Println("❌ 更新失败:", err)
		return
	}
	fmt.Println("✅ 封面图片更新成功!")
}

// 添加相册图片
func addPhotos(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("\n请输入宠物ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64)
	if err != nil {
		fmt.Println("❌ 无效的ID")
		return
	}

	// 获取当前相册
	var pet Pet
	err = db.QueryRow("SELECT id, name, photos FROM pets WHERE id = ?", id).
		Scan(&pet.ID, &pet.Name, &pet.Photos)
	if err != nil {
		fmt.Println("❌ 未找到该宠物")
		return
	}

	var currentPhotos []string
	if pet.Photos.Valid && pet.Photos.String != "" {
		json.Unmarshal([]byte(pet.Photos.String), &currentPhotos)
	}

	fmt.Printf("当前宠物: %s\n", pet.Name)
	fmt.Printf("当前相册: %d张图片\n", len(currentPhotos))
	for i, p := range currentPhotos {
		fmt.Printf("  %d. %s\n", i+1, p)
	}

	fmt.Println("\n请输入图片URL (每行一个，输入空行结束):")
	for {
		url, _ := reader.ReadString('\n')
		url = strings.TrimSpace(url)
		if url == "" {
			break
		}
		currentPhotos = append(currentPhotos, url)
		fmt.Printf("  ✓ 已添加: %s\n", url)
	}

	photosJSON, _ := json.Marshal(currentPhotos)
	_, err = db.Exec("UPDATE pets SET photos = ? WHERE id = ?", string(photosJSON), id)
	if err != nil {
		fmt.Println("❌ 更新失败:", err)
		return
	}
	fmt.Printf("✅ 相册更新成功! 当前共 %d 张图片\n", len(currentPhotos))
}

// 清空相册
func clearPhotos(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("\n请输入宠物ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64)
	if err != nil {
		fmt.Println("❌ 无效的ID")
		return
	}

	fmt.Print("确定要清空该宠物的相册吗? (y/n): ")
	confirm, _ := reader.ReadString('\n')
	if strings.TrimSpace(strings.ToLower(confirm)) != "y" {
		fmt.Println("⚠️  已取消")
		return
	}

	_, err = db.Exec("UPDATE pets SET photos = NULL WHERE id = ?", id)
	if err != nil {
		fmt.Println("❌ 清空失败:", err)
		return
	}
	fmt.Println("✅ 相册已清空!")
}

// 从文件批量更新
func batchUpdateFromFile(db *sql.DB, reader *bufio.Reader) {
	fmt.Println("\n文件格式说明:")
	fmt.Println("每行格式: 宠物ID,封面URL,相册URL1,相册URL2,...")
	fmt.Println("示例: 1,https://example.com/cover.jpg,https://example.com/1.jpg,https://example.com/2.jpg")
	fmt.Println()

	fmt.Print("请输入文件路径 (默认: pet_photos.csv): ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)
	if filePath == "" {
		filePath = "pet_photos.csv"
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("❌ 无法打开文件:", err)
		fmt.Println("\n提示: 请先创建文件，格式如下:")
		fmt.Println("1,https://example.com/cover1.jpg,https://example.com/album1.jpg")
		fmt.Println("2,https://example.com/cover2.jpg")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	successCount := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			fmt.Printf("⚠️  第%d行格式错误，跳过\n", lineNum)
			continue
		}

		id, err := strconv.ParseUint(strings.TrimSpace(parts[0]), 10, 64)
		if err != nil {
			fmt.Printf("⚠️  第%d行ID无效，跳过\n", lineNum)
			continue
		}

		coverPhoto := strings.TrimSpace(parts[1])

		var photos []string
		for i := 2; i < len(parts); i++ {
			url := strings.TrimSpace(parts[i])
			if url != "" {
				photos = append(photos, url)
			}
		}

		// 更新数据库
		var photosJSON *string
		if len(photos) > 0 {
			jsonBytes, _ := json.Marshal(photos)
			jsonStr := string(jsonBytes)
			photosJSON = &jsonStr
		}

		_, err = db.Exec("UPDATE pets SET cover_photo = ?, photos = ? WHERE id = ?",
			coverPhoto, photosJSON, id)
		if err != nil {
			fmt.Printf("❌ 第%d行更新失败: %v\n", lineNum, err)
			continue
		}

		fmt.Printf("✅ ID=%d 更新成功 (封面: %s, 相册: %d张)\n", id, coverPhoto, len(photos))
		successCount++
	}

	fmt.Printf("\n🎉 批量更新完成! 成功更新 %d 条记录\n", successCount)
}
