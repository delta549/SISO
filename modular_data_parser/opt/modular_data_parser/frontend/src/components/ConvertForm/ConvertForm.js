import './ConvertForm.css';
import { Component } from 'react';

class ConvertForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      valueOne: 'JSON',
      valueTwo: 'JSON'
    };
    this.handleChangeFirst = this.handleChangeFirst.bind(this);
    this.handleChangeSecond = this.handleChangeSecond.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChangeFirst(event) {
    this.setState({valueOne: event.target.value});
    console.log(this.state.valueOne)
  }

  handleChangeSecond(event) {
    this.setState({valueTwo: event.target.value});
  }

  handleSubmit(event) {
    alert('Your first selection is: ' + this.state.valueOne + ' Your second is: ' + this.state.valueTwo);

    event.preventDefault();
  }

  render() {
    return (
      <div className='ConvertForm'>
        <form onSubmit={this.handleSubmit}>
          <div className='Uploader'>
            <input type="file" id="myFile" name="filename"/>
          </div>
          <label>
            <select value={this.state.valueOne} onChange={this.handleChangeFirst}>
              <option value="CSV">CSV</option>
              <option value="JSON">JSON</option>
              <option value="TSV">TSV</option>
            </select>
          </label>
          <label>
            <select value={this.state.valueTwo} onChange={this.handleChangeSecond}>
              <option value="CSV">CSV</option>
              <option value="JSON">JSON</option>
              <option value="TSV">TSV</option>
            </select>
          </label>
          <input type="submit" value="Convert" />
        </form>
      </div>
    );
  }
}

export default ConvertForm