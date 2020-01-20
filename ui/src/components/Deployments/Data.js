import React from "react";
import { history } from "configureStore";
import Status from "utils/Status";
import Time from "utils/Time";
import Statistics from "utils/Statistics";
import BoxLine from "../BoxLine";
import PropTypes from "prop-types";

/**
 * Render Deployment status
 */
export default class Data extends React.Component {
  
  static propTypes = {
    ID: PropTypes.string,
    JobID: PropTypes.string,
    Name: PropTypes.string,
    DeployTime: PropTypes.number,
    Status: PropTypes.string
  };

  /**
   * On click go to deployment page.
   */
  goToDeployments = () => {
    Statistics.action({
      type: "go to job details",
      job: this.props.JobID,
      timestamp: this.props.DeployTime
    });
    history.push(`/deployments/${this.props.JobID}/${this.props.DeployTime}`);
  }

  render() {
    return (
      <div onClick={this.goToDeployments}>
        <BoxLine
          LeftText={Time.FormatUnixTime(this.props.DeployTime)}
          RightText={this.props.Status}
          RightClass={`alert-${Status.getColorByStatus(this.props.Status)}`}
        />
      </div>
    );
  }
}