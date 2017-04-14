import React from 'react'
import { render } from 'react-dom'

import { applyMiddleware, compose, createStore, combineReducers } from 'redux'
import { Provider } from 'react-redux'
import { Router, Route, IndexRoute, browserHistory } from 'react-router'
import { syncHistoryWithStore, routerReducer } from 'react-router-redux'

import * as reducers from './reducers'
import { App } from './components'
import { JobsListContainer } from './containers'

import ajaxMiddleware from './middleware/ajaxMiddleware'

const reducer = combineReducers({
  ...reducers,
  routing: routerReducer
})

const finalCreateStore = compose(
  applyMiddleware(ajaxMiddleware)
)(createStore)

const store = finalCreateStore(reducer)

const history = syncHistoryWithStore(browserHistory, store)

render(
  <Provider store={store}>
    <Router history={history}>
      <Route path="/" component={App}>
        <IndexRoute component={JobsListContainer} />
      </Route>
    </Router>
  </Provider>
  ,
  document.getElementById('root')
)