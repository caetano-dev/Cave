import React, { useState, useEffect } from "react";
import SideBar from "./components/SideBar/SideBar";
import File from "./components/File/File";
import Welcome from "./components/Welcome/Welcome";
import fetchFiles from "./utils/utils"

import "./app.css";

export function App() {
  const [toggle, setToggle] = useState("open");
  const [filename, setFilename] = useState("");
  const [data, setData] = useState([]);
  const [id, setId] = useState(null);
  const [tags, setTags] = useState("");
  const [content, setContent] = useState("");
  const [fileIndex, setFileIndex] = useState("");

  const toggleState = () => setToggle((toggle) => (toggle === "open" ? "close" : "open"));

  useEffect(() => {
    fetchFiles(setData)
  }, []);
console.log(data)

  return (
    <>
      <SideBar
        toggle={toggle}
        toggleState={toggleState}
        data={data}
        setData={setData}
        setFilename={setFilename}
        setId={setId}
        setTags={setTags}
        setContent={setContent}
        setFileIndex={setFileIndex}
      />
      {id ? (
        <File
          toggle={toggle}
          setFilename={setFilename}
          filename={filename}
          id={id}
          tags={tags}
          content={content}
          data={data}
          setData={setData}
          setContent={setContent}
          fileIndex={fileIndex}
        />
      ) : (
        <Welcome toggle={toggle} />
      )}
    </>
  );
}
