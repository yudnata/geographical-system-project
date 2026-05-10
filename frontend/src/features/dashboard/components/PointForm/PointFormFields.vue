<script setup lang="ts">
import { useMapPointsStore, type GeoPoint } from '@/stores/mapPoints'

const props = defineProps<{
  modelValue: Partial<GeoPoint>
}>()

const emit = defineEmits(['update:modelValue'])
const store = useMapPointsStore()

// Helper to update individual fields with type safety
const updateField = <K extends keyof GeoPoint>(field: K, value: GeoPoint[K]) => {
  emit('update:modelValue', { ...props.modelValue, [field]: value })
}
</script>

<template>
  <div class="p-6 overflow-y-auto w-full space-y-5">
    <!-- Coordination Info -->
    <div class="bg-primary/5 text-primary text-[10px] px-4 py-3 rounded-xl flex gap-4 font-black border border-primary/10 shadow-sm">
      <div class="flex items-center gap-1.5">
        <span class="opacity-40 uppercase tracking-tighter">Lat:</span> 
        {{ modelValue.latitude?.toFixed(6) || 0 }}
      </div>
      <div class="flex items-center gap-1.5">
        <span class="opacity-40 uppercase tracking-tighter">Lng:</span> 
        {{ modelValue.longitude?.toFixed(6) || 0 }}
      </div>
    </div>

    <!-- Nama Bangunan -->
    <div class="space-y-1.5">
      <label class="text-[11px] font-black text-gray-500 uppercase tracking-wider ml-1">Nama Bangunan</label>
      <div class="relative group">
        <span class="absolute inset-y-0 left-3.5 flex items-center text-gray-400 group-focus-within:text-primary transition-colors">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 21h19.5m-18-18v18m10.5-18v18m6-13.5V21M6.75 6.75h.75m-.75 3h.75m-.75 3h.75m3-6h.75m-.75 3h.75m-.75 3h.75M6.75 21v-3.375c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21M3 3h12m-.75 4.5H21m-3.75 3.75h.008v.008h-.008v-.008zm0 3h.008v.008h-.008v-.008zm0 3h.008v.008h-.008v-.008z" />
          </svg>
        </span>
        <input 
          :value="modelValue.name" 
          @input="e => updateField('name', (e.target as HTMLInputElement).value)"
          type="text" 
          class="w-full pl-10 pr-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:bg-white focus:border-primary focus:ring-4 focus:ring-primary/10 transition-all text-sm font-bold text-gray-800 outline-none"
          placeholder="Contoh: Gedung Kantor Utama">
      </div>
    </div>

    <!-- Kategori -->
    <div class="space-y-1.5">
      <label class="text-[11px] font-black text-gray-500 uppercase tracking-wider ml-1">Tipe / Kategori</label>
      <div class="relative">
        <select 
          :value="modelValue.type_id"
          @change="e => updateField('type_id', parseInt((e.target as HTMLSelectElement).value))"
          class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:bg-white focus:border-primary focus:ring-4 focus:ring-primary/10 transition-all text-sm font-bold text-gray-700 appearance-none outline-none">
          <option v-for="type in store.objectTypes" :value="type.id" :key="type.id">{{ type.name.toUpperCase() }}</option>
        </select>
        <span class="absolute inset-y-0 right-3 flex items-center pointer-events-none text-gray-400">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="m19.5 8.25-7.5 7.5-7.5-7.5" />
          </svg>
        </span>
      </div>
    </div>

    <!-- Alamat -->
    <div class="space-y-1.5">
      <label class="text-[11px] font-black text-gray-500 flex items-center justify-between uppercase tracking-wider ml-1">
        Alamat Spasial
        <span class="text-[9px] bg-emerald-100 text-emerald-700 px-2 py-0.5 rounded-full shadow-sm normal-case font-black">AUTO-DETECT</span>
      </label>
      <textarea 
        :value="modelValue.address"
        @input="e => updateField('address', (e.target as HTMLTextAreaElement).value)"
        rows="3"
        class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:bg-white focus:border-primary focus:ring-4 focus:ring-primary/10 transition-all text-xs font-bold text-gray-600 leading-relaxed outline-none"
        placeholder="Alamat akan terdeteksi otomatis saat Anda memilih titik di peta..."></textarea>
    </div>

    <!-- Grid: Tahun & Kepemilikan -->
    <div class="grid grid-cols-2 gap-5">
      <div class="space-y-1.5">
        <label class="text-[11px] font-black text-gray-500 uppercase tracking-wider ml-1">Tahun Berdiri</label>
        <input 
          :value="modelValue.tahun_berdiri"
          @input="e => updateField('tahun_berdiri', parseInt((e.target as HTMLInputElement).value))"
          type="number"
          class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:bg-white focus:border-primary focus:ring-4 focus:ring-primary/10 transition-all text-sm font-bold text-gray-800 outline-none">
      </div>

      <div class="space-y-1.5">
        <label class="text-[11px] font-black text-gray-500 uppercase tracking-wider ml-1">Status Milik</label>
        <div class="relative">
          <select 
            :value="modelValue.status_kepemilikan"
            @change="e => updateField('status_kepemilikan', (e.target as HTMLSelectElement).value)"
            class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl focus:bg-white focus:border-primary focus:ring-4 focus:ring-primary/10 transition-all text-sm font-bold text-gray-700 appearance-none outline-none">
            <option value="Pemerintah">PEMERINTAH</option>
            <option value="Swasta">SWASTA</option>
            <option value="Pribadi">PRIBADI</option>
            <option value="Yayasan">YAYASAN</option>
          </select>
          <span class="absolute inset-y-0 right-3 flex items-center pointer-events-none text-gray-400">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5" stroke="currentColor" class="w-4 h-4">
              <path stroke-linecap="round" stroke-linejoin="round" d="m19.5 8.25-7.5 7.5-7.5-7.5" />
            </svg>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>
