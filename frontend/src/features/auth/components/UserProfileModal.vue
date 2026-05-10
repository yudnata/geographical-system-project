<script setup lang="ts">
import { ref, computed, watch } from 'vue'


import { useAuthStore, type User } from '@/stores/auth'

import { useMapUIStore } from '@/stores/mapUI'
import { useMapPointsStore } from '@/stores/mapPoints'
import { useNotificationStore } from '@/stores/notifications'

const authStore = useAuthStore()
const uiStore = useMapUIStore()
const pointsStore = useMapPointsStore()
const notificationStore = useNotificationStore()

const fileInput = ref<HTMLInputElement | null>(null)
const previewImage = ref<string | null>(null)
const selectedFile = ref<File | null>(null)
const isPreviewOpen = ref(false)

const realContributionCount = computed(() => {
  if (!authStore.user) return 0
  return pointsStore.points.filter(p => p.owner_id === authStore.user?.id).length
})

const handleClose = () => {
  uiStore.closeProfileModal()
  isChangingPassword.value = false
  oldPassword.value = ''
  newPassword.value = ''
  confirmPassword.value = ''
  errorMsg.value = ''
}


const isChangingPassword = ref(false)
const isEditingProfile = ref(false)
const fullName = ref('')
const avatarURL = ref('')
const oldPassword = ref('')
const newPassword = ref('')
const confirmPassword = ref('')
const isLoading = ref(false)
const errorMsg = ref('')

watch(() => authStore.user, (u: User | null) => {

  if (u) {
    fullName.value = u.full_name || ''
    avatarURL.value = u.avatar_url || ''
  }
}, { immediate: true })


const handleFileSelect = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files[0]) {
    const file = target.files[0]
    selectedFile.value = file
    previewImage.value = URL.createObjectURL(file)
    isPreviewOpen.value = true
  }
}

const triggerFileInput = () => {
  fileInput.value?.click()
}


