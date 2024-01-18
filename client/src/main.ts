import { createApp } from 'vue'
import "./main.css"
import App from './App.vue'
import HomePage from "./pages/home.vue"
import {createRouter, createWebHashHistory} from 'vue-router'
import ErrorPage from './pages/error.vue'
import RegisterPage from './pages/register.vue'


// 2. Define some routes
// Each route should map to a component.
// We'll talk about nested routes later.
const routes = [
  { path: '/', component: HomePage },
  { path: '/error',component: ErrorPage },
  { path: '/register', component: RegisterPage },
]

// 3. Create the router instance and pass the `routes` option
// You can pass in additional options here, but let's
// keep it simple for now.
const router = createRouter({
  // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
  history: createWebHashHistory(),
  routes, // short for `routes: routes`
})

const app = createApp(App)
app.use(router);

app.mount('#app')
