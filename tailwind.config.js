/** @type {import('tailwindcss').Config} */

const { addDynamicIconSelectors } = require("@iconify/tailwind");
const { withAnimations } = require("animated-tailwindcss");
module.exports = withAnimations({
  content: [
    "./internal/views/*.{html,js,templ}",
    "./internal/components/*.{html,js,templ}",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require("@tailwindcss/typography"),
    require("daisyui"),
    addDynamicIconSelectors(),
  ],
  daisyui: {
    themes: ["light", "dark", "cupcake", "dracula"],
  },
});
