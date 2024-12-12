<template>
  <div class="flex flex-col items-start gap-y-1">
    <label
      @click="() => {$refs.loginInput.focus()}"
      class="text-lg font-bold ml-2"
      :class="{ 'login-label': !focusLabel, 'login-label-focus': focusLabel }"
    >{{ text }}</label>

    <div
      class="flex flex-row items-end gap-x-2 w-full transition-input hover:border-sky-500"
      :class="{ 'border-sky-500': isFocused, 'border-slate-500': !isFocused, 'border-red-600': error}"
    >
      <div class=" flex flex-row justify-between header-shadow rounded w-full py-2 px-3 text-gray-700 leading-tight">
        <input
          :type="inputType"
          class="bg-transparent outline-none w-11/12"
          :placeholder
          v-model="inputValue"
          @focusin="onFocus"
          @focusout="unFocus"
          ref="loginInput"
        />

        <div
          v-if="type === 'password'"
          @mousedown="onHold"
          @mouseup="unHold"
          @touchstart="onHold"
          @touchend="unHold"
          @dragend="unHold"
        >
          <img
            v-if="inputType !== type"
            src="../assets/icons/icon-password-hide.svg"
            alt="Показать пароль"
            class="w-6 h-6 cursor-pointer mr-1 float-end"
          />
          <img
            v-else
            src="../assets/icons/icon-password-show.svg"
            class="w-6 h-6 cursor-pointer mr-1"
          />
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
export default {
  props: {
    type: {
      type: String,
      required: true,
    },
    text: {
      type: String,
      required: true,
    },
    error: {
      type: Boolean,
      required: false,
      default: false,
    },
    placeholder: {
      type: String,
      reuired: false,
      default: '',
    }
  },
  emits: ['inputChange'],
  data() {
    return {
      inputType: this.type,
      isFocused: false,
      inputValue: "",
    };
  },
  computed: {
    focusLabel() {
      return this.isFocused || this.inputValue !== "";
    },
  },
  methods: {
    onHold() {
      this.inputType = "text";
    },
    unHold() {
      this.inputType = this.type;
    },
    onFocus() {
      this.isFocused = true;
    },
    unFocus() {
      this.isFocused = false;
    },
  },
  watch:{
    'inputValue'(val){
      this.$emit('inputChange', val);
    }
  }
};
</script>