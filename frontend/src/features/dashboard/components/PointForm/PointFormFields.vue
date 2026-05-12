<script setup lang="ts">
import { useMapPointsStore } from '@/stores/mapPoints'
import type { GeoPoint } from '@/types/map'
import { QuillEditor } from '@vueup/vue-quill'
import BlotFormatter from 'quill-blot-formatter'
import '@vueup/vue-quill/dist/vue-quill.snow.css'
import { ref } from 'vue'

const props = defineProps<{
  modelValue: Partial<GeoPoint>
  blogContent: { title: string; content: string; cover_photo: string }
  activeTab: 'data' | 'blog'
}>()

const emit = defineEmits(['update:modelValue', 'update:blogContent'])

const store = useMapPointsStore()
const isUploading = ref(false)
const isBlogImageUploading = ref(false)
const quillRef = ref()
const cursorBounds = ref<{ top: number; left: number } | null>(null)

const modules = {
  name: 'blotFormatter',
  module: BlotFormatter,
}

const updateField = <K extends keyof GeoPoint>(field: K, value: GeoPoint[K]) => {
  emit('update:modelValue', { ...props.modelValue, [field]: value })
}

const handleFileUpload = async (e: Event) => {
  const target = e.target as HTMLInputElement
  if (!target.files?.length) return

  isUploading.value = true
  const formData = new FormData()
  formData.append('image', target.files[0]!)
  formData.append('folder', 'points')

  try {
    const res = await fetch(`${import.meta.env.VITE_API_URL || 'http://localhost:8080/api'}/upload`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${localStorage.getItem('auth_token')}`,
      },
      body: formData,
    })
    const json = await res.json()
    if (json.success) {
      updateField('cover_image', json.data.url)
    }
  } finally {
    isUploading.value = false
  }
}

const uploadImageToCloudinary = (file: File): Promise<string | null> => {
  return new Promise((resolve) => {
    const formData = new FormData()
    formData.append('image', file)
    formData.append('folder', 'blogs')

    isBlogImageUploading.value = true

    fetch(`${import.meta.env.VITE_API_URL || 'http://localhost:8080/api'}/upload`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${localStorage.getItem('auth_token')}`,
      },
      body: formData,
    })
      .then(res => res.json())
      .then(json => {
        if (json.success) resolve(json.data.url)
        else resolve(null)
      })
      .catch(() => resolve(null))
      .finally(() => {
        isBlogImageUploading.value = false
        cursorBounds.value = null
      })
  })
}

const updateCursorBounds = () => {
  const quill = quillRef.value?.getQuill()
  if (!quill) return
  const range = quill.getSelection(true)
  if (range) {
    cursorBounds.value = quill.getBounds(range.index)
  }
}

const imageHandler = () => {
  const input = document.createElement('input')
  input.type = 'file'
  input.accept = 'image/*'
  input.click()

  input.onchange = async () => {
    const file = input.files?.[0]
    if (!file) return

    updateCursorBounds()
    const url = await uploadImageToCloudinary(file)
    if (!url) return

    const quill = quillRef.value?.getQuill()
    if (!quill) return

    const range = quill.getSelection(true)
    quill.insertEmbed(range.index, 'image', url)
    quill.setSelection(range.index + 1)
  }
}

const onEditorReady = () => {
  const quill = quillRef.value?.getQuill()
  if (!quill) return

  quill.getModule('toolbar').addHandler('image', imageHandler)

  quill.root.addEventListener('paste', async (e: ClipboardEvent) => {
    const items = e.clipboardData?.items
    if (!items) return

    for (const item of Array.from(items)) {
      if (item.type.indexOf('image') !== -1) {
        e.preventDefault()
        const file = item.getAsFile()
        if (file) {
          updateCursorBounds()
          const url = await uploadImageToCloudinary(file)
          if (url) {
            const range = quill.getSelection(true)
            quill.insertEmbed(range.index, 'image', url)
            quill.setSelection(range.index + 1)
          }
        }
      }
    }
  }, false)

  quill.root.addEventListener('drop', async (e: DragEvent) => {
    const files = e.dataTransfer?.files
    if (!files?.length) return

    for (const file of Array.from(files)) {
      if (file.type.indexOf('image') !== -1) {
        e.preventDefault()
        updateCursorBounds()
        const url = await uploadImageToCloudinary(file)
        if (url) {
          const range = quill.getSelection(true)
          quill.insertEmbed(range.index, 'image', url)
          quill.setSelection(range.index + 1)
        }
      }
    }
  }, false)
}
</script>

