<template>
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
    await this.userInfoStore.Authenticate();
  },
  methods: {

  },
};
</script>
