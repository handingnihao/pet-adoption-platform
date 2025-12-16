import { useQuery } from '@tanstack/react-query'
import { Link } from 'react-router-dom'
import { Heart, Gift, HandHelping, Plus } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Card, CardContent } from '../components/ui/card'
import { donationApi } from '../lib/api'
import type { Donation } from '../lib/api'

const statusColors: Record<number, string> = {
  0: 'bg-yellow-100 text-yellow-800',
  1: 'bg-blue-100 text-blue-800',
  2: 'bg-green-100 text-green-800',
  3: 'bg-gray-100 text-gray-600',
}

const typeIcons = {
  money: Heart,
  supply: Gift,
  service: HandHelping,
}

const typeLabels = {
  money: '资金捐赠',
  supply: '物资捐赠',
  service: '服务捐赠',
}

export function MyDonations() {
  const { data, isLoading } = useQuery({
    queryKey: ['my-donations'],
    queryFn: () => donationApi.getMyDonations({ page: 1, page_size: 50 }),
  })

  const donations = data?.data?.list || []

  return (
    <div className="container py-8">
      <div className="flex items-center justify-between mb-6">
        <div>
          <h1 className="text-2xl font-bold">我的捐赠</h1>
          <p className="text-muted-foreground">查看您的爱心捐赠记录</p>
        </div>
        <Button asChild>
          <Link to="/donate">
            <Plus className="w-4 h-4 mr-2" />
            发起捐赠
          </Link>
        </Button>
      </div>

      {isLoading ? (
        <div className="text-center py-12 text-muted-foreground">加载中...</div>
      ) : donations.length === 0 ? (
        <Card>
          <CardContent className="py-12 text-center">
            <Heart className="w-12 h-12 text-muted-foreground mx-auto mb-4" />
            <p className="text-muted-foreground mb-4">您还没有捐赠记录</p>
            <Button asChild>
              <Link to="/donate">发起捐赠</Link>
            </Button>
          </CardContent>
        </Card>
      ) : (
        <div className="space-y-4">
          {donations.map((donation: Donation) => {
            const Icon = typeIcons[donation.type]
            return (
              <Card key={donation.id}>
                <CardContent className="p-4">
                  <div className="flex items-start gap-4">
                    <div className="w-12 h-12 rounded-full bg-primary/10 flex items-center justify-center flex-shrink-0">
                      <Icon className="w-6 h-6 text-primary" />
                    </div>
                    <div className="flex-1 min-w-0">
                      <div className="flex items-center gap-2 mb-1">
                        <span className="font-medium">{typeLabels[donation.type]}</span>
                        <span className={`text-xs px-2 py-0.5 rounded-full ${statusColors[donation.status]}`}>
                          {donation.status_text}
                        </span>
                      </div>
                      
                      {donation.type === 'money' && (
                        <p className="text-2xl font-bold text-primary">¥{donation.amount}</p>
                      )}
                      
                      {donation.type === 'supply' && donation.supply_items.length > 0 && (
                        <div className="flex flex-wrap gap-1 mt-1">
                          {donation.supply_items.map((item, i) => (
                            <span key={i} className="text-xs px-2 py-0.5 bg-muted rounded">
                              {item}
                            </span>
                          ))}
                        </div>
                      )}
                      
                      {donation.type === 'service' && donation.service_desc && (
                        <p className="text-sm text-muted-foreground line-clamp-2">
                          {donation.service_desc}
                        </p>
                      )}
                      
                      {donation.message && (
                        <p className="text-sm text-muted-foreground mt-2 italic">
                          "{donation.message}"
                        </p>
                      )}
                      
                      {donation.organization_name && (
                        <p className="text-xs text-muted-foreground mt-2">
                          捐赠给：{donation.organization_name}
                        </p>
                      )}
                      
                      <p className="text-xs text-muted-foreground mt-2">
                        {new Date(donation.created_at).toLocaleString('zh-CN')}
                      </p>
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
