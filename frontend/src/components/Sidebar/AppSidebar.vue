<script setup lang="ts">
import { useMapPointsStore } from '@/stores/mapPoints'
import SidebarBrand from './SidebarBrand.vue'
import SidebarNav from './SidebarNav.vue'
import SidebarFooter from './SidebarFooter.vue'

const store = useMapPointsStore()
</script>

<template>
  <aside @mouseenter="store.isSidebarExpanded = true" @mouseleave="store.isSidebarExpanded = false"
    class="group absolute top-4 left-4 bottom-4 w-16 hover:w-64 flex-col hidden md:flex shrink-0 transition-[width] duration-500 ease-in-out z-[1001] bg-transparent will-change-[width]">
    <div class="absolute inset-y-0 -left-4 -right-8 z-[-1] cursor-pointer"></div>

    <div class="flex-1 flex flex-col bg-white/80 backdrop-blur-md overflow-hidden shadow-[0_8px_32px_rgba(0,0,0,0.12)] border border-gray-100 rounded-2xl">
      <SidebarBrand />
      <SidebarNav />
      <SidebarFooter />
    </div>
  </aside>
</template>
