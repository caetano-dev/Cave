import React, { useState } from 'react';
import './File.css';

function File({ toggle, filename, content }) {

    return <><h1 className={"title " + toggle}>{filename}</h1><p>{content}</p></>
}

export default File;

