<template>
  <div class="h-full w-full flex flex-col items-center scrollable">
    <div class="flex flex-col lg:flex-row w-full lg:w-10/12 justify-start lg:justify-center items-center lg:items-start gap-2 px-2 mt-4 relative">

      <!--модальное окно истории формул-->
      <Transition>
        <section v-if="showHistoryMW" class="absolute flex flex-col gap-y-4 items-center z-30 p-4 rounded-xl min-w-[500px] bg-white">
          <div class="absolute right-0 top-0 cursor-pointer" @click="hideHistory">
            <svg class="w-10 h-10" viewBox="0 0 52 52" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M34.6668 17.3334L17.3335 34.6667M34.6668 34.6667L17.3335 17.3334" stroke="#8F0101" stroke-width="2" stroke-linecap="round"/>
            </svg>
          </div>
  
          <h1 class="text-2xl text-color-theme cursor-default">История просмотра формул</h1>
  
          <article v-show="formulaHistoryList.length !== 0" id="formulaHistoryScrollWrapper" class="flex flex-col items-center gap-y-3 scrollable rounded-lg max-h-[500px] w-full p-2 bg-gray-100">
            <div v-for="formula of formulaHistoryList" class="w-full">
              <HistoryFormulaItem :formula="formula"/>
            </div>
          </article>

          <article v-if="formulaHistoryList.length === 0" class="flex flex-col items-center cursor-default">
            <div class="px-4 py-2 bg-gray-300 rounded-lg">
              <p class="text-lg">У Вас пока что нет сохраненных формул</p>
            </div>
          </article>
        </section>
      </Transition>

      <!--модальное окно сохранения формулы-->
      <Transition>
        <section v-if="showSaveFormulaMW" class="absolute flex flex-col gap-y-4 items-center z-30 p-4 rounded-xl min-w-[500px] bg-white">
          <div class="absolute right-0 top-0 cursor-pointer" @click="hideSaveFormula">
            <svg class="w-10 h-10" viewBox="0 0 52 52" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M34.6668 17.3334L17.3335 34.6667M34.6668 34.6667L17.3335 17.3334" stroke="#8F0101" stroke-width="2" stroke-linecap="round"/>
            </svg>
          </div>

          <h1 class="text-2xl text-color-theme cursor-default">Сохранение формулы</h1>

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

            <div @click="saveFormula" class="flex flex-row gap-x-2 items-center btn cursor-pointer rounded-lg px-2 py-1 my-2">
              <svg class="w-9 h-9" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M34 42V26H14V42M14 6V16H30M38 42H10C8.93913 42 7.92172 41.5786 7.17157 40.8284C6.42143 40.0783 6 39.0609 6 38V10C6 8.93913 6.42143 7.92172 7.17157 7.17157C7.92172 6.42143 8.93913 6 10 6H32L42 16V38C42 39.0609 41.5786 40.0783 40.8284 40.8284C40.0783 41.5786 39.0609 42 38 42Z" stroke="white" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
              <p class="text-xl text-white">Сохранить формулу</p>
            </div>
          </article>
        </section>
      </Transition>

      <section class="flex flex-col items-center gap-y-6 flex-grow">
        <h1 class="text-center w-full text-2xl">Редактор формул</h1>

        <!--окно с отрендеренной формулой-->
        <article class="flex flex-col gap-2 items-start">
          <p class=" text-lg font-medium ml-1">Ввод с помощью калькулятора: </p>
          <div class="flex flex-row max-h-[500px]">
            <span 
              id="formula-area" 
              v-html="formulaHTML" 
              class="flex flex-row justify-start min-w-[500px] items-center px-2 py-1 text-xl outline-none flex-grow rounded-l-lg border border-solid border-gray-400" 
              contenteditable="true" 
              @input="updateFormulaFromHTML" 
              @focusout="updateFormula">
            </span>

            <div class="btn rounded-r-lg p-1 cursor-pointer">
              <svg class="w-9 h-9" viewBox="0 0 52 52" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M18.8519 21.4883V29.4795C18.8519 33.4283 22.0531 36.6295 26.0019 36.6295C29.9508 36.6295 33.1519 33.4283 33.1519 29.4795V19.3854C33.1519 17.2948 31.4572 15.6001 29.3667 15.6001C27.2761 15.6001 25.5814 17.2948 25.5814 19.3854V29.0589M14.2998 5.20007H37.7003C40.5722 5.20007 42.9004 7.52825 42.9003 10.4002L42.8997 41.6002C42.8997 44.472 40.5716 46.8001 37.6997 46.8L14.2996 46.8C11.4277 46.7999 9.09962 44.4718 9.09964 41.5999L9.09984 10.4C9.09986 7.52818 11.428 5.20007 14.2998 5.20007Z" stroke="white" stroke-width="4" stroke-linecap="round"/>
              </svg>
            </div>
          </div>
        </article>

        <!--клавиатура-->
        <article class="flex flex-row gap-x-4 items-start">
          <!--выбор типа клавиатуры-->
          <aside class="flex flex-col gap-y-1">
            <div 
              @click="setButtonsType(0)" 
              class="flex flex-col w-full px-2 items-center justify-center h-10 rounded button-style" 
              :class="{'button-style-selected': calculatorStore.currentTypeButtons === 0}">
              Стандартные
            </div>

            <div 
              @click="setButtonsType(1)" 
              class="flex flex-col w-full px-2 items-center justify-center h-10 rounded button-style" 
              :class="{'button-style-selected': calculatorStore.currentTypeButtons === 1}">
              Спец. символы
            </div>

            <div 
              @click="setButtonsType(2)" 
              class="flex flex-col w-full px-2 items-center justify-center h-10 rounded button-style" 
              :class="{'button-style-selected': calculatorStore.currentTypeButtons === 2}">
              Тригонометрические
            </div>

            <div 
              @click="setButtonsType(3)" 
              class="flex flex-col w-full px-2 items-center justify-center h-10 rounded button-style" 
              :class="{'button-style-selected': calculatorStore.currentTypeButtons === 3}">
              Дифференцирование
            </div>

            <div 
              @click="setButtonsType(4)" 
              class="flex flex-col w-full px-2 items-center justify-center h-10 rounded button-style" 
              :class="{'button-style-selected': calculatorStore.currentTypeButtons === 4}">
              Доп. знаки
            </div>
          </aside>
          <!--кнопки-->
          <section class="flex flex-row gap-x-1">
            <div v-for="column of getButtonsPreset" class="flex flex-col gap-y-1">

              <div v-for="button of column">
                <CalculatorButtonClaster :id="button.id" :show-extra-btns="button.alternatives.length !== 0">
                  
                  <template #main_btn>
                    <CalculatorButtonItem :is-empty="button.formula === ''" @click.prevent="updateFormulaFromButton(button.argument)">
                      <div v-html="calculatorStore.getButtonByID(button.id, button.formula)"></div>
                    </CalculatorButtonItem>
                  </template>
  
                  <template #extra_btns v-if="button.alternatives.length !== 0">
                    <div v-for="extraButton of button.alternatives">
                      <CalculatorButtonItem @click.prevent="updateFormulaFromButton(extraButton.argument)">
                        <div v-html="calculatorStore.getButtonByID(extraButton.id, extraButton.formula)"></div>
                      </CalculatorButtonItem>
                    </div>
                  </template>
  
                </CalculatorButtonClaster>
              </div>

            </div>
          </section>
        </article>

        <!--окно с LaTeX представлением формулы-->
        <article class="flex flex-col gap-2 items-start">
          <p class=" text-lg font-medium ml-1">LaTeX: </p>
          <div class="flex flex-row">
            <input 
              type="text" 
              class="max-w-none flex-grow w-[500px] outline-none text-lg px-2 py-1 rounded-l-lg border border-solid border-gray-400 focus:border-sky-500" 
              v-model="formula"
            />
            <div @click="copyFormula" class="btn rounded-r-lg p-1 cursor-pointer">
              <svg class="w-9 h-9" viewBox="0 0 52 52" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M43.3333 28.4375L43.3333 9.49979C43.3333 7.84292 41.9902 6.49977 40.3333 6.49979L21.3958 6.50001M30.3333 45.5L15.7083 45.5C13.016 45.5 10.8333 43.3174 10.8333 40.625L10.8333 19.5C10.8333 16.8076 13.016 14.625 15.7083 14.625L30.3333 14.625C33.0257 14.625 35.2083 16.8076 35.2083 19.5L35.2083 40.625C35.2083 43.3174 33.0257 45.5 30.3333 45.5Z" stroke="white" stroke-width="4" stroke-linecap="round"/>
              </svg>
            </div>
          </div>
        </article>

        <article class="flex flex-row gap-x-4">
          <div @click="showHistory" class="flex flex-row gap-x-2 px-2 py-1 cursor-pointer rounded-xl select-none btn">
            <svg class="w-9" viewBox="0 0 52 52" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M42.2601 26C42.2601 36.7696 33.5274 45.5 22.7551 45.5C11.9827 45.5 3.25 36.7696 3.25 26C3.25 15.2304 11.9827 6.5 22.7551 6.5C29.9747 6.5 36.2782 10.4214 39.6507 16.25M36.9093 28.1399L41.7856 23.2649L46.6619 28.1399M30.0625 31.3173L22.75 28.8798V18.6875" stroke="white" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
            <p class="text-white text-xl">История просмотра</p>
          </div>

          <div @click="showSaveFormula" class="flex flex-row gap-x-2 px-2 py-1 cursor-pointer rounded-xl select-none btn">
            <svg class="w-9" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M34 42V26H14V42M14 6V16H30M38 42H10C8.93913 42 7.92172 41.5786 7.17157 40.8284C6.42143 40.0783 6 39.0609 6 38V10C6 8.93913 6.42143 7.92172 7.17157 7.17157C7.92172 6.42143 8.93913 6 10 6H32L42 16V38C42 39.0609 41.5786 40.0783 40.8284 40.8284C40.0783 41.5786 39.0609 42 38 42Z" stroke="white" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
            <p class="text-white text-xl">Сохранить формулу</p>
          </div>
        </article>

      </section>

      <section class="flex-grow flex flex-col gap-y-6 items-center">
        <h1 class="text-center w-full text-2xl">Анализ формул</h1>
                
        <!--окно с отрендеренной формулой и анализом-->
        <article class="flex flex-col gap-2 items-start">

          <div class="flex flex-row mt-9">
            <span 
              id="analyse-formula" 
              v-html="formulaHTML" 
              class="flex flex-row w-[500px] justify-start items-center px-2 py-1 text-xl outline-none flex-grow rounded-l-lg border border-solid border-gray-400" >
            </span>

            <div class="btn rounded-r-lg p-1 cursor-pointer">
              <svg class="w-9 h-9" viewBox="0 0 52 52" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M12.175 26.9457L9.40572 29.715M9.11635 19.5599H5.19995M9.40572 9.40669L12.175 12.176M19.5608 5.19995V9.11635M29.714 9.40669L26.9447 12.176M34.7047 34.6568L44.6763 31.1759C46.5454 30.5234 46.6178 27.9062 44.7868 27.1795L21.9729 19.2346C20.2579 18.554 18.5335 20.2363 19.1709 21.9681L26.6413 45.4173C27.3216 47.2652 29.9375 47.2594 30.6377 45.4084L34.7047 34.6568Z" stroke="white" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            </div>
          </div>
        </article>

        <article class="flex flex-col gap-2 rounded-lg p-4 bg-gray-50 border border-solid border-gray-200">
          <h2 class="text-xl font-bold text-center">В базе сохранённых формул найдена похожая формула</h2>

          <span 
            id="analyse-formula" 
            v-html="sameFormulaHTML" 
            class="flex flex-row justify-center items-center px-2 py-1 text-xl" >
          </span>

          <h3 class="text-xl font-bold">Процент совпадения: 100%</h3>

          <div>
            <p class="text-lg font-bold">Название формулы:</p>
            <p class="text-lg">Основное тригонометрическое тождество</p>
          </div>
          <div>
            <p class="text-lg font-bold">Автор формулы:</p>
            <p class="text-lg">Семёнов Семён Семёнович</p>
          </div>
          <div>
            <p class="text-lg font-bold">Дата загрузки:</p>
            <p class="text-lg">15.12.2024</p>
          </div>
        </article>
      </section>
    </div>
  </div>

