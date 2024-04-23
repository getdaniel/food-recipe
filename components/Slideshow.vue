<template>
    <div class="swiper-container p-4">
        <Swiper class="swiper-wrapper"
            :modules="[Navigation, Pagination, SwiperAutoplay, SwiperEffectCreative]"
            :slides-per-view="2"
            :loop="true"
            :space-between="10" navigation
            :pagination="{ clickable: true }"
            :effect="'creative'"
            :autoplay="{
                delay: 2000,
                disableOnInteraction: false,
            }" 
            :creative-effect="{
                prev: {
                    shadow: false,
                    translate: ['-20%', 0, -1],
                },
                next: {
                    translate: ['100%', 0, 0],
                },
            }">
            <SwiperSlide v-for="slide in slides" :key="slide.id" class="swiper-slide">
                <div class="h-96 w-full p-4 flex justify-center items-center mx-auto border border-gray-400 rounded-lg">
                    <img :src="slide.image" alt="Slide Image" class="w-full h-full">
                </div>
            </SwiperSlide>
        </Swiper>

        <div class="container mx-auto">
            <div class="mt-8 border border-gray-700 rounded-lg p-10">
                <h2><strong>Food Categories:</strong></h2>
                <ul class="flex flex-wrap justify-center items-center mb-10 list-unstyled">
                    <li v-for="category in categories" :key="category.id" class="m-3">
                        <NuxtLink :to="`/recipes?category=${category.id}`"
                            class="p-2 px-4 rounded-lg bg-gray-200 hover:bg-teal-500 text-gray-700 font-medium">
                            {{ category.name }} ({{ getCategoryRecipeCount(category.id) }})
                        </NuxtLink>
                    </li>
                </ul>
                <h2><strong>Ingredients:</strong></h2>
                <ul class="flex flex-wrap justify-center items-center list-unstyled">
                    <li v-for="ingredient in ingredients" :key="ingredient.id" class="m-3">
                        <NuxtLink :to="`/recipes?ingredient=${ingredient.id}`"
                            class="p-2 px-4 rounded-lg bg-gray-200 hover:bg-teal-500 text-gray-700 font-medium">
                            {{ ingredient.name }}
                        </NuxtLink>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { Navigation, Pagination } from 'swiper/modules';
import query from '~/queries/HomePage.gql';

// Define types for your data
interface Category {
    id: number;
    name: string;
}

interface Ingredient {
    id: number;
    name: string;
}

interface Recipe {
    category_id: number;
    title: string;
}

interface QueryData {
    categories: Category[];
    ingredients: Ingredient[];
    recipes: Recipe[];
}

// Import your images (replace with your actual images)
import injeraImage from '../assets/images/injera.jpg';
import wafflesImage from '../assets/images/beyaynet.jpg';
import oatmealImage from '../assets/images/egg.jpg';
import saladImage from '../assets/images/meat.jpg';
import sandwichImage from '../assets/images/mitmita-meat.jpg';
import soupImage from '../assets/images/logo.png';

const slides = [
    { id: 1, image: wafflesImage },
    { id: 2, image: oatmealImage },
    { id: 3, image: saladImage },
    { id: 1, image: injeraImage },
    { id: 2, image: sandwichImage },
    { id: 3, image: soupImage },
    // Add more images as needed
];

const { data } = await useAsyncQuery<QueryData>(query);
const categories = data?.value?.categories ?? [];
const ingredients = data?.value?.ingredients ?? [];
const recipes = data?.value?.recipes ?? [];

// Function to get the count of recipes for a given category
const getCategoryRecipeCount = (categoryId: number): number => {
    return recipes.filter(recipe => recipe.category_id === categoryId).length;
};

</script>