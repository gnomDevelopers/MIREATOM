<template>
  <div class="flex flex-row items-stretch h-14 w-full px-5 header-shadow z-10">
    <div class="h-full grid content-center mx-16">
      <div class="font-semibold text-3xl cursor-pointer">SITELOGO</div>
    </div>
    <div class="h-full flex flex-row gap-x-2 flex-grow">
      <div
        class="h-full min-w-16 grid cursor-pointer justify-center content-center border-bottom border-gray-500 hover:border-b-sky-500"
      >
        <a>Home</a>
      </div>
      <div
        class="h-full min-w-16 grid cursor-pointer justify-center content-center border-bottom border-gray-500 hover:border-b-sky-500"
      >
        <a>About</a>
      </div>
      <div
        class="h-full min-w-16 grid cursor-pointer justify-center content-center border-bottom border-gray-500 hover:border-b-sky-500"
      >
        <a>Main</a>
      </div>
      <div
        class="h-full min-w-16 grid cursor-pointer justify-center content-center border-bottom border-gray-500 hover:border-b-sky-500"
      >
        <a>F.A.Q.</a>
      </div>
    </div>
    <div class="flex flex-row gap-x-2 items-center mr-16">
      <img
        class="rounded-full w-10 h-10"
        src="../assets/icons/user-avatar.png"
      />
      <div class="cursor-pointer">UserName</div>
    </div>
  </div>
</template>
<script lang="ts">
export default {};
</script>
