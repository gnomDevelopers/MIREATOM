/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    screens: {
      'us': '450px',
      'mb': '576px',
      'sm':	'640px',
      'md':	'768px',
      'lg':	'1024px',
      'xl':	'1280px',
    },
    extend: {},
  },
  plugins: [],
}

