package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pet-adoption-platform/config"
)

func main() {
	// 使用config包加载配置
	if err := config.Init(); err != nil {
		log.Fatal("加载配置失败:", err)
	}
	
	// 获取DSN
	dsn := config.AppConfig.Database.MySQL.GetDSN()
	fmt.Println("连接字符串:", dsn)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	fmt.Println("✓ 数据库连接成功")

	// 1. 删除旧表（如果存在）
	fmt.Println("\n[1] 删除旧表...")
	// 先删除外键约束
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	// 删除所有可能引用pets的表
	db.Exec("DROP TABLE IF EXISTS adoption_applications")
	db.Exec("DROP TABLE IF EXISTS adoptions")
	db.Exec("DROP TABLE IF EXISTS favorites")
	db.Exec("DROP TABLE IF EXISTS pets")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	fmt.Println("✓ 旧表已删除")

	// 2. 创建新表
	fmt.Println("\n[2] 创建pets表...")
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")  // 创建时也关闭外键检查
	createTableSQL := "CREATE TABLE pets (" +
		"id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY," +
		"name VARCHAR(100) NOT NULL COMMENT '宠物名称'," +
		"type VARCHAR(20) NOT NULL COMMENT '宠物类型'," +
		"breed VARCHAR(100) COMMENT '品种'," +
		"gender VARCHAR(10) COMMENT '性别'," +
		"age INT COMMENT '年龄(月)'," +
		"size VARCHAR(20) COMMENT '体型'," +
		"color VARCHAR(50) COMMENT '毛色'," +
		"weight DECIMAL(5,2) COMMENT '体重(kg)'," +
		"is_vaccinated BOOLEAN DEFAULT FALSE COMMENT '是否接种疫苗'," +
		"is_sterilized BOOLEAN DEFAULT FALSE COMMENT '是否绝育'," +
		"health_status VARCHAR(200) COMMENT '健康状况'," +
		"description TEXT COMMENT '详细描述'," +
		"`character` TEXT COMMENT '性格特点'," +
		"photos TEXT COMMENT '照片URL'," +
		"cover_photo VARCHAR(500) COMMENT '封面图'," +
		"province VARCHAR(50) COMMENT '省份'," +
		"city VARCHAR(50) COMMENT '城市'," +
		"district VARCHAR(50) COMMENT '区县'," +
		"address VARCHAR(200) COMMENT '详细地址'," +
		"user_id BIGINT NOT NULL COMMENT '发布用户ID'," +
		"status TINYINT DEFAULT 0 COMMENT '状态'," +
		"view_count INT DEFAULT 0 COMMENT '浏览次数'," +
		"favorite_count INT DEFAULT 0 COMMENT '收藏次数'," +
		"adopted_by BIGINT COMMENT '领养人ID'," +
		"adopted_at DATETIME COMMENT '领养时间'," +
		"reviewed_by BIGINT COMMENT '审核人ID'," +
		"reviewed_at DATETIME COMMENT '审核时间'," +
		"reject_reason VARCHAR(500) COMMENT '拒绝原因'," +
		"created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'," +
		"updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'," +
		"deleted_at DATETIME COMMENT '删除时间'," +
		"INDEX (user_id)," +
		"INDEX (status)," +
		"INDEX (city)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci"
	if err := db.Exec(createTableSQL).Error; err != nil {
		log.Fatal("创建表失败:", err)
	}
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")  // 恢复外键检查
	fmt.Println("✓ pets表创建成功")

	// 3. 插入测试数据
	fmt.Println("\n[3] 插入测试数据...")
	insertSQL := `
INSERT INTO pets (name, type, breed, gender, age, size, color, weight, is_vaccinated, is_sterilized, health_status, description, ` + "`character`" + `, cover_photo, province, city, district, address, user_id, status) VALUES
('小白', 'dog', '金毛寻回犬', 'male', 18, 'large', '金黄色', 28.50, TRUE, FALSE, '健康状况良好，已完成所有疫苗接种', '小白是一只温顺可爱的金毛，非常喜欢和人玩耍，性格活泼开朗。适合有院子的家庭饲养。', '温顺、活泼、聪明、忠诚', 'https://example.com/pets/golden-retriever.jpg', '广东省', '深圳市', '南山区', '科技园', 2, 1),
('咪咪', 'cat', '英国短毛猫', 'female', 12, 'medium', '银渐层', 4.20, TRUE, TRUE, '健康，已绝育，已接种疫苗', '咪咪是一只可爱的银渐层英短，性格温和，喜欢安静的环境。非常适合公寓饲养。', '安静、温和、独立、爱干净', 'https://example.com/pets/british-shorthair.jpg', '广东省', '深圳市', '福田区', 'CBD', 2, 1),
('旺财', 'dog', '中华田园犬', 'male', 24, 'medium', '黄白相间', 15.00, TRUE, TRUE, '健康活泼', '旺财是一只非常忠诚的田园犬，看家护院的好帮手。已经两岁了，非常懂事。', '忠诚、勇敢、机警、看家好手', 'https://example.com/pets/chinese-rural-dog.jpg', '广东省', '深圳市', '龙岗区', '布吉', 2, 1),
('雪球', 'rabbit', '安哥拉兔', 'female', 6, 'small', '纯白色', 2.50, TRUE, FALSE, '健康', '雪球是一只纯白色的安哥拉兔，毛发柔软，非常可爱。性格温顺，适合陪伴。', '温顺、可爱、安静', 'https://example.com/pets/angora-rabbit.jpg', '广东省', '深圳市', '宝安区', '西乡', 4, 1),
('虎妞', 'cat', '橘猫', 'female', 8, 'medium', '橘色', 5.80, TRUE, FALSE, '健康，食欲旺盛', '虎妞是一只可爱的橘猫，性格活泼好动，喜欢吃东西。需要主人有耐心陪伴。', '活泼、贪吃、粘人', 'https://example.com/pets/orange-cat.jpg', '广东省', '深圳市', '南山区', '后海', 2, 0)
`
	if err := db.Exec(insertSQL).Error; err != nil {
		log.Fatal("插入数据失败:", err)
	}
	fmt.Println("✓ 测试数据插入成功")

	// 4. 验证数据
	fmt.Println("\n[4] 验证数据...")
	var count int64
	db.Raw("SELECT COUNT(*) FROM pets").Scan(&count)
	fmt.Printf("✓ 当前宠物总数: %d\n", count)

	// 5. 显示数据
	fmt.Println("\n[5] 宠物列表:")
	type PetInfo struct {
		ID     uint   `gorm:"column:id"`
		Name   string `gorm:"column:name"`
		Type   string `gorm:"column:type"`
		Breed  string `gorm:"column:breed"`
		City   string `gorm:"column:city"`
		Status int    `gorm:"column:status"`
	}
	var pets []PetInfo
	db.Raw("SELECT id, name, type, breed, city, status FROM pets ORDER BY id").Scan(&pets)
	
	fmt.Println("\nID\t名称\t类型\t品种\t\t城市\t\t状态")
	fmt.Println("-------------------------------------------------------")
	for _, pet := range pets {
		statusText := "待审核"
		if pet.Status == 1 {
			statusText = "可领养"
		} else if pet.Status == 2 {
			statusText = "已领养"
		}
		fmt.Printf("%d\t%s\t%s\t%s\t\t%s\t%s\n", pet.ID, pet.Name, pet.Type, pet.Breed, pet.City, statusText)
	}

	fmt.Println("\n✅ 数据库设置完成！")
	fmt.Println("现在可以测试宠物API了")
}
