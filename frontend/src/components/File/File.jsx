import React, { useEffect } from "react";
import "./File.css";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";
import modules from "../../modules/quill";
import TagsList from "../TagList/TagsList";

function File({ toggle, filename, id, content, setContent, tags }) {
  const editContent = async (id) => {
    try {
      const response = await fetch("http://localhost:3000/fileEditContent", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id: id,
          content: content,
        }),
      });
      if (!response.ok) {
        throw new Error("Failed to edit file content");
      }
    } catch (error) {
      console.error(error);
      //handle error here. We can display a message to the user.
    }
  };

  useEffect(() => {
    function handleSaveShortcutKeyDown(event) {
      if (event.ctrlKey && event.key === "s") {
        event.preventDefault();
        editContent(id);
        setContent(content);
        console.log("content updated: " + content);
      }
    }
    document.addEventListener("keydown", handleSaveShortcutKeyDown);
    return () => {
      document.removeEventListener("keydown", handleSaveShortcutKeyDown);
    };
  }, [content]);

  return (
    <main className={toggle}>
      <input className="title" placeholder={filename} value={filename}/> 
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
