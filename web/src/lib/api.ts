import axios, { AxiosError } from 'axios'
import type { InternalAxiosRequestConfig } from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1'

export const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// 请求拦截器
api.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const token = localStorage.getItem('token')
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => Promise.reject(error)
)

// 响应拦截器
api.interceptors.response.use(
  (response) => response.data,
  (error: AxiosError<{ message?: string }>) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      window.location.href = '/login'
    }
    const message = error.response?.data?.message || '请求失败'
    return Promise.reject(new Error(message))
  }
)

// API 响应类型
export interface ApiResponse<T = unknown> {
  code: number
  message: string
  data: T
}

export interface PageResponse<T> {
  list: T[]
  total?: number
  page?: number
  page_size?: number
  pagination?: {
    total: number
    page: number
    page_size: number
    total_pages: number
  }
}

// 用户相关 API
export const userApi = {
  login: (data: { username: string; password: string }) =>
    api.post<unknown, ApiResponse<{ token: string; user: User }>>('/users/login', data),
  register: (data: { username: string; password: string; email?: string; phone?: string }) =>
    api.post<unknown, ApiResponse>('/users/register', data),
  getProfile: () => api.get<unknown, ApiResponse<User>>('/users/profile'),
  updateProfile: (data: Partial<User>) =>
    api.put<unknown, ApiResponse>('/users/profile', data),
  updatePassword: (data: { old_password: string; new_password: string }) =>
    api.put<unknown, ApiResponse>('/users/password', data),
}

