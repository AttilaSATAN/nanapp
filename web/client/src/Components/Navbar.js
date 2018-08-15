import React from 'react';
import {Link} from 'react-router-dom';

import './Navbar.css';

/**
 * @description Navbar holds the navigation buttons.
 */
export default class Navbar extends React.Component {
  render(){
    return <header className="Navbar"><Link to={"/"} className="Navbar-brand">NanApp</Link><Link className="Navbar-button" to={"/campaign/list"}>Campaigns</Link></header>
  }
}