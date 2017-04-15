import React from 'react'
import moment from 'moment'

const JobsListItem = ({company_name, created_at, title}) => {
  return (
    <article className="job-teaser">
      <a href="#" className="job-teaser--link">
        <div className="job-teaser--part">
          <h3 className="job-teaser--title">{company_name}</h3>
          <p className="job-teaser--subtitle">{title}</p>
        </div>
        <div className="job-teaser--part">
          <p className="job-teaser--date">{moment(created_at).fromNow()}</p>
        </div>
      </a>
    </article>
  )
}

export default JobsListItem