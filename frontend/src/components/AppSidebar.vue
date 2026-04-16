<script setup lang="ts">
import { RouterLink, useRoute } from 'vue-router'

const route = useRoute()

const navItems = [
  {
    path: '/',
    name: 'Dashboard',
    icon: 'M2.25 12l8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25',
  },
  {
    path: '/add',
    name: 'Tambah Fasilitas',
    icon: 'M12 9v6m3-3H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z',
  },
  {
    path: '/list',
    name: 'Daftar Fasilitas',
    icon: 'M8.25 6.75h12M8.25 12h12m-12 5.25h12M3.75 6.75h.007v.008H3.75V6.75zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zM3.75 12h.007v.008H3.75V12zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm-.375 5.25h.007v.008H3.75v-.008zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z',
  },
]

const isItemActive = (path: string) => {
  if (path === '/') return route.path === '/'
  if (path === '/add') return route.path === '/add' || route.path.startsWith('/edit')
  return route.path === path
}
</script>

<template>
  <aside class="w-64 bg-primary text-white flex-col hidden md:flex shrink-0">
    <div class="p-6">
      <h1 class="text-xl font-bold leading-tight">SIG Faskes</h1>
      <p class="text-surface/80 text-xs mt-1">Sistem Pemetaan Bali</p>
    </div>
    <nav class="flex-1 px-4 space-y-2 mt-4">
      <RouterLink v-for="item in navItems" :key="item.path" :to="item.path" :class="[
        'flex items-center gap-3 px-4 py-2.5 rounded transition',
        isItemActive(item.path)
          ? 'bg-white/20 text-white font-bold'
          : 'text-white/80 hover:bg-white/10 hover:text-white',
      ]">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
          <path stroke-linecap="round" stroke-linejoin="round" :d="item.icon" />
        </svg>
        <span class="font-medium">{{ item.name }}</span>
      </RouterLink>
    </nav>
  </aside>
</template>
