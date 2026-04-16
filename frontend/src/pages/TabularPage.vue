<script setup lang="ts">
import { onMounted } from 'vue'
import { useMapPointsStore } from '@/stores/mapPoints'

const store = useMapPointsStore()

onMounted(async () => {
  await store.fetchPoints()
})

const getTypeName = (typeId: number) => {
  const type = store.objectTypes.find(t => t.id === typeId)
  return type ? type.name : 'Unknown'
}
</script>

<template>
  <div class="h-full flex flex-col bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden mx-4 my-6">
    <div class="px-6 py-5 border-b border-gray-100 flex items-center justify-between bg-gray-50/50">
      <div>
        <h2 class="text-xl font-bold text-gray-800">Daftar Bangunan & Infrastruktur</h2>
        <p class="text-sm text-gray-500 mt-1">Sistem melihat data secara tabular. Anda dapat melihat koordinat spasial di sini.</p>
      </div>
      <div>
        <input v-model="store.searchQuery" type="text" placeholder="Cari bangunan atau alamat..."
          class="px-4 py-2 w-64 bg-white border border-gray-200 rounded-lg focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm shadow-sm" />
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
          <tr v-for="point in store.filteredPoints" :key="point.id" class="hover:bg-blue-50/40 transition-colors group">
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
              <button @click="store.openModal(point)"
                class="px-3 py-1.5 text-xs font-bold text-primary bg-primary/10 rounded-lg hover:bg-primary/20 transition-colors opacity-0 group-hover:opacity-100">Edit Data</button>
            </td>
          </tr>
          <tr v-if="store.filteredPoints.length === 0">
            <td colspan="7" class="py-12 text-center text-gray-400 text-sm">Tidak ada data bangunan yang ditemukan.</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
