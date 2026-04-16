<script setup lang="ts">
import { RouterView, RouterLink, useRoute } from 'vue-router'
import { useFacilityStore } from '@/stores/facility'
import { onMounted, computed } from 'vue'
import AppSidebar from '@/components/AppSidebar.vue'

const store = useFacilityStore()
onMounted(() => {
  store.initFacilities()
})

const route = useRoute()

const pageTitle = computed(() => {
  if (route.path === '/') return 'Dashboard'
  if (route.path === '/add') return 'Tambah Fasilitas Baru'
  if (route.path === '/list') return 'Daftar Fasilitas Kesehatan'
  if (route.path.startsWith('/edit')) return 'Edit Fasilitas'
  return 'Sistem Informasi Geografis'
})
</script>

<template>
  <div class="flex h-screen bg-surface font-sans text-text antialiased overflow-hidden">
    <!-- Sidebar Desktop -->
    <AppSidebar />

    <!-- Main Content -->
    <div class="flex-1 flex flex-col h-screen overflow-hidden relative">
      <!-- Mobile Header & Nav -->
      <header class="bg-primary text-white shadow-sm md:hidden p-4 flex flex-col gap-2">
        <div>
          <h1 class="text-lg font-bold">SIG Faskes Bali</h1>
        </div>
        <nav class="flex overflow-auto text-sm space-x-2 pb-1">
          <RouterLink to="/" class="px-3 py-1.5 rounded text-white/80" active-class="bg-white/20 !text-white font-bold">Dashboard</RouterLink>
          <RouterLink to="/add" class="px-3 py-1.5 rounded shrink-0 text-white/80" active-class="bg-white/20 !text-white font-bold">Tambah Data</RouterLink>
          <RouterLink to="/list" class="px-3 py-1.5 rounded shrink-0 text-white/80" active-class="bg-white/20 !text-white font-bold">Daftar Data</RouterLink>
        </nav>
      </header>

      <!-- Desktop Header -->
      <header class="hidden md:flex bg-white border-b border-text/10 h-16 items-center px-6 shrink-0 shadow-sm z-10">
        <h2 class="text-lg font-bold text-text">{{ pageTitle }}</h2>
      </header>

      <!-- Scrollable Main Area -->
      <main class="flex-1 overflow-auto p-4 md:p-6 relative bg-surface bg-grid-pattern">
        <RouterView />
      </main>
    </div>
  </div>
</template>
