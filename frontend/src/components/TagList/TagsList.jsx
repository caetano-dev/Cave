import React from "react";
import "./TagsList.css";
import PropTypes from 'prop-types'

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

TagsList.propTypes = {
  tags: PropTypes.arrayOf(PropTypes.string)
};
export default TagsList;
