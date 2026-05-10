import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useAuthStore } from './auth'
import { useNotificationStore } from './notifications'
import { useMapUIStore } from './mapUI'

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

export interface GeoPoint {
  id?: number
  name: string
  latitude: number
  longitude: number
  type_id: number
  address?: string
  tahun_berdiri?: number
  status_kepemilikan?: string
  description?: string
  owner_id?: string
  owner_name?: string
  is_active?: boolean
  status?: 'draft' | 'pending' | 'approved' | 'rejected'
  rejection_note?: string
  created_at?: string
  updated_at?: string
}

export interface ObjectType {
  id: number
  name: string
  icon: string
}

export const useMapPointsStore = defineStore('mapPoints', () => {
  const points = ref<GeoPoint[]>([])
  const objectTypes = ref<ObjectType[]>([])
  const isModalOpen = ref(false)
  const activePoint = ref<GeoPoint | null>(null)

  const confirmState = ref({
    isOpen: false,
    title: '',
    message: '',
    onConfirm: () => {},
  })

  const notificationStore = useNotificationStore()

  const fetchPoints = async () => {
    try {
      const res = await fetch(`${API_URL}/points`)
      const json = await res.json()
      if (json.success) points.value = json.data
    } catch {
      notificationStore.error('Gagal mengambil data bangunan')
    }
  }

  const fetchCategories = async () => {
    try {
      const res = await fetch(`${API_URL}/categories`)
      const json = await res.json()
      if (json.success) objectTypes.value = json.data
    } catch {
      notificationStore.error('Gagal mengambil data kategori')
    }
  }

  const savePoint = async (pointData: GeoPoint) => {
    const authStore = useAuthStore()
    const token = authStore.token
    const isEdit = !!pointData.id
    const method = isEdit ? 'PUT' : 'POST'
    const url = isEdit ? `${API_URL}/points/${pointData.id}` : `${API_URL}/points`

    try {
      const res = await fetch(url, {
        method,
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(pointData),
      })
      const json = await res.json()
      if (json.success) {
        const savedPoint = json.data
        if (!savedPoint.owner_name && authStore.user) savedPoint.owner_name = authStore.user.name

        if (isEdit) {
          const index = points.value.findIndex((p) => p.id === pointData.id)
          if (index !== -1) points.value[index] = savedPoint
        } else {
          points.value.push(savedPoint)
        }
        notificationStore.success(isEdit ? 'Data berhasil diperbarui' : 'Data berhasil ditambahkan')
        closeModal()
        return savedPoint
      } else {
        notificationStore.error(json.message || 'Gagal menyimpan data')
      }
    } catch {
      notificationStore.error('Terjadi kesalahan jaringan')
    }
    return null
  }

  const deletePoint = async (id: number) => {
    const authStore = useAuthStore()
    const token = authStore.token

    try {
      const res = await fetch(`${API_URL}/points/${id}`, {
        method: 'DELETE',
        headers: { Authorization: `Bearer ${token}` },
      })
      const json = await res.json()
      if (json.success) {
        points.value = points.value.filter((p) => p.id !== id)
        notificationStore.success('Data berhasil dihapus')
      } else {
        notificationStore.error(json.message || 'Gagal menghapus data')
      }
    } catch {
      notificationStore.error('Terjadi kesalahan jaringan')
    }
  }

  const saveBlog = async (
    pointId: number,
    blogData: { title: string; content: string; cover_photo?: string },
  ) => {
    const authStore = useAuthStore()
    try {
      const res = await fetch(`${API_URL}/points/${pointId}/blog`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${authStore.token}`,
        },
        body: JSON.stringify(blogData),
      })
      const json = await res.json()
      if (json.success) {
        notificationStore.success('Konten blog berhasil disimpan')
        return json.data
      } else {
        notificationStore.error(json.message || 'Gagal menyimpan blog')
      }
    } catch {
      notificationStore.error('Kesalahan jaringan saat menyimpan blog')
    }
    return null
  }

  const getBlog = async (pointId: number) => {
    try {
      const res = await fetch(`${API_URL}/points/${pointId}/blog`)
      const json = await res.json()
      if (json.success) return json.data
    } catch {
      // It's okay if blog is not found
    }
    return null
  }

  const openModal = (point?: GeoPoint) => {
    activePoint.value = point ? { ...point } : null
    isModalOpen.value = true
  }

  const closeModal = () => {
    isModalOpen.value = false
    activePoint.value = null
  }

  const requestConfirm = (title: string, message: string, onConfirm: () => void) => {
    confirmState.value = { isOpen: true, title, message, onConfirm }
  }

  const cancelConfirm = () => {
    confirmState.value.isOpen = false
  }

  const filteredPoints = computed(() => {
    const uiStore = useMapUIStore()
    const authStore = useAuthStore()

    return points.value.filter((point) => {
      const matchSearch =
        point.name.toLowerCase().includes(uiStore.searchQuery.toLowerCase()) ||
        (point.address && point.address.toLowerCase().includes(uiStore.searchQuery.toLowerCase()))
      const matchType = !uiStore.filterTypeId || point.type_id === uiStore.filterTypeId
      const matchOwner = !uiStore.filterMyPoints || point.owner_id === authStore.user?.id

      return matchSearch && matchType && matchOwner
    })
  })

  return {
    points,
    objectTypes,
    isModalOpen,
    activePoint,
    confirmState,
    filteredPoints,
    fetchPoints,
    fetchCategories,
    savePoint,
    deletePoint,
    saveBlog,
    getBlog,
    openModal,
    closeModal,
    requestConfirm,
    cancelConfirm,
  }
})
