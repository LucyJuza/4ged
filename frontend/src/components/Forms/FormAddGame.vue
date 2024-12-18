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
          color="secondaryContainer"
        />
      </v-col>

      <v-col cols="12">
        <v-file-input
          v-model="form.image"
          label="Image"
          placeholder="File input"
          variant="outlined"
          density="comfortable"
          accept="image/*"
          prepend-icon="mdi-camera"
          color="secondaryContainer"
          @update:model-value="handleImageChange"
        >
          <template v-slot:prepend>
            <div class="mr-2">
              <v-avatar v-if="imagePreview" size="40" rounded>
                <v-img :src="imagePreview" cover />
              </v-avatar>
            </div>
          </template>
        </v-file-input>
      </v-col>

      <v-col cols="12">
        <div class="mb-2">
          <label class="secondaryContainer">Genres</label>
        </div>
        <v-combobox
          v-model="form.genres"
          chips
          multiple
          closable-chips
          placeholder="Ajouter des genres"
          variant="outlined"
          density="comfortable"
          color="secondaryContainer"
        >
          <template v-slot:chip="{ props, item }">
            <v-chip
              v-bind="props"
            >
              {{ item.raw }}
            </v-chip>
          </template>
        </v-combobox>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref, watch } from 'vue'

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