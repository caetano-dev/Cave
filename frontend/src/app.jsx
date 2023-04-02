import React, { useState, useEffect } from "react";
import SideBar from "./components/SideBar/SideBar";
import File from "./components/File/File";
import Welcome from "./components/Welcome/Welcome";

import "./app.css";

export function App() {
  const [toggle, setToggle] = useState("open");
  const toggleState = () =>
    toggle === "open" ? setToggle("close") : setToggle("open");

  const [filename, setFilename] = useState("");
  const [data, setData] = useState([]);
  const [id, setId] = useState(null);
  const [tags, setTags] = useState("");
  const [content, setContent] = useState("");
  const [fileIndex, setFileIndex] = useState("");

  useEffect(() => {
    const savedData = JSON.parse(localStorage.getItem("data"));
    if (!navigator.onLine) {
      setData(savedData.files);
      return;
    }

    fetch("http://localhost:3000/files")
      .then((response) => {
        if (response.status === 200) {
          return response.json();
        } else {
          throw new Error("Server error: " + response.status);
        }
      })
      .then((json) => {
        setData(json.files);
        localStorage.setItem("data", JSON.stringify(json));
      })
      .catch((error) => {
        console.error(error);
        setData(savedData.files);
      });
  }, []);
    

  return (
    <>
      <SideBar
        toggle={toggle}
        toggleState={toggleState}
        data={data}
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
