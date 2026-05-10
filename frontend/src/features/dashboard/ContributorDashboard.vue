<script setup lang="ts">
import { onMounted, computed } from 'vue'
import { useMapPointsStore } from '@/stores/mapPoints'
import { useMapUIStore } from '@/stores/mapUI'
import { useAuthStore } from '@/stores/auth'

const store = useMapPointsStore()
const uiStore = useMapUIStore()
const authStore = useAuthStore()

onMounted(async () => {
  await store.fetchCategories()
  await store.fetchPoints(false, true)
})


const myPoints = computed(() => {
  return store.points.filter(p => p.owner_id === authStore.user?.id)
})

const getStatusBadge = (status?: string) => {
  switch (status) {
    case 'approved': return 'bg-emerald-100 text-emerald-700 border-emerald-200'
    case 'pending': return 'bg-amber-100 text-amber-700 border-amber-200'
    case 'rejected': return 'bg-rose-100 text-rose-700 border-rose-200'
    default: return 'bg-slate-100 text-slate-600 border-slate-200'
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
      <!-- Header -->
      <div class="px-6 py-5 border-b border-gray-100 flex items-center justify-between bg-gray-50/40 shrink-0">
        <div>
          <h2 class="text-xl font-extrabold text-gray-900 tracking-tight">Dashboard Kontributor Budaya Bali</h2>
          <p class="text-xs text-gray-500 mt-1 font-medium">Kelola dan pantau status pengajuan objek wisata Anda.</p>
        </div>
        <button @click="store.openModal()"
          class="flex items-center gap-2 px-4 py-2 bg-primary hover:bg-primary-dark text-white font-bold rounded-xl text-xs transition-all active:scale-95 shadow-lg shadow-primary/20">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="3">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
          </svg>
          Tambah Objek Baru
        </button>
      </div>

      <!-- Content Scrollable -->
      <div class="flex-1 overflow-auto p-6 space-y-8">
        <!-- Stats Summary -->
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <div v-for="stat in [
            { label: 'Total Pengajuan', count: myPoints.length, color: 'primary' },
            { label: 'Disetujui', count: myPoints.filter(p => p.status === 'approved').length, color: 'emerald-500' },
            { label: 'Menunggu', count: myPoints.filter(p => p.status === 'pending').length, color: 'amber-500' },
            { label: 'Draf / Ditolak', count: myPoints.filter(p => p.status === 'draft' || p.status === 'rejected' || !p.status).length, color: 'gray-500' }
          ]" :key="stat.label" class="bg-white p-5 rounded-2xl border border-gray-100 shadow-sm group hover:border-primary/30 transition-all duration-300">
            <p class="text-[10px] font-black text-gray-400 uppercase tracking-widest">{{ stat.label }}</p>
            <p class="text-3xl font-black text-gray-900 mt-2 tracking-tighter">{{ stat.count }}</p>
          </div>
        </div>

        <!-- Table -->
        <div class="bg-white rounded-3xl border border-gray-100 shadow-xl shadow-gray-200/40 overflow-hidden flex flex-col">
          <table class="w-full text-left border-collapse min-w-max">
            <thead class="bg-gray-50/50 sticky top-0 z-10 border-b border-gray-100">
              <tr>
                <th class="px-6 py-4 text-[10px] font-black text-gray-400 uppercase tracking-widest">Objek</th>
                <th class="px-6 py-4 text-[10px] font-black text-gray-400 uppercase tracking-widest">Kategori</th>
                <th class="px-6 py-4 text-[10px] font-black text-gray-400 uppercase tracking-widest">Status</th>
                <th class="px-6 py-4 text-[10px] font-black text-gray-400 uppercase tracking-widest">Update Terakhir</th>
                <th class="px-6 py-4 text-[10px] font-black text-gray-400 uppercase tracking-widest text-right">Aksi</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-50">
              <tr v-for="point in myPoints" :key="point.id" class="hover:bg-blue-50/40 transition-colors group">
                <td class="px-6 py-5">
                  <div class="font-black text-gray-900 text-sm tracking-tight group-hover:text-primary transition-colors">{{ point.name }}</div>
                  <div class="text-[10px] text-gray-400 font-bold uppercase tracking-tighter mt-1">{{ point.latitude.toFixed(4) }}, {{ point.longitude.toFixed(4) }}</div>
                </td>
                <td class="px-6 py-5">
                  <span class="text-[10px] font-black text-gray-500 bg-gray-100 px-3 py-1.5 rounded-xl uppercase tracking-wider shadow-sm">{{ getTypeName(point.category_id) }}</span>
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
                <td class="px-6 py-5 text-right">
                  <button @click="store.openModal(point)"
                    class="px-4 py-2 bg-gray-900 text-white hover:bg-primary text-[10px] font-black uppercase tracking-widest rounded-xl transition-all duration-300 shadow-md shadow-gray-200">
                    Edit / Detail
                  </button>
                </td>
              </tr>
              <tr v-if="myPoints.length === 0">
                <td colspan="5" class="px-6 py-20 text-center text-gray-400 text-sm italic font-medium">
                  Belum ada kontribusi yang ditambahkan.
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
