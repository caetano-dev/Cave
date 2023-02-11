import "./FileBar.css";

function FileBar() {
  return (
    <div className="sidebar">
      <ul className="file-tree">
        <li className="folder">
          <span className="folder-name">Folder 1</span>
          <ul className="folder-contents">
            <li className="file">File 1</li>
            <li className="file">File 2</li>
          </ul>
        </li>
        <li className="folder">
          <span className="folder-name">Folder 2</span>
          <ul className="folder-contents">
            <li className="file">File 3</li>
            <li className="file">File 4</li>
          </ul>
        </li>
      </ul>
    </div>
  );
}

export default FileBar;
