import React from 'react';
import ReactDOM from 'react-dom';
import logo from './mylogo.png';
import './App.css';

class App extends React.Component {
    constructor(props) {
    super(props);
    this.state = {email: ''};
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit(event) {
    event.preventDefault();
    fetch('http://localhost:8090/insertEmail', {
      mode: 'no-cors',
      method: 'post',
      body: {
        'email': this.email.value
      }
    })
    .then((resp) => resp.text())
    .then((text)=>  alert(text))
    .catch((err) => alert(err))
  }
  
  getInitialState() {
    return {email: ''}
  }
  handleChange(e) {
    this.setState({email: e.target.value})
  }
  
  render() {
    return (
      <div>
        <img className="logo" img src={logo} alt="My logo" />
        <form onSubmit={this.handleSubmit} name="sign up for beta form">
          <div className="header">
            <p>Coming soon.</p>
          </div>
          
          <div className="input">
            <input ref={(ref) => {this.email = ref}} type="email" className="button" id="email" value={this.state.value} onChange={this.handleChange} name="email" placeholder="Your email (it's a newsletter)" />
            <input type="submit" className="button" id="submit" disabled={!this.state.email} defaultValue="Submit" />
          </div>
        </form>
      </div>
    );
  }
};

export default App
