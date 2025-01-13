<template>
</template>
<route lang="yaml">
meta:
  layout: login
</route>
<script setup>
import { useAuthentication } from '@/composables/authentication';
import { useRouter } from 'vue-router';

const router = useRouter()
const { isAuthenticated } = useAuthentication()
console.log(isAuthenticated.value)
if (isAuthenticated.value) {
  router.push('/home')
}else {
  router.push('/login')
}
</script>