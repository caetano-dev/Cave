import React from "react";
import "./TagsList.css";

function TagsList({ tags }) {
  return tags.length > 0 ? (
    <div>
      {tags.map((tag) => (
        <React.Fragment key={tag}>
          <span className="tagsList">{`#${tag}`}</span>
        </React.Fragment>
      ))}
    </div>
  ) : null;
}

export default TagsList;
