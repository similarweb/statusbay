module.exports = function mocks(app) {


  app.get("/api/v1/alerts", (req, res) => {
    setTimeout(function() {
      res.json(200,
        [
          {
            "URL": "",
            "ID": 3636989,
            "Name": "Lite: Website Analysis (us-west-2)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560855567,
                "EndUnix": 1560886013,
                "Description": "8 Hours 27 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560855341,
                "EndUnix": 1560855567,
                "Description": "3 Minutes "
              }
            ]
          },
          {
            "URL": "",
            "ID": 3657560,
            "Name": "HttpMonitoring: Is healthy (us-west-2)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3588226,
            "Name": "AutoComplete: Web (us-east-1)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3588353,
            "Name": "links.similarweb.com/healthcheck",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3637137,
            "Name": "Lite: Top Websites (us-west-2)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560855247,
                "EndUnix": 1560886013,
                "Description": "8 Hours 32 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560855056,
                "EndUnix": 1560855247,
                "Description": "3 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560854665,
                "EndUnix": 1560855056,
                "Description": "6 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560854603,
                "EndUnix": 1560854665,
                "Description": "1 Minutes "
              }
            ]
          },
          {
            "URL": "",
            "ID": 3637201,
            "Name": "Lite: Top Apps: App Store (us-east-1)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560855048,
                "EndUnix": 1560886013,
                "Description": "8 Hours 36 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560854935,
                "EndUnix": 1560855048,
                "Description": "1 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560854702,
                "EndUnix": 1560854935,
                "Description": "3 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560854658,
                "EndUnix": 1560854702,
                "Description": "44 Seconds "
              },
              {
                "Status": "Up",
                "StartUnix": 1560853491,
                "EndUnix": 1560854658,
                "Description": "19 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560853397,
                "EndUnix": 1560853491,
                "Description": "1 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560849584,
                "EndUnix": 1560853397,
                "Description": "1 Hours 3 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560849536,
                "EndUnix": 1560849584,
                "Description": "48 Seconds "
              }
            ]
          },
          {
            "URL": "",
            "ID": 3637075,
            "Name": "Lite: Mobile Analysis - App Store (us-east-1)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560855346,
                "EndUnix": 1560886013,
                "Description": "8 Hours 31 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560855184,
                "EndUnix": 1560855346,
                "Description": "2 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560855010,
                "EndUnix": 1560855184,
                "Description": "2 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560854854,
                "EndUnix": 1560855010,
                "Description": "2 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560854715,
                "EndUnix": 1560854854,
                "Description": "2 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560854599,
                "EndUnix": 1560854715,
                "Description": "1 Minutes "
              }
            ]
          },
          {
            "URL": "",
            "ID": 3637121,
            "Name": "Lite: Top Websites (us-east-1)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560855291,
                "EndUnix": 1560886013,
                "Description": "8 Hours 32 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560855192,
                "EndUnix": 1560855291,
                "Description": "1 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560855048,
                "EndUnix": 1560855192,
                "Description": "2 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560854902,
                "EndUnix": 1560855048,
                "Description": "2 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560853309,
                "EndUnix": 1560854902,
                "Description": "26 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560853144,
                "EndUnix": 1560853309,
                "Description": "2 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560851557,
                "EndUnix": 1560853144,
                "Description": "26 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560851303,
                "EndUnix": 1560851557,
                "Description": "4 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560849893,
                "EndUnix": 1560851303,
                "Description": "23 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560849784,
                "EndUnix": 1560849893,
                "Description": "1 Minutes "
              }
            ]
          },
          {
            "URL": "",
            "ID": 3637252,
            "Name": "Lite: Mobile Analysis - Android (us-east-1)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560855450,
                "EndUnix": 1560886013,
                "Description": "8 Hours 29 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560855183,
                "EndUnix": 1560855450,
                "Description": "4 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560855007,
                "EndUnix": 1560855183,
                "Description": "2 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560854816,
                "EndUnix": 1560855007,
                "Description": "3 Minutes "
              }
            ]
          },
          {
            "URL": "",
            "ID": 3637161,
            "Name": "Lite: Top Apps: Android (us-east-1)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560855349,
                "EndUnix": 1560886013,
                "Description": "8 Hours 31 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560855228,
                "EndUnix": 1560855349,
                "Description": "2 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560854996,
                "EndUnix": 1560855228,
                "Description": "3 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560854821,
                "EndUnix": 1560854996,
                "Description": "2 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560854713,
                "EndUnix": 1560854821,
                "Description": "1 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560854605,
                "EndUnix": 1560854713,
                "Description": "1 Minutes "
              }
            ]
          },
          {
            "URL": "",
            "ID": 3667702,
            "Name": "Lite: Autocomplete Apps (us-east-1)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3636684,
            "Name": "Lite: Website Analysis (us-east-1)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560855023,
                "EndUnix": 1560886013,
                "Description": "8 Hours 36 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560854848,
                "EndUnix": 1560855023,
                "Description": "2 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560854720,
                "EndUnix": 1560854848,
                "Description": "2 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560854532,
                "EndUnix": 1560854720,
                "Description": "3 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560852867,
                "EndUnix": 1560854532,
                "Description": "27 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560852731,
                "EndUnix": 1560852867,
                "Description": "2 Minutes "
              }
            ]
          },
          {
            "URL": "",
            "ID": 3588224,
            "Name": "Internal API: Mobile Analysis - App Store (us-east-1)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560854420,
                "EndUnix": 1560886013,
                "Description": "8 Hours 46 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560853891,
                "EndUnix": 1560854420,
                "Description": "8 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560853790,
                "EndUnix": 1560853891,
                "Description": "1 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560852757,
                "EndUnix": 1560853790,
                "Description": "17 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560852638,
                "EndUnix": 1560852757,
                "Description": "1 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560849999,
                "EndUnix": 1560852638,
                "Description": "43 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560849590,
                "EndUnix": 1560849999,
                "Description": "6 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560849410,
                "EndUnix": 1560849590,
                "Description": "3 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560849277,
                "EndUnix": 1560849410,
                "Description": "2 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560849144,
                "EndUnix": 1560849277,
                "Description": "2 Minutes "
              }
            ]
          },
          {
            "URL": "",
            "ID": 3637220,
            "Name": "Lite: Top Apps: App Store (us-west-2)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3667540,
            "Name": "Lite: AutoComplete: Universal (us-east-1)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3637279,
            "Name": "Lite: Mobile Analysis - Android (us-west-2)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3667511,
            "Name": "Lite: AutoComplete: Web (us-east-1)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3588229,
            "Name": "Lite: Home Page (us-east-1)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3637083,
            "Name": "Lite: Mobile Analysis - App Store (us-west-2)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3657554,
            "Name": "HttpMonitoring: Is healthy (us-east-1)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3637173,
            "Name": "Lite: Top Apps: Android (us-west-2)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3588234,
            "Name": "Internal API: Top Apps: App Store (us-east-1)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560854421,
                "EndUnix": 1560886013,
                "Description": "8 Hours 46 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560853991,
                "EndUnix": 1560854421,
                "Description": "7 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560853794,
                "EndUnix": 1560853991,
                "Description": "3 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560853425,
                "EndUnix": 1560853794,
                "Description": "6 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560853313,
                "EndUnix": 1560853425,
                "Description": "1 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560852678,
                "EndUnix": 1560853313,
                "Description": "10 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560852508,
                "EndUnix": 1560852678,
                "Description": "2 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560850886,
                "EndUnix": 1560852508,
                "Description": "27 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560850823,
                "EndUnix": 1560850886,
                "Description": "1 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560850297,
                "EndUnix": 1560850823,
                "Description": "8 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560850206,
                "EndUnix": 1560850297,
                "Description": "1 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560850082,
                "EndUnix": 1560850206,
                "Description": "2 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560849924,
                "EndUnix": 1560850082,
                "Description": "2 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560849885,
                "EndUnix": 1560849924,
                "Description": "39 Seconds "
              },
              {
                "Status": "Up",
                "StartUnix": 1560849628,
                "EndUnix": 1560849885,
                "Description": "4 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560849507,
                "EndUnix": 1560849628,
                "Description": "2 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560849133,
                "EndUnix": 1560849507,
                "Description": "6 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560849091,
                "EndUnix": 1560849133,
                "Description": "42 Seconds "
              }
            ]
          },
          {
            "URL": "",
            "ID": 3588235,
            "Name": "AutoComplete: Apps (us-east-1)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3588228,
            "Name": "Internal API: Top Websites (us-east-1)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560854421,
                "EndUnix": 1560886013,
                "Description": "8 Hours 46 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560854063,
                "EndUnix": 1560854421,
                "Description": "5 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560854018,
                "EndUnix": 1560854063,
                "Description": "45 Seconds "
              },
              {
                "Status": "Down",
                "StartUnix": 1560853969,
                "EndUnix": 1560854018,
                "Description": "49 Seconds "
              },
              {
                "Status": "Up",
                "StartUnix": 1560853749,
                "EndUnix": 1560853969,
                "Description": "3 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560853490,
                "EndUnix": 1560853749,
                "Description": "4 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560853191,
                "EndUnix": 1560853490,
                "Description": "4 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560853111,
                "EndUnix": 1560853191,
                "Description": "1 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560852929,
                "EndUnix": 1560853111,
                "Description": "3 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560852887,
                "EndUnix": 1560852929,
                "Description": "42 Seconds "
              },
              {
                "Status": "Up",
                "StartUnix": 1560852582,
                "EndUnix": 1560852887,
                "Description": "5 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560851591,
                "EndUnix": 1560852582,
                "Description": "16 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560851414,
                "EndUnix": 1560851591,
                "Description": "2 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560851142,
                "EndUnix": 1560851414,
                "Description": "4 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560850879,
                "EndUnix": 1560851142,
                "Description": "4 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560850703,
                "EndUnix": 1560850879,
                "Description": "2 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560850558,
                "EndUnix": 1560850703,
                "Description": "2 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560850494,
                "EndUnix": 1560850558,
                "Description": "1 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560849925,
                "EndUnix": 1560850494,
                "Description": "9 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560849839,
                "EndUnix": 1560849925,
                "Description": "1 Minutes "
              }
            ]
          },
          {
            "URL": "",
            "ID": 3588231,
            "Name": "Internal API: Top Apps: Android (us-east-1)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560854419,
                "EndUnix": 1560886013,
                "Description": "8 Hours 46 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560853482,
                "EndUnix": 1560854419,
                "Description": "15 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560853210,
                "EndUnix": 1560853482,
                "Description": "4 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560853024,
                "EndUnix": 1560853210,
                "Description": "3 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560852915,
                "EndUnix": 1560853024,
                "Description": "1 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560852837,
                "EndUnix": 1560852915,
                "Description": "1 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560852361,
                "EndUnix": 1560852837,
                "Description": "7 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560850885,
                "EndUnix": 1560852361,
                "Description": "24 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560850810,
                "EndUnix": 1560850885,
                "Description": "1 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560850214,
                "EndUnix": 1560850810,
                "Description": "9 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560849815,
                "EndUnix": 1560850214,
                "Description": "6 Minutes "
              },
              {
                "Status": "Up",
                "StartUnix": 1560849593,
                "EndUnix": 1560849815,
                "Description": "3 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560849487,
                "EndUnix": 1560849593,
                "Description": "1 Minutes "
              }
            ]
          },
          {
            "URL": "",
            "ID": 3588237,
            "Name": "Lite: Captcha Page (us-east-1)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3588243,
            "Name": "Lite: Pricing Page (us-west-2)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3588245,
            "Name": "Internal API: Top Websites (us-west-2)",
            "Periods": []
          },
          {
            "URL": "",
            "ID": 3588246,
            "Name": "Internal API: Top Apps: App Store (us-west-2)",
            "Periods": [
              {
                "Status": "Up",
                "StartUnix": 1560855287,
                "EndUnix": 1560886013,
                "Description": "8 Hours 32 Minutes "
              },
              {
                "Status": "Down",
                "StartUnix": 1560855169,
                "EndUnix": 1560855287,
                "Description": "1 Minutes "
              }
            ]
          }
        ]
      );
    }, 0)
    
  })

  app.get("/api/v1/jobs", (req, res) => {
    setTimeout(function() {
      res.json(200,
        [
          {"Job": "b2c-api-blocksite-staging-op-us-east-1"},
          {"Job": "b2c-b2c-stuff-staging-op-us-east-1"},
          {"Job": "b2c-polygraph-server-production-op-us-east-1"},
          {"Job": "b2c-polygraph-server-staging-op-us-east-1"},
          {"Job": "b2c-selenium-grid-beta-production-op-us-east-1"},
          {"Job": "b2c-simsites-ext-production-op-us-east-1"},
          {"Job": "b2c-userstyles-hb-service-production-op-us-east-1"},
          {"Job": "b2c-userstyles-production-op-us-east-1"},
          {"Job": "b2c-wp-blocksite-co-produciton-op-us-east-1"},
          {"Job": "b2c-wp-blocksite-co-staging-op-us-east-1"},
          {"Job": "b2c-wp-poperblocker-production-op-us-east-1"},
          {"Job": "bi-portal-production-op-us-east-1"},
          {"Job": "bi-portal-staging-op-us-east-1"},
          {"Job": "bigdata-veneur-staging-us-east-1"},
          {"Job": "branchstack-production-op-us-east-1"},
          {"Job": "branchstack-staging-op-us-east-1"},
          {"Job": "datacol-tapi-consumers-fetchcore-aws-production-op-us-east-1"},
          {"Job": "datacol-tapi-consumers-fetchcore-s3-parquet-staging-op-us-east-1"},
          {"Job": "datacol-tapi-consumers-ga-metadata-production-op-us-east-1"},
          {"Job": "datacol-tapi-consumers-mobile-telemetry-prod-production-op-us-east-1"},
          {"Job": "datacollection-ads-html5-staging-op-us-east-1"},
          {"Job": "datacollection-ads-server-dev-op-us-east-1"},
          {"Job": "datacollection-ads-server-staging-op-us-east-1"},
          {"Job": "datacollection-ads-titles-staging-op-us-east-1"},
          {"Job": "datacollection-fcfetcher-staging-op-us-east-1"},
          {"Job": "datacollection-fcparser-staging-op-us-east-1"},
          {"Job": "datacollection-node-chrome-staging-op-us-east-1"},
          {"Job": "datacollection-scrape-service"},
          {"Job": "datacollection-scrape-service-production-op-us-east-1"},
          {"Job": "datacollection-selenium-hub-staging-op-us-east-1"},
          {"Job": "datacollection-tapi-staging-op-us-east-1"},
          {"Job": "datacollection-teamcity-server-us-east-1"},
          {"Job": "datacollection-vpn-e2e-test-us-east-1"},
          {"Job": "datacollection-warden-production-op-us-east-1"},
          {"Job": "di-admin360-production-op-us-east-1"},
          {"Job": "di-admin360-staging-op-us-east-1"},
          {"Job": "di-worker360-production-op-us-east-1"},
          {"Job": "di-worker360-staging-op-us-east-1"},
          {"Job": "lemonbot-production-op-us-east-1"},
          {"Job": "pe-c-test-staging-op-us-east-1"},
          {"Job": "pe-costya-test-production-op-us-east-1"},
          {"Job": "pe-kibana-deprecation-notice-nginx"},
          {"Job": "pe-master-builder-api-production-op-us-east-1"},
          {"Job": "pe-opsdog-production-op-us-east-1"},
          {"Job": "pe-restbase-aws-mrp-us-east-1"},
          {"Job": "pe-restbase-production-us-east-1"},
          {"Job": "pe-restbase-production-us-west-2"},
          {"Job": "pe-s3base-production-op-us-east-1"},
          {"Job": "pe-statusbay-production-op-us-east-1"},
          {"Job": "pe-statusbay-ui-production-op-us-east-1"},
          {"Job": "pe-year-in-review-similarweb-us-east-1"},
          {"Job": "published-data-lake-collections-production-op-us-east-1"},
          {"Job": "published-data-lake-collections-staging-op-us-east-1"},
          {"Job": "web-account-production-op-us-east-1"},
          {"Job": "web-account-production-op-us-west-2"},
          {"Job": "web-account-sandbox-24688-op-us-east-1"},
          {"Job": "web-account-sandbox-24760-op-us-east-1"},
          {"Job": "web-account-sandbox-25098-op-us-east-1"},
          {"Job": "web-account-sandbox-25098dev-op-us-east-1"},
          {"Job": "web-account-sandbox-25098s-op-us-east-1"},
          {"Job": "web-account-sandbox-25193-op-us-east-1"},
          {"Job": "web-account-sandbox-25399-op-us-east-1"},
          {"Job": "web-account-sandbox-addont-op-us-east-1"},
          {"Job": "web-account-sandbox-amazoncat2-op-us-east-1"},
          {"Job": "web-account-sandbox-amazoncat3-op-us-east-1"},
          {"Job": "web-account-sandbox-amazoncatp-op-us-east-1"},
          {"Job": "web-account-sandbox-bisons-op-us-east-1"},
          {"Job": "web-account-sandbox-buyer-op-us-east-1"},
          {"Job": "web-account-sandbox-dads-op-us-east-1"},
          {"Job": "web-account-sandbox-filterfix-op-us-east-1"},
          {"Job": "web-account-sandbox-freeze-op-us-east-1"},
          {"Job": "web-account-sandbox-frozen-op-us-east-1"},
          {"Job": "web-account-sandbox-gpc-op-us-east-1"},
          {"Job": "web-account-sandbox-hakimi25304-op-us-east-1"},
          {"Job": "web-account-sandbox-hakimi25305-op-us-east-1"},
          {"Job": "web-account-sandbox-juliak-op-us-east-1"},
          {"Job": "web-account-sandbox-kwrec-op-us-east-1"},
          {"Job": "web-account-sandbox-ligers-op-us-east-1"},
          {"Job": "web-account-sandbox-maup-op-us-east-1"},
          {"Job": "web-account-sandbox-moster106-op-us-east-1"},
          {"Job": "web-account-sandbox-moster107-op-us-east-1"},
          {"Job": "web-account-sandbox-moster112-op-us-east-1"},
          {"Job": "web-account-sandbox-moster115-op-us-east-1"},
          {"Job": "web-account-sandbox-moster116-op-us-east-1"},
          {"Job": "web-account-sandbox-moster118-op-us-east-1"},
          {"Job": "web-account-sandbox-moster119-op-us-east-1"},
          {"Job": "web-account-sandbox-moster120-op-us-east-1"},
          {"Job": "web-account-sandbox-moster123-op-us-east-1"},
          {"Job": "web-account-sandbox-moster124-op-us-east-1"},
          {"Job": "web-account-sandbox-moster126-op-us-east-1"},
          {"Job": "web-account-sandbox-moster127-op-us-east-1"},
          {"Job": "web-account-sandbox-moster133-op-us-east-1"},
          {"Job": "web-account-sandbox-moster134-op-us-east-1"},
          {"Job": "web-account-sandbox-moster135-op-us-east-1"},
          {"Job": "web-account-sandbox-moster136-op-us-east-1"},
          {"Job": "web-account-sandbox-ptb-op-us-east-1"},
          {"Job": "web-account-sandbox-raccoons-op-us-east-1"},
          {"Job": "web-account-sandbox-sergey-op-us-east-1"},
          {"Job": "web-account-sandbox-services601-op-us-east-1"},
          {"Job": "web-account-sandbox-sim25140-op-us-east-1"},
          {"Job": "web-account-sandbox-skellig-op-us-east-1"},
          {"Job": "web-account-sandbox-snapshot-op-us-east-1"},
          {"Job": "web-account-sandbox-staging-op-us-east-1"},
          {"Job": "web-account-sandbox-tar-op-us-east-1"},
          {"Job": "web-addon-hadorons-production-op-us-east-1"},
          {"Job": "web-addon-production-op-us-east-1"},
          {"Job": "web-addon-production-op-us-west-2"},
          {"Job": "web-addon-sandbox-addongor-op-us-east-1"},
          {"Job": "web-addon-sandbox-addont-op-us-east-1"},
          {"Job": "web-addon-sandbox-freeze-op-us-east-1"},
          {"Job": "web-addon-sandbox-rescache-op-us-east-1"},
          {"Job": "web-addon-sandbox-staging-op-us-east-1"},
          {"Job": "web-api-nginx-production-op-us-west-2"},
          {"Job": "web-api-production-op-us-east-1"},
          {"Job": "web-api-production-op-us-west-2"},
          {"Job": "web-api-sandbox-addont-op-us-east-1"},
          {"Job": "web-api-sandbox-buyer-op-us-east-1"},
          {"Job": "web-api-sandbox-freeze-op-us-east-1"},
          {"Job": "web-api-sandbox-mau-op-us-east-1"},
          {"Job": "web-api-sandbox-maup-op-us-east-1"},
          {"Job": "web-api-sandbox-services601-op-us-east-1"},
          {"Job": "web-billing-sandbox-addont-op-us-east-1"},
          {"Job": "web-billing-sandbox-freeze-op-us-east-1"},
          {"Job": "web-billing-sandbox-services601-op-us-east-1"},
          {"Job": "web-billing-staging-op-us-east-1"},
          {"Job": "web-dtank-production-op-us-east-1"},
          {"Job": "web-dtank-production-op-us-west-2"},
          {"Job": "web-dtank-sandbox-25257-op-us-east-1"},
          {"Job": "web-dtank-staging-op-us-east-1"},
          {"Job": "web-hadorons-production-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-addont-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-bisons-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-freeze-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-frozen-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-igorsitemaps-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster108-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster112-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster115-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster116-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster118-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster119-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster120-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster123-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster124-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster126-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster127-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster133-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster134-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster135-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-moster136-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-services601-op-us-east-1"},
          {"Job": "web-internalapi-sandbox-staging-op-us-east-1"},
          {"Job": "web-lite-nginx-staging-op-us-east-1"},
          {"Job": "web-mixpanel-proxy-nginx-staging-op-us-east-1"},
          {"Job": "web-octopus-service-external-op-us-east-1"},
          {"Job": "web-octopus-service-external-op-us-west-2"},
          {"Job": "web-pdf-sandbox-freeze-op-us-east-1"},
          {"Job": "web-pro-production-op-us-east-1"},
          {"Job": "web-pro-production-op-us-west-2"},
          {"Job": "web-pro-sandbox-24606-op-us-east-1"},
          {"Job": "web-pro-sandbox-24688-op-us-east-1"},
          {"Job": "web-pro-sandbox-24760-op-us-east-1"},
          {"Job": "web-pro-sandbox-24910u-op-us-east-1"},
          {"Job": "web-pro-sandbox-25098-op-us-east-1"},
          {"Job": "web-pro-sandbox-25098dev-op-us-east-1"},
          {"Job": "web-pro-sandbox-25098s-op-us-east-1"},
          {"Job": "web-pro-sandbox-25193-op-us-east-1"},
          {"Job": "web-pro-sandbox-25257-op-us-east-1"},
          {"Job": "web-pro-sandbox-25399-op-us-east-1"},
          {"Job": "web-pro-sandbox-addont-op-us-east-1"},
          {"Job": "web-pro-sandbox-amazoncat2-op-us-east-1"},
          {"Job": "web-pro-sandbox-amazoncat3-op-us-east-1"},
          {"Job": "web-pro-sandbox-amazoncatp-op-us-east-1"},
          {"Job": "web-pro-sandbox-avishai-op-us-east-1"},
          {"Job": "web-pro-sandbox-bads-op-us-east-1"},
          {"Job": "web-pro-sandbox-bisons-op-us-east-1"},
          {"Job": "web-pro-sandbox-buyer-op-us-east-1"},
          {"Job": "web-pro-sandbox-dads-op-us-east-1"},
          {"Job": "web-pro-sandbox-dadss-op-us-east-1"},
          {"Job": "web-pro-sandbox-filterfix-op-us-east-1"},
          {"Job": "web-pro-sandbox-freeze-op-us-east-1"},
          {"Job": "web-pro-sandbox-freeze2-op-us-east-1"},
          {"Job": "web-pro-sandbox-frozen-op-us-east-1"},
          {"Job": "web-pro-sandbox-gatoggle-op-us-east-1"},
          {"Job": "web-pro-sandbox-gpc-op-us-east-1"},
          {"Job": "web-pro-sandbox-hakimi25304-op-us-east-1"},
          {"Job": "web-pro-sandbox-hakimi25305-op-us-east-1"},
          {"Job": "web-pro-sandbox-juliak-op-us-east-1"},
          {"Job": "web-pro-sandbox-kwrec-op-us-east-1"},
          {"Job": "web-pro-sandbox-ligers-op-us-east-1"},
          {"Job": "web-pro-sandbox-mau-op-us-east-1"},
          {"Job": "web-pro-sandbox-maup-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster106-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster107-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster108-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster112-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster113-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster114-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster115-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster116-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster117-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster118-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster119-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster120-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster121-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster122-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster123-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster124-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster126-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster127-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster133-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster134-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster135-op-us-east-1"},
          {"Job": "web-pro-sandbox-moster136-op-us-east-1"},
          {"Job": "web-pro-sandbox-ortest-op-us-east-1"},
          {"Job": "web-pro-sandbox-ptb-op-us-east-1"},
          {"Job": "web-pro-sandbox-raccoons-op-us-east-1"},
          {"Job": "web-pro-sandbox-sergey-op-us-east-1"},
          {"Job": "web-pro-sandbox-services601-op-us-east-1"},
          {"Job": "web-pro-sandbox-shadi-op-us-east-1"},
          {"Job": "web-pro-sandbox-sim25140-op-us-east-1"},
          {"Job": "web-pro-sandbox-skellig-op-us-east-1"},
          {"Job": "web-pro-sandbox-slingo-op-us-east-1"},
          {"Job": "web-pro-sandbox-snapshot-op-us-east-1"},
          {"Job": "web-pro-sandbox-staging-op-us-east-1"},
          {"Job": "web-pro-sandbox-statcheck-op-us-east-1"},
          {"Job": "web-pro-sandbox-stats2-op-us-east-1"},
          {"Job": "web-pro-sandbox-tar-op-us-east-1"},
          {"Job": "web-reportsmanager-sandbox-24688-op-us-east-1"},
          {"Job": "web-reportsmanager-sandbox-25257-op-us-east-1"},
          {"Job": "web-reportsmanager-sandbox-25399-op-us-east-1"},
          {"Job": "web-reportsmanager-sandbox-addont-op-us-east-1"},
          {"Job": "web-reportsmanager-sandbox-freeze-op-us-east-1"},
          {"Job": "web-reportsmanager-staging-op-us-east-1"},
          {"Job": "web-scheduler-production-op-us-east-1"},
          {"Job": "web-scheduler-sandbox-addont-op-us-east-1"},
          {"Job": "web-scheduler-sandbox-bisons-op-us-east-1"},
          {"Job": "web-scheduler-sandbox-freeze-op-us-east-1"},
          {"Job": "web-scheduler-sandbox-services601-op-us-east-1"},
          {"Job": "web-scheduler-staging-op-us-east-1"},
          {"Job": "web-staging-addon-production-op-us-east-1"},
          {"Job": "web-staging-selenium-grid-duo-staging-op-us-east-1"},
          {"Job": "web-staging-selenium-grid-quattour-staging-op-us-east-1"},
          {"Job": "web-staging-selenium-grid-tribus-staging-op-us-east-1"},
          {"Job": "web-staging-selenium-grid-unus-staging-op-us-east-1"},
          {"Job": "web-userdata-production-op-us-east-1"},
          {"Job": "web-userdata-production-op-us-west-2"},
          {"Job": "web-userdata-sandbox-24688-op-us-east-1"},
          {"Job": "web-userdata-sandbox-24760-op-us-east-1"},
          {"Job": "web-userdata-sandbox-25224-op-us-east-1"},
          {"Job": "web-userdata-sandbox-25399-op-us-east-1"},
          {"Job": "web-userdata-sandbox-addont-op-us-east-1"},
          {"Job": "web-userdata-sandbox-amazoncat2-op-us-east-1"},
          {"Job": "web-userdata-sandbox-amazoncat3-op-us-east-1"},
          {"Job": "web-userdata-sandbox-bads-op-us-east-1"},
          {"Job": "web-userdata-sandbox-bisons-op-us-east-1"},
          {"Job": "web-userdata-sandbox-buyer-op-us-east-1"},
          {"Job": "web-userdata-sandbox-dads-op-us-east-1"},
          {"Job": "web-userdata-sandbox-dadss-op-us-east-1"},
          {"Job": "web-userdata-sandbox-freeze-op-us-east-1"},
          {"Job": "web-userdata-sandbox-frozen-op-us-east-1"},
          {"Job": "web-userdata-sandbox-ligers-op-us-east-1"},
          {"Job": "web-userdata-sandbox-mau-op-us-east-1"},
          {"Job": "web-userdata-sandbox-maup-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster105-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster107-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster111-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster113-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster115-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster116-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster118-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster119-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster120-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster123-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster124-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster126-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster127-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster133-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster134-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster135-op-us-east-1"},
          {"Job": "web-userdata-sandbox-moster136-op-us-east-1"},
          {"Job": "web-userdata-sandbox-services601-op-us-east-1"},
          {"Job": "web-userdata-sandbox-snapshot-op-us-east-1"},
          {"Job": "web-userdata-staging-op-us-east-1"},
          {"Job": "web-wp-similarweb-corp-production-op-us-east-1"},
          {"Job": "web-wp-similarweb-corp-staging-op-us-east-1"},
          ]
      );
    }, 2000)
    
  })

  app.get("/api/v1/deployments/:id/:time", (req, res) => {
    setTimeout(function() {
      res.json(200,
        {
          "Diff": {
            "Type": "Edited",
            "ID": "web-api-nginx-production-op-us-east-1",
            "Fields": null,
            "Objects": null,
            "TaskGroups": [
                {
                    "Type": "Edited",
                    "Name": "web-api-nginx-group",
                    "Fields": [
                        {
                            "Type": "Edited",
                            "Name": "Count",
                            "Old": "3",
                            "New": "4",
                            "Annotations": null
                        },
                        {
                            "Type": "Deleted",
                            "Name": "Meta[test]",
                            "Old": "1",
                            "New": "",
                            "Annotations": null
                        }
                    ],
                    "Objects": null,
                    "Tasks": [
                        {
                            "Type": "Edited",
                            "Name": "web-api-nginx-task",
                            "Fields": [
                                {
                                    "Type": "Deleted",
                                    "Name": "Env[CONSUL_ADDR2]",
                                    "Old": "1234",
                                    "New": "",
                                    "Annotations": null
                                }
                            ],
                            "Objects": [
                                {
                                    "Type": "Edited",
                                    "Name": "Config",
                                    "Fields": [
                                        {
                                            "Type": "None",
                                            "Name": "image",
                                            "Old": "nginx:1.12-alpine",
                                            "New": "nginx:1.12-alpine",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "None",
                                            "Name": "port_map[0][http]",
                                            "Old": "80",
                                            "New": "80",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "None",
                                            "Name": "volumes[0]",
                                            "Old": "local/nginx.conf:/etc/nginx/conf.d/default.conf",
                                            "New": "local/nginx.conf:/etc/nginx/conf.d/default.conf",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Deleted",
                                            "Name": "volumes[1]",
                                            "Old": "local/nginx.conf:/etc/nginx/conf.d/default.conf",
                                            "New": "",
                                            "Annotations": null
                                        }
                                    ],
                                    "Objects": null
                                },
                                {
                                    "Type": "Edited",
                                    "Name": "Resources",
                                    "Fields": [
                                        {
                                            "Type": "Edited",
                                            "Name": "CPU",
                                            "Old": "1023",
                                            "New": "1024",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "None",
                                            "Name": "DiskMB",
                                            "Old": "0",
                                            "New": "0",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "None",
                                            "Name": "IOPS",
                                            "Old": "0",
                                            "New": "0",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "None",
                                            "Name": "MemoryMB",
                                            "Old": "1024",
                                            "New": "1024",
                                            "Annotations": null
                                        }
                                    ],
                                    "Objects": null
                                },
                                {
                                    "Type": "Added",
                                    "Name": "Template",
                                    "Fields": [
                                        {
                                            "Type": "Added",
                                            "Name": "ChangeMode",
                                            "Old": "",
                                            "New": "restart",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "None",
                                            "Name": "ChangeSignal",
                                            "Old": "",
                                            "New": "",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Added",
                                            "Name": "DestPath",
                                            "Old": "",
                                            "New": "local/nginx.conf",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Added",
                                            "Name": "EmbeddedTmpl",
                                            "Old": "",
                                            "New": "server {\n     listen 80;\n     server_name _;\n\n      \n      location / {\n        add_header Content-Type text/plain;\n        return 200 'My ngnix response2';\n\n      }\n    }\n       ",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Added",
                                            "Name": "Envvars",
                                            "Old": "",
                                            "New": "false",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Added",
                                            "Name": "LeftDelim",
                                            "Old": "",
                                            "New": "{{",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Added",
                                            "Name": "Perms",
                                            "Old": "",
                                            "New": "0644",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Added",
                                            "Name": "RightDelim",
                                            "Old": "",
                                            "New": "}}",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "None",
                                            "Name": "SourcePath",
                                            "Old": "",
                                            "New": "",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Added",
                                            "Name": "Splay",
                                            "Old": "",
                                            "New": "5000000000",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Added",
                                            "Name": "VaultGrace",
                                            "Old": "",
                                            "New": "15000000000",
                                            "Annotations": null
                                        }
                                    ],
                                    "Objects": null
                                },
                                {
                                    "Type": "Deleted",
                                    "Name": "Template",
                                    "Fields": [
                                        {
                                            "Type": "Deleted",
                                            "Name": "ChangeMode",
                                            "Old": "restart",
                                            "New": "",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "None",
                                            "Name": "ChangeSignal",
                                            "Old": "",
                                            "New": "",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Deleted",
                                            "Name": "DestPath",
                                            "Old": "local/nginx.conf",
                                            "New": "",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Deleted",
                                            "Name": "EmbeddedTmpl",
                                            "Old": "server {\n     listen 80;\n     server_name _;\n\n      \n      location / {\n        add_header Content-Type text/plain;\n        return 200 'My ngnix response';\n\n      }\n    }\n       ",
                                            "New": "",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Deleted",
                                            "Name": "Envvars",
                                            "Old": "false",
                                            "New": "",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Deleted",
                                            "Name": "LeftDelim",
                                            "Old": "{{",
                                            "New": "",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Deleted",
                                            "Name": "Perms",
                                            "Old": "0644",
                                            "New": "",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Deleted",
                                            "Name": "RightDelim",
                                            "Old": "}}",
                                            "New": "",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "None",
                                            "Name": "SourcePath",
                                            "Old": "",
                                            "New": "",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Deleted",
                                            "Name": "Splay",
                                            "Old": "5000000000",
                                            "New": "",
                                            "Annotations": null
                                        },
                                        {
                                            "Type": "Deleted",
                                            "Name": "VaultGrace",
                                            "Old": "15000000000",
                                            "New": "",
                                            "Annotations": null
                                        }
                                    ],
                                    "Objects": null
                                }
                            ],
                            "Annotations": null
                        }
                    ],
                    "Updates": null
                }
            ]
        }, 
          "Name": "web-staging-selenium-grid-unus-staging-op-us-east-1",
          "ID": "d6ac8ffe-09bb-bad8-d907-6a7fac5e2801",
          "Status": "failed",
          "StatusDescription": "Failed due to progress deadline - no stable job version to auto revert to",
          "DeployBy": "",
          "ReportTo": [
            "#elad_test"
          ],
          "AlertTags": "",
          "Url": "http://nomad-server-production.service.op-us-east-1.consul:4646/ui/jobs/web-staging-selenium-grid-unus-staging-op-us-east-1",
          "Region": "op-us-east-1",
          "DeployTime": 1560852297,
          "MetricsRange": 60,
          "Events": {
            "Task": [
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "0f461421-2189-1665-f64a-ae614028a98f",
                "Events": [
                  {
                    "Message": "Task started by client",
                    "Time": 1559880353218453161,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326823995439,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326804835549,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326804562516,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": false
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "264dbbe9-230e-b952-9b19-d14748cf7e1c",
                "Events": [
                  {
                    "Message": "Task not running for min_healthy_time of 30s by deadline",
                    "Time": 1559880386358774062,
                    "Marked": true,
                    "MarkDescriptions": [
                      "Parameter \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#min_healthy_time\" target=\"_blank\"\u003emin_healthy_time\u003c/a\u003e sets the minimum time the allocation must pass its healthcheck in order to be marked healthy, you can mitigate it by: \u003cul\u003e\n  \u003cli\u003eIn most cases, this is caused by \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#healthy_deadline\" target=\"_blank\"\u003ehealthy_deadline\u003c/a\u003e/\u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#progress_deadline\" target=\"_blank\"\u003eprogress_deadline\u003c/a\u003e which are set too low, inspect these parameters compared to the min_healthy_time, increasing them, might help.\u003c/li\u003e\n  \u003cli\u003eDepended systems are responding slowly, for instance, downloading the docker image from the registry, check if something change with the 'download time'.\u003c/li\u003e\n  \u003cli\u003eDocker image size increased significantly, check if something change in this area.\u003c/a\u003e \n\u003c/ul\u003e"
                    ]
                  },
                  {
                    "Message": "Task started by client",
                    "Time": 1559880356574999195,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326374427682,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326358783742,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326358493582,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": true
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "335bb2d7-0ec2-d2ad-1535-ecab2cf7f747",
                "Events": [
                  {
                    "Message": "Task started by client",
                    "Time": 1559880348362517589,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326831604984,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326816006414,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326815756585,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": false
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "413c6384-9aaf-f2a5-1915-3b3d590bf0ed",
                "Events": [
                  {
                    "Message": "Task not running for min_healthy_time of 30s by deadline",
                    "Time": 1559880386813144641,
                    "Marked": true,
                    "MarkDescriptions": [
                      "Parameter \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#min_healthy_time\" target=\"_blank\"\u003emin_healthy_time\u003c/a\u003e sets the minimum time the allocation must pass its healthcheck in order to be marked healthy, you can mitigate it by: \u003cul\u003e\n  \u003cli\u003eIn most cases, this is caused by \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#healthy_deadline\" target=\"_blank\"\u003ehealthy_deadline\u003c/a\u003e/\u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#progress_deadline\" target=\"_blank\"\u003eprogress_deadline\u003c/a\u003e which are set too low, inspect these parameters compared to the min_healthy_time, increasing them, might help.\u003c/li\u003e\n  \u003cli\u003eDepended systems are responding slowly, for instance, downloading the docker image from the registry, check if something change with the 'download time'.\u003c/li\u003e\n  \u003cli\u003eDocker image size increased significantly, check if something change in this area.\u003c/a\u003e \n\u003c/ul\u003e"
                    ]
                  },
                  {
                    "Message": "Task started by client",
                    "Time": 1559880357542358657,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326827743271,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326812949553,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326812737063,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": true
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "469f51d1-0513-4a17-6f4e-e562c4e14918",
                "Events": [
                  {
                    "Message": "Task not running for min_healthy_time of 30s by deadline",
                    "Time": 1559880386370919589,
                    "Marked": true,
                    "MarkDescriptions": [
                      "Parameter \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#min_healthy_time\" target=\"_blank\"\u003emin_healthy_time\u003c/a\u003e sets the minimum time the allocation must pass its healthcheck in order to be marked healthy, you can mitigate it by: \u003cul\u003e\n  \u003cli\u003eIn most cases, this is caused by \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#healthy_deadline\" target=\"_blank\"\u003ehealthy_deadline\u003c/a\u003e/\u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#progress_deadline\" target=\"_blank\"\u003eprogress_deadline\u003c/a\u003e which are set too low, inspect these parameters compared to the min_healthy_time, increasing them, might help.\u003c/li\u003e\n  \u003cli\u003eDepended systems are responding slowly, for instance, downloading the docker image from the registry, check if something change with the 'download time'.\u003c/li\u003e\n  \u003cli\u003eDocker image size increased significantly, check if something change in this area.\u003c/a\u003e \n\u003c/ul\u003e"
                    ]
                  },
                  {
                    "Message": "Task started by client",
                    "Time": 1559880356611185748,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326387721818,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326370966754,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326370692343,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": true
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "511a551a-ec11-df51-182c-05cfced70f4b",
                "Events": [
                  {
                    "Message": "Task started by client",
                    "Time": 1559880347713682082,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326373967047,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326358541438,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326358302811,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": false
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "6d3fa520-ba5f-e7b6-6e46-216d2ce7a437",
                "Events": [
                  {
                    "Message": "Task not running for min_healthy_time of 30s by deadline",
                    "Time": 1559880386312595756,
                    "Marked": true,
                    "MarkDescriptions": [
                      "Parameter \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#min_healthy_time\" target=\"_blank\"\u003emin_healthy_time\u003c/a\u003e sets the minimum time the allocation must pass its healthcheck in order to be marked healthy, you can mitigate it by: \u003cul\u003e\n  \u003cli\u003eIn most cases, this is caused by \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#healthy_deadline\" target=\"_blank\"\u003ehealthy_deadline\u003c/a\u003e/\u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#progress_deadline\" target=\"_blank\"\u003eprogress_deadline\u003c/a\u003e which are set too low, inspect these parameters compared to the min_healthy_time, increasing them, might help.\u003c/li\u003e\n  \u003cli\u003eDepended systems are responding slowly, for instance, downloading the docker image from the registry, check if something change with the 'download time'.\u003c/li\u003e\n  \u003cli\u003eDocker image size increased significantly, check if something change in this area.\u003c/a\u003e \n\u003c/ul\u003e"
                    ]
                  },
                  {
                    "Message": "Task started by client",
                    "Time": 1559880359638024721,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326329271588,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326312821307,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326312510415,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": true
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "70021742-f5b1-d4d5-49bc-d6bc288f1f51",
                "Events": [
                  {
                    "Message": "Task not running for min_healthy_time of 30s by deadline",
                    "Time": 1559880386800988512,
                    "Marked": true,
                    "MarkDescriptions": [
                      "Parameter \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#min_healthy_time\" target=\"_blank\"\u003emin_healthy_time\u003c/a\u003e sets the minimum time the allocation must pass its healthcheck in order to be marked healthy, you can mitigate it by: \u003cul\u003e\n  \u003cli\u003eIn most cases, this is caused by \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#healthy_deadline\" target=\"_blank\"\u003ehealthy_deadline\u003c/a\u003e/\u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#progress_deadline\" target=\"_blank\"\u003eprogress_deadline\u003c/a\u003e which are set too low, inspect these parameters compared to the min_healthy_time, increasing them, might help.\u003c/li\u003e\n  \u003cli\u003eDepended systems are responding slowly, for instance, downloading the docker image from the registry, check if something change with the 'download time'.\u003c/li\u003e\n  \u003cli\u003eDocker image size increased significantly, check if something change in this area.\u003c/a\u003e \n\u003c/ul\u003e"
                    ]
                  },
                  {
                    "Message": "Task started by client",
                    "Time": 1559880357597689628,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326816563854,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326800851866,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326800583312,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": true
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "73d9efb6-c53e-e170-48bd-e0935ec50d25",
                "Events": [
                  {
                    "Message": "Task started by client",
                    "Time": 1559880347148229804,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326816699991,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326801047711,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326800842719,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": false
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "917aff8f-e896-d997-a1f1-2fd87eae9c3d",
                "Events": [
                  {
                    "Message": "Task not running for min_healthy_time of 30s by deadline",
                    "Time": 1559880386355726840,
                    "Marked": true,
                    "MarkDescriptions": [
                      "Parameter \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#min_healthy_time\" target=\"_blank\"\u003emin_healthy_time\u003c/a\u003e sets the minimum time the allocation must pass its healthcheck in order to be marked healthy, you can mitigate it by: \u003cul\u003e\n  \u003cli\u003eIn most cases, this is caused by \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#healthy_deadline\" target=\"_blank\"\u003ehealthy_deadline\u003c/a\u003e/\u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#progress_deadline\" target=\"_blank\"\u003eprogress_deadline\u003c/a\u003e which are set too low, inspect these parameters compared to the min_healthy_time, increasing them, might help.\u003c/li\u003e\n  \u003cli\u003eDepended systems are responding slowly, for instance, downloading the docker image from the registry, check if something change with the 'download time'.\u003c/li\u003e\n  \u003cli\u003eDocker image size increased significantly, check if something change in this area.\u003c/a\u003e \n\u003c/ul\u003e"
                    ]
                  },
                  {
                    "Message": "Task started by client",
                    "Time": 1559880358313154969,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326383727460,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326355846906,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326355611653,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": true
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "96d2f760-feda-34d2-99e5-1de71a59308b",
                "Events": [
                  {
                    "Message": "Task started by client",
                    "Time": 1559880347963032209,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326371236590,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326356337905,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326356079865,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": false
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "9e90e179-c3dc-84c4-350e-4b3d97175c78",
                "Events": [
                  {
                    "Message": "Task started by client",
                    "Time": 1559880355080888624,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326391170273,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326372338896,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326372097931,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": false
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "a05c7df2-b3a1-3a5e-48d4-4ec747b1a570",
                "Events": [
                  {
                    "Message": "Task not running for min_healthy_time of 30s by deadline",
                    "Time": 1559880386816376211,
                    "Marked": true,
                    "MarkDescriptions": [
                      "Parameter \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#min_healthy_time\" target=\"_blank\"\u003emin_healthy_time\u003c/a\u003e sets the minimum time the allocation must pass its healthcheck in order to be marked healthy, you can mitigate it by: \u003cul\u003e\n  \u003cli\u003eIn most cases, this is caused by \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#healthy_deadline\" target=\"_blank\"\u003ehealthy_deadline\u003c/a\u003e/\u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#progress_deadline\" target=\"_blank\"\u003eprogress_deadline\u003c/a\u003e which are set too low, inspect these parameters compared to the min_healthy_time, increasing them, might help.\u003c/li\u003e\n  \u003cli\u003eDepended systems are responding slowly, for instance, downloading the docker image from the registry, check if something change with the 'download time'.\u003c/li\u003e\n  \u003cli\u003eDocker image size increased significantly, check if something change in this area.\u003c/a\u003e \n\u003c/ul\u003e"
                    ]
                  },
                  {
                    "Message": "Task started by client",
                    "Time": 1559880358844108410,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326834301524,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326817666074,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326817009170,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": true
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "adba3565-dbc7-5dbf-df83-fa3ec3dedc67",
                "Events": [
                  {
                    "Message": "Task not running for min_healthy_time of 30s by deadline",
                    "Time": 1559880386369510392,
                    "Marked": true,
                    "MarkDescriptions": [
                      "Parameter \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#min_healthy_time\" target=\"_blank\"\u003emin_healthy_time\u003c/a\u003e sets the minimum time the allocation must pass its healthcheck in order to be marked healthy, you can mitigate it by: \u003cul\u003e\n  \u003cli\u003eIn most cases, this is caused by \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#healthy_deadline\" target=\"_blank\"\u003ehealthy_deadline\u003c/a\u003e/\u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#progress_deadline\" target=\"_blank\"\u003eprogress_deadline\u003c/a\u003e which are set too low, inspect these parameters compared to the min_healthy_time, increasing them, might help.\u003c/li\u003e\n  \u003cli\u003eDepended systems are responding slowly, for instance, downloading the docker image from the registry, check if something change with the 'download time'.\u003c/li\u003e\n  \u003cli\u003eDocker image size increased significantly, check if something change in this area.\u003c/a\u003e \n\u003c/ul\u003e"
                    ]
                  },
                  {
                    "Message": "Task started by client",
                    "Time": 1559880358359989713,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326390502271,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326369684260,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326369410481,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": true
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "bcbbb697-c038-d869-6c85-b3d4cf6079e5",
                "Events": [
                  {
                    "Message": "Task not running for min_healthy_time of 30s by deadline",
                    "Time": 1559880386357992197,
                    "Marked": true,
                    "MarkDescriptions": [
                      "Parameter \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#min_healthy_time\" target=\"_blank\"\u003emin_healthy_time\u003c/a\u003e sets the minimum time the allocation must pass its healthcheck in order to be marked healthy, you can mitigate it by: \u003cul\u003e\n  \u003cli\u003eIn most cases, this is caused by \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#healthy_deadline\" target=\"_blank\"\u003ehealthy_deadline\u003c/a\u003e/\u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#progress_deadline\" target=\"_blank\"\u003eprogress_deadline\u003c/a\u003e which are set too low, inspect these parameters compared to the min_healthy_time, increasing them, might help.\u003c/li\u003e\n  \u003cli\u003eDepended systems are responding slowly, for instance, downloading the docker image from the registry, check if something change with the 'download time'.\u003c/li\u003e\n  \u003cli\u003eDocker image size increased significantly, check if something change in this area.\u003c/a\u003e \n\u003c/ul\u003e"
                    ]
                  },
                  {
                    "Message": "Task started by client",
                    "Time": 1559880356852379556,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326383621152,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326357953189,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326357715088,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": true
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "e5bbbd70-57ba-f125-42cf-affbb6d03547",
                "Events": [
                  {
                    "Message": "Task not running for min_healthy_time of 30s by deadline",
                    "Time": 1559880386803475874,
                    "Marked": true,
                    "MarkDescriptions": [
                      "Parameter \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#min_healthy_time\" target=\"_blank\"\u003emin_healthy_time\u003c/a\u003e sets the minimum time the allocation must pass its healthcheck in order to be marked healthy, you can mitigate it by: \u003cul\u003e\n  \u003cli\u003eIn most cases, this is caused by \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#healthy_deadline\" target=\"_blank\"\u003ehealthy_deadline\u003c/a\u003e/\u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#progress_deadline\" target=\"_blank\"\u003eprogress_deadline\u003c/a\u003e which are set too low, inspect these parameters compared to the min_healthy_time, increasing them, might help.\u003c/li\u003e\n  \u003cli\u003eDepended systems are responding slowly, for instance, downloading the docker image from the registry, check if something change with the 'download time'.\u003c/li\u003e\n  \u003cli\u003eDocker image size increased significantly, check if something change in this area.\u003c/a\u003e \n\u003c/ul\u003e"
                    ]
                  },
                  {
                    "Message": "Task started by client",
                    "Time": 1559880358806296210,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326820770852,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326803715211,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326803421202,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": true
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "ead0eabc-0ccf-70ef-a4e4-0f2595a07896",
                "Events": [
                  {
                    "Message": "Task started by client",
                    "Time": 1559880348523611006,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326818871334,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326804114075,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326803883163,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": false
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "f2087b68-065f-7b4f-a805-9d77a9ca57f8",
                "Events": [
                  {
                    "Message": "Task not running for min_healthy_time of 30s by deadline",
                    "Time": 1559880386300132381,
                    "Marked": true,
                    "MarkDescriptions": [
                      "Parameter \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#min_healthy_time\" target=\"_blank\"\u003emin_healthy_time\u003c/a\u003e sets the minimum time the allocation must pass its healthcheck in order to be marked healthy, you can mitigate it by: \u003cul\u003e\n  \u003cli\u003eIn most cases, this is caused by \u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#healthy_deadline\" target=\"_blank\"\u003ehealthy_deadline\u003c/a\u003e/\u003ca href=\"https://www.nomadproject.io/docs/job-specification/update.html#progress_deadline\" target=\"_blank\"\u003eprogress_deadline\u003c/a\u003e which are set too low, inspect these parameters compared to the min_healthy_time, increasing them, might help.\u003c/li\u003e\n  \u003cli\u003eDepended systems are responding slowly, for instance, downloading the docker image from the registry, check if something change with the 'download time'.\u003c/li\u003e\n  \u003cli\u003eDocker image size increased significantly, check if something change in this area.\u003c/a\u003e \n\u003c/ul\u003e"
                    ]
                  },
                  {
                    "Message": "Task started by client",
                    "Time": 1559880359468349311,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326317376892,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326300254572,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326300002491,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": true
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "f5946d84-5d0d-78d2-57a1-9f693f68bd82",
                "Events": [
                  {
                    "Message": "Task started by client",
                    "Time": 1559880347956841319,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326388446364,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326368672914,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326368393142,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": false
              },
              {
                "Name": "web-staging-selenium-grid-unus-chrome-node-task",
                "AllocationID": "f5bee92c-4968-17ac-2c8e-70a6ac7e1990",
                "Events": [
                  {
                    "Message": "Task started by client",
                    "Time": 1559880354954874204,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image docker-registry.similarweb.io/selenium-chrome-node-google-chrome-stable:B38.master.027ed80",
                    "Time": 1559880326380201180,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326359618491,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326359416017,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": false
              },
              {
                "Name": "web-staging-selenium-grid-unus-hub-task",
                "AllocationID": "f6128f19-6d35-84de-286c-500cec0480c0",
                "Events": [
                  {
                    "Message": "Task started by client",
                    "Time": 1559880352403024398,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Downloading image selenium/hub:3.8.1-erbium",
                    "Time": 1559880326832483108,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Building Task Directory",
                    "Time": 1559880326818573350,
                    "Marked": false,
                    "MarkDescriptions": []
                  },
                  {
                    "Message": "Task received by client",
                    "Time": 1559880326817848287,
                    "Marked": false,
                    "MarkDescriptions": []
                  }
                ],
                "Marked": false
              }
            ]
          },
          "DatadogLinks": [
            {
              "Group": "web-staging-selenium-grid-unus-hub-group",
              "Task": "web-staging-selenium-grid-unus-hub-task",
              "Service": "web-staging-selenium-hub-unus-staging",
              "Url": "https://app.datadoghq.com/dashboard/frg-4hs-xa2/pe-nomad-job-inspector?\u0026tpl_var_dc=op-us-east-1\u0026tpl_var_job=web-staging-selenium-grid-unus-staging-op-us-east-1\u0026tpl_var_group=web-staging-selenium-grid-unus-hub-group\u0026tpl_var_task=web-staging-selenium-grid-unus-hub-task\u0026tpl_var_haproxy_backend=be_web-staging-selenium-hub-unus-staging\u0026tpl_var_worker_group=web-staging"
            },
            {
              "Group": "web-staging-selenium-grid-unus-chrome-node-group",
              "Task": "web-staging-selenium-grid-unus-chrome-node-task",
              "Service": "web-staging-selenium-chrome-node-unus-staging",
              "Url": "https://app.datadoghq.com/dashboard/frg-4hs-xa2/pe-nomad-job-inspector?\u0026tpl_var_dc=op-us-east-1\u0026tpl_var_job=web-staging-selenium-grid-unus-staging-op-us-east-1\u0026tpl_var_group=web-staging-selenium-grid-unus-chrome-node-group\u0026tpl_var_task=web-staging-selenium-grid-unus-chrome-node-task\u0026tpl_var_haproxy_backend=be_web-staging-selenium-chrome-node-unus-staging\u0026tpl_var_worker_group=web-staging"
            }
          ],
          "OctopusLinks": [
            "https://web-staging-selenium-hub-unus-staging.op-us-east-1.web-staging-grid.int.similarweb.io",
            "https://web-staging-selenium-chrome-node-unus-staging.op-us-east-1.web-staging-grid.int.similarweb.io"
          ],
          "LogLinks": [
            {
              "Query": "application: pro.similarweb AND dc: op-us-east-1 AND deployEnv: production",
              "URL": "https://app.logz.io/#/dashboard/kibana/discover?_a=(columns:!(short_message,message),query:(query:'application: pro.similarweb AND dc: op-us-east-1 AND deployEnv: production'),sort:!('@timestamp',desc))\u0026_g=(time:(from:'2019-07-21T07:08:15Z',to:'2019-07-21T09:08:15Z'))"
            },
            {
              "Query": "application: pro.similarweb AND dc: op-us-east-1 AND deployEnv: production",
              "URL": "https://app.logz.io/#/dashboard/kibana/discover?_a=(columns:!(short_message,message),query:(query:'application: pro.similarweb AND dc: op-us-east-1 AND deployEnv: production'),sort:!('@timestamp',desc))\u0026_g=(time:(from:'2019-07-21T07:08:15Z',to:'2019-07-21T09:08:15Z'))"
            }
          ],
          "Metrics": {
            "web-staging-selenium-grid-unus-chrome-node-task": [
              {
                "Query": "sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,worker_group:web-staging,service:be_web-staging-selenium-chrome-node-unus-staging}.as_count(),sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,worker_group:web-staging,service:be_web-staging-selenium-chrome-node-unus-staging}.as_count()",
                "Title": "HTTP 5xx/4xx",
                "SubTitle": "web-staging-selenium-chrome-node-unus-staging"
              },
              {
                "Query": "sum:sw.haproxy.backend.response.2xx{dc:op-us-east-1,worker_group:web-staging,service:be_web-staging-selenium-chrome-node-unus-staging}.as_count()",
                "Title": "HTTP 2xx",
                "SubTitle": "web-staging-selenium-chrome-node-unus-staging"
              },
              {
                "Query": "avg:sw.haproxy.backend.response.time{dc:op-us-east-1,worker_group:web-staging,service:be_web-staging-selenium-chrome-node-unus-staging}",
                "Title": "AVG Response Time",
                "SubTitle": "web-staging-selenium-chrome-node-unus-staging"
              },
              {
                "Query": "max:nomad.client.allocs.memory.rss{dc:op-us-east-1,worker_group:web-staging,task:web-staging-selenium-grid-unus-chrome-node-task,task_group:web-staging-selenium-grid-unus-chrome-node-group,job:web-staging-selenium-grid-unus-staging-op-us-east-1} by {alloc_id}.fill(0),max:docker.mem.limit{dc:op-us-east-1,worker_group:web-staging,task:web-staging-selenium-grid-unus-chrome-node-task,task_group:web-staging-selenium-grid-unus-chrome-node-group,job:web-staging-selenium-grid-unus-staging-op-us-east-1}",
                "Title": "Memory usage by allocation (RSS)",
                "SubTitle": ""
              }
            ],
            "web-staging-selenium-grid-unus-hub-task": [
              {
                "Query": "sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,worker_group:web-staging,service:be_web-staging-selenium-hub-unus-staging}.as_count(),sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,worker_group:web-staging,service:be_web-staging-selenium-hub-unus-staging}.as_count()",
                "Title": "HTTP 5xx/4xx",
                "SubTitle": "web-staging-selenium-hub-unus-staging"
              },
              {
                "Query": "sum:sw.haproxy.backend.response.2xx{dc:op-us-east-1,worker_group:web-staging,service:be_web-staging-selenium-hub-unus-staging}.as_count()",
                "Title": "HTTP 2xx",
                "SubTitle": "web-staging-selenium-hub-unus-staging"
              },
              {
                "Query": "avg:sw.haproxy.backend.response.time{dc:op-us-east-1,worker_group:web-staging,service:be_web-staging-selenium-hub-unus-staging}",
                "Title": "AVG Response Time",
                "SubTitle": "web-staging-selenium-hub-unus-staging"
              },
              {
                "Query": "max:nomad.client.allocs.memory.rss{dc:op-us-east-1,worker_group:web-staging,task:web-staging-selenium-grid-unus-hub-task,task_group:web-staging-selenium-grid-unus-hub-group,job:web-staging-selenium-grid-unus-staging-op-us-east-1} by {alloc_id}.fill(0),max:docker.mem.limit{dc:op-us-east-1,worker_group:web-staging,task:web-staging-selenium-grid-unus-hub-task,task_group:web-staging-selenium-grid-unus-hub-group,job:web-staging-selenium-grid-unus-staging-op-us-east-1}",
                "Title": "Memory usage by allocation (RSS)",
                "SubTitle": ""
              }
            ]
          }
        }
        
        
      );
    }, 2000);
  });
  app.get("/api/v1/deployments/:id", (req, res) => {
    setTimeout(function() {
      res.json(200,
        [
          {
            "Name": "nginx",
            "DeployTime": 1558187242,
            "Status": "successful"
          },
          {
            "Name": "nginx",
            "DeployTime": 1558187240,
            "Status": "successful"
          },
          {
            "Name": "nginx",
            "DeployTime": 1558187204,
            "Status": "successful"
          }
        ]
        
      );
    }, 2000);
  });
  app.get("/api/v1/deployments", (req, res) => {
    setTimeout(function() {
      res.json(200,
        [
          {
            "Name": "web-pro-sandbox-ligers-op-us-east-1",
            "DeployTime": 1559801014,
            "Status": "running"
          },
          {
            "Name": "web-userdata-sandbox-ligers-op-us-east-1",
            "DeployTime": 1559800819,
            "Status": "successful"
          },
          {
            "Name": "web-pro-sandbox-ligers-op-us-east-1",
            "DeployTime": 1559800798,
            "Status": "successful"
          },
          {
            "Name": "web-userdata-sandbox-24688-op-us-east-1",
            "DeployTime": 1559800285,
            "Status": "successful"
          },
          {
            "Name": "web-userdata-sandbox-24760-op-us-east-1",
            "DeployTime": 1559800280,
            "Status": "successful"
          },
          {
            "Name": "web-account-sandbox-freeze-op-us-east-1",
            "DeployTime": 1559799977,
            "Status": "successful"
          },
          {
            "Name": "web-reportsmanager-sandbox-freeze-op-us-east-1",
            "DeployTime": 1559799972,
            "Status": "successful"
          },
          {
            "Name": "web-pro-sandbox-freeze-op-us-east-1",
            "DeployTime": 1559799962,
            "Status": "successful"
          },
          {
            "Name": "web-userdata-sandbox-freeze-op-us-east-1",
            "DeployTime": 1559799727,
            "Status": "successful"
          },
          {
            "Name": "web-account-sandbox-freeze-op-us-east-1",
            "DeployTime": 1559799691,
            "Status": "successful"
          },
          {
            "Name": "web-reportsmanager-sandbox-freeze-op-us-east-1",
            "DeployTime": 1559799651,
            "Status": "successful"
          },
          {
            "Name": "web-pro-sandbox-freeze-op-us-east-1",
            "DeployTime": 1559799611,
            "Status": "successful"
          },
          {
            "Name": "web-internalapi-sandbox-staging-op-us-east-1",
            "DeployTime": 1559799412,
            "Status": "successful"
          },
          {
            "Name": "web-pro-sandbox-bisons-op-us-east-1",
            "DeployTime": 1559799157,
            "Status": "successful"
          },
          {
            "Name": "di-worker360-staging-op-us-east-1",
            "DeployTime": 1559798852,
            "Status": "successful"
          },
          {
            "Name": "pe-statusbay-production-op-us-east-1",
            "DeployTime": 1559798521,
            "Status": "successful"
          },
          {
            "Name": "web-pro-sandbox-sergey-op-us-east-1",
            "DeployTime": 1559798087,
            "Status": "successful"
          },
          {
            "Name": "web-userdata-sandbox-freeze-op-us-east-1",
            "DeployTime": 1559797986,
            "Status": "successful"
          },
          {
            "Name": "web-pro-sandbox-freeze-op-us-east-1",
            "DeployTime": 1559797986,
            "Status": "successful"
          },
          {
            "Name": "pe-statusbay-production-op-us-east-1",
            "DeployTime": 1559797974,
            "Status": "successful"
          }
        ]
        
        
      );
    }, 2000);
  });
  app.get("/api/v1/metric", (req, res) => {
    setTimeout(function() {
      res.json(200,
        {
          "Response": [
            {
              "metric": "haproxy.backend.response.5xx",
              "display_name": "haproxy.backend.response.5xx",
              "pointlist": [
                [
                  1557750360000,
                  78.23333287239075
                ],
                [
                  1557750390000,
                  83.73333350817363
                ],
                [
                  1557750420000,
                  89.73333326975505
                ],
                [
                  1557750450000,
                  87.56666652361551
                ],
                [
                  1557750480000,
                  99.7666662534078
                ],
                [
                  1557750510000,
                  92.46666701634724
                ],
                [
                  1557750540000,
                  76.7000000476837
                ],
                [
                  1557750570000,
                  86.33333349227905
                ],
                [
                  1557750600000,
                  82.23333406448364
                ],
                [
                  1557750630000,
                  90.23333279291789
                ],
                [
                  1557750660000,
                  81.63333336512248
                ],
                [
                  1557750690000,
                  86.59999974568684
                ],
                [
                  1557750720000,
                  90.83333317438762
                ],
                [
                  1557750750000,
                  96.59999990463258
                ],
                [
                  1557750780000,
                  89.86666758855183
                ],
                [
                  1557750810000,
                  89.76666688919067
                ],
                [
                  1557750840000,
                  104.99999984105426
                ],
                [
                  1557750870000,
                  92.56666723887125
                ],
                [
                  1557750900000,
                  86.36666647593181
                ],
                [
                  1557750930000,
                  80.4333333969116
                ],
                [
                  1557750960000,
                  85.43333355585735
                ],
                [
                  1557750990000,
                  83.099999666214
                ],
                [
                  1557751020000,
                  83.56666700045268
                ],
                [
                  1557751050000,
                  86.86666703224184
                ],
                [
                  1557751080000,
                  88.69999996821086
                ],
                [
                  1557751110000,
                  87.93333371480307
                ],
                [
                  1557751140000,
                  84.33333333333334
                ],
                [
                  1557751170000,
                  98.66666682561237
                ],
                [
                  1557751200000,
                  101.33333349227905
                ],
                [
                  1557751230000,
                  103.86666599909464
                ],
                [
                  1557751260000,
                  98.63333288828534
                ],
                [
                  1557751290000,
                  105.96666685740153
                ],
                [
                  1557751320000,
                  101.13333447774252
                ],
                [
                  1557751350000,
                  99.33333396911621
                ],
                [
                  1557751380000,
                  95.0666663646698
                ],
                [
                  1557751410000,
                  93.6666661898295
                ],
                [
                  1557751440000,
                  98.4333338737488
                ],
                [
                  1557751470000,
                  93.73333311080933
                ],
                [
                  1557751500000,
                  111.33333285649617
                ],
                [
                  1557751530000,
                  95.66666634877522
                ],
                [
                  1557751560000,
                  95.3666668732961
                ],
                [
                  1557751590000,
                  97.3333330154419
                ],
                [
                  1557751620000,
                  99.26666633288066
                ],
                [
                  1557751650000,
                  94.89999993642171
                ],
                [
                  1557751680000,
                  97.03333314259847
                ],
                [
                  1557751710000,
                  98.1666661898295
                ],
                [
                  1557751740000,
                  93.83333349227904
                ],
                [
                  1557751770000,
                  91.89999985694887
                ],
                [
                  1557751800000,
                  98.73333422342937
                ],
                [
                  1557751830000,
                  97.49999984105428
                ],
                [
                  1557751860000,
                  102.30000019073488
                ],
                [
                  1557751890000,
                  106.90000009536745
                ],
                [
                  1557751920000,
                  108.96666638056438
                ],
                [
                  1557751950000,
                  106.33333333333334
                ],
                [
                  1557751980000,
                  105.2333337465922
                ],
                [
                  1557752010000,
                  100.33333285649616
                ],
                [
                  1557752040000,
                  114.26666688919067
                ],
                [
                  1557752070000,
                  109.29999939600627
                ],
                [
                  1557752100000,
                  104.33333381017047
                ],
                [
                  1557752130000,
                  110.53333377838135
                ],
                [
                  1557752160000,
                  111.19999965031941
                ],
                [
                  1557752190000,
                  100.399999777476
                ],
                [
                  1557752220000,
                  104.43333371480307
                ],
                [
                  1557752250000,
                  101.76666688919067
                ],
                [
                  1557752280000,
                  101.6666661898295
                ],
                [
                  1557752310000,
                  105.43333371480306
                ],
                [
                  1557752340000,
                  98.13333304723102
                ],
                [
                  1557752370000,
                  103.83333301544188
                ],
                [
                  1557752400000,
                  108.10000101725262
                ],
                [
                  1557752430000,
                  115.80000027020772
                ],
                [
                  1557752460000,
                  120.53333298365276
                ],
                [
                  1557752490000,
                  123.46666669845582
                ],
                [
                  1557752520000,
                  101.93333403269449
                ],
                [
                  1557752550000,
                  105.03333314259848
                ],
                [
                  1557752580000,
                  101.20000044504803
                ],
                [
                  1557752610000,
                  97.70000012715657
                ],
                [
                  1557752640000,
                  108.06666707992554
                ],
                [
                  1557752670000,
                  114.83333349227907
                ],
                [
                  1557752700000,
                  106.36666742960612
                ],
                [
                  1557752730000,
                  117.73333358764648
                ],
                [
                  1557752760000,
                  116.76666657129924
                ],
                [
                  1557752790000,
                  112.56666628519693
                ],
                [
                  1557752820000,
                  97.16666682561238
                ],
                [
                  1557752850000,
                  101.99999999999999
                ],
                [
                  1557752880000,
                  99.63333336512248
                ],
                [
                  1557752910000,
                  101.13333368301393
                ],
                [
                  1557752940000,
                  110.46666669845582
                ],
                [
                  1557752970000,
                  103.70000012715658
                ],
                [
                  1557753000000,
                  106.06666723887128
                ],
                [
                  1557753030000,
                  111.53333314259848
                ],
                [
                  1557753060000,
                  102.13333304723103
                ],
                [
                  1557753090000,
                  110.63333336512248
                ],
                [
                  1557753120000,
                  104.100000222524
                ],
                [
                  1557753150000,
                  108.23333358764648
                ],
                [
                  1557753180000,
                  101.10000038146974
                ],
                [
                  1557753210000,
                  96.53333346048991
                ],
                [
                  1557753240000,
                  102.50000063578287
                ],
                [
                  1557753270000,
                  94.79999979337056
                ],
                [
                  1557753300000,
                  107.7999997138977
                ],
                [
                  1557753330000,
                  112.2999997138977
                ],
                [
                  1557753360000,
                  110.46666653951009
                ],
                [
                  1557753390000,
                  105.46666653951009
                ],
                [
                  1557753420000,
                  103.73333319028217
                ],
                [
                  1557753450000,
                  106.70000076293945
                ],
                [
                  1557753480000,
                  107.46666669845581
                ],
                [
                  1557753510000,
                  107.60000022252402
                ],
                [
                  1557753540000,
                  107.03333282470702
                ],
                [
                  1557753570000,
                  105.53333393732707
                ],
                [
                  1557753600000,
                  106.13333304723106
                ],
                [
                  1557753630000,
                  115.56666707992554
                ],
                [
                  1557753660000,
                  109.99999952316284
                ],
                [
                  1557753690000,
                  107.13333336512247
                ],
                [
                  1557753720000,
                  112.60000006357829
                ],
                [
                  1557753750000,
                  112.70000044504802
                ],
                [
                  1557753780000,
                  113.63333336512248
                ],
                [
                  1557753810000,
                  123.93333292007445
                ],
                [
                  1557753840000,
                  132.46666701634726
                ],
                [
                  1557753870000,
                  114.73333350817363
                ],
                [
                  1557753900000,
                  111.89999945958455
                ],
                [
                  1557753930000,
                  112.09999974568684
                ],
                [
                  1557753960000,
                  111.13333288828531
                ],
                [
                  1557753990000,
                  111.13333288828531
                ],
                [
                  1557754020000,
                  111.13333288828531  
                ],
                [
                  1557754080000,
                  111.13333288828531
                ],
                [
                  1557754140000,
                  111.13333288828531
                ]
              ],
              "start": 1557750360000,
              "end": 1557753989000,
              "interval": 30,
              "aggr": "sum",
              "length": 121,
              "scope": "dc:op-us-east-1,service:be_web-api-nginx-production,worker_group:web",
              "expression": "sum:haproxy.backend.response.5xx{dc:op-us-east-1,service:be_web-api-nginx-production,worker_group:web}",
              "unit": [
                {
                  "family": "network",
                  "scale_factor": 1,
                  "name": "response",
                  "short_name": "rsp",
                  "plural": "responses",
                  "id": 28
                },
                null
              ]
            },
            {
              "metric": "haproxy.backend.response.4xx",
              "display_name": "haproxy.backend.response.4xx",
              "pointlist": [
                [
                  1557750360000,
                  56.0666667620341
                ],
                [
                  1557750390000,
                  45.30000034968058
                ],
                [
                  1557750420000,
                  34.16666650772095
                ],
                [
                  1557750450000,
                  32.6666667064031
                ],
                [
                  1557750480000,
                  37.566666762034096
                ],
                [
                  1557750510000,
                  39.600000222524
                ],
                [
                  1557750540000,
                  32.23333323001861
                ],
                [
                  1557750570000,
                  42.93333359559377
                ],
                [
                  1557750600000,
                  34.26666657129924
                ],
                [
                  1557750630000,
                  39.49999992052714
                ],
                [
                  1557750660000,
                  44.60000006357829
                ],
                [
                  1557750690000,
                  38.300000190734856
                ],
                [
                  1557750720000,
                  35.16666668653488
                ],
                [
                  1557750750000,
                  38.500000039736435
                ],
                [
                  1557750780000,
                  39.666666547457375
                ],
                [
                  1557750810000,
                  41.76666641235351
                ],
                [
                  1557750840000,
                  37.19999965031942
                ],
                [
                  1557750870000,
                  35.566666563351944
                ],
                [
                  1557750900000,
                  27.93333331743876
                ],
                [
                  1557750930000,
                  36.93333315849304
                ],
                [
                  1557750960000,
                  20.3999999264876
                ],
                [
                  1557750990000,
                  16.36666669448217
                ],
                [
                  1557751020000,
                  19.899999876817066
                ],
                [
                  1557751050000,
                  17.400000025828678
                ],
                [
                  1557751080000,
                  21.499999999999996
                ],
                [
                  1557751110000,
                  25.333333452542618
                ],
                [
                  1557751140000,
                  25.70000004768372
                ],
                [
                  1557751170000,
                  34.03333342075348
                ],
                [
                  1557751200000,
                  23.53333347042401
                ],
                [
                  1557751230000,
                  28.13333332538605
                ],
                [
                  1557751260000,
                  43.299999833107
                ],
                [
                  1557751290000,
                  24.666666805744175
                ],
                [
                  1557751320000,
                  29.633333603541054
                ],
                [
                  1557751350000,
                  26.16666664679845
                ],
                [
                  1557751380000,
                  23.666666413346928
                ],
                [
                  1557751410000,
                  30.43333333730698
                ],
                [
                  1557751440000,
                  24.53333342075348
                ],
                [
                  1557751470000,
                  28.76666686932246
                ],
                [
                  1557751500000,
                  37.033333321412414
                ],
                [
                  1557751530000,
                  44.49999996026358
                ],
                [
                  1557751560000,
                  26.1333331267039
                ],
                [
                  1557751590000,
                  17.000000029802322
                ],
                [
                  1557751620000,
                  26.066666742165882
                ],
                [
                  1557751650000,
                  34.53333334128062
                ],
                [
                  1557751680000,
                  37.79999983310699
                ],
                [
                  1557751710000,
                  42.23333319028219
                ],
                [
                  1557751740000,
                  36.133333057165146
                ],
                [
                  1557751770000,
                  43.43333331743876
                ],
                [
                  1557751800000,
                  38.03333334128062
                ],
                [
                  1557751830000,
                  38.89999987681706
                ],
                [
                  1557751860000,
                  40.86666655540466
                ],
                [
                  1557751890000,
                  50.633333245913185
                ],
                [
                  1557751920000,
                  55.43333371480307
                ],
                [
                  1557751950000,
                  44.23333330949147
                ],
                [
                  1557751980000,
                  44.4333332379659
                ],
                [
                  1557752010000,
                  48.60000002384187
                ],
                [
                  1557752040000,
                  38.666666785875954
                ],
                [
                  1557752070000,
                  49.86666651566824
                ],
                [
                  1557752100000,
                  45.9333332379659
                ],
                [
                  1557752130000,
                  50.09999998410543
                ],
                [
                  1557752160000,
                  38.966666499773666
                ],
                [
                  1557752190000,
                  36.133333365122475
                ],
                [
                  1557752220000,
                  37.69999996821086
                ],
                [
                  1557752250000,
                  34.36666639645894
                ],
                [
                  1557752280000,
                  35.96666653951009
                ],
                [
                  1557752310000,
                  41.16666670640309
                ],
                [
                  1557752340000,
                  33.23333330949148
                ],
                [
                  1557752370000,
                  34.30000023047129
                ],
                [
                  1557752400000,
                  46.56666684150695
                ],
                [
                  1557752430000,
                  44.90000001589458
                ],
                [
                  1557752460000,
                  52.49999992052714
                ],
                [
                  1557752490000,
                  44.233333071072906
                ],
                [
                  1557752520000,
                  47.099999825159706
                ],
                [
                  1557752550000,
                  47.46666673819224
                ],
                [
                  1557752580000,
                  49.99999996026358
                ],
                [
                  1557752610000,
                  49.566666682561234
                ],
                [
                  1557752640000,
                  50.166666507720954
                ],
                [
                  1557752670000,
                  49.06666684150696
                ],
                [
                  1557752700000,
                  56.56666644414266
                ],
                [
                  1557752730000,
                  55.033332983652755
                ],
                [
                  1557752760000,
                  45.06666660308838
                ],
                [
                  1557752790000,
                  74.53333322207132
                ],
                [
                  1557752820000,
                  48.599999864896134
                ],
                [
                  1557752850000,
                  44.966666579246514
                ],
                [
                  1557752880000,
                  37.63333304723104
                ],
                [
                  1557752910000,
                  42.83333293596903
                ],
                [
                  1557752940000,
                  39.10000010331471
                ],
                [
                  1557752970000,
                  43.23333356777827
                ],
                [
                  1557753000000,
                  46.566666444142655
                ],
                [
                  1557753030000,
                  44.199999968210854
                ],
                [
                  1557753060000,
                  50.93333327770233
                ],
                [
                  1557753090000,
                  48.900000174840294
                ],
                [
                  1557753120000,
                  50.13333344459534
                ],
                [
                  1557753150000,
                  49.03333322207133
                ],
                [
                  1557753180000,
                  44.23333350817362
                ],
                [
                  1557753210000,
                  38.26666649182637
                ],
                [
                  1557753240000,
                  42.63333352406819
                ],
                [
                  1557753270000,
                  58.9333332379659
                ],
                [
                  1557753300000,
                  48.66666698455811
                ],
                [
                  1557753330000,
                  46.099999626477555
                ],
                [
                  1557753360000,
                  41.30000007152558
                ],
                [
                  1557753390000,
                  40.99999992052714
                ],
                [
                  1557753420000,
                  23.600000063578285
                ],
                [
                  1557753450000,
                  27.033333082993828
                ],
                [
                  1557753480000,
                  40.86666639645894
                ],
                [
                  1557753510000,
                  47.566666762034096
                ],
                [
                  1557753540000,
                  36.40000025431315
                ],
                [
                  1557753570000,
                  35.79999979337057
                ],
                [
                  1557753600000,
                  42.90000009536743
                ],
                [
                  1557753630000,
                  49.36666667461395
                ],
                [
                  1557753660000,
                  38.7999999721845
                ],
                [
                  1557753690000,
                  48.399999936421715
                ],
                [
                  1557753720000,
                  50.966666539510086
                ],
                [
                  1557753750000,
                  46.96666661898295
                ],
                [
                  1557753780000,
                  41.76666669050852
                ],
                [
                  1557753810000,
                  52.33333333333333
                ],
                [
                  1557753840000,
                  38.73333330949148
                ],
                [
                  1557753870000,
                  50.39999977747599
                ],
                [
                  1557753900000,
                  73.49999952316284
                ],
                [
                  1557753930000,
                  45.133333524068206
                ],
                [
                  1557753960000,
                  43.36666683355968
                ]
              ],
              "start": 1557750360000,
              "end": 1557753989000,
              "interval": 30,
              "aggr": "sum",
              "length": 121,
              "scope": "dc:op-us-east-1,service:be_web-api-nginx-production,worker_group:web",
              "expression": "sum:haproxy.backend.response.4xx{dc:op-us-east-1,service:be_web-api-nginx-production,worker_group:web}",
              "unit": [
                {
                  "family": "network",
                  "scale_factor": 1,
                  "name": "response",
                  "short_name": "rsp",
                  "plural": "responses",
                  "id": 28
                },
                null
              ]
            }
          ],
          "Error": null
        }
        
      );
    }, 2000);
  });
};


