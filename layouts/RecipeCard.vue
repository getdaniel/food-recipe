<template>
    <div class="container mx-auto px-2 py-5">
        <template v-for="category in recipeCategories" :key="category.name">
            <h2 class="text-xl font-semibold text-gray-100 mb-2 pt-12">{{ category.name }}</h2>

            <div class="flex flex-wrap -m-2">
                <div v-for="recipe in category.recipes" :key="recipe" class="p-2 mx-auto aspect-square">
                    <div class="recipe-card w-72 h-auto bg-gray-200 shadow-md rounded-lg relative">
                        <img src="../assets/images/logo.png" alt="Recipe Photo"
                            class="w-full h-48 object-cover rounded-t-lg" />

                        <div class="px-4 py-2 flex flex-col justify-between">
                            <h2 class="text-xl font-semibold text-gray-800 mb-2 text-center">{{ recipe.name }}</h2>
                            <p class="text-base text-gray-600 mb-2 overflow-hidden h-18"> {{
                                recipe.description }}
                            </p>
                            <div class="border-b border-gray-300"></div>
                        </div>

                        <div class="px-4 py-1 flex flex-col justify-between">
                            <div class="flex items-center text-gray-800">
                                <Icon name="uil:star" size="21" class="text-yellow-900"></Icon>
                                <span class="px-1">4.5</span>
                                <span class="ml-1">(120 ratings)</span>
                            </div>
                        </div>

                        <div class="absolute bottom-2 right-2">
                            <Icon name="uil:heart" size="21"></Icon>
                        </div>
                    </div>
                </div>
            </div>
        </template>
    </div>
</template>

<script>
export default {
    data() {
        return {
            recipeCategories: [
                {
                    name: "Breakfast",
                    recipes: [
                        { name: "Injera", description: "Spongy, slightly sour flatbread, staple food of Ethiopia and Eritrea." },
                        { name: "Waffles", description: "Light and fluffy breakfast treat, often served with syrup and fruit." },
                        { name: "Oatmeal", description: "Warm and hearty porridge made with rolled oats." }
                    ]
                },
                {
                    name: "Lunch",
                    recipes: [
                        { name: "Salads", description: "Fresh and healthy combination of greens, vegetables, and protein." },
                        { name: "Sandwiches", description: "Meal option with various bread, filling, and topping combinations." },
                        { name: "Soups", description: "Comforting and nourishing liquid meal, perfect for a light lunch." }
                    ]
                },
                // ... other categories
            ]
        };
    }
};
</script>