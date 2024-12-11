<template>
  <div v-if="!isLoginPage" class="flex flex-row justify-between items-stretch mb:h-24 uus:h-16 shrink-0 w-full px-5 header-shadow border-b-2 z-10 bg-color-light">
    <div class="h-full grid content-center md:ml-16">
      <div @click="$router.push({name: 'MainPage'})"><img class="w-8 mb:w-12 h-8 mb:h-12" src="../assets/icons/icon-sigma.svg"/></div>
    </div>
    <!--  это для норм размера и вкладки для студентоффф-->
    <div v-if="isMainPage" class="flex flex-row justify-center items-center gap-4 mt-2">
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" >Формулы</button>
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'ArticlesPage'})">Статьи</button>
    </div>

    <div v-if="isArticlesPage" class="flex flex-row justify-center items-center gap-4 mt-2">
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'MainPage'})">Формулы</button>
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300">Статьи</button>
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
    isLoginPage(){
      return this.$route.fullPath === '/login';
    },
    isMainPage(){
      return this.$route.fullPath === '/';
    },
    isArticlesPage(){
      return this.$route.fullPath === '/articles';
    },

  },

};
</script>