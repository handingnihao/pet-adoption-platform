import { useParams, Link } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { ArrowLeft, Heart, MessageCircle, Eye, Calendar, User } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Card, CardContent } from '../components/ui/card'
import { communityApi } from '../lib/api'

export function PostDetail() {
  const { id } = useParams<{ id: string }>()

  const { data, isLoading, error } = useQuery({
    queryKey: ['post', id],
    queryFn: () => communityApi.getPost(Number(id)),
    enabled: !!id,
  })

  if (isLoading) {
    return (
      <div className="container py-8">
        <div className="max-w-3xl mx-auto">
          <div className="animate-pulse space-y-4">
            <div className="h-8 bg-muted rounded w-3/4"></div>
            <div className="h-4 bg-muted rounded w-1/4"></div>
            <div className="h-64 bg-muted rounded"></div>
          </div>
        </div>
      </div>
    )
  }

  if (error || !data?.data) {
    return (
      <div className="container py-8">
        <div className="max-w-3xl mx-auto text-center">
          <h1 className="text-2xl font-bold mb-4">帖子不存在</h1>
          <p className="text-muted-foreground mb-6">该帖子可能已被删除或不存在</p>
          <Link to="/community">
            <Button>返回社区</Button>
          </Link>
        </div>
      </div>
    )
  }

  const post = data.data

  const typeLabels: Record<string, string> = {
    story: '领养故事',
    knowledge: '养宠知识',
    daily: '日常分享',
    other: '其他',
  }

  const typeColors: Record<string, string> = {
    story: 'bg-pink-100 text-pink-700',
    knowledge: 'bg-blue-100 text-blue-700',
    daily: 'bg-green-100 text-green-700',
    other: 'bg-gray-100 text-gray-700',
  }

  return (
    <div className="container py-8">
      <div className="max-w-3xl mx-auto">
        {/* 返回按钮 */}
        <Link to="/community" className="inline-flex items-center text-muted-foreground hover:text-foreground mb-6">
          <ArrowLeft className="h-4 w-4 mr-2" />
          返回社区
        </Link>

        <Card>
          <CardContent className="p-6 md:p-8">
            {/* 标签 */}
            <span className={`inline-block px-3 py-1 rounded-full text-sm font-medium mb-4 ${typeColors[post.type] || typeColors.other}`}>
              {typeLabels[post.type] || '其他'}
            </span>

            {/* 标题 */}
            <h1 className="text-2xl md:text-3xl font-bold mb-4">
              {post.title || '无标题'}
            </h1>

            {/* 作者信息 */}
            <div className="flex items-center gap-4 text-sm text-muted-foreground mb-6 pb-6 border-b">
              <div className="flex items-center gap-2">
                <div className="w-10 h-10 rounded-full bg-primary/10 flex items-center justify-center">
                  {post.user_avatar ? (
                    <img src={post.user_avatar} alt="" className="w-10 h-10 rounded-full object-cover" />
                  ) : (
                    <User className="h-5 w-5 text-primary" />
                  )}
                </div>
                <span className="font-medium text-foreground">{post.username || '匿名用户'}</span>
              </div>
              <div className="flex items-center gap-1">
                <Calendar className="h-4 w-4" />
                {new Date(post.created_at).toLocaleDateString('zh-CN', {
                  year: 'numeric',
                  month: 'long',
                  day: 'numeric',
                })}
              </div>
            </div>

            {/* 图片 */}
            {post.images && post.images.length > 0 && (
              <div className="mb-6">
                <div className={`grid gap-4 ${post.images.length === 1 ? 'grid-cols-1' : 'grid-cols-2'}`}>
                  {post.images.map((img: string, index: number) => (
                    <img
                      key={index}
                      src={img}
                      alt={`图片${index + 1}`}
                      className="w-full rounded-lg object-cover max-h-96"
                    />
                  ))}
                </div>
              </div>
            )}

            {/* 内容 */}
            <div className="prose prose-sm md:prose max-w-none mb-8">
              {post.content.split('\n').map((paragraph: string, index: number) => {
                // 处理Markdown格式
                if (paragraph.startsWith('**') && paragraph.endsWith('**')) {
                  return (
                    <h3 key={index} className="font-bold text-lg mt-6 mb-2">
                      {paragraph.replace(/\*\*/g, '')}
                    </h3>
                  )
                }
                if (paragraph.startsWith('- ')) {
                  return (
                    <li key={index} className="ml-4">
                      {paragraph.substring(2)}
                    </li>
                  )
                }
                if (/^\d+\.\s/.test(paragraph)) {
                  return (
                    <li key={index} className="ml-4 list-decimal">
                      {paragraph.replace(/^\d+\.\s/, '')}
                    </li>
                  )
                }
                if (paragraph.trim() === '') {
                  return <br key={index} />
                }
                return (
                  <p key={index} className="mb-3 leading-relaxed">
                    {paragraph}
                  </p>
                )
              })}
            </div>

            {/* 统计信息 */}
            <div className="flex items-center gap-6 pt-6 border-t text-muted-foreground">
              <div className="flex items-center gap-2">
                <Eye className="h-5 w-5" />
                <span>{post.view_count} 浏览</span>
              </div>
              <div className="flex items-center gap-2">
                <Heart className="h-5 w-5" />
                <span>{post.like_count} 点赞</span>
              </div>
              <div className="flex items-center gap-2">
                <MessageCircle className="h-5 w-5" />
                <span>{post.comment_count} 评论</span>
              </div>
            </div>
          </CardContent>
        </Card>

        {/* 评论区占位 */}
        <Card className="mt-6">
          <CardContent className="p-6">
            <h3 className="font-bold text-lg mb-4">评论区</h3>
            <p className="text-muted-foreground text-center py-8">
              评论功能开发中...
            </p>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
