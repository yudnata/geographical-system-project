<script setup lang="ts">
import { onMounted, computed, ref } from 'vue'
import { useMapPointsStore } from '@/stores/mapPoints'
import { useMapUIStore } from '@/stores/mapUI'
import { useAuthStore } from '@/stores/auth'

const store = useMapPointsStore()
const uiStore = useMapUIStore()
const authStore = useAuthStore()

const activeFilter = ref<'all' | 'draft' | 'pending' | 'approved' | 'rejected'>('all')

onMounted(async () => {
  await store.fetchCategories()
  await store.fetchPoints(false, true)
  uiStore.filterMyPoints = true
  uiStore.statusFilter = 'all'
})

const filteredPoints = computed(() => {
  const points = store.points.filter(p => p.owner_id === authStore.user?.id)
  if (activeFilter.value === 'all') return points
  if (activeFilter.value === 'draft') return points.filter(p => p.status === 'draft' || !p.status)
  return points.filter(p => p.status === activeFilter.value)
})

const getStatusBadge = (status?: string) => {
  switch (status) {
    case 'approved': return 'text-emerald-700 border-none'
    case 'pending': return 'text-amber-700 border-none'
    case 'rejected': return 'text-rose-700 border-none'
    case 'draft': return 'text-slate-600 border-none'
    default: return 'text-slate-600 border-none'
  }
}

const getTypeName = (categoryId: number) => {
  return store.objectTypes.find(t => t.id === categoryId)?.name || 'Kategori'
}
</script>

<template>
  <div :class="[
    'h-full w-full flex flex-col transition-[padding] duration-500 ease-in-out overflow-hidden',
    uiStore.isSidebarExpanded ? 'pl-[288px]' : 'pl-24',
    'pt-6 pb-6 pr-6'
  ]">
    <div class="flex-1 flex flex-col bg-white rounded-2xl shadow-[0_8px_32px_rgba(0,0,0,0.06)] border border-gray-100 overflow-hidden">
      <div class="px-6 py-4 border-b border-gray-100 flex items-center justify-between bg-gray-50/40 shrink-0">
        <div class="flex items-center gap-2 bg-white/60 p-1 rounded-2xl border border-gray-100 shadow-sm">
          <button v-for="f in [
            { id: 'all', label: 'Semua' },
            { id: 'draft', label: 'Draft' },
            { id: 'pending', label: 'Menunggu' },
            { id: 'approved', label: 'Terverifikasi' },
            { id: 'rejected', label: 'Ditolak' }
          ]" :key="f.id" @click="activeFilter = f.id as any" :class="[
            'px-5 py-2 text-[10px] font-black uppercase tracking-widest rounded-xl transition-all duration-300',
            activeFilter === f.id
              ? 'bg-primary text-white shadow-lg shadow-primary/30 scale-105'
              : 'text-gray-400 hover:text-gray-600 hover:bg-gray-100'
          ]">
            {{ f.label }}
          </button>
        </div>
      </div>
      <div class="flex-1 overflow-auto p-6 space-y-8">
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <div v-for="stat in [
            { label: 'Total Pengajuan', count: store.points.filter(p => p.owner_id === authStore.user?.id).length, color: 'primary' },
            { label: 'Disetujui', count: store.points.filter(p => p.owner_id === authStore.user?.id && p.status === 'approved').length, color: 'emerald-500' },
            { label: 'Menunggu', count: store.points.filter(p => p.owner_id === authStore.user?.id && p.status === 'pending').length, color: 'amber-500' },
            { label: 'Draf', count: store.points.filter(p => p.owner_id === authStore.user?.id && (p.status === 'draft' || !p.status)).length, color: 'gray-500' }
          ]" :key="stat.label" class="bg-white p-5 rounded-2xl border border-gray-100 shadow-sm group hover:border-primary/30 transition-all duration-300">
            <p class="text-[10px] font-black text-gray-400 uppercase tracking-widest">{{ stat.label }}</p>
            <p class="text-3xl font-black text-gray-900 mt-2 tracking-tighter">{{ stat.count }}</p>
          </div>
        </div>
        <div class="bg-white rounded-3xl border border-gray-100 shadow-xl shadow-gray-200/40 overflow-hidden flex flex-col">
          <table class="w-full text-left border-collapse min-w-max">
            <thead class="bg-gray-50/50 sticky top-0 z-10 border-b border-gray-100">
              <tr>
                <th class="px-6 py-4 text-[10px] font-black text-gray-400 uppercase tracking-widest">Objek</th>
                <th class="px-6 py-4 text-[10px] font-black text-gray-400 uppercase tracking-widest">Kategori</th>
                <th class="px-6 py-4 text-[10px] font-black text-gray-400 uppercase tracking-widest">Status</th>
                <th class="px-6 py-4 text-[10px] font-black text-gray-400 uppercase tracking-widest">Update Terakhir</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-50">
              <tr v-for="point in filteredPoints" :key="point.id" class="hover:bg-blue-50/40 transition-colors group">
                <td class="px-6 py-5">
                  <div class="font-black text-gray-900 text-sm tracking-tight group-hover:text-primary transition-colors">{{ point.name }}</div>
                  <div class="text-[10px] text-gray-400 font-bold tracking-tighter mt-1 line-clamp-1 max-w-[200px]" :title="point.address">
                    {{ point.address || 'Alamat belum diisi' }}
                  </div>
                </td>
                <td class="px-6 py-5">
                  <span class="text-[10px] font-black text-gray-500 bg-gray-100 px-3 py-1.5 rounded-xl uppercase tracking-wider shadow-sm">{{ getTypeName(point.category_id || 0) }}</span>
                </td>

                <td class="px-6 py-5">
                  <div :class="['inline-flex items-center gap-2 px-3 py-1.5 rounded-full border text-[10px] font-black uppercase tracking-widest', getStatusBadge(point.status)]">
                    <span class="w-1.5 h-1.5 rounded-full bg-current animate-pulse"></span>
                    {{ point.status || 'Draft' }}
                  </div>
                  <p v-if="point.status === 'rejected' && point.rejection_note" class="text-[10px] text-rose-500 mt-2 font-bold italic max-w-[200px] leading-relaxed">
                    Ket: {{ point.rejection_note }}
                  </p>
                </td>
                <td class="px-6 py-5 text-xs text-gray-500 font-bold uppercase tracking-tighter">
                  {{ new Date(point.updated_at || '').toLocaleDateString('id-ID', { day: '2-digit', month: 'short', year: 'numeric' }) }}
                </td>
              </tr>
              <tr v-if="filteredPoints.length === 0">
                <td colspan="4" class="px-6 py-20 text-center text-gray-400 text-sm italic font-medium">
                  Tidak ada data dengan status ini.
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
