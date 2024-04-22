<template>
  <div class="min-h-screen bg-gray-300">
    <NuxtLayout>
      <header-comp />
      <NuxtPage />
      <footer-comp />
    </NuxtLayout>
  </div>
</template>

<script>

import HeaderComp from '~/layouts/Header.vue'
import FooterComp from '~/layouts/Footer.vue'

export default {
  components: {
    HeaderComp,
    FooterComp,
  },
}
</script>