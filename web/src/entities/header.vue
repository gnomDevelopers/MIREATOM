<template>
  <div 
    v-if="$route.fullPath !== '/login'" 
    class="flex flex-row gap-x-10 justify-between items-stretch mb:h-16 h-12 shrink-0 w-full px-5 header-shadow border-b-2 z-10 bg-color-light">

    <div class="h-full grid content-center md:ml-16 cursor-pointer">
      <div @click="$router.push({name: 'MainPage'})"><img class="w-8 mb:w-12 h-8 mb:h-12" src="../assets/icons/icon-sigma.svg"/></div>
    </div>

    <div class="flex justify-center items-center gap-x-6 ">
      <div
        @click="$router.push({name: 'MyFormulesPage'})"
        class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl text-color border-2 rounded-lg cursor-pointer transition-bg hover:bg-red-50 active:bg-red-100"
        :class="{'border-hr-color': $route.path === '/myformules'}">
        
        <a class="text-xl">Мои формулы</a>
      </div>
      <div
        @click="$router.push({name: 'MainPage'})"
        class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl text-color border-2 rounded-lg cursor-pointer transition-bg hover:bg-red-50 active:bg-red-100"
        :class="{'border-hr-color': $route.path === '/'}">
        
        <a class="text-xl">Формулы</a>
      </div>
      <div
        @click="$router.push({name: 'ArticlesPage'})"
        class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl text-color border-2 rounded-lg cursor-pointer transition-bg hover:bg-red-50 active:bg-red-100"
        :class="{'border-hr-color': $route.path === '/articles'}">
        <!--box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300-->
        <a class="text-xl">Статьи</a>
      </div>
    </div>

    <div class="flex flex-row gap-x-2 justify-center items-center md:mr-16">

      <img class="w-8 mb:w-10 h-8 mb:h-10" src="../assets/icons/icon-profile.svg"/>

      <div class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'LoginPage'})">
        Вход
      </div>
    </div>
  </div>




</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useUserInfoStore } from '@/stores/userInfoStore';

export default {

  computed: {
    ...mapStores(useUserInfoStore),
  },

};
</script>