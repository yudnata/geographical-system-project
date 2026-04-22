<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useMapPointsStore } from '@/stores/mapPoints'

const route = useRoute()
const authStore = useAuthStore()
const store = useMapPointsStore()

const navItems = computed(() => {
  const all = [
    {
      path: '/',
      name: 'Peta Publik',
      icon: 'M9 6.75V15m6-6v8.25m.503 3.498l4.875-2.437c.381-.19.622-.58.622-1.006V4.82c0-.836-.88-1.38-1.628-1.006l-3.869 1.934c-.317.159-.69.159-1.006 0L9.503 3.252a1.125 1.125 0 00-1.006 0L3.622 5.689C3.24 5.88 3 6.27 3 6.695V19.18c0 .836.88 1.38 1.628 1.006l3.869-1.934c.317-.159.69-.159 1.006 0l4.994 2.497c.317.158.69.158 1.006 0z',
      requiresAuth: false,
    },
    {
      path: '/dashboard',
      name: 'Editor Peta',
      icon: 'M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L6.832 19.82a4.5 4.5 0 01-1.897 1.13l-2.685.8.8-2.685a4.5 4.5 0 011.13-1.897L16.863 4.487zm0 0L19.5 7.125',
      requiresAuth: true,
    },
    {
      path: '/tabular',
      name: 'Tabel Master',
      icon: 'M3.75 12h16.5m-16.5 3.75h16.5M3.75 19.5h16.5M5.625 4.5h12.75a1.875 1.875 0 011.875 1.875v11.25a1.875 1.875 0 01-1.875 1.875H5.625a1.875 1.875 0 01-1.875 1.875H5.625a1.875 1.875 0 01-1.875-1.875V6.375c0-1.036.84-1.875 1.875-1.875z',
      requiresAuth: true,
    },
  ]
  return all.filter(item => !item.requiresAuth || !!authStore.user)
})

const isItemActive = (path: string) => {
  return route.path === path
}
</script>

<template>
  <nav :class="['flex-1 py-6 space-y-2 overflow-y-auto overflow-x-hidden flex flex-col transition-[padding] duration-500', store.isSidebarExpanded ? 'px-6' : 'px-3']">
    <div class="space-y-2">
      <RouterLink v-for="item in navItems" :key="item.path" :to="item.path" :class="[
        'flex items-center gap-4 px-2.5 py-2.5 rounded-xl transition-colors duration-200',
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
  </nav>
</template>
