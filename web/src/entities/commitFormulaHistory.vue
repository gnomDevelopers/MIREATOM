<template>
  <div v-for="commit of getFormulaCommits" :key="commit.id">
    <article>
      <h3>Коммит: {{ commit.hash }}</h3>

      <div>
        <div v-html="convertLatexToHTML(commit.value)"></div>
      </div>
    </article>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useFormulsStore } from '@/stores/formulsStore';
import { StatusCodes } from '@/helpers/constants';
import { API_Get_Formula_Commits } from '@/api/api';
import katex from 'katex';

export default {
  data(){
    return{

    }
  },
  computed: {
    ...mapStores(useStatusWindowStore, useFormulsStore),

    getFormulaCommits(){
      return this.formulsStore.selectedFormulaCommits;
    },
  },
  methods: {
    loadFormulaCommits(){
      if(this.formulsStore.selectedFormulaID === null) return;

      const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Получаем историю изменений формулы...', -1);
      
      //получаем все окммиты формулы по id
      API_Get_Formula_Commits(this.formulsStore.selectedFormulaID)
      .then(async (response: any) => {
        this.statusWindowStore.deteleStatusWindow(stID);

        //очищаем список коммитов
        this.formulsStore.selectedFormulaCommits = [];
        for(const commit of response.data){
          try{
            console.log('try to parse: ', commit.difference);
            //пытаемсчя распарсить строку в JSON
            const diff = await JSON.parse(commit.difference);
            console.log(diff);

            const newFormulaCommit = {
              id: commit.id,
              hash: commit.hash, 
              value: '',
              created_at: commit.created_at,
              code_name: commit.code_name,
            };
            //собираем формулу из массива изменений
            let commitValue = '\\large{';
            for(const part of diff){
              switch (part.type){
                case 'equal': commitValue += `${part.content}`; break;
                case 'insert': commitValue += `\\textcolor{#6df274}{ ${part.content} } `; break;
                case 'delete': commitValue += `\\textcolor{#f26161}{ ${part.content} } `; break;
                case 'init': commitValue += `\\textcolor{#e3ca4f}{$${part.content}$} `; break;
                default: commitValue += `${part.content}`; break;
              }
            }
            commitValue += '}';
            //сохраняем рскрашенную формулу
            newFormulaCommit.value = commitValue;
            // newFormulaCommit.value = `\\textcolor{#228B22}{F=ma}`;
            console.log('value: ', commitValue);
            //добавляем коммит в списко коммитов
            this.formulsStore.selectedFormulaCommits.push(newFormulaCommit);
          }
          catch(error){
            console.log(error);
          }
        }
        // this.formulsStore.selectedFormulaCommits.reverse();
      })
      .catch(error => {
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при получении истории изменения формулы!');
      })
    },
    convertLatexToHTML(value: string){
      return katex.renderToString(value, {
        throwOnError: true,
        displayMode: false,
        output: 'mathml',
        trust: false,
      });
    }
  },
  watch: {
    'formulsStore.selectedFormulaID': {
      handler(val){
        this.loadFormulaCommits();
      },
      immediate: true,
    }
  }
};
</script>