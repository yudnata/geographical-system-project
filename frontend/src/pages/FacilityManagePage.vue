<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useFacilityStore } from '@/stores/facility'
import type { Facility } from '@/stores/facility'

import MapComponent from '@/components/MapComponent.vue'
import FacilityForm from '@/components/FacilityForm.vue'

const store = useFacilityStore()
const router = useRouter()
const route = useRoute()

const mapFlyTo = ref<{ lat: number; lng: number } | null>(null)

const formFacility = ref<Facility | null>(null)

const handleMapClick = (coords: { lat: number; lng: number }) => {
  if (formFacility.value) {
    formFacility.value = {
      ...formFacility.value,
      lat: Number(coords.lat.toFixed(6)),
      lng: Number(coords.lng.toFixed(6)),
    }
  } else {
    formFacility.value = {
      id: 0,
      name: '',
      type: 'Puskesmas',
      address: '',
      district: '',
      phone: '',
      hours: '',
      lat: Number(coords.lat.toFixed(6)),
      lng: Number(coords.lng.toFixed(6)),
    }
  }
}

const handleSave = (data: Omit<Facility, 'id'>) => {
  store.addFacility(data)
  router.push('/list')
}

const handleUpdate = (data: Facility) => {
  store.updateFacility(data.id, data)
  router.push('/list')
}

const handleCancel = () => {
  router.push('/list')
}

onMounted(() => {
  const editId = route.params.id
  if (editId) {
    const existing = store.facilities.find((f) => f.id === Number(editId))
    if (existing) {
      formFacility.value = { ...existing }
      mapFlyTo.value = { lat: existing.lat, lng: existing.lng }
    }
  }
})
</script>

<template>
  <div class="max-w-7xl mx-auto space-y-6">
    <div class="grid grid-cols-1 lg:grid-cols-5 gap-6">
      <!-- Input Form: 2 cols -->
      <div class="lg:col-span-2 order-2 lg:order-1">
        <FacilityForm
          :edit-data="formFacility"
          @save="handleSave"
          @update="handleUpdate"
          @cancel="handleCancel"
        />
      </div>

      <!-- Map: 3 cols -->
      <div
        class="lg:col-span-3 order-1 lg:order-2 bg-white rounded-xl shadow-sm border border-text/10 overflow-hidden relative"
        style="height: calc(100vh - 150px); min-height: 500px"
      >
        <MapComponent :fly-to="mapFlyTo" @map-clicked="handleMapClick" class="w-full h-full" />
      </div>
    </div>
  </div>
</template>
