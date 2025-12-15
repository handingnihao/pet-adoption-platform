import { useState } from 'react'
import { useSearchParams, useNavigate } from 'react-router-dom'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Heart, MessageCircle, Eye, Plus, Search } from 'lucide-react'
import { Button } from '../components/ui/button'
import { Input } from '../components/ui/input'
import { Card, CardContent } from '../components/ui/card'
import { communityApi } from '../lib/api'
import type { Post } from '../lib/api'
import { useAuthStore } from '../store/auth'

const postTypes = [
  { value: '', label: '全部' },
  { value: 'story', label: '领养故事' },
  { value: 'knowledge', label: '养宠知识' },
  { value: 'daily', label: '日常分享' },
  { value: 'other', label: '其他' },
]

export function Community() {
  const [searchParams, setSearchParams] = useSearchParams()
  const navigate = useNavigate()
  const queryClient = useQueryClient()
  const { isAuthenticated } = useAuthStore()
  const [keyword, setKeyword] = useState('')

  const currentType = searchParams.get('type') || ''
  const currentPage = parseInt(searchParams.get('page') || '1')

  const { data, isLoading } = useQuery({
    queryKey: ['posts', currentType, currentPage],
    queryFn: () => communityApi.listPosts({
      page: currentPage,
      page_size: 10,
      type: currentType || undefined,
    }),
  })

  const likeMutation = useMutation({
    mutationFn: async ({ id, isLiked }: { id: number; isLiked: boolean }) => {
      if (isLiked) {
        return communityApi.unlikePost(id)
      }
      return communityApi.likePost(id)
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['posts'] })
    },
  })

  const posts = data?.data?.list || []
  const total = data?.data?.pagination?.total || data?.data?.total || 0
  const totalPages = Math.ceil(total / 10)

  const handleTypeChange = (type: string) => {
    if (type) {
      setSearchParams({ type })
    } else {
      setSearchParams({})
    }
  }

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault()
    if (keyword.trim()) {
      navigate(`/community/search?keyword=${encodeURIComponent(keyword.trim())}`)
    }
  }

  const handleLike = (post: Post) => {
    if (!isAuthenticated) {
      navigate('/login')
      return
    }
    likeMutation.mutate({ id: post.id, isLiked: post.is_liked })
  }

  return (
    <div className="container py-8">
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold">社区动态</h1>
        {isAuthenticated && (
          <Button onClick={() => navigate('/community/create')}>
            <Plus className="w-4 h-4 mr-2" />
            发布动态
          </Button>
        )}
      </div>

      {/* 搜索和筛选 */}
      <div className="flex flex-col md:flex-row gap-4 mb-6">
        <form onSubmit={handleSearch} className="flex gap-2 flex-1 max-w-md">
          <div className="relative flex-1">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="搜索动态..."
              value={keyword}
              onChange={(e) => setKeyword(e.target.value)}
              className="pl-10"
            />
          </div>
          <Button type="submit" variant="outline">搜索</Button>
        </form>

        <div className="flex gap-2 flex-wrap">
          {postTypes.map((type) => (
            <Button
              key={type.value}
              variant={currentType === type.value ? 'default' : 'outline'}
              size="sm"
              onClick={() => handleTypeChange(type.value)}
            >
              {type.label}
            </Button>
          ))}
        </div>
      </div>

      {/* 动态列表 */}
      {isLoading ? (
        <div className="space-y-4">
          {[...Array(3)].map((_, i) => (
            <Card key={i} className="animate-pulse">
              <CardContent className="p-6">
                <div className="flex gap-3 mb-4">
                  <div className="w-10 h-10 rounded-full bg-muted" />
                  <div className="space-y-2">
                    <div className="h-4 w-24 bg-muted rounded" />
                    <div className="h-3 w-16 bg-muted rounded" />
                  </div>
                </div>
                <div className="h-20 bg-muted rounded" />
              </CardContent>
            </Card>
          ))}
        </div>
      ) : posts.length === 0 ? (
        <div className="text-center py-16">
          <div className="text-6xl mb-4">📝</div>
          <h3 className="text-lg font-medium mb-2">暂无动态</h3>
          <p className="text-muted-foreground">成为第一个分享故事的人吧</p>
        </div>
      ) : (
        <div className="space-y-4">
          {posts.map((post: Post) => (
            <PostCard
              key={post.id}
              post={post}
              onLike={() => handleLike(post)}
              onClick={() => navigate(`/community/${post.id}`)}
            />
          ))}

          {/* 分页 */}
          {totalPages > 1 && (
            <div className="flex justify-center gap-2 mt-8">
              <Button
                variant="outline"
                disabled={currentPage <= 1}
                onClick={() => setSearchParams({ ...Object.fromEntries(searchParams), page: String(currentPage - 1) })}
              >
                上一页
              </Button>
              <span className="flex items-center px-4">
                {currentPage} / {totalPages}
              </span>
              <Button
                variant="outline"
                disabled={currentPage >= totalPages}
                onClick={() => setSearchParams({ ...Object.fromEntries(searchParams), page: String(currentPage + 1) })}
              >
                下一页
              </Button>
            </div>
          )}
        </div>
      )}
    </div>
  )
}

