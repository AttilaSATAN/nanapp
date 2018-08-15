import React, { Component } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import './App.css';
import { Provider } from './CampaignContext';
import Navbar from './Navbar';
import Home from './Home';
import CampaignList from './CampaignList';
import Campaign from './Campaign';
import Platform from './Platform';

/**
 * @description Main wrapper and hub of all application. All main DOM wrapped by `Router` and `CampaignContext`.
 * `App` also responsible for the routing.
 * 
 */
class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <Provider>
          <Navbar />
          <div className="App">
            <Switch>
              <Route exact path='/' component={Home}></Route>
              <Route exect path='/campaign/list' component={CampaignList}></Route>
              <Route exact path='/campaign/id/:campaignId' component={Campaign}></Route>
              <Route exact path='/campaign/id/:campaignId/platform/:platformName' component={Platform}></Route>
            </Switch>
          </div>
          <footer><a href="github.com/attilasatan/nanapp">@github</a></footer>
        </Provider>
      </BrowserRouter>
    );
  }
}

export default App;
