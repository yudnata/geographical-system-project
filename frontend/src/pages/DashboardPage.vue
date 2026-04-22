<script setup lang="ts">
import { useMapPointsStore } from '@/stores/mapPoints'
import { useAuthStore } from '@/stores/auth'
import { useNotificationStore } from '@/stores/notifications'
import MapContainer from '@/components/Map/MapContainer.vue'
import PointFormModal from '@/components/Modals/PointFormModal.vue'
import ProfileCompletionModal from '@/components/Modals/ProfileCompletionModal.vue'

const store = useMapPointsStore()
const authStore = useAuthStore()
const notificationStore = useNotificationStore()

const handleMapClick = (data: { lat: number; lng: number; address?: string }) => {
  notificationStore.info('Membuka form...')
  store.openModal({
    id: 0,
    type_id: 1,
    name: '',
    latitude: data.lat,
    longitude: data.lng,
    address: data.address || 'Sedang mencari alamat...',
    owner_id: authStore.user?.id || '',
    is_active: true,
    tahun_berdiri: new Date().getFullYear(),
    status_kepemilikan: 'Pemerintah',
  })
}
</script>

<template>
  <div class="h-full w-full relative flex flex-col">

    <ProfileCompletionModal />

    <MapContainer @map-clicked="handleMapClick" class="flex-1" />

    <PointFormModal />
  </div>
</template>
