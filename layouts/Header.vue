<template>
  <div class="bg-sky-950 text-gray-100 py-1.5 px-6 shadow md:flex justify-between items-center">
    <NuxtLink to="/">
      <div class="flex items-center">
        <span class="text-green-500 text-xl mr-1">
          <div class="rounded-full overflow-hidden mr-4">
            <img src="../assets/images/logo.png" alt="Logo" class="h-10 w-14">
          </div>
        </span>
        <h1 class="text-center font-bold text-gray-300">
          <span class="text-green-500 mb-0 text-xs font-serif">Food</span> <br>
          <span class="italic mt-0">Recipes</span>
        </h1>
      </div>
    </NuxtLink>

    <button class="absolute md:hidden right-6 top-4 cursor-pointer" @click="open = !open">
      <Icon v-if="open" name="X" class="w-6 h-6" />
      <Icon v-else name="☰" class="w-6 h-6" />
    </button>

    <nav>
      <ul class="md:flex md:items-center md:px-0 px-3 md:pb-0 pb-10 md:static absolute bg-sky-950 md:w-auto w-full top-14 duration-700 ease-in" :class="[open ? 'left-0' : 'left-[-100%]']">
        <li class="md:mx-4 md:my-0 my-6" v-for="link in Links">
          <NuxtLink :to="link.link" class="text-xl hover:text-green-500" :class="{ 'text-teal-400': $route.path === link.link }">{{ link.name }}</NuxtLink>
        </li>

        <NuxtLink to="/sign">
          <button class="hover: bg-green-500 duration: delay-300 font-sans rounded py-1.5 px-4 text-white">Get Started</button>
        </NuxtLink>
      </ul>
    </nav>
  </div>
</template>

<script>
export default {
  setup() {
    let open = ref(false);
    let Links = [
      { name: "Home", link: '/' },
      { name: "Recipes", link: '/recipes' },
      { name: "About Us", link: '/about-us' },
    ]

    return { open, Links };
  }
}
</script>

<style>
.text-special-color {
  color: #FF6347; /* Or any other color you prefer */
}
</style>
