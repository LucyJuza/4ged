<template>
  <v-text-field v-model="form.name" label="Nom" placeholder="Nom de la personne" variant="outlined"
    density="comfortable" color="secondaryContainer" />

  <v-text-field v-model="form.image" label="Image" placeholder="Lien vers l'image pour la personne" variant="outlined"
    density="comfortable" />
</template>

<script setup>
import { ref, watch } from "vue";

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({}),
  },
});

const emit = defineEmits(["update:modelValue"]);

const form = ref({
  name: "",
  image: null,
});

const imagePreview = ref(null);

const handleImageChange = (file) => {
  if (!file) {
    imagePreview.value = null;
    return;
  }

  const reader = new FileReader();
  reader.onload = (e) => {
    imagePreview.value = e.target.result;
  };
  reader.readAsDataURL(file);
};

watch(
  form,
  (newValue) => {
    emit("update:modelValue", newValue);
  },
  { deep: true }
);

watch(
  () => props.modelValue,
  (newValue) => {
    if (Object.keys(newValue).length) {
      form.value = { ...newValue };
      if (newValue.image && newValue.image instanceof File) {
        handleImageChange(newValue.image);
      }
    }
  },
  { immediate: true }
);
</script>
