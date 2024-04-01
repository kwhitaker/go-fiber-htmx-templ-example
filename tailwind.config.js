/** @type {import('tailwindcss').Config} */
export default {
  content: ["./view/**/*.templ"], // this is where our templates are located
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: ["retro", "synthwave"],
  },
};
