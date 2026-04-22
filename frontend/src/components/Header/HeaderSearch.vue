<script setup lang="ts">
import { useMapPointsStore, type GeoPoint } from '@/stores/mapPoints'

const store = useMapPointsStore()

const handleSelectPoint = (point: GeoPoint) => {
  store.flyTo(point.latitude, point.longitude)
  store.searchQuery = ''
}
</script>

<template>
  <div class="flex-1 max-w-sm relative group">
    <div class="absolute inset-y-0 left-3 flex items-center text-gray-500 group-focus-within:text-primary transition-colors">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5" stroke="currentColor" class="w-4 h-4">
        <path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
      </svg>
    </div>
    <input v-model="store.searchQuery" type="text" placeholder="Cari lokasi bangunan..."
      class="w-full h-9 pl-9 pr-4 bg-white border border-gray-300 rounded-xl focus:border-primary focus:ring-4 focus:ring-primary/10 transition-all text-xs font-semibold text-gray-800 placeholder:text-gray-400 outline-none" />

    <div v-if="store.searchQuery && store.filteredPoints.length > 0"
      class="absolute top-full left-0 right-0 mt-2 bg-white/90 backdrop-blur-xl border border-white/20 shadow-[0_8px_32px_rgba(0,0,0,0.12)] rounded-2xl overflow-hidden z-[2000]">
      <div class="max-h-64 overflow-y-auto">
        <div v-for="point in store.filteredPoints.slice(0, 5)" :key="point.id" @click="handleSelectPoint(point)"
          class="px-4 py-3 hover:bg-primary/5 cursor-pointer flex items-center gap-3 transition-colors border-b border-gray-50 last:border-0">
          <div class="w-7 h-7 rounded-lg bg-primary/10 text-primary flex items-center justify-center shrink-0">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-3.5 h-3.5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z" />
              <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1 1 15 0Z" />
            </svg>
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-xs font-extrabold text-gray-900 truncate tracking-tight">{{ point.name }}</p>
            <p class="text-[10px] text-gray-500 truncate font-medium">{{ point.address }}</p>
          </div>
        </div>
      </div>
      <div v-if="store.filteredPoints.length > 5" class="px-4 py-2 bg-gray-50/50 text-[10px] font-bold text-gray-400 uppercase tracking-widest text-center border-t border-gray-50">
        +{{ store.filteredPoints.length - 5 }} lainnya
      </div>
    </div>
  </div>
</template>
