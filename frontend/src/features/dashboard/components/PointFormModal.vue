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
const activeTab = ref<'data' | 'blog'>('data')
const blogContent = ref({ title: '', content: '', cover_photo: '' })

watch(() => store.activePoint, async (newPoint) => {
  if (newPoint) {
    formData.value = { ...newPoint }
    if (newPoint.id) {
      const blog = await store.getBlog(newPoint.id)
      if (blog) {
        blogContent.value = { title: blog.title, content: blog.content, cover_photo: blog.cover_photo || '' }
      } else {
        blogContent.value = { title: newPoint.name, content: '', cover_photo: '' }
      }
    }
  } else {
    formData.value = getDefaultForm()
    blogContent.value = { title: '', content: '', cover_photo: '' }
  }
  activeTab.value = 'data'
}, { immediate: true })


const submitForm = async (status: 'draft' | 'pending') => {
  if (!formData.value.name || !formData.value.latitude || !formData.value.longitude) {
    notificationStore.warning('Nama dan Koordinat wajib diisi!')
    return
  }

  if (isSubmitting.value) return
  isSubmitting.value = true
  
  formData.value.status = status


  try {
    const savedPoint = await store.savePoint(formData.value as GeoPoint)
    if (savedPoint && savedPoint.id) {
      // Also save blog if content exists
      await store.saveBlog(savedPoint.id, {
        title: formData.value.name || 'Ulasan',
        content: blogContent.value.content,
        cover_photo: formData.value.description // Reusing description as cover photo for now
      })
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
  <Transition name="modal-fade">
    <div v-if="store.isModalOpen" class="fixed inset-0 z-[9999] flex items-center justify-center p-4">
      <!-- Backdrop -->
      <div class="absolute inset-0 bg-gray-900/60 backdrop-blur-md" @click="store.closeModal()"></div>

      <!-- Modal Content -->
      <div class="relative bg-white rounded-[2rem] shadow-[0_20px_50px_rgba(0,0,0,0.3)] w-full max-w-lg flex flex-col max-h-[90vh] overflow-hidden transform transition-all border border-white/20">

        <PointFormHeader :is-edit="!!formData.id" />

        <!-- Tab Switcher -->
        <div class="px-6 py-2 border-b border-slate-100 flex gap-6">
          <button 
            @click="activeTab = 'data'"
            :class="[activeTab === 'data' ? 'text-indigo-600 border-b-2 border-indigo-600 font-bold' : 'text-slate-400 font-semibold']"
            class="pb-2 text-xs transition-all"
          >
            Data Objek
          </button>
          <button 
            @click="activeTab = 'blog'"
            :class="[activeTab === 'blog' ? 'text-indigo-600 border-b-2 border-indigo-600 font-bold' : 'text-slate-400 font-semibold']"
            class="pb-2 text-xs transition-all"
          >
            Konten Blog
          </button>
        </div>

        <PointFormFields v-model="formData" v-model:blog-content="blogContent" :active-tab="activeTab" />


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
