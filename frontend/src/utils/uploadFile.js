import fetchFiles from "./fetchFiles";

const createFile = async (data, setData) => {
  const savedData = JSON.parse(localStorage.getItem("data")) || [];
  const newFile = {
    FileInformation: {
      id: savedData.length + 1,
      hash: "",
      type: "file",
      created_at: new Date().toISOString(),
      filename: "New file",
      filepath: "files/New file",
      tags: []
    },
    Content: ""
  };
  savedData.push(newFile)
  localStorage.setItem("data", JSON.stringify(savedData));

  //upload file (can be a new function)
  const host = "http://localhost:3000";
  const url = `${host}/files/create`;
  const requestParamns = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      filename: "New file",
      tags: ""
    }),
  };
  fetch(url, requestParamns).catch(console.error);
  //end upload file

  //TODO: when file is created, we need to update the local storage, specially offline.
  fetchFiles(setData)
};
export default createFile;

