<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useNotificationStore } from '@/stores/notifications'

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'
const authStore = useAuthStore()
const notify = useNotificationStore()

interface Category {
  id: number
  name: string
  icon: string
}

const categories = ref<Category[]>([])
const isLoading = ref(false)
const isSubmitting = ref(false)
const showForm = ref(false)
const isEditing = ref(false)
const deleteConfirmId = ref<number | null>(null)

const form = ref({ id: 0, name: '', icon: '' })

const iconOptions = [
  { label: 'Temple (Pura)', value: 'temple' },
  { label: 'Waterfall (Air Terjun)', value: 'waterfall' },
  { label: 'Beach (Pantai)', value: 'beach' },
  { label: 'Mountain (Gunung)', value: 'mountain' },
  { label: 'Hill (Bukit)', value: 'hill' },
  { label: 'Lake (Danau)', value: 'lake' },
  { label: 'Forest (Hutan Wisata)', value: 'forest' },
  { label: 'Village (Desa Wisata)', value: 'village' },
  { label: 'Museum', value: 'museum' },
  { label: 'Market (Pasar Seni)', value: 'market' },
]

const iconEmoji: Record<string, string> = {
  temple: '🛕', waterfall: '🌊', beach: '🏖️', mountain: '⛰️',
  hill: '🌄', lake: '🏞️', forest: '🌲', village: '🏘️',
  museum: '🏛️', market: '🎨',
}

const fetchCategories = async () => {
  isLoading.value = true
  try {
    const res = await fetch(`${API_URL}/categories`)
    const data = await res.json()
    if (data.success) categories.value = data.data
  } catch {
    notify.error('Gagal memuat kategori')
  } finally {
    isLoading.value = false
  }
}

const openCreate = () => {
  form.value = { id: 0, name: '', icon: iconOptions[0]?.value ?? 'temple' }
  isEditing.value = false
  showForm.value = true
}

const openEdit = (cat: Category) => {
  form.value = { id: cat.id, name: cat.name, icon: cat.icon }
  isEditing.value = true
  showForm.value = true
}

