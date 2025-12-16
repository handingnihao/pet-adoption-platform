import { useState } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Search, UserCheck, UserX } from 'lucide-react'
import { Button } from '../../components/ui/button'
import { Input } from '../../components/ui/input'
import { Card, CardContent, CardHeader, CardTitle } from '../../components/ui/card'
import { adminApi } from '../../lib/adminApi'
import type { User } from '../../lib/api'

export function UserManagement() {
  const [keyword, setKeyword] = useState('')
  const [searchKeyword, setSearchKeyword] = useState('')
  const [page, setPage] = useState(1)
  const queryClient = useQueryClient()

  const { data, isLoading } = useQuery({
    queryKey: ['admin', 'users', searchKeyword, page],
    queryFn: () => {
      if (searchKeyword) {
        return adminApi.searchUsers({ keyword: searchKeyword, page, page_size: 10 })
      }
      return adminApi.getUsers({ page, page_size: 10 })
    },
  })

  const disableMutation = useMutation({
    mutationFn: (id: number) => adminApi.disableUser(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'users'] })
    },
  })

  const enableMutation = useMutation({
    mutationFn: (id: number) => adminApi.enableUser(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'users'] })
    },
  })

  const users = data?.data?.list || []
  const total = data?.data?.pagination?.total || data?.data?.total || 0
  const totalPages = Math.ceil(total / 10)

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault()
    setSearchKeyword(keyword)
    setPage(1)
  }

  const handleToggleStatus = (user: User) => {
    if (user.status === 1) {
      if (confirm(`确定要禁用用户 "${user.username}" 吗？`)) {
        disableMutation.mutate(user.id)
      }
    } else {
      enableMutation.mutate(user.id)
    }
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold">用户管理</h1>
          <p className="text-muted-foreground">管理平台用户</p>
        </div>
      </div>

      {/* 搜索 */}
      <Card>
        <CardContent className="p-4">
          <form onSubmit={handleSearch} className="flex gap-4">
            <div className="relative flex-1 max-w-md">
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
              <Input
                placeholder="搜索用户名、邮箱、手机号..."
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

      {/* 用户列表 */}
      <Card>
        <CardHeader>
          <CardTitle>用户列表 ({total})</CardTitle>
        </CardHeader>
        <CardContent>
          {isLoading ? (
            <div className="space-y-4">
              {[...Array(5)].map((_, i) => (
                <div key={i} className="h-16 bg-muted animate-pulse rounded" />
              ))}
            </div>
          ) : users.length === 0 ? (
            <div className="text-center py-8 text-muted-foreground">
              暂无用户数据
            </div>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-full">
                <thead>
                  <tr className="border-b">
                    <th className="text-left py-3 px-4">ID</th>
                    <th className="text-left py-3 px-4">用户名</th>
                    <th className="text-left py-3 px-4">昵称</th>
                    <th className="text-left py-3 px-4">邮箱</th>
                    <th className="text-left py-3 px-4">角色</th>
                    <th className="text-left py-3 px-4">状态</th>
                    <th className="text-left py-3 px-4">注册时间</th>
                    <th className="text-left py-3 px-4">操作</th>
                  </tr>
                </thead>
                <tbody>
                  {users.map((user: User) => (
                    <tr key={user.id} className="border-b hover:bg-muted/50">
                      <td className="py-3 px-4">{user.id}</td>
                      <td className="py-3 px-4 font-medium">{user.username}</td>
                      <td className="py-3 px-4">{user.nickname || '-'}</td>
                      <td className="py-3 px-4">{user.email || '-'}</td>
                      <td className="py-3 px-4">
                        <span className={`px-2 py-1 rounded text-xs ${
                          user.role === 'admin' ? 'bg-purple-100 text-purple-700' : 'bg-gray-100 text-gray-700'
                        }`}>
                          {user.role === 'admin' ? '管理员' : '普通用户'}
                        </span>
                      </td>
                      <td className="py-3 px-4">
                        <span className={`px-2 py-1 rounded text-xs ${
                          user.status === 1 ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'
                        }`}>
                          {user.status === 1 ? '正常' : '禁用'}
                        </span>
                      </td>
                      <td className="py-3 px-4 text-sm text-muted-foreground">
                        {user.created_at ? new Date(user.created_at).toLocaleString('zh-CN') : '-'}
                      </td>
                      <td className="py-3 px-4">
                        {user.role !== 'admin' && (
                          <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => handleToggleStatus(user)}
                            disabled={disableMutation.isPending || enableMutation.isPending}
                          >
                            {user.status === 1 ? (
                              <>
                                <UserX className="h-4 w-4 mr-1 text-red-500" />
                                禁用
                              </>
                            ) : (
                              <>
                                <UserCheck className="h-4 w-4 mr-1 text-green-500" />
                                启用
                              </>
                            )}
                          </Button>
                        )}
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}

          {/* 分页 */}
          {totalPages > 1 && (
            <div className="flex justify-center gap-2 mt-4">
              <Button
                variant="outline"
                size="sm"
                disabled={page <= 1}
                onClick={() => setPage(page - 1)}
              >
                上一页
              </Button>
              <span className="flex items-center px-4 text-sm">
                {page} / {totalPages}
              </span>
              <Button
                variant="outline"
                size="sm"
                disabled={page >= totalPages}
                onClick={() => setPage(page + 1)}
              >
                下一页
              </Button>
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  )
}
