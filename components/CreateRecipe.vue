<template>
    <div v-if="visible"
        class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full flex justify-center items-center">
        <div class="bg-white p-4 rounded-lg max-w-lg w-full">
            <h2 class="text-xl font-bold mb-2">Create New Recipe</h2>
            <VeeForm @submit.prevent="submitRecipe">
                <div class="mb-4">
                    <label for="title" class="block text-gray-700 text-sm font-bold mb-2">
                        Title
                    </label>
                    <VeeField name="title" v-slot="{ field, errors }" rules="required">
                        <input :id="field.id" :name="field.name" v-bind="field" type="text"
                            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline">
                        <VeeErrorMessage :name="field.name" v-slot="{ errors }">
                            <p class="text-red-500 text-xs italic">{{ errors[0] }}</p>
                        </VeeErrorMessage>
                    </VeeField>
                </div>
                <!-- Additional fields can be added here following the same pattern -->
                <button type="submit"
                    class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
                    Submit
                </button>
                <button type="button" @click="close"
                    class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded ml-4">
                    Cancel
                </button>
            </VeeForm>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

// Importing VeeValidate components with custom names
import { Form as VeeForm, Field as VeeField, ErrorMessage as VeeErrorMessage } from 'vee-validate';

const props = defineProps({
    visible: Boolean
});

const emit = defineEmits(['update:visible']);

const recipe = ref({
    title: ''
});

function submitRecipe() {
    console.log("Recipe submitted:", recipe.value);
    close();
}

function close() {
    emit('update:visible', false);
}
</script>