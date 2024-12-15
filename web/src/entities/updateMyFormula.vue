<template>
  <section class=" fixed inset-0 flex justify-center items-center z-30 p-4">
    <div class="flex flex-col gap-y-4 items-center z-30 p-4 rounded-xl min-w-[500px] bg-white relative">
      <div class="absolute right-0 top-0 cursor-pointer" @click="$emit('closeWindow')">
        <svg class="w-10 h-10" viewBox="0 0 52 52" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M34.6668 17.3334L17.3335 34.6667M34.6668 34.6667L17.3335 17.3334" stroke="#8F0101" stroke-width="2" stroke-linecap="round"/>
        </svg>
      </div>

      <h1 class="text-2xl text-color-theme cursor-default">Изменить формулу</h1>

      <article class="flex flex-col items-center gap-y-4 rounded-lg max-h-[500px] w-full py-2 px-4 bg-gray-100">
        <span v-html="formulaHTML" class="text-xl my-4"></span>

        <div class="flex flex-col gap-y-2 w-full py-2 px-4 rounded-lg bg-gray-200">
          <label class="text-lg cursor-pointer" for="formulaName">Введите название формулы</label>
          <input 
          type="text" 
          id="formulaName"
          placeholder="Название формулы" 
          v-model="formulaTitle"
          class="text-lg px-2 py-1 outline-none rounded border border-solid border-gray-300 focus:border-sky-500"/>
        </div>

        <div class="flex flex-col gap-y-2 w-full py-2 px-4 rounded-lg bg-gray-200">
          <label class="text-lg cursor-pointer" for="formulaName">Измените формулу</label>        
          <input 
          type="text" 
          id="formulaText"
          placeholder="Содержание формулы" 
          v-model=formula
          class="text-lg px-2 py-1 outline-none rounded border border-solid border-gray-300 focus:border-sky-500"/>
        </div>

        <div @click="saveFormula" class="flex flex-row gap-x-2 items-center btn cursor-pointer rounded-lg px-2 py-1 my-2">
          <svg class="w-9 h-9" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M34 42V26H14V42M14 6V16H30M38 42H10C8.93913 42 7.92172 41.5786 7.17157 40.8284C6.42143 40.0783 6 39.0609 6 38V10C6 8.93913 6.42143 7.92172 7.17157 7.17157C7.92172 6.42143 8.93913 6 10 6H32L42 16V38C42 39.0609 41.5786 40.0783 40.8284 40.8284C40.0783 41.5786 39.0609 42 38 42Z" stroke="white" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
          <p class="text-xl text-white">Сохранить изменения</p>
        </div>
      </article>
    </div>
    
  </section>
</template>
<script lang="ts">
import katex from 'katex';
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useFormulsStore } from '@/stores/formulsStore';
import { StatusCodes } from '@/helpers/constants';
import { API_Update_Formula } from '@/api/api';

export default {
  emits: ['closeWindow'],
  props: {
    id: {
      type: Number,
      required: true,
    },
    value: {
      type: String,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
  },
  data(){
    return {
      formula: this.value,
      formulaHTML: '',
      formulaTitle: this.title,
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore, useFormulsStore),
  },
  mounted() {
    this.formulaHTML = katex.renderToString(this.formula, {
      throwOnError: true,
      displayMode: false,
      output: 'mathml',
      trust: false,
    });
  },
  methods: {
    saveFormula(){
      //проверка на пустое название формулы
      if(this.formulaTitle === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите название формулы!');
        return;
      }
      //проверка на пустую формулу
      if(this.formula === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите формулу!');
        return;
      }
      //проверка на отсутствие изменений
      if(this.formula === this.value && this.formulaTitle === this.title){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Внесите изменения в формулу!');
        return;
      }

      //создаем объект для отправки
      const data = {
        id: this.id,
        title: this.formulaTitle,
        value: this.formula,
      }
      //выводим окно отправки
      const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Обновляем формулу...', -1);
      //запрос на изменение
      API_Update_Formula(data)
      .then((response: any) => {
        //если все ок - выводим сообщение что все ок
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Формула обновлена!');
        //обновляем в хранилище формулу с текущим id
        for(const item of this.formulsStore.formulsList){
          if(item.id === this.id) {
            item.title = this.formulaTitle;
            item.value = this.formula;
            break;
          }
        }
        //заркываем окно обновления формулы
        this.$emit('closeWindow');
      })
      .catch(error => {
        //если что-то не так - сообщаем об ошибке
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при обновлении формулы!');
      })
    }
  }
};
</script>