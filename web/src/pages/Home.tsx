import { Link } from 'react-router-dom'
import { PawPrint, Heart, Users, Search, ArrowRight } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Card, CardContent } from '../components/ui/card'

export function Home() {
  const features = [
    {
      icon: Search,
      title: '智能匹配',
      description: '根据您的生活方式和偏好，为您推荐最适合的宠物伙伴。'
    },
    {
      icon: Heart,
      title: '爱心领养',
      description: '给流浪动物一个温暖的家，让爱传递下去。'
    },
    {
      icon: Users,
      title: '社区交流',
      description: '与其他宠物爱好者分享经验，获取养宠建议。'
    }
  ]

  const petTypes = [
    { name: '猫咪', image: '🐱', count: 128 },
    { name: '狗狗', image: '🐕', count: 256 },
    { name: '兔子', image: '🐰', count: 45 },
    { name: '仓鼠', image: '🐹', count: 32 },
    { name: '鸟类', image: '🐦', count: 28 },
    { name: '其他', image: '🐾', count: 15 }
  ]

  return (
    <div className="flex flex-col">
      {/* Hero Section */}
      <section className="relative bg-gradient-to-br from-primary/10 via-background to-secondary/10 py-20 lg:py-32">
        <div className="container">
          <div className="grid lg:grid-cols-2 gap-12 items-center">
            <div className="space-y-8">
              <div className="inline-flex items-center px-4 py-2 bg-primary/10 rounded-full text-sm text-primary">
                <PawPrint className="h-4 w-4 mr-2" />
                让爱找到归宿
              </div>
              <h1 className="text-4xl lg:text-6xl font-bold leading-tight">
                给毛孩子一个
                <span className="text-primary">温暖的家</span>
              </h1>
              <p className="text-lg text-muted-foreground max-w-lg">
                我们致力于帮助流浪动物找到爱它们的主人。在这里，每一个生命都值得被温柔以待。
              </p>
              <div className="flex flex-col sm:flex-row gap-4">
                <Button size="lg" asChild>
                  <Link to="/pets">
                    开始领养
                    <ArrowRight className="ml-2 h-4 w-4" />
                  </Link>
                </Button>
                <Button size="lg" variant="outline" asChild>
                  <Link to="/guide">了解更多</Link>
                </Button>
              </div>
              <div className="flex items-center gap-8 pt-4">
                <div>
                  <div className="text-3xl font-bold text-primary">500+</div>
                  <div className="text-sm text-muted-foreground">待领养宠物</div>
                </div>
                <div>
                  <div className="text-3xl font-bold text-primary">1200+</div>
                  <div className="text-sm text-muted-foreground">成功领养</div>
                </div>
                <div>
                  <div className="text-3xl font-bold text-primary">50+</div>
                  <div className="text-sm text-muted-foreground">合作机构</div>
                </div>
              </div>
            </div>
            <div className="relative hidden lg:block">
              <div className="absolute inset-0 bg-gradient-to-br from-primary/20 to-secondary/20 rounded-3xl blur-3xl" />
              <div className="relative bg-gradient-to-br from-primary/5 to-secondary/5 rounded-3xl p-8 border">
                <div className="grid grid-cols-2 gap-4">
                  <div className="aspect-square rounded-2xl bg-muted flex items-center justify-center text-6xl">
                    🐱
                  </div>
                  <div className="aspect-square rounded-2xl bg-muted flex items-center justify-center text-6xl">
                    🐕
                  </div>
                  <div className="aspect-square rounded-2xl bg-muted flex items-center justify-center text-6xl">
                    🐰
                  </div>
                  <div className="aspect-square rounded-2xl bg-muted flex items-center justify-center text-6xl">
                    🐹
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Pet Types */}
      <section className="py-16 bg-muted/30">
        <div className="container">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold mb-4">寻找你的新伙伴</h2>
            <p className="text-muted-foreground">选择你感兴趣的宠物类型</p>
          </div>
          <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4">
            {petTypes.map((type) => (
              <Link
                key={type.name}
                to={`/pets?type=${type.name}`}
                className="group"
              >
                <Card className="hover:border-primary transition-colors">
                  <CardContent className="p-6 text-center">
                    <div className="text-5xl mb-3">{type.image}</div>
                    <div className="font-medium group-hover:text-primary transition-colors">
                      {type.name}
                    </div>
                    <div className="text-sm text-muted-foreground">
                      {type.count} 只待领养
                    </div>
                  </CardContent>
                </Card>
              </Link>
            ))}
          </div>
        </div>
      </section>

      {/* Features */}
      <section className="py-16">
        <div className="container">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold mb-4">为什么选择我们</h2>
            <p className="text-muted-foreground">我们提供全方位的领养服务</p>
          </div>
          <div className="grid md:grid-cols-3 gap-8">
            {features.map((feature) => (
              <Card key={feature.title} className="border-0 shadow-lg">
                <CardContent className="p-8 text-center">
                  <div className="w-16 h-16 mx-auto mb-6 bg-primary/10 rounded-2xl flex items-center justify-center">
                    <feature.icon className="h-8 w-8 text-primary" />
                  </div>
                  <h3 className="text-xl font-semibold mb-3">{feature.title}</h3>
                  <p className="text-muted-foreground">{feature.description}</p>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* CTA */}
      <section className="py-16 bg-primary text-primary-foreground">
        <div className="container text-center">
          <h2 className="text-3xl font-bold mb-4">准备好迎接新成员了吗？</h2>
          <p className="text-primary-foreground/80 mb-8 max-w-2xl mx-auto">
            每一只等待领养的宠物都在期盼着一个温暖的家。开始您的领养之旅，让爱改变生命。
          </p>
          <Button size="lg" variant="secondary" asChild>
            <Link to="/pets">
              浏览待领养宠物
              <ArrowRight className="ml-2 h-4 w-4" />
            </Link>
          </Button>
        </div>
      </section>
    </div>
  )
}
