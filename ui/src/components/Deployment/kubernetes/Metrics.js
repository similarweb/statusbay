import React from 'react'
import PropTypes from 'prop-types';
import { Row } from "react-bootstrap";
import Metric from 'components/Graph/Metric';
import moment from 'moment'

/**
 * Render Metrics
 */
export default class Metrics extends React.Component {

  static propTypes = {    
    /**
     * Example: [{Query:"", Title:""}]
     */
    Metrics : PropTypes.array, 

    /**
     * Deployment time
     */
    DeployTime : PropTypes.number, 

    /**
     * Metrics limit range
     */
    MetricsRange : PropTypes.number, 
  };

  state = {
    /**
     * Get metric data from date
     */
    from: moment.unix(this.props.DeployTime).subtract(this.props.MetricsRange, 'm'),

    /**
     * Get metric data to date
     */
    to: moment.unix(this.props.DeployTime).add(this.props.MetricsRange, 'm'),
    
    /**
     * draw static line. 
     */
    plotLines:[{  color: 'red', width: 2,value: moment.unix(this.props.DeployTime),dashStyle: 'longdashdot'}]
  }

  render() {
    return ( 
      <div>
        <div className="section-title">
          Metrics (
            <span>{this.state.from.format("MM/DD/YYYY HH:mm")}</span> - <span>{this.state.to.format("MM/DD/YYYY HH:mm")}</span>)
        </div>
            <Row>
              { this.props.Metrics.map((metric,i) => 
                <Metric key={i} Title={metric.Title} 
                Query={metric.Query}
                SubTitle={metric.SubTitle} 
                From={this.state.from} 
                To={this.state.to}
                PlotLines={this.state.plotLines}
                />
              )}
          </Row>
      </div>
      )
  } 
}