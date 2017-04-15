import React from 'react'

import Header from './Header'
import ReturnBar from './ReturnBar'

const App = ({children}) => {
  return (
    <div>
      <ReturnBar />
      <Header />
        <div className="main--wrapper">
          <main className="main">
            {children}
          </main>
        </div>
    </div>
  )
}

export default App