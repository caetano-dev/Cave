import React, {useState} from "react";
import "./SideBar.css";
import CreateFileButton from "../CreateFileButton/CreateFileButton";
import ToggleButton from "../ToggleButton/ToggleButton";
import PopUp from "../PopUp/PopUp";
import PropTypes from 'prop-types'
import deleteFile from "../../utils/deleteFile"


function SideBar(props) {
  const [popUpPosition, setPopUpPosition] = useState({x:0,y:0})
  const [selectedFile, setSelectedFile] = useState(null);

  const setVariables = (filename, id, tags, content, index) => {
    props.setFilename(filename);
    props.setId(id);
    props.setTags(tags);
    props.setContent(content);
    props.setFileIndex(index);
  };

  return (
    <aside className={"sidebar " + props.toggle}>
      {props.data.length > 0 ? (
        <>
          <div>
            <CreateFileButton setData={props.setData} />
            <ToggleButton toggle={props.toggle} toggleState={props.toggleState} />
          </div>
          <ul>
            {props.data.map((files, index) => (
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
                onContextMenu={(event) => {
                  event.preventDefault();
                  setSelectedFile(index);
                  setPopUpPosition({ x: event.clientX, y: event.clientY });
                }}
              >
                {selectedFile !== null && (
                  <PopUp
                  text={"Delete file"}
                    position={popUpPosition}
                    onClickFunction={() => {
                      const id = props.data[selectedFile].FileInformation.id;
                      deleteFile(id);
                      props.setData([
                        ...props.data.slice(0, selectedFile),
                        ...props.data.slice(selectedFile + 1),
                      ]);
                      setSelectedFile(null);
                      localStorage.setItem("data", JSON.stringify(props.data));
                    }}
                  />
                )}

                <React.Fragment key={files.FileInformation.id}>
                  <p>{files.FileInformation.filename}</p>
                </React.Fragment>
              </li>
            ))}
          </ul>
        </>
      ) : (
        <div className="noFiles">
          <CreateFileButton setData={props.setData} />
          <p style="color:white; margin: 0 0 0 1.5rem">You have no files</p>
        </div>
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
