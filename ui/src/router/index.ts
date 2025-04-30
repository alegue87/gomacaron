import MainLayout from '../layouts/MainLayout.vue'
import HomePage from '../pages/HomePage.vue'
import LoginPage from '../pages/LoginPage.vue'
import DevicePage from '../pages/DevicePage.vue'
import DashboardPage from '../pages/DashboardPage.vue'
import DetailsPage from '../pages/DetailsPage.vue'
import ParametersPage from '../pages/ParametersPage.vue'

import { createRouter, createMemoryHistory } from 'vue-router'

import { useAuthStore } from '../store'
import type { RouteRecordRaw } from 'vue-router';

// Routes registration
const routes = [
    { 
      path: '/', 
      component: MainLayout,
      children: [
        { path: '', component: HomePage},
        {
          path: '/device/:id/', 
          component: DevicePage,
          children: [
            { path: 'dashboard', component: DashboardPage },
            { path: 'details', component: DetailsPage },
            { path: 'parameters', component: ParametersPage },
          ]
        },
      ],
      meta: {
        requiresAuth: true // Add meta field to indicate protected route
      }
    },
    { path: '/login', component: LoginPage },
  ]
  
// Create Router
const router = createRouter({
  history: createMemoryHistory(),
  routes: routes as RouteRecordRaw[],
})
  
router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    //const token = localStorage.getItem('token');
    const { loggedIn } = useAuthStore()

    if (loggedIn) {
      // User is authenticated, proceed to the route
      //localStorage.setItem('isAdmin', true)
      next();
    } else {
      // User is not authenticated, redirect to login
      //localStorage.setItem('isAdmin', false)
      next('/login');
    }
  } else {
      // Non-protected route, allow access
      next();
  }
  })

export default router