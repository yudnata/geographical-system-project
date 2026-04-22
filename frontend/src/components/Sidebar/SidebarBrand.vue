<script setup lang="ts">
import { useMapPointsStore } from '@/stores/mapPoints'

const store = useMapPointsStore()
</script>

<template>
  <div :class="['h-16 flex items-center shrink-0 border-b border-gray-200 whitespace-nowrap box-border transition-[padding] duration-500', store.isSidebarExpanded ? 'px-6' : 'px-4']">
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
</template>
