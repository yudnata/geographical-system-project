<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useMapPointsStore } from '@/stores/mapPoints'
import UserProfileModal from '@/components/Modals/UserProfileModal.vue'

const router = useRouter()
const authStore = useAuthStore()
const store = useMapPointsStore()

const isProfileModalOpen = ref(false)

const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<template>
  <div :class="['mt-auto pb-10 pt-4 border-t border-gray-200 space-y-1 transition-[padding] duration-500', store.isSidebarExpanded ? 'px-6' : 'px-3']">
    <button @click="isProfileModalOpen = true" class="flex items-center gap-4 px-2.5 py-2.5 rounded-xl transition-colors duration-200 text-gray-400 hover:bg-gray-50 hover:text-gray-600 w-full"
      title="Pengaturan Profil">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5 shrink-0">
        <path stroke-linecap="round" stroke-linejoin="round"
          d="M9.594 3.94c.09-.542.56-.94 1.11-.94h2.593c.55 0 1.02.398 1.11.94l.213 1.281c.063.374.313.686.645.87.074.04.147.083.22.127.324.196.72.257 1.075.124l1.217-.456a1.125 1.125 0 0 1 1.37.49l1.296 2.247a1.125 1.125 0 0 1-.26 1.431l-1.003.827c-.293.24-.438.613-.431.992a6.759 6.759 0 0 1 0 .255c-.007.378.138.75.43.99l1.005.828c.424.35.534.954.26 1.43l-1.298 2.247a1.125 1.125 0 0 1-1.369.491l-1.217-.456c-.355-.133-.75-.072-1.076.124a6.57 6.57 0 0 1-.22.127c-.332.183-.582.495-.644.869l-.213 1.28c-.09.543-.56.941-1.11.941h-2.594c-.55 0-1.02-.398-1.11-.94l-.213-1.281c-.062-.374-.312-.686-.644-.87a6.52 6.52 0 0 1-.22-.127c-.325-.196-.72-.257-1.076-.124l-1.217.456a1.125 1.125 0 0 1-1.369-.49l-1.297-2.247a1.125 1.125 0 0 1 .26-1.431l1.004-.827c.292-.24.437-.613.43-.992a6.932 6.932 0 0 1 0-.255c.007-.378-.138-.75-.43-.99l-1.004-.828a1.125 1.125 0 0 1-.26-1.43l1.297-2.247a1.125 1.125 0 0 1 1.37-.491l1.216.456c.356.133.751.072 1.076-.124.072-.044.146-.087.22-.127.332-.184.582-.496.645-.869L9.594 3.94ZM12 15.75a3.75 3.75 0 1 0 0-7.5 3.75 3.75 0 0 0 0 7.5Z" />
      </svg>
      <span class="text-[10px] font-bold whitespace-nowrap opacity-0 group-hover:opacity-100 transition-opacity duration-300 tracking-wide uppercase">Pengaturan</span>
    </button>

    <button v-if="authStore.user" @click="handleLogout" class="flex items-center gap-4 px-2.5 py-2.5 rounded-xl transition-colors duration-200 text-gray-400 hover:bg-red-50 hover:text-red-500 w-full"
      title="Keluar Aplikasi">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5 shrink-0">
        <path stroke-linecap="round" stroke-linejoin="round"
          d="M15.75 9V5.25A2.25 2.25 0 0 0 13.5 3h-6a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h6a2.25 2.25 0 0 0 2.25-2.25V15m3 0 3-3m0 0-3-3m3 3H9" />
      </svg>
      <span class="text-[10px] font-bold whitespace-nowrap opacity-0 group-hover:opacity-100 transition-opacity duration-300 tracking-wide uppercase">Logout</span>
    </button>

    <div class="pt-4 px-2.5 opacity-0 group-hover:opacity-100 transition-opacity duration-500 overflow-hidden">
      <p class="text-[9px] font-bold text-gray-400 whitespace-nowrap tracking-widest uppercase">
        &copy; {{ new Date().getFullYear() }} by Yudhi Nata
      </p>
    </div>
  </div>

  <UserProfileModal :is-open="isProfileModalOpen" @close="isProfileModalOpen = false" />
</template>