<template>
  <div class="p-6 overflow-y-auto w-full space-y-5">
    <template v-if="activeTab === 'data'">

      <div class="space-y-1.5">
        <label class="text-[11px] font-black text-slate-500 uppercase tracking-wider ml-1">Nama Objek / Destinasi</label>
        <input :value="modelValue.name" @input="e => updateField('name', (e.target as HTMLInputElement).value)" type="text"
          class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200 rounded-xl focus:bg-white focus:border-primary transition-all text-sm font-bold text-slate-800 outline-none"
          placeholder="Nama objek wisata...">
      </div>

      <div class="space-y-1.5">
        <label class="text-[11px] font-black text-slate-500 uppercase tracking-wider ml-1">Kategori Objek</label>
        <select :value="modelValue.category_id" @change="e => updateField('category_id', parseInt((e.target as HTMLSelectElement).value))"
          class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200 rounded-xl focus:bg-white focus:border-primary transition-all text-sm font-bold text-slate-700 appearance-none outline-none">
          <option v-for="type in store.objectTypes" :value="type.id" :key="type.id">{{ type.name.toUpperCase() }}</option>
        </select>
      </div>

      <div class="space-y-1.5">
        <label class="text-[11px] font-black text-slate-500 uppercase tracking-wider ml-1">Estimasi Masa / Sejarah</label>
        <input :value="modelValue.tahun_berdiri" @input="e => updateField('tahun_berdiri', (e.target as HTMLInputElement).value)" type="text"
          class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200 rounded-xl focus:bg-white focus:border-primary transition-all text-sm font-bold text-slate-800 outline-none"
          placeholder="Contoh: Abad ke-14 atau Masa Kerajaan Majapahit">
      </div>

      <div class="space-y-1.5">
        <label class="text-[11px] font-black text-slate-500 uppercase tracking-wider ml-1">Foto Utama</label>
        <div class="flex gap-4 items-center">
          <div class="relative w-20 h-20 shrink-0">
            <div v-if="modelValue.cover_image" class="w-full h-full rounded-xl overflow-hidden border-2 border-primary/20">
              <img :src="modelValue.cover_image" class="w-full h-full object-cover">
            </div>
            <div v-else class="w-full h-full rounded-xl bg-slate-100 border-2 border-dashed border-slate-200 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-slate-300" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round"
                  d="M6.827 6.175A2.31 2.31 0 0 1 5.186 7.23c-.38.054-.757.112-1.134.175C2.999 7.58 2.25 8.507 2.25 9.574V18a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9.574c0-1.067-.75-1.994-1.802-2.169a47.865 47.865 0 0 0-1.134-.175 2.31 2.31 0 0 1-1.64-1.055l-.822-1.316a2.192 2.192 0 0 0-1.736-1.039 48.774 48.774 0 0 0-5.232 0 2.192 2.192 0 0 0-1.736 1.039l-.821 1.316Z" />
                <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 12.75a4.5 4.5 0 1 1-9 0 4.5 4.5 0 0 1 9 0ZM18.75 10.5h.008v.008h-.008V10.5Z" />
              </svg>
            </div>
            <!-- Cover Upload Loading Spinner -->
            <div v-if="isUploading" class="absolute inset-0 bg-white/60 backdrop-blur-[1px] rounded-xl flex items-center justify-center">
              <div class="w-6 h-6 border-2 border-primary/20 border-t-primary rounded-full animate-spin"></div>
            </div>
          </div>

          <div class="flex-1">
            <input type="file" @change="handleFileUpload" class="hidden" id="photo-upload" accept="image/*">
            <label for="photo-upload"
              class="inline-flex items-center px-4 py-2 bg-primary hover:bg-primary/90 text-white text-xs font-bold rounded-lg cursor-pointer transition-all disabled:opacity-50">
              {{ isUploading ? 'Mohon Tunggu...' : 'Ganti Foto Utama' }}
            </label>
            <p class="text-[10px] text-slate-400 mt-1">Gunakan foto terbaik sebagai cover destinasi.</p>
          </div>
        </div>
      </div>

      <div class="space-y-1.5">
        <label class="text-[11px] font-black text-slate-500 uppercase tracking-wider ml-1">Alamat Lengkap</label>
        <textarea :value="modelValue.address" @input="e => updateField('address', (e.target as HTMLTextAreaElement).value)" rows="2"
          class="w-full px-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:bg-white focus:border-primary transition-all text-xs font-bold text-slate-600 outline-none"
          placeholder="Alamat objek..."></textarea>
      </div>

    </template>

    <template v-else>
      <div class="flex-1 flex flex-col min-h-0 relative">
        <div class="flex items-center justify-between mb-2 px-1">
          <label class="text-[11px] font-black text-slate-500 uppercase tracking-wider">Isi Artikel / Ulasan</label>
        </div>

        <div class="flex-1 flex flex-col border border-slate-200 rounded-2xl bg-white relative overflow-hidden shadow-inner">
          <!-- Precision Loading Spinner at Cursor -->
          <div v-if="isBlogImageUploading && cursorBounds"
            class="absolute z-50 pointer-events-none flex items-center gap-2 px-2 py-1 bg-white/80 backdrop-blur shadow-sm rounded-lg border border-primary/20 animate-in fade-in zoom-in duration-300"
            :style="{
              top: `${cursorBounds.top + (quillRef?.getQuill()?.container?.querySelector('.ql-toolbar')?.offsetHeight || 0) + 20}px`,
              left: `${cursorBounds.left + 20}px`
            }">
            <div class="w-4 h-4 border-2 border-primary/20 border-t-primary rounded-full animate-spin"></div>
            <span class="text-[9px] font-black text-primary uppercase tracking-widest">Uploading...</span>
          </div>

          <QuillEditor ref="quillRef" :content="blogContent.content" contentType="html" @update:content="val => emit('update:blogContent', { ...blogContent, content: val })" @ready="onEditorReady"
            theme="snow" :toolbar="[
              ['bold', 'italic', 'underline', 'strike'],
              [{ 'header': 1 }, { 'header': 2 }],
              [{ 'list': 'ordered' }, { 'list': 'bullet' }],
              [{ 'align': [] }],
              ['image', 'link', 'clean']
            ]" :modules="modules" placeholder="Tuliskan ulasan mendalam tentang destinasi ini..." class="flex-1 overflow-hidden flex flex-col" />
        </div>
      </div>
    </template>

  </div>
