import React, { useEffect, useState } from "react";
import "./File.css";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";

function File({ toggle, filename, id, content, setContent }) {
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

  const modules = {
    toolbar: [
      [{ header: [1, 2, 3, 4, 5, 6, false] }, { font: [] }],
      [{ size: [] }],
      [{ color: [] }, { background: [] }],
      ["bold", "italic", "underline", "strike", "blockquote", "code-block"],
      [
        { list: "ordered" },
        { list: "bullet" },
        { indent: "-1" },
        { indent: "+1" },
        { direction: "rtl" },
      ],
      ["link", "image", "video"],
      ["clean"],
    ],
  };

  return (
    <main className={toggle}>
      <h1 className="title">{filename}</h1>
      <ReactQuill
        modules={modules}
        theme="snow"
        value={content}
        onChange={setContent}
      />
    </main>
  );
}

export default File;
