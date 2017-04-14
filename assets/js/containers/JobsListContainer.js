import React, { Component } from 'react'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import * as actions from '../actions/jobsActions'

import JobsList from '../components/JobsList'

class JobsListContainer extends Component {
  constructor(props) {
    super(props)
  }

  componentDidMount() {
    if (!this.props.jobs.data.length) {
      this.props.actions.fetchJobs()
    }
  }

  render() {
    return <JobsList jobs={this.props.jobs.data}/>
  }
}

export default connect(mapStateToProps, mapDisptachToProps)(JobsListContainer)

function mapStateToProps(state) {
  return {jobs: state.jobs}
}

function mapDisptachToProps(dispatch) {
  return { actions: bindActionCreators(actions, dispatch) }
}