</template>

<style>
.ql-toolbar.ql-snow {
  position: sticky !important;
  top: -20px;
  z-index: 50;
  border: none !important;
  border-bottom: 1px solid #f1f5f9 !important;
  background: rgba(255, 255, 255, 0.95) !important;
  backdrop-filter: blur(8px);
  padding: 12px 15px !important;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
}

.ql-container.ql-snow {
  border: none !important;
  font-family: 'Urbanist', sans-serif;
}

.ql-editor {
  min-height: 500px;
  padding: 35px !important;
  line-height: 1.8 !important;
}

.flex-1.flex.flex-col.border.border-slate-200.rounded-2xl {
  overflow: visible !important;
}

.prose-bali,
.ql-editor {
  font-family: 'Urbanist', sans-serif !important;
  font-size: 16px !important;
  color: #334155 !important;
}

.prose-bali h1,
.ql-editor h1 {
  font-size: 2.25rem !important;
  font-weight: 800 !important;
  margin-top: 1.5rem !important;
  margin-bottom: 1rem !important;
  letter-spacing: -0.025em !important;
  color: #0f172a !important;
}

.prose-bali h2,
.ql-editor h2 {
  font-size: 1.5rem !important;
  font-weight: 700 !important;
  margin-top: 1.25rem !important;
  margin-bottom: 0.75rem !important;
  color: #1e293b !important;
}

.prose-bali p,
.ql-editor p {
  margin-bottom: 1.25rem !important;
}
</style>
