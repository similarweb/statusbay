class Statistics {

    /**
     * Report page view to analytics system
     * @param {object} params - parameters list
     */
    pageView(params){
        if (typeof params !== 'object'){
            params = {}
        }
            
        this._track("page_view", params)
    }

    /**
     * Report user action to analytics system
     * @param {object} params - parameters list
     */
    action(params){
       this._track("action", params)
    }

    /**
     * Implement track report
     * @param {string} action - action key
     * @param {object} params - parameters list
     */
    _track(action, params){
        params['source'] = "UI"
        params['uri'] = window.location.pathname.replace(/\/([0-9]+)/g, '/*');        
        try{
            window.mixpanel.track(action, params);
        } catch(err) {
            // I don't care about the message
        }
    }
}

export default new Statistics()