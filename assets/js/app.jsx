// -*- JavaScript -*-


class PersonItem extends React.Component {
  render() {
    return (
      <tr>
        <td> {this.props.id}    </td>
        <td> {this.props.first} </td>
        <td> {this.props.last}  </td>
      </tr>
    );
  }
}

class DictionaryItem extends React.Component {
  render() {
    return (
      <tr>
        <td> {this.props.source_category}    </td>
        <td> {this.props.source_contractor} </td>
        <td> {this.props.category}  </td>
        <td> {this.props.contractor}  </td>
        <td> {this.props.comment}  </td>
      </tr>
    );
  }
}

class PeopleList extends React.Component {
  constructor(props) {
    super(props);
    this.state = { people: [] };
  }

  componentDidMount() {
    this.serverRequest =
      axios
        .get("/people")
        .then((result) => {
           this.setState({ people: result.data });
        });
  }

  render() {
    const people = this.state.people.map((person, i) => {
      return (
        <PersonItem key={i} id={person.Id} first={person.First} last={person.Last} />
      );
    });

    return (
      <div>
        <table><tbody>
          <tr><th>Id</th><th>First</th><th>Last</th></tr>
          {people}
        </tbody></table>

      </div>
    );
  }
}

class Dictionary extends React.Component {
  constructor(props) {
    super(props);
    this.state = { dictionary: [] };
  }

  componentDidMount() {
    this.serverRequest =
      axios
        .get("/dictionary")
        .then((result) => {
           this.setState({ dictionary: result.data });
        });
  }

  render() {
    const dictionary = this.state.dictionary.map((item, i) => {
      return (
        <DictionaryItem key={i}
        source_category={item.sourceCategory}
        source_contractor={item.sourceContractor}
        category={item.category}
        contractor={item.contractor}
        comment={item.comment} />
      );
    });

    return (
      <div>
        <table><tbody>
          <tr><th>SourceCategory</th><th>SourceContractor</th><th>Category</th><th>Contractor</th><th>Comment</th></tr>
          {dictionary}
        </tbody></table>

      </div>
    );
  }
}

// ReactDOM.render( <PeopleList/>, document.querySelector("#root"));
ReactDOM.render( <Dictionary/>, document.querySelector("#root"));
