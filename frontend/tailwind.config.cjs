module.exports = {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Urbanist', 'sans-serif'],
      },
      colors: {
        primary: '#2D6A4F',
        accent: '#52B788',
        surface: '#F1F8F4',
        danger: '#D62828',
        text: '#1B1B1B',
      },
    },
  },
  plugins: [],
}
