<script setup lang="ts">
import { RouterView } from 'vue-router'
import NotificationToast from '@/components/UI/NotificationToast.vue'
import ConfirmModal from '@/components/Modals/ConfirmModal.vue'
import PointFormModal from '@/features/dashboard/components/PointFormModal.vue'
import ProfileCompletionModal from '@/features/auth/components/ProfileCompletionModal.vue'
import UserProfileModal from '@/features/auth/components/UserProfileModal.vue'

</script>

<template>
  <RouterView />
  <NotificationToast />
  <ConfirmModal />
  <PointFormModal />
  <ProfileCompletionModal />
  <UserProfileModal />
</template>
