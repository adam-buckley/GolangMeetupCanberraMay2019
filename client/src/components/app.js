import React from 'react'

export default class App extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      isLoading: true,
      timer: null
    }
  }
  componentDidMount() {
	WebAssembly.instantiateStreaming(fetch("http://localhost:3000"), go.importObject).then(async (result) => {
		go.run(result.instance)
    this.setState({ isLoading: false })
    
    let timer = window.setInterval(() => {
      UpdateDots()
    }, 100);
    this.setState({timer: timer})
	});
  }

  componentWillUnmount() {
    window.clearTimeout(this.state.timer)
  }
  render() {
    return this.state.isLoading ? <div>Loading</div> : <canvas height={1080} width={1080}></canvas>
  }
}