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

  handleChange(event) {
    this.setState({value: event.target.email});
  }

  handleSubmit(event) {
    alert('Votre email: ' + this.state.email + " a bien été prise en compte. Vous êtes bien inscrit à notre Newsletter, merci !");
    event.preventDefault();
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
        <form action="/index.html" onSubmit={this.handleSubmit} method="post" name="sign up for beta form">
          <div className="header">
            <p>Bientôt de retour, en beta finale.</p>
          </div>
          
          <div className="input">
            <input type="email" className="button" id="email" value={this.state.value} onChange={this.handleChange} name="email" placeholder="Votre email (c'est une newsletter)" />
            <input type="submit" className="button" id="submit" disabled={!this.state.email} defaultValue="Envoyer" />
          </div>
        </form>
      </div>
    );
  }
};

export default App