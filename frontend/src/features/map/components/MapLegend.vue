<script setup lang="ts">
import { ref } from 'vue'
import { useMapPointsStore } from '@/stores/mapPoints'

const store = useMapPointsStore()
const isExpanded = ref(true)

const ICON_MAP: Record<string, string> = {
  'temple': 'M12 2l-4 3h8l-4-3zm0 4l-6 4h12l-6-4zm0 5l-8 6h16l-8-6zm-2 6h4v5h-4z',
  'waterfall': 'M12 2v16M9 6c-2 0-3 1-3 3v13M15 6c2 0 3 1 3 3v13M12 22v-2',
  'beach': 'M12 10a4 4 0 100-8 4 4 0 000 8z M2 17c2 0 3-1 5-1s3 1 5 1 3-1 5-1 3 1 5 1 M2 21c2 0 3-1 5-1s3 1 5 1 3-1 5-1 3 1 5 1',
  'mountain': 'M3 20l9-16 9 16H3z M8 11l4-3 4 3',
  'hill': 'M2 20c4-6 10-6 14 0 M12 20c4-4 8-4 10 0',
  'lake': 'M12 21c-5 0-9-2-9-5s4-5 9-5 9 2 9 5-4 5-9 5z M12 17a3 3 0 110-6 3 3 0 010 6z',
  'forest': 'M12 2l3 5H9l3-5zm0 6l4 7H8l4-7zm0 7v5',
  'village': 'M3 9l9-7 9 7v11a2 2 0 01-2 2H5a2 2 0 01-2-2V9z',
  'museum': 'M4 22V10h16v12M2 10l10-8 10 8M9 14v4M15 14v4',
  'market': 'M3 3h18v2H3V3zm3 4v14h12V7H6zM10 7v14M14 7v14'
}


const getIconPath = (iconName?: string) => {
  if (!iconName) return 'M12 2v20M2 12h20'
  return ICON_MAP[iconName] || iconName
}

const colors = [

  '#2D6A4F', '#0077B6', '#D62828', '#F4A261', '#9D4EDD',
  '#FFB703', '#E63946', '#219EBC', '#023047', '#FB8500',
  '#8ECAE6', '#606C38'
]
</script>

<template>
  <div
    class="absolute bottom-16 right-6 z-[1000] bg-white/90 backdrop-blur-md rounded-2xl shadow-[0_8px_32px_rgba(0,0,0,0.12)] border border-white/20 min-w-[140px] overflow-hidden transition-all duration-300">
    <div @click="isExpanded = !isExpanded" class="px-3 py-2.5 bg-gray-50/50 border-b border-gray-100 flex items-center justify-between cursor-pointer hover:bg-gray-100 transition-colors group">
      <h4 class="text-[9px] font-black text-gray-400 uppercase tracking-[0.2em]">Legenda Tipe</h4>
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="3" stroke="currentColor"
        :class="['w-3 h-3 text-gray-400 group-hover:text-primary transition-transform duration-300', isExpanded ? '' : '-rotate-180']">
        <path stroke-linecap="round" stroke-linejoin="round" d="m19.5 8.25-7.5 7.5-7.5-7.5" />
      </svg>
    </div>

    <div v-show="isExpanded" class="p-3 space-y-2.5 max-h-[250px] overflow-y-auto transition-all duration-300">
      <div v-for="(type, index) in store.objectTypes" :key="type.id" class="flex items-center gap-2.5 group cursor-default">
        <div class="w-5 h-5 rounded-lg flex items-center justify-center text-white p-1 shadow-sm transition-transform group-hover:scale-110"
          :style="{ backgroundColor: colors[index % colors.length] }">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round" class="w-full h-full">
            <path :d="getIconPath(type.icon)" />
          </svg>

        </div>
        <span class="text-[10px] font-bold text-gray-600 tracking-tight uppercase">{{ type.name }}</span>
      </div>
    </div>
  </div>
</template>
