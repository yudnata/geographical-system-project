<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const handleClose = () => {
  emit('close')
  isChangingPassword.value = false
  oldPassword.value = ''
  newPassword.value = ''
  confirmPassword.value = ''
  errorMsg.value = ''
}

import { ref } from 'vue'
import { useNotificationStore } from '@/stores/notifications'

const notificationStore = useNotificationStore()

const isChangingPassword = ref(false)
const oldPassword = ref('')
const newPassword = ref('')
const confirmPassword = ref('')
const isLoading = ref(false)
const errorMsg = ref('')

const handleChangePassword = async () => {
  errorMsg.value = ''
  if (authStore.user?.has_password && !oldPassword.value) {
    errorMsg.value = 'Password lama wajib diisi.'
    return
  }
  if (!newPassword.value || newPassword.value.length < 6) {
    errorMsg.value = 'Password baru minimal 6 karakter.'
    return
  }
  if (newPassword.value !== confirmPassword.value) {
    errorMsg.value = 'Konfirmasi password tidak cocok.'
    return
  }

  isLoading.value = true
  try {
    const res = await fetch(`${import.meta.env.VITE_API_URL || 'http://localhost:8080/api'}/auth/password`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${authStore.token}`,
      },
      body: JSON.stringify({
        old_password: oldPassword.value || undefined,
        new_password: newPassword.value,
      }),
    })
    const json = await res.json()
    if (json.success) {
      notificationStore.success('Password berhasil diperbarui!')
      if (authStore.user) authStore.user.has_password = true
      isChangingPassword.value = false
      oldPassword.value = ''
      newPassword.value = ''
      confirmPassword.value = ''
    } else {
      errorMsg.value = json.message || 'Gagal memperbarui password.'
    }
  } catch (err) {
    console.error(err)
    errorMsg.value = 'Terjadi kesalahan jaringan.'
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div v-if="isOpen" class="fixed inset-0 z-[2000] flex items-center justify-center p-4">

    <div class="absolute inset-0 bg-black/20 backdrop-blur-md transition-opacity" @click="handleClose"></div>


    <div
      class="relative w-full max-w-sm bg-white/90 backdrop-blur-2xl rounded-[2.5rem] shadow-[0_32px_64px_-16px_rgba(0,0,0,0.1)] overflow-hidden transform transition-all animate-in fade-in zoom-in slide-in-from-bottom-8 duration-500 border border-white/50 p-8">


      <button @click="handleClose" class="absolute top-6 right-6 p-2 text-gray-400 hover:text-gray-900 transition-colors">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
          <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>

      <div class="flex flex-col items-center">

        <div class="relative group">
          <div class="w-24 h-24 rounded-3xl overflow-hidden bg-gray-50 border border-gray-100 flex items-center justify-center shadow-inner">
            <img v-if="authStore.user?.avatar_url" :src="authStore.user.avatar_url" class="w-full h-full object-cover transition-transform group-hover:scale-110 duration-500" />
            <div v-else class="text-gray-300">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
                <path fill-rule="evenodd"
                  d="M7.5 6a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM3.751 20.105a8.251 8.251 0 0116.498 0 .75.75 0 01-.437.695A18.683 18.683 0 0112 22.5c-2.786 0-5.433-.608-7.812-1.7a.75.75 0 01-.437-.695z"
                  clip-rule="evenodd" />
              </svg>
            </div>
          </div>
          <div v-if="authStore.user" class="absolute -bottom-1 -right-1 w-6 h-6 bg-green-500 border-4 border-white rounded-full"></div>
        </div>


        <div class="mt-6 text-center">
          <h2 class="text-xl font-bold text-gray-900 tracking-tight">
            {{ authStore.user ? authStore.user.name : 'Guest User' }}
          </h2>
          <p class="text-xs font-semibold text-gray-400 mt-1 uppercase tracking-widest">
            {{ authStore.user ? authStore.user.email : 'Sistem Pemetaan' }}
          </p>
        </div>


        <div class="mt-8 w-full flex items-center justify-center gap-2">
          <div class="px-3 py-1 bg-gray-50 border border-gray-100 rounded-full text-[10px] font-bold text-gray-500 uppercase tracking-tight">
            {{ authStore.user ? 'Contributor' : 'Viewer Only' }}
          </div>
          <div v-if="authStore.user" class="px-3 py-1 bg-primary/10 border border-primary/10 rounded-full text-[10px] font-bold text-primary uppercase tracking-tight">
            Verified
          </div>
        </div>


        <div class="mt-8 w-full space-y-3">
          <div class="flex items-center justify-between p-4 bg-gray-50/50 rounded-2xl border border-gray-300">
            <span class="text-xs font-bold text-gray-400 uppercase">Kontribusi</span>
            <span class="text-sm font-bold text-gray-900">{{ authStore.user ? '12 Titik' : '0' }}</span>
          </div>
          <div class="flex items-center justify-between p-4 bg-gray-50/50 rounded-2xl border border-gray-300">
            <span class="text-xs font-bold text-gray-400 uppercase">Status</span>
            <span class="text-sm font-bold text-gray-900">{{ authStore.user ? 'Aktif' : 'Terbatas' }}</span>
          </div>
          <div v-if="authStore.user" class="flex flex-col gap-2 p-4 bg-gray-50/50 rounded-2xl border border-gray-300">
            <div class="flex items-center justify-between">
              <span class="text-xs font-bold text-gray-400 uppercase">Keamanan</span>
              <button @click="isChangingPassword = !isChangingPassword" class="text-xs font-bold text-primary hover:underline">
                {{ isChangingPassword ? 'Batal' : (authStore.user?.has_password ? 'Ubah Password' : 'Set Password') }}
              </button>
            </div>

            <div v-if="isChangingPassword" class="mt-2 space-y-3">
              <div v-if="errorMsg" class="text-xs text-red-600 bg-red-50 p-2 rounded border border-red-100">
                {{ errorMsg }}
              </div>
              <div v-if="authStore.user?.has_password">
                <label class="block text-[10px] font-bold text-gray-500 uppercase mb-1">Password Lama</label>
                <input v-model="oldPassword" type="password"
                  class="w-full px-3 py-1.5 text-sm bg-white border border-gray-200 rounded-lg focus:border-primary focus:ring-1 focus:ring-primary outline-none" />
              </div>
              <div>
                <label class="block text-[10px] font-bold text-gray-500 uppercase mb-1">Password Baru</label>
                <input v-model="newPassword" type="password"
                  class="w-full px-3 py-1.5 text-sm bg-white border border-gray-200 rounded-lg focus:border-primary focus:ring-1 focus:ring-primary outline-none" />
              </div>
              <div>
                <label class="block text-[10px] font-bold text-gray-500 uppercase mb-1">Konfirmasi Password Baru</label>
                <input v-model="confirmPassword" type="password"
                  class="w-full px-3 py-1.5 text-sm bg-white border border-gray-200 rounded-lg focus:border-primary focus:ring-1 focus:ring-primary outline-none" />
              </div>
              <button @click="handleChangePassword" :disabled="isLoading"
                class="w-full mt-2 py-2 bg-primary text-white rounded-lg font-bold text-xs hover:bg-primary-hover transition-colors disabled:opacity-50">
                {{ isLoading ? 'Menyimpan...' : 'Simpan Password' }}
              </button>
            </div>
          </div>
        </div>


        <div v-if="!authStore.user" class="mt-6 w-full text-center">
          <p class="text-[10px] text-gray-400 font-medium">Masuk untuk menyimpan data pemetaan Anda secara permanen</p>
        </div>


        <button @click="handleClose" class="mt-10 w-full py-4 bg-gray-900 text-white rounded-2xl font-bold text-sm hover:translate-y-[-2px] hover:shadow-lg active:translate-y-0 transition-all">
          {{ authStore.user ? 'Tutup' : 'Kembali' }}
        </button>
      </div>
    </div>
  </div>
</template>
