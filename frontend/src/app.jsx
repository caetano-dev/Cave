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

  useEffect(() => {
    const savedData = JSON.parse(localStorage.getItem("data"));
    if (savedData) {
      setData(savedData);
    } else {
      fetch("http://localhost:3000/files")
        .then((response) => response.json())
        .then((json) => {
          setData(json);
          localStorage.setItem("data", JSON.stringify(json));
        });
    }
  }, []);

  const fetchContent = async (id) => {
    try {
      //TODO: make filecontent return the id of the files and get all the contents from all of them at once.
      const response = await fetch("http://localhost:3000/filecontent", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id,
        }),
      });
      if (response.status === 200) {
        const data = await response.json();
        setContent(data.Content);
        console.log(data)
        console.log("data content: " + data.Content);
      } else {
        console.log(`Failed to fetch content for id ${id}`);
      }
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    if (id) {
      console.log("fetching content for id " + id);
      fetchContent(id);
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
        setTags={setTags}
      />
      {id ? (
        <File
          toggle={toggle}
          setFilename={setFilename}
          filename={filename}
          id={id}
          tags={tags}
          content={content}
          setContent={setContent}
        />
      ) : (
        <Welcome toggle={toggle} />
      )}
    </>
  );
}
