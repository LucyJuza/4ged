<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-text-field
          v-model="form.selectedGame"
          label="Jeu joué"
          placeholder="Nom du jeu"
          variant="outlined"
          density="comfortable"
          color="secondaryContainer"
        />
      </v-col>

      <v-col cols="12">
        <v-text-field
          v-model="form.date"
          label="Date de la partie"
          placeholder="dd/mm/yyyy"
          variant="outlined"
          density="comfortable"
          color="secondaryContainer"
        />
      </v-col>

      <v-col cols="12">
        <v-text-field
          v-model="form.location"
          label="Lieu"
          placeholder="Lieu ou s'est déroulée la partie"
          variant="outlined"
          density="comfortable"
          color="secondaryContainer"
        />
      </v-col>

      <v-col cols="12">
        <v-text-field
          v-model="form.duration"
          label="Durée (minutes)"
          placeholder="Durée de la partie en minutes"
          variant="outlined"
          density="comfortable"
          type="number"
          color="secondaryContainer"
        />
      </v-col>

      <v-col cols="12">
        <v-autocomplete
          v-model="form.players"
          :items="formattedPlayers"
          label="Participant·e·s"
          chips
          multiple
          placeholder="Ajouter des participant·e·s"
          variant="outlined"
          item-title="title"
          item-value="value"
        >
          <template v-slot:chip="{ props, item }">
            <v-chip
                  v-bind="props"
              class="bg-secondaryContainer"
                >
              {{ item.raw.title }}
                  </v-chip>
          </template>
        </v-autocomplete>
      </v-col>

      <v-col cols="12">
        <v-autocomplete
          v-model="form.winners"
          :items="formattedPlayers"
          label="Gagnant·e·s"
          chips
          multiple
          placeholder="Sélectionner les gagnant·e·s"
          variant="outlined"
          item-title="title"
          item-value="value"
        >
          <template v-slot:chip="{ props, item }">
                  <v-chip
              v-bind="props"
              class="bg-secondaryContainer"
                  >
              {{ item.raw.title }}
                  </v-chip>
          </template>
        </v-autocomplete>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref, computed, watch } from "vue";

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({}),
  },
});

const emit = defineEmits(["update:modelValue"]);

const availablePlayers = ref([
  { id: 1, name: "Joueur 1" },
  { id: 2, name: "Joueur 2" },
]);

const formattedPlayers = computed(() => {
  return availablePlayers.value.map(player => ({
    title: player.name,
    value: player.id
  }));
});

const form = ref({
  selectedGame: "",
  date: "",
  location: "",
  duration: "",
  players: [],
  winners: [],
});

const isValid = computed(() => {
  return form.value.selectedGame &&
         form.value.date &&
         form.value.location &&
         form.value.duration &&
         form.value.players.length > 0;
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
