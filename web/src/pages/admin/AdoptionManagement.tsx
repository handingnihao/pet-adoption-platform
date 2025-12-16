import { useState } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Check, X, Eye, Clock, CheckCircle, XCircle } from 'lucide-react'
import { Button } from '../../components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '../../components/ui/card'
import { adminApi } from '../../lib/adminApi'
import type { AdoptionApplication } from '../../lib/api'

const statusTabs = [
  { value: 'all', label: '全部' },
  { value: 'pending', label: '待审核' },
  { value: 'approved', label: '已通过' },
  { value: 'rejected', label: '已拒绝' },
]

const statusConfig = {
  'pending': { label: '待审核', icon: Clock, color: 'text-yellow-500', bg: 'bg-yellow-50' },
  'reviewing': { label: '审核中', icon: Clock, color: 'text-blue-500', bg: 'bg-blue-50' },
  'interview': { label: '待面试', icon: Clock, color: 'text-purple-500', bg: 'bg-purple-50' },
  'home_visit': { label: '待家访', icon: Clock, color: 'text-indigo-500', bg: 'bg-indigo-50' },
  'approved': { label: '已通过', icon: CheckCircle, color: 'text-green-500', bg: 'bg-green-50' },
  'rejected': { label: '已拒绝', icon: XCircle, color: 'text-red-500', bg: 'bg-red-50' },
  'cancelled': { label: '已取消', icon: XCircle, color: 'text-gray-500', bg: 'bg-gray-50' },
}

