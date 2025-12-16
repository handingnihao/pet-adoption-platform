import { useQuery } from '@tanstack/react-query'
import { Link } from 'react-router-dom'
import { Building2, Plus, MapPin, Phone, Mail, Clock } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Card, CardContent } from '../components/ui/card'
import { organizationApi } from '../lib/api'
import type { Organization } from '../lib/api'

const statusColors: Record<number, string> = {
  0: 'bg-yellow-100 text-yellow-800',
  1: 'bg-green-100 text-green-800',
  2: 'bg-red-100 text-red-800',
}

const statusLabels: Record<number, string> = {
  0: '待审核',
  1: '已认证',
  2: '已拒绝',
}

const typeLabels: Record<string, string> = {
  shelter: '动物收容所',
  rescue: '救助站',
  hospital: '宠物医院',
  association: '动物保护协会',
  other: '其他',
}

export function MyOrganizations() {
  const { data, isLoading } = useQuery({
    queryKey: ['my-organizations'],
    queryFn: () => organizationApi.getMyOrganizations(),
  })

  const organizations = data?.data || []

  return (
    <div className="container py-8">
      <div className="flex items-center justify-between mb-6">
        <div>
          <h1 className="text-2xl font-bold">我的机构</h1>
          <p className="text-muted-foreground">管理您创建或入驻的机构</p>
        </div>
        <Button asChild>
          <Link to="/organizations/apply">
            <Plus className="w-4 h-4 mr-2" />
            申请入驻
          </Link>
        </Button>
      </div>

      {isLoading ? (
        <div className="text-center py-12 text-muted-foreground">加载中...</div>
      ) : organizations.length === 0 ? (
        <Card>
          <CardContent className="py-12 text-center">
            <Building2 className="w-12 h-12 text-muted-foreground mx-auto mb-4" />
            <p className="text-muted-foreground mb-4">您还没有入驻任何机构</p>
            <Button asChild>
              <Link to="/organizations/apply">申请入驻</Link>
            </Button>
          </CardContent>
        </Card>
      ) : (
        <div className="grid md:grid-cols-2 gap-6">
          {organizations.map((org: Organization) => (
            <Card key={org.id} className="overflow-hidden">
              <CardContent className="p-6">
                <div className="flex items-start gap-4">
                  <div className="w-16 h-16 rounded-lg bg-primary/10 flex items-center justify-center flex-shrink-0">
                    {org.logo ? (
                      <img src={org.logo} alt={org.name} className="w-16 h-16 rounded-lg object-cover" />
                    ) : (
                      <Building2 className="w-8 h-8 text-primary" />
                    )}
                  </div>
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center gap-2 mb-1">
                      <h3 className="font-bold text-lg truncate">{org.name}</h3>
                      <span className={`text-xs px-2 py-0.5 rounded-full ${statusColors[org.status]}`}>
                        {statusLabels[org.status]}
                      </span>
                    </div>
                    <p className="text-sm text-muted-foreground mb-2">
                      {typeLabels[org.type] || org.type}
                    </p>
                    {org.description && (
                      <p className="text-sm text-muted-foreground line-clamp-2 mb-3">
                        {org.description}
                      </p>
                    )}
                    <div className="space-y-1 text-xs text-muted-foreground">
                      {(org.province || org.city) && (
                        <div className="flex items-center gap-1">
                          <MapPin className="w-3 h-3" />
                          <span>{org.province} {org.city} {org.address}</span>
                        </div>
                      )}
                      {org.contact_phone && (
                        <div className="flex items-center gap-1">
                          <Phone className="w-3 h-3" />
                          <span>{org.contact_phone}</span>
                        </div>
                      )}
                      {org.contact_email && (
                        <div className="flex items-center gap-1">
                          <Mail className="w-3 h-3" />
                          <span>{org.contact_email}</span>
                        </div>
                      )}
                      <div className="flex items-center gap-1">
                        <Clock className="w-3 h-3" />
                        <span>申请于 {new Date(org.created_at).toLocaleDateString('zh-CN')}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      )}
    </div>
  )
}
