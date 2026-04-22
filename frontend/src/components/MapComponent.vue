<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch, nextTick } from 'vue'
import L from 'leaflet'
import 'leaflet/dist/leaflet.css'
import { useMapPointsStore } from '@/stores/mapPoints'
import { useAuthStore } from '@/stores/auth'
import { storeToRefs } from 'pinia'

const emit = defineEmits<{
  (e: 'map-clicked', data: { lat: number; lng: number; address?: string }): void
}>()

const store = useMapPointsStore()
const authStore = useAuthStore()
const { filteredPoints, objectTypes, isEditMode, isSidebarExpanded } = storeToRefs(store)

const mapContainer = ref<HTMLElement | null>(null)
let map: L.Map | null = null
let markerLayer: L.LayerGroup | null = null

const reverseGeocode = async (lat: number, lng: number) => {
  try {
    const res = await fetch(`https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=${lat}&lon=${lng}&accept-language=id`)
    const data = await res.json()
    return data.display_name || 'Alamat tidak ditemukan'
  } catch (error) {
    console.error('Reverse geocoding error:', error)
    return 'Gagal mendeteksi alamat'
  }
}

onMounted(async () => {
  if (!mapContainer.value) return

  await store.fetchPoints()

  map = L.map(mapContainer.value, { zoomControl: false }).setView([-8.65, 115.2166], 12)
  L.control.zoom({ position: 'bottomleft' }).addTo(map)

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; OpenStreetMap contributors',
    maxZoom: 19,
  }).addTo(map)

  markerLayer = L.layerGroup().addTo(map)

  map.on('click', async (e: L.LeafletMouseEvent) => {
    console.log('[MapComponent] Map clicked!', e.latlng, 'isEditMode:', store.isEditMode)
    if (!store.isEditMode) return
    console.log('[MapComponent] Edit mode active, emitting map-clicked...')

    emit('map-clicked', { lat: e.latlng.lat, lng: e.latlng.lng })

    if (map) {
      L.popup()
        .setLatLng(e.latlng)
        .setContent('<p class="font-bold text-sm text-primary m-0">⏳ Mendeteksi Alamat...</p>')
        .openOn(map)
    }

    const address = await reverseGeocode(e.latlng.lat, e.latlng.lng)

    if (store.activePoint && store.activePoint.latitude === e.latlng.lat && store.activePoint.longitude === e.latlng.lng) {
      store.activePoint = { ...store.activePoint, address }
    }

    if (map) {
      L.popup()
        .setLatLng(e.latlng)
        .setContent(`<div class="font-bold text-sm text-primary max-w-[200px] leading-tight"><p class="m-0 mb-1">📍 ${address}</p></div>`)
        .openOn(map)
    }
  })

  const legend = new L.Control({ position: 'bottomright' })
  legend.onAdd = function () {
    const container = L.DomUtil.create('div', 'bg-white rounded-2xl shadow-xl border border-gray-100 overflow-hidden transition-all duration-300 min-w-[160px] font-sans')
    container.id = 'map-legend'

    const header = L.DomUtil.create('div', 'flex items-center justify-between px-3 py-2.5 bg-gray-50/50 border-b border-gray-100 cursor-pointer group', container)
    header.innerHTML = `
      <span class="font-extrabold text-gray-800 tracking-tight uppercase text-[10px]">Legenda Tipe</span>
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="3" stroke="currentColor" class="w-3 h-3 text-gray-400 group-hover:text-primary transition-transform duration-300 transform" id="legend-chevron">
        <path stroke-linecap="round" stroke-linejoin="round" d="m19.5 8.25-7.5 7.5-7.5-7.5" />
      </svg>
    `

    const body = L.DomUtil.create('div', 'p-3 max-h-[300px] overflow-y-auto transition-all duration-300', container)
    body.id = 'legend-body'

    const colors = [
      '#2D6A4F', '#0077B6', '#D62828', '#F4A261', '#9D4EDD',
      '#FFB703', '#E63946', '#219EBC', '#023047', '#FB8500',
      '#8ECAE6', '#606C38'
    ]

    objectTypes.value.forEach((type, index) => {
      const color = colors[index % colors.length]
      body.innerHTML += `
        <div class="flex items-center mb-2 last:mb-0">
            <div class="w-6 h-6 rounded-lg mr-2 flex items-center justify-center text-white p-1.5 shadow-sm" style="background:${color}">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round" class="w-full h-full">
                <path d="${type.icon}" />
              </svg>
            </div>
            <span class="text-gray-600 font-bold text-[11px] uppercase tracking-wide">${type.name}</span>
        </div>
      `
    })

    let isExpanded = true
    L.DomEvent.on(header, 'click', (e) => {
      L.DomEvent.stopPropagation(e)
      isExpanded = !isExpanded
      const chevron = document.getElementById('legend-chevron')
      if (isExpanded) {
        body.classList.remove('hidden')
        chevron?.classList.remove('-rotate-180')
      } else {
        body.classList.add('hidden')
        chevron?.classList.add('-rotate-180')
      }
    })

    L.DomEvent.disableScrollPropagation(container)
    L.DomEvent.disableClickPropagation(container)

    return container
  }
  legend.addTo(map)

  renderMarkers()

  setTimeout(() => map?.invalidateSize(), 100)
  setTimeout(() => map?.invalidateSize(), 500)
})

