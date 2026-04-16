<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useMapPointsStore } from '@/stores/mapPoints'
import UserProfileModal from '@/components/UserProfileModal.vue'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const store = useMapPointsStore()

const isProfileModalOpen = ref(false)

const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}

const navItems = [
  {
    path: '/',
    name: 'Peta Publik',
    icon: 'M9 6.75V15m6-6v8.25m.503 3.498l4.875-2.437c.381-.19.622-.58.622-1.006V4.82c0-.836-.88-1.38-1.628-1.006l-3.869 1.934c-.317.159-.69.159-1.006 0L9.503 3.252a1.125 1.125 0 00-1.006 0L3.622 5.689C3.24 5.88 3 6.27 3 6.695V19.18c0 .836.88 1.38 1.628 1.006l3.869-1.934c.317-.159.69-.159 1.006 0l4.994 2.497c.317.158.69.158 1.006 0z',
  },
  {
    path: '/dashboard',
    name: 'Dashboard (Edit)',
    icon: 'M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L6.832 19.82a4.5 4.5 0 01-1.897 1.13l-2.685.8.8-2.685a4.5 4.5 0 011.13-1.897L16.863 4.487zm0 0L19.5 7.125',
  },
  {
    path: '/tabular',
    name: 'Tabel Master',
    icon: 'M3.75 12h16.5m-16.5 3.75h16.5M3.75 19.5h16.5M5.625 4.5h12.75a1.875 1.875 0 011.875 1.875v11.25a1.875 1.875 0 01-1.875 1.875H5.625a1.875 1.875 0 01-1.875-1.875V6.375c0-1.036.84-1.875 1.875-1.875z',
  },
]

const isItemActive = (path: string) => {
  return route.path === path
}
</script>

<template>
  <aside @mouseenter="store.isSidebarExpanded = true" @mouseleave="store.isSidebarExpanded = false"
    class="group absolute top-4 left-4 bottom-4 w-16 hover:w-64 bg-white/80 backdrop-blur-md flex-col hidden md:flex shrink-0 transition-all duration-500 ease-in-out z-[1001] overflow-hidden shadow-[0_8px_32px_rgba(0,0,0,0.12)] border border-gray-100 rounded-2xl">
    <div :class="['h-16 flex items-center shrink-0 border-b border-gray-50 whitespace-nowrap box-border transition-all duration-500', store.isSidebarExpanded ? 'px-6' : 'px-4']">
      <div class="w-8 h-8 rounded-lg bg-primary/10 text-primary flex items-center justify-center shrink-0">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-5 h-5">
          <path fill-rule="evenodd"
            d="M11.54 22.351l.07.04.028.016a.76.76 0 00.723 0l.028-.015.071-.041a16.975 16.975 0 001.144-.742 19.58 19.58 0 002.683-2.282c1.944-1.99 3.963-4.98 3.963-8.827a8.25 8.25 0 00-16.5 0c0 3.846 2.02 6.837 3.963 8.827a19.58 19.58 0 002.682 2.282 16.975 16.975 0 001.145.742zM12 13.5a3 3 0 100-6 3 3 0 000 6z"
            clip-rule="evenodd" />
        </svg>
      </div>
      <div class="ml-4 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
        <h1 class="text-sm font-extrabold leading-tight text-gray-900 tracking-tight">Geo System</h1>
        <p class="text-gray-400 text-[10px] font-medium mt-0.5">Sistem Pemetaan</p>
      </div>
    </div>
    <nav :class="['flex-1 py-6 space-y-2 overflow-y-auto overflow-x-hidden flex flex-col h-full transition-all duration-500', store.isSidebarExpanded ? 'px-6' : 'px-3']">
      <div class="space-y-2">
        <RouterLink v-for="item in navItems" :key="item.path" :to="item.path" :class="[
          'flex items-center gap-4 px-2.5 py-2.5 rounded-xl transition-all duration-200',
          isItemActive(item.path)
            ? 'bg-primary/10 text-primary font-bold shadow-sm shadow-primary/5'
            : 'text-gray-400 hover:bg-gray-50 hover:text-gray-600',
        ]" :title="item.name">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5 shrink-0">
            <path stroke-linecap="round" stroke-linejoin="round" :d="item.icon" />
          </svg>
          <span class="text-xs font-semibold whitespace-nowrap opacity-0 group-hover:opacity-100 transition-opacity duration-300 tracking-wide">{{ item.name }}</span>
        </RouterLink>
      </div>

      <div class="mt-auto pt-4 border-t border-gray-50 space-y-1">
        <button @click="isProfileModalOpen = true" class="flex items-center gap-4 px-2.5 py-2.5 rounded-xl transition-all duration-200 text-gray-400 hover:bg-gray-50 hover:text-gray-600 w-full"
          title="Pengaturan Profil">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5 shrink-0">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M9.594 3.94c.09-.542.56-.94 1.11-.94h2.593c.55 0 1.02.398 1.11.94l.213 1.281c.063.374.313.686.645.87.074.04.147.083.22.127.324.196.72.257 1.075.124l1.217-.456a1.125 1.125 0 0 1 1.37.49l1.296 2.247a1.125 1.125 0 0 1-.26 1.431l-1.003.827c-.293.24-.438.613-.431.992a6.759 6.759 0 0 1 0 .255c-.007.378.138.75.43.99l1.005.828c.424.35.534.954.26 1.43l-1.298 2.247a1.125 1.125 0 0 1-1.369.491l-1.217-.456c-.355-.133-.75-.072-1.076.124a6.57 6.57 0 0 1-.22.127c-.332.183-.582.495-.644.869l-.213 1.28c-.09.543-.56.941-1.11.941h-2.594c-.55 0-1.02-.398-1.11-.94l-.213-1.281c-.062-.374-.312-.686-.644-.87a6.52 6.52 0 0 1-.22-.127c-.325-.196-.72-.257-1.076-.124l-1.217.456a1.125 1.125 0 0 1-1.369-.49l-1.297-2.247a1.125 1.125 0 0 1 .26-1.431l1.004-.827c.292-.24.437-.613.43-.992a6.932 6.932 0 0 1 0-.255c.007-.378-.138-.75-.43-.99l-1.004-.828a1.125 1.125 0 0 1-.26-1.43l1.297-2.247a1.125 1.125 0 0 1 1.37-.491l1.216.456c.356.133.751.072 1.076-.124.072-.044.146-.087.22-.127.332-.184.582-.496.645-.869L9.594 3.94ZM12 15.75a3.75 3.75 0 1 0 0-7.5 3.75 3.75 0 0 0 0 7.5Z" />
          </svg>
          <span class="text-[10px] font-bold whitespace-nowrap opacity-0 group-hover:opacity-100 transition-opacity duration-300 tracking-wide uppercase">Pengaturan</span>
        </button>

        <button v-if="authStore.user" @click="handleLogout" class="flex items-center gap-4 px-2.5 py-2.5 rounded-xl transition-all duration-200 text-gray-400 hover:bg-red-50 hover:text-red-500 w-full"
          title="Keluar Aplikasi">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5 shrink-0">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M15.75 9V5.25A2.25 2.25 0 0 0 13.5 3h-6a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h6a2.25 2.25 0 0 0 2.25-2.25V15m3 0 3-3m0 0-3-3m3 3H9" />
          </svg>
          <span class="text-[10px] font-bold whitespace-nowrap opacity-0 group-hover:opacity-100 transition-opacity duration-300 tracking-wide uppercase">Logout</span>
        </button>
      </div>
    </nav>
  </aside>

  <UserProfileModal :is-open="isProfileModalOpen" @close="isProfileModalOpen = false" />
</template>
