
// import React from 'react'
// import ReactDom from 'react-dom'

// import App from './components/app'

// ReactDom.render(<App />, document.getElementById('app'))
let timer;

document.addEventListener("DOMContentLoaded", function() {
	WebAssembly.instantiateStreaming(fetch("http://localhost:3000"), go.importObject).then(async (result) => {
		go.run(result.instance)

		// let timer = window.setInterval(() => {
		// 	go.UpdateDots()
		// }, 100);
	}).then(() => {
		timer = window.setInterval(() => updateDots(), 50)

		var canvas = document.getElementById("canvas")
		canvas.addEventListener("click", () => {
			addDot(100)
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
