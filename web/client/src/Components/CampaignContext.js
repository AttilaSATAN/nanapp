import React from 'react';

/**
 * @description Data holding context for campaign.
 */
export const CampaignContext = React.createContext();

/**
 * @description Context provider for campaign data.
 */
export class Provider extends React.Component {

  componentWillMount() {
    fetch("/api/campaign")
      .then(res => res.json())
      .then(data => {
        this.setState({
          campaigns: data,
          brandColors: {
            facebook: '#4267b2',
            instagram: '#000000',
            google: '#ea4335'
          }
        })
      });
  }
  render() {
    return <CampaignContext.Provider value={this.state}>{this.props.children}</CampaignContext.Provider>
  }
}
