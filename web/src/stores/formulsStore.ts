import { defineStore } from "pinia";

export const useFormulsStore = defineStore('formuls', {
  state() {
    return{
      formulsList: [] as {id: number, title: string, value: string, user_id: number}[],
    }
  },
  actions: {
    
  }
});