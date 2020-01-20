const Strings = {
    /**
     * Convert full nomad id to prefix. 0e44bd2a-39ba-923a-8d91-3598a6b03247 -> 0e44bd2a
     * 
     * @param {string} first The First Number
     * @returns {string}
     */
    convertToNomadPrefix: function(id){
        const regex=/([^\]/[-]+)+/;
        var groups = id.match(regex);
        
        return (groups.length > 0) ? groups[0] : id
    },
  
}

export default Strings;