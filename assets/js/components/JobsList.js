import React from 'react'

import JobsListItem from './JobsListItem'
import Promo from './Promo'

const JobsList = ({jobs}) => {
  return (
    <div>
      <Promo />
      <section className="current-listings">
        <h2 className="current-listings--title">Current Listings</h2>
        <ul className="job-list">
          {
            jobs.map((job, index) => {
              return (
                <li key={`job-${index}`} className="job-list--item">
                  <JobsListItem {...job} />
                </li>
              )
            })
          }
        </ul>
      </section>
    </div>
  )
}

export default JobsList