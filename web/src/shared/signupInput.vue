<template>
  <div class="flex flex-col items-start relative">
    <label
        @click="() => {$refs.signupInput.focus()}"
        class="absolute transition-all block text-lg font-bold mb-2"
        :class="{ 'login-label': !focusLabel, 'login-label-focus': focusLabel }"
    >{{ text }}</label>

    <div
        class="flex flex-row items-end gap-x-2 w-full h-16 transition-input hover:border-sky-500"
        :class="{ 'border-sky-500': isFocused, 'border-slate-500': !isFocused, 'border-red-600': error}"
    >
      <div class=" flex flex-row justify-between shadow appearance-none rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none ">
        <input
            v-if="type === 'password'"
            placeholder="your_password"
            class="bg-transparent outline-none w-11/12"
            :type="inputType"
            v-model="inputValue"
            @focusin="onFocus"
            @focusout="unFocus"
            ref="signupInput"
        />
        <input
            v-if="type === 'login'"
            placeholder="email@example.com"
            class="bg-transparent outline-none w-11/12"
            :type="inputType"
            v-model="inputValue"
            @focusin="onFocus"
            @focusout="unFocus"
            ref="signupInput"
        />
        <input
            v-if="type === 'name'"
            placeholder="Иванов Иван Иванович"
            class="bg-transparent outline-none w-11/12"
            :type="inputType"
            v-model="inputValue"
            @focusin="onFocus"
            @focusout="unFocus"
            ref="signupInput"
        />

        <input
            v-if="type === 'repPassword'"
            placeholder="your_password"
            class="bg-transparent outline-none w-11/12"
            :type="inputType"
            v-model="inputValue"
            @focusin="onFocus"
            @focusout="unFocus"
            ref="signupInput"
        />

        <div
            v-if="type === 'password' || type === 'repPassword'"
            @mousedown="onHold"
            @mouseup="unHold"
            @touchstart="onHold"
            @touchend="unHold"
            @dragend="unHold"
        >
          <img
              v-if="inputType === 'password'"
              src="../assets/icons/icon-password-hide.svg"
              alt="Показать пароль"
              class="w-6 h-6 cursor-pointer mr-1 float-end"
          />
          <img
              v-else
              src="../assets/icons/icon-password-show.svg"
              alt="Скрыть пароль"
              class="w-6 h-6 cursor-pointer mr-1 float-end"
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
    }
  },
  emits: ['inputChange'],
  data() {
    return {
      inputType: this.type === 'password' || this.type === 'repPassword' ? 'password' : this.type, // Установка начального состояния
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
      if (this.type === 'password' || this.type === 'repPassword') {
        this.inputType = 'text';
      }
    },
    unHold() {
      if (this.type === 'password' || this.type === 'repPassword') {
        this.inputType = 'password';
      }
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