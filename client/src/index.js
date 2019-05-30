
let timer;
let canvas = document.getElementById("canvas")
let BOUNDS_X = 1920
let BOUNDS_Y = 902
let fps = document.getElementById("fps")
let dots_counter = document.getElementById("dots_counter")
window.t0 = 0

document.addEventListener("DOMContentLoaded", function() {
	WebAssembly.instantiateStreaming(fetch("http://localhost:3000"), go.importObject).then(async (result) => {
		go.run(result.instance)
	}).then(() => {
		addDot(1)

		window.drawCanvas = (buffer) => {

			// (attempt to) Limit fps to 60
			// setTimeout(() => {
				window.requestAnimationFrame(updateDots)

				// Draw black canvas
				var ctx = canvas.getContext("2d");
				ctx.fillStyle = "black"
				ctx.fillRect(0, 0, BOUNDS_X, BOUNDS_Y)

				ctx.fillStyle = "white"

				// Get dots list from WASM
				const dots = JSON.parse(buffer);
				dots_counter.innerHTML = (dots.length / 2) + " dots"
				if (dots && dots.length) {
					
					// Draw to screen
					for (i = 0; i < dots.length; i += 2) {
						ctx.fillRect(dots[i], dots[i + 1], 1, 1)
					}
				}

				var t1 = performance.now()
				fps.innerHTML = Math.round(1000 / (t1 - window.t0)) + " fps"
			// }, 1000/60)
		}
		
		window.requestAnimationFrame(updateDots) // ()

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
