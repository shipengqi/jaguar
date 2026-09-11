import { create } from 'zustand'
import { devtools, persist } from 'zustand/middleware'
import { immer } from 'zustand/middleware/immer'

interface User {
  id: string
  name: string
  email: string
}

interface AuthState {
  user: User | null
  token: string | null
  isAuthenticated: boolean
  setUser: (user: User) => void
  setToken: (token: string) => void
  logout: () => void
}

export const useAuthStore = create<AuthState>()(
  devtools(
    persist(
      immer((set) => ({
        user: null,
        token: null,
        isAuthenticated: false,

        setUser: (user) =>
          set((state) => {
            state.user = user
            state.isAuthenticated = true
          }),

        setToken: (token) =>
          set((state) => {
            state.token = token
          }),

        logout: () =>
          set((state) => {
            state.user = null
            state.token = null
            state.isAuthenticated = false
          }),
      })),
      { name: 'auth-storage' }
    )
  )
)
