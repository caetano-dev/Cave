import React, { useState } from 'react';

interface ToolbarProps{
    brushSize: number;
    brushColor: string;
    onBrushSizeChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
    onBrushColorChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
}

const Toolbar: React.FC<ToolbarProps> = ({ brushSize, brushColor, onBrushSizeChange, onBrushColorChange }) => (
  <div className="toolbar">
    <label htmlFor="brush-size">Brush size:</label>
    <input
      type="number"
      id="brush-size"
      value={brushSize}
      onChange={onBrushSizeChange}
    />
    <label htmlFor="brush-color">Brush color:</label>
    <input
      type="color"
      id="brush-color"
      value={brushColor}
      onChange={onBrushColorChange}
    />
  </div>
);

const Canvas = () => {
  const [brushSize, setBrushSize] = useState(1);
  const [brushColor, setBrushColor] = useState('black');

  const handleBrushSizeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setBrushSize(Number(event.target.value));
  };

  const handleBrushColorChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setBrushColor(event.target.value);
  };

  const draw = (event: React.MouseEvent<HTMLCanvasElement, MouseEvent>) => {
    const canvas = event.target as HTMLCanvasElement;
    const context = canvas.getContext('2d');

    if(!context) return;

    context.fillStyle = brushColor;
    context.fillRect(event.clientX, event.clientY, brushSize, brushSize);
  };

  return (
    <div>
        <Toolbar brushSize={brushSize} brushColor={brushColor} onBrushSizeChange={handleBrushSizeChange} onBrushColorChange={handleBrushColorChange} />
      <canvas
        width={window.innerWidth}
        height={window.innerHeight}
        onClick={draw}
      />
    </div>
  );
};

export default Canvas;
