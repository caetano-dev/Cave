const uploadFile = (file) => {
  if (navigator.onLine) {
    const host = "http://localhost:3000";
    const url = `${host}/files/create`;
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(file),
    };
    fetch(url, options)
      .then((response) => {
        if (response.ok) {
          console.log("upload successful!");
        } else {
          console.log("upload failed!");
        }
      })
      .catch((error) => {
        console.error(error);
      });
  }
};

export default uploadFile;
