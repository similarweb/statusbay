import { combineReducers } from 'redux'
import { connectRouter } from 'connected-react-router'
import { deployment } from '../reducers/deployment.reducer';

const rootReducer = (history) => combineReducers({
  deployment,
  router: connectRouter(history)
})

export default rootReducer
