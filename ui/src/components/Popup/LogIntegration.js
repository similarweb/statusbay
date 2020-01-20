import React from "react";
import { Modal, Button } from "react-bootstrap";
import Statistics from "utils/Statistics";
import PropTypes from "prop-types";

/**
 * Show log integration popup
 */
export default class LogIntegration extends React.Component {

  static propTypes = {
    Show: PropTypes.any
  };

  state = {
    /**
     * Is popup is open
     */
    show: false
  };

  /**
   * When component mount
   */
  componentDidMount() {
    Statistics.action({ type: "Popup: view log integration" });
    this.props.Show(this.handleShow);
  }

  /**
   * Click for close popup
   */
  handleClose = () => {
    this.setState({ show: false });
  }

  /**
   * Click for show popup
   */
  handleShow = () => {
    this.setState({ show: true });
  }

  /**
   * Nomad example
   */
  nomadExample() {
    return `
<code>
job "job-name" {
  meta {
    log-query = "application: appName
    log-query-2 = "application: appName AND environment: <[ job_environment|default('staging') ]>"
  }
}
</code>
`;
  }

  /**
   * Kubernetes example
   */
  kubernetesExample() {
    return `
<code>
---
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    statusbay.io/kibana-query: "application: statusbay AND mode: watcher"
    statusbay.io/kibana-query-1: "application: statusbay AND mode: client"
    statusbay.io/kibana-query-1: "@_facility: fcfetcher AND environment: production" 
</code>
`;
  }

  render() {
    return (
      <Modal size="lg" show={this.state.show} onHide={this.handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Log integration</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <pre>
            {
              <div>
                <div>
                  <h5>Nomad:</h5>
                  Add Meta{" "}
                  <a target={`_blank`} href="https://www.nomadproject.io/docs/job-specification/meta.html">
                    Stanza
                  </a>{" "}
                  to the Job file
                  <div dangerouslySetInnerHTML={{ __html: this.nomadExample() }} />
                </div>


                <div>
                  <h5>Kubernetes:</h5>
                  Add{" "}
                  <a target={`_blank`} href="https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/">
                  Annotation
                  </a>{" "}
                  to the deployment
                  <div dangerouslySetInnerHTML={{ __html: this.kubernetesExample() }} />
                </div>
              </div>
            }
          </pre>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={this.handleClose}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
    );
  }
}
