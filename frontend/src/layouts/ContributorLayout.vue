<script setup lang="ts">
import { RouterView, RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const router = useRouter()

const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<template>
  <div class="flex h-screen w-full bg-[#f8fafc] font-sans text-slate-800 antialiased overflow-hidden">
    <!-- Contributor Sidebar -->
    <aside class="w-64 flex-shrink-0 flex flex-col bg-white border-r border-slate-200">
      <!-- Logo -->
      <div class="flex items-center gap-3 px-6 py-5 border-b border-slate-100">
        <div class="w-8 h-8 rounded-lg bg-indigo-600 flex items-center justify-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-white" viewBox="0 0 24 24" fill="currentColor">
            <path d="M12 2C8.13 2 5 5.13 5 9c0 5.25 7 13 7 13s7-7.75 7-13c0-3.87-3.13-7-7-7zm0 9.5c-1.38 0-2.5-1.12-2.5-2.5s1.12-2.5 2.5-2.5 2.5 1.12 2.5 2.5-1.12 2.5-2.5 2.5z" />
          </svg>
        </div>
        <div>
          <h1 class="text-sm font-bold text-slate-900">GIS Bali</h1>
          <p class="text-[10px] text-indigo-600 font-medium uppercase tracking-wider">Kontributor</p>
        </div>
      </div>

      <!-- Nav -->
      <nav class="flex-1 px-3 py-4 space-y-1 overflow-y-auto">
        <p class="text-[10px] font-semibold text-slate-400 uppercase tracking-widest px-3 mb-2">Workspace</p>

        <RouterLink to="/contributor/dashboard" class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm text-slate-600 hover:text-slate-900 hover:bg-slate-50 transition-all"
          active-class="!text-indigo-600 !bg-indigo-50 !font-semibold">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="m2.25 12 8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25" />
          </svg>
          Kontribusi Saya
        </RouterLink>

        <RouterLink to="/contributor/map" class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm text-slate-600 hover:text-slate-900 hover:bg-slate-50 transition-all"
          active-class="!text-indigo-600 !bg-indigo-50 !font-semibold">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M9 6.75V15m6-10.5v.008H15V4.5Zm0 3v.008H15V7.5Zm0 3v.008H15v-.008Zm0 3v.008H15V13.5Zm-6-6v.008H9V7.5Zm0 3v.008H9v-.008Zm0 3v.008H9V13.5Zm0 3v.008H9v-.008Z" />
            <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 12.75V12A9.75 9.75 0 0 1 12 2.25h.75A9.75 9.75 0 0 1 22.5 12v.75A9.75 9.75 0 0 1 12.75 22.5H12A9.75 9.75 0 0 1 2.25 12.75Z" />
          </svg>
          Peta Kontribusi
        </RouterLink>


        <div class="border-t border-slate-100 my-3"></div>
        <p class="text-[10px] font-semibold text-slate-400 uppercase tracking-widest px-3 mb-2">Aplikasi</p>

        <RouterLink to="/" class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm text-slate-600 hover:text-slate-900 hover:bg-slate-50 transition-all">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M9 6.75V15m6-10.5v.008H15V4.5Zm0 3v.008H15V7.5Zm0 3v.008H15v-.008Zm0 3v.008H15V13.5Zm-6-6v.008H9V7.5Zm0 3v.008H9v-.008Zm0 3v.008H9V13.5Zm0 3v.008H9v-.008Z" />
            <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 12.75V12A9.75 9.75 0 0 1 12 2.25h.75A9.75 9.75 0 0 1 22.5 12v.75A9.75 9.75 0 0 1 12.75 22.5H12A9.75 9.75 0 0 1 2.25 12.75Z" />
          </svg>
          Peta Utama
        </RouterLink>
      </nav>

      <!-- User info -->
      <div class="px-3 py-4 border-t border-slate-100">
        <div class="flex items-center gap-3 p-3 rounded-lg bg-slate-50">
          <div class="w-8 h-8 rounded-full bg-indigo-600 flex items-center justify-center text-xs font-bold text-white flex-shrink-0">
            {{ authStore.user?.full_name?.charAt(0).toUpperCase() || 'K' }}
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-xs font-semibold text-slate-900 truncate">{{ authStore.user?.full_name || 'Kontributor' }}</p>
            <p class="text-[10px] text-slate-500 capitalize">{{ authStore.user?.role }}</p>
          </div>
          <button @click="handleLogout" title="Logout" class="text-slate-400 hover:text-red-500 transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round"
                d="M8.25 9V5.25A2.25 2.25 0 0 1 10.5 3h6a2.25 2.25 0 0 1 2.25 2.25v13.5A2.25 2.25 0 0 1 16.5 21h-6a2.25 2.25 0 0 1-2.25-2.25V15m-3 0-3-3m0 0 3-3m-3 3H15" />
            </svg>
          </button>
        </div>
      </div>
    </aside>

    <!-- Main Content -->
    <main class="flex-1 overflow-y-auto">
      <RouterView />
    </main>
  </div>
</template>
