import { 
  Heart, CheckCircle, Home, Clock, FileText, Phone, 
  AlertCircle, PawPrint, Stethoscope, ShieldCheck,
  Users, MapPin, Calendar, BadgeCheck
} from 'lucide-react'
import { Card, CardContent, CardHeader, CardTitle } from '../components/ui/card'
import { Button } from '../components/ui/button'
import { Link } from 'react-router-dom'

export function Guide() {
  return (
    <div className="min-h-screen">
      {/* Hero Section */}
      <section className="bg-gradient-to-br from-primary/10 via-background to-secondary/10 py-16">
        <div className="container text-center">
          <div className="inline-flex items-center px-4 py-2 bg-primary/10 rounded-full text-sm text-primary mb-6">
            <Heart className="h-4 w-4 mr-2" />
            领养代替购买
          </div>
          <h1 className="text-4xl font-bold mb-4">领养指南</h1>
          <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
            感谢您选择领养！这份指南将帮助您了解领养流程，做好充分准备，给毛孩子一个温暖的家。
          </p>
        </div>
      </section>

      <div className="container py-12 space-y-16">
        {/* 领养前的准备 */}
        <section>
          <div className="flex items-center gap-3 mb-8">
            <div className="p-3 bg-primary/10 rounded-xl">
              <CheckCircle className="h-6 w-6 text-primary" />
            </div>
            <div>
              <h2 className="text-2xl font-bold">领养前的准备</h2>
              <p className="text-muted-foreground">在决定领养之前，请确保您已做好以下准备</p>
            </div>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2 text-lg">
                  <Home className="h-5 w-5 text-primary" />
                  居住环境
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-3 text-sm">
                <p>• 确保您的住所允许饲养宠物</p>
                <p>• 有足够的空间供宠物活动</p>
                <p>• 准备好宠物的专属区域（休息区、饮食区）</p>
                <p>• 检查并消除家中的安全隐患（电线、有毒植物等）</p>
                <p>• 如果是租房，需获得房东同意</p>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2 text-lg">
                  <Clock className="h-5 w-5 text-primary" />
                  时间投入
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-3 text-sm">
                <p>• 每天有足够时间陪伴宠物</p>
                <p>• 狗狗需要每天遛弯1-2小时</p>
                <p>• 猫咪需要每天互动玩耍时间</p>
                <p>• 定期清理、洗澡、梳毛等护理</p>
                <p>• 出差或旅行时能妥善安排照顾</p>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2 text-lg">
                  <Users className="h-5 w-5 text-primary" />
                  家庭共识
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-3 text-sm">
                <p>• 所有家庭成员同意养宠</p>
                <p>• 确认家人无宠物过敏</p>
                <p>• 明确日常照顾责任分工</p>
                <p>• 家中有老人或小孩需特别考虑</p>
                <p>• 已有宠物需考虑兼容性</p>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2 text-lg">
                  <Stethoscope className="h-5 w-5 text-primary" />
                  经济准备
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-3 text-sm">
                <p>• 日常食物费用：200-500元/月</p>
                <p>• 定期疫苗和驱虫：500-1000元/年</p>
                <p>• 日用品（猫砂、玩具等）</p>
                <p>• 预留医疗应急资金</p>
                <p>• 绝育手术费用（如未绝育）</p>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2 text-lg">
                  <ShieldCheck className="h-5 w-5 text-primary" />
                  心理准备
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-3 text-sm">
                <p>• 宠物平均寿命10-20年，是长期承诺</p>
                <p>• 做好应对掉毛、异味等问题的准备</p>
                <p>• 理解宠物可能有行为问题需耐心纠正</p>
                <p>• 接受宠物生老病死是自然规律</p>
                <p>• 不因任何原因遗弃宠物</p>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2 text-lg">
                  <PawPrint className="h-5 w-5 text-primary" />
                  必备物品
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-3 text-sm">
                <p>• 食盆和水盆</p>
                <p>• 优质宠物粮</p>
                <p>• 舒适的窝/垫子</p>
                <p>• 牵引绳和项圈（狗狗）</p>
                <p>• 猫砂盆和猫砂（猫咪）</p>
                <p>• 玩具和磨爪器</p>
              </CardContent>
            </Card>
          </div>
        </section>

        {/* 领养流程 */}
        <section>
          <div className="flex items-center gap-3 mb-8">
            <div className="p-3 bg-primary/10 rounded-xl">
              <FileText className="h-6 w-6 text-primary" />
            </div>
            <div>
              <h2 className="text-2xl font-bold">领养流程</h2>
              <p className="text-muted-foreground">简单几步，开启您的领养之旅</p>
            </div>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6">
            {[
              { step: 1, title: '浏览宠物', desc: '在平台浏览待领养的宠物，根据您的喜好和条件筛选合适的伙伴', icon: PawPrint },
              { step: 2, title: '提交申请', desc: '找到心仪的宠物后，填写领养申请表，详细说明您的领养意愿和条件', icon: FileText },
              { step: 3, title: '审核沟通', desc: '工作人员会审核您的申请，可能会电话或上门回访确认情况', icon: Phone },
              { step: 4, title: '签署协议', desc: '审核通过后，签署领养协议，支付必要费用（如疫苗费），带宠物回家', icon: BadgeCheck },
            ].map((item) => (
              <Card key={item.step} className="relative overflow-hidden">
                <div className="absolute top-4 right-4 text-6xl font-bold text-primary/10">
                  {item.step}
                </div>
                <CardContent className="pt-6">
                  <div className="p-3 bg-primary/10 rounded-xl w-fit mb-4">
                    <item.icon className="h-6 w-6 text-primary" />
                  </div>
                  <h3 className="font-bold text-lg mb-2">{item.title}</h3>
                  <p className="text-sm text-muted-foreground">{item.desc}</p>
                </CardContent>
              </Card>
            ))}
          </div>
        </section>

        {/* 领养条件 */}
        <section>
          <div className="flex items-center gap-3 mb-8">
            <div className="p-3 bg-primary/10 rounded-xl">
              <BadgeCheck className="h-6 w-6 text-primary" />
            </div>
            <div>
              <h2 className="text-2xl font-bold">领养条件</h2>
              <p className="text-muted-foreground">为确保宠物得到良好照顾，领养人需满足以下基本条件</p>
            </div>
          </div>

          <div className="grid md:grid-cols-2 gap-8">
            <Card className="border-green-200 bg-green-50/50">
              <CardHeader>
                <CardTitle className="flex items-center gap-2 text-green-700">
                  <CheckCircle className="h-5 w-5" />
                  基本要求
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-3">
                <div className="flex items-start gap-3">
                  <CheckCircle className="h-5 w-5 text-green-600 mt-0.5 flex-shrink-0" />
                  <p>年满18周岁，具有完全民事行为能力</p>
                </div>
                <div className="flex items-start gap-3">
                  <CheckCircle className="h-5 w-5 text-green-600 mt-0.5 flex-shrink-0" />
                  <p>有稳定的住所和经济来源</p>
                </div>
                <div className="flex items-start gap-3">
                  <CheckCircle className="h-5 w-5 text-green-600 mt-0.5 flex-shrink-0" />
                  <p>家人一致同意领养</p>
                </div>
                <div className="flex items-start gap-3">
                  <CheckCircle className="h-5 w-5 text-green-600 mt-0.5 flex-shrink-0" />
                  <p>承诺科学喂养、定期免疫</p>
                </div>
                <div className="flex items-start gap-3">
                  <CheckCircle className="h-5 w-5 text-green-600 mt-0.5 flex-shrink-0" />
                  <p>接受必要的回访</p>
                </div>
                <div className="flex items-start gap-3">
                  <CheckCircle className="h-5 w-5 text-green-600 mt-0.5 flex-shrink-0" />
                  <p>签署领养协议并遵守约定</p>
                </div>
              </CardContent>
            </Card>

            <Card className="border-red-200 bg-red-50/50">
              <CardHeader>
                <CardTitle className="flex items-center gap-2 text-red-700">
                  <AlertCircle className="h-5 w-5" />
                  不建议领养的情况
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-3">
                <div className="flex items-start gap-3">
                  <AlertCircle className="h-5 w-5 text-red-500 mt-0.5 flex-shrink-0" />
                  <p>学生党（经济和时间不稳定）</p>
                </div>
                <div className="flex items-start gap-3">
                  <AlertCircle className="h-5 w-5 text-red-500 mt-0.5 flex-shrink-0" />
                  <p>经常出差或工作繁忙无暇照顾</p>
                </div>
                <div className="flex items-start gap-3">
                  <AlertCircle className="h-5 w-5 text-red-500 mt-0.5 flex-shrink-0" />
                  <p>家人有宠物过敏或强烈反对</p>
                </div>
                <div className="flex items-start gap-3">
                  <AlertCircle className="h-5 w-5 text-red-500 mt-0.5 flex-shrink-0" />
                  <p>将宠物作为礼物送人</p>
                </div>
                <div className="flex items-start gap-3">
                  <AlertCircle className="h-5 w-5 text-red-500 mt-0.5 flex-shrink-0" />
                  <p>租房且房东不允许养宠</p>
                </div>
                <div className="flex items-start gap-3">
                  <AlertCircle className="h-5 w-5 text-red-500 mt-0.5 flex-shrink-0" />
                  <p>一时冲动、没有长期规划</p>
                </div>
              </CardContent>
            </Card>
          </div>
        </section>

        {/* 领养后须知 */}
        <section>
          <div className="flex items-center gap-3 mb-8">
            <div className="p-3 bg-primary/10 rounded-xl">
              <Calendar className="h-6 w-6 text-primary" />
            </div>
            <div>
              <h2 className="text-2xl font-bold">领养后须知</h2>
              <p className="text-muted-foreground">将宠物带回家后，请注意以下事项</p>
            </div>
          </div>

          <div className="grid md:grid-cols-3 gap-6">
            <Card>
              <CardHeader>
                <CardTitle className="text-lg">第一周：适应期</CardTitle>
              </CardHeader>
              <CardContent className="space-y-2 text-sm text-muted-foreground">
                <p>• 给宠物一个安静的空间适应新环境</p>
                <p>• 不要急于抱或过度亲近</p>
                <p>• 保持原有的饮食习惯</p>
                <p>• 准备好躲避处让它感到安全</p>
                <p>• 暂不进行洗澡等应激行为</p>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="text-lg">第一个月：建立信任</CardTitle>
              </CardHeader>
              <CardContent className="space-y-2 text-sm text-muted-foreground">
                <p>• 带宠物进行全面健康检查</p>
                <p>• 按时完成疫苗接种</p>
                <p>• 开始基础训练（定点排便等）</p>
                <p>• 建立固定的作息和喂食时间</p>
                <p>• 耐心陪伴，不要责骂</p>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="text-lg">长期照顾</CardTitle>
              </CardHeader>
              <CardContent className="space-y-2 text-sm text-muted-foreground">
                <p>• 定期驱虫（每3个月）</p>
                <p>• 每年接种疫苗</p>
                <p>• 定期体检（每年1-2次）</p>
                <p>• 适龄绝育手术</p>
                <p>• 配合平台进行回访</p>
              </CardContent>
            </Card>
          </div>
        </section>

        {/* 常见问题 */}
        <section>
          <div className="flex items-center gap-3 mb-8">
            <div className="p-3 bg-primary/10 rounded-xl">
              <AlertCircle className="h-6 w-6 text-primary" />
            </div>
            <div>
              <h2 className="text-2xl font-bold">常见问题</h2>
              <p className="text-muted-foreground">关于领养的一些常见疑问</p>
            </div>
          </div>

          <div className="grid md:grid-cols-2 gap-6">
            {[
              { q: '领养需要付费吗？', a: '领养本身不收费，但可能需要支付疫苗、绝育等医疗费用，具体以实际情况为准。这些费用用于保障宠物的健康。' },
              { q: '领养审核需要多长时间？', a: '一般审核时间为3-7个工作日。如果需要上门回访，时间可能会稍长。请耐心等待。' },
              { q: '领养后可以退回吗？', a: '我们不鼓励退养行为。但如果确实无法继续照顾，请第一时间联系我们，切勿遗弃或转送他人。' },
              { q: '为什么会被拒绝领养？', a: '常见原因包括：居住条件不符合、家人不同意、无法配合回访等。被拒绝并不代表您不适合养宠，可以改善条件后再次申请。' },
              { q: '可以指定领养某只宠物吗？', a: '可以。您可以在平台上浏览宠物信息，选择心仪的宠物后提交领养申请。' },
              { q: '领养后需要做什么检查？', a: '建议在领养后一周内带宠物进行全面体检，包括血常规、传染病筛查等，确保健康状况。' },
            ].map((item, index) => (
              <Card key={index}>
                <CardContent className="pt-6">
                  <h3 className="font-bold mb-2">{item.q}</h3>
                  <p className="text-sm text-muted-foreground">{item.a}</p>
                </CardContent>
              </Card>
            ))}
          </div>
        </section>

        {/* CTA */}
        <section className="text-center py-12 bg-gradient-to-r from-primary/10 to-secondary/10 rounded-2xl">
          <div className="max-w-2xl mx-auto px-4">
            <h2 className="text-2xl font-bold mb-4">准备好了吗？</h2>
            <p className="text-muted-foreground mb-6">
              如果您已经做好准备，欢迎浏览我们的宠物列表，找到您的命中注定！
            </p>
            <div className="flex justify-center gap-4">
              <Button size="lg" asChild>
                <Link to="/pets">
                  <PawPrint className="h-4 w-4 mr-2" />
                  浏览待领养宠物
                </Link>
              </Button>
              <Button size="lg" variant="outline" asChild>
                <Link to="/community">
                  加入社区交流
                </Link>
              </Button>
            </div>
          </div>
        </section>

        {/* 联系方式 */}
        <section className="text-center">
          <Card>
            <CardContent className="py-8">
              <div className="flex items-center justify-center gap-2 mb-4">
                <MapPin className="h-5 w-5 text-primary" />
                <span className="font-medium">联系我们</span>
              </div>
              <p className="text-muted-foreground mb-2">
                如有任何疑问，欢迎通过以下方式联系我们
              </p>
              <div className="flex justify-center gap-8 text-sm">
                <span>📧 contact@petadopt.com</span>
                <span>📞 400-123-4567</span>
                <span>🕐 工作日 9:00-18:00</span>
              </div>
            </CardContent>
          </Card>
        </section>
      </div>
    </div>
  )
}
