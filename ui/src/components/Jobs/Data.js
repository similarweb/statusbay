import React from "react";
import { history } from "configureStore";
import BoxLine from "components/BoxLine";
import PropTypes from "prop-types";
import Time from "utils/Time";
import Statistics from "utils/Statistics";

/**
 * Render Deployment status
 */
export default  class Data extends React.Component {

  static propTypes = {
    Name: PropTypes.string,
    NameIcon: PropTypes.string,
    DeployTime: PropTypes.number,
    Status: PropTypes.string
  };

  /**
   * Go to deployment page
   */
  goToDeployments = () =>  {
    Statistics.action({ type: "go to job deployment", job: this.props.Name });
    history.push(`/deployments/${this.props.Name}`);
  }

  /**
   * Return element with last deploy time
   * @param {int} e - timestamp
   */
  lastDeployment(time) {
    return (
      <div>
        <span>Last Deploy:</span> <span>{Time.FormatUnixTime(time)}</span>
      </div>
    );
  }

  render() {
    return (
      <div onClick={this.goToDeployments}>
        <BoxLine
          LeftText={this.props.Name}
          LeftTextIcon={this.props.NameIcon}
          LeftSubText={this.lastDeployment(this.props.DeployTime)}
          RightText={this.props.Status}
        />
      </div>
    );
  }
}
