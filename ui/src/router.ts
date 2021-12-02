import { createRouter, createWebHistory } from 'vue-router'

import routes from 'virtual:generated-pages'

for (let i = 0; i < routes.length; i++) {
  routes[i].path = routes[i].path
}
// console.log(routes) // uncomment to view routes generated
const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
