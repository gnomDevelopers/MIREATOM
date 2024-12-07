<template>
  <div class="relative" @click.right.prevent="toggleExtraBtns">

    <!--основная кнопка-->
    <slot name="main_btn"></slot>

    <!--значок что есть доп кнопки-->
    <div v-if="showExtraBtns" class="absolute right-1 bottom-1 rounded-full bg-red-600 w-1 h-1"></div>
    
    <!--доп кнопки-->
    <div v-show="calculatorStore.currentOpenedButtonID === id" class="absolute flex flex-row gap-x-1 -top-11 -left-1 p-1 rounded z-10 bg-gray-400">
      <slot name="extra_btns"></slot>
    </div>

  </div>
</template>

<script lang="ts">
import { mapStores } from 'pinia';
import { useCalculatorStore } from '@/stores/calculatorStore';

export default {
  props: {
    id: {
      type: Number,
      required: true,
    },
    showExtraBtns: {
      type: Boolean,
      required: false,
      default: false,
    }
  },
  computed: {
    ...mapStores(useCalculatorStore),
  },
  methods: {
    // переключатель доп клавиш
    toggleExtraBtns(){
      if(!this.showExtraBtns) return;
      if(this.calculatorStore.currentOpenedButtonID === this.id) this.calculatorStore.currentOpenedButtonID = null;
      else this.calculatorStore.currentOpenedButtonID = this.id;
    }
  }
};
</script>