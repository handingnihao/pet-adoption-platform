import { create } from 'zustand'
import { persist } from 'zustand/middleware'
import { userApi } from '../lib/api'
import type { User } from '../lib/api'

interface AuthState {
  user: User | null
  token: string | null
  isAuthenticated: boolean
  login: (username: string, password: string) => Promise<void>
  register: (data: { username: string; password: string; email?: string; phone?: string }) => Promise<void>
  logout: () => void
  refreshUser: () => Promise<void>
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      user: null,
      token: null,
      isAuthenticated: false,

      login: async (username: string, password: string) => {
        const response = await userApi.login({ username, password })
        const { token, user } = response.data
        localStorage.setItem('token', token)
        set({ user, token, isAuthenticated: true })
      },

      register: async (data) => {
        await userApi.register(data)
      },

      logout: () => {
        localStorage.removeItem('token')
        set({ user: null, token: null, isAuthenticated: false })
      },

      refreshUser: async () => {
        try {
          const response = await userApi.getProfile()
          set({ user: response.data })
        } catch {
          set({ user: null, token: null, isAuthenticated: false })
        }
      },
    }),
    {
      name: 'auth-storage',
      partialize: (state) => ({ token: state.token, user: state.user, isAuthenticated: state.isAuthenticated }),
    }
  )
)
