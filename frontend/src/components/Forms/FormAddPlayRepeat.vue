<template>
  <v-autocomplete required v-model="form.selectedPlay" :items="formatedPlays" label="Partie à répéter"
    placeholder="Sélectionner la partie jouée" variant="outlined" item-title="title" item-value="value">
    <template v-slot="{ props, item }">
      <v-chip v-bind="props" class="bg-secondaryContainer">
        {{ item.title }}
      </v-chip>
    </template>
  </v-autocomplete>

  <v-text-field required type="date" v-model="form.date" label="Date de la partie" placeholder="dd/mm/yyyy"
    variant="outlined" density="comfortable" color="secondaryContainer" />

  <v-text-field required v-model="form.location" label="Lieu" placeholder="Lieu ou s'est déroulée la partie"
    variant="outlined" density="comfortable" color="secondaryContainer" />

  <v-text-field required v-model="form.duration" label="Durée (minutes)" placeholder="Durée de la partie en minutes"
    variant="outlined" density="comfortable" type="number" color="secondaryContainer" />

  <v-autocomplete required v-model="form.players" :items="formattedPlayers" label="Participant·e·s" chips multiple
    placeholder="Ajouter des participant·e·s" variant="outlined" item-title="title" item-value="value">
    <template v-slot:chip="{ props, item }">
      <v-chip v-bind="props" class="bg-secondaryContainer">
        {{ item.raw.title }}
      </v-chip>
    </template>
  </v-autocomplete>

  <v-autocomplete v-model="form.winners" :items="winnablePlayers" label="Gagnant·e·s" chips multiple
    placeholder="Sélectionner les gagnant·e·s" variant="outlined" item-title="title" item-value="value">
    <template v-slot:chip="{ props, item }">
      <v-chip v-bind="props" class="bg-secondaryContainer">
        {{ item.title }}
      </v-chip>
    </template>
  </v-autocomplete>
</template>

<script setup>
import { useAppStore } from "@/stores/app";
import { storeToRefs } from "pinia";
import { ref, computed, watch } from "vue";
const appStore = useAppStore()
const reactiveAppStore = storeToRefs(appStore)
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({}),
  },
});

const emit = defineEmits(["update:modelValue"]);

const availablePlayers = reactiveAppStore.userData.value.persons
const availablePlays = reactiveAppStore.userData.value.plays

const formattedPlayers = computed(() => {
  return availablePlayers.map(player => ({
    title: player.name,
    value: player.id
  }));
});

const formatedPlays = computed(() => {
  return availablePlays?.map(play => ({
    title: `Partie du ${(new Date(play.date)).toLocaleDateString()}`,
    value: play.id
  }));
});
const form = ref({
  selectedPlay: "",
  date: "",
  location: "",
  duration: "",
  players: [],
  winners: [],
});

const winnablePlayers = computed(() => {
  return [...formattedPlayers.value]?.filter(p => form.value.players?.includes(p.value))
})

const isValid = computed(() => {
  return form.value.selectedGame &&
    form.value.date &&
    form.value.location &&
    form.value.duration &&
    form.value.players?.length > 0;
});

watch(
  form,
  (newValue) => {
    emit("update:modelValue", { ...newValue, isValid: isValid.value });
  },
  { deep: true }
);

watch(
  () => props.modelValue,
  (newValue) => {
    if (Object.keys(newValue).length) {
      form.value = { ...newValue };
    }
  },
  { immediate: true }
);
</script>
