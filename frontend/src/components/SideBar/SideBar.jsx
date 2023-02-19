import React, { useState, useEffect} from 'react';
import './SideBar.css';

function SideBar({ toggle, toggleState }) {
  const [data, setData] = useState([
    {
      name: "Foobar",
      type: "folder",
      children: [
        {
          name: "foo.txt",
          type: "file",
        },
        {
          name: "bar.txt",
          type: "file",
        },
      ],
    },
    {
      name: "Lorem Impsum",
      type: "folder",
      children: [
        {
          name: "lorem.txt",
          type: "file",
        },
        {
          name: "inpsum.txt",
          type: "file",
        },
        {
          name: "Latin Class",
          type: "folder",
          children: [
            {
              name: "latin.txt",
              type: "file",
            },
          ],
        },
      ],
    },
    {
      name: "file.txt",
      type: "file",
    },
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
  const renderFile = (file) => <li className="file">{file.name}</li>;

  const renderNode = (content) => <>{content.type == 'folder' ? renderFolder(content) : renderFile(content)}</>

  return <>
    <aside className={"sidebar " + toggle}>
      <button className="sideBarButton" onClick={toggleState}>☰</button>
      <ul>{data.map(renderNode)}</ul>
    </aside>
  </>
}

export default SideBar;

