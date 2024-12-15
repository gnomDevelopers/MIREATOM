<template>
  <div class="flex flex-col items-center h-full w-full">

    <!--модальное окно формул в статье-->
    <Transition>
      <section v-if="articleStore.showArticleFormulsMW" class="absolute flex flex-col gap-y-4 items-center z-30 p-4 rounded-xl min-w-[500px] bg-white">
        <div class="absolute right-0 top-0 cursor-pointer" @click="hideArticleFormuls">
          <svg class="w-10 h-10" viewBox="0 0 52 52" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M34.6668 17.3334L17.3335 34.6667M34.6668 34.6667L17.3335 17.3334" stroke="#8F0101" stroke-width="2" stroke-linecap="round"/>
          </svg>
        </div>

        <h1 class="text-2xl text-color-theme cursor-default">Формулы, приведённые в статье</h1>

        <article class="flex flex-col items-center gap-y-3 scrollable rounded-lg max-h-[500px] w-full p-2 bg-gray-100">

          <div v-for="formula of articleStore.selectedArticleFormuls">
            <HistoryFormulaItem :formula="formula"/>
          </div>

        </article>
      </section>
    </Transition>

    <!-- модальное окно для загрузки статьи -->

    <!-- Исправить ошибки!! -->
    <Transition>
      <UploadArticle v-if="showUploadArticleMW" @close-window="hideUploadArticle" @update-articles="() => { getArticles(); getMyArticles(); }"/>
    </Transition>

    <div class="flex flex-row w-full h-full items-start">

      <aside class="flex flex-col h-full justify-start gap-y-4 items-center w-72 border-r-2 p-4 border-solid border-gray-300">
        <h1 class="text-2xl">Фильтры</h1>

        <section class="flex flex-col">
          <label class="text-lg ml-2">Название статьи</label>

          <div class="flex flex-row items-stretch">
            <input 
              type="text" 
              placeholder="Высшая математика"
              v-model="filters.articleTitle"
              class="border border-solid rounded-l-lg outline-none px-2 border-gray-300 focus:border-red-800"/>
              <div @click="filter" class="px-2 py-1 rounded-r-lg cursor-pointer btn">
              <svg class="w-6 h-6" viewBox="0 0 26 26" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M23.8333 23.8334L21.6667 21.6667M22.75 12.4584C22.75 18.1423 18.1423 22.7501 12.4583 22.7501C6.7744 22.7501 2.16667 18.1423 2.16667 12.4584C2.16667 6.77448 6.7744 2.16675 12.4583 2.16675C18.1423 2.16675 22.75 6.77448 22.75 12.4584Z" stroke="#FAFCFF" stroke-width="1.625" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            </div>
          </div>

        </section>

        <section class="flex flex-col">
          <label class="text-lg ml-2">Автор статьи</label>

          <div class="flex flex-row items-stretch">
            <input 
              type="text" 
              placeholder="Иванов Иван"
              v-model="filters.author"
              class="border border-solid rounded-l-lg outline-none px-2 border-gray-300 focus:border-red-800"/>
            <div @click="filter" class="px-2 py-1 rounded-r-lg cursor-pointer btn">
              <svg class="w-6 h-6" viewBox="0 0 26 26" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M23.8333 23.8334L21.6667 21.6667M22.75 12.4584C22.75 18.1423 18.1423 22.7501 12.4583 22.7501C6.7744 22.7501 2.16667 18.1423 2.16667 12.4584C2.16667 6.77448 6.7744 2.16675 12.4583 2.16675C18.1423 2.16675 22.75 6.77448 22.75 12.4584Z" stroke="#FAFCFF" stroke-width="1.625" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            </div>
          </div>
          
        </section>

        <section class="flex flex-col">
          <label class="text-lg ml-2">Наука</label>

          <div class="flex flex-row items-stretch">
            <input 
              type="text" 
              placeholder="Физика"
              v-model="newScience"
              class="border border-solid rounded-l-lg outline-none px-2 border-gray-300 focus:border-red-800"/>
            <div 
              class="px-2 py-1 rounded-r-lg cursor-pointer btn"
              @click="addScience">
              <svg class="w-6 h-6" viewBox="0 0 26 26" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M6.5 13H19.5M13 19.5V6.5" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            </div>
          </div>

          <ul class="list-none flex flex-col gap-y-2 mt-2 w-full">
            <li 
              v-for="(science, index) in sciences" 
              :key="index" 
              class="self-start flex flex-row items-center gap-x-2 px-2 cursor-default rounded-md border border-solid border-red-600">
              <div class="cursor-pointer" @click="removeScience(index)">
                <svg class="w-4 h-4" viewBox="0 0 14 14" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M10.5 3.5L3.50002 10.5M3.50002 3.5L10.5 10.5" stroke="#1E1E1E" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </div>

              <p>{{ science }}</p>
            </li>
          </ul>
        </section>
        <section class="flex flex-col">
          <label class="text-lg ml-2">Раздел науки</label>

          <div class="flex flex-row items-stretch">
            <input 
              type="text" 
              placeholder="Механика"
              v-model="newSciencePart"
              class="border border-solid rounded-l-lg outline-none px-2 border-gray-300 focus:border-red-800"/>
            <div 
              class="px-2 py-1 rounded-r-lg cursor-pointer btn"
              @click="addSciencePart">
              <svg class="w-6 h-6" viewBox="0 0 26 26" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M6.5 13H19.5M13 19.5V6.5" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            </div>
          </div>

          <ul class="list-none flex flex-col gap-y-2 mt-2 w-full">
            <li 
              v-for="(sciencePart, index) in scienceParts" 
              :key="index" 
              class="self-start flex flex-row items-center gap-x-2 px-2 cursor-default rounded-md border border-solid border-red-600">
              <div class="cursor-pointer" @click="removeSciencePart(index)">
                <svg class="w-4 h-4" viewBox="0 0 14 14" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M10.5 3.5L3.50002 10.5M3.50002 3.5L10.5 10.5" stroke="#1E1E1E" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </div>

              <p>{{ sciencePart }}</p>
            </li>
          </ul>
        </section>

        <div @click="filter" class="btn px-4 py-1 rounded-lg cursor-pointer">
          <p class="text-xl text-white select-none">Поиск</p>
        </div>
      </aside>

      <main class="flex flex-row items-start w-full h-full">
        <section class="flex flex-col w-1/2 gap-y-6 h-full py-4 border-r-2 border-solid border-gray-300">
          <h1 class=" text-center w-full text-2xl">Статьи пользователей</h1>

          <div class="flex flex-col w-full px-4 h-full scrollable" style="height: calc(100svh - 62px - 32px - 32px - 32px);">
            <div class="flex flex-col gap-y-4 w-full">
              <div v-for="article in filteredArticles" :key="article.id">
                <ArticleItem :title="article.title" :author="article.full_name" :science="article.science" :science-type="article.section"/>
              </div>
            </div>
          </div>
        </section>

        <section class="flex flex-col w-1/2 h-full py-4 gap-y-6">
          <h1 class=" text-center w-full text-2xl">Мои статьи</h1>

          <article v-if="userInfoStore.userID !== null" class="flex flex-col w-full px-4 h-full scrollable">
            <div v-if="myArticles.length !== 0" class="flex flex-col gap-y-4 w-full">
              <div v-for="myArticles in myArticles" :key="myArticles.id">
                <ArticleItem :title="myArticles.title" :author="myArticles.full_name" :science="myArticles.science" :science-type="myArticles.section"/>
              </div>
            </div>
            
            <div v-else class="flex flex-col items-center cursor-default">
              <div class="px-4 py-2 rounded-lg bg-gray-300">
                <p class="text-lg">У Вас пока что нет собственных статей</p>
              </div>
            </div>
          </article>

          <article v-else class="w-full items-center rounded-lg p-2 cursor-default border border-solid border-gray-300 bg-gray-100">
            <p class="w-full text-center text-xl ">Войдите, чтобы просматривать и добавлять свои статьи!</p>
          </article>

          <article @click="showUploadArticle()" class="self-end flex flex-row gap-x-2 items-center btn cursor-pointer mr-10 rounded-xl p-2">
            <div>
              <svg class="w-9 h-9" viewBox="0 0 57 57" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M30.5 15.5C30.5 14.3954 29.6046 13.5 28.5 13.5C27.3954 13.5 26.5 14.3954 26.5 15.5H30.5ZM26.5 41.5C26.5 42.6046 27.3954 43.5 28.5 43.5C29.6046 43.5 30.5 42.6046 30.5 41.5H26.5ZM41.5 30.5C42.6046 30.5 43.5 29.6046 43.5 28.5C43.5 27.3954 42.6046 26.5 41.5 26.5V30.5ZM15.5 26.5C14.3954 26.5 13.5 27.3954 13.5 28.5C13.5 29.6046 14.3954 30.5 15.5 30.5V26.5ZM52.5 28.5C52.5 41.7548 41.7548 52.5 28.5 52.5V56.5C43.964 56.5 56.5 43.964 56.5 28.5H52.5ZM28.5 52.5C15.2452 52.5 4.5 41.7548 4.5 28.5H0.5C0.5 43.964 13.036 56.5 28.5 56.5V52.5ZM4.5 28.5C4.5 15.2452 15.2452 4.5 28.5 4.5V0.5C13.036 0.5 0.5 13.036 0.5 28.5H4.5ZM28.5 4.5C41.7548 4.5 52.5 15.2452 52.5 28.5H56.5C56.5 13.036 43.964 0.5 28.5 0.5V4.5ZM26.5 15.5L26.5 28.5H30.5L30.5 15.5H26.5ZM26.5 28.5V41.5H30.5V28.5H26.5ZM41.5 26.5H28.5V30.5H41.5V26.5ZM28.5 26.5H15.5V30.5H28.5V26.5Z" fill="white"/>
              </svg>
            </div>
            <p class="text-xl text-white select-none">Загрузить статью</p>
            <input type="file" class=" appearance-none w-0 h-0 hidden" ref="uploadArticleFile"/>
          </article>

        </section>
      </main>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useBlurStore } from '@/stores/blurStore';
