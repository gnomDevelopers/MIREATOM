import { createRouter, createWebHistory } from "vue-router";
import { useUserInfoStore } from "@/stores/userInfoStore";
import { watchEffect } from 'vue';

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: 'MainPage',
      path: '/',
      component: () => import('@/pages/MainPage.vue'),
      meta: { requiresAuth: false },
    },
    {
      name: 'LoginPage',
      path: '/login',
      component: () => import('@/pages/LoginPage.vue'),
      meta: { requiresAuth: false },
    },
    {
      name: 'ArticlesPage',
      path: '/articles',
      component: () => import('@/pages/ArticlesPage.vue'),
      meta: { requiresAuth: false },
    },
    {
      name: 'MyFormulesPage',
      path: '/myformules',
      component: () => import('@/pages/MyFormulesPage.vue'),
      meta: { requiresAuth: false },
    },
  ],
});

router.beforeEach(async (to, from, next) => {
  const userInfoStore = useUserInfoStore();

  if(to.meta.requiresAuth){
    if(userInfoStore.authorized === null){
      await new Promise<void>((resolve) => {
        const unwatch = watchEffect(() => {
          if (userInfoStore.authorized !== null) {
            unwatch();
            resolve();
          }
        });
      });
    }
    if(!userInfoStore.authorized) {
      next({name: 'LoginPage'});
      // next();
      return;
    }
  }

  // if(to.name === 'LoginPage' && userInfoStore.authorized){
  //   next({name: 'MainPage'});
  //   return;
  // }

  next();
});