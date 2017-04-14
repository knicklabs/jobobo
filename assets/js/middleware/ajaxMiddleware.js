import axios from 'axios'

export default function ajaxMiddleware() {
  return dispatch => action => {
    const {
      fetch,
      type,
      method = 'GET',
      data,
      params,
      ...rest } = action

    if (!fetch) return dispatch(action)

    const LOADING = type + '_LOADING'
    const SUCCESS = type + '_SUCCESS'
    const ERROR = type + '_ERROR'

    dispatch({...rest, type: LOADING})

    return axios({
        url: fetch,
        method: method,
        data: data,
        params: params
      })
      .then((response) => {
        dispatch({...rest, response: response.data, type: SUCCESS})
      })
      .catch((error) => {
        dispatch({...rest, error, type: ERROR})
      })
  }
}