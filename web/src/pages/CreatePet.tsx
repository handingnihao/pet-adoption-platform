import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useMutation } from '@tanstack/react-query'
import { ArrowLeft, Image, X } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Input } from '../components/ui/input'
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '../components/ui/card'
import { petApi } from '../lib/api'
import { useAuthStore } from '../store/auth'

const petTypes = [
  { value: 'cat', label: '猫咪' },
  { value: 'dog', label: '狗狗' },
  { value: 'rabbit', label: '兔子' },
  { value: 'hamster', label: '仓鼠' },
  { value: 'bird', label: '鸟类' },
  { value: 'other', label: '其他' },
]

const sizes = [
  { value: 'small', label: '小型' },
  { value: 'medium', label: '中型' },
  { value: 'large', label: '大型' },
]

export function CreatePet() {
  const navigate = useNavigate()
  const { isAuthenticated } = useAuthStore()

  const [formData, setFormData] = useState({
    name: '',
    type: 'cat',
    breed: '',
    gender: 'male',
    age: '',
    size: 'medium',
    color: '',
    weight: '',
    is_vaccinated: false,
    is_sterilized: false,
    health_status: '',
    description: '',
    character: '',
    cover_photo: '',
    photos: [] as string[],
    province: '',
    city: '',
    district: '',
    address: '',
  })

  const [photoUrl, setPhotoUrl] = useState('')

  const mutation = useMutation({
    mutationFn: () => petApi.create({
      ...formData,
      age: parseInt(formData.age) || 0,
      weight: parseFloat(formData.weight) || undefined,
      photos: formData.photos.length > 0 ? formData.photos : undefined,
    }),
    onSuccess: () => {
      navigate('/my-pets', { state: { message: '宠物发布成功，等待审核' } })
    },
  })

  if (!isAuthenticated) {
    navigate('/login')
    return null
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value, type } = e.target
    if (type === 'checkbox') {
      setFormData({ ...formData, [name]: (e.target as HTMLInputElement).checked })
    } else {
      setFormData({ ...formData, [name]: value })
    }
  }

  const handleAddPhoto = () => {
    if (photoUrl.trim() && formData.photos.length < 9) {
      setFormData({ ...formData, photos: [...formData.photos, photoUrl.trim()] })
      setPhotoUrl('')
    }
  }

  const handleRemovePhoto = (index: number) => {
    setFormData({ ...formData, photos: formData.photos.filter((_, i) => i !== index) })
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (!formData.name || !formData.type || !formData.age || !formData.description) {
      alert('请填写必填项（名称、类型、年龄、详细介绍）')
      return
    }
    mutation.mutate()
  }

  return (
    <div className="container py-8">
      <div className="max-w-3xl mx-auto">
        <Button variant="ghost" className="mb-6" onClick={() => navigate(-1)}>
          <ArrowLeft className="h-4 w-4 mr-2" />
          返回
        </Button>

        <Card>
          <CardHeader>
            <CardTitle>发布宠物</CardTitle>
            <CardDescription>填写宠物信息，帮助它找到新家</CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-6">
              {mutation.isError && (
                <div className="p-3 text-sm text-destructive bg-destructive/10 rounded-md">
                  {mutation.error instanceof Error ? mutation.error.message : '发布失败'}
                </div>
              )}

              {/* 基本信息 */}
              <div className="space-y-4">
                <h3 className="font-medium">基本信息</h3>
                <div className="grid md:grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">宠物名称 *</label>
                    <Input name="name" value={formData.name} onChange={handleChange} placeholder="给它起个名字" required />
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">宠物类型 *</label>
                    <select name="type" value={formData.type} onChange={handleChange} className="w-full h-10 px-3 border rounded-md">
                      {petTypes.map(t => <option key={t.value} value={t.value}>{t.label}</option>)}
                    </select>
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">品种</label>
                    <Input name="breed" value={formData.breed} onChange={handleChange} placeholder="如：金毛、英短等" />
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">性别 *</label>
                    <select name="gender" value={formData.gender} onChange={handleChange} className="w-full h-10 px-3 border rounded-md">
                      <option value="male">公</option>
                      <option value="female">母</option>
                    </select>
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">年龄（月） *</label>
                    <Input name="age" type="number" value={formData.age} onChange={handleChange} placeholder="月龄" required />
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">体型</label>
                    <select name="size" value={formData.size} onChange={handleChange} className="w-full h-10 px-3 border rounded-md">
                      {sizes.map(s => <option key={s.value} value={s.value}>{s.label}</option>)}
                    </select>
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">毛色</label>
                    <Input name="color" value={formData.color} onChange={handleChange} placeholder="毛色描述" />
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">体重（kg）</label>
                    <Input name="weight" type="number" step="0.1" value={formData.weight} onChange={handleChange} placeholder="体重" />
                  </div>
                </div>
              </div>

              {/* 健康信息 */}
              <div className="space-y-4">
                <h3 className="font-medium">健康信息</h3>
                <div className="flex gap-6">
                  <label className="flex items-center gap-2">
                    <input type="checkbox" name="is_vaccinated" checked={formData.is_vaccinated} onChange={handleChange} className="rounded" />
                    <span className="text-sm">已接种疫苗</span>
                  </label>
                  <label className="flex items-center gap-2">
                    <input type="checkbox" name="is_sterilized" checked={formData.is_sterilized} onChange={handleChange} className="rounded" />
                    <span className="text-sm">已绝育</span>
                  </label>
                </div>
                <div className="space-y-2">
                  <label className="text-sm font-medium">健康状况</label>
                  <Input name="health_status" value={formData.health_status} onChange={handleChange} placeholder="如：健康、有轻微皮肤病等" />
                </div>
              </div>

              {/* 性格描述 */}
              <div className="space-y-4">
                <h3 className="font-medium">详细描述</h3>
                <div className="space-y-2">
                  <label className="text-sm font-medium">性格特点</label>
                  <Input name="character" value={formData.character} onChange={handleChange} placeholder="如：温顺、活泼、粘人等" />
                </div>
                <div className="space-y-2">
                  <label className="text-sm font-medium">详细介绍 *</label>
                  <textarea
                    name="description"
                    value={formData.description}
                    onChange={handleChange}
                    className="w-full min-h-[120px] px-3 py-2 border rounded-md text-sm resize-none focus:outline-none focus:ring-2 focus:ring-primary"
                    placeholder="详细描述宠物的情况、故事等..."
                  />
                </div>
              </div>

              {/* 图片 */}
              <div className="space-y-4">
                <h3 className="font-medium">宠物照片</h3>
                <div className="space-y-2">
                  <label className="text-sm font-medium">封面图片</label>
                  <Input name="cover_photo" value={formData.cover_photo} onChange={handleChange} placeholder="输入图片URL" />
                </div>
                <div className="space-y-2">
                  <label className="text-sm font-medium">更多图片（最多9张）</label>
                  {formData.photos.length > 0 && (
                    <div className="flex gap-2 flex-wrap mb-2">
                      {formData.photos.map((photo, idx) => (
                        <div key={idx} className="relative w-20 h-20">
                          <img src={photo} alt="" className="w-full h-full object-cover rounded-lg" />
                          <button type="button" className="absolute -top-2 -right-2 p-1 bg-destructive text-white rounded-full" onClick={() => handleRemovePhoto(idx)}>
                            <X className="w-3 h-3" />
                          </button>
                        </div>
                      ))}
                    </div>
                  )}
                  {formData.photos.length < 9 && (
                    <div className="flex gap-2">
                      <Input value={photoUrl} onChange={(e) => setPhotoUrl(e.target.value)} placeholder="输入图片URL" />
                      <Button type="button" variant="outline" onClick={handleAddPhoto}>
                        <Image className="w-4 h-4 mr-1" />添加
                      </Button>
                    </div>
                  )}
                </div>
              </div>

              {/* 位置 */}
              <div className="space-y-4">
                <h3 className="font-medium">位置信息</h3>
                <div className="grid md:grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">省份</label>
                    <Input name="province" value={formData.province} onChange={handleChange} placeholder="如：广东省" />
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">城市</label>
                    <Input name="city" value={formData.city} onChange={handleChange} placeholder="如：深圳市" />
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">区县</label>
                    <Input name="district" value={formData.district} onChange={handleChange} placeholder="如：南山区" />
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">详细地址</label>
                    <Input name="address" value={formData.address} onChange={handleChange} placeholder="详细地址" />
                  </div>
                </div>
              </div>

              {/* 提交 */}
              <div className="flex gap-4 pt-4">
                <Button type="submit" className="flex-1" disabled={mutation.isPending}>
                  {mutation.isPending ? '发布中...' : '发布宠物'}
                </Button>
                <Button type="button" variant="outline" onClick={() => navigate(-1)}>取消</Button>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
