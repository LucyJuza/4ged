<template>
  <v-container fluid fill-height class="align-center justify-center">
    <v-row justify="center">
      <v-col cols="12" sm="8" md="6" lg="4">
        <v-card class="bg-tertiaryContainer">
          <v-toolbar color="primary">
            <v-toolbar-title>{{ isLogin ? 'Connexion' : 'Inscription' }}</v-toolbar-title>
          </v-toolbar>
          
          <v-card-text>
            <v-form ref="form" @update:model-value="onFormValidityChange">
              <v-container>
                <v-text-field
                  v-model="username"
                  :rules="usernameRules"
                  label="Nom d'utilisateur"
                  prepend-icon="mdi-account"
                  variant="outlined"
                  required
                  class="mb-6"
                ></v-text-field>

                <v-text-field
                  v-model="password"
                  :rules="passwordRules"
                  label="Mot de passe"
                  prepend-icon="mdi-lock"
                  :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                  @click:append="togglePassword"
                  :type="showPassword ? 'text' : 'password'"
                  required
                  variant="outlined"
                ></v-text-field>

                <v-text-field
                  v-if="!isLogin"
                  v-model="confirmPassword"
                  :rules="confirmPasswordRules"
                  label="Confirmer le mot de passe"
                  prepend-icon="mdi-lock"
                  :append-icon="showConfirmPassword ? 'mdi-eye' : 'mdi-eye-off'"
                  @click:append="toggleConfirmPassword"
                  :type="showConfirmPassword ? 'text' : 'password'"
                  required
                  class="mt-6"
                  variant="outlined"
                ></v-text-field>
              </v-container>
            </v-form>
          </v-card-text>

          <v-card-actions class="px-4 pb-4">
            <v-btn
              block
              color="primary"
              :loading="loading"
              :disabled="!formIsValid"
              @click="onSubmit"
              variant="flat"
            >
              {{ isLogin ? 'Se connecter' : "S'inscrire" }}
            </v-btn>
          </v-card-actions>

          <v-card-text class="text-center">
            <a href="#" @click.prevent="toggleForm">
              {{ isLogin ? "Pas encore de compte ? S'inscrire" : 'Déjà un compte ? Se connecter' }}
            </a>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-snackbar v-model="snackbar" :color="snackbarColor" timeout="3000">
      {{ snackbarText }}
      <template v-slot:action="{ attrs }">
        <v-btn text v-bind="attrs" @click="snackbar = false">
          Fermer
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script setup>
import { ref, computed } from 'vue';

const form = ref(null);
const formIsValid = ref(false);
const loading = ref(false);
const isLogin = ref(true);
const showPassword = ref(false);
const showConfirmPassword = ref(false);

const username = ref('');
const password = ref('');
const confirmPassword = ref('');

const snackbar = ref(false);
const snackbarText = ref('');
const snackbarColor = ref('success');

const usernameRules = [
  v => !!v || 'Ce champ est requis',
  v => v.length <= 50 || 'Le nom doit faire moins de 50 caractères'
];

const passwordRules = [
  v => !!v || 'Le mot de passe est requis',
  v => v.length >= 8 || 'Le mot de passe doit faire au moins 8 caractères',
  v => /\d/.test(v) || 'Le mot de passe doit contenir au moins un chiffre',
  v => /[a-z]/.test(v) || 'Le mot de passe doit contenir au moins une minuscule',
  v => /[A-Z]/.test(v) || 'Le mot de passe doit contenir au moins une majuscule'
];

const confirmPasswordRules = computed(() => [
  v => !!v || 'La confirmation du mot de passe est requise',
  v => v === password.value || 'Les mots de passe ne correspondent pas'
]);

const onFormValidityChange = (value) => {
  formIsValid.value = value;
};

const onSubmit = async () => {
  if (!form.value?.validate()) return;

  loading.value = true;

  try {
    if (isLogin.value) {
      await login();
    } else {
      await register();
    }
  } catch (error) {
    showError(error.message);
  } finally {
    loading.value = false;
  }
};

const login = async () => {
  await new Promise(resolve => setTimeout(resolve, 1000));
  showSuccess('Connexion réussie !');
};

const register = async () => {
  await new Promise(resolve => setTimeout(resolve, 1000));
  showSuccess('Inscription réussie !');
};

const toggleForm = () => {
  isLogin.value = !isLogin.value;
  form.value?.reset();
  formIsValid.value = false;
};

const togglePassword = () => {
  showPassword.value = !showPassword.value;
};

const toggleConfirmPassword = () => {
  showConfirmPassword.value = !showConfirmPassword.value;
};

const showSuccess = (message) => {
  snackbarColor.value = 'success';
  snackbarText.value = message;
  snackbar.value = true;
};

const showError = (message) => {
  snackbarColor.value = 'error';
  snackbarText.value = message;
  snackbar.value = true;
};
</script>

<style>
.v-label.v-field-label--floating {
  background-color: rgb(var(--v-theme-secondaryContainer)) !important;
  border-radius: 4px !important;
  padding: 0 4px !important;
}

.v-field__outline {
  color: rgb(var(--v-theme-secondary)) !important;
}

.v-card--variant-elevated {
 box-shadow: none !important;
}
</style>