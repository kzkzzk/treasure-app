import { h, Component } from "preact";
import firebase from "./firebase";
import { getPrivateMessage, getPublicMessage, getArticles, createArticle } from "./api";

class App extends Component {
  constructor() {
    super();
    this.state.user = null;
    this.state.message = "";
    this.state.errorMessage = "";

    this.state.article = {
      id: 1,
      title: "x",
      body: "y"
    }
    this.state.newarticle = {
      title: "",
      body: ""
    }

    this.state.tags = ["ex1", "ex2"]
  }

  // DOMの初期化時
  componentDidMount() {
    firebase.auth().onAuthStateChanged(user => {
      if (user) {
        this.setState({ user });
      } else {
        this.setState({
          user: null
        });
      }
    });
  }

  getPrivateMessage() {
    this.state.user
      .getIdToken()
      .then(token => {
        return getPrivateMessage(token);
      })
      .then(resp => {
        this.setState({
          message: resp.message
        });
      })
      .catch(error => {
        this.setState({
          errorMessage: error.toString()
        });
      });
  }

  changeId(e) {
    this.setState({article: {id: e.target.value}})
  }

  createTitle(e) {
    this.setState({newarticle: {title: e.target.value}})
  }

  createBody(e) {
    this.setState({newarticle: {body: e.target.value}})
  }

  getArticles() {
    this.state.user
      .getIdToken()
      .then(token => {
        return getArticles(this.state.article.id, token);
      })
      .then(resp => {
        let tags = []
        resp.tags.forEach((tag) => {
          tags.push(tag.name);
        })
        this.setState({
          errorMessage: "",
          tags: tags,
          article: {
            id: resp.id,
            title: resp.title,
            body: resp.body
          }
        });
      })
      .catch(error => {
        this.setState({
          errorMessage: error.toString()
        });
      });
  }

  submitArticle(e) {
    e.preventDefault();
    createArticle(this.state.newarticle)
      .then(resp => {
        let tags = []
        resp.tags.forEach((tag) => {
          tags.push(tag.name);
        })
        this.setState({
          errorMessage: "",
          tags: tags,
          article: {
            id: resp.id,
            title: resp.title,
            body: resp.body
          }
        });
      })
      .catch(error => {
        this.setState({
          errorMessage: error.toString()
        });
      });
  }

  render(props, state) {
    if (state.user === null) {
      return <button onClick={firebase.login}>Please login</button>;
    }
    return (
      <div>
        <div>
          <div>{state.message}</div>
          <p style="color:red;">{state.errorMessage}</p>
          <button onClick={this.getPrivateMessage.bind(this)}>
            Get Private Message
          </button>
          <button onClick={firebase.logout}>Logout</button>
          <div>
            <div>
              <input type="text" value={this.state.article.id} onChange={this.changeId.bind(this)}/>
              <button onClick={this.getArticles.bind(this)}>show</button>
            </div>
            <div>
              <h2>{state.article.title}</h2>
              <p>{state.article.body}</p>
              <span>{state.tags}</span>
            </div>
          </div>
        </div>
        <div>
        <h2>新規作成</h2>
        <form onSubmit={this.submitArticle.bind(this)}>
          <label>
            <input type="text" value={this.state.newarticle.title} onChange={this.createTitle.bind(this)}/>
          </label>
          <label>
            <input type="text" value={this.state.newarticle.body} onChange={this.createBody.bind(this)}/>
          </label>
          <input type="submit" value="新規作成" />
        </form>
      </div>
    </div>
    );
  }
}

export default App;
