<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useMapPointsStore, type GeoPoint } from '@/stores/mapPoints'
import { useNotificationStore } from '@/stores/notifications'

const store = useMapPointsStore()
const notificationStore = useNotificationStore()
const pendingPoints = ref<GeoPoint[]>([])
const isLoading = ref(false)

const selectedPoint = ref<GeoPoint | null>(null)
const blogData = ref<{ title: string; content: string; cover_photo: string } | null>(null)
const isPreviewOpen = ref(false)
const showRejectModal = ref(false)
const rejectionNote = ref('')
const isProcessing = ref(false)

const fetchPending = async () => {
  isLoading.value = true
  try {
    const res = await fetch(`${import.meta.env.VITE_API_URL || 'http://localhost:8080/api'}/points/pending`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      },
    })
    const json = await res.json()
    if (json.success) {
      pendingPoints.value = json.data
    }
  } catch {
    notificationStore.error('Gagal mengambil data pengajuan')
  } finally {
    isLoading.value = false
  }
}

onMounted(fetchPending)

const openPreview = async (point: GeoPoint) => {
  selectedPoint.value = point
  isPreviewOpen.value = true
  if (point.id) {
    const blog = await store.getBlog(point.id)
    if (blog) {
      blogData.value = { title: blog.title, content: blog.content, cover_photo: blog.cover_photo || '' }
    } else {
      blogData.value = null
    }
  }
}


const verifyPoint = async (status: 'approved' | 'rejected') => {
  if (!selectedPoint.value) return

  isProcessing.value = true
  try {
    const res = await fetch(`${import.meta.env.VITE_API_URL || 'http://localhost:8080/api'}/points/${selectedPoint.value.id}/verify`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      },
      body: JSON.stringify({
        status,
        rejection_note: status === 'rejected' ? rejectionNote.value : ''
      }),
    })
    const json = await res.json()
    if (json.success) {
      notificationStore.success(`Berhasil memverifikasi: ${status}`)
      isPreviewOpen.value = false
      showRejectModal.value = false
      rejectionNote.value = ''
      fetchPending()
    } else {
      notificationStore.error(json.message)
    }
  } catch {
    notificationStore.error('Gagal memproses verifikasi')
  } finally {
    isProcessing.value = false
  }
}

const getCategoryName = (id?: number) => {
  return store.objectTypes.find(t => t.id === id)?.name || 'Kategori'
}
</script>

