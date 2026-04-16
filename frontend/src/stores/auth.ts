import { defineStore } from 'pinia'
import { ref } from 'vue'

export interface User {
  id: string
  name: string
  email?: string
  sso_provider?: string
  sso_id?: string
  avatar_url?: string
  phone?: string
  institution?: string
  is_profile_completed?: boolean
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('auth_token'))
  const user = ref<User | null>(null)

  const login = (jwt: string, userData: User) => {
    token.value = jwt
    user.value = userData
    localStorage.setItem('auth_token', jwt)
  }

  const logout = () => {
    token.value = null
    user.value = null
    localStorage.removeItem('auth_token')
  }

  const isAuthenticated = () => !!token.value

  return { token, user, login, logout, isAuthenticated }
})
