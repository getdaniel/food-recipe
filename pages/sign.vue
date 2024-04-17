<template>
    <div class="flex h-screen w-full items-center justify-center">
        <div class="w-full max-w-md bg-white rounded-lg shadow-md overflow-hidden">
            <div class="flex items-center border-b border-gray-200 px-4 py-3">
                <button :class="{ 'text-blue-600 font-bold': activeTab === 'signin' }" @click="setActiveTab('signin')">
                    Sign In
                </button>
                <button :class="{ 'text-blue-600 font-bold': activeTab === 'signup' }" @click="setActiveTab('signup')">
                    Sign Up
                </button>
            </div>
            <div v-if="activeTab === 'signin'" class="p-4">
                <form @submit.prevent="handleSignIn">
                    <div class="mb-3">
                        <label for="email" class="form-label block mb-2 text-sm font-medium text-gray-700">
                            Email
                        </label>
                        <input type="email" id="email"
                            class="form-control block w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                            v-model="signin.email" required />
                    </div>
                    <div class="mb-3">
                        <label for="password" class="form-label block mb-2 text-sm font-medium text-gray-700">
                            Password
                        </label>
                        <input type="password" id="password"
                            class="form-control block w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                            v-model="signin.password" required />
                    </div>
                    <button type="submit"
                        class="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
                        Sign In
                    </button>
                </form>
            </div>
            <div v-else class="p-4" style="display: none;">
                <form @submit.prevent="handleSignUp">
                    <div class="mb-3">
                        <label for="name" class="form-label block mb-2 text-sm font-medium text-gray-700">
                            Name
                        </label>
                        <input type="text" id="name"
                            class="form-control block w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                            v-model="signup.name" required />
                    </div>
                    <div class="mb-3">
                        <label for="email" class="form-label block mb-2 text-sm font-medium text-gray-700">
                            Email
                        </label>
                        <input type="email" id="email"
                            class="form-control block w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                            v-model="signup.email" required />
                    </div>
                    <div class="mb-3">
                        <label for="password" class="form-label block mb-2 text-sm font-medium text-gray-700">
                            Password
                        </label>
                        <input type="password" id="password"
                            class="form-control block w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                            v-model="signup.password" required />
                    </div>
                    <button type="submit"
                        class="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
                        Sign Up
                    </button>
                </form>
            </div>
        </div>
    </div>
</template>

<script setup>
// Define your reactive data properties
const activeTab = ref('signin'); // Initial active tab

const signin = ref({ email: '', password: '' });
const signup = ref({ name: '', email: '', password: '' });

// Define methods for handling actions
const setActiveTab = (newTab) => {
  activeTab.value = newTab; // Update the activeTab property
};

const handleSignIn = () => {
  // Implement your sign-in logic here (e.g., send data to backend API)
  console.log('Sign in data:', signin.value); // Example logging for demonstration
};

const handleSignUp = () => {
  // Implement your sign-up logic here (e.g., send data to backend API)
  console.log('Sign up data:', signup.value); // Example logging for demonstration
};
</script>
