/** @type {import('tailwindcss').Config} */
module.exports = {
	theme: {
		fontFamily: {
			'sans': ['"Open Sans"', 'sans-serif'],
		},
	},
	content: ["./views/**/*.templ}", "./**/*.templ"],
	safelist: [],
	plugins: [require("daisyui")],
	daisyui: {
		themes: ["black"]
	}
}