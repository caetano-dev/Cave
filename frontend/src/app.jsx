import React, { useState } from 'react';

import SideBar from './components/SideBar/SideBar'
import File from './components/File/File'

import './app.css'

export function App() {
  const [toggle, setToggle] = useState("open") //open/close
  const toggleState = () => toggle == "open" ? setToggle("close") : setToggle("open");

  const [filename, setFilename] = useState("Title")
  const [content, setContent] = useState("content")

  return (
    <>

      {/*clica -> baixa automaticamente (pode ser cache)-> edita -> salva -> upload*/}
      <File toggle={toggle} filename={filename} content={content} />
      <SideBar toggle={toggle} toggleState={toggleState} />
    </>
  );
}
