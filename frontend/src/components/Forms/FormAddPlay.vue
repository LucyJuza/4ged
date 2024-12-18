<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-text-field
          v-model="form.selectedGame"
          label="Partie sélectionnée"
          placeholder="partie du dd/mm/yyyy"
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
          label="Durée"
          placeholder="Durée de la partie en minutes"
          variant="outlined"
          density="comfortable"
          type="number"
          color="secondaryContainer"
        />
      </v-col>

      <v-col cols="12">
        <v-text-field
          label="Participant·e·s"
          readonly
          variant="outlined"
          density="comfortable"
          color="secondaryContainer"
          :model-value="null"
        >
          <template v-slot:append>
            <v-menu>
              <template v-slot:activator="{ props }">
                <v-btn
                  icon
                  variant="text"
                  v-bind="props"
                  color="secondaryContainer"
                >
                  <v-icon>mdi-plus</v-icon>
                </v-btn>
              </template>

              <v-list>
                <v-list-item
                  v-for="player in availablePlayers"
                  :key="player.id"
                  :value="player"
                  @click="togglePlayer(player)"
                >
                  <v-list-item-title>{{ player.name }}</v-list-item-title>
                </v-list-item>
              </v-list>
            </v-menu>
          </template>

          <template v-slot:default>
            <v-slide-group
              v-model="form.players"
              multiple
              show-arrows
              class="mt-1"
            >
              <v-slide-group-item
                v-for="player in form.players"
                :key="player.id"
                v-slot="{ isSelected, toggle }"
              >
                <v-scale-transition>
                  <v-chip
                    label
                    color="secondaryContainer"
                    variant="elevated"
                    class="text-body-2 mx-1"
                    closable
                    @click="toggle"
                    @click:close="removePlayer(player)"
                  >
                    {{ player.name }}
                  </v-chip>
                </v-scale-transition>
              </v-slide-group-item>
            </v-slide-group>
          </template>
        </v-text-field>
      </v-col>

      <v-col cols="12">
        <v-text-field
          label="Gagnant·e·s"
          readonly
          variant="outlined"
          density="comfortable"
          color="secondaryContainer"
          :model-value="null"
        >
          <template v-slot:default>
            <v-slide-group
              v-model="form.winners"
              multiple
              show-arrows
              class="mt-1"
            >
              <v-slide-group-item
                v-for="player in form.players"
                :key="player.id"
                v-slot="{ isSelected, toggle }"
              >
                <v-scale-transition>
                  <v-chip
                    label
                    color="secondaryContainer"
                    variant="elevated"
                    class="text-body-2 mx-1"
                    @click="toggleWinner(player)"
                  >
                    {{ player.name }}
                  </v-chip>
                </v-scale-transition>
              </v-slide-group-item>
            </v-slide-group>
          </template>
        </v-text-field>
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

const form = ref({
  selectedGame: "",
  date: "",
  location: "",
  duration: "",
  players: [],
  winners: [],
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
