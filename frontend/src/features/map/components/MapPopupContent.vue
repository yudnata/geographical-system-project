<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { useMapPointsStore, type GeoPoint } from '@/stores/mapPoints'
import { useMapUIStore } from '@/stores/mapUI'

const props = defineProps<{
  point: GeoPoint
  typeName: string
}>()

const authStore = useAuthStore()
const store = useMapPointsStore()
const uiStore = useMapUIStore()

const handleEdit = () => {
  store.openModal(props.point)
}

const handleDelete = () => {
  store.requestConfirm(
    'Hapus Bangunan?',
    `Apakah Anda yakin ingin menghapus data bangunan "${props.point.name}" secara permanen?`,
    async () => {
      if (props.point.id) {
        await store.deletePoint(props.point.id)
      }
    }
  )
}
</script>

<template>
  <div class="p-3 min-w-[220px] font-sans">
    <h4 class="font-black text-gray-900 text-sm mb-1 tracking-tight">{{ point.name }}</h4>

    <div class="space-y-1.5 text-[11px]">
      <p class="flex items-start gap-2">
        <span class="w-20 shrink-0 font-semibold text-gray-400">Kategori:</span>
        <span class="bg-gray-100 px-1.5 rounded text-gray-700 font-bold">{{ typeName }}</span>
      </p>

      <p class="flex items-start gap-2">
        <span class="w-20 shrink-0 font-semibold text-gray-400">Alamat:</span>
        <span class="text-gray-600 leading-normal">{{ point.address || 'Tanpa Alamat' }}</span>
      </p>

      <p v-if="point.tahun_berdiri" class="flex items-start gap-2">
        <span class="w-20 shrink-0 font-semibold text-gray-400">Th. Berdiri:</span>
        <span class="text-gray-800">{{ point.tahun_berdiri }}</span>
      </p>

      <p class="flex items-start gap-2">
        <span class="w-20 shrink-0 font-semibold text-gray-400">Kepemilikan:</span>
        <span class="px-1.5 bg-gray-100 rounded text-gray-700 font-bold">{{ point.status_kepemilikan || 'Tidak Diketahui' }}</span>
      </p>

      <div class="mt-2 pt-2 border-t border-dotted border-gray-200 flex items-center gap-1.5 opacity-70 italic text-[10px]">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-3 h-3">
          <path d="M10 8a3 3 0 100-6 3 3 0 000 6zM3.465 14.493a1.23 1.23 0 00.41 1.412A9.957 9.957 0 0010 18c2.31 0 4.438-.784 6.131-2.1.43-.333.604-.903.408-1.41a7.002 7.002 0 00-13.074.003z" />
        </svg>
        Kontributor: {{ point.owner_name || 'Sistem' }}
      </div>

      <div v-if="authStore.isAuthenticated() && uiStore.isEditMode && point.owner_id === authStore.user?.id" class="mt-3 flex gap-2 pt-2 border-t border-gray-100">
        <button @click="handleEdit" class="flex-1 bg-primary/10 hover:bg-primary/20 text-primary font-bold py-1.5 rounded text-[10px] uppercase tracking-wide transition-colors">
          Edit
        </button>
        <button @click="handleDelete" class="flex-1 bg-red-50 hover:bg-red-100 text-red-600 font-bold py-1.5 rounded text-[10px] uppercase tracking-wide transition-colors">
          Hapus
        </button>
      </div>
    </div>
  </div>
</template>
