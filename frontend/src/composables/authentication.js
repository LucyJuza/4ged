import { useAppStore } from '@/stores/app'
import { storeToRefs } from 'pinia'
import { ref, watchEffect } from 'vue'
const appStore = useAppStore()
const { userId } = storeToRefs(appStore) 

export function useAuthentication() {
  const isAuthenticated = ref(false)
  const checkAuth = () => {
    if( userId.value ){
      isAuthenticated.value = true
    }else {
      isAuthenticated.value = false
    }
  }
  watchEffect(() => {
    checkAuth()
  })
  return { isAuthenticated }
}