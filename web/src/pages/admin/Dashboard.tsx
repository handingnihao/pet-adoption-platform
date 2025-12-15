import { useQuery } from '@tanstack/react-query'
import { PawPrint, Heart, TrendingUp } from 'lucide-react'
import { Card, CardContent, CardHeader, CardTitle } from '../../components/ui/card'
import { adminApi } from '../../lib/adminApi'

export function AdminDashboard() {
  const { data: petStats } = useQuery({
    queryKey: ['admin', 'pet-statistics'],
    queryFn: () => adminApi.getPetStatistics(),
  })

  const { data: adoptionStats } = useQuery({
    queryKey: ['admin', 'adoption-statistics'],
    queryFn: () => adminApi.getAdoptionStatistics(),
  })

  const stats = [
    {
      title: '待审核宠物',
      value: petStats?.data?.pending || 0,
      icon: PawPrint,
      color: 'text-yellow-500',
      bg: 'bg-yellow-50',
    },
    {
      title: '已发布宠物',
      value: petStats?.data?.approved || 0,
      icon: PawPrint,
      color: 'text-green-500',
      bg: 'bg-green-50',
    },
    {
      title: '待审核申请',
      value: adoptionStats?.data?.pending_applications || 0,
      icon: Heart,
      color: 'text-blue-500',
      bg: 'bg-blue-50',
    },
    {
      title: '成功领养',
      value: adoptionStats?.data?.total_adoptions || 0,
      icon: TrendingUp,
      color: 'text-purple-500',
      bg: 'bg-purple-50',
    },
  ]

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold">仪表盘</h1>
        <p className="text-muted-foreground">欢迎使用宠物领养平台管理后台</p>
      </div>

      {/* 统计卡片 */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        {stats.map((stat) => (
          <Card key={stat.title}>
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm text-muted-foreground">{stat.title}</p>
                  <p className="text-3xl font-bold mt-1">{stat.value}</p>
                </div>
                <div className={`p-3 rounded-full ${stat.bg}`}>
                  <stat.icon className={`h-6 w-6 ${stat.color}`} />
                </div>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>

      {/* 宠物类型分布 */}
      <div className="grid md:grid-cols-2 gap-6">
        <Card>
          <CardHeader>
            <CardTitle>宠物类型分布</CardTitle>
          </CardHeader>
          <CardContent>
            {petStats?.data?.by_type ? (
              <div className="space-y-3">
                {Object.entries(petStats.data.by_type).map(([type, count]) => (
                  <div key={type} className="flex items-center justify-between">
                    <span className="capitalize">{type === 'dog' ? '狗狗' : type === 'cat' ? '猫咪' : type}</span>
                    <div className="flex items-center gap-2">
                      <div className="w-32 h-2 bg-muted rounded-full overflow-hidden">
                        <div 
                          className="h-full bg-primary rounded-full"
                          style={{ width: `${(count / (petStats.data?.total || 1)) * 100}%` }}
                        />
                      </div>
                      <span className="text-sm text-muted-foreground w-8">{count}</span>
                    </div>
                  </div>
                ))}
              </div>
            ) : (
              <p className="text-muted-foreground">暂无数据</p>
            )}
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>领养申请概览</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-4">
              <div className="flex justify-between items-center">
                <span>总申请数</span>
                <span className="font-bold">{adoptionStats?.data?.total_applications || 0}</span>
              </div>
              <div className="flex justify-between items-center">
                <span>待审核</span>
                <span className="text-yellow-500 font-bold">{adoptionStats?.data?.pending_applications || 0}</span>
              </div>
              <div className="flex justify-between items-center">
                <span>已通过</span>
                <span className="text-green-500 font-bold">{adoptionStats?.data?.approved_applications || 0}</span>
              </div>
              <div className="flex justify-between items-center">
                <span>已拒绝</span>
                <span className="text-red-500 font-bold">{adoptionStats?.data?.rejected_applications || 0}</span>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
