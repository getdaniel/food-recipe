<template>
    <div class="flex h-screen w-full items-center justify-center">
        <div class="w-full max-w-md bg-white rounded-lg shadow-md overflow-hidden">
            <div class="flex items-center border-b border-gray-200 px-4 py-3">
                <div class="w-full flex justify-between">
                    <button :class="{ 'text-blue-600 font-bold': activeTab === 'signin' }"
                        @click="setActiveTab('signin')" class="w-1/2 px-2 py-1 bg-gray-200 hover:bg-gray-300">
                        Sign In
                    </button>
                    <button :class="{ 'text-blue-600 font-bold': activeTab === 'signup' }"
                        @click="setActiveTab('signup')" class="w-1/2 px-2 py-1 bg-gray-200 hover:bg-gray-300">
                        Sign Up
                    </button>
                </div>
            </div>

            <div class="p-4">
                <VeeForm @submit.prevent="handleSignIn" v-show="activeTab === 'signin'" ref="signInForm">
                    <div class="mb-3">
                        <label for="email" class="form-label block mb-2 text-sm font-medium text-gray-700">
                            Email
                        </label>
                        <VeeField name="email" :rules="{ required: true}">
                            <input type="email" id="email" placeholder="Enter your email" v-model="signin.email"
                                class="form-control block w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500" />
                                <span v-error> {{ error }} </span>
                        </VeeField>
                    </div>
                    <div class="mb-3">
                        <label for="password" class="form-label block mb-2 text-sm font-medium text-gray-700">
                            Password
                        </label>
                        <VeeField name="password" :rules="{ required: true, min: 6 }">
                            <input type="password" id="password" placeholder="Enter your password"
                                v-model="signin.password"
                                class="form-control block w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500" />
                                <span v-error> {{ error }} </span>
                        </VeeField>
                    </div>
                    <div class="flex items-center justify-center">
                        <button type="submit" class="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">Sign In</button>
                    </div>
                </VeeForm>
                <VeeForm @submit.prevent="handleSignUp" v-show="activeTab === 'signup'" ref="signUpForm">
                    <div class="mb-3">
                        <label for="name" class="form-label block mb-2 text-sm font-medium text-gray-700">
                            Name
                        </label>
                        <VeeField name="name" :rules="{ required: true }">
                            <input type="text" id="name" placeholder="Enter your full name" v-model="signup.name"
                                class="form-control block w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500" />
                                <span v-error> {{ error }} </span>
                        </VeeField>
                    </div>
                    <div class="mb-3">
                        <label for="email" class="form-label block mb-2 text-sm font-medium text-gray-700">
                            Email
                        </label>
                        <VeeField name="email" :rules="{ required: true, email: true }">
                            <input type="email" id="email" placeholder="Enter your email" v-model="signup.email"
                                class="form-control block w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500" />
                                <span v-error> {{ error }} </span>
                        </VeeField>
                    </div>
                    <div class="mb-3">
                        <label for="password" class="form-label block mb-2 text-sm font-medium text-gray-700">
                            Password
                        </label>
                        <VeeField name="password" :rules="{ required: true, min: 6 }">
                            <input type="password" id="password" placeholder="Enter your password"
                                v-model="signup.password"
                                class="form-control block w-full px-3 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500" />
                                <span v-error> {{ error }} </span>
                        </VeeField>
                    </div>
                    <div class="mb-3">
                        <label for="confirmPassword" class="block text-sm font-medium text-gray-700">Confirm
                            Password</label>
                        <VeeField name="confirmPassword" :rules="{ required: true, same: signup.password }"> <input
                                type="password" id="confirmPassword" placeholder="Confirm your password"
                                v-model="signup.confirmPassword"
                                class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500" />
                                <span v-error> {{ error }} </span>
                        </VeeField>
                    </div>

                    <div class="flex items-center justify-center">
                        <button type="submit" class="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">Sign Up</button>
                    </div>
                </VeeForm>
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
