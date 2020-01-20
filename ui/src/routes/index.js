import React from 'react'
import { Route, Switch } from 'react-router'
import Deployment from '../components/Deployment/Index'
import Jobs from '../components/Jobs/Index'
import Search from '../components/Search/Index'
import Deployments from '../components/Deployments/Index'
import Header from '../components/Header'
import NotFound from '../components/NotFound'
import { Container } from 'react-bootstrap';
import Statistics from '../utils/Statistics'
import { history } from "../configureStore";


export default class Routes extends React.Component {

  constructor(props) {
    super(props);

    history.listen(() => {
      Statistics.pageView();
    });
  }

    /**
   * When component mount
   */
  componentDidMount() {
    Statistics.pageView();
  }

  render(){
    return(
      <div>
        <Header />
        <Container>
          <div id="main">
            <Switch>
              <Route exact path="/" component={Jobs} />
              <Route exact path="/search" component={Search} />
              <Route exact path="/deployments/:job" component={Deployments} />
              <Route exact path="/deployments/:job/:time" component={Deployment} />
              <Route path="*" component={NotFound}/>
            </Switch>
          </div>
          
        </Container>
      </div>
    );
  }
}



