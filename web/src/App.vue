<template>
  <!--блюр-->
  <Transition>
    <div v-if="blurStore.showBlur" class="absolute w-svw h-svh z-20" style="background-color:  rgba(0, 0, 0, 0.45);"></div>
  </Transition>

  <StatusWindow />
  <Header />
  <AcceptCookie />
  <RouterView />
</template>
<script lang="ts">
import { mapStores } from "pinia";
import { useUserInfoStore } from "./stores/userInfoStore";
import { useBlurStore } from "./stores/blurStore";

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
    ...mapStores(useUserInfoStore, useBlurStore),
  },
  async mounted() {
    // проверка авторизации
    await this.userInfoStore.Authenticate();
  },
  methods: {

  },
};
</script>
<style v-global>
  .v-enter-active, .v-leave-active {
    transition: opacity 150ms ease-in-out;
  }
  .v-enter-from, .v-leave-to {
    opacity: 0;
  }
  .v-enter-to, .v-leave-from {
    opacity: 1;
  }

  .list-enter-active,
  .list-leave-active {
    transition: all 0.5s ease-out;
  }
  .list-enter-from,
  .list-leave-to {
    opacity: 0;
    transform: translateY(-10px);
  }

</style>