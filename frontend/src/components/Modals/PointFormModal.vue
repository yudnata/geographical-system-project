<script setup lang="ts">
import { ref, watch, onMounted, onUnmounted } from 'vue'
import { useMapPointsStore, type GeoPoint } from '@/stores/mapPoints'
import { useNotificationStore } from '@/stores/notifications'

// Sub-components
import PointFormHeader from './PointForm/PointFormHeader.vue'
import PointFormFields from './PointForm/PointFormFields.vue'
import PointFormActions from './PointForm/PointFormActions.vue'

const store = useMapPointsStore()
const notificationStore = useNotificationStore()

const getDefaultForm = (): Partial<GeoPoint> => ({
  type_id: store.objectTypes.length > 0 ? store.objectTypes[0]?.id || 1 : 1,
  name: '',
  latitude: 0,
  longitude: 0,
  address: '',
  tahun_berdiri: new Date().getFullYear(),
  status_kepemilikan: 'Pemerintah',
  description: ''
})

const formData = ref<Partial<GeoPoint>>(getDefaultForm())
const isSubmitting = ref(false)

watch(() => store.activePoint, (newPoint) => {
  if (newPoint) {
    formData.value = { ...newPoint }
  } else {
    formData.value = getDefaultForm()
  }
}, { immediate: true })

const submitForm = async () => {
  if (!formData.value.name || !formData.value.latitude || !formData.value.longitude) {
    notificationStore.warning('Nama dan Koordinat wajib diisi!')
    return
  }

  if (isSubmitting.value) return
  isSubmitting.value = true

  try {
    await store.savePoint(formData.value as GeoPoint)
  } finally {
    isSubmitting.value = false
  }
}

const handleKeydown = (e: KeyboardEvent) => {
  if (e.key === 'Escape' && store.isModalOpen) store.closeModal()
}

onMounted(() => window.addEventListener('keydown', handleKeydown))
onUnmounted(() => window.removeEventListener('keydown', handleKeydown))
</script>

<template>
  <Transition name="modal-fade">
    <div v-if="store.isModalOpen" class="fixed inset-0 z-[9999] flex items-center justify-center p-4">
      <!-- Backdrop -->
      <div class="absolute inset-0 bg-gray-900/60 backdrop-blur-md" @click="store.closeModal()"></div>

      <!-- Modal Content -->
      <div class="relative bg-white rounded-[2rem] shadow-[0_20px_50px_rgba(0,0,0,0.3)] w-full max-w-lg flex flex-col max-h-[90vh] overflow-hidden transform transition-all border border-white/20">

        <PointFormHeader :is-edit="!!formData.id" />

        <PointFormFields v-model="formData" />

        <PointFormActions :is-edit="!!formData.id" :is-submitting="isSubmitting" @submit="submitForm" />

      </div>
    </div>
  </Transition>
</template>

<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
  transform: scale(0.9) translateY(20px);
}
</style>
