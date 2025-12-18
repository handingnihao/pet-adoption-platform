import { describe, it, expect, beforeEach } from 'vitest'
import { useAuthStore } from '../auth'

describe('authStore', () => {
  beforeEach(() => {
    // 重置store
    useAuthStore.setState({
      user: null,
      token: null,
      isAuthenticated: false
    })
  })

  it('初始状态', () => {
    const state = useAuthStore.getState()
    expect(state.user).toBeNull()
    expect(state.token).toBeNull()
    expect(state.isAuthenticated).toBe(false)
  })

  it('设置登录状态', () => {
    const mockUser = { 
      id: 1, 
      username: 'testuser',
      role: 'user',
      status: 1,
      created_at: '2024-01-01'
    }
    const mockToken = 'test-token'

    // 模拟登录
    useAuthStore.setState({
      user: mockUser,
      token: mockToken,
      isAuthenticated: true
    })

    const state = useAuthStore.getState()
    expect(state.user).toEqual(mockUser)
    expect(state.token).toBe(mockToken)
    expect(state.isAuthenticated).toBe(true)
  })

  it('登出', () => {
    // 先设置登录状态
    useAuthStore.setState({
      user: { 
        id: 1, 
        username: 'test',
        role: 'user',
        status: 1,
        created_at: '2024-01-01'
      },
      token: 'token',
      isAuthenticated: true
    })

    // 执行登出
    useAuthStore.getState().logout()

    const state = useAuthStore.getState()
    expect(state.user).toBeNull()
    expect(state.token).toBeNull()
    expect(state.isAuthenticated).toBe(false)
  })
})
