import React from "react";
import "./SideBar.css";

function SideBar({ toggle, toggleState, data, setFilename, setId, setTags, setContent, setFileIndex }) {
  const setVariables = (filename, id, tags, content, index) => {
    setFilename(filename);
    setId(id);
    setTags(tags);
    setContent(content);
    setFileIndex(index)
  };

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
          <ul>
            {data.map((files, index) => (
              <li
                className="file"
                onClick={() =>
                  setVariables(
                    files.FileInformation.filename,
                    files.FileInformation.id,
                    files.FileInformation.tags,
                    files.Content,
                    index,
                  )
                }
              >
                <React.Fragment key={files.FileInformation.id}>
                  <p>{files.FileInformation.filename}</p>
                </React.Fragment>
              </li>
            ))}
          </ul>
        </>
      ) : (
        <p style="color:white; margin: 0 0 0 1.5rem">You have no files</p>
      )}
    </aside>
  );
}
export default SideBar;