onUnmounted(() => {
  if (map) {
    map.remove()
    map = null
  }
})

const renderMarkers = () => {
  if (!map || !markerLayer) return
  markerLayer.clearLayers()

  const colors = [
    '#2D6A4F', '#0077B6', '#D62828', '#F4A261', '#9D4EDD',
    '#FFB703', '#E63946', '#219EBC', '#023047', '#FB8500',
    '#8ECAE6', '#606C38'
  ]

  filteredPoints.value.forEach((point) => {
    const typeObj = objectTypes.value.find(t => t.id === point.type_id)
    const typeName = typeObj ? typeObj.name : 'Unknown'
    const typeIconPath = typeObj ? typeObj.icon : 'M12 2v20M2 12h20'
    const color = colors[(point.type_id - 1) % colors.length]

    const customIcon = L.divIcon({
      className: 'custom-div-icon',
      html: `
        <div class="relative group">
          <div class="w-8 h-8 rounded-xl flex items-center justify-center text-white shadow-lg transition-all duration-300 group-hover:scale-110 group-hover:rotate-3 active:scale-95" style="background:${color}; border: 2px solid white;">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round" class="w-4 h-4">
              <path d="${typeIconPath}" />
            </svg>
          </div>
          <div class="absolute -top-1 -right-1 w-3 h-3 bg-white rounded-full border-2 border-[${color}] shadow-sm animate-pulse"></div>
        </div>
      `,
      iconSize: [32, 32],
      iconAnchor: [16, 16],
      popupAnchor: [0, -24],
    })

    const marker = L.marker([point.latitude, point.longitude], { icon: customIcon }).addTo(markerLayer!)

    marker.bindPopup(`
      <div class="p-3 min-w-[200px] font-sans">
        <h4 class="font-black text-gray-900 text-sm mb-1 tracking-tight">${point.name}</h4>
        <div class="space-y-1.5 text-[11px]">
          <p class="flex items-start gap-2"><span class="w-20 shrink-0 font-semibold text-gray-400">Kategori:</span> <span class="bg-gray-100 px-1.5 rounded text-gray-700 font-bold">${typeName}</span></p>
          <p class="flex items-start gap-2"><span class="w-20 shrink-0 font-semibold text-gray-400">Alamat:</span> <span class="text-gray-600 leading-normal">${point.address || 'Tanpa Alamat'}</span></p>
          ${point.tahun_berdiri ? `<p class="flex items-start gap-2"><span class="w-20 shrink-0 font-semibold text-gray-400">Th. Berdiri:</span> <span class="text-gray-800">${point.tahun_berdiri}</span></p>` : ''}
          <p class="flex items-start gap-2"><span class="w-20 shrink-0 font-semibold text-gray-400">Kepemilikan:</span> <span class="px-1.5 bg-gray-100 rounded text-gray-700 font-bold">${point.status_kepemilikan || 'Tidak Diketahui'}</span></p>
          <div class="mt-2 pt-2 border-t border-dotted border-gray-200 flex items-center gap-1.5 opacity-70 italic text-[10px]">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-3 h-3">
              <path d="M10 8a3 3 0 100-6 3 3 0 000 6zM3.465 14.493a1.23 1.23 0 00.41 1.412A9.957 9.957 0 0010 18c2.31 0 4.438-.784 6.131-2.1.43-.333.604-.903.408-1.41a7.002 7.002 0 00-13.074.003z" />
            </svg>
            Kontributor: ${point.owner_name || 'Sistem'}
          </div>
          ${authStore.isAuthenticated() && store.isEditMode && point.owner_id === authStore.user?.id ? `
          <div class="mt-2 flex gap-2 pt-2 border-t border-gray-100">
            <button class="edit-point-btn flex-1 bg-primary/10 hover:bg-primary/20 text-primary font-bold py-1.5 rounded text-[10px] uppercase tracking-wide transition-colors" data-id="${point.id}">Edit</button>
            <button class="delete-point-btn flex-1 bg-red-50 hover:bg-red-100 text-red-600 font-bold py-1.5 rounded text-[10px] uppercase tracking-wide transition-colors" data-id="${point.id}">Hapus</button>
          </div>
          ` : ''}
        </div>
      </div>
    `, {
      maxWidth: 300,
      className: 'premium-popup'
    })

    marker.on('popupopen', () => {
      const editBtn = document.querySelector('.edit-point-btn')
      const deleteBtn = document.querySelector('.delete-point-btn')

      if (editBtn) {
        editBtn.addEventListener('click', () => {
          store.openModal(point)
        })
      }

      if (deleteBtn) {
        deleteBtn.addEventListener('click', () => {
          store.requestConfirm(
            'Hapus Bangunan?',
            `Apakah Anda yakin ingin menghapus data bangunan "${point.name}" secara permanen?`,
            async () => {
              if (point.id) {
                await store.deletePoint(point.id)
              }
            }
          )
        })
      }
    })
  })
}

