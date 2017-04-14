import * as types from '../constants/types'

const initialState = {
  data: [],
  loading: false
}

export default function jobs(state = initialState, action) {
  switch (action.type) {
    case types.FETCH_JOBS_LOADING:
      return {
        ...state,
        loading: true
      }
    case types.FETCH_JOBS_SUCCESS:
      return {
        ...state,
        data: action.response
      }
    default:
      return state
  }
}