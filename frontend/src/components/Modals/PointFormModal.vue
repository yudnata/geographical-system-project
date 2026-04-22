<script setup lang="ts">
import { ref, watch, onMounted, onUnmounted } from 'vue'
import { useMapPointsStore, type GeoPoint } from '@/stores/mapPoints'
import { useNotificationStore } from '@/stores/notifications'

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
    const success = await store.savePoint(formData.value)
    if (success) {
      store.closeModal()
    }
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
  <Transition name="fade">
    <div v-if="store.isModalOpen" class="fixed inset-0 z-[9999] flex items-center justify-center">

      <div class="absolute inset-0 bg-gray-900/40 backdrop-blur-sm" @click="store.closeModal()"></div>


      <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-lg mx-4 flex flex-col max-h-[90vh] overflow-hidden transform transition-all">

        <div class="px-6 py-4 border-b border-gray-100 flex items-center justify-between shrink-0 bg-gray-50/50">
          <h3 class="text-lg font-bold text-gray-800">
            {{ formData.id ? 'Edit Data Bangunan' : 'Tambah Bangunan Baru' }}
          </h3>
          <button @click="store.closeModal()" class="text-gray-400 hover:text-red-500 transition-colors p-1 rounded-lg hover:bg-red-50">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>


        <div class="p-6 overflow-y-auto w-full space-y-4">

          <div class="bg-blue-50 text-blue-800 text-xs px-4 py-2.5 rounded-lg flex gap-3 font-mono border border-blue-100">
            <div><strong class="uppercase text-blue-900/60">Lat:</strong> {{ formData.latitude?.toFixed(6) || 0 }}</div>
            <div><strong class="uppercase text-blue-900/60">Lng:</strong> {{ formData.longitude?.toFixed(6) || 0 }}</div>
          </div>

          <div class="space-y-1">
            <label class="text-xs font-semibold text-gray-600 uppercase">Nama Bangunan *</label>
            <input v-model="formData.name" type="text" required
              class="w-full px-4 py-2 bg-gray-50 border border-gray-200 rounded-lg focus:bg-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm"
              placeholder="Contoh: RSUD Wangaya">
          </div>

          <div class="space-y-1">
            <label class="text-xs font-semibold text-gray-600 uppercase">Tipe Objek Master *</label>
            <select v-model="formData.type_id"
              class="w-full px-4 py-2 bg-gray-50 border border-gray-200 rounded-lg focus:bg-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm appearance-none">
              <option v-for="type in store.objectTypes" :value="type.id" :key="type.id">{{ type.name }}</option>
            </select>
          </div>

          <div class="space-y-1">
            <label class="text-xs font-semibold text-gray-600 flex items-center justify-between uppercase">
              Alamat Lengkap *
              <span class="text-[10px] bg-green-100 text-green-700 px-2 py-0.5 rounded shadow-sm normal-case font-bold">Auto</span>
            </label>
            <textarea v-model="formData.address" rows="2"
              class="w-full px-4 py-2 bg-gray-50 border border-gray-200 rounded-lg focus:bg-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm"
              placeholder="Alamat lengkap (Terisi otomatis dari klik peta)..."></textarea>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div class="space-y-1">
              <label class="text-xs font-semibold text-gray-600 uppercase">Tahun Berdiri</label>
              <input v-model="formData.tahun_berdiri" type="number"
                class="w-full px-4 py-2 bg-gray-50 border border-gray-200 rounded-lg focus:bg-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm">
            </div>

            <div class="space-y-1">
              <label class="text-xs font-semibold text-gray-600 uppercase">Kepemilikan</label>
              <select v-model="formData.status_kepemilikan"
                class="w-full px-4 py-2 bg-gray-50 border border-gray-200 rounded-lg focus:bg-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm appearance-none">
                <option value="Pemerintah">Pemerintah</option>
                <option value="Swasta">Swasta</option>
                <option value="Pribadi">Pribadi</option>
                <option value="Yayasan">Yayasan</option>
              </select>
            </div>
          </div>
        </div>


        <div class="px-6 py-4 border-t border-gray-100 bg-gray-50/50 shrink-0 flex items-center justify-end gap-3 rounded-b-2xl">
          <button @click="store.closeModal()" :disabled="isSubmitting"
            class="px-5 py-2.5 text-sm font-semibold text-gray-600 bg-white border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors disabled:opacity-50">Batal</button>
          <button @click="submitForm" :disabled="isSubmitting"
            class="px-5 py-2.5 text-sm font-bold text-white bg-primary rounded-lg hover:bg-primary-dark transition-colors shadow-sm active:scale-[0.98] disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2">
            <svg v-if="isSubmitting" class="animate-spin h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ formData.id ? 'Simpan Perubahan' : 'Tambahkan Poin' }}
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
