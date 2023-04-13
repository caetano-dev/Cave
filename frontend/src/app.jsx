import React, { useState, useEffect } from "react";
import SideBar from "./components/SideBar/SideBar";
import File from "./components/File/File";
import Welcome from "./components/Welcome/Welcome";
import fetchFiles from "./utils/fetchFiles";

import "./app.css";

export function App() {
  const [toggle, setToggle] = useState("open");
  const [filename, setFilename] = useState("");
  const [data, setData] = useState([]);
  const [id, setId] = useState(null);
  const [tags, setTags] = useState("");
  const [content, setContent] = useState("");
  const [fileIndex, setFileIndex] = useState("");

  const toggleState = () =>
    setToggle((toggle) => (toggle === "open" ? "close" : "open"));
  const props = {
    id: id,
    setId: setId,
    toggle: toggle,
    toggleState: toggleState,
    data: data,
    filename: filename,
    setFilename: setFilename,
    setData: setData,
    tags: tags,
    setTags: setTags,
    setContent: setContent,
    content: content,
    fileIndex: fileIndex,
    setFileIndex: setFileIndex,
  };

  useEffect(() => {
    fetchFiles(setData);
    console.log("data")
    console.log(data)
    setData(data)
    //syncFiles(setData);
  }, []);
  console.log(data);

  return (
    <>
      <SideBar {...props} />
      {id ? <File {...props} /> : <Welcome toggle={toggle} />}
    </>
  );
}
