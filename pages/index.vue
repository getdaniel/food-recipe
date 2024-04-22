<template>
  <div class="text-center min-h-screen mt-10">
    <p v-if="data?.users?.length === 0">No users available.</p>
    <div v-else>
      <p v-for="user in data?.users" :key="user.id">
        Username: {{ user.username }}, <br /> Email: {{ user.email }}
      </p>
    </div>
  </div>
</template>

<script lang="ts" setup>

interface User {
  id: string;
  username: string;
  email: string;
}

interface UserData {
  users: User[];
}

const query = gql`
  query MyQuery {
    users {
      id
      username
      email
    }
  }
`;

const { data } = await useAsyncQuery<UserData>(query);
</script>