<script setup lang="ts">
import { onMounted, computed } from 'vue'
import { useMapPointsStore } from '@/stores/mapPoints'
import { useAuthStore } from '@/stores/auth'

const store = useMapPointsStore()
const authStore = useAuthStore()

onMounted(async () => {
  await store.fetchCategories()
  await store.fetchPoints()
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

const getTypeName = (typeId: number) => {
  return store.objectTypes.find(t => t.id === typeId)?.name || 'Kategori'
}
</script>

<template>
  <div class="p-8 space-y-8">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-slate-900">Kontribusi Saya</h2>
        <p class="text-sm text-slate-500 mt-1">Kelola dan pantau status pengajuan objek wisata Anda.</p>
      </div>
      <button @click="store.openModal()"
        class="flex items-center gap-2 px-5 py-2.5 bg-indigo-600 text-white text-sm font-bold rounded-xl hover:bg-indigo-700 transition-all shadow-lg shadow-indigo-100">
        <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        Tambah Objek Baru
      </button>
    </div>

    <!-- Stats Summary -->
    <div class="grid grid-cols-4 gap-6">
      <div v-for="stat in [
        { label: 'Total', count: myPoints.length, color: 'indigo' },
        { label: 'Disetujui', count: myPoints.filter(p => p.status === 'approved').length, color: 'emerald' },
        { label: 'Menunggu', count: myPoints.filter(p => p.status === 'pending').length, color: 'amber' },
        { label: 'Draft', count: myPoints.filter(p => p.status === 'draft' || !p.status).length, color: 'slate' }
      ]" :key="stat.label" class="bg-white p-5 rounded-2xl border border-slate-100 shadow-sm">
        <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ stat.label }}</p>
        <p class="text-2xl font-black text-slate-900 mt-1">{{ stat.count }}</p>
      </div>
    </div>

    <!-- Table -->
    <div class="bg-white rounded-2xl border border-slate-100 shadow-sm overflow-hidden">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-slate-50/50 border-b border-slate-100">
            <th class="px-6 py-4 text-[10px] font-bold text-slate-400 uppercase tracking-widest">Objek</th>
            <th class="px-6 py-4 text-[10px] font-bold text-slate-400 uppercase tracking-widest">Kategori</th>
            <th class="px-6 py-4 text-[10px] font-bold text-slate-400 uppercase tracking-widest">Status</th>
            <th class="px-6 py-4 text-[10px] font-bold text-slate-400 uppercase tracking-widest">Update Terakhir</th>
            <th class="px-6 py-4 text-[10px] font-bold text-slate-400 uppercase tracking-widest text-right">Aksi</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-50">
          <tr v-for="point in myPoints" :key="point.id" class="hover:bg-slate-50/50 transition-colors">
            <td class="px-6 py-4">
              <div class="font-bold text-slate-900 text-sm">{{ point.name }}</div>
              <div class="text-[10px] text-slate-400 font-mono mt-0.5">{{ point.latitude.toFixed(4) }}, {{ point.longitude.toFixed(4) }}</div>
            </td>
            <td class="px-6 py-4">
              <span class="text-[10px] font-bold text-slate-500 bg-slate-100 px-2 py-1 rounded-md uppercase">{{ getTypeName(point.type_id) }}</span>
            </td>
            <td class="px-6 py-4">
              <div :class="['inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full border text-[10px] font-bold uppercase tracking-wider', getStatusBadge(point.status)]">
                <span class="w-1 h-1 rounded-full bg-current"></span>
                {{ point.status || 'Draft' }}
              </div>
              <p v-if="point.status === 'rejected' && point.rejection_note" class="text-[10px] text-rose-500 mt-1 italic max-w-[200px]">
                Ket: {{ point.rejection_note }}
              </p>
            </td>
            <td class="px-6 py-4 text-xs text-slate-500 font-medium">
              {{ new Date(point.updated_at || '').toLocaleDateString('id-ID', { day: '2-digit', month: 'short', year: 'numeric' }) }}
            </td>
            <td class="px-6 py-4 text-right">
              <button @click="store.openModal(point)" class="text-indigo-600 hover:text-indigo-800 text-xs font-bold transition-colors">
                Edit / Detail
              </button>
            </td>
          </tr>
          <tr v-if="myPoints.length === 0">
            <td colspan="5" class="px-6 py-12 text-center text-slate-400 text-sm italic">
              Belum ada kontribusi yang ditambahkan.
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