watch([filteredPoints, isEditMode, () => authStore.user], () => {
  renderMarkers()
})

watch(isEditMode, (newMode) => {
  if (mapContainer.value) {
    if (newMode) {
      mapContainer.value.style.cursor = 'crosshair'
    } else {
      mapContainer.value.style.cursor = 'grab'
    }
  }
})

watch(isSidebarExpanded, () => {
  nextTick(() => {
    setTimeout(() => map?.invalidateSize(), 100)
    setTimeout(() => map?.invalidateSize(), 550)
  })
})

watch(() => store.flyToCoords, (newCoords) => {
  if (newCoords && map) {
    map.flyTo([newCoords.lat, newCoords.lng], 16)

    markerLayer?.eachLayer((layer: L.Layer) => {
      if (layer instanceof L.Marker) {
        const latLng = layer.getLatLng()
        if (latLng.lat === newCoords.lat && latLng.lng === newCoords.lng) {
          layer.openPopup()
        }
      }
    })
  }
})
</script>

<template>
  <div ref="mapContainer" class="w-full h-full bg-slate-100 z-0 flex-1"></div>
</template>

<style>
.leaflet-container {
  font-family: inherit;
}

.premium-popup .leaflet-popup-content-wrapper {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(8px);
  border: 1px solid rgba(0, 0, 0, 0.05);
  border-radius: 12px;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1);
  padding: 0;
}

.premium-popup .leaflet-popup-content {
  margin: 0;
}

.premium-popup .leaflet-popup-tip {
  background: white;
}
</style>
