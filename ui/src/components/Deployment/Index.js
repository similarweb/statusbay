import React from 'react'
import {connect} from 'react-redux';
import Loader from '../Loader';
import BoxError from '../BoxError';
import KubernetesDeployment from './kubernetes/Index'
import { jobService } from "../../services/jobs.service"
import PropTypes from 'prop-types';

/**
 * Render Deployment status
 */
@connect()
export default class Deployment extends React.Component {

  static propTypes = {    
    /**
     * Redux store
     */
    dispatch : PropTypes.func,

   /**
     * Match object
     */
    match : PropTypes.object, 
  };

  state = {
    /**
     * Ajax call for deployment data
     */
    requested: true,

    /**
     * Indicator for ajax error response 
     */
    requestedError: false,

    /**
     * Timeout id for pulling updated data (when deployment is running)
     */
    timeoutAjaxCall: null,

    orchestrator: null,
  }

  /*
   * When component mount
   */
  componentWillMount(){
    this.timeoutAjaxCall = null
    this.getDeployment(this.props.match.params.job, this.props.match.params.time)
  }

  /**
   * When component unmount
   */
  componentWillUnmount(){
    if (this.timeoutAjaxCall != null) {
         clearTimeout(this.timeoutAjaxCall);
    }
  }

  /**
   * Get deployment detail
   * @param {string} job - Job name
   * @param {int} time - Deploy time
   */
  getDeployment(job, time){
    jobService.deployment(job, time).then(
      response => {        
        this.props.dispatch({ type: 'NEW', deployment:response})
        this.setState({requested: false, orchestrator: response.Orchestrator})        
        if (response.Status == "running"){
          this.timeoutAjaxCall = setTimeout(() => { 
            this.getDeployment(job, time); 
          }, 2000);
        }
      },
      () => {
        this.setState({requestedError: true, requested: false})
      }
    )
  }
  
  RenderOrchestrator() {
    return <KubernetesDeployment />;
  }

  render() {
      return (
      <div className="deployment-page">      
       
        { this.state.requestedError && <BoxError Message={`data not found for job ${this.props.match.params.job}` }  BackgroundClass="none" /> }
        { this.state.requested && <div className="center"><Loader /></div> }

        { !this.state.requested  && !this.state.requestedError &&  
          <div>
            {this.RenderOrchestrator()}
          </div>}
      </div>
      )
  }
}