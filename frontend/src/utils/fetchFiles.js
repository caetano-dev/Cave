async function fetchFiles(setData) {
  //get data from local storage
  const savedData = JSON.parse(localStorage.getItem("data")) || { files: [] };

  if (navigator.onLine) {
    //if it is online, we need to compare with local storage. if local storage is different, we upload it.
    //fetch files from server
    await fetch("http://localhost:3000/files")
      .then((response) => {
        if (response.status === 200) {
          return response.json();
        } else {
          throw new Error("Server error: " + response.status);
        }
      })
      .then((json) => {
        //save to data and local storage
        setData(json.files);
        localStorage.setItem("data", JSON.stringify(json.files)); //similar to data
      })
      .catch((error) => {
        //if there is an error, use local storage
        console.error(error);
        setData(savedData);
      });
  } else { //if it's offline, use local storage
    setData(savedData);
  }
};
export default fetchFiles;
