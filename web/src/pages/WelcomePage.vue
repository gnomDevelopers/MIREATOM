<template>
  <div class="h-full w-full flex flex-col items-center">
    <div class="flex flex-col w-10/12 items-center gap-y-2">
      <section class="w-full mt-4">
        <div>
          <p class=" text-lg font-medium">
            Formula: 
            <span id="formula-area"></span>
          </p>
        </div>
        <!--клавиатура-->
        <div class="flex flex-row gap-x-4 items-start">
          <!--выбор типа клавиатуры-->
          <aside class="flex flex-col gap-y-1">
            <div @click="setButtonsType(1)" class="flex flex-col w-full px-2 items-center justify-center h-10 rounded button-style" :class="{'bg-slate-100': calculatorStore.currentTypeButtons === 1}">Стандартные</div>
            <div @click="setButtonsType(2)" class="flex flex-col w-full px-2 items-center justify-center h-10 rounded button-style" :class="{'bg-slate-100': calculatorStore.currentTypeButtons === 2}">Спец. символы</div>
            <div @click="setButtonsType(3)" class="flex flex-col w-full px-2 items-center justify-center h-10 rounded button-style" :class="{'bg-slate-100': calculatorStore.currentTypeButtons === 3}">Тригонометрические</div>
            <div @click="setButtonsType(4)" class="flex flex-col w-full px-2 items-center justify-center h-10 rounded button-style" :class="{'bg-slate-100': calculatorStore.currentTypeButtons === 4}">Дифференцирование</div>
            <div @click="setButtonsType(5)" class="flex flex-col w-full px-2 items-center justify-center h-10 rounded button-style" :class="{'bg-slate-100': calculatorStore.currentTypeButtons === 5}">Доп. знаки</div>
          </aside>
          <!--кнопки-->
          <section class="flex flex-row gap-x-1">
            <div v-for="column of getButtonsPreset" class="flex flex-col gap-y-1">

              <div v-for="button of column">
                <CalculatorButtonClaster :id="button.id" :show-extra-btns="button.alternatives.length !== 0">
                  
                  <template #main_btn>
                    <CalculatorButtonItem>
                      <div v-html="calculatorStore.getButtonByID(button.id, button.formula)"></div>
                    </CalculatorButtonItem>
                  </template>
  
                  <template #extra_btns v-if="button.alternatives.length !== 0">
                    <div v-for="extraButton of button.alternatives">
                      <CalculatorButtonItem>
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
    </div>
  </div>

</template>
<script lang="ts">
import katex from 'katex';
import { mapStores } from 'pinia';
import { useCalculatorStore } from '@/stores/calculatorStore';

import CalculatorButtonClaster from '@/shared/calculatorButtonClaster.vue';
import CalculatorButtonItem from '@/shared/calculatorButtonItem.vue';

const StandartButtons = [
  [
    {id: 1110, formula: '( )', argument: '( )', alternatives: []},
    {id: 1120, formula: '\\cfrac{a}{b}', argument: '\\cfrac{a}{b}', alternatives: []},
    {id: 1130, formula: 'x^2', argument: 'x^2', alternatives: [
      {id: 1131, formula: 'x^2', argument: 'x^2'},
      {id: 1132, formula: 'x^3', argument: 'x^3'},
      {id: 1133, formula: 'x^n', argument: 'x^n'},
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
      {id: 1222, formula: '\\sqrt[3]{x}', argument: '\\sqrt[3]{x}'},
      {id: 1223, formula: '\\sqrt[n]{x}', argument: '\\sqrt[n]{x}'},
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
];

export default {
  components:{
    CalculatorButtonClaster,
    CalculatorButtonItem,
  },
  data() {
    return{
      formulaContainer: null as null | HTMLElement,
      buttonsID: 1,
    }
  },
  computed: {
    ...mapStores(useCalculatorStore),

    getButtonsPreset() {
      return StandartButtons;
    }
  },
  mounted(){
    this.formulaContainer = document.getElementById('formula-area');
    katex.render("c = \\pm\\sqrt{a^2 + b^2}", this.formulaContainer!, {
      throwOnError: true,
      displayMode: false,
      output: 'mathml',
      trust: false,
    });

    // this.calculatorStore.renderButtons();
    // console.log(this.calculatorStore.buttonsList);
  },
  methods: {
    setButtonsType(type: number){
      this.calculatorStore.currentTypeButtons = type;
    }
  }
};
</script>