const submitForm = async () => {
  if (!form.value.name.trim() || !form.value.icon.trim()) {
    notify.warning('Nama dan icon wajib diisi')
    return
  }
  isSubmitting.value = true
  try {
    const method = isEditing.value ? 'PUT' : 'POST'
    const url = isEditing.value ? `${API_URL}/categories/${form.value.id}` : `${API_URL}/categories`
    const res = await fetch(url, {
      method,
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${authStore.token}`,
      },
      body: JSON.stringify({ name: form.value.name, icon: form.value.icon }),
    })
    const data = await res.json()
    if (data.success) {
      notify.success(isEditing.value ? 'Kategori diperbarui!' : 'Kategori ditambahkan!')
      showForm.value = false
      fetchCategories()
    } else {
      notify.error(data.message || 'Terjadi kesalahan')
    }
  } finally {
    isSubmitting.value = false
  }
}

const deleteCategory = async (id: number) => {
  try {
    const res = await fetch(`${API_URL}/categories/${id}`, {
      method: 'DELETE',
      headers: { Authorization: `Bearer ${authStore.token}` },
    })
    const data = await res.json()
    if (data.success) {
      notify.success('Kategori dihapus')
      deleteConfirmId.value = null
      fetchCategories()
    } else {
      notify.error(data.message || 'Gagal menghapus')
    }
  } catch {
    notify.error('Gagal menghapus kategori')
  }
}

onMounted(fetchCategories)
</script>

<template>
  <div class="p-8 min-h-screen bg-[#0f1117]">
    <!-- Header -->
    <div class="flex items-center justify-between mb-8">
      <div>
        <h1 class="text-2xl font-bold text-white">Manajemen Kategori</h1>
        <p class="text-white/40 text-sm mt-1">Kelola kategori objek yang tampil di peta interaktif</p>
      </div>
      <button @click="openCreate" class="flex items-center gap-2 px-4 py-2.5 bg-amber-500 hover:bg-amber-400 text-white font-semibold rounded-xl text-sm transition-all active:scale-95">
        <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
          <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
        </svg>
        Tambah Kategori
      </button>
    </div>

    <!-- Loading -->
    <div v-if="isLoading" class="flex items-center justify-center py-20 text-white/40">
      <svg class="animate-spin w-6 h-6 mr-3" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z" />
      </svg>
      Memuat kategori...
    </div>

    <!-- Category Grid -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
      <div v-for="cat in categories" :key="cat.id" class="group bg-[#161b27] border border-white/5 rounded-2xl p-5 hover:border-amber-500/30 transition-all">
        <div class="flex items-start justify-between mb-4">
          <div class="w-12 h-12 rounded-xl bg-amber-500/10 flex items-center justify-center text-2xl">
            {{ iconEmoji[cat.icon] || '📍' }}
          </div>
          <div class="flex gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
            <button @click="openEdit(cat)" class="p-1.5 rounded-lg text-white/40 hover:text-amber-400 hover:bg-amber-500/10 transition-all" title="Edit">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round"
                  d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897l8.932-8.931Zm0 0L19.5 7.125" />
              </svg>
            </button>
            <button @click="deleteConfirmId = cat.id" class="p-1.5 rounded-lg text-white/40 hover:text-red-400 hover:bg-red-500/10 transition-all" title="Hapus">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round"
                  d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0" />
              </svg>
            </button>
          </div>
        </div>
        <h3 class="font-semibold text-white text-sm">{{ cat.name }}</h3>
        <p class="text-white/30 text-xs mt-0.5 font-mono">icon: {{ cat.icon }}</p>
      </div>
    </div>

    <!-- Form Modal -->
    <Transition name="modal-fade">
      <div v-if="showForm" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/70 backdrop-blur-sm" @click="showForm = false"></div>
        <div class="relative bg-[#161b27] border border-white/10 rounded-2xl shadow-2xl w-full max-w-md p-6">
          <h2 class="text-lg font-bold text-white mb-5">
            {{ isEditing ? 'Edit Kategori' : 'Tambah Kategori Baru' }}
          </h2>

          <div class="space-y-4">
            <div>
              <label class="block text-xs font-semibold text-white/50 uppercase tracking-wider mb-1.5">Nama Kategori</label>
              <input v-model="form.name" type="text" placeholder="Contoh: Pura"
                class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-2.5 text-sm text-white placeholder-white/20 focus:outline-none focus:border-amber-500/50 transition-colors"
                @keyup.enter="submitForm" />
            </div>
            <div>
              <label class="block text-xs font-semibold text-white/50 uppercase tracking-wider mb-1.5">Icon</label>
              <select v-model="form.icon" class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-2.5 text-sm text-white focus:outline-none focus:border-amber-500/50 transition-colors">
                <option v-for="opt in iconOptions" :key="opt.value" :value="opt.value" class="bg-[#161b27]">
                  {{ iconEmoji[opt.value] }} {{ opt.label }}
                </option>
              </select>
            </div>
          </div>

          <div class="flex gap-3 mt-6">
            <button @click="showForm = false" class="flex-1 py-2.5 rounded-xl border border-white/10 text-white/60 hover:text-white hover:bg-white/5 text-sm font-medium transition-all">
              Batal
            </button>
            <button @click="submitForm" :disabled="isSubmitting"
              class="flex-1 py-2.5 rounded-xl bg-amber-500 hover:bg-amber-400 text-white text-sm font-semibold transition-all disabled:opacity-50 flex items-center justify-center gap-2">
              <svg v-if="isSubmitting" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z" />
              </svg>
              {{ isEditing ? 'Simpan Perubahan' : 'Tambahkan' }}
            </button>
          </div>
        </div>
      </div>
    </Transition>

    <!-- Delete Confirmation Modal -->
    <Transition name="modal-fade">
      <div v-if="deleteConfirmId !== null" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/70 backdrop-blur-sm" @click="deleteConfirmId = null"></div>
        <div class="relative bg-[#161b27] border border-red-500/20 rounded-2xl shadow-2xl w-full max-w-sm p-6 text-center">
          <div class="w-12 h-12 rounded-full bg-red-500/10 flex items-center justify-center mx-auto mb-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-red-400" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round"
                d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z" />
            </svg>
          </div>
          <h3 class="text-white font-bold mb-2">Hapus Kategori?</h3>
          <p class="text-white/40 text-sm mb-6">Tindakan ini tidak dapat dibatalkan. Pastikan tidak ada objek peta yang menggunakan kategori ini.</p>
          <div class="flex gap-3">
            <button @click="deleteConfirmId = null" class="flex-1 py-2.5 rounded-xl border border-white/10 text-white/60 text-sm font-medium hover:bg-white/5 transition-all">Batal</button>
            <button @click="deleteCategory(deleteConfirmId!)" class="flex-1 py-2.5 rounded-xl bg-red-500 hover:bg-red-400 text-white text-sm font-semibold transition-all">Hapus</button>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: all 0.2s ease;
}

.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
  transform: scale(0.96);
}
</style>
