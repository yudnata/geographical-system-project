<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useMapPointsStore, type GeoPoint } from '@/stores/mapPoints'

const route = useRoute()
const authStore = useAuthStore()
const store = useMapPointsStore()

const handleSelectPoint = (point: GeoPoint) => {
  store.flyTo(point.latitude, point.longitude)
  store.searchQuery = ''
}

const pageTitle = computed(() => {
  if (route.path === '/') return 'Peta Publik'
  if (route.path === '/dashboard') return 'Peta Author'
  if (route.path === '/tabular') return 'Tabel Data Master'
  return 'Sistem Informasi Geografis'
})

const isDashboard = computed(() => route.path === '/dashboard')

const avatarError = ref(false)
watch(() => authStore.user, () => {
  avatarError.value = false
})
</script>

<template>
  <header :class="[
    'hidden md:flex absolute top-4 right-4 bg-white/80 backdrop-blur-md border border-gray-100 h-16 items-center px-6 shrink-0 shadow-[0_8px_32px_rgba(0,0,0,0.12)] z-[1000] justify-between rounded-2xl transition-[left] duration-500 ease-in-out',
    store.isSidebarExpanded ? 'left-[288px]' : 'left-24'
  ]">
    <div class="flex items-center gap-4 shrink-0">
      <h2 class="text-lg font-bold text-gray-800 whitespace-nowrap">{{ pageTitle }}</h2>
    </div>

    <div class="flex items-center gap-2.5 flex-1 px-6">
      <!-- Search Bar -->
      <div class="flex-1 max-w-sm relative group">
        <div class="absolute inset-y-0 left-3 flex items-center text-gray-500 group-focus-within:text-primary transition-colors">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
          </svg>
        </div>
        <input v-model="store.searchQuery" type="text" placeholder="Cari lokasi bangunan..."
          class="w-full h-9 pl-9 pr-4 bg-white border border-gray-300 rounded-xl focus:border-primary focus:ring-4 focus:ring-primary/10 transition-all text-xs font-semibold text-gray-800 placeholder:text-gray-400 outline-none" />

        <!-- Search Results Dropdown -->
        <div v-if="store.searchQuery && store.filteredPoints.length > 0"
          class="absolute top-full left-0 right-0 mt-2 bg-white/90 backdrop-blur-xl border border-white/20 shadow-[0_8px_32px_rgba(0,0,0,0.12)] rounded-2xl overflow-hidden z-[2000]">
          <div class="max-h-64 overflow-y-auto">
            <div v-for="point in store.filteredPoints.slice(0, 5)" :key="point.id" @click="handleSelectPoint(point)"
              class="px-4 py-3 hover:bg-primary/5 cursor-pointer flex items-center gap-3 transition-colors border-b border-gray-50 last:border-0">
              <div class="w-7 h-7 rounded-lg bg-primary/10 text-primary flex items-center justify-center shrink-0">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-3.5 h-3.5">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z" />
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1 1 15 0Z" />
                </svg>
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-xs font-extrabold text-gray-900 truncate tracking-tight">{{ point.name }}</p>
                <p class="text-[10px] text-gray-500 truncate font-medium">{{ point.address }}</p>
              </div>
            </div>
          </div>
          <div v-if="store.filteredPoints.length > 5" class="px-4 py-2 bg-gray-50/50 text-[10px] font-bold text-gray-400 uppercase tracking-widest text-center border-t border-gray-50">
            +{{ store.filteredPoints.length - 5 }} lainnya
          </div>
        </div>
      </div>

      <!-- Category Filter -->
      <div class="relative shrink-0">
        <select v-model="store.filterTypeId"
          class="h-9 bg-white border border-gray-300 pl-3 pr-7 rounded-xl text-[11px] font-bold text-gray-700 appearance-none focus:ring-4 focus:ring-primary/10 focus:border-primary outline-none transition-all cursor-pointer uppercase tracking-tight">
          <option :value="null">SEMUA KATEGORI</option>
          <option v-for="type in store.objectTypes" :key="type.id" :value="type.id">
            {{ type.name.toUpperCase() }}
          </option>
        </select>
        <span class="absolute inset-y-0 right-2 flex items-center pointer-events-none text-gray-500">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-3 h-3">
            <path stroke-linecap="round" stroke-linejoin="round" d="m19.5 8.25-7.5 7.5-7.5-7.5" />
          </svg>
        </span>
      </div>

      <!-- Dashboard-only: Filter & Edit Mode -->
      <template v-if="isDashboard">
        <div class="h-5 w-px bg-gray-300 shrink-0"></div>

        <!-- Filter Toggle -->
        <div class="flex bg-gray-200/80 rounded-xl p-0.5 shrink-0 h-9 items-center">
          <button @click="store.filterMyPoints = false" :class="!store.filterMyPoints ? 'bg-white shadow-md text-primary font-extrabold' : 'text-gray-500 hover:text-gray-700'"
            class="px-3 h-7 text-[11px] font-bold rounded-lg transition-all whitespace-nowrap">Semua</button>
          <button @click="store.filterMyPoints = true" :class="store.filterMyPoints ? 'bg-white shadow-md text-primary font-extrabold' : 'text-gray-500 hover:text-gray-700'"
            class="px-3 h-7 text-[11px] font-bold rounded-lg transition-all whitespace-nowrap">Milik Saya</button>
        </div>

        <!-- Edit Mode Toggle -->
        <div class="flex items-center gap-2.5 h-9 px-3 rounded-xl border border-gray-300 bg-white cursor-pointer shrink-0" @click="store.toggleEditMode()">
          <span class="text-[11px] font-bold uppercase tracking-wider select-none" :class="store.isEditMode ? 'text-emerald-700' : 'text-gray-500'">Edit</span>
          <div :class="[
            'relative w-11 h-6 rounded-full transition-colors duration-300 shadow-inner',
            store.isEditMode ? 'bg-emerald-500' : 'bg-gray-300'
          ]">
            <div :class="[
              'absolute top-0.5 w-5 h-5 bg-white rounded-full shadow-md transition-all duration-300',
              store.isEditMode ? 'left-[22px]' : 'left-0.5'
            ]"></div>
          </div>
        </div>
      </template>
    </div>

    <div class="flex items-center gap-2">
      <div v-if="authStore.user" class="flex items-center gap-3 px-3 py-1.5 rounded-full hover:bg-gray-50 transition-colors cursor-pointer border border-transparent hover:border-gray-100">
        <div class="text-right hidden sm:block">
          <p class="text-xs font-bold text-gray-900 leading-tight">{{ authStore.user.name }}</p>
          <p class="text-[10px] text-gray-400 font-medium">{{ authStore.user.email || 'Administrator' }}</p>
        </div>

        <div class="relative">
          <img v-if="authStore.user.avatar_url && !avatarError" :src="authStore.user.avatar_url" class="w-9 h-9 rounded-full border-2 border-white shadow-sm object-cover"
            @error="avatarError = true" />
          <img v-else :src="authStore.defaultAvatarUrl" class="w-9 h-9 rounded-full border-2 border-white shadow-sm object-cover" />
          <div class="absolute -bottom-0.5 -right-0.5 w-3 h-3 bg-green-500 border-2 border-white rounded-full"></div>
        </div>
      </div>

      <div v-else class="flex items-center gap-2 px-2 py-1 rounded-full hover:bg-gray-50 transition-colors cursor-pointer border border-transparent hover:border-gray-100">
        <p class="text-xs font-bold text-gray-500 hidden sm:block">Guest</p>
        <img :src="authStore.defaultAvatarUrl" class="w-8 h-8 rounded-full border-2 border-white shadow-sm object-cover" />
      </div>


      <RouterLink v-if="!authStore.user" to="/login"
        class="px-5 py-2 bg-primary text-white text-[10px] font-bold rounded-full hover:bg-primary/90 transition-all shadow-sm shadow-primary/20 tracking-tighter uppercase">
        Login
      </RouterLink>
    </div>
  </header>
</template>
