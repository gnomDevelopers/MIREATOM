import { defineStore } from "pinia";

export const useArticleStore = defineStore('article', {
  state() {
    return{
      showArticleFormulsMW: false,

      selectedArticleFormuls: ['\\sin^{2}{(\\alpha)} + \\cos^{2}{(\\alpha)} = 1', '\\sin^{2}{(\\alpha)} + \\cos^{2}{(\\alpha)} = 1', '\\sin^{2}{(\\alpha)} + \\cos^{2}{(\\alpha)} = 1'] as string[], 
    }
  },
});