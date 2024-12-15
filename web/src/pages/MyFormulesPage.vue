<template>
  <!-- модальное окно для редактирования формулы -->
  <Transition>
    <UpdateMyFormula v-if="showUpdateFormulaMW" :id="updateFormulaID" :value="updateFormulaValue" :title="updateFormulaName" @close-window="hideEditFormula"/>
  </Transition>

  <div class="flex flex-row h-full scrollable">
    <div class="flex flex-col w-1/2 border-r-2 border-gray-200">
      <div class="flex justify-center m-6">
        <p class="text-2xl">Сохранённые формулы</p>
      </div>
      
      <div v-if="getFormulsList.length !== 0" class="flex flex-col h-full mb-4 mx-6 scrollable" id="myFormulaScrollWrapper">
        <div class="flex flex-col justify-center gap-4">
          <div>
            <div v-for="formulaItem in getFormulsList" >
              <MyFormulaItem
                :formula="formulaItem.value"
                :title="formulaItem.title"
                :id="formulaItem.id"
                :class="{'border-2 border-gray-400': formulsStore.selectedFormulaID === formulaItem.id}"
                @click="selectFormula(formulaItem.id)"
                @edit-formula="showEditFormula(formulaItem.value, formulaItem.title, formulaItem.id)"
              />
            </div>
          </div>
        </div>
      </div>

      <div v-else class="w-full flex flex-col items-center cursor-default">
        <div class="px-4 py-2 rounded-lg bg-gray-300">
          <p class=text-lg>У Вас пока что нет сохраненных формул</p>
        </div>
      </div>
    </div>

    <div class="flex flex-col w-1/2">
      <div class="flex justify-center m-6">
        <p class="text-2xl">История изменений формулы</p>
      </div>

      <div v-if="formulsStore.selectedFormulaID !== null" class="flex flex-col h-screen mb-4 mx-6 scrollable">
        <div class="flex flex-col">
          <CommitFormulaHistory />
        </div>
      </div>
      <div v-else class="flex flex-col items-center">
        <div class="px-4 py-2 rounded-lg bg-gray-300">
          <p class="text-lg">Выберите формулу, чтобы посмотреть историю её изменений</p>
        </div>
      </div>
    </div>
  </div>
</template>
  
<script lang="ts">
import { mapStores } from 'pinia';
import { useBlurStore } from '@/stores/blurStore';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { useFormulsStore } from '@/stores/formulsStore';
import { StatusCodes, type TMaybeNumber, type TMaybeString } from '@/helpers/constants';
import { API_Get_Formuls_History } from '@/api/api';

import MyFormulaItem from '@/shared/myFormulaItem.vue';
import UpdateMyFormula from '@/entities/updateMyFormula.vue';
import CommitFormulaHistory from '@/entities/commitFormulaHistory.vue';

export default {
  components: {
    MyFormulaItem,
    UpdateMyFormula,
    CommitFormulaHistory,
  },
  data() {
    return {
      formulaItems: [] as {id: number, title: string, value: string, user_id: number}[],

      showUpdateFormulaMW: false,
      updateFormulaID: -1,
      updateFormulaName: '',
      updateFormulaValue: '',

      formulasPage: 1,

      selectedFormulaCommits: [],
    };
  },
  computed: {
    ...mapStores(useBlurStore, useStatusWindowStore, useUserInfoStore, useFormulsStore),

    getFormulsList(){
      return this.formulsStore.formulsList;
    }
  },
  mounted(){
    if(this.userInfoStore.userID === null) return;
    //получение первой страницы истории
    API_Get_Formuls_History(this.userInfoStore.userID, this.formulasPage)
    .then((response:any) => {
      this.formulsStore.formulsList = response.data;

      //если получено меньше 20 элементов - значит больше формул нет
      if(response.data.length < 20) return;

      //иначе находим элемент-обертку списка
      const myFormulasHTML = document.getElementById('myFormulaScrollWrapper');
      //если не нашли - выходим
      if(myFormulasHTML === null) return;
      //добавляем слушатель события скролл
      myFormulasHTML.addEventListener('scroll', this.handleMyFromulaScroll);
    })
    .catch(error => {
      this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при получении формул!');
    })
  },
  methods: {
    selectFormula(formulaID: number) {
      this.formulsStore.selectedFormulaID = formulaID;
    },
    showEditFormula(formula: string, title: string, formulaID: number) {
      this.updateFormulaValue = formula;  
      this.updateFormulaName = title; 
      this.updateFormulaID = formulaID;

      this.showUpdateFormulaMW = true; 
      this.blurStore.showBlur = true;
    },
    hideEditFormula() {
      this.showUpdateFormulaMW = false; 
      this.blurStore.showBlur = false;
    },
    handleMyFromulaScroll(event: any){
      const scrollHeight = event.target.scrollHeight;
      const scrollTop = event.target.scrollTop;
      const clientHeight = event.target.clientHeight;

      if (scrollTop + clientHeight >= scrollHeight) {    

        if(this.userInfoStore.userID === null) return;

        this.formulasPage++;
        API_Get_Formuls_History(this.userInfoStore.userID, this.formulasPage)
        .then((response:any) => {
          //сохраняем полученные формулы в массив
          for(const item of response.data){
            this.formulsStore.formulsList.push(item);
          }
          //если получено меньше 20 элементов - значит больше формул нет
          if(response.data.length < 20) event.target.removeEventListener('scroll', this.handleMyFromulaScroll);
        });
      }
    },
  }
};
</script>