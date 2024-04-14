<!-- The default layout (header, footer included here) -->
<template>
  <div>
    <TheHeader />
    <div class="min-h-screen">
      <TheSearch />
      <RecipeCard />
    </div>
    <TheFooter />
  </div>
</template>

<script setup>
import TheHeader from '~/components/TheHeader.vue'
import TheSearch from '~/components/SearchBar.vue'
import TheFooter from '~/components/TheFooter.vue'
import RecipeCard from '~/components/RecipeCard.vue'
</script>