</template>
<script lang="ts">
import katex from 'katex';
import { mapStores } from 'pinia';
import { useCalculatorStore } from '@/stores/calculatorStore';
import { useBlurStore } from '@/stores/blurStore';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { API_Get_Formuls_History, API_Save_Formula } from '@/api/api';
import { garbageCollector, getFirstMrow, insertEmptyElementsInHTML, insertHTMLBeforeCursor, parseLatexFromHTML, renderKatex } from '@/helpers/latexHTMLParser';
import { StatusCodes } from '@/helpers/constants';
import { nextTick } from 'vue';

import CalculatorButtonClaster from '@/shared/calculatorButtonClaster.vue';
import CalculatorButtonItem from '@/shared/calculatorButtonItem.vue';
import HistoryFormulaItem from '@/shared/historyFormulaItem.vue';

const StandartButtons = [
  [ // standar buttons
    [
      {id: 1110, formula: '( )', argument: '( )', alternatives: [
        {id: 1111, formula: '( )', argument: '( )'},
        {id: 1112, formula: '(', argument: '('},
        {id: 1113, formula: ')', argument: ')'},
      ]},
      {id: 1120, formula: '\\cfrac{a}{b}', argument: '\\cfrac{a}{b}', alternatives: []},
      {id: 1130, formula: 'x^2', argument: '{x}^{2}', alternatives: [
        {id: 1131, formula: 'x^2', argument: '{x}^{2}'},
        {id: 1132, formula: 'x^3', argument: '{x}^{3}'},
        {id: 1133, formula: 'x^n', argument: '{x}^{n}'},
      ]},
      {id: 1140, formula: '\\pi', argument: '\\pi', alternatives: []},
    ],
    [
      {id: 1210, formula: '>', argument: '>', alternatives: [
        {id: 1211, formula: '>', argument: '>'},
        {id: 1212, formula: '\\ge', argument: '\\ge'},
        {id: 1213, formula: '<', argument: '<'},
        {id: 1214, formula: '\\le', argument: '\\le'},
      ]},
      {id: 1220, formula: '\\sqrt{\\smash[b]{x}}', argument: '\\sqrt{x}', alternatives: [
        {id: 1221, formula: '\\sqrt{x}', argument: '\\sqrt{x}'},
        {id: 1222, formula: '\\sqrt[3]{x}', argument: '\\sqrt[{3}]{x}'},
        {id: 1223, formula: '\\sqrt[n]{x}', argument: '\\sqrt[{n}]{x}'},
      ]},
      {id: 1230, formula: 'x', argument: 'x', alternatives: [
        {id: 1231, formula: 'x', argument: 'x'},
        {id: 1232, formula: 'y', argument: 'y'},
        {id: 1233, formula: 'z', argument: 'z'},
      ]},
      {id: 1240, formula: '\\%', argument: '\\%', alternatives: []},
    ],
    [
      {id: 1310, formula: '7', argument: '7', alternatives: []},
      {id: 1320, formula: '4', argument: '4', alternatives: []},
      {id: 1330, formula: '1', argument: '1', alternatives: []},
      {id: 1340, formula: '0', argument: '0', alternatives: []},
    ],
    [
      {id: 1410, formula: '8', argument: '8', alternatives: []},
      {id: 1420, formula: '5', argument: '5', alternatives: []},
      {id: 1430, formula: '2', argument: '2', alternatives: []},
      {id: 1440, formula: '.', argument: '.', alternatives: []},
    ],
    [
      {id: 1510, formula: '9', argument: '9', alternatives: []},
      {id: 1520, formula: '6', argument: '6', alternatives: []},
      {id: 1530, formula: '3', argument: '3', alternatives: []},
      {id: 1540, formula: '=', argument: '=', alternatives: []},
    ],
    [
      {id: 1610, formula: '/', argument: '/', alternatives: []},
      {id: 1620, formula: '*', argument: '*', alternatives: []},
      {id: 1630, formula: '-', argument: '-', alternatives: []},
      {id: 1640, formula: '+', argument: '+', alternatives: []},
    ],
  ],
  [ // special buttons
    [
      {id: 2110, formula: 'a', argument: 'a', alternatives: []},
      {id: 2120, formula: 'i', argument: 'i', alternatives: []},
      {id: 2130, formula: 'q', argument: 'q', alternatives: []},
      {id: 2140, formula: 'y', argument: 'y', alternatives: []},
      {id: 2150, formula: '\\alpha', argument: '\\alpha', alternatives: []},
      {id: 2160, formula: '\\iota', argument: '\\iota', alternatives: []},
      {id: 2170, formula: '\\rho', argument: '\\rho', alternatives: []},
      {id: 2180, formula: '\\varepsilon', argument: '\\varepsilon', alternatives: []},
    ],
    [
      {id: 2210, formula: 'b', argument: 'b', alternatives: []},
      {id: 2220, formula: 'j', argument: 'j', alternatives: []},
      {id: 2230, formula: 'r', argument: 'r', alternatives: []},
      {id: 2240, formula: 'z', argument: 'z', alternatives: []},
      {id: 2250, formula: '\\beta', argument: '\\beta', alternatives: []},
      {id: 2260, formula: '\\kappa', argument: '\\kappa', alternatives: []},
      {id: 2270, formula: '\\sigma', argument: '\\sigma', alternatives: []},
      {id: 2280, formula: '\\varkappa', argument: '\\varkappa', alternatives: []},
    ],
    [
      {id: 2310, formula: 'c', argument: 'c', alternatives: []},
      {id: 2320, formula: 'k', argument: 'k', alternatives: []},
      {id: 2330, formula: 's', argument: 's', alternatives: []},
      {id: 2340, formula: '', argument: '', alternatives: []},
      {id: 2350, formula: '\\gamma', argument: '\\gamma', alternatives: []},
      {id: 2360, formula: '\\lambda', argument: '\\lambda', alternatives: []},
      {id: 2370, formula: '\\tau', argument: '\\tau', alternatives: []},
      {id: 2380, formula: '\\vartheta', argument: '\\vartheta', alternatives: []},
    ],
    [
      {id: 2410, formula: 'd', argument: 'd', alternatives: []},
      {id: 2420, formula: 'l', argument: 'l', alternatives: []},
      {id: 2430, formula: 't', argument: 't', alternatives: []},
      {id: 2440, formula: '', argument: '', alternatives: []},
      {id: 2450, formula: '\\delta', argument: '\\delta', alternatives: []},
      {id: 2460, formula: '\\mu', argument: '\\mu', alternatives: []},
      {id: 2470, formula: '\\upsilon', argument: '\\upsilon', alternatives: []},
      {id: 2480, formula: '\\thetasym', argument: '\\thetasym', alternatives: []},
    ],
    [
      {id: 2510, formula: 'e', argument: 'e', alternatives: []},
      {id: 2520, formula: 'm', argument: 'm', alternatives: []},
      {id: 2530, formula: 'u', argument: 'u', alternatives: []},
      {id: 2540, formula: '', argument: '', alternatives: []},
      {id: 2550, formula: '\\epsilon', argument: '\\epsilon', alternatives: []},
      {id: 2560, formula: '\\nu', argument: '\\nu', alternatives: []},
      {id: 2570, formula: '\\phi', argument: '\\phi', alternatives: []},
      {id: 2580, formula: '\\varpi', argument: '\\varpi', alternatives: []},
    ],
    [
      {id: 2610, formula: 'f', argument: 'f', alternatives: []},
      {id: 2620, formula: 'n', argument: 'n', alternatives: []},
      {id: 2630, formula: 'v', argument: 'v', alternatives: []},
      {id: 2640, formula: '', argument: '', alternatives: []},
      {id: 2650, formula: '\\zeta', argument: '\\zeta', alternatives: []},
      {id: 2660, formula: '\\xi', argument: '\\xi', alternatives: []},
      {id: 2670, formula: '\\chi', argument: '\\chi', alternatives: []},
      {id: 2680, formula: '\\varrho', argument: '\\varrho', alternatives: []},
    ],
    [
      {id: 2710, formula: 'g', argument: 'g', alternatives: []},
      {id: 2720, formula: 'o', argument: 'o', alternatives: []},
      {id: 2730, formula: 'w', argument: 'w', alternatives: []},
      {id: 2740, formula: '', argument: '', alternatives: []},
      {id: 2750, formula: '\\eta', argument: '\\eta', alternatives: []},
      {id: 2760, formula: '\\omicron', argument: '\\omicron', alternatives: []},
      {id: 2770, formula: '\\psi', argument: '\\psi', alternatives: []},
      {id: 2780, formula: '\\varsigma', argument: '\\varsigma', alternatives: []},
    ],
    [
      {id: 2810, formula: 'h', argument: 'h', alternatives: []},
      {id: 2820, formula: 'p', argument: 'p', alternatives: []},
      {id: 2830, formula: 'x', argument: 'x', alternatives: []},
      {id: 2840, formula: '', argument: '', alternatives: []},
      {id: 2850, formula: '\\theta', argument: '\\theta', alternatives: []},
      {id: 2860, formula: '\\pi', argument: '\\pi', alternatives: []},
      {id: 2870, formula: '\\omega', argument: '\\omega', alternatives: []},
      {id: 2880, formula: '\\varphi', argument: '\\varphi', alternatives: []},
    ],
  ],
  [ // trigonometrical buttons
    [
      {id: 3110, formula: '\\sin', argument: '\\sin{x}', alternatives: []},
      {id: 3120, formula: '\\arcsin', argument: '\\arcsin{x}', alternatives: []},
      {id: 3130, formula: '\\sinh', argument: '\\sinh{x}', alternatives: []},
      {id: 3140, formula: '\\operatorname{arsinh}', argument: '\\operatorname{arsinh}', alternatives: []},
      {id: 3150, formula: 'x^\\circ', argument: '{x}^{\\circ}', alternatives: []},
    ],
    [
      {id: 3210, formula: '\\cos', argument: '\\cos{x}', alternatives: []},
      {id: 3220, formula: '\\arccos', argument: '\\arccos{x}', alternatives: []},
      {id: 3230, formula: '\\cosh', argument: '\\cosh{x}', alternatives: []},
      {id: 3240, formula: '\\operatorname{arcosh}', argument: '\\operatorname{arcosh}', alternatives: []},
    ],
    [
      {id: 3410, formula: '\\tg', argument: '\\tg{x}', alternatives: []},
      {id: 3420, formula: '\\arctg', argument: '\\arctg{x}', alternatives: []},
      {id: 3430, formula: '\\th', argument: '\\th{x}', alternatives: []},
      {id: 3440, formula: '\\operatorname{arth}', argument: '\\operatorname{arth}', alternatives: []},
    ],
    [
      {id: 3510, formula: '\\ctg', argument: '\\ctg{x}', alternatives: []},
      {id: 3520, formula: '\\arcctg', argument: '\\arcctg{x}', alternatives: []},
      {id: 3530, formula: '\\cth', argument: '\\cth{x}', alternatives: []},
      {id: 3540, formula: '\\operatorname{arcth}', argument: '\\operatorname{arcth}', alternatives: []},
    ],
    [
      {id: 3610, formula: '\\sec', argument: '\\sec{x}', alternatives: []},
      {id: 3620, formula: '\\operatorname{arcsec}', argument: '\\operatorname{arcsec}', alternatives: []},
      {id: 3630, formula: '\\operatorname{sech}', argument: '\\operatorname{sech}', alternatives: []},
      {id: 3640, formula: '\\operatorname{arsech}', argument: '\\operatorname{arsech}', alternatives: []},
    ],
    [
      {id: 3710, formula: '\\csc', argument: '\\csc{x}', alternatives: []},
      {id: 3720, formula: '\\operatorname{arccsc}', argument: '\\operatorname{arccsc}', alternatives: []},
      {id: 3730, formula: '\\operatorname{csch}', argument: '\\operatorname{csch}', alternatives: []},
      {id: 3740, formula: '\\operatorname{arcsch}', argument: '\\operatorname{arcsch}', alternatives: []},
    ],
  ],
  [
    [
      {id: 4110, formula: '\\lim_{a \\to b} ', argument: '\\lim_{{a} \\to {b}}{}', alternatives: []},
      {id: 4120, formula: '\\lim_{a \\to b^-}', argument: '\\lim_{{a} \\to {{b}^{-}}}{}', alternatives: []},
      {id: 4130, formula: '\\lim_{a \\to b^+}', argument: '\\lim_{{a} \\to {{b}^{+}}}{}', alternatives: []},
      {id: 4140, formula: '\\infin', argument: '\\infin', alternatives: []},
    ],
    [
      {id: 4210, formula: '\\cfrac{d}{dx}', argument: '\\cfrac{d}{dx}', alternatives: []},
      {id: 4220, formula: '\\int', argument: '\\int{}', alternatives: []},
      {id: 4230, formula: '\\int_a^b', argument: '\\int_{a}^{b}{}', alternatives: []},
      {id: 4240, formula: '\\sum_{i=a}^b', argument: '\\sum_{i=a}^{b}{}', alternatives: []},
    ],
    [
      {id: 4310, formula: '\\log_{10}', argument: '\\log_{10}{}', alternatives: []},
      {id: 4320, formula: '\\log_2', argument: '\\log_{2}{}', alternatives: []},
      {id: 4330, formula: '\\log_x', argument: '\\log_{x}{}', alternatives: []},
      {id: 4340, formula: '\\sum_{i=a}^b', argument: '\\sum_{i=a}^{b}{}', alternatives: []},
    ],
  ],
];

