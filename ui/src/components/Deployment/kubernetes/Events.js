import React from 'react'
import PropTypes from 'prop-types';
import { Row, Col ,Collapse} from "react-bootstrap";
import { Scrollbars } from 'react-custom-scrollbars';
import Timeline from 'components/Timeline'
import CollapseOpen from "styles/icons/chevron-down.svg"
import CollapseClose from "styles/icons/chevron-up.svg"
import SVG from 'react-inlinesvg';

/*
* Render deployment data
*/
export default class Events extends React.Component {

  static propTypes = {

    /**
     * Box title
     */
    Title: PropTypes.string,
    /**
     * Event list
     */
    Events: PropTypes.array,
    
  };


  /**
   * section toggle.
   */
  toggleContent = () => {

    this.setState({
      open: !this.state.open,
    });
  }

  state = {
    open: false

  }

  render() {

    return (  

      <div>
        <div className="section-title pointer"  onClick={this.toggleContent.bind(this)} >
          {this.props.Title}
          <div className="pull-right">
            {this.state.open? <SVG src={CollapseOpen} width="24px" /> : <SVG src={CollapseClose} width="24px" />}
          </div>
        </div>
        <div className="box events">
          <Collapse in={this.state.open}>
            <div className="content ">
              <Row>
                <Col sm="12">
                  <Scrollbars style={{ height: 300 }}>
                    <Timeline Events={this.props.Events} />
                  </Scrollbars>
                </Col>
              </Row>
            </div>
          </Collapse>
        </div>
      </div>
      )
  }
}