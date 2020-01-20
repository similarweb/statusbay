import React from 'react'
import PropTypes from 'prop-types';
import { Row, Col, Table } from "react-bootstrap";
import { Scrollbars } from 'react-custom-scrollbars';
import Statistics from 'utils/Statistics'
import Timeline from 'components/Timeline'

/*
* Render deployment data
*/
export default class Pods extends React.Component {

  static propTypes = {  
    Events: PropTypes.object,
  };

  state = {
    selectedPod: Object.keys(this.props.Events)[0],
  }


  changeTimelineEvents = (podId) => {
    Statistics.action({type: "change timeline"})

    this.setState({
      selectedPod: podId,
    });
  }
  
   /**
   *  Render timeline by selected allocation
   * 
   */
  renderEventTimeline() {
      return <Timeline Events={this.props.Events[this.state.selectedPod].Events}/>     
  }
  

  render() {

    return (  
      <div>
        <div className="section-title">
              Pod Events
            </div>
            <div className="box events">
            <div className="content "> 
              <Row>              
                <Col sm="4">
                <Scrollbars style={{height: 300 }}>
                      <Table  hover className="table-allocations">
                        <thead>
                          <tr>
                            <th>Pod <span className="subtitle">({Object.keys(this.props.Events).length})</span></th>
                            <th>Status</th>
                          </tr>
                        </thead>
                        <tbody>
                        {Object.keys(this.props.Events).map((podId)  =>
                          <tr key={podId} className={this.state.selectedPod == podId  ? "selected": ""} onClick={this.changeTimelineEvents.bind(this, podId)}>
                            <td>
                              <p >{this.props.Events[podId].Marked && <span className="circle-red"></span>} {podId} </p>
                            </td>
                            <td >
                              <p >{this.props.Events[podId].Phase} </p>
                            </td>
                           
                          </tr>
                        )}
                        </tbody>
                      </Table>
                      </Scrollbars>
                </Col>
                <Col sm="8">
                    <h3> Events</h3>
                    <Scrollbars style={{height: 300 }}>
                    {this.renderEventTimeline()}
                    </Scrollbars>
                </Col>
                </Row>
                </div>
            </div>
      </div>
      )
  }
}