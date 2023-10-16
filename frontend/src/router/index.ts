import { createRouter, createWebHashHistory } from 'vue-router'

/**
 * 基础路由
 * @type { *[] }
 */
const router: any[] = [
  {
    path: '/',
    name: 'role',
    component: () => import('../views/Role.vue')
  }
]

const Router = createRouter({
  history: createWebHashHistory(),
  routes: router,
})

export default Router
