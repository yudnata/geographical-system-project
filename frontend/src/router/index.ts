import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import MainLayout from '@/layouts/MainLayout.vue'
import AdminLayout from '@/layouts/AdminLayout.vue'
import ContributorLayout from '@/layouts/ContributorLayout.vue'

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
    // Admin pages - wrapped inside AdminLayout
    {
      path: '/admin',
      component: AdminLayout,
      meta: { requiresAuth: true, requiresAdmin: true },
      children: [
        {
          path: 'categories',
          name: 'admin-categories',
          component: () => import('@/features/admin/CategoriesPage.vue'),
        },
        {
          path: 'users',
          name: 'admin-users',
          component: () => import('@/features/admin/UsersPage.vue'),
        },
        {
          path: 'verification',
          name: 'admin-verification',
          component: () => import('@/features/admin/VerificationPage.vue'),
        },
      ],
    },
    // Contributor pages - wrapped inside ContributorLayout
    {
      path: '/contributor',
      component: ContributorLayout,
      meta: { requiresAuth: true },
      children: [
        {
          path: 'dashboard',
          name: 'contributor-dashboard',
          component: () => import('@/features/dashboard/ContributorDashboard.vue'),
        },
        {
          path: 'map',
          name: 'contributor-map',
          component: () => import('@/features/dashboard/DashboardPage.vue'),
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
