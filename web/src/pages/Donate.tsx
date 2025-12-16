import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useMutation, useQuery } from '@tanstack/react-query'
import { Heart, Gift, HandHelping, ArrowLeft, Check } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Input } from '../components/ui/input'
import { Card, CardContent, CardHeader, CardTitle } from '../components/ui/card'
import { donationApi, organizationApi } from '../lib/api'
import type { DonationCreateRequest, DonationType } from '../lib/api'
import { useAuthStore } from '../store/auth'

const donationTypes = [
  { value: 'money' as DonationType, label: '资金捐赠', icon: Heart, description: '直接捐赠资金，帮助购买宠物食品、医疗用品等' },
  { value: 'supply' as DonationType, label: '物资捐赠', icon: Gift, description: '捐赠猫粮、狗粮、玩具、药品等物资' },
  { value: 'service' as DonationType, label: '服务捐赠', icon: HandHelping, description: '提供志愿服务，如遛狗、清洁、医疗服务等' },
]

const amountOptions = [10, 50, 100, 200, 500, 1000]

export function Donate() {
  const navigate = useNavigate()
  const { isAuthenticated } = useAuthStore()
  const [formData, setFormData] = useState<DonationCreateRequest>({
    type: 'money',
    amount: 0,
    supply_items: [],
    service_desc: '',
    message: '',
    is_anonymous: false,
  })
  const [customAmount, setCustomAmount] = useState('')
  const [supplyInput, setSupplyInput] = useState('')
  const [step, setStep] = useState(1) // 1: 选择类型, 2: 填写详情, 3: 确认

  // 获取机构列表
  const { data: orgsData } = useQuery({
    queryKey: ['organizations-for-donate'],
    queryFn: () => organizationApi.list({ page: 1, page_size: 100 }),
  })

  // 获取捐赠统计
  const { data: statsData } = useQuery({
    queryKey: ['donation-statistics'],
    queryFn: () => donationApi.getStatistics(),
  })

  const createMutation = useMutation({
    mutationFn: (data: DonationCreateRequest) => donationApi.create(data),
    onSuccess: () => {
      alert('捐赠创建成功！感谢您的爱心！')
      navigate('/my-donations')
    },
    onError: (error: Error) => {
      alert(error.message)
    },
  })

  const handleTypeSelect = (type: DonationType) => {
    setFormData({ ...formData, type })
    setStep(2)
  }

  const handleAmountSelect = (amount: number) => {
    setFormData({ ...formData, amount })
    setCustomAmount('')
  }

  const handleCustomAmountChange = (value: string) => {
    setCustomAmount(value)
    const amount = parseFloat(value) || 0
    setFormData({ ...formData, amount })
  }

  const addSupplyItem = () => {
    if (supplyInput.trim()) {
      setFormData({
        ...formData,
        supply_items: [...(formData.supply_items || []), supplyInput.trim()],
      })
      setSupplyInput('')
    }
  }

  const removeSupplyItem = (index: number) => {
    const items = [...(formData.supply_items || [])]
    items.splice(index, 1)
    setFormData({ ...formData, supply_items: items })
  }

  const handleSubmit = () => {
    if (!isAuthenticated) {
      alert('请先登录')
      navigate('/login')
      return
    }

    // 验证
    if (formData.type === 'money' && (!formData.amount || formData.amount <= 0)) {
      alert('请输入捐赠金额')
      return
    }
    if (formData.type === 'supply' && (!formData.supply_items || formData.supply_items.length === 0)) {
      alert('请添加至少一项物资')
      return
    }
    if (formData.type === 'service' && !formData.service_desc) {
      alert('请填写服务描述')
      return
    }

    createMutation.mutate(formData)
  }

  const stats = statsData?.data

  return (
    <div className="min-h-screen bg-gradient-to-b from-primary/5 to-background">
      {/* Header */}
      <section className="py-12 bg-gradient-to-r from-primary/10 to-primary/5">
        <div className="container">
          <div className="max-w-3xl mx-auto text-center">
            <Heart className="w-16 h-16 text-primary mx-auto mb-4" />
            <h1 className="text-4xl font-bold mb-4">爱心捐赠</h1>
            <p className="text-lg text-muted-foreground">
              您的每一份爱心，都将帮助更多流浪动物获得温暖与关爱
            </p>
            
            {/* 统计数据 */}
            {stats && (
              <div className="flex justify-center gap-8 mt-8">
                <div className="text-center">
                  <div className="text-3xl font-bold text-primary">{stats.total_donors}</div>
                  <div className="text-sm text-muted-foreground">爱心人士</div>
                </div>
                <div className="text-center">
                  <div className="text-3xl font-bold text-primary">¥{stats.total_amount.toFixed(0)}</div>
                  <div className="text-sm text-muted-foreground">累计捐赠</div>
                </div>
                <div className="text-center">
                  <div className="text-3xl font-bold text-primary">{stats.total_donations}</div>
                  <div className="text-sm text-muted-foreground">捐赠次数</div>
                </div>
              </div>
            )}
          </div>
        </div>
      </section>

      <div className="container py-12">
        <div className="max-w-2xl mx-auto">
          {/* 步骤指示器 */}
          <div className="flex items-center justify-center gap-4 mb-8">
            {[1, 2, 3].map((s) => (
              <div key={s} className="flex items-center">
                <div className={`w-8 h-8 rounded-full flex items-center justify-center text-sm font-medium ${
                  step >= s ? 'bg-primary text-primary-foreground' : 'bg-muted text-muted-foreground'
                }`}>
                  {step > s ? <Check className="w-4 h-4" /> : s}
                </div>
                {s < 3 && <div className={`w-12 h-0.5 ${step > s ? 'bg-primary' : 'bg-muted'}`} />}
              </div>
            ))}
          </div>

          {/* Step 1: 选择捐赠类型 */}
          {step === 1 && (
            <div className="space-y-4">
              <h2 className="text-2xl font-bold text-center mb-6">选择捐赠方式</h2>
              <div className="grid gap-4">
                {donationTypes.map((type) => (
                  <Card
                    key={type.value}
                    className="cursor-pointer hover:border-primary transition-colors"
                    onClick={() => handleTypeSelect(type.value)}
                  >
                    <CardContent className="p-6 flex items-center gap-4">
                      <div className="w-12 h-12 rounded-full bg-primary/10 flex items-center justify-center">
                        <type.icon className="w-6 h-6 text-primary" />
                      </div>
                      <div className="flex-1">
                        <h3 className="font-bold text-lg">{type.label}</h3>
                        <p className="text-sm text-muted-foreground">{type.description}</p>
                      </div>
                    </CardContent>
                  </Card>
                ))}
              </div>
            </div>
          )}

          {/* Step 2: 填写详情 */}
          {step === 2 && (
            <Card>
              <CardHeader>
                <Button variant="ghost" size="sm" onClick={() => setStep(1)} className="w-fit">
                  <ArrowLeft className="w-4 h-4 mr-2" />
                  返回
                </Button>
                <CardTitle>
                  {formData.type === 'money' && '资金捐赠'}
                  {formData.type === 'supply' && '物资捐赠'}
                  {formData.type === 'service' && '服务捐赠'}
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-6">
                {/* 资金捐赠 */}
                {formData.type === 'money' && (
                  <div className="space-y-4">
                    <label className="text-sm font-medium">选择金额</label>
                    <div className="grid grid-cols-3 gap-3">
                      {amountOptions.map((amount) => (
                        <Button
                          key={amount}
                          variant={formData.amount === amount && !customAmount ? 'default' : 'outline'}
                          onClick={() => handleAmountSelect(amount)}
                        >
                          ¥{amount}
                        </Button>
                      ))}
                    </div>
                    <div className="space-y-2">
                      <label className="text-sm font-medium">自定义金额</label>
                      <Input
                        type="number"
                        placeholder="输入自定义金额"
                        value={customAmount}
                        onChange={(e) => handleCustomAmountChange(e.target.value)}
                      />
                    </div>
                  </div>
                )}

                {/* 物资捐赠 */}
                {formData.type === 'supply' && (
                  <div className="space-y-4">
                    <label className="text-sm font-medium">物资清单</label>
                    <div className="flex gap-2">
                      <Input
                        placeholder="输入物资名称，如：猫粮10kg"
                        value={supplyInput}
                        onChange={(e) => setSupplyInput(e.target.value)}
                        onKeyPress={(e) => e.key === 'Enter' && addSupplyItem()}
                      />
                      <Button onClick={addSupplyItem}>添加</Button>
                    </div>
                    {formData.supply_items && formData.supply_items.length > 0 && (
                      <div className="flex flex-wrap gap-2">
                        {formData.supply_items.map((item, index) => (
                          <span
                            key={index}
                            className="inline-flex items-center gap-1 px-3 py-1 bg-primary/10 rounded-full text-sm"
                          >
                            {item}
                            <button
                              onClick={() => removeSupplyItem(index)}
                              className="text-muted-foreground hover:text-destructive"
                            >
                              ×
                            </button>
                          </span>
                        ))}
                      </div>
                    )}
                  </div>
                )}

                {/* 服务捐赠 */}
                {formData.type === 'service' && (
                  <div className="space-y-2">
                    <label className="text-sm font-medium">服务描述</label>
                    <textarea
                      className="w-full min-h-[100px] p-3 border rounded-md"
                      placeholder="请描述您可以提供的服务，如：每周末可以帮忙遛狗2小时"
                      value={formData.service_desc}
                      onChange={(e) => setFormData({ ...formData, service_desc: e.target.value })}
                    />
                  </div>
                )}

                {/* 选择机构 */}
                <div className="space-y-2">
                  <label className="text-sm font-medium">捐赠给（可选）</label>
                  <select
                    className="w-full p-2 border rounded-md"
                    value={formData.organization_id || ''}
                    onChange={(e) => setFormData({
                      ...formData,
                      organization_id: e.target.value ? parseInt(e.target.value) : undefined
                    })}
                  >
                    <option value="">平台统一分配</option>
                    {orgsData?.data?.list?.map((org) => (
                      <option key={org.id} value={org.id}>{org.name}</option>
                    ))}
                  </select>
                </div>

                {/* 留言 */}
                <div className="space-y-2">
                  <label className="text-sm font-medium">留言（可选）</label>
                  <textarea
                    className="w-full min-h-[80px] p-3 border rounded-md"
                    placeholder="写下您想说的话..."
                    value={formData.message}
                    onChange={(e) => setFormData({ ...formData, message: e.target.value })}
                  />
                </div>

                {/* 匿名选项 */}
                <label className="flex items-center gap-2 cursor-pointer">
                  <input
                    type="checkbox"
                    checked={formData.is_anonymous}
                    onChange={(e) => setFormData({ ...formData, is_anonymous: e.target.checked })}
                    className="w-4 h-4"
                  />
                  <span className="text-sm">匿名捐赠</span>
                </label>

                <Button className="w-full" size="lg" onClick={() => setStep(3)}>
                  下一步
                </Button>
              </CardContent>
            </Card>
          )}

          {/* Step 3: 确认 */}
          {step === 3 && (
            <Card>
              <CardHeader>
                <Button variant="ghost" size="sm" onClick={() => setStep(2)} className="w-fit">
                  <ArrowLeft className="w-4 h-4 mr-2" />
                  返回
                </Button>
                <CardTitle>确认捐赠信息</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="bg-muted/50 rounded-lg p-4 space-y-3">
                  <div className="flex justify-between">
                    <span className="text-muted-foreground">捐赠类型</span>
                    <span className="font-medium">
                      {formData.type === 'money' && '资金捐赠'}
                      {formData.type === 'supply' && '物资捐赠'}
                      {formData.type === 'service' && '服务捐赠'}
                    </span>
                  </div>
                  
                  {formData.type === 'money' && (
                    <div className="flex justify-between">
                      <span className="text-muted-foreground">捐赠金额</span>
                      <span className="font-bold text-primary text-xl">¥{formData.amount}</span>
                    </div>
                  )}
                  
                  {formData.type === 'supply' && formData.supply_items && (
                    <div>
                      <span className="text-muted-foreground">物资清单</span>
                      <div className="mt-2 flex flex-wrap gap-2">
                        {formData.supply_items.map((item, i) => (
                          <span key={i} className="px-2 py-1 bg-primary/10 rounded text-sm">{item}</span>
                        ))}
                      </div>
                    </div>
                  )}
                  
                  {formData.type === 'service' && (
                    <div>
                      <span className="text-muted-foreground">服务描述</span>
                      <p className="mt-1 text-sm">{formData.service_desc}</p>
                    </div>
                  )}
                  
                  {formData.message && (
                    <div>
                      <span className="text-muted-foreground">留言</span>
                      <p className="mt-1 text-sm">{formData.message}</p>
                    </div>
                  )}
                  
                  <div className="flex justify-between">
                    <span className="text-muted-foreground">是否匿名</span>
                    <span>{formData.is_anonymous ? '是' : '否'}</span>
                  </div>
                </div>

                <Button
                  className="w-full"
                  size="lg"
                  onClick={handleSubmit}
                  disabled={createMutation.isPending}
                >
                  {createMutation.isPending ? '提交中...' : '确认捐赠'}
                </Button>
                
                <p className="text-xs text-center text-muted-foreground">
                  提交后，平台将与您联系确认捐赠详情
                </p>
              </CardContent>
            </Card>
          )}
        </div>
      </div>
    </div>
  )
}
