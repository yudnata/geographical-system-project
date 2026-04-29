import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useMapUIStore = defineStore('mapUI', () => {
  const isSidebarExpanded = ref(false)
  const isProfileModalOpen = ref(false)

  // Map Interaction State
  const isEditMode = ref(false)
  const flyToCoords = ref<{ lat: number; lng: number } | null>(null)

  // Search & Filter State
  const searchQuery = ref('')
  const filterTypeId = ref<number | null>(null)
  const filterMyPoints = ref(false)

  // Toggle Actions
  const toggleSidebar = () => {
    isSidebarExpanded.value = !isSidebarExpanded.value
  }

  const toggleEditMode = () => {
    isEditMode.value = !isEditMode.value
  }

  const openProfileModal = () => {
    isProfileModalOpen.value = true
  }

  const closeProfileModal = () => {
    isProfileModalOpen.value = false
  }

  const flyTo = (lat: number, lng: number) => {
    flyToCoords.value = { lat, lng }
  }

  return {
    isSidebarExpanded,
    isProfileModalOpen,
    isEditMode,
    flyToCoords,
    searchQuery,
    filterTypeId,
    filterMyPoints,
    toggleSidebar,
    toggleEditMode,
    openProfileModal,
    closeProfileModal,
    flyTo,
  }
})
