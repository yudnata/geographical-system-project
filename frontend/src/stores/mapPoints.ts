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
  const objectTypes = ref<ObjectType[]>([
    {
      id: 1,
      name: 'Rumah',
      icon: 'M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6',
    },
    {
      id: 2,
      name: 'Kantor',
      icon: 'M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4',
    },
    {
      id: 3,
      name: 'Rumah Sakit',
      icon: 'M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z',
    },
    {
      id: 4,
      name: 'Sekolah',
      icon: 'M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253',
    },
    { id: 5, name: 'Tempat Ibadah', icon: 'M21 12a9 9 0 11-18 0 9 9 0 0118 0z M12 8v8 M8 12h8' },
    {
      id: 6,
      name: 'Objek Wisata',
      icon: 'M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z M15 11a3 3 0 11-6 0 3 3 0 016 0z',
    },
    {
      id: 7,
      name: 'Kuliner',
      icon: 'M10.5 3.5a2.121 2.121 0 00-3 0 3.84 3.84 0 00-1.05 3.8c.215.845.663 1.59 1.25 2.15a13.56 13.56 0 014.25 5.5c.34 1.15.66 2.3.95 3.45.1.35.45.6.85.6h.4c.4 0 .75-.25.85-.6.29-1.15.61-2.3.95-3.45a13.56 13.56 0 014.25-5.5 3.84 3.84 0 00.2-5.95 2.121 2.121 0 00-3 0l-2.5 2.5-2.5-2.5z',
    },
    { id: 8, name: 'Mall', icon: 'M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z' },
  ])

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
      } else {
        notificationStore.error(json.message || 'Gagal menyimpan data')
      }
    } catch {
      notificationStore.error('Terjadi kesalahan jaringan')
    }
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
    savePoint,
    deletePoint,
    openModal,
    closeModal,
    requestConfirm,
    cancelConfirm,
  }
})
