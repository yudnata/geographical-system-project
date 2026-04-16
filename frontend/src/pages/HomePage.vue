<script setup lang="ts">
import { ref } from 'vue'
import DashboardStats from '@/components/DashboardStats.vue'
import MapComponent from '@/components/MapComponent.vue'

const mapFlyTo = ref<{ lat: number; lng: number } | null>(null)
</script>

<template>
  <div class="space-y-6 max-w-7xl mx-auto">
    <div>
      <DashboardStats />
    </div>

    <!-- Peta Area -->
    <div
      class="bg-white rounded-xl shadow-sm border border-text/10 overflow-hidden relative"
      style="height: calc(100vh - 250px); min-height: 500px"
    >
      <MapComponent :fly-to="mapFlyTo" class="w-full h-full" />
    </div>
  </div>
</template>