function PostCard({ post, onLike, onClick }: { post: Post; onLike: () => void; onClick: () => void }) {
  return (
    <Card className="hover:shadow-md transition-shadow cursor-pointer" onClick={onClick}>
      <CardContent className="p-6">
        {/* 用户信息 */}
        <div className="flex items-center gap-3 mb-4">
          <div className="w-10 h-10 rounded-full bg-primary/10 flex items-center justify-center">
            {post.user_avatar ? (
              <img src={post.user_avatar} alt="" className="w-full h-full rounded-full object-cover" />
            ) : (
              <span className="text-primary font-medium">
                {post.username?.[0] || '?'}
              </span>
            )}
          </div>
          <div>
            <p className="font-medium">{post.username || '匿名用户'}</p>
            <p className="text-xs text-muted-foreground">
              {new Date(post.created_at).toLocaleDateString()}
            </p>
          </div>
          <span className="ml-auto px-2 py-1 text-xs bg-muted rounded">
            {postTypes.find(t => t.value === post.type)?.label || '其他'}
          </span>
        </div>

        {/* 内容 */}
        {post.title && (
          <h3 className="font-semibold mb-2">{post.title}</h3>
        )}
        <p className="text-muted-foreground line-clamp-3 mb-4">{post.content}</p>

        {/* 图片 */}
        {post.images && post.images.length > 0 && (
          <div className="flex gap-2 mb-4 overflow-hidden">
            {post.images.slice(0, 3).map((img, idx) => (
              <div key={idx} className="w-24 h-24 rounded-lg overflow-hidden bg-muted flex-shrink-0">
                <img src={img} alt="" className="w-full h-full object-cover" />
              </div>
            ))}
            {post.images.length > 3 && (
              <div className="w-24 h-24 rounded-lg bg-muted flex items-center justify-center text-muted-foreground">
                +{post.images.length - 3}
              </div>
            )}
          </div>
        )}

        {/* 互动数据 */}
        <div className="flex items-center gap-6 text-sm text-muted-foreground">
          <button
            className={`flex items-center gap-1 hover:text-primary ${post.is_liked ? 'text-red-500' : ''}`}
            onClick={(e) => {
              e.stopPropagation()
              onLike()
            }}
          >
            <Heart className={`w-4 h-4 ${post.is_liked ? 'fill-current' : ''}`} />
            {post.like_count || 0}
          </button>
          <span className="flex items-center gap-1">
            <MessageCircle className="w-4 h-4" />
            {post.comment_count || 0}
          </span>
          <span className="flex items-center gap-1">
            <Eye className="w-4 h-4" />
            {post.view_count || 0}
          </span>
        </div>
      </CardContent>
    </Card>
  )
}