export function AdoptionManagement() {
  const [statusFilter, setStatusFilter] = useState<string>('all')
  const [page, setPage] = useState(1)
  const [selectedApp, setSelectedApp] = useState<AdoptionApplication | null>(null)
  const queryClient = useQueryClient()

  const { data, isLoading } = useQuery({
    queryKey: ['admin', 'applications', statusFilter, page],
    queryFn: () => {
      if (statusFilter === 'pending') {
        return adminApi.getPendingApplications({ page, page_size: 10 })
      }
      return adminApi.getAllApplications({ 
        page, 
        page_size: 10,
        status: statusFilter !== 'all' ? statusFilter : undefined
      })
    },
  })

  const reviewMutation = useMutation({
    mutationFn: ({ id, action }: { id: number; action: string }) =>
      adminApi.reviewApplication(id, { action, comment: '', reason: '' }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'applications'] })
      queryClient.invalidateQueries({ queryKey: ['admin', 'adoption-statistics'] })
      setSelectedApp(null)
    },
  })

  const applications = data?.data?.list || []
  const total = data?.data?.pagination?.total || data?.data?.total || 0
  const totalPages = Math.ceil(total / 10)

  const handleApprove = (app: AdoptionApplication) => {
    if (confirm('确定要通过这个领养申请吗？')) {
      reviewMutation.mutate({ id: app.id, action: 'approve' })
    }
  }

  const handleReject = (app: AdoptionApplication) => {
    const reason = prompt('请输入拒绝原因：')
    if (reason) {
      reviewMutation.mutate({ id: app.id, action: 'reject' })
    }
  }

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold">领养管理</h1>
        <p className="text-muted-foreground">审核和管理领养申请</p>
      </div>

      {/* 状态筛选 */}
      <div className="flex gap-2">
        {statusTabs.map((tab) => (
          <Button
            key={tab.value}
            variant={statusFilter === tab.value ? 'default' : 'outline'}
            size="sm"
            onClick={() => { setStatusFilter(tab.value); setPage(1); }}
          >
            {tab.label}
          </Button>
        ))}
      </div>

      <div className="grid lg:grid-cols-3 gap-6">
        {/* 列表 */}
        <div className="lg:col-span-2">
          <Card>
            <CardHeader>
              <CardTitle>申请列表 ({total})</CardTitle>
            </CardHeader>
            <CardContent>
              {isLoading ? (
                <div className="space-y-4">
                  {[...Array(5)].map((_, i) => (
                    <div key={i} className="h-24 bg-muted animate-pulse rounded" />
                  ))}
                </div>
              ) : applications.length === 0 ? (
                <div className="text-center py-8 text-muted-foreground">
                  暂无申请数据
                </div>
              ) : (
                <div className="space-y-4">
                  {applications.map((app: AdoptionApplication) => {
                    const status = statusConfig[app.status as keyof typeof statusConfig]
                    const StatusIcon = status?.icon || Clock
                    return (
                      <div
                        key={app.id}
                        className={`flex gap-4 p-4 border rounded-lg cursor-pointer transition-colors ${
                          selectedApp?.id === app.id ? 'border-primary bg-primary/5' : 'hover:bg-muted/50'
                        }`}
                        onClick={() => setSelectedApp(app)}
                      >
                        <div className="w-16 h-16 rounded-lg overflow-hidden bg-muted flex-shrink-0">
                          {app.pet?.cover_photo ? (
                            <img src={app.pet.cover_photo} alt={app.pet.name} className="w-full h-full object-cover" />
                          ) : (
                            <div className="w-full h-full flex items-center justify-center text-xl">🐾</div>
                          )}
                        </div>
                        <div className="flex-1 min-w-0">
                          <div className="flex items-center gap-2">
                            <h3 className="font-medium">{app.pet?.name || '未知宠物'}</h3>
                            <span className={`flex items-center gap-1 px-2 py-0.5 rounded text-xs ${status?.bg} ${status?.color}`}>
                              <StatusIcon className="h-3 w-3" />
                              {status?.label}
                            </span>
                          </div>
                          <p className="text-sm text-muted-foreground">
                            申请人: {app.user?.nickname || app.user?.username || '未知'}
                          </p>
                          <p className="text-xs text-muted-foreground mt-1">
                            {new Date(app.created_at).toLocaleString()}
                          </p>
                        </div>
                        {app.status === 'pending' && (
                          <div className="flex items-center gap-2">
                            <Button
                              size="sm"
                              variant="outline"
                              className="text-green-600"
                              onClick={(e) => { e.stopPropagation(); handleApprove(app); }}
                              disabled={reviewMutation.isPending}
                            >
                              <Check className="h-4 w-4" />
                            </Button>
                            <Button
                              size="sm"
                              variant="outline"
                              className="text-red-600"
                              onClick={(e) => { e.stopPropagation(); handleReject(app); }}
                              disabled={reviewMutation.isPending}
                            >
                              <X className="h-4 w-4" />
                            </Button>
                          </div>
                        )}
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

        {/* 详情 */}
        <div>
          <Card className="sticky top-20">
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Eye className="h-5 w-5" />
                申请详情
              </CardTitle>
            </CardHeader>
            <CardContent>
              {selectedApp ? (
                <div className="space-y-4">
                  <div className="p-3 bg-muted rounded-lg">
                    <h4 className="font-medium mb-2">宠物信息</h4>
                    <p className="text-sm">{selectedApp.pet?.name} ({selectedApp.pet?.breed || selectedApp.pet?.type})</p>
                  </div>

                  <div className="p-3 bg-muted rounded-lg">
                    <h4 className="font-medium mb-2">申请人</h4>
                    <p className="text-sm">{selectedApp.user?.nickname || selectedApp.user?.username}</p>
                    <p className="text-xs text-muted-foreground">{selectedApp.user?.email || selectedApp.user?.phone}</p>
                  </div>

                  <div>
                    <h4 className="font-medium mb-2">申请编号</h4>
                    <p className="text-sm bg-muted p-3 rounded-lg">{selectedApp.application_no}</p>
                  </div>

                  <div>
                    <h4 className="font-medium mb-2">申请人信息</h4>
                    <div className="text-sm bg-muted p-3 rounded-lg space-y-1">
                      <p>姓名：{selectedApp.applicant_name}</p>
                      <p>电话：{selectedApp.applicant_phone}</p>
                      <p>地址：{selectedApp.applicant_address}</p>
                    </div>
                  </div>

                  <div>
                    <h4 className="font-medium mb-2">住房情况</h4>
                    <div className="text-sm bg-muted p-3 rounded-lg space-y-1">
                      <p>类型：{selectedApp.housing_type}</p>
                      <p>面积：{selectedApp.housing_area}㎡</p>
                      <p>院子：{selectedApp.has_yard ? '有' : '无'}</p>
                    </div>
                  </div>

                  <div>
                    <h4 className="font-medium mb-2">家庭情况</h4>
                    <div className="text-sm bg-muted p-3 rounded-lg space-y-1">
                      <p>成员数：{selectedApp.family_members}人</p>
                      <p>有孩子：{selectedApp.has_children ? '是' : '否'}</p>
                      {selectedApp.children_age && <p>孩子年龄：{selectedApp.children_age}</p>}
                    </div>
                  </div>

                  <div>
                    <h4 className="font-medium mb-2">领养原因</h4>
                    <p className="text-sm bg-muted p-3 rounded-lg">{selectedApp.adoption_reason}</p>
                  </div>

                  <div>
                    <h4 className="font-medium mb-2">照顾计划</h4>
                    <p className="text-sm bg-muted p-3 rounded-lg">{selectedApp.how_to_care}</p>
                  </div>

                  <div>
                    <h4 className="font-medium mb-2">应急计划</h4>
                    <p className="text-sm bg-muted p-3 rounded-lg">{selectedApp.emergency_plan}</p>
                  </div>

                  <div className="flex items-center gap-2 text-sm">
                    <span className="text-muted-foreground">家人同意：</span>
                    {selectedApp.family_agree ? (
                      <span className="text-green-600">是</span>
                    ) : (
                      <span className="text-red-600">否</span>
                    )}
                  </div>

                  {selectedApp.status === 'pending' && (
                    <div className="flex gap-2 pt-4">
                      <Button
                        className="flex-1"
                        onClick={() => handleApprove(selectedApp)}
                        disabled={reviewMutation.isPending}
                      >
                        <Check className="h-4 w-4 mr-2" />
                        通过
                      </Button>
                      <Button
                        variant="destructive"
                        className="flex-1"
                        onClick={() => handleReject(selectedApp)}
                        disabled={reviewMutation.isPending}
                      >
                        <X className="h-4 w-4 mr-2" />
                        拒绝
                      </Button>
                    </div>
                  )}
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
