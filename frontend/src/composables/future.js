// taken from https://vuejs.org/guide/reusability/composables#async-state-example
import { ref } from 'vue'

export function useFuture(promise) {
  const data = ref(null)
  const error = ref(null)

  promise
  .then(val => data.value = val)
  .catch(err => data.error)

  return { data, error }
}