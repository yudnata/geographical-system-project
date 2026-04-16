import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface ObjectType {
  id: number
  name: string
  icon: string
}

export interface GeoPoint {
  id: number
  type_id: number
  name: string
  latitude: number
  longitude: number
  address: string
  owner_id: string
  tahun_berdiri?: number
  status_kepemilikan?: string
  description?: string
  is_active: boolean
}

export const useMapPointsStore = defineStore('mapPoints', () => {
  const objectTypes = ref<ObjectType[]>([
    { id: 1, name: 'Rumah', icon: 'M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z' },
    { id: 2, name: 'Kantor', icon: 'M3 21h18M9 8h6M9 12h6M9 16h6M5 21V3h14v18' },
    { id: 3, name: 'Rumah Sakit', icon: 'M12 7v10M7 12h10' },
    { id: 4, name: 'Sekolah', icon: 'M22 10L12 5 2 10l10 5 10-5zM6 12v5c3.5 3 8.5 3 12 0v-5' },
    { id: 5, name: 'Tempat Ibadah', icon: 'M12 2L2 7v10l10 5 10-5V7L12 2zm0 18v-8' },
    {
      id: 6,
      name: 'Objek Wisata',
      icon: 'M12 21a9 9 0 1 0 0-18 9 9 0 0 0 0 18zM12 8l1.5 3.5L17 13l-3.5 1.5L12 18l-1.5-3.5L7 13l3.5-1.5L12 8z',
    },
    { id: 7, name: 'Kuliner', icon: 'M3 11l1-9h16l1 9c0 3-3 5-3 5s-1 6-1 7H7s-1-7-1-7-3-2-3-5z' },
    {
      id: 8,
      name: 'Mall',
      icon: 'M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4zM3 6h18M16 10a4 4 0 0 1-8 0',
    },
    { id: 9, name: 'Kantor Polisi', icon: 'M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z' },
    {
      id: 10,
      name: 'Pemadam Kebakaran',
      icon: 'M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3-1.072-2.143-.224-4.054 2-6 .5 2.5 2 4.9 4 6.5 2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.295 1-3a2.5 2.5 0 0 0 2.5 2.5z',
    },
    { id: 11, name: 'Taman', icon: 'M12 19V5M5 11l7-7 7 7M5 19h14' },
  ])

  const points = ref<GeoPoint[]>([])

  const isEditMode = ref(false)
  const isModalOpen = ref(false)
  const activePoint = ref<GeoPoint | null>(null)
  const isSidebarExpanded = ref(false)

  const searchQuery = ref('')
  const filterTypeId = ref<number | null>(null)
  const filterMyPoints = ref<boolean>(false)
  const currentUserId = ref<string>('')

  const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

  const fetchPoints = async () => {
    try {
      const res = await fetch(`${API_URL}/points`)
      const json = await res.json()
      if (json.success) {
        points.value = json.data
      }
    } catch (err) {
      console.error('Failed to fetch points:', err)
    }
  }

  const savePoint = async (pointData: Partial<GeoPoint>) => {
    try {
      const token = localStorage.getItem('auth_token') || ''
      const res = await fetch(`${API_URL}/points`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(pointData),
      })
      const json = await res.json()
      if (json.success) {
        points.value.push(json.data)
        return true
      }
      alert(json.message)
      return false
    } catch (err) {
      console.error('Failed to save point:', err)
      return false
    }
  }

  const filteredPoints = computed(() => {
    return points.value.filter((p) => {
      if (!p.is_active) return false
      if (filterTypeId.value && p.type_id !== filterTypeId.value) return false
      if (searchQuery.value) {
        const query = searchQuery.value.toLowerCase()
        if (!p.name.toLowerCase().includes(query) && !p.address.toLowerCase().includes(query))
          return false
      }
      return true
    })
  })

  const toggleEditMode = () => (isEditMode.value = !isEditMode.value)
  const openModal = (point?: GeoPoint) => {
    activePoint.value = point || null
    isModalOpen.value = true
  }
  const closeModal = () => {
    isModalOpen.value = false
    activePoint.value = null
  }

  return {
    objectTypes,
    points,
    isEditMode,
    isModalOpen,
    activePoint,
    isSidebarExpanded,
    searchQuery,
    filterTypeId,
    filterMyPoints,
    filteredPoints,
    currentUserId,
    fetchPoints,
    savePoint,
    toggleEditMode,
    openModal,
    closeModal,
  }
})
