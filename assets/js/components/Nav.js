import React from 'react'

const Nav = () => {
  return (
    <nav className="site-nav">
      <ul className="site-nav--menu">
        <li className="site-nav--menu--item site-nav--menu--item__first">
          <a href="/" className="site-nav--home">
            <img src="/assets/images/sn-jobs.png" alt="Software Niagara Jobs" className="site-nav--logo"/>
          </a>
        </li>
        <li className="site-nav--menu--item site-nav--menu--item__last">
          <a href="" className="site-nav--button">Post Job</a>
        </li>
      </ul>
    </nav>
  )
}

export default Nav