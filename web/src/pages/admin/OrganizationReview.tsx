import { useState } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Check, X, Building2, Clock, CheckCircle, XCircle } from 'lucide-react'
import { Button } from '../../components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '../../components/ui/card'
import { organizationApi } from '../../lib/api'
import { adminApi } from '../../lib/adminApi'
import type { Organization } from '../../lib/api'

const statusConfig = {
  0: { label: '待审核', icon: Clock, color: 'text-yellow-500', bg: 'bg-yellow-50' },
  1: { label: '已通过', icon: CheckCircle, color: 'text-green-500', bg: 'bg-green-50' },
  2: { label: '已拒绝', icon: XCircle, color: 'text-red-500', bg: 'bg-red-50' },
}

export function OrganizationReview() {
  const [page, setPage] = useState(1)
  const [selectedOrg, setSelectedOrg] = useState<Organization | null>(null)
  const queryClient = useQueryClient()

  const { data, isLoading } = useQuery({
    queryKey: ['admin', 'organizations', page],
    queryFn: () => organizationApi.list({ page, page_size: 10 }),
  })

  const reviewMutation = useMutation({
    mutationFn: ({ id, status }: { id: number; status: number }) =>
      adminApi.updateOrganizationStatus(id, { status }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'organizations'] })
      setSelectedOrg(null)
    },
  })

  const organizations = data?.data?.list || []
  const total = data?.data?.total || 0
  const totalPages = Math.ceil(total / 10)

  const handleApprove = (org: Organization) => {
    if (confirm(`确定要通过机构 "${org.name}" 的审核吗？`)) {
      reviewMutation.mutate({ id: org.id, status: 1 })
    }
  }

  const handleReject = (org: Organization) => {
    if (confirm(`确定要拒绝机构 "${org.name}" 的审核吗？`)) {
      reviewMutation.mutate({ id: org.id, status: 2 })
    }
  }

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold">机构审核</h1>
        <p className="text-muted-foreground">审核机构入驻申请</p>
      </div>

      <div className="grid lg:grid-cols-3 gap-6">
        <div className="lg:col-span-2">
          <Card>
            <CardHeader>
              <CardTitle>机构列表 ({total})</CardTitle>
            </CardHeader>
            <CardContent>
              {isLoading ? (
                <div className="space-y-4">
                  {[...Array(5)].map((_, i) => (
                    <div key={i} className="h-20 bg-muted animate-pulse rounded" />
                  ))}
                </div>
              ) : organizations.length === 0 ? (
                <div className="text-center py-8 text-muted-foreground">
                  暂无机构数据
                </div>
              ) : (
                <div className="space-y-4">
                  {organizations.map((org: Organization) => {
                    const status = statusConfig[org.status as keyof typeof statusConfig]
                    const StatusIcon = status?.icon || Clock
                    return (
                      <div
                        key={org.id}
                        className={`flex gap-4 p-4 border rounded-lg cursor-pointer transition-colors ${
                          selectedOrg?.id === org.id ? 'border-primary bg-primary/5' : 'hover:bg-muted/50'
                        }`}
                        onClick={() => setSelectedOrg(org)}
                      >
                        <div className="w-16 h-16 rounded-lg overflow-hidden bg-muted flex-shrink-0 flex items-center justify-center">
                          {org.logo ? (
                            <img src={org.logo} alt={org.name} className="w-full h-full object-cover" />
                          ) : (
                            <Building2 className="h-8 w-8 text-muted-foreground" />
                          )}
                        </div>
                        <div className="flex-1 min-w-0">
                          <div className="flex items-center gap-2">
                            <h3 className="font-medium">{org.name}</h3>
                            <span className={`flex items-center gap-1 px-2 py-0.5 rounded text-xs ${status?.bg} ${status?.color}`}>
                              <StatusIcon className="h-3 w-3" />
                              {status?.label}
                            </span>
                          </div>
                          <p className="text-sm text-muted-foreground">{org.type}</p>
                          <p className="text-xs text-muted-foreground mt-1">
                            {org.city} {org.district}
                          </p>
                        </div>
                        {org.status === 0 && (
                          <div className="flex items-center gap-2">
                            <Button
                              size="sm"
                              variant="outline"
                              className="text-green-600"
                              onClick={(e) => { e.stopPropagation(); handleApprove(org); }}
                            >
                              <Check className="h-4 w-4" />
                            </Button>
                            <Button
                              size="sm"
                              variant="outline"
                              className="text-red-600"
                              onClick={(e) => { e.stopPropagation(); handleReject(org); }}
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

        <div>
          <Card className="sticky top-20">
            <CardHeader>
              <CardTitle>机构详情</CardTitle>
            </CardHeader>
            <CardContent>
              {selectedOrg ? (
                <div className="space-y-4">
                  {selectedOrg.logo && (
                    <img src={selectedOrg.logo} alt={selectedOrg.name} className="w-full h-32 object-cover rounded-lg" />
                  )}
                  <div>
                    <h3 className="font-bold text-lg">{selectedOrg.name}</h3>
                    <p className="text-sm text-muted-foreground">{selectedOrg.type}</p>
                  </div>
                  {selectedOrg.description && (
                    <p className="text-sm">{selectedOrg.description}</p>
                  )}
                  <div className="text-sm space-y-2">
                    <p><span className="text-muted-foreground">联系人：</span>{selectedOrg.contact_name || '-'}</p>
                    <p><span className="text-muted-foreground">电话：</span>{selectedOrg.contact_phone || '-'}</p>
                    <p><span className="text-muted-foreground">邮箱：</span>{selectedOrg.contact_email || '-'}</p>
                    <p><span className="text-muted-foreground">地址：</span>{selectedOrg.province} {selectedOrg.city} {selectedOrg.district} {selectedOrg.address}</p>
                  </div>

                  {selectedOrg.status === 0 && (
                    <div className="flex gap-2 pt-4">
                      <Button className="flex-1" onClick={() => handleApprove(selectedOrg)}>
                        <Check className="h-4 w-4 mr-2" />通过
                      </Button>
                      <Button variant="destructive" className="flex-1" onClick={() => handleReject(selectedOrg)}>
                        <X className="h-4 w-4 mr-2" />拒绝
                      </Button>
                    </div>
                  )}
                </div>
              ) : (
                <p className="text-center text-muted-foreground py-8">点击左侧列表查看详情</p>
              )}
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
