import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'

// Importuj PocketBase
import PocketBase from 'pocketbase'

// Inicializuj PocketBase
const pb = new PocketBase('http://localhost:8090')

// Môžeš použiť `pb` kdekoľvek vo svojej aplikácii
// Napríklad: pb.collection('users').getList()

const app = createApp(App)

app.use(createPinia())

// Pridaj PocketBase do globálnych vlastností aplikácie
app.config.globalProperties.$pb = pb

app.mount('#app')
