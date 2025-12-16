import { useState } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { useQuery, useMutation } from '@tanstack/react-query'
import { ArrowLeft } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Input } from '../components/ui/input'
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '../components/ui/card'
import { petApi, adoptionApi } from '../lib/api'
import { useAuthStore } from '../store/auth'

export function AdoptApply() {
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()
  const { isAuthenticated } = useAuthStore()

  const [formData, setFormData] = useState({
    applicant_name: '',
    applicant_phone: '',
    applicant_address: '',
    housing_type: 'apartment' as 'apartment' | 'house' | 'villa' | 'other',
    housing_area: 50,
    has_yard: false,
    family_members: 1,
    has_children: false,
    children_age: '',
    family_agree: false,
    has_pet_experience: false,
    pet_experience: '',
    current_pets: '',
    adoption_reason: '',
    how_to_care: '',
    emergency_plan: '',
  })

  const { data: petData, isLoading: petLoading } = useQuery({
    queryKey: ['pet', id],
    queryFn: () => petApi.getById(Number(id)),
    enabled: !!id,
  })

  const mutation = useMutation({
    mutationFn: () => adoptionApi.createApplication({
      pet_id: Number(id),
      organization_id: 1,
      applicant_name: formData.applicant_name,
      applicant_phone: formData.applicant_phone,
      applicant_address: formData.applicant_address,
      housing_type: formData.housing_type,
      housing_area: formData.housing_area,
      has_yard: formData.has_yard,
      family_members: formData.family_members,
      has_children: formData.has_children,
      children_age: formData.children_age,
      family_agree: formData.family_agree,
      adoption_reason: formData.adoption_reason,
      how_to_care: formData.how_to_care,
      emergency_plan: formData.emergency_plan,
    }),
    onSuccess: () => {
      navigate('/my-applications', { state: { message: '申请已提交，请等待审核' } })
    },
  })

  const pet = petData?.data

  if (!isAuthenticated) {
    navigate('/login')
    return null
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value, type } = e.target
    let newValue: string | number | boolean = value
    
    if (type === 'checkbox') {
      newValue = (e.target as HTMLInputElement).checked
    } else if (type === 'number') {
      newValue = parseInt(value, 10) || 0
    }
    
    setFormData({
      ...formData,
      [name]: newValue,
    })
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    mutation.mutate()
  }

  if (petLoading) {
    return (
      <div className="container py-8">
        <div className="max-w-2xl mx-auto animate-pulse">
          <div className="h-8 w-32 bg-muted rounded mb-6" />
          <div className="h-64 bg-muted rounded" />
        </div>
      </div>
    )
  }

  if (!pet) {
    return (
      <div className="container py-8 text-center">
        <p>宠物不存在</p>
        <Button onClick={() => navigate('/pets')} className="mt-4">返回列表</Button>
      </div>
    )
  }

  return (
    <div className="container py-8">
      <div className="max-w-2xl mx-auto">
        <Button variant="ghost" className="mb-6" onClick={() => navigate(-1)}>
          <ArrowLeft className="h-4 w-4 mr-2" />
          返回
        </Button>

        <Card>
          <CardHeader>
            <CardTitle>申请领养 - {pet.name}</CardTitle>
            <CardDescription>
              请认真填写以下信息，我们会尽快审核您的申请
            </CardDescription>
          </CardHeader>
          <CardContent>
            {/* 宠物信息卡片 */}
            <div className="flex items-center gap-4 p-4 bg-muted rounded-lg mb-6">
              <div className="w-16 h-16 rounded-lg overflow-hidden bg-background">
                {pet.cover_photo ? (
                  <img src={pet.cover_photo} alt={pet.name} className="w-full h-full object-cover" />
                ) : (
                  <div className="w-full h-full flex items-center justify-center text-2xl">
                    {pet.type === 'cat' ? '🐱' : '🐕'}
                  </div>
                )}
              </div>
              <div>
                <h3 className="font-medium">{pet.name}</h3>
                <p className="text-sm text-muted-foreground">
                  {pet.breed || pet.type} · {pet.age}个月
                </p>
              </div>
            </div>

            <form onSubmit={handleSubmit} className="space-y-6">
              {mutation.isError && (
                <div className="p-3 text-sm text-destructive bg-destructive/10 rounded-md">
                  {mutation.error instanceof Error ? mutation.error.message : '提交失败'}
                </div>
              )}

              {/* 申请人信息 */}
              <div className="space-y-4">
                <h3 className="font-medium text-lg border-b pb-2">申请人信息</h3>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">
                      姓名 <span className="text-destructive">*</span>
                    </label>
                    <Input
                      name="applicant_name"
                      placeholder="请输入您的姓名"
                      value={formData.applicant_name}
                      onChange={handleChange}
                      required
                    />
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">
                      联系电话 <span className="text-destructive">*</span>
                    </label>
                    <Input
                      name="applicant_phone"
                      placeholder="请输入11位手机号"
                      value={formData.applicant_phone}
                      onChange={handleChange}
                      required
                      maxLength={11}
                    />
                  </div>
                </div>
                <div className="space-y-2">
                  <label className="text-sm font-medium">
                    详细地址 <span className="text-destructive">*</span>
                  </label>
                  <Input
                    name="applicant_address"
                    placeholder="请输入您的详细居住地址"
                    value={formData.applicant_address}
                    onChange={handleChange}
                    required
                  />
                </div>
              </div>

              {/* 住房情况 */}
              <div className="space-y-4">
                <h3 className="font-medium text-lg border-b pb-2">住房情况</h3>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">
                      住房类型 <span className="text-destructive">*</span>
                    </label>
                    <select
                      name="housing_type"
                      className="w-full px-3 py-2 border rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-primary"
                      value={formData.housing_type}
                      onChange={handleChange}
                      required
                    >
                      <option value="apartment">公寓</option>
                      <option value="house">住宅</option>
                      <option value="villa">别墅</option>
                      <option value="other">其他</option>
                    </select>
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">
                      住房面积(㎡) <span className="text-destructive">*</span>
                    </label>
                    <Input
                      type="number"
                      name="housing_area"
                      placeholder="请输入住房面积"
                      value={formData.housing_area}
                      onChange={handleChange}
                      required
                      min={10}
                      max={10000}
                    />
                  </div>
                </div>
                <div className="flex items-center gap-4">
                  <label className="flex items-center gap-2">
                    <input
                      type="checkbox"
                      name="has_yard"
                      checked={formData.has_yard}
                      onChange={handleChange}
                      className="rounded"
                    />
                    <span className="text-sm">有院子/阳台</span>
                  </label>
                </div>
              </div>

              {/* 家庭情况 */}
              <div className="space-y-4">
                <h3 className="font-medium text-lg border-b pb-2">家庭情况</h3>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">
                      家庭成员数 <span className="text-destructive">*</span>
                    </label>
                    <Input
                      type="number"
                      name="family_members"
                      placeholder="家庭成员人数"
                      value={formData.family_members}
                      onChange={handleChange}
                      required
                      min={1}
                      max={20}
                    />
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">孩子年龄</label>
                    <Input
                      name="children_age"
                      placeholder="如有孩子，请填写年龄"
                      value={formData.children_age}
                      onChange={handleChange}
                    />
                  </div>
                </div>
                <div className="flex items-center gap-4 flex-wrap">
                  <label className="flex items-center gap-2">
                    <input
                      type="checkbox"
                      name="has_children"
                      checked={formData.has_children}
                      onChange={handleChange}
                      className="rounded"
                    />
                    <span className="text-sm">家中有孩子</span>
                  </label>
                  <label className="flex items-center gap-2">
                    <input
                      type="checkbox"
                      name="family_agree"
                      checked={formData.family_agree}
                      onChange={handleChange}
                      className="rounded"
                    />
                    <span className="text-sm">家人已同意领养 <span className="text-destructive">*</span></span>
                  </label>
                </div>
              </div>

              {/* 领养计划 */}
              <div className="space-y-4">
                <h3 className="font-medium text-lg border-b pb-2">领养计划</h3>
                <div className="space-y-2">
                  <label className="text-sm font-medium">
                    领养原因 <span className="text-destructive">*</span>
                  </label>
                  <textarea
                    name="adoption_reason"
                    className="w-full min-h-[100px] px-3 py-2 border rounded-md text-sm resize-none focus:outline-none focus:ring-2 focus:ring-primary"
                    placeholder="请说明您想要领养这只宠物的原因（至少10个字）..."
                    value={formData.adoption_reason}
                    onChange={handleChange}
                    required
                    minLength={10}
                  />
                </div>
                <div className="space-y-2">
                  <label className="text-sm font-medium">
                    照顾计划 <span className="text-destructive">*</span>
                  </label>
                  <textarea
                    name="how_to_care"
                    className="w-full min-h-[80px] px-3 py-2 border rounded-md text-sm resize-none focus:outline-none focus:ring-2 focus:ring-primary"
                    placeholder="请描述您计划如何照顾宠物（饮食、运动、陪伴等，至少10个字）..."
                    value={formData.how_to_care}
                    onChange={handleChange}
                    required
                    minLength={10}
                  />
                </div>
                <div className="space-y-2">
                  <label className="text-sm font-medium">
                    应急计划 <span className="text-destructive">*</span>
                  </label>
                  <textarea
                    name="emergency_plan"
                    className="w-full min-h-[80px] px-3 py-2 border rounded-md text-sm resize-none focus:outline-none focus:ring-2 focus:ring-primary"
                    placeholder="如遇紧急情况（如搬家、出差、生病等），您将如何安置宠物？（至少10个字）..."
                    value={formData.emergency_plan}
                    onChange={handleChange}
                    required
                    minLength={10}
                  />
                </div>
              </div>

              <div className="flex gap-4 pt-4">
                <Button
                  type="submit"
                  className="flex-1"
                  disabled={mutation.isPending || !formData.applicant_name || !formData.adoption_reason || !formData.family_agree}
                >
                  {mutation.isPending ? '提交中...' : '提交申请'}
                </Button>
                <Button type="button" variant="outline" onClick={() => navigate(-1)}>
                  取消
                </Button>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
