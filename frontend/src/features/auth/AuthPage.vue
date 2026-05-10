<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { googleTokenLogin } from 'vue3-google-login'
import { useAuthStore } from '@/stores/auth'
import { useNotificationStore } from '@/stores/notifications'

const router = useRouter()
const authStore = useAuthStore()
const notificationStore = useNotificationStore()

const mode = ref<'login' | 'register'>('login')
const isLoading = ref(false)
const showPassword = ref(false)
const errorMsg = ref('')

const form = ref({
  email: '',
  name: '',
  password: '',
})

const handleEmailSubmit = async () => {
  errorMsg.value = ''
  if (!form.value.email || !form.value.password) {
    errorMsg.value = 'Email dan password wajib diisi!'
    return
  }
  if (mode.value === 'register' && !form.value.name) {
    errorMsg.value = 'Nama wajib diisi saat mendaftar!'
    return
  }

  isLoading.value = true
  try {
    let result
    if (mode.value === 'login') {
      result = await authStore.loginWithEmail(form.value.email, form.value.password)
    } else {
      result = await authStore.registerWithEmail(form.value.name, form.value.email, form.value.password)
    }

    if (result.success) {
      if (mode.value === 'login') {
        notificationStore.success('Selamat datang kembali!')
        router.push('/')
      } else {
        notificationStore.success('Akun berhasil dibuat! Silakan login.')
        mode.value = 'login'
        form.value.name = ''
        form.value.password = ''
      }
    } else {
      errorMsg.value = result.message || 'Terjadi kesalahan.'
    }
  } catch (err) {
    console.error('Auth Error:', err)
    errorMsg.value = 'Tidak dapat terhubung ke server.'
  } finally {
    isLoading.value = false
  }
}

