import React, { useEffect } from "react";
import "./File.css";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";
import modules from "../../modules/quill";
import TagsList from "../TagList/TagsList";
import PropTypes from 'prop-types'

function File(props) {
  const editFileInServer = async (id, field, value) => {
    const host = "http://localhost:3000";
    const url =
      field === "content" ? `${host}/fileEditContent` : `${host}/fileEditName`;

    const fetchParamns = {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ id, value }),
    };

    fetch(url, fetchParamns).catch(console.error);
  };

  function handleEditFilename(event) {
    const newFilename = event.target.value;
    props.setFilename(newFilename);
    props.data[props.fileIndex].FileInformation.filename = newFilename;
    localStorage.setItem("data", JSON.stringify(props.data));
    editFileInServer(props.id, "filename", newFilename);
  }

  useEffect(() => {
    function handleSaveShortcutKeyDown(event) {
      if (event.ctrlKey && event.key === "s") {
        event.preventDefault();
        props.setContent(props.content);
        props.data[props.fileIndex].Content = props.content;
        localStorage.setItem("data", JSON.stringify(props.data));
        editFileInServer(props.id, "content", props.content);
      }
    }
    document.addEventListener("keydown", handleSaveShortcutKeyDown);
    return () => {
      document.removeEventListener("keydown", handleSaveShortcutKeyDown);
    };
  }, [props.content]);

  return (
    <main className={props.toggle}>
      <input onChange={handleEditFilename} className="title" value={props.filename} />
      <TagsList tags={props.tags} />
      <ReactQuill
        className="quill"
        modules={modules}
        theme="snow"
        value={props.content}
        onChange={props.setContent}
      />
    </main>
  );
}

File.propTypes = {
  toggle: PropTypes.string.isRequired,
  filename: PropTypes.string.isRequired,
  setFilename: PropTypes.func.isRequired,
  setContent: PropTypes.func.isRequired,
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
  setData: PropTypes.func.isRequired,
  id: PropTypes.number.isRequired,
  content: PropTypes.string.isRequired,
  tags: PropTypes.arrayOf(PropTypes.string).isRequired,
  fileIndex: PropTypes.number.isRequired,
};
export default File;
