<template>
  <v-container fluid fill-height class="align-center justify-center">
    <v-row justify="center">
      <v-col cols="12" sm="8" md="6" lg="4">
        <v-card class="elevation-12">
          <v-toolbar color="primary">
            <v-toolbar-title>{{ isLogin ? 'Connexion' : 'Inscription' }}</v-toolbar-title>
          </v-toolbar>
          
          <v-card-text>
            <v-form ref="form" v-model="valid" @submit.prevent="onSubmit">
              <v-container>
                <v-text-field
                  v-model="username"
                  :rules="usernameRules"
                  label="Nom d'utilisateur"
                  prepend-icon="mdi-account"
                  required
                ></v-text-field>

                <v-text-field
                  v-model="password"
                  :rules="passwordRules"
                  label="Mot de passe"
                  prepend-icon="mdi-lock"
                  :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                  @click:append="showPassword = !showPassword"
                  :type="showPassword ? 'text' : 'password'"
                  required
                ></v-text-field>

                <!-- Confirmation mot de passe uniquement pour l'inscription -->
                <v-text-field
                  v-if="!isLogin"
                  v-model="confirmPassword"
                  :rules="confirmPasswordRules"
                  label="Confirmer le mot de passe"
                  prepend-icon="mdi-lock"
                  :append-icon="showConfirmPassword ? 'mdi-eye' : 'mdi-eye-off'"
                  @click:append="showConfirmPassword = !showConfirmPassword"
                  :type="showConfirmPassword ? 'text' : 'password'"
                  required
                ></v-text-field>
              </v-container>
            </v-form>
          </v-card-text>

          <v-card-actions class="px-4 pb-4">
            <v-btn
              block
              color="primary"
              :loading="loading"
              :disabled="!valid"
              @click="onSubmit"
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

    <!-- Snackbar pour les messages -->
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

<script>
export default {
  name: 'LoginRegister',
  data: () => ({
    valid: false,
    loading: false,
    isLogin: true,
    showPassword: false,
    showConfirmPassword: false,
    
    // Form fields
    username: '',
    password: '',
    confirmPassword: '',

    // Snackbar
    snackbar: false,
    snackbarText: '',
    snackbarColor: 'success',

    // Validation rules
    usernameRules: [
      v => !!v || 'Ce champ est requis',
      v => v.length <= 50 || 'Le nom doit faire moins de 50 caractères'
    ],
    passwordRules: [
      v => !!v || 'Le mot de passe est requis',
      v => v.length >= 8 || 'Le mot de passe doit faire au moins 8 caractères',
      v => /\d/.test(v) || 'Le mot de passe doit contenir au moins un chiffre',
      v => /[a-z]/.test(v) || 'Le mot de passe doit contenir au moins une minuscule',
      v => /[A-Z]/.test(v) || 'Le mot de passe doit contenir au moins une majuscule'
    ],
    confirmPasswordRules: [
      v => !!v || 'La confirmation du mot de passe est requise',
      v => v === this.password || 'Les mots de passe ne correspondent pas'
    ]
  }),

  methods: {
    async onSubmit() {
      if (!this.$refs.form.validate()) return;

      this.loading = true;

      try {
        if (this.isLogin) {
          // Logique de connexion
          await this.login();
        } else {
          // Logique d'inscription
          await this.register();
        }
      } catch (error) {
        this.showError(error.message);
      } finally {
        this.loading = false;
      }
    },

    async login() {
      // Simulation d'une requête API
      await new Promise(resolve => setTimeout(resolve, 1000));
      this.showSuccess('Connexion réussie !');
      // Redirection ou autre logique après connexion
    },

    async register() {
      // Simulation d'une requête API
      await new Promise(resolve => setTimeout(resolve, 1000));
      this.showSuccess('Inscription réussie !');
      // Redirection ou autre logique après inscription
    },

    toggleForm() {
      this.isLogin = !this.isLogin;
      this.$refs.form.reset();
    },

    showSuccess(message) {
      this.snackbarColor = 'success';
      this.snackbarText = message;
      this.snackbar = true;
    },

    showError(message) {
      this.snackbarColor = 'error';
      this.snackbarText = message;
      this.snackbar = true;
    }
  }
}
</script>