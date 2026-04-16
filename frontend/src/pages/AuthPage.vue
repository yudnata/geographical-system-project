<script setup lang="ts">
import { useRouter } from 'vue-router'
import { decodeCredential } from 'vue3-google-login'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

interface GoogleCredentialResponse {
  credential: string;
}

interface GoogleUserData {
  email: string;
  name: string;
  sub: string;
  picture: string;
}

const handleSSOSuccess = async (response: GoogleCredentialResponse) => {
  try {
    const userData = decodeCredential(response.credential) as GoogleUserData

    const res = await fetch(`${API_URL}/auth/sso`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        email: userData.email,
        name: userData.name,
        sso_provider: 'google',
        sso_id: userData.sub,
        avatar_url: userData.picture
      })
    })

    const result = await res.json()
    if (result.success) {
      authStore.login(result.data.token, result.data.user)
      router.push('/dashboard')
    } else {
      alert(result.message || 'Login SSO gagal di sisi server.')
    }
  } catch (err) {
    console.error('SSO Error:', err)
    alert('Terjadi kesalahan saat memproses login Google.')
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 -mt-20">
    <div class="bg-white p-8 rounded-xl shadow-[0_4px_24px_rgba(0,0,0,0.06)] w-full max-w-sm">
      <h1 class="text-2xl font-bold text-center text-primary mb-2">Masuk</h1>
      <p class="text-sm text-center text-gray-500 mb-6">Akses kontrol SIG Tata Ruang Anda</p>

      <div class="space-y-4">

        <div>
          <label class="block text-xs font-semibold text-gray-600 uppercase mb-1">Username</label>
          <input type="text" class="w-full px-4 py-2 bg-gray-100 border-transparent rounded-lg focus:bg-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm"
            placeholder="Masukkan username">
        </div>
        <div>
          <label class="block text-xs font-semibold text-gray-600 uppercase mb-1">Password</label>
          <input type="password" class="w-full px-4 py-2 bg-gray-100 border-transparent rounded-lg focus:bg-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm"
            placeholder="••••••••">
        </div>
        <button class="w-full bg-primary text-white font-bold py-2.5 rounded-lg hover:bg-primary-hover active:scale-[0.98] transition-all">
          Masuk Sekarang
        </button>

        <div class="relative flex py-2 items-center">
          <div class="flex-grow border-t border-gray-200"></div>
          <span class="flex-shrink-0 mx-4 text-gray-400 text-xs">ATAU</span>
          <div class="flex-grow border-t border-gray-200"></div>
        </div>

        <div class="flex justify-center w-full">

          <GoogleLogin :callback="handleSSOSuccess" />
        </div>
      </div>
    </div>
  </div>
</template>
