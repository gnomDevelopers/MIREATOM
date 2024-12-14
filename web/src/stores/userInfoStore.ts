import { defineStore } from "pinia";
import { type TMaybeBoolean, type TMaybeNumber, type TMaybeString } from "@/helpers/constants";
import { API_Authenticate, API_Get_User_Data } from "@/api/api";

export const useUserInfoStore = defineStore('userInfo', {
  state() {
    return{
      authorized: null as TMaybeBoolean, // проверка авторизованности
      userID: null as TMaybeNumber,

      email: null as TMaybeString,
      name: null as TMaybeString,
      surname: null as TMaybeString,
      third_name: null as TMaybeString,
    }
  },
  actions: {
    async Authenticate(){
      try{
        const response = await API_Authenticate();
        this.onAuthorized(response);

      }catch (error){
        this.authorized = false;
      }
    },
    onAuthorized(response: any){
      document.cookie = `access_token=${response.data.access_token}; max-age=${60 * 60 * 2}; secure; samesite=strict`;
      document.cookie = `refresh_token=${response.data.refresh_token}; max-age=${60 * 60 * 6}; secure; samesite=strict`;
      
      // устанавливаем авторизацию, чтобы роутер перекинул на страницы 
      this.authorized = true;
      this.userID = response.data.id;
      console.log('userID: ', this.userID);
      //получаем инфу о пользователе
      this.getUserData();
    },
    async getUserData(){
      if(!this.authorized || this.userID === null) return;

      API_Get_User_Data(this.userID).then((response:any) => {
        this.email =  response.data.email;
        this.name =  response.data.name;
        this.surname =  response.data.surname;
        this.third_name =  response.data.third_name;
      })
    }
  }
});