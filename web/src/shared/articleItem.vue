<template>
  <article class="grid grid-cols-12 w-full items-center rounded-lg px-2 py-1 cursor-default border border-solid border-gray-300 bg-gray-100">
    <div class="flex flex-col gap-y-1 items-start col-span-7">
      <p class="text-xl">{{ title }}</p>
      <p class="text-gray-700">Автор: {{ author }}</p>
    </div>
    <div class="flex flex-col gap-y-1 items-center col-span-3">
      <p class="text-lg font-semibold">{{ science }}</p>
      <p class="text-gray-700">{{ scienceType }}</p>
    </div>
    <div class="flex flex-row justify-end gap-x-2 items-center col-span-2">
      <div @click="showArticleFormuls" class="flex flex-col items-center justify-center p-2 btn cursor-pointer rounded-lg" title="Посмотреть все формулы в статье">
        <svg class="w-7 h-7" viewBox="0 0 35 35" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M7.93285 9.78963L6.76357 11.0999L6.87849 11.5955H9.88931C9.27451 15.1217 8.81484 18.3004 7.93285 23.1995C6.87849 29.4487 6.2723 30.9839 5.91318 31.545C5.59716 32.1004 5.14324 32.3881 4.52843 32.3881C3.8533 32.3881 2.6323 31.8583 1.97727 31.2573C1.74169 31.0978 1.46302 31.1377 1.16711 31.3314C0.575286 31.827 0.000700559 32.5903 0.000700559 33.2739C-0.0337746 34.1968 1.2102 35 2.39672 35C3.44247 35 4.98236 34.3506 6.60269 32.8239C8.81771 30.7418 10.4754 27.8849 11.8601 21.7098C12.7507 17.7677 13.1702 15.2043 13.7965 11.5983L17.537 11.2622L18.3472 9.78963H14.1355C15.2243 2.99927 16.0747 1.99382 17.1205 1.99382C17.8157 1.99382 18.6259 2.5236 19.5567 3.56608C19.8325 3.94206 20.2462 3.90218 20.5622 3.64583C21.0937 3.33252 21.7488 2.5293 21.7918 1.80013C21.8235 0.996908 20.8581 0 19.1976 0C17.6922 0 15.3852 0.996908 13.446 2.95654C11.7452 4.74243 10.8546 6.97266 10.2398 9.78963H7.93285ZM19.9015 14.3355C21.0621 12.8031 21.7574 12.2933 22.1165 12.2933C22.4871 12.2933 22.783 12.655 23.4179 14.6887L24.5039 18.1608C22.4038 21.3452 20.8668 23.1197 19.9331 23.1197C19.6228 23.1197 19.3096 23.0229 19.0884 22.8064C18.8701 22.5899 18.6288 22.4105 18.3932 22.4105C17.6232 22.4105 16.6694 23.3362 16.6522 24.4841C16.6349 25.6547 17.4623 26.4921 18.5483 26.4921C20.4128 26.4921 21.9901 24.6777 25.0354 19.8669L25.926 22.8434C26.6902 25.4012 27.5865 26.4921 28.707 26.4921C30.0917 26.4921 31.9534 25.3215 33.9846 22.157L33.1342 21.1943C31.9132 22.6526 31.103 23.3362 30.6232 23.3362C30.0889 23.3362 29.6206 22.5273 29.0029 20.5391L27.7015 16.3436C28.4714 15.2185 29.2299 14.2729 29.8964 13.5181C30.6893 12.6209 31.3012 12.2392 31.7178 12.2392C32.0683 12.2392 32.3728 12.3873 32.5625 12.5952C32.8124 12.8687 32.9618 12.9968 33.2577 12.9968C33.9271 12.9968 34.9613 12.1537 34.9987 11.0372C35.0332 10.0004 34.3839 9.11743 33.2577 9.11743C31.5569 9.11743 30.0573 10.5701 27.17 14.8254L26.5753 12.9997C25.7421 10.4362 25.1905 9.11743 24.0299 9.11743C22.6796 9.11743 20.8237 10.7638 19.008 13.3728L19.9015 14.3355Z" fill="white"/>
        </svg>
      </div>
      <div class="flex flex-col items-center justify-center p-2 btn cursor-pointer rounded-lg">
        <svg class="w-7 h-7" viewBox="0 0 35 35" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M30.625 21.875V27.7083C30.625 28.4819 30.3177 29.2237 29.7707 29.7707C29.2237 30.3177 28.4819 30.625 27.7083 30.625H7.29167C6.51812 30.625 5.77625 30.3177 5.22927 29.7707C4.68229 29.2237 4.375 28.4819 4.375 27.7083V21.875M10.2083 14.5833L17.5 21.875M17.5 21.875L24.7917 14.5833M17.5 21.875V4.375" stroke="white" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      </div>
    </div>
  </article>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useArticleStore } from '@/stores/articleStore';
import { useBlurStore } from '@/stores/blurStore';

export default {
  props: {
    title: {
      type: String,
      required: true,
    },
    author: {
      type: String,
      required: true,
    },
    science: {
      type: String,
      required: true,
    },
    scienceType: {
      type: String,
      required: true,
    },
  },
  computed: {
    ...mapStores(useArticleStore, useBlurStore),
  },
  methods: {
    showArticleFormuls(){
      this.blurStore.showBlur = true;
      this.articleStore.showArticleFormulsMW = true;
    }
  }
};
</script>