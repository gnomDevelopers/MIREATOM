<template>
  <div v-for="commit of getFormulaCommits" :key="commit.id">

    <article class="flex flex-col justify-center p-3 rounded-lg m-2 bg-gray-100">
      <div class="flex flex-col gap-y-2 bg-gray-200 w-full rounded-lg p-2">

        <div class="flex flex-row items-center px-2 py-1 min-h-10 bg-white rounded-lg">
          <span class="text-lg" v-html="convertLatexToHTML(commit.value, commit.newValue)"></span>
        </div>

        <div class="flex flex-row items-center min-h-10 bg-white rounded-lg">
          <div class="px-2 py-1 flex-grow">
            <p class="text-lg">{{ commit.newValue }}</p>
          </div>

          <div @click="copyFormula(commit.newValue)" class="btn rounded-r-lg cursor-pointer p-1">
            <svg class="w-8 h-8" viewBox="0 0 40 40" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M33.3333 21.8749L33.3332 7.99979C33.3332 6.34292 31.9901 4.99977 30.3332 4.99979L16.4582 4.99995M23.3333 34.9999L12.0833 35C10.0122 35 8.33325 33.321 8.33325 31.25L8.33325 14.9999C8.33325 12.9289 10.0122 11.2499 12.0832 11.2499L23.3332 11.2499C25.4043 11.2499 27.0832 12.9289 27.0832 14.9999L27.0833 31.2499C27.0833 33.321 25.4043 34.9999 23.3333 34.9999Z" stroke="white" stroke-width="3" stroke-linecap="round"/>
            </svg>
          </div>
        </div>

        <h3 class="text-base font-bold">Коммит: <span class="text-gray-600">{{ commit.hash }}</span></h3>
      </div>
    </article>

  </div>
</template>
<script lang="ts">
import katex from 'katex';
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useFormulsStore } from '@/stores/formulsStore';
import { StatusCodes } from '@/helpers/constants';

export default {
  computed: {
    ...mapStores(useStatusWindowStore, useFormulsStore),

    getFormulaCommits(){
      return this.formulsStore.selectedFormulaCommits;
    },
  },
  methods: {
    convertLatexToHTML(value: string, secondFormula: string){
      try{
        const render = katex.renderToString(value, {
          throwOnError: true,
          displayMode: false,
          output: 'mathml',
          trust: false,
        });

        return render;

      }catch(error){
        return katex.renderToString(secondFormula, {
          throwOnError: true,
          displayMode: false,
          output: 'mathml',
          trust: false,
        });
      }
    },
    copyFormula(formula: string){
      navigator.clipboard.writeText(formula).then(() => {
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Формула скопирована!', 1500);
      })
      .catch((err) => {
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при копировании формулы!');
        console.error('Could not copy text: ', err);
      });
    }
  },
  watch: {
    'formulsStore.selectedFormulaID': {
      handler(val){
        this.formulsStore.loadFormulaCommits();
      },
      immediate: true,
    }
  }
};
</script>