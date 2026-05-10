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
      component: () => import('@/features/auth/AuthPage.vue'),
    },
    // All app pages - wrapped inside MainLayout
    {
      path: '/',
      component: MainLayout,
      children: [
        {
          path: '',
          name: 'public-map',
          component: () => import('@/features/map/PublicMapPage.vue'),
        },
        {
          path: 'dashboard',
          name: 'dashboard',
          component: () => import('@/features/dashboard/DashboardPage.vue'),
          meta: { requiresAuth: true },
        },
        {
          path: 'tabular',
          name: 'tabular-view',
          component: () => import('@/features/dashboard/TabularPage.vue'),
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

  if (to.meta.requiresAdmin && !authStore.isAdmin()) {
    return { name: 'public-map' } // Redirect to home if not admin
  }
})

export default router
