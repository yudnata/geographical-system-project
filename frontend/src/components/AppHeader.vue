<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useMapPointsStore } from '@/stores/mapPoints'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const store = useMapPointsStore()

const pageTitle = computed(() => {
  if (route.path === '/') return 'Peta Publik'
  if (route.path === '/dashboard') return 'Dashboard Private'
  if (route.path === '/tabular') return 'Tabel Data Master'
  return 'Sistem Informasi Geografis'
})

const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<template>
  <header :class="[
    'hidden md:flex absolute top-4 right-4 bg-white/80 backdrop-blur-md border border-gray-100 h-16 items-center px-6 shrink-0 shadow-[0_8px_32px_rgba(0,0,0,0.12)] z-[1000] justify-between rounded-2xl transition-[left] duration-500 ease-in-out',
    store.isSidebarExpanded ? 'left-[288px]' : 'left-24'
  ]">
    <div class="flex items-center gap-4">
      <h2 class="text-lg font-bold text-gray-800">{{ pageTitle }}</h2>
    </div>

    <div class="flex items-center gap-4 flex-1 max-w-2xl px-8">
      <div class="relative flex-1 group">
        <span class="absolute inset-y-0 left-3 flex items-center text-gray-500 group-focus-within:text-primary transition-colors">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
            <path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
          </svg>
        </span>
        <input v-model="store.searchQuery" type="text" placeholder="Cari lokasi atau kontributor..."
          class="w-full bg-white border border-gray-200 py-2.5 pl-10 pr-4 rounded-xl text-xs focus:ring-4 focus:ring-primary/10 focus:border-primary outline-none transition-all placeholder:text-gray-500 font-bold text-gray-800" />
      </div>

      <div class="relative min-w-[150px]">
        <select v-model="store.filterTypeId"
          class="w-full bg-white border border-gray-200 py-2.5 px-3 rounded-xl text-xs font-extrabold text-gray-700 appearance-none focus:ring-4 focus:ring-primary/10 focus:border-primary outline-none transition-all cursor-pointer uppercase tracking-tight">
          <option :value="null">SEMUA KATEGORI</option>
          <option v-for="type in store.objectTypes" :key="type.id" :value="type.id">
            {{ type.name.toUpperCase() }}
          </option>
        </select>
        <span class="absolute inset-y-0 right-3 flex items-center pointer-events-none text-gray-600">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="m19.5 8.25-7.5 7.5-7.5-7.5" />
          </svg>
        </span>
      </div>
    </div>

    <div class="flex items-center gap-2">
      <div v-if="authStore.user" class="flex items-center gap-3 px-3 py-1.5 rounded-full hover:bg-gray-50 transition-colors cursor-pointer border border-transparent hover:border-gray-100">
        <div class="text-right hidden sm:block">
          <p class="text-xs font-bold text-gray-900 leading-tight">{{ authStore.user.name }}</p>
          <p class="text-[10px] text-gray-400 font-medium">{{ authStore.user.email || 'Administrator' }}</p>
        </div>

        <div class="relative">
          <img v-if="authStore.user.avatar_url" :src="authStore.user.avatar_url" class="w-9 h-9 rounded-full border-2 border-white shadow-sm object-cover" />
          <img v-else :src="authStore.defaultAvatarUrl" class="w-9 h-9 rounded-full border-2 border-white shadow-sm object-cover" />
          <div class="absolute -bottom-0.5 -right-0.5 w-3 h-3 bg-green-500 border-2 border-white rounded-full"></div>
        </div>
      </div>

      <div v-else class="flex items-center gap-2 px-3 py-2.5 bg-white border border-gray-200 rounded-xl">
        <div class="text-right hidden sm:block">
          <p class="text-xs font-bold text-gray-700 leading-tight tracking-tight uppercase">Guest</p>
        </div>
        <div class="flex items-center justify-center">
          <img :src="authStore.defaultAvatarUrl" class="w-9 h-9 rounded-full border-2 border-white shadow-sm object-cover" />
        </div>
      </div>

      <div class="h-8 w-px bg-gray-100"></div>

      <button v-if="authStore.user" @click="handleLogout" class="p-2 text-gray-400 hover:text-red-500 hover:bg-red-50 rounded-lg transition-all" title="Keluar">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
          <path stroke-linecap="round" stroke-linejoin="round"
            d="M15.75 9V5.25A2.25 2.25 0 0 0 13.5 3h-6a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h6a2.25 2.25 0 0 0 2.25-2.25V15m3 0 3-3m0 0-3-3m3 3H9" />
        </svg>
      </button>
      <RouterLink v-else to="/login"
        class="px-5 py-2 bg-primary text-white text-[10px] font-bold rounded-full hover:bg-primary/90 transition-all shadow-sm shadow-primary/20 tracking-tighter uppercase">
        Login
      </RouterLink>
    </div>
  </header>
</template>
