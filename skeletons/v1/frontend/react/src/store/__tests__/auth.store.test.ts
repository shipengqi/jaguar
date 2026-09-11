import { act, renderHook } from '@testing-library/react'
import { useAuthStore } from '../auth.store'

describe('AuthStore', () => {
  beforeEach(() => {
    useAuthStore.setState({ user: null, token: null, isAuthenticated: false })
  })

  it('should set user and mark authenticated', () => {
    const { result } = renderHook(() => useAuthStore())
    const mockUser = { id: '1', name: 'Test', email: 'test@example.com' }

    act(() => {
      result.current.setUser(mockUser)
    })

    expect(result.current.user).toEqual(mockUser)
    expect(result.current.isAuthenticated).toBe(true)
  })

  it('should clear state on logout', () => {
    const { result } = renderHook(() => useAuthStore())

    act(() => {
      result.current.setUser({ id: '1', name: 'Test', email: 'test@test.com' })
      result.current.logout()
    })

    expect(result.current.user).toBeNull()
    expect(result.current.isAuthenticated).toBe(false)
  })
})
