import * as types from '../constants/types'
import * as routes from '../constants/routes'

export function fetchJobs() {
  return {
    type: types.FETCH_JOBS,
    fetch: routes.buildRoute(routes.API_JOBS_CREATE)
  }
}