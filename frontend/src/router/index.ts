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
          name: 'home',
          component: () => import('@/features/home/HomePage.vue'),
        },
        {
          path: 'explore',
          name: 'public-map',
          component: () => import('@/features/explore/PublicMapPage.vue'),
        },
        {
          path: 'dashboard',
          name: 'dashboard',
          component: () => import('@/features/management/pages/AdminDashboard.vue'),
          meta: { requiresAuth: true },
        },
        {
          path: 'tabular',
          name: 'tabular-view',
          component: () => import('@/features/points/pages/PointListPage.vue'),
          meta: { requiresAuth: true },
        },
        {
          path: 'blog/:id',
          name: 'blog-detail',
          component: () => import('@/features/blog/BlogDetailPage.vue'),
        },
        // Admin Features
        {
          path: 'admin/categories',
          name: 'admin-categories',
          component: () => import('@/features/admin/CategoriesPage.vue'),
          meta: { requiresAuth: true, requiresAdmin: true },
        },
        {
          path: 'admin/users',
          name: 'admin-users',
          component: () => import('@/features/admin/UsersPage.vue'),
          meta: { requiresAuth: true, requiresAdmin: true },
        },
        {
          path: 'admin/verification',
          name: 'admin-verification',
          component: () => import('@/features/admin/VerificationPage.vue'),
          meta: { requiresAuth: true, requiresAdmin: true },
        },
        // Contributor Features
        {
          path: 'contributor/dashboard',
          name: 'contributor-dashboard',
          component: () => import('@/features/management/pages/ContributorDashboard.vue'),
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
    return { name: 'public-map' }
  }
})

export default router
