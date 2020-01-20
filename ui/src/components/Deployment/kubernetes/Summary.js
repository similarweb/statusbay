import React from 'react'
import PropTypes from 'prop-types';
import { Row, Col } from "react-bootstrap";

/*
* Render deployment data
*/
export default class Deployment extends React.Component {

  static propTypes = {  
    DesiredState: PropTypes.number,
    Replicasets: PropTypes.object,
  };

  state = {

    replicas: 0,
    updatedReplicas: 0,
    readyReplicas: 0,
    availableReplicas: 0,
    UnavailableReplicas: 0,

  };


  /*
   * When component mount
   */
  componentWillMount(){
    const replicasets = this.props.Replicasets
    var replicas = 0
    var updatedReplicas = 0
    var readyReplicas = 0
    var availableReplicas = 0
    var UnavailableReplicas = 0
    Object.keys(replicasets).map(function(key) {
      const replicaStatus = replicasets[key].Status

      replicas += (replicaStatus.replicas)? replicaStatus.replicas : 0
      updatedReplicas += (replicaStatus.updatedReplicas)? replicaStatus.updatedReplicas : 0
      readyReplicas += (replicaStatus.readyReplicas)? replicaStatus.readyReplicas : 0
      availableReplicas += (replicaStatus.availableReplicas)? replicaStatus.availableReplicas : 0
      UnavailableReplicas += (replicaStatus.UnavailableReplicas)? replicaStatus.UnavailableReplicas : 0
    });
    this.setState({replicas,updatedReplicas,readyReplicas,availableReplicas,UnavailableReplicas})

  }

  render() {
    
    return (  
      <div>
        <Row>
          <Col sm="6" className="pr-md-1">
            <div className="box-number box">
            <h3 className="label">Desired Replicas</h3>
            <p className="value">{this.props.DesiredState}</p>
            </div>
          </Col>

          <Col sm="6" className="pr-md-1">
              <div className="box-number box">
                <h3 className="label">Current Replicas</h3>
                <p className="value">{this.state.replicas}</p>
              </div>
          </Col>
         
          </Row>
          <Row>
            <Col sm="3" className="pr-md-1">
                <div className="box-number box">
                  <h3 className="label">Updated Replicas</h3>
                  <p className="value">{this.state.updatedReplicas}</p>
                </div>
            </Col>
            <Col sm="3" className="pr-md-1">
                <div className="box-number box">
                  <h3 className="label">Ready Replicas</h3>
                  <p className="value">{this.state.readyReplicas}</p>
                </div>
            </Col>
            <Col sm="3" className="pr-md-1">
                <div className="box-number box">
                  <h3 className="label">Available Replicas</h3>
                  <p className="value">{this.state.availableReplicas}</p>
                </div>
            </Col>
            <Col sm="3" className="pr-md-1">
                <div className="box-number box">
                  <h3 className="label">Unavailable Replicas</h3>                
                  <p className="value">{this.state.UnavailableReplicas}</p>
                </div>
            </Col>
         </Row>       
         </div>
      )
  }
}