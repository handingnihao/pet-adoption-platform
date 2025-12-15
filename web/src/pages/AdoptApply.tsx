import { useState } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { useQuery, useMutation } from '@tanstack/react-query'
import { ArrowLeft } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '../components/ui/card'
import { petApi, adoptionApi } from '../lib/api'
import { useAuthStore } from '../store/auth'

export function AdoptApply() {
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()
  const { isAuthenticated } = useAuthStore()

  const [formData, setFormData] = useState({
    reason: '',
    experience: '',
    living_condition: '',
    family_agreement: false,
  })

  const { data: petData, isLoading: petLoading } = useQuery({
    queryKey: ['pet', id],
    queryFn: () => petApi.getById(Number(id)),
    enabled: !!id,
  })

  const mutation = useMutation({
    mutationFn: () => adoptionApi.createApplication({
      pet_id: Number(id),
      ...formData,
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

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value, type } = e.target
    setFormData({
      ...formData,
      [name]: type === 'checkbox' ? (e.target as HTMLInputElement).checked : value,
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

              <div className="space-y-2">
                <label className="text-sm font-medium">
                  领养原因 <span className="text-destructive">*</span>
                </label>
                <textarea
                  name="reason"
                  className="w-full min-h-[100px] px-3 py-2 border rounded-md text-sm resize-none focus:outline-none focus:ring-2 focus:ring-primary"
                  placeholder="请说明您想要领养这只宠物的原因..."
                  value={formData.reason}
                  onChange={handleChange}
                  required
                />
              </div>

              <div className="space-y-2">
                <label className="text-sm font-medium">养宠经验</label>
                <textarea
                  name="experience"
                  className="w-full min-h-[80px] px-3 py-2 border rounded-md text-sm resize-none focus:outline-none focus:ring-2 focus:ring-primary"
                  placeholder="请描述您之前的养宠经验（如果有）..."
                  value={formData.experience}
                  onChange={handleChange}
                />
              </div>

              <div className="space-y-2">
                <label className="text-sm font-medium">居住条件</label>
                <textarea
                  name="living_condition"
                  className="w-full min-h-[80px] px-3 py-2 border rounded-md text-sm resize-none focus:outline-none focus:ring-2 focus:ring-primary"
                  placeholder="请描述您的居住环境（如：住房类型、面积、是否允许养宠等）..."
                  value={formData.living_condition}
                  onChange={handleChange}
                />
              </div>

              <div className="flex items-center gap-2">
                <input
                  type="checkbox"
                  id="family_agreement"
                  name="family_agreement"
                  checked={formData.family_agreement}
                  onChange={handleChange}
                  className="rounded"
                />
                <label htmlFor="family_agreement" className="text-sm">
                  家人已同意领养
                </label>
              </div>

              <div className="flex gap-4">
                <Button
                  type="submit"
                  className="flex-1"
                  disabled={mutation.isPending || !formData.reason}
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
