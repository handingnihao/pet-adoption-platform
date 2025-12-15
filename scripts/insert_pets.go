package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

// 配置结构
type Config struct {
	Database struct {
		MySQL struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
		} `yaml:"mysql"`
	} `yaml:"database"`
}

func loadConfig() (*Config, error) {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// 宠物数据模板
type PetTemplate struct {
	Name         string
	Type         string
	Breed        string
	Gender       string
	AgeMin       int
	AgeMax       int
	Size         string
	Colors       []string
	WeightMin    float64
	WeightMax    float64
	Characters   []string
	Descriptions []string
	CoverPhotos  []string
}

var petTemplates = map[string]PetTemplate{
	"cat": {
		Name: "猫咪",
		Type: "cat",
		Breed: "",
		Size: "small",
		Colors: []string{"橘色", "白色", "黑色", "灰色", "三花", "虎斑"},
		Characters: []string{
			"活泼好动，喜欢玩耍",
			"温顺粘人，喜欢被抚摸",
			"独立安静，适合上班族",
			"聪明伶俐，会开门",
			"胆小怕生，需要耐心",
		},
		Descriptions: []string{
			"毛发柔软顺滑，眼睛明亮有神。已完成疫苗接种和驱虫，身体健康。",
			"性格温和，不挑食，适应能力强。会使用猫砂盆，很爱干净。",
			"从小被救助，经过精心照料已完全康复。期待一个温暖的家。",
			"喜欢晒太阳和玩逗猫棒，晚上会安静地陪伴主人。",
			"与其他猫咪相处融洽，可以多猫家庭领养。",
		},
		CoverPhotos: []string{
			"https://images.unsplash.com/photo-1514888286974-6c03e2ca1dba?w=400",
			"https://images.unsplash.com/photo-1573865526739-10659fec78a5?w=400",
			"https://images.unsplash.com/photo-1495360010541-f48722b34f7d?w=400",
			"https://images.unsplash.com/photo-1518791841217-8f162f1e1131?w=400",
			"https://images.unsplash.com/photo-1574158622682-e40e69881006?w=400",
		},
	},
	"dog": {
		Name: "狗狗",
		Type: "dog",
		Breed: "",
		Size: "medium",
		Colors: []string{"金色", "黑色", "白色", "棕色", "黑白", "黄色"},
		Characters: []string{
			"忠诚护主，非常聪明",
			"活泼好动，精力充沛",
			"温顺友善，适合家庭",
			"胆大勇敢，看家好手",
			"乖巧听话，已学会基本指令",
		},
		Descriptions: []string{
			"毛发浓密有光泽，体格健壮。已完成全部疫苗接种，定期驱虫。",
			"喜欢户外活动和散步，每天需要一定运动量。会基本的坐下、握手指令。",
			"性格友善，对小朋友很温柔，是理想的家庭伴侣犬。",
			"从救助站被救回，经过训练已经很乖巧，期待找到爱它的主人。",
			"喜欢玩飞盘和球，精力充沛，适合有院子的家庭。",
		},
		CoverPhotos: []string{
			"https://images.unsplash.com/photo-1587300003388-59208cc962cb?w=400",
			"https://images.unsplash.com/photo-1561037404-61cd46aa615b?w=400",
			"https://images.unsplash.com/photo-1517849845537-4d257902454a?w=400",
			"https://images.unsplash.com/photo-1588943211346-0908a1fb0b01?w=400",
			"https://images.unsplash.com/photo-1552053831-71594a27632d?w=400",
		},
	},
	"rabbit": {
		Name: "兔子",
		Type: "rabbit",
		Breed: "",
		Size: "small",
		Colors: []string{"白色", "灰色", "棕色", "黑色", "花色"},
		Characters: []string{
			"温顺安静，适合公寓",
			"活泼好奇，喜欢探索",
			"胆小敏感，需要安静环境",
			"亲人粘人，喜欢被抱",
			"独立性强，不需要太多陪伴",
		},
		Descriptions: []string{
			"毛发蓬松柔软，耳朵直立可爱。身体健康，已做健康检查。",
			"喜欢吃新鲜蔬菜和干草，饮食习惯良好，不挑食。",
			"性格温和，不会咬人，适合有小朋友的家庭。",
			"已经适应室内生活，会使用厕所，很爱干净。",
			"喜欢在家里跳来跳去，给家庭带来很多欢乐。",
		},
		CoverPhotos: []string{
			"https://images.unsplash.com/photo-1585110396000-c9ffd4e4b308?w=400",
			"https://images.unsplash.com/photo-1535241749838-299277c6fc91?w=400",
			"https://images.unsplash.com/photo-1452857297128-d9c29adba80b?w=400",
			"https://images.unsplash.com/photo-1559214369-a6b1d7919865?w=400",
			"https://images.unsplash.com/photo-1518882605630-8b7f2be2b9e1?w=400",
		},
	},
	"hamster": {
		Name: "仓鼠",
		Type: "hamster",
		Breed: "",
		Size: "small",
		Colors: []string{"金色", "白色", "灰色", "棕色", "奶茶色"},
		Characters: []string{
			"活泼好动，喜欢跑轮",
			"胆小怕生，需要时间适应",
			"贪吃可爱，囤粮能手",
			"夜间活跃，白天睡觉",
			"温顺亲人，不咬手",
		},
		Descriptions: []string{
			"毛发柔软圆润，小巧可爱。身体健康，活力十足。",
			"喜欢囤积食物，看它塞满腮帮子非常可爱。",
			"已经适应人类，可以上手，不会咬人。",
			"需要准备跑轮和木屑，饲养成本较低。",
			"适合空间有限的家庭，是入门级的小宠物。",
		},
		CoverPhotos: []string{
			"https://images.unsplash.com/photo-1425082661705-1834bfd09dca?w=400",
			"https://images.unsplash.com/photo-1548767797-d8c844163c4c?w=400",
			"https://images.unsplash.com/photo-1612943117237-51cb8d tried?w=400",
			"https://images.unsplash.com/photo-1591382696684-38c427c7547a?w=400",
			"https://images.unsplash.com/photo-1606567595334-d39972c85dfd?w=400",
		},
	},
	"bird": {
		Name: "鸟类",
		Type: "bird",
		Breed: "",
		Size: "small",
		Colors: []string{"绿色", "黄色", "蓝色", "白色", "彩色"},
		Characters: []string{
			"活泼好动，叫声悦耳",
			"聪明伶俐，可以学说话",
			"胆小敏感，需要安静环境",
			"亲人粘人，喜欢互动",
			"独立性强，容易饲养",
		},
		Descriptions: []string{
			"羽毛鲜艳漂亮，眼睛明亮有神。身体健康，食欲良好。",
			"叫声清脆悦耳，可以给家里带来生机。",
			"已经适应家庭环境，不会乱飞乱撞。",
			"喜欢吃小米和水果，饲养简单。",
			"适合喜欢观赏性宠物的家庭。",
		},
		CoverPhotos: []string{
			"https://images.unsplash.com/photo-1552728089-57bdde30beb3?w=400",
			"https://images.unsplash.com/photo-1444464666168-49d633b86797?w=400",
			"https://images.unsplash.com/photo-1522926193341-e9ffd686c60f?w=400",
			"https://images.unsplash.com/photo-1480044965905-02098d419e96?w=400",
			"https://images.unsplash.com/photo-1591198936750-16d8e15edb9e?w=400",
		},
	},
	"other": {
		Name: "其他",
		Type: "other",
		Breed: "",
		Size: "small",
		Colors: []string{"棕色", "绿色", "黄色", "黑色", "花色"},
		Characters: []string{
			"安静温顺，容易饲养",
			"独特可爱，适合观赏",
			"互动性强，有趣好玩",
			"好奇心强，喜欢探索",
			"适应力强，不挑环境",
		},
		Descriptions: []string{
			"是一只特别的小宠物，身体健康，状态良好。",
			"饲养简单，不需要太多专业知识。",
			"适合想要尝试不同宠物的家庭。",
			"可以给生活带来不一样的乐趣。",
			"期待找到懂得欣赏它的主人。",
		},
		CoverPhotos: []string{
			"https://images.unsplash.com/photo-1597633425046-08f5110420b5?w=400",
			"https://images.unsplash.com/photo-1559253664-ca249d4608c6?w=400",
			"https://images.unsplash.com/photo-1516467508483-a7212febe31a?w=400",
			"https://images.unsplash.com/photo-1504450874802-0ba2bcd9b5ae?w=400",
			"https://images.unsplash.com/photo-1590691566903-692bf5ca7493?w=400",
		},
	},
}

var catBreeds = []string{"中华田园猫", "英短蓝猫", "美短虎斑", "布偶猫", "橘猫"}
var dogBreeds = []string{"中华田园犬", "金毛寻回犬", "拉布拉多", "柯基", "泰迪"}
var rabbitBreeds = []string{"荷兰垂耳兔", "安哥拉兔", "侏儒兔", "狮子兔", "雷克斯兔"}
var hamsterBreeds = []string{"金丝熊", "三线仓鼠", "银狐仓鼠", "布丁仓鼠", "奶茶仓鼠"}
var birdBreeds = []string{"虎皮鹦鹉", "玄凤鹦鹉", "文鸟", "珍珠鸟", "八哥"}
var otherBreeds = []string{"乌龟", "金鱼", "蜥蜴", "刺猬", "龙猫"}

var catNames = []string{"小橘", "咪咪", "花花", "团子", "豆豆"}
var dogNames = []string{"旺财", "大黄", "小黑", "毛毛", "球球"}
var rabbitNames = []string{"雪球", "棉花", "小白", "萝卜", "跳跳"}
var hamsterNames = []string{"小仓", "球球", "肉肉", "奶酪", "瓜子"}
var birdNames = []string{"小翠", "鸣鸣", "彩虹", "羽毛", "飞飞"}
var otherNames = []string{"小宝", "奇奇", "乐乐", "萌萌", "可可"}

var provinces = []string{"广东省", "浙江省", "江苏省", "北京市", "上海市"}
var cities = []string{"深圳市", "杭州市", "南京市", "北京市", "上海市"}
var districts = []string{"南山区", "西湖区", "鼓楼区", "朝阳区", "浦东新区"}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 加载配置
	config, err := loadConfig()
	if err != nil {
		fmt.Println("❌ 加载配置失败:", err)
		fmt.Println("请确保在项目根目录运行此脚本，且config/config.yaml文件存在")
		return
	}

	// 连接数据库
	mysql := config.Database.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.Database)

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
	fmt.Printf("✅ 数据库连接成功 (%s@%s:%d/%s)\n", mysql.Username, mysql.Host, mysql.Port, mysql.Database)

	// 获取用户ID
	var userID int64
	err = db.QueryRow("SELECT id FROM users LIMIT 1").Scan(&userID)
	if err != nil {
		fmt.Println("❌ 没有找到任何用户，请先创建用户")
		return
	}
	fmt.Printf("✅ 使用用户ID: %d\n\n", userID)

	// 插入宠物
	insertSQL := "INSERT INTO pets (user_id, name, type, breed, gender, age, size, color, weight, is_vaccinated, is_sterilized, health_status, description, `character`, cover_photo, province, city, district, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	typeBreeds := map[string][]string{
		"cat":     catBreeds,
		"dog":     dogBreeds,
		"rabbit":  rabbitBreeds,
		"hamster": hamsterBreeds,
		"bird":    birdBreeds,
		"other":   otherBreeds,
	}

	typeNames := map[string][]string{
		"cat":     catNames,
		"dog":     dogNames,
		"rabbit":  rabbitNames,
		"hamster": hamsterNames,
		"bird":    birdNames,
		"other":   otherNames,
	}

	typeEmoji := map[string]string{
		"cat":     "🐱",
		"dog":     "🐕",
		"rabbit":  "🐰",
		"hamster": "🐹",
		"bird":    "🐦",
		"other":   "🐾",
	}

	successCount := 0
	petTypes := []string{"cat", "dog", "rabbit", "hamster", "bird", "other"}

	for _, petType := range petTypes {
		template := petTemplates[petType]
		breeds := typeBreeds[petType]
		names := typeNames[petType]

		fmt.Printf("\n%s 创建 %s...\n", typeEmoji[petType], template.Name)

		for i := 0; i < 5; i++ {
			name := names[i]
			breed := breeds[i]
			gender := "male"
			if rand.Intn(2) == 0 {
				gender = "female"
			}
			age := rand.Intn(36) + 3 // 3-38个月
			color := template.Colors[rand.Intn(len(template.Colors))]
			
			var weight float64
			switch petType {
			case "cat":
				weight = 2.0 + rand.Float64()*6.0 // 2-8kg
			case "dog":
				weight = 5.0 + rand.Float64()*25.0 // 5-30kg
			case "rabbit":
				weight = 1.0 + rand.Float64()*3.0 // 1-4kg
			case "hamster":
				weight = 0.03 + rand.Float64()*0.07 // 30-100g
			case "bird":
				weight = 0.02 + rand.Float64()*0.1 // 20-120g
			default:
				weight = 0.5 + rand.Float64()*2.0 // 0.5-2.5kg
			}

			isVaccinated := rand.Intn(2) == 1
			isSterilized := rand.Intn(2) == 1
			healthStatus := "健康"
			character := template.Characters[i%len(template.Characters)]
			description := template.Descriptions[i%len(template.Descriptions)]
			coverPhoto := template.CoverPhotos[i%len(template.CoverPhotos)]

			locIdx := rand.Intn(len(provinces))
			province := provinces[locIdx]
			city := cities[locIdx]
			district := districts[locIdx]

			status := 1 // 已发布
			now := time.Now().Add(-time.Duration(rand.Intn(30*24)) * time.Hour)

			_, err := db.Exec(insertSQL,
				userID, name, petType, breed, gender, age, template.Size, color, weight,
				isVaccinated, isSterilized, healthStatus, description, character,
				coverPhoto, province, city, district, status, now, now,
			)

			if err != nil {
				fmt.Printf("  ❌ 创建失败 [%s]: %v\n", name, err)
				continue
			}

			genderStr := "♂"
			if gender == "female" {
				genderStr = "♀"
			}
			fmt.Printf("  ✅ %s (%s) - %s %s %d个月\n", name, breed, genderStr, color, age)
			successCount++
		}
	}

	fmt.Printf("\n🎉 完成！成功创建 %d 只宠物\n", successCount)
}
