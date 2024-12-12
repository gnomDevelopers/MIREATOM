import { defineStore } from "pinia";
import type { TMaybeBoolean } from "@/helpers/constants";
import { API_Authenticate } from "@/api/api";

export const useUserInfoStore = defineStore('userInfo', {
  state() {
    return{
      authorized: null as TMaybeBoolean, // проверка авторизованности
      
    }
  },
  actions: {
    async Authenticate(){
      try{
        const response = await API_Authenticate();
        await this.onAuthorized(response);

      }catch (error){
        this.authorized = false;
      }
    },
    async onAuthorized(response: any){
      document.cookie = `access_token=${response.data.access_token}; max-age=${60 * 60 * 2}; secure; samesite=strict`;
      document.cookie = `refresh_token=${response.data.refresh_token}; max-age=${60 * 60 * 6}; secure; samesite=strict`;
      
      // устанавливаем авторизацию, чтобы роутер перекинул на страницы 
      this.authorized = true;
    },
  }
});