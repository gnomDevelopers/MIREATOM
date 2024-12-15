<template>
  <section class=" fixed inset-0 flex justify-center items-center z-30 p-4">
    <div class="flex flex-col gap-y-4 items-center z-30 p-4 rounded-xl min-w-[500px] bg-white relative">
      <div class="absolute right-0 top-0 cursor-pointer" @click="$emit('closeWindow')">
        <svg class="w-10 h-10" viewBox="0 0 52 52" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M34.6668 17.3334L17.3335 34.6667M34.6668 34.6667L17.3335 17.3334" stroke="#8F0101" stroke-width="2" stroke-linecap="round"/>
        </svg>
      </div>

      <h1 class="text-2xl text-color-theme cursor-default">Загрузить статью</h1>

      <article class="flex flex-col items-center gap-y-4 rounded-lg max-h-[500px] w-full py-2 px-4 bg-gray-100">

        <div class="flex flex-col gap-y-2 w-full py-2 px-4 rounded-lg bg-gray-200">
          <label class="text-lg cursor-pointer" for="articleTitle">Введите название статьи</label>
          <input
            type="text"
            id="articleTitle"
            placeholder="Название статьи"
            v-model="uploadArticleTitle"
            class="text-lg px-2 py-1 outline-none rounded border border-solid border-gray-300 focus:border-sky-500"/>
        </div>

        <div class="flex flex-col gap-y-2 w-full py-2 px-4 rounded-lg bg-gray-200">
          <label class="text-lg cursor-pointer" for="articleScience">Введите название науки для статьи</label>
          <input
            type="text"
            id="articleScience"
            placeholder="Наука"
            v-model="uploadArticleScience"
            class="text-lg px-2 py-1 outline-none rounded border border-solid border-gray-300 focus:border-sky-500"/>
        </div>

        <div class="flex flex-col gap-y-2 w-full py-2 px-4 rounded-lg bg-gray-200">
          <label class="text-lg cursor-pointer" for="articleSection">Введите название раздела науки для статьи</label>
          <input
            type="text"
            id="articleSection"
            placeholder="Раздел науки"
            v-model="uploadArticleSection"
            class="text-lg px-2 py-1 outline-none rounded border border-solid border-gray-300 focus:border-sky-500"/>
        </div>

        <article @click="() => {$refs.uploadArticleFile.click()}" class="self-end flex flex-row gap-x-2 items-center btn cursor-pointer mr-10 rounded-xl p-2">
          <div>
            <svg class="w-9 h-9" viewBox="0 0 57 57" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M30.5 15.5C30.5 14.3954 29.6046 13.5 28.5 13.5C27.3954 13.5 26.5 14.3954 26.5 15.5H30.5ZM26.5 41.5C26.5 42.6046 27.3954 43.5 28.5 43.5C29.6046 43.5 30.5 42.6046 30.5 41.5H26.5ZM41.5 30.5C42.6046 30.5 43.5 29.6046 43.5 28.5C43.5 27.3954 42.6046 26.5 41.5 26.5V30.5ZM15.5 26.5C14.3954 26.5 13.5 27.3954 13.5 28.5C13.5 29.6046 14.3954 30.5 15.5 30.5V26.5ZM52.5 28.5C52.5 41.7548 41.7548 52.5 28.5 52.5V56.5C43.964 56.5 56.5 43.964 56.5 28.5H52.5ZM28.5 52.5C15.2452 52.5 4.5 41.7548 4.5 28.5H0.5C0.5 43.964 13.036 56.5 28.5 56.5V52.5ZM4.5 28.5C4.5 15.2452 15.2452 4.5 28.5 4.5V0.5C13.036 0.5 0.5 13.036 0.5 28.5H4.5ZM28.5 4.5C41.7548 4.5 52.5 15.2452 52.5 28.5H56.5C56.5 13.036 43.964 0.5 28.5 0.5V4.5ZM26.5 15.5L26.5 28.5H30.5L30.5 15.5H26.5ZM26.5 28.5V41.5H30.5V28.5H26.5ZM41.5 26.5H28.5V30.5H41.5V26.5ZM28.5 26.5H15.5V30.5H28.5V26.5Z" fill="white"/>
            </svg>
          </div>
          <p class="text-xl text-white select-none">Загрузить файл (.docx или .tex)</p>
          <input
            type="file"
            class="appearance-none w-0 h-0 hidden"
            ref="uploadArticleFile"
            accept=".tex, .docx"
            @change="handleFileSelection"
          />
        </article>

        <div @click="uploadArticle" class="flex flex-row gap-x-2 items-center btn cursor-pointer rounded-lg px-2 py-1 mt-2">
          <svg class="w-9 h-9" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M34 42V26H14V42M14 6V16H30M38 42H10C8.93913 42 7.92172 41.5786 7.17157 40.8284C6.42143 40.0783 6 39.0609 6 38V10C6 8.93913 6.42143 7.92172 7.17157 7.17157C7.92172 6.42143 8.93913 6 10 6H32L42 16V38C42 39.0609 41.5786 40.0783 40.8284 40.8284C40.0783 41.5786 39.0609 42 38 42Z" stroke="white" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
          <p class="text-xl text-white">Загрузить статью</p>
        </div>
      </article>
    </div>

  </section>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useFormulsStore } from '@/stores/formulsStore';
import { StatusCodes } from '@/helpers/constants';
import { API_Article_Post } from '@/api/api';

export default {
  emits: ['closeWindow', 'updateArticles'],
  data() {
    return {
      uploadArticleTitle: '',
      uploadArticleScience: '',
      uploadArticleSection: '',
      uploadArticleFile: null as File | null, // Для файла
      showUploadArticleMW: false, // Для модального окна
    };
  },
  computed: {
    ...mapStores(useStatusWindowStore, useFormulsStore),
  },

  methods: {
    handleFileSelection(event: Event) {
      //очевидно
      const input = event.target as HTMLInputElement;
      if (input.files && input.files[0]) {
        if(!input.files[0].name.endsWith('.tex') && !input.files[0].name.endsWith('.docx')){
          this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Неверное расширение файла!');
          return;
        }
        this.uploadArticleFile = input.files[0]; // Сохраняем выбранный файл
      }

    },
    uploadArticle() {
      //проверка на пустое название статьи
      if(this.uploadArticleTitle === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите название статьи!');
        return;
      }
      //проверка на пустоту науки
      if(this.uploadArticleScience === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите название науки!');
        return;
      }
      //проверка на пустоту раздела науки
      if(this.uploadArticleSection === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите название раздела науки!');
        return;
      }
      //проверка на добавление файла со статьей
      if (!this.uploadArticleFile) {
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Загрузите файл со статьей!');
        return;
      }

      //создаем объект formData для отправки файла стратьи
      const formData = new FormData();
      formData.append('title', this.uploadArticleTitle);
      formData.append('science', this.uploadArticleScience);
      formData.append('section', this.uploadArticleSection);
      formData.append('file', this.uploadArticleFile, this.uploadArticleFile.name);

      //выводим окно отправки
      const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Добавляем статью...', -1);

      //запрос на изменение
      API_Article_Post(formData)
      .then((response: any) => {
        //если все ок - выводим сообщение что все ок
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Статья добавлена!');
        //добавляем статью
        this.$emit('updateArticles');
        //закрываем окно создания статьи
        this.$emit('closeWindow');
      })
      .catch(error => {
        //если что-то не так - сообщаем об ошибке
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при добавлении статьи!');
      });
    }
  }
};
</script>