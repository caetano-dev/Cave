import React, { useState } from 'react';
import './SideBar.css';

function SideBar() {
  const [data, setData] = useState([
    {
      name: 'Foobar',
      type: 'folder',
      children: [
        {
          name: 'foo.txt',
          type: 'file'
        },
        {
          name: 'bar.txt',
          type: 'file'
        }
      ]
    },
    {
      name: 'Lorem Impsum',
      type: 'folder',
      children: [
        {
          name: 'lorem.txt',
          type: 'file'
        },
        {
          name: 'inpsum.txt',
          type: 'file'
        },
        {
          name: 'Latin Class',
          type: 'folder',
          children: [
            {
              name: 'latin.txt',
              type: 'file'
            },

          ]
        }
      ]
    },
    {
      name: 'file.txt',
      type: 'file'
    }
  ]);

  const renderFolder = (folder) => <li className='folder'>{folder.name} <ul>{folder.children.map(renderNode)}</ul></li>
  const renderFile = (file) => <li className='file'>{file.name}</li>

  const renderNode = (content) => <>{content.type == 'folder' ? renderFolder(content) : renderFile(content)}</>

  return <aside className='sidebar close'>Folders<ul>{data.map(renderNode)}</ul></aside>
}

export default SideBar;

