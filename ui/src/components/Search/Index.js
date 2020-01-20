import React from "react";
import { Col, Row } from "react-bootstrap";
import { jobService } from "services/jobs.service";
import Loader from "components/Loader";
import BoxLine from "components/BoxLine";
import BoxAlert from "components/BoxAlert";
import { DebounceInput } from "react-debounce-input";
import {history} from 'configureStore'
import Statistics from 'utils/Statistics'

/**
 * Render Deployment status
 */
export default class Search extends React.Component {

  state = {
    /**
     * Search imput typing
     */
    search: "",

    /**
     * Ajax job request
     */
    requested: true,

    /**
     * Job list
     */
    baseJobs: [],


    jobs: []
  };

  /**
   * When component mount
   */
  componentDidMount(){
    this.getJobs();
  }

  /**
   * On typing search input
   * @param {object} e - element
   */
  handleChangeText = (e) => {
    this.setState({ search: e.target.value });
    const filterJobs = this.state.baseJobs.filter(function(key) {
      return key.Job.includes(e.target.value);
    });
    this.setState({ jobs: filterJobs });
  }

  getJobs() {
    jobService.JobNames().then(
      data => {
        this.setState({ jobs: data, baseJobs: data, requested: false });
      },
      () => {
        this.setState({ requested: false });
      }
    );
  }

  goToDeployments(job) {
    Statistics.action({type: "go to job deployment",job: job})
    history.push(`/deployments/${job}`)
  }

  render() {
    return (
      <div className="job-search">
        <div className="page-title" />
        {/* Request time, show loader */ this.state.requested ? (
          <div className="center">
            <Loader />
          </div>
        ) : (
          <div>
            <Row>
              <Col sm="7" className="col-centered">
                <DebounceInput
                  debounceTimeout={500}
                  className="search-field"
                  placeholder="Search Job"
                  onChange={this.handleChangeText}
                />
              </Col>
            </Row>
            <Row>
              <Col sm="7" className="col-centered">
                {this.state.search.length > 0 && (
                  <div className="box small no-title">
                    <div className="content search-results">
                      {/* If user search job that not found */ !this.state
                        .requested &&
                        this.state.jobs.length == 0 &&
                        this.state.baseJobs.length > 0 && (
                          <BoxAlert Message="JOBS NOT FOUND" />
                        )}
                      {this.state.jobs.map((row, i) => (
                        <div className="line" key={i} onClick={this.goToDeployments.bind(this, row.Job)}>
                          <BoxLine  LeftText={row.Job}/>

                          </div>
                      ))}
                    </div>
                  </div>
                )}
              </Col>
            </Row>
          </div>
        )}
      </div>
    );
  }
}
