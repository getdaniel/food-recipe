<template>
    <div class="flex justify-center items-center h-full m-1">
        <VeeForm>
            <VeeField name="search" rules="required">
                <input type="text" placeholder="Enter your search keyword"
                    class="border border-gray-600 bg-gray-300 rounded-md px-4 py-2 focus:outline-none focus:border-blue-500 w-96 placeholder-gray-500">
                <VeeErrorMessage name="search" />
            </VeeField>
        </VeeForm>
    </div>
</template>