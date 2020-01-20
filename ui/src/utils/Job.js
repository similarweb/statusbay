import React from 'react';

class Job { 
    
    GetMarker(type){
        let markerType = ""
        switch (type.toLowerCase()){
            case "edited":
                markerType = "+/-"
                break;
            case "deleted":
                markerType = "-"
                break;
            case "added":
                markerType = "+"
                break;
        }
        return <span className={`marker is-${type.toLowerCase()}`}>{markerType}</span>
    }

    
    GetBaseURL(region){
        return `${"<<NOMAD_ADDRESS>>".replace("service.consul", `service.${region}.consul`)}/ui`
    }
    
    JobURL(region, id,){
        return `${this.GetBaseURL(region)}/jobs/${id}`
    }

    GetAllocationLogURL(region, allocation, task){
        
        return `${this.GetBaseURL(region)}/allocations/${allocation}/${task}/logs`
    }
  
}

export default new Job();