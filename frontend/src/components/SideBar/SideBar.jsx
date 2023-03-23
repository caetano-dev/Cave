import React from "react";
import "./SideBar.css";

function SideBar({ toggle, toggleState, data, setFilename, setId, setTags }) {
  const renderFolder = (folder) => (
    <li className="folder">
      {folder.name} <ul>{folder.children.map(renderNode)}</ul>
    </li>
  );

  const setVariables = (filename, id, tags) => {
    setFilename(filename);
    setId(id);
    setTags(tags);
  };

  const renderFile = (file) => (
    <li
      className="file"
      onClick={() => setVariables(file.filename, file.id, file.tags)}
    >
      {file.filename}
    </li>
  );
  const renderNode = (content) => (
    <>
      {content.type === "folder" ? renderFolder(content) : renderFile(content)}
    </>
  );

  return (
    <aside className={"sidebar " + toggle}>
      {data.length > 0 ? (
        <>
          <div>
            <button
              className={"sideBarButton " + toggle}
              title="Toggle sidebar"
              onClick={toggleState}
            >
              ☰
            </button>
          </div>
          <ul>{data.map(renderNode)}</ul>
        </>
      ) : (
        <p style="color:white; margin: 0 0 0 1.5rem">You have no files.</p>
      )}
    </aside>
  );
}

export default SideBar;
