<template>
    <div class="text-center min-h-screen">
        Hello... I am comming soon!
        <p>{{ data }}</p>
    </div>
</template>

<script lang="ts" setup> 
const query = gql`
query {
  users {
    id
    username
    email
  }
}`

const { data } = await useAsyncQuery(query)

</script>