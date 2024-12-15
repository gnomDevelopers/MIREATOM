import { defineStore } from "pinia";
import { useStatusWindowStore } from "./statusWindowStore";
import { type TMaybeNumber, StatusCodes } from "@/helpers/constants";
import { API_Get_Formula_Commits } from "@/api/api";

export const useFormulsStore = defineStore('formuls', {
  state() {
    return{
      formulsList: [] as {id: number, title: string, value: string, user_id: number}[],

      selectedFormulaID: null as TMaybeNumber,

      selectedFormulaCommits: [] as {id: number, hash: string, value: string, created_at: string, code_name: string}[],
    }
  },
  actions: {
    loadFormulaCommits(){
      if(this.selectedFormulaID === null) return;

      const statusWindowStore = useStatusWindowStore();

      const stID = statusWindowStore.showStatusWindow(StatusCodes.loading, 'Получаем историю изменений формулы...', -1);
      
      //получаем все окммиты формулы по id
      API_Get_Formula_Commits(this.selectedFormulaID)
      .then(async (response: any) => {
        statusWindowStore.deteleStatusWindow(stID);

        //очищаем список коммитов
        this.selectedFormulaCommits = [];
        for(const commit of response.data){
          try{
            //пытаемсчя распарсить строку в JSON
            const diff = await JSON.parse(commit.difference);

            const newFormulaCommit = {
              id: commit.id,
              hash: commit.hash, 
              value: '',
              newValue: '',
              created_at: commit.created_at,
              code_name: commit.code_name,
            };
            //собираем формулу из массива изменений
            let commitValue = '';
            let newValue = '';
            for(const part of diff){
              switch (part.type){
                case 'equal': commitValue += `${part.content}`; newValue += part.content; break;
                case 'insert': commitValue += `\\textcolor{#6df274}{ ${part.content} } `; newValue += part.content; break;
                case 'delete': commitValue += `\\textcolor{#f26161}{ ${part.content} } `; break;
                case 'init': commitValue += `\\textcolor{#e3ca4f}{$${part.content}$} `; newValue += part.content; break;
                default: commitValue += `${part.content}`; newValue += part.content; break;
              }
            }
            //сохраняем рскрашенную формулу
            newFormulaCommit.value = commitValue;
            //сохраняем новую latex формулу
            newFormulaCommit.newValue = newValue;
            //добавляем коммит в списко коммитов
            this.selectedFormulaCommits.push(newFormulaCommit);
          }
          catch(error){
            console.log(error);
          }
        }
        // this.formulsStore.selectedFormulaCommits.reverse();
      })
      .catch(error => {
        statusWindowStore.deteleStatusWindow(stID);
        statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при получении истории изменения формулы!');
      })
    },
  }
});