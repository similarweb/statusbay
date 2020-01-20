import React from "react";
import { Modal, Button } from "react-bootstrap";
import Statistics from "utils/Statistics";
import PropTypes from "prop-types";

/**
 * Show alerts integration popup
 */
export default class AlertsIntegration extends React.Component {

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
    Statistics.action({ type: "Popup: view alert integration" });
    this.props.Show(this.handleShow);
  }

  /**
   * Click for close popup
   */
  handleClose = () =>  {
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
    alerts-tags="tag,anotherTag"
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
  statusbay.io/alerts-statuscake-tags: nginx
</code>
`;
  }

  render() {
    return (
      <Modal show={this.state.show} onHide={this.handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Statuscake integration</Modal.Title>
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
                to deployment
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
