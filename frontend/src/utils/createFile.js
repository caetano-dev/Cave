import fetchFiles from "./fetchFiles";
import uploadFile from "./uploadFile";

const createFile = async (setData) => {
  const savedData = JSON.parse(localStorage.getItem("data")) || [];
  const newFile = {
    FileInformation: {
      id: savedData.length + 1,
      hash: "",
      type: "file",
      created_at: new Date().toISOString(),
      filename: "New_file",
      filepath: "files/New file",
      tags: [],
    },
    Content: "",
  };
  savedData.push(newFile);
  console.log(savedData);
  localStorage.setItem("data", JSON.stringify(savedData));
  uploadFile(newFile);
  fetchFiles(setData);
};
export default createFile;
