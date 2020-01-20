import React from 'react'
import LoaderSVG from "styles/icons/loader.svg"
import SVG from 'react-inlinesvg';

/**
 * Render loader
 */
export default class Loader extends React.Component {

  render() {
    return (
        <SVG src={LoaderSVG} className="loader"/>
      )
  }
}



