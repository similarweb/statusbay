import React from "react";
import { alertsService } from "services/alerts.service";
import Loader from "components/Loader";
import BoxError from "components/BoxError";
import moment from "moment";
import { Row } from "react-bootstrap";
import Highcharts from "highcharts/highcharts-gantt";
import HighchartsReact from "highcharts-react-official";
import AlertsIntegration from "components/Popup/AlertsIntegration"
import PropTypes from "prop-types";
import ThumbsUp from "styles/icons/thumbs-up.svg"
import SVG from 'react-inlinesvg';

/**
 * 
 */
export default class Alerts extends React.Component {
  
  static propTypes = {
    DeployTime: PropTypes.number,
    MetricsRange: PropTypes.number,
    AlertTags: PropTypes.string
  };

  state = {
    /**
     * Ajax alert request 
     */
    requested: true,

    /**
     * If ajax request return error
     */
    requestError: false,

    /**
     * List of alerts
     */
    alerts: [],

    /**
     * Default count alerts
     */
    alertCount: 0
  };

  /**
   * When component mount
   */
  componentDidMount() {
    const from = moment.unix(this.props.DeployTime).subtract(this.props.MetricsRange, "m").format("x") / 1000;
    const to = moment.unix(this.props.DeployTime).add(this.props.MetricsRange, "m").format("x") / 1000;

    console.log(this.props.AlertTags)
    if (this.isSetTags()) {
      this.getAlerts(this.props.AlertTags, from, to);
    }
  }

  /**
   * Is tags sent to deployment
   */
  isSetTags() {    
    return this.props.AlertTags != undefined && this.props.AlertTags != "";
  }
  
  /**
   * Show alert integration popup
   */
  alertConfigurationHowTo = () => {
    this.clickChild()
  }
  
  /**
   * Get alerts
   * @param {string} tags - tags list
   * @param {int} from - unix time stemp
   * @param {int} to - unix time step  
   */
  getAlerts(tags, from, to) {
    alertsService.alerts(tags, from, to).then(
      data => {
        const categories = [];
        const seriesData = [];
        let alertCount = 0;

        data.map(check => {
          if (check.Periods.length > 0) {
            categories.push({url: check.URL, name:check.Name});
            check.Periods.map(Period => {
              if (Period.Status == "Down") {
                seriesData.push({
                  x: Period.StartUnix * 1000, 
                  x2: Period.EndUnix * 1000,
                  description: Period.Description,
                  y: alertCount,
                  color: "red"
                });
              }
            });
            alertCount += 1;
          }
        });

        this.setState({alertCount, seriesData})

        const chartOptions = {
          chart: {type: "xrange" },
          title: {text: ""},
          xAxis: {
            type: "datetime",
            labels: {
              formatter: function() {
                return moment(new Date(this.value)).format("HH:mm");
              }
            },
            plotLines: [{
                color: "red",
                width: 2,
                value: moment.unix(this.props.DeployTime),
                dashStyle: "longdashdot"
              },
              {
                color: "blue",
                width: 2,
                value: moment(),
                
                label: {
                  text: 'Now',
                }
               
              }
            ],
            min: parseInt(moment.unix(this.props.DeployTime).subtract(this.props.MetricsRange, "m").format("x"),10),
            max: parseInt(moment.unix(this.props.DeployTime).add(this.props.MetricsRange, "m").format("x"),10)
          },
          tooltip: {
            shared: false,
            formatter: function() {
              return `<b>${this.yCategory.name}</b><br/><b>Downtime:  </b>${this.point.description}`;
            }
          },
          yAxis: {
            title: {
              text: ""
            },
            labels: { 
              formatter: function() { 
                return `<a target="_blank" href="${this.value.url}" ">${this.value.name}</span>`; 
              }, 
              useHTML: true 
              },
            categories: categories,
            reversed: true
          },
          series: [{
              pointPadding: 0,
              groupPadding: 10,
              borderColor: "gray",
              pointWidth: 20,
              showInLegend: false,
              data: seriesData
            }
          ]
        };

        this.setState({
          requested: false,
          alerts: seriesData,
          chartOptions: chartOptions
        });
      },
      () => {
        this.setState({ requested: false, requestError: true });        
      }
    );
  }

  render() {
    return (
      <div>
        {!this.isSetTags() && (
          <Row>
            <BoxError TextClick={this.alertConfigurationHowTo} ColWidth={12} Message={`Click to integration`} />
            <AlertsIntegration Show={click => this.clickChild = click}/>
          </Row>
        )}
        
        {this.state.requested && this.isSetTags() && <div className="center"><Loader /></div>}
        {this.state.requestError && 
          <Row>
            <BoxError ColWidth={12} Message={`NO DATA AVAILABLE`} />
          </Row>
        }

        {!this.state.requested && !this.state.requestError && this.isSetTags() &&

            <div className="box">
              <div className="box-title">Statuscake</div>
              <div className="content alert-graph">
                {this.state.alerts.length > 0 ? 
                <HighchartsReact
                  highcharts={Highcharts}
                  containerProps={{ style: { height: `${80* this.state.alertCount}px` } }}
                  options={this.state.chartOptions}
                />
                : 
                <div className="center thumbs-up">
                  <SVG src={ThumbsUp} width="24"/> <span>Alerts not found</span>
                  </div>
              }
              </div>
            </div>
        }   
      </div>
    );
  }
}
