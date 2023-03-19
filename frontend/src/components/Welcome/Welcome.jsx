import React from "react";
import "./Welcome.css"
import noteIcon from "../../assets/note-icon.png"

function Welcome({toggle}) {
  return (
    <main className={toggle + " welcome"}>
      <h1>No file selected.</h1>
      <img src={noteIcon}/>
      <h2>Open a file or create a new one.</h2>
    </main>
  );
}

export default Welcome;
