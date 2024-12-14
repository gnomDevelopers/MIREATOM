<template>
    <!-- модальное окно для редактирования формулы -->
    <Transition>
      <section v-if="showEditFormulaMW" class=" fixed inset-0 flex justify-center items-center z-30 p-4">
        <div class="flex flex-col gap-y-4 items-center z-30 p-4 rounded-xl min-w-[500px] bg-white relative">
            <div class="absolute right-0 top-0 cursor-pointer" @click="hideEditFormula">
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
                    v-model="saveFormulaName"
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
    </Transition>
  

    <div class="flex h-screen scrollable">
      <div class="flex flex-col w-1/2 border-r border-gray-200">
        <div class="flex justify-center m-12">
          <p class="text-2xl">Сохранённые формулы</p>
        </div>
        <div class="flex flex-col h-screen mx-6 scrollable">
          <div class="flex flex-col justify-center gap-4">
            <MyFormulaItem 
              v-for="(formulaItem, index) in formulaItems" 
              :key="index"
              :formula="formulaItem.formula" 
              :name="formulaItem.name"
              :index="index" 
              :class="{
                'border-2 border-gray-400': selectedFormulaIndex === index
              }"
              @click="selectFormula(index)"
              @edit-formula="showEditFormula(formulaItem.formula, formulaItem.name)" 
            />
          </div>
        </div>
      </div>
  
      <div class="flex flex-col w-1/2 border-l border-gray-200">
        <div class="flex justify-center m-12">
          <p class="text-2xl">История изменений формулы</p>
        </div>
  
        <div v-if="selectedFormulaIndex !== null" class="flex flex-col h-screen mx-6 scrollable">
          <div class="flex flex-col">
            <div class="flex justify-center mt-4 mx-12 rounded-lg bg-gray-300">
              <p>Формула была изменена 15.02.2024</p>
            </div>
            <MyFormulaItem 
              :formula="formulaItems[selectedFormulaIndex].formula" 
              :name="formulaItems[selectedFormulaIndex].name"
              :index="selectedFormulaIndex" 
            />
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import MyFormulaItem from '@/shared/myFormulaItem.vue';
  import { mapStores } from 'pinia';
  import { useBlurStore } from '@/stores/blurStore';
  import { useStatusWindowStore } from '@/stores/statusWindowStore';
  import { StatusCodes } from '@/helpers/constants';
  
  export default {
    components: {
      MyFormulaItem
    },
    data() {
      return {
        formulaItems: [
          { formula: '\\sin^{2}{(\\alpha)} + \\cos^{2}{(\\alpha)} = 1', name: 'Основное тригонометрическое тождество' },
          { formula: '\\int_{0}^{1} x^2 dx', name: 'Интеграл от x^2' },
          { formula: 'E = mc^2', name: 'Уравнение Эйнштейна' },
          { formula: 'a^2 + b^2 = c^2', name: 'Теорема Пифагора' },
          { formula: '\\log(x) + \\log(y) = \\log(xy)', name: 'Логарифмическое тождество' },
          { formula: 'F = ma', name: 'Закон Ньютона' },
        ],
        selectedFormulaIndex: null as number | null,
        formulaContainer: null as null | HTMLElement,
        formula: '',
        formulaHTML: '',
        sameFormulaHTML: '',
  
        showHistoryMW: false,
        showEditFormulaMW: false,
  
        saveFormulaName: '',
        saveFormulaText: 'ВАававыаыва',
      };
    },
    computed: {
      ...mapStores(useBlurStore, useStatusWindowStore)
    },
    methods: {
      selectFormula(index: number) {
        this.selectedFormulaIndex = index;
      },
      showEditFormula(formula: string, name: string) {
        this.formula = formula;  
        this.saveFormulaName = name;  
        this.showEditFormulaMW = true; 
        this.blurStore.showBlur = true;
      },
  
      hideEditFormula() {
        this.showEditFormulaMW = false; 
        this.blurStore.showBlur = false;
      },

      saveFormula() {
        if (this.saveFormulaName === '') {
          this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Назовите свою формулу!');
          return;
        }

        this.formulaItems[this.selectedFormulaIndex!].name = this.saveFormulaName;  
        this.showEditFormulaMW = false; 
        this.blurStore.showBlur = false;
      }
    }
  };
  </script>
  