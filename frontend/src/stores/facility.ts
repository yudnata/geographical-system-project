import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface Facility {
  id: number
  name: string
  type: string
  address: string
  district: string
  phone: string
  hours: string
  lat: number
  lng: number
}

const defaultData: Facility[] = [
  {
    id: 1,
    name: 'RSUP Prof. Dr. I.G.N.G. Ngoerah',
    type: 'Rumah Sakit',
    address: 'Jl. Diponegoro, Dauh Puri Klod',
    district: 'Denpasar Barat',
    phone: '(0361) 227911',
    hours: '24 Jam',
    lat: -8.6756,
    lng: 115.2052,
  },
  {
    id: 2,
    name: 'Puskesmas Denpasar Selatan I',
    type: 'Puskesmas',
    address: 'Jl. Gurita No.8, Sesetan',
    district: 'Denpasar Selatan',
    phone: '(0361) 723522',
    hours: '08:00 - 14:00',
    lat: -8.6922,
    lng: 115.2255,
  },
  {
    id: 3,
    name: 'Klinik Padma Bahtera',
    type: 'Klinik',
    address: 'Jl. Gatot Subroto Timur No.296',
    district: 'Denpasar Timur',
    phone: '(0361) 461122',
    hours: '08:00 - 21:00',
    lat: -8.636,
    lng: 115.239,
  },
  {
    id: 4,
    name: 'Apotek Kimia Farma Teuku Umar',
    type: 'Apotek',
    address: 'Jl. Teuku Umar No.69',
    district: 'Denpasar Barat',
    phone: '(0361) 223590',
    hours: '07:00 - 22:00',
    lat: -8.674,
    lng: 115.208,
  },
  {
    id: 5,
    name: 'RSUD Wangaya',
    type: 'Rumah Sakit',
    address: 'Jl. Kartini No.133',
    district: 'Denpasar Utara',
    phone: '(0361) 222141',
    hours: '24 Jam',
    lat: -8.6475,
    lng: 115.2066,
  },
]

export const useFacilityStore = defineStore('facility', () => {
  const facilities = ref<Facility[]>([])

  function initFacilities() {
    const saved = localStorage.getItem('sig_facilities')
    if (saved) {
      facilities.value = JSON.parse(saved)
    } else {
      facilities.value = [...defaultData]
      saveToLocal()
    }
  }

  function saveToLocal() {
    localStorage.setItem('sig_facilities', JSON.stringify(facilities.value))
  }

  function addFacility(facility: Omit<Facility, 'id'>) {
    facilities.value.push({ ...facility, id: Date.now() })
    saveToLocal()
  }

  function updateFacility(id: number, updatedData: Omit<Facility, 'id'>) {
    const index = facilities.value.findIndex((f) => f.id === id)
    if (index !== -1) {
      facilities.value[index] = { ...updatedData, id }
      saveToLocal()
    }
  }

  function deleteFacility(id: number) {
    facilities.value = facilities.value.filter((f) => f.id !== id)
    saveToLocal()
  }

  const statsByType = computed(() => {
    return facilities.value.reduce(
      (acc, curr) => {
        acc[curr.type] = (acc[curr.type] || 0) + 1
        return acc
      },
      {} as Record<string, number>,
    )
  })

  return {
    facilities,
    initFacilities,
    addFacility,
    updateFacility,
    deleteFacility,
    statsByType,
  }
})
