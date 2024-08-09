/** @type {import('tailwindcss').Config} */
module.exports = {
	theme: {
		container: {
			screens: {
				sm: '600px',
				md: '728px',
				lg: '884px',
				xl: '1000px',
				'2xl': '1496px',
			},
		},
		fontFamily: {
			'sans': ['"JetBrains Mono"', 'monospace'],
			'jetbrains': ['"JetBrains Mono"', 'monospace'],
			'oswald': ['Oswald', 'sans-serif'],

		},
	},
	content: ["./views/**/*.templ}", "./**/*.templ"],
	safelist: [],
	plugins: [require("daisyui")],
	daisyui: {
		themes: ["white"]
	}
}