const handleUpdateName = async () => {
  if (!fullName.value) {
    errorMsg.value = 'Nama lengkap wajib diisi.'
    return
  }

  isLoading.value = true
  try {
    const res = await fetch(`${import.meta.env.VITE_API_URL || 'http://localhost:8080/api'}/auth/profile`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${authStore.token}`,
      },
      body: JSON.stringify({
        full_name: fullName.value,
        phone: authStore.user?.phone || '',
        avatar_url: authStore.user?.avatar_url || '',
      }),
    })
    const json = await res.json()
    if (json.success) {
      notificationStore.success('Nama berhasil diperbarui!')
      authStore.user = json.data
      isEditingProfile.value = false
    } else {
      errorMsg.value = json.message || 'Gagal memperbarui profil.'
    }
  } catch (err: unknown) {
    console.error("Profile update error:", err)
    errorMsg.value = err instanceof Error ? err.message : 'Terjadi kesalahan jaringan.'
  } finally {
    isLoading.value = false
  }
}

const handleUploadAvatar = async () => {
  if (!selectedFile.value) return

  isLoading.value = true
  errorMsg.value = ''
  try {
    const formData = new FormData()
    formData.append('avatar', selectedFile.value)

    const uploadRes = await fetch(`${import.meta.env.VITE_API_URL || 'http://localhost:8080/api'}/auth/avatar`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${authStore.token}`,
      },
      body: formData,
    })

    if (!uploadRes.ok) {
      let errMsg = 'Gagal mengunggah foto'
      try {
        const errorJson = await uploadRes.json()
        errMsg = errorJson.message || errMsg
      } catch { }
      throw new Error(errMsg)
    }

    const uploadJson = await uploadRes.json()
    if (uploadJson.success) {
      const finalAvatarURL = uploadJson.data.url
      const res = await fetch(`${import.meta.env.VITE_API_URL || 'http://localhost:8080/api'}/auth/profile`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${authStore.token}`,
        },
        body: JSON.stringify({
          full_name: authStore.user?.full_name || '',
          phone: authStore.user?.phone || '',
          avatar_url: finalAvatarURL,
        }),
      })
      const json = await res.json()
      if (json.success) {
        notificationStore.success('Foto profil diperbarui!')
        authStore.user = json.data
        selectedFile.value = null
        previewImage.value = null
        isPreviewOpen.value = false
      } else {
        throw new Error(json.message || 'Gagal menyimpan URL avatar ke profil')
      }
    }
  } catch (err: unknown) {
    console.error("Avatar update error:", err)
    errorMsg.value = err instanceof Error ? err.message : 'Gagal mengunggah foto'
  } finally {
    isLoading.value = false
  }
}

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
  <Teleport to="body" v-if="uiStore.isProfileModalOpen">
    <div class="fixed inset-0 z-[9999] flex items-center justify-center p-4">

      <div class="absolute inset-0 bg-black/20 backdrop-blur-md transition-opacity" @click="handleClose"></div>


      <div class="relative w-full max-w-sm bg-white rounded-[2.5rem] shadow-2xl overflow-hidden p-8 border border-gray-100 animate-fade-in-up">


        <button @click="handleClose" class="absolute top-6 right-6 p-2 text-gray-400 hover:text-gray-900 transition-colors">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>

        <div class="flex flex-col items-center">
          <input type="file" ref="fileInput" class="hidden" accept="image/*" @change="handleFileSelect" />

          <div class="relative group cursor-pointer" @click="triggerFileInput">
            <div class="w-24 h-24 rounded-full overflow-hidden bg-gray-50 border border-gray-100 flex items-center justify-center shadow-inner relative">
              <img v-if="previewImage || authStore.user?.avatar_url" :src="previewImage || authStore.user?.avatar_url"
                class="w-full h-full object-cover transition-transform group-hover:scale-110 duration-500" />
              <div v-else class="text-gray-300">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
                  <path fill-rule="evenodd"
                    d="M7.5 6a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM3.751 20.105a8.251 8.251 0 0116.498 0 .75.75 0 01-.437.695A18.683 18.683 0 0112 22.5c-2.786 0-5.433-.608-7.812-1.7a.75.75 0 01-.437-.695z"
                    clip-rule="evenodd" />
                </svg>
              </div>

              <!-- Edit Overlay -->
              <div class="absolute inset-0 bg-black/40 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-6 h-6 text-white">
                  <path stroke-linecap="round" stroke-linejoin="round"
                    d="M6.827 6.175A2.31 2.31 0 0 1 5.186 7.23c-.38.054-.757.112-1.134.175C2.999 7.58 2.25 8.507 2.25 9.574V18a2.25 2.25 0 0 0 2.25 2.25h15a2.25 2.25 0 0 0 2.25-2.25V9.574c0-1.067-.75-1.994-1.802-2.169a47.865 47.865 0 0 0-1.134-.175 2.31 2.31 0 0 1-1.64-1.055l-.822-1.316a2.192 2.192 0 0 0-1.736-1.039 48.774 48.774 0 0 0-5.232 0 2.192 2.192 0 0 0-1.736 1.039l-.821 1.316Z" />
                  <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 12.75a4.5 4.5 0 1 1-9 0 4.5 4.5 0 0 1 9 0ZM18.75 10.5h.008v.008h-.008V10.5Z" />
                </svg>
              </div>
            </div>
            <div v-if="authStore.user" class="absolute -bottom-1 -right-1 w-6 h-6 bg-green-500 border-4 border-white rounded-full"></div>
          </div>



          <div class="mt-6 text-center">
            <h2 class="text-xl font-bold text-gray-900 tracking-tight">
              {{ authStore.user ? authStore.user.full_name : 'Guest User' }}
            </h2>
            <p class="text-xs font-semibold text-gray-400 mt-1 uppercase tracking-widest">
              {{ authStore.user ? authStore.user.email : 'Sistem Pemetaan' }}
            </p>
          </div>


          <div class="mt-8 w-full flex items-center justify-center gap-2">
            <div class="px-3 py-1 bg-gray-50 border border-gray-100 rounded-full text-[10px] font-bold text-gray-500 uppercase tracking-tight">
              {{ authStore.user ? authStore.user.role : 'Viewer Only' }}
            </div>
            <div v-if="authStore.user" class="px-3 py-1 bg-primary/10 border border-primary/10 rounded-full text-[10px] font-bold text-primary uppercase tracking-tight">
              Verified
            </div>
          </div>


          <div class="mt-8 w-full space-y-3">
            <div v-if="errorMsg" class="text-xs text-red-600 bg-red-50 p-3 rounded-2xl border border-red-100 mb-4 animate-shake">
              {{ errorMsg }}
            </div>

            <div class="flex items-center justify-between p-4 bg-gray-50/50 rounded-2xl border border-gray-300">
              <span class="text-xs font-bold text-gray-400 uppercase">Kontribusi</span>
              <span class="text-sm font-bold text-gray-900">{{ realContributionCount }} Titik</span>
            </div>
            <div class="flex items-center justify-between p-4 bg-gray-50/50 rounded-2xl border border-gray-300">
              <span class="text-xs font-bold text-gray-400 uppercase">Status</span>
              <span class="text-sm font-bold text-gray-900">{{ authStore.user ? 'Aktif' : 'Terbatas' }}</span>
            </div>
            <div class="flex flex-col gap-2 p-4 bg-gray-50/50 rounded-2xl border border-gray-300">
              <div class="flex items-center justify-between">
                <span class="text-xs font-bold text-gray-400 uppercase">Profil</span>
                <button @click="isEditingProfile = !isEditingProfile" class="text-xs font-bold text-primary hover:underline">
                  {{ isEditingProfile ? 'Batal' : 'Edit Nama' }}
                </button>
              </div>

              <div v-if="isEditingProfile" class="mt-2 space-y-3">
                <div>
                  <label class="block text-[10px] font-bold text-gray-500 uppercase mb-1">Nama Lengkap</label>
                  <input v-model="fullName" type="text"
                    class="w-full px-3 py-1.5 text-sm bg-white border border-gray-200 rounded-lg focus:border-primary focus:ring-1 focus:ring-primary outline-none" />
                </div>
                <button @click="handleUpdateName" :disabled="isLoading"
                  class="w-full mt-2 py-2 bg-primary text-white rounded-lg font-bold text-xs hover:bg-primary-hover transition-colors disabled:opacity-50 shadow-md shadow-primary/20">
                  {{ isLoading ? 'Menyimpan...' : 'Simpan Nama Baru' }}
                </button>
              </div>

            </div>

            <div v-if="authStore.user" class="flex flex-col gap-2 p-4 bg-gray-50/50 rounded-2xl border border-gray-300">
              <div class="flex items-center justify-between">
                <span class="text-xs font-bold text-gray-400 uppercase">Keamanan</span>
                <button @click="isChangingPassword = !isChangingPassword" class="text-xs font-bold text-primary hover:underline">
                  {{ isChangingPassword ? 'Batal' : (authStore.user?.has_password ? 'Ubah Password' : 'Set Password') }}
                </button>
              </div>

              <div v-if="isChangingPassword" class="mt-2 space-y-3">
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


          <button @click="handleClose"
            class="mt-10 w-full py-4 bg-primary text-white rounded-2xl font-bold text-sm hover:translate-y-[-2px] hover:shadow-lg active:translate-y-0 transition-all shadow-lg shadow-primary/20">
            {{ authStore.user ? 'Tutup' : 'Kembali' }}
          </button>

        </div>
      </div>
    </div>

    <!-- Image Preview Modal -->
    <div v-if="isPreviewOpen && previewImage" class="fixed inset-0 z-[10000] flex items-center justify-center p-6 bg-black/60 backdrop-blur-sm animate-fade-in">
      <div class="bg-white rounded-[2.5rem] p-8 w-full max-w-sm flex flex-col items-center animate-fade-in-up shadow-2xl">
        <h3 class="text-lg font-bold text-gray-900 mb-6">Pratinjau Foto</h3>

        <div v-if="errorMsg" class="w-full text-xs text-red-600 bg-red-50 p-3 rounded-2xl border border-red-100 mb-6 animate-shake">
          {{ errorMsg }}
        </div>

        <div class="w-48 h-48 rounded-full overflow-hidden border-4 border-gray-100 shadow-xl mb-8">
          <img :src="previewImage" class="w-full h-full object-cover" />
        </div>

        <div class="flex gap-3 w-full">
          <button @click="isPreviewOpen = false; errorMsg = ''" class="flex-1 py-3 bg-gray-100 text-gray-600 rounded-xl font-bold text-xs hover:bg-gray-200 transition-colors">
            Tutup Pratinjau
          </button>

          <button @click="handleUploadAvatar" :disabled="isLoading"
            class="flex-1 py-3 bg-primary text-white rounded-xl font-bold text-xs hover:bg-primary/90 transition-colors shadow-lg shadow-primary/20">
            {{ isLoading ? 'Menyimpan...' : 'Gunakan Foto' }}
          </button>
        </div>
      </div>
    </div>
  </Teleport>
</template>


<style scoped>
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px) scale(0.95);
  }

  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.animate-fade-in-up {
  animation: fadeInUp 0.4s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}
</style>
