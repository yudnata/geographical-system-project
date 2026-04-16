<script setup lang="ts">
import { ref, computed } from 'vue'
import { useFacilityStore } from '@/stores/facility'
import { storeToRefs } from 'pinia'
import type { Facility } from '@/stores/facility'

const emit = defineEmits<{
  (e: 'edit', facility: Facility): void
}>()

const store = useFacilityStore()
const { facilities } = storeToRefs(store)

const filterType = ref('')

const filteredFacilities = computed(() => {
  if (!filterType.value) return facilities.value
  return facilities.value.filter((f) => f.type === filterType.value)
})

const getTypeBadgeClass = (type: string) => {
  switch (type) {
    case 'Puskesmas':
      return 'bg-primary/10 text-primary border-primary/20'
    case 'Klinik':
      return 'bg-accent/10 text-accent border-accent/20'
    case 'Rumah Sakit':
      return 'bg-danger/10 text-danger border-danger/20'
    case 'Apotek':
      return 'bg-text/10 text-text border-text/20'
    default:
      return 'bg-surface text-text/80 border-text/10'
  }
}

const editItem = (item: Facility) => {
  emit('edit', item)
}

const deleteItem = (id: number) => {
  if (
    confirm('Yakin ingin menghapus data fasilitas kesehatan ini? Data tidak dapat dikembalikan.')
  ) {
    store.deleteFacility(id)
  }
}
</script>

<template>
  <section class="bg-white rounded-lg shadow overflow-hidden border border-text/10 pb-2 mt-6">
    <div
      class="p-4 border-b flex flex-col sm:flex-row justify-between items-center gap-4 bg-surface"
    >
      <h2 class="text-lg font-bold text-text">Daftar Fasilitas Kesehatan</h2>
      <div class="flex items-center gap-2">
        <label class="text-sm font-medium text-text/80">Filter Jenis:</label>
        <select
          v-model="filterType"
          class="rounded border-text/20 border py-1.5 px-3 text-sm focus:ring-primary focus:border-primary bg-white shadow-sm"
        >
          <option value="">Semua Fasilitas</option>
          <option value="Puskesmas">Puskesmas</option>
          <option value="Klinik">Klinik</option>
          <option value="Rumah Sakit">Rumah Sakit</option>
          <option value="Apotek">Apotek</option>
        </select>
      </div>
    </div>
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-text/10">
        <thead class="bg-surface">
          <tr>
            <th
              class="px-6 py-3 text-left text-xs font-semibold text-text/70 uppercase tracking-wider"
            >
              Nama & Detail
            </th>
            <th
              class="px-6 py-3 text-left text-xs font-semibold text-text/70 uppercase tracking-wider"
            >
              Jenis
            </th>
            <th
              class="px-6 py-3 text-left text-xs font-semibold text-text/70 uppercase tracking-wider"
            >
              Kecamatan
            </th>
            <th
              class="px-6 py-3 text-left text-xs font-semibold text-text/70 uppercase tracking-wider"
            >
              Telepon
            </th>
            <th
              class="px-6 py-3 text-right text-xs font-semibold text-text/70 uppercase tracking-wider"
            >
              Aksi
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-text/10">
          <tr v-if="filteredFacilities.length === 0">
            <td colspan="5" class="px-6 py-8 text-center text-sm text-text/60 bg-surface">
              Tidak ada data fasilitas yang ditemukan.
            </td>
          </tr>
          <tr
            v-for="item in filteredFacilities"
            :key="item.id"
            class="hover:bg-primary/5 transition"
          >
            <td class="px-6 py-4">
              <div class="text-sm font-bold text-text">{{ item.name }}</div>
              <div class="text-xs text-text/60 mt-1 truncate max-w-xs">{{ item.address }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span
                :class="getTypeBadgeClass(item.type)"
                class="px-2.5 py-1 inline-flex text-xs leading-5 font-semibold rounded-full border border-opacity-20"
              >
                {{ item.type }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-text/80">{{ item.district }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-text/80">
              {{ item.phone || '-' }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <button
                @click="editItem(item)"
                class="text-primary hover:text-primary/80 bg-primary/10 px-3 py-1 rounded mr-2 transition"
              >
                Edit
              </button>
              <button
                @click="deleteItem(item.id)"
                class="text-danger hover:text-danger/80 bg-danger/10 px-3 py-1 rounded transition"
              >
                Hapus
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>
