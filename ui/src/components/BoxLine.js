import React from "react";
import PropTypes from "prop-types";
import Status from "utils/Status";
import { OverlayTrigger, Tooltip } from "react-bootstrap";

/**
 * Render box alert view
 */
export default class BoxLine extends React.Component {
  
  static propTypes = {
    LeftText: PropTypes.string,
    LeftTextIcon: PropTypes.string,
    NameIcon: PropTypes.string,
    LeftSubText: PropTypes.object,
    RightText: PropTypes.oneOfType([PropTypes.object, PropTypes.string]),
    Title: PropTypes.oneOfType([PropTypes.bool]),
    RightClass: PropTypes.string
  };

  render() {
    return (
      <div
        className={`line ${this.props.Title ? "title" : ""}`}
        onClick={this.goToDeployments}
      >
        <div className="left col-md-11">
          <div className="box-line">
            <p className="box-line-title">{this.props.LeftText} </p>
            <span className="box-line-title-subtitle">
              {this.props.LeftSubText}
            </span>
          </div>
        </div>
        <div className="right col-md-1">
          <OverlayTrigger
            key="top"
            placement="top"
            overlay={
              <Tooltip id={`tooltip-top`}>
                <span>{this.props.RightText}</span>
              </Tooltip>
            }
          >
            <span>{Status.getIconByStatus(this.props.RightText)}</span>
          </OverlayTrigger>
        </div>
      </div>
    );
  }
}