<template>
  <div class="p-8 space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-white">Verifikasi Konten</h2>
        <p class="text-sm text-white/40 mt-1">Tinjau pengajuan dari kontributor sebelum ditayangkan di peta publik.</p>
      </div>
      <div class="bg-amber-500/10 px-4 py-2 rounded-xl border border-amber-500/20">
        <span class="text-amber-500 font-bold text-sm">{{ pendingPoints.length }} Menunggu Verifikasi</span>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="isLoading" class="flex flex-col items-center justify-center py-20 gap-4">
      <div class="w-10 h-10 border-4 border-amber-500/20 border-t-amber-500 rounded-full animate-spin"></div>
      <p class="text-white/40 text-sm font-medium">Memuat pengajuan...</p>
    </div>

    <!-- Empty State -->
    <div v-else-if="pendingPoints.length === 0" class="bg-[#161b27] border border-white/5 rounded-2xl p-12 text-center">
      <div class="w-16 h-16 bg-white/5 rounded-full flex items-center justify-center mx-auto mb-4 text-white/20">
        <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
      </div>
      <h3 class="text-white font-bold">Semua Bersih!</h3>
      <p class="text-white/40 text-sm mt-1">Tidak ada pengajuan baru yang perlu diverifikasi.</p>
    </div>

    <!-- Grid Pengajuan -->
    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="point in pendingPoints" :key="point.id"
        class="group bg-[#161b27] border border-white/5 rounded-2xl overflow-hidden hover:border-amber-500/30 transition-all flex flex-col shadow-xl shadow-black/20">
        <!-- Image Preview -->
        <div class="h-40 bg-slate-800 relative overflow-hidden">
          <img v-if="point.description" :src="point.description" class="w-full h-full object-cover">
          <div v-else class="w-full h-full flex items-center justify-center text-white/10 italic text-xs">Tanpa Foto</div>
          <div class="absolute top-3 left-3 bg-black/60 backdrop-blur px-2 py-1 rounded-md border border-white/10">
            <span class="text-[10px] font-bold text-white uppercase">{{ getCategoryName(point.type_id) }}</span>
          </div>
        </div>


        <div class="p-5 flex-1 flex flex-col gap-4">
          <div>
            <h4 class="text-white font-bold text-lg group-hover:text-amber-400 transition-colors">{{ point.name }}</h4>
            <p class="text-white/40 text-xs mt-1 line-clamp-2">{{ point.address }}</p>
          </div>

          <div class="flex items-center gap-3 pt-4 border-t border-white/5">
            <div class="w-8 h-8 rounded-full bg-amber-500/20 flex items-center justify-center text-amber-500">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
            </div>
            <div>
              <p class="text-white/60 text-xs font-bold">{{ point.owner_name }}</p>
              <p class="text-[10px] text-white/20 uppercase tracking-widest">{{ new Date(point.created_at || '').toLocaleDateString('id-ID') }}</p>
            </div>

          </div>

          <button @click="openPreview(point)" class="w-full py-2.5 bg-white/5 hover:bg-white/10 text-white text-xs font-bold rounded-xl transition-all border border-white/10">
            Tinjau Detail
          </button>
        </div>
      </div>
    </div>

    <!-- Modal Preview -->
    <div v-if="isPreviewOpen && selectedPoint" class="fixed inset-0 z-[2000] flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-black/80 backdrop-blur-sm" @click="isPreviewOpen = false"></div>

      <div class="relative bg-[#0f141e] border border-white/10 rounded-3xl w-full max-w-4xl max-h-[90vh] overflow-hidden flex flex-col shadow-2xl">
        <!-- Header -->
        <div class="px-8 py-4 border-b border-white/5 flex items-center justify-between shrink-0">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-xl bg-amber-500/10 flex items-center justify-center text-amber-500 font-bold">
              ?
            </div>
            <div>
              <h3 class="text-white font-bold">Peninjauan Pengajuan</h3>
              <p class="text-[10px] text-white/40 uppercase tracking-widest">ID: {{ selectedPoint.id }} • Dari: {{ selectedPoint.owner_name }}</p>
            </div>
          </div>
          <button @click="isPreviewOpen = false" class="text-white/20 hover:text-white transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Scrollable Content -->
        <div class="flex-1 overflow-y-auto p-8 space-y-8 custom-scrollbar">
          <!-- Main Info -->
          <div class="grid grid-cols-3 gap-8">
            <div class="col-span-1 space-y-6">
              <div class="aspect-video rounded-2xl bg-white/5 overflow-hidden border border-white/10">
                <img v-if="selectedPoint.description" :src="selectedPoint.description" class="w-full h-full object-cover">
              </div>
              <div class="space-y-4">
                <div v-for="info in [
                  { label: 'Nama', value: selectedPoint.name },
                  { label: 'Kategori', value: getCategoryName(selectedPoint.type_id) },
                  { label: 'Lokasi', value: `${selectedPoint.latitude}, ${selectedPoint.longitude}` },
                  { label: 'Alamat', value: selectedPoint.address }
                ]" :key="info.label">

                  <p class="text-[10px] font-black text-white/30 uppercase tracking-widest">{{ info.label }}</p>
                  <p class="text-sm text-white/80 font-bold mt-1">{{ info.value }}</p>
                </div>
              </div>
            </div>

            <div class="col-span-2 bg-white/5 border border-white/10 rounded-2xl p-6">
              <p class="text-[10px] font-black text-amber-500 uppercase tracking-widest mb-4">Konten Blog / Artikel</p>
              <div v-if="blogData" class="space-y-4">
                <h1 class="text-2xl font-black text-white">{{ blogData.title }}</h1>
                <div class="prose prose-invert max-w-none text-white/60 text-sm leading-relaxed" v-html="blogData.content"></div>
              </div>
              <div v-else class="h-64 flex flex-col items-center justify-center text-white/10">
                <p class="italic text-sm font-medium">Kontributor belum mengisi konten blog.</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer Actions -->
        <div class="px-8 py-6 border-t border-white/5 bg-white/5 shrink-0 flex items-center justify-end gap-4">
          <button @click="showRejectModal = true" class="px-6 py-2.5 border border-red-500/30 text-red-400 hover:bg-red-500/10 text-xs font-black uppercase tracking-widest rounded-xl transition-all">
            Tolak Pengajuan
          </button>
          <button @click="verifyPoint('approved')" :disabled="isProcessing"
            class="px-8 py-2.5 bg-emerald-600 hover:bg-emerald-500 text-white text-xs font-black uppercase tracking-widest rounded-xl transition-all shadow-lg shadow-emerald-900/40 disabled:opacity-50">
            {{ isProcessing ? 'Memproses...' : 'Setujui & Publikasikan' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Reject Modal -->
    <div v-if="showRejectModal" class="fixed inset-0 z-[2100] flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-black/90 backdrop-blur-md" @click="showRejectModal = false"></div>
      <div class="relative bg-[#161b27] border border-white/10 rounded-2xl w-full max-w-md p-6 shadow-2xl">
        <h3 class="text-white font-bold text-lg mb-2">Alasan Penolakan</h3>
        <p class="text-white/40 text-xs mb-4">Berikan catatan mengapa pengajuan ini ditolak agar kontributor dapat memperbaikinya.</p>
        <textarea v-model="rejectionNote" rows="4" class="w-full bg-black/40 border border-white/10 rounded-xl p-4 text-white text-sm focus:border-red-500/50 outline-none transition-all"
          placeholder="Tulis alasan di sini..."></textarea>
        <div class="flex gap-3 mt-6">
          <button @click="showRejectModal = false" class="flex-1 py-2.5 text-white/40 hover:text-white text-xs font-bold transition-all">Batal</button>
          <button @click="verifyPoint('rejected')" :disabled="!rejectionNote || isProcessing"
            class="flex-1 py-2.5 bg-red-600 hover:bg-red-500 text-white text-xs font-bold rounded-xl transition-all disabled:opacity-50">
            {{ isProcessing ? '...' : 'Konfirmasi Tolak' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.05);
  border-radius: 10px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.1);
}
</style>
