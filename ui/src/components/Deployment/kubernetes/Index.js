import React from 'react'
import {connect} from 'react-redux';
import PropTypes from 'prop-types';
import Status from 'utils/Status'
import { Tabs, Tab, Row } from "react-bootstrap";
import Pods from 'components/Deployment/kubernetes/Pods'
import Replicasets from 'components/Deployment/kubernetes/Replicasets'
import Events from 'components/Deployment/kubernetes/Events'
import Summary from 'components/Deployment/kubernetes/Summary'
import Metrics from 'components/Deployment/kubernetes/Metrics'
// import { Row, Col } from "react-bootstrap";
import BoxError from 'components/BoxError';
import Loader from 'components/Loader';
import LogLinks from 'components/Deployment/kubernetes/LogLinks';
import Alerts from 'components/Graph/Alert';

/**
 * Render Deployment status
 */
@connect(state => ({
  Status: state.deployment.Status,
  Name: state.deployment.Data.Application,
  CreationTimestamp: state.deployment.Data.CreationTimestamp,
  Deployments: state.deployment.Data.Deployments,
  DeploymentDescription: state.deployment.Data.DeploymentDescription,
}))
export default class KubernetesDeployment extends React.Component {

  static propTypes = {    
    /**
     * Application name
     */
    Name : PropTypes.string, 

    /**
     * Application status
     */
    Status : PropTypes.string, 

    /**
     * Application Deployment creation time
     */
    CreationTimestamp : PropTypes.number,

    /**
     * List of deployments
     */
    Deployments : PropTypes.object,

    /**
     * Application status description
     */
    DeploymentDescription : PropTypes.string,
   
  };

 
  state = {
    selectedTab: Object.keys(this.props.Deployments)[0],
  }

  render() {    
      return (
      <>
        <div className="page-title">
            <h1 onClick={this.onClickJobURL}>{this.props.Name}</h1>
        </div>

        <div className={`alert alert-${Status.getColorByStatus(this.props.Status)}`}>
          <p><span>{Status.getIconByStatus(this.props.Status)}</span> {this.props.DeploymentDescription} </p>                        
        </div>


        <Tabs defaultActiveKey={this.state.selectedTab} transition={false}>
          {Object.keys(this.props.Deployments).map((name) => 
            <Tab key={name} eventKey={name} title={name}>

              <Summary DesiredState={this.props.Deployments[name].Deployment.DesiredState} Replicasets={this.props.Deployments[name].Replicaset} /> 
              {this.props.Deployments[name].Pods == undefined || (typeof(this.props.Deployments[name].Pods) == "object" && Object.keys(this.props.Deployments[name].Pods).length == 0  ) ?
                this.props.Status == "running" ? <div className="center"><Loader />Waiting for pods data... please wait</div>
                :
                <Row><BoxError ColWidth="12" Message="Pods data not found"  /></Row>
                :
                <Pods Events={this.props.Deployments[name].Pods}/>
              }


            {this.props.Deployments[name].DeploymentEvents == undefined || (typeof(this.props.Deployments[name].DeploymentEvents) == "object" && Object.keys(this.props.Deployments[name].DeploymentEvents).length == 0  ) ?
              this.props.Status == "running" ? <div className="center"><Loader />Waiting for replicaset data... please wait</div>
              :
              <Row><BoxError ColWidth="12" Message="Deployment data not found"  /></Row>
              :
              <Events Title="Deployment Events" Events={this.props.Deployments[name].DeploymentEvents}/> 
            }

              <Replicasets Replicasets={this.props.Deployments[name].Replicaset} />
              <Metrics Metrics={this.props.Deployments[name].Metrics} DeployTime={this.props.CreationTimestamp} MetricsRange={60}/>

              <div className="section-title">Links</div>
              <Row>
                {<LogLinks Links={this.props.Deployments[name].LogsLinks}/> }
              </Row>
              <div className="section-title">Alerts</div>
              
              <Alerts AlertTags={this.props.Deployments[name].AlertTags} DeployTime={this.props.CreationTimestamp} MetricsRange={60}/>
            </Tab>
          )}
        </Tabs>


      </>
      )
  }
}