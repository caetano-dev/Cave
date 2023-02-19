import React, { useState } from 'react';
import './File.css';

function File({ toggle }) {

    return <h1 className={"title " + toggle}>Hello World</h1>
}

export default File;

