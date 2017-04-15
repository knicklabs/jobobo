import React from 'react'

const Promo = () => {
  return (
    <article className="promo">
      <header className="promo--header">
        <i className="fa fa-handshake-o promo--icon" aria-hidden="true"></i>
        <h1 className="promo--title">
          <strong>Software Niagara Jobs</strong> is a great place to find and list technology jobs in the Greater
          Niagara Area
        </h1>
      </header>
      <div className="promo--main">
        <p>Work local. Hire local.</p>
      </div>
      <footer className="promo--footer">
        <a href="#" className="promo--button">Post a Job</a>
        <p>Free for 30-days</p>
      </footer>
    </article>
  )
}

export default Promo