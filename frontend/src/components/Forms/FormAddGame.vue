<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-text-field
          v-model="form.name"
          label="Nom"
          placeholder="Mon super jeu custom"
          variant="outlined"
          density="comfortable"
        />
      </v-col>

      <v-col cols="12">
        <v-text-field
          v-model="form.image"
          label="Image"
          placeholder="Lien vers l'image du jeu"
          variant="outlined"
          density="comfortable"
        />
      </v-col>

      <v-col cols="12">
        <v-autocomplete
          v-model="form.genres"
          :items="items"
          label="Genres"
          chips
          multiple
          placeholder="Ajouter des genres"
          variant="outlined"
        >
          <template v-slot:chip="{ props, item }">
              <v-chip
                v-bind="props"
                class="bg-secondaryContainer"
              >
                {{ item.raw }}
              </v-chip>
            </template>
        </v-autocomplete>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { useAppStore } from '@/stores/app'
import { ref, watch } from 'vue'
const appStore = useAppStore()
const items = appStore.genres.map(g => g.name)
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['update:modelValue'])

const form = ref({
  name: '',
  image: null,
  genres: []
})

const imagePreview = ref(null)

const handleImageChange = (file) => {
  if (!file) {
    imagePreview.value = null
    return
  }

  const reader = new FileReader()
  reader.onload = (e) => {
    imagePreview.value = e.target.result
  }
  reader.readAsDataURL(file)
}

// Synchronise les changements avec le parent via v-model
watch(form, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })

// Initialise le formulaire avec les valeurs existantes si présentes
watch(() => props.modelValue, (newValue) => {
  if (Object.keys(newValue).length) {
    form.value = { ...newValue }
    if (newValue.image && newValue.image instanceof File) {
      handleImageChange(newValue.image)
    }
  }
}, { immediate: true })
</script>