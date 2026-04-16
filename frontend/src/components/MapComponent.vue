<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue'
import L from 'leaflet'
import { useFacilityStore } from '@/stores/facility'
import { storeToRefs } from 'pinia'

const props = defineProps<{
  flyTo?: { lat: number; lng: number } | null
}>()

const emit = defineEmits<{
  (e: 'map-clicked', coords: { lat: number; lng: number }): void
}>()

const store = useFacilityStore()
const { facilities } = storeToRefs(store)

const mapContainer = ref<HTMLElement | null>(null)
let map: L.Map | null = null
let markerLayer: L.LayerGroup | null = null

const typeColors: Record<string, string> = {
  Puskesmas: '#2D6A4F',
  Klinik: '#52B788',
  'Rumah Sakit': '#D62828',
  Apotek: '#1B1B1B',
}

onMounted(() => {
  if (!mapContainer.value) return

  map = L.map(mapContainer.value).setView([-8.65, 115.2166], 12)

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; OpenStreetMap contributors',
    maxZoom: 19,
  }).addTo(map)

  markerLayer = L.layerGroup().addTo(map)

  map.on('click', (e: L.LeafletMouseEvent) => {
    emit('map-clicked', { lat: e.latlng.lat, lng: e.latlng.lng })
    if (map) {
      L.popup()
        .setLatLng(e.latlng)
        .setContent('<p class="font-bold text-sm text-primary m-0">âœ“ Koordinat Diambil!</p>')
        .openOn(map)
    }
  })

  const legend = new L.Control({ position: 'bottomright' })
  legend.onAdd = function () {
    const div = L.DomUtil.create('div', 'bg-white p-3 rounded shadow text-sm leading-relaxed')
    const types = ['Puskesmas', 'Klinik', 'Rumah Sakit', 'Apotek']
    div.innerHTML = '<h4 class="font-bold mb-2 border-b pb-1 text-text">Keterangan</h4>'
    types.forEach((type) => {
      div.innerHTML += `
        <div class="flex items-center mb-1">
            <div class="w-4 h-4 rounded-full mr-2 border-2 border-white shadow-sm" style="background:${typeColors[type]}"></div>
            <span class="text-text/80">${type}</span>
        </div>
      `
    })
    return div
  }
  legend.addTo(map)

  renderMarkers()

  setTimeout(() => {
    map?.invalidateSize()
  }, 100)
  setTimeout(() => {
    map?.invalidateSize()
  }, 500)
  setTimeout(() => {
    map?.invalidateSize()
  }, 1000)

  const resizeObserver = new ResizeObserver(() => {
    if (map) {
      map.invalidateSize()
    }
  })
  resizeObserver.observe(mapContainer.value)

  onUnmounted(() => {
    resizeObserver.disconnect()
    if (map) {
      map.remove()
      map = null
    }
  })
})

const renderMarkers = () => {
  if (!markerLayer || !map) return
  markerLayer.clearLayers()

  facilities.value.forEach((facility) => {
    const color = typeColors[facility.type] || '#888'

    const customIcon = L.divIcon({
      className: 'custom-pin',
      html: `<div style="background-color:${color}; width: 18px; height: 18px; border-radius: 50%; border: 2.5px solid white; box-shadow: 0 0 5px rgba(0,0,0,0.6);"></div>`,
      iconSize: [20, 20],
      iconAnchor: [10, 10],
      popupAnchor: [0, -12],
    })

    const popupContent = `
      <div class="px-1 py-0 min-w-[200px]">
          <h3 class="font-bold text-base mb-1 text-text border-b pb-1">${facility.name}</h3>
          <div class="inline-block px-2 py-0.5 rounded text-[10px] font-bold text-white mb-2" style="background-color:${color}">${facility.type.toUpperCase()}</div>
          <div class="text-xs text-text/80 space-y-1.5 mt-1">
              <p class="flex items-start"><strong class="w-16 shrink-0">Area:</strong> ${facility.district}</p>
              <p class="flex items-start"><strong class="w-16 shrink-0">Alamat:</strong> ${facility.address}</p>
              <p class="flex items-start"><strong class="w-16 shrink-0">Telepon:</strong> ${facility.phone || '-'}</p>
              <p class="flex items-start"><strong class="w-16 shrink-0">Buka:</strong> ${facility.hours || '-'}</p>
          </div>
      </div>
    `

    L.marker([facility.lat, facility.lng], { icon: customIcon })
      .bindPopup(popupContent, { maxWidth: 300 })
      .addTo(markerLayer!)
  })
}

watch(
  facilities,
  () => {
    renderMarkers()
  },
  { deep: true },
)

watch(
  () => props.flyTo,
  (newVal) => {
    if (newVal && map) {
      map.flyTo([newVal.lat, newVal.lng], 14, { duration: 1.5 })
    }
  },
)
</script>

<template>
  <div
    ref="mapContainer"
    class="z-10 cursor-crosshair"
    style="width: 100%; height: 100%; min-height: 500px"
  ></div>
</template>
