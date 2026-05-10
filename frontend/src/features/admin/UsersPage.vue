<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useNotificationStore } from '@/stores/notifications'

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'
const authStore = useAuthStore()
const notify = useNotificationStore()

interface User {
  id: string
  name: string
  email: string
  role: string
  avatar_url?: string
  phone?: string
  institution?: string
  is_profile_completed: boolean
  created_at: string
}

const users = ref<User[]>([])
const isLoading = ref(false)
const searchQuery = ref('')

const filteredUsers = computed(() => {
  const q = searchQuery.value.toLowerCase()
  if (!q) return users.value
  return users.value.filter(u =>
    u.name.toLowerCase().includes(q) || u.email.toLowerCase().includes(q)
  )
})

const fetchUsers = async () => {
  isLoading.value = true
  try {
    const res = await fetch(`${API_URL}/auth/users`, {
      headers: { Authorization: `Bearer ${authStore.token}` }
    })
    const data = await res.json()
    if (data.success) users.value = data.data
    else notify.error(data.message || 'Gagal memuat pengguna')
  } catch {
    notify.error('Gagal terhubung ke server')
  } finally {
    isLoading.value = false
  }
}

const roleBadge = (role: string) => {
  if (role === 'admin') return 'bg-amber-500/15 text-amber-400 border-amber-500/30'
  return 'bg-blue-500/15 text-blue-400 border-blue-500/30'
}

const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })
}

onMounted(fetchUsers)
</script>


<template>
  <div class="p-8 min-h-screen bg-[#0f1117]">
    <!-- Header -->
    <div class="flex items-center justify-between mb-8">
      <div>
        <h1 class="text-2xl font-bold text-white">Manajemen Pengguna</h1>
        <p class="text-white/40 text-sm mt-1">Daftar seluruh akun yang terdaftar di sistem</p>
      </div>
      <div class="text-right">
        <p class="text-2xl font-bold text-white">{{ users.length }}</p>
        <p class="text-white/30 text-xs">Total Pengguna</p>
      </div>
    </div>

    <!-- Search -->
    <div class="relative mb-6">
      <svg xmlns="http://www.w3.org/2000/svg" class="absolute left-3.5 top-1/2 -translate-y-1/2 w-4 h-4 text-white/30" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
      </svg>
      <input v-model="searchQuery" type="text" placeholder="Cari nama atau email..."
        class="w-full bg-[#161b27] border border-white/10 rounded-xl pl-10 pr-4 py-2.5 text-sm text-white placeholder-white/20 focus:outline-none focus:border-amber-500/50 transition-colors" />
    </div>

    <!-- Loading -->
    <div v-if="isLoading" class="flex items-center justify-center py-20 text-white/40">
      <svg class="animate-spin w-6 h-6 mr-3" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z" />
      </svg>
      Memuat pengguna...
    </div>

    <!-- Table -->
    <div v-else class="bg-[#161b27] border border-white/5 rounded-2xl overflow-hidden">
      <table class="w-full">
        <thead>
          <tr class="border-b border-white/5">
            <th class="text-left px-5 py-3.5 text-xs font-semibold text-white/30 uppercase tracking-wider">Pengguna</th>
            <th class="text-left px-5 py-3.5 text-xs font-semibold text-white/30 uppercase tracking-wider">Email</th>
            <th class="text-left px-5 py-3.5 text-xs font-semibold text-white/30 uppercase tracking-wider">Peran</th>
            <th class="text-left px-5 py-3.5 text-xs font-semibold text-white/30 uppercase tracking-wider">Institusi</th>
            <th class="text-left px-5 py-3.5 text-xs font-semibold text-white/30 uppercase tracking-wider">Bergabung</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-white/5">
          <tr v-if="filteredUsers.length === 0">
            <td colspan="5" class="text-center py-12 text-white/30 text-sm">Tidak ada pengguna ditemukan</td>
          </tr>
          <tr v-for="user in filteredUsers" :key="user.id" class="hover:bg-white/2 transition-colors">
            <td class="px-5 py-4">
              <div class="flex items-center gap-3">
                <div class="w-8 h-8 rounded-full bg-gradient-to-br from-amber-400 to-orange-500 flex items-center justify-center text-xs font-bold text-white flex-shrink-0">
                  {{ user.name?.charAt(0).toUpperCase() || '?' }}
                </div>
                <div>
                  <p class="text-sm font-medium text-white">{{ user.name }}</p>
                  <p v-if="!user.is_profile_completed" class="text-[10px] text-yellow-500/70">Profil belum lengkap</p>
                </div>
              </div>
            </td>
            <td class="px-5 py-4 text-sm text-white/50">{{ user.email }}</td>
            <td class="px-5 py-4">
              <span :class="['text-[11px] font-semibold px-2.5 py-1 rounded-full border capitalize', roleBadge(user.role)]">
                {{ user.role }}
              </span>
            </td>
            <td class="px-5 py-4 text-sm text-white/50">{{ user.institution || '-' }}</td>
            <td class="px-5 py-4 text-sm text-white/40">{{ formatDate(user.created_at) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
