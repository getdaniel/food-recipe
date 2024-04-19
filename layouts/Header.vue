<template>

  <head>
    <title>Food Recipes</title>
  </head>
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
      <Icon v-if="open" name="x" class="w-6 h-6" />
      <Icon v-else name="menu-deep" class="w-6 h-6" />
    </button>

    <nav>
      <ul
        class="md:flex md:items-center md:px-0 px-3 md:pb-0 pb-10 md:static absolute bg-sky-950 md:w-auto w-full top-14 duration-700 ease-in"
        :class="[open ? 'left-0' : 'left-[-100%]']">
        <li class="md:mx-4 md:my-0 my-6" v-for="link in Links">
          <NuxtLink :to="link.link" class="text-xl hover:text-green-500">{{ link.name }}</NuxtLink>
        </li>

        <li class="md:mx-4 md:my-0 my-6 relative">
          <NuxtLink to="/recipes" class="text-xl hover:text-green-500">
            Recipes
          </NuxtLink>
        </li>

        <NuxtLink to="/about-us">
          <li class="md:mx-4 md:my-0 my-6">About Us</li>
        </NuxtLink>
        <NuxtLink to="/sign">
          <button class="hover: bg-green-500 duration: delay-300 font-sans rounded py-1.5 px-4 text-white">Get
            Started</button>
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
      // Removed Recipes link as it's now part of the dropdown
    ]

    return { Links, open };
  }
}

</script>