import { useArticleStore } from '@/stores/articleStore';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { API_Articles_Get, API_Article_Get_ByID } from '@/api/api';
import { StatusCodes, type Article } from '@/helpers/constants';

import HistoryFormulaItem from '@/shared/historyFormulaItem.vue';
import UpdateMyFormula from "@/entities/updateMyFormula.vue";
import UploadArticle from "@/entities/uploadArticle.vue";
import ArticleItem from '@/shared/articleItem.vue';

export default {
  components: {
    UploadArticle,
    UpdateMyFormula,
    ArticleItem,
    HistoryFormulaItem,
  },
  data(){
    return {
      articles: [] as Article[],
      myArticles: [] as Article[],
      filteredArticles: [] as Article[],
      filters: {
        articleTitle: '',
        author: '',
        science: '',
        section: ''
      },
      newScience: '', 
      sciences: [] as string[],
      newSciencePart: '',
      scienceParts: [] as string[],
      showUploadArticleMW: false,
    }
  },
  computed: {
    ...mapStores(useUserInfoStore, useBlurStore, useArticleStore, useStatusWindowStore),
  },
  mounted() {
    this.getArticles();
    this.getMyArticles();
  },
  methods: {
    hideArticleFormuls(){
      this.blurStore.showBlur = false;
      this.articleStore.showArticleFormulsMW = false;
    },
    async getArticles() {
      try {
        const response = await API_Articles_Get(); // Вызов API
        this.filteredArticles = response;
        this.articles = response;
      } catch (error) {
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при получении статей:');
      }
    },
    async getMyArticles() {
      try {
        const response = await API_Article_Get_ByID(this.userInfoStore.userID!); // Вызов API
        this.myArticles = response;
      } catch (error) {
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при получении Ваших статей:');
      }
    },
    showUploadArticle() {
      this.showUploadArticleMW = true;
      this.blurStore.showBlur = true;
    },
    hideUploadArticle() {
      this.showUploadArticleMW = false;
      this.blurStore.showBlur = false;
    },
    filter() {
      this.filteredArticles = this.articles.filter(article => {
        const matchesTitle = !this.filters.articleTitle || article.title.toLowerCase().includes(this.filters.articleTitle.toLowerCase());
        const matchesAuthor = !this.filters.author || article.full_name.toLowerCase().includes(this.filters.author.toLowerCase());
        const matchesScience = (this.sciences.length === 0 || this.sciences.some(science => article.science.toLowerCase().includes(science.toLowerCase())));
        const matchesSection = (this.scienceParts.length === 0 || this.scienceParts.some(part => article.section.toLowerCase().includes(part.toLowerCase())));

        return matchesTitle && matchesAuthor && matchesScience && matchesSection;
      });
    },
    addScience() {
      if (this.newScience.trim() !== '') {
        this.sciences.push(this.newScience.trim());
        this.newScience = '';
        this.filters.science = '';  
        this.filter();  
      }
    },
    removeScience(index: number) {
      this.sciences.splice(index, 1);
    },
    addSciencePart() {
      if (this.newSciencePart.trim() !== '') {
        this.scienceParts.push(this.newSciencePart.trim());
        this.newSciencePart = '';
        this.filters.section = '';  
        this.filter();  
      }
    },
    removeSciencePart(index: number) {
      this.scienceParts.splice(index, 1);
    }
  }
}
</script>