import React from 'react'
import moment from 'moment'
import PropTypes from 'prop-types';

/**
 * Render timeline view
 */
export default class Timeline extends React.Component {

  static propTypes = { 
    /**
     * list of events. [{Message: "Task started by client", Time: 1557262721147300000, Mark: false, MarkDescriptions: Array(0)}]
     */
    Events : PropTypes.array, 
  };

  /**
   * Show descriptions text on marked event
   * 
   * @param {array} descriptions list of notes
   * @returns {HTML}
   */
  getTooltipDescription = (descriptions) => {
   
    return ( <div className={`alert alert-info`}><p>How to fix it:</p><ul>{descriptions.map((desc, i) =><li key={i}><div dangerouslySetInnerHTML={{ __html: desc }} /></li>)}</ul></div>)
  }

  render() {
    return (
      <div className="timeline">
        {this.props.Events.length > 0 ? (
          <ul className="">
            {this.props.Events.map((event, index) => (
              <li key={`${index}_task`} className={event.Marked ? "marked" : ""}>
                <span className="float-right">
                  {moment(event.Time / 1000000).format("MM/DD/YYYY HH:mm:ss")}
                </span>
                {event.Marked ? (
                  <div>
                    <span>{event.Message}</span>
                    {event.MarkDescriptions.length > 0 &&
                      this.getTooltipDescription(event.MarkDescriptions)}
                  </div>
                ) : (
                  <span>{event.Message}</span>
                )}
              </li>
            ))}
          </ul>
        ) : (
          <div className="empty-box">
            <div>
               <span>Events not received</span>
            </div>
          </div>
        )}
      </div>
      )
  }
}