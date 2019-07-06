import React from 'react';
import './App.css';
import {Form} from 'react-bootstrap';
import {Button} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.css';


class App extends React.Component {

  renderForm() {
    return <InputForm />;
  }

  render(){
    return (
      <div>
        <header>
        <script type="text/babel"></script>
        </header>
        <div className="App">
          {this.renderForm()}
        </div>
      </div>
    );
  }
}

class InputForm extends React.Component{
  constructor(props) {
    super(props);
    this.state= {longlink: ''};
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleChange= this.handleChange.bind(this);
  }
  
  handleSubmit(event){
    alert("submitted" + this.state.longlink); 
    event.preventDefault(); // avoid reloading page after submission
  }
  handleChange(event){
    this.setState({longlink:event.target.value});
  }

  render(){
    return (
      <center>
      <Form 
        onSubmit={e => this.handleSubmit(e)}
        style={{ width: '50%' }}
      >
        <Form.Group controlId="formBasicEmail">
          <Form.Control onChange={this.handleChange} type="text" placeholder="Shorten your link" required />
        </Form.Group>
        <Button variant="primary" type="submit">Submit</Button>
      </Form>
      </center>
    );
  }
}

export default App;
