import { Heart, Users, Target, Mail, Github, MessageCircle } from 'lucide-react'
import { Card, CardContent } from '../components/ui/card'

// 团队成员信息 - 可自行修改
const teamMembers = [
  {
    name: '唐非',
    role: '总负责人与后端开发',
    avatar: 'https://page-image.tos-cn-guangzhou.volces.com/hajimi-project/tf.jpg', // 头像URL，留空使用默认
    description: `项目整体规划与协调\n后端架构设计与开发`,
  },
  {
    name: '梁展图',
    role: '运维与后端开发',
    avatar: 'https://page-image.tos-cn-guangzhou.volces.com/hajimi-project/lzt.jpg',
    description: '后端开发\n服务器运维',
  },
  {
    name: '王烨枫',
    role: '前端与静态资源维护',
    avatar: 'https://page-image.tos-cn-guangzhou.volces.com/hajimi-project/wyf.jpg',
    description: '前端界面开发\nCOS静态资源维护',
  },
  {
      name: '林訉毅',
      role: '前端与项目设计',
      avatar: 'https://page-image.tos-cn-guangzhou.volces.com/hajimi-project/lfy.jpg',
      description: '前后端联调与接口测试\n前端界面设计',
  },
]

// 项目里程碑 - 可自行修改
const milestones = [
  { date: '2025年12月8日', title: '项目启动', description: '确定项目方向，组建团队' },
  { date: '2025年12月9日', title: '需求分析', description: '完成需求调研与技术选型' },
  { date: '2025年12月9日-15日', title: '开发阶段', description: '核心功能开发与测试' },
  { date: '2025年12月', title: '项目上线', description: '完成部署，正式发布' },
]

