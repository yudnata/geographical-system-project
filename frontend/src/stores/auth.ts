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
  const defaultAvatarUrl = ref<string>('https://res.cloudinary.com/di7pw938v/image/upload/v1706764503/user_3_pvyc8x.png')
  const token = ref<string | null>(localStorage.getItem('auth_token'))
  const user = ref<User | null>(null)

  const fetchConfig = async () => {
    try {
      const resp = await fetch(`${import.meta.env.VITE_API_URL}/api/auth/config`)
      const data = await resp.json()
      if (data.success) {
        defaultAvatarUrl.value = data.data.default_avatar_url
      }
    } catch (e) {
      console.error('Failed to fetch backend config', e)
    }
  }

  fetchConfig()

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

  return { token, user, defaultAvatarUrl, login, logout, isAuthenticated }
})
