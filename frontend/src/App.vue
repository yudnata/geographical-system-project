<script setup lang="ts">
import { RouterView } from 'vue-router'
import NotificationToast from '@/components/UI/NotificationToast.vue'
import ConfirmModal from '@/components/Modals/ConfirmModal.vue'
import PointFormModal from '@/components/Modals/PointFormModal.vue'
import ProfileCompletionModal from '@/components/Modals/ProfileCompletionModal.vue'
import UserProfileModal from '@/components/Modals/UserProfileModal.vue'

</script>

<template>
  <RouterView />
  <NotificationToast />
  <ConfirmModal />
  <PointFormModal />
  <ProfileCompletionModal />
  <UserProfileModal />
</template>
