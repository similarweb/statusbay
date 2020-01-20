import React from "react";
import PropTypes from "prop-types";
import { Col } from "react-bootstrap";
import BoxAlert from "components/BoxAlert";

/**
 * Render box alert view
 */
export default class BoxError extends React.Component {
  
  static propTypes = {
    ColWidth: PropTypes.number,
    Message: PropTypes.string,
    BackgroundClass: PropTypes.string,
    TextClick: PropTypes.any
  };

  render() {
    const backgroundClass = this.props.BackgroundClass
      ? this.props.BackgroundClass
      : "";
    const col = this.props.ColWidth == undefined ? 7 : this.props.ColWidth;
    return (
      <Col sm={col} className={`col-centered `}>
        <div className={`box ${backgroundClass}`}>
          <div className="content">
            <BoxAlert TextClick={this.props.TextClick}  Message={this.props.Message} /> 
          </div>
        </div>
      </Col>
    );
  }
}