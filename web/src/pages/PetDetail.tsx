import { useState } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { ArrowLeft, MapPin, CheckCircle, XCircle } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '../components/ui/card'
import { petApi } from '../lib/api'
import { useAuthStore } from '../store/auth'

export function PetDetail() {
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()
  const { isAuthenticated } = useAuthStore()
  const [selectedImage, setSelectedImage] = useState(0)

  const { data, isLoading, error } = useQuery({
    queryKey: ['pet', id],
    queryFn: () => petApi.getById(Number(id)),
    enabled: !!id,
  })

  const pet = data?.data

  if (isLoading) {
    return (
      <div className="container py-8">
        <div className="animate-pulse">
          <div className="h-8 w-32 bg-muted rounded mb-6" />
          <div className="grid lg:grid-cols-2 gap-8">
            <div className="aspect-square bg-muted rounded-lg" />
            <div className="space-y-4">
              <div className="h-8 bg-muted rounded w-1/2" />
              <div className="h-4 bg-muted rounded w-1/3" />
              <div className="h-32 bg-muted rounded" />
            </div>
          </div>
        </div>
      </div>
    )
  }

  if (error || !pet) {
    return (
      <div className="container py-8 text-center">
        <div className="text-6xl mb-4">😿</div>
        <h2 className="text-xl font-semibold mb-2">宠物不存在</h2>
        <p className="text-muted-foreground mb-4">该宠物可能已被领养或下架</p>
        <Button onClick={() => navigate('/pets')}>返回列表</Button>
      </div>
    )
  }

  const images = pet.photos?.length ? pet.photos : (pet.cover_photo ? [pet.cover_photo] : [])
  const genderText = pet.gender === 'male' ? '公' : pet.gender === 'female' ? '母' : '未知'

  const handleAdopt = () => {
    if (!isAuthenticated) {
      navigate('/login', { state: { from: `/pets/${id}` } })
      return
    }
    navigate(`/adopt/${id}`)
  }

  return (
    <div className="container py-8">
      {/* 返回按钮 */}
      <Button variant="ghost" className="mb-6" onClick={() => navigate(-1)}>
        <ArrowLeft className="h-4 w-4 mr-2" />
        返回
      </Button>

      <div className="grid lg:grid-cols-2 gap-8">
        {/* 图片展示 */}
        <div className="space-y-4">
          <div className="aspect-square rounded-lg overflow-hidden bg-muted">
            {images.length > 0 ? (
              <img
                src={images[selectedImage]}
                alt={pet.name}
                className="w-full h-full object-cover"
              />
            ) : (
              <div className="w-full h-full flex items-center justify-center text-8xl">
                {pet.type === 'cat' ? '🐱' : pet.type === 'dog' ? '🐕' : '🐾'}
              </div>
            )}
          </div>
          {images.length > 1 && (
            <div className="flex gap-2 overflow-x-auto pb-2">
              {images.map((img, idx) => (
                <button
                  key={idx}
                  onClick={() => setSelectedImage(idx)}
                  className={`w-20 h-20 rounded-md overflow-hidden flex-shrink-0 border-2 ${selectedImage === idx ? 'border-primary' : 'border-transparent'
                    }`}
                >
                  <img src={img} alt="" className="w-full h-full object-cover" />
                </button>
              ))}
            </div>
          )}
        </div>

        {/* 详情信息 */}
        <div className="space-y-6">
          <div>
            <div className="flex items-start justify-between">
              <div>
                <h1 className="text-3xl font-bold mb-2">{pet.name}</h1>
                <p className="text-muted-foreground">
                  {pet.breed || pet.type} · {genderText} · {pet.age}个月
                </p>
              </div>
            </div>
          </div>

          {/* 位置 */}
          {(pet.province || pet.city) && (
            <div className="flex items-center text-muted-foreground">
              <MapPin className="h-4 w-4 mr-2" />
              {pet.province} {pet.city} {pet.district}
            </div>
          )}

          {/* 基本信息 */}
          <Card>
            <CardHeader>
              <CardTitle className="text-lg">基本信息</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="grid grid-cols-2 gap-4 text-sm">
                <div>
                  <span className="text-muted-foreground">体型：</span>
                  {pet.size || '未知'}
                </div>
                <div>
                  <span className="text-muted-foreground">毛色：</span>
                  {pet.color || '未知'}
                </div>
                <div>
                  <span className="text-muted-foreground">体重：</span>
                  {pet.weight ? `${pet.weight}kg` : '未知'}
                </div>
                <div>
                  <span className="text-muted-foreground">健康状况：</span>
                  {pet.health_status || '良好'}
                </div>
                <div className="flex items-center gap-1">
                  {pet.is_vaccinated ? (
                    <CheckCircle className="h-4 w-4 text-green-500" />
                  ) : (
                    <XCircle className="h-4 w-4 text-muted-foreground" />
                  )}
                  已疫苗
                </div>
                <div className="flex items-center gap-1">
                  {pet.is_sterilized ? (
                    <CheckCircle className="h-4 w-4 text-green-500" />
                  ) : (
                    <XCircle className="h-4 w-4 text-muted-foreground" />
                  )}
                  已绝育
                </div>
              </div>
            </CardContent>
          </Card>

          {/* 性格特点 */}
          {pet.character && (
            <Card>
              <CardHeader>
                <CardTitle className="text-lg">性格特点</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-sm">{pet.character}</p>
              </CardContent>
            </Card>
          )}

          {/* 详细描述 */}
          {pet.description && (
            <Card>
              <CardHeader>
                <CardTitle className="text-lg">详细介绍</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-sm whitespace-pre-line">{pet.description}</p>
              </CardContent>
            </Card>
          )}

          {/* 操作按钮 */}
          <div className="flex gap-4">
            <Button className="flex-1" size="lg" onClick={handleAdopt}>
              申请领养
            </Button>
          </div>

          {/* 发布者信息 */}
          {pet.user && (
            <Card>
              <CardContent className="p-4">
                <div className="flex items-center gap-3">
                  <div className="w-10 h-10 rounded-full bg-primary/10 flex items-center justify-center">
                    <span className="text-primary font-medium">
                      {pet.user.nickname?.[0] || pet.user.username[0]}
                    </span>
                  </div>
                  <div>
                    <p className="font-medium">{pet.user.nickname || pet.user.username}</p>
                    <p className="text-xs text-muted-foreground">发布于 {new Date(pet.created_at).toLocaleDateString()}</p>
                  </div>
                </div>
              </CardContent>
            </Card>
          )}
        </div>
      </div>
    </div>
  )
}
