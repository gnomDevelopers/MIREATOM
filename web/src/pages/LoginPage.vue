<template>
  <div class="flex h-screen">
    <div class="flex flex-col md:w-1/2 uus:w-full h-full items-start justify-start p-4">
      <img class="w-8 mb:w-12 h-8 mb:h-12 mb: m-4" src="../assets/icons/icon-sigma.svg"/>
      <div class="flex flex-col w-full h-full items-center justify-center">
        <div class="flex flex-col w-10/12 h-full items-center m-10">
          <!-- Переключатель вход/регистрация -->
          <div class="toggle-container">
            <div class="toggle">
              <button
                  :class="{ active: entranceType === 'login' }"
                  @click="entranceType = 'login'"
                  class="toggle-btn"
              >
                Вход
              </button>
              <button
                  :class="{ active: entranceType === 'signup' }"
                  @click="entranceType = 'signup'"
                  class="toggle-btn"
              >
                Регистрация
              </button>
            </div>
          </div>

          <!-- Форма для входа -->
          <div v-if="entranceType === 'login'" class="w-full h-1/2">
            <p class="auth-description">Войдите с помощью своей учётной записи</p>

            <div class="mb-4 h-full">

              <loginInput type="text"  text="Электронная почта:" @input-change="checkLogin"/><br>
              <loginInput type="password" text="Ваш пароль:" @input-change="checkPassword"/>
              <div class="w-full flex justify-center mt-16">
                <submitButton value="Войти" class="btn w-5/6" @click="sendLogin"/>
              </div>


            </div>


          </div>

          <!--Форма для регистрации-->
          <div v-if="entranceType === 'signup'" class="w-full h-1/2">
            <div class="flex flex-row items-stretch justify-center gap-2">
              <p class="auth-description">Зарегистрируйтесь, если впервые в Sigma</p>
            </div>
<!--ПРОПИСАТЬ ПРОВЕРКУ ДЛЯ РЕГИСТРАЦИЙ-->

            <div class="mb-4 h-full">
              <signupInput type="name"  text="ФИО:" /><br> <!--@input-change="checkLogin"-->
              <signupInput type="login"  text="Электронная почта:" /><br>
              <signupInput type="password" text="Ваш пароль:" /><br>
              <signupInput type="repPassword" text="Повторите пароль:" /> <!--ПРОВЕРИТЬ СОВПАДАЮТ ЛИ-->
              <div class="w-full flex justify-center mt-16">
                <submitButton value="Зарегистрироваться" class="btn w-5/6" @click="sendReg"/>
              </div>
            </div>

          </div>
        </div>
      </div>
    </div>
    <div class="login-bg w-1/2 uus:hidden md:block"></div>
  </div>
</template>

<script lang="ts">

import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useUserInfoStore } from '@/stores/userInfoStore';
import loginInput from '../shared/loginInput.vue';
import signupInput from "@/shared/signupInput.vue";
import submitButton from '../shared/submitButton.vue';
import { ValidUserLogin, ValidUserPassword } from '../helpers/validator';
import { type IValidAnswer, StatusCodes, type IAPI_Login } from '../helpers/constants';
import {API_Login} from '@/api/api';
import { defineStore } from 'pinia';


export default {

  components:{
    signupInput,
    loginInput,
    submitButton,
  },

  data(){
    return{
      entranceType: 'login',
      login: {value: '', error: ''} as IValidAnswer,
      password: {value: '', error: ''} as IValidAnswer,
      name: {value: '', error: ''} as IValidAnswer,
      repPassword: {value: '', error: ''} as IValidAnswer,

      showPassword: false, // Для отображения пароля
    }
  },
  computed:{
    ...mapStores(useStatusWindowStore, useUserInfoStore),

  },

  methods: {
    sendLogin(){
      if(this.login.value !== '' && this.password.value !== ''){
        const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);

        const data:IAPI_Login = { login: this.login.value, password: this.password.value };

        API_Login(data)
            .then(async (response:any) => {
              await this.userInfoStore.onAuthorized(response);

              this.statusWindowStore.deteleStatusWindow(stID);
              this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Авторизация успешна!');
              this.$router.push('/');
              //ЛОГИКА ВХОДА
            })
            .catch(error => {
              this.statusWindowStore.deteleStatusWindow(stID);

              if(error.status === 500 || error.status === 400) this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Неверный логин или пароль!');
              else this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при авторизации!');
            });
        return;
      }

      if(this.login.value === ''){
        if(this.login.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите логин!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.login.error);
      }
      if(this.password.value === ''){
        if(this.password.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите пароль!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.password.error);
      }
    },


    checkLogin(value: string){
      this.login = ValidUserLogin(value);
      if(value === '') this.login.error = '';
    },
    checkPassword(value: string){
      this.password = ValidUserPassword(value);
      if(value === '') this.password.error = '';
    },

    sendReg(){
      if(this.login.value !== '' && this.password.value !== '' && this.name.value !== '' && this.repPassword.value !== ''){
        const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);

        const data:IAPI_Login = { login: this.login.value, password: this.password.value };

        API_Login(data)
            .then(async (response:any) => {
              await this.userInfoStore.onAuthorized(response);

              this.statusWindowStore.deteleStatusWindow(stID);
              this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Авторизация успешна!');
              this.$router.push('/');
              //ЛОГИКА ВХОДА
            })
            .catch(error => {
              this.statusWindowStore.deteleStatusWindow(stID);

              if(error.status === 500 || error.status === 400) this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Неверный логин или пароль!');
              else this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при авторизации!');
            });
        return;
      }

      if(this.login.value === ''){
        if(this.login.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите логин!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.login.error);
      }
      if(this.password.value === ''){
        if(this.password.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите пароль!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.password.error);
      }
    },


  },
};
</script>

