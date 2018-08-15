import React from 'react';
import { Link } from 'react-router-dom';
import './CampaignButton.css';

export default class CampaignButton extends React.Component {

  constructor(props){
    super()
    this.state = {
      campaign: props.campaign
    }
  }

  render(){
    return <li><Link className="CampaignButton-link" to={`/campaign/id/${this.state.campaign._id}`}><strong>{this.state.campaign.name}</strong> ({this.state.campaign.status})</Link></li>
  }
}