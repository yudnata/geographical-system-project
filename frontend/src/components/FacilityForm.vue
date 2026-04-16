<script setup lang="ts">
import { ref, watch, computed } from 'vue'
import type { Facility } from '@/stores/facility'

const props = defineProps<{
  editData: Facility | null
}>()

const emit = defineEmits<{
  (e: 'save', data: Omit<Facility, 'id'>): void
  (e: 'update', data: Facility): void
  (e: 'cancel'): void
}>()

const isEditing = computed(() => !!props.editData)

const form = ref<Omit<Facility, 'id'>>({
  name: '',
  type: 'Puskesmas',
  address: '',
  district: '',
  phone: '',
  hours: '',
  lat: 0,
  lng: 0,
})

const resetForm = () => {
  form.value = {
    name: '',
    type: 'Puskesmas',
    address: '',
    district: '',
    phone: '',
    hours: '',
    lat: 0,
    lng: 0,
  }
}

const cancelEdit = () => {
  resetForm()
  emit('cancel')
}

watch(
  () => props.editData,
  (newData) => {
    if (newData) {
      form.value = { ...newData }
    } else {
      resetForm()
    }
  },
  { immediate: true, deep: true },
)

const submitForm = () => {
  if (isEditing.value && props.editData) {
    emit('update', { ...form.value, id: props.editData.id })
  } else {
    emit('save', { ...form.value })
  }
  resetForm()
}
</script>

<template>
  <div
    class="bg-white rounded-xl shadow-sm border border-text/10 p-6 flex flex-col h-full overflow-hidden"
  >
    <h2 class="text-xl font-bold mb-4 border-b pb-2 flex items-center justify-between shrink-0">
      {{ isEditing ? 'Edit Fasilitas' : 'Tambah Fasilitas Baru' }}
      <span
        v-if="isEditing"
        class="text-xs bg-danger/10 text-danger px-2 py-1 rounded cursor-pointer hover:bg-danger/20 transition"
        @click="cancelEdit"
        >Batal Edit</span
      >
    </h2>
    <form @submit.prevent="submitForm" class="flex-1 overflow-auto space-y-4 pr-2">
      <div>
        <label class="block text-sm font-medium text-text/80">Nama Fasilitas</label>
        <input
          v-model="form.name"
          type="text"
          required
          class="mt-1 block w-full rounded border-text/20 border p-2 focus:ring-primary focus:border-primary text-sm"
        />
      </div>
      <div>
        <label class="block text-sm font-medium text-text/80">Jenis</label>
        <select
          v-model="form.type"
          required
          class="mt-1 block w-full rounded border-text/20 border p-2 focus:ring-primary focus:border-primary text-sm"
        >
          <option value="Puskesmas">Puskesmas</option>
          <option value="Klinik">Klinik</option>
          <option value="Rumah Sakit">Rumah Sakit</option>
          <option value="Apotek">Apotek</option>
        </select>
      </div>
      <div>
        <label class="block text-sm font-medium text-text/80">Kecamatan</label>
        <input
          v-model="form.district"
          type="text"
          required
          class="mt-1 block w-full rounded border-text/20 border p-2 focus:ring-primary focus:border-primary text-sm"
        />
      </div>
      <div>
        <label class="block text-sm font-medium text-text/80">Alamat Lengkap</label>
        <textarea
          v-model="form.address"
          required
          rows="2"
          class="mt-1 block w-full rounded border-text/20 border p-2 focus:ring-primary focus:border-primary text-sm"
        ></textarea>
      </div>
      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-text/80">Telepon</label>
          <input
            v-model="form.phone"
            type="text"
            placeholder="(0361) xxx"
            class="mt-1 block w-full rounded border-text/20 border p-2 focus:ring-primary focus:border-primary text-sm"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-text/80">Jam Operasi</label>
          <input
            v-model="form.hours"
            type="text"
            placeholder="Mis: 08:00 - 20:00"
            class="mt-1 block w-full rounded border-text/20 border p-2 focus:ring-primary focus:border-primary text-sm"
          />
        </div>
      </div>
      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-text/80">Latitude</label>
          <input
            v-model="form.lat"
            type="number"
            step="any"
            required
            class="mt-1 block w-full rounded border-text/20 border p-2 bg-surface focus:ring-primary focus:border-primary text-sm"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-text/80">Longitude</label>
          <input
            v-model="form.lng"
            type="number"
            step="any"
            required
            class="mt-1 block w-full rounded border-text/20 border p-2 bg-surface focus:ring-primary focus:border-primary text-sm"
          />
        </div>
      </div>
      <p
        class="text-xs text-primary bg-primary/10 p-2 rounded border border-primary/20 flex gap-2 items-start"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="w-4 h-4 mt-0.5 shrink-0"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M15.042 21.672L13.684 16.6m0 0l-2.51 2.225.569-9.47 5.227 7.917-3.286-.672zm-7.518-.267A8.25 8.25 0 1120.25 10.5M8.288 14.212A5.25 5.25 0 1117.25 10.5"
          />
        </svg>
        Klik pada peta untuk mengisi koordinat secara otomatis.
      </p>

      <div class="pt-4 shrink-0 mt-auto">
        <button
          type="submit"
          class="w-full bg-primary text-white rounded py-2.5 px-4 hover:bg-primary/90 transition font-bold shadow-sm"
        >
          {{ isEditing ? 'Simpan Perubahan' : 'Tambah Data' }}
        </button>
      </div>
    </form>
  </div>
</template>
