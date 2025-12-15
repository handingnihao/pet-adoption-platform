import { Link, useNavigate } from 'react-router-dom'
import { PawPrint, Menu, X, User, LogOut, Settings, Heart, LayoutDashboard } from 'lucide-react'
import { useState } from 'react'
import { Button } from '../ui/button'
import { useAuthStore } from '../../store/auth'

export function Header() {
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false)
  const { isAuthenticated, user, logout } = useAuthStore()
  const navigate = useNavigate()

  const handleLogout = () => {
    logout()
    navigate('/')
  }

  const navItems = [
    { name: '首页', href: '/' },
    { name: '找宠物', href: '/pets' },
    { name: '领养指南', href: '/guide' },
    { name: '社区', href: '/community' },
    { name: '关于我们', href: '/about' },
  ]

  return (
    <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <nav className="container flex h-16 items-center justify-between">
        {/* Logo */}
        <Link to="/" className="flex items-center space-x-2">
          <PawPrint className="h-8 w-8 text-primary" />
          <span className="text-xl font-bold">宠爱有家</span>
        </Link>

        {/* Desktop Navigation */}
        <div className="hidden md:flex items-center space-x-6">
          {navItems.map((item) => (
            <Link
              key={item.name}
              to={item.href}
              className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
            >
              {item.name}
            </Link>
          ))}
        </div>

        {/* User Actions */}
        <div className="hidden md:flex items-center space-x-4">
          {isAuthenticated ? (
            <div className="flex items-center space-x-4">
              <Link to="/favorites" className="text-muted-foreground hover:text-primary">
                <Heart className="h-5 w-5" />
              </Link>
              <div className="relative group">
                <button className="flex items-center space-x-2 text-sm">
                  <div className="h-8 w-8 rounded-full bg-primary/10 flex items-center justify-center">
                    <User className="h-4 w-4 text-primary" />
                  </div>
                  <span>{user?.nickname || user?.username}</span>
                </button>
                {/* Dropdown */}
                <div className="absolute right-0 mt-2 w-48 py-2 bg-background border rounded-lg shadow-lg opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all">
                  <Link
                    to="/profile"
                    className="flex items-center px-4 py-2 text-sm hover:bg-muted"
                  >
                    <Settings className="h-4 w-4 mr-2" />
                    个人中心
                  </Link>
                  <Link
                    to="/my-pets"
                    className="flex items-center px-4 py-2 text-sm hover:bg-muted"
                  >
                    <PawPrint className="h-4 w-4 mr-2" />
                    我的宠物
                  </Link>
                  <Link
                    to="/my-applications"
                    className="flex items-center px-4 py-2 text-sm hover:bg-muted"
                  >
                    <Heart className="h-4 w-4 mr-2" />
                    我的申请
                  </Link>
                  {user?.role === 'admin' && (
                    <>
                      <hr className="my-2" />
                      <Link
                        to="/admin"
                        className="flex items-center px-4 py-2 text-sm hover:bg-muted text-primary"
                      >
                        <LayoutDashboard className="h-4 w-4 mr-2" />
                        管理后台
                      </Link>
                    </>
                  )}
                  <hr className="my-2" />
                  <button
                    onClick={handleLogout}
                    className="flex items-center w-full px-4 py-2 text-sm text-destructive hover:bg-muted"
                  >
                    <LogOut className="h-4 w-4 mr-2" />
                    退出登录
                  </button>
                </div>
              </div>
            </div>
          ) : (
            <div className="flex items-center space-x-2">
              <Button variant="ghost" asChild>
                <Link to="/login">登录</Link>
              </Button>
              <Button asChild>
                <Link to="/register">注册</Link>
              </Button>
            </div>
          )}
        </div>

        {/* Mobile Menu Button */}
        <button
          className="md:hidden"
          onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
        >
          {mobileMenuOpen ? (
            <X className="h-6 w-6" />
          ) : (
            <Menu className="h-6 w-6" />
          )}
        </button>
      </nav>

      {/* Mobile Menu */}
      {mobileMenuOpen && (
        <div className="md:hidden border-t">
          <div className="container py-4 space-y-4">
            {navItems.map((item) => (
              <Link
                key={item.name}
                to={item.href}
                className="block text-sm font-medium"
                onClick={() => setMobileMenuOpen(false)}
              >
                {item.name}
              </Link>
            ))}
            <hr />
            {isAuthenticated ? (
              <div className="space-y-2">
                <Link to="/profile" className="block text-sm" onClick={() => setMobileMenuOpen(false)}>
                  个人中心
                </Link>
                <button onClick={handleLogout} className="text-sm text-destructive">
                  退出登录
                </button>
              </div>
            ) : (
              <div className="flex space-x-2">
                <Button variant="outline" asChild className="flex-1">
                  <Link to="/login">登录</Link>
                </Button>
                <Button asChild className="flex-1">
                  <Link to="/register">注册</Link>
                </Button>
              </div>
            )}
          </div>
        </div>
      )}
    </header>
  )
}
