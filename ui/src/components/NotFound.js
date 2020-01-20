import React from 'react'
import { connect } from "react-redux";
import Statistics from 'utils/Statistics'
import {history} from 'configureStore'


/**
 * Render box alert view
 */
@connect()
export default class NotFound extends React.Component {

  /**
   * When component mount
   */
  componentDidMount() {
    Statistics.action({type: "page not found", path: window.location.href})
    history.push("/")
  }
 
  render() {
    return ("")
  }
}