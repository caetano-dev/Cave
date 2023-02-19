import React, { useState, useEffect } from 'react';
import './SideBar.css';

function SideBar({ toggle, toggleState }) {
  const [data, setData] = useState([
  ]);
  useEffect(() => {
    fetch("http://localhost:3000/files")
      .then((response) => response.json())
      .then((json) => setData(json))
  }, []);

  console.log(data)

  const renderFolder = (folder) => (
    <li className="folder">
      {folder.name} <ul>{folder.children.map(renderNode)}</ul>
    </li>
  );
  const renderFile = (file) => <li className="file">{file.filename}</li>;
  const renderNode = (content) => <>{content.type == 'folder' ? renderFolder(content) : renderFile(content)}</>

  return <>
    <aside className={"sidebar " + toggle}>
      <button className="sideBarButton" onClick={toggleState}>☰</button>
      <ul>{data.map(renderNode)}</ul>
    </aside>
  </>
}

export default SideBar;

