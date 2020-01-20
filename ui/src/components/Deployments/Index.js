import React from "react";
import { Col, Row } from "react-bootstrap";
import PropTypes from "prop-types";
import { jobService } from "services/jobs.service";
import Data from "components/Deployments/Data";
import Loader from "components/Loader";
import BoxAlert from "components/BoxAlert";
import BoxLine from "components/BoxLine";
import { Scrollbars } from "react-custom-scrollbars";

/**
 * Render Deployment status
 */
export default class Deployments extends React.Component {

  static propTypes = {
    match: PropTypes.object
  };

  state = {
    /**
     * Ajax deployment request
     */
    isRequested: true,

    /**
     * Deployment list
     */
    deployments: []
  };

  /**
   * When component mount
   */
  componentDidMount() {
    this.getDeployments();
  }

  /**
   * Getting last deployment for a job
   */
  getDeployments() {
    jobService.deployments(this.props.match.params.job, 20).then(
      data => {
        this.setState({ deployments: data, isRequested: false });
      },
      () => {
        this.setState({ isRequested: false });
      }
    );
  }

  render() {
    return (
      <div className="deployments-page">
        <div className="page-title">
          <h1>{this.props.match.params.job}</h1>
        </div>
        <Row>
          <Col sm="7" className="col-centered">
            <div className="box">
              <div className="content">
                {/* Request time, show loader */ this.state.isRequested && (
                  <div className="center">
                    <Loader />
                  </div>
                )}
                {/* If data returned from the server with empty deployments  */ !this
                  .state.isRequested &&
                  this.state.deployments.length == 0 && (
                    <BoxAlert Message="DEPLOYMENTS NOT FOUND" />
                  )}

                {this.state.deployments.length > 0 && (
                  <div>
                    <BoxLine
                      LeftText="Previous Deployments"
                      RightText={<p>Status</p>}
                      Title={true}
                    />

                    <Scrollbars style={{ height: 300 }}>
                      {this.state.deployments.map((job, i) => (
                        <Data
                          key={i}
                          JobID={this.props.match.params.job}
                          ID={job["ID"]}
                          DeployTime={job["DeployTime"]}
                          Status={job["Status"]}
                        />
                      ))}
                    </Scrollbars>
                  </div>
                )}
              </div>
            </div>
          </Col>
        </Row>
      </div>
    );
  }
}
