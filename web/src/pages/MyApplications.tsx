import { useNavigate, useLocation } from 'react-router-dom'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Clock, CheckCircle, XCircle, Trash2 } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Card, CardContent } from '../components/ui/card'
import { adoptionApi } from '../lib/api'
import type { AdoptionApplication } from '../lib/api'
import { useAuthStore } from '../store/auth'

const statusConfig = {
  'pending': { label: '待审核', icon: Clock, color: 'text-yellow-500', bg: 'bg-yellow-50' },
  'reviewing': { label: '审核中', icon: Clock, color: 'text-blue-500', bg: 'bg-blue-50' },
  'interview': { label: '待面试', icon: Clock, color: 'text-purple-500', bg: 'bg-purple-50' },
  'home_visit': { label: '待家访', icon: Clock, color: 'text-indigo-500', bg: 'bg-indigo-50' },
  'approved': { label: '已通过', icon: CheckCircle, color: 'text-green-500', bg: 'bg-green-50' },
  'rejected': { label: '已拒绝', icon: XCircle, color: 'text-red-500', bg: 'bg-red-50' },
  'cancelled': { label: '已取消', icon: XCircle, color: 'text-gray-500', bg: 'bg-gray-50' },
}

export function MyApplications() {
  const navigate = useNavigate()
  const location = useLocation()
  const queryClient = useQueryClient()
  const { isAuthenticated } = useAuthStore()

  const message = location.state?.message

  const { data, isLoading } = useQuery({
    queryKey: ['my-applications'],
    queryFn: async () => {
      const result = await adoptionApi.getMyApplications({ page: 1, page_size: 50 })
      console.log('我的申请API返回:', result)
      console.log('申请列表数据:', result?.data)
      return result
    },
    enabled: isAuthenticated,
  })

  const cancelMutation = useMutation({
    mutationFn: (id: number) => adoptionApi.cancelApplication(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['my-applications'] })
    },
  })

  if (!isAuthenticated) {
    navigate('/login')
    return null
  }

  const applications = data?.data?.list || []

  return (
    <div className="container py-8">
      <h1 className="text-2xl font-bold mb-6">我的领养申请</h1>

      {message && (
        <div className="p-4 mb-6 bg-green-50 text-green-700 rounded-lg">
          {message}
        </div>
      )}

      {isLoading ? (
        <div className="space-y-4">
          {[...Array(3)].map((_, i) => (
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
      ) : applications.length === 0 ? (
        <div className="text-center py-16">
          <div className="text-6xl mb-4">📋</div>
          <h3 className="text-lg font-medium mb-2">暂无申请记录</h3>
          <p className="text-muted-foreground mb-4">快去找到你心仪的宠物吧</p>
          <Button onClick={() => navigate('/pets')}>浏览宠物</Button>
        </div>
      ) : (
        <div className="space-y-4">
          {applications.map((app: AdoptionApplication) => {
            const status = statusConfig[app.status as keyof typeof statusConfig] || statusConfig['pending']
            const StatusIcon = status.icon

            return (
              <Card key={app.id}>
                <CardContent className="p-4">
                  <div className="flex gap-4">
                    {/* 宠物图片 */}
                    <div
                      className="w-24 h-24 rounded-lg overflow-hidden bg-muted cursor-pointer flex-shrink-0"
                      onClick={() => app.pet_info && navigate(`/pets/${app.pet_info.id}`)}
                    >
                      {app.pet_info?.cover_photo ? (
                        <img
                          src={app.pet_info.cover_photo}
                          alt={app.pet_info.name}
                          className="w-full h-full object-cover"
                        />
                      ) : (
                        <div className="w-full h-full flex items-center justify-center text-3xl">
                          🐾
                        </div>
                      )}
                    </div>

                    {/* 申请信息 */}
                    <div className="flex-1 min-w-0">
                      <div className="flex items-start justify-between">
                        <div>
                          <h3 className="font-medium">
                            {app.pet_info?.name || '未知宠物'}
                          </h3>
                          <p className="text-sm text-muted-foreground">
                            {app.pet_info?.breed || app.pet_info?.type} · {app.pet_info?.age}个月
                          </p>
                        </div>
                        <div className={`flex items-center gap-1 px-2 py-1 rounded text-sm ${status.bg} ${status.color}`}>
                          <StatusIcon className="w-4 h-4" />
                          {status.label}
                        </div>
                      </div>

                      <div className="mt-2 space-y-1">
                        <p className="text-sm text-muted-foreground line-clamp-2">
                          <span className="font-medium">申请人：</span>{app.applicant_name} · {app.applicant_phone}
                        </p>
                        <p className="text-sm text-muted-foreground line-clamp-2">
                          <span className="font-medium">领养理由：</span>{app.adoption_reason}
                        </p>
                        <p className="text-sm text-muted-foreground">
                          <span className="font-medium">申请编号：</span>{app.application_no}
                        </p>
                      </div>

                      <div className="flex items-center justify-between mt-3">
                        <span className="text-xs text-muted-foreground">
                          申请时间：{new Date(app.created_at).toLocaleDateString()}
                        </span>

                        {app.status === 'pending' && (
                          <Button
                            variant="ghost"
                            size="sm"
                            className="text-destructive hover:text-destructive"
                            onClick={() => {
                              if (confirm('确定要取消这个申请吗？')) {
                                cancelMutation.mutate(app.id)
                              }
                            }}
                          >
                            <Trash2 className="w-4 h-4 mr-1" />
                            取消申请
                          </Button>
                        )}
                      </div>
                    </div>
                  </div>
                </CardContent>
              </Card>
            )
          })}
        </div>
      )}
    </div>
  )
}
