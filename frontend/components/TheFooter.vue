<template>
    <footer class="text-center py-2 bg-sky-950">
        <p class="text-gray-100 mb-5">
            Food Recipes Site
        </p>
        <!-- Social Media Links -->
        <div class="flex justify-center mt-2 space-x-4">
            <a href="https://github.com/getdaniel" target="_blank" title="GitHub">
                <Icon name="uil:github" color="black" />
            </a>
            <a href="https://www.linkedin.com/in/daniel-getaneh" target="_blank" title="LinkedIn">
                <Icon name="uil:linkedin" color="black" />
            </a>
            <a href="https://t.me/dgety" target="_blank" title="Telegram">
                <Icon name="uil:telegram" color="black" />
            </a>
            <a href="https://discord.com/invite/getdaniel" target="_blank" title="Discord">
                <Icon name="uil:discord" color="black" />
            </a>
            <a href="https://x.com/DanielGetanehS" target="_blank" title="X">
                <Icon name="uil:twitter" color="black" />
            </a>
        </div>
        {{ currentYear }} &copy; All Rights Reserved.
    </footer>
</template>

<script>
export default {
    data() {
        return {
            currentYear: new Date().getFullYear()
        };
    }
}
</script>