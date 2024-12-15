import type { TMaybeNumber } from "@/helpers/constants";
import { defineStore } from "pinia";

export const useFormulsStore = defineStore('formuls', {
  state() {
    return{
      formulsList: [] as {id: number, title: string, value: string, user_id: number}[],

      selectedFormulaID: null as TMaybeNumber,

      selectedFormulaCommits: [] as {id: number, hash: string, value: string, created_at: string, code_name: string}[],
    }
  },
  actions: {
    
  }
});