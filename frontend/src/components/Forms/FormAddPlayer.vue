<template>
  <v-text-field required v-model="form.name" label="Nom" placeholder="Nom de la personne" variant="outlined"
    density="comfortable" color="secondaryContainer" />
  <v-file-input v-model="form.previewImage" accept="image/*" label="Image" prepend-icon="mdi-camera" show-size
    variant="outlined" required @update:model-value="handleImageChange" />
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
  previewImage: null,
});

const handleImageChange = async (file) => {
  if (file) {
    try {
      form.value.image = await convertToBase64(file);
      if (form.value.previewImage) {
        URL.revokeObjectURL(previewImage);
      }
      form.value.previewImage = URL.createObjectURL(file);
    } catch (error) {
      showError('Erreur lors du traitement de l\'image');
    }
  } else {
    if (form.value.previewImage) {
      URL.revokeObjectURL(form.value.previewImage);
    }
    form.value.previewImage = null;
    form.value.image = null;
  }
};

const convertToBase64 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
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
