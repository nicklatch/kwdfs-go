/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
    darkMode: ['selector', '[data-mode="dark"]'],
    content: ["./cmd/web/templates/**/*.{html,gohtml,js}"],
    theme: {
        fontFamily: {
            'sans': ['"Noto Sans"', ...defaultTheme.fontFamily.sans]
        }, extend: {},
    }, plugins: ['@tailwindcss/forms'],
}

