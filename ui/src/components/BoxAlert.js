import React from "react";
import PropTypes from "prop-types";
import Warning from "styles/icons/alert-triangle.svg"
import SVG from 'react-inlinesvg';

/**
 * Render box alert view
 */
export default class BoxAlert extends React.Component {

  static propTypes = {
    Message: PropTypes.string,
    TextClick: PropTypes.any
  
  };

  /**
   * Check is send function was sent
   */
  isClickBox(){
    return typeof this.props.TextClick == "function"
  }

  /**
   * If function is set and user click on the text, execute the function
   */
  clickBox = () => {  
    if (this.isClickBox()) {
      this.props.TextClick()
    }
  }

  render() {

    return (
      <div className="center empty-data" >
        <SVG src={Warning} className="" width="26" />
        <p className={`message ${this.isClickBox()? "link" : ""}`} onClick={this.clickBox}>{this.props.Message}</p>
      </div>
    );
  }
}
