import React, { useState } from 'react';

import SideBar from './components/SideBar/SideBar'
import File from './components/File/File'

import './app.css'

export function App() {
  const [toggle, setToggle] = useState("open") //open/close
  const toggleState = () => toggle == "open" ? setToggle("close") : setToggle("open");

  return (
    <>
      <File toggle={toggle} />
      <SideBar toggle={toggle} toggleState={toggleState} />
    </>
  );
}
