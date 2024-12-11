<template>
  <div class="h-full w-full flex flex-col items-center">
    <div class="flex flex-col w-10/12 items-center gap-y-2">
      <section class="flex flex-col gap-y-4 w-full mt-4">

        <div class="flex flex-col gap-y-2">
          <div class="flex flex-row gap-x-2 items-center">
            <p class=" text-lg font-medium">Latex код: </p>
            <input 
              type="text" 
              class="max-w-none w-[400px] outline-none text-lg px-2 py-1 rounded-lg border border-solid border-gray-400 focus:border-sky-500" 
              v-model="formula"
            />
          </div>
          <div class="flex flex-row gap-x-2 items-center">
            <p class=" text-lg font-medium">Формула: </p>
            <span 
              id="formula-area" 
              ref="formulaArea"
              v-html="formulaHTML" 
              class="text-xl" 
              contenteditable="true" 
              @input="updateFormulaFromHTML" 
              @focusout="updateFormula">
            
            </span>
          </div>
        </div>

        <!--клавиатура-->
        <div class="flex flex-row gap-x-4 items-start">
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
        </div>
      </section>
      <button @click="APIRequest" class="border border-solid border-gray-500 rounded px-4 py-2 hover:bg-green-200 active:bg-green-400 transition-colors">Кинуть Z</button>
    </div>
  </div>

</template>
<script lang="ts">
import katex from 'katex';
import { mapStores } from 'pinia';
import { useCalculatorStore } from '@/stores/calculatorStore';

import CalculatorButtonClaster from '@/shared/calculatorButtonClaster.vue';
import CalculatorButtonItem from '@/shared/calculatorButtonItem.vue';
import { API_Health } from '@/api/api';
import { insertHTMLBeforeCursor, parseLatexFromHTML } from '@/helpers/latexHTMLParser';

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
      {id: 4110, formula: '\\lim_{a \\to b} ', argument: '\\lim_{{a} \\to {b}} ', alternatives: []},
      {id: 4120, formula: '\\lim_{a \\to b^-}', argument: '\\lim_{{a} \\to {{b}^{-}}}', alternatives: []},
      {id: 4130, formula: '\\lim_{a \\to b^+}', argument: '\\lim_{{a} \\to {{b}^{+}}}', alternatives: []},
      {id: 4140, formula: '\\infin', argument: '\\infin', alternatives: []},
    ],
    [
      {id: 4210, formula: '\\cfrac{d}{dx}', argument: '\\cfrac{d}{dx}', alternatives: []},
      {id: 4220, formula: '\\int', argument: '\\int', alternatives: []},
      {id: 4230, formula: '\\int_a^b', argument: '\\int_{a}^{b}', alternatives: []},
      {id: 4240, formula: '\\sum_{i=a}^b', argument: '\\sum_{i=a}^{b}', alternatives: []},
    ],
    [
      {id: 4310, formula: '\\log_{10}', argument: '\\log_{10}', alternatives: []},
      {id: 4320, formula: '\\log_2', argument: '\\log_{2}', alternatives: []},
      {id: 4330, formula: '\\log_x', argument: '\\log_{x}', alternatives: []},
      {id: 4340, formula: '\\sum_{i=a}^b', argument: '\\sum_{i=a}^{b}', alternatives: []},
    ],
  ],
];

export default {
  components:{
    CalculatorButtonClaster,
    CalculatorButtonItem,
  },
  data() {
    return{
      formulaContainer: null as null | HTMLElement,
      formula: '',
      formulaHTML: '',

      needUpdateformula: true,
    }
  },
  computed: {
    ...mapStores(useCalculatorStore),

    getButtonsPreset() {
      return StandartButtons[this.calculatorStore.currentTypeButtons];
    },

  },
  mounted(){
    this.formulaContainer = document.getElementById('formula-area');
  },
  methods: {
    setButtonsType(type: number){
      this.calculatorStore.currentTypeButtons = type;
    },
    updateFormulaFromButton(newFormulaPart: string){
      //вставляем в DOM введеную с кнопки формулу
      insertHTMLBeforeCursor(this.formulaContainer!, newFormulaPart);
      //обновляем формулу
      this.formula = parseLatexFromHTML(this.formulaContainer!);
    },
    updateFormulaFromHTML(event: any){
      console.log('ZOV: ', event.target);
      this.needUpdateformula = false;
      this.formula = parseLatexFromHTML(event.target);
    },
    updateFormula(){
      this.formulaHTML = katex.renderToString(this.formula, {
        throwOnError: true,
        displayMode: false,
        output: 'mathml',
        trust: false,
      });
    },
    APIRequest(){
      API_Health();
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