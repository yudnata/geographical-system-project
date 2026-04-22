<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useMapPointsStore } from '@/stores/mapPoints'
import HeaderSearch from './HeaderSearch.vue'
import HeaderCategoryFilter from './HeaderCategoryFilter.vue'
import HeaderEditorControls from './HeaderEditorControls.vue'
import HeaderUserInfo from './HeaderUserInfo.vue'

const route = useRoute()
const store = useMapPointsStore()

const pageTitle = computed(() => {
  if (route.path === '/') return 'Peta Publik'
  if (route.path === '/dashboard') return 'Peta Author'
  if (route.path === '/tabular') return 'Tabel Data Master'
  return 'Sistem Informasi Geografis'
})

const isDashboard = computed(() => route.path === '/dashboard')
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
      <HeaderSearch />
      <HeaderCategoryFilter />
      <HeaderEditorControls v-if="isDashboard" />
    </div>

    <HeaderUserInfo />
  </header>
</template>