export function About() {
  const handleEmailClick = () => {
    const email = 'tyumaoo@qq.com'
    // 复制到剪贴板
    navigator.clipboard.writeText(email).then(() => {
      alert(`联系邮箱：${email}\n已复制到剪贴板`)
    }).catch(() => {
      alert(`联系邮箱：${email}`)
    })
  }

  return (
    <div className="min-h-screen">
      {/* Hero Section */}
      <section className="relative bg-gradient-to-br from-primary/10 via-primary/5 to-background py-20">
        <div className="container">
          <div className="max-w-3xl mx-auto text-center">
            <div className="inline-flex items-center justify-center w-20 h-20 rounded-full bg-primary/10 mb-6">
              <Heart className="w-10 h-10 text-primary" />
            </div>
            <h1 className="text-4xl md:text-5xl font-bold mb-6">
              关于我们
            </h1>
            <p className="text-xl text-muted-foreground leading-relaxed">
              我们是一支充满热情的初创团队，致力于通过技术手段
              <br className="hidden md:block" />
              让每一个流浪的小生命都能找到温暖的家
            </p>
          </div>
        </div>
        
        {/* 装饰元素 */}
        <div className="absolute top-10 left-10 w-20 h-20 bg-primary/5 rounded-full blur-2xl" />
        <div className="absolute bottom-10 right-10 w-32 h-32 bg-primary/5 rounded-full blur-3xl" />
      </section>

      {/* 项目介绍 */}
      <section className="py-16 bg-background">
        <div className="container">
          <div className="max-w-4xl mx-auto">
            <div className="grid md:grid-cols-2 gap-8">
              <Card className="border-none shadow-lg">
                <CardContent className="p-8">
                  <div className="w-12 h-12 rounded-lg bg-blue-100 flex items-center justify-center mb-4">
                    <Target className="w-6 h-6 text-blue-600" />
                  </div>
                  <h3 className="text-xl font-bold mb-3">项目愿景</h3>
                  <p className="text-muted-foreground leading-relaxed">
                    打造一个便捷、安全、透明的宠物领养平台，连接爱心人士与待领养宠物，
                    推广"领养代替购买"的理念，减少流浪动物数量。
                  </p>
                </CardContent>
              </Card>
              
              <Card className="border-none shadow-lg">
                <CardContent className="p-8">
                  <div className="w-12 h-12 rounded-lg bg-green-100 flex items-center justify-center mb-4">
                    <Users className="w-6 h-6 text-green-600" />
                  </div>
                  <h3 className="text-xl font-bold mb-3">我们的使命</h3>
                  <p className="text-muted-foreground leading-relaxed">
                    通过技术创新，降低领养门槛，提高领养效率，
                    为每一只流浪动物创造被爱的机会，为每一位爱宠人士提供可靠的领养渠道。
                  </p>
                </CardContent>
              </Card>
            </div>
          </div>
        </div>
      </section>

      {/* 团队成员 */}
      <section className="py-16 bg-muted/30">
        <div className="container">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold mb-4">团队成员</h2>
            <p className="text-muted-foreground">一群热爱技术、热爱动物的小伙伴</p>
          </div>
          
          <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-6 max-w-5xl mx-auto">
            {teamMembers.map((member, index) => (
              <Card key={index} className="border-none shadow-md hover:shadow-lg transition-shadow">
                <CardContent className="p-6 text-center">
                  <div className="w-20 h-20 rounded-full bg-gradient-to-br from-primary/20 to-primary/10 flex items-center justify-center mx-auto mb-4">
                    {member.avatar ? (
                      <img src={member.avatar} alt={member.name} className="w-20 h-20 rounded-full object-cover" />
                    ) : (
                      <span className="text-2xl font-bold text-primary">
                        {member.name[0]}
                      </span>
                    )}
                  </div>
                  <h3 className="font-bold text-lg mb-1">{member.name}</h3>
                  <p className="text-primary text-sm mb-2">{member.role}</p>
                  <p className="text-muted-foreground text-sm whitespace-pre-line">{member.description}</p>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* 项目历程 */}
      <section className="py-16 bg-background">
        <div className="container">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold mb-4">项目历程</h2>
            <p className="text-muted-foreground">从想法到现实的每一步</p>
          </div>
          
          <div className="max-w-3xl mx-auto">
            <div className="relative">
              {/* 时间线 */}
              <div className="absolute left-4 md:left-1/2 top-0 bottom-0 w-0.5 bg-primary/20 transform md:-translate-x-1/2" />
              
              {milestones.map((milestone, index) => (
                <div key={index} className={`relative flex items-start mb-8 ${index % 2 === 0 ? 'md:flex-row-reverse' : ''}`}>
                  {/* 时间点 */}
                  <div className="absolute left-4 md:left-1/2 w-3 h-3 bg-primary rounded-full transform -translate-x-1/2 mt-1.5 ring-4 ring-background" />
                  
                  {/* 内容 */}
                  <div className={`ml-12 md:ml-0 md:w-1/2 ${index % 2 === 0 ? 'md:pr-12 md:text-right' : 'md:pl-12'}`}>
                    <Card className="border-none shadow-sm hover:shadow-md transition-shadow inline-block">
                      <CardContent className="p-4">
                        <span className="text-sm text-primary font-medium">{milestone.date}</span>
                        <h4 className="font-bold mt-1">{milestone.title}</h4>
                        <p className="text-sm text-muted-foreground mt-1">{milestone.description}</p>
                      </CardContent>
                    </Card>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* 技术栈 */}
      <section className="py-16 bg-muted/30">
        <div className="container">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold mb-4">技术栈</h2>
            <p className="text-muted-foreground">我们使用的核心技术</p>
          </div>
          
          <div className="grid grid-cols-2 md:grid-cols-4 gap-4 max-w-3xl mx-auto">
            {[
              { name: 'React', category: '前端框架' },
              { name: 'TypeScript', category: '开发语言' },
              { name: 'TailwindCSS', category: 'UI样式' },
              { name: 'Go', category: '后端语言' },
              { name: 'Gin', category: 'Web框架' },
              { name: 'MySQL', category: '数据库' },
              { name: 'Redis', category: '缓存' },
              { name: 'Docker', category: '容器化' },
            ].map((tech, index) => (
              <div key={index} className="bg-background rounded-lg p-4 text-center shadow-sm hover:shadow-md transition-shadow">
                <p className="font-bold">{tech.name}</p>
                <p className="text-xs text-muted-foreground mt-1">{tech.category}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* 联系我们 */}
      <section className="py-16 bg-background">
        <div className="container">
          <div className="max-w-2xl mx-auto text-center">
            <h2 className="text-3xl font-bold mb-4">联系我们</h2>
            <p className="text-muted-foreground mb-8">
              如有任何问题或建议，欢迎通过以下方式联系我们
            </p>
            
            <div className="flex flex-wrap justify-center gap-4">
              <button
                onClick={handleEmailClick}
                className="inline-flex items-center gap-2 px-6 py-3 bg-primary text-primary-foreground rounded-lg hover:bg-primary/90 transition-colors"
              >
                <Mail className="w-5 h-5" />
                邮箱联系
              </button>
              <a
                href="https://gitee.com/olrain/pet-adoption-platform"
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center gap-2 px-6 py-3 bg-muted text-foreground rounded-lg hover:bg-muted/80 transition-colors"
              >
                <Github className="w-5 h-5" />
                Gitee 仓库
              </a>
              <a
                href="#"
                className="inline-flex items-center gap-2 px-6 py-3 bg-muted text-foreground rounded-lg hover:bg-muted/80 transition-colors"
              >
                <MessageCircle className="w-5 h-5" />
                在线反馈
              </a>
            </div>
          </div>
        </div>
      </section>

      {/* 底部装饰 */}
      <section className="py-12 bg-gradient-to-t from-primary/5 to-background">
        <div className="container text-center">
          <p className="text-muted-foreground">
            🐾 用爱心连接每一个生命 · 领养代替购买 🐾
          </p>
        </div>
      </section>
    </div>
  )
}
