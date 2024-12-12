<template>
  <!--блюр-->
  <div v-if="false" class="absolute w-svw h-svh opacity-45 bg-black z-20"></div>

  <StatusWindow />
  <Header />
  <AcceptCookie />
  <RouterView />
</template>
<script lang="ts">
import { mapStores } from "pinia";
import { useUserInfoStore } from "./stores/userInfoStore";

import StatusWindow from "./entities/statusWindow.vue";
import Header from "./entities/header.vue";
import AcceptCookie from "./widgets/acceptCookie.vue";
import { API_Health } from "./api/api";

export default {
  components: {
    StatusWindow,
    Header,
    AcceptCookie,
  },
  computed: {
    ...mapStores(useUserInfoStore),
  },
  async mounted() {
    //проверка авторизации
    // await this.userInfoStore.Authenticate();
    // API_Health();
  },
  methods: {

  },
};
</script>
