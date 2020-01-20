import React from 'react';
import IconCheck from "styles/icons/check.svg"
import IconX from "styles/icons/x.svg"
import IconArrowUp from "styles/icons/arrow-up.svg"
import IconPause from "styles/icons/pause.svg"
import IconTrash from "styles/icons/trash.svg"
import SVG from 'react-inlinesvg';


const Status = {
    
    /**
     * Return class name for deployment status
     * @param {string} status deployment status
     */
    getColorByStatus: function(status){
        const className = "status-icon"
        switch (status) {
            case 'successful':
                return `success ${className}`;
            case 'failed':
                return `danger ${className}`;
            case 'running':
                return `primary ${className}`;
            case 'paused':
                return `secondary ${className}`;
            case 'cancelled':
                return `warning ${className}`;
            default:
                return className;
        }
    },
    /**
     * Return icon element by deployment status
     * @param {string} status deployment status
     */
    getIconByStatus: function(status){
        switch (status) {
            case 'successful':
                return <SVG src={IconCheck} className={this.getColorByStatus(status)} />;
            case 'failed':
                return <SVG src={IconX} className={this.getColorByStatus(status)} />;
            case 'running':
                return <SVG src={IconArrowUp} className={this.getColorByStatus(status)} />;
            case 'paused':
                return <SVG src={IconPause} className={this.getColorByStatus(status)} />;
            case 'cancelled':
                return <SVG src={IconTrash} className={this.getColorByStatus(status)} />;
            default:
                return status;
        }
    },
  
}

export default Status;