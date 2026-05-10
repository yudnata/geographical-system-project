<script setup lang="ts">
import { useMapPointsStore, type GeoPoint } from '@/stores/mapPoints'
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
      })
  })
}

const imageHandler = () => {
  const input = document.createElement('input')
  input.type = 'file'
  input.accept = 'image/*'
  input.click()

  input.onchange = async () => {
    const file = input.files?.[0]
    if (!file) return

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
  if (quill) {
    quill.getModule('toolbar').addHandler('image', imageHandler)
  }
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
          <div v-if="modelValue.cover_image" class="w-20 h-20 rounded-xl overflow-hidden border-2 border-primary/20 shrink-0">
            <img :src="modelValue.cover_image" class="w-full h-full object-cover">
          </div>
          <div v-else class="w-20 h-20 rounded-xl bg-slate-100 border-2 border-dashed border-slate-200 flex items-center justify-center shrink-0">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-slate-300" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round"
                d="M6.827 6.175A2.31 2.31 0 0 1 5.186 7.23c-.38.054-.757.112-1.134.175C2.999 7.58 2.25 8.507 2.25 9.574V18a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9.574c0-1.067-.75-1.994-1.802-2.169a47.865 47.865 0 0 0-1.134-.175 2.31 2.31 0 0 1-1.64-1.055l-.822-1.316a2.192 2.192 0 0 0-1.736-1.039 48.774 48.774 0 0 0-5.232 0 2.192 2.192 0 0 0-1.736 1.039l-.821 1.316Z" />
              <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 12.75a4.5 4.5 0 1 1-9 0 4.5 4.5 0 0 1 9 0ZM18.75 10.5h.008v.008h-.008V10.5Z" />
            </svg>
          </div>
          <div class="flex-1">
            <input type="file" @change="handleFileUpload" class="hidden" id="photo-upload" accept="image/*">
            <label for="photo-upload"
              class="inline-flex items-center px-4 py-2 bg-primary hover:bg-primary/90 text-white text-xs font-bold rounded-lg cursor-pointer transition-all disabled:opacity-50">
              {{ isUploading ? 'Uploading...' : 'Unggah Foto' }}
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
      <div class="h-full flex flex-col">
        <div class="flex items-center justify-between mb-2 px-1">
          <label class="text-[11px] font-black text-slate-500 uppercase tracking-wider">Isi Artikel / Ulasan</label>
          <span v-if="isBlogImageUploading" class="text-[10px] text-primary font-bold animate-pulse flex items-center gap-1">
            <svg class="w-3 h-3 animate-spin" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
            </svg>
            Mengunggah gambar...
          </span>
        </div>
        <div class="flex-1 min-h-[450px] border border-slate-200 rounded-2xl overflow-hidden shadow-inner">
          <QuillEditor ref="quillRef" :content="blogContent.content" contentType="html" @update:content="val => emit('update:blogContent', { ...blogContent, content: val })" @ready="onEditorReady"
            theme="snow" :toolbar="[
              ['bold', 'italic', 'underline', 'strike'],
              [{ 'header': 1 }, { 'header': 2 }],
              [{ 'size': ['small', false, 'large', 'huge'] }],
              [{ 'list': 'ordered' }, { 'list': 'bullet' }],
              [{ 'align': [] }],
              ['image', 'clean']
            ]" :modules="modules" placeholder="Tuliskan ulasan mendalam tentang destinasi ini... Anda bisa menyisipkan gambar juga." class="h-full prose-bali" />
        </div>
      </div>
    </template>

  </div>
</template>

<style>
.prose-bali,
.ql-editor {
  font-family: 'Urbanist', sans-serif !important;
  font-size: 16px !important;
  line-height: 1.8 !important;
  color: #334155 !important;
}

.prose-bali h1,
.ql-editor h1 {
  font-size: 2.25rem !important;
  font-weight: 800 !important;
  margin-top: 1.5rem !important;
  margin-bottom: 1rem !important;
  letter-spacing: -0.025em !important;
}

.prose-bali h2,
.ql-editor h2 {
  font-size: 1.5rem !important;
  font-weight: 700 !important;
  margin-top: 1.25rem !important;
  margin-bottom: 0.75rem !important;
}

.prose-bali p,
.ql-editor p {
  margin-bottom: 1.25rem !important;
}

.ql-editor {
  caret-color: #000000 !important;
  min-height: 400px;
}

.flex-1.min-h-\[450px\]:focus-within {
  border-color: #cbd5e1 !important;
  box-shadow: 0 4px 12px -2px rgba(0, 0, 0, 0.05) !important;
}

.ql-container.ql-snow {
  border: none !important;
  font-family: 'Urbanist', sans-serif;
}

.ql-toolbar.ql-snow {
  border: none !important;
  border-bottom: 1px solid #f1f5f9 !important;
  background: #f8fafc;
  padding: 10px 15px !important;
}

.ql-editor.ql-blank::before {
  color: #94a3b8 !important;
  font-style: normal !important;
}
</style>
