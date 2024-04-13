<template>
  <div class="bg-sky-950 text-gray-100 py-1.5 px-6 shadow md:flex justify-between items-center">

    <!-- Logo and Title -->
    <div class="flex items-center">
      <span class="text-green-500 text-xl mr-1">
        <div class="rounded-full overflow-hidden mr-4">
          <img src="../statics/logo.png" alt="Logo" class="h-10 w-14">
        </div>
      </span>
      <h1 class="text-center font-bold text-gray-300">
        <span class="text-green-500 mb-0 text-xs font-serif">Food</span> <br>
        <span class="italic mt-0">Recipes</span>
      </h1>

    </div>

    <!-- Navigation Toggle Button -->
    <button class="absolute md:hidden right-6 top-4 cursor-pointer" @click="open = !open">
      <svg v-if="open" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
      </svg>
      <svg v-else class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7"></path>
      </svg>
    </button>

    <!-- Navigation Menu -->
    <nav>
      <ul
        class="md:flex md:items-center md:px-0 px-3 md:pb-0 pb-10 md:static absolute bg-sky-950 md:w-auto w-full top-14 duration-700 ease-in"
        :class="[open ? 'left-0' : 'left-[-100%]']">
        <!-- Main Navigation Links -->
        <li class="md:mx-4 md:my-0 my-6" v-for="link in Links">
          <a href={{ link.link }} class="text-xl hover:text-green-500">{{ link.name }}</a>
        </li>
        <BaseButton> Get Started </BaseButton>
      </ul>
    </nav>

  </div>
</template>

<script>

import BaseButton from './BaseButton';

export default {

  components: {
    BaseButton
  },

  setup() {
    let open = ref(false);
    let Links = [
      { name: "Home", link: "#" },
      { name: "Recipes", link: "#" },
      { name: "Categories", link: "#" },
    ]

    return { Links, open };
  }
}

</script>