import { api } from './api'
import type { ApiResponse, PageResponse, User, Pet, AdoptionApplication } from './api'

// 管理员 API
export const adminApi = {
  // 用户管理
  getUsers: (params?: { page?: number; page_size?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<User>>>('/users', { params }),
  searchUsers: (params: { keyword: string; page?: number; page_size?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<User>>>('/users/search', { params }),
  getUserById: (id: number) =>
    api.get<unknown, ApiResponse<User>>(`/users/${id}`),
  disableUser: (id: number) =>
    api.put<unknown, ApiResponse>(`/users/${id}/disable`),
  enableUser: (id: number) =>
    api.put<unknown, ApiResponse>(`/users/${id}/enable`),

  // 宠物管理
  getPendingPets: (params?: { page?: number; page_size?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<Pet>>>('/pets/pending', { params }),
  approvePet: (id: number) =>
    api.put<unknown, ApiResponse>(`/pets/${id}/approve`),
  rejectPet: (id: number, reason?: string) =>
    api.put<unknown, ApiResponse>(`/pets/${id}/reject`, { reason }),
  getPetStatistics: () =>
    api.get<unknown, ApiResponse<PetStatistics>>('/pets/statistics'),

  // 领养管理
  getAllApplications: (params?: { page?: number; page_size?: number; status?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<AdoptionApplication>>>('/adoptions/applications', { params }),
  getPendingApplications: (params?: { page?: number; page_size?: number }) =>
    api.get<unknown, ApiResponse<PageResponse<AdoptionApplication>>>('/adoptions/applications/pending', { params }),
  reviewApplication: (id: number, data: { status: number; remark?: string }) =>
    api.put<unknown, ApiResponse>(`/adoptions/applications/${id}/review`, data),
  getAdoptionStatistics: () =>
    api.get<unknown, ApiResponse<AdoptionStatistics>>('/adoptions/statistics'),

  // 机构管理
  updateOrganizationStatus: (id: number, data: { status: number; remark?: string }) =>
    api.put<unknown, ApiResponse>(`/organizations/${id}/status`, data),
}

// 统计类型
export interface PetStatistics {
  total: number
  pending: number
  approved: number
  rejected: number
  by_type: Record<string, number>
}

export interface AdoptionStatistics {
  total_applications: number
  pending_applications: number
  approved_applications: number
  rejected_applications: number
  total_adoptions: number
}
