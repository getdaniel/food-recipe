<template>
  <header class="bg-green-600 text-white p-4">
    <div class="container mx-auto flex items-center justify-between">

      <!-- Logo and Title -->
      <div class="flex items-center">
        <div class="rounded-full overflow-hidden mr-4">
          <img src="../statics/logo.png" alt="Logo" class="h-12 w-12">
        </div>
      </div>

      <!-- Navigation Menu -->
      <nav class="md:flex space-x-4 mt-4 md:mt-0">
        <ul class="flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4">
          <!-- Main Navigation Links -->
          <li><a href="#" class="hover:text-gray-300">Home</a></li>
          <li><a href="#" class="hover:text-gray-300">Browse Recipes</a></li>
          <li><a href="#" class="hover:text-gray-300">Categories</a></li>
          <li><a href="#" class="hover:text-gray-300">Sign Up</a></li>
          <li><a href="#" class="hover:text-gray-300">Log In</a></li>
        </ul>
      </nav>

      <!-- Navigation Toggle Button -->
      <button class="block md:hidden focus:outline-none">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7"></path>
        </svg>
      </button>
    </div>
  </header>
</template>