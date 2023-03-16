import React, { useState, useEffect } from "react";
import SideBar from "./components/SideBar/SideBar";
import File from "./components/File/File";

import "./app.css";

export function App() {
  const [toggle, setToggle] = useState("open"); //open/close
  const toggleState = () =>
    toggle == "open" ? setToggle("close") : setToggle("open");

  const [filename, setFilename] = useState("");
  const [data, setData] = useState([]);
  const [id, setId] = useState(null);
  const [content, setContent] = useState("");

  useEffect(() => {
    fetch("http://localhost:3000/files")
      .then((response) => response.json())
      .then((json) => setData(json));
  }, []);

  const fetchContent = async (id) => {
    try {
      const response = await fetch("http://localhost:3000/filecontent", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id,
        }),
      });
      const data = await response.json();
      setContent(data.Content);
      console.log("data content: "+data.Content)
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    if (id) {
      console.log("fetching content for id " + id);
      fetchContent(id);
      console.log("content app.jsx: "+ content);
    }
  }, [id]);

  console.log(data);

  return (
    <>
      <SideBar
        toggle={toggle}
        toggleState={toggleState}
        data={data}
        setFilename={setFilename}
        setId={setId}
      />
      <File
        toggle={toggle}
        filename={filename}
        id={id}
        content={content}
        setContent={setContent}
      />
    </>
  );
}
