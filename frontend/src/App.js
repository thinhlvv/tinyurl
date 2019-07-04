import React from 'react';
import logo from './logo.svg';
import './App.css';

class App extends React.Component {
  constructor() {
    super();
  }

  renderForm() {
    return <InputForm />;
  }

  render(){
    return (
      <div>
        <header>
        <script type="text/babel"></script>
        </header>
        <body>
        <div className="App">
          {this.renderForm()}
        </div>
        </body>
      </div>
    );
  }
}

class InputForm extends React.Component{
  constructor(props) {
    super(props);
    this.status = {value: ''};
    // this.handleSubmit = this.postLonglinkToServer.bind(this);
  }
  
  postLonglinkToServer(event){
    alert("submitted" + this.status.value); 
    event.preventDefault();
  }

  render(){
    return (
      <center>
        <form onSubmit={this.postLonglinkToServer}>
        <label>
          <input type="text" value={this.status.value} placeholder="Shorten your link"/>
        </label>
        <input type="submit" value="Shorten" />
        </form>
      </center>
    );
  }
}

export default App;
