<template>
  <div v-if="!isEmpty" class="min-w-10 w-full h-10 flex flex-col justify-center items-center rounded text-lg p-1 button-style">
    <slot></slot>
  </div>
  <div v-else class="min-w-10 w-full h-10">

  </div>
</template>
<script lang="ts">
  export default {
    props: {
      isEmpty: {
        type: Boolean,
        require: false,
        default: false,
      }
    }
  };
</script>