<script setup lang="ts">
import { useRouter } from 'vue-router'
import FacilityTable from '@/components/FacilityTable.vue'
import type { Facility } from '@/stores/facility'

const router = useRouter()

const handlingEdit = (facility: Facility) => {
  router.push(`/edit/${facility.id}`)
}
</script>

<template>
  <div class="max-w-7xl mx-auto space-y-6">
    <div class="flex items-center justify-between">
      <button
        @click="router.push(`/add`)"
        class="bg-primary hover:bg-primary/90 text-surface font-medium px-4 py-2 rounded-lg shadow-sm text-sm transition"
      >
        + Tambah Data
      </button>
    </div>

    <div class="bg-white rounded-xl shadow-sm border border-text/10 p-6">
      <FacilityTable @edit="handlingEdit" />
    </div>
  </div>
</template>
