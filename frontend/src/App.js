import React from "react";
import "./App.css";
import { Form, Jumbotron } from "react-bootstrap";
import {
  Container,
  Row,
  Col,
  InputGroup,
  FormControl,
  Button
} from "react-bootstrap";
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
    this.state = { longLink: "" };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }

  handleSubmit(event) {
    let url = this.state.longLink;
    let valid = /^(ftp|http|https):\/\/[^ "]+$/.test(url);
    if (!valid) {
      alert("Invalid url submitted.");
    } else {
      alert("Submitted: " + this.state.longLink);
    }
    event.preventDefault(); // avoid reloading page after submission
  }
  handleChange(event) {
    this.setState({ longLink: event.target.value });
  }

  render() {
    return (
      <Container>
        <Row className="pt-5 mt-5">
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
                      style={{ width: "75%" }}
                    >
                      <Row>
                        <Col md={9}>
                          <InputGroup>
                            <FormControl
                              onChange={this.handleChange}
                              type="text"
                              placeholder="Shorten your link"
                              aria-label="Shorten your link"
                              aria-describedby="url-input"
                              required
                            />
                          </InputGroup>
                        </Col>
                        <Col md={3}>
                          <Button
                            className="btn"
                            type="submit"
                            block
                          >
                            Shorten
                          </Button>
                        </Col>
                      </Row>
                    </Form>
                  </center>
                </Col>
              </Row>
            </Jumbotron>
          </Col>
        </Row>
      </Container>
    );
  }
}

export default App;
