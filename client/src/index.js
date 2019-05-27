
// import React from 'react'
// import ReactDom from 'react-dom'

// import App from './components/app'

// ReactDom.render(<App />, document.getElementById('app'))
let timer;
let canvas = document.getElementById("canvas")
let BOUNDS_X = 1920
let BOUNDS_Y = 902

document.addEventListener("DOMContentLoaded", function() {
	WebAssembly.instantiateStreaming(fetch("http://localhost:3000"), go.importObject).then(async (result) => {
		go.run(result.instance)

		// let timer = window.setInterval(() => {
		// 	go.UpdateDots()
		// }, 100);
	}).then(() => {
		addDot(1)
		timer = window.setInterval(() => {
			var ctx = canvas.getContext("2d");
			ctx.fillStyle = "black"
			ctx.fillRect(0, 0, BOUNDS_X, BOUNDS_Y)

			ctx.fillStyle = "white"

			const r_dots = updateDots()
			const dots = JSON.parse(r_dots)
			// debugger;
			if (dots && dots.length) {
				for (i = 0; i < dots.length; i++) {
					// console.log(dots[i].x, dots[i].y)
					ctx.fillRect(dots[i].x, dots[i].y, 1, 1)
				}
			}
		}, 50)

		canvas.addEventListener("click", () => {
			addDot(1)
		})

		document.getElementById("add_1000").addEventListener("click", () => {
			addDot(1000)
		})

		document.getElementById("add_10000").addEventListener("click", () => {
			addDot(10000)
		})

		document.getElementById("add_a_lot").addEventListener("click", () => {
			addDot(1000000)
		})
	})
});
