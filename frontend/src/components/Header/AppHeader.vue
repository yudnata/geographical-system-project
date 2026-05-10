<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useMapUIStore } from '@/stores/mapUI'
import HeaderSearch from './HeaderSearch.vue'
import HeaderCategoryFilter from './HeaderCategoryFilter.vue'
import HeaderEditorControls from './HeaderEditorControls.vue'
import HeaderUserInfo from './HeaderUserInfo.vue'

const route = useRoute()
const uiStore = useMapUIStore()

const pageTitle = computed(() => {
  if (route.path === '/') return 'Eksplorasi Budaya Bali'
  if (route.path === '/dashboard') return 'Kontribusi Saya'
  if (route.path === '/tabular') return 'Pustaka Data Master'
  if (route.path === '/admin/verification') return 'Verifikasi Objek'
  if (route.path === '/admin/categories') return 'Kelola Kategori'
  if (route.path === '/admin/users') return 'Manajemen Tim'
  return 'Budaya Bali: Digital Mapping'
})

const isDashboard = computed(() => route.path === '/dashboard')
</script>

<template>
  <header :class="[
    'hidden md:flex absolute top-4 right-4 bg-white/95 backdrop-blur-md border border-gray-100 h-16 items-center px-6 shrink-0 shadow-[0_8px_32px_rgba(0,0,0,0.12)] z-[1000] justify-between rounded-2xl transition-all duration-500 ease-in-out',
    uiStore.isSidebarExpanded ? 'left-[288px]' : 'left-24'
  ]">

    <div class="flex items-center gap-4 shrink-0">
      <h2 class="text-lg font-bold text-gray-800 whitespace-nowrap">{{ pageTitle }}</h2>
    </div>

    <div class="flex items-center gap-2.5 flex-1 px-6">
      <HeaderSearch />
      <HeaderCategoryFilter />
      <HeaderEditorControls v-if="isDashboard" />
    </div>

    <HeaderUserInfo />
  </header>
</template>
