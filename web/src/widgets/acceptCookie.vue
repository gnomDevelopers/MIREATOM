<template>
  <div v-if="showAcceptCookie" class="absolute bottom-5 left-5 flex flex-col justify-center items-center p-4 rounded-xl bg-color-light header-shadow">
    <div class="flex flex-col justify-between items-stretch w-80 gap-y-2">
      <div class="flex flex-row gap-x-2 justify-center items-center">
        <img class="w-7 h-7" src="../assets/icons/icon-cookie.svg"/>
        <p class="text-2xl">Использование Cookie</p>
      </div>
      <p class="w-full text-justify text-lg">
        Этот веб-сайт использует файлы cookie, чтобы обеспечить вам наиболее приятный опыт использования веб-сайта
      </p>
      <div class="flex flex-row justify-evenly">
        <div @click="setCookiePolicy(true)" class="py-2 px-4 text-xl rounded-xl border-2 border-solid cursor-pointer font-semibold select-none text-white bg-sky-700 border-sky-700 hover:bg-sky-600 hover:border-sky-600 active:bg-sky-500 active:border-sky-500">Принять</div>
        <div @click="setCookiePolicy(false)" class="py-2 px-4 text-xl rounded-xl border-2 border-solid cursor-pointer font-semibold select-none text-sky-700 bg-white border-sky-700 hover:bg-slate-50 hover:border-sky-600 active:bg-slate-100 active:border-sky-500">Отклонить</div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { CHECK_COOKIE_EXIST, StatusCodes } from '@/helpers/constants';

export default {
  data(){
    return {
      showAcceptCookie: false,
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore),
  },
  mounted(){
    if(!CHECK_COOKIE_EXIST('acceptCookie')){
      this.showAcceptCookie = true;
      return;
    }
    // this.showAcceptCookie = true;
  },
  methods: {
    setCookiePolicy(acceptCookie: boolean){
      document.cookie = `acceptCookie=${acceptCookie}; max-age=${60 * 60 * 24 * 365};`;

      this.showAcceptCookie = false;

      if(acceptCookie) this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Cookie файлы включены!');
      else {
        //если есть установленные ранее куки, удаляем
        if(CHECK_COOKIE_EXIST('access_token')){
          document.cookie = `access_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC;`;
        }
        if(CHECK_COOKIE_EXIST('refresh_token')){
          document.cookie = `refresh_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC;`;
        }
        
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Cookie файлы отключены!');
      }
    }
  }
};
</script>