// 宠物相关 API
export const petApi = {
  list: (params?: { page?: number; page_size?: number; type?: string; status?: string }) =>
    api.get<unknown, ApiResponse<PageResponse<Pet>>>('/pets/query', { params }),
  search: (params: { keyword: string; page?: number; page_size?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<Pet>>>('/pets/search', { params }),
  getById: (id: number) => api.get<unknown, ApiResponse<Pet>>(`/pets/${id}`),
  create: (data: PetCreateRequest) => api.post<unknown, ApiResponse<Pet>>('/pets', data),
  update: (id: number, data: Partial<PetCreateRequest>) =>
    api.put<unknown, ApiResponse>(`/pets/${id}`, data),
  delete: (id: number) => api.delete<unknown, ApiResponse>(`/pets/${id}`),
  getMyPets: (params?: { page?: number; page_size?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<Pet>>>('/pets/my', { params }),
  getRecommended: () => api.get<unknown, ApiResponse<Pet[]>>('/pets/recommended'),
  getStatistics: () => api.get<unknown, ApiResponse<PetStatistics>>('/pets/statistics'),
}

// 宠物统计类型
export interface PetStatistics {
  total: number
  available: number
  pending: number
  adopted: number
  offline: number
  by_type: Record<string, number>
}

// 领养相关 API
export const adoptionApi = {
  createApplication: (data: AdoptionApplicationRequest) =>
    api.post<unknown, ApiResponse>('/adoptions/applications', data),
  getMyApplications: (params?: { page?: number; page_size?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<AdoptionApplication>>>('/adoptions/applications/my', { params }),
  getApplicationById: (id: number) =>
    api.get<unknown, ApiResponse<AdoptionApplication>>(`/adoptions/applications/${id}`),
  cancelApplication: (id: number) =>
    api.delete<unknown, ApiResponse>(`/adoptions/applications/${id}`),
  getMyRecords: (params?: { page?: number; page_size?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<AdoptionRecord>>>('/adoptions/records/my', { params }),
}

// 社区相关 API
export const communityApi = {
  listPosts: (params?: { page?: number; page_size?: number; type?: string }) =>
    api.get<unknown, ApiResponse<PageResponse<Post>>>('/community/posts', { params }),
  searchPosts: (params: { keyword: string; page?: number; page_size?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<Post>>>('/community/posts/search', { params }),
  getPost: (id: number) => api.get<unknown, ApiResponse<Post>>(`/community/posts/${id}`),
  createPost: (data: PostCreateRequest) =>
    api.post<unknown, ApiResponse<Post>>('/community/posts', data),
  updatePost: (id: number, data: Partial<PostCreateRequest>) =>
    api.put<unknown, ApiResponse>(`/community/posts/${id}`, data),
  deletePost: (id: number) => api.delete<unknown, ApiResponse>(`/community/posts/${id}`),
  getMyPosts: (params?: { page?: number; page_size?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<Post>>>('/community/posts/my', { params }),
  likePost: (id: number) => api.post<unknown, ApiResponse>(`/community/posts/${id}/like`),
  unlikePost: (id: number) => api.delete<unknown, ApiResponse>(`/community/posts/${id}/like`),
  getComments: (postId: number, params?: { page?: number; page_size?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<Comment>>>(`/community/posts/${postId}/comments`, { params }),
  createComment: (data: CommentCreateRequest) =>
    api.post<unknown, ApiResponse<Comment>>('/community/comments', data),
  deleteComment: (id: number) => api.delete<unknown, ApiResponse>(`/community/comments/${id}`),
  likeComment: (id: number) => api.post<unknown, ApiResponse>(`/community/comments/${id}/like`),
  unlikeComment: (id: number) => api.delete<unknown, ApiResponse>(`/community/comments/${id}/like`),
}

// 机构相关 API
export const organizationApi = {
  list: (params?: { page?: number; page_size?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<Organization>>>('/organizations', { params }),
  getById: (id: number) => api.get<unknown, ApiResponse<Organization>>(`/organizations/${id}`),
  create: (data: OrganizationCreateRequest) =>
    api.post<unknown, ApiResponse<Organization>>('/organizations', data),
  update: (id: number, data: Partial<OrganizationCreateRequest>) =>
    api.put<unknown, ApiResponse>(`/organizations/${id}`, data),
  delete: (id: number) => api.delete<unknown, ApiResponse>(`/organizations/${id}`),
  getMyOrganizations: () =>
    api.get<unknown, ApiResponse<Organization[]>>('/organizations/my'),
}

// 类型定义
export interface User {
  id: number
  username: string
  nickname?: string
  email?: string
  phone?: string
  avatar?: string
  gender?: number
  role: string
  status: number
  created_at: string
}

export interface Pet {
  id: number
  user_id: number
  name: string
  type: string
  breed?: string
  gender: string
  age: number
  size?: string
  color?: string
  weight?: number
  is_vaccinated: boolean
  is_sterilized: boolean
  health_status?: string
  description?: string
  character?: string
  cover_photo?: string
  photos?: string[]
  province?: string
  city?: string
  district?: string
  address?: string
  status: number
  view_count: number
  created_at: string
  user?: User
}

export interface PetCreateRequest {
  name: string
  type: string
  breed?: string
  gender: string
  age: number
  size?: string
  color?: string
  weight?: number
  is_vaccinated?: boolean
  is_sterilized?: boolean
  health_status?: string
  description?: string
  character?: string
  cover_photo?: string
  photos?: string[]
  province?: string
  city?: string
  district?: string
  address?: string
}

export interface AdoptionApplication {
  id: number
  pet_id: number
  user_id: number
  reason: string
  experience?: string
  living_condition?: string
  family_agreement?: boolean
  status: number
  created_at: string
  pet?: Pet
  user?: User
}

export interface AdoptionApplicationRequest {
  pet_id: number
  reason: string
  experience?: string
  living_condition?: string
  family_agreement?: boolean
}

export interface AdoptionRecord {
  id: number
  pet_id: number
  adopter_id: number
  status: number
  adopted_at: string
  pet?: Pet
  adopter?: User
}

export interface Post {
  id: number | string
  user_id: number
  username?: string
  user_avatar?: string
  title?: string
  content: string
  images?: string[]
  video_url?: string
  type: string
  view_count: number
  like_count: number
  comment_count: number
  share_count: number
  is_liked: boolean
  is_top: number
  created_at: string
}

export interface PostCreateRequest {
  title?: string
  content: string
  images?: string[]
  video_url?: string
  type?: string
  pet_id?: number
}

export interface Comment {
  id: number
  post_id: number
  user_id: number
  username?: string
  user_avatar?: string
  parent_id: number
  reply_to_user_id?: number
  reply_to_username?: string
  content: string
  like_count: number
  is_liked: boolean
  created_at: string
  replies?: Comment[]
}

export interface CommentCreateRequest {
  post_id: number
  content: string
  parent_id?: number
  reply_to_user_id?: number
}

export interface Organization {
  id: number
  name: string
  type: string
  description?: string
  logo?: string
  contact_name?: string
  contact_phone?: string
  contact_email?: string
  province?: string
  city?: string
  district?: string
  address?: string
  status: number
  created_at: string
}

export interface OrganizationCreateRequest {
  name: string
  type: string
  description?: string
  logo?: string
  contact_name?: string
  contact_phone?: string
  contact_email?: string
  province?: string
  city?: string
  district?: string
  address?: string
}
