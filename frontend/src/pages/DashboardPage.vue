<script setup lang="ts">
import { useMapPointsStore } from '@/stores/mapPoints'
import MapComponent from '@/components/MapComponent.vue'
import PointFormModal from '@/components/PointFormModal.vue'
import ProfileCompletionModal from '@/components/ProfileCompletionModal.vue'

const store = useMapPointsStore()

const handleMapClick = (data: { lat: number; lng: number; address?: string }) => {
  store.openModal({
    id: 0,
    type_id: 1,
    name: '',
    latitude: data.lat,
    longitude: data.lng,
    address: data.address || '',
    owner_id: store.currentUserId,
    is_active: true,
    tahun_berdiri: new Date().getFullYear(),
    status_kepemilikan: 'Pemerintah',
  })
}
</script>

<template>
  <div class="h-full w-full relative flex flex-col">

    <ProfileCompletionModal />

    <MapComponent @map-clicked="handleMapClick" class="flex-1" />


    <div class="absolute top-4 right-4 z-20 flex gap-3">

      <div class="bg-white px-3 py-2 rounded-xl shadow-[0_4px_24px_rgba(0,0,0,0.1)] border border-gray-100 flex items-center gap-3">
        <label class="text-xs font-bold text-gray-400 uppercase tracking-widest">Tampilan:</label>
        <div class="flex bg-gray-100 rounded-lg p-0.5">
          <button @click="store.filterMyPoints = false" :class="!store.filterMyPoints ? 'bg-white shadow-sm text-gray-800' : 'text-gray-500 hover:text-gray-700'"
            class="px-3 py-1 text-xs font-bold rounded-md transition-all">Semua Bangunan</button>
          <button @click="store.filterMyPoints = true" :class="store.filterMyPoints ? 'bg-white shadow-sm text-gray-800' : 'text-gray-500 hover:text-gray-700'"
            class="px-3 py-1 text-xs font-bold rounded-md transition-all">Milik Saya Saja</button>
        </div>
      </div>


      <button @click="store.toggleEditMode()" :class="store.isEditMode ? 'bg-primary text-white border-primary shadow-primary/30' : 'bg-white text-gray-700 border-gray-200 hover:bg-gray-50'"
        class="px-4 py-2 border rounded-xl shadow-[0_4px_24px_rgba(0,0,0,0.1)] flex items-center gap-2 font-bold text-sm transition-all">
        <div class="relative flex h-3 w-3">
          <span v-if="store.isEditMode" class="animate-ping absolute inline-flex h-full w-full rounded-full bg-white opacity-75"></span>
          <span class="relative inline-flex rounded-full h-3 w-3" :class="store.isEditMode ? 'bg-white' : 'bg-gray-300'"></span>
        </div>
        MODE EDIT
      </button>
    </div>


    <PointFormModal />
  </div>
</template>
