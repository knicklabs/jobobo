import React from 'react'

import JobsListItem from './JobsListItem'

const JobsList = ({jobs}) => {
  return (
    <ul>
      {
        jobs.map((job, index) => {
          return (
            <li key={`job-${index}`}>
              <JobsListItem {...job} />
            </li>
          )
        })
      }
    </ul>
  )
}

export default JobsList