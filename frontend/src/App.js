import React from "react";
import "./App.css";
import { Form, Jumbotron } from "react-bootstrap";
import { Row, Col, InputGroup, FormControl, Button } from "react-bootstrap";
import "bootstrap/dist/css/bootstrap.css";

class App extends React.Component {
  renderForm() {
    return <InputForm />;
  }

  render() {
    return (
      <div>
        <header>
          <script type="text/babel" />
        </header>
        <div className="App">{this.renderForm()}</div>
      </div>
    );
  }
}

class InputForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = { longlink: "" };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }

  handleSubmit(event) {
    alert("submitted" + this.state.longlink);
    event.preventDefault(); // avoid reloading page after submission
  }
  handleChange(event) {
    this.setState({ longlink: event.target.value });
  }

  render() {
    const btn = {backgroundColor: "#6d9ce8", border: "none"}
    return (
      <container>
        <Row className="pl-5 pr-5 pt-5 mt-5">
          <Col>
            <Jumbotron style={{ backgroundColor: "#f9f1f1" }}>
              <Row className="pt-2 pb-3">
                <Col>
                  <h1>A smile is the universal welcome.</h1>
                  <h3>
                    <i>Max Eastman</i>
                  </h3>
                </Col>
              </Row>
              <br />
              <Row>
                <Col>
                  <center>
                    <Form
                      onSubmit={e => this.handleSubmit(e)}
                      style={{ width: "50%" }}
                    >
                      <InputGroup className="mb-3">
                        <FormControl
                          onChange={this.handleChange}
                          type="text"
                          placeholder="Shorten your link"
                          aria-label="Shorten your link"
                          aria-describedby="url-input"
                          required
                        />
                        <InputGroup.Append>
                          <Button className="pl-5 pr-5" style={btn} variant="primary" type="submit">Submit</Button>
                        </InputGroup.Append>
                      </InputGroup>
                    </Form>
                  </center>
                </Col>
              </Row>
            </Jumbotron>
          </Col>
        </Row>
      </container>
    );
  }
}

export default App;
