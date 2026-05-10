<script setup lang="ts">
import { onMounted } from 'vue'
import { useMapPointsStore } from '@/stores/mapPoints'
import { useMapUIStore } from '@/stores/mapUI'
import { useNotificationStore } from '@/stores/notifications'
import MapContainer from '@/features/map/components/MapContainer.vue'

const store = useMapPointsStore()
const uiStore = useMapUIStore()
const notificationStore = useNotificationStore()

onMounted(() => {
  uiStore.filterMyPoints = true
  uiStore.isEditMode = true
})

const handleMapClick = (data: { lat: number; lng: number; address?: string }) => {
  notificationStore.info('Membuka form...')
  store.openModal({
    id: 0,
    name: '',
    latitude: data.lat,
    longitude: data.lng,
    type_id: 1,
    address: data.address || '',
  })
}
</script>

<template>
  <div class="h-full w-full flex flex-col overflow-hidden relative">
    <MapContainer @map-clicked="handleMapClick" />
  </div>
</template>
