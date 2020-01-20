import React from "react";
import { Col, Row } from "react-bootstrap";
import { jobService } from "../../services/jobs.service";
import Loader from "../Loader";
import BoxAlert from "../BoxAlert";
import Data from "./Data";
import BoxLine from "../BoxLine";
import { Scrollbars } from "react-custom-scrollbars";
import Statistics from "../../utils/Statistics";
import { history } from "../../configureStore";

/**
 * Render Deployment status
 */
export default class Jobs extends React.Component {

  state = {
    /**
     * Search input typing
     */
    search: "",

    /**
     * Search job ajax request
     */
    requested: true,

    /**
     * List of all jobs
     */
    baseJobs: [],

    /**
     * Filter jobs
     */
    jobs: []
  };

  /**
   * When component mount
   */
  componentDidMount() {
    this.getJobs();
  }

  /**
   * Getting last job deployments
   */
  getJobs() {
    jobService.jobs(20).then(
      data => {
        this.setState({ jobs: data, baseJobs: data, requested: false });
      },
      () => {
        this.setState({ requested: false });
      }
    );
  }

  /**
   * On typing search input
   * @param {object} e - element
   */
  handleChangeText = (e) =>  {
    this.setState({ search: e.target.value });
    const filterJobs = this.state.baseJobs.filter(function(key) {
      return key.Name.includes(e.target.value);
    });
    this.setState({ jobs: filterJobs });
  }

  goToSearchPage = () => {
    Statistics.action({ type: "go to search page" });
    history.push(`/search`);
  }

  render() {
    return (
      <div className="jobs-page">
        <div className="page-title">
          <h1>Last Deployments</h1>
        </div>
        <Row>
          <Col sm="7" className="col-centered">
            <input
              className="search-field no-margin"
              placeholder="Search Deployment"
              onChange={this.handleChangeText}
            />
            <span
              onClick={this.goToSearchPage}
              className="bottom-input-text pull-right"
            >
              Search Job
            </span>
          </Col>
        </Row>
        <Row>
          <Col sm="7" className="col-centered">
            <div className="box">
              <div className="content search-results">
                {/* Request time, show loader */ this.state.requested && (
                  <div className="center">
                    <Loader />
                  </div>
                )}
                {/* If data returned from the server with empty jobs  */ !this
                  .state.requested &&
                  this.state.baseJobs.length == 0 && (
                    <BoxAlert Message="JOBS NOT FOUND" />
                  )}
                {this.state.jobs.length > 0 && (
                  <BoxLine
                    LeftText="Last Deployments"
                    RightText={<p>Status</p>}
                    Title={true}
                  />
                )}
                {/* If user search job that not found */ !this.state
                  .requested &&
                  this.state.jobs.length == 0 &&
                  this.state.baseJobs.length > 0 && (
                    <BoxAlert Message="JOB DEPLOYMENT NOT FOUND" />
                  )}
                <Scrollbars style={{ height: 300 }}>
                  {this.state.jobs.map((job, i) => (
                    <Data
                      key={i}
                      Name={job["Name"]}
                      DeployTime={job["DeployTime"]}
                      Status={job["Status"]}
                    />
                  ))}
                </Scrollbars>
              </div>
            </div>
          </Col>
        </Row>
      </div>
    );
  }
}