import { defineStore } from "pinia";
import { DEVMODE, type TMaybeNumber } from "@/helpers/constants";
import katex from 'katex';

export const useCalculatorStore = defineStore('calculator', {
  state() {
    return{ 
      currentTypeButtons: 3, // 0 - стандартные, 1 - спец. символы и тд.
      currentOpenedButtonID: null as TMaybeNumber,
      buttonsMap: new Map<number, string>(),
    }
  },
  actions: {
    getButtonByID(buttonID: number, buttonFormula?: string):string{
      //try found button in memory
      if(this.buttonsMap.has(buttonID)) return this.buttonsMap.get(buttonID)!;
      //if not found and no formula - error
      if(buttonFormula === undefined){
        if(DEVMODE) console.error(`Button with id: ${buttonID} not found!`);
        return '';
      }
      //else - create a new button
      const newButton = katex.renderToString(buttonFormula, {
        displayMode: false,
        output: 'mathml',
      });
      //save new button to memory
      this.buttonsMap.set(buttonID, newButton);
      //return new button
      return newButton;
    }
  }
});