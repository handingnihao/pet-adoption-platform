package main

import (
	"database/sql"
	"encoding/json"
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

// 社区故事数据
var stories = []struct {
	Title   string
	Content string
	Type    string
	Images  []string
}{
	{
		Title: "小橘来我家一周年啦！",
		Content: `还记得一年前的今天，我在爱心宠物领养平台上看到了小橘的照片。那时候它还在救助站里，瘦瘦小小的，眼神里带着一丝怯懦。

当时我刚搬进新家，一个人住总觉得有些冷清。看到小橘的第一眼，我就知道，它就是我要找的小伙伴。

申请领养的过程很顺利，工作人员非常负责，详细了解了我的居住环境和工作情况。一周后，我终于把小橘接回了家。

刚到家的时候，小橘躲在沙发底下不肯出来，我就静静地坐在旁边，给它时间适应。第三天，它终于主动蹭了蹭我的脚。那一刻，我知道我们的缘分开始了。

现在的小橘已经是一只8斤重的大橘猫了，每天早上准时叫我起床，晚上必须抱着我的手臂才能睡着。

感谢这个平台，让我遇见了小橘。领养代替购买，每一个生命都值得被温柔以待。🧡`,
		Type:   "story",
		Images: []string{"https://images.unsplash.com/photo-1514888286974-6c03e2ca1dba?w=600"},
	},
	{
		Title: "从流浪到被爱，豆豆的故事",
		Content: `豆豆是一只被遗弃的田园犬，被发现的时候在垃圾桶旁边翻找食物，瘦得皮包骨。

救助站的志愿者把它带回去，经过一个月的精心照料，豆豆慢慢恢复了健康。它很聪明，学会了坐下、握手、趴下等基本指令。

三个月前，我通过平台看到了豆豆的信息。视频里的它摇着尾巴，眼睛亮晶晶的，看起来特别阳光。

办完领养手续，我带豆豆回家的路上，它一直安静地趴在车后座，偶尔抬头看看我，好像在确认这一切是真的。

现在豆豆每天最开心的事就是去公园散步。它喜欢和其他狗狗玩耍，见到小朋友会主动摇尾巴。曾经流浪街头的它，现在是全小区最受欢迎的狗狗。

每个毛孩子都值得一个温暖的家。如果你有条件，请考虑领养，而不是购买。`,
		Type:   "story",
		Images: []string{"https://images.unsplash.com/photo-1587300003388-59208cc962cb?w=600"},
	},
	{
		Title: "养猫新手必看：猫咪入门指南",
		Content: `作为一个养猫三年的铲屎官，分享一些给新手的建议：

**一、基础准备**
- 猫粮：建议选择正规品牌，根据猫咪年龄选择幼猫粮或成猫粮
- 猫砂盆：封闭式或开放式都可以，重要的是每天清理
- 猫抓板：保护家具的必备品
- 猫碗和饮水机：不锈钢或陶瓷材质更健康

**二、健康注意**
- 按时接种疫苗（猫三联）
- 定期驱虫（体内外）
- 适龄绝育，对健康有益
- 定期体检，预防胜于治疗

**三、行为习惯**
- 猫咪是夜行动物，凌晨跑酷是正常的
- 猫咪喜欢高处，准备一个猫爬架
- 每天抽时间和猫咪互动，逗猫棒是好东西
- 猫咪舔毛会产生毛球，准备化毛膏

**四、禁忌食物**
- 洋葱、大蒜（会导致溶血）
- 巧克力（含有可可碱）
- 葡萄、葡萄干（可能导致肾衰竭）
- 生鱼、生肉（可能有寄生虫）

希望每一位新手铲屎官都能和猫主子幸福生活！`,
		Type:   "knowledge",
		Images: []string{},
	},
	{
		Title: "三只流浪猫变成家猫的温馨日常",
		Content: `半年前，小区里来了三只流浪猫——猫妈妈带着两只小猫。

起初我只是偶尔喂它们一些猫粮，慢慢地它们开始信任我，愿意靠近我。

冬天来了，看着它们在寒风中瑟瑟发抖，我实在不忍心。经过一番纠结，我决定把它们都带回家。

TNR（抓捕-绝育-放归）之后，我发现猫妈妈其实很亲人，两只小猫更是活泼好动。它们很快适应了室内生活。

现在，猫妈妈最喜欢的地方是阳台的猫窝，晒着太阳打盹。两只小猫整天追逐打闹，把家里搞得热闹非凡。

从流浪到家猫，它们的生活发生了翻天覆地的变化。而我，也因为它们的陪伴，生活变得更加充实。

流浪动物需要我们的关注和帮助，哪怕只是一碗水、一把猫粮，都是莫大的温暖。`,
		Type:   "story",
		Images: []string{"https://images.unsplash.com/photo-1573865526739-10659fec78a5?w=600", "https://images.unsplash.com/photo-1495360010541-f48722b34f7d?w=600"},
	},
	{
		Title: "新手养狗注意事项",
		Content: `养狗是一件幸福但也需要责任心的事情。分享一些新手必知的要点：

**疫苗接种**
- 幼犬：45天后开始接种，共3针+狂犬
- 成犬：每年加强一次
- 打完疫苗前不要带出门、不要洗澡

**饮食管理**
- 幼犬一天3-4餐，成犬一天2餐
- 狗粮要根据体型和年龄选择
- 不能吃的：巧克力、葡萄、洋葱、木糖醇

**日常护理**
- 定期洗澡（2-4周一次）
- 每天梳毛，保持毛发整洁
- 定期修剪指甲
- 清洁耳朵和牙齿

**运动需求**
- 每天至少30分钟户外活动
- 大型犬需要更多运动量
- 遛狗一定要牵绑带

**训练建议**
- 从基础指令开始：坐、趴、等待
- 奖励为主，避免体罚
- 保持耐心，狗狗都能学会

祝大家都能成为合格的铲屎官！`,
		Type:   "knowledge",
		Images: []string{},
	},
	{
		Title: "记录和布丁的第100天",
		Content: `今天是布丁来我家的第100天，必须记录一下这个特别的日子！

布丁是一只英短蓝猫，之前的主人因为工作调动无法继续养它，通过平台找到了我。

刚接回来的时候，布丁有些认生，总是躲在床底下。我就每天坐在床边，轻声和它说话，给它最喜欢的冻干。

大概一周后，布丁终于愿意出来了。第一次主动跳上我的膝盖时，我激动得差点哭出来。

现在的布丁简直是个小霸王：
🐱 早上6点准时踩脸叫醒我
🐱 吃饭必须人陪着
🐱 看电视必须坐在我旁边
🐱 睡觉必须挨着我

虽然它霸道，但我乐在其中。它治愈了我独居的孤独，让家变成了真正的家。

感恩遇见你，布丁！这是我们的第100天，未来还有无数个100天在等着我们。💙`,
		Type:   "daily",
		Images: []string{"https://images.unsplash.com/photo-1574158622682-e40e69881006?w=600"},
	},
	{
		Title: "夏季宠物防暑小贴士",
		Content: `夏天到了，如何帮助毛孩子们度过炎热的夏季呢？

**🐕 狗狗篇**
1. 避免高温时段遛狗，选择早晚凉爽时
2. 出门带水，随时补充水分
3. 不要把狗狗单独留在车内！！
4. 可以适当剃毛，但不要剃光
5. 准备凉垫或冰垫

**🐱 猫咪篇**
1. 保持室内通风，必要时开空调
2. 多放几个水碗，鼓励喝水
3. 长毛猫可以修剪腹部毛发
4. 避免阳光直射的地方
5. 冰块加水可以降温

**⚠️ 中暑症状**
- 呼吸急促、喘气
- 流口水增多
- 精神萎靡
- 体温升高
- 牙龈发红或发白

如果发现中暑，立即转移到阴凉处，用湿毛巾擦拭身体（不要用冰水），并尽快送医。

让我们一起保护毛孩子，安全度夏！`,
		Type:   "knowledge",
		Images: []string{},
	},
	{
		Title: "遇见你是最美的意外——我和旺财的故事",
		Content: `从来没想过我会养狗，直到遇见旺财。

那是一个下雨天，我在回家路上看到一只浑身湿透的小狗蜷缩在路边。它那可怜巴巴的眼神，让我无法视而不见。

本来只是想带它躲躲雨，没想到这一躲，就躲进了我的生活。

带它去医院检查，医生说它营养不良，可能流浪了很久。我决定先养着，等它康复了再找领养。

结果，根本舍不得送走。

旺财很聪明，一个星期就学会了定点上厕所。它还特别粘人，我去哪它跟哪，上厕所都要守在门口。

有一次我加班到很晚，回家时发现它就趴在门口等我，看到我开门，尾巴摇得像螺旋桨一样。那一刻，我的心都化了。

现在旺财是我最好的朋友。我们一起散步、一起看日落、一起宅家。它让我明白，有时候，最美的相遇都是意外。

如果你也想养一只小动物，不妨给流浪的毛孩子一个机会。也许你也会收获一份意想不到的温暖。`,
		Type:   "story",
		Images: []string{"https://images.unsplash.com/photo-1561037404-61cd46aa615b?w=600"},
	},
	{
		Title: "周末带毛孩子去了宠物友好咖啡馆",
		Content: `发现了一家超棒的宠物友好咖啡馆！必须分享给大家！

店名就不说了（怕广告嫌疑），位置在市中心，装修很温馨。

店里有专门的宠物活动区，地面铺的是防滑的软木地板，很贴心。还有给毛孩子喝的专用饮水点。

我带着我家二哈去的，本来担心它会拆家，没想到它遇到了几个小伙伴，玩得不亦乐乎。

咖啡和甜点也很好吃，我点了一杯拿铁和提拉米苏。店员还送了我家狗子一个小饼干，它开心坏了。

整个下午，我一边喝咖啡一边看它们玩耍，简直是最幸福的时光。

现在越来越多的地方对宠物友好了，希望这样的地方越来越多！毕竟，毛孩子也是家人啊。

有带毛孩子出去玩的经历吗？评论区分享一下吧！`,
		Type:   "daily",
		Images: []string{},
	},
	{
		Title: "三年后，我终于鼓起勇气再养一只猫",
		Content: `三年前，我的猫咪因病离开了我。那之后，我发誓再也不养猫了。

可是最近，在平台上看到一只和它长得很像的猫咪，一只橘白色的小猫，正等待领养。

我纠结了很久，害怕再经历一次失去的痛苦。但我也知道，还有那么多毛孩子在等待一个家。

思来想去，我决定去见见它。

见面的那一刻，小猫就像认识我一样，直接跳到了我怀里，呼噜呼噜叫个不停。

我想，这就是缘分吧。

把它带回家后，我给它取名叫"小福"，希望它福气满满。

小福很快就适应了新家，它喜欢窝在之前那只猫最爱的窗台上晒太阳。

我不知道它是不是它的转世，但我相信，它带着爱来到我身边。

生命的意义，也许就是不断地爱，不断地被爱。

感谢这个平台，让我有勇气再次打开心扉。❤️`,
		Type:   "story",
		Images: []string{"https://images.unsplash.com/photo-1518791841217-8f162f1e1131?w=600"},
	},
}

func main() {
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

	// 获取用户ID（用于发布帖子）
	var userID int64
	err = db.QueryRow("SELECT id FROM users WHERE role = 'user' LIMIT 1").Scan(&userID)
	if err != nil {
		fmt.Println("❌ 获取用户失败，尝试使用管理员账户")
		err = db.QueryRow("SELECT id FROM users LIMIT 1").Scan(&userID)
		if err != nil {
			fmt.Println("❌ 没有找到任何用户，请先创建用户")
			return
		}
	}
	fmt.Printf("✅ 使用用户ID: %d\n\n", userID)

	// 随机打乱顺序
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(stories), func(i, j int) {
		stories[i], stories[j] = stories[j], stories[i]
	})

	// 插入帖子
	insertSQL := `INSERT INTO posts (user_id, title, content, images, type, status, view_count, like_count, comment_count, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, 1, ?, ?, ?, ?, ?)`

	successCount := 0
	baseTime := time.Now()

	for i, story := range stories {
		// 随机生成一些互动数据
		viewCount := rand.Intn(500) + 50
		likeCount := rand.Intn(viewCount/5) + 5
		commentCount := rand.Intn(likeCount/2) + 1

		// 时间往前推，让帖子有不同的发布时间
		createdAt := baseTime.Add(-time.Duration(i*24+rand.Intn(48)) * time.Hour)

		// 转换图片为JSON
		imagesJSON, _ := json.Marshal(story.Images)

		_, err := db.Exec(insertSQL,
			userID,
			story.Title,
			story.Content,
			string(imagesJSON),
			story.Type,
			viewCount,
			likeCount,
			commentCount,
			createdAt,
			createdAt,
		)

		if err != nil {
			fmt.Printf("❌ 插入失败 [%s]: %v\n", story.Title, err)
			continue
		}

		typeEmoji := map[string]string{
			"story":     "📖",
			"knowledge": "📚",
			"daily":     "☀️",
		}
		fmt.Printf("%s 已添加: %s (浏览:%d 点赞:%d 评论:%d)\n",
			typeEmoji[story.Type], story.Title, viewCount, likeCount, commentCount)
		successCount++
	}

	fmt.Printf("\n🎉 完成！成功添加 %d 篇社区故事\n", successCount)
}
