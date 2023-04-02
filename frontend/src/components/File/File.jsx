import React, { useEffect } from "react";
import "./File.css";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";
import modules from "../../modules/quill";
import TagsList from "../TagList/TagsList";
import PropTypes from 'prop-types'

function File({
  id,
  tags,
  toggle,
  filename,
  setFilename,
  fileIndex,
  content,
  setContent,
  data,
  setData,
}) {
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
    setData(data);
    localStorage.setItem("data", JSON.stringify(data));
  };

  function handleEditFilename(event) {
    const newFilename = event.target.value;
    setFilename(newFilename);
    data[fileIndex].FileInformation.filename = newFilename;
    editFileInServer(id, "filename", newFilename);
  }

  useEffect(() => {
    function handleSaveShortcutKeyDown(event) {
      if (event.ctrlKey && event.key === "s") {
        event.preventDefault();
        setContent(content);
        data[fileIndex].Content = content;
        editFileInServer(id, "content", content);
      }
    }
    document.addEventListener("keydown", handleSaveShortcutKeyDown);
    return () => {
      document.removeEventListener("keydown", handleSaveShortcutKeyDown);
    };
  }, [content]);

  return (
    <main className={toggle}>
      <input onChange={handleEditFilename} className="title" value={filename} />
      <TagsList tags={tags} />
      <ReactQuill
        className="quill"
        modules={modules}
        theme="snow"
        value={content}
        onChange={setContent}
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
