const deleteFile = (id) => {
  const host = "http://localhost:3000/files/delete";
  fetch(host, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ ID: id }),
  })
    .then((response) => {
      if (response.status != 200) {
        console.error("Failed to delete file: " + response.statusText);
      }
    })
    .catch((error) => {
      console.error(error);
    });
};

export default deleteFile;
