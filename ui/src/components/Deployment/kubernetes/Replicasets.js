import React from 'react'
import PropTypes from 'prop-types';
import Events from 'components/Deployment/kubernetes/Events'

/*
* Render deployment data
*/
export default class Replicasets extends React.Component {

  static propTypes = {  
    Replicasets: PropTypes.object,
  };

  state = {

  }


  render() {

    return (  
      <div>
         {Object.keys(this.props.Replicasets).map((replicaName) => 
            <div key={`replica-event-${replicaName}`}>{(this.props.Replicasets[replicaName].Events.length > 0) &&  <Events Title={`Replicaset ${replicaName} events`} Events={this.props.Replicasets[replicaName].Events}/> }</div>
         )}
      </div>
      )
  }
}