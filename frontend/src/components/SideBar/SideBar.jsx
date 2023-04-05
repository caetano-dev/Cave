import React from "react";
import "./SideBar.css";
import CreateFileButton from "../CreateFileButton/CreateFileButton";
import PropTypes from 'prop-types'


function SideBar({
  toggle,
  toggleState,
  data,
  setFilename,
  setData,
  setId,
  setTags,
  setContent,
  setFileIndex,
}) {
  const setVariables = (filename, id, tags, content, index) => {
    setFilename(filename);
    setId(id);
    setTags(tags);
    setContent(content);
    setFileIndex(index);
  };

  return (
    <aside className={"sidebar " + toggle}>
      {data.length > 0 ? (
        <>
          <div>
            <CreateFileButton setData={setData} />
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
                    index
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

SideBar.propTypes = {
  toggle: PropTypes.string.isRequired,
  toggleState: PropTypes.func.isRequired,
  data: PropTypes.arrayOf(
    PropTypes.shape({
      FileInformation: PropTypes.shape({
        filename: PropTypes.string.isRequired,
        id: PropTypes.number.isRequired,
        tags: PropTypes.arrayOf(PropTypes.string).isRequired,
      }).isRequired,
      Content: PropTypes.string.isRequired,
    })
  ).isRequired,
  setFilename: PropTypes.func.isRequired,
  setId: PropTypes.func.isRequired,
  setTags: PropTypes.func.isRequired,
  setContent: PropTypes.func.isRequired,
  setFileIndex: PropTypes.func.isRequired,
};
export default SideBar;
