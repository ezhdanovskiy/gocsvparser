// -*- JavaScript -*-


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

ReactDOM.render( <Dictionary/>, document.querySelector("#root"));
