import React from "react";
import fetchFiles from "../../utils/utils"

function CreateFileButton({setData}) {
  const createFile = async () => {
    const host = "http://localhost:3000";
    const url = `${host}/files/create`;
    const requestParamns = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
    };

    fetch(url, requestParamns).catch(console.error);
    fetchFiles(setData)
  };

  return <button onClick={createFile}>Add file</button>;
}

export default CreateFileButton;
