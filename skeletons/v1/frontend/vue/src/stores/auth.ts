import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

interface User {
  id: string
  name: string
  email: string
}

export const useAuthStore = defineStore(
  'auth',
  () => {
    const user = ref<User | null>(null)
    const token = ref<string | null>(null)
    const isAuthenticated = computed(() => !!user.value)

    function setUser(newUser: User) {
      user.value = newUser
    }
    function setToken(newToken: string) {
      token.value = newToken
    }
    function logout() {
      user.value = null
      token.value = null
    }

    return { user, token, isAuthenticated, setUser, setToken, logout }
  },
  { persist: true }
)
