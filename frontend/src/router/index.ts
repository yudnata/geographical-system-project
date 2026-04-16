import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: () => import('@/pages/HomePage.vue'),
    },
    {
      path: '/add',
      name: 'add-facility',
      component: () => import('@/pages/FacilityManagePage.vue'),
    },
    {
      path: '/edit/:id',
      name: 'edit-facility',
      component: () => import('@/pages/FacilityManagePage.vue'),
    },
    {
      path: '/list',
      name: 'facility-list',
      component: () => import('@/pages/FacilityListPage.vue'),
    },
  ],
})

export default router
