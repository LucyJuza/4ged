// taken from https://vuejs.org/guide/reusability/composables#async-state-example
import { ref } from 'vue'

export function useFutureList(promiseList) {
  const data = ref(null)
  const error = ref(null)
  
  Promise.all(promiseList)
  .then(vals => data.value = vals)
  .catch(err => data.error)
  

  return { data, error }
}