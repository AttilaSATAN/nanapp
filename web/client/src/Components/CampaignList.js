import React from 'react';
import './CampaignList.css';
import CampaignButton from './CampaignButton';
import { CampaignContext } from './CampaignContext';

/**
 * Listing component for campaigns.
 */
export default class CampaignList extends React.Component {
  render() {
    return (
      <CampaignContext.Consumer>
        {value => {
          return (
            value && value.campaigns ? (
              <React.Fragment><ul>{value.campaigns.map(campaign => <CampaignButton campaign={campaign} />)}</ul></React.Fragment>
            ) : (
                <React.Fragment>Loading...</React.Fragment>
              )
          )
        }}
      </CampaignContext.Consumer>)
  }
} 