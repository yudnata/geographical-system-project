<script setup lang="ts">
import { computed, watch } from 'vue'
import { RouterView, RouterLink, useRoute } from 'vue-router'
import AppSidebar from '@/components/Sidebar/AppSidebar.vue'
import AppHeader from '@/components/Header/AppHeader.vue'
import { useAuthStore } from '@/stores/auth'
import { useMapUIStore } from '@/stores/mapUI'

const route = useRoute()
const authStore = useAuthStore()
const uiStore = useMapUIStore()

watch(() => route.path, () => {
  uiStore.setSelectedPreviewPoint(null)
})

const hideHeader = computed(() => {
  const routesWithoutHeader = ['/admin/categories', '/admin/users', '/admin/verification', '/contributor/dashboard', '/blog']
  return routesWithoutHeader.some(path => route.path.startsWith(path))
})

</script>



<template>
  <div class="flex h-screen w-full bg-surface font-sans text-text antialiased overflow-hidden">
    <AppSidebar />

    <div class="flex-1 flex flex-col h-screen overflow-hidden relative">
      <!-- Mobile Header -->
      <header class="bg-primary text-white shadow-sm md:hidden p-4 flex flex-col gap-2">
        <div>
          <h1 class="text-lg font-bold">Budaya Bali</h1>

        </div>
        <nav class="flex overflow-auto text-sm space-x-2 pb-1">
          <RouterLink to="/" class="px-3 py-1.5 rounded text-white/80" active-class="bg-white/20 !text-white font-bold">Peta Publik</RouterLink>
          <RouterLink v-if="authStore.user" to="/dashboard" class="px-3 py-1.5 rounded shrink-0 text-white/80" active-class="bg-white/20 !text-white font-bold">Editor Peta</RouterLink>
          <RouterLink v-if="authStore.user" to="/tabular" class="px-3 py-1.5 rounded shrink-0 text-white/80" active-class="bg-white/20 !text-white font-bold">Tabel Master</RouterLink>
        </nav>
      </header>

      <main class="flex-1 relative overflow-hidden bg-surface h-full flex flex-col min-h-0">
        <AppHeader v-if="!hideHeader" />
        <RouterView />
      </main>

    </div>
  </div>
</template>
