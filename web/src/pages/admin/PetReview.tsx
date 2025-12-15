import { useState } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Check, X, Eye } from 'lucide-react'
import { Button } from '../../components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '../../components/ui/card'
import { adminApi } from '../../lib/adminApi'
import type { Pet } from '../../lib/api'

export function PetReview() {
  const [page, setPage] = useState(1)
  const [selectedPet, setSelectedPet] = useState<Pet | null>(null)
  const queryClient = useQueryClient()

  const { data, isLoading } = useQuery({
    queryKey: ['admin', 'pending-pets', page],
    queryFn: () => adminApi.getPendingPets({ page, page_size: 10 }),
  })

  const approveMutation = useMutation({
    mutationFn: (id: number) => adminApi.approvePet(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'pending-pets'] })
      queryClient.invalidateQueries({ queryKey: ['admin', 'pet-statistics'] })
      setSelectedPet(null)
    },
  })

  const rejectMutation = useMutation({
    mutationFn: (id: number) => adminApi.rejectPet(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'pending-pets'] })
      queryClient.invalidateQueries({ queryKey: ['admin', 'pet-statistics'] })
      setSelectedPet(null)
    },
  })

  const pets = data?.data?.list || []
  const total = data?.data?.total || 0
  const totalPages = Math.ceil(total / 10)

  const handleApprove = (pet: Pet) => {
    if (confirm(`确定要通过宠物 "${pet.name}" 的审核吗？`)) {
      approveMutation.mutate(pet.id)
    }
  }

  const handleReject = (pet: Pet) => {
    if (confirm(`确定要拒绝宠物 "${pet.name}" 的审核吗？`)) {
      rejectMutation.mutate(pet.id)
    }
  }

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold">宠物审核</h1>
        <p className="text-muted-foreground">审核待发布的宠物信息</p>
      </div>

      <div className="grid lg:grid-cols-3 gap-6">
        {/* 列表 */}
        <div className="lg:col-span-2">
          <Card>
            <CardHeader>
              <CardTitle>待审核列表 ({total})</CardTitle>
            </CardHeader>
            <CardContent>
              {isLoading ? (
                <div className="space-y-4">
                  {[...Array(5)].map((_, i) => (
                    <div key={i} className="h-24 bg-muted animate-pulse rounded" />
                  ))}
                </div>
              ) : pets.length === 0 ? (
                <div className="text-center py-8">
                  <div className="text-4xl mb-2">🎉</div>
                  <p className="text-muted-foreground">暂无待审核宠物</p>
                </div>
              ) : (
                <div className="space-y-4">
                  {pets.map((pet: Pet) => (
                    <div
                      key={pet.id}
                      className={`flex gap-4 p-4 border rounded-lg cursor-pointer transition-colors ${
                        selectedPet?.id === pet.id ? 'border-primary bg-primary/5' : 'hover:bg-muted/50'
                      }`}
                      onClick={() => setSelectedPet(pet)}
                    >
                      <div className="w-20 h-20 rounded-lg overflow-hidden bg-muted flex-shrink-0">
                        {pet.cover_photo ? (
                          <img src={pet.cover_photo} alt={pet.name} className="w-full h-full object-cover" />
                        ) : (
                          <div className="w-full h-full flex items-center justify-center text-2xl">
                            {pet.type === 'cat' ? '🐱' : pet.type === 'dog' ? '🐕' : '🐾'}
                          </div>
                        )}
                      </div>
                      <div className="flex-1 min-w-0">
                        <h3 className="font-medium">{pet.name}</h3>
                        <p className="text-sm text-muted-foreground">
                          {pet.breed || pet.type} · {pet.gender === 'male' ? '公' : '母'} · {pet.age}个月
                        </p>
                        <p className="text-xs text-muted-foreground mt-1">
                          发布者: {pet.user?.nickname || pet.user?.username || '未知'}
                        </p>
                      </div>
                      <div className="flex items-center gap-2">
                        <Button
                          size="sm"
                          variant="outline"
                          className="text-green-600"
                          onClick={(e) => { e.stopPropagation(); handleApprove(pet); }}
                          disabled={approveMutation.isPending}
                        >
                          <Check className="h-4 w-4" />
                        </Button>
                        <Button
                          size="sm"
                          variant="outline"
                          className="text-red-600"
                          onClick={(e) => { e.stopPropagation(); handleReject(pet); }}
                          disabled={rejectMutation.isPending}
                        >
                          <X className="h-4 w-4" />
                        </Button>
                      </div>
                    </div>
                  ))}
                </div>
              )}

              {/* 分页 */}
              {totalPages > 1 && (
                <div className="flex justify-center gap-2 mt-4">
                  <Button variant="outline" size="sm" disabled={page <= 1} onClick={() => setPage(page - 1)}>
                    上一页
                  </Button>
                  <span className="flex items-center px-4 text-sm">{page} / {totalPages}</span>
                  <Button variant="outline" size="sm" disabled={page >= totalPages} onClick={() => setPage(page + 1)}>
                    下一页
                  </Button>
                </div>
              )}
            </CardContent>
          </Card>
        </div>

        {/* 详情预览 */}
        <div>
          <Card className="sticky top-20">
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Eye className="h-5 w-5" />
                详情预览
              </CardTitle>
            </CardHeader>
            <CardContent>
              {selectedPet ? (
                <div className="space-y-4">
                  {selectedPet.cover_photo && (
                    <img
                      src={selectedPet.cover_photo}
                      alt={selectedPet.name}
                      className="w-full aspect-square object-cover rounded-lg"
                    />
                  )}
                  <div>
                    <h3 className="font-bold text-lg">{selectedPet.name}</h3>
                    <p className="text-sm text-muted-foreground">
                      {selectedPet.breed || selectedPet.type} · {selectedPet.age}个月
                    </p>
                  </div>
                  <div className="text-sm space-y-2">
                    <p><span className="text-muted-foreground">性别：</span>{selectedPet.gender === 'male' ? '公' : '母'}</p>
                    <p><span className="text-muted-foreground">体型：</span>{selectedPet.size || '未知'}</p>
                    <p><span className="text-muted-foreground">颜色：</span>{selectedPet.color || '未知'}</p>
                    <p><span className="text-muted-foreground">疫苗：</span>{selectedPet.is_vaccinated ? '已接种' : '未接种'}</p>
                    <p><span className="text-muted-foreground">绝育：</span>{selectedPet.is_sterilized ? '已绝育' : '未绝育'}</p>
                    <p><span className="text-muted-foreground">位置：</span>{selectedPet.city} {selectedPet.district}</p>
                  </div>
                  {selectedPet.description && (
                    <div>
                      <p className="text-sm text-muted-foreground mb-1">描述：</p>
                      <p className="text-sm">{selectedPet.description}</p>
                    </div>
                  )}
                  <div className="flex gap-2 pt-4">
                    <Button
                      className="flex-1"
                      onClick={() => handleApprove(selectedPet)}
                      disabled={approveMutation.isPending}
                    >
                      <Check className="h-4 w-4 mr-2" />
                      通过
                    </Button>
                    <Button
                      variant="destructive"
                      className="flex-1"
                      onClick={() => handleReject(selectedPet)}
                      disabled={rejectMutation.isPending}
                    >
                      <X className="h-4 w-4 mr-2" />
                      拒绝
                    </Button>
                  </div>
                </div>
              ) : (
                <p className="text-center text-muted-foreground py-8">
                  点击左侧列表查看详情
                </p>
              )}
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
