<script setup lang="ts">
import { useFacilityStore } from '@/stores/facility'
import { toRefs } from 'vue'

const store = useFacilityStore()
const { statsByType } = toRefs(store)
</script>

<template>
  <section class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
    <div
      class="bg-white rounded-xl shadow-sm border border-text/10 p-5 flex items-center justify-between"
    >
      <div>
        <h3 class="text-text/60 text-sm font-medium">Puskesmas</h3>
        <p class="text-3xl font-bold text-text mt-2">{{ statsByType['Puskesmas'] || 0 }}</p>
      </div>
    </div>
    <div
      class="bg-white rounded-xl shadow-sm border border-text/10 p-5 flex items-center justify-between"
    >
      <div>
        <h3 class="text-text/60 text-sm font-medium">Klinik</h3>
        <p class="text-3xl font-bold text-text mt-2">{{ statsByType['Klinik'] || 0 }}</p>
      </div>
    </div>
    <div
      class="bg-white rounded-xl shadow-sm border border-text/10 p-5 flex items-center justify-between"
    >
      <div>
        <h3 class="text-text/60 text-sm font-medium">Rumah Sakit</h3>
        <p class="text-3xl font-bold text-text mt-2">{{ statsByType['Rumah Sakit'] || 0 }}</p>
      </div>
    </div>
    <div
      class="bg-white rounded-xl shadow-sm border border-text/10 p-5 flex items-center justify-between"
    >
      <div>
        <h3 class="text-text/60 text-sm font-medium">Apotek</h3>
        <p class="text-3xl font-bold text-text mt-2">{{ statsByType['Apotek'] || 0 }}</p>
      </div>
    </div>
  </section>
</template>
