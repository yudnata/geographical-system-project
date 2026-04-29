import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import MainLayout from '@/layouts/MainLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // Login page - completely standalone, no layout
    {
      path: '/login',
      name: 'login',
      component: () => import('@/pages/AuthPage.vue'),
    },
    // All app pages - wrapped inside MainLayout
    {
      path: '/',
      component: MainLayout,
      children: [
        {
          path: '',
          name: 'public-map',
          component: () => import('@/pages/PublicMapPage.vue'),
        },
        {
          path: 'dashboard',
          name: 'dashboard',
          component: () => import('@/pages/DashboardPage.vue'),
          meta: { requiresAuth: true },
        },
        {
          path: 'tabular',
          name: 'tabular-view',
          component: () => import('@/pages/TabularPage.vue'),
          meta: { requiresAuth: true },
        },
      ],
    },
  ],
})

router.beforeEach((to) => {
  const authStore = useAuthStore()
  if (to.meta.requiresAuth && !authStore.isAuthenticated()) {
    return { name: 'login' }
  }
})

export default router
