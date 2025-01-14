/**
 * plugins/vuetify.js
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'
import materialExport from "../material-theme-ihm2.json";

// Composables
import { createVuetify } from 'vuetify'
// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
  theme: {
    defaultTheme: 'light',
    themes: {
      "light": { dark: false, colors: materialExport.schemes.light },
      "light-high-contrast": { dark: false, colors: materialExport.schemes['light-high-contrast'] },
      "light-medium-contrast": { dark: false, colors: materialExport.schemes['light-medium-contrast'] },
      "dark": { dark: true, colors: materialExport.schemes.dark },
      "dark-high-contrast": { dark: true, colors: materialExport.schemes['dark-high-contrast'] },
      "dark-medium-contrast": { dark: true, colors: materialExport.schemes['dark-medium-contrast'] },
    }
  },
})
