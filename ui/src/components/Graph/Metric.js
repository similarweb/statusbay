import React from "react";
import PropTypes from "prop-types";
import { Col } from "react-bootstrap";
import { metricService } from "services/metrics.service";
import Loader from "components/Loader";
import BoxAlert from "components/BoxAlert";
import Highcharts from "highcharts";
import HighchartsReact from "highcharts-react-official";
import moment from "moment";

/**
 * Render graph
 */
export default class Metric extends React.Component {

  static propTypes = {
    Title: PropTypes.string, // Graph title
    SubTitle: PropTypes.string, // Box subtitle
    Query: PropTypes.string, // Graph query - avg:haproxy.backend.response.time{dc:op-us-east-1}
    From: PropTypes.object, // filter from date. moment object
    To: PropTypes.object, // filter to date. moment object
    PlotLines: PropTypes.array // marked lines - {color: "red", width: 2, value: {MOMENT OBJECT}, dashStyle: "longdashdot"}
  };

  state = {
    /**
     * Graph settings
     */
    chartOptions: {},

    /**
     * If metric requested
     */
    requested: true,
    
    /**
     * Show empty graph data
     */
    showEmptyData: false
  };

  
  /**
   * When component mount
   */
  componentDidMount() {
    //Getting metric data
    metricService
      .metric(
        this.props.Query,
        this.props.From.format("x") / 1000,
        this.props.To.format("x") / 1000
      )
      .then(
        data => {
          if (
            data.error ||
            data.Response == null ||
            (typeof data.Response == "object" && data.Response.length === 0)
          ) {
            this.setState({
              requested: false,
              showEmptyData: true
            });
            return;
          }

          const series = [];
          let interval = 30;
          let scaleFactorText = "";

          // Go over on all metrics response
          data.Response.map(data => {
            const datapoint = [];
            let scale_factor = 1;
            let unitName = "";

            if (scaleFactorText === "") {
              scaleFactorText = this.scaleFactorText(data.metric);
            }
            //If metric have units
            if (data.unit) {
              scale_factor = data.unit[0].scale_factor;
              unitName = data.unit[0].short_name;
            }
            interval = data.interval;

            //Getting all data points
            data.pointlist.map(points => {
              datapoint.push({
                x: points[0],
                y: this.calculateScaleFactor(
                  points[1],
                  scale_factor,
                  data.metric
                ),
                unitName: scaleFactorText === "" ? unitName : scaleFactorText
              });
            });

            const seriesData = this.seriesData(data.metric);

            series.push({
              ...seriesData,
              ...{ name: data.expression, data: datapoint }
            });
          });

          const chartOptions = {
            //General options for the chart.
            chart: {
              type: "spline",
              height: "200",
              zoomType: "x"
            },
            //The chart's main title.
            title: {
              text: ""
            },
            //Normally this is the vertical axis
            yAxis: {
              title: false,
              labels: {
                format: `{value} ${scaleFactorText}`
              }
            },
            //Accessibility options for an axis
            tooltip: {
              shared: false,
              formatter: function() {
                return `<b>${Highcharts.numberFormat(this.y, 2)} ${
                  this.point.unitName
                }</b><br/><b>Time:</b>${moment(this.x).format("HH:mm:ss")}`;
              }
            },
            //Normally this is the horizontal axis
            xAxis: {
              plotLines: this.props.PlotLines,
              min: parseInt(this.props.From.format("x"), 10),
              max: parseInt(this.props.To.format("x"), 10),
              labels: {
                tickInterval: interval,
                formatter: function() {
                  return moment(this.value).format("HH:mm");
                }
              }
            },
            //Series options for specific data and the data itself.
            series: series
          };

          this.setState({
            chartOptions: chartOptions,
            requested: false
          });
        },
        () => {
          this.setState({
            requested: false,
            showEmptyData: true
          });
        }
      );
  }

  seriesData(metricName) {
    var config = {};
    switch (metricName) {
      case "docker.mem.limit":
        config = {
          color: "#FF0000",
          dashStyle: "LongDash",
          marker: {
            enabled: false
          }
        };
        break;
    }
    return config;
  }

  /*
   * Return scale factor text
   */
  scaleFactorText(metricName) {
    var text = "";
    switch (metricName) {
      case "nomad.client.allocs.memory.rss":
      case "docker.mem.limit":
        text = "G";
        break;
    }
    return text;
  }

  /*
   * Calculate date point with scalFactor. by default calculating by Datadog scale factor
   */
  calculateScaleFactor(pointValue, scaleFactor, metricName) {
    var value = pointValue;
    switch (metricName) {
      case "nomad.client.allocs.memory.rss":
      case "docker.mem.limit":
        value = value / 1024 / 1024 / 1024;
        break;
      default:
        value = value * scaleFactor;
    }
    return value;
  }

  /**
   * Show empty data
   */
  showEmptyData() {
    return <BoxAlert Message="NO DATA AVAILABLE" />;
  }

  /**
   * Getting sub text subtitle
   */
  getSubTitleText() {
    if (this.props.SubTitle == "") {
      return "";
    }
    return <span className="sub-title">(service: {this.props.SubTitle})</span>;
  }
  render() {
    return (
      <Col sm="6" className="chart">
        <div className="box">
          <div className="box-title">
            {this.props.Title} {this.getSubTitleText()}
          </div>
          <div className="content ">
            {this.state.requested && (
              <div className="center">
                <Loader />
              </div>
            )}
            {this.state.chartOptions.chart && (
              <HighchartsReact
                highcharts={Highcharts}
                options={this.state.chartOptions}
              />
            )}
            {this.state.showEmptyData && this.showEmptyData()}
          </div>
        </div>
      </Col>
    );
  }
}



