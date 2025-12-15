import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useMutation } from '@tanstack/react-query'
import { ArrowLeft, Image, X } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Input } from '../components/ui/input'
import { Card, CardContent, CardHeader, CardTitle } from '../components/ui/card'
import { communityApi } from '../lib/api'
import { useAuthStore } from '../store/auth'

const postTypes = [
  { value: 'story', label: '领养故事' },
  { value: 'knowledge', label: '养宠知识' },
  { value: 'daily', label: '日常分享' },
  { value: 'other', label: '其他' },
]

export function CreatePost() {
  const navigate = useNavigate()
  const { isAuthenticated } = useAuthStore()

  const [formData, setFormData] = useState({
    title: '',
    content: '',
    type: 'daily',
    images: [] as string[],
  })
  const [imageUrl, setImageUrl] = useState('')

  const mutation = useMutation({
    mutationFn: () => communityApi.createPost({
      title: formData.title || undefined,
      content: formData.content,
      type: formData.type,
      images: formData.images.length > 0 ? formData.images : undefined,
    }),
    onSuccess: () => {
      navigate('/community', { state: { message: '发布成功' } })
    },
  })

  if (!isAuthenticated) {
    navigate('/login')
    return null
  }

  const handleAddImage = () => {
    if (imageUrl.trim() && formData.images.length < 9) {
      setFormData({
        ...formData,
        images: [...formData.images, imageUrl.trim()],
      })
      setImageUrl('')
    }
  }

  const handleRemoveImage = (index: number) => {
    setFormData({
      ...formData,
      images: formData.images.filter((_, i) => i !== index),
    })
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (!formData.content.trim()) return
    mutation.mutate()
  }

  return (
    <div className="container py-8">
      <div className="max-w-2xl mx-auto">
        <Button variant="ghost" className="mb-6" onClick={() => navigate(-1)}>
          <ArrowLeft className="h-4 w-4 mr-2" />
          返回
        </Button>

        <Card>
          <CardHeader>
            <CardTitle>发布动态</CardTitle>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-6">
              {mutation.isError && (
                <div className="p-3 text-sm text-destructive bg-destructive/10 rounded-md">
                  {mutation.error instanceof Error ? mutation.error.message : '发布失败'}
                </div>
              )}

              {/* 类型选择 */}
              <div className="space-y-2">
                <label className="text-sm font-medium">动态类型</label>
                <div className="flex gap-2 flex-wrap">
                  {postTypes.map((type) => (
                    <Button
                      key={type.value}
                      type="button"
                      variant={formData.type === type.value ? 'default' : 'outline'}
                      size="sm"
                      onClick={() => setFormData({ ...formData, type: type.value })}
                    >
                      {type.label}
                    </Button>
                  ))}
                </div>
              </div>

              {/* 标题 */}
              <div className="space-y-2">
                <label className="text-sm font-medium">标题（可选）</label>
                <Input
                  placeholder="给你的动态起个标题吧"
                  value={formData.title}
                  onChange={(e) => setFormData({ ...formData, title: e.target.value })}
                />
              </div>

              {/* 内容 */}
              <div className="space-y-2">
                <label className="text-sm font-medium">
                  内容 <span className="text-destructive">*</span>
                </label>
                <textarea
                  className="w-full min-h-[200px] px-3 py-2 border rounded-md text-sm resize-none focus:outline-none focus:ring-2 focus:ring-primary"
                  placeholder="分享你和毛孩子的故事..."
                  value={formData.content}
                  onChange={(e) => setFormData({ ...formData, content: e.target.value })}
                  required
                />
                <p className="text-xs text-muted-foreground text-right">
                  {formData.content.length} 字
                </p>
              </div>

              {/* 图片 */}
              <div className="space-y-2">
                <label className="text-sm font-medium">图片（最多9张）</label>
                
                {/* 已添加的图片 */}
                {formData.images.length > 0 && (
                  <div className="flex gap-2 flex-wrap mb-2">
                    {formData.images.map((img, idx) => (
                      <div key={idx} className="relative w-20 h-20">
                        <img
                          src={img}
                          alt=""
                          className="w-full h-full object-cover rounded-lg"
                        />
                        <button
                          type="button"
                          className="absolute -top-2 -right-2 p-1 bg-destructive text-white rounded-full"
                          onClick={() => handleRemoveImage(idx)}
                        >
                          <X className="w-3 h-3" />
                        </button>
                      </div>
                    ))}
                  </div>
                )}

                {/* 添加图片URL */}
                {formData.images.length < 9 && (
                  <div className="flex gap-2">
                    <Input
                      placeholder="输入图片URL"
                      value={imageUrl}
                      onChange={(e) => setImageUrl(e.target.value)}
                    />
                    <Button type="button" variant="outline" onClick={handleAddImage}>
                      <Image className="w-4 h-4 mr-1" />
                      添加
                    </Button>
                  </div>
                )}
                <p className="text-xs text-muted-foreground">
                  提示：请输入图片的网络地址
                </p>
              </div>

              {/* 提交按钮 */}
              <div className="flex gap-4">
                <Button
                  type="submit"
                  className="flex-1"
                  disabled={mutation.isPending || !formData.content.trim()}
                >
                  {mutation.isPending ? '发布中...' : '发布'}
                </Button>
                <Button type="button" variant="outline" onClick={() => navigate(-1)}>
                  取消
                </Button>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
