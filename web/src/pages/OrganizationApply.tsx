import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useMutation } from '@tanstack/react-query'
import { Building2, ArrowLeft } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Input } from '../components/ui/input'
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '../components/ui/card'
import { organizationApi } from '../lib/api'
import type { OrganizationCreateRequest } from '../lib/api'

const orgTypes = [
  { value: 'shelter', label: '动物收容所' },
  { value: 'rescue', label: '救助站' },
  { value: 'hospital', label: '宠物医院' },
  { value: 'association', label: '动物保护协会' },
  { value: 'other', label: '其他' },
]

export function OrganizationApply() {
  const navigate = useNavigate()
  const [formData, setFormData] = useState<OrganizationCreateRequest>({
    name: '',
    type: 'shelter',
    description: '',
    contact_name: '',
    contact_phone: '',
    contact_email: '',
    province: '',
    city: '',
    address: '',
  })

  const createMutation = useMutation({
    mutationFn: (data: OrganizationCreateRequest) => organizationApi.create(data),
    onSuccess: () => {
      alert('入驻申请提交成功！请等待审核')
      navigate('/my-organizations')
    },
    onError: (error: Error) => {
      alert(error.message)
    },
  })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target
    setFormData({ ...formData, [name]: value })
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    
    if (!formData.name.trim()) {
      alert('请输入机构名称')
      return
    }
    if (!formData.contact_name?.trim()) {
      alert('请输入联系人姓名')
      return
    }
    if (!formData.contact_phone?.trim()) {
      alert('请输入联系电话')
      return
    }

    createMutation.mutate(formData)
  }

  return (
    <div className="container py-8">
      <Button variant="ghost" onClick={() => navigate(-1)} className="mb-4">
        <ArrowLeft className="w-4 h-4 mr-2" />
        返回
      </Button>

      <div className="max-w-2xl mx-auto">
        <Card>
          <CardHeader className="text-center">
            <div className="w-16 h-16 rounded-full bg-primary/10 flex items-center justify-center mx-auto mb-4">
              <Building2 className="w-8 h-8 text-primary" />
            </div>
            <CardTitle className="text-2xl">机构入驻申请</CardTitle>
            <CardDescription>
              填写以下信息申请成为平台合作机构，审核通过后可发布宠物领养信息
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-6">
              <div className="space-y-2">
                <label className="text-sm font-medium">机构名称 *</label>
                <Input
                  name="name"
                  placeholder="请输入机构全称"
                  value={formData.name}
                  onChange={handleChange}
                  required
                />
              </div>

              <div className="space-y-2">
                <label className="text-sm font-medium">机构类型 *</label>
                <select
                  name="type"
                  className="w-full p-2 border rounded-md"
                  value={formData.type}
                  onChange={handleChange}
                >
                  {orgTypes.map((type) => (
                    <option key={type.value} value={type.value}>
                      {type.label}
                    </option>
                  ))}
                </select>
              </div>

              <div className="space-y-2">
                <label className="text-sm font-medium">机构简介</label>
                <textarea
                  name="description"
                  className="w-full min-h-[100px] p-3 border rounded-md"
                  placeholder="请简要介绍您的机构"
                  value={formData.description}
                  onChange={handleChange}
                />
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <label className="text-sm font-medium">联系人姓名 *</label>
                  <Input
                    name="contact_name"
                    placeholder="请输入联系人姓名"
                    value={formData.contact_name}
                    onChange={handleChange}
                    required
                  />
                </div>
                <div className="space-y-2">
                  <label className="text-sm font-medium">联系电话 *</label>
                  <Input
                    name="contact_phone"
                    placeholder="请输入联系电话"
                    value={formData.contact_phone}
                    onChange={handleChange}
                    required
                  />
                </div>
              </div>

              <div className="space-y-2">
                <label className="text-sm font-medium">联系邮箱</label>
                <Input
                  name="contact_email"
                  type="email"
                  placeholder="请输入联系邮箱"
                  value={formData.contact_email}
                  onChange={handleChange}
                />
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <label className="text-sm font-medium">省份</label>
                  <Input
                    name="province"
                    placeholder="如：广东省"
                    value={formData.province}
                    onChange={handleChange}
                  />
                </div>
                <div className="space-y-2">
                  <label className="text-sm font-medium">城市</label>
                  <Input
                    name="city"
                    placeholder="如：广州市"
                    value={formData.city}
                    onChange={handleChange}
                  />
                </div>
              </div>

              <div className="space-y-2">
                <label className="text-sm font-medium">详细地址</label>
                <Input
                  name="address"
                  placeholder="请输入详细地址"
                  value={formData.address}
                  onChange={handleChange}
                />
              </div>

              <Button
                type="submit"
                className="w-full"
                size="lg"
                disabled={createMutation.isPending}
              >
                {createMutation.isPending ? '提交中...' : '提交申请'}
              </Button>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
