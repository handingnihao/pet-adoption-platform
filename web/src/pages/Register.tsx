import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { PawPrint, Eye, EyeOff, CheckCircle, XCircle } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Input } from '../components/ui/input'
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '../components/ui/card'
import { useAuthStore } from '../store/auth'

// 手机号正则：中国大陆手机号
const PHONE_REGEX = /^1[3-9]\d{9}$/

// 密码验证规则
const PASSWORD_RULES = {
  minLength: { test: (pwd: string) => pwd.length >= 8, label: '至少8位字符' },
  hasUppercase: { test: (pwd: string) => /[A-Z]/.test(pwd), label: '包含大写字母' },
  hasLowercase: { test: (pwd: string) => /[a-z]/.test(pwd), label: '包含小写字母' },
  hasNumber: { test: (pwd: string) => /\d/.test(pwd), label: '包含数字' },
  hasSpecial: { test: (pwd: string) => /[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(pwd), label: '包含特殊字符' },
}

// 邮箱正则
const EMAIL_REGEX = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

export function Register() {
  const [formData, setFormData] = useState({
    username: '',
    password: '',
    confirmPassword: '',
    email: '',
    phone: '',
  })
  const [showPassword, setShowPassword] = useState(false)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')
  const [fieldErrors, setFieldErrors] = useState<Record<string, string>>({})

  const { register } = useAuthStore()
  const navigate = useNavigate()

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setFormData({ ...formData, [name]: value })
    // 清除对应字段的错误
    if (fieldErrors[name]) {
      setFieldErrors({ ...fieldErrors, [name]: '' })
    }
  }

  // 验证单个字段
  const validateField = (name: string, value: string): string => {
    switch (name) {
      case 'username':
        if (!value.trim()) return '请输入用户名'
        if (value.length < 2) return '用户名至少2个字符'
        if (value.length > 20) return '用户名不能超过20个字符'
        return ''
      case 'phone':
        if (!value.trim()) return '请输入手机号'
        if (!PHONE_REGEX.test(value)) return '请输入有效的手机号码'
        return ''
      case 'email':
        if (!value.trim()) return '请输入邮箱'
        if (!EMAIL_REGEX.test(value)) return '请输入有效的邮箱地址'
        return ''
      case 'password':
        if (!value) return '请输入密码'
        const failedRules = Object.values(PASSWORD_RULES).filter(rule => !rule.test(value))
        if (failedRules.length > 0) return `密码需要${failedRules.map(r => r.label).join('、')}`
        return ''
      case 'confirmPassword':
        if (!value) return '请确认密码'
        if (value !== formData.password) return '两次输入的密码不一致'
        return ''
      default:
        return ''
    }
  }

  // 字段失去焦点时验证
  const handleBlur = (e: React.FocusEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    const error = validateField(name, value)
    setFieldErrors({ ...fieldErrors, [name]: error })
  }

  // 检查密码强度
  const getPasswordStrength = () => {
    return Object.entries(PASSWORD_RULES).map(([key, rule]) => ({
      key,
      label: rule.label,
      passed: rule.test(formData.password),
    }))
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')

    // 验证所有必填字段
    const errors: Record<string, string> = {}
    errors.username = validateField('username', formData.username)
    errors.phone = validateField('phone', formData.phone)
    errors.email = validateField('email', formData.email)
    errors.password = validateField('password', formData.password)
    errors.confirmPassword = validateField('confirmPassword', formData.confirmPassword)

    // 过滤掉空错误
    const hasErrors = Object.values(errors).some(e => e !== '')
    if (hasErrors) {
      setFieldErrors(errors)
      setError('请检查输入信息')
      return
    }

    setLoading(true)
    try {
      await register({
        username: formData.username,
        password: formData.password,
        email: formData.email || undefined,
        phone: formData.phone || undefined,
      })
      navigate('/login', { state: { message: '注册成功，请登录' } })
    } catch (err) {
      setError(err instanceof Error ? err.message : '注册失败')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-[calc(100vh-4rem)] flex items-center justify-center py-12 px-4">
      <Card className="w-full max-w-md">
        <CardHeader className="text-center">
          <div className="mx-auto mb-4">
            <PawPrint className="h-12 w-12 text-primary" />
          </div>
          <CardTitle className="text-2xl">创建账号</CardTitle>
          <CardDescription>加入我们，开启您的领养之旅</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-4">
            {error && (
              <div className="p-3 text-sm text-destructive bg-destructive/10 rounded-md">
                {error}
              </div>
            )}

            <div className="space-y-2">
              <label className="text-sm font-medium">用户名 *</label>
              <Input
                name="username"
                placeholder="请输入用户名（2-20个字符）"
                value={formData.username}
                onChange={handleChange}
                onBlur={handleBlur}
                className={fieldErrors.username ? 'border-destructive' : ''}
                required
              />
              {fieldErrors.username && <p className="text-xs text-destructive">{fieldErrors.username}</p>}
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">密码 *</label>
              <div className="relative">
                <Input
                  name="password"
                  type={showPassword ? 'text' : 'password'}
                  placeholder="请输入密码"
                  value={formData.password}
                  onChange={handleChange}
                  onBlur={handleBlur}
                  className={fieldErrors.password ? 'border-destructive' : ''}
                  required
                />
                <button
                  type="button"
                  className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground"
                  onClick={() => setShowPassword(!showPassword)}
                >
                  {showPassword ? <EyeOff className="h-4 w-4" /> : <Eye className="h-4 w-4" />}
                </button>
              </div>
              {/* 密码强度提示 */}
              {formData.password && (
                <div className="grid grid-cols-2 gap-1 text-xs mt-2">
                  {getPasswordStrength().map(({ key, label, passed }) => (
                    <div key={key} className={`flex items-center gap-1 ${passed ? 'text-green-600' : 'text-muted-foreground'}`}>
                      {passed ? <CheckCircle className="h-3 w-3" /> : <XCircle className="h-3 w-3" />}
                      {label}
                    </div>
                  ))}
                </div>
              )}
              {fieldErrors.password && <p className="text-xs text-destructive">{fieldErrors.password}</p>}
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">确认密码 *</label>
              <Input
                name="confirmPassword"
                type="password"
                placeholder="请再次输入密码"
                value={formData.confirmPassword}
                onChange={handleChange}
                onBlur={handleBlur}
                className={fieldErrors.confirmPassword ? 'border-destructive' : ''}
                required
              />
              {fieldErrors.confirmPassword && <p className="text-xs text-destructive">{fieldErrors.confirmPassword}</p>}
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">邮箱 *</label>
              <Input
                name="email"
                type="email"
                placeholder="请输入邮箱"
                value={formData.email}
                onChange={handleChange}
                onBlur={handleBlur}
                className={fieldErrors.email ? 'border-destructive' : ''}
                required
              />
              {fieldErrors.email && <p className="text-xs text-destructive">{fieldErrors.email}</p>}
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">手机号 *</label>
              <Input
                name="phone"
                placeholder="请输入手机号"
                value={formData.phone}
                onChange={handleChange}
                onBlur={handleBlur}
                className={fieldErrors.phone ? 'border-destructive' : ''}
                required
              />
              {fieldErrors.phone && <p className="text-xs text-destructive">{fieldErrors.phone}</p>}
            </div>

            <Button type="submit" className="w-full" disabled={loading}>
              {loading ? '注册中...' : '注册'}
            </Button>

            <div className="text-center text-sm text-muted-foreground">
              已有账号？{' '}
              <Link to="/login" className="text-primary hover:underline">
                立即登录
              </Link>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}
