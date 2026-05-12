<script setup lang="ts">
import { onMounted, computed } from 'vue'
import { usePointsStore } from '@/stores/pointsStore'
import { useUIStore } from '@/stores/uiStore'
import type { GeoPoint } from '@/types/pointTypes'
import { useAuthStore } from '@/stores/auth'
import { useRouter } from 'vue-router'

const store = usePointsStore()
const uiStore = useUIStore()
const authStore = useAuthStore()
const router = useRouter()


onMounted(async () => {
  if (authStore.isAdmin()) {
    uiStore.filterMyPoints = false
  } else {
    uiStore.filterMyPoints = true
  }
  await store.fetchCategories()
  await store.fetchPoints()
})


const tabularPoints = computed(() => {
  return store.points.filter((point) => {
    const query = uiStore.searchQuery.toLowerCase()
    const matchSearch =
      point.name.toLowerCase().includes(query) ||
      (point.address && point.address.toLowerCase().includes(query))
    const matchType = !uiStore.filterTypeId || point.category_id === uiStore.filterTypeId
    const matchOwner = !uiStore.filterMyPoints || point.owner_id === authStore.user?.id

    return matchSearch && matchType && matchOwner
  })
})

const getTypeName = (categoryId: number) => {
  const type = store.objectTypes.find(t => t.id === categoryId)
  return type ? type.name : 'Unknown'
}

const getStatusLabel = (status?: string) => {
  switch (status) {
    case 'approved': return { label: 'Terverifikasi', class: 'text-emerald-600', dot: 'bg-emerald-500' }
    case 'pending': return { label: 'Menunggu', class: 'text-amber-500', dot: 'bg-amber-500' }
    case 'rejected': return { label: 'Ditolak', class: 'text-rose-600', dot: 'bg-rose-500' }
    default: return { label: 'Draft', class: 'text-gray-500', dot: 'bg-gray-400' }
  }
}

const handleDelete = (point: GeoPoint) => {
  store.requestConfirm(
    'Hapus Data?',
    `Apakah Anda yakin ingin menghapus data situs "${point.name}" secara permanen?`,
    async () => {
      if (point.id) {
        await store.deletePoint(point.id)
      }
    }
  )
}
</script>


<template>
  <div :class="[
    'h-full w-full flex flex-col transition-[padding] duration-500 ease-in-out overflow-hidden',
    uiStore.isSidebarExpanded ? 'pl-[288px]' : 'pl-24',
    'pt-24 pb-6 pr-6'
  ]">
    <div class="flex-1 flex flex-col bg-white rounded-2xl shadow-[0_8px_32px_rgba(0,0,0,0.06)] border border-gray-100 overflow-hidden">
      <!-- No header here anymore, directly table or empty space padding if needed -->

      <div class="flex-1 overflow-auto">
        <table class="w-full text-left border-separate border-spacing-0 min-w-max">
          <thead class="bg-gray-50/80 sticky top-0 z-10">
            <tr>
              <th class="py-4 px-6 text-[10px] font-black text-gray-400 uppercase tracking-widest border-b border-gray-200">ID</th>
              <th class="py-4 px-6 text-[10px] font-black text-gray-400 uppercase tracking-widest border-b border-gray-200">Nama Objek / Situs</th>
              <th class="py-4 px-6 text-[10px] font-black text-gray-400 uppercase tracking-widest border-b border-gray-200">Kategori</th>
              <th class="py-4 px-6 text-[10px] font-black text-gray-400 uppercase tracking-widest border-b border-gray-200">Estimasi Masa / Sejarah</th>
              <th class="py-4 px-6 text-[10px] font-black text-gray-400 uppercase tracking-widest border-b border-gray-200">Status</th>
              <th class="py-4 px-6 text-[10px] font-black text-gray-400 uppercase tracking-widest text-center border-b border-gray-200">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-100">
            <tr v-for="point in tabularPoints" :key="point.id" class="hover:bg-blue-50/40 transition-colors group">
              <td class="py-5 px-6 text-xs text-gray-500 font-mono border-b border-gray-50">#{{ point.id }}</td>
              <td class="py-5 px-6 border-b border-gray-50">
                <div class="font-black text-gray-900 text-sm tracking-tight group-hover:text-primary transition-colors">{{ point.name }}</div>
                <div class="text-[10px] text-gray-400 font-bold uppercase tracking-tighter mt-1 line-clamp-1 max-w-xs" :title="point.address">{{ point.address }}</div>
              </td>
              <td class="py-5 px-6 border-b border-gray-50">
                <span class="px-3 py-1.5 text-[10px] font-black tracking-widest rounded-xl bg-slate-100 text-slate-600 border border-slate-200 uppercase shadow-sm">
                  {{ getTypeName(point.category_id) }}
                </span>
              </td>
              <td class="py-5 px-6 text-xs text-gray-600 font-black border-b border-gray-50">{{ point.tahun_berdiri || '-' }}</td>
              <td class="py-5 px-6 border-b border-gray-50">
                <span :class="[getStatusLabel(point.status).class, 'text-[10px] font-black uppercase tracking-widest flex items-center gap-2']">
                  <span :class="[getStatusLabel(point.status).dot, 'w-1.5 h-1.5 rounded-full animate-pulse']"></span>
                  {{ getStatusLabel(point.status).label }}
                </span>
              </td>

              <td class="py-5 px-6 text-center border-b border-gray-50">
                <div class="flex items-center justify-center gap-2">
                  <button v-if="point.owner_id === authStore.user?.id || authStore.isAdmin()" @click="store.openModal(point)"
                    class="px-4 py-2 text-[10px] font-black uppercase tracking-widest rounded-xl transition-all text-primary bg-primary/10 hover:bg-primary/20">
                    Edit
                  </button>
                  <button v-if="point.owner_id === authStore.user?.id || authStore.isAdmin()" @click="handleDelete(point)"
                    class="px-4 py-2 text-[10px] font-black uppercase tracking-widest rounded-xl transition-all text-rose-600 bg-rose-50 hover:bg-rose-100">
                    Hapus
                  </button>
                  <button @click="router.push(`/blog/${point.id}`)"
                    class="px-4 py-2 text-[10px] font-black uppercase tracking-widest rounded-xl transition-all text-white bg-primary hover:bg-emerald-600 shadow-md shadow-primary/20">
                    Baca Blog
                  </button>
                </div>
              </td>

            </tr>
            <tr v-if="tabularPoints.length === 0">
              <td colspan="6" class="py-20 text-center text-gray-400 text-[10px] font-black uppercase tracking-[0.2em] italic">Data tidak ditemukan.</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>

</template>
