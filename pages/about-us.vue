<template>
    <div class="min-h-screen py-10">
        <!-- Welcome Text -->
        <div class="text-center px-12 mb-6 lg:mb-0 lg:mr-6">
            <h1 class="text-xl font-bold text-blue-700 mb-6 underline">Welcome to Food Recipes Site!</h1>
            <p class="text-lg text-gray-800 leading-relaxed mb-6">Food Recipes is a project aimed at providing a
                platform for food enthusiasts to explore and share delicious
                recipes from around the world. Our goal is to inspire people to cook and experiment with new flavors,
                ingredients, and techniques.</p>
            <p class="text-lg text-gray-800 leading-relaxed mb-6">We are proud to collaborate with:</p>

            <!-- Logos -->
            <div class="flex justify-center items-center mb-6">
                <div class="mr-4">
                    <img src="../assets/images/logo.png" alt="Food Recipes Logo" class="h-28 w-28 rounded-full">
                </div>
                <div class="mr-4">
                    <img src="../assets/images/minab.jpeg" alt="Minab IT Solution Logo" class="h-28 w-28 rounded-full">
                </div>
            </div>

            <p class="text-lg text-gray-800 leading-relaxed mb-6">Food Recipes is proud to collaborate with <span
                    class="italic font-semibold">
                    <NuxtLink to="https://www.minabtech.com/" target="_blank">Minab IT Solution</NuxtLink>
                </span> on this internship project. Minab IT Solution
                is a leading
                technology company dedicated to delivering innovative solutions to its clients.</p>
            <div class="border-t border-gray-300 my-6"></div>
            <p class="text-lg text-gray-800 leading-relaxed mb-6">If you have any questions, feedback, or suggestions,
                please feel free to contact us:</p>
        </div>

        <!-- Get in Touch and Contact Form Container -->
        <div class="flex flex-col lg:flex-row px-4 lg:px-10">
            <!-- Get in Touch -->
            <div class="bg-gray-200 p-6 rounded-lg shadow-md mb-6 lg:mb-0 lg:mr-4 lg:w-1/2">
                <h2 class="text-lg font-bold mb-4 text-center">Get in Touch</h2>
                <div class="mb-2">
                    <p class="text-gray-800 mb-1"><strong>Phone:</strong></p>
                    <p class="text-gray-600">+251935976224</p>
                </div>
                <div class="mb-2">
                    <p class="text-gray-800 mb-1"><strong>Email:</strong></p>
                    <p class="text-gray-600">danielgetaneh2011@gmail.com</p>
                </div>
                <div class="mb-2">
                    <p class="text-gray-800 mb-1"><strong>Location:</strong></p>
                    <!-- You can embed Google Maps here -->
                    <!-- For example: -->
                    <iframe
                        src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d127772.83722517936!2d38.65102105762962!3d9.020865221426748!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x163561add5fbc56b%3A0xaec7c5d3848f799b!2sAddis%20Ababa%2C%20Ethiopia!5e0!3m2!1sen!2sus!4v1649904164560!5m2!1sen!2sus"
                        width="100%" height="200" style="border:0;" loading="lazy"></iframe>
                </div>
            </div>

            <!-- Contact Form -->
            <form @submit.prevent="" class="bg-gray-200 p-6 rounded-lg shadow-md lg:w-1/2">
                <div class="mt-2">
                    <label for="name" class="block text-sm font-medium text-gray-700">Name:</label>
                    <input type="text" v-validate="'required'" name="name" id="name" placeholder="Enter your name"
                        class="mt-1 block w-full h-12 px-4 rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200">
                    <span class="text-red-500"></span>
                </div>
                <div class="mt-2">
                    <label for="email" class="block text-sm font-medium text-gray-700">Email:</label>
                    <input type="email" v-validate="'required|email'" name="email" id="email"
                        placeholder="Enter your email address"
                        class="mt-1 block w-full h-12 px-4 rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200">
                    <span class="text-red-500"></span>
                </div>
                <div class="mt-2">
                    <label for="message" class="block text-sm font-medium text-gray-700">Message:</label>
                    <textarea v-validate="'required'" name="message" id="message" rows="5"
                        placeholder="Enter your message here"
                        class="mt-1 block w-full h-36 px-4 py-2 rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200"></textarea>
                    <span class="text-red-500"></span>
                </div>
                <div class="mt-4 text-center">
                    <button type="submit"
                        class="inline-block bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded-lg shadow-md">Submit</button>
                </div>
            </form>
        </div>
    </div>
</template>