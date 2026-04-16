<script setup lang="ts">
import { useMapPointsStore } from '@/stores/mapPoints'
import MapComponent from '@/components/MapComponent.vue'

const store = useMapPointsStore()
store.isEditMode = false
</script>

<template>
  <div class="h-full w-full relative">
    <MapComponent />
  </div>
</template>
