<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useNotificationStore } from '@/stores/notifications'
import type { GeoPoint } from '@/stores/mapPoints'

interface BlogDetail {
  blog: {
    id: number
    title: string
    content: string
    cover_photo: string | null
    updated_at: string
  }
  point: GeoPoint
}

const route = useRoute()
const router = useRouter()
const notificationStore = useNotificationStore()
const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

const detail = ref<BlogDetail | null>(null)
const isLoading = ref(true)

const fetchBlogDetail = async () => {
  try {
    const res = await fetch(`${API_URL}/public/blogs/${route.params.id}`)
    const json = await res.json()
    if (json.success) {
      detail.value = json.data
    } else {
      notificationStore.error('Gagal memuat detail ulasan')
      router.push('/')
    }
  } catch {
    notificationStore.error('Terjadi kesalahan jaringan')
    router.push('/')
  } finally {

    isLoading.value = false
  }
}

onMounted(fetchBlogDetail)
</script>

<template>
  <div class="min-h-screen bg-slate-50 pb-20">
    <!-- Navbar / Back Button -->
    <nav class="fixed top-0 left-0 right-0 h-16 bg-white/80 backdrop-blur-md border-b border-slate-100 z-50 flex items-center px-6 justify-between">
      <button @click="router.back()" class="flex items-center gap-2 text-slate-600 hover:text-primary font-bold transition-colors">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5" stroke="currentColor" class="w-5 h-5">
          <path stroke-linecap="round" stroke-linejoin="round" d="M10.5 19.5L3 12m0 0l7.5-7.5M3 12h18" />
        </svg>
        Kembali ke Peta
      </button>
      <div class="hidden md:block">
        <h1 v-if="detail" class="text-sm font-black text-slate-400 uppercase tracking-widest">{{ detail.point.name }}</h1>
      </div>
      <div class="w-24"></div>
    </nav>

    <!-- Content -->
    <div v-if="isLoading" class="flex flex-col items-center justify-center h-[80vh]">
      <div class="w-12 h-12 border-4 border-primary border-t-transparent rounded-full animate-spin"></div>
      <p class="mt-4 text-slate-400 font-bold animate-pulse">Memuat Ulasan...</p>
    </div>

    <div v-else-if="detail" class="pt-24 max-w-4xl mx-auto px-6">
      <!-- Header Section -->
      <header class="mb-12">
        <div class="flex items-center gap-3 mb-6">
          <span class="px-3 py-1 bg-amber-100 text-amber-700 text-[10px] font-black uppercase tracking-widest rounded-full border border-amber-200">
            Destinasi Pilihan
          </span>
          <span class="text-slate-400 text-[10px] font-bold uppercase tracking-wider">
            Ditinjau pada {{ new Date(detail.blog.updated_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' }) }}
          </span>
        </div>

        <h1 class="text-4xl md:text-5xl font-black text-slate-900 leading-tight mb-8">
          {{ detail.blog.title }}
        </h1>

        <div class="flex items-center gap-4 p-4 bg-white rounded-2xl border border-slate-100 shadow-sm">
          <div class="w-12 h-12 bg-indigo-600 rounded-xl flex items-center justify-center text-white shrink-0 shadow-lg shadow-indigo-100">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-6 h-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" />
              <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" />
            </svg>
          </div>
          <div>
            <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-none mb-1">Lokasi Objek</p>
            <p class="text-sm font-extrabold text-slate-800">{{ detail.point.name }}</p>
            <p class="text-xs text-slate-500 font-medium line-clamp-1">{{ detail.point.address }}</p>
          </div>
        </div>
      </header>

      <!-- Cover Photo -->
      <div v-if="detail.blog.cover_photo" class="relative group rounded-3xl overflow-hidden shadow-2xl mb-16 aspect-video">
        <img :src="detail.blog.cover_photo" class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105" alt="Cover Photo">
        <div class="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
      </div>

      <!-- Article Body -->
      <article
        class="prose prose-slate lg:prose-xl max-w-none prose-headings:font-black prose-headings:text-slate-900 prose-p:text-slate-600 prose-p:leading-relaxed prose-img:rounded-3xl prose-img:shadow-lg">
        <div v-html="detail.blog.content"></div>
      </article>

      <!-- Footer Info -->
      <footer class="mt-20 pt-12 border-t border-slate-100 flex flex-col md:flex-row items-center justify-between gap-8">
        <div class="flex items-center gap-4">
          <div class="w-14 h-14 bg-slate-200 rounded-full overflow-hidden border-2 border-white shadow-md">
            <div class="w-full h-full flex items-center justify-center font-black text-slate-400 text-xl">
              {{ detail.point.owner_name?.charAt(0) || 'S' }}
            </div>
          </div>
          <div>
            <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-none mb-1">Ditulis oleh</p>
            <p class="text-lg font-black text-slate-900">{{ detail.point.owner_name || 'Sistem' }}</p>
            <p class="text-xs text-slate-500 font-medium italic">Kontributor Warisan Budaya Bali</p>
          </div>
        </div>

        <button @click="router.push('/')" class="px-8 py-4 bg-slate-900 text-white font-bold rounded-2xl hover:bg-primary transition-all shadow-xl shadow-slate-200 hover:shadow-primary/20">
          Jelajahi Peta Bali
        </button>
      </footer>
    </div>
  </div>
</template>

<style>
/* Custom prose styles for Quill content */
.prose img {
  @apply mx-auto my-12;
}

.prose h1,
.prose h2,
.prose h3 {
  @apply mt-16 mb-8;
}

.prose p {
  @apply mb-6;
}

.prose strong {
  @apply text-slate-900;
}

.prose blockquote {
  @apply border-l-4 border-amber-400 bg-amber-50/50 p-8 italic rounded-r-2xl;
}
</style>
