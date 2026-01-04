import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Plus, Trash2, EyeOff, Clock, CheckCircle, XCircle } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Card, CardContent } from '../components/ui/card'
import { petApi } from '../lib/api'
import type { Pet } from '../lib/api'
import { useAuthStore } from '../store/auth'

const statusConfig: Record<string | number, { label: string; icon: typeof Clock; color: string; bg: string }> = {
  // 数字状态 (旧版本兼容)
  0: { label: '待审核', icon: Clock, color: 'text-yellow-500', bg: 'bg-yellow-50' },
  1: { label: '已发布', icon: CheckCircle, color: 'text-green-500', bg: 'bg-green-50' },
  2: { label: '已下架', icon: EyeOff, color: 'text-gray-500', bg: 'bg-gray-50' },
  3: { label: '已拒绝', icon: XCircle, color: 'text-red-500', bg: 'bg-red-50' },
  // 字符串状态 (后端实际返回)
  'pending': { label: '待审核', icon: Clock, color: 'text-yellow-500', bg: 'bg-yellow-50' },
  'available': { label: '已发布', icon: CheckCircle, color: 'text-green-500', bg: 'bg-green-50' },
  'adopted': { label: '已领养', icon: CheckCircle, color: 'text-blue-500', bg: 'bg-blue-50' },
  'offline': { label: '已下架', icon: EyeOff, color: 'text-gray-500', bg: 'bg-gray-50' },
  'rejected': { label: '已拒绝', icon: XCircle, color: 'text-red-500', bg: 'bg-red-50' },
}

export function MyPets() {
  const navigate = useNavigate()
  const { isAuthenticated } = useAuthStore()
  const queryClient = useQueryClient()
  const [page, setPage] = useState(1)

  const { data, isLoading } = useQuery({
    queryKey: ['my-pets', page],
    queryFn: () => petApi.getMyPets({ page, page_size: 10 }),
    enabled: isAuthenticated,
  })

  const deleteMutation = useMutation({
    mutationFn: (id: number) => petApi.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['my-pets'] })
    },
  })

  if (!isAuthenticated) {
    navigate('/login')
    return null
  }

  const pets = data?.data?.list || []
  const total = data?.data?.pagination?.total || data?.data?.total || 0
  const totalPages = Math.ceil(total / 10)

  const handleDelete = (pet: Pet) => {
    if (confirm(`确定要删除宠物 "${pet.name}" 吗？此操作不可恢复。`)) {
      deleteMutation.mutate(pet.id)
    }
  }

  return (
    <div className="container py-8">
      <div className="flex items-center justify-between mb-6">
        <div>
          <h1 className="text-2xl font-bold">我的宠物</h1>
          <p className="text-muted-foreground">管理您发布的宠物信息</p>
        </div>
        <Button onClick={() => navigate('/pets/create')}>
          <Plus className="h-4 w-4 mr-2" />
          发布宠物
        </Button>
      </div>

      {isLoading ? (
        <div className="grid md:grid-cols-2 gap-4">
          {[...Array(4)].map((_, i) => (
            <Card key={i} className="animate-pulse">
              <CardContent className="p-4">
                <div className="flex gap-4">
                  <div className="w-24 h-24 bg-muted rounded" />
                  <div className="flex-1 space-y-2">
                    <div className="h-5 bg-muted rounded w-1/3" />
                    <div className="h-4 bg-muted rounded w-1/2" />
                  </div>
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      ) : pets.length === 0 ? (
        <div className="text-center py-16">
          <div className="text-6xl mb-4">🐾</div>
          <h3 className="text-lg font-medium mb-2">您还没有发布宠物</h3>
          <p className="text-muted-foreground mb-4">发布您的宠物，帮助它们找到新家</p>
          <Button onClick={() => navigate('/pets/create')}>
            <Plus className="h-4 w-4 mr-2" />
            发布宠物
          </Button>
        </div>
      ) : (
        <>
          <div className="grid md:grid-cols-2 gap-4">
            {pets.map((pet: Pet) => {
              const status = statusConfig[pet.status] || statusConfig[0]
              const StatusIcon = status.icon
              return (
                <Card key={pet.id}>
                  <CardContent className="p-4">
                    <div className="flex gap-4">
                      <div
                        className="w-24 h-24 rounded-lg overflow-hidden bg-muted cursor-pointer flex-shrink-0"
                        onClick={() => navigate(`/pets/${pet.id}`)}
                      >
                        {pet.cover_photo ? (
                          <img src={pet.cover_photo} alt={pet.name} className="w-full h-full object-cover" />
                        ) : (
                          <div className="w-full h-full flex items-center justify-center text-3xl">
                            {pet.type === 'cat' ? '🐱' : pet.type === 'dog' ? '🐕' : '🐾'}
                          </div>
                        )}
                      </div>
                      <div className="flex-1 min-w-0">
                        <div className="flex items-center gap-2 mb-1">
                          <h3 className="font-medium truncate">{pet.name}</h3>
                          <span className={`flex items-center gap-1 px-2 py-0.5 rounded text-xs ${status.bg} ${status.color}`}>
                            <StatusIcon className="h-3 w-3" />
                            {status.label}
                          </span>
                        </div>
                        <p className="text-sm text-muted-foreground">
                          {pet.breed || pet.type} · {pet.gender === 'male' ? '公' : '母'} · {pet.age}个月
                        </p>
                        <p className="text-xs text-muted-foreground mt-1">
                          浏览 {pet.view_count} 次 · {new Date(pet.created_at).toLocaleDateString()}
                        </p>
                        <div className="flex gap-2 mt-3">
                          <Button
                            size="sm"
                            variant="outline"
                            className="text-destructive"
                            onClick={() => handleDelete(pet)}
                            disabled={deleteMutation.isPending}
                          >
                            <Trash2 className="h-3 w-3 mr-1" />
                            删除
                          </Button>
                        </div>
                      </div>
                    </div>
                  </CardContent>
                </Card>
              )
            })}
          </div>

          {totalPages > 1 && (
            <div className="flex justify-center gap-2 mt-6">
              <Button variant="outline" disabled={page <= 1} onClick={() => setPage(page - 1)}>
                上一页
              </Button>
              <span className="flex items-center px-4">{page} / {totalPages}</span>
              <Button variant="outline" disabled={page >= totalPages} onClick={() => setPage(page + 1)}>
                下一页
              </Button>
            </div>
          )}
        </>
      )}
    </div>
  )
}
