import React from 'react'
import PropTypes from 'prop-types';
import { Button,  Col} from "react-bootstrap";
import BoxAlert from 'components/BoxAlert';
import LogIntegration from "components/Popup/LogIntegration"
import Statistics from 'utils/Statistics'

/**
 * Render Log links
 */
export default class LogLinks extends React.Component {

  static propTypes = {  
    /**
     * List of logs link
     */
    Links : PropTypes.array,
  };

  /**
   * Open Link in new tab
   */
  onGoTo = (url) => {
    console.log(url)
    Statistics.action({type: "open application log"})
    window.open(url, '_blank');
  }

  /**
   * Show log integration popup
   */
  logConfigurationHowTo = () => {
    this.clickChild()
  }


  render() {
    const showError = (this.props.Links == undefined || this.props.Links.length == 0)? true : false
    
    return (
      <Col sm="4">
      <div className="box">
        <div className="box-title">
          Kibana Logs
        </div>
        <div className="content">
        { showError && 
          <div>
          <BoxAlert TextClick={this.logConfigurationHowTo} Message="Log query not configured"/>
          <LogIntegration Show={click => this.clickChild = click}/>
          </div>

        }
        {!showError && this.props.Links.map((link, index) => 
          <Button key={index} onClick={() => this.onGoTo(link.Url)} variant="link">{link.Query}</Button>
        )}
        </div>
      </div>
    </Col>
      )
  }
}