export default {
  components:{
    CalculatorButtonClaster,
    CalculatorButtonItem,
    HistoryFormulaItem,
  },
  data() {
    return{
      formulaContainer: null as null | HTMLElement,
      formula: '',
      formulaHTML: '',
      analyseFormula: '\\sin^{2}{(\\alpha)} + \\cos^{2}{(\\alpha)} = 1',
      analyseFormulaHTML: '',
      sameFormula: '\\sin^{2}{(\\alpha)} + \\cos^{2}{(\\alpha)} = 1',
      sameFormulaHTML: '',

      needUpdateformula: true,

      showHistoryMW: false,
      showSaveFormulaMW: false,

      saveFormulaName: '',

      formulaHistoryList: [] as string[],
      formulaHistoryPage: 1,
    }
  },
  computed: {
    ...mapStores(useCalculatorStore, useBlurStore, useStatusWindowStore, useUserInfoStore),

    getButtonsPreset() {
      return StandartButtons[this.calculatorStore.currentTypeButtons];
    },

  },
  mounted(){
    this.formulaContainer = document.getElementById('formula-area');

    //рисуем заглущечные формулы
    this.analyseFormulaHTML = katex.renderToString(this.analyseFormula, {
      throwOnError: true,
      displayMode: false,
      output: 'mathml',
      trust: false,
    });

    this.sameFormulaHTML = katex.renderToString(this.sameFormula, {
      throwOnError: true,
      displayMode: false,
      output: 'mathml',
      trust: false,
    });
  },
  methods: {
    setButtonsType(type: number){
      this.calculatorStore.currentTypeButtons = type;
    },
    updateFormulaFromButton(newFormulaPart: string){
      //скрываем экстра кнопки
      this.calculatorStore.currentOpenedButtonID = null;

      //вставляем в DOM введеную с кнопки формулу
      insertHTMLBeforeCursor(this.formulaContainer!, newFormulaPart);

      //парсим html в Latex формулу
      this.formula = parseLatexFromHTML(this.formulaContainer!);
    },
    updateFormulaFromHTML(event: any){
      //отключаем перерисовку html при изменении формулы
      this.needUpdateformula = false;

      //парсим html в Latex формулу
      this.formula = parseLatexFromHTML(event.target);

      //в следующем кадре, после отрисовки, настраиваем скроллы у формулы
      nextTick(() => {
        this.handleFormulaWindowSize();
      });
    },
    updateFormula(){
      //рендерим формулу
      this.formulaHTML = katex.renderToString(this.formula, {
        throwOnError: true,
        displayMode: false,
        output: 'mathml',
        trust: false,
      });

      nextTick(() => {
        //удаляем пустые элементы
        garbageCollector(this.formulaContainer!);

        //отключаем перерисовку html при изменении формулы
        this.needUpdateformula = false;

        //повторно обновляем формулу
        this.formula = parseLatexFromHTML(this.formulaContainer!);

        //вставка пустых элементов для курсора
        const mrow = getFirstMrow(this.formulaContainer!);
        if(mrow === undefined) return;
        insertEmptyElementsInHTML(mrow);

        //настраиваем скроллы у формулы
        this.handleFormulaWindowSize();
      })
    },
    handleFormulaWindowSize(){
      if(this.formulaContainer!.offsetWidth > 500){
        this.formulaContainer!.classList.add('scroll-x');
      }
      else{
        this.formulaContainer!.classList.remove('scroll-x');
      }

      if(this.formulaContainer!.offsetHeight >= 500){
        this.formulaContainer!.classList.add('scroll-y');
      }
      else{
        this.formulaContainer!.classList.remove('scroll-y');
      }
    },
    showHistory(){
      //если пользователь не атворизован - выход
      if(this.userInfoStore.userID === null){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Войдите в аккаунт, чтобы просматривать историю ваших формул!');
        return;
      }
      //открываем окно истории
      this.blurStore.showBlur = true;
      this.showHistoryMW = true;

      //сбро до заводских настроек
      this.formulaHistoryPage = 1;
      this.formulaHistoryList = [];

      const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Подгружаем историю формул...', -1);
        
      //получение первой страницы истории
      API_Get_Formuls_History(this.userInfoStore.userID, this.formulaHistoryPage)
      .then((response:any) => {
        this.statusWindowStore.deteleStatusWindow(stID);
        //сохраняем полученные формулы в массив
        for(const item of response.data){
          this.formulaHistoryList.push(item.value);
        }
        //если получено меньше 20 элементов - значит больше формул нет
        if(response.data.length < 20) return;
        
        //получаем обертку для истории формул
        const historyHTML = document.getElementById('formulaHistoryScrollWrapper');
        //если ничего нет - выход
        if(!historyHTML) return;

        //иначе добавляем слушатель на скролл
        historyHTML.addEventListener('scroll', this.handleHistoryScroll);
      })
      .catch(error => {
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при получении истории!');
      })
    },
    hideHistory(){
      this.blurStore.showBlur = false;
      this.showHistoryMW = false;
    },
    showSaveFormula(){
      if(this.userInfoStore.userID === null){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Войдите в аккаунт, чтобы сохранять свои формулы!');
        return;
      }
      if(this.formula === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Вы не ввели формулу!', 1500);
        return;
      }
      this.blurStore.showBlur = true;
      this.showSaveFormulaMW = true;
    },
    hideSaveFormula(){
      this.blurStore.showBlur = false;
      this.showSaveFormulaMW = false;
    },
    saveFormula(){
      if(this.saveFormulaName === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Назовите свою формулу!');
        return;
      }
      //API
      const data = {title: this.saveFormulaName, value: this.formula};
      const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Сохраняем формулу...', -1);

      API_Save_Formula(data)
      .then((response:any) => {
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Формула сохранена!');

        this.hideSaveFormula();
      })
      .catch(error => {
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при сохранении формулы!');
      })
    },
    handleHistoryScroll(event:any){
      const scrollHeight = event.target.scrollHeight;
      const scrollTop = event.target.scrollTop;
      const clientHeight = event.target.clientHeight;

      if (scrollTop + clientHeight >= scrollHeight) {    
        console.log("Scrolled to bottom!");

        if(this.userInfoStore.userID === null) return;

        this.formulaHistoryPage++;
        API_Get_Formuls_History(this.userInfoStore.userID, this.formulaHistoryPage)
        .then((response:any) => {
          //сохраняем полученные формулы в массив
          for(const item of response.data){
            this.formulaHistoryList.push(item.value);
          }

          //если получено меньше 20 элементов - значит больше формул нет
          if(response.data.length < 20) event.target.removeEventListener('scroll', this.handleHistoryScroll);
        });
      }
    },
    copyFormula(){
      navigator.clipboard.writeText(this.formula).then(() => {
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Формула скопирована!', 1500);
      })
      .catch((err) => {
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при копировании формулы!');
        console.error('Could not copy text: ', err);
      });
    }
  },
  watch: {
    formula(val){
      if(this.needUpdateformula){
        this.updateFormula();
      }
      else this.needUpdateformula = true;
    },
  }
};
</script>