import React, { useEffect } from "react";
import "./File.css";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";
import modules from "../../modules/quill";
import TagsList from "../TagList/TagsList";

function File({
  toggle,
  filename,
  setFilename,
  id,
  content,
  setContent,
  setData,
  data,
  tags,
  fileIndex,
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

export default File;
