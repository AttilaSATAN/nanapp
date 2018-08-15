import React from 'react';
import { CampaignContext } from './CampaignContext';
import { Link } from 'react-router-dom';

/**
 * @description Responsible of displaying one campaign.
 */
export default class Campaign extends React.Component {
  constructor(props) {
    super();
    this.state = { campaignId: props.match.params.campaignId };
  }

  render() {
    return (
      <CampaignContext.Consumer>
        {value => {
          if (value && value.campaigns) {
            let camp = value.campaigns.filter(c => c._id === this.state.campaignId)[0];

            return (
              <div>
                <h2>{camp.name}</h2>
                <p>Goal: {camp.goal}</p>
                <p>Status: {camp.status}</p>
                <p>Campaign Budget: {camp.total_budget}</p>
                <strong>Platforms</strong>
                <hr/>
                {Object.keys(camp.platforms).map(platform => camp.platforms[platform] ? <p><Link style={{color: value.brandColors[platform]}} to={`/campaign/id/${this.state.campaignId}/platform/${platform}`}>{platform.replace(/^\w/, c => c.toUpperCase())}</Link></p> : '')}
              </div>
            )
          }
          return <div>Loading...</div>
        }}

      </CampaignContext.Consumer>
    )
  }
}