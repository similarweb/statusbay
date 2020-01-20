import moment from 'moment'

const Time = {
    /**
     * convert unix time step format
     * @param {int} time unix timestamp
     */
    FormatUnixTime: function(time){
        return moment.unix(time).format("MM/DD/YYYY HH:mm:ss");
        }
}

export default Time;