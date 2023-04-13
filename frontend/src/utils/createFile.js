import fetchFiles from "./fetchFiles";

const createFile = async (setData) => {
  const host = "http://localhost:3000";
  const url = `${host}/files/create`;
  const requestParamns = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
  };
  fetch(url, requestParamns).catch(console.error);
  console.log("new file created")
  //TODO: when file is created, we need to update the local storage.
  fetchFiles(setData)
};
export default createFile;
