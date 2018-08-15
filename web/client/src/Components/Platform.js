import React from 'react';
import { CampaignContext } from './CampaignContext';
import Moment from 'react-moment';

import './Platform.css'

export default class Platform extends React.Component {
  constructor(props) {
    super()
    this.state = {
      ...props.match.params
    }

    console.log(this.state, props.match)
  }
  render() {
    return (
      <CampaignContext.Consumer>
        {value => {
          if (value && value.campaigns) {
            let camp = value.campaigns.filter(c => c._id === this.state.campaignId)[0];

            let platform = value.campaigns.filter(c => c._id === this.state.campaignId)[0].platforms[this.state.platformName];
            return (
              <div className="Platform">
                <div className="Platform-data">
                <h1 style={{color: value.brandColors[this.state.platformName]}}>{this.state.platformName.toUpperCase()}</h1>
                  <p>{platform.creatives.header}</p>
                  <p>{platform.creatives.header_1}</p>
                  <p>{platform.creatives.header_2}</p>
                  <p>Start Date: <Moment>{platform.StartDate}</Moment></p>
                  <p>End Date: <Moment>{platform.EndDate}</Moment></p>
                  <p>Impressions: {platform.insights.impressions}</p>
                  <p>Clicks: {platform.insights.clicks}</p>
                  <p>Web Site Visits: {platform.insights.website_visits}</p>
                  <p>Cost Per Click: $ {platform.insights.cost_per_click}</p>
                  <p>Click Through Rate: {platform.insights.click_through_rate}</p>
                  <p>Advanced KPI 1: {platform.insights.advanced_kpi_1}</p>
                </div>
                <div className="Platform-image"> <img src={`/images/${platform.creatives.image}`} /></div>
              </div>
            )
          }
          return <div>Loading...</div>
        }}
      </CampaignContext.Consumer>
    )
  }
}