const loginWithGoogle = async () => {
  errorMsg.value = ''
  isLoading.value = true
  try {
    const tokenResponse = await googleTokenLogin()
    const userInfoRes = await fetch('https://www.googleapis.com/oauth2/v3/userinfo', {
      headers: { Authorization: `Bearer ${tokenResponse.access_token}` },
    })
    const userInfo = await userInfoRes.json()

    const result = await authStore.loginWithSSO(userInfo)

    if (result.success) {
      notificationStore.success('Selamat datang kembali!')
      router.push('/')
    } else {
      errorMsg.value = result.message || 'Login SSO gagal di sisi server.'
    }
  } catch (err) {
    console.error('SSO Error:', err)
    errorMsg.value = 'Terjadi kesalahan saat login Google.'
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen h-screen flex items-center justify-center bg-gray-50">
    <div class="bg-white p-8 rounded-xl shadow-[0_4px_24px_rgba(0,0,0,0.06)] w-full max-w-sm">

      <!-- Header -->
      <h1 class="text-2xl font-bold text-center text-primary mb-1">
        {{ mode === 'login' ? 'Masuk' : 'Daftar Akun' }}
      </h1>
      <p class="text-sm text-center text-gray-500 mb-6">
        {{ mode === 'login' ? 'Akses kontrol SIG Tata Ruang Anda' : 'Buat akun baru dengan email' }}
      </p>

      <div class="space-y-4">

        <!-- Field Nama (hanya saat Register) -->
        <div v-if="mode === 'register'">
          <label class="block text-xs font-semibold text-gray-600 uppercase mb-1">Nama Lengkap</label>
          <input v-model="form.name" type="text"
            class="w-full px-4 py-2 bg-gray-100 border-transparent rounded-lg focus:bg-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm"
            placeholder="Contoh: Budi Santoso" @keyup.enter="handleEmailSubmit">
        </div>

        <div>
          <label class="block text-xs font-semibold text-gray-600 uppercase mb-1">Email</label>
          <input v-model="form.email" type="email"
            class="w-full px-4 py-2 bg-gray-100 border-transparent rounded-lg focus:bg-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm"
            placeholder="contoh@email.com" @keyup.enter="handleEmailSubmit">
        </div>

        <div>
          <label class="block text-xs font-semibold text-gray-600 uppercase mb-1">Password</label>
          <div class="relative">
            <input v-model="form.password" :type="showPassword ? 'text' : 'password'"
              class="w-full px-4 py-2 pr-10 bg-gray-100 border-transparent rounded-lg focus:bg-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm"
              placeholder="••••••••" @keyup.enter="handleEmailSubmit">
            <button type="button" class="absolute inset-y-0 right-3 flex items-center text-gray-400 hover:text-gray-600" @click="showPassword = !showPassword">
              <svg v-if="!showPassword" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
                <path stroke-linecap="round" stroke-linejoin="round"
                  d="M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178Z" />
                <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z" />
              </svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
                <path stroke-linecap="round" stroke-linejoin="round"
                  d="M3.98 8.223A10.477 10.477 0 0 0 1.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.451 10.451 0 0 1 12 4.5c4.756 0 8.773 3.162 10.065 7.498a10.522 10.522 0 0 1-4.293 5.774M6.228 6.228 3 3m3.228 3.228 3.65 3.65m7.894 7.894L21 21m-3.228-3.228-3.65-3.65m0 0a3 3 0 1 0-4.243-4.243m4.242 4.242L9.88 9.88" />
              </svg>
            </button>
          </div>
        </div>

        <div v-if="errorMsg" class="flex items-start gap-2.5 bg-red-50 border border-red-200 text-red-700 text-xs rounded-lg px-4 py-3">
          <svg class="w-4 h-4 mt-0.5 shrink-0 text-red-500" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z" />
          </svg>
          <span>{{ errorMsg }}</span>
        </div>

        <button @click="handleEmailSubmit" :disabled="isLoading"
          class="w-full bg-primary text-white font-bold py-2.5 rounded-lg hover:bg-primary-hover active:scale-[0.98] transition-all disabled:opacity-60 disabled:cursor-not-allowed flex items-center justify-center gap-2">
          <svg v-if="isLoading" class="animate-spin w-4 h-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z" />
          </svg>
          {{ mode === 'login' ? 'Masuk Sekarang' : 'Buat Akun' }}
        </button>

        <p class="text-center text-xs text-gray-500">
          <span v-if="mode === 'login'">
            Belum punya akun?
            <button @click="mode = 'register'; form.password = ''" class="text-primary font-bold hover:underline">Daftar di sini</button>
          </span>
          <span v-else>
            Sudah punya akun?
            <button @click="mode = 'login'; form.password = ''" class="text-primary font-bold hover:underline">Masuk</button>
          </span>
        </p>

        <div class="relative flex py-1 items-center">
          <div class="flex-grow border-t border-gray-200"></div>
          <span class="flex-shrink-0 mx-4 text-gray-400 text-xs">ATAU</span>
          <div class="flex-grow border-t border-gray-200"></div>
        </div>

        <button @click="loginWithGoogle" :disabled="isLoading"
          class="w-full flex items-center justify-center gap-3 px-4 py-2.5 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 active:scale-[0.98] transition-all shadow-sm disabled:opacity-60 disabled:cursor-not-allowed">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" class="w-5 h-5 shrink-0">
            <path fill="#FFC107"
              d="M43.6 20.1H42V20H24v8h11.3C33.7 32.7 29.2 36 24 36c-6.6 0-12-5.4-12-12s5.4-12 12-12c3.1 0 5.8 1.2 7.9 3l5.7-5.7C34 6.5 29.3 4 24 4 12.9 4 4 12.9 4 24s8.9 20 20 20 20-8.9 20-20c0-1.3-.1-2.7-.4-3.9z" />
            <path fill="#FF3D00" d="m6.3 14.7 6.6 4.8C14.7 16 19 13 24 13c3.1 0 5.8 1.2 7.9 3l5.7-5.7C34 6.5 29.3 4 24 4 16.3 4 9.7 8.3 6.3 14.7z" />
            <path fill="#4CAF50" d="M24 44c5.2 0 9.9-2 13.4-5.2l-6.2-5.2C29.3 35.5 26.8 36 24 36c-5.2 0-9.6-3.3-11.3-7.9l-6.5 5C9.5 39.6 16.2 44 24 44z" />
            <path fill="#1976D2" d="M43.6 20.1H42V20H24v8h11.3c-.8 2.3-2.3 4.2-4.3 5.6l6.2 5.2C36.9 37.3 44 32 44 24c0-1.3-.1-2.7-.4-3.9z" />
          </svg>
          <span class="text-sm font-semibold text-gray-700">
            {{ mode === 'login' ? 'Masuk' : 'Daftar' }} dengan Google
          </span>
        </button>

        <p class="text-center text-[11px] text-gray-400 leading-relaxed">
          Akun yang terdaftar via Google tidak memerlukan password.
        </p>

      </div>
    </div>
  </div>
</template>
