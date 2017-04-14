import { vsprintf } from 'sprintf-js'

export const API_BASE =  '/api'

export const API_JOBS_LIST = '/jobs'
export const API_JOBS_SHOW = '/jobs/%s'
export const API_JOBS_CREATE = '/jobs'
export const API_JOBS_UPDATE = '/jobs/%s'
export const API_JOBS_DELETE = '/jobs/%s'

export function buildRoute(route, ...params) {
  return API_BASE + vsprintf(route, params);
}
