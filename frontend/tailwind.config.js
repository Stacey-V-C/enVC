/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,tsx}"],
  theme: {
    colors: {
      "pale-green": "rgba(232, 243, 239, 0.87)",
      "pale-grey": "rgba(216, 216, 216, 0.87)",
      "off-white": "#F9F9F9",
    },
    extend: {
      transition: {
        "fade": "opacity 1s ease-in-out",
      }
    },
  },
  plugins: [],
}

