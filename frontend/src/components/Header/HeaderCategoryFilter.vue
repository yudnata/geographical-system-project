<script setup lang="ts">
import { useMapPointsStore } from '@/stores/mapPoints'

const store = useMapPointsStore()
</script>

<template>
  <div class="relative shrink-0">
    <select v-model="store.filterTypeId"
      class="h-9 bg-white border border-gray-300 pl-3 pr-7 rounded-xl text-[11px] font-bold text-gray-700 appearance-none focus:ring-4 focus:ring-primary/10 focus:border-primary outline-none transition-all cursor-pointer uppercase tracking-tight">
      <option :value="null">SEMUA KATEGORI</option>
      <option v-for="type in store.objectTypes" :key="type.id" :value="type.id">
        {{ type.name.toUpperCase() }}
      </option>
    </select>
    <span class="absolute inset-y-0 right-2 flex items-center pointer-events-none text-gray-500">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-3 h-3">
        <path stroke-linecap="round" stroke-linejoin="round" d="m19.5 8.25-7.5 7.5-7.5-7.5" />
      </svg>
    </span>
  </div>
</template>
