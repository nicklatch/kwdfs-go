/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: ['selector', '[data-mode="dark"]'],
    content: ["./cmd/web/templates/**/*.{html,js}"],
    theme: {
        extend: {},
    },
    plugins: [],
}

