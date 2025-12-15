import { useState } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Check, X, Eye, Clock, CheckCircle, XCircle } from 'lucide-react'
import { Button } from '../../components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '../../components/ui/card'
import { adminApi } from '../../lib/adminApi'
import type { AdoptionApplication } from '../../lib/api'

const statusTabs = [
  { value: -1, label: '全部' },
  { value: 0, label: '待审核' },
  { value: 1, label: '已通过' },
  { value: 2, label: '已拒绝' },
]

const statusConfig = {
  0: { label: '待审核', icon: Clock, color: 'text-yellow-500', bg: 'bg-yellow-50' },
  1: { label: '已通过', icon: CheckCircle, color: 'text-green-500', bg: 'bg-green-50' },
  2: { label: '已拒绝', icon: XCircle, color: 'text-red-500', bg: 'bg-red-50' },
}

export function AdoptionManagement() {
  const [statusFilter, setStatusFilter] = useState(-1)
  const [page, setPage] = useState(1)
  const [selectedApp, setSelectedApp] = useState<AdoptionApplication | null>(null)
  const queryClient = useQueryClient()

  const { data, isLoading } = useQuery({
    queryKey: ['admin', 'applications', statusFilter, page],
    queryFn: () => {
      if (statusFilter === 0) {
        return adminApi.getPendingApplications({ page, page_size: 10 })
      }
      return adminApi.getAllApplications({ 
        page, 
        page_size: 10,
        status: statusFilter >= 0 ? statusFilter : undefined
      })
    },
  })

  const reviewMutation = useMutation({
    mutationFn: ({ id, status }: { id: number; status: number }) =>
      adminApi.reviewApplication(id, { status }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'applications'] })
      queryClient.invalidateQueries({ queryKey: ['admin', 'adoption-statistics'] })
      setSelectedApp(null)
    },
  })

  const applications = data?.data?.list || []
  const total = data?.data?.total || 0
  const totalPages = Math.ceil(total / 10)

  const handleApprove = (app: AdoptionApplication) => {
    if (confirm('确定要通过这个领养申请吗？')) {
      reviewMutation.mutate({ id: app.id, status: 1 })
    }
  }

  const handleReject = (app: AdoptionApplication) => {
    if (confirm('确定要拒绝这个领养申请吗？')) {
      reviewMutation.mutate({ id: app.id, status: 2 })
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
                        {app.status === 0 && (
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
                    <h4 className="font-medium mb-2">领养原因</h4>
                    <p className="text-sm bg-muted p-3 rounded-lg">{selectedApp.reason}</p>
                  </div>

                  {selectedApp.experience && (
                    <div>
                      <h4 className="font-medium mb-2">养宠经验</h4>
                      <p className="text-sm bg-muted p-3 rounded-lg">{selectedApp.experience}</p>
                    </div>
                  )}

                  {selectedApp.living_condition && (
                    <div>
                      <h4 className="font-medium mb-2">居住条件</h4>
                      <p className="text-sm bg-muted p-3 rounded-lg">{selectedApp.living_condition}</p>
                    </div>
                  )}

                  <div className="flex items-center gap-2 text-sm">
                    <span className="text-muted-foreground">家人同意：</span>
                    {selectedApp.family_agreement ? (
                      <span className="text-green-600">是</span>
                    ) : (
                      <span className="text-red-600">否</span>
                    )}
                  </div>

                  {selectedApp.status === 0 && (
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
