<template>
  <div class="flex h-screen">
    <div class="flex flex-col md:w-1/2 uus:w-full h-full items-start justify-start p-4">
      <img @click="$router.push('/')" class="w-8 mb:w-12 h-8 mb:h-12 mb:m-4" src="../assets/icons/icon-sigma.svg"/>
      <div class="flex flex-col w-full h-full items-center justify-center">
        <div class="flex flex-col w-10/12 h-full items-center m-10">

          <!-- Переключатель вход/регистрация -->
          <section class="toggle-container">
            <div class="toggle">
              <button
                :class="{ active: entranceType === 'login' }"
                @click="entranceType = 'login'"
                class="toggle-btn">

                Вход
              </button>

              <button
                :class="{ active: entranceType === 'signup' }"
                @click="entranceType = 'signup'"
                class="toggle-btn">

                Регистрация
              </button>
            </div>
          </section>

          <!-- Форма для входа -->
          <section v-if="entranceType === 'login'" class="w-full h-1/2">
            <p class="auth-description">Войдите с помощью своей учётной записи</p>

            <div class="mb-4 h-full">
              <loginInput type="text" text="Электронная почта:" placeholder="email@example.com" @input-change="checkLogin"/><br>
              <loginInput type="password" text="Ваш пароль:" placeholder="your_password" @input-change="checkPassword"/>
              <div class="w-full flex justify-center mt-16">
                <submitButton value="Войти" class="btn w-5/6" @click="sendLogin"/>
              </div>
            </div>
          </section>

          <!--Форма для регистрации-->
          <section v-if="entranceType === 'signup'" class="w-full h-1/2">
            <div class="flex flex-row items-stretch justify-center gap-2">
              <p class="auth-description">Зарегистрируйтесь, если впервые в Sigma</p>
            </div>
            <div class="mb-4 h-full">
              <signupInput type="name"  text="ФИО:" @input-change="checkRegFIO"/><br> <!--@input-change="checkLogin"-->
              <signupInput type="login"  text="Электронная почта:" @input-change="checkRegLogin"/><br>
              <signupInput type="password" text="Ваш пароль:" @input-change="checkRegPassword"/><br>
              <signupInput type="repPassword" text="Повторите пароль:" @input-change="checkRegRepeatPassword"/> <!--ПРОВЕРИТЬ СОВПАДАЮТ ЛИ-->
              <div class="w-full flex justify-center mt-16">
                <submitButton value="Зарегистрироваться" class="btn w-5/6" @click="sendReg"/>
              </div>
            </div>
          </section>

        </div>
      </div>
    </div>
    <div class="login-bg w-1/2 hidden md:block"></div>
  </div>
</template>

<script lang="ts">

import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { ValidUserLogin, ValidUserPassword, ValidUserName} from '../helpers/validator';
import { type IValidAnswer, StatusCodes, type IAPI_Login, type IAPI_Register } from '../helpers/constants';
import {API_Login, API_Register} from '@/api/api';

import loginInput from '../shared/loginInput.vue';
import signupInput from "@/shared/signupInput.vue";
import submitButton from '../shared/submitButton.vue';

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

      regLogin: {value: '', error: ''} as IValidAnswer,
      regFullname: {value: '', error: ''} as IValidAnswer,
      regPassword: {value: '', error: ''} as IValidAnswer,
      regRepeatPassword: {value: '', error: ''} as IValidAnswer,

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
        const data:IAPI_Login = { email: this.login.value, password: this.password.value };

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
    sendReg(){
      if(this.regLogin.value !== '' && this.regPassword.value !== '' && this.regFullname.value !== '' && this.regRepeatPassword.value !== ''){
        const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);
        
        const data:IAPI_Register = { 
          email: this.regLogin.value, 
          password: this.regPassword.value, 
          fullname: this.regFullname.value
        };

        API_Register(data)
        .then(async (response:any) => {
          await this.userInfoStore.onAuthorized(response);

          this.statusWindowStore.deteleStatusWindow(stID);
          this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Регистрация успешна!');
          
          this.$router.push('/');
          //ЛОГИКА ВХОДА
        })
        .catch(error => {
          this.statusWindowStore.deteleStatusWindow(stID);

          if(error.status === 500 || error.status === 400) this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Неверные данные!');
          else this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при авторизации!');
        });
        return;
      }

      if(this.regLogin.value === ''){
        if(this.regLogin.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите логин!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.regLogin.error);
      }
      if(this.regPassword.value === ''){
        if(this.regPassword.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите пароль!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.regPassword.error);
      }
      if(this.regFullname.value === ''){
        if(this.regFullname.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите ФИО!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.regFullname.error);
      }
      if(this.regRepeatPassword.value === ''){
        if(this.regRepeatPassword.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите повторный пароль!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.regRepeatPassword.error);
      }
    },

    // проверки для входа
    checkLogin(value: string){
      this.login = ValidUserLogin(value);
      if(value === '') this.login.error = '';
    },
    checkPassword(value: string){
      this.password = ValidUserPassword(value);
      if(value === '') this.password.error = '';
    },

    // проверки для регистрации
    checkRegLogin(value: string){
      this.regLogin = ValidUserLogin(value);
      if(value === '') this.regLogin.error = '';
    },
    checkRegPassword(value: string){
      this.regPassword = ValidUserPassword(value);
      if(value === '') this.regPassword.error = '';
    },
    checkRegRepeatPassword(value: string){
      // this.regRepeatPassword = ValidUserPassword(value);
      // if(value === '') this.regRepeatPassword.error = '';
      if(value !== this.regPassword.value) {
        this.regRepeatPassword = { value: '', error: 'Повторный пароль не совпадает!' };
        return;
      }
      this.regRepeatPassword = { value: value, error: '' };
    },
    checkRegFIO(value: string){
      this.regFullname = ValidUserName(value);
      if(value === '') this.regFullname.error = '';
    },
  },
};
</script>

