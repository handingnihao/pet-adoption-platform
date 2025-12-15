import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { User, Mail, Phone, Camera, Save } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Input } from '../components/ui/input'
import { Card, CardContent, CardHeader, CardTitle } from '../components/ui/card'
import { useAuthStore } from '../store/auth'
import { userApi } from '../lib/api'

export function Profile() {
  const { user, isAuthenticated, refreshUser } = useAuthStore()
  const navigate = useNavigate()
  
  const [formData, setFormData] = useState({
    nickname: '',
    email: '',
    phone: '',
    avatar: '',
    real_name: '',
  })
  const [loading, setLoading] = useState(false)
  const [message, setMessage] = useState({ type: '', text: '' })

  useEffect(() => {
    if (!isAuthenticated) {
      navigate('/login')
      return
    }
    if (user) {
      setFormData({
        nickname: user.nickname || '',
        email: user.email || '',
        phone: user.phone || '',
        avatar: user.avatar || '',
        real_name: '',
      })
    }
  }, [user, isAuthenticated, navigate])

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value })
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setMessage({ type: '', text: '' })

    try {
      await userApi.updateProfile(formData)
      await refreshUser()
      setMessage({ type: 'success', text: '保存成功' })
    } catch (err) {
      setMessage({ type: 'error', text: err instanceof Error ? err.message : '保存失败' })
    } finally {
      setLoading(false)
    }
  }

  if (!user) {
    return null
  }

  return (
    <div className="container py-8">
      <div className="max-w-2xl mx-auto">
        <h1 className="text-2xl font-bold mb-6">个人中心</h1>

        <div className="grid gap-6">
          {/* 头像卡片 */}
          <Card>
            <CardHeader>
              <CardTitle className="text-lg">头像</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="flex items-center gap-6">
                <div className="relative">
                  <div className="w-24 h-24 rounded-full bg-primary/10 flex items-center justify-center overflow-hidden">
                    {user.avatar ? (
                      <img src={user.avatar} alt="头像" className="w-full h-full object-cover" />
                    ) : (
                      <User className="w-12 h-12 text-primary" />
                    )}
                  </div>
                  <button className="absolute bottom-0 right-0 p-2 bg-primary text-white rounded-full hover:bg-primary/90">
                    <Camera className="w-4 h-4" />
                  </button>
                </div>
                <div>
                  <p className="font-medium">{user.username}</p>
                  <p className="text-sm text-muted-foreground">
                    {user.role === 'admin' ? '管理员' : '普通用户'}
                  </p>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* 个人信息表单 */}
          <Card>
            <CardHeader>
              <CardTitle className="text-lg">基本信息</CardTitle>
            </CardHeader>
            <CardContent>
              <form onSubmit={handleSubmit} className="space-y-4">
                {message.text && (
                  <div className={`p-3 text-sm rounded-md ${
                    message.type === 'success' 
                      ? 'bg-green-100 text-green-700' 
                      : 'bg-destructive/10 text-destructive'
                  }`}>
                    {message.text}
                  </div>
                )}

                <div className="grid md:grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">用户名</label>
                    <Input value={user.username} disabled />
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-medium">昵称</label>
                    <Input
                      name="nickname"
                      placeholder="请输入昵称"
                      value={formData.nickname}
                      onChange={handleChange}
                    />
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-medium flex items-center gap-2">
                      <Mail className="w-4 h-4" /> 邮箱
                    </label>
                    <Input
                      name="email"
                      type="email"
                      placeholder="请输入邮箱"
                      value={formData.email}
                      onChange={handleChange}
                    />
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-medium flex items-center gap-2">
                      <Phone className="w-4 h-4" /> 手机号
                    </label>
                    <Input
                      name="phone"
                      placeholder="请输入手机号"
                      value={formData.phone}
                      onChange={handleChange}
                    />
                  </div>
                </div>

                <div className="flex justify-end">
                  <Button type="submit" disabled={loading}>
                    <Save className="w-4 h-4 mr-2" />
                    {loading ? '保存中...' : '保存修改'}
                  </Button>
                </div>
              </form>
            </CardContent>
          </Card>

          {/* 修改密码 */}
          <Card>
            <CardHeader>
              <CardTitle className="text-lg">安全设置</CardTitle>
            </CardHeader>
            <CardContent>
              <Button variant="outline" onClick={() => navigate('/change-password')}>
                修改密码
              </Button>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
