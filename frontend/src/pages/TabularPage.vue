<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { useMapPointsStore } from '@/stores/mapPoints'
import { useMapUIStore } from '@/stores/mapUI'
import { useAuthStore } from '@/stores/auth'

const store = useMapPointsStore()
const uiStore = useMapUIStore()
const authStore = useAuthStore()

onMounted(async () => {
  uiStore.filterMyPoints = true
  await store.fetchPoints()
})

const localSearchQuery = ref('')

const tabularPoints = computed(() => {
  return store.points.filter((point) => {
    const matchSearch =
      point.name.toLowerCase().includes(localSearchQuery.value.toLowerCase()) ||
      (point.address && point.address.toLowerCase().includes(localSearchQuery.value.toLowerCase()))
    const matchType = !uiStore.filterTypeId || point.type_id === uiStore.filterTypeId
    const matchOwner = !uiStore.filterMyPoints || point.owner_id === authStore.user?.id

    return matchSearch && matchType && matchOwner
  })
})

const getTypeName = (typeId: number) => {
  const type = store.objectTypes.find(t => t.id === typeId)
  return type ? type.name : 'Unknown'
}
</script>

<template>
  <div :class="[
    'h-full w-full flex flex-col transition-[padding] duration-500 ease-in-out overflow-hidden',
    uiStore.isSidebarExpanded ? 'pl-[288px]' : 'pl-24',
    'pt-24 pb-6 pr-6'
  ]">
    <div class="flex-1 flex flex-col bg-white rounded-2xl shadow-[0_8px_32px_rgba(0,0,0,0.06)] border border-gray-100 overflow-hidden">
      <div class="px-6 py-5 border-b border-gray-100 flex items-center justify-between bg-gray-50/40">
        <div>
          <h2 class="text-xl font-extrabold text-gray-900 tracking-tight">Daftar Bangunan & Infrastruktur</h2>
          <p class="text-xs text-gray-500 mt-1 font-medium">Sistem melihat data secara tabular. Anda dapat melihat koordinat spasial di sini.</p>
        </div>
        <div class="flex items-center gap-4">
          <!-- Toggle Filter My Points -->
          <div class="flex bg-gray-100 p-1 rounded-xl shadow-inner shrink-0">
            <button @click="uiStore.filterMyPoints = false" :class="[!uiStore.filterMyPoints ? 'bg-white text-primary shadow-sm' : 'text-gray-500 hover:text-gray-700']"
              class="px-4 py-1.5 text-xs font-bold rounded-lg transition-all flex items-center gap-2">
              Semua Data
            </button>
            <button @click="uiStore.filterMyPoints = true" :class="[uiStore.filterMyPoints ? 'bg-white text-primary shadow-sm' : 'text-gray-500 hover:text-gray-700']"
              class="px-4 py-1.5 text-xs font-bold rounded-lg transition-all flex items-center gap-2">
              Kontribusi Saya
            </button>
          </div>

          <div class="relative group">
            <span class="absolute inset-y-0 left-3 flex items-center text-gray-400 group-focus-within:text-primary transition-colors">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-4 h-4">
                <path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
              </svg>
            </span>
            <input v-model="localSearchQuery" type="text" placeholder="Cari bangunan atau alamat..."
              class="pl-9 pr-4 py-2 w-64 bg-white border border-gray-200 rounded-xl focus:border-primary focus:ring-4 focus:ring-primary/10 transition-all text-xs font-bold shadow-sm" />
          </div>
        </div>
      </div>

      <div class="flex-1 overflow-auto">
        <table class="w-full text-left border-collapse min-w-max">
          <thead class="bg-gray-50 sticky top-0 shadow-sm outline outline-1 outline-gray-200 z-10">
            <tr>
              <th class="py-3 px-6 text-xs font-semibold text-gray-600 uppercase tracking-wide">ID</th>
              <th class="py-3 px-6 text-xs font-semibold text-gray-600 uppercase tracking-wide">Nama Bangunan</th>
              <th class="py-3 px-6 text-xs font-semibold text-gray-600 uppercase tracking-wide">Tipe</th>
              <th class="py-3 px-6 text-xs font-semibold text-gray-600 uppercase tracking-wide">Tahun Berdiri</th>
              <th class="py-3 px-6 text-xs font-semibold text-gray-600 uppercase tracking-wide">Koordinat</th>
              <th class="py-3 px-6 text-xs font-semibold text-gray-600 uppercase tracking-wide">Status</th>
              <th class="py-3 px-6 text-xs font-semibold text-gray-600 uppercase tracking-wide text-center">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-100">
            <tr v-for="point in tabularPoints" :key="point.id" class="hover:bg-blue-50/40 transition-colors group">
              <td class="py-4 px-6 text-sm text-gray-500 font-mono">#{{ point.id }}</td>
              <td class="py-4 px-6">
                <div class="font-bold text-gray-800">{{ point.name }}</div>
                <div class="text-xs text-gray-500 line-clamp-1 max-w-xs mt-0.5" :title="point.address">{{ point.address }}</div>
              </td>
              <td class="py-4 px-6">
                <span class="px-2.5 py-1 text-[10px] font-bold tracking-wider rounded-full bg-blue-100 text-blue-700 uppercase">
                  {{ getTypeName(point.type_id) }}
                </span>
              </td>
              <td class="py-4 px-6 text-sm text-gray-600 font-medium">{{ point.tahun_berdiri || '-' }}</td>
              <td class="py-4 px-6 text-xs font-mono text-gray-500">
                <div class="flex flex-col gap-0.5">
                  <span>Lat: {{ point.latitude?.toFixed(5) }}</span>
                  <span>Lng: {{ point.longitude?.toFixed(5) }}</span>
                </div>
              </td>
              <td class="py-4 px-6">
                <span v-if="point.is_active" class="text-emerald-600 text-xs font-bold flex items-center gap-1.5"><span class="w-1.5 h-1.5 rounded-full bg-emerald-500"></span> Aktif</span>
                <span v-else class="text-red-500 text-xs font-bold flex items-center gap-1.5"><span class="w-1.5 h-1.5 rounded-full bg-red-500"></span> Nonaktif</span>
              </td>
              <td class="py-4 px-6 text-center">
                <button @click="point.owner_id === authStore.user?.id ? store.openModal(point) : null" :disabled="point.owner_id !== authStore.user?.id" :class="[
                  'px-3 py-1.5 text-xs font-bold rounded-lg transition-all',
                  point.owner_id === authStore.user?.id
                    ? 'text-primary bg-primary/10 hover:bg-primary/20'
                    : 'text-gray-400 bg-gray-100 cursor-not-allowed'
                ]">
                  {{ point.owner_id === authStore.user?.id ? 'Edit Data' : 'Hanya Lihat' }}
                </button>
              </td>
            </tr>
            <tr v-if="tabularPoints.length === 0">
              <td colspan="7" class="py-12 text-center text-gray-400 text-sm">Tidak ada data bangunan yang ditemukan.</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
