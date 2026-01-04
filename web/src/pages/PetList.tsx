import { useState } from 'react'
import { Link, useSearchParams } from 'react-router-dom'
import { Search, Filter, MapPin } from 'lucide-react'
import { useQuery } from '@tanstack/react-query'
import { Button } from '../components/ui/button'
import { Input } from '../components/ui/input'
import { Card, CardContent } from '../components/ui/card'
import { petApi } from '../lib/api'
import type { Pet } from '../lib/api'

const petTypes = [
  { value: '', label: '全部' },
  { value: 'cat', label: '猫咪' },
  { value: 'dog', label: '狗狗' },
  { value: 'rabbit', label: '兔子' },
  { value: 'hamster', label: '仓鼠' },
  { value: 'bird', label: '鸟类' },
  { value: 'other', label: '其他' },
]

export function PetList() {
  const [searchParams, setSearchParams] = useSearchParams()
  const [keyword, setKeyword] = useState(searchParams.get('keyword') || '')

  const currentType = searchParams.get('type') || ''
  const currentPage = parseInt(searchParams.get('page') || '1')

  const { data, isLoading } = useQuery({
    queryKey: ['pets', currentType, currentPage, searchParams.get('keyword')],
    queryFn: () => {
      const kw = searchParams.get('keyword')
      if (kw) {
        return petApi.search({ keyword: kw, page: currentPage, page_size: 12 })
      }
      return petApi.list({ page: currentPage, page_size: 12, type: currentType || undefined })
    },
  })

  const pets = data?.data?.list || []
  const total = data?.data?.pagination?.total || data?.data?.total || 0
  const totalPages = Math.ceil(total / 12)

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault()
    if (keyword.trim()) {
      setSearchParams({ keyword: keyword.trim() })
    } else {
      searchParams.delete('keyword')
      setSearchParams(searchParams)
    }
  }

  const handleTypeChange = (type: string) => {
    if (type) {
      setSearchParams({ type })
    } else {
      searchParams.delete('type')
      searchParams.delete('keyword')
      setSearchParams(searchParams)
    }
  }

  const handlePageChange = (page: number) => {
    searchParams.set('page', page.toString())
    setSearchParams(searchParams)
  }

  return (
    <div className="container py-8">
      {/* 搜索栏 */}
      <div className="mb-8">
        <form onSubmit={handleSearch} className="flex gap-4 max-w-xl">
          <div className="relative flex-1">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="搜索宠物名称、品种..."
              value={keyword}
              onChange={(e) => setKeyword(e.target.value)}
              className="pl-10"
            />
          </div>
          <Button type="submit">搜索</Button>
        </form>
      </div>

      {/* 筛选标签 */}
      <div className="flex items-center gap-2 mb-6 flex-wrap">
        <Filter className="h-4 w-4 text-muted-foreground" />
        {petTypes.map((type) => (
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

      {/* 宠物列表 */}
      {isLoading ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          {[...Array(8)].map((_, i) => (
            <Card key={i} className="animate-pulse">
              <div className="aspect-square bg-muted" />
              <CardContent className="p-4">
                <div className="h-5 bg-muted rounded mb-2" />
                <div className="h-4 bg-muted rounded w-2/3" />
              </CardContent>
            </Card>
          ))}
        </div>
      ) : pets.length === 0 ? (
        <div className="text-center py-16">
          <div className="text-6xl mb-4">🐾</div>
          <h3 className="text-lg font-medium mb-2">暂无宠物</h3>
          <p className="text-muted-foreground">换个条件试试吧</p>
        </div>
      ) : (
        <>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            {pets.map((pet: Pet) => (
              <PetCard key={pet.id} pet={pet} />
            ))}
          </div>

          {/* 分页 */}
          {totalPages > 1 && (
            <div className="flex justify-center gap-2 mt-8">
              <Button
                variant="outline"
                disabled={currentPage <= 1}
                onClick={() => handlePageChange(currentPage - 1)}
              >
                上一页
              </Button>
              <span className="flex items-center px-4">
                {currentPage} / {totalPages}
              </span>
              <Button
                variant="outline"
                disabled={currentPage >= totalPages}
                onClick={() => handlePageChange(currentPage + 1)}
              >
                下一页
              </Button>
            </div>
          )}
        </>
      )}
    </div>
  )
}

function PetCard({ pet }: { pet: Pet }) {
  const genderText = pet.gender === 'male' ? '♂' : pet.gender === 'female' ? '♀' : ''
  const genderColor = pet.gender === 'male' ? 'text-blue-500' : 'text-pink-500'

  return (
    <Link to={`/pets/${pet.id}`}>
      <Card className="overflow-hidden hover:shadow-lg transition-shadow group">
        <div className="aspect-square relative bg-muted">
          {pet.cover_photo ? (
            <img
              src={pet.cover_photo}
              alt={pet.name}
              className="w-full h-full object-cover group-hover:scale-105 transition-transform"
            />
          ) : (
            <div className="w-full h-full flex items-center justify-center text-6xl">
              {pet.type === 'cat' ? '🐱' : pet.type === 'dog' ? '🐕' : '🐾'}
            </div>
          )}
        </div>
        <CardContent className="p-4">
          <div className="flex items-center justify-between mb-2">
            <h3 className="font-semibold truncate">{pet.name}</h3>
            <span className={`text-lg ${genderColor}`}>{genderText}</span>
          </div>
          <p className="text-sm text-muted-foreground mb-2">
            {pet.breed || pet.type} · {pet.age}个月
          </p>
          {(pet.city || pet.district) && (
            <p className="text-xs text-muted-foreground flex items-center">
              <MapPin className="h-3 w-3 mr-1" />
              {pet.city} {pet.district}
            </p>
          )}
        </CardContent>
      </Card>
    </Link>
  )
}
