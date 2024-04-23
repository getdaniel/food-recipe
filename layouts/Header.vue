<template>
  <div class="bg-sky-950 text-gray-100 py-1.5 px-6 shadow md:flex justify-between items-center relative">
    <!-- Ensure relative positioning here for proper stacking context -->
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
      <ul
        class="md:flex md:items-center md:px-0 px-3 md:pb-0 pb-10 md:static absolute bg-sky-950 md:w-auto w-full top-14 duration-700 ease-in z-10"
        :class="[open ? 'left-0' : 'left-[-100%]']">
        <li class="md:mx-4 md:my-0 my-6" v-for="link in Links">
          <NuxtLink :to="link.link" class="text-xl hover:text-green-500"
            :class="{ 'text-teal-400': $route.path === link.link }">{{ link.name }}</NuxtLink>
        </li>

        <!-- Conditional rendering based on user authentication -->
        <template v-if="isAuthenticated">
          <li class="relative">
            <div @click="toggleProfileMenu"
              class="rounded-full overflow-hidden w-10 h-10 border-2 border-green-500 cursor-pointer profile-menu-trigger">
              <img src="~/assets/images/me.png" alt="Profile" class="w-full h-full object-cover profile-menu-trigger">
            </div>
            
            <!-- Profile menu: My Profile and Logout -->
            <ul v-if="profileMenuOpen"
              class="absolute right-0 mt-2 bg-sky-950 shadow-lg rounded-xl text-grey-100 w-28 py-1 z-50">
              <li>
                <NuxtLink to="/profile" class="block px-4 py-2 hover:bg-teal-400" @click="closeProfileMenu">My Profile
                </NuxtLink>
              </li>
              <li>
                <button @click="logout" class="block w-full text-left px-4 py-2 hover:bg-red-500">Logout</button>
              </li>
            </ul>
          </li>
        </template>
        <template v-else>
          <!-- Get Started Button -->
          <li>
            <NuxtLink to="/sign">
              <button class="hover:bg-green-500 duration-300 font-sans rounded py-1.5 px-4 text-white">Get Started
              </button>
            </NuxtLink>
          </li>
        </template>
      </ul>
    </nav>
  </div>
</template>

<script>
export default {
  setup() {
    const open = ref(false);
    const profileMenuOpen = ref(false);
    const isAuthenticated = ref(true);

    let Links = [
      { name: "Home", link: '/' },
      { name: "Recipes", link: '/recipes' },
      { name: "About Us", link: '/about-us' },
    ]

    const logout = () => {
      isAuthenticated.value = false;
      profileMenuOpen.value = false;
    }

    const toggleProfileMenu = () => {
      profileMenuOpen.value = !profileMenuOpen.value;
    }

    const closeProfileMenu = () => {
      profileMenuOpen.value = false;
    };

    const handleClickOutside = (event) => {
      if (!event.target.closest('.profile-menu-trigger')) {
        closeProfileMenu();
      }
    };

    onMounted(() => {
      document.addEventListener('click', handleClickOutside);
    });

    onUnmounted(() => {
      document.removeEventListener('click', handleClickOutside);
    });

    return { open, Links, profileMenuOpen, isAuthenticated, logout, toggleProfileMenu, closeProfileMenu };
  }
}
</script>