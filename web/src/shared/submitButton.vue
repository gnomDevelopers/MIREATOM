<template>
  <input
    class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-xl font-semibold text-slate-100"
    type="button"
    :value
    @click.prevent="$emit('submit')"
  />
</template>
<script lang="ts">
export default {
  props: {
    value: {
      type: String,
      required: true,
    },
  },
  emits: ["submit"],
};
</script>
