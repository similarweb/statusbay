import React from "react";
import { Link } from "react-router-dom";
import IconLogo from "styles/icons/logo.svg"
import SVG from 'react-inlinesvg';

/**
 * Render loader
 */
export default class Header extends React.Component {

  render() {
    return (
      <div id="header-login">
        <Link to="/">
          <SVG src={IconLogo} />
          <span>statusbay</span>
        </Link>
      </div>
    );
  }
}
