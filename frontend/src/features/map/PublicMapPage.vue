<script setup lang="ts">
import { useMapUIStore } from '@/stores/mapUI'
import MapContainer from '@/features/map/components/MapContainer.vue'

const uiStore = useMapUIStore()

uiStore.isEditMode = false
uiStore.filterMyPoints = false
</script>

<template>
  <div class="h-full w-full relative">
    <MapContainer />
  </div>
</template>
