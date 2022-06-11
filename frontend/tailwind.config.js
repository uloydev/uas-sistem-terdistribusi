/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./src/components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: "#0d6efd",
        dark: "rgb(30,30,47)",
        "dark-light": "#27293D",
      },
    },
  },
  plugins: [],
};
