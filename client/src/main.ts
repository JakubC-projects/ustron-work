import { createApp } from 'vue'
import "./main.css"
import App from './App.vue'
import HomePage from "./pages/home.vue"
import {createRouter, createWebHashHistory} from 'vue-router'
import ErrorPage from './pages/error.vue'
import RegisterPage from './pages/register.vue'
import AdminPage from './pages/admin.vue'
import { createPinia } from 'pinia'
import 'vue-toast-notification/dist/theme-bootstrap.css';
import ToastPlugin from 'vue-toast-notification'



// 2. Define some routes
// Each route should map to a component.
// We'll talk about nested routes later.
const routes = [
  { path: '/', component: HomePage },
  { path: '/error',component: ErrorPage },
  { path: '/register', component: RegisterPage },
  { path: '/admin', component: AdminPage },
]

// 3. Create the router instance and pass the `routes` option
// You can pass in additional options here, but let's
// keep it simple for now.
const router = createRouter({
  // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
  history: createWebHashHistory(),
  routes, // short for `routes: routes`
})

const pinia = createPinia()

const app = createApp(App)
app.use(router);
app.use(pinia);
app.use(ToastPlugin);

app.mount('#app')
