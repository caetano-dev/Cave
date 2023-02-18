import React, { useState } from 'react';
import './SideBar.css';

function SideBar() {
  const [data, setData] = useState([
    {
      name: 'Folder 1',
      type: 'folder',
      children: [
        {
          name: 'File 1',
          type: 'file'
        },
        {
          name: 'File 2',
          type: 'file'
        }
      ]
    },
    {
      name: 'Folder 2',
      type: 'folder',
      children: [
        {
          name: 'File 3',
          type: 'file'
        },
        {
          name: 'File 4',
          type: 'file'
        },
        {
          name: 'Folder 3',
          type: 'folder',
          children: [
            {
              name: 'File 5',
              type: 'file'
            }
          ]
        }
      ]
    },
    {
      name: 'File 6',
      type: 'file'
    }
  ]);

  const handleClick = (node) => {
    if (node.type === 'folder') {
      const updatedNode = {...node, isOpen: !node.isOpen};
      const updatedData = data.map((n) => (n === node ? updatedNode : n));
      setData(updatedData)
    }
  };

  const renderNode = (node, level = 0) => {
    const isFolder = node.type === 'folder';
    const marginLeft = level * 16;

    return (
      <div key={node.name} style={{ marginLeft }} className="node">
        <div className={isFolder ? "folder" + (node.isOpen ? " open" : "") : "file"} onClick={() => handleClick(node)}>{node.name}</div>
        {isFolder && node.children && (
          <div className={"children"}>
            {node.children.map(childNode => renderNode(childNode, level + 1))}
          </div>
        )}
      </div>
    );
  };

  return <div className={"sidebar"}>{data.map(node => renderNode(node))}</div>;
}

export default SideBar;