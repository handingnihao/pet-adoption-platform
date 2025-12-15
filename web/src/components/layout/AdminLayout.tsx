import { Outlet, Link, useLocation, useNavigate } from 'react-router-dom'
import { 
  LayoutDashboard, Users, PawPrint, Heart, Building2, 
  ChevronLeft, LogOut, Menu, X
} from 'lucide-react'
import { useState } from 'react'
import { Button } from '../ui/button'
import { useAuthStore } from '../../store/auth'

const menuItems = [
  { path: '/admin', icon: LayoutDashboard, label: '仪表盘' },
  { path: '/admin/users', icon: Users, label: '用户管理' },
  { path: '/admin/pets', icon: PawPrint, label: '宠物管理' },
  { path: '/admin/adoptions', icon: Heart, label: '领养管理' },
  { path: '/admin/organizations', icon: Building2, label: '机构审核' },
]

export function AdminLayout() {
  const location = useLocation()
  const navigate = useNavigate()
  const { user, logout } = useAuthStore()
  const [sidebarOpen, setSidebarOpen] = useState(true)

  // 检查是否是管理员
  if (user?.role !== 'admin') {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4">无权限访问</h1>
          <p className="text-muted-foreground mb-4">您没有管理员权限</p>
          <Button onClick={() => navigate('/')}>返回首页</Button>
        </div>
      </div>
    )
  }

  const handleLogout = () => {
    logout()
    navigate('/login')
  }

  return (
    <div className="min-h-screen bg-muted/30">
      {/* 侧边栏 */}
      <aside className={`fixed top-0 left-0 z-40 h-screen transition-transform ${
        sidebarOpen ? 'translate-x-0' : '-translate-x-full'
      } md:translate-x-0`}>
        <div className="h-full w-64 bg-background border-r flex flex-col">
          {/* Logo */}
          <div className="h-16 flex items-center justify-between px-4 border-b">
            <Link to="/admin" className="flex items-center gap-2">
              <PawPrint className="h-6 w-6 text-primary" />
              <span className="font-bold">管理后台</span>
            </Link>
            <button className="md:hidden" onClick={() => setSidebarOpen(false)}>
              <X className="h-5 w-5" />
            </button>
          </div>

          {/* 菜单 */}
          <nav className="flex-1 p-4 space-y-1">
            {menuItems.map((item) => {
              const isActive = location.pathname === item.path || 
                (item.path !== '/admin' && location.pathname.startsWith(item.path))
              return (
                <Link
                  key={item.path}
                  to={item.path}
                  className={`flex items-center gap-3 px-3 py-2 rounded-lg transition-colors ${
                    isActive 
                      ? 'bg-primary text-primary-foreground' 
                      : 'hover:bg-muted'
                  }`}
                >
                  <item.icon className="h-5 w-5" />
                  <span>{item.label}</span>
                </Link>
              )
            })}
          </nav>

          {/* 底部 */}
          <div className="p-4 border-t space-y-2">
            <Link
              to="/"
              className="flex items-center gap-3 px-3 py-2 rounded-lg hover:bg-muted transition-colors"
            >
              <ChevronLeft className="h-5 w-5" />
              <span>返回前台</span>
            </Link>
            <button
              onClick={handleLogout}
              className="flex items-center gap-3 px-3 py-2 rounded-lg hover:bg-muted transition-colors w-full text-left text-destructive"
            >
              <LogOut className="h-5 w-5" />
              <span>退出登录</span>
            </button>
          </div>
        </div>
      </aside>

      {/* 主内容区 */}
      <div className={`${sidebarOpen ? 'md:ml-64' : ''}`}>
        {/* 顶部栏 */}
        <header className="h-16 bg-background border-b flex items-center justify-between px-4 sticky top-0 z-30">
          <button className="md:hidden" onClick={() => setSidebarOpen(true)}>
            <Menu className="h-5 w-5" />
          </button>
          <div className="flex items-center gap-4 ml-auto">
            <span className="text-sm text-muted-foreground">
              管理员: {user?.nickname || user?.username}
            </span>
          </div>
        </header>

        {/* 页面内容 */}
        <main className="p-6">
          <Outlet />
        </main>
      </div>
    </div>
  )
}
