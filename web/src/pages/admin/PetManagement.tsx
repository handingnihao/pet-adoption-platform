import { useState } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Search, Edit, Trash2, Check, X, Eye, Clock, CheckCircle, EyeOff } from 'lucide-react'
import { Button } from '../../components/ui/button'
import { Input } from '../../components/ui/input'
import { Card, CardContent, CardHeader, CardTitle } from '../../components/ui/card'
import { petApi } from '../../lib/api'
import { adminApi } from '../../lib/adminApi'
import type { Pet } from '../../lib/api'

const statusConfig: Record<string, { label: string; icon: typeof Clock; color: string; bg: string }> = {
  pending: { label: '待审核', icon: Clock, color: 'text-yellow-500', bg: 'bg-yellow-50' },
  available: { label: '已发布', icon: CheckCircle, color: 'text-green-500', bg: 'bg-green-50' },
  adopted: { label: '已领养', icon: CheckCircle, color: 'text-blue-500', bg: 'bg-blue-50' },
  offline: { label: '已下架', icon: EyeOff, color: 'text-gray-500', bg: 'bg-gray-50' },
}

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

export function PetManagement() {
  const [keyword, setKeyword] = useState('')
  const [searchKeyword, setSearchKeyword] = useState('')
  const [page, setPage] = useState(1)
  const [editingPet, setEditingPet] = useState<Pet | null>(null)
  const [editForm, setEditForm] = useState<Partial<Pet>>({})
  const queryClient = useQueryClient()

  const { data, isLoading } = useQuery({
    queryKey: ['admin', 'all-pets', searchKeyword, page],
    queryFn: () => {
      if (searchKeyword) {
        return petApi.search({ keyword: searchKeyword, page, page_size: 10 })
      }
      // 管理员默认查看待审核的宠物列表
      return adminApi.getPendingPets({ page, page_size: 10 })
    },
  })

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: number; data: Partial<Pet> }) =>
      adminApi.updatePet(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'all-pets'] })
      setEditingPet(null)
      setEditForm({})
    },
    onError: (error) => {
      alert('更新失败: ' + (error instanceof Error ? error.message : '未知错误'))
    },
  })

  const deleteMutation = useMutation({
    mutationFn: (id: number) => petApi.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'all-pets'] })
    },
  })

  const approveMutation = useMutation({
    mutationFn: (id: number) => adminApi.approvePet(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'all-pets'] })
      alert('审核通过成功')
    },
    onError: (error: Error) => {
      alert('审核失败: ' + error.message)
    },
  })

  const rejectMutation = useMutation({
    mutationFn: (id: number) => adminApi.rejectPet(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'all-pets'] })
    },
  })

  const pets = data?.data?.list || []
  const total = data?.data?.pagination?.total || data?.data?.total || 0
  const totalPages = Math.ceil(total / 10)

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault()
    setSearchKeyword(keyword)
    setPage(1)
  }

  const handleEdit = (pet: Pet) => {
    setEditingPet(pet)
    setEditForm({
      name: pet.name,
      type: pet.type,
      breed: pet.breed,
      gender: pet.gender,
      age: pet.age,
      size: pet.size,
      color: pet.color,
      weight: pet.weight,
      is_vaccinated: pet.is_vaccinated,
      is_sterilized: pet.is_sterilized,
      health_status: pet.health_status,
      description: pet.description,
      character: pet.character,
      cover_photo: pet.cover_photo,
      province: pet.province,
      city: pet.city,
      district: pet.district,
    })
  }

  const handleSave = () => {
    if (editingPet) {
      updateMutation.mutate({ id: editingPet.id, data: editForm })
    }
  }

  const handleDelete = (pet: Pet) => {
    if (confirm(`确定要删除宠物 "${pet.name}" 吗？此操作不可恢复。`)) {
      deleteMutation.mutate(pet.id)
    }
  }

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold">宠物管理</h1>
        <p className="text-muted-foreground">管理所有宠物信息</p>
      </div>

      {/* 搜索 */}
      <Card>
        <CardContent className="p-4">
          <form onSubmit={handleSearch} className="flex gap-4">
            <div className="relative flex-1 max-w-md">
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
              <Input
                placeholder="搜索宠物名称、品种..."
                value={keyword}
                onChange={(e) => setKeyword(e.target.value)}
                className="pl-10"
              />
            </div>
            <Button type="submit">搜索</Button>
            {searchKeyword && (
              <Button type="button" variant="outline" onClick={() => { setSearchKeyword(''); setKeyword(''); }}>
                清除
              </Button>
            )}
          </form>
        </CardContent>
      </Card>

      <div className="grid lg:grid-cols-3 gap-6">
        {/* 列表 */}
        <div className="lg:col-span-2">
          <Card>
            <CardHeader>
              <CardTitle>宠物列表 ({total})</CardTitle>
            </CardHeader>
            <CardContent>
              {isLoading ? (
                <div className="space-y-4">
                  {[...Array(5)].map((_, i) => (
                    <div key={i} className="h-20 bg-muted animate-pulse rounded" />
                  ))}
                </div>
              ) : pets.length === 0 ? (
                <div className="text-center py-8 text-muted-foreground">
                  暂无宠物数据
                </div>
              ) : (
                <div className="space-y-3">
                  {pets.map((pet: Pet) => {
                    const status = statusConfig[pet.status] || statusConfig['pending']
                    const StatusIcon = status.icon
                    return (
                      <div
                        key={pet.id}
                        className={`flex gap-4 p-4 border rounded-lg transition-colors ${editingPet?.id === pet.id ? 'border-primary bg-primary/5' : 'hover:bg-muted/50'
                          }`}
                      >
                        <div className="w-16 h-16 rounded-lg overflow-hidden bg-muted flex-shrink-0">
                          {pet.cover_photo ? (
                            <img src={pet.cover_photo} alt={pet.name} className="w-full h-full object-cover" />
                          ) : (
                            <div className="w-full h-full flex items-center justify-center text-xl">
                              {pet.type === 'cat' ? '🐱' : pet.type === 'dog' ? '🐕' : '🐾'}
                            </div>
                          )}
                        </div>
                        <div className="flex-1 min-w-0">
                          <div className="flex items-center gap-2">
                            <h3 className="font-medium">{pet.name}</h3>
                            <span className={`flex items-center gap-1 px-2 py-0.5 rounded text-xs ${status.bg} ${status.color}`}>
                              <StatusIcon className="h-3 w-3" />
                              {status.label}
                            </span>
                          </div>
                          <p className="text-sm text-muted-foreground">
                            {pet.breed || pet.type} · {pet.gender === 'male' ? '公' : '母'} · {pet.age}个月
                          </p>
                          <p className="text-xs text-muted-foreground">
                            发布者: {pet.user?.nickname || pet.user?.username || '未知'} · ID: {pet.id}
                          </p>
                        </div>
                        <div className="flex items-center gap-1">
                          {pet.status === 'pending' && (
                            <>
                              <Button
                                size="sm"
                                variant="ghost"
                                className="text-green-600"
                                onClick={() => approveMutation.mutate(pet.id)}
                                disabled={approveMutation.isPending}
                                title="发布宠物"
                              >
                                <Check className="h-4 w-4" />
                              </Button>
                              <Button
                                size="sm"
                                variant="ghost"
                                className="text-red-600"
                                onClick={() => rejectMutation.mutate(pet.id)}
                                disabled={rejectMutation.isPending}
                                title="拒绝"
                              >
                                <X className="h-4 w-4" />
                              </Button>
                            </>
                          )}
                          {pet.status === 'available' && (
                            <Button
                              size="sm"
                              variant="ghost"
                              className="text-orange-600"
                              onClick={() => {
                                if (confirm(`确定要下架宠物 "${pet.name}" 吗？`)) {
                                  updateMutation.mutate({ id: pet.id, data: { status: 'offline' } })
                                }
                              }}
                              disabled={updateMutation.isPending}
                              title="下架"
                            >
                              <EyeOff className="h-4 w-4" />
                            </Button>
                          )}
                          <Button size="sm" variant="ghost" onClick={() => handleEdit(pet)}>
                            <Edit className="h-4 w-4" />
                          </Button>
                          <Button
                            size="sm"
                            variant="ghost"
                            className="text-destructive"
                            onClick={() => handleDelete(pet)}
                            disabled={deleteMutation.isPending}
                          >
                            <Trash2 className="h-4 w-4" />
                          </Button>
                        </div>
                      </div>
                    )
                  })}
                </div>
              )}

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

        {/* 编辑面板 */}
        <div>
          <Card className="sticky top-20">
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                {editingPet ? (
                  <>
                    <Edit className="h-5 w-5" />
                    编辑宠物
                  </>
                ) : (
                  <>
                    <Eye className="h-5 w-5" />
                    选择编辑
                  </>
                )}
              </CardTitle>
            </CardHeader>
            <CardContent>
              {editingPet ? (
                <div className="space-y-4 max-h-[calc(100vh-200px)] overflow-y-auto pr-2">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">名称</label>
                    <Input
                      value={editForm.name || ''}
                      onChange={(e) => setEditForm({ ...editForm, name: e.target.value })}
                    />
                  </div>

                  <div className="grid grid-cols-2 gap-2">
                    <div className="space-y-2">
                      <label className="text-sm font-medium">类型</label>
                      <select
                        className="w-full h-10 px-3 border rounded-md text-sm"
                        value={editForm.type || ''}
                        onChange={(e) => setEditForm({ ...editForm, type: e.target.value })}
                      >
                        {petTypes.map(t => <option key={t.value} value={t.value}>{t.label}</option>)}
                      </select>
                    </div>
                    <div className="space-y-2">
                      <label className="text-sm font-medium">性别</label>
                      <select
                        className="w-full h-10 px-3 border rounded-md text-sm"
                        value={editForm.gender || ''}
                        onChange={(e) => setEditForm({ ...editForm, gender: e.target.value })}
                      >
                        <option value="male">公</option>
                        <option value="female">母</option>
                      </select>
                    </div>
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-medium">品种</label>
                    <Input
                      value={editForm.breed || ''}
                      onChange={(e) => setEditForm({ ...editForm, breed: e.target.value })}
                    />
                  </div>

                  <div className="grid grid-cols-2 gap-2">
                    <div className="space-y-2">
                      <label className="text-sm font-medium">年龄(月)</label>
                      <Input
                        type="number"
                        value={editForm.age || ''}
                        onChange={(e) => setEditForm({ ...editForm, age: parseInt(e.target.value) })}
                      />
                    </div>
                    <div className="space-y-2">
                      <label className="text-sm font-medium">体型</label>
                      <select
                        className="w-full h-10 px-3 border rounded-md text-sm"
                        value={editForm.size || ''}
                        onChange={(e) => setEditForm({ ...editForm, size: e.target.value })}
                      >
                        <option value="">未知</option>
                        {sizes.map(s => <option key={s.value} value={s.value}>{s.label}</option>)}
                      </select>
                    </div>
                  </div>

                  <div className="grid grid-cols-2 gap-2">
                    <div className="space-y-2">
                      <label className="text-sm font-medium">颜色</label>
                      <Input
                        value={editForm.color || ''}
                        onChange={(e) => setEditForm({ ...editForm, color: e.target.value })}
                      />
                    </div>
                    <div className="space-y-2">
                      <label className="text-sm font-medium">体重(kg)</label>
                      <Input
                        type="number"
                        step="0.1"
                        value={editForm.weight || ''}
                        onChange={(e) => setEditForm({ ...editForm, weight: parseFloat(e.target.value) })}
                      />
                    </div>
                  </div>

                  <div className="flex gap-4">
                    <label className="flex items-center gap-2 text-sm">
                      <input
                        type="checkbox"
                        checked={editForm.is_vaccinated || false}
                        onChange={(e) => setEditForm({ ...editForm, is_vaccinated: e.target.checked })}
                        className="rounded"
                      />
                      已疫苗
                    </label>
                    <label className="flex items-center gap-2 text-sm">
                      <input
                        type="checkbox"
                        checked={editForm.is_sterilized || false}
                        onChange={(e) => setEditForm({ ...editForm, is_sterilized: e.target.checked })}
                        className="rounded"
                      />
                      已绝育
                    </label>
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-medium">健康状况</label>
                    <Input
                      value={editForm.health_status || ''}
                      onChange={(e) => setEditForm({ ...editForm, health_status: e.target.value })}
                    />
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-medium">性格特点</label>
                    <Input
                      value={editForm.character || ''}
                      onChange={(e) => setEditForm({ ...editForm, character: e.target.value })}
                    />
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-medium">描述</label>
                    <textarea
                      className="w-full min-h-[80px] px-3 py-2 border rounded-md text-sm resize-none"
                      value={editForm.description || ''}
                      onChange={(e) => setEditForm({ ...editForm, description: e.target.value })}
                    />
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-medium">封面图片URL</label>
                    <Input
                      value={editForm.cover_photo || ''}
                      onChange={(e) => setEditForm({ ...editForm, cover_photo: e.target.value })}
                    />
                  </div>

                  <div className="grid grid-cols-3 gap-2">
                    <div className="space-y-2">
                      <label className="text-sm font-medium">省份</label>
                      <Input
                        value={editForm.province || ''}
                        onChange={(e) => setEditForm({ ...editForm, province: e.target.value })}
                      />
                    </div>
                    <div className="space-y-2">
                      <label className="text-sm font-medium">城市</label>
                      <Input
                        value={editForm.city || ''}
                        onChange={(e) => setEditForm({ ...editForm, city: e.target.value })}
                      />
                    </div>
                    <div className="space-y-2">
                      <label className="text-sm font-medium">区县</label>
                      <Input
                        value={editForm.district || ''}
                        onChange={(e) => setEditForm({ ...editForm, district: e.target.value })}
                      />
                    </div>
                  </div>

                  <div className="flex gap-2 pt-4">
                    <Button
                      className="flex-1"
                      onClick={handleSave}
                      disabled={updateMutation.isPending}
                    >
                      {updateMutation.isPending ? '保存中...' : '保存修改'}
                    </Button>
                    <Button
                      variant="outline"
                      onClick={() => { setEditingPet(null); setEditForm({}); }}
                    >
                      取消
                    </Button>
                  </div>
                </div>
              ) : (
                <p className="text-center text-muted-foreground py-8">
                  点击列表中的编辑按钮修改宠物信息
                </p>
              )}
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
