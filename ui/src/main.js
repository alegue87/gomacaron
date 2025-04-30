import { createApp, markRaw } from 'vue'
import { Quasar } from 'quasar'
import App from './App.vue'
import ApiService from './services/api.service'
import { createPinia } from 'pinia'
import { StorageTokenService } from './services/storage.service'
import router from './router'
import { useAuthStore } from '@/store';

import './style.css'
// Import icon libraries
import '@quasar/extras/material-icons/material-icons.css'
// Import Quasar css
import 'quasar/src/css/index.sass'

// Set the base URL of the API
ApiService.init(import.meta.env.VUE_APP_MOQUI_API_ENDPOINT)

// If token exists set header
if (StorageTokenService.getToken()) {
  ApiService.setHeader()
}

const pinia = createPinia()


pinia.use(({ store }) => {
  store.$router = markRaw(router)
});

// Create app
const myApp = createApp(App) // Root Component

myApp
  .use(router)
  .use(pinia)


const quasar = myApp.use(Quasar, {
    plugins: {}, // import Quasar plugins and add here
  })
  
myApp.mount('#app') // Html root element


export default myApp
export { pinia, quasar } 
